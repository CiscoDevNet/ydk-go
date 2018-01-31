// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-hsrp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hsrp: HSRP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_hsrp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_hsrp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-hsrp-oper hsrp}", reflect.TypeOf(Hsrp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-hsrp-oper:hsrp", reflect.TypeOf(Hsrp{}))
}

// HsrpVmacState represents Hsrp vmac state
type HsrpVmacState string

const (
    // VMAC stored locally
    HsrpVmacState_stored HsrpVmacState = "stored"

    // VMAC reserved in mac table
    HsrpVmacState_reserved HsrpVmacState = "reserved"

    // VMAC active in mac table
    HsrpVmacState_active HsrpVmacState = "active"

    // VMAC not yet reserved in mac table
    HsrpVmacState_reserving HsrpVmacState = "reserving"
)

// StandbyGrpState represents Standby grp state
type StandbyGrpState string

const (
    // Initial
    StandbyGrpState_state_initial StandbyGrpState = "state-initial"

    // Learn
    StandbyGrpState_state_learn StandbyGrpState = "state-learn"

    // Listen
    StandbyGrpState_state_listen StandbyGrpState = "state-listen"

    // Speak
    StandbyGrpState_state_speak StandbyGrpState = "state-speak"

    // Standby
    StandbyGrpState_state_standby StandbyGrpState = "state-standby"

    // Active
    StandbyGrpState_state_active StandbyGrpState = "state-active"
)

// HsrpStateChangeReason represents Hsrp state change reason
type HsrpStateChangeReason string

const (
    // BFD session down
    HsrpStateChangeReason_state_change_bfd_down HsrpStateChangeReason = "state-change-bfd-down"

    // Virtual IP learnt
    HsrpStateChangeReason_state_change_vip_learnt HsrpStateChangeReason = "state-change-vip-learnt"

    // Interface IP update
    HsrpStateChangeReason_state_change_interface_ip HsrpStateChangeReason = "state-change-interface-ip"

    // Delay timer expired
    HsrpStateChangeReason_state_change_delay_timer HsrpStateChangeReason = "state-change-delay-timer"

    // Ready on startup
    HsrpStateChangeReason_state_change_startup HsrpStateChangeReason = "state-change-startup"

    // HSRP shut down
    HsrpStateChangeReason_state_change_shutdown HsrpStateChangeReason = "state-change-shutdown"

    // Interface Up update
    HsrpStateChangeReason_state_change_interface_up HsrpStateChangeReason = "state-change-interface-up"

    // Interface Down update
    HsrpStateChangeReason_state_change_interface_down HsrpStateChangeReason = "state-change-interface-down"

    // Active timer expired
    HsrpStateChangeReason_state_change_active_timer HsrpStateChangeReason = "state-change-active-timer"

    // Standby timer expired
    HsrpStateChangeReason_state_change_standby_timer HsrpStateChangeReason = "state-change-standby-timer"

    // Resign received
    HsrpStateChangeReason_state_change_resign HsrpStateChangeReason = "state-change-resign"

    // Coup received
    HsrpStateChangeReason_state_change_coup HsrpStateChangeReason = "state-change-coup"

    // Higher priority speak received
    HsrpStateChangeReason_state_change_higher_priority_speak HsrpStateChangeReason = "state-change-higher-priority-speak"

    // Higher priority standby received
    HsrpStateChangeReason_state_change_higher_priority_standby HsrpStateChangeReason = "state-change-higher-priority-standby"

    // Lower priority standby received
    HsrpStateChangeReason_state_change_lower_priority_standby HsrpStateChangeReason = "state-change-lower-priority-standby"

    // Higher priority active received
    HsrpStateChangeReason_state_change_higher_priority_active HsrpStateChangeReason = "state-change-higher-priority-active"

    // Lower priority active received
    HsrpStateChangeReason_state_change_lower_priority_active HsrpStateChangeReason = "state-change-lower-priority-active"

    // Virtual IP configured
    HsrpStateChangeReason_state_change_virtual_ip_configured HsrpStateChangeReason = "state-change-virtual-ip-configured"

    // Virtual IP lost
    HsrpStateChangeReason_state_change_virtual_ip_lost HsrpStateChangeReason = "state-change-virtual-ip-lost"

    // Recovered from checkpoint
    HsrpStateChangeReason_state_change_recovered_from_checkpoint HsrpStateChangeReason = "state-change-recovered-from-checkpoint"

    // MAC address update
    HsrpStateChangeReason_state_change_mac_update HsrpStateChangeReason = "state-change-mac-update"

    // Forwarder Admin state change
    HsrpStateChangeReason_state_change_admin HsrpStateChangeReason = "state-change-admin"

    // MGO parent change
    HsrpStateChangeReason_state_change_parent HsrpStateChangeReason = "state-change-parent"

    // Checkpoint update from Primary HSRP instance
    HsrpStateChangeReason_state_change_chkpt_update HsrpStateChangeReason = "state-change-chkpt-update"

    // Resync following ISSU primary event
    HsrpStateChangeReason_state_change_issu_resync HsrpStateChangeReason = "state-change-issu-resync"

    // Reset to learn parameters
    HsrpStateChangeReason_state_change_reset_to_learn HsrpStateChangeReason = "state-change-reset-to-learn"

    // Maximum reason in enumeration
    HsrpStateChangeReason_state_change_max HsrpStateChangeReason = "state-change-max"
)

// HsrpBAf represents Hsrp b af
type HsrpBAf string

const (
    // IPv4 Address Family
    HsrpBAf_ipv4 HsrpBAf = "ipv4"

    // IPv6 Address Family
    HsrpBAf_ipv6 HsrpBAf = "ipv6"

    // The number of supported address families
    HsrpBAf_count HsrpBAf = "count"
)

// HsrpBfdSessionState represents Hsrp bfd session state
type HsrpBfdSessionState string

const (
    // None
    HsrpBfdSessionState_bfd_state_none HsrpBfdSessionState = "bfd-state-none"

    // Inactive
    HsrpBfdSessionState_bfd_state_inactive HsrpBfdSessionState = "bfd-state-inactive"

    // Up
    HsrpBfdSessionState_bfd_state_up HsrpBfdSessionState = "bfd-state-up"

    // Down
    HsrpBfdSessionState_bfd_state_down HsrpBfdSessionState = "bfd-state-down"
)

// Hsrp
// HSRP operational data
type Hsrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 HSRP information.
    Ipv4 Hsrp_Ipv4

    // HSRP MGO session table.
    MgoSessions Hsrp_MgoSessions

    // IPv6 HSRP information.
    Ipv6 Hsrp_Ipv6

    // The table of HSRP BFD Sessions.
    BfdSessions Hsrp_BfdSessions

    // HSRP summary statistics.
    Summary Hsrp_Summary
}

func (hsrp *Hsrp) GetFilter() yfilter.YFilter { return hsrp.YFilter }

func (hsrp *Hsrp) SetFilter(yf yfilter.YFilter) { hsrp.YFilter = yf }

func (hsrp *Hsrp) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    if yname == "mgo-sessions" { return "MgoSessions" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "bfd-sessions" { return "BfdSessions" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (hsrp *Hsrp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-hsrp-oper:hsrp"
}

func (hsrp *Hsrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &hsrp.Ipv4
    }
    if childYangName == "mgo-sessions" {
        return &hsrp.MgoSessions
    }
    if childYangName == "ipv6" {
        return &hsrp.Ipv6
    }
    if childYangName == "bfd-sessions" {
        return &hsrp.BfdSessions
    }
    if childYangName == "summary" {
        return &hsrp.Summary
    }
    return nil
}

func (hsrp *Hsrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &hsrp.Ipv4
    children["mgo-sessions"] = &hsrp.MgoSessions
    children["ipv6"] = &hsrp.Ipv6
    children["bfd-sessions"] = &hsrp.BfdSessions
    children["summary"] = &hsrp.Summary
    return children
}

func (hsrp *Hsrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hsrp *Hsrp) GetBundleName() string { return "cisco_ios_xr" }

func (hsrp *Hsrp) GetYangName() string { return "hsrp" }

func (hsrp *Hsrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hsrp *Hsrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hsrp *Hsrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hsrp *Hsrp) SetParent(parent types.Entity) { hsrp.parent = parent }

func (hsrp *Hsrp) GetParent() types.Entity { return hsrp.parent }

func (hsrp *Hsrp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-hsrp-oper" }

// Hsrp_Ipv4
// IPv4 HSRP information
type Hsrp_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP standby group table.
    Groups Hsrp_Ipv4_Groups

    // The HSRP tracked interfaces table.
    TrackedInterfaces Hsrp_Ipv4_TrackedInterfaces

    // The HSRP interface information table.
    Interfaces Hsrp_Ipv4_Interfaces
}

func (ipv4 *Hsrp_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Hsrp_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Hsrp_Ipv4) GetGoName(yname string) string {
    if yname == "groups" { return "Groups" }
    if yname == "tracked-interfaces" { return "TrackedInterfaces" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv4 *Hsrp_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Hsrp_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &ipv4.Groups
    }
    if childYangName == "tracked-interfaces" {
        return &ipv4.TrackedInterfaces
    }
    if childYangName == "interfaces" {
        return &ipv4.Interfaces
    }
    return nil
}

func (ipv4 *Hsrp_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &ipv4.Groups
    children["tracked-interfaces"] = &ipv4.TrackedInterfaces
    children["interfaces"] = &ipv4.Interfaces
    return children
}

func (ipv4 *Hsrp_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Hsrp_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Hsrp_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Hsrp_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Hsrp_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Hsrp_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Hsrp_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Hsrp_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Hsrp_Ipv4) GetParentYangName() string { return "hsrp" }

// Hsrp_Ipv4_Groups
// The HSRP standby group table
type Hsrp_Ipv4_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An HSRP standby group. The type is slice of Hsrp_Ipv4_Groups_Group.
    Group []Hsrp_Ipv4_Groups_Group
}

func (groups *Hsrp_Ipv4_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Hsrp_Ipv4_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Hsrp_Ipv4_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Hsrp_Ipv4_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Hsrp_Ipv4_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv4_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Hsrp_Ipv4_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Hsrp_Ipv4_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Hsrp_Ipv4_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Hsrp_Ipv4_Groups) GetYangName() string { return "groups" }

func (groups *Hsrp_Ipv4_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Hsrp_Ipv4_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Hsrp_Ipv4_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Hsrp_Ipv4_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Hsrp_Ipv4_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Hsrp_Ipv4_Groups) GetParentYangName() string { return "ipv4" }

// Hsrp_Ipv4_Groups_Group
// An HSRP standby group
type Hsrp_Ipv4_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The HSRP group number. The type is interface{}
    // with range: -2147483648..2147483647.
    GroupNumber interface{}

    // Authentication string. The type is string with length: 0..9.
    AuthenticationString interface{}

    // Virtual mac address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // HSRP Group number. The type is interface{} with range: 0..4294967295.
    HsrpGroupNumber interface{}

    // Address family. The type is HsrpBAf.
    AddressFamily interface{}

    // HSRP Protocol Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Session Name. The type is string with length: 0..16.
    SessionName interface{}

    // Number of slaves following state. The type is interface{} with range:
    // 0..4294967295.
    Slaves interface{}

    // Group is a slave group. The type is bool.
    IsSlave interface{}

    // Followed Session Name. The type is string with length: 0..16.
    FollowedSessionName interface{}

    // Configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // Preempt delay time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    PreemptDelay interface{}

    // Preempt time remaining in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    PreemptTimerSecs interface{}

    // Hellotime in msecs. The type is interface{} with range: 0..4294967295.
    // Units are millisecond.
    HelloTime interface{}

    // Holdtime in msecs. The type is interface{} with range: 0..4294967295. Units
    // are millisecond.
    HoldTime interface{}

    // Learned hellotime in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    LearnedHelloTime interface{}

    // Learned holdtime in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    LearnedHoldTime interface{}

    // Minimum delay time in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    MinDelayTime interface{}

    // Reload delay time in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    ReloadDelayTime interface{}

    // Configured Virtual IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    VirtualIpAddress interface{}

    // Virtual linklocal IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualLinklocalIpv6Address interface{}

    // Active router's IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ActiveIpAddress interface{}

    // Active router's IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ActiveIpv6Address interface{}

    // Active router's interface MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    ActiveMacAddress interface{}

    // Standby router's IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StandbyIpAddress interface{}

    // Standby router's IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StandbyIpv6Address interface{}

    // Standby router's interface MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    StandbyMacAddress interface{}

    // HSRP router state. The type is StandbyGrpState.
    HsrpRouterState interface{}

    // Interface Name. The type is string with length: 0..64.
    InterfaceNameXr interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Priority of the router. The type is interface{} with range: 0..255.
    RouterPriority interface{}

    // Priority of the Active router. The type is interface{} with range: 0..255.
    ActivePriority interface{}

    // Active timer running flag. The type is bool.
    ActiveTimerFlag interface{}

    // Active timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    ActiveTimerSecs interface{}

    // Active timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    ActiveTimerMsecs interface{}

    // Standby timer running flag. The type is bool.
    StandbyTimerFlag interface{}

    // Standby timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    StandbyTimerSecs interface{}

    // Standby timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    StandbyTimerMsecs interface{}

    // Hello timer running flag. The type is bool.
    HelloTimerFlag interface{}

    // Hello timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    HelloTimerSecs interface{}

    // Hello timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    HelloTimerMsecs interface{}

    // Delay timer running flag. The type is bool.
    DelayTimerFlag interface{}

    // Delay timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    DelayTimerSecs interface{}

    // Delay timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    DelayTimerMsecs interface{}

    // Time in current state secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    CurrentStateTimerSecs interface{}

    // Number of state changes. The type is interface{} with range: 0..4294967295.
    StateChangeCount interface{}

    // Number of tracked interfaces. The type is interface{} with range:
    // 0..4294967295.
    TrackedInterfaceCount interface{}

    // Number of tracked interfaces up. The type is interface{} with range:
    // 0..4294967295.
    TrackedInterfaceUpCount interface{}

    // Preempt enabled. The type is bool.
    PreemptEnabled interface{}

    // Use configured timers. The type is bool.
    UseConfiguredTimers interface{}

    // Use configured virtual IP. The type is bool.
    UseConfiguredVirtualIp interface{}

    // Use burnt in MAC address configured. The type is bool.
    UseBiaConfigured interface{}

    // Non-default timers are configured. The type is bool.
    ConfiguredTimers interface{}

    // MAC address configured. The type is bool.
    ConfiguredMacAddress interface{}

    // HSRP redirects disabled. The type is bool.
    RedirectsDisabled interface{}

    // HSRP BFD fast failover. The type is bool.
    BfdEnabled interface{}

    // BFD Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    BfdInterface interface{}

    // BFD Peer IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BfdPeerIpAddress interface{}

    // BFD Peer IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    BfdPeerIpv6Address interface{}

    // BFD session state. The type is HsrpBfdSessionState.
    BfdSessionState interface{}

    // BFD packet send interval. The type is interface{} with range:
    // 0..4294967295.
    BfdInterval interface{}

    // BFD multiplier. The type is interface{} with range: 0..4294967295.
    BfdMultiplier interface{}

    // Virtual mac address state. The type is HsrpVmacState.
    VirtualMacAddressState interface{}

    // Secondary virtual IP addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SecondaryAddress []interface{}

    // Time last resign was sent.
    ResignSentTime Hsrp_Ipv4_Groups_Group_ResignSentTime

    // Time last resign was received.
    ResignReceivedTime Hsrp_Ipv4_Groups_Group_ResignReceivedTime

    // Time last coup was sent.
    CoupSentTime Hsrp_Ipv4_Groups_Group_CoupSentTime

    // Time last coup was received.
    CoupReceivedTime Hsrp_Ipv4_Groups_Group_CoupReceivedTime

    // HSRP Group statistics.
    Statistics Hsrp_Ipv4_Groups_Group_Statistics

    // Global virtual IPv6 addresses. The type is slice of
    // Hsrp_Ipv4_Groups_Group_GlobalAddress.
    GlobalAddress []Hsrp_Ipv4_Groups_Group_GlobalAddress

    // State change history. The type is slice of
    // Hsrp_Ipv4_Groups_Group_StateChangeHistory.
    StateChangeHistory []Hsrp_Ipv4_Groups_Group_StateChangeHistory
}

