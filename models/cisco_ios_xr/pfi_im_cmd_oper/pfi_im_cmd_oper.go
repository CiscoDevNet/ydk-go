// This module contains a collection of YANG definitions
// for Cisco IOS-XR pfi-im-cmd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   interfaces: Interface operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package pfi_im_cmd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pfi_im_cmd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pfi-im-cmd-oper interfaces}", reflect.TypeOf(Interfaces{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pfi-im-cmd-oper:interfaces", reflect.TypeOf(Interfaces{}))
}

// ImCmdIntfTypeEnum represents Im cmd intf type enum
type ImCmdIntfTypeEnum string

const (
    // srp
    ImCmdIntfTypeEnum_srp ImCmdIntfTypeEnum = "srp"

    // tunnel
    ImCmdIntfTypeEnum_tunnel ImCmdIntfTypeEnum = "tunnel"

    // bundle
    ImCmdIntfTypeEnum_bundle ImCmdIntfTypeEnum = "bundle"

    // serial
    ImCmdIntfTypeEnum_serial ImCmdIntfTypeEnum = "serial"

    // sonet pos
    ImCmdIntfTypeEnum_sonet_pos ImCmdIntfTypeEnum = "sonet-pos"

    // tunnel gre
    ImCmdIntfTypeEnum_tunnel_gre ImCmdIntfTypeEnum = "tunnel-gre"

    // pseudowire head end
    ImCmdIntfTypeEnum_pseudowire_head_end ImCmdIntfTypeEnum = "pseudowire-head-end"

    // cem
    ImCmdIntfTypeEnum_cem ImCmdIntfTypeEnum = "cem"

    // gcc
    ImCmdIntfTypeEnum_gcc ImCmdIntfTypeEnum = "gcc"
)

// VlanSwitchedMode represents VLAN-Switched mode
type VlanSwitchedMode string

const (
    // Disabled
    VlanSwitchedMode_none VlanSwitchedMode = "none"

    // Trunk port
    VlanSwitchedMode_trunk_port VlanSwitchedMode = "trunk-port"

    // Access port
    VlanSwitchedMode_access_port VlanSwitchedMode = "access-port"
)

// ImCmdStatsEnum represents List of different interface stats structures
type ImCmdStatsEnum string

const (
    // full
    ImCmdStatsEnum_full ImCmdStatsEnum = "full"

    // basic
    ImCmdStatsEnum_basic ImCmdStatsEnum = "basic"
)

// SrpMgmtFailureStateEt represents SRP failure state type
type SrpMgmtFailureStateEt string

const (
    // Idle
    SrpMgmtFailureStateEt_idle_failure_state SrpMgmtFailureStateEt = "idle-failure-state"

    // Wait To Restore
    SrpMgmtFailureStateEt_wait_to_restore_failure_state SrpMgmtFailureStateEt = "wait-to-restore-failure-state"

    // Manual Switch
    SrpMgmtFailureStateEt_manual_switch_failure_state SrpMgmtFailureStateEt = "manual-switch-failure-state"

    // Signal Degrade
    SrpMgmtFailureStateEt_signal_degrade_failure_state SrpMgmtFailureStateEt = "signal-degrade-failure-state"

    // Signal Fail
    SrpMgmtFailureStateEt_signal_fail_failure_state SrpMgmtFailureStateEt = "signal-fail-failure-state"

    // Forced Switch
    SrpMgmtFailureStateEt_forced_switch_failure_state SrpMgmtFailureStateEt = "forced-switch-failure-state"

    // Shutdown
    SrpMgmtFailureStateEt_shutdown_failure_state SrpMgmtFailureStateEt = "shutdown-failure-state"

    // Invalid
    SrpMgmtFailureStateEt_invalid_failure_state SrpMgmtFailureStateEt = "invalid-failure-state"

    // Unknown
    SrpMgmtFailureStateEt_unknown_failure_state SrpMgmtFailureStateEt = "unknown-failure-state"
)

// GccDerState represents Gcc der state
type GccDerState string

const (
    // In Service
    GccDerState_in_service GccDerState = "in-service"

    // Out Of Service
    GccDerState_out_of_service GccDerState = "out-of-service"

    // Maintenance
    GccDerState_maintenance GccDerState = "maintenance"

    // Automatic In Service
    GccDerState_ais GccDerState = "ais"
)

// EfpTagEtype represents Tag ethertype
type EfpTagEtype string

const (
    // Untagged
    EfpTagEtype_untagged EfpTagEtype = "untagged"

    // Dot1Q
    EfpTagEtype_dot1q EfpTagEtype = "dot1q"

    // Dot1ad
    EfpTagEtype_dot1ad EfpTagEtype = "dot1ad"
)

// TunnelGreMode represents Tunnel gre mode
type TunnelGreMode string

const (
    // Tunnel GRE mode is Unknown
    TunnelGreMode_unknown TunnelGreMode = "unknown"

    // Tunnel GRE Mode is IPv4
    TunnelGreMode_gr_eo_ipv4 TunnelGreMode = "gr-eo-ipv4"

    // Tunnel GRE Mode is IPv6
    TunnelGreMode_gr_eo_ipv6 TunnelGreMode = "gr-eo-ipv6"

    // Tunnel MGRE Mode is IPv4
    TunnelGreMode_mgr_eo_ipv4 TunnelGreMode = "mgr-eo-ipv4"

    // Tunnel MGRE Mode is IPv6
    TunnelGreMode_mgr_eo_ipv6 TunnelGreMode = "mgr-eo-ipv6"

    // Tunnel Mode is IPv4
    TunnelGreMode_ipv4 TunnelGreMode = "ipv4"

    // Tunnel Mode is IPv6
    TunnelGreMode_ipv6 TunnelGreMode = "ipv6"
)

// GccSecState represents Gcc sec state
type GccSecState string

const (
    // Normal
    GccSecState_normal GccSecState = "normal"

    // Maintenance
    GccSecState_maintenance GccSecState = "maintenance"

    // Automatic In Service
    GccSecState_ais GccSecState = "ais"
)

// SrpMgmtIpsWrapState represents SRP IPS side wrap state
type SrpMgmtIpsWrapState string

const (
    // Idle
    SrpMgmtIpsWrapState_idle_wrap_state SrpMgmtIpsWrapState = "idle-wrap-state"

    // Wrapped
    SrpMgmtIpsWrapState_wrapped_state SrpMgmtIpsWrapState = "wrapped-state"

    // Locked out
    SrpMgmtIpsWrapState_locked_out_wrap_state SrpMgmtIpsWrapState = "locked-out-wrap-state"

    // UNKNOWN
    SrpMgmtIpsWrapState_unknown_wrap_state SrpMgmtIpsWrapState = "unknown-wrap-state"
)

// StatsCounter represents Stats counter
type StatsCounter string

const (
    // stats counter rate
    StatsCounter_stats_counter_rate StatsCounter = "stats-counter-rate"

    // stats counter uint32
    StatsCounter_stats_counter_uint32 StatsCounter = "stats-counter-uint32"

    // stats counter uint64
    StatsCounter_stats_counter_uint64 StatsCounter = "stats-counter-uint64"

    // stats counter generic
    StatsCounter_stats_counter_generic StatsCounter = "stats-counter-generic"

    // stats counter proto
    StatsCounter_stats_counter_proto StatsCounter = "stats-counter-proto"

    // stats counter srp
    StatsCounter_stats_counter_srp StatsCounter = "stats-counter-srp"

    // stats counter ipv4 prec
    StatsCounter_stats_counter_ipv4_prec StatsCounter = "stats-counter-ipv4-prec"

    // stats counter ipv4 dscp
    StatsCounter_stats_counter_ipv4_dscp StatsCounter = "stats-counter-ipv4-dscp"

    // stats counter mpls exp
    StatsCounter_stats_counter_mpls_exp StatsCounter = "stats-counter-mpls-exp"

    // stats counter ipv4 bgp pa
    StatsCounter_stats_counter_ipv4_bgp_pa StatsCounter = "stats-counter-ipv4-bgp-pa"

    // stats counter src bgp pa
    StatsCounter_stats_counter_src_bgp_pa StatsCounter = "stats-counter-src-bgp-pa"

    // stats counter basic
    StatsCounter_stats_counter_basic StatsCounter = "stats-counter-basic"

    // stats counter comp generic
    StatsCounter_stats_counter_comp_generic StatsCounter = "stats-counter-comp-generic"

    // stats counter comp proto
    StatsCounter_stats_counter_comp_proto StatsCounter = "stats-counter-comp-proto"

    // stats counter comp basic
    StatsCounter_stats_counter_comp_basic StatsCounter = "stats-counter-comp-basic"

    // stats counter accounting
    StatsCounter_stats_counter_accounting StatsCounter = "stats-counter-accounting"

    // stats counter comp accounting
    StatsCounter_stats_counter_comp_accounting StatsCounter = "stats-counter-comp-accounting"

    // stats counter flow
    StatsCounter_stats_counter_flow StatsCounter = "stats-counter-flow"

    // stats counter comp flow
    StatsCounter_stats_counter_comp_flow StatsCounter = "stats-counter-comp-flow"
)

// SonetApsEt represents APS states
type SonetApsEt string

const (
    // APS not configured on port
    SonetApsEt_not_configured SonetApsEt = "not-configured"

    // Working port is up 
    SonetApsEt_working_active SonetApsEt = "working-active"

    // Protect port is up  
    SonetApsEt_protect_active SonetApsEt = "protect-active"

    // Working port is down 
    SonetApsEt_working_inactive SonetApsEt = "working-inactive"

    // Protect port is down  
    SonetApsEt_protect_inactive SonetApsEt = "protect-inactive"
)

// ImAttrDuplex represents Im attr duplex
type ImAttrDuplex string

const (
    // im attr duplex unknown
    ImAttrDuplex_im_attr_duplex_unknown ImAttrDuplex = "im-attr-duplex-unknown"

    // im attr duplex half
    ImAttrDuplex_im_attr_duplex_half ImAttrDuplex = "im-attr-duplex-half"

    // im attr duplex full
    ImAttrDuplex_im_attr_duplex_full ImAttrDuplex = "im-attr-duplex-full"
)

// SrpMgmtIpsPathInd represents SRP IPS path indication
type SrpMgmtIpsPathInd string

const (
    // SHORT
    SrpMgmtIpsPathInd_short_path SrpMgmtIpsPathInd = "short-path"

    // LONG
    SrpMgmtIpsPathInd_long_path SrpMgmtIpsPathInd = "long-path"

    // UNKNOWN
    SrpMgmtIpsPathInd_unknown_path SrpMgmtIpsPathInd = "unknown-path"
)

// PppFsmState represents Ppp fsm state
type PppFsmState string

const (
    // Connection Idle
    PppFsmState_ppp_fsm_state_initial_0 PppFsmState = "ppp-fsm-state-initial-0"

    // This layer required, but lower layer down
    PppFsmState_ppp_fsm_state_starting_1 PppFsmState = "ppp-fsm-state-starting-1"

    // Lower layer up, but this layer not required
    PppFsmState_ppp_fsm_state_closed_2 PppFsmState = "ppp-fsm-state-closed-2"

    // Listening for a Config Request
    PppFsmState_ppp_fsm_state_stopped_3 PppFsmState = "ppp-fsm-state-stopped-3"

    // Shutting down due to local change
    PppFsmState_ppp_fsm_state_closing_4 PppFsmState = "ppp-fsm-state-closing-4"

    // Shutting down due to peer's actions
    PppFsmState_ppp_fsm_state_stopping_5 PppFsmState = "ppp-fsm-state-stopping-5"

    // Config Request Sent
    PppFsmState_ppp_fsm_state_req_sent_6 PppFsmState = "ppp-fsm-state-req-sent-6"

    // Config Ack Received
    PppFsmState_ppp_fsm_state_ack_rcvd_7 PppFsmState = "ppp-fsm-state-ack-rcvd-7"

    // Config Ack Sent
    PppFsmState_ppp_fsm_state_ack_sent_8 PppFsmState = "ppp-fsm-state-ack-sent-8"

    // Connection Open
    PppFsmState_ppp_fsm_state_opened_9 PppFsmState = "ppp-fsm-state-opened-9"
)

// EfpTagPriority represents Priority
type EfpTagPriority string

const (
    // Priority 0
    EfpTagPriority_priority0 EfpTagPriority = "priority0"

    // Priority 1
    EfpTagPriority_priority1 EfpTagPriority = "priority1"

    // Priority 2
    EfpTagPriority_priority2 EfpTagPriority = "priority2"

    // Priority 3
    EfpTagPriority_priority3 EfpTagPriority = "priority3"

    // Priority 4
    EfpTagPriority_priority4 EfpTagPriority = "priority4"

    // Priority 5
    EfpTagPriority_priority5 EfpTagPriority = "priority5"

    // Priority 6
    EfpTagPriority_priority6 EfpTagPriority = "priority6"

    // Priority 7
    EfpTagPriority_priority7 EfpTagPriority = "priority7"

    // Any priority
    EfpTagPriority_priority_any EfpTagPriority = "priority-any"
)

// ImCmdLoopbackEnum represents Im cmd loopback enum
type ImCmdLoopbackEnum string

const (
    // no loopback
    ImCmdLoopbackEnum_no_loopback ImCmdLoopbackEnum = "no-loopback"

    // internal loopback
    ImCmdLoopbackEnum_internal_loopback ImCmdLoopbackEnum = "internal-loopback"

    // external loopback
    ImCmdLoopbackEnum_external_loopback ImCmdLoopbackEnum = "external-loopback"

    // line loopback
    ImCmdLoopbackEnum_line_loopback ImCmdLoopbackEnum = "line-loopback"
)

// ImCmdFrTypeEnum represents Im cmd fr type enum
type ImCmdFrTypeEnum string

const (
    // frame relay cisco
    ImCmdFrTypeEnum_frame_relay_cisco ImCmdFrTypeEnum = "frame-relay-cisco"

    // frame relay ietf
    ImCmdFrTypeEnum_frame_relay_ietf ImCmdFrTypeEnum = "frame-relay-ietf"
)

// ImCmdLmiTypeEnum represents Im cmd lmi type enum
type ImCmdLmiTypeEnum string

const (
    // lmi type auto
    ImCmdLmiTypeEnum_lmi_type_auto ImCmdLmiTypeEnum = "lmi-type-auto"

    // lmi type ansi
    ImCmdLmiTypeEnum_lmi_type_ansi ImCmdLmiTypeEnum = "lmi-type-ansi"

    // lmi type ccitt
    ImCmdLmiTypeEnum_lmi_type_ccitt ImCmdLmiTypeEnum = "lmi-type-ccitt"

    // lmi type cisco
    ImCmdLmiTypeEnum_lmi_type_cisco ImCmdLmiTypeEnum = "lmi-type-cisco"
)

// SrpMgmtSrrFailure represents SRP SRR failure type
type SrpMgmtSrrFailure string

const (
    // Idle
    SrpMgmtSrrFailure_idle_srr_failure SrpMgmtSrrFailure = "idle-srr-failure"

    // Wait To Restore
    SrpMgmtSrrFailure_wait_to_restore_srr_failure SrpMgmtSrrFailure = "wait-to-restore-srr-failure"

    // Signal Fail
    SrpMgmtSrrFailure_signal_fail_srr_failure SrpMgmtSrrFailure = "signal-fail-srr-failure"

    // Forced Switch
    SrpMgmtSrrFailure_forced_switch_srr_failure SrpMgmtSrrFailure = "forced-switch-srr-failure"

    // UNKNOWN
    SrpMgmtSrrFailure_unknown_srr_failure SrpMgmtSrrFailure = "unknown-srr-failure"
)

// ImStateEnum represents Im state enum
type ImStateEnum string

const (
    // im state not ready
    ImStateEnum_im_state_not_ready ImStateEnum = "im-state-not-ready"

    // im state admin down
    ImStateEnum_im_state_admin_down ImStateEnum = "im-state-admin-down"

    // im state down
    ImStateEnum_im_state_down ImStateEnum = "im-state-down"

    // im state up
    ImStateEnum_im_state_up ImStateEnum = "im-state-up"

    // im state shutdown
    ImStateEnum_im_state_shutdown ImStateEnum = "im-state-shutdown"

    // im state err disable
    ImStateEnum_im_state_err_disable ImStateEnum = "im-state-err-disable"

    // im state down immediate
    ImStateEnum_im_state_down_immediate ImStateEnum = "im-state-down-immediate"

    // im state down immediate admin
    ImStateEnum_im_state_down_immediate_admin ImStateEnum = "im-state-down-immediate-admin"

    // im state down graceful
    ImStateEnum_im_state_down_graceful ImStateEnum = "im-state-down-graceful"

    // im state begin shutdown
    ImStateEnum_im_state_begin_shutdown ImStateEnum = "im-state-begin-shutdown"

    // im state end shutdown
    ImStateEnum_im_state_end_shutdown ImStateEnum = "im-state-end-shutdown"

    // im state begin error disable
    ImStateEnum_im_state_begin_error_disable ImStateEnum = "im-state-begin-error-disable"

    // im state end error disable
    ImStateEnum_im_state_end_error_disable ImStateEnum = "im-state-end-error-disable"

    // im state begin down graceful
    ImStateEnum_im_state_begin_down_graceful ImStateEnum = "im-state-begin-down-graceful"

    // im state reset
    ImStateEnum_im_state_reset ImStateEnum = "im-state-reset"

    // im state operational
    ImStateEnum_im_state_operational ImStateEnum = "im-state-operational"

    // im state not operational
    ImStateEnum_im_state_not_operational ImStateEnum = "im-state-not-operational"

    // im state unknown
    ImStateEnum_im_state_unknown ImStateEnum = "im-state-unknown"

    // im state last
    ImStateEnum_im_state_last ImStateEnum = "im-state-last"
)

// StatsTypeContents represents Stats type contents
type StatsTypeContents string

const (
    // stats type single
    StatsTypeContents_stats_type_single StatsTypeContents = "stats-type-single"

    // stats type variable
    StatsTypeContents_stats_type_variable StatsTypeContents = "stats-type-variable"
)

// ImAttrFlowControl represents Im attr flow control
type ImAttrFlowControl string

const (
    // im attr flow control off
    ImAttrFlowControl_im_attr_flow_control_off ImAttrFlowControl = "im-attr-flow-control-off"

    // im attr flow control on
    ImAttrFlowControl_im_attr_flow_control_on ImAttrFlowControl = "im-attr-flow-control-on"

    // im attr flow control not sup
    ImAttrFlowControl_im_attr_flow_control_not_sup ImAttrFlowControl = "im-attr-flow-control-not-sup"

    // im attr flow control priority
    ImAttrFlowControl_im_attr_flow_control_priority ImAttrFlowControl = "im-attr-flow-control-priority"
)

// StatsId represents Stats id
type StatsId string

const (
    // stats id type unknown
    StatsId_stats_id_type_unknown StatsId = "stats-id-type-unknown"

    // stats id type min
    StatsId_stats_id_type_min StatsId = "stats-id-type-min"

    // stats id type spare
    StatsId_stats_id_type_spare StatsId = "stats-id-type-spare"

    // stats id type node
    StatsId_stats_id_type_node StatsId = "stats-id-type-node"

    // stats id type other
    StatsId_stats_id_type_other StatsId = "stats-id-type-other"

    // stats id type feature
    StatsId_stats_id_type_feature StatsId = "stats-id-type-feature"

    // stats id type max
    StatsId_stats_id_type_max StatsId = "stats-id-type-max"
)

// TunlPfiAfId represents Tunl pfi af id
type TunlPfiAfId string

const (
    // Unspecified AFI
    TunlPfiAfId_tunl_pfi_af_id_none TunlPfiAfId = "tunl-pfi-af-id-none"

    // IPv4 AFI
    TunlPfiAfId_tunl_pfi_af_id_ipv4 TunlPfiAfId = "tunl-pfi-af-id-ipv4"

    // IPv6 AFI
    TunlPfiAfId_tunl_pfi_af_id_ipv6 TunlPfiAfId = "tunl-pfi-af-id-ipv6"
)

// TunnelKaDfState represents Tunnel ka df state
type TunnelKaDfState string

const (
    // Tunnel GRE KA State is Disabled
    TunnelKaDfState_disable TunnelKaDfState = "disable"

    // Tunnel GRE KA State is Enabled
    TunnelKaDfState_enable TunnelKaDfState = "enable"
)

// BmdMemberTypeEnum represents Bmd member type enum
type BmdMemberTypeEnum string

const (
    // Member has been configured on the local device
    BmdMemberTypeEnum_bmd_mbr_local BmdMemberTypeEnum = "bmd-mbr-local"

    // Member has been configured on an mLACP peer
    // device
    BmdMemberTypeEnum_bmd_mbr_foreign BmdMemberTypeEnum = "bmd-mbr-foreign"

    // Member's type is unknown
    BmdMemberTypeEnum_bmd_mbr_unknown BmdMemberTypeEnum = "bmd-mbr-unknown"
)

// TunnelKeyState represents Tunnel key state
type TunnelKeyState string

const (
    // Tunnel GRE Key is not present
    TunnelKeyState_absent TunnelKeyState = "absent"

    // Tunnel GRE Key is present
    TunnelKeyState_present TunnelKeyState = "present"
)

// BmMbrStateReason represents Bm mbr state reason
type BmMbrStateReason string

const (
    // Reason unavailable (diagnostics error)
    BmMbrStateReason_bm_mbr_state_reason_unknown BmMbrStateReason = "bm-mbr-state-reason-unknown"

    // Link cannot be used (unknown reason)
    BmMbrStateReason_bm_mbr_state_reason_unselectable_unknown BmMbrStateReason = "bm-mbr-state-reason-unselectable-unknown"

    // Link is down
    BmMbrStateReason_bm_mbr_state_reason_link_down BmMbrStateReason = "bm-mbr-state-reason-link-down"

    // Link is being removed from the bundle
    BmMbrStateReason_bm_mbr_state_reason_link_deleting BmMbrStateReason = "bm-mbr-state-reason-link-deleting"

    // Link is in the process of being created
    BmMbrStateReason_bm_mbr_state_reason_creating BmMbrStateReason = "bm-mbr-state-reason-creating"

    // Bundle is in the process of being created
    BmMbrStateReason_bm_mbr_state_reason_bundle_creating BmMbrStateReason = "bm-mbr-state-reason-bundle-creating"

    // Bundle is in the process of being deleted
    BmMbrStateReason_bm_mbr_state_reason_bundle_deleting BmMbrStateReason = "bm-mbr-state-reason-bundle-deleting"

    // Bundle has been shut down
    BmMbrStateReason_bm_mbr_state_reason_bundle_admin_down BmMbrStateReason = "bm-mbr-state-reason-bundle-admin-down"

    // Bundle is in the process of being replicated to
    // this location
    BmMbrStateReason_bm_mbr_state_reason_replicating BmMbrStateReason = "bm-mbr-state-reason-replicating"

    // Incompatible with other links in the bundle
    // (bandwidth out of range)
    BmMbrStateReason_bm_mbr_state_reason_bandwidth BmMbrStateReason = "bm-mbr-state-reason-bandwidth"

    // Loopback: Actor and Partner have the same
    // System ID and Key
    BmMbrStateReason_bm_mbr_state_reason_loop_back BmMbrStateReason = "bm-mbr-state-reason-loop-back"

    // Incompatible with other links in the bundle
    // (LACP vs non-LACP)
    BmMbrStateReason_bm_mbr_state_reason_activity_type BmMbrStateReason = "bm-mbr-state-reason-activity-type"

    // Bundle shutdown is configured for the bundle
    BmMbrStateReason_bm_mbr_state_reason_bundle_shutdown BmMbrStateReason = "bm-mbr-state-reason-bundle-shutdown"

    // Not enough links available to meet
    // minimum-active threshold
    BmMbrStateReason_bm_mbr_state_reason_min_selected BmMbrStateReason = "bm-mbr-state-reason-min-selected"

    // Link is Standby due to maximum-active links
    // configuration
    BmMbrStateReason_bm_mbr_state_reason_max_selected BmMbrStateReason = "bm-mbr-state-reason-max-selected"

    // Bundle has too many member links configured
    BmMbrStateReason_bm_mbr_state_reason_link_limit BmMbrStateReason = "bm-mbr-state-reason-link-limit"

    // Bundle has reached maximum supported number of
    // active links
    BmMbrStateReason_bm_mbr_state_reason_active_limit BmMbrStateReason = "bm-mbr-state-reason-active-limit"

    // Link is Standby (unknown reason)
    BmMbrStateReason_bm_mbr_state_reason_standby_unknown BmMbrStateReason = "bm-mbr-state-reason-standby-unknown"

    // Link is Expired; LACPDUs are not being received
    // from the partner
    BmMbrStateReason_bm_mbr_state_reason_expired BmMbrStateReason = "bm-mbr-state-reason-expired"

    // Link is Defaulted; LACPDUs are not being
    // received from the partner
    BmMbrStateReason_bm_mbr_state_reason_defaulted BmMbrStateReason = "bm-mbr-state-reason-defaulted"

    // Link is Not Aggregatable (unknown reason)
    BmMbrStateReason_bm_mbr_state_reason_act_or_not_agg BmMbrStateReason = "bm-mbr-state-reason-act-or-not-agg"

    // Partner has marked the link as Not Aggregatable
    BmMbrStateReason_bm_mbr_state_reason_partner_not_agg BmMbrStateReason = "bm-mbr-state-reason-partner-not-agg"

    // Partner System ID/Key do not match that of the
    // Selected links
    BmMbrStateReason_bm_mbr_state_reason_lagid BmMbrStateReason = "bm-mbr-state-reason-lagid"

    // Bundle interface is not present in
    // configuration
    BmMbrStateReason_bm_mbr_state_reason_bundle_not_cfgd BmMbrStateReason = "bm-mbr-state-reason-bundle-not-cfgd"

    // Wait-while timer is running
    BmMbrStateReason_bm_mbr_state_reason_bundle_not_ready BmMbrStateReason = "bm-mbr-state-reason-bundle-not-ready"

    // Partner has not echoed the correct parameters
    // for this link
    BmMbrStateReason_bm_mbr_state_reason_partner_ood BmMbrStateReason = "bm-mbr-state-reason-partner-ood"

    // Partner is not Synchronized (Waiting, Standby,
    // or LAG ID mismatch)
    BmMbrStateReason_bm_mbr_state_reason_partner_not_in_sync BmMbrStateReason = "bm-mbr-state-reason-partner-not-in-sync"

    // Partner is not Synchronized (Waiting, not
    // Selected, or out-of-date)
    BmMbrStateReason_bm_mbr_state_reason_foreign_partner_oos BmMbrStateReason = "bm-mbr-state-reason-foreign-partner-oos"

    // Link is Attached and has not gone Collecting
    // (unknown reason)
    BmMbrStateReason_bm_mbr_state_reason_attach_unknown BmMbrStateReason = "bm-mbr-state-reason-attach-unknown"

    // Partner has not advertized that it is
    // Collecting
    BmMbrStateReason_bm_mbr_state_reason_partner_not_collecting BmMbrStateReason = "bm-mbr-state-reason-partner-not-collecting"

    // Link is Collecting and has not gone
    // Distributing (unknown reason)
    BmMbrStateReason_bm_mbr_state_reason_collect_unknown BmMbrStateReason = "bm-mbr-state-reason-collect-unknown"

    // Link is marked as Standby by mLACP peer
    BmMbrStateReason_bm_mbr_state_reason_standby_foreign BmMbrStateReason = "bm-mbr-state-reason-standby-foreign"

    // Link is waiting for BFD session to start
    BmMbrStateReason_bm_mbr_state_reason_bfd_starting BmMbrStateReason = "bm-mbr-state-reason-bfd-starting"

    // BFD state of this link is Down
    BmMbrStateReason_bm_mbr_state_reason_bfd_down BmMbrStateReason = "bm-mbr-state-reason-bfd-down"

    // BFD session is unconfigured on the remote end
    BmMbrStateReason_bm_mbr_state_reason_bfd_nbr_unconfig BmMbrStateReason = "bm-mbr-state-reason-bfd-nbr-unconfig"

    // Link is not operational as a result of mLACP
    // negotiations
    BmMbrStateReason_bm_mbr_state_reason_mlacp BmMbrStateReason = "bm-mbr-state-reason-mlacp"

    // ICCP group is isolated from the core network
    BmMbrStateReason_bm_mbr_state_reason_pe_isolated BmMbrStateReason = "bm-mbr-state-reason-pe-isolated"

    // Forced switchover to the mLACP peer
    BmMbrStateReason_bm_mbr_state_reason_forced_switchover BmMbrStateReason = "bm-mbr-state-reason-forced-switchover"

    // Link is error disabled (unknown reason)
    BmMbrStateReason_bm_mbr_state_reason_errdis_unknown BmMbrStateReason = "bm-mbr-state-reason-errdis-unknown"

    // Waiting for member state information from mLACP
    // peer
    BmMbrStateReason_bm_mbr_state_reason_mlacp_no_mbr_state_info BmMbrStateReason = "bm-mbr-state-reason-mlacp-no-mbr-state-info"

    // Link is Active
    BmMbrStateReason_bm_mbr_state_reason_active BmMbrStateReason = "bm-mbr-state-reason-active"

    // Waiting for bundle state information from mLACP
    // peer
    BmMbrStateReason_bm_mbr_state_reason_mlacp_no_bdl_state_info BmMbrStateReason = "bm-mbr-state-reason-mlacp-no-bdl-state-info"

    // Waiting for bundle configuration information
    // from mLACP peer
    BmMbrStateReason_bm_mbr_state_reason_mlacp_no_bdl_config_info BmMbrStateReason = "bm-mbr-state-reason-mlacp-no-bdl-config-info"

    // Waiting for bundle to complete initial
    // synchronization with mLACP peer
    BmMbrStateReason_bm_mbr_state_reason_mlacp_no_bdl_sync BmMbrStateReason = "bm-mbr-state-reason-mlacp-no-bdl-sync"

    // mLACP bundle does not have a peer device
    BmMbrStateReason_bm_mbr_state_reason_mlacp_bdl_has_no_peer BmMbrStateReason = "bm-mbr-state-reason-mlacp-bdl-has-no-peer"

    // Link is being ignored due to an inconsistency
    // with mLACP peer
    BmMbrStateReason_bm_mbr_state_reason_mlacp_nak BmMbrStateReason = "bm-mbr-state-reason-mlacp-nak"

    // ICCP transport is unavailable
    BmMbrStateReason_bm_mbr_state_reason_mlacp_transport_unavailable BmMbrStateReason = "bm-mbr-state-reason-mlacp-transport-unavailable"

    // ICCP Group is not fully configured
    BmMbrStateReason_bm_mbr_state_reason_mlacp_not_configured BmMbrStateReason = "bm-mbr-state-reason-mlacp-not-configured"

    // mLACP recovery delay timer is running
    BmMbrStateReason_bm_mbr_state_reason_recovery_timer BmMbrStateReason = "bm-mbr-state-reason-recovery-timer"

    // mLACP peer is active
    BmMbrStateReason_bm_mbr_state_reason_mlacp_standby BmMbrStateReason = "bm-mbr-state-reason-mlacp-standby"

    // mLACP peer has more links/bandwidth available
    BmMbrStateReason_bm_mbr_state_reason_maximized_out BmMbrStateReason = "bm-mbr-state-reason-maximized-out"

    // mLACP peer has one or more links Selected
    BmMbrStateReason_bm_mbr_state_reason_mlacp_peer_selected BmMbrStateReason = "bm-mbr-state-reason-mlacp-peer-selected"

    // mLACP bundle does not have a peer device
    // (connect timer running)
    BmMbrStateReason_bm_mbr_state_reason_mlacp_connect_timer_running BmMbrStateReason = "bm-mbr-state-reason-mlacp-connect-timer-running"

    // Bundle is not configured to run mLACP
    BmMbrStateReason_bm_mbr_state_reason_bundle_not_mlacp BmMbrStateReason = "bm-mbr-state-reason-bundle-not-mlacp"

    // Bundle has too many working links configured
    // (more than the maximum-active limit)
    BmMbrStateReason_bm_mbr_state_reason_no_lon BmMbrStateReason = "bm-mbr-state-reason-no-lon"

    // Additional bandwidth from link would exceed
    // load balancing capabilities
    BmMbrStateReason_bm_mbr_state_reason_cumul_rel_bw_limit BmMbrStateReason = "bm-mbr-state-reason-cumul-rel-bw-limit"

    // No MAC address available for the bundle
    BmMbrStateReason_bm_mbr_state_reason_no_mac BmMbrStateReason = "bm-mbr-state-reason-no-mac"

    // No system ID available for use by this bundle
    BmMbrStateReason_bm_mbr_state_reason_no_system_id BmMbrStateReason = "bm-mbr-state-reason-no-system-id"

    // Link is shutdown
    BmMbrStateReason_bm_mbr_state_reason_link_shutdown BmMbrStateReason = "bm-mbr-state-reason-link-shutdown"

    // Non-LACP link in mLACP bundle
    BmMbrStateReason_bm_mbr_state_reason_activity_mlacp BmMbrStateReason = "bm-mbr-state-reason-activity-mlacp"

    // LACP link in inter-chassis bundle
    BmMbrStateReason_bm_mbr_state_reason_activity_iccp BmMbrStateReason = "bm-mbr-state-reason-activity-iccp"

    // Parent bundle is both inter-chassis and
    // configured for mLACP
    BmMbrStateReason_bm_mbr_state_reason_bundle_icpe_mlacp BmMbrStateReason = "bm-mbr-state-reason-bundle-icpe-mlacp"

    // Too many bundle members in system; no link
    // number available
    BmMbrStateReason_bm_mbr_state_reason_no_link_num BmMbrStateReason = "bm-mbr-state-reason-no-link-num"

    // mLACP peer has a higher priority link
    BmMbrStateReason_bm_mbr_state_reason_standby_peer_higher_prio BmMbrStateReason = "bm-mbr-state-reason-standby-peer-higher-prio"

    // Link is in standby redundancy state
    BmMbrStateReason_bm_mbr_state_reason_red_state_standby BmMbrStateReason = "bm-mbr-state-reason-red-state-standby"

    // One or more links in the bundle are in standby
    // redundancy state
    BmMbrStateReason_bm_mbr_state_reason_other_red_state_standby BmMbrStateReason = "bm-mbr-state-reason-other-red-state-standby"

    // Holding down temporary to avoid churn after
    // restart
    BmMbrStateReason_bm_mbr_state_reason_hold_ing BmMbrStateReason = "bm-mbr-state-reason-hold-ing"

    // Bundle has been error-disabled
    BmMbrStateReason_bm_mbr_state_reason_bundle_error_disabled BmMbrStateReason = "bm-mbr-state-reason-bundle-error-disabled"

    // Bundle has been disabled by EFD
    BmMbrStateReason_bm_mbr_state_reason_bundle_efd_disabled BmMbrStateReason = "bm-mbr-state-reason-bundle-efd-disabled"

    // Singleton ICCP group is isolated from the core
    // network
    BmMbrStateReason_bm_mbr_state_reason_singleton_pe_isolated BmMbrStateReason = "bm-mbr-state-reason-singleton-pe-isolated"

    // Link is waiting for BFDv6 session to start
    BmMbrStateReason_bm_mbr_state_reason_bfd_ipv6_starting BmMbrStateReason = "bm-mbr-state-reason-bfd-ipv6-starting"

    // BFDv6 state of this link is Down
    BmMbrStateReason_bm_mbr_state_reason_bfd_ipv6_down BmMbrStateReason = "bm-mbr-state-reason-bfd-ipv6-down"

    // BFDv6 session is unconfigured on the remote end
    BmMbrStateReason_bm_mbr_state_reason_bfd_ipv6_nbr_unconfig BmMbrStateReason = "bm-mbr-state-reason-bfd-ipv6-nbr-unconfig"

    // LACP delay timer is running
    BmMbrStateReason_bm_mbr_state_reason_timer_running BmMbrStateReason = "bm-mbr-state-reason-timer-running"

    // Client has configured the bundle state Down
    BmMbrStateReason_bm_mbr_state_reason_client_bundle_ctrl BmMbrStateReason = "bm-mbr-state-reason-client-bundle-ctrl"

    // Enumeration maximum value
    BmMbrStateReason_bm_mbr_state_reason_count BmMbrStateReason = "bm-mbr-state-reason-count"
)

// BmSeverity represents Severity of the member state reason
type BmSeverity string

const (
    // OK
    BmSeverity_ok BmSeverity = "ok"

    // Information
    BmSeverity_information BmSeverity = "information"

    // Misconfiguration
    BmSeverity_misconfiguration BmSeverity = "misconfiguration"

    // Warning
    BmSeverity_warning BmSeverity = "warning"

    // Error
    BmSeverity_error_ BmSeverity = "error"
)

// SrpMgmtIpsReq represents SRP IPS request type
type SrpMgmtIpsReq string

const (
    // Idle
    SrpMgmtIpsReq_idle_ips_request SrpMgmtIpsReq = "idle-ips-request"

    // Wait To Restore
    SrpMgmtIpsReq_wait_to_restore_ips_request SrpMgmtIpsReq = "wait-to-restore-ips-request"

    // Manual Switch
    SrpMgmtIpsReq_manual_switch_ips_request SrpMgmtIpsReq = "manual-switch-ips-request"

    // Signal Degrade
    SrpMgmtIpsReq_signal_degrade_ips_request SrpMgmtIpsReq = "signal-degrade-ips-request"

    // Signal Fail
    SrpMgmtIpsReq_signal_fail_ips_request SrpMgmtIpsReq = "signal-fail-ips-request"

    // Forced Switch
    SrpMgmtIpsReq_forced_switch_ips_request SrpMgmtIpsReq = "forced-switch-ips-request"

    // UNKNOWN
    SrpMgmtIpsReq_unknown_ips_request SrpMgmtIpsReq = "unknown-ips-request"
)

// SrpMgmtFailureEt represents SRP failure type
type SrpMgmtFailureEt string

