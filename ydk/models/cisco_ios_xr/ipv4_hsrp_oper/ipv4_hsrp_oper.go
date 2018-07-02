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
    EntityData types.CommonEntityData
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

func (hsrp *Hsrp) GetEntityData() *types.CommonEntityData {
    hsrp.EntityData.YFilter = hsrp.YFilter
    hsrp.EntityData.YangName = "hsrp"
    hsrp.EntityData.BundleName = "cisco_ios_xr"
    hsrp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-hsrp-oper"
    hsrp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-hsrp-oper:hsrp"
    hsrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hsrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hsrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hsrp.EntityData.Children = types.NewOrderedMap()
    hsrp.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &hsrp.Ipv4})
    hsrp.EntityData.Children.Append("mgo-sessions", types.YChild{"MgoSessions", &hsrp.MgoSessions})
    hsrp.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &hsrp.Ipv6})
    hsrp.EntityData.Children.Append("bfd-sessions", types.YChild{"BfdSessions", &hsrp.BfdSessions})
    hsrp.EntityData.Children.Append("summary", types.YChild{"Summary", &hsrp.Summary})
    hsrp.EntityData.Leafs = types.NewOrderedMap()

    hsrp.EntityData.YListKeys = []string {}

    return &(hsrp.EntityData)
}

// Hsrp_Ipv4
// IPv4 HSRP information
type Hsrp_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP standby group table.
    Groups Hsrp_Ipv4_Groups

    // The HSRP tracked interfaces table.
    TrackedInterfaces Hsrp_Ipv4_TrackedInterfaces

    // The HSRP interface information table.
    Interfaces Hsrp_Ipv4_Interfaces
}

func (ipv4 *Hsrp_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "hsrp"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("groups", types.YChild{"Groups", &ipv4.Groups})
    ipv4.EntityData.Children.Append("tracked-interfaces", types.YChild{"TrackedInterfaces", &ipv4.TrackedInterfaces})
    ipv4.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv4.Interfaces})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Hsrp_Ipv4_Groups
// The HSRP standby group table
type Hsrp_Ipv4_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An HSRP standby group. The type is slice of Hsrp_Ipv4_Groups_Group.
    Group []*Hsrp_Ipv4_Groups_Group
}

func (groups *Hsrp_Ipv4_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "ipv4"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Hsrp_Ipv4_Groups_Group
// An HSRP standby group
type Hsrp_Ipv4_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The HSRP group number. The type is interface{}
    // with range: 0..4294967295.
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
    GlobalAddress []*Hsrp_Ipv4_Groups_Group_GlobalAddress

    // State change history. The type is slice of
    // Hsrp_Ipv4_Groups_Group_StateChangeHistory.
    StateChangeHistory []*Hsrp_Ipv4_Groups_Group_StateChangeHistory
}

func (group *Hsrp_Ipv4_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.InterfaceName, "interface-name") + types.AddKeyToken(group.GroupNumber, "group-number")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("resign-sent-time", types.YChild{"ResignSentTime", &group.ResignSentTime})
    group.EntityData.Children.Append("resign-received-time", types.YChild{"ResignReceivedTime", &group.ResignReceivedTime})
    group.EntityData.Children.Append("coup-sent-time", types.YChild{"CoupSentTime", &group.CoupSentTime})
    group.EntityData.Children.Append("coup-received-time", types.YChild{"CoupReceivedTime", &group.CoupReceivedTime})
    group.EntityData.Children.Append("statistics", types.YChild{"Statistics", &group.Statistics})
    group.EntityData.Children.Append("global-address", types.YChild{"GlobalAddress", nil})
    for i := range group.GlobalAddress {
        group.EntityData.Children.Append(types.GetSegmentPath(group.GlobalAddress[i]), types.YChild{"GlobalAddress", group.GlobalAddress[i]})
    }
    group.EntityData.Children.Append("state-change-history", types.YChild{"StateChangeHistory", nil})
    for i := range group.StateChangeHistory {
        group.EntityData.Children.Append(types.GetSegmentPath(group.StateChangeHistory[i]), types.YChild{"StateChangeHistory", group.StateChangeHistory[i]})
    }
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", group.InterfaceName})
    group.EntityData.Leafs.Append("group-number", types.YLeaf{"GroupNumber", group.GroupNumber})
    group.EntityData.Leafs.Append("authentication-string", types.YLeaf{"AuthenticationString", group.AuthenticationString})
    group.EntityData.Leafs.Append("virtual-mac-address", types.YLeaf{"VirtualMacAddress", group.VirtualMacAddress})
    group.EntityData.Leafs.Append("hsrp-group-number", types.YLeaf{"HsrpGroupNumber", group.HsrpGroupNumber})
    group.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", group.AddressFamily})
    group.EntityData.Leafs.Append("version", types.YLeaf{"Version", group.Version})
    group.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", group.SessionName})
    group.EntityData.Leafs.Append("slaves", types.YLeaf{"Slaves", group.Slaves})
    group.EntityData.Leafs.Append("is-slave", types.YLeaf{"IsSlave", group.IsSlave})
    group.EntityData.Leafs.Append("followed-session-name", types.YLeaf{"FollowedSessionName", group.FollowedSessionName})
    group.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", group.ConfiguredPriority})
    group.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", group.PreemptDelay})
    group.EntityData.Leafs.Append("preempt-timer-secs", types.YLeaf{"PreemptTimerSecs", group.PreemptTimerSecs})
    group.EntityData.Leafs.Append("hello-time", types.YLeaf{"HelloTime", group.HelloTime})
    group.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", group.HoldTime})
    group.EntityData.Leafs.Append("learned-hello-time", types.YLeaf{"LearnedHelloTime", group.LearnedHelloTime})
    group.EntityData.Leafs.Append("learned-hold-time", types.YLeaf{"LearnedHoldTime", group.LearnedHoldTime})
    group.EntityData.Leafs.Append("min-delay-time", types.YLeaf{"MinDelayTime", group.MinDelayTime})
    group.EntityData.Leafs.Append("reload-delay-time", types.YLeaf{"ReloadDelayTime", group.ReloadDelayTime})
    group.EntityData.Leafs.Append("virtual-ip-address", types.YLeaf{"VirtualIpAddress", group.VirtualIpAddress})
    group.EntityData.Leafs.Append("virtual-linklocal-ipv6-address", types.YLeaf{"VirtualLinklocalIpv6Address", group.VirtualLinklocalIpv6Address})
    group.EntityData.Leafs.Append("active-ip-address", types.YLeaf{"ActiveIpAddress", group.ActiveIpAddress})
    group.EntityData.Leafs.Append("active-ipv6-address", types.YLeaf{"ActiveIpv6Address", group.ActiveIpv6Address})
    group.EntityData.Leafs.Append("active-mac-address", types.YLeaf{"ActiveMacAddress", group.ActiveMacAddress})
    group.EntityData.Leafs.Append("standby-ip-address", types.YLeaf{"StandbyIpAddress", group.StandbyIpAddress})
    group.EntityData.Leafs.Append("standby-ipv6-address", types.YLeaf{"StandbyIpv6Address", group.StandbyIpv6Address})
    group.EntityData.Leafs.Append("standby-mac-address", types.YLeaf{"StandbyMacAddress", group.StandbyMacAddress})
    group.EntityData.Leafs.Append("hsrp-router-state", types.YLeaf{"HsrpRouterState", group.HsrpRouterState})
    group.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", group.InterfaceNameXr})
    group.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", group.Interface})
    group.EntityData.Leafs.Append("router-priority", types.YLeaf{"RouterPriority", group.RouterPriority})
    group.EntityData.Leafs.Append("active-priority", types.YLeaf{"ActivePriority", group.ActivePriority})
    group.EntityData.Leafs.Append("active-timer-flag", types.YLeaf{"ActiveTimerFlag", group.ActiveTimerFlag})
    group.EntityData.Leafs.Append("active-timer-secs", types.YLeaf{"ActiveTimerSecs", group.ActiveTimerSecs})
    group.EntityData.Leafs.Append("active-timer-msecs", types.YLeaf{"ActiveTimerMsecs", group.ActiveTimerMsecs})
    group.EntityData.Leafs.Append("standby-timer-flag", types.YLeaf{"StandbyTimerFlag", group.StandbyTimerFlag})
    group.EntityData.Leafs.Append("standby-timer-secs", types.YLeaf{"StandbyTimerSecs", group.StandbyTimerSecs})
    group.EntityData.Leafs.Append("standby-timer-msecs", types.YLeaf{"StandbyTimerMsecs", group.StandbyTimerMsecs})
    group.EntityData.Leafs.Append("hello-timer-flag", types.YLeaf{"HelloTimerFlag", group.HelloTimerFlag})
    group.EntityData.Leafs.Append("hello-timer-secs", types.YLeaf{"HelloTimerSecs", group.HelloTimerSecs})
    group.EntityData.Leafs.Append("hello-timer-msecs", types.YLeaf{"HelloTimerMsecs", group.HelloTimerMsecs})
    group.EntityData.Leafs.Append("delay-timer-flag", types.YLeaf{"DelayTimerFlag", group.DelayTimerFlag})
    group.EntityData.Leafs.Append("delay-timer-secs", types.YLeaf{"DelayTimerSecs", group.DelayTimerSecs})
    group.EntityData.Leafs.Append("delay-timer-msecs", types.YLeaf{"DelayTimerMsecs", group.DelayTimerMsecs})
    group.EntityData.Leafs.Append("current-state-timer-secs", types.YLeaf{"CurrentStateTimerSecs", group.CurrentStateTimerSecs})
    group.EntityData.Leafs.Append("state-change-count", types.YLeaf{"StateChangeCount", group.StateChangeCount})
    group.EntityData.Leafs.Append("tracked-interface-count", types.YLeaf{"TrackedInterfaceCount", group.TrackedInterfaceCount})
    group.EntityData.Leafs.Append("tracked-interface-up-count", types.YLeaf{"TrackedInterfaceUpCount", group.TrackedInterfaceUpCount})
    group.EntityData.Leafs.Append("preempt-enabled", types.YLeaf{"PreemptEnabled", group.PreemptEnabled})
    group.EntityData.Leafs.Append("use-configured-timers", types.YLeaf{"UseConfiguredTimers", group.UseConfiguredTimers})
    group.EntityData.Leafs.Append("use-configured-virtual-ip", types.YLeaf{"UseConfiguredVirtualIp", group.UseConfiguredVirtualIp})
    group.EntityData.Leafs.Append("use-bia-configured", types.YLeaf{"UseBiaConfigured", group.UseBiaConfigured})
    group.EntityData.Leafs.Append("configured-timers", types.YLeaf{"ConfiguredTimers", group.ConfiguredTimers})
    group.EntityData.Leafs.Append("configured-mac-address", types.YLeaf{"ConfiguredMacAddress", group.ConfiguredMacAddress})
    group.EntityData.Leafs.Append("redirects-disabled", types.YLeaf{"RedirectsDisabled", group.RedirectsDisabled})
    group.EntityData.Leafs.Append("bfd-enabled", types.YLeaf{"BfdEnabled", group.BfdEnabled})
    group.EntityData.Leafs.Append("bfd-interface", types.YLeaf{"BfdInterface", group.BfdInterface})
    group.EntityData.Leafs.Append("bfd-peer-ip-address", types.YLeaf{"BfdPeerIpAddress", group.BfdPeerIpAddress})
    group.EntityData.Leafs.Append("bfd-peer-ipv6-address", types.YLeaf{"BfdPeerIpv6Address", group.BfdPeerIpv6Address})
    group.EntityData.Leafs.Append("bfd-session-state", types.YLeaf{"BfdSessionState", group.BfdSessionState})
    group.EntityData.Leafs.Append("bfd-interval", types.YLeaf{"BfdInterval", group.BfdInterval})
    group.EntityData.Leafs.Append("bfd-multiplier", types.YLeaf{"BfdMultiplier", group.BfdMultiplier})
    group.EntityData.Leafs.Append("virtual-mac-address-state", types.YLeaf{"VirtualMacAddressState", group.VirtualMacAddressState})
    group.EntityData.Leafs.Append("secondary-address", types.YLeaf{"SecondaryAddress", group.SecondaryAddress})

    group.EntityData.YListKeys = []string {"InterfaceName", "GroupNumber"}

    return &(group.EntityData)
}

