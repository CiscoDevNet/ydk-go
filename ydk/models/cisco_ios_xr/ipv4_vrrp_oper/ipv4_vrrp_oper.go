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
    parent types.Entity
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

func (vrrp *Vrrp) GetFilter() yfilter.YFilter { return vrrp.YFilter }

func (vrrp *Vrrp) SetFilter(yf yfilter.YFilter) { vrrp.YFilter = yf }

func (vrrp *Vrrp) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "mgo-sessions" { return "MgoSessions" }
    return ""
}

func (vrrp *Vrrp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-vrrp-oper:vrrp"
}

func (vrrp *Vrrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &vrrp.Summary
    }
    if childYangName == "ipv6" {
        return &vrrp.Ipv6
    }
    if childYangName == "ipv4" {
        return &vrrp.Ipv4
    }
    if childYangName == "mgo-sessions" {
        return &vrrp.MgoSessions
    }
    return nil
}

func (vrrp *Vrrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &vrrp.Summary
    children["ipv6"] = &vrrp.Ipv6
    children["ipv4"] = &vrrp.Ipv4
    children["mgo-sessions"] = &vrrp.MgoSessions
    return children
}

func (vrrp *Vrrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrrp *Vrrp) GetBundleName() string { return "cisco_ios_xr" }

func (vrrp *Vrrp) GetYangName() string { return "vrrp" }

func (vrrp *Vrrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrrp *Vrrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrrp *Vrrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrrp *Vrrp) SetParent(parent types.Entity) { vrrp.parent = parent }

func (vrrp *Vrrp) GetParent() types.Entity { return vrrp.parent }

func (vrrp *Vrrp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-vrrp-oper" }

// Vrrp_Summary
// VRRP summary statistics
type Vrrp_Summary struct {
    parent types.Entity
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
    Ipv6BfdSessionsUp interface{}

    // Number of VRRP IPv6 BFD sessions in DOWN state. The type is interface{}
    // with range: 0..4294967295.
    Ipv6BfdSessionsDown interface{}

    // Number of VRRP IPv6 BFD sessions in INACTIVE state. The type is interface{}
    // with range: 0..4294967295.
    Ipv6BfdSessionInactive interface{}
}

func (summary *Vrrp_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Vrrp_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Vrrp_Summary) GetGoName(yname string) string {
    if yname == "ipv4-sessions-master-owner" { return "Ipv4SessionsMasterOwner" }
    if yname == "ipv4-sessions-master" { return "Ipv4SessionsMaster" }
    if yname == "ipv4-sessions-backup" { return "Ipv4SessionsBackup" }
    if yname == "ipv4-sessions-init" { return "Ipv4SessionsInit" }
    if yname == "ipv4-slaves-master" { return "Ipv4SlavesMaster" }
    if yname == "ipv4-slaves-backup" { return "Ipv4SlavesBackup" }
    if yname == "ipv4-slaves-init" { return "Ipv4SlavesInit" }
    if yname == "ipv4-virtual-ip-addresses-master-owner-up" { return "Ipv4VirtualIpAddressesMasterOwnerUp" }
    if yname == "ipv4-virtual-ip-addresses-master-owner-down" { return "Ipv4VirtualIpAddressesMasterOwnerDown" }
    if yname == "ipv4-virtual-ip-addresses-master-up" { return "Ipv4VirtualIpAddressesMasterUp" }
    if yname == "ipv4-virtual-ip-addresses-master-down" { return "Ipv4VirtualIpAddressesMasterDown" }
    if yname == "ipv4-virtual-ip-addresses-backup-up" { return "Ipv4VirtualIpAddressesBackupUp" }
    if yname == "ipv4-virtual-ip-addresses-backup-down" { return "Ipv4VirtualIpAddressesBackupDown" }
    if yname == "ipv4-virtual-ip-addresses-init-up" { return "Ipv4VirtualIpAddressesInitUp" }
    if yname == "ipv4-virtual-ip-addresses-init-down" { return "Ipv4VirtualIpAddressesInitDown" }
    if yname == "ipv6-sessions-master-owner" { return "Ipv6SessionsMasterOwner" }
    if yname == "ipv6-sessions-master" { return "Ipv6SessionsMaster" }
    if yname == "ipv6-sessions-backup" { return "Ipv6SessionsBackup" }
    if yname == "ipv6-sessions-init" { return "Ipv6SessionsInit" }
    if yname == "ipv6-slaves-master" { return "Ipv6SlavesMaster" }
    if yname == "ipv6-slaves-backup" { return "Ipv6SlavesBackup" }
    if yname == "ipv6-slaves-init" { return "Ipv6SlavesInit" }
    if yname == "ipv6-virtual-ip-addresses-master-owner-up" { return "Ipv6VirtualIpAddressesMasterOwnerUp" }
    if yname == "ipv6-virtual-ip-addresses-master-owner-down" { return "Ipv6VirtualIpAddressesMasterOwnerDown" }
    if yname == "ipv6-virtual-ip-addresses-master-up" { return "Ipv6VirtualIpAddressesMasterUp" }
    if yname == "ipv6-virtual-ip-addresses-master-down" { return "Ipv6VirtualIpAddressesMasterDown" }
    if yname == "ipv6-virtual-ip-addresses-backup-up" { return "Ipv6VirtualIpAddressesBackupUp" }
    if yname == "ipv6-virtual-ip-addresses-backup-down" { return "Ipv6VirtualIpAddressesBackupDown" }
    if yname == "ipv6-virtual-ip-addresses-init-up" { return "Ipv6VirtualIpAddressesInitUp" }
    if yname == "ipv6-virtual-ip-addresses-init-down" { return "Ipv6VirtualIpAddressesInitDown" }
    if yname == "interfaces-ipv4-state-up" { return "InterfacesIpv4StateUp" }
    if yname == "interfaces-ipv4-state-down" { return "InterfacesIpv4StateDown" }
    if yname == "tracked-interfaces-ipv4-state-up" { return "TrackedInterfacesIpv4StateUp" }
    if yname == "tracked-interfaces-ipv4-state-down" { return "TrackedInterfacesIpv4StateDown" }
    if yname == "interfaces-ipv6-state-up" { return "InterfacesIpv6StateUp" }
    if yname == "interfaces-ipv6-state-down" { return "InterfacesIpv6StateDown" }
    if yname == "tracked-interfaces-ipv6-state-up" { return "TrackedInterfacesIpv6StateUp" }
    if yname == "tracked-interfaces-ipv6-state-down" { return "TrackedInterfacesIpv6StateDown" }
    if yname == "tracked-objects-state-up" { return "TrackedObjectsStateUp" }
    if yname == "tracked-objects-state-down" { return "TrackedObjectsStateDown" }
    if yname == "bfd-sessions-up" { return "BfdSessionsUp" }
    if yname == "bfd-sessions-down" { return "BfdSessionsDown" }
    if yname == "bfd-session-inactive" { return "BfdSessionInactive" }
    if yname == "ipv6bfd-sessions-up" { return "Ipv6BfdSessionsUp" }
    if yname == "ipv6bfd-sessions-down" { return "Ipv6BfdSessionsDown" }
    if yname == "ipv6bfd-session-inactive" { return "Ipv6BfdSessionInactive" }
    return ""
}

func (summary *Vrrp_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Vrrp_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Vrrp_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Vrrp_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-sessions-master-owner"] = summary.Ipv4SessionsMasterOwner
    leafs["ipv4-sessions-master"] = summary.Ipv4SessionsMaster
    leafs["ipv4-sessions-backup"] = summary.Ipv4SessionsBackup
    leafs["ipv4-sessions-init"] = summary.Ipv4SessionsInit
    leafs["ipv4-slaves-master"] = summary.Ipv4SlavesMaster
    leafs["ipv4-slaves-backup"] = summary.Ipv4SlavesBackup
    leafs["ipv4-slaves-init"] = summary.Ipv4SlavesInit
    leafs["ipv4-virtual-ip-addresses-master-owner-up"] = summary.Ipv4VirtualIpAddressesMasterOwnerUp
    leafs["ipv4-virtual-ip-addresses-master-owner-down"] = summary.Ipv4VirtualIpAddressesMasterOwnerDown
    leafs["ipv4-virtual-ip-addresses-master-up"] = summary.Ipv4VirtualIpAddressesMasterUp
    leafs["ipv4-virtual-ip-addresses-master-down"] = summary.Ipv4VirtualIpAddressesMasterDown
    leafs["ipv4-virtual-ip-addresses-backup-up"] = summary.Ipv4VirtualIpAddressesBackupUp
    leafs["ipv4-virtual-ip-addresses-backup-down"] = summary.Ipv4VirtualIpAddressesBackupDown
    leafs["ipv4-virtual-ip-addresses-init-up"] = summary.Ipv4VirtualIpAddressesInitUp
    leafs["ipv4-virtual-ip-addresses-init-down"] = summary.Ipv4VirtualIpAddressesInitDown
    leafs["ipv6-sessions-master-owner"] = summary.Ipv6SessionsMasterOwner
    leafs["ipv6-sessions-master"] = summary.Ipv6SessionsMaster
    leafs["ipv6-sessions-backup"] = summary.Ipv6SessionsBackup
    leafs["ipv6-sessions-init"] = summary.Ipv6SessionsInit
    leafs["ipv6-slaves-master"] = summary.Ipv6SlavesMaster
    leafs["ipv6-slaves-backup"] = summary.Ipv6SlavesBackup
    leafs["ipv6-slaves-init"] = summary.Ipv6SlavesInit
    leafs["ipv6-virtual-ip-addresses-master-owner-up"] = summary.Ipv6VirtualIpAddressesMasterOwnerUp
    leafs["ipv6-virtual-ip-addresses-master-owner-down"] = summary.Ipv6VirtualIpAddressesMasterOwnerDown
    leafs["ipv6-virtual-ip-addresses-master-up"] = summary.Ipv6VirtualIpAddressesMasterUp
    leafs["ipv6-virtual-ip-addresses-master-down"] = summary.Ipv6VirtualIpAddressesMasterDown
    leafs["ipv6-virtual-ip-addresses-backup-up"] = summary.Ipv6VirtualIpAddressesBackupUp
    leafs["ipv6-virtual-ip-addresses-backup-down"] = summary.Ipv6VirtualIpAddressesBackupDown
    leafs["ipv6-virtual-ip-addresses-init-up"] = summary.Ipv6VirtualIpAddressesInitUp
    leafs["ipv6-virtual-ip-addresses-init-down"] = summary.Ipv6VirtualIpAddressesInitDown
    leafs["interfaces-ipv4-state-up"] = summary.InterfacesIpv4StateUp
    leafs["interfaces-ipv4-state-down"] = summary.InterfacesIpv4StateDown
    leafs["tracked-interfaces-ipv4-state-up"] = summary.TrackedInterfacesIpv4StateUp
    leafs["tracked-interfaces-ipv4-state-down"] = summary.TrackedInterfacesIpv4StateDown
    leafs["interfaces-ipv6-state-up"] = summary.InterfacesIpv6StateUp
    leafs["interfaces-ipv6-state-down"] = summary.InterfacesIpv6StateDown
    leafs["tracked-interfaces-ipv6-state-up"] = summary.TrackedInterfacesIpv6StateUp
    leafs["tracked-interfaces-ipv6-state-down"] = summary.TrackedInterfacesIpv6StateDown
    leafs["tracked-objects-state-up"] = summary.TrackedObjectsStateUp
    leafs["tracked-objects-state-down"] = summary.TrackedObjectsStateDown
    leafs["bfd-sessions-up"] = summary.BfdSessionsUp
    leafs["bfd-sessions-down"] = summary.BfdSessionsDown
    leafs["bfd-session-inactive"] = summary.BfdSessionInactive
    leafs["ipv6bfd-sessions-up"] = summary.Ipv6BfdSessionsUp
    leafs["ipv6bfd-sessions-down"] = summary.Ipv6BfdSessionsDown
    leafs["ipv6bfd-session-inactive"] = summary.Ipv6BfdSessionInactive
    return leafs
}

func (summary *Vrrp_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Vrrp_Summary) GetYangName() string { return "summary" }

func (summary *Vrrp_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Vrrp_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Vrrp_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Vrrp_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Vrrp_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Vrrp_Summary) GetParentYangName() string { return "vrrp" }

// Vrrp_Ipv6
// IPv6 VRRP configuration
type Vrrp_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP tracked item table.
    TrackItems Vrrp_Ipv6_TrackItems

    // The VRRP virtual router table.
    VirtualRouters Vrrp_Ipv6_VirtualRouters

    // The VRRP interface table.
    Interfaces Vrrp_Ipv6_Interfaces
}