func (group *Hsrp_Ipv4_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Hsrp_Ipv4_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Hsrp_Ipv4_Groups_Group) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "group-number" { return "GroupNumber" }
    if yname == "authentication-string" { return "AuthenticationString" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "hsrp-group-number" { return "HsrpGroupNumber" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "version" { return "Version" }
    if yname == "session-name" { return "SessionName" }
    if yname == "slaves" { return "Slaves" }
    if yname == "is-slave" { return "IsSlave" }
    if yname == "followed-session-name" { return "FollowedSessionName" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "preempt-timer-secs" { return "PreemptTimerSecs" }
    if yname == "hello-time" { return "HelloTime" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "learned-hello-time" { return "LearnedHelloTime" }
    if yname == "learned-hold-time" { return "LearnedHoldTime" }
    if yname == "min-delay-time" { return "MinDelayTime" }
    if yname == "reload-delay-time" { return "ReloadDelayTime" }
    if yname == "virtual-ip-address" { return "VirtualIpAddress" }
    if yname == "virtual-linklocal-ipv6-address" { return "VirtualLinklocalIpv6Address" }
    if yname == "active-ip-address" { return "ActiveIpAddress" }
    if yname == "active-ipv6-address" { return "ActiveIpv6Address" }
    if yname == "active-mac-address" { return "ActiveMacAddress" }
    if yname == "standby-ip-address" { return "StandbyIpAddress" }
    if yname == "standby-ipv6-address" { return "StandbyIpv6Address" }
    if yname == "standby-mac-address" { return "StandbyMacAddress" }
    if yname == "hsrp-router-state" { return "HsrpRouterState" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface" { return "Interface" }
    if yname == "router-priority" { return "RouterPriority" }
    if yname == "active-priority" { return "ActivePriority" }
    if yname == "active-timer-flag" { return "ActiveTimerFlag" }
    if yname == "active-timer-secs" { return "ActiveTimerSecs" }
    if yname == "active-timer-msecs" { return "ActiveTimerMsecs" }
    if yname == "standby-timer-flag" { return "StandbyTimerFlag" }
    if yname == "standby-timer-secs" { return "StandbyTimerSecs" }
    if yname == "standby-timer-msecs" { return "StandbyTimerMsecs" }
    if yname == "hello-timer-flag" { return "HelloTimerFlag" }
    if yname == "hello-timer-secs" { return "HelloTimerSecs" }
    if yname == "hello-timer-msecs" { return "HelloTimerMsecs" }
    if yname == "delay-timer-flag" { return "DelayTimerFlag" }
    if yname == "delay-timer-secs" { return "DelayTimerSecs" }
    if yname == "delay-timer-msecs" { return "DelayTimerMsecs" }
    if yname == "current-state-timer-secs" { return "CurrentStateTimerSecs" }
    if yname == "state-change-count" { return "StateChangeCount" }
    if yname == "tracked-interface-count" { return "TrackedInterfaceCount" }
    if yname == "tracked-interface-up-count" { return "TrackedInterfaceUpCount" }
    if yname == "preempt-enabled" { return "PreemptEnabled" }
    if yname == "use-configured-timers" { return "UseConfiguredTimers" }
    if yname == "use-configured-virtual-ip" { return "UseConfiguredVirtualIp" }
    if yname == "use-bia-configured" { return "UseBiaConfigured" }
    if yname == "configured-timers" { return "ConfiguredTimers" }
    if yname == "configured-mac-address" { return "ConfiguredMacAddress" }
    if yname == "redirects-disabled" { return "RedirectsDisabled" }
    if yname == "bfd-enabled" { return "BfdEnabled" }
    if yname == "bfd-interface" { return "BfdInterface" }
    if yname == "bfd-peer-ip-address" { return "BfdPeerIpAddress" }
    if yname == "bfd-peer-ipv6-address" { return "BfdPeerIpv6Address" }
    if yname == "bfd-session-state" { return "BfdSessionState" }
    if yname == "bfd-interval" { return "BfdInterval" }
    if yname == "bfd-multiplier" { return "BfdMultiplier" }
    if yname == "virtual-mac-address-state" { return "VirtualMacAddressState" }
    if yname == "secondary-address" { return "SecondaryAddress" }
    if yname == "resign-sent-time" { return "ResignSentTime" }
    if yname == "resign-received-time" { return "ResignReceivedTime" }
    if yname == "coup-sent-time" { return "CoupSentTime" }
    if yname == "coup-received-time" { return "CoupReceivedTime" }
    if yname == "statistics" { return "Statistics" }
    if yname == "global-address" { return "GlobalAddress" }
    if yname == "state-change-history" { return "StateChangeHistory" }
    return ""
}

func (group *Hsrp_Ipv4_Groups_Group) GetSegmentPath() string {
    return "group" + "[interface-name='" + fmt.Sprintf("%v", group.InterfaceName) + "']" + "[group-number='" + fmt.Sprintf("%v", group.GroupNumber) + "']"
}

func (group *Hsrp_Ipv4_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "resign-sent-time" {
        return &group.ResignSentTime
    }
    if childYangName == "resign-received-time" {
        return &group.ResignReceivedTime
    }
    if childYangName == "coup-sent-time" {
        return &group.CoupSentTime
    }
    if childYangName == "coup-received-time" {
        return &group.CoupReceivedTime
    }
    if childYangName == "statistics" {
        return &group.Statistics
    }
    if childYangName == "global-address" {
        for _, c := range group.GlobalAddress {
            if group.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv4_Groups_Group_GlobalAddress{}
        group.GlobalAddress = append(group.GlobalAddress, child)
        return &group.GlobalAddress[len(group.GlobalAddress)-1]
    }
    if childYangName == "state-change-history" {
        for _, c := range group.StateChangeHistory {
            if group.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv4_Groups_Group_StateChangeHistory{}
        group.StateChangeHistory = append(group.StateChangeHistory, child)
        return &group.StateChangeHistory[len(group.StateChangeHistory)-1]
    }
    return nil
}

func (group *Hsrp_Ipv4_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["resign-sent-time"] = &group.ResignSentTime
    children["resign-received-time"] = &group.ResignReceivedTime
    children["coup-sent-time"] = &group.CoupSentTime
    children["coup-received-time"] = &group.CoupReceivedTime
    children["statistics"] = &group.Statistics
    for i := range group.GlobalAddress {
        children[group.GlobalAddress[i].GetSegmentPath()] = &group.GlobalAddress[i]
    }
    for i := range group.StateChangeHistory {
        children[group.StateChangeHistory[i].GetSegmentPath()] = &group.StateChangeHistory[i]
    }
    return children
}

func (group *Hsrp_Ipv4_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = group.InterfaceName
    leafs["group-number"] = group.GroupNumber
    leafs["authentication-string"] = group.AuthenticationString
    leafs["virtual-mac-address"] = group.VirtualMacAddress
    leafs["hsrp-group-number"] = group.HsrpGroupNumber
    leafs["address-family"] = group.AddressFamily
    leafs["version"] = group.Version
    leafs["session-name"] = group.SessionName
    leafs["slaves"] = group.Slaves
    leafs["is-slave"] = group.IsSlave
    leafs["followed-session-name"] = group.FollowedSessionName
    leafs["configured-priority"] = group.ConfiguredPriority
    leafs["preempt-delay"] = group.PreemptDelay
    leafs["preempt-timer-secs"] = group.PreemptTimerSecs
    leafs["hello-time"] = group.HelloTime
    leafs["hold-time"] = group.HoldTime
    leafs["learned-hello-time"] = group.LearnedHelloTime
    leafs["learned-hold-time"] = group.LearnedHoldTime
    leafs["min-delay-time"] = group.MinDelayTime
    leafs["reload-delay-time"] = group.ReloadDelayTime
    leafs["virtual-ip-address"] = group.VirtualIpAddress
    leafs["virtual-linklocal-ipv6-address"] = group.VirtualLinklocalIpv6Address
    leafs["active-ip-address"] = group.ActiveIpAddress
    leafs["active-ipv6-address"] = group.ActiveIpv6Address
    leafs["active-mac-address"] = group.ActiveMacAddress
    leafs["standby-ip-address"] = group.StandbyIpAddress
    leafs["standby-ipv6-address"] = group.StandbyIpv6Address
    leafs["standby-mac-address"] = group.StandbyMacAddress
    leafs["hsrp-router-state"] = group.HsrpRouterState
    leafs["interface-name-xr"] = group.InterfaceNameXr
    leafs["interface"] = group.Interface
    leafs["router-priority"] = group.RouterPriority
    leafs["active-priority"] = group.ActivePriority
    leafs["active-timer-flag"] = group.ActiveTimerFlag
    leafs["active-timer-secs"] = group.ActiveTimerSecs
    leafs["active-timer-msecs"] = group.ActiveTimerMsecs
    leafs["standby-timer-flag"] = group.StandbyTimerFlag
    leafs["standby-timer-secs"] = group.StandbyTimerSecs
    leafs["standby-timer-msecs"] = group.StandbyTimerMsecs
    leafs["hello-timer-flag"] = group.HelloTimerFlag
    leafs["hello-timer-secs"] = group.HelloTimerSecs
    leafs["hello-timer-msecs"] = group.HelloTimerMsecs
    leafs["delay-timer-flag"] = group.DelayTimerFlag
    leafs["delay-timer-secs"] = group.DelayTimerSecs
    leafs["delay-timer-msecs"] = group.DelayTimerMsecs
    leafs["current-state-timer-secs"] = group.CurrentStateTimerSecs
    leafs["state-change-count"] = group.StateChangeCount
    leafs["tracked-interface-count"] = group.TrackedInterfaceCount
    leafs["tracked-interface-up-count"] = group.TrackedInterfaceUpCount
    leafs["preempt-enabled"] = group.PreemptEnabled
    leafs["use-configured-timers"] = group.UseConfiguredTimers
    leafs["use-configured-virtual-ip"] = group.UseConfiguredVirtualIp
    leafs["use-bia-configured"] = group.UseBiaConfigured
    leafs["configured-timers"] = group.ConfiguredTimers
    leafs["configured-mac-address"] = group.ConfiguredMacAddress
    leafs["redirects-disabled"] = group.RedirectsDisabled
    leafs["bfd-enabled"] = group.BfdEnabled
    leafs["bfd-interface"] = group.BfdInterface
    leafs["bfd-peer-ip-address"] = group.BfdPeerIpAddress
    leafs["bfd-peer-ipv6-address"] = group.BfdPeerIpv6Address
    leafs["bfd-session-state"] = group.BfdSessionState
    leafs["bfd-interval"] = group.BfdInterval
    leafs["bfd-multiplier"] = group.BfdMultiplier
    leafs["virtual-mac-address-state"] = group.VirtualMacAddressState
    leafs["secondary-address"] = group.SecondaryAddress
    return leafs
}

func (group *Hsrp_Ipv4_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Hsrp_Ipv4_Groups_Group) GetYangName() string { return "group" }

func (group *Hsrp_Ipv4_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Hsrp_Ipv4_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Hsrp_Ipv4_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Hsrp_Ipv4_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Hsrp_Ipv4_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Hsrp_Ipv4_Groups_Group) GetParentYangName() string { return "groups" }

// Hsrp_Ipv4_Groups_Group_ResignSentTime
// Time last resign was sent
type Hsrp_Ipv4_Groups_Group_ResignSentTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetFilter() yfilter.YFilter { return resignSentTime.YFilter }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) SetFilter(yf yfilter.YFilter) { resignSentTime.YFilter = yf }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetSegmentPath() string {
    return "resign-sent-time"
}

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resignSentTime.Seconds
    leafs["nanoseconds"] = resignSentTime.Nanoseconds
    return leafs
}

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetBundleName() string { return "cisco_ios_xr" }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetYangName() string { return "resign-sent-time" }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) SetParent(parent types.Entity) { resignSentTime.parent = parent }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetParent() types.Entity { return resignSentTime.parent }

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetParentYangName() string { return "group" }

// Hsrp_Ipv4_Groups_Group_ResignReceivedTime
// Time last resign was received
type Hsrp_Ipv4_Groups_Group_ResignReceivedTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetFilter() yfilter.YFilter { return resignReceivedTime.YFilter }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) SetFilter(yf yfilter.YFilter) { resignReceivedTime.YFilter = yf }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetSegmentPath() string {
    return "resign-received-time"
}

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resignReceivedTime.Seconds
    leafs["nanoseconds"] = resignReceivedTime.Nanoseconds
    return leafs
}

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetBundleName() string { return "cisco_ios_xr" }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetYangName() string { return "resign-received-time" }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) SetParent(parent types.Entity) { resignReceivedTime.parent = parent }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetParent() types.Entity { return resignReceivedTime.parent }

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetParentYangName() string { return "group" }

// Hsrp_Ipv4_Groups_Group_CoupSentTime
// Time last coup was sent
type Hsrp_Ipv4_Groups_Group_CoupSentTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetFilter() yfilter.YFilter { return coupSentTime.YFilter }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) SetFilter(yf yfilter.YFilter) { coupSentTime.YFilter = yf }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetSegmentPath() string {
    return "coup-sent-time"
}

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = coupSentTime.Seconds
    leafs["nanoseconds"] = coupSentTime.Nanoseconds
    return leafs
}

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetBundleName() string { return "cisco_ios_xr" }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetYangName() string { return "coup-sent-time" }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) SetParent(parent types.Entity) { coupSentTime.parent = parent }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetParent() types.Entity { return coupSentTime.parent }

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetParentYangName() string { return "group" }