// Hsrp_Ipv4_Groups_Group_ResignSentTime
// Time last resign was sent
type Hsrp_Ipv4_Groups_Group_ResignSentTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignSentTime *Hsrp_Ipv4_Groups_Group_ResignSentTime) GetEntityData() *types.CommonEntityData {
    resignSentTime.EntityData.YFilter = resignSentTime.YFilter
    resignSentTime.EntityData.YangName = "resign-sent-time"
    resignSentTime.EntityData.BundleName = "cisco_ios_xr"
    resignSentTime.EntityData.ParentYangName = "group"
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

// Hsrp_Ipv4_Groups_Group_ResignReceivedTime
// Time last resign was received
type Hsrp_Ipv4_Groups_Group_ResignReceivedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignReceivedTime *Hsrp_Ipv4_Groups_Group_ResignReceivedTime) GetEntityData() *types.CommonEntityData {
    resignReceivedTime.EntityData.YFilter = resignReceivedTime.YFilter
    resignReceivedTime.EntityData.YangName = "resign-received-time"
    resignReceivedTime.EntityData.BundleName = "cisco_ios_xr"
    resignReceivedTime.EntityData.ParentYangName = "group"
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

// Hsrp_Ipv4_Groups_Group_CoupSentTime
// Time last coup was sent
type Hsrp_Ipv4_Groups_Group_CoupSentTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (coupSentTime *Hsrp_Ipv4_Groups_Group_CoupSentTime) GetEntityData() *types.CommonEntityData {
    coupSentTime.EntityData.YFilter = coupSentTime.YFilter
    coupSentTime.EntityData.YangName = "coup-sent-time"
    coupSentTime.EntityData.BundleName = "cisco_ios_xr"
    coupSentTime.EntityData.ParentYangName = "group"
    coupSentTime.EntityData.SegmentPath = "coup-sent-time"
    coupSentTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coupSentTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coupSentTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coupSentTime.EntityData.Children = types.NewOrderedMap()
    coupSentTime.EntityData.Leafs = types.NewOrderedMap()
    coupSentTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", coupSentTime.Seconds})
    coupSentTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", coupSentTime.Nanoseconds})

    coupSentTime.EntityData.YListKeys = []string {}

    return &(coupSentTime.EntityData)
}

// Hsrp_Ipv4_Groups_Group_CoupReceivedTime
// Time last coup was received
type Hsrp_Ipv4_Groups_Group_CoupReceivedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (coupReceivedTime *Hsrp_Ipv4_Groups_Group_CoupReceivedTime) GetEntityData() *types.CommonEntityData {
    coupReceivedTime.EntityData.YFilter = coupReceivedTime.YFilter
    coupReceivedTime.EntityData.YangName = "coup-received-time"
    coupReceivedTime.EntityData.BundleName = "cisco_ios_xr"
    coupReceivedTime.EntityData.ParentYangName = "group"
    coupReceivedTime.EntityData.SegmentPath = "coup-received-time"
    coupReceivedTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coupReceivedTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coupReceivedTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coupReceivedTime.EntityData.Children = types.NewOrderedMap()
    coupReceivedTime.EntityData.Leafs = types.NewOrderedMap()
    coupReceivedTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", coupReceivedTime.Seconds})
    coupReceivedTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", coupReceivedTime.Nanoseconds})

    coupReceivedTime.EntityData.YListKeys = []string {}

    return &(coupReceivedTime.EntityData)
}

// Hsrp_Ipv4_Groups_Group_Statistics
// HSRP Group statistics
type Hsrp_Ipv4_Groups_Group_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Hsrp_Ipv4_Groups_Group_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "group"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("active-transitions", types.YLeaf{"ActiveTransitions", statistics.ActiveTransitions})
    statistics.EntityData.Leafs.Append("standby-transitions", types.YLeaf{"StandbyTransitions", statistics.StandbyTransitions})
    statistics.EntityData.Leafs.Append("speak-transitions", types.YLeaf{"SpeakTransitions", statistics.SpeakTransitions})
    statistics.EntityData.Leafs.Append("listen-transitions", types.YLeaf{"ListenTransitions", statistics.ListenTransitions})
    statistics.EntityData.Leafs.Append("learn-transitions", types.YLeaf{"LearnTransitions", statistics.LearnTransitions})
    statistics.EntityData.Leafs.Append("init-transitions", types.YLeaf{"InitTransitions", statistics.InitTransitions})
    statistics.EntityData.Leafs.Append("hello-packets-sent", types.YLeaf{"HelloPacketsSent", statistics.HelloPacketsSent})
    statistics.EntityData.Leafs.Append("resign-packets-sent", types.YLeaf{"ResignPacketsSent", statistics.ResignPacketsSent})
    statistics.EntityData.Leafs.Append("coup-packets-sent", types.YLeaf{"CoupPacketsSent", statistics.CoupPacketsSent})
    statistics.EntityData.Leafs.Append("hello-packets-received", types.YLeaf{"HelloPacketsReceived", statistics.HelloPacketsReceived})
    statistics.EntityData.Leafs.Append("resign-packets-received", types.YLeaf{"ResignPacketsReceived", statistics.ResignPacketsReceived})
    statistics.EntityData.Leafs.Append("coup-packets-received", types.YLeaf{"CoupPacketsReceived", statistics.CoupPacketsReceived})
    statistics.EntityData.Leafs.Append("auth-fail-received", types.YLeaf{"AuthFailReceived", statistics.AuthFailReceived})
    statistics.EntityData.Leafs.Append("invalid-timer-received", types.YLeaf{"InvalidTimerReceived", statistics.InvalidTimerReceived})
    statistics.EntityData.Leafs.Append("mismatch-virtual-ip-address-received", types.YLeaf{"MismatchVirtualIpAddressReceived", statistics.MismatchVirtualIpAddressReceived})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Hsrp_Ipv4_Groups_Group_GlobalAddress
// Global virtual IPv6 addresses
type Hsrp_Ipv4_Groups_Group_GlobalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (globalAddress *Hsrp_Ipv4_Groups_Group_GlobalAddress) GetEntityData() *types.CommonEntityData {
    globalAddress.EntityData.YFilter = globalAddress.YFilter
    globalAddress.EntityData.YangName = "global-address"
    globalAddress.EntityData.BundleName = "cisco_ios_xr"
    globalAddress.EntityData.ParentYangName = "group"
    globalAddress.EntityData.SegmentPath = "global-address"
    globalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalAddress.EntityData.Children = types.NewOrderedMap()
    globalAddress.EntityData.Leafs = types.NewOrderedMap()
    globalAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", globalAddress.Ipv6Address})

    globalAddress.EntityData.YListKeys = []string {}

    return &(globalAddress.EntityData)
}

// Hsrp_Ipv4_Groups_Group_StateChangeHistory
// State change history
type Hsrp_Ipv4_Groups_Group_StateChangeHistory struct {
    EntityData types.CommonEntityData
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

func (stateChangeHistory *Hsrp_Ipv4_Groups_Group_StateChangeHistory) GetEntityData() *types.CommonEntityData {
    stateChangeHistory.EntityData.YFilter = stateChangeHistory.YFilter
    stateChangeHistory.EntityData.YangName = "state-change-history"
    stateChangeHistory.EntityData.BundleName = "cisco_ios_xr"
    stateChangeHistory.EntityData.ParentYangName = "group"
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

// Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time
// Time of state change
type Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (time *Hsrp_Ipv4_Groups_Group_StateChangeHistory_Time) GetEntityData() *types.CommonEntityData {
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

// Hsrp_Ipv4_TrackedInterfaces
// The HSRP tracked interfaces table
type Hsrp_Ipv4_TrackedInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An HSRP tracked interface entry. The type is slice of
    // Hsrp_Ipv4_TrackedInterfaces_TrackedInterface.
    TrackedInterface []*Hsrp_Ipv4_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Ipv4_TrackedInterfaces) GetEntityData() *types.CommonEntityData {
    trackedInterfaces.EntityData.YFilter = trackedInterfaces.YFilter
    trackedInterfaces.EntityData.YangName = "tracked-interfaces"
    trackedInterfaces.EntityData.BundleName = "cisco_ios_xr"
    trackedInterfaces.EntityData.ParentYangName = "ipv4"
    trackedInterfaces.EntityData.SegmentPath = "tracked-interfaces"
    trackedInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterfaces.EntityData.Children = types.NewOrderedMap()
    trackedInterfaces.EntityData.Children.Append("tracked-interface", types.YChild{"TrackedInterface", nil})
    for i := range trackedInterfaces.TrackedInterface {
        trackedInterfaces.EntityData.Children.Append(types.GetSegmentPath(trackedInterfaces.TrackedInterface[i]), types.YChild{"TrackedInterface", trackedInterfaces.TrackedInterface[i]})
    }
    trackedInterfaces.EntityData.Leafs = types.NewOrderedMap()

    trackedInterfaces.EntityData.YListKeys = []string {}

    return &(trackedInterfaces.EntityData)
}

// Hsrp_Ipv4_TrackedInterfaces_TrackedInterface
// An HSRP tracked interface entry
type Hsrp_Ipv4_TrackedInterfaces_TrackedInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name of the interface. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The HSRP group number. The type is interface{}
    // with range: 0..4294967295.
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

func (trackedInterface *Hsrp_Ipv4_TrackedInterfaces_TrackedInterface) GetEntityData() *types.CommonEntityData {
    trackedInterface.EntityData.YFilter = trackedInterface.YFilter
    trackedInterface.EntityData.YangName = "tracked-interface"
    trackedInterface.EntityData.BundleName = "cisco_ios_xr"
    trackedInterface.EntityData.ParentYangName = "tracked-interfaces"
    trackedInterface.EntityData.SegmentPath = "tracked-interface" + types.AddKeyToken(trackedInterface.InterfaceName, "interface-name") + types.AddKeyToken(trackedInterface.GroupNumber, "group-number") + types.AddKeyToken(trackedInterface.TrackedInterfaceName, "tracked-interface-name")
    trackedInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterface.EntityData.Children = types.NewOrderedMap()
    trackedInterface.EntityData.Leafs = types.NewOrderedMap()
    trackedInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", trackedInterface.InterfaceName})
    trackedInterface.EntityData.Leafs.Append("group-number", types.YLeaf{"GroupNumber", trackedInterface.GroupNumber})
    trackedInterface.EntityData.Leafs.Append("tracked-interface-name", types.YLeaf{"TrackedInterfaceName", trackedInterface.TrackedInterfaceName})
    trackedInterface.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", trackedInterface.Interface})
    trackedInterface.EntityData.Leafs.Append("hsrp-group-number", types.YLeaf{"HsrpGroupNumber", trackedInterface.HsrpGroupNumber})
    trackedInterface.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", trackedInterface.PriorityDecrement})
    trackedInterface.EntityData.Leafs.Append("interface-up-flag", types.YLeaf{"InterfaceUpFlag", trackedInterface.InterfaceUpFlag})
    trackedInterface.EntityData.Leafs.Append("tracked-interface-name-xr", types.YLeaf{"TrackedInterfaceNameXr", trackedInterface.TrackedInterfaceNameXr})
    trackedInterface.EntityData.Leafs.Append("is-object", types.YLeaf{"IsObject", trackedInterface.IsObject})

    trackedInterface.EntityData.YListKeys = []string {"InterfaceName", "GroupNumber", "TrackedInterfaceName"}

    return &(trackedInterface.EntityData)
}

// Hsrp_Ipv4_Interfaces
// The HSRP interface information table
type Hsrp_Ipv4_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A HSRP interface entry. The type is slice of
    // Hsrp_Ipv4_Interfaces_Interface.
    Interface []*Hsrp_Ipv4_Interfaces_Interface
}

func (interfaces *Hsrp_Ipv4_Interfaces) GetEntityData() *types.CommonEntityData {
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

// Hsrp_Ipv4_Interfaces_Interface
// A HSRP interface entry
type Hsrp_Ipv4_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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

func (self *Hsrp_Ipv4_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("statistics", types.YChild{"Statistics", &self.Statistics})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("use-bia-flag", types.YLeaf{"UseBiaFlag", self.UseBiaFlag})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Hsrp_Ipv4_Interfaces_Interface_Statistics
// HSRP Interface Statistics
type Hsrp_Ipv4_Interfaces_Interface_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Hsrp_Ipv4_Interfaces_Interface_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "interface"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("advert-packets-sent", types.YLeaf{"AdvertPacketsSent", statistics.AdvertPacketsSent})
    statistics.EntityData.Leafs.Append("advert-packets-received", types.YLeaf{"AdvertPacketsReceived", statistics.AdvertPacketsReceived})
    statistics.EntityData.Leafs.Append("long-packets-received", types.YLeaf{"LongPacketsReceived", statistics.LongPacketsReceived})
    statistics.EntityData.Leafs.Append("short-packets-received", types.YLeaf{"ShortPacketsReceived", statistics.ShortPacketsReceived})
    statistics.EntityData.Leafs.Append("invalid-version-received", types.YLeaf{"InvalidVersionReceived", statistics.InvalidVersionReceived})
    statistics.EntityData.Leafs.Append("invalid-operation-code-received", types.YLeaf{"InvalidOperationCodeReceived", statistics.InvalidOperationCodeReceived})
    statistics.EntityData.Leafs.Append("unknown-group-received", types.YLeaf{"UnknownGroupReceived", statistics.UnknownGroupReceived})
    statistics.EntityData.Leafs.Append("inoperational-group-received", types.YLeaf{"InoperationalGroupReceived", statistics.InoperationalGroupReceived})
    statistics.EntityData.Leafs.Append("conflict-source-ip-address-received", types.YLeaf{"ConflictSourceIpAddressReceived", statistics.ConflictSourceIpAddressReceived})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Hsrp_MgoSessions
// HSRP MGO session table
type Hsrp_MgoSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HSRP MGO session. The type is slice of Hsrp_MgoSessions_MgoSession.
    MgoSession []*Hsrp_MgoSessions_MgoSession
}

func (mgoSessions *Hsrp_MgoSessions) GetEntityData() *types.CommonEntityData {
    mgoSessions.EntityData.YFilter = mgoSessions.YFilter
    mgoSessions.EntityData.YangName = "mgo-sessions"
    mgoSessions.EntityData.BundleName = "cisco_ios_xr"
    mgoSessions.EntityData.ParentYangName = "hsrp"
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

// Hsrp_MgoSessions_MgoSession
// HSRP MGO session
type Hsrp_MgoSessions_MgoSession struct {
    EntityData types.CommonEntityData
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
    Slave []*Hsrp_MgoSessions_MgoSession_Slave
}

func (mgoSession *Hsrp_MgoSessions_MgoSession) GetEntityData() *types.CommonEntityData {
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

// Hsrp_MgoSessions_MgoSession_Slave
// List of slaves following this primary session
type Hsrp_MgoSessions_MgoSession_Slave struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface of slave group. The type is string with length: 0..64.
    SlaveGroupInterface interface{}

    // Group number of slave group. The type is interface{} with range:
    // 0..4294967295.
    SlaveGroupNumber interface{}
}

func (slave *Hsrp_MgoSessions_MgoSession_Slave) GetEntityData() *types.CommonEntityData {
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
    slave.EntityData.Leafs.Append("slave-group-interface", types.YLeaf{"SlaveGroupInterface", slave.SlaveGroupInterface})
    slave.EntityData.Leafs.Append("slave-group-number", types.YLeaf{"SlaveGroupNumber", slave.SlaveGroupNumber})

    slave.EntityData.YListKeys = []string {}

    return &(slave.EntityData)
}

// Hsrp_Ipv6
// IPv6 HSRP information
type Hsrp_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP tracked interfaces table.
    TrackedInterfaces Hsrp_Ipv6_TrackedInterfaces

    // The HSRP standby group table.
    Groups Hsrp_Ipv6_Groups

    // The HSRP interface information table.
    Interfaces Hsrp_Ipv6_Interfaces
}

func (ipv6 *Hsrp_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "hsrp"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("tracked-interfaces", types.YChild{"TrackedInterfaces", &ipv6.TrackedInterfaces})
    ipv6.EntityData.Children.Append("groups", types.YChild{"Groups", &ipv6.Groups})
    ipv6.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv6.Interfaces})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Hsrp_Ipv6_TrackedInterfaces
// The HSRP tracked interfaces table
type Hsrp_Ipv6_TrackedInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An HSRP tracked interface entry. The type is slice of
    // Hsrp_Ipv6_TrackedInterfaces_TrackedInterface.
    TrackedInterface []*Hsrp_Ipv6_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Ipv6_TrackedInterfaces) GetEntityData() *types.CommonEntityData {
    trackedInterfaces.EntityData.YFilter = trackedInterfaces.YFilter
    trackedInterfaces.EntityData.YangName = "tracked-interfaces"
    trackedInterfaces.EntityData.BundleName = "cisco_ios_xr"
    trackedInterfaces.EntityData.ParentYangName = "ipv6"
    trackedInterfaces.EntityData.SegmentPath = "tracked-interfaces"
    trackedInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterfaces.EntityData.Children = types.NewOrderedMap()
    trackedInterfaces.EntityData.Children.Append("tracked-interface", types.YChild{"TrackedInterface", nil})
    for i := range trackedInterfaces.TrackedInterface {
        trackedInterfaces.EntityData.Children.Append(types.GetSegmentPath(trackedInterfaces.TrackedInterface[i]), types.YChild{"TrackedInterface", trackedInterfaces.TrackedInterface[i]})
    }
    trackedInterfaces.EntityData.Leafs = types.NewOrderedMap()

    trackedInterfaces.EntityData.YListKeys = []string {}

    return &(trackedInterfaces.EntityData)
}

// Hsrp_Ipv6_TrackedInterfaces_TrackedInterface
// An HSRP tracked interface entry
type Hsrp_Ipv6_TrackedInterfaces_TrackedInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name of the interface. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The HSRP group number. The type is interface{}
    // with range: 0..4294967295.
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

func (trackedInterface *Hsrp_Ipv6_TrackedInterfaces_TrackedInterface) GetEntityData() *types.CommonEntityData {
    trackedInterface.EntityData.YFilter = trackedInterface.YFilter
    trackedInterface.EntityData.YangName = "tracked-interface"
    trackedInterface.EntityData.BundleName = "cisco_ios_xr"
    trackedInterface.EntityData.ParentYangName = "tracked-interfaces"
    trackedInterface.EntityData.SegmentPath = "tracked-interface" + types.AddKeyToken(trackedInterface.InterfaceName, "interface-name") + types.AddKeyToken(trackedInterface.GroupNumber, "group-number") + types.AddKeyToken(trackedInterface.TrackedInterfaceName, "tracked-interface-name")
    trackedInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterface.EntityData.Children = types.NewOrderedMap()
    trackedInterface.EntityData.Leafs = types.NewOrderedMap()
    trackedInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", trackedInterface.InterfaceName})
    trackedInterface.EntityData.Leafs.Append("group-number", types.YLeaf{"GroupNumber", trackedInterface.GroupNumber})
    trackedInterface.EntityData.Leafs.Append("tracked-interface-name", types.YLeaf{"TrackedInterfaceName", trackedInterface.TrackedInterfaceName})
    trackedInterface.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", trackedInterface.Interface})
    trackedInterface.EntityData.Leafs.Append("hsrp-group-number", types.YLeaf{"HsrpGroupNumber", trackedInterface.HsrpGroupNumber})
    trackedInterface.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", trackedInterface.PriorityDecrement})
    trackedInterface.EntityData.Leafs.Append("interface-up-flag", types.YLeaf{"InterfaceUpFlag", trackedInterface.InterfaceUpFlag})
    trackedInterface.EntityData.Leafs.Append("tracked-interface-name-xr", types.YLeaf{"TrackedInterfaceNameXr", trackedInterface.TrackedInterfaceNameXr})
    trackedInterface.EntityData.Leafs.Append("is-object", types.YLeaf{"IsObject", trackedInterface.IsObject})

    trackedInterface.EntityData.YListKeys = []string {"InterfaceName", "GroupNumber", "TrackedInterfaceName"}

    return &(trackedInterface.EntityData)
}

// Hsrp_Ipv6_Groups
// The HSRP standby group table
type Hsrp_Ipv6_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An HSRP standby group. The type is slice of Hsrp_Ipv6_Groups_Group.
    Group []*Hsrp_Ipv6_Groups_Group
}