func (ipv6 *Vrrp_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Vrrp_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Vrrp_Ipv6) GetGoName(yname string) string {
    if yname == "track-items" { return "TrackItems" }
    if yname == "virtual-routers" { return "VirtualRouters" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv6 *Vrrp_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Vrrp_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-items" {
        return &ipv6.TrackItems
    }
    if childYangName == "virtual-routers" {
        return &ipv6.VirtualRouters
    }
    if childYangName == "interfaces" {
        return &ipv6.Interfaces
    }
    return nil
}

func (ipv6 *Vrrp_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-items"] = &ipv6.TrackItems
    children["virtual-routers"] = &ipv6.VirtualRouters
    children["interfaces"] = &ipv6.Interfaces
    return children
}

func (ipv6 *Vrrp_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Vrrp_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Vrrp_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Vrrp_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Vrrp_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Vrrp_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Vrrp_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Vrrp_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Vrrp_Ipv6) GetParentYangName() string { return "vrrp" }

// Vrrp_Ipv6_TrackItems
// The VRRP tracked item table
type Vrrp_Ipv6_TrackItems struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A configured VRRP IP address entry. The type is slice of
    // Vrrp_Ipv6_TrackItems_TrackItem.
    TrackItem []Vrrp_Ipv6_TrackItems_TrackItem
}

func (trackItems *Vrrp_Ipv6_TrackItems) GetFilter() yfilter.YFilter { return trackItems.YFilter }

func (trackItems *Vrrp_Ipv6_TrackItems) SetFilter(yf yfilter.YFilter) { trackItems.YFilter = yf }

func (trackItems *Vrrp_Ipv6_TrackItems) GetGoName(yname string) string {
    if yname == "track-item" { return "TrackItem" }
    return ""
}

func (trackItems *Vrrp_Ipv6_TrackItems) GetSegmentPath() string {
    return "track-items"
}

func (trackItems *Vrrp_Ipv6_TrackItems) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-item" {
        for _, c := range trackItems.TrackItem {
            if trackItems.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv6_TrackItems_TrackItem{}
        trackItems.TrackItem = append(trackItems.TrackItem, child)
        return &trackItems.TrackItem[len(trackItems.TrackItem)-1]
    }
    return nil
}

func (trackItems *Vrrp_Ipv6_TrackItems) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackItems.TrackItem {
        children[trackItems.TrackItem[i].GetSegmentPath()] = &trackItems.TrackItem[i]
    }
    return children
}

func (trackItems *Vrrp_Ipv6_TrackItems) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackItems *Vrrp_Ipv6_TrackItems) GetBundleName() string { return "cisco_ios_xr" }

func (trackItems *Vrrp_Ipv6_TrackItems) GetYangName() string { return "track-items" }

func (trackItems *Vrrp_Ipv6_TrackItems) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackItems *Vrrp_Ipv6_TrackItems) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackItems *Vrrp_Ipv6_TrackItems) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackItems *Vrrp_Ipv6_TrackItems) SetParent(parent types.Entity) { trackItems.parent = parent }

func (trackItems *Vrrp_Ipv6_TrackItems) GetParent() types.Entity { return trackItems.parent }

func (trackItems *Vrrp_Ipv6_TrackItems) GetParentYangName() string { return "ipv6" }

// Vrrp_Ipv6_TrackItems_TrackItem
// A configured VRRP IP address entry
type Vrrp_Ipv6_TrackItems_TrackItem struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name to track. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The VRRP virtual router id. The type is
    // interface{} with range: -2147483648..2147483647.
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

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetFilter() yfilter.YFilter { return trackItem.YFilter }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) SetFilter(yf yfilter.YFilter) { trackItem.YFilter = yf }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "tracked-interface-name" { return "TrackedInterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "virtual-router-id-xr" { return "VirtualRouterIdXr" }
    if yname == "tracked-item-type" { return "TrackedItemType" }
    if yname == "tracked-item-index" { return "TrackedItemIndex" }
    if yname == "state" { return "State" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetSegmentPath() string {
    return "track-item" + "[interface-name='" + fmt.Sprintf("%v", trackItem.InterfaceName) + "']" + "[virtual-router-id='" + fmt.Sprintf("%v", trackItem.VirtualRouterId) + "']" + "[tracked-interface-name='" + fmt.Sprintf("%v", trackItem.TrackedInterfaceName) + "']"
}

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = trackItem.InterfaceName
    leafs["virtual-router-id"] = trackItem.VirtualRouterId
    leafs["tracked-interface-name"] = trackItem.TrackedInterfaceName
    leafs["interface"] = trackItem.Interface
    leafs["virtual-router-id-xr"] = trackItem.VirtualRouterIdXr
    leafs["tracked-item-type"] = trackItem.TrackedItemType
    leafs["tracked-item-index"] = trackItem.TrackedItemIndex
    leafs["state"] = trackItem.State
    leafs["priority"] = trackItem.Priority
    return leafs
}

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetBundleName() string { return "cisco_ios_xr" }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetYangName() string { return "track-item" }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) SetParent(parent types.Entity) { trackItem.parent = parent }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetParent() types.Entity { return trackItem.parent }

func (trackItem *Vrrp_Ipv6_TrackItems_TrackItem) GetParentYangName() string { return "track-items" }