// Hsrp_Ipv4_Groups_Group_CoupReceivedTime
// Time last coup was received
type Hsrp_Ipv4_Groups_Group_CoupReceivedTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetFilter() yfilter.YFilter { return coupReceivedTime.YFilter }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) SetFilter(yf yfilter.YFilter) { coupReceivedTime.YFilter = yf }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetSegmentPath() string {
    return "coup-received-time"
}

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = coupReceivedTime.Seconds
    leafs["nanoseconds"] = coupReceivedTime.Nanoseconds
    return leafs
}

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetBundleName() string { return "cisco_ios_xr" }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetYangName() string { return "coup-received-time" }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) SetParent(parent types.Entity) { coupReceivedTime.parent = parent }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetParent() types.Entity { return coupReceivedTime.parent }

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetParentYangName() string { return "group" }

// Hsrp_Ipv4_Groups_Group_Statistics
// HSRP Group statistics
type Hsrp_Ipv4_Groups_Group_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of transitions to Active State. The type is interface{} with range:
    // 0..4294967295.
    ActiveTransitions interface{}

    // Number of transitions to Standby State. The type is interface{} with range:
    // 0..4294967295.
    StandbyTransitions interface{}

    // Number of transitions to Speak State. The type is interface{} with range:
    // 0..4294967295.
    SpeakTransitions interface{}

    // Number of transitions to Listen State. The type is interface{} with range:
    // 0..4294967295.
    ListenTransitions interface{}

    // Number of transitions to Learn State. The type is interface{} with range:
    // 0..4294967295.
    LearnTransitions interface{}

    // Number of transitions to Init State. The type is interface{} with range:
    // 0..4294967295.
    InitTransitions interface{}

    // Number of Hello Packets sent (NB: Bundles only). The type is interface{}
    // with range: 0..4294967295.
    HelloPacketsSent interface{}

    // Number of Resign Packets sent. The type is interface{} with range:
    // 0..4294967295.
    ResignPacketsSent interface{}

    // Number of Coup Packets sent. The type is interface{} with range:
    // 0..4294967295.
    CoupPacketsSent interface{}

    // Number of Hello Packets received. The type is interface{} with range:
    // 0..4294967295.
    HelloPacketsReceived interface{}

    // Number of Resign Packets received. The type is interface{} with range:
    // 0..4294967295.
    ResignPacketsReceived interface{}

    // Number of Coup Packets received. The type is interface{} with range:
    // 0..4294967295.
    CoupPacketsReceived interface{}

    // Number of Packets received that failed authentication. The type is
    // interface{} with range: 0..4294967295.
    AuthFailReceived interface{}

    // Number of packets received with invalid Hello Time value. The type is
    // interface{} with range: 0..4294967295.
    InvalidTimerReceived interface{}

    // Number of packets received with mismatching virtual IP address. The type is
    // interface{} with range: 0..4294967295.
    MismatchVirtualIpAddressReceived interface{}
}

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetGoName(yname string) string {
    if yname == "active-transitions" { return "ActiveTransitions" }
    if yname == "standby-transitions" { return "StandbyTransitions" }
    if yname == "speak-transitions" { return "SpeakTransitions" }
    if yname == "listen-transitions" { return "ListenTransitions" }
    if yname == "learn-transitions" { return "LearnTransitions" }
    if yname == "init-transitions" { return "InitTransitions" }
    if yname == "hello-packets-sent" { return "HelloPacketsSent" }
    if yname == "resign-packets-sent" { return "ResignPacketsSent" }
    if yname == "coup-packets-sent" { return "CoupPacketsSent" }
    if yname == "hello-packets-received" { return "HelloPacketsReceived" }
    if yname == "resign-packets-received" { return "ResignPacketsReceived" }
    if yname == "coup-packets-received" { return "CoupPacketsReceived" }
    if yname == "auth-fail-received" { return "AuthFailReceived" }
    if yname == "invalid-timer-received" { return "InvalidTimerReceived" }
    if yname == "mismatch-virtual-ip-address-received" { return "MismatchVirtualIpAddressReceived" }
    return ""
}

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-transitions"] = statistics.ActiveTransitions
    leafs["standby-transitions"] = statistics.StandbyTransitions
    leafs["speak-transitions"] = statistics.SpeakTransitions
    leafs["listen-transitions"] = statistics.ListenTransitions
    leafs["learn-transitions"] = statistics.LearnTransitions
    leafs["init-transitions"] = statistics.InitTransitions
    leafs["hello-packets-sent"] = statistics.HelloPacketsSent
    leafs["resign-packets-sent"] = statistics.ResignPacketsSent
    leafs["coup-packets-sent"] = statistics.CoupPacketsSent
    leafs["hello-packets-received"] = statistics.HelloPacketsReceived
    leafs["resign-packets-received"] = statistics.ResignPacketsReceived
    leafs["coup-packets-received"] = statistics.CoupPacketsReceived
    leafs["auth-fail-received"] = statistics.AuthFailReceived
    leafs["invalid-timer-received"] = statistics.InvalidTimerReceived
    leafs["mismatch-virtual-ip-address-received"] = statistics.MismatchVirtualIpAddressReceived
    return leafs
}

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetYangName() string { return "statistics" }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetParentYangName() string { return "group" }

// Hsrp_Ipv4_Groups_Group_GlobalAddress
// Global virtual IPv6 addresses
type Hsrp_Ipv4_Groups_Group_GlobalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetFilter() yfilter.YFilter { return globalAddress.YFilter }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) SetFilter(yf yfilter.YFilter) { globalAddress.YFilter = yf }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetSegmentPath() string {
    return "global-address"
}

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = globalAddress.Ipv6Address
    return leafs
}

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetYangName() string { return "global-address" }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) SetParent(parent types.Entity) { globalAddress.parent = parent }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetParent() types.Entity { return globalAddress.parent }

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetParentYangName() string { return "group" }

// Hsrp_Ipv4_Groups_Group_StateChangeHistory
// State change history
type Hsrp_Ipv4_Groups_Group_StateChangeHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Old State. The type is StandbyGrpState.
    OldState interface{}

    // New State. The type is StandbyGrpState.
    NewState interface{}

    // Reason for state change. The type is HsrpStateChangeReason.
    Reason interface{}

    // Time of state change.
    Time Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time
}

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetFilter() yfilter.YFilter { return stateChangeHistory.YFilter }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) SetFilter(yf yfilter.YFilter) { stateChangeHistory.YFilter = yf }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetGoName(yname string) string {
    if yname == "old-state" { return "OldState" }
    if yname == "new-state" { return "NewState" }
    if yname == "reason" { return "Reason" }
    if yname == "time" { return "Time" }
    return ""
}

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetSegmentPath() string {
    return "state-change-history"
}

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "time" {
        return &stateChangeHistory.Time
    }
    return nil
}

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["time"] = &stateChangeHistory.Time
    return children
}

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["old-state"] = stateChangeHistory.OldState
    leafs["new-state"] = stateChangeHistory.NewState
    leafs["reason"] = stateChangeHistory.Reason
    return leafs
}

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetBundleName() string { return "cisco_ios_xr" }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetYangName() string { return "state-change-history" }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) SetParent(parent types.Entity) { stateChangeHistory.parent = parent }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetParent() types.Entity { return stateChangeHistory.parent }

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetParentYangName() string { return "group" }

// Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time
// Time of state change
type Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetFilter() yfilter.YFilter { return time.YFilter }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) SetFilter(yf yfilter.YFilter) { time.YFilter = yf }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetSegmentPath() string {
    return "time"
}

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = time.Seconds
    leafs["nanoseconds"] = time.Nanoseconds
    return leafs
}

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetBundleName() string { return "cisco_ios_xr" }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetYangName() string { return "time" }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) SetParent(parent types.Entity) { time.parent = parent }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetParent() types.Entity { return time.parent }

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetParentYangName() string { return "state-change-history" }

// Hsrp_Ipv4_TrackedInterfaces
// The HSRP tracked interfaces table
type Hsrp_Ipv4_TrackedInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An HSRP tracked interface entry. The type is slice of
    // Hsrp_Ipv4_TrackedInterfaces_TrackedInterface.
    TrackedInterface []Hsrp_Ipv4_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetFilter() yfilter.YFilter { return trackedInterfaces.YFilter }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) SetFilter(yf yfilter.YFilter) { trackedInterfaces.YFilter = yf }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetGoName(yname string) string {
    if yname == "tracked-interface" { return "TrackedInterface" }
    return ""
}

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetSegmentPath() string {
    return "tracked-interfaces"
}

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-interface" {
        for _, c := range trackedInterfaces.TrackedInterface {
            if trackedInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv4_TrackedInterfaces_TrackedInterface{}
        trackedInterfaces.TrackedInterface = append(trackedInterfaces.TrackedInterface, child)
        return &trackedInterfaces.TrackedInterface[len(trackedInterfaces.TrackedInterface)-1]
    }
    return nil
}

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedInterfaces.TrackedInterface {
        children[trackedInterfaces.TrackedInterface[i].GetSegmentPath()] = &trackedInterfaces.TrackedInterface[i]
    }
    return children
}

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetYangName() string { return "tracked-interfaces" }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) SetParent(parent types.Entity) { trackedInterfaces.parent = parent }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetParent() types.Entity { return trackedInterfaces.parent }

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetParentYangName() string { return "ipv4" }

// Hsrp_Ipv4_TrackedInterfaces_TrackedInterface
// An HSRP tracked interface entry
type Hsrp_Ipv4_TrackedInterfaces_TrackedInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name of the interface. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The HSRP group number. The type is interface{}
    // with range: -2147483648..2147483647.
    GroupNumber interface{}

    // This attribute is a key. The interface name of the interface being tracked.
    // The type is string with pattern: [a-zA-Z0-9./-]+.
    TrackedInterfaceName interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // HSRP Group number. The type is interface{} with range: 0..4294967295.
    HsrpGroupNumber interface{}

    // Priority weighting. The type is interface{} with range: 0..4294967295.
    PriorityDecrement interface{}

    // Interface up flag. The type is bool.
    InterfaceUpFlag interface{}

    // Tracked Interface Name. The type is string with length: 0..64.
    TrackedInterfaceNameXr interface{}

    // Tracked Object Flag. The type is bool.
    IsObject interface{}
}

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetFilter() yfilter.YFilter { return trackedInterface.YFilter }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) SetFilter(yf yfilter.YFilter) { trackedInterface.YFilter = yf }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "group-number" { return "GroupNumber" }
    if yname == "tracked-interface-name" { return "TrackedInterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "hsrp-group-number" { return "HsrpGroupNumber" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    if yname == "interface-up-flag" { return "InterfaceUpFlag" }
    if yname == "tracked-interface-name-xr" { return "TrackedInterfaceNameXr" }
    if yname == "is-object" { return "IsObject" }
    return ""
}

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetSegmentPath() string {
    return "tracked-interface" + "[interface-name='" + fmt.Sprintf("%v", trackedInterface.InterfaceName) + "']" + "[group-number='" + fmt.Sprintf("%v", trackedInterface.GroupNumber) + "']" + "[tracked-interface-name='" + fmt.Sprintf("%v", trackedInterface.TrackedInterfaceName) + "']"
}

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = trackedInterface.InterfaceName
    leafs["group-number"] = trackedInterface.GroupNumber
    leafs["tracked-interface-name"] = trackedInterface.TrackedInterfaceName
    leafs["interface"] = trackedInterface.Interface
    leafs["hsrp-group-number"] = trackedInterface.HsrpGroupNumber
    leafs["priority-decrement"] = trackedInterface.PriorityDecrement
    leafs["interface-up-flag"] = trackedInterface.InterfaceUpFlag
    leafs["tracked-interface-name-xr"] = trackedInterface.TrackedInterfaceNameXr
    leafs["is-object"] = trackedInterface.IsObject
    return leafs
}

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetYangName() string { return "tracked-interface" }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) SetParent(parent types.Entity) { trackedInterface.parent = parent }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetParent() types.Entity { return trackedInterface.parent }

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetParentYangName() string { return "tracked-interfaces" }

// Hsrp_Ipv4_Interfaces
// The HSRP interface information table
type Hsrp_Ipv4_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A HSRP interface entry. The type is slice of
    // Hsrp_Ipv4_Interfaces_Interface.
    Interface []Hsrp_Ipv4_Interfaces_Interface
}

func (interfaces *Hsrp_Ipv4_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Hsrp_Ipv4_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Hsrp_Ipv4_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Hsrp_Ipv4_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Hsrp_Ipv4_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv4_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Hsrp_Ipv4_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Hsrp_Ipv4_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Hsrp_Ipv4_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Hsrp_Ipv4_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Hsrp_Ipv4_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Hsrp_Ipv4_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Hsrp_Ipv4_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Hsrp_Ipv4_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Hsrp_Ipv4_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Hsrp_Ipv4_Interfaces) GetParentYangName() string { return "ipv4" }

// Hsrp_Ipv4_Interfaces_Interface
// A HSRP interface entry
type Hsrp_Ipv4_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Use burnt in mac address flag. The type is bool.
    UseBiaFlag interface{}

    // HSRP Interface Statistics.
    Statistics Hsrp_Ipv4_Interfaces_Interface_Statistics
}