const (
    // Hardware missing
    SrpMgmtFailureEt_hardware_missing_failure SrpMgmtFailureEt = "hardware-missing-failure"

    // L1 admin state
    SrpMgmtFailureEt_layer1_admin_state_failure SrpMgmtFailureEt = "layer1-admin-state-failure"

    // Layer 1 error
    SrpMgmtFailureEt_layer1_error_failure SrpMgmtFailureEt = "layer1-error-failure"

    // Keepalive missed
    SrpMgmtFailureEt_keepalive_missed_failure SrpMgmtFailureEt = "keepalive-missed-failure"

    // Link quality degraded
    SrpMgmtFailureEt_link_quality_degraded_failure SrpMgmtFailureEt = "link-quality-degraded-failure"

    // Mate problem
    SrpMgmtFailureEt_mate_problem_failure SrpMgmtFailureEt = "mate-problem-failure"

    // Side mismatch
    SrpMgmtFailureEt_side_mismatch_failure SrpMgmtFailureEt = "side-mismatch-failure"

    // Unknown
    SrpMgmtFailureEt_unknown_failure SrpMgmtFailureEt = "unknown-failure"
)

// ImAttrTransportMode represents Im attr transport mode
type ImAttrTransportMode string

const (
    // im attr transport mode unknown
    ImAttrTransportMode_im_attr_transport_mode_unknown ImAttrTransportMode = "im-attr-transport-mode-unknown"

    // im attr transport mode lan
    ImAttrTransportMode_im_attr_transport_mode_lan ImAttrTransportMode = "im-attr-transport-mode-lan"

    // im attr transport mode wan
    ImAttrTransportMode_im_attr_transport_mode_wan ImAttrTransportMode = "im-attr-transport-mode-wan"

    // im attr transport mode otn bt opu1e
    ImAttrTransportMode_im_attr_transport_mode_otn_bt_opu1e ImAttrTransportMode = "im-attr-transport-mode-otn-bt-opu1e"

    // im attr transport mode otn bt opu2e
    ImAttrTransportMode_im_attr_transport_mode_otn_bt_opu2e ImAttrTransportMode = "im-attr-transport-mode-otn-bt-opu2e"

    // im attr transport mode otn opu3
    ImAttrTransportMode_im_attr_transport_mode_otn_opu3 ImAttrTransportMode = "im-attr-transport-mode-otn-opu3"

    // im attr transport mode otn opu4
    ImAttrTransportMode_im_attr_transport_mode_otn_opu4 ImAttrTransportMode = "im-attr-transport-mode-otn-opu4"
)

// TunlIpModeDir represents Tunl ip mode dir
type TunlIpModeDir string

const (
    // tunl ip mode dir none
    TunlIpModeDir_tunl_ip_mode_dir_none TunlIpModeDir = "tunl-ip-mode-dir-none"

    // tunl ip mode dir decap
    TunlIpModeDir_tunl_ip_mode_dir_decap TunlIpModeDir = "tunl-ip-mode-dir-decap"

    // tunl ip mode dir encap
    TunlIpModeDir_tunl_ip_mode_dir_encap TunlIpModeDir = "tunl-ip-mode-dir-encap"

    // tunl ip mode dir max
    TunlIpModeDir_tunl_ip_mode_dir_max TunlIpModeDir = "tunl-ip-mode-dir-max"
)

// ImCmdEncapsEnum represents Im cmd encaps enum
type ImCmdEncapsEnum string

const (
    // frame relay
    ImCmdEncapsEnum_frame_relay ImCmdEncapsEnum = "frame-relay"

    // vlan
    ImCmdEncapsEnum_vlan ImCmdEncapsEnum = "vlan"

    // ppp
    ImCmdEncapsEnum_ppp ImCmdEncapsEnum = "ppp"
)

// BmMuxstate represents Bm muxstate
type BmMuxstate string

const (
    // Port is not attached to a bundle
    BmMuxstate_detached BmMuxstate = "detached"

    // Port has chosen bundle and is waiting to join
    BmMuxstate_waiting BmMuxstate = "waiting"

    // Port is attached to the bundle but not active
    BmMuxstate_attached BmMuxstate = "attached"

    // Port is ready to receive data
    BmMuxstate_collecting BmMuxstate = "collecting"

    // Port is distributing data
    BmMuxstate_distributing BmMuxstate = "distributing"

    // Port is active and can send and receive data
    BmMuxstate_collecting_distributing BmMuxstate = "collecting-distributing"
)

// NcpIdent represents Ncp ident
type NcpIdent string

const (
    // CDP control protocol
    NcpIdent_cdpcp NcpIdent = "cdpcp"

    // IPv4 control protocol
    NcpIdent_ipcp NcpIdent = "ipcp"

    // IPv4 Interworking control protocol
    NcpIdent_ipcpiw NcpIdent = "ipcpiw"

    // IPv6 control protocol
    NcpIdent_ipv6cp NcpIdent = "ipv6cp"

    // MPLS control protocol
    NcpIdent_mplscp NcpIdent = "mplscp"

    // OSI (CLNS) control protocol
    NcpIdent_osicp NcpIdent = "osicp"
)

// BmdMemberState represents Bmd member state
type BmdMemberState string

const (
    // Member is configured
    BmdMemberState_bmd_mbr_state_configured BmdMemberState = "bmd-mbr-state-configured"

    // Member is standby
    BmdMemberState_bmd_mbr_state_standby BmdMemberState = "bmd-mbr-state-standby"

    // Member is hot standby
    BmdMemberState_bmd_mbr_state_hot_standby BmdMemberState = "bmd-mbr-state-hot-standby"

    // Member is negotiating
    BmdMemberState_bmd_mbr_state_negotiating BmdMemberState = "bmd-mbr-state-negotiating"

    // Member has a BFD session running
    BmdMemberState_bmd_mbr_state_bfd_running BmdMemberState = "bmd-mbr-state-bfd-running"

    // Member is active
    BmdMemberState_bmd_mbr_state_active BmdMemberState = "bmd-mbr-state-active"
)

// BmMuxreason represents Bm muxreason
type BmMuxreason string

const (
    // Selection logic has not yet been run for the
    // bundle this link is a member of
    BmMuxreason_bm_mux_reason_no_reason BmMuxreason = "bm-mux-reason-no-reason"

    // Link is down
    BmMuxreason_bm_mux_reason_link_down BmMuxreason = "bm-mux-reason-link-down"

    // Link is being removed from the bundle
    BmMuxreason_bm_mux_reason_link_deleted BmMuxreason = "bm-mux-reason-link-deleted"

    // Link has wrong duplexity
    BmMuxreason_bm_mux_reason_duplex BmMuxreason = "bm-mux-reason-duplex"

    // Link has wrong bandwidth
    BmMuxreason_bm_mux_reason_bandwidth BmMuxreason = "bm-mux-reason-bandwidth"

    // Link is a loopback interface
    BmMuxreason_bm_mux_reason_loop_back BmMuxreason = "bm-mux-reason-loop-back"

    // Link has wrong activity type
    BmMuxreason_bm_mux_reason_activity_type BmMuxreason = "bm-mux-reason-activity-type"

    // Link's bundle already has maximum number of
    // members allowed
    BmMuxreason_bm_mux_reason_link_limit BmMuxreason = "bm-mux-reason-link-limit"

    // Link is attached to a shared medium
    BmMuxreason_bm_mux_reason_shared BmMuxreason = "bm-mux-reason-shared"

    // Link has wrong LAG ID
    BmMuxreason_bm_mux_reason_lagid BmMuxreason = "bm-mux-reason-lagid"

    // Link's bundle does not exist
    BmMuxreason_bm_mux_reason_no_bundle BmMuxreason = "bm-mux-reason-no-bundle"

    // Link's bundle has no primary link
    BmMuxreason_bm_mux_reason_no_primary BmMuxreason = "bm-mux-reason-no-primary"

    // Link's bundle is shut down
    BmMuxreason_bm_mux_reason_bundle_down BmMuxreason = "bm-mux-reason-bundle-down"

    // Link is marked individual by partner
    BmMuxreason_bm_mux_reason_individual BmMuxreason = "bm-mux-reason-individual"

    // Link is Defaulted, suggesting it is not
    // receiving LACPDUs from the peer
    BmMuxreason_bm_mux_reason_defaulted BmMuxreason = "bm-mux-reason-defaulted"

    // Link is in InSync state
    BmMuxreason_bm_mux_reason_in_sync BmMuxreason = "bm-mux-reason-in-sync"

    // Link is in Collecting state
    BmMuxreason_bm_mux_reason_collecting BmMuxreason = "bm-mux-reason-collecting"

    // Link exceeds maximum active limit
    BmMuxreason_bm_mux_reason_active_link_limit BmMuxreason = "bm-mux-reason-active-link-limit"

    // Link is in Distributing state
    BmMuxreason_bm_mux_reason_distributing BmMuxreason = "bm-mux-reason-distributing"

    // Enumeration maximum value
    BmMuxreason_bm_mux_reason_count BmMuxreason = "bm-mux-reason-count"
)

// ImAttrLink represents Im attr link
type ImAttrLink string

const (
    // im attr link type auto
    ImAttrLink_im_attr_link_type_auto ImAttrLink = "im-attr-link-type-auto"

    // im attr link type force
    ImAttrLink_im_attr_link_type_force ImAttrLink = "im-attr-link-type-force"
)

// VlanEncaps represents VLAN encapsulation
type VlanEncaps string

const (
    // No encapsulation
    VlanEncaps_no_encapsulation VlanEncaps = "no-encapsulation"

    // IEEE 802.1Q encapsulation
    VlanEncaps_dot1q VlanEncaps = "dot1q"

    // Double 802.1Q encapsulation
    VlanEncaps_qinq VlanEncaps = "qinq"

    // Double 802.1Q wildcarded encapsulation
    VlanEncaps_qin_any VlanEncaps = "qin-any"

    // IEEE 802.1Q native VLAN encapsulation
    VlanEncaps_dot1q_native VlanEncaps = "dot1q-native"

    // IEEE 802.1ad encapsulation
    VlanEncaps_dot1ad VlanEncaps = "dot1ad"

    // IEEE 802.1ad native VLAN encapsulation
    VlanEncaps_dot1ad_native VlanEncaps = "dot1ad-native"

    // Ethernet Service Instance
    VlanEncaps_service_instance VlanEncaps = "service-instance"

    // IEEE 802.1ad 802.1Q encapsulation
    VlanEncaps_dot1ad_dot1q VlanEncaps = "dot1ad-dot1q"

    // IEEE 802.1ad wildcard 802.1Q encapsulation
    VlanEncaps_dot1ad_any VlanEncaps = "dot1ad-any"
)

// EfpPayloadEtype represents Payload ethertype match
type EfpPayloadEtype string

const (
    // Any
    EfpPayloadEtype_payload_ethertype_any EfpPayloadEtype = "payload-ethertype-any"

    // IP
    EfpPayloadEtype_payload_ethertype_ip EfpPayloadEtype = "payload-ethertype-ip"

    // PPPoE
    EfpPayloadEtype_payload_ethertype_pppoe EfpPayloadEtype = "payload-ethertype-pppoe"
)

// InterfaceTypeSet represents Interface type set
type InterfaceTypeSet string

const (
    // Restrict the output to hardware interfaces only
    InterfaceTypeSet_hardware_interfaces InterfaceTypeSet = "hardware-interfaces"
)

// BmStateReasonTarget represents Scope of the state reason
type BmStateReasonTarget string

const (
    // Member applicable reason
    BmStateReasonTarget_member_reason BmStateReasonTarget = "member-reason"

    // Bundle applicable reason
    BmStateReasonTarget_bundle_reason BmStateReasonTarget = "bundle-reason"
)

// ImAttrMedia represents Im attr media
type ImAttrMedia string