// Vrrp_Ipv6_VirtualRouters
// The VRRP virtual router table
type Vrrp_Ipv6_VirtualRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP virtual router. The type is slice of
    // Vrrp_Ipv6_VirtualRouters_VirtualRouter.
    VirtualRouter []Vrrp_Ipv6_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetFilter() yfilter.YFilter { return virtualRouters.YFilter }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) SetFilter(yf yfilter.YFilter) { virtualRouters.YFilter = yf }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetGoName(yname string) string {
    if yname == "virtual-router" { return "VirtualRouter" }
    return ""
}

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetSegmentPath() string {
    return "virtual-routers"
}

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-router" {
        for _, c := range virtualRouters.VirtualRouter {
            if virtualRouters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv6_VirtualRouters_VirtualRouter{}
        virtualRouters.VirtualRouter = append(virtualRouters.VirtualRouter, child)
        return &virtualRouters.VirtualRouter[len(virtualRouters.VirtualRouter)-1]
    }
    return nil
}

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range virtualRouters.VirtualRouter {
        children[virtualRouters.VirtualRouter[i].GetSegmentPath()] = &virtualRouters.VirtualRouter[i]
    }
    return children
}

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetYangName() string { return "virtual-routers" }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) SetParent(parent types.Entity) { virtualRouters.parent = parent }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetParent() types.Entity { return virtualRouters.parent }

func (virtualRouters *Vrrp_Ipv6_VirtualRouters) GetParentYangName() string { return "ipv6" }

// Vrrp_Ipv6_VirtualRouters_VirtualRouter
// A VRRP virtual router
type Vrrp_Ipv6_VirtualRouters_VirtualRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The VRRP virtual router id. The type is
    // interface{} with range: -2147483648..2147483647.
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
    Ipv6OperationalAddress []Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress

    // IPv6 Configured but Down VRRP addresses. The type is slice of
    // Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress.
    Ipv6ConfiguredDownAddress []Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress

    // State change history. The type is slice of
    // Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory.
    StateChangeHistory []Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory
}

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetFilter() yfilter.YFilter { return virtualRouter.YFilter }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) SetFilter(yf yfilter.YFilter) { virtualRouter.YFilter = yf }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "virtual-router-id-xr" { return "VirtualRouterIdXr" }
    if yname == "version" { return "Version" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "session-name" { return "SessionName" }
    if yname == "slaves" { return "Slaves" }
    if yname == "is-slave" { return "IsSlave" }
    if yname == "followed-session-name" { return "FollowedSessionName" }
    if yname == "secondary-address-count" { return "SecondaryAddressCount" }
    if yname == "operational-address-count" { return "OperationalAddressCount" }
    if yname == "primary-virtual-ip" { return "PrimaryVirtualIp" }
    if yname == "configured-down-address-count" { return "ConfiguredDownAddressCount" }
    if yname == "virtual-linklocal-ipv6-address" { return "VirtualLinklocalIpv6Address" }
    if yname == "primary-state" { return "PrimaryState" }
    if yname == "master-ip-address" { return "MasterIpAddress" }
    if yname == "master-ipv6-address" { return "MasterIpv6Address" }
    if yname == "master-priority" { return "MasterPriority" }
    if yname == "vrrp-state" { return "VrrpState" }
    if yname == "authentication-type" { return "AuthenticationType" }
    if yname == "authentication-string" { return "AuthenticationString" }
    if yname == "configured-advertize-time" { return "ConfiguredAdvertizeTime" }
    if yname == "oper-advertize-time" { return "OperAdvertizeTime" }
    if yname == "min-delay-time" { return "MinDelayTime" }
    if yname == "reload-delay-time" { return "ReloadDelayTime" }
    if yname == "delay-timer-flag" { return "DelayTimerFlag" }
    if yname == "delay-timer-secs" { return "DelayTimerSecs" }
    if yname == "delay-timer-msecs" { return "DelayTimerMsecs" }
    if yname == "authentication-flag" { return "AuthenticationFlag" }
    if yname == "force-timer-flag" { return "ForceTimerFlag" }
    if yname == "preempt-flag" { return "PreemptFlag" }
    if yname == "ip-address-owner-flag" { return "IpAddressOwnerFlag" }
    if yname == "is-accept-mode" { return "IsAcceptMode" }
    if yname == "preempt-delay-time" { return "PreemptDelayTime" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "operational-priority" { return "OperationalPriority" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    if yname == "tracked-interface-count" { return "TrackedInterfaceCount" }
    if yname == "tracked-interface-up-count" { return "TrackedInterfaceUpCount" }
    if yname == "tracked-item-count" { return "TrackedItemCount" }
    if yname == "tracked-item-up-count" { return "TrackedItemUpCount" }
    if yname == "time-in-current-state" { return "TimeInCurrentState" }
    if yname == "state-change-count" { return "StateChangeCount" }
    if yname == "time-vrouter-up" { return "TimeVrouterUp" }
    if yname == "master-count" { return "MasterCount" }
    if yname == "adverts-received-count" { return "AdvertsReceivedCount" }
    if yname == "advert-interval-error-count" { return "AdvertIntervalErrorCount" }
    if yname == "adverts-sent-count" { return "AdvertsSentCount" }
    if yname == "authentication-fail-count" { return "AuthenticationFailCount" }
    if yname == "ttl-error-count" { return "TtlErrorCount" }
    if yname == "priority-zero-received-count" { return "PriorityZeroReceivedCount" }
    if yname == "priority-zero-sent-count" { return "PriorityZeroSentCount" }
    if yname == "invalid-packet-count" { return "InvalidPacketCount" }
    if yname == "address-list-error-count" { return "AddressListErrorCount" }
    if yname == "invalid-auth-type-count" { return "InvalidAuthTypeCount" }
    if yname == "auth-type-mismatch-count" { return "AuthTypeMismatchCount" }
    if yname == "pkt-length-errors-count" { return "PktLengthErrorsCount" }
    if yname == "time-stats-discontinuity" { return "TimeStatsDiscontinuity" }
    if yname == "bfd-session-state" { return "BfdSessionState" }
    if yname == "bfd-interval" { return "BfdInterval" }
    if yname == "bfd-multiplier" { return "BfdMultiplier" }
    if yname == "bfd-cfg-remote-ip" { return "BfdCfgRemoteIp" }
    if yname == "bfd-configured-remote-ipv6-address" { return "BfdConfiguredRemoteIpv6Address" }
    if yname == "state-from-checkpoint" { return "StateFromCheckpoint" }
    if yname == "interface-ipv4-address" { return "InterfaceIpv4Address" }
    if yname == "interface-ipv6-address" { return "InterfaceIpv6Address" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "virtual-mac-address-state" { return "VirtualMacAddressState" }
    if yname == "operational-address" { return "OperationalAddress" }
    if yname == "ipv4-configured-down-address" { return "Ipv4ConfiguredDownAddress" }
    if yname == "resign-sent-time" { return "ResignSentTime" }
    if yname == "resign-received-time" { return "ResignReceivedTime" }
    if yname == "ipv6-operational-address" { return "Ipv6OperationalAddress" }
    if yname == "ipv6-configured-down-address" { return "Ipv6ConfiguredDownAddress" }
    if yname == "state-change-history" { return "StateChangeHistory" }
    return ""
}

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetSegmentPath() string {
    return "virtual-router" + "[interface-name='" + fmt.Sprintf("%v", virtualRouter.InterfaceName) + "']" + "[virtual-router-id='" + fmt.Sprintf("%v", virtualRouter.VirtualRouterId) + "']"
}

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "resign-sent-time" {
        return &virtualRouter.ResignSentTime
    }
    if childYangName == "resign-received-time" {
        return &virtualRouter.ResignReceivedTime
    }
    if childYangName == "ipv6-operational-address" {
        for _, c := range virtualRouter.Ipv6OperationalAddress {
            if virtualRouter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress{}
        virtualRouter.Ipv6OperationalAddress = append(virtualRouter.Ipv6OperationalAddress, child)
        return &virtualRouter.Ipv6OperationalAddress[len(virtualRouter.Ipv6OperationalAddress)-1]
    }
    if childYangName == "ipv6-configured-down-address" {
        for _, c := range virtualRouter.Ipv6ConfiguredDownAddress {
            if virtualRouter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress{}
        virtualRouter.Ipv6ConfiguredDownAddress = append(virtualRouter.Ipv6ConfiguredDownAddress, child)
        return &virtualRouter.Ipv6ConfiguredDownAddress[len(virtualRouter.Ipv6ConfiguredDownAddress)-1]
    }
    if childYangName == "state-change-history" {
        for _, c := range virtualRouter.StateChangeHistory {
            if virtualRouter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory{}
        virtualRouter.StateChangeHistory = append(virtualRouter.StateChangeHistory, child)
        return &virtualRouter.StateChangeHistory[len(virtualRouter.StateChangeHistory)-1]
    }
    return nil
}

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["resign-sent-time"] = &virtualRouter.ResignSentTime
    children["resign-received-time"] = &virtualRouter.ResignReceivedTime
    for i := range virtualRouter.Ipv6OperationalAddress {
        children[virtualRouter.Ipv6OperationalAddress[i].GetSegmentPath()] = &virtualRouter.Ipv6OperationalAddress[i]
    }
    for i := range virtualRouter.Ipv6ConfiguredDownAddress {
        children[virtualRouter.Ipv6ConfiguredDownAddress[i].GetSegmentPath()] = &virtualRouter.Ipv6ConfiguredDownAddress[i]
    }
    for i := range virtualRouter.StateChangeHistory {
        children[virtualRouter.StateChangeHistory[i].GetSegmentPath()] = &virtualRouter.StateChangeHistory[i]
    }
    return children
}

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = virtualRouter.InterfaceName
    leafs["virtual-router-id"] = virtualRouter.VirtualRouterId
    leafs["interface-name-xr"] = virtualRouter.InterfaceNameXr
    leafs["virtual-router-id-xr"] = virtualRouter.VirtualRouterIdXr
    leafs["version"] = virtualRouter.Version
    leafs["address-family"] = virtualRouter.AddressFamily
    leafs["session-name"] = virtualRouter.SessionName
    leafs["slaves"] = virtualRouter.Slaves
    leafs["is-slave"] = virtualRouter.IsSlave
    leafs["followed-session-name"] = virtualRouter.FollowedSessionName
    leafs["secondary-address-count"] = virtualRouter.SecondaryAddressCount
    leafs["operational-address-count"] = virtualRouter.OperationalAddressCount
    leafs["primary-virtual-ip"] = virtualRouter.PrimaryVirtualIp
    leafs["configured-down-address-count"] = virtualRouter.ConfiguredDownAddressCount
    leafs["virtual-linklocal-ipv6-address"] = virtualRouter.VirtualLinklocalIpv6Address
    leafs["primary-state"] = virtualRouter.PrimaryState
    leafs["master-ip-address"] = virtualRouter.MasterIpAddress
    leafs["master-ipv6-address"] = virtualRouter.MasterIpv6Address
    leafs["master-priority"] = virtualRouter.MasterPriority
    leafs["vrrp-state"] = virtualRouter.VrrpState
    leafs["authentication-type"] = virtualRouter.AuthenticationType
    leafs["authentication-string"] = virtualRouter.AuthenticationString
    leafs["configured-advertize-time"] = virtualRouter.ConfiguredAdvertizeTime
    leafs["oper-advertize-time"] = virtualRouter.OperAdvertizeTime
    leafs["min-delay-time"] = virtualRouter.MinDelayTime
    leafs["reload-delay-time"] = virtualRouter.ReloadDelayTime
    leafs["delay-timer-flag"] = virtualRouter.DelayTimerFlag
    leafs["delay-timer-secs"] = virtualRouter.DelayTimerSecs
    leafs["delay-timer-msecs"] = virtualRouter.DelayTimerMsecs
    leafs["authentication-flag"] = virtualRouter.AuthenticationFlag
    leafs["force-timer-flag"] = virtualRouter.ForceTimerFlag
    leafs["preempt-flag"] = virtualRouter.PreemptFlag
    leafs["ip-address-owner-flag"] = virtualRouter.IpAddressOwnerFlag
    leafs["is-accept-mode"] = virtualRouter.IsAcceptMode
    leafs["preempt-delay-time"] = virtualRouter.PreemptDelayTime
    leafs["configured-priority"] = virtualRouter.ConfiguredPriority
    leafs["operational-priority"] = virtualRouter.OperationalPriority
    leafs["priority-decrement"] = virtualRouter.PriorityDecrement
    leafs["tracked-interface-count"] = virtualRouter.TrackedInterfaceCount
    leafs["tracked-interface-up-count"] = virtualRouter.TrackedInterfaceUpCount
    leafs["tracked-item-count"] = virtualRouter.TrackedItemCount
    leafs["tracked-item-up-count"] = virtualRouter.TrackedItemUpCount
    leafs["time-in-current-state"] = virtualRouter.TimeInCurrentState
    leafs["state-change-count"] = virtualRouter.StateChangeCount
    leafs["time-vrouter-up"] = virtualRouter.TimeVrouterUp
    leafs["master-count"] = virtualRouter.MasterCount
    leafs["adverts-received-count"] = virtualRouter.AdvertsReceivedCount
    leafs["advert-interval-error-count"] = virtualRouter.AdvertIntervalErrorCount
    leafs["adverts-sent-count"] = virtualRouter.AdvertsSentCount
    leafs["authentication-fail-count"] = virtualRouter.AuthenticationFailCount
    leafs["ttl-error-count"] = virtualRouter.TtlErrorCount
    leafs["priority-zero-received-count"] = virtualRouter.PriorityZeroReceivedCount
    leafs["priority-zero-sent-count"] = virtualRouter.PriorityZeroSentCount
    leafs["invalid-packet-count"] = virtualRouter.InvalidPacketCount
    leafs["address-list-error-count"] = virtualRouter.AddressListErrorCount
    leafs["invalid-auth-type-count"] = virtualRouter.InvalidAuthTypeCount
    leafs["auth-type-mismatch-count"] = virtualRouter.AuthTypeMismatchCount
    leafs["pkt-length-errors-count"] = virtualRouter.PktLengthErrorsCount
    leafs["time-stats-discontinuity"] = virtualRouter.TimeStatsDiscontinuity
    leafs["bfd-session-state"] = virtualRouter.BfdSessionState
    leafs["bfd-interval"] = virtualRouter.BfdInterval
    leafs["bfd-multiplier"] = virtualRouter.BfdMultiplier
    leafs["bfd-cfg-remote-ip"] = virtualRouter.BfdCfgRemoteIp
    leafs["bfd-configured-remote-ipv6-address"] = virtualRouter.BfdConfiguredRemoteIpv6Address
    leafs["state-from-checkpoint"] = virtualRouter.StateFromCheckpoint
    leafs["interface-ipv4-address"] = virtualRouter.InterfaceIpv4Address
    leafs["interface-ipv6-address"] = virtualRouter.InterfaceIpv6Address
    leafs["virtual-mac-address"] = virtualRouter.VirtualMacAddress
    leafs["virtual-mac-address-state"] = virtualRouter.VirtualMacAddressState
    leafs["operational-address"] = virtualRouter.OperationalAddress
    leafs["ipv4-configured-down-address"] = virtualRouter.Ipv4ConfiguredDownAddress
    return leafs
}

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetYangName() string { return "virtual-router" }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) SetParent(parent types.Entity) { virtualRouter.parent = parent }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetParent() types.Entity { return virtualRouter.parent }

func (virtualRouter *Vrrp_Ipv6_VirtualRouters_VirtualRouter) GetParentYangName() string { return "virtual-routers" }

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime
// Time last resign was sent
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetFilter() yfilter.YFilter { return resignSentTime.YFilter }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) SetFilter(yf yfilter.YFilter) { resignSentTime.YFilter = yf }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetSegmentPath() string {
    return "resign-sent-time"
}

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resignSentTime.Seconds
    leafs["nanoseconds"] = resignSentTime.Nanoseconds
    return leafs
}

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetBundleName() string { return "cisco_ios_xr" }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetYangName() string { return "resign-sent-time" }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) SetParent(parent types.Entity) { resignSentTime.parent = parent }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetParent() types.Entity { return resignSentTime.parent }