func (self *Hsrp_Ipv4_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Hsrp_Ipv4_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Hsrp_Ipv4_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "use-bia-flag" { return "UseBiaFlag" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (self *Hsrp_Ipv4_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Hsrp_Ipv4_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &self.Statistics
    }
    return nil
}

func (self *Hsrp_Ipv4_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &self.Statistics
    return children
}

func (self *Hsrp_Ipv4_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["use-bia-flag"] = self.UseBiaFlag
    return leafs
}

func (self *Hsrp_Ipv4_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Hsrp_Ipv4_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Hsrp_Ipv4_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Hsrp_Ipv4_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Hsrp_Ipv4_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Hsrp_Ipv4_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Hsrp_Ipv4_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Hsrp_Ipv4_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Hsrp_Ipv4_Interfaces_Interface_Statistics
// HSRP Interface Statistics
type Hsrp_Ipv4_Interfaces_Interface_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of advertisement packets sent. The type is interface{} with range:
    // 0..4294967295.
    AdvertPacketsSent interface{}

    // Number of advertisement packets received. The type is interface{} with
    // range: 0..4294967295.
    AdvertPacketsReceived interface{}

    // Number of packets received that were too Long. The type is interface{} with
    // range: 0..4294967295.
    LongPacketsReceived interface{}

    // Number of packets received that were too short. The type is interface{}
    // with range: 0..4294967295.
    ShortPacketsReceived interface{}

    // Number of packets received with invalid version. The type is interface{}
    // with range: 0..4294967295.
    InvalidVersionReceived interface{}

    // Number of packets received with invalid operation code. The type is
    // interface{} with range: 0..4294967295.
    InvalidOperationCodeReceived interface{}

    // Number of packets received for an unknown group id. The type is interface{}
    // with range: 0..4294967295.
    UnknownGroupReceived interface{}

    // Number of packets received for an inoperational group. The type is
    // interface{} with range: 0..4294967295.
    InoperationalGroupReceived interface{}

    // Number of packets received from a conflicting Source IP address. The type
    // is interface{} with range: 0..4294967295.
    ConflictSourceIpAddressReceived interface{}
}

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetGoName(yname string) string {
    if yname == "advert-packets-sent" { return "AdvertPacketsSent" }
    if yname == "advert-packets-received" { return "AdvertPacketsReceived" }
    if yname == "long-packets-received" { return "LongPacketsReceived" }
    if yname == "short-packets-received" { return "ShortPacketsReceived" }
    if yname == "invalid-version-received" { return "InvalidVersionReceived" }
    if yname == "invalid-operation-code-received" { return "InvalidOperationCodeReceived" }
    if yname == "unknown-group-received" { return "UnknownGroupReceived" }
    if yname == "inoperational-group-received" { return "InoperationalGroupReceived" }
    if yname == "conflict-source-ip-address-received" { return "ConflictSourceIpAddressReceived" }
    return ""
}

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["advert-packets-sent"] = statistics.AdvertPacketsSent
    leafs["advert-packets-received"] = statistics.AdvertPacketsReceived
    leafs["long-packets-received"] = statistics.LongPacketsReceived
    leafs["short-packets-received"] = statistics.ShortPacketsReceived
    leafs["invalid-version-received"] = statistics.InvalidVersionReceived
    leafs["invalid-operation-code-received"] = statistics.InvalidOperationCodeReceived
    leafs["unknown-group-received"] = statistics.UnknownGroupReceived
    leafs["inoperational-group-received"] = statistics.InoperationalGroupReceived
    leafs["conflict-source-ip-address-received"] = statistics.ConflictSourceIpAddressReceived
    return leafs
}

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetYangName() string { return "statistics" }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetParentYangName() string { return "interface" }

// Hsrp_MgoSessions
// HSRP MGO session table
type Hsrp_MgoSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // HSRP MGO session. The type is slice of Hsrp_MgoSessions_MgoSession.
    MgoSession []Hsrp_MgoSessions_MgoSession
}

func (mgoSessions *Hsrp_MgoSessions) GetFilter() yfilter.YFilter { return mgoSessions.YFilter }

func (mgoSessions *Hsrp_MgoSessions) SetFilter(yf yfilter.YFilter) { mgoSessions.YFilter = yf }

func (mgoSessions *Hsrp_MgoSessions) GetGoName(yname string) string {
    if yname == "mgo-session" { return "MgoSession" }
    return ""
}

func (mgoSessions *Hsrp_MgoSessions) GetSegmentPath() string {
    return "mgo-sessions"
}

func (mgoSessions *Hsrp_MgoSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mgo-session" {
        for _, c := range mgoSessions.MgoSession {
            if mgoSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_MgoSessions_MgoSession{}
        mgoSessions.MgoSession = append(mgoSessions.MgoSession, child)
        return &mgoSessions.MgoSession[len(mgoSessions.MgoSession)-1]
    }
    return nil
}

func (mgoSessions *Hsrp_MgoSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mgoSessions.MgoSession {
        children[mgoSessions.MgoSession[i].GetSegmentPath()] = &mgoSessions.MgoSession[i]
    }
    return children
}

func (mgoSessions *Hsrp_MgoSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mgoSessions *Hsrp_MgoSessions) GetBundleName() string { return "cisco_ios_xr" }

func (mgoSessions *Hsrp_MgoSessions) GetYangName() string { return "mgo-sessions" }

func (mgoSessions *Hsrp_MgoSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mgoSessions *Hsrp_MgoSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mgoSessions *Hsrp_MgoSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mgoSessions *Hsrp_MgoSessions) SetParent(parent types.Entity) { mgoSessions.parent = parent }

func (mgoSessions *Hsrp_MgoSessions) GetParent() types.Entity { return mgoSessions.parent }

func (mgoSessions *Hsrp_MgoSessions) GetParentYangName() string { return "hsrp" }

// Hsrp_MgoSessions_MgoSession
// HSRP MGO session
type Hsrp_MgoSessions_MgoSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP MGO session name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    SessionName interface{}

    // Session Name. The type is string with length: 0..16.
    PrimarySessionName interface{}

    // Interface of primary session. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    PrimarySessionInterface interface{}

    // Address family of primary session. The type is HsrpBAf.
    PrimaryAfName interface{}

    // Group number of primary session. The type is interface{} with range:
    // 0..4294967295.
    PrimarySessionNumber interface{}

    // State of primary session. The type is StandbyGrpState.
    PrimarySessionState interface{}

    // List of slaves following this primary session. The type is slice of
    // Hsrp_MgoSessions_MgoSession_Slave.
    Slave []Hsrp_MgoSessions_MgoSession_Slave
}

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetFilter() yfilter.YFilter { return mgoSession.YFilter }

func (mgoSession *Hsrp_MgoSessions_MgoSession) SetFilter(yf yfilter.YFilter) { mgoSession.YFilter = yf }

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetGoName(yname string) string {
    if yname == "session-name" { return "SessionName" }
    if yname == "primary-session-name" { return "PrimarySessionName" }
    if yname == "primary-session-interface" { return "PrimarySessionInterface" }
    if yname == "primary-af-name" { return "PrimaryAfName" }
    if yname == "primary-session-number" { return "PrimarySessionNumber" }
    if yname == "primary-session-state" { return "PrimarySessionState" }
    if yname == "slave" { return "Slave" }
    return ""
}

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetSegmentPath() string {
    return "mgo-session" + "[session-name='" + fmt.Sprintf("%v", mgoSession.SessionName) + "']"
}

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slave" {
        for _, c := range mgoSession.Slave {
            if mgoSession.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_MgoSessions_MgoSession_Slave{}
        mgoSession.Slave = append(mgoSession.Slave, child)
        return &mgoSession.Slave[len(mgoSession.Slave)-1]
    }
    return nil
}

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mgoSession.Slave {
        children[mgoSession.Slave[i].GetSegmentPath()] = &mgoSession.Slave[i]
    }
    return children
}

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-name"] = mgoSession.SessionName
    leafs["primary-session-name"] = mgoSession.PrimarySessionName
    leafs["primary-session-interface"] = mgoSession.PrimarySessionInterface
    leafs["primary-af-name"] = mgoSession.PrimaryAfName
    leafs["primary-session-number"] = mgoSession.PrimarySessionNumber
    leafs["primary-session-state"] = mgoSession.PrimarySessionState
    return leafs
}

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetBundleName() string { return "cisco_ios_xr" }

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetYangName() string { return "mgo-session" }

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mgoSession *Hsrp_MgoSessions_MgoSession) SetParent(parent types.Entity) { mgoSession.parent = parent }

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetParent() types.Entity { return mgoSession.parent }

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetParentYangName() string { return "mgo-sessions" }

// Hsrp_MgoSessions_MgoSession_Slave
// List of slaves following this primary session
type Hsrp_MgoSessions_MgoSession_Slave struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface of slave group. The type is string with length: 0..64.
    SlaveGroupInterface interface{}

    // Group number of slave group. The type is interface{} with range:
    // 0..4294967295.
    SlaveGroupNumber interface{}
}

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetFilter() yfilter.YFilter { return slave.YFilter }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) SetFilter(yf yfilter.YFilter) { slave.YFilter = yf }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetGoName(yname string) string {
    if yname == "slave-group-interface" { return "SlaveGroupInterface" }
    if yname == "slave-group-number" { return "SlaveGroupNumber" }
    return ""
}

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetSegmentPath() string {
    return "slave"
}

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slave-group-interface"] = slave.SlaveGroupInterface
    leafs["slave-group-number"] = slave.SlaveGroupNumber
    return leafs
}

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetBundleName() string { return "cisco_ios_xr" }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetYangName() string { return "slave" }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) SetParent(parent types.Entity) { slave.parent = parent }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetParent() types.Entity { return slave.parent }

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetParentYangName() string { return "mgo-session" }

// Hsrp_Ipv6
// IPv6 HSRP information
type Hsrp_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP tracked interfaces table.
    TrackedInterfaces Hsrp_Ipv6_TrackedInterfaces

    // The HSRP standby group table.
    Groups Hsrp_Ipv6_Groups

    // The HSRP interface information table.
    Interfaces Hsrp_Ipv6_Interfaces
}

func (ipv6 *Hsrp_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Hsrp_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Hsrp_Ipv6) GetGoName(yname string) string {
    if yname == "tracked-interfaces" { return "TrackedInterfaces" }
    if yname == "groups" { return "Groups" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv6 *Hsrp_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Hsrp_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-interfaces" {
        return &ipv6.TrackedInterfaces
    }
    if childYangName == "groups" {
        return &ipv6.Groups
    }
    if childYangName == "interfaces" {
        return &ipv6.Interfaces
    }
    return nil
}

func (ipv6 *Hsrp_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tracked-interfaces"] = &ipv6.TrackedInterfaces
    children["groups"] = &ipv6.Groups
    children["interfaces"] = &ipv6.Interfaces
    return children
}

func (ipv6 *Hsrp_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Hsrp_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Hsrp_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Hsrp_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Hsrp_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Hsrp_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Hsrp_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Hsrp_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Hsrp_Ipv6) GetParentYangName() string { return "hsrp" }

// Hsrp_Ipv6_TrackedInterfaces
// The HSRP tracked interfaces table
type Hsrp_Ipv6_TrackedInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An HSRP tracked interface entry. The type is slice of
    // Hsrp_Ipv6_TrackedInterfaces_TrackedInterface.
    TrackedInterface []Hsrp_Ipv6_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetFilter() yfilter.YFilter { return trackedInterfaces.YFilter }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) SetFilter(yf yfilter.YFilter) { trackedInterfaces.YFilter = yf }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetGoName(yname string) string {
    if yname == "tracked-interface" { return "TrackedInterface" }
    return ""
}

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetSegmentPath() string {
    return "tracked-interfaces"
}

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-interface" {
        for _, c := range trackedInterfaces.TrackedInterface {
            if trackedInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv6_TrackedInterfaces_TrackedInterface{}
        trackedInterfaces.TrackedInterface = append(trackedInterfaces.TrackedInterface, child)
        return &trackedInterfaces.TrackedInterface[len(trackedInterfaces.TrackedInterface)-1]
    }
    return nil
}

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedInterfaces.TrackedInterface {
        children[trackedInterfaces.TrackedInterface[i].GetSegmentPath()] = &trackedInterfaces.TrackedInterface[i]
    }
    return children
}

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetYangName() string { return "tracked-interfaces" }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) SetParent(parent types.Entity) { trackedInterfaces.parent = parent }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetParent() types.Entity { return trackedInterfaces.parent }

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetParentYangName() string { return "ipv6" }

// Hsrp_Ipv6_TrackedInterfaces_TrackedInterface
// An HSRP tracked interface entry
type Hsrp_Ipv6_TrackedInterfaces_TrackedInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name of the interface. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The HSRP group number. The type is interface{}
    // with range: -2147483648..2147483647.
    GroupNumber interface{}

    // This attribute is a key. The interface name of the interface being tracked.
    // The type is string with pattern: [a-zA-Z0-9./-]+.
    TrackedInterfaceName interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // HSRP Group number. The type is interface{} with range: 0..4294967295.
    HsrpGroupNumber interface{}

    // Priority weighting. The type is interface{} with range: 0..4294967295.
    PriorityDecrement interface{}

    // Interface up flag. The type is bool.
    InterfaceUpFlag interface{}

    // Tracked Interface Name. The type is string with length: 0..64.
    TrackedInterfaceNameXr interface{}

    // Tracked Object Flag. The type is bool.
    IsObject interface{}
}

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetFilter() yfilter.YFilter { return trackedInterface.YFilter }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) SetFilter(yf yfilter.YFilter) { trackedInterface.YFilter = yf }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "group-number" { return "GroupNumber" }
    if yname == "tracked-interface-name" { return "TrackedInterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "hsrp-group-number" { return "HsrpGroupNumber" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    if yname == "interface-up-flag" { return "InterfaceUpFlag" }
    if yname == "tracked-interface-name-xr" { return "TrackedInterfaceNameXr" }
    if yname == "is-object" { return "IsObject" }
    return ""
}

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetSegmentPath() string {
    return "tracked-interface" + "[interface-name='" + fmt.Sprintf("%v", trackedInterface.InterfaceName) + "']" + "[group-number='" + fmt.Sprintf("%v", trackedInterface.GroupNumber) + "']" + "[tracked-interface-name='" + fmt.Sprintf("%v", trackedInterface.TrackedInterfaceName) + "']"
}

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = trackedInterface.InterfaceName
    leafs["group-number"] = trackedInterface.GroupNumber
    leafs["tracked-interface-name"] = trackedInterface.TrackedInterfaceName
    leafs["interface"] = trackedInterface.Interface
    leafs["hsrp-group-number"] = trackedInterface.HsrpGroupNumber
    leafs["priority-decrement"] = trackedInterface.PriorityDecrement
    leafs["interface-up-flag"] = trackedInterface.InterfaceUpFlag
    leafs["tracked-interface-name-xr"] = trackedInterface.TrackedInterfaceNameXr
    leafs["is-object"] = trackedInterface.IsObject
    return leafs
}

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetYangName() string { return "tracked-interface" }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) SetParent(parent types.Entity) { trackedInterface.parent = parent }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetParent() types.Entity { return trackedInterface.parent }

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetParentYangName() string { return "tracked-interfaces" }

// Hsrp_Ipv6_Groups
// The HSRP standby group table
type Hsrp_Ipv6_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An HSRP standby group. The type is slice of Hsrp_Ipv6_Groups_Group.
    Group []Hsrp_Ipv6_Groups_Group
}

func (groups *Hsrp_Ipv6_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Hsrp_Ipv6_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Hsrp_Ipv6_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Hsrp_Ipv6_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Hsrp_Ipv6_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv6_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Hsrp_Ipv6_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Hsrp_Ipv6_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Hsrp_Ipv6_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Hsrp_Ipv6_Groups) GetYangName() string { return "groups" }