const (
    // im attr media other
    ImAttrMedia_im_attr_media_other ImAttrMedia = "im-attr-media-other"

    // im attr media unknown
    ImAttrMedia_im_attr_media_unknown ImAttrMedia = "im-attr-media-unknown"

    // im attr media aui
    ImAttrMedia_im_attr_media_aui ImAttrMedia = "im-attr-media-aui"

    // im attr media 10base5
    ImAttrMedia_im_attr_media_10base5 ImAttrMedia = "im-attr-media-10base5"

    // im attr media foirl
    ImAttrMedia_im_attr_media_foirl ImAttrMedia = "im-attr-media-foirl"

    // im attr media 10base2
    ImAttrMedia_im_attr_media_10base2 ImAttrMedia = "im-attr-media-10base2"

    // im attr media 10broad36
    ImAttrMedia_im_attr_media_10broad36 ImAttrMedia = "im-attr-media-10broad36"

    // im attr media 10base
    ImAttrMedia_im_attr_media_10base ImAttrMedia = "im-attr-media-10base"

    // im attr media 10base thd
    ImAttrMedia_im_attr_media_10base_thd ImAttrMedia = "im-attr-media-10base-thd"

    // im attr media 10base tfd
    ImAttrMedia_im_attr_media_10base_tfd ImAttrMedia = "im-attr-media-10base-tfd"

    // im attr media 10base fp
    ImAttrMedia_im_attr_media_10base_fp ImAttrMedia = "im-attr-media-10base-fp"

    // im attr media 10base fb
    ImAttrMedia_im_attr_media_10base_fb ImAttrMedia = "im-attr-media-10base-fb"

    // im attr media 10base fl
    ImAttrMedia_im_attr_media_10base_fl ImAttrMedia = "im-attr-media-10base-fl"

    // im attr media 10base flhd
    ImAttrMedia_im_attr_media_10base_flhd ImAttrMedia = "im-attr-media-10base-flhd"

    // im attr media 10base flfd
    ImAttrMedia_im_attr_media_10base_flfd ImAttrMedia = "im-attr-media-10base-flfd"

    // im attr media 100base t4
    ImAttrMedia_im_attr_media_100base_t4 ImAttrMedia = "im-attr-media-100base-t4"

    // im attr media 100base tx
    ImAttrMedia_im_attr_media_100base_tx ImAttrMedia = "im-attr-media-100base-tx"

    // im attr media 100base txhd
    ImAttrMedia_im_attr_media_100base_txhd ImAttrMedia = "im-attr-media-100base-txhd"

    // im attr media 100base txfd
    ImAttrMedia_im_attr_media_100base_txfd ImAttrMedia = "im-attr-media-100base-txfd"

    // im attr media 100base fx
    ImAttrMedia_im_attr_media_100base_fx ImAttrMedia = "im-attr-media-100base-fx"

    // im attr media 100base fxhd
    ImAttrMedia_im_attr_media_100base_fxhd ImAttrMedia = "im-attr-media-100base-fxhd"

    // im attr media 100base fxfd
    ImAttrMedia_im_attr_media_100base_fxfd ImAttrMedia = "im-attr-media-100base-fxfd"

    // im attr media 100base ex
    ImAttrMedia_im_attr_media_100base_ex ImAttrMedia = "im-attr-media-100base-ex"

    // im attr media 100base exhd
    ImAttrMedia_im_attr_media_100base_exhd ImAttrMedia = "im-attr-media-100base-exhd"

    // im attr media 100base exfd
    ImAttrMedia_im_attr_media_100base_exfd ImAttrMedia = "im-attr-media-100base-exfd"

    // im attr media 100base t2
    ImAttrMedia_im_attr_media_100base_t2 ImAttrMedia = "im-attr-media-100base-t2"

    // im attr media 100base t2hd
    ImAttrMedia_im_attr_media_100base_t2hd ImAttrMedia = "im-attr-media-100base-t2hd"

    // im attr media 100base t2fd
    ImAttrMedia_im_attr_media_100base_t2fd ImAttrMedia = "im-attr-media-100base-t2fd"

    // im attr media 1000base x
    ImAttrMedia_im_attr_media_1000base_x ImAttrMedia = "im-attr-media-1000base-x"

    // im attr media 1000base xhdx
    ImAttrMedia_im_attr_media_1000base_xhdx ImAttrMedia = "im-attr-media-1000base-xhdx"

    // im attr media 1000base xfd
    ImAttrMedia_im_attr_media_1000base_xfd ImAttrMedia = "im-attr-media-1000base-xfd"

    // im attr media 1000base lx
    ImAttrMedia_im_attr_media_1000base_lx ImAttrMedia = "im-attr-media-1000base-lx"

    // im attr media 1000base lxhd
    ImAttrMedia_im_attr_media_1000base_lxhd ImAttrMedia = "im-attr-media-1000base-lxhd"

    // im attr media 1000base lxfdx
    ImAttrMedia_im_attr_media_1000base_lxfdx ImAttrMedia = "im-attr-media-1000base-lxfdx"

    // im attr media 1000base sx
    ImAttrMedia_im_attr_media_1000base_sx ImAttrMedia = "im-attr-media-1000base-sx"

    // im attr media 1000base sxhd
    ImAttrMedia_im_attr_media_1000base_sxhd ImAttrMedia = "im-attr-media-1000base-sxhd"

    // im attr media 1000base sxfd
    ImAttrMedia_im_attr_media_1000base_sxfd ImAttrMedia = "im-attr-media-1000base-sxfd"

    // im attr media 1000base cx
    ImAttrMedia_im_attr_media_1000base_cx ImAttrMedia = "im-attr-media-1000base-cx"

    // im attr media 1000base cxhdx
    ImAttrMedia_im_attr_media_1000base_cxhdx ImAttrMedia = "im-attr-media-1000base-cxhdx"

    // im attr media 1000base cxfd
    ImAttrMedia_im_attr_media_1000base_cxfd ImAttrMedia = "im-attr-media-1000base-cxfd"

    // im attr media 1000base
    ImAttrMedia_im_attr_media_1000base ImAttrMedia = "im-attr-media-1000base"

    // im attr media 1000base thd
    ImAttrMedia_im_attr_media_1000base_thd ImAttrMedia = "im-attr-media-1000base-thd"

    // im attr media 1000base tfd
    ImAttrMedia_im_attr_media_1000base_tfd ImAttrMedia = "im-attr-media-1000base-tfd"

    // im attr media 10gbase x
    ImAttrMedia_im_attr_media_10gbase_x ImAttrMedia = "im-attr-media-10gbase-x"

    // im attr media 10gbase lx4
    ImAttrMedia_im_attr_media_10gbase_lx4 ImAttrMedia = "im-attr-media-10gbase-lx4"

    // im attr media 10gbase r
    ImAttrMedia_im_attr_media_10gbase_r ImAttrMedia = "im-attr-media-10gbase-r"

    // im attr media 10gbase er
    ImAttrMedia_im_attr_media_10gbase_er ImAttrMedia = "im-attr-media-10gbase-er"

    // im attr media 10gbase lr
    ImAttrMedia_im_attr_media_10gbase_lr ImAttrMedia = "im-attr-media-10gbase-lr"

    // im attr media 10gbase sr
    ImAttrMedia_im_attr_media_10gbase_sr ImAttrMedia = "im-attr-media-10gbase-sr"

    // im attr media 10gbase w
    ImAttrMedia_im_attr_media_10gbase_w ImAttrMedia = "im-attr-media-10gbase-w"

    // im attr media 10gbase ew
    ImAttrMedia_im_attr_media_10gbase_ew ImAttrMedia = "im-attr-media-10gbase-ew"

    // im attr media 10gbase lw
    ImAttrMedia_im_attr_media_10gbase_lw ImAttrMedia = "im-attr-media-10gbase-lw"

    // im attr media 10gbase sw
    ImAttrMedia_im_attr_media_10gbase_sw ImAttrMedia = "im-attr-media-10gbase-sw"

    // im attr media 10gbase zr
    ImAttrMedia_im_attr_media_10gbase_zr ImAttrMedia = "im-attr-media-10gbase-zr"

    // im attr media 802 9a
    ImAttrMedia_im_attr_media_802_9a ImAttrMedia = "im-attr-media-802-9a"

    // im attr media rj45
    ImAttrMedia_im_attr_media_rj45 ImAttrMedia = "im-attr-media-rj45"

    // im attr media 1000base zx
    ImAttrMedia_im_attr_media_1000base_zx ImAttrMedia = "im-attr-media-1000base-zx"

    // im attr media 1000base cwdm
    ImAttrMedia_im_attr_media_1000base_cwdm ImAttrMedia = "im-attr-media-1000base-cwdm"

    // im attr media 1000base cwdm 1470
    ImAttrMedia_im_attr_media_1000base_cwdm_1470 ImAttrMedia = "im-attr-media-1000base-cwdm-1470"

    // im attr media 1000base cwdm 1490
    ImAttrMedia_im_attr_media_1000base_cwdm_1490 ImAttrMedia = "im-attr-media-1000base-cwdm-1490"

    // im attr media 1000base cwdm 1510
    ImAttrMedia_im_attr_media_1000base_cwdm_1510 ImAttrMedia = "im-attr-media-1000base-cwdm-1510"

    // im attr media 1000base cwdm 1530
    ImAttrMedia_im_attr_media_1000base_cwdm_1530 ImAttrMedia = "im-attr-media-1000base-cwdm-1530"

    // im attr media 1000base cwdm 1550
    ImAttrMedia_im_attr_media_1000base_cwdm_1550 ImAttrMedia = "im-attr-media-1000base-cwdm-1550"

    // im attr media 1000base cwdm 1570
    ImAttrMedia_im_attr_media_1000base_cwdm_1570 ImAttrMedia = "im-attr-media-1000base-cwdm-1570"

    // im attr media 1000base cwdm 1590
    ImAttrMedia_im_attr_media_1000base_cwdm_1590 ImAttrMedia = "im-attr-media-1000base-cwdm-1590"

    // im attr media 1000base cwdm 1610
    ImAttrMedia_im_attr_media_1000base_cwdm_1610 ImAttrMedia = "im-attr-media-1000base-cwdm-1610"

    // im attr media 10gbase dwdm
    ImAttrMedia_im_attr_media_10gbase_dwdm ImAttrMedia = "im-attr-media-10gbase-dwdm"

    // im attr media 100gbase lr4
    ImAttrMedia_im_attr_media_100gbase_lr4 ImAttrMedia = "im-attr-media-100gbase-lr4"

    // im attr media 1000base dwdm
    ImAttrMedia_im_attr_media_1000base_dwdm ImAttrMedia = "im-attr-media-1000base-dwdm"

    // im attr media 1000base dwdm 1533
    ImAttrMedia_im_attr_media_1000base_dwdm_1533 ImAttrMedia = "im-attr-media-1000base-dwdm-1533"

    // im attr media 1000base dwdm 1537
    ImAttrMedia_im_attr_media_1000base_dwdm_1537 ImAttrMedia = "im-attr-media-1000base-dwdm-1537"

    // im attr media 1000base dwdm 1541
    ImAttrMedia_im_attr_media_1000base_dwdm_1541 ImAttrMedia = "im-attr-media-1000base-dwdm-1541"

    // im attr media 1000base dwdm 1545
    ImAttrMedia_im_attr_media_1000base_dwdm_1545 ImAttrMedia = "im-attr-media-1000base-dwdm-1545"

    // im attr media 1000base dwdm 1549
    ImAttrMedia_im_attr_media_1000base_dwdm_1549 ImAttrMedia = "im-attr-media-1000base-dwdm-1549"

    // im attr media 1000base dwdm 1553
    ImAttrMedia_im_attr_media_1000base_dwdm_1553 ImAttrMedia = "im-attr-media-1000base-dwdm-1553"

    // im attr media 1000base dwdm 1557
    ImAttrMedia_im_attr_media_1000base_dwdm_1557 ImAttrMedia = "im-attr-media-1000base-dwdm-1557"

    // im attr media 1000base dwdm 1561
    ImAttrMedia_im_attr_media_1000base_dwdm_1561 ImAttrMedia = "im-attr-media-1000base-dwdm-1561"

    // im attr media 40gbase lr4
    ImAttrMedia_im_attr_media_40gbase_lr4 ImAttrMedia = "im-attr-media-40gbase-lr4"

    // im attr media 40gbase er4
    ImAttrMedia_im_attr_media_40gbase_er4 ImAttrMedia = "im-attr-media-40gbase-er4"

    // im attr media 100gbase er4
    ImAttrMedia_im_attr_media_100gbase_er4 ImAttrMedia = "im-attr-media-100gbase-er4"

    // im attr media 1000base ex
    ImAttrMedia_im_attr_media_1000base_ex ImAttrMedia = "im-attr-media-1000base-ex"

    // im attr media 1000base bx10 d
    ImAttrMedia_im_attr_media_1000base_bx10_d ImAttrMedia = "im-attr-media-1000base-bx10-d"

    // im attr media 1000base bx10 u
    ImAttrMedia_im_attr_media_1000base_bx10_u ImAttrMedia = "im-attr-media-1000base-bx10-u"

    // im attr media 1000base dwdm 1561 42
    ImAttrMedia_im_attr_media_1000base_dwdm_1561_42 ImAttrMedia = "im-attr-media-1000base-dwdm-1561-42"

    // im attr media 1000base dwdm 1560 61
    ImAttrMedia_im_attr_media_1000base_dwdm_1560_61 ImAttrMedia = "im-attr-media-1000base-dwdm-1560-61"

    // im attr media 1000base dwdm 1559 79
    ImAttrMedia_im_attr_media_1000base_dwdm_1559_79 ImAttrMedia = "im-attr-media-1000base-dwdm-1559-79"

    // im attr media 1000base dwdm 1558 98
    ImAttrMedia_im_attr_media_1000base_dwdm_1558_98 ImAttrMedia = "im-attr-media-1000base-dwdm-1558-98"

    // im attr media 1000base dwdm 1558 17
    ImAttrMedia_im_attr_media_1000base_dwdm_1558_17 ImAttrMedia = "im-attr-media-1000base-dwdm-1558-17"

    // im attr media 1000base dwdm 1557 36
    ImAttrMedia_im_attr_media_1000base_dwdm_1557_36 ImAttrMedia = "im-attr-media-1000base-dwdm-1557-36"

    // im attr media 1000base dwdm 1556 55
    ImAttrMedia_im_attr_media_1000base_dwdm_1556_55 ImAttrMedia = "im-attr-media-1000base-dwdm-1556-55"

    // im attr media 1000base dwdm 1555 75
    ImAttrMedia_im_attr_media_1000base_dwdm_1555_75 ImAttrMedia = "im-attr-media-1000base-dwdm-1555-75"

    // im attr media 1000base dwdm 1554 94
    ImAttrMedia_im_attr_media_1000base_dwdm_1554_94 ImAttrMedia = "im-attr-media-1000base-dwdm-1554-94"

    // im attr media 1000base dwdm 1554 13
    ImAttrMedia_im_attr_media_1000base_dwdm_1554_13 ImAttrMedia = "im-attr-media-1000base-dwdm-1554-13"

    // im attr media 1000base dwdm 1553 33
    ImAttrMedia_im_attr_media_1000base_dwdm_1553_33 ImAttrMedia = "im-attr-media-1000base-dwdm-1553-33"

    // im attr media 1000base dwdm 1552 52
    ImAttrMedia_im_attr_media_1000base_dwdm_1552_52 ImAttrMedia = "im-attr-media-1000base-dwdm-1552-52"

    // im attr media 1000base dwdm 1551 72
    ImAttrMedia_im_attr_media_1000base_dwdm_1551_72 ImAttrMedia = "im-attr-media-1000base-dwdm-1551-72"

    // im attr media 1000base dwdm 1550 92
    ImAttrMedia_im_attr_media_1000base_dwdm_1550_92 ImAttrMedia = "im-attr-media-1000base-dwdm-1550-92"

    // im attr media 1000base dwdm 1550 12
    ImAttrMedia_im_attr_media_1000base_dwdm_1550_12 ImAttrMedia = "im-attr-media-1000base-dwdm-1550-12"

    // im attr media 1000base dwdm 1549 32
    ImAttrMedia_im_attr_media_1000base_dwdm_1549_32 ImAttrMedia = "im-attr-media-1000base-dwdm-1549-32"

    // im attr media 1000base dwdm 1548 51
    ImAttrMedia_im_attr_media_1000base_dwdm_1548_51 ImAttrMedia = "im-attr-media-1000base-dwdm-1548-51"

    // im attr media 1000base dwdm 1547 72
    ImAttrMedia_im_attr_media_1000base_dwdm_1547_72 ImAttrMedia = "im-attr-media-1000base-dwdm-1547-72"

    // im attr media 1000base dwdm 1546 92
    ImAttrMedia_im_attr_media_1000base_dwdm_1546_92 ImAttrMedia = "im-attr-media-1000base-dwdm-1546-92"

    // im attr media 1000base dwdm 1546 12
    ImAttrMedia_im_attr_media_1000base_dwdm_1546_12 ImAttrMedia = "im-attr-media-1000base-dwdm-1546-12"

    // im attr media 1000base dwdm 1545 32
    ImAttrMedia_im_attr_media_1000base_dwdm_1545_32 ImAttrMedia = "im-attr-media-1000base-dwdm-1545-32"

    // im attr media 1000base dwdm 1544 53
    ImAttrMedia_im_attr_media_1000base_dwdm_1544_53 ImAttrMedia = "im-attr-media-1000base-dwdm-1544-53"

    // im attr media 1000base dwdm 1543 73
    ImAttrMedia_im_attr_media_1000base_dwdm_1543_73 ImAttrMedia = "im-attr-media-1000base-dwdm-1543-73"

    // im attr media 1000base dwdm 1542 94
    ImAttrMedia_im_attr_media_1000base_dwdm_1542_94 ImAttrMedia = "im-attr-media-1000base-dwdm-1542-94"

    // im attr media 1000base dwdm 1542 14
    ImAttrMedia_im_attr_media_1000base_dwdm_1542_14 ImAttrMedia = "im-attr-media-1000base-dwdm-1542-14"

    // im attr media 1000base dwdm 1541 35
    ImAttrMedia_im_attr_media_1000base_dwdm_1541_35 ImAttrMedia = "im-attr-media-1000base-dwdm-1541-35"

    // im attr media 1000base dwdm 1540 56
    ImAttrMedia_im_attr_media_1000base_dwdm_1540_56 ImAttrMedia = "im-attr-media-1000base-dwdm-1540-56"

    // im attr media 1000base dwdm 1539 77
    ImAttrMedia_im_attr_media_1000base_dwdm_1539_77 ImAttrMedia = "im-attr-media-1000base-dwdm-1539-77"

    // im attr media 1000base dwdm 1538 98
    ImAttrMedia_im_attr_media_1000base_dwdm_1538_98 ImAttrMedia = "im-attr-media-1000base-dwdm-1538-98"

    // im attr media 1000base dwdm 1538 19
    ImAttrMedia_im_attr_media_1000base_dwdm_1538_19 ImAttrMedia = "im-attr-media-1000base-dwdm-1538-19"

    // im attr media 1000base dwdm 1537 40
    ImAttrMedia_im_attr_media_1000base_dwdm_1537_40 ImAttrMedia = "im-attr-media-1000base-dwdm-1537-40"

    // im attr media 1000base dwdm 1536 61
    ImAttrMedia_im_attr_media_1000base_dwdm_1536_61 ImAttrMedia = "im-attr-media-1000base-dwdm-1536-61"

    // im attr media 1000base dwdm 1535 82
    ImAttrMedia_im_attr_media_1000base_dwdm_1535_82 ImAttrMedia = "im-attr-media-1000base-dwdm-1535-82"

    // im attr media 1000base dwdm 1535 04
    ImAttrMedia_im_attr_media_1000base_dwdm_1535_04 ImAttrMedia = "im-attr-media-1000base-dwdm-1535-04"

    // im attr media 1000base dwdm 1534 25
    ImAttrMedia_im_attr_media_1000base_dwdm_1534_25 ImAttrMedia = "im-attr-media-1000base-dwdm-1534-25"

    // im attr media 1000base dwdm 1533 47
    ImAttrMedia_im_attr_media_1000base_dwdm_1533_47 ImAttrMedia = "im-attr-media-1000base-dwdm-1533-47"

    // im attr media 1000base dwdm 1532 68
    ImAttrMedia_im_attr_media_1000base_dwdm_1532_68 ImAttrMedia = "im-attr-media-1000base-dwdm-1532-68"

    // im attr media 1000base dwdm 1531 90
    ImAttrMedia_im_attr_media_1000base_dwdm_1531_90 ImAttrMedia = "im-attr-media-1000base-dwdm-1531-90"

    // im attr media 1000base dwdm 1531 12
    ImAttrMedia_im_attr_media_1000base_dwdm_1531_12 ImAttrMedia = "im-attr-media-1000base-dwdm-1531-12"

    // im attr media 1000base dwdm 1530 33
    ImAttrMedia_im_attr_media_1000base_dwdm_1530_33 ImAttrMedia = "im-attr-media-1000base-dwdm-1530-33"

    // im attr media 1000base dwdm tunable
    ImAttrMedia_im_attr_media_1000base_dwdm_tunable ImAttrMedia = "im-attr-media-1000base-dwdm-tunable"

    // im attr media 10gbase dwdm 1561 42
    ImAttrMedia_im_attr_media_10gbase_dwdm_1561_42 ImAttrMedia = "im-attr-media-10gbase-dwdm-1561-42"

    // im attr media 10gbase dwdm 1560 61
    ImAttrMedia_im_attr_media_10gbase_dwdm_1560_61 ImAttrMedia = "im-attr-media-10gbase-dwdm-1560-61"

    // im attr media 10gbase dwdm 1559 79
    ImAttrMedia_im_attr_media_10gbase_dwdm_1559_79 ImAttrMedia = "im-attr-media-10gbase-dwdm-1559-79"

    // im attr media 10gbase dwdm 1558 98
    ImAttrMedia_im_attr_media_10gbase_dwdm_1558_98 ImAttrMedia = "im-attr-media-10gbase-dwdm-1558-98"

    // im attr media 10gbase dwdm 1558 17
    ImAttrMedia_im_attr_media_10gbase_dwdm_1558_17 ImAttrMedia = "im-attr-media-10gbase-dwdm-1558-17"

    // im attr media 10gbase dwdm 1557 36
    ImAttrMedia_im_attr_media_10gbase_dwdm_1557_36 ImAttrMedia = "im-attr-media-10gbase-dwdm-1557-36"

    // im attr media 10gbase dwdm 1556 55
    ImAttrMedia_im_attr_media_10gbase_dwdm_1556_55 ImAttrMedia = "im-attr-media-10gbase-dwdm-1556-55"

    // im attr media 10gbase dwdm 1555 75
    ImAttrMedia_im_attr_media_10gbase_dwdm_1555_75 ImAttrMedia = "im-attr-media-10gbase-dwdm-1555-75"

    // im attr media 10gbase dwdm 1554 94
    ImAttrMedia_im_attr_media_10gbase_dwdm_1554_94 ImAttrMedia = "im-attr-media-10gbase-dwdm-1554-94"

    // im attr media 10gbase dwdm 1554 13
    ImAttrMedia_im_attr_media_10gbase_dwdm_1554_13 ImAttrMedia = "im-attr-media-10gbase-dwdm-1554-13"

    // im attr media 10gbase dwdm 1553 33
    ImAttrMedia_im_attr_media_10gbase_dwdm_1553_33 ImAttrMedia = "im-attr-media-10gbase-dwdm-1553-33"

    // im attr media 10gbase dwdm 1552 52
    ImAttrMedia_im_attr_media_10gbase_dwdm_1552_52 ImAttrMedia = "im-attr-media-10gbase-dwdm-1552-52"

    // im attr media 10gbase dwdm 1551 72
    ImAttrMedia_im_attr_media_10gbase_dwdm_1551_72 ImAttrMedia = "im-attr-media-10gbase-dwdm-1551-72"

    // im attr media 10gbase dwdm 1550 92
    ImAttrMedia_im_attr_media_10gbase_dwdm_1550_92 ImAttrMedia = "im-attr-media-10gbase-dwdm-1550-92"

    // im attr media 10gbase dwdm 1550 12
    ImAttrMedia_im_attr_media_10gbase_dwdm_1550_12 ImAttrMedia = "im-attr-media-10gbase-dwdm-1550-12"

    // im attr media 10gbase dwdm 1549 32
    ImAttrMedia_im_attr_media_10gbase_dwdm_1549_32 ImAttrMedia = "im-attr-media-10gbase-dwdm-1549-32"

    // im attr media 10gbase dwdm 1548 51
    ImAttrMedia_im_attr_media_10gbase_dwdm_1548_51 ImAttrMedia = "im-attr-media-10gbase-dwdm-1548-51"

    // im attr media 10gbase dwdm 1547 72
    ImAttrMedia_im_attr_media_10gbase_dwdm_1547_72 ImAttrMedia = "im-attr-media-10gbase-dwdm-1547-72"

    // im attr media 10gbase dwdm 1546 92
    ImAttrMedia_im_attr_media_10gbase_dwdm_1546_92 ImAttrMedia = "im-attr-media-10gbase-dwdm-1546-92"

    // im attr media 10gbase dwdm 1546 12
    ImAttrMedia_im_attr_media_10gbase_dwdm_1546_12 ImAttrMedia = "im-attr-media-10gbase-dwdm-1546-12"

    // im attr media 10gbase dwdm 1545 32
    ImAttrMedia_im_attr_media_10gbase_dwdm_1545_32 ImAttrMedia = "im-attr-media-10gbase-dwdm-1545-32"

    // im attr media 10gbase dwdm 1544 53
    ImAttrMedia_im_attr_media_10gbase_dwdm_1544_53 ImAttrMedia = "im-attr-media-10gbase-dwdm-1544-53"

    // im attr media 10gbase dwdm 1543 73
    ImAttrMedia_im_attr_media_10gbase_dwdm_1543_73 ImAttrMedia = "im-attr-media-10gbase-dwdm-1543-73"

    // im attr media 10gbase dwdm 1542 94
    ImAttrMedia_im_attr_media_10gbase_dwdm_1542_94 ImAttrMedia = "im-attr-media-10gbase-dwdm-1542-94"

    // im attr media 10gbase dwdm 1542 14
    ImAttrMedia_im_attr_media_10gbase_dwdm_1542_14 ImAttrMedia = "im-attr-media-10gbase-dwdm-1542-14"

    // im attr media 10gbase dwdm 1541 35
    ImAttrMedia_im_attr_media_10gbase_dwdm_1541_35 ImAttrMedia = "im-attr-media-10gbase-dwdm-1541-35"

    // im attr media 10gbase dwdm 1540 56
    ImAttrMedia_im_attr_media_10gbase_dwdm_1540_56 ImAttrMedia = "im-attr-media-10gbase-dwdm-1540-56"

    // im attr media 10gbase dwdm 1539 77
    ImAttrMedia_im_attr_media_10gbase_dwdm_1539_77 ImAttrMedia = "im-attr-media-10gbase-dwdm-1539-77"

    // im attr media 10gbase dwdm 1538 98
    ImAttrMedia_im_attr_media_10gbase_dwdm_1538_98 ImAttrMedia = "im-attr-media-10gbase-dwdm-1538-98"

    // im attr media 10gbase dwdm 1538 19
    ImAttrMedia_im_attr_media_10gbase_dwdm_1538_19 ImAttrMedia = "im-attr-media-10gbase-dwdm-1538-19"

    // im attr media 10gbase dwdm 1537 40
    ImAttrMedia_im_attr_media_10gbase_dwdm_1537_40 ImAttrMedia = "im-attr-media-10gbase-dwdm-1537-40"

    // im attr media 10gbase dwdm 1536 61
    ImAttrMedia_im_attr_media_10gbase_dwdm_1536_61 ImAttrMedia = "im-attr-media-10gbase-dwdm-1536-61"

    // im attr media 10gbase dwdm 1535 82
    ImAttrMedia_im_attr_media_10gbase_dwdm_1535_82 ImAttrMedia = "im-attr-media-10gbase-dwdm-1535-82"

    // im attr media 10gbase dwdm 1535 04
    ImAttrMedia_im_attr_media_10gbase_dwdm_1535_04 ImAttrMedia = "im-attr-media-10gbase-dwdm-1535-04"

    // im attr media 10gbase dwdm 1534 25
    ImAttrMedia_im_attr_media_10gbase_dwdm_1534_25 ImAttrMedia = "im-attr-media-10gbase-dwdm-1534-25"

    // im attr media 10gbase dwdm 1533 47
    ImAttrMedia_im_attr_media_10gbase_dwdm_1533_47 ImAttrMedia = "im-attr-media-10gbase-dwdm-1533-47"

    // im attr media 10gbase dwdm 1532 68
    ImAttrMedia_im_attr_media_10gbase_dwdm_1532_68 ImAttrMedia = "im-attr-media-10gbase-dwdm-1532-68"

    // im attr media 10gbase dwdm 1531 90
    ImAttrMedia_im_attr_media_10gbase_dwdm_1531_90 ImAttrMedia = "im-attr-media-10gbase-dwdm-1531-90"

    // im attr media 10gbase dwdm 1531 12
    ImAttrMedia_im_attr_media_10gbase_dwdm_1531_12 ImAttrMedia = "im-attr-media-10gbase-dwdm-1531-12"

    // im attr media 10gbase dwdm 1530 33
    ImAttrMedia_im_attr_media_10gbase_dwdm_1530_33 ImAttrMedia = "im-attr-media-10gbase-dwdm-1530-33"

    // im attr media 10gbase dwdm tunable
    ImAttrMedia_im_attr_media_10gbase_dwdm_tunable ImAttrMedia = "im-attr-media-10gbase-dwdm-tunable"

    // im attr media 40gbase dwdm 1561 42
    ImAttrMedia_im_attr_media_40gbase_dwdm_1561_42 ImAttrMedia = "im-attr-media-40gbase-dwdm-1561-42"

    // im attr media 40gbase dwdm 1560 61
    ImAttrMedia_im_attr_media_40gbase_dwdm_1560_61 ImAttrMedia = "im-attr-media-40gbase-dwdm-1560-61"

    // im attr media 40gbase dwdm 1559 79
    ImAttrMedia_im_attr_media_40gbase_dwdm_1559_79 ImAttrMedia = "im-attr-media-40gbase-dwdm-1559-79"

    // im attr media 40gbase dwdm 1558 98
    ImAttrMedia_im_attr_media_40gbase_dwdm_1558_98 ImAttrMedia = "im-attr-media-40gbase-dwdm-1558-98"

    // im attr media 40gbase dwdm 1558 17
    ImAttrMedia_im_attr_media_40gbase_dwdm_1558_17 ImAttrMedia = "im-attr-media-40gbase-dwdm-1558-17"

    // im attr media 40gbase dwdm 1557 36
    ImAttrMedia_im_attr_media_40gbase_dwdm_1557_36 ImAttrMedia = "im-attr-media-40gbase-dwdm-1557-36"

    // im attr media 40gbase dwdm 1556 55
    ImAttrMedia_im_attr_media_40gbase_dwdm_1556_55 ImAttrMedia = "im-attr-media-40gbase-dwdm-1556-55"

    // im attr media 40gbase dwdm 1555 75
    ImAttrMedia_im_attr_media_40gbase_dwdm_1555_75 ImAttrMedia = "im-attr-media-40gbase-dwdm-1555-75"

    // im attr media 40gbase dwdm 1554 94
    ImAttrMedia_im_attr_media_40gbase_dwdm_1554_94 ImAttrMedia = "im-attr-media-40gbase-dwdm-1554-94"

    // im attr media 40gbase dwdm 1554 13
    ImAttrMedia_im_attr_media_40gbase_dwdm_1554_13 ImAttrMedia = "im-attr-media-40gbase-dwdm-1554-13"

    // im attr media 40gbase dwdm 1553 33
    ImAttrMedia_im_attr_media_40gbase_dwdm_1553_33 ImAttrMedia = "im-attr-media-40gbase-dwdm-1553-33"

    // im attr media 40gbase dwdm 1552 52
    ImAttrMedia_im_attr_media_40gbase_dwdm_1552_52 ImAttrMedia = "im-attr-media-40gbase-dwdm-1552-52"

    // im attr media 40gbase dwdm 1551 72
    ImAttrMedia_im_attr_media_40gbase_dwdm_1551_72 ImAttrMedia = "im-attr-media-40gbase-dwdm-1551-72"

    // im attr media 40gbase dwdm 1550 92
    ImAttrMedia_im_attr_media_40gbase_dwdm_1550_92 ImAttrMedia = "im-attr-media-40gbase-dwdm-1550-92"

    // im attr media 40gbase dwdm 1550 12
    ImAttrMedia_im_attr_media_40gbase_dwdm_1550_12 ImAttrMedia = "im-attr-media-40gbase-dwdm-1550-12"

    // im attr media 40gbase dwdm 1549 32
    ImAttrMedia_im_attr_media_40gbase_dwdm_1549_32 ImAttrMedia = "im-attr-media-40gbase-dwdm-1549-32"

    // im attr media 40gbase dwdm 1548 51
    ImAttrMedia_im_attr_media_40gbase_dwdm_1548_51 ImAttrMedia = "im-attr-media-40gbase-dwdm-1548-51"

    // im attr media 40gbase dwdm 1547 72
    ImAttrMedia_im_attr_media_40gbase_dwdm_1547_72 ImAttrMedia = "im-attr-media-40gbase-dwdm-1547-72"

    // im attr media 40gbase dwdm 1546 92
    ImAttrMedia_im_attr_media_40gbase_dwdm_1546_92 ImAttrMedia = "im-attr-media-40gbase-dwdm-1546-92"

    // im attr media 40gbase dwdm 1546 12
    ImAttrMedia_im_attr_media_40gbase_dwdm_1546_12 ImAttrMedia = "im-attr-media-40gbase-dwdm-1546-12"

    // im attr media 40gbase dwdm 1545 32
    ImAttrMedia_im_attr_media_40gbase_dwdm_1545_32 ImAttrMedia = "im-attr-media-40gbase-dwdm-1545-32"

    // im attr media 40gbase dwdm 1544 53
    ImAttrMedia_im_attr_media_40gbase_dwdm_1544_53 ImAttrMedia = "im-attr-media-40gbase-dwdm-1544-53"

    // im attr media 40gbase dwdm 1543 73
    ImAttrMedia_im_attr_media_40gbase_dwdm_1543_73 ImAttrMedia = "im-attr-media-40gbase-dwdm-1543-73"

    // im attr media 40gbase dwdm 1542 94
    ImAttrMedia_im_attr_media_40gbase_dwdm_1542_94 ImAttrMedia = "im-attr-media-40gbase-dwdm-1542-94"

    // im attr media 40gbase dwdm 1542 14
    ImAttrMedia_im_attr_media_40gbase_dwdm_1542_14 ImAttrMedia = "im-attr-media-40gbase-dwdm-1542-14"

    // im attr media 40gbase dwdm 1541 35
    ImAttrMedia_im_attr_media_40gbase_dwdm_1541_35 ImAttrMedia = "im-attr-media-40gbase-dwdm-1541-35"

    // im attr media 40gbase dwdm 1540 56
    ImAttrMedia_im_attr_media_40gbase_dwdm_1540_56 ImAttrMedia = "im-attr-media-40gbase-dwdm-1540-56"

    // im attr media 40gbase dwdm 1539 77
    ImAttrMedia_im_attr_media_40gbase_dwdm_1539_77 ImAttrMedia = "im-attr-media-40gbase-dwdm-1539-77"

    // im attr media 40gbase dwdm 1538 98
    ImAttrMedia_im_attr_media_40gbase_dwdm_1538_98 ImAttrMedia = "im-attr-media-40gbase-dwdm-1538-98"

    // im attr media 40gbase dwdm 1538 19
    ImAttrMedia_im_attr_media_40gbase_dwdm_1538_19 ImAttrMedia = "im-attr-media-40gbase-dwdm-1538-19"

    // im attr media 40gbase dwdm 1537 40
    ImAttrMedia_im_attr_media_40gbase_dwdm_1537_40 ImAttrMedia = "im-attr-media-40gbase-dwdm-1537-40"

    // im attr media 40gbase dwdm 1536 61
    ImAttrMedia_im_attr_media_40gbase_dwdm_1536_61 ImAttrMedia = "im-attr-media-40gbase-dwdm-1536-61"

    // im attr media 40gbase dwdm 1535 82
    ImAttrMedia_im_attr_media_40gbase_dwdm_1535_82 ImAttrMedia = "im-attr-media-40gbase-dwdm-1535-82"

    // im attr media 40gbase dwdm 1535 04
    ImAttrMedia_im_attr_media_40gbase_dwdm_1535_04 ImAttrMedia = "im-attr-media-40gbase-dwdm-1535-04"

    // im attr media 40gbase dwdm 1534 25
    ImAttrMedia_im_attr_media_40gbase_dwdm_1534_25 ImAttrMedia = "im-attr-media-40gbase-dwdm-1534-25"

    // im attr media 40gbase dwdm 1533 47
    ImAttrMedia_im_attr_media_40gbase_dwdm_1533_47 ImAttrMedia = "im-attr-media-40gbase-dwdm-1533-47"

    // im attr media 40gbase dwdm 1532 68
    ImAttrMedia_im_attr_media_40gbase_dwdm_1532_68 ImAttrMedia = "im-attr-media-40gbase-dwdm-1532-68"

    // im attr media 40gbase dwdm 1531 90
    ImAttrMedia_im_attr_media_40gbase_dwdm_1531_90 ImAttrMedia = "im-attr-media-40gbase-dwdm-1531-90"

    // im attr media 40gbase dwdm 1531 12
    ImAttrMedia_im_attr_media_40gbase_dwdm_1531_12 ImAttrMedia = "im-attr-media-40gbase-dwdm-1531-12"

    // im attr media 40gbase dwdm 1530 33
    ImAttrMedia_im_attr_media_40gbase_dwdm_1530_33 ImAttrMedia = "im-attr-media-40gbase-dwdm-1530-33"

    // im attr media 40gbase dwdm tunable
    ImAttrMedia_im_attr_media_40gbase_dwdm_tunable ImAttrMedia = "im-attr-media-40gbase-dwdm-tunable"

    // im attr media 100gbase dwdm 1561 42
    ImAttrMedia_im_attr_media_100gbase_dwdm_1561_42 ImAttrMedia = "im-attr-media-100gbase-dwdm-1561-42"

    // im attr media 100gbase dwdm 1560 61
    ImAttrMedia_im_attr_media_100gbase_dwdm_1560_61 ImAttrMedia = "im-attr-media-100gbase-dwdm-1560-61"

    // im attr media 100gbase dwdm 1559 79
    ImAttrMedia_im_attr_media_100gbase_dwdm_1559_79 ImAttrMedia = "im-attr-media-100gbase-dwdm-1559-79"

    // im attr media 100gbase dwdm 1558 98
    ImAttrMedia_im_attr_media_100gbase_dwdm_1558_98 ImAttrMedia = "im-attr-media-100gbase-dwdm-1558-98"

    // im attr media 100gbase dwdm 1558 17
    ImAttrMedia_im_attr_media_100gbase_dwdm_1558_17 ImAttrMedia = "im-attr-media-100gbase-dwdm-1558-17"

    // im attr media 100gbase dwdm 1557 36
    ImAttrMedia_im_attr_media_100gbase_dwdm_1557_36 ImAttrMedia = "im-attr-media-100gbase-dwdm-1557-36"

    // im attr media 100gbase dwdm 1556 55
    ImAttrMedia_im_attr_media_100gbase_dwdm_1556_55 ImAttrMedia = "im-attr-media-100gbase-dwdm-1556-55"

    // im attr media 100gbase dwdm 1555 75
    ImAttrMedia_im_attr_media_100gbase_dwdm_1555_75 ImAttrMedia = "im-attr-media-100gbase-dwdm-1555-75"

    // im attr media 100gbase dwdm 1554 94
    ImAttrMedia_im_attr_media_100gbase_dwdm_1554_94 ImAttrMedia = "im-attr-media-100gbase-dwdm-1554-94"

    // im attr media 100gbase dwdm 1554 13
    ImAttrMedia_im_attr_media_100gbase_dwdm_1554_13 ImAttrMedia = "im-attr-media-100gbase-dwdm-1554-13"

    // im attr media 100gbase dwdm 1553 33
    ImAttrMedia_im_attr_media_100gbase_dwdm_1553_33 ImAttrMedia = "im-attr-media-100gbase-dwdm-1553-33"

    // im attr media 100gbase dwdm 1552 52
    ImAttrMedia_im_attr_media_100gbase_dwdm_1552_52 ImAttrMedia = "im-attr-media-100gbase-dwdm-1552-52"

    // im attr media 100gbase dwdm 1551 72
    ImAttrMedia_im_attr_media_100gbase_dwdm_1551_72 ImAttrMedia = "im-attr-media-100gbase-dwdm-1551-72"

    // im attr media 100gbase dwdm 1550 92
    ImAttrMedia_im_attr_media_100gbase_dwdm_1550_92 ImAttrMedia = "im-attr-media-100gbase-dwdm-1550-92"

    // im attr media 100gbase dwdm 1550 12
    ImAttrMedia_im_attr_media_100gbase_dwdm_1550_12 ImAttrMedia = "im-attr-media-100gbase-dwdm-1550-12"

    // im attr media 100gbase dwdm 1549 32
    ImAttrMedia_im_attr_media_100gbase_dwdm_1549_32 ImAttrMedia = "im-attr-media-100gbase-dwdm-1549-32"

    // im attr media 100gbase dwdm 1548 51
    ImAttrMedia_im_attr_media_100gbase_dwdm_1548_51 ImAttrMedia = "im-attr-media-100gbase-dwdm-1548-51"

    // im attr media 100gbase dwdm 1547 72
    ImAttrMedia_im_attr_media_100gbase_dwdm_1547_72 ImAttrMedia = "im-attr-media-100gbase-dwdm-1547-72"

    // im attr media 100gbase dwdm 1546 92
    ImAttrMedia_im_attr_media_100gbase_dwdm_1546_92 ImAttrMedia = "im-attr-media-100gbase-dwdm-1546-92"

    // im attr media 100gbase dwdm 1546 12
    ImAttrMedia_im_attr_media_100gbase_dwdm_1546_12 ImAttrMedia = "im-attr-media-100gbase-dwdm-1546-12"

    // im attr media 100gbase dwdm 1545 32
    ImAttrMedia_im_attr_media_100gbase_dwdm_1545_32 ImAttrMedia = "im-attr-media-100gbase-dwdm-1545-32"

    // im attr media 100gbase dwdm 1544 53
    ImAttrMedia_im_attr_media_100gbase_dwdm_1544_53 ImAttrMedia = "im-attr-media-100gbase-dwdm-1544-53"

    // im attr media 100gbase dwdm 1543 73
    ImAttrMedia_im_attr_media_100gbase_dwdm_1543_73 ImAttrMedia = "im-attr-media-100gbase-dwdm-1543-73"

    // im attr media 100gbase dwdm 1542 94
    ImAttrMedia_im_attr_media_100gbase_dwdm_1542_94 ImAttrMedia = "im-attr-media-100gbase-dwdm-1542-94"

    // im attr media 100gbase dwdm 1542 14
    ImAttrMedia_im_attr_media_100gbase_dwdm_1542_14 ImAttrMedia = "im-attr-media-100gbase-dwdm-1542-14"

    // im attr media 100gbase dwdm 1541 35
    ImAttrMedia_im_attr_media_100gbase_dwdm_1541_35 ImAttrMedia = "im-attr-media-100gbase-dwdm-1541-35"

    // im attr media 100gbase dwdm 1540 56
    ImAttrMedia_im_attr_media_100gbase_dwdm_1540_56 ImAttrMedia = "im-attr-media-100gbase-dwdm-1540-56"

    // im attr media 100gbase dwdm 1539 77
    ImAttrMedia_im_attr_media_100gbase_dwdm_1539_77 ImAttrMedia = "im-attr-media-100gbase-dwdm-1539-77"

    // im attr media 100gbase dwdm 1538 98
    ImAttrMedia_im_attr_media_100gbase_dwdm_1538_98 ImAttrMedia = "im-attr-media-100gbase-dwdm-1538-98"

    // im attr media 100gbase dwdm 1538 19
    ImAttrMedia_im_attr_media_100gbase_dwdm_1538_19 ImAttrMedia = "im-attr-media-100gbase-dwdm-1538-19"

    // im attr media 100gbase dwdm 1537 40
    ImAttrMedia_im_attr_media_100gbase_dwdm_1537_40 ImAttrMedia = "im-attr-media-100gbase-dwdm-1537-40"

    // im attr media 100gbase dwdm 1536 61
    ImAttrMedia_im_attr_media_100gbase_dwdm_1536_61 ImAttrMedia = "im-attr-media-100gbase-dwdm-1536-61"

    // im attr media 100gbase dwdm 1535 82
    ImAttrMedia_im_attr_media_100gbase_dwdm_1535_82 ImAttrMedia = "im-attr-media-100gbase-dwdm-1535-82"

    // im attr media 100gbase dwdm 1535 04
    ImAttrMedia_im_attr_media_100gbase_dwdm_1535_04 ImAttrMedia = "im-attr-media-100gbase-dwdm-1535-04"

    // im attr media 100gbase dwdm 1534 25
    ImAttrMedia_im_attr_media_100gbase_dwdm_1534_25 ImAttrMedia = "im-attr-media-100gbase-dwdm-1534-25"

    // im attr media 100gbase dwdm 1533 47
    ImAttrMedia_im_attr_media_100gbase_dwdm_1533_47 ImAttrMedia = "im-attr-media-100gbase-dwdm-1533-47"

    // im attr media 100gbase dwdm 1532 68
    ImAttrMedia_im_attr_media_100gbase_dwdm_1532_68 ImAttrMedia = "im-attr-media-100gbase-dwdm-1532-68"

    // im attr media 100gbase dwdm 1531 90
    ImAttrMedia_im_attr_media_100gbase_dwdm_1531_90 ImAttrMedia = "im-attr-media-100gbase-dwdm-1531-90"

    // im attr media 100gbase dwdm 1531 12
    ImAttrMedia_im_attr_media_100gbase_dwdm_1531_12 ImAttrMedia = "im-attr-media-100gbase-dwdm-1531-12"

    // im attr media 100gbase dwdm 1530 33
    ImAttrMedia_im_attr_media_100gbase_dwdm_1530_33 ImAttrMedia = "im-attr-media-100gbase-dwdm-1530-33"

    // im attr media 100gbase dwdm tunable
    ImAttrMedia_im_attr_media_100gbase_dwdm_tunable ImAttrMedia = "im-attr-media-100gbase-dwdm-tunable"

    // im attr media 40gbase kr4
    ImAttrMedia_im_attr_media_40gbase_kr4 ImAttrMedia = "im-attr-media-40gbase-kr4"

    // im attr media 40gbase cr4
    ImAttrMedia_im_attr_media_40gbase_cr4 ImAttrMedia = "im-attr-media-40gbase-cr4"

    // im attr media 40gbase sr4
    ImAttrMedia_im_attr_media_40gbase_sr4 ImAttrMedia = "im-attr-media-40gbase-sr4"

    // im attr media 40gbase fr
    ImAttrMedia_im_attr_media_40gbase_fr ImAttrMedia = "im-attr-media-40gbase-fr"

    // im attr media 100gbase cr10
    ImAttrMedia_im_attr_media_100gbase_cr10 ImAttrMedia = "im-attr-media-100gbase-cr10"

    // im attr media 100gbase sr10
    ImAttrMedia_im_attr_media_100gbase_sr10 ImAttrMedia = "im-attr-media-100gbase-sr10"

    // im attr media 40gbase csr4
    ImAttrMedia_im_attr_media_40gbase_csr4 ImAttrMedia = "im-attr-media-40gbase-csr4"

    // im attr media 10gbase cwdm
    ImAttrMedia_im_attr_media_10gbase_cwdm ImAttrMedia = "im-attr-media-10gbase-cwdm"

    // im attr media 10gbase cwdm tunable
    ImAttrMedia_im_attr_media_10gbase_cwdm_tunable ImAttrMedia = "im-attr-media-10gbase-cwdm-tunable"

    // im attr media 10gbase cwdm 1470
    ImAttrMedia_im_attr_media_10gbase_cwdm_1470 ImAttrMedia = "im-attr-media-10gbase-cwdm-1470"

    // im attr media 10gbase cwdm 1490
    ImAttrMedia_im_attr_media_10gbase_cwdm_1490 ImAttrMedia = "im-attr-media-10gbase-cwdm-1490"

    // im attr media 10gbase cwdm 1510
    ImAttrMedia_im_attr_media_10gbase_cwdm_1510 ImAttrMedia = "im-attr-media-10gbase-cwdm-1510"

    // im attr media 10gbase cwdm 1530
    ImAttrMedia_im_attr_media_10gbase_cwdm_1530 ImAttrMedia = "im-attr-media-10gbase-cwdm-1530"

    // im attr media 10gbase cwdm 1550
    ImAttrMedia_im_attr_media_10gbase_cwdm_1550 ImAttrMedia = "im-attr-media-10gbase-cwdm-1550"

    // im attr media 10gbase cwdm 1570
    ImAttrMedia_im_attr_media_10gbase_cwdm_1570 ImAttrMedia = "im-attr-media-10gbase-cwdm-1570"

    // im attr media 10gbase cwdm 1590
    ImAttrMedia_im_attr_media_10gbase_cwdm_1590 ImAttrMedia = "im-attr-media-10gbase-cwdm-1590"

    // im attr media 10gbase cwdm 1610
    ImAttrMedia_im_attr_media_10gbase_cwdm_1610 ImAttrMedia = "im-attr-media-10gbase-cwdm-1610"

    // im attr media 40gbase cwdm
    ImAttrMedia_im_attr_media_40gbase_cwdm ImAttrMedia = "im-attr-media-40gbase-cwdm"

    // im attr media 40gbase cwdm tunable
    ImAttrMedia_im_attr_media_40gbase_cwdm_tunable ImAttrMedia = "im-attr-media-40gbase-cwdm-tunable"

    // im attr media 40gbase cwdm 1470
    ImAttrMedia_im_attr_media_40gbase_cwdm_1470 ImAttrMedia = "im-attr-media-40gbase-cwdm-1470"

    // im attr media 40gbase cwdm 1490
    ImAttrMedia_im_attr_media_40gbase_cwdm_1490 ImAttrMedia = "im-attr-media-40gbase-cwdm-1490"

    // im attr media 40gbase cwdm 1510
    ImAttrMedia_im_attr_media_40gbase_cwdm_1510 ImAttrMedia = "im-attr-media-40gbase-cwdm-1510"

    // im attr media 40gbase cwdm 1530
    ImAttrMedia_im_attr_media_40gbase_cwdm_1530 ImAttrMedia = "im-attr-media-40gbase-cwdm-1530"

    // im attr media 40gbase cwdm 1550
    ImAttrMedia_im_attr_media_40gbase_cwdm_1550 ImAttrMedia = "im-attr-media-40gbase-cwdm-1550"

    // im attr media 40gbase cwdm 1570
    ImAttrMedia_im_attr_media_40gbase_cwdm_1570 ImAttrMedia = "im-attr-media-40gbase-cwdm-1570"

    // im attr media 40gbase cwdm 1590
    ImAttrMedia_im_attr_media_40gbase_cwdm_1590 ImAttrMedia = "im-attr-media-40gbase-cwdm-1590"

    // im attr media 40gbase cwdm 1610
    ImAttrMedia_im_attr_media_40gbase_cwdm_1610 ImAttrMedia = "im-attr-media-40gbase-cwdm-1610"

    // im attr media 100gbase cwdm
    ImAttrMedia_im_attr_media_100gbase_cwdm ImAttrMedia = "im-attr-media-100gbase-cwdm"

    // im attr media 100gbase cwdm tunable
    ImAttrMedia_im_attr_media_100gbase_cwdm_tunable ImAttrMedia = "im-attr-media-100gbase-cwdm-tunable"

    // im attr media 100gbase cwdm 1470
    ImAttrMedia_im_attr_media_100gbase_cwdm_1470 ImAttrMedia = "im-attr-media-100gbase-cwdm-1470"

    // im attr media 100gbase cwdm 1490
    ImAttrMedia_im_attr_media_100gbase_cwdm_1490 ImAttrMedia = "im-attr-media-100gbase-cwdm-1490"

    // im attr media 100gbase cwdm 1510
    ImAttrMedia_im_attr_media_100gbase_cwdm_1510 ImAttrMedia = "im-attr-media-100gbase-cwdm-1510"

    // im attr media 100gbase cwdm 1530
    ImAttrMedia_im_attr_media_100gbase_cwdm_1530 ImAttrMedia = "im-attr-media-100gbase-cwdm-1530"

    // im attr media 100gbase cwdm 1550
    ImAttrMedia_im_attr_media_100gbase_cwdm_1550 ImAttrMedia = "im-attr-media-100gbase-cwdm-1550"

    // im attr media 100gbase cwdm 1570
    ImAttrMedia_im_attr_media_100gbase_cwdm_1570 ImAttrMedia = "im-attr-media-100gbase-cwdm-1570"

    // im attr media 100gbase cwdm 1590
    ImAttrMedia_im_attr_media_100gbase_cwdm_1590 ImAttrMedia = "im-attr-media-100gbase-cwdm-1590"

    // im attr media 100gbase cwdm 1610
    ImAttrMedia_im_attr_media_100gbase_cwdm_1610 ImAttrMedia = "im-attr-media-100gbase-cwdm-1610"

    // im attr media 40gbase elpb
    ImAttrMedia_im_attr_media_40gbase_elpb ImAttrMedia = "im-attr-media-40gbase-elpb"

    // im attr media 100gbase elpb
    ImAttrMedia_im_attr_media_100gbase_elpb ImAttrMedia = "im-attr-media-100gbase-elpb"

    // im attr media 100gbase lr10
    ImAttrMedia_im_attr_media_100gbase_lr10 ImAttrMedia = "im-attr-media-100gbase-lr10"

    // im attr media 40gbase
    ImAttrMedia_im_attr_media_40gbase ImAttrMedia = "im-attr-media-40gbase"

    // im attr media 100gbase kp4
    ImAttrMedia_im_attr_media_100gbase_kp4 ImAttrMedia = "im-attr-media-100gbase-kp4"

    // im attr media 100gbase kr4
    ImAttrMedia_im_attr_media_100gbase_kr4 ImAttrMedia = "im-attr-media-100gbase-kr4"

    // im attr media 10gbase lrm
    ImAttrMedia_im_attr_media_10gbase_lrm ImAttrMedia = "im-attr-media-10gbase-lrm"

    // im attr media 10gbase cx4
    ImAttrMedia_im_attr_media_10gbase_cx4 ImAttrMedia = "im-attr-media-10gbase-cx4"

    // im attr media 10gbase
    ImAttrMedia_im_attr_media_10gbase ImAttrMedia = "im-attr-media-10gbase"

    // im attr media 10gbase kx4
    ImAttrMedia_im_attr_media_10gbase_kx4 ImAttrMedia = "im-attr-media-10gbase-kx4"

    // im attr media 10gbase kr
    ImAttrMedia_im_attr_media_10gbase_kr ImAttrMedia = "im-attr-media-10gbase-kr"

    // im attr media 10gbase pr
    ImAttrMedia_im_attr_media_10gbase_pr ImAttrMedia = "im-attr-media-10gbase-pr"

    // im attr media 100base lx
    ImAttrMedia_im_attr_media_100base_lx ImAttrMedia = "im-attr-media-100base-lx"

    // im attr media 100base zx
    ImAttrMedia_im_attr_media_100base_zx ImAttrMedia = "im-attr-media-100base-zx"

    // im attr media 1000base bx d
    ImAttrMedia_im_attr_media_1000base_bx_d ImAttrMedia = "im-attr-media-1000base-bx-d"

    // im attr media 1000base bx u
    ImAttrMedia_im_attr_media_1000base_bx_u ImAttrMedia = "im-attr-media-1000base-bx-u"

    // im attr media 1000base bx20 d
    ImAttrMedia_im_attr_media_1000base_bx20_d ImAttrMedia = "im-attr-media-1000base-bx20-d"

    // im attr media 1000base bx20 u
    ImAttrMedia_im_attr_media_1000base_bx20_u ImAttrMedia = "im-attr-media-1000base-bx20-u"

    // im attr media 1000base bx40 d
    ImAttrMedia_im_attr_media_1000base_bx40_d ImAttrMedia = "im-attr-media-1000base-bx40-d"

    // im attr media 1000base bx40 da
    ImAttrMedia_im_attr_media_1000base_bx40_da ImAttrMedia = "im-attr-media-1000base-bx40-da"

    // im attr media 1000base bx40 u
    ImAttrMedia_im_attr_media_1000base_bx40_u ImAttrMedia = "im-attr-media-1000base-bx40-u"

    // im attr media 1000base bx80 d
    ImAttrMedia_im_attr_media_1000base_bx80_d ImAttrMedia = "im-attr-media-1000base-bx80-d"

    // im attr media 1000base bx80 u
    ImAttrMedia_im_attr_media_1000base_bx80_u ImAttrMedia = "im-attr-media-1000base-bx80-u"

    // im attr media 1000base bx120 d
    ImAttrMedia_im_attr_media_1000base_bx120_d ImAttrMedia = "im-attr-media-1000base-bx120-d"

    // im attr media 1000base bx120 u
    ImAttrMedia_im_attr_media_1000base_bx120_u ImAttrMedia = "im-attr-media-1000base-bx120-u"

    // im attr media 10gbase bx d
    ImAttrMedia_im_attr_media_10gbase_bx_d ImAttrMedia = "im-attr-media-10gbase-bx-d"

    // im attr media 10gbase bx u
    ImAttrMedia_im_attr_media_10gbase_bx_u ImAttrMedia = "im-attr-media-10gbase-bx-u"

    // im attr media 10gbase bx10 d
    ImAttrMedia_im_attr_media_10gbase_bx10_d ImAttrMedia = "im-attr-media-10gbase-bx10-d"

    // im attr media 10gbase bx10 u
    ImAttrMedia_im_attr_media_10gbase_bx10_u ImAttrMedia = "im-attr-media-10gbase-bx10-u"

    // im attr media 10gbase bx20 d
    ImAttrMedia_im_attr_media_10gbase_bx20_d ImAttrMedia = "im-attr-media-10gbase-bx20-d"

    // im attr media 10gbase bx20 u
    ImAttrMedia_im_attr_media_10gbase_bx20_u ImAttrMedia = "im-attr-media-10gbase-bx20-u"

    // im attr media 10gbase bx40 d
    ImAttrMedia_im_attr_media_10gbase_bx40_d ImAttrMedia = "im-attr-media-10gbase-bx40-d"

    // im attr media 10gbase bx40 u
    ImAttrMedia_im_attr_media_10gbase_bx40_u ImAttrMedia = "im-attr-media-10gbase-bx40-u"

    // im attr media 10gbase bx80 d
    ImAttrMedia_im_attr_media_10gbase_bx80_d ImAttrMedia = "im-attr-media-10gbase-bx80-d"

    // im attr media 10gbase bx80 u
    ImAttrMedia_im_attr_media_10gbase_bx80_u ImAttrMedia = "im-attr-media-10gbase-bx80-u"

    // im attr media 10gbase bx120 d
    ImAttrMedia_im_attr_media_10gbase_bx120_d ImAttrMedia = "im-attr-media-10gbase-bx120-d"

    // im attr media 10gbase bx120 u
    ImAttrMedia_im_attr_media_10gbase_bx120_u ImAttrMedia = "im-attr-media-10gbase-bx120-u"

    // im attr media 1000base dr lx
    ImAttrMedia_im_attr_media_1000base_dr_lx ImAttrMedia = "im-attr-media-1000base-dr-lx"

    // im attr media 100gbase er4l
    ImAttrMedia_im_attr_media_100gbase_er4l ImAttrMedia = "im-attr-media-100gbase-er4l"

    // im attr media 100gbase sr4
    ImAttrMedia_im_attr_media_100gbase_sr4 ImAttrMedia = "im-attr-media-100gbase-sr4"

    // im attr media 40gbase sr bd
    ImAttrMedia_im_attr_media_40gbase_sr_bd ImAttrMedia = "im-attr-media-40gbase-sr-bd"

    // im attr media 25gbase cr
    ImAttrMedia_im_attr_media_25gbase_cr ImAttrMedia = "im-attr-media-25gbase-cr"

    // im attr media 25gbase cr s
    ImAttrMedia_im_attr_media_25gbase_cr_s ImAttrMedia = "im-attr-media-25gbase-cr-s"

    // im attr media 25gbase kr
    ImAttrMedia_im_attr_media_25gbase_kr ImAttrMedia = "im-attr-media-25gbase-kr"

    // im attr media 25gbase kr s
    ImAttrMedia_im_attr_media_25gbase_kr_s ImAttrMedia = "im-attr-media-25gbase-kr-s"

    // im attr media 25gbase r
    ImAttrMedia_im_attr_media_25gbase_r ImAttrMedia = "im-attr-media-25gbase-r"

    // im attr media 25gbase sr
    ImAttrMedia_im_attr_media_25gbase_sr ImAttrMedia = "im-attr-media-25gbase-sr"

    // im attr media 25gbase dwdm
    ImAttrMedia_im_attr_media_25gbase_dwdm ImAttrMedia = "im-attr-media-25gbase-dwdm"

    // im attr media 25gbase dwdm tunable
    ImAttrMedia_im_attr_media_25gbase_dwdm_tunable ImAttrMedia = "im-attr-media-25gbase-dwdm-tunable"

    // im attr media 25gbase cwdm
    ImAttrMedia_im_attr_media_25gbase_cwdm ImAttrMedia = "im-attr-media-25gbase-cwdm"

    // im attr media 25gbase cwdm tunable
    ImAttrMedia_im_attr_media_25gbase_cwdm_tunable ImAttrMedia = "im-attr-media-25gbase-cwdm-tunable"

    // im attr media 100gbase psm4
    ImAttrMedia_im_attr_media_100gbase_psm4 ImAttrMedia = "im-attr-media-100gbase-psm4"

    // im attr media 100gbase er10
    ImAttrMedia_im_attr_media_100gbase_er10 ImAttrMedia = "im-attr-media-100gbase-er10"

    // im attr media 100gbase er10l
    ImAttrMedia_im_attr_media_100gbase_er10l ImAttrMedia = "im-attr-media-100gbase-er10l"

    // im attr media 100gbase acc
    ImAttrMedia_im_attr_media_100gbase_acc ImAttrMedia = "im-attr-media-100gbase-acc"

    // im attr media 100gbase aoc
    ImAttrMedia_im_attr_media_100gbase_aoc ImAttrMedia = "im-attr-media-100gbase-aoc"

    // im attr media 100gbase cwdm4
    ImAttrMedia_im_attr_media_100gbase_cwdm4 ImAttrMedia = "im-attr-media-100gbase-cwdm4"

    // im attr media 40gbase psm4
    ImAttrMedia_im_attr_media_40gbase_psm4 ImAttrMedia = "im-attr-media-40gbase-psm4"

    // im attr media 100gbase cr4
    ImAttrMedia_im_attr_media_100gbase_cr4 ImAttrMedia = "im-attr-media-100gbase-cr4"

    // im attr media 100gbase act loop
    ImAttrMedia_im_attr_media_100gbase_act_loop ImAttrMedia = "im-attr-media-100gbase-act-loop"

    // im attr media 100gbase pas loop
    ImAttrMedia_im_attr_media_100gbase_pas_loop ImAttrMedia = "im-attr-media-100gbase-pas-loop"

    // im attr media 50gbase cr2
    ImAttrMedia_im_attr_media_50gbase_cr2 ImAttrMedia = "im-attr-media-50gbase-cr2"

    // im attr media 50gbase sr2
    ImAttrMedia_im_attr_media_50gbase_sr2 ImAttrMedia = "im-attr-media-50gbase-sr2"

    // im attr media 50gbase psm2
    ImAttrMedia_im_attr_media_50gbase_psm2 ImAttrMedia = "im-attr-media-50gbase-psm2"

    // im attr media 200gbase cr4
    ImAttrMedia_im_attr_media_200gbase_cr4 ImAttrMedia = "im-attr-media-200gbase-cr4"

    // im attr media 400gbase fr4
    ImAttrMedia_im_attr_media_400gbase_fr4 ImAttrMedia = "im-attr-media-400gbase-fr4"

    // im attr media 400gbase dr4
    ImAttrMedia_im_attr_media_400gbase_dr4 ImAttrMedia = "im-attr-media-400gbase-dr4"

    // im attr media 400gbase cr4
    ImAttrMedia_im_attr_media_400gbase_cr4 ImAttrMedia = "im-attr-media-400gbase-cr4"

    // im attr media 10gbase cr
    ImAttrMedia_im_attr_media_10gbase_cr ImAttrMedia = "im-attr-media-10gbase-cr"

    // im attr media 10gbase aoc
    ImAttrMedia_im_attr_media_10gbase_aoc ImAttrMedia = "im-attr-media-10gbase-aoc"

    // im attr media 40gbase aoc
    ImAttrMedia_im_attr_media_40gbase_aoc ImAttrMedia = "im-attr-media-40gbase-aoc"

    // im attr media 40gbase acu
    ImAttrMedia_im_attr_media_40gbase_acu ImAttrMedia = "im-attr-media-40gbase-acu"

    // im attr media 100gbase acu
    ImAttrMedia_im_attr_media_100gbase_acu ImAttrMedia = "im-attr-media-100gbase-acu"

    // im attr media 10gbase cu1m
    ImAttrMedia_im_attr_media_10gbase_cu1m ImAttrMedia = "im-attr-media-10gbase-cu1m"

    // im attr media 10gbase cu3m
    ImAttrMedia_im_attr_media_10gbase_cu3m ImAttrMedia = "im-attr-media-10gbase-cu3m"

    // im attr media 10gbase cu5m
    ImAttrMedia_im_attr_media_10gbase_cu5m ImAttrMedia = "im-attr-media-10gbase-cu5m"

    // im attr media 10gbase acu7m
    ImAttrMedia_im_attr_media_10gbase_acu7m ImAttrMedia = "im-attr-media-10gbase-acu7m"

    // im attr media 10gbase acu10m
    ImAttrMedia_im_attr_media_10gbase_acu10m ImAttrMedia = "im-attr-media-10gbase-acu10m"

    // im attr media 4x10g base lr
    ImAttrMedia_im_attr_media_4x10g_base_lr ImAttrMedia = "im-attr-media-4x10g-base-lr"

    // im attr media 10gbase cu1 5m
    ImAttrMedia_im_attr_media_10gbase_cu1_5m ImAttrMedia = "im-attr-media-10gbase-cu1-5m"

    // im attr media 10gbase cu2m
    ImAttrMedia_im_attr_media_10gbase_cu2m ImAttrMedia = "im-attr-media-10gbase-cu2m"

    // im attr media 10gbase cu2 5m
    ImAttrMedia_im_attr_media_10gbase_cu2_5m ImAttrMedia = "im-attr-media-10gbase-cu2-5m"

    // im attr media 10gbase aoc1m
    ImAttrMedia_im_attr_media_10gbase_aoc1m ImAttrMedia = "im-attr-media-10gbase-aoc1m"

    // im attr media 10gbase aoc2m
    ImAttrMedia_im_attr_media_10gbase_aoc2m ImAttrMedia = "im-attr-media-10gbase-aoc2m"

    // im attr media 10gbase aoc3m
    ImAttrMedia_im_attr_media_10gbase_aoc3m ImAttrMedia = "im-attr-media-10gbase-aoc3m"

    // im attr media 10gbase aoc5m
    ImAttrMedia_im_attr_media_10gbase_aoc5m ImAttrMedia = "im-attr-media-10gbase-aoc5m"

    // im attr media 10gbase aoc7m
    ImAttrMedia_im_attr_media_10gbase_aoc7m ImAttrMedia = "im-attr-media-10gbase-aoc7m"

    // im attr media 10gbase aoc10m
    ImAttrMedia_im_attr_media_10gbase_aoc10m ImAttrMedia = "im-attr-media-10gbase-aoc10m"

    // im attr media 40gbase acu1m
    ImAttrMedia_im_attr_media_40gbase_acu1m ImAttrMedia = "im-attr-media-40gbase-acu1m"

    // im attr media 40gbase acu3m
    ImAttrMedia_im_attr_media_40gbase_acu3m ImAttrMedia = "im-attr-media-40gbase-acu3m"

    // im attr media 40gbase acu5m
    ImAttrMedia_im_attr_media_40gbase_acu5m ImAttrMedia = "im-attr-media-40gbase-acu5m"

    // im attr media 40gbase acu7m
    ImAttrMedia_im_attr_media_40gbase_acu7m ImAttrMedia = "im-attr-media-40gbase-acu7m"

    // im attr media 40gbase acu10m
    ImAttrMedia_im_attr_media_40gbase_acu10m ImAttrMedia = "im-attr-media-40gbase-acu10m"

    // im attr media 25gbase cu1m
    ImAttrMedia_im_attr_media_25gbase_cu1m ImAttrMedia = "im-attr-media-25gbase-cu1m"

    // im attr media 25gbase cu2m
    ImAttrMedia_im_attr_media_25gbase_cu2m ImAttrMedia = "im-attr-media-25gbase-cu2m"

    // im attr media 25gbase cu3m
    ImAttrMedia_im_attr_media_25gbase_cu3m ImAttrMedia = "im-attr-media-25gbase-cu3m"

    // im attr media 25gbase cu5m
    ImAttrMedia_im_attr_media_25gbase_cu5m ImAttrMedia = "im-attr-media-25gbase-cu5m"

    // im attr media 100gbase sm sr
    ImAttrMedia_im_attr_media_100gbase_sm_sr ImAttrMedia = "im-attr-media-100gbase-sm-sr"
)