func (resignSentTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignSentTime) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime
// Time last resign was received
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetFilter() yfilter.YFilter { return resignReceivedTime.YFilter }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) SetFilter(yf yfilter.YFilter) { resignReceivedTime.YFilter = yf }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetSegmentPath() string {
    return "resign-received-time"
}

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resignReceivedTime.Seconds
    leafs["nanoseconds"] = resignReceivedTime.Nanoseconds
    return leafs
}

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetBundleName() string { return "cisco_ios_xr" }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetYangName() string { return "resign-received-time" }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) SetParent(parent types.Entity) { resignReceivedTime.parent = parent }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetParent() types.Entity { return resignReceivedTime.parent }

func (resignReceivedTime *Vrrp_Ipv6_VirtualRouters_VirtualRouter_ResignReceivedTime) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress
// IPv6 Operational VRRP addresses
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetFilter() yfilter.YFilter { return ipv6OperationalAddress.YFilter }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) SetFilter(yf yfilter.YFilter) { ipv6OperationalAddress.YFilter = yf }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetSegmentPath() string {
    return "ipv6-operational-address"
}

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6OperationalAddress.Ipv6Address
    return leafs
}

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetYangName() string { return "ipv6-operational-address" }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) SetParent(parent types.Entity) { ipv6OperationalAddress.parent = parent }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetParent() types.Entity { return ipv6OperationalAddress.parent }

func (ipv6OperationalAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress
// IPv6 Configured but Down VRRP addresses
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetFilter() yfilter.YFilter { return ipv6ConfiguredDownAddress.YFilter }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) SetFilter(yf yfilter.YFilter) { ipv6ConfiguredDownAddress.YFilter = yf }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetSegmentPath() string {
    return "ipv6-configured-down-address"
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6ConfiguredDownAddress.Ipv6Address
    return leafs
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetYangName() string { return "ipv6-configured-down-address" }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) SetParent(parent types.Entity) { ipv6ConfiguredDownAddress.parent = parent }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetParent() types.Entity { return ipv6ConfiguredDownAddress.parent }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv6_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory
// State change history
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory struct {
    parent types.Entity
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

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetFilter() yfilter.YFilter { return stateChangeHistory.YFilter }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) SetFilter(yf yfilter.YFilter) { stateChangeHistory.YFilter = yf }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetGoName(yname string) string {
    if yname == "old-state" { return "OldState" }
    if yname == "new-state" { return "NewState" }
    if yname == "reason" { return "Reason" }
    if yname == "time" { return "Time" }
    return ""
}

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetSegmentPath() string {
    return "state-change-history"
}

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "time" {
        return &stateChangeHistory.Time
    }
    return nil
}

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["time"] = &stateChangeHistory.Time
    return children
}

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["old-state"] = stateChangeHistory.OldState
    leafs["new-state"] = stateChangeHistory.NewState
    leafs["reason"] = stateChangeHistory.Reason
    return leafs
}

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetBundleName() string { return "cisco_ios_xr" }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetYangName() string { return "state-change-history" }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) SetParent(parent types.Entity) { stateChangeHistory.parent = parent }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetParent() types.Entity { return stateChangeHistory.parent }