func (groups *Hsrp_Ipv6_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Hsrp_Ipv6_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Hsrp_Ipv6_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Hsrp_Ipv6_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Hsrp_Ipv6_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Hsrp_Ipv6_Groups) GetParentYangName() string { return "ipv6" }

// Hsrp_Ipv6_Groups_Group
// An HSRP standby group
type Hsrp_Ipv6_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The HSRP group number. The type is interface{}
    // with range: -2147483648..2147483647.
    GroupNumber interface{}

    // Authentication string. The type is string with length: 0..9.
    AuthenticationString interface{}

    // Virtual mac address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // HSRP Group number. The type is interface{} with range: 0..4294967295.
    HsrpGroupNumber interface{}

    // Address family. The type is HsrpBAf.
    AddressFamily interface{}

    // HSRP Protocol Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Session Name. The type is string with length: 0..16.
    SessionName interface{}

    // Number of slaves following state. The type is interface{} with range:
    // 0..4294967295.
    Slaves interface{}

    // Group is a slave group. The type is bool.
    IsSlave interface{}

    // Followed Session Name. The type is string with length: 0..16.
    FollowedSessionName interface{}

    // Configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // Preempt delay time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    PreemptDelay interface{}

    // Preempt time remaining in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    PreemptTimerSecs interface{}

    // Hellotime in msecs. The type is interface{} with range: 0..4294967295.
    // Units are millisecond.
    HelloTime interface{}

    // Holdtime in msecs. The type is interface{} with range: 0..4294967295. Units
    // are millisecond.
    HoldTime interface{}

    // Learned hellotime in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    LearnedHelloTime interface{}

    // Learned holdtime in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    LearnedHoldTime interface{}

    // Minimum delay time in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    MinDelayTime interface{}

    // Reload delay time in msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    ReloadDelayTime interface{}

    // Configured Virtual IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    VirtualIpAddress interface{}

    // Virtual linklocal IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualLinklocalIpv6Address interface{}

    // Active router's IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ActiveIpAddress interface{}

    // Active router's IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ActiveIpv6Address interface{}

    // Active router's interface MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    ActiveMacAddress interface{}

    // Standby router's IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StandbyIpAddress interface{}

    // Standby router's IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StandbyIpv6Address interface{}

    // Standby router's interface MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    StandbyMacAddress interface{}

    // HSRP router state. The type is StandbyGrpState.
    HsrpRouterState interface{}

    // Interface Name. The type is string with length: 0..64.
    InterfaceNameXr interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Priority of the router. The type is interface{} with range: 0..255.
    RouterPriority interface{}

    // Priority of the Active router. The type is interface{} with range: 0..255.
    ActivePriority interface{}

    // Active timer running flag. The type is bool.
    ActiveTimerFlag interface{}

    // Active timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    ActiveTimerSecs interface{}

    // Active timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    ActiveTimerMsecs interface{}

    // Standby timer running flag. The type is bool.
    StandbyTimerFlag interface{}

    // Standby timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    StandbyTimerSecs interface{}

    // Standby timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    StandbyTimerMsecs interface{}

    // Hello timer running flag. The type is bool.
    HelloTimerFlag interface{}

    // Hello timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    HelloTimerSecs interface{}

    // Hello timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    HelloTimerMsecs interface{}

    // Delay timer running flag. The type is bool.
    DelayTimerFlag interface{}

    // Delay timer running time secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    DelayTimerSecs interface{}

    // Delay timer running time msecs. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    DelayTimerMsecs interface{}

    // Time in current state secs. The type is interface{} with range:
    // 0..4294967295. Units are second.
    CurrentStateTimerSecs interface{}

    // Number of state changes. The type is interface{} with range: 0..4294967295.
    StateChangeCount interface{}

    // Number of tracked interfaces. The type is interface{} with range:
    // 0..4294967295.
    TrackedInterfaceCount interface{}

    // Number of tracked interfaces up. The type is interface{} with range:
    // 0..4294967295.
    TrackedInterfaceUpCount interface{}

    // Preempt enabled. The type is bool.
    PreemptEnabled interface{}

    // Use configured timers. The type is bool.
    UseConfiguredTimers interface{}

    // Use configured virtual IP. The type is bool.
    UseConfiguredVirtualIp interface{}

    // Use burnt in MAC address configured. The type is bool.
    UseBiaConfigured interface{}

    // Non-default timers are configured. The type is bool.
    ConfiguredTimers interface{}

    // MAC address configured. The type is bool.
    ConfiguredMacAddress interface{}

    // HSRP redirects disabled. The type is bool.
    RedirectsDisabled interface{}

    // HSRP BFD fast failover. The type is bool.
    BfdEnabled interface{}

    // BFD Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    BfdInterface interface{}

    // BFD Peer IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BfdPeerIpAddress interface{}

    // BFD Peer IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    BfdPeerIpv6Address interface{}

    // BFD session state. The type is HsrpBfdSessionState.
    BfdSessionState interface{}

    // BFD packet send interval. The type is interface{} with range:
    // 0..4294967295.
    BfdInterval interface{}

    // BFD multiplier. The type is interface{} with range: 0..4294967295.
    BfdMultiplier interface{}

    // Virtual mac address state. The type is HsrpVmacState.
    VirtualMacAddressState interface{}

    // Secondary virtual IP addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SecondaryAddress []interface{}

    // Time last resign was sent.
    ResignSentTime Hsrp_Ipv6_Groups_Group_ResignSentTime

    // Time last resign was received.
    ResignReceivedTime Hsrp_Ipv6_Groups_Group_ResignReceivedTime

    // Time last coup was sent.
    CoupSentTime Hsrp_Ipv6_Groups_Group_CoupSentTime

    // Time last coup was received.
    CoupReceivedTime Hsrp_Ipv6_Groups_Group_CoupReceivedTime

    // HSRP Group statistics.
    Statistics Hsrp_Ipv6_Groups_Group_Statistics

    // Global virtual IPv6 addresses. The type is slice of
    // Hsrp_Ipv6_Groups_Group_GlobalAddress.
    GlobalAddress []Hsrp_Ipv6_Groups_Group_GlobalAddress

    // State change history. The type is slice of
    // Hsrp_Ipv6_Groups_Group_StateChangeHistory.
    StateChangeHistory []Hsrp_Ipv6_Groups_Group_StateChangeHistory
}

func (group *Hsrp_Ipv6_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Hsrp_Ipv6_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Hsrp_Ipv6_Groups_Group) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "group-number" { return "GroupNumber" }
    if yname == "authentication-string" { return "AuthenticationString" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "hsrp-group-number" { return "HsrpGroupNumber" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "version" { return "Version" }
    if yname == "session-name" { return "SessionName" }
    if yname == "slaves" { return "Slaves" }
    if yname == "is-slave" { return "IsSlave" }
    if yname == "followed-session-name" { return "FollowedSessionName" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "preempt-timer-secs" { return "PreemptTimerSecs" }
    if yname == "hello-time" { return "HelloTime" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "learned-hello-time" { return "LearnedHelloTime" }
    if yname == "learned-hold-time" { return "LearnedHoldTime" }
    if yname == "min-delay-time" { return "MinDelayTime" }
    if yname == "reload-delay-time" { return "ReloadDelayTime" }
    if yname == "virtual-ip-address" { return "VirtualIpAddress" }
    if yname == "virtual-linklocal-ipv6-address" { return "VirtualLinklocalIpv6Address" }
    if yname == "active-ip-address" { return "ActiveIpAddress" }
    if yname == "active-ipv6-address" { return "ActiveIpv6Address" }
    if yname == "active-mac-address" { return "ActiveMacAddress" }
    if yname == "standby-ip-address" { return "StandbyIpAddress" }
    if yname == "standby-ipv6-address" { return "StandbyIpv6Address" }
    if yname == "standby-mac-address" { return "StandbyMacAddress" }
    if yname == "hsrp-router-state" { return "HsrpRouterState" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface" { return "Interface" }
    if yname == "router-priority" { return "RouterPriority" }
    if yname == "active-priority" { return "ActivePriority" }
    if yname == "active-timer-flag" { return "ActiveTimerFlag" }
    if yname == "active-timer-secs" { return "ActiveTimerSecs" }
    if yname == "active-timer-msecs" { return "ActiveTimerMsecs" }
    if yname == "standby-timer-flag" { return "StandbyTimerFlag" }
    if yname == "standby-timer-secs" { return "StandbyTimerSecs" }
    if yname == "standby-timer-msecs" { return "StandbyTimerMsecs" }
    if yname == "hello-timer-flag" { return "HelloTimerFlag" }
    if yname == "hello-timer-secs" { return "HelloTimerSecs" }
    if yname == "hello-timer-msecs" { return "HelloTimerMsecs" }
    if yname == "delay-timer-flag" { return "DelayTimerFlag" }
    if yname == "delay-timer-secs" { return "DelayTimerSecs" }
    if yname == "delay-timer-msecs" { return "DelayTimerMsecs" }
    if yname == "current-state-timer-secs" { return "CurrentStateTimerSecs" }
    if yname == "state-change-count" { return "StateChangeCount" }
    if yname == "tracked-interface-count" { return "TrackedInterfaceCount" }
    if yname == "tracked-interface-up-count" { return "TrackedInterfaceUpCount" }
    if yname == "preempt-enabled" { return "PreemptEnabled" }
    if yname == "use-configured-timers" { return "UseConfiguredTimers" }
    if yname == "use-configured-virtual-ip" { return "UseConfiguredVirtualIp" }
    if yname == "use-bia-configured" { return "UseBiaConfigured" }
    if yname == "configured-timers" { return "ConfiguredTimers" }
    if yname == "configured-mac-address" { return "ConfiguredMacAddress" }
    if yname == "redirects-disabled" { return "RedirectsDisabled" }
    if yname == "bfd-enabled" { return "BfdEnabled" }
    if yname == "bfd-interface" { return "BfdInterface" }
    if yname == "bfd-peer-ip-address" { return "BfdPeerIpAddress" }
    if yname == "bfd-peer-ipv6-address" { return "BfdPeerIpv6Address" }
    if yname == "bfd-session-state" { return "BfdSessionState" }
    if yname == "bfd-interval" { return "BfdInterval" }
    if yname == "bfd-multiplier" { return "BfdMultiplier" }
    if yname == "virtual-mac-address-state" { return "VirtualMacAddressState" }
    if yname == "secondary-address" { return "SecondaryAddress" }
    if yname == "resign-sent-time" { return "ResignSentTime" }
    if yname == "resign-received-time" { return "ResignReceivedTime" }
    if yname == "coup-sent-time" { return "CoupSentTime" }
    if yname == "coup-received-time" { return "CoupReceivedTime" }
    if yname == "statistics" { return "Statistics" }
    if yname == "global-address" { return "GlobalAddress" }
    if yname == "state-change-history" { return "StateChangeHistory" }
    return ""
}

func (group *Hsrp_Ipv6_Groups_Group) GetSegmentPath() string {
    return "group" + "[interface-name='" + fmt.Sprintf("%v", group.InterfaceName) + "']" + "[group-number='" + fmt.Sprintf("%v", group.GroupNumber) + "']"
}

func (group *Hsrp_Ipv6_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "resign-sent-time" {
        return &group.ResignSentTime
    }
    if childYangName == "resign-received-time" {
        return &group.ResignReceivedTime
    }
    if childYangName == "coup-sent-time" {
        return &group.CoupSentTime
    }
    if childYangName == "coup-received-time" {
        return &group.CoupReceivedTime
    }
    if childYangName == "statistics" {
        return &group.Statistics
    }
    if childYangName == "global-address" {
        for _, c := range group.GlobalAddress {
            if group.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv6_Groups_Group_GlobalAddress{}
        group.GlobalAddress = append(group.GlobalAddress, child)
        return &group.GlobalAddress[len(group.GlobalAddress)-1]
    }
    if childYangName == "state-change-history" {
        for _, c := range group.StateChangeHistory {
            if group.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv6_Groups_Group_StateChangeHistory{}
        group.StateChangeHistory = append(group.StateChangeHistory, child)
        return &group.StateChangeHistory[len(group.StateChangeHistory)-1]
    }
    return nil
}

func (group *Hsrp_Ipv6_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["resign-sent-time"] = &group.ResignSentTime
    children["resign-received-time"] = &group.ResignReceivedTime
    children["coup-sent-time"] = &group.CoupSentTime
    children["coup-received-time"] = &group.CoupReceivedTime
    children["statistics"] = &group.Statistics
    for i := range group.GlobalAddress {
        children[group.GlobalAddress[i].GetSegmentPath()] = &group.GlobalAddress[i]
    }
    for i := range group.StateChangeHistory {
        children[group.StateChangeHistory[i].GetSegmentPath()] = &group.StateChangeHistory[i]
    }
    return children
}

func (group *Hsrp_Ipv6_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = group.InterfaceName
    leafs["group-number"] = group.GroupNumber
    leafs["authentication-string"] = group.AuthenticationString
    leafs["virtual-mac-address"] = group.VirtualMacAddress
    leafs["hsrp-group-number"] = group.HsrpGroupNumber
    leafs["address-family"] = group.AddressFamily
    leafs["version"] = group.Version
    leafs["session-name"] = group.SessionName
    leafs["slaves"] = group.Slaves
    leafs["is-slave"] = group.IsSlave
    leafs["followed-session-name"] = group.FollowedSessionName
    leafs["configured-priority"] = group.ConfiguredPriority
    leafs["preempt-delay"] = group.PreemptDelay
    leafs["preempt-timer-secs"] = group.PreemptTimerSecs
    leafs["hello-time"] = group.HelloTime
    leafs["hold-time"] = group.HoldTime
    leafs["learned-hello-time"] = group.LearnedHelloTime
    leafs["learned-hold-time"] = group.LearnedHoldTime
    leafs["min-delay-time"] = group.MinDelayTime
    leafs["reload-delay-time"] = group.ReloadDelayTime
    leafs["virtual-ip-address"] = group.VirtualIpAddress
    leafs["virtual-linklocal-ipv6-address"] = group.VirtualLinklocalIpv6Address
    leafs["active-ip-address"] = group.ActiveIpAddress
    leafs["active-ipv6-address"] = group.ActiveIpv6Address
    leafs["active-mac-address"] = group.ActiveMacAddress
    leafs["standby-ip-address"] = group.StandbyIpAddress
    leafs["standby-ipv6-address"] = group.StandbyIpv6Address
    leafs["standby-mac-address"] = group.StandbyMacAddress
    leafs["hsrp-router-state"] = group.HsrpRouterState
    leafs["interface-name-xr"] = group.InterfaceNameXr
    leafs["interface"] = group.Interface
    leafs["router-priority"] = group.RouterPriority
    leafs["active-priority"] = group.ActivePriority
    leafs["active-timer-flag"] = group.ActiveTimerFlag
    leafs["active-timer-secs"] = group.ActiveTimerSecs
    leafs["active-timer-msecs"] = group.ActiveTimerMsecs
    leafs["standby-timer-flag"] = group.StandbyTimerFlag
    leafs["standby-timer-secs"] = group.StandbyTimerSecs
    leafs["standby-timer-msecs"] = group.StandbyTimerMsecs
    leafs["hello-timer-flag"] = group.HelloTimerFlag
    leafs["hello-timer-secs"] = group.HelloTimerSecs
    leafs["hello-timer-msecs"] = group.HelloTimerMsecs
    leafs["delay-timer-flag"] = group.DelayTimerFlag
    leafs["delay-timer-secs"] = group.DelayTimerSecs
    leafs["delay-timer-msecs"] = group.DelayTimerMsecs
    leafs["current-state-timer-secs"] = group.CurrentStateTimerSecs
    leafs["state-change-count"] = group.StateChangeCount
    leafs["tracked-interface-count"] = group.TrackedInterfaceCount
    leafs["tracked-interface-up-count"] = group.TrackedInterfaceUpCount
    leafs["preempt-enabled"] = group.PreemptEnabled
    leafs["use-configured-timers"] = group.UseConfiguredTimers
    leafs["use-configured-virtual-ip"] = group.UseConfiguredVirtualIp
    leafs["use-bia-configured"] = group.UseBiaConfigured
    leafs["configured-timers"] = group.ConfiguredTimers
    leafs["configured-mac-address"] = group.ConfiguredMacAddress
    leafs["redirects-disabled"] = group.RedirectsDisabled
    leafs["bfd-enabled"] = group.BfdEnabled
    leafs["bfd-interface"] = group.BfdInterface
    leafs["bfd-peer-ip-address"] = group.BfdPeerIpAddress
    leafs["bfd-peer-ipv6-address"] = group.BfdPeerIpv6Address
    leafs["bfd-session-state"] = group.BfdSessionState
    leafs["bfd-interval"] = group.BfdInterval
    leafs["bfd-multiplier"] = group.BfdMultiplier
    leafs["virtual-mac-address-state"] = group.VirtualMacAddressState
    leafs["secondary-address"] = group.SecondaryAddress
    return leafs
}

func (group *Hsrp_Ipv6_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Hsrp_Ipv6_Groups_Group) GetYangName() string { return "group" }

func (group *Hsrp_Ipv6_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Hsrp_Ipv6_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Hsrp_Ipv6_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Hsrp_Ipv6_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Hsrp_Ipv6_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Hsrp_Ipv6_Groups_Group) GetParentYangName() string { return "groups" }

// Hsrp_Ipv6_Groups_Group_ResignSentTime
// Time last resign was sent
type Hsrp_Ipv6_Groups_Group_ResignSentTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetFilter() yfilter.YFilter { return resignSentTime.YFilter }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) SetFilter(yf yfilter.YFilter) { resignSentTime.YFilter = yf }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetSegmentPath() string {
    return "resign-sent-time"
}

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resignSentTime.Seconds
    leafs["nanoseconds"] = resignSentTime.Nanoseconds
    return leafs
}

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetBundleName() string { return "cisco_ios_xr" }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetYangName() string { return "resign-sent-time" }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) SetParent(parent types.Entity) { resignSentTime.parent = parent }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetParent() types.Entity { return resignSentTime.parent }

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetParentYangName() string { return "group" }

// Hsrp_Ipv6_Groups_Group_ResignReceivedTime
// Time last resign was received
type Hsrp_Ipv6_Groups_Group_ResignReceivedTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetFilter() yfilter.YFilter { return resignReceivedTime.YFilter }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) SetFilter(yf yfilter.YFilter) { resignReceivedTime.YFilter = yf }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetSegmentPath() string {
    return "resign-received-time"
}

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resignReceivedTime.Seconds
    leafs["nanoseconds"] = resignReceivedTime.Nanoseconds
    return leafs
}

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetBundleName() string { return "cisco_ios_xr" }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetYangName() string { return "resign-received-time" }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) SetParent(parent types.Entity) { resignReceivedTime.parent = parent }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetParent() types.Entity { return resignReceivedTime.parent }

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetParentYangName() string { return "group" }