// SrpMgmtSrrNodeState represents SRP SRR node state
type SrpMgmtSrrNodeState string

const (
    // Idle
    SrpMgmtSrrNodeState_idle_srr_state SrpMgmtSrrNodeState = "idle-srr-state"

    // Discovery
    SrpMgmtSrrNodeState_discovery_srr_state SrpMgmtSrrNodeState = "discovery-srr-state"

    // UNKNOWN
    SrpMgmtSrrNodeState_unknown_srr_state SrpMgmtSrrNodeState = "unknown-srr-state"
)

// Interfaces
// Interface operational data
type Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed operational data for interfaces and configured features.
    InterfaceXr Interfaces_InterfaceXr

    // Node and/or interface type specific view of interface summary data.
    NodeTypeSets Interfaces_NodeTypeSets

    // Brief operational data for interfaces.
    InterfaceBriefs Interfaces_InterfaceBriefs

    // Inventory summary information.
    InventorySummary Interfaces_InventorySummary

    // Descriptions for interfaces.
    Interfaces Interfaces_Interfaces

    // Interface summary information.
    InterfaceSummary Interfaces_InterfaceSummary
}

func (interfaces *Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "Cisco-IOS-XR-pfi-im-cmd-oper"
    interfaces.EntityData.SegmentPath = "Cisco-IOS-XR-pfi-im-cmd-oper:interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface-xr", types.YChild{"InterfaceXr", &interfaces.InterfaceXr})
    interfaces.EntityData.Children.Append("node-type-sets", types.YChild{"NodeTypeSets", &interfaces.NodeTypeSets})
    interfaces.EntityData.Children.Append("interface-briefs", types.YChild{"InterfaceBriefs", &interfaces.InterfaceBriefs})
    interfaces.EntityData.Children.Append("inventory-summary", types.YChild{"InventorySummary", &interfaces.InventorySummary})
    interfaces.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaces.Interfaces})
    interfaces.EntityData.Children.Append("interface-summary", types.YChild{"InterfaceSummary", &interfaces.InterfaceSummary})
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Interfaces_InterfaceXr
// Detailed operational data for interfaces and
// configured features
type Interfaces_InterfaceXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed operational data for a particular interface. The type is slice of
    // Interfaces_InterfaceXr_Interface.
    Interface []*Interfaces_InterfaceXr_Interface
}

func (interfaceXr *Interfaces_InterfaceXr) GetEntityData() *types.CommonEntityData {
    interfaceXr.EntityData.YFilter = interfaceXr.YFilter
    interfaceXr.EntityData.YangName = "interface-xr"
    interfaceXr.EntityData.BundleName = "cisco_ios_xr"
    interfaceXr.EntityData.ParentYangName = "interfaces"
    interfaceXr.EntityData.SegmentPath = "interface-xr"
    interfaceXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceXr.EntityData.Children = types.NewOrderedMap()
    interfaceXr.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceXr.Interface {
        interfaceXr.EntityData.Children.Append(types.GetSegmentPath(interfaceXr.Interface[i]), types.YChild{"Interface", interfaceXr.Interface[i]})
    }
    interfaceXr.EntityData.Leafs = types.NewOrderedMap()

    interfaceXr.EntityData.YListKeys = []string {}

    return &(interfaceXr.EntityData)
}

// Interfaces_InterfaceXr_Interface
// Detailed operational data for a particular
// interface
type Interfaces_InterfaceXr_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceHandle interface{}

    // Interface type. The type is string.
    InterfaceType interface{}

    // Hardware type description string. The type is string with length: 0..64.
    HardwareTypeString interface{}

    // Interface state. The type is ImStateEnum.
    State interface{}

    // Line protocol state. The type is ImStateEnum.
    LineState interface{}

    // Interface encapsulation. The type is string.
    Encapsulation interface{}

    // Interface encapsulation description string. The type is string with length:
    // 0..32.
    EncapsulationTypeString interface{}

    // MTU in bytes. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    Mtu interface{}

    // L2 transport flag. The type is bool.
    IsL2TransportEnabled interface{}

    // The number of times the state has changed. The type is interface{} with
    // range: 0..4294967295.
    StateTransitionCount interface{}

    // The time elasped after the last state transition. The type is interface{}
    // with range: 0..4294967295.
    LastStateTransitionTime interface{}

    // Dampening enabled flag. The type is bool.
    IsDampeningEnabled interface{}

    // Interface speed (Kb/s). The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Cyclic Redundancy Check length. The type is interface{} with range:
    // 0..4294967295.
    CrcLength interface{}

    // Interface scramble config. The type is bool.
    IsScrambleEnabled interface{}

    // Interface duplexity. The type is ImAttrDuplex.
    Duplexity interface{}

    // Interface media type. The type is ImAttrMedia.
    MediaType interface{}

    // Interface link type. The type is ImAttrLink.
    LinkType interface{}

    // Input flow control configuration. The type is ImAttrFlowControl.
    InFlowControl interface{}

    // Output flow control configuration. The type is ImAttrFlowControl.
    OutFlowControl interface{}

    // Interface bandwidth (Kb/s). The type is interface{} with range:
    // 0..4294967295.
    Bandwidth interface{}

    // Maximum Interface bandwidth (Kb/s). The type is interface{} with range:
    // 0..4294967295.
    MaxBandwidth interface{}

    // Interface keepalive time (s). The type is interface{} with range:
    // 0..4294967295.
    Keepalive interface{}

    // Loopback detected by layer 2. The type is bool.
    IsL2Looped interface{}

    // Parent interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    ParentInterfaceName interface{}

    // Interface loopback configuration. The type is ImCmdLoopbackEnum.
    LoopbackConfiguration interface{}

    // Interface description string. The type is string.
    Description interface{}

    // Maintenance embargo flag. The type is bool.
    IsMaintenanceEnabled interface{}

    // Data invert flag. The type is bool.
    IsDataInverted interface{}

    // Interface transport mode. The type is ImAttrTransportMode.
    TransportMode interface{}

    // Fast Shutdown flag. The type is bool.
    FastShutdown interface{}

    // This is not supposed to be used. It is a dummy attribute to support ifindex
    // for OC model. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // State dampening information.
    DampeningInformation Interfaces_InterfaceXr_Interface_DampeningInformation

    // Interface MAC address.
    MacAddress Interfaces_InterfaceXr_Interface_MacAddress

    // Interface burned in address.
    BurnedInAddress Interfaces_InterfaceXr_Interface_BurnedInAddress

    // Carrier Delay.
    CarrierDelay Interfaces_InterfaceXr_Interface_CarrierDelay

    // Interface ARP type and timeout.
    ArpInformation Interfaces_InterfaceXr_Interface_ArpInformation

    // Interface IP address info.
    IpInformation Interfaces_InterfaceXr_Interface_IpInformation

    // Information specific to the encapsulation.
    EncapsulationInformation Interfaces_InterfaceXr_Interface_EncapsulationInformation

    // Information specific to the interface type.
    InterfaceTypeInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation

    // Packet and byte rates.
    DataRates Interfaces_InterfaceXr_Interface_DataRates

    // Packet, byte and error counters.
    InterfaceStatistics Interfaces_InterfaceXr_Interface_InterfaceStatistics

    // L2 Protocol Statistics.
    L2InterfaceStatistics Interfaces_InterfaceXr_Interface_L2InterfaceStatistics
}

func (self *Interfaces_InterfaceXr_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-xr"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("dampening-information", types.YChild{"DampeningInformation", &self.DampeningInformation})
    self.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &self.MacAddress})
    self.EntityData.Children.Append("burned-in-address", types.YChild{"BurnedInAddress", &self.BurnedInAddress})
    self.EntityData.Children.Append("carrier-delay", types.YChild{"CarrierDelay", &self.CarrierDelay})
    self.EntityData.Children.Append("arp-information", types.YChild{"ArpInformation", &self.ArpInformation})
    self.EntityData.Children.Append("ip-information", types.YChild{"IpInformation", &self.IpInformation})
    self.EntityData.Children.Append("encapsulation-information", types.YChild{"EncapsulationInformation", &self.EncapsulationInformation})
    self.EntityData.Children.Append("interface-type-information", types.YChild{"InterfaceTypeInformation", &self.InterfaceTypeInformation})
    self.EntityData.Children.Append("data-rates", types.YChild{"DataRates", &self.DataRates})
    self.EntityData.Children.Append("interface-statistics", types.YChild{"InterfaceStatistics", &self.InterfaceStatistics})
    self.EntityData.Children.Append("l2-interface-statistics", types.YChild{"L2InterfaceStatistics", &self.L2InterfaceStatistics})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", self.InterfaceHandle})
    self.EntityData.Leafs.Append("interface-type", types.YLeaf{"InterfaceType", self.InterfaceType})
    self.EntityData.Leafs.Append("hardware-type-string", types.YLeaf{"HardwareTypeString", self.HardwareTypeString})
    self.EntityData.Leafs.Append("state", types.YLeaf{"State", self.State})
    self.EntityData.Leafs.Append("line-state", types.YLeaf{"LineState", self.LineState})
    self.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", self.Encapsulation})
    self.EntityData.Leafs.Append("encapsulation-type-string", types.YLeaf{"EncapsulationTypeString", self.EncapsulationTypeString})
    self.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", self.Mtu})
    self.EntityData.Leafs.Append("is-l2-transport-enabled", types.YLeaf{"IsL2TransportEnabled", self.IsL2TransportEnabled})
    self.EntityData.Leafs.Append("state-transition-count", types.YLeaf{"StateTransitionCount", self.StateTransitionCount})
    self.EntityData.Leafs.Append("last-state-transition-time", types.YLeaf{"LastStateTransitionTime", self.LastStateTransitionTime})
    self.EntityData.Leafs.Append("is-dampening-enabled", types.YLeaf{"IsDampeningEnabled", self.IsDampeningEnabled})
    self.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", self.Speed})
    self.EntityData.Leafs.Append("crc-length", types.YLeaf{"CrcLength", self.CrcLength})
    self.EntityData.Leafs.Append("is-scramble-enabled", types.YLeaf{"IsScrambleEnabled", self.IsScrambleEnabled})
    self.EntityData.Leafs.Append("duplexity", types.YLeaf{"Duplexity", self.Duplexity})
    self.EntityData.Leafs.Append("media-type", types.YLeaf{"MediaType", self.MediaType})
    self.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", self.LinkType})
    self.EntityData.Leafs.Append("in-flow-control", types.YLeaf{"InFlowControl", self.InFlowControl})
    self.EntityData.Leafs.Append("out-flow-control", types.YLeaf{"OutFlowControl", self.OutFlowControl})
    self.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", self.Bandwidth})
    self.EntityData.Leafs.Append("max-bandwidth", types.YLeaf{"MaxBandwidth", self.MaxBandwidth})
    self.EntityData.Leafs.Append("keepalive", types.YLeaf{"Keepalive", self.Keepalive})
    self.EntityData.Leafs.Append("is-l2-looped", types.YLeaf{"IsL2Looped", self.IsL2Looped})
    self.EntityData.Leafs.Append("parent-interface-name", types.YLeaf{"ParentInterfaceName", self.ParentInterfaceName})
    self.EntityData.Leafs.Append("loopback-configuration", types.YLeaf{"LoopbackConfiguration", self.LoopbackConfiguration})
    self.EntityData.Leafs.Append("description", types.YLeaf{"Description", self.Description})
    self.EntityData.Leafs.Append("is-maintenance-enabled", types.YLeaf{"IsMaintenanceEnabled", self.IsMaintenanceEnabled})
    self.EntityData.Leafs.Append("is-data-inverted", types.YLeaf{"IsDataInverted", self.IsDataInverted})
    self.EntityData.Leafs.Append("transport-mode", types.YLeaf{"TransportMode", self.TransportMode})
    self.EntityData.Leafs.Append("fast-shutdown", types.YLeaf{"FastShutdown", self.FastShutdown})
    self.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", self.IfIndex})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Interfaces_InterfaceXr_Interface_DampeningInformation
// State dampening information
type Interfaces_InterfaceXr_Interface_DampeningInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening penalty of the interface. The type is interface{} with range:
    // 0..4294967295.
    Penalty interface{}

    // Flag showing if state is suppressed. The type is bool.
    IsSuppressedEnabled interface{}

    // Remaining period of suppression in secs. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsRemaining interface{}

    // Configured decay half life in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    HalfLife interface{}

    // Configured reuse threshold. The type is interface{} with range:
    // 0..4294967295.
    ReuseThreshold interface{}

    // Value of suppress threshold. The type is interface{} with range:
    // 0..4294967295.
    SuppressThreshold interface{}

    // Maximum suppress time in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    MaximumSuppressTime interface{}

    // Configured restart penalty. The type is interface{} with range:
    // 0..4294967295.
    RestartPenalty interface{}
}

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetEntityData() *types.CommonEntityData {
    dampeningInformation.EntityData.YFilter = dampeningInformation.YFilter
    dampeningInformation.EntityData.YangName = "dampening-information"
    dampeningInformation.EntityData.BundleName = "cisco_ios_xr"
    dampeningInformation.EntityData.ParentYangName = "interface"
    dampeningInformation.EntityData.SegmentPath = "dampening-information"
    dampeningInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dampeningInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dampeningInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dampeningInformation.EntityData.Children = types.NewOrderedMap()
    dampeningInformation.EntityData.Leafs = types.NewOrderedMap()
    dampeningInformation.EntityData.Leafs.Append("penalty", types.YLeaf{"Penalty", dampeningInformation.Penalty})
    dampeningInformation.EntityData.Leafs.Append("is-suppressed-enabled", types.YLeaf{"IsSuppressedEnabled", dampeningInformation.IsSuppressedEnabled})
    dampeningInformation.EntityData.Leafs.Append("seconds-remaining", types.YLeaf{"SecondsRemaining", dampeningInformation.SecondsRemaining})
    dampeningInformation.EntityData.Leafs.Append("half-life", types.YLeaf{"HalfLife", dampeningInformation.HalfLife})
    dampeningInformation.EntityData.Leafs.Append("reuse-threshold", types.YLeaf{"ReuseThreshold", dampeningInformation.ReuseThreshold})
    dampeningInformation.EntityData.Leafs.Append("suppress-threshold", types.YLeaf{"SuppressThreshold", dampeningInformation.SuppressThreshold})
    dampeningInformation.EntityData.Leafs.Append("maximum-suppress-time", types.YLeaf{"MaximumSuppressTime", dampeningInformation.MaximumSuppressTime})
    dampeningInformation.EntityData.Leafs.Append("restart-penalty", types.YLeaf{"RestartPenalty", dampeningInformation.RestartPenalty})

    dampeningInformation.EntityData.YListKeys = []string {}

    return &(dampeningInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_MacAddress
// Interface MAC address
type Interfaces_InterfaceXr_Interface_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "interface"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", macAddress.Address})

    macAddress.EntityData.YListKeys = []string {}

    return &(macAddress.EntityData)
}

// Interfaces_InterfaceXr_Interface_BurnedInAddress
// Interface burned in address
type Interfaces_InterfaceXr_Interface_BurnedInAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetEntityData() *types.CommonEntityData {
    burnedInAddress.EntityData.YFilter = burnedInAddress.YFilter
    burnedInAddress.EntityData.YangName = "burned-in-address"
    burnedInAddress.EntityData.BundleName = "cisco_ios_xr"
    burnedInAddress.EntityData.ParentYangName = "interface"
    burnedInAddress.EntityData.SegmentPath = "burned-in-address"
    burnedInAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    burnedInAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    burnedInAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    burnedInAddress.EntityData.Children = types.NewOrderedMap()
    burnedInAddress.EntityData.Leafs = types.NewOrderedMap()
    burnedInAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", burnedInAddress.Address})

    burnedInAddress.EntityData.YListKeys = []string {}

    return &(burnedInAddress.EntityData)
}

// Interfaces_InterfaceXr_Interface_CarrierDelay
// Carrier Delay
type Interfaces_InterfaceXr_Interface_CarrierDelay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Carrier delay on state up (ms). The type is interface{} with range:
    // 0..4294967295.
    CarrierDelayUp interface{}

    // Carrier delay on state down (ms). The type is interface{} with range:
    // 0..4294967295.
    CarrierDelayDown interface{}
}

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetEntityData() *types.CommonEntityData {
    carrierDelay.EntityData.YFilter = carrierDelay.YFilter
    carrierDelay.EntityData.YangName = "carrier-delay"
    carrierDelay.EntityData.BundleName = "cisco_ios_xr"
    carrierDelay.EntityData.ParentYangName = "interface"
    carrierDelay.EntityData.SegmentPath = "carrier-delay"
    carrierDelay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    carrierDelay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    carrierDelay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    carrierDelay.EntityData.Children = types.NewOrderedMap()
    carrierDelay.EntityData.Leafs = types.NewOrderedMap()
    carrierDelay.EntityData.Leafs.Append("carrier-delay-up", types.YLeaf{"CarrierDelayUp", carrierDelay.CarrierDelayUp})
    carrierDelay.EntityData.Leafs.Append("carrier-delay-down", types.YLeaf{"CarrierDelayDown", carrierDelay.CarrierDelayDown})

    carrierDelay.EntityData.YListKeys = []string {}

    return &(carrierDelay.EntityData)
}