func (stateChangeHistory *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time
// Time of state change
type Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetFilter() yfilter.YFilter { return time.YFilter }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) SetFilter(yf yfilter.YFilter) { time.YFilter = yf }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetSegmentPath() string {
    return "time"
}

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = time.Seconds
    leafs["nanoseconds"] = time.Nanoseconds
    return leafs
}

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetBundleName() string { return "cisco_ios_xr" }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetYangName() string { return "time" }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) SetParent(parent types.Entity) { time.parent = parent }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetParent() types.Entity { return time.parent }

func (time *Vrrp_Ipv6_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetParentYangName() string { return "state-change-history" }

// Vrrp_Ipv6_Interfaces
// The VRRP interface table
type Vrrp_Ipv6_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP interface entry. The type is slice of
    // Vrrp_Ipv6_Interfaces_Interface.
    Interface []Vrrp_Ipv6_Interfaces_Interface
}

func (interfaces *Vrrp_Ipv6_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Vrrp_Ipv6_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Vrrp_Ipv6_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Vrrp_Ipv6_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Vrrp_Ipv6_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv6_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Vrrp_Ipv6_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Vrrp_Ipv6_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Vrrp_Ipv6_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Vrrp_Ipv6_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Vrrp_Ipv6_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Vrrp_Ipv6_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Vrrp_Ipv6_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Vrrp_Ipv6_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Vrrp_Ipv6_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Vrrp_Ipv6_Interfaces) GetParentYangName() string { return "ipv6" }

// Vrrp_Ipv6_Interfaces_Interface
// A VRRP interface entry
type Vrrp_Ipv6_Interfaces_Interface struct {
    parent types.Entity
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

func (self *Vrrp_Ipv6_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Vrrp_Ipv6_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Vrrp_Ipv6_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "invalid-checksum-count" { return "InvalidChecksumCount" }
    if yname == "invalid-version-count" { return "InvalidVersionCount" }
    if yname == "invalid-vrid-count" { return "InvalidVridCount" }
    if yname == "invalid-packet-length-count" { return "InvalidPacketLengthCount" }
    return ""
}

func (self *Vrrp_Ipv6_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Vrrp_Ipv6_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Vrrp_Ipv6_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Vrrp_Ipv6_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["invalid-checksum-count"] = self.InvalidChecksumCount
    leafs["invalid-version-count"] = self.InvalidVersionCount
    leafs["invalid-vrid-count"] = self.InvalidVridCount
    leafs["invalid-packet-length-count"] = self.InvalidPacketLengthCount
    return leafs
}

func (self *Vrrp_Ipv6_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Vrrp_Ipv6_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Vrrp_Ipv6_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Vrrp_Ipv6_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Vrrp_Ipv6_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Vrrp_Ipv6_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Vrrp_Ipv6_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Vrrp_Ipv6_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Vrrp_Ipv4
// IPv4 VRRP configuration
type Vrrp_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP interface table.
    Interfaces Vrrp_Ipv4_Interfaces

    // The VRRP tracked item table.
    TrackItems Vrrp_Ipv4_TrackItems

    // The VRRP virtual router table.
    VirtualRouters Vrrp_Ipv4_VirtualRouters
}

func (ipv4 *Vrrp_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Vrrp_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Vrrp_Ipv4) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "track-items" { return "TrackItems" }
    if yname == "virtual-routers" { return "VirtualRouters" }
    return ""
}

func (ipv4 *Vrrp_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Vrrp_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &ipv4.Interfaces
    }
    if childYangName == "track-items" {
        return &ipv4.TrackItems
    }
    if childYangName == "virtual-routers" {
        return &ipv4.VirtualRouters
    }
    return nil
}

func (ipv4 *Vrrp_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &ipv4.Interfaces
    children["track-items"] = &ipv4.TrackItems
    children["virtual-routers"] = &ipv4.VirtualRouters
    return children
}

func (ipv4 *Vrrp_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Vrrp_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Vrrp_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Vrrp_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Vrrp_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Vrrp_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Vrrp_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Vrrp_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Vrrp_Ipv4) GetParentYangName() string { return "vrrp" }

// Vrrp_Ipv4_Interfaces
// The VRRP interface table
type Vrrp_Ipv4_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP interface entry. The type is slice of
    // Vrrp_Ipv4_Interfaces_Interface.
    Interface []Vrrp_Ipv4_Interfaces_Interface
}

func (interfaces *Vrrp_Ipv4_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Vrrp_Ipv4_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Vrrp_Ipv4_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Vrrp_Ipv4_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Vrrp_Ipv4_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv4_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Vrrp_Ipv4_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Vrrp_Ipv4_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Vrrp_Ipv4_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Vrrp_Ipv4_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Vrrp_Ipv4_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Vrrp_Ipv4_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Vrrp_Ipv4_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Vrrp_Ipv4_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Vrrp_Ipv4_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Vrrp_Ipv4_Interfaces) GetParentYangName() string { return "ipv4" }

// Vrrp_Ipv4_Interfaces_Interface
// A VRRP interface entry
type Vrrp_Ipv4_Interfaces_Interface struct {
    parent types.Entity
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

func (self *Vrrp_Ipv4_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Vrrp_Ipv4_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Vrrp_Ipv4_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "invalid-checksum-count" { return "InvalidChecksumCount" }
    if yname == "invalid-version-count" { return "InvalidVersionCount" }
    if yname == "invalid-vrid-count" { return "InvalidVridCount" }
    if yname == "invalid-packet-length-count" { return "InvalidPacketLengthCount" }
    return ""
}

func (self *Vrrp_Ipv4_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Vrrp_Ipv4_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Vrrp_Ipv4_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Vrrp_Ipv4_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["invalid-checksum-count"] = self.InvalidChecksumCount
    leafs["invalid-version-count"] = self.InvalidVersionCount
    leafs["invalid-vrid-count"] = self.InvalidVridCount
    leafs["invalid-packet-length-count"] = self.InvalidPacketLengthCount
    return leafs
}

func (self *Vrrp_Ipv4_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Vrrp_Ipv4_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Vrrp_Ipv4_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Vrrp_Ipv4_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Vrrp_Ipv4_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Vrrp_Ipv4_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Vrrp_Ipv4_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Vrrp_Ipv4_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Vrrp_Ipv4_TrackItems
// The VRRP tracked item table
type Vrrp_Ipv4_TrackItems struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A configured VRRP IP address entry. The type is slice of
    // Vrrp_Ipv4_TrackItems_TrackItem.
    TrackItem []Vrrp_Ipv4_TrackItems_TrackItem
}

func (trackItems *Vrrp_Ipv4_TrackItems) GetFilter() yfilter.YFilter { return trackItems.YFilter }

func (trackItems *Vrrp_Ipv4_TrackItems) SetFilter(yf yfilter.YFilter) { trackItems.YFilter = yf }

func (trackItems *Vrrp_Ipv4_TrackItems) GetGoName(yname string) string {
    if yname == "track-item" { return "TrackItem" }
    return ""
}

func (trackItems *Vrrp_Ipv4_TrackItems) GetSegmentPath() string {
    return "track-items"
}

func (trackItems *Vrrp_Ipv4_TrackItems) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-item" {
        for _, c := range trackItems.TrackItem {
            if trackItems.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv4_TrackItems_TrackItem{}
        trackItems.TrackItem = append(trackItems.TrackItem, child)
        return &trackItems.TrackItem[len(trackItems.TrackItem)-1]
    }
    return nil
}

func (trackItems *Vrrp_Ipv4_TrackItems) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackItems.TrackItem {
        children[trackItems.TrackItem[i].GetSegmentPath()] = &trackItems.TrackItem[i]
    }
    return children
}

func (trackItems *Vrrp_Ipv4_TrackItems) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackItems *Vrrp_Ipv4_TrackItems) GetBundleName() string { return "cisco_ios_xr" }

func (trackItems *Vrrp_Ipv4_TrackItems) GetYangName() string { return "track-items" }

func (trackItems *Vrrp_Ipv4_TrackItems) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackItems *Vrrp_Ipv4_TrackItems) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackItems *Vrrp_Ipv4_TrackItems) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackItems *Vrrp_Ipv4_TrackItems) SetParent(parent types.Entity) { trackItems.parent = parent }

func (trackItems *Vrrp_Ipv4_TrackItems) GetParent() types.Entity { return trackItems.parent }

func (trackItems *Vrrp_Ipv4_TrackItems) GetParentYangName() string { return "ipv4" }