// Hsrp_Ipv6_Groups_Group_CoupSentTime
// Time last coup was sent
type Hsrp_Ipv6_Groups_Group_CoupSentTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetFilter() yfilter.YFilter { return coupSentTime.YFilter }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) SetFilter(yf yfilter.YFilter) { coupSentTime.YFilter = yf }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetSegmentPath() string {
    return "coup-sent-time"
}

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = coupSentTime.Seconds
    leafs["nanoseconds"] = coupSentTime.Nanoseconds
    return leafs
}

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetBundleName() string { return "cisco_ios_xr" }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetYangName() string { return "coup-sent-time" }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) SetParent(parent types.Entity) { coupSentTime.parent = parent }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetParent() types.Entity { return coupSentTime.parent }

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetParentYangName() string { return "group" }

// Hsrp_Ipv6_Groups_Group_CoupReceivedTime
// Time last coup was received
type Hsrp_Ipv6_Groups_Group_CoupReceivedTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetFilter() yfilter.YFilter { return coupReceivedTime.YFilter }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) SetFilter(yf yfilter.YFilter) { coupReceivedTime.YFilter = yf }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetSegmentPath() string {
    return "coup-received-time"
}

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = coupReceivedTime.Seconds
    leafs["nanoseconds"] = coupReceivedTime.Nanoseconds
    return leafs
}

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetBundleName() string { return "cisco_ios_xr" }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetYangName() string { return "coup-received-time" }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) SetParent(parent types.Entity) { coupReceivedTime.parent = parent }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetParent() types.Entity { return coupReceivedTime.parent }

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetParentYangName() string { return "group" }

// Hsrp_Ipv6_Groups_Group_Statistics
// HSRP Group statistics
type Hsrp_Ipv6_Groups_Group_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of transitions to Active State. The type is interface{} with range:
    // 0..4294967295.
    ActiveTransitions interface{}

    // Number of transitions to Standby State. The type is interface{} with range:
    // 0..4294967295.
    StandbyTransitions interface{}

    // Number of transitions to Speak State. The type is interface{} with range:
    // 0..4294967295.
    SpeakTransitions interface{}

    // Number of transitions to Listen State. The type is interface{} with range:
    // 0..4294967295.
    ListenTransitions interface{}

    // Number of transitions to Learn State. The type is interface{} with range:
    // 0..4294967295.
    LearnTransitions interface{}

    // Number of transitions to Init State. The type is interface{} with range:
    // 0..4294967295.
    InitTransitions interface{}

    // Number of Hello Packets sent (NB: Bundles only). The type is interface{}
    // with range: 0..4294967295.
    HelloPacketsSent interface{}

    // Number of Resign Packets sent. The type is interface{} with range:
    // 0..4294967295.
    ResignPacketsSent interface{}

    // Number of Coup Packets sent. The type is interface{} with range:
    // 0..4294967295.
    CoupPacketsSent interface{}

    // Number of Hello Packets received. The type is interface{} with range:
    // 0..4294967295.
    HelloPacketsReceived interface{}

    // Number of Resign Packets received. The type is interface{} with range:
    // 0..4294967295.
    ResignPacketsReceived interface{}

    // Number of Coup Packets received. The type is interface{} with range:
    // 0..4294967295.
    CoupPacketsReceived interface{}

    // Number of Packets received that failed authentication. The type is
    // interface{} with range: 0..4294967295.
    AuthFailReceived interface{}

    // Number of packets received with invalid Hello Time value. The type is
    // interface{} with range: 0..4294967295.
    InvalidTimerReceived interface{}

    // Number of packets received with mismatching virtual IP address. The type is
    // interface{} with range: 0..4294967295.
    MismatchVirtualIpAddressReceived interface{}
}

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetGoName(yname string) string {
    if yname == "active-transitions" { return "ActiveTransitions" }
    if yname == "standby-transitions" { return "StandbyTransitions" }
    if yname == "speak-transitions" { return "SpeakTransitions" }
    if yname == "listen-transitions" { return "ListenTransitions" }
    if yname == "learn-transitions" { return "LearnTransitions" }
    if yname == "init-transitions" { return "InitTransitions" }
    if yname == "hello-packets-sent" { return "HelloPacketsSent" }
    if yname == "resign-packets-sent" { return "ResignPacketsSent" }
    if yname == "coup-packets-sent" { return "CoupPacketsSent" }
    if yname == "hello-packets-received" { return "HelloPacketsReceived" }
    if yname == "resign-packets-received" { return "ResignPacketsReceived" }
    if yname == "coup-packets-received" { return "CoupPacketsReceived" }
    if yname == "auth-fail-received" { return "AuthFailReceived" }
    if yname == "invalid-timer-received" { return "InvalidTimerReceived" }
    if yname == "mismatch-virtual-ip-address-received" { return "MismatchVirtualIpAddressReceived" }
    return ""
}

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-transitions"] = statistics.ActiveTransitions
    leafs["standby-transitions"] = statistics.StandbyTransitions
    leafs["speak-transitions"] = statistics.SpeakTransitions
    leafs["listen-transitions"] = statistics.ListenTransitions
    leafs["learn-transitions"] = statistics.LearnTransitions
    leafs["init-transitions"] = statistics.InitTransitions
    leafs["hello-packets-sent"] = statistics.HelloPacketsSent
    leafs["resign-packets-sent"] = statistics.ResignPacketsSent
    leafs["coup-packets-sent"] = statistics.CoupPacketsSent
    leafs["hello-packets-received"] = statistics.HelloPacketsReceived
    leafs["resign-packets-received"] = statistics.ResignPacketsReceived
    leafs["coup-packets-received"] = statistics.CoupPacketsReceived
    leafs["auth-fail-received"] = statistics.AuthFailReceived
    leafs["invalid-timer-received"] = statistics.InvalidTimerReceived
    leafs["mismatch-virtual-ip-address-received"] = statistics.MismatchVirtualIpAddressReceived
    return leafs
}

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetYangName() string { return "statistics" }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetParentYangName() string { return "group" }

// Hsrp_Ipv6_Groups_Group_GlobalAddress
// Global virtual IPv6 addresses
type Hsrp_Ipv6_Groups_Group_GlobalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetFilter() yfilter.YFilter { return globalAddress.YFilter }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) SetFilter(yf yfilter.YFilter) { globalAddress.YFilter = yf }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetSegmentPath() string {
    return "global-address"
}

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = globalAddress.Ipv6Address
    return leafs
}

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetYangName() string { return "global-address" }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) SetParent(parent types.Entity) { globalAddress.parent = parent }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetParent() types.Entity { return globalAddress.parent }

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetParentYangName() string { return "group" }

// Hsrp_Ipv6_Groups_Group_StateChangeHistory
// State change history
type Hsrp_Ipv6_Groups_Group_StateChangeHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Old State. The type is StandbyGrpState.
    OldState interface{}

    // New State. The type is StandbyGrpState.
    NewState interface{}

    // Reason for state change. The type is HsrpStateChangeReason.
    Reason interface{}

    // Time of state change.
    Time Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time
}

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetFilter() yfilter.YFilter { return stateChangeHistory.YFilter }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) SetFilter(yf yfilter.YFilter) { stateChangeHistory.YFilter = yf }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetGoName(yname string) string {
    if yname == "old-state" { return "OldState" }
    if yname == "new-state" { return "NewState" }
    if yname == "reason" { return "Reason" }
    if yname == "time" { return "Time" }
    return ""
}

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetSegmentPath() string {
    return "state-change-history"
}

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "time" {
        return &stateChangeHistory.Time
    }
    return nil
}

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["time"] = &stateChangeHistory.Time
    return children
}

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["old-state"] = stateChangeHistory.OldState
    leafs["new-state"] = stateChangeHistory.NewState
    leafs["reason"] = stateChangeHistory.Reason
    return leafs
}

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetBundleName() string { return "cisco_ios_xr" }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetYangName() string { return "state-change-history" }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) SetParent(parent types.Entity) { stateChangeHistory.parent = parent }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetParent() types.Entity { return stateChangeHistory.parent }

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetParentYangName() string { return "group" }

// Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time
// Time of state change
type Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetFilter() yfilter.YFilter { return time.YFilter }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) SetFilter(yf yfilter.YFilter) { time.YFilter = yf }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetSegmentPath() string {
    return "time"
}

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = time.Seconds
    leafs["nanoseconds"] = time.Nanoseconds
    return leafs
}

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetBundleName() string { return "cisco_ios_xr" }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetYangName() string { return "time" }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) SetParent(parent types.Entity) { time.parent = parent }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetParent() types.Entity { return time.parent }

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetParentYangName() string { return "state-change-history" }

// Hsrp_Ipv6_Interfaces
// The HSRP interface information table
type Hsrp_Ipv6_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A HSRP interface entry. The type is slice of
    // Hsrp_Ipv6_Interfaces_Interface.
    Interface []Hsrp_Ipv6_Interfaces_Interface
}

func (interfaces *Hsrp_Ipv6_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Hsrp_Ipv6_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Hsrp_Ipv6_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Hsrp_Ipv6_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Hsrp_Ipv6_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Ipv6_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Hsrp_Ipv6_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Hsrp_Ipv6_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Hsrp_Ipv6_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Hsrp_Ipv6_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Hsrp_Ipv6_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Hsrp_Ipv6_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Hsrp_Ipv6_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Hsrp_Ipv6_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Hsrp_Ipv6_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Hsrp_Ipv6_Interfaces) GetParentYangName() string { return "ipv6" }

// Hsrp_Ipv6_Interfaces_Interface
// A HSRP interface entry
type Hsrp_Ipv6_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // IM Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Use burnt in mac address flag. The type is bool.
    UseBiaFlag interface{}

    // HSRP Interface Statistics.
    Statistics Hsrp_Ipv6_Interfaces_Interface_Statistics
}