func (groups *Hsrp_Ipv6_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "ipv6"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Hsrp_Ipv6_Groups_Group
// An HSRP standby group
type Hsrp_Ipv6_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. The HSRP group number. The type is interface{}
    // with range: 0..4294967295.
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
    GlobalAddress []*Hsrp_Ipv6_Groups_Group_GlobalAddress

    // State change history. The type is slice of
    // Hsrp_Ipv6_Groups_Group_StateChangeHistory.
    StateChangeHistory []*Hsrp_Ipv6_Groups_Group_StateChangeHistory
}

func (group *Hsrp_Ipv6_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.InterfaceName, "interface-name") + types.AddKeyToken(group.GroupNumber, "group-number")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("resign-sent-time", types.YChild{"ResignSentTime", &group.ResignSentTime})
    group.EntityData.Children.Append("resign-received-time", types.YChild{"ResignReceivedTime", &group.ResignReceivedTime})
    group.EntityData.Children.Append("coup-sent-time", types.YChild{"CoupSentTime", &group.CoupSentTime})
    group.EntityData.Children.Append("coup-received-time", types.YChild{"CoupReceivedTime", &group.CoupReceivedTime})
    group.EntityData.Children.Append("statistics", types.YChild{"Statistics", &group.Statistics})
    group.EntityData.Children.Append("global-address", types.YChild{"GlobalAddress", nil})
    for i := range group.GlobalAddress {
        group.EntityData.Children.Append(types.GetSegmentPath(group.GlobalAddress[i]), types.YChild{"GlobalAddress", group.GlobalAddress[i]})
    }
    group.EntityData.Children.Append("state-change-history", types.YChild{"StateChangeHistory", nil})
    for i := range group.StateChangeHistory {
        group.EntityData.Children.Append(types.GetSegmentPath(group.StateChangeHistory[i]), types.YChild{"StateChangeHistory", group.StateChangeHistory[i]})
    }
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", group.InterfaceName})
    group.EntityData.Leafs.Append("group-number", types.YLeaf{"GroupNumber", group.GroupNumber})
    group.EntityData.Leafs.Append("authentication-string", types.YLeaf{"AuthenticationString", group.AuthenticationString})
    group.EntityData.Leafs.Append("virtual-mac-address", types.YLeaf{"VirtualMacAddress", group.VirtualMacAddress})
    group.EntityData.Leafs.Append("hsrp-group-number", types.YLeaf{"HsrpGroupNumber", group.HsrpGroupNumber})
    group.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", group.AddressFamily})
    group.EntityData.Leafs.Append("version", types.YLeaf{"Version", group.Version})
    group.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", group.SessionName})
    group.EntityData.Leafs.Append("slaves", types.YLeaf{"Slaves", group.Slaves})
    group.EntityData.Leafs.Append("is-slave", types.YLeaf{"IsSlave", group.IsSlave})
    group.EntityData.Leafs.Append("followed-session-name", types.YLeaf{"FollowedSessionName", group.FollowedSessionName})
    group.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", group.ConfiguredPriority})
    group.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", group.PreemptDelay})
    group.EntityData.Leafs.Append("preempt-timer-secs", types.YLeaf{"PreemptTimerSecs", group.PreemptTimerSecs})
    group.EntityData.Leafs.Append("hello-time", types.YLeaf{"HelloTime", group.HelloTime})
    group.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", group.HoldTime})
    group.EntityData.Leafs.Append("learned-hello-time", types.YLeaf{"LearnedHelloTime", group.LearnedHelloTime})
    group.EntityData.Leafs.Append("learned-hold-time", types.YLeaf{"LearnedHoldTime", group.LearnedHoldTime})
    group.EntityData.Leafs.Append("min-delay-time", types.YLeaf{"MinDelayTime", group.MinDelayTime})
    group.EntityData.Leafs.Append("reload-delay-time", types.YLeaf{"ReloadDelayTime", group.ReloadDelayTime})
    group.EntityData.Leafs.Append("virtual-ip-address", types.YLeaf{"VirtualIpAddress", group.VirtualIpAddress})
    group.EntityData.Leafs.Append("virtual-linklocal-ipv6-address", types.YLeaf{"VirtualLinklocalIpv6Address", group.VirtualLinklocalIpv6Address})
    group.EntityData.Leafs.Append("active-ip-address", types.YLeaf{"ActiveIpAddress", group.ActiveIpAddress})
    group.EntityData.Leafs.Append("active-ipv6-address", types.YLeaf{"ActiveIpv6Address", group.ActiveIpv6Address})
    group.EntityData.Leafs.Append("active-mac-address", types.YLeaf{"ActiveMacAddress", group.ActiveMacAddress})
    group.EntityData.Leafs.Append("standby-ip-address", types.YLeaf{"StandbyIpAddress", group.StandbyIpAddress})
    group.EntityData.Leafs.Append("standby-ipv6-address", types.YLeaf{"StandbyIpv6Address", group.StandbyIpv6Address})
    group.EntityData.Leafs.Append("standby-mac-address", types.YLeaf{"StandbyMacAddress", group.StandbyMacAddress})
    group.EntityData.Leafs.Append("hsrp-router-state", types.YLeaf{"HsrpRouterState", group.HsrpRouterState})
    group.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", group.InterfaceNameXr})
    group.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", group.Interface})
    group.EntityData.Leafs.Append("router-priority", types.YLeaf{"RouterPriority", group.RouterPriority})
    group.EntityData.Leafs.Append("active-priority", types.YLeaf{"ActivePriority", group.ActivePriority})
    group.EntityData.Leafs.Append("active-timer-flag", types.YLeaf{"ActiveTimerFlag", group.ActiveTimerFlag})
    group.EntityData.Leafs.Append("active-timer-secs", types.YLeaf{"ActiveTimerSecs", group.ActiveTimerSecs})
    group.EntityData.Leafs.Append("active-timer-msecs", types.YLeaf{"ActiveTimerMsecs", group.ActiveTimerMsecs})
    group.EntityData.Leafs.Append("standby-timer-flag", types.YLeaf{"StandbyTimerFlag", group.StandbyTimerFlag})
    group.EntityData.Leafs.Append("standby-timer-secs", types.YLeaf{"StandbyTimerSecs", group.StandbyTimerSecs})
    group.EntityData.Leafs.Append("standby-timer-msecs", types.YLeaf{"StandbyTimerMsecs", group.StandbyTimerMsecs})
    group.EntityData.Leafs.Append("hello-timer-flag", types.YLeaf{"HelloTimerFlag", group.HelloTimerFlag})
    group.EntityData.Leafs.Append("hello-timer-secs", types.YLeaf{"HelloTimerSecs", group.HelloTimerSecs})
    group.EntityData.Leafs.Append("hello-timer-msecs", types.YLeaf{"HelloTimerMsecs", group.HelloTimerMsecs})
    group.EntityData.Leafs.Append("delay-timer-flag", types.YLeaf{"DelayTimerFlag", group.DelayTimerFlag})
    group.EntityData.Leafs.Append("delay-timer-secs", types.YLeaf{"DelayTimerSecs", group.DelayTimerSecs})
    group.EntityData.Leafs.Append("delay-timer-msecs", types.YLeaf{"DelayTimerMsecs", group.DelayTimerMsecs})
    group.EntityData.Leafs.Append("current-state-timer-secs", types.YLeaf{"CurrentStateTimerSecs", group.CurrentStateTimerSecs})
    group.EntityData.Leafs.Append("state-change-count", types.YLeaf{"StateChangeCount", group.StateChangeCount})
    group.EntityData.Leafs.Append("tracked-interface-count", types.YLeaf{"TrackedInterfaceCount", group.TrackedInterfaceCount})
    group.EntityData.Leafs.Append("tracked-interface-up-count", types.YLeaf{"TrackedInterfaceUpCount", group.TrackedInterfaceUpCount})
    group.EntityData.Leafs.Append("preempt-enabled", types.YLeaf{"PreemptEnabled", group.PreemptEnabled})
    group.EntityData.Leafs.Append("use-configured-timers", types.YLeaf{"UseConfiguredTimers", group.UseConfiguredTimers})
    group.EntityData.Leafs.Append("use-configured-virtual-ip", types.YLeaf{"UseConfiguredVirtualIp", group.UseConfiguredVirtualIp})
    group.EntityData.Leafs.Append("use-bia-configured", types.YLeaf{"UseBiaConfigured", group.UseBiaConfigured})
    group.EntityData.Leafs.Append("configured-timers", types.YLeaf{"ConfiguredTimers", group.ConfiguredTimers})
    group.EntityData.Leafs.Append("configured-mac-address", types.YLeaf{"ConfiguredMacAddress", group.ConfiguredMacAddress})
    group.EntityData.Leafs.Append("redirects-disabled", types.YLeaf{"RedirectsDisabled", group.RedirectsDisabled})
    group.EntityData.Leafs.Append("bfd-enabled", types.YLeaf{"BfdEnabled", group.BfdEnabled})
    group.EntityData.Leafs.Append("bfd-interface", types.YLeaf{"BfdInterface", group.BfdInterface})
    group.EntityData.Leafs.Append("bfd-peer-ip-address", types.YLeaf{"BfdPeerIpAddress", group.BfdPeerIpAddress})
    group.EntityData.Leafs.Append("bfd-peer-ipv6-address", types.YLeaf{"BfdPeerIpv6Address", group.BfdPeerIpv6Address})
    group.EntityData.Leafs.Append("bfd-session-state", types.YLeaf{"BfdSessionState", group.BfdSessionState})
    group.EntityData.Leafs.Append("bfd-interval", types.YLeaf{"BfdInterval", group.BfdInterval})
    group.EntityData.Leafs.Append("bfd-multiplier", types.YLeaf{"BfdMultiplier", group.BfdMultiplier})
    group.EntityData.Leafs.Append("virtual-mac-address-state", types.YLeaf{"VirtualMacAddressState", group.VirtualMacAddressState})
    group.EntityData.Leafs.Append("secondary-address", types.YLeaf{"SecondaryAddress", group.SecondaryAddress})

    group.EntityData.YListKeys = []string {"InterfaceName", "GroupNumber"}

    return &(group.EntityData)
}