// Vrrp_Ipv4_TrackItems_TrackItem
// A configured VRRP IP address entry
type Vrrp_Ipv4_TrackItems_TrackItem struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name to track. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The VRRP virtual router id. The type is
    // interface{} with range: -2147483648..2147483647.
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

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetFilter() yfilter.YFilter { return trackItem.YFilter }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) SetFilter(yf yfilter.YFilter) { trackItem.YFilter = yf }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "tracked-interface-name" { return "TrackedInterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "virtual-router-id-xr" { return "VirtualRouterIdXr" }
    if yname == "tracked-item-type" { return "TrackedItemType" }
    if yname == "tracked-item-index" { return "TrackedItemIndex" }
    if yname == "state" { return "State" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetSegmentPath() string {
    return "track-item" + "[interface-name='" + fmt.Sprintf("%v", trackItem.InterfaceName) + "']" + "[virtual-router-id='" + fmt.Sprintf("%v", trackItem.VirtualRouterId) + "']" + "[tracked-interface-name='" + fmt.Sprintf("%v", trackItem.TrackedInterfaceName) + "']"
}

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = trackItem.InterfaceName
    leafs["virtual-router-id"] = trackItem.VirtualRouterId
    leafs["tracked-interface-name"] = trackItem.TrackedInterfaceName
    leafs["interface"] = trackItem.Interface
    leafs["virtual-router-id-xr"] = trackItem.VirtualRouterIdXr
    leafs["tracked-item-type"] = trackItem.TrackedItemType
    leafs["tracked-item-index"] = trackItem.TrackedItemIndex
    leafs["state"] = trackItem.State
    leafs["priority"] = trackItem.Priority
    return leafs
}

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetBundleName() string { return "cisco_ios_xr" }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetYangName() string { return "track-item" }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) SetParent(parent types.Entity) { trackItem.parent = parent }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetParent() types.Entity { return trackItem.parent }

func (trackItem *Vrrp_Ipv4_TrackItems_TrackItem) GetParentYangName() string { return "track-items" }

// Vrrp_Ipv4_VirtualRouters
// The VRRP virtual router table
type Vrrp_Ipv4_VirtualRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP virtual router. The type is slice of
    // Vrrp_Ipv4_VirtualRouters_VirtualRouter.
    VirtualRouter []Vrrp_Ipv4_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetFilter() yfilter.YFilter { return virtualRouters.YFilter }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) SetFilter(yf yfilter.YFilter) { virtualRouters.YFilter = yf }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetGoName(yname string) string {
    if yname == "virtual-router" { return "VirtualRouter" }
    return ""
}

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetSegmentPath() string {
    return "virtual-routers"
}

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-router" {
        for _, c := range virtualRouters.VirtualRouter {
            if virtualRouters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv4_VirtualRouters_VirtualRouter{}
        virtualRouters.VirtualRouter = append(virtualRouters.VirtualRouter, child)
        return &virtualRouters.VirtualRouter[len(virtualRouters.VirtualRouter)-1]
    }
    return nil
}

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range virtualRouters.VirtualRouter {
        children[virtualRouters.VirtualRouter[i].GetSegmentPath()] = &virtualRouters.VirtualRouter[i]
    }
    return children
}

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetYangName() string { return "virtual-routers" }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) SetParent(parent types.Entity) { virtualRouters.parent = parent }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetParent() types.Entity { return virtualRouters.parent }

func (virtualRouters *Vrrp_Ipv4_VirtualRouters) GetParentYangName() string { return "ipv4" }