func (self *Hsrp_Ipv6_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Hsrp_Ipv6_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Hsrp_Ipv6_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "use-bia-flag" { return "UseBiaFlag" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (self *Hsrp_Ipv6_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Hsrp_Ipv6_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &self.Statistics
    }
    return nil
}

func (self *Hsrp_Ipv6_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &self.Statistics
    return children
}

func (self *Hsrp_Ipv6_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["use-bia-flag"] = self.UseBiaFlag
    return leafs
}

func (self *Hsrp_Ipv6_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Hsrp_Ipv6_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Hsrp_Ipv6_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Hsrp_Ipv6_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Hsrp_Ipv6_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Hsrp_Ipv6_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Hsrp_Ipv6_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Hsrp_Ipv6_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Hsrp_Ipv6_Interfaces_Interface_Statistics
// HSRP Interface Statistics
type Hsrp_Ipv6_Interfaces_Interface_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of advertisement packets sent. The type is interface{} with range:
    // 0..4294967295.
    AdvertPacketsSent interface{}

    // Number of advertisement packets received. The type is interface{} with
    // range: 0..4294967295.
    AdvertPacketsReceived interface{}

    // Number of packets received that were too Long. The type is interface{} with
    // range: 0..4294967295.
    LongPacketsReceived interface{}

    // Number of packets received that were too short. The type is interface{}
    // with range: 0..4294967295.
    ShortPacketsReceived interface{}

    // Number of packets received with invalid version. The type is interface{}
    // with range: 0..4294967295.
    InvalidVersionReceived interface{}

    // Number of packets received with invalid operation code. The type is
    // interface{} with range: 0..4294967295.
    InvalidOperationCodeReceived interface{}

    // Number of packets received for an unknown group id. The type is interface{}
    // with range: 0..4294967295.
    UnknownGroupReceived interface{}

    // Number of packets received for an inoperational group. The type is
    // interface{} with range: 0..4294967295.
    InoperationalGroupReceived interface{}

    // Number of packets received from a conflicting Source IP address. The type
    // is interface{} with range: 0..4294967295.
    ConflictSourceIpAddressReceived interface{}
}

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetGoName(yname string) string {
    if yname == "advert-packets-sent" { return "AdvertPacketsSent" }
    if yname == "advert-packets-received" { return "AdvertPacketsReceived" }
    if yname == "long-packets-received" { return "LongPacketsReceived" }
    if yname == "short-packets-received" { return "ShortPacketsReceived" }
    if yname == "invalid-version-received" { return "InvalidVersionReceived" }
    if yname == "invalid-operation-code-received" { return "InvalidOperationCodeReceived" }
    if yname == "unknown-group-received" { return "UnknownGroupReceived" }
    if yname == "inoperational-group-received" { return "InoperationalGroupReceived" }
    if yname == "conflict-source-ip-address-received" { return "ConflictSourceIpAddressReceived" }
    return ""
}

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["advert-packets-sent"] = statistics.AdvertPacketsSent
    leafs["advert-packets-received"] = statistics.AdvertPacketsReceived
    leafs["long-packets-received"] = statistics.LongPacketsReceived
    leafs["short-packets-received"] = statistics.ShortPacketsReceived
    leafs["invalid-version-received"] = statistics.InvalidVersionReceived
    leafs["invalid-operation-code-received"] = statistics.InvalidOperationCodeReceived
    leafs["unknown-group-received"] = statistics.UnknownGroupReceived
    leafs["inoperational-group-received"] = statistics.InoperationalGroupReceived
    leafs["conflict-source-ip-address-received"] = statistics.ConflictSourceIpAddressReceived
    return leafs
}

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetYangName() string { return "statistics" }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetParentYangName() string { return "interface" }

// Hsrp_BfdSessions
// The table of HSRP BFD Sessions
type Hsrp_BfdSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An HSRP BFD Session. The type is slice of Hsrp_BfdSessions_BfdSession.
    BfdSession []Hsrp_BfdSessions_BfdSession
}

func (bfdSessions *Hsrp_BfdSessions) GetFilter() yfilter.YFilter { return bfdSessions.YFilter }

func (bfdSessions *Hsrp_BfdSessions) SetFilter(yf yfilter.YFilter) { bfdSessions.YFilter = yf }

func (bfdSessions *Hsrp_BfdSessions) GetGoName(yname string) string {
    if yname == "bfd-session" { return "BfdSession" }
    return ""
}

func (bfdSessions *Hsrp_BfdSessions) GetSegmentPath() string {
    return "bfd-sessions"
}

func (bfdSessions *Hsrp_BfdSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd-session" {
        for _, c := range bfdSessions.BfdSession {
            if bfdSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_BfdSessions_BfdSession{}
        bfdSessions.BfdSession = append(bfdSessions.BfdSession, child)
        return &bfdSessions.BfdSession[len(bfdSessions.BfdSession)-1]
    }
    return nil
}

func (bfdSessions *Hsrp_BfdSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bfdSessions.BfdSession {
        children[bfdSessions.BfdSession[i].GetSegmentPath()] = &bfdSessions.BfdSession[i]
    }
    return children
}

func (bfdSessions *Hsrp_BfdSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bfdSessions *Hsrp_BfdSessions) GetBundleName() string { return "cisco_ios_xr" }

func (bfdSessions *Hsrp_BfdSessions) GetYangName() string { return "bfd-sessions" }

func (bfdSessions *Hsrp_BfdSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdSessions *Hsrp_BfdSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdSessions *Hsrp_BfdSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdSessions *Hsrp_BfdSessions) SetParent(parent types.Entity) { bfdSessions.parent = parent }

func (bfdSessions *Hsrp_BfdSessions) GetParent() types.Entity { return bfdSessions.parent }

func (bfdSessions *Hsrp_BfdSessions) GetParentYangName() string { return "hsrp" }

// Hsrp_BfdSessions_BfdSession
// An HSRP BFD Session
type Hsrp_BfdSessions_BfdSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. Destination IP Address of BFD Session. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // BFD Interface Name. The type is string with length: 0..64.
    BfdInterfaceName interface{}

    // Session Address family. The type is HsrpBAf.
    SessionAddressFamily interface{}

    // BFD destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // BFD IPv6 destination address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationIpv6Address interface{}

    // BFD session state. The type is HsrpBfdSessionState.
    BfdSessionState interface{}

    // BFD packet send interval. The type is interface{} with range:
    // 0..4294967295.
    BfdInterval interface{}

    // BFD multiplier. The type is interface{} with range: 0..4294967295.
    BfdMultiplier interface{}

    // HSRP Groups tracking the BFD session. The type is slice of
    // Hsrp_BfdSessions_BfdSession_Group.
    Group []Hsrp_BfdSessions_BfdSession_Group
}

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetFilter() yfilter.YFilter { return bfdSession.YFilter }

func (bfdSession *Hsrp_BfdSessions_BfdSession) SetFilter(yf yfilter.YFilter) { bfdSession.YFilter = yf }

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "bfd-interface-name" { return "BfdInterfaceName" }
    if yname == "session-address-family" { return "SessionAddressFamily" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-ipv6-address" { return "DestinationIpv6Address" }
    if yname == "bfd-session-state" { return "BfdSessionState" }
    if yname == "bfd-interval" { return "BfdInterval" }
    if yname == "bfd-multiplier" { return "BfdMultiplier" }
    if yname == "group" { return "Group" }
    return ""
}

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetSegmentPath() string {
    return "bfd-session" + "[interface-name='" + fmt.Sprintf("%v", bfdSession.InterfaceName) + "']" + "[ip-address='" + fmt.Sprintf("%v", bfdSession.IpAddress) + "']"
}

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range bfdSession.Group {
            if bfdSession.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_BfdSessions_BfdSession_Group{}
        bfdSession.Group = append(bfdSession.Group, child)
        return &bfdSession.Group[len(bfdSession.Group)-1]
    }
    return nil
}

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bfdSession.Group {
        children[bfdSession.Group[i].GetSegmentPath()] = &bfdSession.Group[i]
    }
    return children
}

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdSession.InterfaceName
    leafs["ip-address"] = bfdSession.IpAddress
    leafs["bfd-interface-name"] = bfdSession.BfdInterfaceName
    leafs["session-address-family"] = bfdSession.SessionAddressFamily
    leafs["destination-address"] = bfdSession.DestinationAddress
    leafs["destination-ipv6-address"] = bfdSession.DestinationIpv6Address
    leafs["bfd-session-state"] = bfdSession.BfdSessionState
    leafs["bfd-interval"] = bfdSession.BfdInterval
    leafs["bfd-multiplier"] = bfdSession.BfdMultiplier
    return leafs
}

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetBundleName() string { return "cisco_ios_xr" }

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetYangName() string { return "bfd-session" }

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdSession *Hsrp_BfdSessions_BfdSession) SetParent(parent types.Entity) { bfdSession.parent = parent }

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetParent() types.Entity { return bfdSession.parent }

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetParentYangName() string { return "bfd-sessions" }

// Hsrp_BfdSessions_BfdSession_Group
// HSRP Groups tracking the BFD session
type Hsrp_BfdSessions_BfdSession_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // HSRP Group number. The type is interface{} with range: 0..4294967295.
    HsrpGroupNumber interface{}
}

func (group *Hsrp_BfdSessions_BfdSession_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Hsrp_BfdSessions_BfdSession_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Hsrp_BfdSessions_BfdSession_Group) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "hsrp-group-number" { return "HsrpGroupNumber" }
    return ""
}

func (group *Hsrp_BfdSessions_BfdSession_Group) GetSegmentPath() string {
    return "group"
}