// Interfaces_InterfaceXr_Interface_ArpInformation
// Interface ARP type and timeout
type Interfaces_InterfaceXr_Interface_ArpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ARP timeout in seconds. Only valid if 'ARPIsLearningDisabled' is 'false'.
    // The type is interface{} with range: 0..4294967295. Units are second.
    ArpTimeout interface{}

    // ARP type name. The type is string.
    ArpTypeName interface{}

    // Whether the interface has dynamic learning disabled. The type is bool.
    ArpIsLearningDisabled interface{}
}

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetEntityData() *types.CommonEntityData {
    arpInformation.EntityData.YFilter = arpInformation.YFilter
    arpInformation.EntityData.YangName = "arp-information"
    arpInformation.EntityData.BundleName = "cisco_ios_xr"
    arpInformation.EntityData.ParentYangName = "interface"
    arpInformation.EntityData.SegmentPath = "arp-information"
    arpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arpInformation.EntityData.Children = types.NewOrderedMap()
    arpInformation.EntityData.Leafs = types.NewOrderedMap()
    arpInformation.EntityData.Leafs.Append("arp-timeout", types.YLeaf{"ArpTimeout", arpInformation.ArpTimeout})
    arpInformation.EntityData.Leafs.Append("arp-type-name", types.YLeaf{"ArpTypeName", arpInformation.ArpTypeName})
    arpInformation.EntityData.Leafs.Append("arp-is-learning-disabled", types.YLeaf{"ArpIsLearningDisabled", arpInformation.ArpIsLearningDisabled})

    arpInformation.EntityData.YListKeys = []string {}

    return &(arpInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_IpInformation
// Interface IP address info
type Interfaces_InterfaceXr_Interface_IpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Interface subnet mask length. The type is interface{} with range:
    // 0..4294967295.
    SubnetMaskLength interface{}
}

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetEntityData() *types.CommonEntityData {
    ipInformation.EntityData.YFilter = ipInformation.YFilter
    ipInformation.EntityData.YangName = "ip-information"
    ipInformation.EntityData.BundleName = "cisco_ios_xr"
    ipInformation.EntityData.ParentYangName = "interface"
    ipInformation.EntityData.SegmentPath = "ip-information"
    ipInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipInformation.EntityData.Children = types.NewOrderedMap()
    ipInformation.EntityData.Leafs = types.NewOrderedMap()
    ipInformation.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", ipInformation.IpAddress})
    ipInformation.EntityData.Leafs.Append("subnet-mask-length", types.YLeaf{"SubnetMaskLength", ipInformation.SubnetMaskLength})

    ipInformation.EntityData.YListKeys = []string {}

    return &(ipInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation
// Information specific to the encapsulation
type Interfaces_InterfaceXr_Interface_EncapsulationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EncapsulationType. The type is ImCmdEncapsEnum.
    EncapsulationType interface{}

    // Frame Relay information.
    FrameRelayInformation Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation

    // VLAN 802.1q information.
    Dot1qInformation Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation

    // PPP information.
    PppInformation Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation
}

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetEntityData() *types.CommonEntityData {
    encapsulationInformation.EntityData.YFilter = encapsulationInformation.YFilter
    encapsulationInformation.EntityData.YangName = "encapsulation-information"
    encapsulationInformation.EntityData.BundleName = "cisco_ios_xr"
    encapsulationInformation.EntityData.ParentYangName = "interface"
    encapsulationInformation.EntityData.SegmentPath = "encapsulation-information"
    encapsulationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encapsulationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encapsulationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encapsulationInformation.EntityData.Children = types.NewOrderedMap()
    encapsulationInformation.EntityData.Children.Append("frame-relay-information", types.YChild{"FrameRelayInformation", &encapsulationInformation.FrameRelayInformation})
    encapsulationInformation.EntityData.Children.Append("dot1q-information", types.YChild{"Dot1qInformation", &encapsulationInformation.Dot1qInformation})
    encapsulationInformation.EntityData.Children.Append("ppp-information", types.YChild{"PppInformation", &encapsulationInformation.PppInformation})
    encapsulationInformation.EntityData.Leafs = types.NewOrderedMap()
    encapsulationInformation.EntityData.Leafs.Append("encapsulation-type", types.YLeaf{"EncapsulationType", encapsulationInformation.EncapsulationType})

    encapsulationInformation.EntityData.YListKeys = []string {}

    return &(encapsulationInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation
// Frame Relay information
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Frame Relay encapsulation type. The type is ImCmdFrTypeEnum.
    FrEncapsulationType interface{}

    // The LMI type: Autosense, ANSI, CCITT or CISCO. The type is
    // ImCmdLmiTypeEnum.
    LmiType interface{}

    // LMI DLCI. The type is interface{} with range: 0..4294967295.
    Lmidlci interface{}

    // The NNI LMI interface type. The type is bool.
    IsNni interface{}

    // The DTE/DCE LMI interface type. The type is bool.
    IsDte interface{}

    // Flag indicating whether the LMI  DTE/DCE/NNI-DTE state is UP. The type is
    // bool.
    IsLmiUp interface{}

    // Flag indicating whether the LMI  NNI-DCE state is UP. The type is bool.
    IsLmiNniDceUp interface{}

    // The status of FR LMI for an interface. The type is bool.
    IsLmiEnabled interface{}

    // Number of enquiry messages received. The type is interface{} with range:
    // 0..4294967295.
    EnquiriesReceived interface{}

    // Number of enquiry messages sent. The type is interface{} with range:
    // 0..4294967295.
    EnquiriesSent interface{}

    // Number of status messages received. The type is interface{} with range:
    // 0..4294967295.
    StatusReceived interface{}

    // Number of status messages sent. The type is interface{} with range:
    // 0..4294967295.
    StatusSent interface{}

    // Number of update status messages received. The type is interface{} with
    // range: 0..4294967295.
    UpdateStatusReceived interface{}

    // Number of update status messages sent. The type is interface{} with range:
    // 0..4294967295.
    UpdateStatusSent interface{}
}

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetEntityData() *types.CommonEntityData {
    frameRelayInformation.EntityData.YFilter = frameRelayInformation.YFilter
    frameRelayInformation.EntityData.YangName = "frame-relay-information"
    frameRelayInformation.EntityData.BundleName = "cisco_ios_xr"
    frameRelayInformation.EntityData.ParentYangName = "encapsulation-information"
    frameRelayInformation.EntityData.SegmentPath = "frame-relay-information"
    frameRelayInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frameRelayInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frameRelayInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frameRelayInformation.EntityData.Children = types.NewOrderedMap()
    frameRelayInformation.EntityData.Leafs = types.NewOrderedMap()
    frameRelayInformation.EntityData.Leafs.Append("fr-encapsulation-type", types.YLeaf{"FrEncapsulationType", frameRelayInformation.FrEncapsulationType})
    frameRelayInformation.EntityData.Leafs.Append("lmi-type", types.YLeaf{"LmiType", frameRelayInformation.LmiType})
    frameRelayInformation.EntityData.Leafs.Append("lmidlci", types.YLeaf{"Lmidlci", frameRelayInformation.Lmidlci})
    frameRelayInformation.EntityData.Leafs.Append("is-nni", types.YLeaf{"IsNni", frameRelayInformation.IsNni})
    frameRelayInformation.EntityData.Leafs.Append("is-dte", types.YLeaf{"IsDte", frameRelayInformation.IsDte})
    frameRelayInformation.EntityData.Leafs.Append("is-lmi-up", types.YLeaf{"IsLmiUp", frameRelayInformation.IsLmiUp})
    frameRelayInformation.EntityData.Leafs.Append("is-lmi-nni-dce-up", types.YLeaf{"IsLmiNniDceUp", frameRelayInformation.IsLmiNniDceUp})
    frameRelayInformation.EntityData.Leafs.Append("is-lmi-enabled", types.YLeaf{"IsLmiEnabled", frameRelayInformation.IsLmiEnabled})
    frameRelayInformation.EntityData.Leafs.Append("enquiries-received", types.YLeaf{"EnquiriesReceived", frameRelayInformation.EnquiriesReceived})
    frameRelayInformation.EntityData.Leafs.Append("enquiries-sent", types.YLeaf{"EnquiriesSent", frameRelayInformation.EnquiriesSent})
    frameRelayInformation.EntityData.Leafs.Append("status-received", types.YLeaf{"StatusReceived", frameRelayInformation.StatusReceived})
    frameRelayInformation.EntityData.Leafs.Append("status-sent", types.YLeaf{"StatusSent", frameRelayInformation.StatusSent})
    frameRelayInformation.EntityData.Leafs.Append("update-status-received", types.YLeaf{"UpdateStatusReceived", frameRelayInformation.UpdateStatusReceived})
    frameRelayInformation.EntityData.Leafs.Append("update-status-sent", types.YLeaf{"UpdateStatusSent", frameRelayInformation.UpdateStatusSent})

    frameRelayInformation.EntityData.YListKeys = []string {}

    return &(frameRelayInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation
// VLAN 802.1q information
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation type and tag stack.
    EncapsulationDetails Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails

    // VLAN-Switched information.
    VlanSwitched Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched
}

func (dot1qInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation) GetEntityData() *types.CommonEntityData {
    dot1qInformation.EntityData.YFilter = dot1qInformation.YFilter
    dot1qInformation.EntityData.YangName = "dot1q-information"
    dot1qInformation.EntityData.BundleName = "cisco_ios_xr"
    dot1qInformation.EntityData.ParentYangName = "encapsulation-information"
    dot1qInformation.EntityData.SegmentPath = "dot1q-information"
    dot1qInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1qInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1qInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1qInformation.EntityData.Children = types.NewOrderedMap()
    dot1qInformation.EntityData.Children.Append("encapsulation-details", types.YChild{"EncapsulationDetails", &dot1qInformation.EncapsulationDetails})
    dot1qInformation.EntityData.Children.Append("vlan-switched", types.YChild{"VlanSwitched", &dot1qInformation.VlanSwitched})
    dot1qInformation.EntityData.Leafs = types.NewOrderedMap()

    dot1qInformation.EntityData.YListKeys = []string {}

    return &(dot1qInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails
// Encapsulation type and tag stack
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLANEncapsulation. The type is VlanEncaps.
    VlanEncapsulation interface{}

    // Tag value. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Native tag value. The type is interface{} with range: 0..65535.
    NativeTag interface{}

    // 802.1ad tag value. The type is interface{} with range: 0..65535.
    Dot1adTag interface{}

    // 802.1ad native tag value. The type is interface{} with range: 0..65535.
    Dot1adNativeTag interface{}

    // 802.1ad Outer tag value. The type is interface{} with range: 0..65535.
    Dot1adOuterTag interface{}

    // Stack value.
    Stack Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_Stack

    // Service Instance encapsulation.
    ServiceInstanceDetails Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails

    // 802.1ad 802.1Q stack value.
    Dot1adDot1qStack Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_Dot1adDot1qStack
}

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails) GetEntityData() *types.CommonEntityData {
    encapsulationDetails.EntityData.YFilter = encapsulationDetails.YFilter
    encapsulationDetails.EntityData.YangName = "encapsulation-details"
    encapsulationDetails.EntityData.BundleName = "cisco_ios_xr"
    encapsulationDetails.EntityData.ParentYangName = "dot1q-information"
    encapsulationDetails.EntityData.SegmentPath = "encapsulation-details"
    encapsulationDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encapsulationDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encapsulationDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encapsulationDetails.EntityData.Children = types.NewOrderedMap()
    encapsulationDetails.EntityData.Children.Append("stack", types.YChild{"Stack", &encapsulationDetails.Stack})
    encapsulationDetails.EntityData.Children.Append("service-instance-details", types.YChild{"ServiceInstanceDetails", &encapsulationDetails.ServiceInstanceDetails})
    encapsulationDetails.EntityData.Children.Append("dot1ad-dot1q-stack", types.YChild{"Dot1adDot1qStack", &encapsulationDetails.Dot1adDot1qStack})
    encapsulationDetails.EntityData.Leafs = types.NewOrderedMap()
    encapsulationDetails.EntityData.Leafs.Append("vlan-encapsulation", types.YLeaf{"VlanEncapsulation", encapsulationDetails.VlanEncapsulation})
    encapsulationDetails.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", encapsulationDetails.Tag})
    encapsulationDetails.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", encapsulationDetails.OuterTag})
    encapsulationDetails.EntityData.Leafs.Append("native-tag", types.YLeaf{"NativeTag", encapsulationDetails.NativeTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-tag", types.YLeaf{"Dot1adTag", encapsulationDetails.Dot1adTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-native-tag", types.YLeaf{"Dot1adNativeTag", encapsulationDetails.Dot1adNativeTag})
    encapsulationDetails.EntityData.Leafs.Append("dot1ad-outer-tag", types.YLeaf{"Dot1adOuterTag", encapsulationDetails.Dot1adOuterTag})

    encapsulationDetails.EntityData.YListKeys = []string {}

    return &(encapsulationDetails.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_Stack
// Stack value
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_Stack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_Stack) GetEntityData() *types.CommonEntityData {
    stack.EntityData.YFilter = stack.YFilter
    stack.EntityData.YangName = "stack"
    stack.EntityData.BundleName = "cisco_ios_xr"
    stack.EntityData.ParentYangName = "encapsulation-details"
    stack.EntityData.SegmentPath = "stack"
    stack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stack.EntityData.Children = types.NewOrderedMap()
    stack.EntityData.Leafs = types.NewOrderedMap()
    stack.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", stack.OuterTag})
    stack.EntityData.Leafs.Append("second-tag", types.YLeaf{"SecondTag", stack.SecondTag})

    stack.EntityData.YListKeys = []string {}

    return &(stack.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails
// Service Instance encapsulation
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload Ethertype to match. The type is EfpPayloadEtype.
    PayloadEthertype interface{}

    // Number of tags popped on ingress. The type is interface{} with range:
    // 0..65535.
    TagsPopped interface{}

    // Whether the packet must match the encapsulation exactly, with no further
    // inner tags. The type is interface{} with range: -2147483648..2147483647.
    IsExactMatch interface{}

    // Whether this represents the native VLAN on the port. The type is
    // interface{} with range: -2147483648..2147483647.
    IsNativeVlan interface{}

    // Whether the native VLAN is customer-tag preserving. The type is interface{}
    // with range: -2147483648..2147483647.
    IsNativePreserving interface{}

    // The source MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacMatch interface{}

    // The destination MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DestinationMacMatch interface{}

    // VLAN tags for locally-sourced traffic.
    LocalTrafficStack Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack

    // Tags to match on ingress packets. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch.
    TagsToMatch []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe.
    Pushe []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe
}

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails) GetEntityData() *types.CommonEntityData {
    serviceInstanceDetails.EntityData.YFilter = serviceInstanceDetails.YFilter
    serviceInstanceDetails.EntityData.YangName = "service-instance-details"
    serviceInstanceDetails.EntityData.BundleName = "cisco_ios_xr"
    serviceInstanceDetails.EntityData.ParentYangName = "encapsulation-details"
    serviceInstanceDetails.EntityData.SegmentPath = "service-instance-details"
    serviceInstanceDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceInstanceDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceInstanceDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceInstanceDetails.EntityData.Children = types.NewOrderedMap()
    serviceInstanceDetails.EntityData.Children.Append("local-traffic-stack", types.YChild{"LocalTrafficStack", &serviceInstanceDetails.LocalTrafficStack})
    serviceInstanceDetails.EntityData.Children.Append("tags-to-match", types.YChild{"TagsToMatch", nil})
    for i := range serviceInstanceDetails.TagsToMatch {
        serviceInstanceDetails.EntityData.Children.Append(types.GetSegmentPath(serviceInstanceDetails.TagsToMatch[i]), types.YChild{"TagsToMatch", serviceInstanceDetails.TagsToMatch[i]})
    }
    serviceInstanceDetails.EntityData.Children.Append("pushe", types.YChild{"Pushe", nil})
    for i := range serviceInstanceDetails.Pushe {
        serviceInstanceDetails.EntityData.Children.Append(types.GetSegmentPath(serviceInstanceDetails.Pushe[i]), types.YChild{"Pushe", serviceInstanceDetails.Pushe[i]})
    }
    serviceInstanceDetails.EntityData.Leafs = types.NewOrderedMap()
    serviceInstanceDetails.EntityData.Leafs.Append("payload-ethertype", types.YLeaf{"PayloadEthertype", serviceInstanceDetails.PayloadEthertype})
    serviceInstanceDetails.EntityData.Leafs.Append("tags-popped", types.YLeaf{"TagsPopped", serviceInstanceDetails.TagsPopped})
    serviceInstanceDetails.EntityData.Leafs.Append("is-exact-match", types.YLeaf{"IsExactMatch", serviceInstanceDetails.IsExactMatch})
    serviceInstanceDetails.EntityData.Leafs.Append("is-native-vlan", types.YLeaf{"IsNativeVlan", serviceInstanceDetails.IsNativeVlan})
    serviceInstanceDetails.EntityData.Leafs.Append("is-native-preserving", types.YLeaf{"IsNativePreserving", serviceInstanceDetails.IsNativePreserving})
    serviceInstanceDetails.EntityData.Leafs.Append("source-mac-match", types.YLeaf{"SourceMacMatch", serviceInstanceDetails.SourceMacMatch})
    serviceInstanceDetails.EntityData.Leafs.Append("destination-mac-match", types.YLeaf{"DestinationMacMatch", serviceInstanceDetails.DestinationMacMatch})

    serviceInstanceDetails.EntityData.YListKeys = []string {}

    return &(serviceInstanceDetails.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
}

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetEntityData() *types.CommonEntityData {
    localTrafficStack.EntityData.YFilter = localTrafficStack.YFilter
    localTrafficStack.EntityData.YangName = "local-traffic-stack"
    localTrafficStack.EntityData.BundleName = "cisco_ios_xr"
    localTrafficStack.EntityData.ParentYangName = "service-instance-details"
    localTrafficStack.EntityData.SegmentPath = "local-traffic-stack"
    localTrafficStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficStack.EntityData.Children = types.NewOrderedMap()
    localTrafficStack.EntityData.Children.Append("local-traffic-tag", types.YChild{"LocalTrafficTag", nil})
    for i := range localTrafficStack.LocalTrafficTag {
        localTrafficStack.EntityData.Children.Append(types.GetSegmentPath(localTrafficStack.LocalTrafficTag[i]), types.YChild{"LocalTrafficTag", localTrafficStack.LocalTrafficTag[i]})
    }
    localTrafficStack.EntityData.Leafs = types.NewOrderedMap()

    localTrafficStack.EntityData.YListKeys = []string {}

    return &(localTrafficStack.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
// VLAN tags for locally-sourced traffic
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetEntityData() *types.CommonEntityData {
    localTrafficTag.EntityData.YFilter = localTrafficTag.YFilter
    localTrafficTag.EntityData.YangName = "local-traffic-tag"
    localTrafficTag.EntityData.BundleName = "cisco_ios_xr"
    localTrafficTag.EntityData.ParentYangName = "local-traffic-stack"
    localTrafficTag.EntityData.SegmentPath = "local-traffic-tag"
    localTrafficTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficTag.EntityData.Children = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", localTrafficTag.Ethertype})
    localTrafficTag.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", localTrafficTag.VlanId})

    localTrafficTag.EntityData.YListKeys = []string {}

    return &(localTrafficTag.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch
// Tags to match on ingress packets
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag to match. The type is EfpTagEtype.
    Ethertype interface{}

    // Priority to match. The type is EfpTagPriority.
    Priority interface{}

    // VLAN Ids to match. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange.
    VlanRange []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
}

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetEntityData() *types.CommonEntityData {
    tagsToMatch.EntityData.YFilter = tagsToMatch.YFilter
    tagsToMatch.EntityData.YangName = "tags-to-match"
    tagsToMatch.EntityData.BundleName = "cisco_ios_xr"
    tagsToMatch.EntityData.ParentYangName = "service-instance-details"
    tagsToMatch.EntityData.SegmentPath = "tags-to-match"
    tagsToMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagsToMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagsToMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagsToMatch.EntityData.Children = types.NewOrderedMap()
    tagsToMatch.EntityData.Children.Append("vlan-range", types.YChild{"VlanRange", nil})
    for i := range tagsToMatch.VlanRange {
        tagsToMatch.EntityData.Children.Append(types.GetSegmentPath(tagsToMatch.VlanRange[i]), types.YChild{"VlanRange", tagsToMatch.VlanRange[i]})
    }
    tagsToMatch.EntityData.Leafs = types.NewOrderedMap()
    tagsToMatch.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", tagsToMatch.Ethertype})
    tagsToMatch.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", tagsToMatch.Priority})

    tagsToMatch.EntityData.YListKeys = []string {}

    return &(tagsToMatch.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
// VLAN Ids to match
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN ID Low. The type is interface{} with range: 0..65535.
    VlanIdLow interface{}

    // VLAN ID High. The type is interface{} with range: 0..65535.
    VlanIdHigh interface{}
}

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetEntityData() *types.CommonEntityData {
    vlanRange.EntityData.YFilter = vlanRange.YFilter
    vlanRange.EntityData.YangName = "vlan-range"
    vlanRange.EntityData.BundleName = "cisco_ios_xr"
    vlanRange.EntityData.ParentYangName = "tags-to-match"
    vlanRange.EntityData.SegmentPath = "vlan-range"
    vlanRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRange.EntityData.Children = types.NewOrderedMap()
    vlanRange.EntityData.Leafs = types.NewOrderedMap()
    vlanRange.EntityData.Leafs.Append("vlan-id-low", types.YLeaf{"VlanIdLow", vlanRange.VlanIdLow})
    vlanRange.EntityData.Leafs.Append("vlan-id-high", types.YLeaf{"VlanIdHigh", vlanRange.VlanIdHigh})

    vlanRange.EntityData.YListKeys = []string {}

    return &(vlanRange.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe
// VLAN tags pushed on egress
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetEntityData() *types.CommonEntityData {
    pushe.EntityData.YFilter = pushe.YFilter
    pushe.EntityData.YangName = "pushe"
    pushe.EntityData.BundleName = "cisco_ios_xr"
    pushe.EntityData.ParentYangName = "service-instance-details"
    pushe.EntityData.SegmentPath = "pushe"
    pushe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pushe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pushe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pushe.EntityData.Children = types.NewOrderedMap()
    pushe.EntityData.Leafs = types.NewOrderedMap()
    pushe.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", pushe.Ethertype})
    pushe.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", pushe.VlanId})

    pushe.EntityData.YListKeys = []string {}

    return &(pushe.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_Dot1adDot1qStack
// 802.1ad 802.1Q stack value
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_Dot1adDot1qStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (dot1adDot1qStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_EncapsulationDetails_Dot1adDot1qStack) GetEntityData() *types.CommonEntityData {
    dot1adDot1qStack.EntityData.YFilter = dot1adDot1qStack.YFilter
    dot1adDot1qStack.EntityData.YangName = "dot1ad-dot1q-stack"
    dot1adDot1qStack.EntityData.BundleName = "cisco_ios_xr"
    dot1adDot1qStack.EntityData.ParentYangName = "encapsulation-details"
    dot1adDot1qStack.EntityData.SegmentPath = "dot1ad-dot1q-stack"
    dot1adDot1qStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1adDot1qStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1adDot1qStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1adDot1qStack.EntityData.Children = types.NewOrderedMap()
    dot1adDot1qStack.EntityData.Leafs = types.NewOrderedMap()
    dot1adDot1qStack.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", dot1adDot1qStack.OuterTag})
    dot1adDot1qStack.EntityData.Leafs.Append("second-tag", types.YLeaf{"SecondTag", dot1adDot1qStack.SecondTag})

    dot1adDot1qStack.EntityData.YListKeys = []string {}

    return &(dot1adDot1qStack.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched
// VLAN-Switched information
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN-Switched mode. The type is VlanSwitchedMode.
    Mode interface{}

    // VLAN-Switched Access VLAN. The type is interface{} with range: 0..65535.
    AccessVlan interface{}

    // VLAN-Switched Trunk VLAN ranges.
    TrunkVlanRanges Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges
}

func (vlanSwitched *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched) GetEntityData() *types.CommonEntityData {
    vlanSwitched.EntityData.YFilter = vlanSwitched.YFilter
    vlanSwitched.EntityData.YangName = "vlan-switched"
    vlanSwitched.EntityData.BundleName = "cisco_ios_xr"
    vlanSwitched.EntityData.ParentYangName = "dot1q-information"
    vlanSwitched.EntityData.SegmentPath = "vlan-switched"
    vlanSwitched.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanSwitched.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanSwitched.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanSwitched.EntityData.Children = types.NewOrderedMap()
    vlanSwitched.EntityData.Children.Append("trunk-vlan-ranges", types.YChild{"TrunkVlanRanges", &vlanSwitched.TrunkVlanRanges})
    vlanSwitched.EntityData.Leafs = types.NewOrderedMap()
    vlanSwitched.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", vlanSwitched.Mode})
    vlanSwitched.EntityData.Leafs.Append("access-vlan", types.YLeaf{"AccessVlan", vlanSwitched.AccessVlan})

    vlanSwitched.EntityData.YListKeys = []string {}

    return &(vlanSwitched.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges
// VLAN-Switched Trunk VLAN ranges
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Payload Ethertype to match. The type is EfpPayloadEtype.
    PayloadEthertype interface{}

    // Number of tags popped on ingress. The type is interface{} with range:
    // 0..65535.
    TagsPopped interface{}

    // Whether the packet must match the encapsulation exactly, with no further
    // inner tags. The type is interface{} with range: -2147483648..2147483647.
    IsExactMatch interface{}

    // Whether this represents the native VLAN on the port. The type is
    // interface{} with range: -2147483648..2147483647.
    IsNativeVlan interface{}

    // Whether the native VLAN is customer-tag preserving. The type is interface{}
    // with range: -2147483648..2147483647.
    IsNativePreserving interface{}

    // The source MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacMatch interface{}

    // The destination MAC address to match on ingress. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DestinationMacMatch interface{}

    // VLAN tags for locally-sourced traffic.
    LocalTrafficStack Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack

    // Tags to match on ingress packets. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch.
    TagsToMatch []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_Pushe.
    Pushe []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_Pushe
}

func (trunkVlanRanges *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges) GetEntityData() *types.CommonEntityData {
    trunkVlanRanges.EntityData.YFilter = trunkVlanRanges.YFilter
    trunkVlanRanges.EntityData.YangName = "trunk-vlan-ranges"
    trunkVlanRanges.EntityData.BundleName = "cisco_ios_xr"
    trunkVlanRanges.EntityData.ParentYangName = "vlan-switched"
    trunkVlanRanges.EntityData.SegmentPath = "trunk-vlan-ranges"
    trunkVlanRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunkVlanRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunkVlanRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunkVlanRanges.EntityData.Children = types.NewOrderedMap()
    trunkVlanRanges.EntityData.Children.Append("local-traffic-stack", types.YChild{"LocalTrafficStack", &trunkVlanRanges.LocalTrafficStack})
    trunkVlanRanges.EntityData.Children.Append("tags-to-match", types.YChild{"TagsToMatch", nil})
    for i := range trunkVlanRanges.TagsToMatch {
        trunkVlanRanges.EntityData.Children.Append(types.GetSegmentPath(trunkVlanRanges.TagsToMatch[i]), types.YChild{"TagsToMatch", trunkVlanRanges.TagsToMatch[i]})
    }
    trunkVlanRanges.EntityData.Children.Append("pushe", types.YChild{"Pushe", nil})
    for i := range trunkVlanRanges.Pushe {
        trunkVlanRanges.EntityData.Children.Append(types.GetSegmentPath(trunkVlanRanges.Pushe[i]), types.YChild{"Pushe", trunkVlanRanges.Pushe[i]})
    }
    trunkVlanRanges.EntityData.Leafs = types.NewOrderedMap()
    trunkVlanRanges.EntityData.Leafs.Append("payload-ethertype", types.YLeaf{"PayloadEthertype", trunkVlanRanges.PayloadEthertype})
    trunkVlanRanges.EntityData.Leafs.Append("tags-popped", types.YLeaf{"TagsPopped", trunkVlanRanges.TagsPopped})
    trunkVlanRanges.EntityData.Leafs.Append("is-exact-match", types.YLeaf{"IsExactMatch", trunkVlanRanges.IsExactMatch})
    trunkVlanRanges.EntityData.Leafs.Append("is-native-vlan", types.YLeaf{"IsNativeVlan", trunkVlanRanges.IsNativeVlan})
    trunkVlanRanges.EntityData.Leafs.Append("is-native-preserving", types.YLeaf{"IsNativePreserving", trunkVlanRanges.IsNativePreserving})
    trunkVlanRanges.EntityData.Leafs.Append("source-mac-match", types.YLeaf{"SourceMacMatch", trunkVlanRanges.SourceMacMatch})
    trunkVlanRanges.EntityData.Leafs.Append("destination-mac-match", types.YLeaf{"DestinationMacMatch", trunkVlanRanges.DestinationMacMatch})

    trunkVlanRanges.EntityData.YListKeys = []string {}

    return &(trunkVlanRanges.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag
}

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack) GetEntityData() *types.CommonEntityData {
    localTrafficStack.EntityData.YFilter = localTrafficStack.YFilter
    localTrafficStack.EntityData.YangName = "local-traffic-stack"
    localTrafficStack.EntityData.BundleName = "cisco_ios_xr"
    localTrafficStack.EntityData.ParentYangName = "trunk-vlan-ranges"
    localTrafficStack.EntityData.SegmentPath = "local-traffic-stack"
    localTrafficStack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficStack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficStack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficStack.EntityData.Children = types.NewOrderedMap()
    localTrafficStack.EntityData.Children.Append("local-traffic-tag", types.YChild{"LocalTrafficTag", nil})
    for i := range localTrafficStack.LocalTrafficTag {
        localTrafficStack.EntityData.Children.Append(types.GetSegmentPath(localTrafficStack.LocalTrafficTag[i]), types.YChild{"LocalTrafficTag", localTrafficStack.LocalTrafficTag[i]})
    }
    localTrafficStack.EntityData.Leafs = types.NewOrderedMap()

    localTrafficStack.EntityData.YListKeys = []string {}

    return &(localTrafficStack.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag
// VLAN tags for locally-sourced traffic
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_LocalTrafficStack_LocalTrafficTag) GetEntityData() *types.CommonEntityData {
    localTrafficTag.EntityData.YFilter = localTrafficTag.YFilter
    localTrafficTag.EntityData.YangName = "local-traffic-tag"
    localTrafficTag.EntityData.BundleName = "cisco_ios_xr"
    localTrafficTag.EntityData.ParentYangName = "local-traffic-stack"
    localTrafficTag.EntityData.SegmentPath = "local-traffic-tag"
    localTrafficTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTrafficTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTrafficTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTrafficTag.EntityData.Children = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs = types.NewOrderedMap()
    localTrafficTag.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", localTrafficTag.Ethertype})
    localTrafficTag.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", localTrafficTag.VlanId})

    localTrafficTag.EntityData.YListKeys = []string {}

    return &(localTrafficTag.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch
// Tags to match on ingress packets
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag to match. The type is EfpTagEtype.
    Ethertype interface{}

    // Priority to match. The type is EfpTagPriority.
    Priority interface{}

    // VLAN Ids to match. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange.
    VlanRange []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange
}

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch) GetEntityData() *types.CommonEntityData {
    tagsToMatch.EntityData.YFilter = tagsToMatch.YFilter
    tagsToMatch.EntityData.YangName = "tags-to-match"
    tagsToMatch.EntityData.BundleName = "cisco_ios_xr"
    tagsToMatch.EntityData.ParentYangName = "trunk-vlan-ranges"
    tagsToMatch.EntityData.SegmentPath = "tags-to-match"
    tagsToMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tagsToMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tagsToMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tagsToMatch.EntityData.Children = types.NewOrderedMap()
    tagsToMatch.EntityData.Children.Append("vlan-range", types.YChild{"VlanRange", nil})
    for i := range tagsToMatch.VlanRange {
        tagsToMatch.EntityData.Children.Append(types.GetSegmentPath(tagsToMatch.VlanRange[i]), types.YChild{"VlanRange", tagsToMatch.VlanRange[i]})
    }
    tagsToMatch.EntityData.Leafs = types.NewOrderedMap()
    tagsToMatch.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", tagsToMatch.Ethertype})
    tagsToMatch.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", tagsToMatch.Priority})

    tagsToMatch.EntityData.YListKeys = []string {}

    return &(tagsToMatch.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange
// VLAN Ids to match
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN ID Low. The type is interface{} with range: 0..65535.
    VlanIdLow interface{}

    // VLAN ID High. The type is interface{} with range: 0..65535.
    VlanIdHigh interface{}
}

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_TagsToMatch_VlanRange) GetEntityData() *types.CommonEntityData {
    vlanRange.EntityData.YFilter = vlanRange.YFilter
    vlanRange.EntityData.YangName = "vlan-range"
    vlanRange.EntityData.BundleName = "cisco_ios_xr"
    vlanRange.EntityData.ParentYangName = "tags-to-match"
    vlanRange.EntityData.SegmentPath = "vlan-range"
    vlanRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRange.EntityData.Children = types.NewOrderedMap()
    vlanRange.EntityData.Leafs = types.NewOrderedMap()
    vlanRange.EntityData.Leafs.Append("vlan-id-low", types.YLeaf{"VlanIdLow", vlanRange.VlanIdLow})
    vlanRange.EntityData.Leafs.Append("vlan-id-high", types.YLeaf{"VlanIdHigh", vlanRange.VlanIdHigh})

    vlanRange.EntityData.YListKeys = []string {}

    return &(vlanRange.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_Pushe
// VLAN tags pushed on egress
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_Pushe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1qInformation_VlanSwitched_TrunkVlanRanges_Pushe) GetEntityData() *types.CommonEntityData {
    pushe.EntityData.YFilter = pushe.YFilter
    pushe.EntityData.YangName = "pushe"
    pushe.EntityData.BundleName = "cisco_ios_xr"
    pushe.EntityData.ParentYangName = "trunk-vlan-ranges"
    pushe.EntityData.SegmentPath = "pushe"
    pushe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pushe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pushe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pushe.EntityData.Children = types.NewOrderedMap()
    pushe.EntityData.Leafs = types.NewOrderedMap()
    pushe.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", pushe.Ethertype})
    pushe.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", pushe.VlanId})

    pushe.EntityData.YListKeys = []string {}

    return &(pushe.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation
// PPP information
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LCP State. The type is PppFsmState.
    LcpState interface{}

    // Loopback detected. The type is interface{} with range:
    // -2147483648..2147483647.
    IsLoopbackDetected interface{}

    // Keepalive value. The type is interface{} with range: 0..4294967295.
    KeepalivePeriod interface{}

    // MP Bundle Member. The type is interface{} with range:
    // -2147483648..2147483647.
    IsMpBundleMember interface{}

    // Is Multilink Open. The type is interface{} with range:
    // -2147483648..2147483647.
    IsMultilinkOpen interface{}

    // Array of per-NCP data. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray.
    NcpInfoArray []*Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray
}

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetEntityData() *types.CommonEntityData {
    pppInformation.EntityData.YFilter = pppInformation.YFilter
    pppInformation.EntityData.YangName = "ppp-information"
    pppInformation.EntityData.BundleName = "cisco_ios_xr"
    pppInformation.EntityData.ParentYangName = "encapsulation-information"
    pppInformation.EntityData.SegmentPath = "ppp-information"
    pppInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppInformation.EntityData.Children = types.NewOrderedMap()
    pppInformation.EntityData.Children.Append("ncp-info-array", types.YChild{"NcpInfoArray", nil})
    for i := range pppInformation.NcpInfoArray {
        pppInformation.EntityData.Children.Append(types.GetSegmentPath(pppInformation.NcpInfoArray[i]), types.YChild{"NcpInfoArray", pppInformation.NcpInfoArray[i]})
    }
    pppInformation.EntityData.Leafs = types.NewOrderedMap()
    pppInformation.EntityData.Leafs.Append("lcp-state", types.YLeaf{"LcpState", pppInformation.LcpState})
    pppInformation.EntityData.Leafs.Append("is-loopback-detected", types.YLeaf{"IsLoopbackDetected", pppInformation.IsLoopbackDetected})
    pppInformation.EntityData.Leafs.Append("keepalive-period", types.YLeaf{"KeepalivePeriod", pppInformation.KeepalivePeriod})
    pppInformation.EntityData.Leafs.Append("is-mp-bundle-member", types.YLeaf{"IsMpBundleMember", pppInformation.IsMpBundleMember})
    pppInformation.EntityData.Leafs.Append("is-multilink-open", types.YLeaf{"IsMultilinkOpen", pppInformation.IsMultilinkOpen})

    pppInformation.EntityData.YListKeys = []string {}

    return &(pppInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray
// Array of per-NCP data
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NCP state value. The type is PppFsmState.
    NcpState interface{}

    // NCP state identifier. The type is NcpIdent.
    NcpIdentifier interface{}
}

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetEntityData() *types.CommonEntityData {
    ncpInfoArray.EntityData.YFilter = ncpInfoArray.YFilter
    ncpInfoArray.EntityData.YangName = "ncp-info-array"
    ncpInfoArray.EntityData.BundleName = "cisco_ios_xr"
    ncpInfoArray.EntityData.ParentYangName = "ppp-information"
    ncpInfoArray.EntityData.SegmentPath = "ncp-info-array"
    ncpInfoArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncpInfoArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncpInfoArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncpInfoArray.EntityData.Children = types.NewOrderedMap()
    ncpInfoArray.EntityData.Leafs = types.NewOrderedMap()
    ncpInfoArray.EntityData.Leafs.Append("ncp-state", types.YLeaf{"NcpState", ncpInfoArray.NcpState})
    ncpInfoArray.EntityData.Leafs.Append("ncp-identifier", types.YLeaf{"NcpIdentifier", ncpInfoArray.NcpIdentifier})

    ncpInfoArray.EntityData.YListKeys = []string {}

    return &(ncpInfoArray.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation
// Information specific to the interface type
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // InterfaceTypeInfo. The type is ImCmdIntfTypeEnum.
    InterfaceTypeInfo interface{}

    // SRP interface information.
    SrpInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation

    // Tunnel interface information.
    TunnelInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation

    // Bundle interface information.
    BundleInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation

    // Serial interface information.
    SerialInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation

    // SONET POS interface information.
    SonetPosInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation

    // Tunnel GRE interface information.
    TunnelGreInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation

    // PseudowireHeadEnd interface information.
    PseudowireHeadEndInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation

    // Cem interface information.
    CemInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation

    // GCC interface information.
    GccInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation
}

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetEntityData() *types.CommonEntityData {
    interfaceTypeInformation.EntityData.YFilter = interfaceTypeInformation.YFilter
    interfaceTypeInformation.EntityData.YangName = "interface-type-information"
    interfaceTypeInformation.EntityData.BundleName = "cisco_ios_xr"
    interfaceTypeInformation.EntityData.ParentYangName = "interface"
    interfaceTypeInformation.EntityData.SegmentPath = "interface-type-information"
    interfaceTypeInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTypeInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTypeInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTypeInformation.EntityData.Children = types.NewOrderedMap()
    interfaceTypeInformation.EntityData.Children.Append("srp-information", types.YChild{"SrpInformation", &interfaceTypeInformation.SrpInformation})
    interfaceTypeInformation.EntityData.Children.Append("tunnel-information", types.YChild{"TunnelInformation", &interfaceTypeInformation.TunnelInformation})
    interfaceTypeInformation.EntityData.Children.Append("bundle-information", types.YChild{"BundleInformation", &interfaceTypeInformation.BundleInformation})
    interfaceTypeInformation.EntityData.Children.Append("serial-information", types.YChild{"SerialInformation", &interfaceTypeInformation.SerialInformation})
    interfaceTypeInformation.EntityData.Children.Append("sonet-pos-information", types.YChild{"SonetPosInformation", &interfaceTypeInformation.SonetPosInformation})
    interfaceTypeInformation.EntityData.Children.Append("tunnel-gre-information", types.YChild{"TunnelGreInformation", &interfaceTypeInformation.TunnelGreInformation})
    interfaceTypeInformation.EntityData.Children.Append("pseudowire-head-end-information", types.YChild{"PseudowireHeadEndInformation", &interfaceTypeInformation.PseudowireHeadEndInformation})
    interfaceTypeInformation.EntityData.Children.Append("cem-information", types.YChild{"CemInformation", &interfaceTypeInformation.CemInformation})
    interfaceTypeInformation.EntityData.Children.Append("gcc-information", types.YChild{"GccInformation", &interfaceTypeInformation.GccInformation})
    interfaceTypeInformation.EntityData.Leafs = types.NewOrderedMap()
    interfaceTypeInformation.EntityData.Leafs.Append("interface-type-info", types.YLeaf{"InterfaceTypeInfo", interfaceTypeInformation.InterfaceTypeInfo})

    interfaceTypeInformation.EntityData.YListKeys = []string {}

    return &(interfaceTypeInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation
// SRP interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRP-specific data.
    SrpInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation

    // SRP-specific packet and byte counters.
    SrpStatistics Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetEntityData() *types.CommonEntityData {
    srpInformation.EntityData.YFilter = srpInformation.YFilter
    srpInformation.EntityData.YangName = "srp-information"
    srpInformation.EntityData.BundleName = "cisco_ios_xr"
    srpInformation.EntityData.ParentYangName = "interface-type-information"
    srpInformation.EntityData.SegmentPath = "srp-information"
    srpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srpInformation.EntityData.Children = types.NewOrderedMap()
    srpInformation.EntityData.Children.Append("srp-information", types.YChild{"SrpInformation", &srpInformation.SrpInformation})
    srpInformation.EntityData.Children.Append("srp-statistics", types.YChild{"SrpStatistics", &srpInformation.SrpStatistics})
    srpInformation.EntityData.Leafs = types.NewOrderedMap()

    srpInformation.EntityData.YListKeys = []string {}

    return &(srpInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation
// SRP-specific data
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRP IPS information.
    IpsInfo Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo

    // SRP topology information.
    TopologyInfo Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo

    // SRP SRR information.
    SrrInfo Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo

    // SRP rate limit information.
    RateLimitInfo Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetEntityData() *types.CommonEntityData {
    srpInformation.EntityData.YFilter = srpInformation.YFilter
    srpInformation.EntityData.YangName = "srp-information"
    srpInformation.EntityData.BundleName = "cisco_ios_xr"
    srpInformation.EntityData.ParentYangName = "srp-information"
    srpInformation.EntityData.SegmentPath = "srp-information"
    srpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srpInformation.EntityData.Children = types.NewOrderedMap()
    srpInformation.EntityData.Children.Append("ips-info", types.YChild{"IpsInfo", &srpInformation.IpsInfo})
    srpInformation.EntityData.Children.Append("topology-info", types.YChild{"TopologyInfo", &srpInformation.TopologyInfo})
    srpInformation.EntityData.Children.Append("srr-info", types.YChild{"SrrInfo", &srpInformation.SrrInfo})
    srpInformation.EntityData.Children.Append("rate-limit-info", types.YChild{"RateLimitInfo", &srpInformation.RateLimitInfo})
    srpInformation.EntityData.Leafs = types.NewOrderedMap()

    srpInformation.EntityData.YListKeys = []string {}

    return &(srpInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo
// SRP IPS information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is the interfaceadministratively down. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAdminDown interface{}

    // IPS information. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation.
    LocalInformation []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation
}

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetEntityData() *types.CommonEntityData {
    ipsInfo.EntityData.YFilter = ipsInfo.YFilter
    ipsInfo.EntityData.YangName = "ips-info"
    ipsInfo.EntityData.BundleName = "cisco_ios_xr"
    ipsInfo.EntityData.ParentYangName = "srp-information"
    ipsInfo.EntityData.SegmentPath = "ips-info"
    ipsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipsInfo.EntityData.Children = types.NewOrderedMap()
    ipsInfo.EntityData.Children.Append("local-information", types.YChild{"LocalInformation", nil})
    for i := range ipsInfo.LocalInformation {
        ipsInfo.EntityData.Children.Append(types.GetSegmentPath(ipsInfo.LocalInformation[i]), types.YChild{"LocalInformation", ipsInfo.LocalInformation[i]})
    }
    ipsInfo.EntityData.Leafs = types.NewOrderedMap()
    ipsInfo.EntityData.Leafs.Append("is-admin-down", types.YLeaf{"IsAdminDown", ipsInfo.IsAdminDown})

    ipsInfo.EntityData.YListKeys = []string {}

    return &(ipsInfo.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation
// IPS information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address for node. The type is string.
    MacAddress interface{}

    // Inter card bus enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    IsInterCardBusEnabled interface{}

    // IPS Wait To Restore period in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    WtrTimerPeriod interface{}

    // Side A IPS details.
    SideA Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA

    // Side B IPS details.
    SideB Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetEntityData() *types.CommonEntityData {
    localInformation.EntityData.YFilter = localInformation.YFilter
    localInformation.EntityData.YangName = "local-information"
    localInformation.EntityData.BundleName = "cisco_ios_xr"
    localInformation.EntityData.ParentYangName = "ips-info"
    localInformation.EntityData.SegmentPath = "local-information"
    localInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localInformation.EntityData.Children = types.NewOrderedMap()
    localInformation.EntityData.Children.Append("side-a", types.YChild{"SideA", &localInformation.SideA})
    localInformation.EntityData.Children.Append("side-b", types.YChild{"SideB", &localInformation.SideB})
    localInformation.EntityData.Leafs = types.NewOrderedMap()
    localInformation.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", localInformation.MacAddress})
    localInformation.EntityData.Leafs.Append("is-inter-card-bus-enabled", types.YLeaf{"IsInterCardBusEnabled", localInformation.IsInterCardBusEnabled})
    localInformation.EntityData.Leafs.Append("wtr-timer-period", types.YLeaf{"WtrTimerPeriod", localInformation.WtrTimerPeriod})

    localInformation.EntityData.YListKeys = []string {}

    return &(localInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA
// Side A IPS details
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string.
    MacAddress interface{}

    // Wrap state. The type is SrpMgmtIpsWrapState.
    WrapState interface{}

    // SRP IPS packet send interval in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    PacketSentTimer interface{}

    // Time in seconds remaining until next send of an IPS request. The type is
    // interface{} with range: 0..4294967295. Units are second.
    SendTimerTimeRemaining interface{}

    // Time in seconds until wrap removal. The type is interface{} with range:
    // 0..4294967295. Units are second.
    WtrTimerRemaining interface{}

    // Self Detected Requests. The type is SrpMgmtIpsReq.
    SelfDetectedRequest interface{}

    // Remote Requests. The type is SrpMgmtIpsReq.
    RemoteRequest interface{}

    // Neighbour mac address for received message. The type is string.
    RxNeighborMacAddress interface{}

    // Type of message received. The type is SrpMgmtIpsReq.
    RxMessageType interface{}

    // Short/long path for received message. The type is SrpMgmtIpsPathInd.
    RxPathType interface{}

    // Time to live for received message. The type is interface{} with range:
    // 0..4294967295.
    RxTtl interface{}

    // Test for existence of an RX packet. The type is interface{} with range:
    // -2147483648..2147483647.
    RxPacketTest interface{}

    // Mac address of node receiving TXed messages. The type is string.
    TxNeighborMacAddress interface{}

    // Type of message transmitted. The type is SrpMgmtIpsReq.
    TxMessageType interface{}

    // Short/long path of transmitted message. The type is SrpMgmtIpsPathInd.
    TxPathType interface{}

    // Time to live for transmitted message. The type is interface{} with range:
    // 0..4294967295.
    TxTtl interface{}

    // Test for existence of a TX packet. The type is interface{} with range:
    // -2147483648..2147483647.
    TxPacketTest interface{}

    // Number of milliseconds to wait after an L1 failure is detected before
    // triggering an L2 wrap. The type is interface{} with range: 0..4294967295.
    // Units are millisecond.
    DelayKeepAliveTrigger interface{}

    // Failures presently asserted. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure.
    AssertedFailure []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure
}

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetEntityData() *types.CommonEntityData {
    sideA.EntityData.YFilter = sideA.YFilter
    sideA.EntityData.YangName = "side-a"
    sideA.EntityData.BundleName = "cisco_ios_xr"
    sideA.EntityData.ParentYangName = "local-information"
    sideA.EntityData.SegmentPath = "side-a"
    sideA.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sideA.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sideA.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sideA.EntityData.Children = types.NewOrderedMap()
    sideA.EntityData.Children.Append("asserted-failure", types.YChild{"AssertedFailure", nil})
    for i := range sideA.AssertedFailure {
        sideA.EntityData.Children.Append(types.GetSegmentPath(sideA.AssertedFailure[i]), types.YChild{"AssertedFailure", sideA.AssertedFailure[i]})
    }
    sideA.EntityData.Leafs = types.NewOrderedMap()
    sideA.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", sideA.MacAddress})
    sideA.EntityData.Leafs.Append("wrap-state", types.YLeaf{"WrapState", sideA.WrapState})
    sideA.EntityData.Leafs.Append("packet-sent-timer", types.YLeaf{"PacketSentTimer", sideA.PacketSentTimer})
    sideA.EntityData.Leafs.Append("send-timer-time-remaining", types.YLeaf{"SendTimerTimeRemaining", sideA.SendTimerTimeRemaining})
    sideA.EntityData.Leafs.Append("wtr-timer-remaining", types.YLeaf{"WtrTimerRemaining", sideA.WtrTimerRemaining})
    sideA.EntityData.Leafs.Append("self-detected-request", types.YLeaf{"SelfDetectedRequest", sideA.SelfDetectedRequest})
    sideA.EntityData.Leafs.Append("remote-request", types.YLeaf{"RemoteRequest", sideA.RemoteRequest})
    sideA.EntityData.Leafs.Append("rx-neighbor-mac-address", types.YLeaf{"RxNeighborMacAddress", sideA.RxNeighborMacAddress})
    sideA.EntityData.Leafs.Append("rx-message-type", types.YLeaf{"RxMessageType", sideA.RxMessageType})
    sideA.EntityData.Leafs.Append("rx-path-type", types.YLeaf{"RxPathType", sideA.RxPathType})
    sideA.EntityData.Leafs.Append("rx-ttl", types.YLeaf{"RxTtl", sideA.RxTtl})
    sideA.EntityData.Leafs.Append("rx-packet-test", types.YLeaf{"RxPacketTest", sideA.RxPacketTest})
    sideA.EntityData.Leafs.Append("tx-neighbor-mac-address", types.YLeaf{"TxNeighborMacAddress", sideA.TxNeighborMacAddress})
    sideA.EntityData.Leafs.Append("tx-message-type", types.YLeaf{"TxMessageType", sideA.TxMessageType})
    sideA.EntityData.Leafs.Append("tx-path-type", types.YLeaf{"TxPathType", sideA.TxPathType})
    sideA.EntityData.Leafs.Append("tx-ttl", types.YLeaf{"TxTtl", sideA.TxTtl})
    sideA.EntityData.Leafs.Append("tx-packet-test", types.YLeaf{"TxPacketTest", sideA.TxPacketTest})
    sideA.EntityData.Leafs.Append("delay-keep-alive-trigger", types.YLeaf{"DelayKeepAliveTrigger", sideA.DelayKeepAliveTrigger})

    sideA.EntityData.YListKeys = []string {}

    return &(sideA.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure
// Failures presently asserted
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Failure type. The type is SrpMgmtFailureEt.
    Type interface{}

    // Reported state. The type is SrpMgmtFailureStateEt.
    ReportedState interface{}

    // Debounced state. The type is SrpMgmtFailureStateEt.
    DebouncedState interface{}

    // Current state. The type is SrpMgmtFailureStateEt.
    CurrentState interface{}

    // Stable time. The type is interface{} with range: 0..18446744073709551615.
    StableTime interface{}

    // Debounce delay. The type is interface{} with range: 0..4294967295.
    DebouncedDelay interface{}
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetEntityData() *types.CommonEntityData {
    assertedFailure.EntityData.YFilter = assertedFailure.YFilter
    assertedFailure.EntityData.YangName = "asserted-failure"
    assertedFailure.EntityData.BundleName = "cisco_ios_xr"
    assertedFailure.EntityData.ParentYangName = "side-a"
    assertedFailure.EntityData.SegmentPath = "asserted-failure"
    assertedFailure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    assertedFailure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    assertedFailure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    assertedFailure.EntityData.Children = types.NewOrderedMap()
    assertedFailure.EntityData.Leafs = types.NewOrderedMap()
    assertedFailure.EntityData.Leafs.Append("type", types.YLeaf{"Type", assertedFailure.Type})
    assertedFailure.EntityData.Leafs.Append("reported-state", types.YLeaf{"ReportedState", assertedFailure.ReportedState})
    assertedFailure.EntityData.Leafs.Append("debounced-state", types.YLeaf{"DebouncedState", assertedFailure.DebouncedState})
    assertedFailure.EntityData.Leafs.Append("current-state", types.YLeaf{"CurrentState", assertedFailure.CurrentState})
    assertedFailure.EntityData.Leafs.Append("stable-time", types.YLeaf{"StableTime", assertedFailure.StableTime})
    assertedFailure.EntityData.Leafs.Append("debounced-delay", types.YLeaf{"DebouncedDelay", assertedFailure.DebouncedDelay})

    assertedFailure.EntityData.YListKeys = []string {}

    return &(assertedFailure.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB
// Side B IPS details
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string.
    MacAddress interface{}

    // Wrap state. The type is SrpMgmtIpsWrapState.
    WrapState interface{}

    // SRP IPS packet send interval in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    PacketSentTimer interface{}

    // Time in seconds remaining until next send of an IPS request. The type is
    // interface{} with range: 0..4294967295. Units are second.
    SendTimerTimeRemaining interface{}

    // Time in seconds until wrap removal. The type is interface{} with range:
    // 0..4294967295. Units are second.
    WtrTimerRemaining interface{}

    // Self Detected Requests. The type is SrpMgmtIpsReq.
    SelfDetectedRequest interface{}

    // Remote Requests. The type is SrpMgmtIpsReq.
    RemoteRequest interface{}

    // Neighbour mac address for received message. The type is string.
    RxNeighborMacAddress interface{}

    // Type of message received. The type is SrpMgmtIpsReq.
    RxMessageType interface{}

    // Short/long path for received message. The type is SrpMgmtIpsPathInd.
    RxPathType interface{}

    // Time to live for received message. The type is interface{} with range:
    // 0..4294967295.
    RxTtl interface{}

    // Test for existence of an RX packet. The type is interface{} with range:
    // -2147483648..2147483647.
    RxPacketTest interface{}

    // Mac address of node receiving TXed messages. The type is string.
    TxNeighborMacAddress interface{}

    // Type of message transmitted. The type is SrpMgmtIpsReq.
    TxMessageType interface{}

    // Short/long path of transmitted message. The type is SrpMgmtIpsPathInd.
    TxPathType interface{}

    // Time to live for transmitted message. The type is interface{} with range:
    // 0..4294967295.
    TxTtl interface{}

    // Test for existence of a TX packet. The type is interface{} with range:
    // -2147483648..2147483647.
    TxPacketTest interface{}

    // Number of milliseconds to wait after an L1 failure is detected before
    // triggering an L2 wrap. The type is interface{} with range: 0..4294967295.
    // Units are millisecond.
    DelayKeepAliveTrigger interface{}

    // Failures presently asserted. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure.
    AssertedFailure []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure
}

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetEntityData() *types.CommonEntityData {
    sideB.EntityData.YFilter = sideB.YFilter
    sideB.EntityData.YangName = "side-b"
    sideB.EntityData.BundleName = "cisco_ios_xr"
    sideB.EntityData.ParentYangName = "local-information"
    sideB.EntityData.SegmentPath = "side-b"
    sideB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sideB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sideB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sideB.EntityData.Children = types.NewOrderedMap()
    sideB.EntityData.Children.Append("asserted-failure", types.YChild{"AssertedFailure", nil})
    for i := range sideB.AssertedFailure {
        sideB.EntityData.Children.Append(types.GetSegmentPath(sideB.AssertedFailure[i]), types.YChild{"AssertedFailure", sideB.AssertedFailure[i]})
    }
    sideB.EntityData.Leafs = types.NewOrderedMap()
    sideB.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", sideB.MacAddress})
    sideB.EntityData.Leafs.Append("wrap-state", types.YLeaf{"WrapState", sideB.WrapState})
    sideB.EntityData.Leafs.Append("packet-sent-timer", types.YLeaf{"PacketSentTimer", sideB.PacketSentTimer})
    sideB.EntityData.Leafs.Append("send-timer-time-remaining", types.YLeaf{"SendTimerTimeRemaining", sideB.SendTimerTimeRemaining})
    sideB.EntityData.Leafs.Append("wtr-timer-remaining", types.YLeaf{"WtrTimerRemaining", sideB.WtrTimerRemaining})
    sideB.EntityData.Leafs.Append("self-detected-request", types.YLeaf{"SelfDetectedRequest", sideB.SelfDetectedRequest})
    sideB.EntityData.Leafs.Append("remote-request", types.YLeaf{"RemoteRequest", sideB.RemoteRequest})
    sideB.EntityData.Leafs.Append("rx-neighbor-mac-address", types.YLeaf{"RxNeighborMacAddress", sideB.RxNeighborMacAddress})
    sideB.EntityData.Leafs.Append("rx-message-type", types.YLeaf{"RxMessageType", sideB.RxMessageType})
    sideB.EntityData.Leafs.Append("rx-path-type", types.YLeaf{"RxPathType", sideB.RxPathType})
    sideB.EntityData.Leafs.Append("rx-ttl", types.YLeaf{"RxTtl", sideB.RxTtl})
    sideB.EntityData.Leafs.Append("rx-packet-test", types.YLeaf{"RxPacketTest", sideB.RxPacketTest})
    sideB.EntityData.Leafs.Append("tx-neighbor-mac-address", types.YLeaf{"TxNeighborMacAddress", sideB.TxNeighborMacAddress})
    sideB.EntityData.Leafs.Append("tx-message-type", types.YLeaf{"TxMessageType", sideB.TxMessageType})
    sideB.EntityData.Leafs.Append("tx-path-type", types.YLeaf{"TxPathType", sideB.TxPathType})
    sideB.EntityData.Leafs.Append("tx-ttl", types.YLeaf{"TxTtl", sideB.TxTtl})
    sideB.EntityData.Leafs.Append("tx-packet-test", types.YLeaf{"TxPacketTest", sideB.TxPacketTest})
    sideB.EntityData.Leafs.Append("delay-keep-alive-trigger", types.YLeaf{"DelayKeepAliveTrigger", sideB.DelayKeepAliveTrigger})

    sideB.EntityData.YListKeys = []string {}

    return &(sideB.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure
// Failures presently asserted
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Failure type. The type is SrpMgmtFailureEt.
    Type interface{}

    // Reported state. The type is SrpMgmtFailureStateEt.
    ReportedState interface{}

    // Debounced state. The type is SrpMgmtFailureStateEt.
    DebouncedState interface{}

    // Current state. The type is SrpMgmtFailureStateEt.
    CurrentState interface{}

    // Stable time. The type is interface{} with range: 0..18446744073709551615.
    StableTime interface{}

    // Debounce delay. The type is interface{} with range: 0..4294967295.
    DebouncedDelay interface{}
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetEntityData() *types.CommonEntityData {
    assertedFailure.EntityData.YFilter = assertedFailure.YFilter
    assertedFailure.EntityData.YangName = "asserted-failure"
    assertedFailure.EntityData.BundleName = "cisco_ios_xr"
    assertedFailure.EntityData.ParentYangName = "side-b"
    assertedFailure.EntityData.SegmentPath = "asserted-failure"
    assertedFailure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    assertedFailure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    assertedFailure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    assertedFailure.EntityData.Children = types.NewOrderedMap()
    assertedFailure.EntityData.Leafs = types.NewOrderedMap()
    assertedFailure.EntityData.Leafs.Append("type", types.YLeaf{"Type", assertedFailure.Type})
    assertedFailure.EntityData.Leafs.Append("reported-state", types.YLeaf{"ReportedState", assertedFailure.ReportedState})
    assertedFailure.EntityData.Leafs.Append("debounced-state", types.YLeaf{"DebouncedState", assertedFailure.DebouncedState})
    assertedFailure.EntityData.Leafs.Append("current-state", types.YLeaf{"CurrentState", assertedFailure.CurrentState})
    assertedFailure.EntityData.Leafs.Append("stable-time", types.YLeaf{"StableTime", assertedFailure.StableTime})
    assertedFailure.EntityData.Leafs.Append("debounced-delay", types.YLeaf{"DebouncedDelay", assertedFailure.DebouncedDelay})

    assertedFailure.EntityData.YListKeys = []string {}

    return &(assertedFailure.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo
// SRP topology information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is the interfaceadministratively down. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAdminDown interface{}

    // Detailed SRP topology information. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation.
    LocalInformation []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation
}

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetEntityData() *types.CommonEntityData {
    topologyInfo.EntityData.YFilter = topologyInfo.YFilter
    topologyInfo.EntityData.YangName = "topology-info"
    topologyInfo.EntityData.BundleName = "cisco_ios_xr"
    topologyInfo.EntityData.ParentYangName = "srp-information"
    topologyInfo.EntityData.SegmentPath = "topology-info"
    topologyInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyInfo.EntityData.Children = types.NewOrderedMap()
    topologyInfo.EntityData.Children.Append("local-information", types.YChild{"LocalInformation", nil})
    for i := range topologyInfo.LocalInformation {
        topologyInfo.EntityData.Children.Append(types.GetSegmentPath(topologyInfo.LocalInformation[i]), types.YChild{"LocalInformation", topologyInfo.LocalInformation[i]})
    }
    topologyInfo.EntityData.Leafs = types.NewOrderedMap()
    topologyInfo.EntityData.Leafs.Append("is-admin-down", types.YLeaf{"IsAdminDown", topologyInfo.IsAdminDown})

    topologyInfo.EntityData.YListKeys = []string {}

    return &(topologyInfo.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation
// Detailed SRP topology information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // How often a topology pkt is sent. The type is interface{} with range:
    // 0..4294967295.
    TopologyTimer interface{}

    // Time remaining until next topo pkt sent. The type is interface{} with
    // range: 0..4294967295.
    NextTopologyPacketDelay interface{}

    // Time since last topo pkt was received. The type is interface{} with range:
    // 0..4294967295.
    TimeSinceLastTopologyPacketReceived interface{}

    // Time since last topology change. The type is interface{} with range:
    // 0..4294967295.
    TimeSinceLastTopologyChange interface{}

    // Number of nodes on ring. The type is interface{} with range: 0..65535.
    NumberOfNodesOnRing interface{}

    // List of nodes on the ring info. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode.
    RingNode []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetEntityData() *types.CommonEntityData {
    localInformation.EntityData.YFilter = localInformation.YFilter
    localInformation.EntityData.YangName = "local-information"
    localInformation.EntityData.BundleName = "cisco_ios_xr"
    localInformation.EntityData.ParentYangName = "topology-info"
    localInformation.EntityData.SegmentPath = "local-information"
    localInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localInformation.EntityData.Children = types.NewOrderedMap()
    localInformation.EntityData.Children.Append("ring-node", types.YChild{"RingNode", nil})
    for i := range localInformation.RingNode {
        localInformation.EntityData.Children.Append(types.GetSegmentPath(localInformation.RingNode[i]), types.YChild{"RingNode", localInformation.RingNode[i]})
    }
    localInformation.EntityData.Leafs = types.NewOrderedMap()
    localInformation.EntityData.Leafs.Append("topology-timer", types.YLeaf{"TopologyTimer", localInformation.TopologyTimer})
    localInformation.EntityData.Leafs.Append("next-topology-packet-delay", types.YLeaf{"NextTopologyPacketDelay", localInformation.NextTopologyPacketDelay})
    localInformation.EntityData.Leafs.Append("time-since-last-topology-packet-received", types.YLeaf{"TimeSinceLastTopologyPacketReceived", localInformation.TimeSinceLastTopologyPacketReceived})
    localInformation.EntityData.Leafs.Append("time-since-last-topology-change", types.YLeaf{"TimeSinceLastTopologyChange", localInformation.TimeSinceLastTopologyChange})
    localInformation.EntityData.Leafs.Append("number-of-nodes-on-ring", types.YLeaf{"NumberOfNodesOnRing", localInformation.NumberOfNodesOnRing})

    localInformation.EntityData.YListKeys = []string {}

    return &(localInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode
// List of nodes on the ring info
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outer-ring hops to reach this node. The type is interface{} with range:
    // 0..65535.
    HopCount interface{}

    // MAC address. The type is string.
    MacAddress interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Wrap state. The type is interface{} with range: -2147483648..2147483647.
    IsWrapped interface{}

    // SRR protocol supported. The type is interface{} with range:
    // -2147483648..2147483647.
    IsSrrSupported interface{}

    // Node name. The type is string.
    NodeName interface{}
}

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetEntityData() *types.CommonEntityData {
    ringNode.EntityData.YFilter = ringNode.YFilter
    ringNode.EntityData.YangName = "ring-node"
    ringNode.EntityData.BundleName = "cisco_ios_xr"
    ringNode.EntityData.ParentYangName = "local-information"
    ringNode.EntityData.SegmentPath = "ring-node"
    ringNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ringNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ringNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ringNode.EntityData.Children = types.NewOrderedMap()
    ringNode.EntityData.Leafs = types.NewOrderedMap()
    ringNode.EntityData.Leafs.Append("hop-count", types.YLeaf{"HopCount", ringNode.HopCount})
    ringNode.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", ringNode.MacAddress})
    ringNode.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", ringNode.Ipv4Address})
    ringNode.EntityData.Leafs.Append("is-wrapped", types.YLeaf{"IsWrapped", ringNode.IsWrapped})
    ringNode.EntityData.Leafs.Append("is-srr-supported", types.YLeaf{"IsSrrSupported", ringNode.IsSrrSupported})
    ringNode.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", ringNode.NodeName})

    ringNode.EntityData.YListKeys = []string {}

    return &(ringNode.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo
// SRP SRR information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is the interfaceadministratively down. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAdminDown interface{}

    // SRR enabled. The type is interface{} with range: -2147483648..2147483647.
    IsSrrEnabled interface{}

    // SRP information. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo.
    SrrDetailedInfo []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo
}

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetEntityData() *types.CommonEntityData {
    srrInfo.EntityData.YFilter = srrInfo.YFilter
    srrInfo.EntityData.YangName = "srr-info"
    srrInfo.EntityData.BundleName = "cisco_ios_xr"
    srrInfo.EntityData.ParentYangName = "srp-information"
    srrInfo.EntityData.SegmentPath = "srr-info"
    srrInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srrInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srrInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srrInfo.EntityData.Children = types.NewOrderedMap()
    srrInfo.EntityData.Children.Append("srr-detailed-info", types.YChild{"SrrDetailedInfo", nil})
    for i := range srrInfo.SrrDetailedInfo {
        srrInfo.EntityData.Children.Append(types.GetSegmentPath(srrInfo.SrrDetailedInfo[i]), types.YChild{"SrrDetailedInfo", srrInfo.SrrDetailedInfo[i]})
    }
    srrInfo.EntityData.Leafs = types.NewOrderedMap()
    srrInfo.EntityData.Leafs.Append("is-admin-down", types.YLeaf{"IsAdminDown", srrInfo.IsAdminDown})
    srrInfo.EntityData.Leafs.Append("is-srr-enabled", types.YLeaf{"IsSrrEnabled", srrInfo.IsSrrEnabled})

    srrInfo.EntityData.YListKeys = []string {}

    return &(srrInfo.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo
// SRP information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version number. The type is interface{} with range: 0..4294967295.
    VersionNumber interface{}

    // Wrong version recieved. The type is interface{} with range:
    // -2147483648..2147483647.
    IsWrongVersionReceived interface{}

    // Time that last wrong version message recieved. The type is interface{} with
    // range: 0..4294967295.
    LastWrongVersionReceiveTime interface{}

    // SRR node mac address. The type is string.
    MacAddress interface{}

    // SRR node state. The type is SrpMgmtSrrNodeState.
    NodeState interface{}

    // Is the outer ring in use. The type is interface{} with range:
    // -2147483648..2147483647.
    IsOuterRingInUse interface{}

    // Is the inner ring in use. The type is interface{} with range:
    // -2147483648..2147483647.
    IsInnerRingInUse interface{}

    // Is announcing enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAnnounce interface{}

    // Outer fail type. The type is SrpMgmtSrrFailure.
    OuterFailType interface{}

    // Inner fail type. The type is SrpMgmtSrrFailure.
    InnerFailType interface{}

    // SRR packet send timer interval in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    PacketSendTimer interface{}

    // Time remaining in seconds to next SRR packet send. The type is interface{}
    // with range: 0..4294967295. Units are second.
    NextSrrPacketSendTime interface{}

    // Single ring bandwidth Mbps. The type is interface{} with range:
    // 0..4294967295. Units are Mbit/s.
    SingleRingBw interface{}

    // SRR Wait To Restore interval delay in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    WtrTime interface{}

    // Time remaining in seconds until next outer ring wrap removal. The type is
    // interface{} with range: 0..4294967295. Units are second.
    WtrTimerRemainingOuterRing interface{}

    // Time remaining in seconds until next inner ring wrap removal. The type is
    // interface{} with range: 0..4294967295. Units are second.
    WtrTimerRemainingInnerRing interface{}

    // List of nodes on the ring info. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing.
    NodesOnRing []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing

    // nodes not in topology map. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing.
    NodesNotOnRing []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing
}

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetEntityData() *types.CommonEntityData {
    srrDetailedInfo.EntityData.YFilter = srrDetailedInfo.YFilter
    srrDetailedInfo.EntityData.YangName = "srr-detailed-info"
    srrDetailedInfo.EntityData.BundleName = "cisco_ios_xr"
    srrDetailedInfo.EntityData.ParentYangName = "srr-info"
    srrDetailedInfo.EntityData.SegmentPath = "srr-detailed-info"
    srrDetailedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srrDetailedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srrDetailedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srrDetailedInfo.EntityData.Children = types.NewOrderedMap()
    srrDetailedInfo.EntityData.Children.Append("nodes-on-ring", types.YChild{"NodesOnRing", nil})
    for i := range srrDetailedInfo.NodesOnRing {
        srrDetailedInfo.EntityData.Children.Append(types.GetSegmentPath(srrDetailedInfo.NodesOnRing[i]), types.YChild{"NodesOnRing", srrDetailedInfo.NodesOnRing[i]})
    }
    srrDetailedInfo.EntityData.Children.Append("nodes-not-on-ring", types.YChild{"NodesNotOnRing", nil})
    for i := range srrDetailedInfo.NodesNotOnRing {
        srrDetailedInfo.EntityData.Children.Append(types.GetSegmentPath(srrDetailedInfo.NodesNotOnRing[i]), types.YChild{"NodesNotOnRing", srrDetailedInfo.NodesNotOnRing[i]})
    }
    srrDetailedInfo.EntityData.Leafs = types.NewOrderedMap()
    srrDetailedInfo.EntityData.Leafs.Append("version-number", types.YLeaf{"VersionNumber", srrDetailedInfo.VersionNumber})
    srrDetailedInfo.EntityData.Leafs.Append("is-wrong-version-received", types.YLeaf{"IsWrongVersionReceived", srrDetailedInfo.IsWrongVersionReceived})
    srrDetailedInfo.EntityData.Leafs.Append("last-wrong-version-receive-time", types.YLeaf{"LastWrongVersionReceiveTime", srrDetailedInfo.LastWrongVersionReceiveTime})
    srrDetailedInfo.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", srrDetailedInfo.MacAddress})
    srrDetailedInfo.EntityData.Leafs.Append("node-state", types.YLeaf{"NodeState", srrDetailedInfo.NodeState})
    srrDetailedInfo.EntityData.Leafs.Append("is-outer-ring-in-use", types.YLeaf{"IsOuterRingInUse", srrDetailedInfo.IsOuterRingInUse})
    srrDetailedInfo.EntityData.Leafs.Append("is-inner-ring-in-use", types.YLeaf{"IsInnerRingInUse", srrDetailedInfo.IsInnerRingInUse})
    srrDetailedInfo.EntityData.Leafs.Append("is-announce", types.YLeaf{"IsAnnounce", srrDetailedInfo.IsAnnounce})
    srrDetailedInfo.EntityData.Leafs.Append("outer-fail-type", types.YLeaf{"OuterFailType", srrDetailedInfo.OuterFailType})
    srrDetailedInfo.EntityData.Leafs.Append("inner-fail-type", types.YLeaf{"InnerFailType", srrDetailedInfo.InnerFailType})
    srrDetailedInfo.EntityData.Leafs.Append("packet-send-timer", types.YLeaf{"PacketSendTimer", srrDetailedInfo.PacketSendTimer})
    srrDetailedInfo.EntityData.Leafs.Append("next-srr-packet-send-time", types.YLeaf{"NextSrrPacketSendTime", srrDetailedInfo.NextSrrPacketSendTime})
    srrDetailedInfo.EntityData.Leafs.Append("single-ring-bw", types.YLeaf{"SingleRingBw", srrDetailedInfo.SingleRingBw})
    srrDetailedInfo.EntityData.Leafs.Append("wtr-time", types.YLeaf{"WtrTime", srrDetailedInfo.WtrTime})
    srrDetailedInfo.EntityData.Leafs.Append("wtr-timer-remaining-outer-ring", types.YLeaf{"WtrTimerRemainingOuterRing", srrDetailedInfo.WtrTimerRemainingOuterRing})
    srrDetailedInfo.EntityData.Leafs.Append("wtr-timer-remaining-inner-ring", types.YLeaf{"WtrTimerRemainingInnerRing", srrDetailedInfo.WtrTimerRemainingInnerRing})

    srrDetailedInfo.EntityData.YListKeys = []string {}

    return &(srrDetailedInfo.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing
// List of nodes on the ring info
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node name. The type is string.
    NodeName interface{}

    // Does the SRR information exist for this node. The type is interface{} with
    // range: -2147483648..2147483647.
    SrrEntryExits interface{}

    // node mac address. The type is string.
    MacAddress interface{}

    // Outer failure. The type is SrpMgmtSrrFailure.
    OuterFailure interface{}

    // Inner failure. The type is SrpMgmtSrrFailure.
    InnerFailure interface{}

    // Announce last received ?. The type is interface{} with range:
    // -2147483648..2147483647.
    IsLastAnnounceReceived interface{}

    // Announce last received. The type is interface{} with range: 0..4294967295.
    LastAnnounceReceivedTime interface{}
}

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetEntityData() *types.CommonEntityData {
    nodesOnRing.EntityData.YFilter = nodesOnRing.YFilter
    nodesOnRing.EntityData.YangName = "nodes-on-ring"
    nodesOnRing.EntityData.BundleName = "cisco_ios_xr"
    nodesOnRing.EntityData.ParentYangName = "srr-detailed-info"
    nodesOnRing.EntityData.SegmentPath = "nodes-on-ring"
    nodesOnRing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodesOnRing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodesOnRing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodesOnRing.EntityData.Children = types.NewOrderedMap()
    nodesOnRing.EntityData.Leafs = types.NewOrderedMap()
    nodesOnRing.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", nodesOnRing.NodeName})
    nodesOnRing.EntityData.Leafs.Append("srr-entry-exits", types.YLeaf{"SrrEntryExits", nodesOnRing.SrrEntryExits})
    nodesOnRing.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", nodesOnRing.MacAddress})
    nodesOnRing.EntityData.Leafs.Append("outer-failure", types.YLeaf{"OuterFailure", nodesOnRing.OuterFailure})
    nodesOnRing.EntityData.Leafs.Append("inner-failure", types.YLeaf{"InnerFailure", nodesOnRing.InnerFailure})
    nodesOnRing.EntityData.Leafs.Append("is-last-announce-received", types.YLeaf{"IsLastAnnounceReceived", nodesOnRing.IsLastAnnounceReceived})
    nodesOnRing.EntityData.Leafs.Append("last-announce-received-time", types.YLeaf{"LastAnnounceReceivedTime", nodesOnRing.LastAnnounceReceivedTime})

    nodesOnRing.EntityData.YListKeys = []string {}

    return &(nodesOnRing.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing
// nodes not in topology map
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node name. The type is string.
    NodeName interface{}

    // Does the SRR information exist for this node. The type is interface{} with
    // range: -2147483648..2147483647.
    SrrEntryExits interface{}

    // node mac address. The type is string.
    MacAddress interface{}

    // Outer failure. The type is SrpMgmtSrrFailure.
    OuterFailure interface{}

    // Inner failure. The type is SrpMgmtSrrFailure.
    InnerFailure interface{}

    // Announce last received ?. The type is interface{} with range:
    // -2147483648..2147483647.
    IsLastAnnounceReceived interface{}

    // Announce last received. The type is interface{} with range: 0..4294967295.
    LastAnnounceReceivedTime interface{}
}

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetEntityData() *types.CommonEntityData {
    nodesNotOnRing.EntityData.YFilter = nodesNotOnRing.YFilter
    nodesNotOnRing.EntityData.YangName = "nodes-not-on-ring"
    nodesNotOnRing.EntityData.BundleName = "cisco_ios_xr"
    nodesNotOnRing.EntityData.ParentYangName = "srr-detailed-info"
    nodesNotOnRing.EntityData.SegmentPath = "nodes-not-on-ring"
    nodesNotOnRing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodesNotOnRing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodesNotOnRing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodesNotOnRing.EntityData.Children = types.NewOrderedMap()
    nodesNotOnRing.EntityData.Leafs = types.NewOrderedMap()
    nodesNotOnRing.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", nodesNotOnRing.NodeName})
    nodesNotOnRing.EntityData.Leafs.Append("srr-entry-exits", types.YLeaf{"SrrEntryExits", nodesNotOnRing.SrrEntryExits})
    nodesNotOnRing.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", nodesNotOnRing.MacAddress})
    nodesNotOnRing.EntityData.Leafs.Append("outer-failure", types.YLeaf{"OuterFailure", nodesNotOnRing.OuterFailure})
    nodesNotOnRing.EntityData.Leafs.Append("inner-failure", types.YLeaf{"InnerFailure", nodesNotOnRing.InnerFailure})
    nodesNotOnRing.EntityData.Leafs.Append("is-last-announce-received", types.YLeaf{"IsLastAnnounceReceived", nodesNotOnRing.IsLastAnnounceReceived})
    nodesNotOnRing.EntityData.Leafs.Append("last-announce-received-time", types.YLeaf{"LastAnnounceReceivedTime", nodesNotOnRing.LastAnnounceReceivedTime})

    nodesNotOnRing.EntityData.YListKeys = []string {}

    return &(nodesNotOnRing.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo
// SRP rate limit information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is the interfaceadministratively down. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAdminDown interface{}

    // SRP rate limit information. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo.
    RateLimitDetailedInfo []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo
}

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetEntityData() *types.CommonEntityData {
    rateLimitInfo.EntityData.YFilter = rateLimitInfo.YFilter
    rateLimitInfo.EntityData.YangName = "rate-limit-info"
    rateLimitInfo.EntityData.BundleName = "cisco_ios_xr"
    rateLimitInfo.EntityData.ParentYangName = "srp-information"
    rateLimitInfo.EntityData.SegmentPath = "rate-limit-info"
    rateLimitInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rateLimitInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rateLimitInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rateLimitInfo.EntityData.Children = types.NewOrderedMap()
    rateLimitInfo.EntityData.Children.Append("rate-limit-detailed-info", types.YChild{"RateLimitDetailedInfo", nil})
    for i := range rateLimitInfo.RateLimitDetailedInfo {
        rateLimitInfo.EntityData.Children.Append(types.GetSegmentPath(rateLimitInfo.RateLimitDetailedInfo[i]), types.YChild{"RateLimitDetailedInfo", rateLimitInfo.RateLimitDetailedInfo[i]})
    }
    rateLimitInfo.EntityData.Leafs = types.NewOrderedMap()
    rateLimitInfo.EntityData.Leafs.Append("is-admin-down", types.YLeaf{"IsAdminDown", rateLimitInfo.IsAdminDown})

    rateLimitInfo.EntityData.YListKeys = []string {}

    return &(rateLimitInfo.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo
// SRP rate limit information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum SRP priority for high-priority transmit queue. The type is
    // interface{} with range: 0..65535.
    MinPriorityValue interface{}
}

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetEntityData() *types.CommonEntityData {
    rateLimitDetailedInfo.EntityData.YFilter = rateLimitDetailedInfo.YFilter
    rateLimitDetailedInfo.EntityData.YangName = "rate-limit-detailed-info"
    rateLimitDetailedInfo.EntityData.BundleName = "cisco_ios_xr"
    rateLimitDetailedInfo.EntityData.ParentYangName = "rate-limit-info"
    rateLimitDetailedInfo.EntityData.SegmentPath = "rate-limit-detailed-info"
    rateLimitDetailedInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rateLimitDetailedInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rateLimitDetailedInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rateLimitDetailedInfo.EntityData.Children = types.NewOrderedMap()
    rateLimitDetailedInfo.EntityData.Leafs = types.NewOrderedMap()
    rateLimitDetailedInfo.EntityData.Leafs.Append("min-priority-value", types.YLeaf{"MinPriorityValue", rateLimitDetailedInfo.MinPriorityValue})

    rateLimitDetailedInfo.EntityData.YListKeys = []string {}

    return &(rateLimitDetailedInfo.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics
// SRP-specific packet and byte counters
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data rate interval (5 mins or 30 seconds). The type is interface{} with
    // range: 0..4294967295. Units are second.
    DataRateInterval interface{}

    // Data rates for side A interface.
    SideADataRate Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate

    // Data rates for side B interface.
    SideBDataRate Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate

    // Errors for side A interface.
    SideAErrors Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors

    // Errors for side B interface.
    SideBErrors Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors
}

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetEntityData() *types.CommonEntityData {
    srpStatistics.EntityData.YFilter = srpStatistics.YFilter
    srpStatistics.EntityData.YangName = "srp-statistics"
    srpStatistics.EntityData.BundleName = "cisco_ios_xr"
    srpStatistics.EntityData.ParentYangName = "srp-information"
    srpStatistics.EntityData.SegmentPath = "srp-statistics"
    srpStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srpStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srpStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srpStatistics.EntityData.Children = types.NewOrderedMap()
    srpStatistics.EntityData.Children.Append("side-a-data-rate", types.YChild{"SideADataRate", &srpStatistics.SideADataRate})
    srpStatistics.EntityData.Children.Append("side-b-data-rate", types.YChild{"SideBDataRate", &srpStatistics.SideBDataRate})
    srpStatistics.EntityData.Children.Append("side-a-errors", types.YChild{"SideAErrors", &srpStatistics.SideAErrors})
    srpStatistics.EntityData.Children.Append("side-b-errors", types.YChild{"SideBErrors", &srpStatistics.SideBErrors})
    srpStatistics.EntityData.Leafs = types.NewOrderedMap()
    srpStatistics.EntityData.Leafs.Append("data-rate-interval", types.YLeaf{"DataRateInterval", srpStatistics.DataRateInterval})

    srpStatistics.EntityData.YListKeys = []string {}

    return &(srpStatistics.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate
// Data rates for side A interface
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent bit rate. The type is interface{} with range: 0..4294967295.
    BitRateSent interface{}

    // Sent packet rate. The type is interface{} with range: 0..4294967295.
    PacketRateSent interface{}

    // Received bit rate. The type is interface{} with range: 0..4294967295.
    BitRateReceived interface{}

    // Received packet rate. The type is interface{} with range: 0..4294967295.
    PacketRateReceived interface{}
}

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetEntityData() *types.CommonEntityData {
    sideADataRate.EntityData.YFilter = sideADataRate.YFilter
    sideADataRate.EntityData.YangName = "side-a-data-rate"
    sideADataRate.EntityData.BundleName = "cisco_ios_xr"
    sideADataRate.EntityData.ParentYangName = "srp-statistics"
    sideADataRate.EntityData.SegmentPath = "side-a-data-rate"
    sideADataRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sideADataRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sideADataRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sideADataRate.EntityData.Children = types.NewOrderedMap()
    sideADataRate.EntityData.Leafs = types.NewOrderedMap()
    sideADataRate.EntityData.Leafs.Append("bit-rate-sent", types.YLeaf{"BitRateSent", sideADataRate.BitRateSent})
    sideADataRate.EntityData.Leafs.Append("packet-rate-sent", types.YLeaf{"PacketRateSent", sideADataRate.PacketRateSent})
    sideADataRate.EntityData.Leafs.Append("bit-rate-received", types.YLeaf{"BitRateReceived", sideADataRate.BitRateReceived})
    sideADataRate.EntityData.Leafs.Append("packet-rate-received", types.YLeaf{"PacketRateReceived", sideADataRate.PacketRateReceived})

    sideADataRate.EntityData.YListKeys = []string {}

    return &(sideADataRate.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate
// Data rates for side B interface
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent bit rate. The type is interface{} with range: 0..4294967295.
    BitRateSent interface{}

    // Sent packet rate. The type is interface{} with range: 0..4294967295.
    PacketRateSent interface{}

    // Received bit rate. The type is interface{} with range: 0..4294967295.
    BitRateReceived interface{}

    // Received packet rate. The type is interface{} with range: 0..4294967295.
    PacketRateReceived interface{}
}

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetEntityData() *types.CommonEntityData {
    sideBDataRate.EntityData.YFilter = sideBDataRate.YFilter
    sideBDataRate.EntityData.YangName = "side-b-data-rate"
    sideBDataRate.EntityData.BundleName = "cisco_ios_xr"
    sideBDataRate.EntityData.ParentYangName = "srp-statistics"
    sideBDataRate.EntityData.SegmentPath = "side-b-data-rate"
    sideBDataRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sideBDataRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sideBDataRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sideBDataRate.EntityData.Children = types.NewOrderedMap()
    sideBDataRate.EntityData.Leafs = types.NewOrderedMap()
    sideBDataRate.EntityData.Leafs.Append("bit-rate-sent", types.YLeaf{"BitRateSent", sideBDataRate.BitRateSent})
    sideBDataRate.EntityData.Leafs.Append("packet-rate-sent", types.YLeaf{"PacketRateSent", sideBDataRate.PacketRateSent})
    sideBDataRate.EntityData.Leafs.Append("bit-rate-received", types.YLeaf{"BitRateReceived", sideBDataRate.BitRateReceived})
    sideBDataRate.EntityData.Leafs.Append("packet-rate-received", types.YLeaf{"PacketRateReceived", sideBDataRate.PacketRateReceived})

    sideBDataRate.EntityData.YListKeys = []string {}

    return &(sideBDataRate.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors
// Errors for side A interface
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error packets received. The type is interface{} with range: 0..4294967295.
    ErrorPacketsReceived interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input insufficient resources events. The type is interface{} with range:
    // 0..4294967295.
    InputInsufficientResourceEvents interface{}

    // Aborts received at MAC/RAC. The type is interface{} with range:
    // 0..4294967295.
    MacAbortsReceived interface{}

    // Too small packets received at MAC/RAC. The type is interface{} with range:
    // 0..4294967295.
    MacRuntPacketsReceived interface{}

    // Too large packets received at MAC/RAC. The type is interface{} with range:
    // 0..4294967295.
    MacGiantPacketsReceived interface{}

    // Too small packets received at framer. The type is interface{} with range:
    // 0..4294967295.
    FramerRuntPacketsReceived interface{}

    // Too large packets received at framer. The type is interface{} with range:
    // 0..4294967295.
    FramerGiantPacketsReceived interface{}

    // Aborts received at framer. The type is interface{} with range:
    // 0..4294967295.
    FramerAbortsReceived interface{}
}

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetEntityData() *types.CommonEntityData {
    sideAErrors.EntityData.YFilter = sideAErrors.YFilter
    sideAErrors.EntityData.YangName = "side-a-errors"
    sideAErrors.EntityData.BundleName = "cisco_ios_xr"
    sideAErrors.EntityData.ParentYangName = "srp-statistics"
    sideAErrors.EntityData.SegmentPath = "side-a-errors"
    sideAErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sideAErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sideAErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sideAErrors.EntityData.Children = types.NewOrderedMap()
    sideAErrors.EntityData.Leafs = types.NewOrderedMap()
    sideAErrors.EntityData.Leafs.Append("error-packets-received", types.YLeaf{"ErrorPacketsReceived", sideAErrors.ErrorPacketsReceived})
    sideAErrors.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", sideAErrors.CrcErrors})
    sideAErrors.EntityData.Leafs.Append("input-insufficient-resource-events", types.YLeaf{"InputInsufficientResourceEvents", sideAErrors.InputInsufficientResourceEvents})
    sideAErrors.EntityData.Leafs.Append("mac-aborts-received", types.YLeaf{"MacAbortsReceived", sideAErrors.MacAbortsReceived})
    sideAErrors.EntityData.Leafs.Append("mac-runt-packets-received", types.YLeaf{"MacRuntPacketsReceived", sideAErrors.MacRuntPacketsReceived})
    sideAErrors.EntityData.Leafs.Append("mac-giant-packets-received", types.YLeaf{"MacGiantPacketsReceived", sideAErrors.MacGiantPacketsReceived})
    sideAErrors.EntityData.Leafs.Append("framer-runt-packets-received", types.YLeaf{"FramerRuntPacketsReceived", sideAErrors.FramerRuntPacketsReceived})
    sideAErrors.EntityData.Leafs.Append("framer-giant-packets-received", types.YLeaf{"FramerGiantPacketsReceived", sideAErrors.FramerGiantPacketsReceived})
    sideAErrors.EntityData.Leafs.Append("framer-aborts-received", types.YLeaf{"FramerAbortsReceived", sideAErrors.FramerAbortsReceived})

    sideAErrors.EntityData.YListKeys = []string {}

    return &(sideAErrors.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors
// Errors for side B interface
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error packets received. The type is interface{} with range: 0..4294967295.
    ErrorPacketsReceived interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input insufficient resources events. The type is interface{} with range:
    // 0..4294967295.
    InputInsufficientResourceEvents interface{}

    // Aborts received at MAC/RAC. The type is interface{} with range:
    // 0..4294967295.
    MacAbortsReceived interface{}

    // Too small packets received at MAC/RAC. The type is interface{} with range:
    // 0..4294967295.
    MacRuntPacketsReceived interface{}

    // Too large packets received at MAC/RAC. The type is interface{} with range:
    // 0..4294967295.
    MacGiantPacketsReceived interface{}

    // Too small packets received at framer. The type is interface{} with range:
    // 0..4294967295.
    FramerRuntPacketsReceived interface{}

    // Too large packets received at framer. The type is interface{} with range:
    // 0..4294967295.
    FramerGiantPacketsReceived interface{}

    // Aborts received at framer. The type is interface{} with range:
    // 0..4294967295.
    FramerAbortsReceived interface{}
}

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetEntityData() *types.CommonEntityData {
    sideBErrors.EntityData.YFilter = sideBErrors.YFilter
    sideBErrors.EntityData.YangName = "side-b-errors"
    sideBErrors.EntityData.BundleName = "cisco_ios_xr"
    sideBErrors.EntityData.ParentYangName = "srp-statistics"
    sideBErrors.EntityData.SegmentPath = "side-b-errors"
    sideBErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sideBErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sideBErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sideBErrors.EntityData.Children = types.NewOrderedMap()
    sideBErrors.EntityData.Leafs = types.NewOrderedMap()
    sideBErrors.EntityData.Leafs.Append("error-packets-received", types.YLeaf{"ErrorPacketsReceived", sideBErrors.ErrorPacketsReceived})
    sideBErrors.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", sideBErrors.CrcErrors})
    sideBErrors.EntityData.Leafs.Append("input-insufficient-resource-events", types.YLeaf{"InputInsufficientResourceEvents", sideBErrors.InputInsufficientResourceEvents})
    sideBErrors.EntityData.Leafs.Append("mac-aborts-received", types.YLeaf{"MacAbortsReceived", sideBErrors.MacAbortsReceived})
    sideBErrors.EntityData.Leafs.Append("mac-runt-packets-received", types.YLeaf{"MacRuntPacketsReceived", sideBErrors.MacRuntPacketsReceived})
    sideBErrors.EntityData.Leafs.Append("mac-giant-packets-received", types.YLeaf{"MacGiantPacketsReceived", sideBErrors.MacGiantPacketsReceived})
    sideBErrors.EntityData.Leafs.Append("framer-runt-packets-received", types.YLeaf{"FramerRuntPacketsReceived", sideBErrors.FramerRuntPacketsReceived})
    sideBErrors.EntityData.Leafs.Append("framer-giant-packets-received", types.YLeaf{"FramerGiantPacketsReceived", sideBErrors.FramerGiantPacketsReceived})
    sideBErrors.EntityData.Leafs.Append("framer-aborts-received", types.YLeaf{"FramerAbortsReceived", sideBErrors.FramerAbortsReceived})

    sideBErrors.EntityData.YListKeys = []string {}

    return &(sideBErrors.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation
// Tunnel interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel source name. The type is string.
    SourceName interface{}

    // Tunnel source IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceIpv4Address interface{}

    // Tunnel destination IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationIpv4Address interface{}

    // Tunnel protocol/transport. The type is string.
    TunnelType interface{}

    // GRE tunnel key. The type is interface{} with range: 0..4294967295.
    Key interface{}

    // GRE tunnel TTL. The type is interface{} with range: 0..4294967295.
    Ttl interface{}
}

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetEntityData() *types.CommonEntityData {
    tunnelInformation.EntityData.YFilter = tunnelInformation.YFilter
    tunnelInformation.EntityData.YangName = "tunnel-information"
    tunnelInformation.EntityData.BundleName = "cisco_ios_xr"
    tunnelInformation.EntityData.ParentYangName = "interface-type-information"
    tunnelInformation.EntityData.SegmentPath = "tunnel-information"
    tunnelInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelInformation.EntityData.Children = types.NewOrderedMap()
    tunnelInformation.EntityData.Leafs = types.NewOrderedMap()
    tunnelInformation.EntityData.Leafs.Append("source-name", types.YLeaf{"SourceName", tunnelInformation.SourceName})
    tunnelInformation.EntityData.Leafs.Append("source-ipv4-address", types.YLeaf{"SourceIpv4Address", tunnelInformation.SourceIpv4Address})
    tunnelInformation.EntityData.Leafs.Append("destination-ipv4-address", types.YLeaf{"DestinationIpv4Address", tunnelInformation.DestinationIpv4Address})
    tunnelInformation.EntityData.Leafs.Append("tunnel-type", types.YLeaf{"TunnelType", tunnelInformation.TunnelType})
    tunnelInformation.EntityData.Leafs.Append("key", types.YLeaf{"Key", tunnelInformation.Key})
    tunnelInformation.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", tunnelInformation.Ttl})

    tunnelInformation.EntityData.YListKeys = []string {}

    return &(tunnelInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation
// Bundle interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of bundle members and their properties. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member.
    Member []*Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member
}

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetEntityData() *types.CommonEntityData {
    bundleInformation.EntityData.YFilter = bundleInformation.YFilter
    bundleInformation.EntityData.YangName = "bundle-information"
    bundleInformation.EntityData.BundleName = "cisco_ios_xr"
    bundleInformation.EntityData.ParentYangName = "interface-type-information"
    bundleInformation.EntityData.SegmentPath = "bundle-information"
    bundleInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInformation.EntityData.Children = types.NewOrderedMap()
    bundleInformation.EntityData.Children.Append("member", types.YChild{"Member", nil})
    for i := range bundleInformation.Member {
        bundleInformation.EntityData.Children.Append(types.GetSegmentPath(bundleInformation.Member[i]), types.YChild{"Member", bundleInformation.Member[i]})
    }
    bundleInformation.EntityData.Leafs = types.NewOrderedMap()

    bundleInformation.EntityData.YListKeys = []string {}

    return &(bundleInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member
// List of bundle members and their properties
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Member's interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // The priority of this member. The type is interface{} with range: 0..65535.
    PortPriority interface{}

    // Member's link number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Member's underlying link ID. The type is interface{} with range: 0..65535.
    UnderlyingLinkId interface{}

    // Member's link order number. The type is interface{} with range: 0..65535.
    LinkOrderNumber interface{}

    // Location of member. The type is interface{} with range: 0..4294967295.
    IccpNode interface{}

    // Bandwidth of this member (kbps). The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    Bandwidth interface{}

    // Boolean indicating LACP enabled or not. The type is string.
    LacpEnabled interface{}

    // Member's type (local/foreign). The type is BmdMemberTypeEnum.
    MemberType interface{}

    // Member's (short form) name. The type is string.
    MemberName interface{}

    // Counters data about member link.
    Counters Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters

    // Lacp data about member link.
    LinkData Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData

    // Mux state machine data.
    MemberMuxData Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData

    // MAC address of this member (deprecated).
    MacAddress Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress
}

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetEntityData() *types.CommonEntityData {
    member.EntityData.YFilter = member.YFilter
    member.EntityData.YangName = "member"
    member.EntityData.BundleName = "cisco_ios_xr"
    member.EntityData.ParentYangName = "bundle-information"
    member.EntityData.SegmentPath = "member"
    member.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    member.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    member.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    member.EntityData.Children = types.NewOrderedMap()
    member.EntityData.Children.Append("counters", types.YChild{"Counters", &member.Counters})
    member.EntityData.Children.Append("link-data", types.YChild{"LinkData", &member.LinkData})
    member.EntityData.Children.Append("member-mux-data", types.YChild{"MemberMuxData", &member.MemberMuxData})
    member.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &member.MacAddress})
    member.EntityData.Leafs = types.NewOrderedMap()
    member.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", member.InterfaceName})
    member.EntityData.Leafs.Append("port-priority", types.YLeaf{"PortPriority", member.PortPriority})
    member.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", member.PortNumber})
    member.EntityData.Leafs.Append("underlying-link-id", types.YLeaf{"UnderlyingLinkId", member.UnderlyingLinkId})
    member.EntityData.Leafs.Append("link-order-number", types.YLeaf{"LinkOrderNumber", member.LinkOrderNumber})
    member.EntityData.Leafs.Append("iccp-node", types.YLeaf{"IccpNode", member.IccpNode})
    member.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", member.Bandwidth})
    member.EntityData.Leafs.Append("lacp-enabled", types.YLeaf{"LacpEnabled", member.LacpEnabled})
    member.EntityData.Leafs.Append("member-type", types.YLeaf{"MemberType", member.MemberType})
    member.EntityData.Leafs.Append("member-name", types.YLeaf{"MemberName", member.MemberName})

    member.EntityData.YListKeys = []string {}

    return &(member.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters
// Counters data about member link
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LACPDUs received. The type is interface{} with range: 0..4294967295.
    LacpdUsReceived interface{}

    // LACPDUs transmitted. The type is interface{} with range: 0..4294967295.
    LacpdUsTransmitted interface{}

    // Marker packets received. The type is interface{} with range: 0..4294967295.
    MarkerPacketsReceived interface{}

    // Marker response packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    MarkerResponsesTransmitted interface{}

    // Illegal and unknown packets received. The type is interface{} with range:
    // 0..4294967295.
    IllegalPacketsReceived interface{}

    // LACPDUs received that exceed the rate limit. The type is interface{} with
    // range: 0..4294967295.
    ExcessLacpdUsReceived interface{}

    // Marker packets received that exceed the rate limit. The type is interface{}
    // with range: 0..4294967295.
    ExcessMarkerPacketsReceived interface{}

    // State flag set to Defaulted. The type is interface{} with range:
    // 0..4294967295.
    Defaulted interface{}

    // State flag set to Expired. The type is interface{} with range:
    // 0..4294967295.
    Expired interface{}

    // Last time counters cleared (s) (deprecated). The type is interface{} with
    // range: 0..4294967295.
    LastClearedSec interface{}

    // Last time counters cleared (nsec) (deprecated). The type is interface{}
    // with range: 0..4294967295.
    LastClearedNsec interface{}
}

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "member"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Leafs = types.NewOrderedMap()
    counters.EntityData.Leafs.Append("lacpd-us-received", types.YLeaf{"LacpdUsReceived", counters.LacpdUsReceived})
    counters.EntityData.Leafs.Append("lacpd-us-transmitted", types.YLeaf{"LacpdUsTransmitted", counters.LacpdUsTransmitted})
    counters.EntityData.Leafs.Append("marker-packets-received", types.YLeaf{"MarkerPacketsReceived", counters.MarkerPacketsReceived})
    counters.EntityData.Leafs.Append("marker-responses-transmitted", types.YLeaf{"MarkerResponsesTransmitted", counters.MarkerResponsesTransmitted})
    counters.EntityData.Leafs.Append("illegal-packets-received", types.YLeaf{"IllegalPacketsReceived", counters.IllegalPacketsReceived})
    counters.EntityData.Leafs.Append("excess-lacpd-us-received", types.YLeaf{"ExcessLacpdUsReceived", counters.ExcessLacpdUsReceived})
    counters.EntityData.Leafs.Append("excess-marker-packets-received", types.YLeaf{"ExcessMarkerPacketsReceived", counters.ExcessMarkerPacketsReceived})
    counters.EntityData.Leafs.Append("defaulted", types.YLeaf{"Defaulted", counters.Defaulted})
    counters.EntityData.Leafs.Append("expired", types.YLeaf{"Expired", counters.Expired})
    counters.EntityData.Leafs.Append("last-cleared-sec", types.YLeaf{"LastClearedSec", counters.LastClearedSec})
    counters.EntityData.Leafs.Append("last-cleared-nsec", types.YLeaf{"LastClearedNsec", counters.LastClearedNsec})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData
// Lacp data about member link
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Member's interface handle. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // System priority of actor system. The type is interface{} with range:
    // 0..65535.
    ActorSystemPriority interface{}

    // MAC Address of the actor system. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    ActorSystemMacAddress interface{}

    // Operational key for this port. The type is interface{} with range:
    // 0..65535.
    ActorOperationalKey interface{}

    // System priority of partner system. The type is interface{} with range:
    // 0..65535.
    PartnerSystemPriority interface{}

    // MAC Address used to identify the partner system. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PartnerSystemMacAddress interface{}

    // Operational key for partner port. The type is interface{} with range:
    // 0..65535.
    PartnerOperationalKey interface{}

    // MIB ifindex of selected bundle. The type is interface{} with range:
    // 0..4294967295.
    SelectedAggregatorId interface{}

    // MIB ifindex of attached bundle. The type is interface{} with range:
    // 0..4294967295.
    AttachedAggregatorId interface{}

    // Port number of this port. The type is interface{} with range: 0..65535.
    ActorPortId interface{}

    // Priority of this port. The type is interface{} with range: 0..65535.
    ActorPortPriority interface{}

    // Port number of the partner's port. The type is interface{} with range:
    // 0..65535.
    PartnerPortId interface{}

    // Priority of the partner's port. The type is interface{} with range:
    // 0..65535.
    PartnerPortPriority interface{}

    // LACP state of this port. The type is interface{} with range: 0..255.
    ActorPortState interface{}

    // LACP state of the partner's port. The type is interface{} with range:
    // 0..255.
    PartnerPortState interface{}
}

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xr"
    linkData.EntityData.ParentYangName = "member"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkData.EntityData.Children = types.NewOrderedMap()
    linkData.EntityData.Leafs = types.NewOrderedMap()
    linkData.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", linkData.InterfaceHandle})
    linkData.EntityData.Leafs.Append("actor-system-priority", types.YLeaf{"ActorSystemPriority", linkData.ActorSystemPriority})
    linkData.EntityData.Leafs.Append("actor-system-mac-address", types.YLeaf{"ActorSystemMacAddress", linkData.ActorSystemMacAddress})
    linkData.EntityData.Leafs.Append("actor-operational-key", types.YLeaf{"ActorOperationalKey", linkData.ActorOperationalKey})
    linkData.EntityData.Leafs.Append("partner-system-priority", types.YLeaf{"PartnerSystemPriority", linkData.PartnerSystemPriority})
    linkData.EntityData.Leafs.Append("partner-system-mac-address", types.YLeaf{"PartnerSystemMacAddress", linkData.PartnerSystemMacAddress})
    linkData.EntityData.Leafs.Append("partner-operational-key", types.YLeaf{"PartnerOperationalKey", linkData.PartnerOperationalKey})
    linkData.EntityData.Leafs.Append("selected-aggregator-id", types.YLeaf{"SelectedAggregatorId", linkData.SelectedAggregatorId})
    linkData.EntityData.Leafs.Append("attached-aggregator-id", types.YLeaf{"AttachedAggregatorId", linkData.AttachedAggregatorId})
    linkData.EntityData.Leafs.Append("actor-port-id", types.YLeaf{"ActorPortId", linkData.ActorPortId})
    linkData.EntityData.Leafs.Append("actor-port-priority", types.YLeaf{"ActorPortPriority", linkData.ActorPortPriority})
    linkData.EntityData.Leafs.Append("partner-port-id", types.YLeaf{"PartnerPortId", linkData.PartnerPortId})
    linkData.EntityData.Leafs.Append("partner-port-priority", types.YLeaf{"PartnerPortPriority", linkData.PartnerPortPriority})
    linkData.EntityData.Leafs.Append("actor-port-state", types.YLeaf{"ActorPortState", linkData.ActorPortState})
    linkData.EntityData.Leafs.Append("partner-port-state", types.YLeaf{"PartnerPortState", linkData.PartnerPortState})

    linkData.EntityData.YListKeys = []string {}

    return &(linkData.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData
// Mux state machine data
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current state of this bundle member. The type is BmMuxstate.
    MuxState interface{}

    // Internal value indicating if an error occurred trying to put a link into
    // the desired state. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Reason for last Mux state change. The type is BmMbrStateReason.
    MemberMuxStateReason interface{}

    // Current internal state of this bundle member. The type is BmdMemberState.
    MemberState interface{}

    // Reason for last Mux state change (Deprecated). The type is BmMuxreason.
    MuxStateReason interface{}

    // Data regarding the reason for last Mux state change.
    MemberMuxStateReasonData Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData
}

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetEntityData() *types.CommonEntityData {
    memberMuxData.EntityData.YFilter = memberMuxData.YFilter
    memberMuxData.EntityData.YangName = "member-mux-data"
    memberMuxData.EntityData.BundleName = "cisco_ios_xr"
    memberMuxData.EntityData.ParentYangName = "member"
    memberMuxData.EntityData.SegmentPath = "member-mux-data"
    memberMuxData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberMuxData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberMuxData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberMuxData.EntityData.Children = types.NewOrderedMap()
    memberMuxData.EntityData.Children.Append("member-mux-state-reason-data", types.YChild{"MemberMuxStateReasonData", &memberMuxData.MemberMuxStateReasonData})
    memberMuxData.EntityData.Leafs = types.NewOrderedMap()
    memberMuxData.EntityData.Leafs.Append("mux-state", types.YLeaf{"MuxState", memberMuxData.MuxState})
    memberMuxData.EntityData.Leafs.Append("error", types.YLeaf{"Error", memberMuxData.Error})
    memberMuxData.EntityData.Leafs.Append("member-mux-state-reason", types.YLeaf{"MemberMuxStateReason", memberMuxData.MemberMuxStateReason})
    memberMuxData.EntityData.Leafs.Append("member-state", types.YLeaf{"MemberState", memberMuxData.MemberState})
    memberMuxData.EntityData.Leafs.Append("mux-state-reason", types.YLeaf{"MuxStateReason", memberMuxData.MuxStateReason})

    memberMuxData.EntityData.YListKeys = []string {}

    return &(memberMuxData.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData
// Data regarding the reason for last Mux state
// change
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The item the reason applies to. The type is BmStateReasonTarget.
    ReasonType interface{}

    // The severity of the reason. The type is BmSeverity.
    Severity interface{}
}

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetEntityData() *types.CommonEntityData {
    memberMuxStateReasonData.EntityData.YFilter = memberMuxStateReasonData.YFilter
    memberMuxStateReasonData.EntityData.YangName = "member-mux-state-reason-data"
    memberMuxStateReasonData.EntityData.BundleName = "cisco_ios_xr"
    memberMuxStateReasonData.EntityData.ParentYangName = "member-mux-data"
    memberMuxStateReasonData.EntityData.SegmentPath = "member-mux-state-reason-data"
    memberMuxStateReasonData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberMuxStateReasonData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberMuxStateReasonData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberMuxStateReasonData.EntityData.Children = types.NewOrderedMap()
    memberMuxStateReasonData.EntityData.Leafs = types.NewOrderedMap()
    memberMuxStateReasonData.EntityData.Leafs.Append("reason-type", types.YLeaf{"ReasonType", memberMuxStateReasonData.ReasonType})
    memberMuxStateReasonData.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", memberMuxStateReasonData.Severity})

    memberMuxStateReasonData.EntityData.YListKeys = []string {}

    return &(memberMuxStateReasonData.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress
// MAC address of this member (deprecated)
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "member"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", macAddress.Address})

    macAddress.EntityData.YListKeys = []string {}

    return &(macAddress.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation
// Serial interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeslots separated by : or - from 1 to 31. : indicates individual timeslot
    // and - represents a range. E.g. 1-3:5 represents timeslots 1, 2, 3, and 5.
    // The type is string.
    Timeslots interface{}
}

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetEntityData() *types.CommonEntityData {
    serialInformation.EntityData.YFilter = serialInformation.YFilter
    serialInformation.EntityData.YangName = "serial-information"
    serialInformation.EntityData.BundleName = "cisco_ios_xr"
    serialInformation.EntityData.ParentYangName = "interface-type-information"
    serialInformation.EntityData.SegmentPath = "serial-information"
    serialInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serialInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serialInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serialInformation.EntityData.Children = types.NewOrderedMap()
    serialInformation.EntityData.Leafs = types.NewOrderedMap()
    serialInformation.EntityData.Leafs.Append("timeslots", types.YLeaf{"Timeslots", serialInformation.Timeslots})

    serialInformation.EntityData.YListKeys = []string {}

    return &(serialInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation
// SONET POS interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // APS state. The type is SonetApsEt.
    ApsState interface{}
}

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetEntityData() *types.CommonEntityData {
    sonetPosInformation.EntityData.YFilter = sonetPosInformation.YFilter
    sonetPosInformation.EntityData.YangName = "sonet-pos-information"
    sonetPosInformation.EntityData.BundleName = "cisco_ios_xr"
    sonetPosInformation.EntityData.ParentYangName = "interface-type-information"
    sonetPosInformation.EntityData.SegmentPath = "sonet-pos-information"
    sonetPosInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sonetPosInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sonetPosInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sonetPosInformation.EntityData.Children = types.NewOrderedMap()
    sonetPosInformation.EntityData.Leafs = types.NewOrderedMap()
    sonetPosInformation.EntityData.Leafs.Append("aps-state", types.YLeaf{"ApsState", sonetPosInformation.ApsState})

    sonetPosInformation.EntityData.YListKeys = []string {}

    return &(sonetPosInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation
// Tunnel GRE interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel source name. The type is string.
    SourceName interface{}

    // Tunnel destination IP address length. The type is interface{} with range:
    // 0..255.
    DestinationIpAddressLength interface{}

    // GRE tunnel TOS. The type is interface{} with range: 0..4294967295.
    TunnelTos interface{}

    // GRE tunnel TTL. The type is interface{} with range: 0..4294967295.
    TunnelTtl interface{}

    // Key value for GRE Packet. The type is interface{} with range:
    // 0..4294967295.
    Key interface{}

    // Keepalive period in seconds. The type is interface{} with range: 0..65535.
    // Units are second.
    KeepalivePeriod interface{}

    // Keepalive retry. The type is interface{} with range: 0..255.
    KeepaliveMaximumRetry interface{}

    // Tunnel GRE Mode. The type is TunnelGreMode.
    TunnelMode interface{}

    // Tunnel Mode Direction. The type is TunlIpModeDir.
    TunnelModeDirection interface{}

    // Keepalive State. The type is TunnelKaDfState.
    KeepaliveState interface{}

    // DF Bit State. The type is TunnelKaDfState.
    DfBitState interface{}

    // Key Config State. The type is TunnelKeyState.
    KeyBitState interface{}

    // Tunnel source IP address.
    SourceIpAddress Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress

    // Tunnel destination IP address.
    DestinationIpAddress Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress
}

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetEntityData() *types.CommonEntityData {
    tunnelGreInformation.EntityData.YFilter = tunnelGreInformation.YFilter
    tunnelGreInformation.EntityData.YangName = "tunnel-gre-information"
    tunnelGreInformation.EntityData.BundleName = "cisco_ios_xr"
    tunnelGreInformation.EntityData.ParentYangName = "interface-type-information"
    tunnelGreInformation.EntityData.SegmentPath = "tunnel-gre-information"
    tunnelGreInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelGreInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelGreInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelGreInformation.EntityData.Children = types.NewOrderedMap()
    tunnelGreInformation.EntityData.Children.Append("source-ip-address", types.YChild{"SourceIpAddress", &tunnelGreInformation.SourceIpAddress})
    tunnelGreInformation.EntityData.Children.Append("destination-ip-address", types.YChild{"DestinationIpAddress", &tunnelGreInformation.DestinationIpAddress})
    tunnelGreInformation.EntityData.Leafs = types.NewOrderedMap()
    tunnelGreInformation.EntityData.Leafs.Append("source-name", types.YLeaf{"SourceName", tunnelGreInformation.SourceName})
    tunnelGreInformation.EntityData.Leafs.Append("destination-ip-address-length", types.YLeaf{"DestinationIpAddressLength", tunnelGreInformation.DestinationIpAddressLength})
    tunnelGreInformation.EntityData.Leafs.Append("tunnel-tos", types.YLeaf{"TunnelTos", tunnelGreInformation.TunnelTos})
    tunnelGreInformation.EntityData.Leafs.Append("tunnel-ttl", types.YLeaf{"TunnelTtl", tunnelGreInformation.TunnelTtl})
    tunnelGreInformation.EntityData.Leafs.Append("key", types.YLeaf{"Key", tunnelGreInformation.Key})
    tunnelGreInformation.EntityData.Leafs.Append("keepalive-period", types.YLeaf{"KeepalivePeriod", tunnelGreInformation.KeepalivePeriod})
    tunnelGreInformation.EntityData.Leafs.Append("keepalive-maximum-retry", types.YLeaf{"KeepaliveMaximumRetry", tunnelGreInformation.KeepaliveMaximumRetry})
    tunnelGreInformation.EntityData.Leafs.Append("tunnel-mode", types.YLeaf{"TunnelMode", tunnelGreInformation.TunnelMode})
    tunnelGreInformation.EntityData.Leafs.Append("tunnel-mode-direction", types.YLeaf{"TunnelModeDirection", tunnelGreInformation.TunnelModeDirection})
    tunnelGreInformation.EntityData.Leafs.Append("keepalive-state", types.YLeaf{"KeepaliveState", tunnelGreInformation.KeepaliveState})
    tunnelGreInformation.EntityData.Leafs.Append("df-bit-state", types.YLeaf{"DfBitState", tunnelGreInformation.DfBitState})
    tunnelGreInformation.EntityData.Leafs.Append("key-bit-state", types.YLeaf{"KeyBitState", tunnelGreInformation.KeyBitState})

    tunnelGreInformation.EntityData.YListKeys = []string {}

    return &(tunnelGreInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress
// Tunnel source IP address
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is TunlPfiAfId.
    Afi interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetEntityData() *types.CommonEntityData {
    sourceIpAddress.EntityData.YFilter = sourceIpAddress.YFilter
    sourceIpAddress.EntityData.YangName = "source-ip-address"
    sourceIpAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceIpAddress.EntityData.ParentYangName = "tunnel-gre-information"
    sourceIpAddress.EntityData.SegmentPath = "source-ip-address"
    sourceIpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceIpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceIpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceIpAddress.EntityData.Children = types.NewOrderedMap()
    sourceIpAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceIpAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceIpAddress.Afi})
    sourceIpAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceIpAddress.Ipv4})
    sourceIpAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceIpAddress.Ipv6})

    sourceIpAddress.EntityData.YListKeys = []string {}

    return &(sourceIpAddress.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress
// Tunnel destination IP address
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is TunlPfiAfId.
    Afi interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetEntityData() *types.CommonEntityData {
    destinationIpAddress.EntityData.YFilter = destinationIpAddress.YFilter
    destinationIpAddress.EntityData.YangName = "destination-ip-address"
    destinationIpAddress.EntityData.BundleName = "cisco_ios_xr"
    destinationIpAddress.EntityData.ParentYangName = "tunnel-gre-information"
    destinationIpAddress.EntityData.SegmentPath = "destination-ip-address"
    destinationIpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationIpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationIpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationIpAddress.EntityData.Children = types.NewOrderedMap()
    destinationIpAddress.EntityData.Leafs = types.NewOrderedMap()
    destinationIpAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", destinationIpAddress.Afi})
    destinationIpAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", destinationIpAddress.Ipv4})
    destinationIpAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", destinationIpAddress.Ipv6})

    destinationIpAddress.EntityData.YListKeys = []string {}

    return &(destinationIpAddress.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation
// PseudowireHeadEnd interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface list Name. The type is string.
    InterfaceListName interface{}

    // L2 Overhead. The type is interface{} with range: 0..4294967295.
    L2Overhead interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}
}

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetEntityData() *types.CommonEntityData {
    pseudowireHeadEndInformation.EntityData.YFilter = pseudowireHeadEndInformation.YFilter
    pseudowireHeadEndInformation.EntityData.YangName = "pseudowire-head-end-information"
    pseudowireHeadEndInformation.EntityData.BundleName = "cisco_ios_xr"
    pseudowireHeadEndInformation.EntityData.ParentYangName = "interface-type-information"
    pseudowireHeadEndInformation.EntityData.SegmentPath = "pseudowire-head-end-information"
    pseudowireHeadEndInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireHeadEndInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireHeadEndInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireHeadEndInformation.EntityData.Children = types.NewOrderedMap()
    pseudowireHeadEndInformation.EntityData.Leafs = types.NewOrderedMap()
    pseudowireHeadEndInformation.EntityData.Leafs.Append("interface-list-name", types.YLeaf{"InterfaceListName", pseudowireHeadEndInformation.InterfaceListName})
    pseudowireHeadEndInformation.EntityData.Leafs.Append("l2-overhead", types.YLeaf{"L2Overhead", pseudowireHeadEndInformation.L2Overhead})
    pseudowireHeadEndInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", pseudowireHeadEndInformation.InternalLabel})

    pseudowireHeadEndInformation.EntityData.YListKeys = []string {}

    return &(pseudowireHeadEndInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation
// Cem interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeslots separated by : or - from 1 to 32. : indicates individual timeslot
    // and - represents a range. E.g. 1-3:5 represents timeslots 1, 2, 3, and 5.
    // The type is string.
    Timeslots interface{}

    // Payload size in bytes configured on CEM interface. The type is interface{}
    // with range: 0..65535. Units are byte.
    Payload interface{}

    // Dejitter buffer length configuredin milliseconds. The type is interface{}
    // with range: 0..65535. Units are millisecond.
    DejitterBuffer interface{}

    // If framing is TRUE then the CEM  interface is structure aware ; otherwise
    // it is structure agnostic. The type is interface{} with range:
    // -2147483648..2147483647.
    Framing interface{}
}

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetEntityData() *types.CommonEntityData {
    cemInformation.EntityData.YFilter = cemInformation.YFilter
    cemInformation.EntityData.YangName = "cem-information"
    cemInformation.EntityData.BundleName = "cisco_ios_xr"
    cemInformation.EntityData.ParentYangName = "interface-type-information"
    cemInformation.EntityData.SegmentPath = "cem-information"
    cemInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cemInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cemInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cemInformation.EntityData.Children = types.NewOrderedMap()
    cemInformation.EntityData.Leafs = types.NewOrderedMap()
    cemInformation.EntityData.Leafs.Append("timeslots", types.YLeaf{"Timeslots", cemInformation.Timeslots})
    cemInformation.EntityData.Leafs.Append("payload", types.YLeaf{"Payload", cemInformation.Payload})
    cemInformation.EntityData.Leafs.Append("dejitter-buffer", types.YLeaf{"DejitterBuffer", cemInformation.DejitterBuffer})
    cemInformation.EntityData.Leafs.Append("framing", types.YLeaf{"Framing", cemInformation.Framing})

    cemInformation.EntityData.YListKeys = []string {}

    return &(cemInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation
// GCC interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Derived State. The type is GccDerState.
    DerivedMode interface{}

    // Sec State . The type is GccSecState.
    SecState interface{}
}

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetEntityData() *types.CommonEntityData {
    gccInformation.EntityData.YFilter = gccInformation.YFilter
    gccInformation.EntityData.YangName = "gcc-information"
    gccInformation.EntityData.BundleName = "cisco_ios_xr"
    gccInformation.EntityData.ParentYangName = "interface-type-information"
    gccInformation.EntityData.SegmentPath = "gcc-information"
    gccInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gccInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gccInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gccInformation.EntityData.Children = types.NewOrderedMap()
    gccInformation.EntityData.Leafs = types.NewOrderedMap()
    gccInformation.EntityData.Leafs.Append("derived-mode", types.YLeaf{"DerivedMode", gccInformation.DerivedMode})
    gccInformation.EntityData.Leafs.Append("sec-state", types.YLeaf{"SecState", gccInformation.SecState})

    gccInformation.EntityData.YListKeys = []string {}

    return &(gccInformation.EntityData)
}

// Interfaces_InterfaceXr_Interface_DataRates
// Packet and byte rates
type Interfaces_InterfaceXr_Interface_DataRates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    InputDataRate interface{}

    // Input packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    InputPacketRate interface{}

    // Output data rate in 1000's of bps. The type is interface{} with range:
    // 0..18446744073709551615. Units are bit/s.
    OutputDataRate interface{}

    // Output packets per second. The type is interface{} with range:
    // 0..18446744073709551615. Units are packet/s.
    OutputPacketRate interface{}

    // Peak input data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputDataRate interface{}

    // Peak input packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakInputPacketRate interface{}

    // Peak output data rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputDataRate interface{}

    // Peak output packet rate. The type is interface{} with range:
    // 0..18446744073709551615.
    PeakOutputPacketRate interface{}

    // Bandwidth (in kbps). The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    Bandwidth interface{}

    // Number of 30-sec intervals less one. The type is interface{} with range:
    // 0..4294967295.
    LoadInterval interface{}

    // Output load as fraction of 255. The type is interface{} with range: 0..255.
    OutputLoad interface{}

    // Input load as fraction of 255. The type is interface{} with range: 0..255.
    InputLoad interface{}

    // Reliability coefficient. The type is interface{} with range: 0..255.
    Reliability interface{}
}

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetEntityData() *types.CommonEntityData {
    dataRates.EntityData.YFilter = dataRates.YFilter
    dataRates.EntityData.YangName = "data-rates"
    dataRates.EntityData.BundleName = "cisco_ios_xr"
    dataRates.EntityData.ParentYangName = "interface"
    dataRates.EntityData.SegmentPath = "data-rates"
    dataRates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRates.EntityData.Children = types.NewOrderedMap()
    dataRates.EntityData.Leafs = types.NewOrderedMap()
    dataRates.EntityData.Leafs.Append("input-data-rate", types.YLeaf{"InputDataRate", dataRates.InputDataRate})
    dataRates.EntityData.Leafs.Append("input-packet-rate", types.YLeaf{"InputPacketRate", dataRates.InputPacketRate})
    dataRates.EntityData.Leafs.Append("output-data-rate", types.YLeaf{"OutputDataRate", dataRates.OutputDataRate})
    dataRates.EntityData.Leafs.Append("output-packet-rate", types.YLeaf{"OutputPacketRate", dataRates.OutputPacketRate})
    dataRates.EntityData.Leafs.Append("peak-input-data-rate", types.YLeaf{"PeakInputDataRate", dataRates.PeakInputDataRate})
    dataRates.EntityData.Leafs.Append("peak-input-packet-rate", types.YLeaf{"PeakInputPacketRate", dataRates.PeakInputPacketRate})
    dataRates.EntityData.Leafs.Append("peak-output-data-rate", types.YLeaf{"PeakOutputDataRate", dataRates.PeakOutputDataRate})
    dataRates.EntityData.Leafs.Append("peak-output-packet-rate", types.YLeaf{"PeakOutputPacketRate", dataRates.PeakOutputPacketRate})
    dataRates.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", dataRates.Bandwidth})
    dataRates.EntityData.Leafs.Append("load-interval", types.YLeaf{"LoadInterval", dataRates.LoadInterval})
    dataRates.EntityData.Leafs.Append("output-load", types.YLeaf{"OutputLoad", dataRates.OutputLoad})
    dataRates.EntityData.Leafs.Append("input-load", types.YLeaf{"InputLoad", dataRates.InputLoad})
    dataRates.EntityData.Leafs.Append("reliability", types.YLeaf{"Reliability", dataRates.Reliability})

    dataRates.EntityData.YListKeys = []string {}

    return &(dataRates.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceStatistics
// Packet, byte and error counters
type Interfaces_InterfaceXr_Interface_InterfaceStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // StatsType. The type is ImCmdStatsEnum.
    StatsType interface{}

    // Packet, byte and all error counters.
    FullInterfaceStats Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats

    // Packet, byte and selected error counters.
    BasicInterfaceStats Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats
}

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetEntityData() *types.CommonEntityData {
    interfaceStatistics.EntityData.YFilter = interfaceStatistics.YFilter
    interfaceStatistics.EntityData.YangName = "interface-statistics"
    interfaceStatistics.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatistics.EntityData.ParentYangName = "interface"
    interfaceStatistics.EntityData.SegmentPath = "interface-statistics"
    interfaceStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatistics.EntityData.Children = types.NewOrderedMap()
    interfaceStatistics.EntityData.Children.Append("full-interface-stats", types.YChild{"FullInterfaceStats", &interfaceStatistics.FullInterfaceStats})
    interfaceStatistics.EntityData.Children.Append("basic-interface-stats", types.YChild{"BasicInterfaceStats", &interfaceStatistics.BasicInterfaceStats})
    interfaceStatistics.EntityData.Leafs = types.NewOrderedMap()
    interfaceStatistics.EntityData.Leafs.Append("stats-type", types.YLeaf{"StatsType", interfaceStatistics.StatsType})

    interfaceStatistics.EntityData.YListKeys = []string {}

    return &(interfaceStatistics.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats
// Packet, byte and all error counters
type Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Multicast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsReceived interface{}

    // Broadcast packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsReceived interface{}

    // Multicast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    MulticastPacketsSent interface{}

    // Broadcast packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    BroadcastPacketsSent interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Received runt packets. The type is interface{} with range: 0..4294967295.
    RuntPacketsReceived interface{}

    // Received giant packets. The type is interface{} with range: 0..4294967295.
    GiantPacketsReceived interface{}

    // Received throttled packets. The type is interface{} with range:
    // 0..4294967295.
    ThrottledPacketsReceived interface{}

    // Received parity packets. The type is interface{} with range: 0..4294967295.
    ParityPacketsReceived interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Input CRC errors. The type is interface{} with range: 0..4294967295.
    CrcErrors interface{}

    // Input overruns. The type is interface{} with range: 0..4294967295.
    InputOverruns interface{}

    // Framing-errors received. The type is interface{} with range: 0..4294967295.
    FramingErrorsReceived interface{}

    // Input ignored packets. The type is interface{} with range: 0..4294967295.
    InputIgnoredPackets interface{}

    // Input aborts. The type is interface{} with range: 0..4294967295.
    InputAborts interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Output underruns. The type is interface{} with range: 0..4294967295.
    OutputUnderruns interface{}

    // Output buffer failures. The type is interface{} with range: 0..4294967295.
    OutputBufferFailures interface{}

    // Output buffers swapped out. The type is interface{} with range:
    // 0..4294967295.
    OutputBuffersSwappedOut interface{}

    // Applique. The type is interface{} with range: 0..4294967295.
    Applique interface{}

    // Number of board resets. The type is interface{} with range: 0..4294967295.
    Resets interface{}

    // Carrier transitions. The type is interface{} with range: 0..4294967295.
    CarrierTransitions interface{}

    // Availability bit mask. The type is interface{} with range: 0..4294967295.
    AvailabilityFlag interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetEntityData() *types.CommonEntityData {
    fullInterfaceStats.EntityData.YFilter = fullInterfaceStats.YFilter
    fullInterfaceStats.EntityData.YangName = "full-interface-stats"
    fullInterfaceStats.EntityData.BundleName = "cisco_ios_xr"
    fullInterfaceStats.EntityData.ParentYangName = "interface-statistics"
    fullInterfaceStats.EntityData.SegmentPath = "full-interface-stats"
    fullInterfaceStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fullInterfaceStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fullInterfaceStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fullInterfaceStats.EntityData.Children = types.NewOrderedMap()
    fullInterfaceStats.EntityData.Leafs = types.NewOrderedMap()
    fullInterfaceStats.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", fullInterfaceStats.PacketsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", fullInterfaceStats.BytesReceived})
    fullInterfaceStats.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", fullInterfaceStats.PacketsSent})
    fullInterfaceStats.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", fullInterfaceStats.BytesSent})
    fullInterfaceStats.EntityData.Leafs.Append("multicast-packets-received", types.YLeaf{"MulticastPacketsReceived", fullInterfaceStats.MulticastPacketsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("broadcast-packets-received", types.YLeaf{"BroadcastPacketsReceived", fullInterfaceStats.BroadcastPacketsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("multicast-packets-sent", types.YLeaf{"MulticastPacketsSent", fullInterfaceStats.MulticastPacketsSent})
    fullInterfaceStats.EntityData.Leafs.Append("broadcast-packets-sent", types.YLeaf{"BroadcastPacketsSent", fullInterfaceStats.BroadcastPacketsSent})
    fullInterfaceStats.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", fullInterfaceStats.OutputDrops})
    fullInterfaceStats.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", fullInterfaceStats.OutputQueueDrops})
    fullInterfaceStats.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", fullInterfaceStats.InputDrops})
    fullInterfaceStats.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", fullInterfaceStats.InputQueueDrops})
    fullInterfaceStats.EntityData.Leafs.Append("runt-packets-received", types.YLeaf{"RuntPacketsReceived", fullInterfaceStats.RuntPacketsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("giant-packets-received", types.YLeaf{"GiantPacketsReceived", fullInterfaceStats.GiantPacketsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("throttled-packets-received", types.YLeaf{"ThrottledPacketsReceived", fullInterfaceStats.ThrottledPacketsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("parity-packets-received", types.YLeaf{"ParityPacketsReceived", fullInterfaceStats.ParityPacketsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", fullInterfaceStats.UnknownProtocolPacketsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", fullInterfaceStats.InputErrors})
    fullInterfaceStats.EntityData.Leafs.Append("crc-errors", types.YLeaf{"CrcErrors", fullInterfaceStats.CrcErrors})
    fullInterfaceStats.EntityData.Leafs.Append("input-overruns", types.YLeaf{"InputOverruns", fullInterfaceStats.InputOverruns})
    fullInterfaceStats.EntityData.Leafs.Append("framing-errors-received", types.YLeaf{"FramingErrorsReceived", fullInterfaceStats.FramingErrorsReceived})
    fullInterfaceStats.EntityData.Leafs.Append("input-ignored-packets", types.YLeaf{"InputIgnoredPackets", fullInterfaceStats.InputIgnoredPackets})
    fullInterfaceStats.EntityData.Leafs.Append("input-aborts", types.YLeaf{"InputAborts", fullInterfaceStats.InputAborts})
    fullInterfaceStats.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", fullInterfaceStats.OutputErrors})
    fullInterfaceStats.EntityData.Leafs.Append("output-underruns", types.YLeaf{"OutputUnderruns", fullInterfaceStats.OutputUnderruns})
    fullInterfaceStats.EntityData.Leafs.Append("output-buffer-failures", types.YLeaf{"OutputBufferFailures", fullInterfaceStats.OutputBufferFailures})
    fullInterfaceStats.EntityData.Leafs.Append("output-buffers-swapped-out", types.YLeaf{"OutputBuffersSwappedOut", fullInterfaceStats.OutputBuffersSwappedOut})
    fullInterfaceStats.EntityData.Leafs.Append("applique", types.YLeaf{"Applique", fullInterfaceStats.Applique})
    fullInterfaceStats.EntityData.Leafs.Append("resets", types.YLeaf{"Resets", fullInterfaceStats.Resets})
    fullInterfaceStats.EntityData.Leafs.Append("carrier-transitions", types.YLeaf{"CarrierTransitions", fullInterfaceStats.CarrierTransitions})
    fullInterfaceStats.EntityData.Leafs.Append("availability-flag", types.YLeaf{"AvailabilityFlag", fullInterfaceStats.AvailabilityFlag})
    fullInterfaceStats.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", fullInterfaceStats.LastDataTime})
    fullInterfaceStats.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", fullInterfaceStats.SecondsSinceLastClearCounters})
    fullInterfaceStats.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", fullInterfaceStats.LastDiscontinuityTime})
    fullInterfaceStats.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", fullInterfaceStats.SecondsSincePacketReceived})
    fullInterfaceStats.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", fullInterfaceStats.SecondsSincePacketSent})

    fullInterfaceStats.EntityData.YListKeys = []string {}

    return &(fullInterfaceStats.EntityData)
}

// Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats
// Packet, byte and selected error counters
type Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BytesReceived interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Total input drops. The type is interface{} with range: 0..4294967295.
    InputDrops interface{}

    // Input queue drops. The type is interface{} with range: 0..4294967295.
    InputQueueDrops interface{}

    // Total input errors. The type is interface{} with range: 0..4294967295.
    InputErrors interface{}

    // Unknown protocol packets received. The type is interface{} with range:
    // 0..4294967295.
    UnknownProtocolPacketsReceived interface{}

    // Total output drops. The type is interface{} with range: 0..4294967295.
    OutputDrops interface{}

    // Output queue drops. The type is interface{} with range: 0..4294967295.
    OutputQueueDrops interface{}

    // Total output errors. The type is interface{} with range: 0..4294967295.
    OutputErrors interface{}

    // Time when counters were last written (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    LastDataTime interface{}

    // Number of seconds since last clear counters. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsSinceLastClearCounters interface{}

    // SysUpTime when counters were last reset (in seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastDiscontinuityTime interface{}

    // Seconds since packet received. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketReceived interface{}

    // Seconds since packet sent. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SecondsSincePacketSent interface{}
}

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetEntityData() *types.CommonEntityData {
    basicInterfaceStats.EntityData.YFilter = basicInterfaceStats.YFilter
    basicInterfaceStats.EntityData.YangName = "basic-interface-stats"
    basicInterfaceStats.EntityData.BundleName = "cisco_ios_xr"
    basicInterfaceStats.EntityData.ParentYangName = "interface-statistics"
    basicInterfaceStats.EntityData.SegmentPath = "basic-interface-stats"
    basicInterfaceStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInterfaceStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInterfaceStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInterfaceStats.EntityData.Children = types.NewOrderedMap()
    basicInterfaceStats.EntityData.Leafs = types.NewOrderedMap()
    basicInterfaceStats.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", basicInterfaceStats.PacketsReceived})
    basicInterfaceStats.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", basicInterfaceStats.BytesReceived})
    basicInterfaceStats.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", basicInterfaceStats.PacketsSent})
    basicInterfaceStats.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", basicInterfaceStats.BytesSent})
    basicInterfaceStats.EntityData.Leafs.Append("input-drops", types.YLeaf{"InputDrops", basicInterfaceStats.InputDrops})
    basicInterfaceStats.EntityData.Leafs.Append("input-queue-drops", types.YLeaf{"InputQueueDrops", basicInterfaceStats.InputQueueDrops})
    basicInterfaceStats.EntityData.Leafs.Append("input-errors", types.YLeaf{"InputErrors", basicInterfaceStats.InputErrors})
    basicInterfaceStats.EntityData.Leafs.Append("unknown-protocol-packets-received", types.YLeaf{"UnknownProtocolPacketsReceived", basicInterfaceStats.UnknownProtocolPacketsReceived})
    basicInterfaceStats.EntityData.Leafs.Append("output-drops", types.YLeaf{"OutputDrops", basicInterfaceStats.OutputDrops})
    basicInterfaceStats.EntityData.Leafs.Append("output-queue-drops", types.YLeaf{"OutputQueueDrops", basicInterfaceStats.OutputQueueDrops})
    basicInterfaceStats.EntityData.Leafs.Append("output-errors", types.YLeaf{"OutputErrors", basicInterfaceStats.OutputErrors})
    basicInterfaceStats.EntityData.Leafs.Append("last-data-time", types.YLeaf{"LastDataTime", basicInterfaceStats.LastDataTime})
    basicInterfaceStats.EntityData.Leafs.Append("seconds-since-last-clear-counters", types.YLeaf{"SecondsSinceLastClearCounters", basicInterfaceStats.SecondsSinceLastClearCounters})
    basicInterfaceStats.EntityData.Leafs.Append("last-discontinuity-time", types.YLeaf{"LastDiscontinuityTime", basicInterfaceStats.LastDiscontinuityTime})
    basicInterfaceStats.EntityData.Leafs.Append("seconds-since-packet-received", types.YLeaf{"SecondsSincePacketReceived", basicInterfaceStats.SecondsSincePacketReceived})
    basicInterfaceStats.EntityData.Leafs.Append("seconds-since-packet-sent", types.YLeaf{"SecondsSincePacketSent", basicInterfaceStats.SecondsSincePacketSent})

    basicInterfaceStats.EntityData.YListKeys = []string {}

    return &(basicInterfaceStats.EntityData)
}

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics
// L2 Protocol Statistics
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats type value. The type is interface{} with range: 0..4294967295.
    StatsType interface{}

    // Bag contents. The type is StatsTypeContents.
    Contents interface{}

    // Identifier.
    StatsId Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId

    // Block Array. The type is slice of
    // Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray.
    BlockArray []*Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray

    // Element Array. The type is slice of
    // Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray.
    ElementArray []*Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray
}

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetEntityData() *types.CommonEntityData {
    l2InterfaceStatistics.EntityData.YFilter = l2InterfaceStatistics.YFilter
    l2InterfaceStatistics.EntityData.YangName = "l2-interface-statistics"
    l2InterfaceStatistics.EntityData.BundleName = "cisco_ios_xr"
    l2InterfaceStatistics.EntityData.ParentYangName = "interface"
    l2InterfaceStatistics.EntityData.SegmentPath = "l2-interface-statistics"
    l2InterfaceStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2InterfaceStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2InterfaceStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2InterfaceStatistics.EntityData.Children = types.NewOrderedMap()
    l2InterfaceStatistics.EntityData.Children.Append("stats-id", types.YChild{"StatsId", &l2InterfaceStatistics.StatsId})
    l2InterfaceStatistics.EntityData.Children.Append("block-array", types.YChild{"BlockArray", nil})
    for i := range l2InterfaceStatistics.BlockArray {
        l2InterfaceStatistics.EntityData.Children.Append(types.GetSegmentPath(l2InterfaceStatistics.BlockArray[i]), types.YChild{"BlockArray", l2InterfaceStatistics.BlockArray[i]})
    }
    l2InterfaceStatistics.EntityData.Children.Append("element-array", types.YChild{"ElementArray", nil})
    for i := range l2InterfaceStatistics.ElementArray {
        l2InterfaceStatistics.EntityData.Children.Append(types.GetSegmentPath(l2InterfaceStatistics.ElementArray[i]), types.YChild{"ElementArray", l2InterfaceStatistics.ElementArray[i]})
    }
    l2InterfaceStatistics.EntityData.Leafs = types.NewOrderedMap()
    l2InterfaceStatistics.EntityData.Leafs.Append("stats-type", types.YLeaf{"StatsType", l2InterfaceStatistics.StatsType})
    l2InterfaceStatistics.EntityData.Leafs.Append("contents", types.YLeaf{"Contents", l2InterfaceStatistics.Contents})

    l2InterfaceStatistics.EntityData.YListKeys = []string {}

    return &(l2InterfaceStatistics.EntityData)
}

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId
// Identifier
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // id type. The type is StatsId.
    IdType interface{}

    // Unused. The type is interface{} with range: 0..4294967295.
    Unused interface{}

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceHandle interface{}

    // Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Feature ID. The type is interface{} with range: 0..4294967295.
    FeatureId interface{}

    // ID. The type is interface{} with range: 0..4294967295.
    Id interface{}
}

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetEntityData() *types.CommonEntityData {
    statsId.EntityData.YFilter = statsId.YFilter
    statsId.EntityData.YangName = "stats-id"
    statsId.EntityData.BundleName = "cisco_ios_xr"
    statsId.EntityData.ParentYangName = "l2-interface-statistics"
    statsId.EntityData.SegmentPath = "stats-id"
    statsId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsId.EntityData.Children = types.NewOrderedMap()
    statsId.EntityData.Leafs = types.NewOrderedMap()
    statsId.EntityData.Leafs.Append("id-type", types.YLeaf{"IdType", statsId.IdType})
    statsId.EntityData.Leafs.Append("unused", types.YLeaf{"Unused", statsId.Unused})
    statsId.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", statsId.InterfaceHandle})
    statsId.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statsId.NodeId})
    statsId.EntityData.Leafs.Append("feature-id", types.YLeaf{"FeatureId", statsId.FeatureId})
    statsId.EntityData.Leafs.Append("id", types.YLeaf{"Id", statsId.Id})

    statsId.EntityData.YListKeys = []string {}

    return &(statsId.EntityData)
}

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray
// Block Array
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is StatsCounter.
    Type interface{}

    // count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // data. The type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Data interface{}
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetEntityData() *types.CommonEntityData {
    blockArray.EntityData.YFilter = blockArray.YFilter
    blockArray.EntityData.YangName = "block-array"
    blockArray.EntityData.BundleName = "cisco_ios_xr"
    blockArray.EntityData.ParentYangName = "l2-interface-statistics"
    blockArray.EntityData.SegmentPath = "block-array"
    blockArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockArray.EntityData.Children = types.NewOrderedMap()
    blockArray.EntityData.Leafs = types.NewOrderedMap()
    blockArray.EntityData.Leafs.Append("type", types.YLeaf{"Type", blockArray.Type})
    blockArray.EntityData.Leafs.Append("count", types.YLeaf{"Count", blockArray.Count})
    blockArray.EntityData.Leafs.Append("data", types.YLeaf{"Data", blockArray.Data})

    blockArray.EntityData.YListKeys = []string {}

    return &(blockArray.EntityData)
}

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray
// Element Array
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key. The type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Key interface{}

    // block array. The type is slice of
    // Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray.
    BlockArray []*Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray
}

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetEntityData() *types.CommonEntityData {
    elementArray.EntityData.YFilter = elementArray.YFilter
    elementArray.EntityData.YangName = "element-array"
    elementArray.EntityData.BundleName = "cisco_ios_xr"
    elementArray.EntityData.ParentYangName = "l2-interface-statistics"
    elementArray.EntityData.SegmentPath = "element-array"
    elementArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    elementArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    elementArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    elementArray.EntityData.Children = types.NewOrderedMap()
    elementArray.EntityData.Children.Append("block-array", types.YChild{"BlockArray", nil})
    for i := range elementArray.BlockArray {
        elementArray.EntityData.Children.Append(types.GetSegmentPath(elementArray.BlockArray[i]), types.YChild{"BlockArray", elementArray.BlockArray[i]})
    }
    elementArray.EntityData.Leafs = types.NewOrderedMap()
    elementArray.EntityData.Leafs.Append("key", types.YLeaf{"Key", elementArray.Key})

    elementArray.EntityData.YListKeys = []string {}

    return &(elementArray.EntityData)
}

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray
// block array
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is StatsCounter.
    Type interface{}

    // count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // data. The type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Data interface{}
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetEntityData() *types.CommonEntityData {
    blockArray.EntityData.YFilter = blockArray.YFilter
    blockArray.EntityData.YangName = "block-array"
    blockArray.EntityData.BundleName = "cisco_ios_xr"
    blockArray.EntityData.ParentYangName = "element-array"
    blockArray.EntityData.SegmentPath = "block-array"
    blockArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockArray.EntityData.Children = types.NewOrderedMap()
    blockArray.EntityData.Leafs = types.NewOrderedMap()
    blockArray.EntityData.Leafs.Append("type", types.YLeaf{"Type", blockArray.Type})
    blockArray.EntityData.Leafs.Append("count", types.YLeaf{"Count", blockArray.Count})
    blockArray.EntityData.Leafs.Append("data", types.YLeaf{"Data", blockArray.Data})

    blockArray.EntityData.YListKeys = []string {}

    return &(blockArray.EntityData)
}

// Interfaces_NodeTypeSets
// Node and/or interface type specific view of
// interface summary data
type Interfaces_NodeTypeSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary data for all interfaces on a particular node. The type is slice of
    // Interfaces_NodeTypeSets_NodeTypeSet.
    NodeTypeSet []*Interfaces_NodeTypeSets_NodeTypeSet
}

func (nodeTypeSets *Interfaces_NodeTypeSets) GetEntityData() *types.CommonEntityData {
    nodeTypeSets.EntityData.YFilter = nodeTypeSets.YFilter
    nodeTypeSets.EntityData.YangName = "node-type-sets"
    nodeTypeSets.EntityData.BundleName = "cisco_ios_xr"
    nodeTypeSets.EntityData.ParentYangName = "interfaces"
    nodeTypeSets.EntityData.SegmentPath = "node-type-sets"
    nodeTypeSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeTypeSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeTypeSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeTypeSets.EntityData.Children = types.NewOrderedMap()
    nodeTypeSets.EntityData.Children.Append("node-type-set", types.YChild{"NodeTypeSet", nil})
    for i := range nodeTypeSets.NodeTypeSet {
        nodeTypeSets.EntityData.Children.Append(types.GetSegmentPath(nodeTypeSets.NodeTypeSet[i]), types.YChild{"NodeTypeSet", nodeTypeSets.NodeTypeSet[i]})
    }
    nodeTypeSets.EntityData.Leafs = types.NewOrderedMap()

    nodeTypeSets.EntityData.YListKeys = []string {}

    return &(nodeTypeSets.EntityData)
}

// Interfaces_NodeTypeSets_NodeTypeSet
// Summary data for all interfaces on a particular
// node
type Interfaces_NodeTypeSets_NodeTypeSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The location to filter on. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // The interface type to filter on. The type is InterfaceTypeSet.
    TypeSetName interface{}

    // Interface summary information.
    InterfaceSummary Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary
}

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetEntityData() *types.CommonEntityData {
    nodeTypeSet.EntityData.YFilter = nodeTypeSet.YFilter
    nodeTypeSet.EntityData.YangName = "node-type-set"
    nodeTypeSet.EntityData.BundleName = "cisco_ios_xr"
    nodeTypeSet.EntityData.ParentYangName = "node-type-sets"
    nodeTypeSet.EntityData.SegmentPath = "node-type-set"
    nodeTypeSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeTypeSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeTypeSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeTypeSet.EntityData.Children = types.NewOrderedMap()
    nodeTypeSet.EntityData.Children.Append("interface-summary", types.YChild{"InterfaceSummary", &nodeTypeSet.InterfaceSummary})
    nodeTypeSet.EntityData.Leafs = types.NewOrderedMap()
    nodeTypeSet.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", nodeTypeSet.NodeName})
    nodeTypeSet.EntityData.Leafs.Append("type-set-name", types.YLeaf{"TypeSetName", nodeTypeSet.TypeSetName})

    nodeTypeSet.EntityData.YListKeys = []string {}

    return &(nodeTypeSet.EntityData)
}

// Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary
// Interface summary information
type Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Counts for all interfaces.
    InterfaceCounts Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts

    // List of per interface type summary information. The type is slice of
    // Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType.
    InterfaceType []*Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType
}

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "node-type-set"
    interfaceSummary.EntityData.SegmentPath = "interface-summary"
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = types.NewOrderedMap()
    interfaceSummary.EntityData.Children.Append("interface-counts", types.YChild{"InterfaceCounts", &interfaceSummary.InterfaceCounts})
    interfaceSummary.EntityData.Children.Append("interface-type", types.YChild{"InterfaceType", nil})
    for i := range interfaceSummary.InterfaceType {
        interfaceSummary.EntityData.Children.Append(types.GetSegmentPath(interfaceSummary.InterfaceType[i]), types.YChild{"InterfaceType", interfaceSummary.InterfaceType[i]})
    }
    interfaceSummary.EntityData.Leafs = types.NewOrderedMap()

    interfaceSummary.EntityData.YListKeys = []string {}

    return &(interfaceSummary.EntityData)
}

// Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts
// Counts for all interfaces
type Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Number of interfaces in UP state. The type is interface{} with range:
    // 0..4294967295.
    UpInterfaceCount interface{}

    // Number of interfaces in DOWN state. The type is interface{} with range:
    // 0..4294967295.
    DownInterfaceCount interface{}

    // Number of interfaces in an ADMINDOWN state. The type is interface{} with
    // range: 0..4294967295.
    AdminDownInterfaceCount interface{}
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetEntityData() *types.CommonEntityData {
    interfaceCounts.EntityData.YFilter = interfaceCounts.YFilter
    interfaceCounts.EntityData.YangName = "interface-counts"
    interfaceCounts.EntityData.BundleName = "cisco_ios_xr"
    interfaceCounts.EntityData.ParentYangName = "interface-summary"
    interfaceCounts.EntityData.SegmentPath = "interface-counts"
    interfaceCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceCounts.EntityData.Children = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", interfaceCounts.InterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("up-interface-count", types.YLeaf{"UpInterfaceCount", interfaceCounts.UpInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("down-interface-count", types.YLeaf{"DownInterfaceCount", interfaceCounts.DownInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("admin-down-interface-count", types.YLeaf{"AdminDownInterfaceCount", interfaceCounts.AdminDownInterfaceCount})

    interfaceCounts.EntityData.YListKeys = []string {}

    return &(interfaceCounts.EntityData)
}

// Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType
// List of per interface type summary information
type Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the interface type. The type is string.
    InterfaceTypeName interface{}

    // Description of the interface type. The type is string.
    InterfaceTypeDescription interface{}

    // Counts for interfaces of this type.
    InterfaceCounts Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts
}

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetEntityData() *types.CommonEntityData {
    interfaceType.EntityData.YFilter = interfaceType.YFilter
    interfaceType.EntityData.YangName = "interface-type"
    interfaceType.EntityData.BundleName = "cisco_ios_xr"
    interfaceType.EntityData.ParentYangName = "interface-summary"
    interfaceType.EntityData.SegmentPath = "interface-type"
    interfaceType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceType.EntityData.Children = types.NewOrderedMap()
    interfaceType.EntityData.Children.Append("interface-counts", types.YChild{"InterfaceCounts", &interfaceType.InterfaceCounts})
    interfaceType.EntityData.Leafs = types.NewOrderedMap()
    interfaceType.EntityData.Leafs.Append("interface-type-name", types.YLeaf{"InterfaceTypeName", interfaceType.InterfaceTypeName})
    interfaceType.EntityData.Leafs.Append("interface-type-description", types.YLeaf{"InterfaceTypeDescription", interfaceType.InterfaceTypeDescription})

    interfaceType.EntityData.YListKeys = []string {}

    return &(interfaceType.EntityData)
}

// Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts
// Counts for interfaces of this type
type Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Number of interfaces in UP state. The type is interface{} with range:
    // 0..4294967295.
    UpInterfaceCount interface{}

    // Number of interfaces in DOWN state. The type is interface{} with range:
    // 0..4294967295.
    DownInterfaceCount interface{}

    // Number of interfaces in an ADMINDOWN state. The type is interface{} with
    // range: 0..4294967295.
    AdminDownInterfaceCount interface{}
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetEntityData() *types.CommonEntityData {
    interfaceCounts.EntityData.YFilter = interfaceCounts.YFilter
    interfaceCounts.EntityData.YangName = "interface-counts"
    interfaceCounts.EntityData.BundleName = "cisco_ios_xr"
    interfaceCounts.EntityData.ParentYangName = "interface-type"
    interfaceCounts.EntityData.SegmentPath = "interface-counts"
    interfaceCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceCounts.EntityData.Children = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", interfaceCounts.InterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("up-interface-count", types.YLeaf{"UpInterfaceCount", interfaceCounts.UpInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("down-interface-count", types.YLeaf{"DownInterfaceCount", interfaceCounts.DownInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("admin-down-interface-count", types.YLeaf{"AdminDownInterfaceCount", interfaceCounts.AdminDownInterfaceCount})

    interfaceCounts.EntityData.YListKeys = []string {}

    return &(interfaceCounts.EntityData)
}

// Interfaces_InterfaceBriefs
// Brief operational data for interfaces
type Interfaces_InterfaceBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief operational attributes for a particular interface. The type is slice
    // of Interfaces_InterfaceBriefs_InterfaceBrief.
    InterfaceBrief []*Interfaces_InterfaceBriefs_InterfaceBrief
}

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetEntityData() *types.CommonEntityData {
    interfaceBriefs.EntityData.YFilter = interfaceBriefs.YFilter
    interfaceBriefs.EntityData.YangName = "interface-briefs"
    interfaceBriefs.EntityData.BundleName = "cisco_ios_xr"
    interfaceBriefs.EntityData.ParentYangName = "interfaces"
    interfaceBriefs.EntityData.SegmentPath = "interface-briefs"
    interfaceBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceBriefs.EntityData.Children = types.NewOrderedMap()
    interfaceBriefs.EntityData.Children.Append("interface-brief", types.YChild{"InterfaceBrief", nil})
    for i := range interfaceBriefs.InterfaceBrief {
        interfaceBriefs.EntityData.Children.Append(types.GetSegmentPath(interfaceBriefs.InterfaceBrief[i]), types.YChild{"InterfaceBrief", interfaceBriefs.InterfaceBrief[i]})
    }
    interfaceBriefs.EntityData.Leafs = types.NewOrderedMap()

    interfaceBriefs.EntityData.YListKeys = []string {}

    return &(interfaceBriefs.EntityData)
}

// Interfaces_InterfaceBriefs_InterfaceBrief
// Brief operational attributes for a particular
// interface
type Interfaces_InterfaceBriefs_InterfaceBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    Interface interface{}

    // Parent Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    ParentInterface interface{}

    // Interface type. The type is string.
    Type interface{}

    // Operational state. The type is ImStateEnum.
    State interface{}

    // Operational state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    ActualState interface{}

    // Line protocol state. The type is ImStateEnum.
    LineState interface{}

    // Line protocol state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    ActualLineState interface{}

    // Interface encapsulation. The type is string.
    Encapsulation interface{}

    // Interface encapsulation description string. The type is string with length:
    // 0..32.
    EncapsulationTypeString interface{}

    // MTU in bytes. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    Mtu interface{}

    // Subif MTU overhead. The type is interface{} with range: 0..4294967295.
    SubInterfaceMtuOverhead interface{}

    // L2 transport. The type is bool.
    L2Transport interface{}

    // Interface bandwidth (Kb/s). The type is interface{} with range:
    // 0..4294967295.
    Bandwidth interface{}
}

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetEntityData() *types.CommonEntityData {
    interfaceBrief.EntityData.YFilter = interfaceBrief.YFilter
    interfaceBrief.EntityData.YangName = "interface-brief"
    interfaceBrief.EntityData.BundleName = "cisco_ios_xr"
    interfaceBrief.EntityData.ParentYangName = "interface-briefs"
    interfaceBrief.EntityData.SegmentPath = "interface-brief" + types.AddKeyToken(interfaceBrief.InterfaceName, "interface-name")
    interfaceBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceBrief.EntityData.Children = types.NewOrderedMap()
    interfaceBrief.EntityData.Leafs = types.NewOrderedMap()
    interfaceBrief.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceBrief.InterfaceName})
    interfaceBrief.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", interfaceBrief.Interface})
    interfaceBrief.EntityData.Leafs.Append("parent-interface", types.YLeaf{"ParentInterface", interfaceBrief.ParentInterface})
    interfaceBrief.EntityData.Leafs.Append("type", types.YLeaf{"Type", interfaceBrief.Type})
    interfaceBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", interfaceBrief.State})
    interfaceBrief.EntityData.Leafs.Append("actual-state", types.YLeaf{"ActualState", interfaceBrief.ActualState})
    interfaceBrief.EntityData.Leafs.Append("line-state", types.YLeaf{"LineState", interfaceBrief.LineState})
    interfaceBrief.EntityData.Leafs.Append("actual-line-state", types.YLeaf{"ActualLineState", interfaceBrief.ActualLineState})
    interfaceBrief.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", interfaceBrief.Encapsulation})
    interfaceBrief.EntityData.Leafs.Append("encapsulation-type-string", types.YLeaf{"EncapsulationTypeString", interfaceBrief.EncapsulationTypeString})
    interfaceBrief.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", interfaceBrief.Mtu})
    interfaceBrief.EntityData.Leafs.Append("sub-interface-mtu-overhead", types.YLeaf{"SubInterfaceMtuOverhead", interfaceBrief.SubInterfaceMtuOverhead})
    interfaceBrief.EntityData.Leafs.Append("l2-transport", types.YLeaf{"L2Transport", interfaceBrief.L2Transport})
    interfaceBrief.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", interfaceBrief.Bandwidth})

    interfaceBrief.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceBrief.EntityData)
}

// Interfaces_InventorySummary
// Inventory summary information
type Interfaces_InventorySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Counts for all interfaces.
    InterfaceCounts Interfaces_InventorySummary_InterfaceCounts

    // List of per interface type summary information. The type is slice of
    // Interfaces_InventorySummary_InterfaceType.
    InterfaceType []*Interfaces_InventorySummary_InterfaceType
}

func (inventorySummary *Interfaces_InventorySummary) GetEntityData() *types.CommonEntityData {
    inventorySummary.EntityData.YFilter = inventorySummary.YFilter
    inventorySummary.EntityData.YangName = "inventory-summary"
    inventorySummary.EntityData.BundleName = "cisco_ios_xr"
    inventorySummary.EntityData.ParentYangName = "interfaces"
    inventorySummary.EntityData.SegmentPath = "inventory-summary"
    inventorySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventorySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventorySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventorySummary.EntityData.Children = types.NewOrderedMap()
    inventorySummary.EntityData.Children.Append("interface-counts", types.YChild{"InterfaceCounts", &inventorySummary.InterfaceCounts})
    inventorySummary.EntityData.Children.Append("interface-type", types.YChild{"InterfaceType", nil})
    for i := range inventorySummary.InterfaceType {
        inventorySummary.EntityData.Children.Append(types.GetSegmentPath(inventorySummary.InterfaceType[i]), types.YChild{"InterfaceType", inventorySummary.InterfaceType[i]})
    }
    inventorySummary.EntityData.Leafs = types.NewOrderedMap()

    inventorySummary.EntityData.YListKeys = []string {}

    return &(inventorySummary.EntityData)
}

// Interfaces_InventorySummary_InterfaceCounts
// Counts for all interfaces
type Interfaces_InventorySummary_InterfaceCounts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Number of interfaces in UP state. The type is interface{} with range:
    // 0..4294967295.
    UpInterfaceCount interface{}

    // Number of interfaces in DOWN state. The type is interface{} with range:
    // 0..4294967295.
    DownInterfaceCount interface{}

    // Number of interfaces in an ADMINDOWN state. The type is interface{} with
    // range: 0..4294967295.
    AdminDownInterfaceCount interface{}
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetEntityData() *types.CommonEntityData {
    interfaceCounts.EntityData.YFilter = interfaceCounts.YFilter
    interfaceCounts.EntityData.YangName = "interface-counts"
    interfaceCounts.EntityData.BundleName = "cisco_ios_xr"
    interfaceCounts.EntityData.ParentYangName = "inventory-summary"
    interfaceCounts.EntityData.SegmentPath = "interface-counts"
    interfaceCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceCounts.EntityData.Children = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", interfaceCounts.InterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("up-interface-count", types.YLeaf{"UpInterfaceCount", interfaceCounts.UpInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("down-interface-count", types.YLeaf{"DownInterfaceCount", interfaceCounts.DownInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("admin-down-interface-count", types.YLeaf{"AdminDownInterfaceCount", interfaceCounts.AdminDownInterfaceCount})

    interfaceCounts.EntityData.YListKeys = []string {}

    return &(interfaceCounts.EntityData)
}

// Interfaces_InventorySummary_InterfaceType
// List of per interface type summary information
type Interfaces_InventorySummary_InterfaceType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the interface type. The type is string.
    InterfaceTypeName interface{}

    // Description of the interface type. The type is string.
    InterfaceTypeDescription interface{}

    // Counts for interfaces of this type.
    InterfaceCounts Interfaces_InventorySummary_InterfaceType_InterfaceCounts
}

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetEntityData() *types.CommonEntityData {
    interfaceType.EntityData.YFilter = interfaceType.YFilter
    interfaceType.EntityData.YangName = "interface-type"
    interfaceType.EntityData.BundleName = "cisco_ios_xr"
    interfaceType.EntityData.ParentYangName = "inventory-summary"
    interfaceType.EntityData.SegmentPath = "interface-type"
    interfaceType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceType.EntityData.Children = types.NewOrderedMap()
    interfaceType.EntityData.Children.Append("interface-counts", types.YChild{"InterfaceCounts", &interfaceType.InterfaceCounts})
    interfaceType.EntityData.Leafs = types.NewOrderedMap()
    interfaceType.EntityData.Leafs.Append("interface-type-name", types.YLeaf{"InterfaceTypeName", interfaceType.InterfaceTypeName})
    interfaceType.EntityData.Leafs.Append("interface-type-description", types.YLeaf{"InterfaceTypeDescription", interfaceType.InterfaceTypeDescription})

    interfaceType.EntityData.YListKeys = []string {}

    return &(interfaceType.EntityData)
}

// Interfaces_InventorySummary_InterfaceType_InterfaceCounts
// Counts for interfaces of this type
type Interfaces_InventorySummary_InterfaceType_InterfaceCounts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Number of interfaces in UP state. The type is interface{} with range:
    // 0..4294967295.
    UpInterfaceCount interface{}

    // Number of interfaces in DOWN state. The type is interface{} with range:
    // 0..4294967295.
    DownInterfaceCount interface{}

    // Number of interfaces in an ADMINDOWN state. The type is interface{} with
    // range: 0..4294967295.
    AdminDownInterfaceCount interface{}
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetEntityData() *types.CommonEntityData {
    interfaceCounts.EntityData.YFilter = interfaceCounts.YFilter
    interfaceCounts.EntityData.YangName = "interface-counts"
    interfaceCounts.EntityData.BundleName = "cisco_ios_xr"
    interfaceCounts.EntityData.ParentYangName = "interface-type"
    interfaceCounts.EntityData.SegmentPath = "interface-counts"
    interfaceCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceCounts.EntityData.Children = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", interfaceCounts.InterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("up-interface-count", types.YLeaf{"UpInterfaceCount", interfaceCounts.UpInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("down-interface-count", types.YLeaf{"DownInterfaceCount", interfaceCounts.DownInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("admin-down-interface-count", types.YLeaf{"AdminDownInterfaceCount", interfaceCounts.AdminDownInterfaceCount})

    interfaceCounts.EntityData.YListKeys = []string {}

    return &(interfaceCounts.EntityData)
}

// Interfaces_Interfaces
// Descriptions for interfaces
type Interfaces_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Description for a particular interface. The type is slice of
    // Interfaces_Interfaces_Interface.
    Interface []*Interfaces_Interfaces_Interface
}

func (interfaces *Interfaces_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interfaces"
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

// Interfaces_Interfaces_Interface
// Description for a particular interface
type Interfaces_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    Interface interface{}

    // Operational state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    State interface{}

    // Line protocol state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    LineState interface{}

    // Interface description string. The type is string.
    Description interface{}
}

func (self *Interfaces_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
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
    self.EntityData.Leafs.Append("state", types.YLeaf{"State", self.State})
    self.EntityData.Leafs.Append("line-state", types.YLeaf{"LineState", self.LineState})
    self.EntityData.Leafs.Append("description", types.YLeaf{"Description", self.Description})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Interfaces_InterfaceSummary
// Interface summary information
type Interfaces_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Counts for all interfaces.
    InterfaceCounts Interfaces_InterfaceSummary_InterfaceCounts

    // List of per interface type summary information. The type is slice of
    // Interfaces_InterfaceSummary_InterfaceType.
    InterfaceType []*Interfaces_InterfaceSummary_InterfaceType
}

func (interfaceSummary *Interfaces_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "interfaces"
    interfaceSummary.EntityData.SegmentPath = "interface-summary"
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = types.NewOrderedMap()
    interfaceSummary.EntityData.Children.Append("interface-counts", types.YChild{"InterfaceCounts", &interfaceSummary.InterfaceCounts})
    interfaceSummary.EntityData.Children.Append("interface-type", types.YChild{"InterfaceType", nil})
    for i := range interfaceSummary.InterfaceType {
        interfaceSummary.EntityData.Children.Append(types.GetSegmentPath(interfaceSummary.InterfaceType[i]), types.YChild{"InterfaceType", interfaceSummary.InterfaceType[i]})
    }
    interfaceSummary.EntityData.Leafs = types.NewOrderedMap()

    interfaceSummary.EntityData.YListKeys = []string {}

    return &(interfaceSummary.EntityData)
}

// Interfaces_InterfaceSummary_InterfaceCounts
// Counts for all interfaces
type Interfaces_InterfaceSummary_InterfaceCounts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Number of interfaces in UP state. The type is interface{} with range:
    // 0..4294967295.
    UpInterfaceCount interface{}

    // Number of interfaces in DOWN state. The type is interface{} with range:
    // 0..4294967295.
    DownInterfaceCount interface{}

    // Number of interfaces in an ADMINDOWN state. The type is interface{} with
    // range: 0..4294967295.
    AdminDownInterfaceCount interface{}
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetEntityData() *types.CommonEntityData {
    interfaceCounts.EntityData.YFilter = interfaceCounts.YFilter
    interfaceCounts.EntityData.YangName = "interface-counts"
    interfaceCounts.EntityData.BundleName = "cisco_ios_xr"
    interfaceCounts.EntityData.ParentYangName = "interface-summary"
    interfaceCounts.EntityData.SegmentPath = "interface-counts"
    interfaceCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceCounts.EntityData.Children = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", interfaceCounts.InterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("up-interface-count", types.YLeaf{"UpInterfaceCount", interfaceCounts.UpInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("down-interface-count", types.YLeaf{"DownInterfaceCount", interfaceCounts.DownInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("admin-down-interface-count", types.YLeaf{"AdminDownInterfaceCount", interfaceCounts.AdminDownInterfaceCount})

    interfaceCounts.EntityData.YListKeys = []string {}

    return &(interfaceCounts.EntityData)
}

// Interfaces_InterfaceSummary_InterfaceType
// List of per interface type summary information
type Interfaces_InterfaceSummary_InterfaceType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the interface type. The type is string.
    InterfaceTypeName interface{}

    // Description of the interface type. The type is string.
    InterfaceTypeDescription interface{}

    // Counts for interfaces of this type.
    InterfaceCounts Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts
}

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetEntityData() *types.CommonEntityData {
    interfaceType.EntityData.YFilter = interfaceType.YFilter
    interfaceType.EntityData.YangName = "interface-type"
    interfaceType.EntityData.BundleName = "cisco_ios_xr"
    interfaceType.EntityData.ParentYangName = "interface-summary"
    interfaceType.EntityData.SegmentPath = "interface-type"
    interfaceType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceType.EntityData.Children = types.NewOrderedMap()
    interfaceType.EntityData.Children.Append("interface-counts", types.YChild{"InterfaceCounts", &interfaceType.InterfaceCounts})
    interfaceType.EntityData.Leafs = types.NewOrderedMap()
    interfaceType.EntityData.Leafs.Append("interface-type-name", types.YLeaf{"InterfaceTypeName", interfaceType.InterfaceTypeName})
    interfaceType.EntityData.Leafs.Append("interface-type-description", types.YLeaf{"InterfaceTypeDescription", interfaceType.InterfaceTypeDescription})

    interfaceType.EntityData.YListKeys = []string {}

    return &(interfaceType.EntityData)
}

// Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts
// Counts for interfaces of this type
type Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Number of interfaces in UP state. The type is interface{} with range:
    // 0..4294967295.
    UpInterfaceCount interface{}

    // Number of interfaces in DOWN state. The type is interface{} with range:
    // 0..4294967295.
    DownInterfaceCount interface{}

    // Number of interfaces in an ADMINDOWN state. The type is interface{} with
    // range: 0..4294967295.
    AdminDownInterfaceCount interface{}
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetEntityData() *types.CommonEntityData {
    interfaceCounts.EntityData.YFilter = interfaceCounts.YFilter
    interfaceCounts.EntityData.YangName = "interface-counts"
    interfaceCounts.EntityData.BundleName = "cisco_ios_xr"
    interfaceCounts.EntityData.ParentYangName = "interface-type"
    interfaceCounts.EntityData.SegmentPath = "interface-counts"
    interfaceCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceCounts.EntityData.Children = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs = types.NewOrderedMap()
    interfaceCounts.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", interfaceCounts.InterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("up-interface-count", types.YLeaf{"UpInterfaceCount", interfaceCounts.UpInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("down-interface-count", types.YLeaf{"DownInterfaceCount", interfaceCounts.DownInterfaceCount})
    interfaceCounts.EntityData.Leafs.Append("admin-down-interface-count", types.YLeaf{"AdminDownInterfaceCount", interfaceCounts.AdminDownInterfaceCount})

    interfaceCounts.EntityData.YListKeys = []string {}

    return &(interfaceCounts.EntityData)
}