// Vrrp_Ipv4_VirtualRouters_VirtualRouter
// A VRRP virtual router
type Vrrp_Ipv4_VirtualRouters_VirtualRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The VRRP virtual router id. The type is
    // interface{} with range: -2147483648..2147483647.
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
    Ipv6OperationalAddress []Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress

    // IPv6 Configured but Down VRRP addresses. The type is slice of
    // Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress.
    Ipv6ConfiguredDownAddress []Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress

    // State change history. The type is slice of
    // Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory.
    StateChangeHistory []Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory
}

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetFilter() yfilter.YFilter { return virtualRouter.YFilter }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) SetFilter(yf yfilter.YFilter) { virtualRouter.YFilter = yf }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "virtual-router-id-xr" { return "VirtualRouterIdXr" }
    if yname == "version" { return "Version" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "session-name" { return "SessionName" }
    if yname == "slaves" { return "Slaves" }
    if yname == "is-slave" { return "IsSlave" }
    if yname == "followed-session-name" { return "FollowedSessionName" }
    if yname == "secondary-address-count" { return "SecondaryAddressCount" }
    if yname == "operational-address-count" { return "OperationalAddressCount" }
    if yname == "primary-virtual-ip" { return "PrimaryVirtualIp" }
    if yname == "configured-down-address-count" { return "ConfiguredDownAddressCount" }
    if yname == "virtual-linklocal-ipv6-address" { return "VirtualLinklocalIpv6Address" }
    if yname == "primary-state" { return "PrimaryState" }
    if yname == "master-ip-address" { return "MasterIpAddress" }
    if yname == "master-ipv6-address" { return "MasterIpv6Address" }
    if yname == "master-priority" { return "MasterPriority" }
    if yname == "vrrp-state" { return "VrrpState" }
    if yname == "authentication-type" { return "AuthenticationType" }
    if yname == "authentication-string" { return "AuthenticationString" }
    if yname == "configured-advertize-time" { return "ConfiguredAdvertizeTime" }
    if yname == "oper-advertize-time" { return "OperAdvertizeTime" }
    if yname == "min-delay-time" { return "MinDelayTime" }
    if yname == "reload-delay-time" { return "ReloadDelayTime" }
    if yname == "delay-timer-flag" { return "DelayTimerFlag" }
    if yname == "delay-timer-secs" { return "DelayTimerSecs" }
    if yname == "delay-timer-msecs" { return "DelayTimerMsecs" }
    if yname == "authentication-flag" { return "AuthenticationFlag" }
    if yname == "force-timer-flag" { return "ForceTimerFlag" }
    if yname == "preempt-flag" { return "PreemptFlag" }
    if yname == "ip-address-owner-flag" { return "IpAddressOwnerFlag" }
    if yname == "is-accept-mode" { return "IsAcceptMode" }
    if yname == "preempt-delay-time" { return "PreemptDelayTime" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "operational-priority" { return "OperationalPriority" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    if yname == "tracked-interface-count" { return "TrackedInterfaceCount" }
    if yname == "tracked-interface-up-count" { return "TrackedInterfaceUpCount" }
    if yname == "tracked-item-count" { return "TrackedItemCount" }
    if yname == "tracked-item-up-count" { return "TrackedItemUpCount" }
    if yname == "time-in-current-state" { return "TimeInCurrentState" }
    if yname == "state-change-count" { return "StateChangeCount" }
    if yname == "time-vrouter-up" { return "TimeVrouterUp" }
    if yname == "master-count" { return "MasterCount" }
    if yname == "adverts-received-count" { return "AdvertsReceivedCount" }
    if yname == "advert-interval-error-count" { return "AdvertIntervalErrorCount" }
    if yname == "adverts-sent-count" { return "AdvertsSentCount" }
    if yname == "authentication-fail-count" { return "AuthenticationFailCount" }
    if yname == "ttl-error-count" { return "TtlErrorCount" }
    if yname == "priority-zero-received-count" { return "PriorityZeroReceivedCount" }
    if yname == "priority-zero-sent-count" { return "PriorityZeroSentCount" }
    if yname == "invalid-packet-count" { return "InvalidPacketCount" }
    if yname == "address-list-error-count" { return "AddressListErrorCount" }
    if yname == "invalid-auth-type-count" { return "InvalidAuthTypeCount" }
    if yname == "auth-type-mismatch-count" { return "AuthTypeMismatchCount" }
    if yname == "pkt-length-errors-count" { return "PktLengthErrorsCount" }
    if yname == "time-stats-discontinuity" { return "TimeStatsDiscontinuity" }
    if yname == "bfd-session-state" { return "BfdSessionState" }
    if yname == "bfd-interval" { return "BfdInterval" }
    if yname == "bfd-multiplier" { return "BfdMultiplier" }
    if yname == "bfd-cfg-remote-ip" { return "BfdCfgRemoteIp" }
    if yname == "bfd-configured-remote-ipv6-address" { return "BfdConfiguredRemoteIpv6Address" }
    if yname == "state-from-checkpoint" { return "StateFromCheckpoint" }
    if yname == "interface-ipv4-address" { return "InterfaceIpv4Address" }
    if yname == "interface-ipv6-address" { return "InterfaceIpv6Address" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "virtual-mac-address-state" { return "VirtualMacAddressState" }
    if yname == "operational-address" { return "OperationalAddress" }
    if yname == "ipv4-configured-down-address" { return "Ipv4ConfiguredDownAddress" }
    if yname == "resign-sent-time" { return "ResignSentTime" }
    if yname == "resign-received-time" { return "ResignReceivedTime" }
    if yname == "ipv6-operational-address" { return "Ipv6OperationalAddress" }
    if yname == "ipv6-configured-down-address" { return "Ipv6ConfiguredDownAddress" }
    if yname == "state-change-history" { return "StateChangeHistory" }
    return ""
}

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetSegmentPath() string {
    return "virtual-router" + "[interface-name='" + fmt.Sprintf("%v", virtualRouter.InterfaceName) + "']" + "[virtual-router-id='" + fmt.Sprintf("%v", virtualRouter.VirtualRouterId) + "']"
}

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "resign-sent-time" {
        return &virtualRouter.ResignSentTime
    }
    if childYangName == "resign-received-time" {
        return &virtualRouter.ResignReceivedTime
    }
    if childYangName == "ipv6-operational-address" {
        for _, c := range virtualRouter.Ipv6OperationalAddress {
            if virtualRouter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress{}
        virtualRouter.Ipv6OperationalAddress = append(virtualRouter.Ipv6OperationalAddress, child)
        return &virtualRouter.Ipv6OperationalAddress[len(virtualRouter.Ipv6OperationalAddress)-1]
    }
    if childYangName == "ipv6-configured-down-address" {
        for _, c := range virtualRouter.Ipv6ConfiguredDownAddress {
            if virtualRouter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress{}
        virtualRouter.Ipv6ConfiguredDownAddress = append(virtualRouter.Ipv6ConfiguredDownAddress, child)
        return &virtualRouter.Ipv6ConfiguredDownAddress[len(virtualRouter.Ipv6ConfiguredDownAddress)-1]
    }
    if childYangName == "state-change-history" {
        for _, c := range virtualRouter.StateChangeHistory {
            if virtualRouter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory{}
        virtualRouter.StateChangeHistory = append(virtualRouter.StateChangeHistory, child)
        return &virtualRouter.StateChangeHistory[len(virtualRouter.StateChangeHistory)-1]
    }
    return nil
}

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["resign-sent-time"] = &virtualRouter.ResignSentTime
    children["resign-received-time"] = &virtualRouter.ResignReceivedTime
    for i := range virtualRouter.Ipv6OperationalAddress {
        children[virtualRouter.Ipv6OperationalAddress[i].GetSegmentPath()] = &virtualRouter.Ipv6OperationalAddress[i]
    }
    for i := range virtualRouter.Ipv6ConfiguredDownAddress {
        children[virtualRouter.Ipv6ConfiguredDownAddress[i].GetSegmentPath()] = &virtualRouter.Ipv6ConfiguredDownAddress[i]
    }
    for i := range virtualRouter.StateChangeHistory {
        children[virtualRouter.StateChangeHistory[i].GetSegmentPath()] = &virtualRouter.StateChangeHistory[i]
    }
    return children
}

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = virtualRouter.InterfaceName
    leafs["virtual-router-id"] = virtualRouter.VirtualRouterId
    leafs["interface-name-xr"] = virtualRouter.InterfaceNameXr
    leafs["virtual-router-id-xr"] = virtualRouter.VirtualRouterIdXr
    leafs["version"] = virtualRouter.Version
    leafs["address-family"] = virtualRouter.AddressFamily
    leafs["session-name"] = virtualRouter.SessionName
    leafs["slaves"] = virtualRouter.Slaves
    leafs["is-slave"] = virtualRouter.IsSlave
    leafs["followed-session-name"] = virtualRouter.FollowedSessionName
    leafs["secondary-address-count"] = virtualRouter.SecondaryAddressCount
    leafs["operational-address-count"] = virtualRouter.OperationalAddressCount
    leafs["primary-virtual-ip"] = virtualRouter.PrimaryVirtualIp
    leafs["configured-down-address-count"] = virtualRouter.ConfiguredDownAddressCount
    leafs["virtual-linklocal-ipv6-address"] = virtualRouter.VirtualLinklocalIpv6Address
    leafs["primary-state"] = virtualRouter.PrimaryState
    leafs["master-ip-address"] = virtualRouter.MasterIpAddress
    leafs["master-ipv6-address"] = virtualRouter.MasterIpv6Address
    leafs["master-priority"] = virtualRouter.MasterPriority
    leafs["vrrp-state"] = virtualRouter.VrrpState
    leafs["authentication-type"] = virtualRouter.AuthenticationType
    leafs["authentication-string"] = virtualRouter.AuthenticationString
    leafs["configured-advertize-time"] = virtualRouter.ConfiguredAdvertizeTime
    leafs["oper-advertize-time"] = virtualRouter.OperAdvertizeTime
    leafs["min-delay-time"] = virtualRouter.MinDelayTime
    leafs["reload-delay-time"] = virtualRouter.ReloadDelayTime
    leafs["delay-timer-flag"] = virtualRouter.DelayTimerFlag
    leafs["delay-timer-secs"] = virtualRouter.DelayTimerSecs
    leafs["delay-timer-msecs"] = virtualRouter.DelayTimerMsecs
    leafs["authentication-flag"] = virtualRouter.AuthenticationFlag
    leafs["force-timer-flag"] = virtualRouter.ForceTimerFlag
    leafs["preempt-flag"] = virtualRouter.PreemptFlag
    leafs["ip-address-owner-flag"] = virtualRouter.IpAddressOwnerFlag
    leafs["is-accept-mode"] = virtualRouter.IsAcceptMode
    leafs["preempt-delay-time"] = virtualRouter.PreemptDelayTime
    leafs["configured-priority"] = virtualRouter.ConfiguredPriority
    leafs["operational-priority"] = virtualRouter.OperationalPriority
    leafs["priority-decrement"] = virtualRouter.PriorityDecrement
    leafs["tracked-interface-count"] = virtualRouter.TrackedInterfaceCount
    leafs["tracked-interface-up-count"] = virtualRouter.TrackedInterfaceUpCount
    leafs["tracked-item-count"] = virtualRouter.TrackedItemCount
    leafs["tracked-item-up-count"] = virtualRouter.TrackedItemUpCount
    leafs["time-in-current-state"] = virtualRouter.TimeInCurrentState
    leafs["state-change-count"] = virtualRouter.StateChangeCount
    leafs["time-vrouter-up"] = virtualRouter.TimeVrouterUp
    leafs["master-count"] = virtualRouter.MasterCount
    leafs["adverts-received-count"] = virtualRouter.AdvertsReceivedCount
    leafs["advert-interval-error-count"] = virtualRouter.AdvertIntervalErrorCount
    leafs["adverts-sent-count"] = virtualRouter.AdvertsSentCount
    leafs["authentication-fail-count"] = virtualRouter.AuthenticationFailCount
    leafs["ttl-error-count"] = virtualRouter.TtlErrorCount
    leafs["priority-zero-received-count"] = virtualRouter.PriorityZeroReceivedCount
    leafs["priority-zero-sent-count"] = virtualRouter.PriorityZeroSentCount
    leafs["invalid-packet-count"] = virtualRouter.InvalidPacketCount
    leafs["address-list-error-count"] = virtualRouter.AddressListErrorCount
    leafs["invalid-auth-type-count"] = virtualRouter.InvalidAuthTypeCount
    leafs["auth-type-mismatch-count"] = virtualRouter.AuthTypeMismatchCount
    leafs["pkt-length-errors-count"] = virtualRouter.PktLengthErrorsCount
    leafs["time-stats-discontinuity"] = virtualRouter.TimeStatsDiscontinuity
    leafs["bfd-session-state"] = virtualRouter.BfdSessionState
    leafs["bfd-interval"] = virtualRouter.BfdInterval
    leafs["bfd-multiplier"] = virtualRouter.BfdMultiplier
    leafs["bfd-cfg-remote-ip"] = virtualRouter.BfdCfgRemoteIp
    leafs["bfd-configured-remote-ipv6-address"] = virtualRouter.BfdConfiguredRemoteIpv6Address
    leafs["state-from-checkpoint"] = virtualRouter.StateFromCheckpoint
    leafs["interface-ipv4-address"] = virtualRouter.InterfaceIpv4Address
    leafs["interface-ipv6-address"] = virtualRouter.InterfaceIpv6Address
    leafs["virtual-mac-address"] = virtualRouter.VirtualMacAddress
    leafs["virtual-mac-address-state"] = virtualRouter.VirtualMacAddressState
    leafs["operational-address"] = virtualRouter.OperationalAddress
    leafs["ipv4-configured-down-address"] = virtualRouter.Ipv4ConfiguredDownAddress
    return leafs
}

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetYangName() string { return "virtual-router" }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) SetParent(parent types.Entity) { virtualRouter.parent = parent }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetParent() types.Entity { return virtualRouter.parent }

func (virtualRouter *Vrrp_Ipv4_VirtualRouters_VirtualRouter) GetParentYangName() string { return "virtual-routers" }

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime
// Time last resign was sent
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetFilter() yfilter.YFilter { return resignSentTime.YFilter }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) SetFilter(yf yfilter.YFilter) { resignSentTime.YFilter = yf }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetSegmentPath() string {
    return "resign-sent-time"
}

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resignSentTime.Seconds
    leafs["nanoseconds"] = resignSentTime.Nanoseconds
    return leafs
}

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetBundleName() string { return "cisco_ios_xr" }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetYangName() string { return "resign-sent-time" }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) SetParent(parent types.Entity) { resignSentTime.parent = parent }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetParent() types.Entity { return resignSentTime.parent }

func (resignSentTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignSentTime) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime
// Time last resign was received
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetFilter() yfilter.YFilter { return resignReceivedTime.YFilter }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) SetFilter(yf yfilter.YFilter) { resignReceivedTime.YFilter = yf }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetSegmentPath() string {
    return "resign-received-time"
}

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resignReceivedTime.Seconds
    leafs["nanoseconds"] = resignReceivedTime.Nanoseconds
    return leafs
}

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetBundleName() string { return "cisco_ios_xr" }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetYangName() string { return "resign-received-time" }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) SetParent(parent types.Entity) { resignReceivedTime.parent = parent }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetParent() types.Entity { return resignReceivedTime.parent }