func (group *Hsrp_BfdSessions_BfdSession_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (group *Hsrp_BfdSessions_BfdSession_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (group *Hsrp_BfdSessions_BfdSession_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = group.InterfaceName
    leafs["hsrp-group-number"] = group.HsrpGroupNumber
    return leafs
}

func (group *Hsrp_BfdSessions_BfdSession_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Hsrp_BfdSessions_BfdSession_Group) GetYangName() string { return "group" }

func (group *Hsrp_BfdSessions_BfdSession_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Hsrp_BfdSessions_BfdSession_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Hsrp_BfdSessions_BfdSession_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Hsrp_BfdSessions_BfdSession_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Hsrp_BfdSessions_BfdSession_Group) GetParent() types.Entity { return group.parent }

func (group *Hsrp_BfdSessions_BfdSession_Group) GetParentYangName() string { return "bfd-session" }

// Hsrp_Summary
// HSRP summary statistics
type Hsrp_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of IPv4 sessions in ACTIVE state. The type is interface{} with
    // range: 0..4294967295.
    Ipv4SessionsActive interface{}

    // Number of IPv4 sessions in STANDBY state. The type is interface{} with
    // range: 0..4294967295.
    Ipv4SessionsStandby interface{}

    // Number of IPv4 sessions in SPEAK state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SessionsSpeak interface{}

    // Number of IPv4 sessions in LISTEN state. The type is interface{} with
    // range: 0..4294967295.
    Ipv4SessionsListen interface{}

    // Number of IPv4 sessions in LEARN state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SessionsLearn interface{}

    // Number of IPv4 sessions in INIT state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SessionsInit interface{}

    // Number of IPv4 slaves in ACTIVE state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesActive interface{}

    // Number of IPv4 slaves in STANDBY state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesStandby interface{}

    // Number of IPv4 slaves in SPEAK state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesSpeak interface{}

    // Number of IPv4 slaves in LISTEN state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesListen interface{}

    // Number of IPv4 slaves in LEARN state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesLearn interface{}

    // Number of IPv4 slaves in INIT state. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SlavesInit interface{}

    // Number of UP IPv4 Virtual IP Addresses on groups in ACTIVE state. The type
    // is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesActiveUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on groups in ACTIVE state. The
    // type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesActiveDown interface{}

    // Number of UP IPv4 Virtual IP Addresses on groups in STANDBY state. The type
    // is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesStandbyUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on groups in STANDBY state. The
    // type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesStandbyDown interface{}

    // Number of UP IPv4 Virtual IP Addresses on groups in SPEAK state. The type
    // is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesSpeakUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on groups in SPEAK state. The type
    // is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesSpeakDown interface{}

    // Number of UP IPv4 Virtual IP Addresses on groups in LISTEN state. The type
    // is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesListenUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on groups in LISTEN state. The
    // type is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesListenDown interface{}

    // Number of UP IPv4 Virtual IP Addresses on groups in LEARN state. The type
    // is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesLearnUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on groups in LEARN state. The type
    // is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesLearnDown interface{}

    // Number of UP IPv4 Virtual IP Addresses on groups in INIT state. The type is
    // interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesInitUp interface{}

    // Number of DOWN IPv4 Virtual IP Addresses on groups in INIT state. The type
    // is interface{} with range: 0..4294967295.
    Ipv4VirtualIpAddressesInitDown interface{}

    // Number of IPv6 sessions in ACTIVE state. The type is interface{} with
    // range: 0..4294967295.
    Ipv6SessionsActive interface{}

    // Number of IPv6 sessions in STANDBY state. The type is interface{} with
    // range: 0..4294967295.
    Ipv6SessionsStandby interface{}

    // Number of IPv6 sessions in SPEAK state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SessionsSpeak interface{}

    // Number of IPv6 sessions in LISTEN state. The type is interface{} with
    // range: 0..4294967295.
    Ipv6SessionsListen interface{}

    // Number of IPv6 sessions in LEARN state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SessionsLearn interface{}

    // Number of IPv6 sessions in INIT state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SessionsInit interface{}

    // Number of IPv6 slaves in ACTIVE state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesActive interface{}

    // Number of IPv6 slaves in STANDBY state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesStandby interface{}

    // Number of IPv6 slaves in SPEAK state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesSpeak interface{}

    // Number of IPv6 slaves in LISTEN state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesListen interface{}

    // Number of IPv6 slaves in LEARN state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesLearn interface{}

    // Number of IPv6 slaves in INIT state. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SlavesInit interface{}

    // Number of UP IPv6 Virtual IP Addresses on groups in ACTIVE state. The type
    // is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesActiveUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on groups in ACTIVE state. The
    // type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesActiveDown interface{}

    // Number of UP IPv6 Virtual IP Addresses on groups in STANDBY state. The type
    // is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesStandbyUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on groups in STANDBY state. The
    // type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesStandbyDown interface{}

    // Number of UP IPv6 Virtual IP Addresses on groups in SPEAK state. The type
    // is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesSpeakUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on groups in SPEAK state. The type
    // is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesSpeakDown interface{}

    // Number of UP IPv6 Virtual IP Addresses on groups in LISTEN state. The type
    // is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesListenUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on groups in LISTEN state. The
    // type is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesListenDown interface{}

    // Number of UP IPv6 Virtual IP Addresses on groups in LEARN state. The type
    // is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesLearnUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on groups in LEARN state. The type
    // is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesLearnDown interface{}

    // Number of UP IPv6 Virtual IP Addresses on groups in INIT state. The type is
    // interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesInitUp interface{}

    // Number of DOWN IPv6 Virtual IP Addresses on groups in INIT state. The type
    // is interface{} with range: 0..4294967295.
    Ipv6VirtualIpAddressesInitDown interface{}

    // Number of HSRP interfaces with IPv4 caps in UP state. The type is
    // interface{} with range: 0..4294967295.
    InterfacesIpv4StateUp interface{}

    // Number of HSRP interfaces with IPv4 caps in DOWN state. The type is
    // interface{} with range: 0..4294967295.
    InterfacesIpv4StateDown interface{}

    // Number of tracked interfaces with IPv4 caps in UP state. The type is
    // interface{} with range: 0..4294967295.
    TrackedInterfacesIpv4StateUp interface{}

    // Number of tracked interfaces with IPv4 caps in DOWN state. The type is
    // interface{} with range: 0..4294967295.
    TrackedInterfacesIpv4StateDown interface{}

    // Number of tracked objects in UP state. The type is interface{} with range:
    // 0..4294967295.
    TrackedObjectsUp interface{}

    // Number of tracked objects in DOWN state. The type is interface{} with
    // range: 0..4294967295.
    TrackedObjectsDown interface{}

    // Number of HSRP interfaces with IPv6 caps in UP state. The type is
    // interface{} with range: 0..4294967295.
    InterfacesIpv6StateUp interface{}

    // Number of HSRP interfaces with IPv6 caps in DOWN state. The type is
    // interface{} with range: 0..4294967295.
    InterfacesIpv6StateDown interface{}

    // Number of tracked interfaces with IPv6 caps in UP state. The type is
    // interface{} with range: 0..4294967295.
    TrackedInterfacesIpv6StateUp interface{}

    // Number of tracked interfaces with IPv6 caps in DOWN state. The type is
    // interface{} with range: 0..4294967295.
    TrackedInterfacesIpv6StateDown interface{}

    // Number of HSRP BFD sessions in UP state. The type is interface{} with
    // range: 0..4294967295.
    BfdSessionsUp interface{}

    // Number of HSRP BFD sessions in DOWN state. The type is interface{} with
    // range: 0..4294967295.
    BfdSessionsDown interface{}

    // Number of HSRP BFD sessions in INACTIVE state. The type is interface{} with
    // range: 0..4294967295.
    BfdSessionInactive interface{}
}

func (summary *Hsrp_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Hsrp_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Hsrp_Summary) GetGoName(yname string) string {
    if yname == "ipv4-sessions-active" { return "Ipv4SessionsActive" }
    if yname == "ipv4-sessions-standby" { return "Ipv4SessionsStandby" }
    if yname == "ipv4-sessions-speak" { return "Ipv4SessionsSpeak" }
    if yname == "ipv4-sessions-listen" { return "Ipv4SessionsListen" }
    if yname == "ipv4-sessions-learn" { return "Ipv4SessionsLearn" }
    if yname == "ipv4-sessions-init" { return "Ipv4SessionsInit" }
    if yname == "ipv4-slaves-active" { return "Ipv4SlavesActive" }
    if yname == "ipv4-slaves-standby" { return "Ipv4SlavesStandby" }
    if yname == "ipv4-slaves-speak" { return "Ipv4SlavesSpeak" }
    if yname == "ipv4-slaves-listen" { return "Ipv4SlavesListen" }
    if yname == "ipv4-slaves-learn" { return "Ipv4SlavesLearn" }
    if yname == "ipv4-slaves-init" { return "Ipv4SlavesInit" }
    if yname == "ipv4-virtual-ip-addresses-active-up" { return "Ipv4VirtualIpAddressesActiveUp" }
    if yname == "ipv4-virtual-ip-addresses-active-down" { return "Ipv4VirtualIpAddressesActiveDown" }
    if yname == "ipv4-virtual-ip-addresses-standby-up" { return "Ipv4VirtualIpAddressesStandbyUp" }
    if yname == "ipv4-virtual-ip-addresses-standby-down" { return "Ipv4VirtualIpAddressesStandbyDown" }
    if yname == "ipv4-virtual-ip-addresses-speak-up" { return "Ipv4VirtualIpAddressesSpeakUp" }
    if yname == "ipv4-virtual-ip-addresses-speak-down" { return "Ipv4VirtualIpAddressesSpeakDown" }
    if yname == "ipv4-virtual-ip-addresses-listen-up" { return "Ipv4VirtualIpAddressesListenUp" }
    if yname == "ipv4-virtual-ip-addresses-listen-down" { return "Ipv4VirtualIpAddressesListenDown" }
    if yname == "ipv4-virtual-ip-addresses-learn-up" { return "Ipv4VirtualIpAddressesLearnUp" }
    if yname == "ipv4-virtual-ip-addresses-learn-down" { return "Ipv4VirtualIpAddressesLearnDown" }
    if yname == "ipv4-virtual-ip-addresses-init-up" { return "Ipv4VirtualIpAddressesInitUp" }
    if yname == "ipv4-virtual-ip-addresses-init-down" { return "Ipv4VirtualIpAddressesInitDown" }
    if yname == "ipv6-sessions-active" { return "Ipv6SessionsActive" }
    if yname == "ipv6-sessions-standby" { return "Ipv6SessionsStandby" }
    if yname == "ipv6-sessions-speak" { return "Ipv6SessionsSpeak" }
    if yname == "ipv6-sessions-listen" { return "Ipv6SessionsListen" }
    if yname == "ipv6-sessions-learn" { return "Ipv6SessionsLearn" }
    if yname == "ipv6-sessions-init" { return "Ipv6SessionsInit" }
    if yname == "ipv6-slaves-active" { return "Ipv6SlavesActive" }
    if yname == "ipv6-slaves-standby" { return "Ipv6SlavesStandby" }
    if yname == "ipv6-slaves-speak" { return "Ipv6SlavesSpeak" }
    if yname == "ipv6-slaves-listen" { return "Ipv6SlavesListen" }
    if yname == "ipv6-slaves-learn" { return "Ipv6SlavesLearn" }
    if yname == "ipv6-slaves-init" { return "Ipv6SlavesInit" }
    if yname == "ipv6-virtual-ip-addresses-active-up" { return "Ipv6VirtualIpAddressesActiveUp" }
    if yname == "ipv6-virtual-ip-addresses-active-down" { return "Ipv6VirtualIpAddressesActiveDown" }
    if yname == "ipv6-virtual-ip-addresses-standby-up" { return "Ipv6VirtualIpAddressesStandbyUp" }
    if yname == "ipv6-virtual-ip-addresses-standby-down" { return "Ipv6VirtualIpAddressesStandbyDown" }
    if yname == "ipv6-virtual-ip-addresses-speak-up" { return "Ipv6VirtualIpAddressesSpeakUp" }
    if yname == "ipv6-virtual-ip-addresses-speak-down" { return "Ipv6VirtualIpAddressesSpeakDown" }
    if yname == "ipv6-virtual-ip-addresses-listen-up" { return "Ipv6VirtualIpAddressesListenUp" }
    if yname == "ipv6-virtual-ip-addresses-listen-down" { return "Ipv6VirtualIpAddressesListenDown" }
    if yname == "ipv6-virtual-ip-addresses-learn-up" { return "Ipv6VirtualIpAddressesLearnUp" }
    if yname == "ipv6-virtual-ip-addresses-learn-down" { return "Ipv6VirtualIpAddressesLearnDown" }
    if yname == "ipv6-virtual-ip-addresses-init-up" { return "Ipv6VirtualIpAddressesInitUp" }
    if yname == "ipv6-virtual-ip-addresses-init-down" { return "Ipv6VirtualIpAddressesInitDown" }
    if yname == "interfaces-ipv4-state-up" { return "InterfacesIpv4StateUp" }
    if yname == "interfaces-ipv4-state-down" { return "InterfacesIpv4StateDown" }
    if yname == "tracked-interfaces-ipv4-state-up" { return "TrackedInterfacesIpv4StateUp" }
    if yname == "tracked-interfaces-ipv4-state-down" { return "TrackedInterfacesIpv4StateDown" }
    if yname == "tracked-objects-up" { return "TrackedObjectsUp" }
    if yname == "tracked-objects-down" { return "TrackedObjectsDown" }
    if yname == "interfaces-ipv6-state-up" { return "InterfacesIpv6StateUp" }
    if yname == "interfaces-ipv6-state-down" { return "InterfacesIpv6StateDown" }
    if yname == "tracked-interfaces-ipv6-state-up" { return "TrackedInterfacesIpv6StateUp" }
    if yname == "tracked-interfaces-ipv6-state-down" { return "TrackedInterfacesIpv6StateDown" }
    if yname == "bfd-sessions-up" { return "BfdSessionsUp" }
    if yname == "bfd-sessions-down" { return "BfdSessionsDown" }
    if yname == "bfd-session-inactive" { return "BfdSessionInactive" }
    return ""
}

func (summary *Hsrp_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Hsrp_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Hsrp_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Hsrp_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-sessions-active"] = summary.Ipv4SessionsActive
    leafs["ipv4-sessions-standby"] = summary.Ipv4SessionsStandby
    leafs["ipv4-sessions-speak"] = summary.Ipv4SessionsSpeak
    leafs["ipv4-sessions-listen"] = summary.Ipv4SessionsListen
    leafs["ipv4-sessions-learn"] = summary.Ipv4SessionsLearn
    leafs["ipv4-sessions-init"] = summary.Ipv4SessionsInit
    leafs["ipv4-slaves-active"] = summary.Ipv4SlavesActive
    leafs["ipv4-slaves-standby"] = summary.Ipv4SlavesStandby
    leafs["ipv4-slaves-speak"] = summary.Ipv4SlavesSpeak
    leafs["ipv4-slaves-listen"] = summary.Ipv4SlavesListen
    leafs["ipv4-slaves-learn"] = summary.Ipv4SlavesLearn
    leafs["ipv4-slaves-init"] = summary.Ipv4SlavesInit
    leafs["ipv4-virtual-ip-addresses-active-up"] = summary.Ipv4VirtualIpAddressesActiveUp
    leafs["ipv4-virtual-ip-addresses-active-down"] = summary.Ipv4VirtualIpAddressesActiveDown
    leafs["ipv4-virtual-ip-addresses-standby-up"] = summary.Ipv4VirtualIpAddressesStandbyUp
    leafs["ipv4-virtual-ip-addresses-standby-down"] = summary.Ipv4VirtualIpAddressesStandbyDown
    leafs["ipv4-virtual-ip-addresses-speak-up"] = summary.Ipv4VirtualIpAddressesSpeakUp
    leafs["ipv4-virtual-ip-addresses-speak-down"] = summary.Ipv4VirtualIpAddressesSpeakDown
    leafs["ipv4-virtual-ip-addresses-listen-up"] = summary.Ipv4VirtualIpAddressesListenUp
    leafs["ipv4-virtual-ip-addresses-listen-down"] = summary.Ipv4VirtualIpAddressesListenDown
    leafs["ipv4-virtual-ip-addresses-learn-up"] = summary.Ipv4VirtualIpAddressesLearnUp
    leafs["ipv4-virtual-ip-addresses-learn-down"] = summary.Ipv4VirtualIpAddressesLearnDown
    leafs["ipv4-virtual-ip-addresses-init-up"] = summary.Ipv4VirtualIpAddressesInitUp
    leafs["ipv4-virtual-ip-addresses-init-down"] = summary.Ipv4VirtualIpAddressesInitDown
    leafs["ipv6-sessions-active"] = summary.Ipv6SessionsActive
    leafs["ipv6-sessions-standby"] = summary.Ipv6SessionsStandby
    leafs["ipv6-sessions-speak"] = summary.Ipv6SessionsSpeak
    leafs["ipv6-sessions-listen"] = summary.Ipv6SessionsListen
    leafs["ipv6-sessions-learn"] = summary.Ipv6SessionsLearn
    leafs["ipv6-sessions-init"] = summary.Ipv6SessionsInit
    leafs["ipv6-slaves-active"] = summary.Ipv6SlavesActive
    leafs["ipv6-slaves-standby"] = summary.Ipv6SlavesStandby
    leafs["ipv6-slaves-speak"] = summary.Ipv6SlavesSpeak
    leafs["ipv6-slaves-listen"] = summary.Ipv6SlavesListen
    leafs["ipv6-slaves-learn"] = summary.Ipv6SlavesLearn
    leafs["ipv6-slaves-init"] = summary.Ipv6SlavesInit
    leafs["ipv6-virtual-ip-addresses-active-up"] = summary.Ipv6VirtualIpAddressesActiveUp
    leafs["ipv6-virtual-ip-addresses-active-down"] = summary.Ipv6VirtualIpAddressesActiveDown
    leafs["ipv6-virtual-ip-addresses-standby-up"] = summary.Ipv6VirtualIpAddressesStandbyUp
    leafs["ipv6-virtual-ip-addresses-standby-down"] = summary.Ipv6VirtualIpAddressesStandbyDown
    leafs["ipv6-virtual-ip-addresses-speak-up"] = summary.Ipv6VirtualIpAddressesSpeakUp
    leafs["ipv6-virtual-ip-addresses-speak-down"] = summary.Ipv6VirtualIpAddressesSpeakDown
    leafs["ipv6-virtual-ip-addresses-listen-up"] = summary.Ipv6VirtualIpAddressesListenUp
    leafs["ipv6-virtual-ip-addresses-listen-down"] = summary.Ipv6VirtualIpAddressesListenDown
    leafs["ipv6-virtual-ip-addresses-learn-up"] = summary.Ipv6VirtualIpAddressesLearnUp
    leafs["ipv6-virtual-ip-addresses-learn-down"] = summary.Ipv6VirtualIpAddressesLearnDown
    leafs["ipv6-virtual-ip-addresses-init-up"] = summary.Ipv6VirtualIpAddressesInitUp
    leafs["ipv6-virtual-ip-addresses-init-down"] = summary.Ipv6VirtualIpAddressesInitDown
    leafs["interfaces-ipv4-state-up"] = summary.InterfacesIpv4StateUp
    leafs["interfaces-ipv4-state-down"] = summary.InterfacesIpv4StateDown
    leafs["tracked-interfaces-ipv4-state-up"] = summary.TrackedInterfacesIpv4StateUp
    leafs["tracked-interfaces-ipv4-state-down"] = summary.TrackedInterfacesIpv4StateDown
    leafs["tracked-objects-up"] = summary.TrackedObjectsUp
    leafs["tracked-objects-down"] = summary.TrackedObjectsDown
    leafs["interfaces-ipv6-state-up"] = summary.InterfacesIpv6StateUp
    leafs["interfaces-ipv6-state-down"] = summary.InterfacesIpv6StateDown
    leafs["tracked-interfaces-ipv6-state-up"] = summary.TrackedInterfacesIpv6StateUp
    leafs["tracked-interfaces-ipv6-state-down"] = summary.TrackedInterfacesIpv6StateDown
    leafs["bfd-sessions-up"] = summary.BfdSessionsUp
    leafs["bfd-sessions-down"] = summary.BfdSessionsDown
    leafs["bfd-session-inactive"] = summary.BfdSessionInactive
    return leafs
}

func (summary *Hsrp_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Hsrp_Summary) GetYangName() string { return "summary" }

func (summary *Hsrp_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Hsrp_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Hsrp_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Hsrp_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Hsrp_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Hsrp_Summary) GetParentYangName() string { return "hsrp" }