// Hsrp_Ipv6_Groups_Group_ResignSentTime
// Time last resign was sent
type Hsrp_Ipv6_Groups_Group_ResignSentTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignSentTime *Hsrp_Ipv6_Groups_Group_ResignSentTime) GetEntityData() *types.CommonEntityData {
    resignSentTime.EntityData.YFilter = resignSentTime.YFilter
    resignSentTime.EntityData.YangName = "resign-sent-time"
    resignSentTime.EntityData.BundleName = "cisco_ios_xr"
    resignSentTime.EntityData.ParentYangName = "group"
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

// Hsrp_Ipv6_Groups_Group_ResignReceivedTime
// Time last resign was received
type Hsrp_Ipv6_Groups_Group_ResignReceivedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resignReceivedTime *Hsrp_Ipv6_Groups_Group_ResignReceivedTime) GetEntityData() *types.CommonEntityData {
    resignReceivedTime.EntityData.YFilter = resignReceivedTime.YFilter
    resignReceivedTime.EntityData.YangName = "resign-received-time"
    resignReceivedTime.EntityData.BundleName = "cisco_ios_xr"
    resignReceivedTime.EntityData.ParentYangName = "group"
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

// Hsrp_Ipv6_Groups_Group_CoupSentTime
// Time last coup was sent
type Hsrp_Ipv6_Groups_Group_CoupSentTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (coupSentTime *Hsrp_Ipv6_Groups_Group_CoupSentTime) GetEntityData() *types.CommonEntityData {
    coupSentTime.EntityData.YFilter = coupSentTime.YFilter
    coupSentTime.EntityData.YangName = "coup-sent-time"
    coupSentTime.EntityData.BundleName = "cisco_ios_xr"
    coupSentTime.EntityData.ParentYangName = "group"
    coupSentTime.EntityData.SegmentPath = "coup-sent-time"
    coupSentTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coupSentTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coupSentTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coupSentTime.EntityData.Children = types.NewOrderedMap()
    coupSentTime.EntityData.Leafs = types.NewOrderedMap()
    coupSentTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", coupSentTime.Seconds})
    coupSentTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", coupSentTime.Nanoseconds})

    coupSentTime.EntityData.YListKeys = []string {}

    return &(coupSentTime.EntityData)
}

// Hsrp_Ipv6_Groups_Group_CoupReceivedTime
// Time last coup was received
type Hsrp_Ipv6_Groups_Group_CoupReceivedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (coupReceivedTime *Hsrp_Ipv6_Groups_Group_CoupReceivedTime) GetEntityData() *types.CommonEntityData {
    coupReceivedTime.EntityData.YFilter = coupReceivedTime.YFilter
    coupReceivedTime.EntityData.YangName = "coup-received-time"
    coupReceivedTime.EntityData.BundleName = "cisco_ios_xr"
    coupReceivedTime.EntityData.ParentYangName = "group"
    coupReceivedTime.EntityData.SegmentPath = "coup-received-time"
    coupReceivedTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coupReceivedTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coupReceivedTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coupReceivedTime.EntityData.Children = types.NewOrderedMap()
    coupReceivedTime.EntityData.Leafs = types.NewOrderedMap()
    coupReceivedTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", coupReceivedTime.Seconds})
    coupReceivedTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", coupReceivedTime.Nanoseconds})

    coupReceivedTime.EntityData.YListKeys = []string {}

    return &(coupReceivedTime.EntityData)
}

// Hsrp_Ipv6_Groups_Group_Statistics
// HSRP Group statistics
type Hsrp_Ipv6_Groups_Group_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Hsrp_Ipv6_Groups_Group_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "group"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("active-transitions", types.YLeaf{"ActiveTransitions", statistics.ActiveTransitions})
    statistics.EntityData.Leafs.Append("standby-transitions", types.YLeaf{"StandbyTransitions", statistics.StandbyTransitions})
    statistics.EntityData.Leafs.Append("speak-transitions", types.YLeaf{"SpeakTransitions", statistics.SpeakTransitions})
    statistics.EntityData.Leafs.Append("listen-transitions", types.YLeaf{"ListenTransitions", statistics.ListenTransitions})
    statistics.EntityData.Leafs.Append("learn-transitions", types.YLeaf{"LearnTransitions", statistics.LearnTransitions})
    statistics.EntityData.Leafs.Append("init-transitions", types.YLeaf{"InitTransitions", statistics.InitTransitions})
    statistics.EntityData.Leafs.Append("hello-packets-sent", types.YLeaf{"HelloPacketsSent", statistics.HelloPacketsSent})
    statistics.EntityData.Leafs.Append("resign-packets-sent", types.YLeaf{"ResignPacketsSent", statistics.ResignPacketsSent})
    statistics.EntityData.Leafs.Append("coup-packets-sent", types.YLeaf{"CoupPacketsSent", statistics.CoupPacketsSent})
    statistics.EntityData.Leafs.Append("hello-packets-received", types.YLeaf{"HelloPacketsReceived", statistics.HelloPacketsReceived})
    statistics.EntityData.Leafs.Append("resign-packets-received", types.YLeaf{"ResignPacketsReceived", statistics.ResignPacketsReceived})
    statistics.EntityData.Leafs.Append("coup-packets-received", types.YLeaf{"CoupPacketsReceived", statistics.CoupPacketsReceived})
    statistics.EntityData.Leafs.Append("auth-fail-received", types.YLeaf{"AuthFailReceived", statistics.AuthFailReceived})
    statistics.EntityData.Leafs.Append("invalid-timer-received", types.YLeaf{"InvalidTimerReceived", statistics.InvalidTimerReceived})
    statistics.EntityData.Leafs.Append("mismatch-virtual-ip-address-received", types.YLeaf{"MismatchVirtualIpAddressReceived", statistics.MismatchVirtualIpAddressReceived})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Hsrp_Ipv6_Groups_Group_GlobalAddress
// Global virtual IPv6 addresses
type Hsrp_Ipv6_Groups_Group_GlobalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV6Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (globalAddress *Hsrp_Ipv6_Groups_Group_GlobalAddress) GetEntityData() *types.CommonEntityData {
    globalAddress.EntityData.YFilter = globalAddress.YFilter
    globalAddress.EntityData.YangName = "global-address"
    globalAddress.EntityData.BundleName = "cisco_ios_xr"
    globalAddress.EntityData.ParentYangName = "group"
    globalAddress.EntityData.SegmentPath = "global-address"
    globalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalAddress.EntityData.Children = types.NewOrderedMap()
    globalAddress.EntityData.Leafs = types.NewOrderedMap()
    globalAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", globalAddress.Ipv6Address})

    globalAddress.EntityData.YListKeys = []string {}

    return &(globalAddress.EntityData)
}

// Hsrp_Ipv6_Groups_Group_StateChangeHistory
// State change history
type Hsrp_Ipv6_Groups_Group_StateChangeHistory struct {
    EntityData types.CommonEntityData
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

func (stateChangeHistory *Hsrp_Ipv6_Groups_Group_StateChangeHistory) GetEntityData() *types.CommonEntityData {
    stateChangeHistory.EntityData.YFilter = stateChangeHistory.YFilter
    stateChangeHistory.EntityData.YangName = "state-change-history"
    stateChangeHistory.EntityData.BundleName = "cisco_ios_xr"
    stateChangeHistory.EntityData.ParentYangName = "group"
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

// Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time
// Time of state change
type Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (time *Hsrp_Ipv6_Groups_Group_StateChangeHistory_Time) GetEntityData() *types.CommonEntityData {
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

// Hsrp_Ipv6_Interfaces
// The HSRP interface information table
type Hsrp_Ipv6_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A HSRP interface entry. The type is slice of
    // Hsrp_Ipv6_Interfaces_Interface.
    Interface []*Hsrp_Ipv6_Interfaces_Interface
}

func (interfaces *Hsrp_Ipv6_Interfaces) GetEntityData() *types.CommonEntityData {
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

// Hsrp_Ipv6_Interfaces_Interface
// A HSRP interface entry
type Hsrp_Ipv6_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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

func (self *Hsrp_Ipv6_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("statistics", types.YChild{"Statistics", &self.Statistics})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("use-bia-flag", types.YLeaf{"UseBiaFlag", self.UseBiaFlag})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Hsrp_Ipv6_Interfaces_Interface_Statistics
// HSRP Interface Statistics
type Hsrp_Ipv6_Interfaces_Interface_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Hsrp_Ipv6_Interfaces_Interface_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "interface"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("advert-packets-sent", types.YLeaf{"AdvertPacketsSent", statistics.AdvertPacketsSent})
    statistics.EntityData.Leafs.Append("advert-packets-received", types.YLeaf{"AdvertPacketsReceived", statistics.AdvertPacketsReceived})
    statistics.EntityData.Leafs.Append("long-packets-received", types.YLeaf{"LongPacketsReceived", statistics.LongPacketsReceived})
    statistics.EntityData.Leafs.Append("short-packets-received", types.YLeaf{"ShortPacketsReceived", statistics.ShortPacketsReceived})
    statistics.EntityData.Leafs.Append("invalid-version-received", types.YLeaf{"InvalidVersionReceived", statistics.InvalidVersionReceived})
    statistics.EntityData.Leafs.Append("invalid-operation-code-received", types.YLeaf{"InvalidOperationCodeReceived", statistics.InvalidOperationCodeReceived})
    statistics.EntityData.Leafs.Append("unknown-group-received", types.YLeaf{"UnknownGroupReceived", statistics.UnknownGroupReceived})
    statistics.EntityData.Leafs.Append("inoperational-group-received", types.YLeaf{"InoperationalGroupReceived", statistics.InoperationalGroupReceived})
    statistics.EntityData.Leafs.Append("conflict-source-ip-address-received", types.YLeaf{"ConflictSourceIpAddressReceived", statistics.ConflictSourceIpAddressReceived})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Hsrp_BfdSessions
// The table of HSRP BFD Sessions
type Hsrp_BfdSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An HSRP BFD Session. The type is slice of Hsrp_BfdSessions_BfdSession.
    BfdSession []*Hsrp_BfdSessions_BfdSession
}

func (bfdSessions *Hsrp_BfdSessions) GetEntityData() *types.CommonEntityData {
    bfdSessions.EntityData.YFilter = bfdSessions.YFilter
    bfdSessions.EntityData.YangName = "bfd-sessions"
    bfdSessions.EntityData.BundleName = "cisco_ios_xr"
    bfdSessions.EntityData.ParentYangName = "hsrp"
    bfdSessions.EntityData.SegmentPath = "bfd-sessions"
    bfdSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdSessions.EntityData.Children = types.NewOrderedMap()
    bfdSessions.EntityData.Children.Append("bfd-session", types.YChild{"BfdSession", nil})
    for i := range bfdSessions.BfdSession {
        bfdSessions.EntityData.Children.Append(types.GetSegmentPath(bfdSessions.BfdSession[i]), types.YChild{"BfdSession", bfdSessions.BfdSession[i]})
    }
    bfdSessions.EntityData.Leafs = types.NewOrderedMap()

    bfdSessions.EntityData.YListKeys = []string {}

    return &(bfdSessions.EntityData)
}

// Hsrp_BfdSessions_BfdSession
// An HSRP BFD Session
type Hsrp_BfdSessions_BfdSession struct {
    EntityData types.CommonEntityData
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
    Group []*Hsrp_BfdSessions_BfdSession_Group
}

func (bfdSession *Hsrp_BfdSessions_BfdSession) GetEntityData() *types.CommonEntityData {
    bfdSession.EntityData.YFilter = bfdSession.YFilter
    bfdSession.EntityData.YangName = "bfd-session"
    bfdSession.EntityData.BundleName = "cisco_ios_xr"
    bfdSession.EntityData.ParentYangName = "bfd-sessions"
    bfdSession.EntityData.SegmentPath = "bfd-session" + types.AddKeyToken(bfdSession.InterfaceName, "interface-name") + types.AddKeyToken(bfdSession.IpAddress, "ip-address")
    bfdSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdSession.EntityData.Children = types.NewOrderedMap()
    bfdSession.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range bfdSession.Group {
        bfdSession.EntityData.Children.Append(types.GetSegmentPath(bfdSession.Group[i]), types.YChild{"Group", bfdSession.Group[i]})
    }
    bfdSession.EntityData.Leafs = types.NewOrderedMap()
    bfdSession.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", bfdSession.InterfaceName})
    bfdSession.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", bfdSession.IpAddress})
    bfdSession.EntityData.Leafs.Append("bfd-interface-name", types.YLeaf{"BfdInterfaceName", bfdSession.BfdInterfaceName})
    bfdSession.EntityData.Leafs.Append("session-address-family", types.YLeaf{"SessionAddressFamily", bfdSession.SessionAddressFamily})
    bfdSession.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", bfdSession.DestinationAddress})
    bfdSession.EntityData.Leafs.Append("destination-ipv6-address", types.YLeaf{"DestinationIpv6Address", bfdSession.DestinationIpv6Address})
    bfdSession.EntityData.Leafs.Append("bfd-session-state", types.YLeaf{"BfdSessionState", bfdSession.BfdSessionState})
    bfdSession.EntityData.Leafs.Append("bfd-interval", types.YLeaf{"BfdInterval", bfdSession.BfdInterval})
    bfdSession.EntityData.Leafs.Append("bfd-multiplier", types.YLeaf{"BfdMultiplier", bfdSession.BfdMultiplier})

    bfdSession.EntityData.YListKeys = []string {"InterfaceName", "IpAddress"}

    return &(bfdSession.EntityData)
}

// Hsrp_BfdSessions_BfdSession_Group
// HSRP Groups tracking the BFD session
type Hsrp_BfdSessions_BfdSession_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // HSRP Group number. The type is interface{} with range: 0..4294967295.
    HsrpGroupNumber interface{}
}

func (group *Hsrp_BfdSessions_BfdSession_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "bfd-session"
    group.EntityData.SegmentPath = "group"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", group.InterfaceName})
    group.EntityData.Leafs.Append("hsrp-group-number", types.YLeaf{"HsrpGroupNumber", group.HsrpGroupNumber})

    group.EntityData.YListKeys = []string {}

    return &(group.EntityData)
}

// Hsrp_Summary
// HSRP summary statistics
type Hsrp_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *Hsrp_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "hsrp"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("ipv4-sessions-active", types.YLeaf{"Ipv4SessionsActive", summary.Ipv4SessionsActive})
    summary.EntityData.Leafs.Append("ipv4-sessions-standby", types.YLeaf{"Ipv4SessionsStandby", summary.Ipv4SessionsStandby})
    summary.EntityData.Leafs.Append("ipv4-sessions-speak", types.YLeaf{"Ipv4SessionsSpeak", summary.Ipv4SessionsSpeak})
    summary.EntityData.Leafs.Append("ipv4-sessions-listen", types.YLeaf{"Ipv4SessionsListen", summary.Ipv4SessionsListen})
    summary.EntityData.Leafs.Append("ipv4-sessions-learn", types.YLeaf{"Ipv4SessionsLearn", summary.Ipv4SessionsLearn})
    summary.EntityData.Leafs.Append("ipv4-sessions-init", types.YLeaf{"Ipv4SessionsInit", summary.Ipv4SessionsInit})
    summary.EntityData.Leafs.Append("ipv4-slaves-active", types.YLeaf{"Ipv4SlavesActive", summary.Ipv4SlavesActive})
    summary.EntityData.Leafs.Append("ipv4-slaves-standby", types.YLeaf{"Ipv4SlavesStandby", summary.Ipv4SlavesStandby})
    summary.EntityData.Leafs.Append("ipv4-slaves-speak", types.YLeaf{"Ipv4SlavesSpeak", summary.Ipv4SlavesSpeak})
    summary.EntityData.Leafs.Append("ipv4-slaves-listen", types.YLeaf{"Ipv4SlavesListen", summary.Ipv4SlavesListen})
    summary.EntityData.Leafs.Append("ipv4-slaves-learn", types.YLeaf{"Ipv4SlavesLearn", summary.Ipv4SlavesLearn})
    summary.EntityData.Leafs.Append("ipv4-slaves-init", types.YLeaf{"Ipv4SlavesInit", summary.Ipv4SlavesInit})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-active-up", types.YLeaf{"Ipv4VirtualIpAddressesActiveUp", summary.Ipv4VirtualIpAddressesActiveUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-active-down", types.YLeaf{"Ipv4VirtualIpAddressesActiveDown", summary.Ipv4VirtualIpAddressesActiveDown})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-standby-up", types.YLeaf{"Ipv4VirtualIpAddressesStandbyUp", summary.Ipv4VirtualIpAddressesStandbyUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-standby-down", types.YLeaf{"Ipv4VirtualIpAddressesStandbyDown", summary.Ipv4VirtualIpAddressesStandbyDown})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-speak-up", types.YLeaf{"Ipv4VirtualIpAddressesSpeakUp", summary.Ipv4VirtualIpAddressesSpeakUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-speak-down", types.YLeaf{"Ipv4VirtualIpAddressesSpeakDown", summary.Ipv4VirtualIpAddressesSpeakDown})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-listen-up", types.YLeaf{"Ipv4VirtualIpAddressesListenUp", summary.Ipv4VirtualIpAddressesListenUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-listen-down", types.YLeaf{"Ipv4VirtualIpAddressesListenDown", summary.Ipv4VirtualIpAddressesListenDown})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-learn-up", types.YLeaf{"Ipv4VirtualIpAddressesLearnUp", summary.Ipv4VirtualIpAddressesLearnUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-learn-down", types.YLeaf{"Ipv4VirtualIpAddressesLearnDown", summary.Ipv4VirtualIpAddressesLearnDown})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-init-up", types.YLeaf{"Ipv4VirtualIpAddressesInitUp", summary.Ipv4VirtualIpAddressesInitUp})
    summary.EntityData.Leafs.Append("ipv4-virtual-ip-addresses-init-down", types.YLeaf{"Ipv4VirtualIpAddressesInitDown", summary.Ipv4VirtualIpAddressesInitDown})
    summary.EntityData.Leafs.Append("ipv6-sessions-active", types.YLeaf{"Ipv6SessionsActive", summary.Ipv6SessionsActive})
    summary.EntityData.Leafs.Append("ipv6-sessions-standby", types.YLeaf{"Ipv6SessionsStandby", summary.Ipv6SessionsStandby})
    summary.EntityData.Leafs.Append("ipv6-sessions-speak", types.YLeaf{"Ipv6SessionsSpeak", summary.Ipv6SessionsSpeak})
    summary.EntityData.Leafs.Append("ipv6-sessions-listen", types.YLeaf{"Ipv6SessionsListen", summary.Ipv6SessionsListen})
    summary.EntityData.Leafs.Append("ipv6-sessions-learn", types.YLeaf{"Ipv6SessionsLearn", summary.Ipv6SessionsLearn})
    summary.EntityData.Leafs.Append("ipv6-sessions-init", types.YLeaf{"Ipv6SessionsInit", summary.Ipv6SessionsInit})
    summary.EntityData.Leafs.Append("ipv6-slaves-active", types.YLeaf{"Ipv6SlavesActive", summary.Ipv6SlavesActive})
    summary.EntityData.Leafs.Append("ipv6-slaves-standby", types.YLeaf{"Ipv6SlavesStandby", summary.Ipv6SlavesStandby})
    summary.EntityData.Leafs.Append("ipv6-slaves-speak", types.YLeaf{"Ipv6SlavesSpeak", summary.Ipv6SlavesSpeak})
    summary.EntityData.Leafs.Append("ipv6-slaves-listen", types.YLeaf{"Ipv6SlavesListen", summary.Ipv6SlavesListen})
    summary.EntityData.Leafs.Append("ipv6-slaves-learn", types.YLeaf{"Ipv6SlavesLearn", summary.Ipv6SlavesLearn})
    summary.EntityData.Leafs.Append("ipv6-slaves-init", types.YLeaf{"Ipv6SlavesInit", summary.Ipv6SlavesInit})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-active-up", types.YLeaf{"Ipv6VirtualIpAddressesActiveUp", summary.Ipv6VirtualIpAddressesActiveUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-active-down", types.YLeaf{"Ipv6VirtualIpAddressesActiveDown", summary.Ipv6VirtualIpAddressesActiveDown})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-standby-up", types.YLeaf{"Ipv6VirtualIpAddressesStandbyUp", summary.Ipv6VirtualIpAddressesStandbyUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-standby-down", types.YLeaf{"Ipv6VirtualIpAddressesStandbyDown", summary.Ipv6VirtualIpAddressesStandbyDown})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-speak-up", types.YLeaf{"Ipv6VirtualIpAddressesSpeakUp", summary.Ipv6VirtualIpAddressesSpeakUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-speak-down", types.YLeaf{"Ipv6VirtualIpAddressesSpeakDown", summary.Ipv6VirtualIpAddressesSpeakDown})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-listen-up", types.YLeaf{"Ipv6VirtualIpAddressesListenUp", summary.Ipv6VirtualIpAddressesListenUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-listen-down", types.YLeaf{"Ipv6VirtualIpAddressesListenDown", summary.Ipv6VirtualIpAddressesListenDown})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-learn-up", types.YLeaf{"Ipv6VirtualIpAddressesLearnUp", summary.Ipv6VirtualIpAddressesLearnUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-learn-down", types.YLeaf{"Ipv6VirtualIpAddressesLearnDown", summary.Ipv6VirtualIpAddressesLearnDown})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-init-up", types.YLeaf{"Ipv6VirtualIpAddressesInitUp", summary.Ipv6VirtualIpAddressesInitUp})
    summary.EntityData.Leafs.Append("ipv6-virtual-ip-addresses-init-down", types.YLeaf{"Ipv6VirtualIpAddressesInitDown", summary.Ipv6VirtualIpAddressesInitDown})
    summary.EntityData.Leafs.Append("interfaces-ipv4-state-up", types.YLeaf{"InterfacesIpv4StateUp", summary.InterfacesIpv4StateUp})
    summary.EntityData.Leafs.Append("interfaces-ipv4-state-down", types.YLeaf{"InterfacesIpv4StateDown", summary.InterfacesIpv4StateDown})
    summary.EntityData.Leafs.Append("tracked-interfaces-ipv4-state-up", types.YLeaf{"TrackedInterfacesIpv4StateUp", summary.TrackedInterfacesIpv4StateUp})
    summary.EntityData.Leafs.Append("tracked-interfaces-ipv4-state-down", types.YLeaf{"TrackedInterfacesIpv4StateDown", summary.TrackedInterfacesIpv4StateDown})
    summary.EntityData.Leafs.Append("tracked-objects-up", types.YLeaf{"TrackedObjectsUp", summary.TrackedObjectsUp})
    summary.EntityData.Leafs.Append("tracked-objects-down", types.YLeaf{"TrackedObjectsDown", summary.TrackedObjectsDown})
    summary.EntityData.Leafs.Append("interfaces-ipv6-state-up", types.YLeaf{"InterfacesIpv6StateUp", summary.InterfacesIpv6StateUp})
    summary.EntityData.Leafs.Append("interfaces-ipv6-state-down", types.YLeaf{"InterfacesIpv6StateDown", summary.InterfacesIpv6StateDown})
    summary.EntityData.Leafs.Append("tracked-interfaces-ipv6-state-up", types.YLeaf{"TrackedInterfacesIpv6StateUp", summary.TrackedInterfacesIpv6StateUp})
    summary.EntityData.Leafs.Append("tracked-interfaces-ipv6-state-down", types.YLeaf{"TrackedInterfacesIpv6StateDown", summary.TrackedInterfacesIpv6StateDown})
    summary.EntityData.Leafs.Append("bfd-sessions-up", types.YLeaf{"BfdSessionsUp", summary.BfdSessionsUp})
    summary.EntityData.Leafs.Append("bfd-sessions-down", types.YLeaf{"BfdSessionsDown", summary.BfdSessionsDown})
    summary.EntityData.Leafs.Append("bfd-session-inactive", types.YLeaf{"BfdSessionInactive", summary.BfdSessionInactive})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