func (resignReceivedTime *Vrrp_Ipv4_VirtualRouters_VirtualRouter_ResignReceivedTime) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress
// IPv6 Operational VRRP addresses
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetFilter() yfilter.YFilter { return ipv6OperationalAddress.YFilter }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) SetFilter(yf yfilter.YFilter) { ipv6OperationalAddress.YFilter = yf }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetSegmentPath() string {
    return "ipv6-operational-address"
}

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6OperationalAddress.Ipv6Address
    return leafs
}

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetYangName() string { return "ipv6-operational-address" }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) SetParent(parent types.Entity) { ipv6OperationalAddress.parent = parent }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetParent() types.Entity { return ipv6OperationalAddress.parent }

func (ipv6OperationalAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6OperationalAddress) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress
// IPv6 Configured but Down VRRP addresses
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetFilter() yfilter.YFilter { return ipv6ConfiguredDownAddress.YFilter }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) SetFilter(yf yfilter.YFilter) { ipv6ConfiguredDownAddress.YFilter = yf }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetSegmentPath() string {
    return "ipv6-configured-down-address"
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6ConfiguredDownAddress.Ipv6Address
    return leafs
}

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetYangName() string { return "ipv6-configured-down-address" }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) SetParent(parent types.Entity) { ipv6ConfiguredDownAddress.parent = parent }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetParent() types.Entity { return ipv6ConfiguredDownAddress.parent }

func (ipv6ConfiguredDownAddress *Vrrp_Ipv4_VirtualRouters_VirtualRouter_Ipv6ConfiguredDownAddress) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory
// State change history
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory struct {
    parent types.Entity
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

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetFilter() yfilter.YFilter { return stateChangeHistory.YFilter }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) SetFilter(yf yfilter.YFilter) { stateChangeHistory.YFilter = yf }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetGoName(yname string) string {
    if yname == "old-state" { return "OldState" }
    if yname == "new-state" { return "NewState" }
    if yname == "reason" { return "Reason" }
    if yname == "time" { return "Time" }
    return ""
}

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetSegmentPath() string {
    return "state-change-history"
}

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "time" {
        return &stateChangeHistory.Time
    }
    return nil
}

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["time"] = &stateChangeHistory.Time
    return children
}

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["old-state"] = stateChangeHistory.OldState
    leafs["new-state"] = stateChangeHistory.NewState
    leafs["reason"] = stateChangeHistory.Reason
    return leafs
}

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetBundleName() string { return "cisco_ios_xr" }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetYangName() string { return "state-change-history" }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) SetParent(parent types.Entity) { stateChangeHistory.parent = parent }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetParent() types.Entity { return stateChangeHistory.parent }

func (stateChangeHistory *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory) GetParentYangName() string { return "virtual-router" }

// Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time
// Time of state change
type Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetFilter() yfilter.YFilter { return time.YFilter }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) SetFilter(yf yfilter.YFilter) { time.YFilter = yf }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetSegmentPath() string {
    return "time"
}

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = time.Seconds
    leafs["nanoseconds"] = time.Nanoseconds
    return leafs
}

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetBundleName() string { return "cisco_ios_xr" }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetYangName() string { return "time" }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) SetParent(parent types.Entity) { time.parent = parent }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetParent() types.Entity { return time.parent }

func (time *Vrrp_Ipv4_VirtualRouters_VirtualRouter_StateChangeHistory_Time) GetParentYangName() string { return "state-change-history" }

// Vrrp_MgoSessions
// VRRP MGO Session information
type Vrrp_MgoSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP MGO Session. The type is slice of Vrrp_MgoSessions_MgoSession.
    MgoSession []Vrrp_MgoSessions_MgoSession
}

func (mgoSessions *Vrrp_MgoSessions) GetFilter() yfilter.YFilter { return mgoSessions.YFilter }

func (mgoSessions *Vrrp_MgoSessions) SetFilter(yf yfilter.YFilter) { mgoSessions.YFilter = yf }

func (mgoSessions *Vrrp_MgoSessions) GetGoName(yname string) string {
    if yname == "mgo-session" { return "MgoSession" }
    return ""
}

func (mgoSessions *Vrrp_MgoSessions) GetSegmentPath() string {
    return "mgo-sessions"
}

func (mgoSessions *Vrrp_MgoSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mgo-session" {
        for _, c := range mgoSessions.MgoSession {
            if mgoSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_MgoSessions_MgoSession{}
        mgoSessions.MgoSession = append(mgoSessions.MgoSession, child)
        return &mgoSessions.MgoSession[len(mgoSessions.MgoSession)-1]
    }
    return nil
}

func (mgoSessions *Vrrp_MgoSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mgoSessions.MgoSession {
        children[mgoSessions.MgoSession[i].GetSegmentPath()] = &mgoSessions.MgoSession[i]
    }
    return children
}

func (mgoSessions *Vrrp_MgoSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mgoSessions *Vrrp_MgoSessions) GetBundleName() string { return "cisco_ios_xr" }

func (mgoSessions *Vrrp_MgoSessions) GetYangName() string { return "mgo-sessions" }

func (mgoSessions *Vrrp_MgoSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mgoSessions *Vrrp_MgoSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mgoSessions *Vrrp_MgoSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mgoSessions *Vrrp_MgoSessions) SetParent(parent types.Entity) { mgoSessions.parent = parent }

func (mgoSessions *Vrrp_MgoSessions) GetParent() types.Entity { return mgoSessions.parent }

func (mgoSessions *Vrrp_MgoSessions) GetParentYangName() string { return "vrrp" }

// Vrrp_MgoSessions_MgoSession
// A VRRP MGO Session
type Vrrp_MgoSessions_MgoSession struct {
    parent types.Entity
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
    Slave []Vrrp_MgoSessions_MgoSession_Slave
}

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetFilter() yfilter.YFilter { return mgoSession.YFilter }

func (mgoSession *Vrrp_MgoSessions_MgoSession) SetFilter(yf yfilter.YFilter) { mgoSession.YFilter = yf }

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetGoName(yname string) string {
    if yname == "session-name" { return "SessionName" }
    if yname == "primary-session-name" { return "PrimarySessionName" }
    if yname == "primary-session-interface" { return "PrimarySessionInterface" }
    if yname == "primary-af-name" { return "PrimaryAfName" }
    if yname == "primary-session-number" { return "PrimarySessionNumber" }
    if yname == "primary-session-state" { return "PrimarySessionState" }
    if yname == "slave" { return "Slave" }
    return ""
}

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetSegmentPath() string {
    return "mgo-session" + "[session-name='" + fmt.Sprintf("%v", mgoSession.SessionName) + "']"
}

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slave" {
        for _, c := range mgoSession.Slave {
            if mgoSession.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_MgoSessions_MgoSession_Slave{}
        mgoSession.Slave = append(mgoSession.Slave, child)
        return &mgoSession.Slave[len(mgoSession.Slave)-1]
    }
    return nil
}

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mgoSession.Slave {
        children[mgoSession.Slave[i].GetSegmentPath()] = &mgoSession.Slave[i]
    }
    return children
}

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-name"] = mgoSession.SessionName
    leafs["primary-session-name"] = mgoSession.PrimarySessionName
    leafs["primary-session-interface"] = mgoSession.PrimarySessionInterface
    leafs["primary-af-name"] = mgoSession.PrimaryAfName
    leafs["primary-session-number"] = mgoSession.PrimarySessionNumber
    leafs["primary-session-state"] = mgoSession.PrimarySessionState
    return leafs
}

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetBundleName() string { return "cisco_ios_xr" }

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetYangName() string { return "mgo-session" }

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mgoSession *Vrrp_MgoSessions_MgoSession) SetParent(parent types.Entity) { mgoSession.parent = parent }

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetParent() types.Entity { return mgoSession.parent }

func (mgoSession *Vrrp_MgoSessions_MgoSession) GetParentYangName() string { return "mgo-sessions" }

// Vrrp_MgoSessions_MgoSession_Slave
// List of slaves following this primary session
type Vrrp_MgoSessions_MgoSession_Slave struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface of slave. The type is string with length: 0..64.
    SlaveInterface interface{}

    // VRID of slave. The type is interface{} with range: 0..4294967295.
    SlaveVirtualRouterId interface{}
}

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetFilter() yfilter.YFilter { return slave.YFilter }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) SetFilter(yf yfilter.YFilter) { slave.YFilter = yf }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetGoName(yname string) string {
    if yname == "slave-interface" { return "SlaveInterface" }
    if yname == "slave-virtual-router-id" { return "SlaveVirtualRouterId" }
    return ""
}

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetSegmentPath() string {
    return "slave"
}

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slave-interface"] = slave.SlaveInterface
    leafs["slave-virtual-router-id"] = slave.SlaveVirtualRouterId
    return leafs
}

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetBundleName() string { return "cisco_ios_xr" }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetYangName() string { return "slave" }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) SetParent(parent types.Entity) { slave.parent = parent }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetParent() types.Entity { return slave.parent }

func (slave *Vrrp_MgoSessions_MgoSession_Slave) GetParentYangName() string { return "mgo-session" }

