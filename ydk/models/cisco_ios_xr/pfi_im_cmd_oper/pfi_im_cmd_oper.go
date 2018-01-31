// This module contains a collection of YANG definitions
// for Cisco IOS-XR pfi-im-cmd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   interfaces: Interface operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

    // Maintainance
    GccDerState_maintainance GccDerState = "maintainance"

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

    // Maintainance
    GccSecState_maintainance GccSecState = "maintainance"

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
    BmSeverity_error BmSeverity = "error"
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
    parent types.Entity
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

func (interfaces *Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Interfaces) GetGoName(yname string) string {
    if yname == "interface-xr" { return "InterfaceXr" }
    if yname == "node-type-sets" { return "NodeTypeSets" }
    if yname == "interface-briefs" { return "InterfaceBriefs" }
    if yname == "inventory-summary" { return "InventorySummary" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "interface-summary" { return "InterfaceSummary" }
    return ""
}

func (interfaces *Interfaces) GetSegmentPath() string {
    return "Cisco-IOS-XR-pfi-im-cmd-oper:interfaces"
}

func (interfaces *Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-xr" {
        return &interfaces.InterfaceXr
    }
    if childYangName == "node-type-sets" {
        return &interfaces.NodeTypeSets
    }
    if childYangName == "interface-briefs" {
        return &interfaces.InterfaceBriefs
    }
    if childYangName == "inventory-summary" {
        return &interfaces.InventorySummary
    }
    if childYangName == "interfaces" {
        return &interfaces.Interfaces
    }
    if childYangName == "interface-summary" {
        return &interfaces.InterfaceSummary
    }
    return nil
}

func (interfaces *Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-xr"] = &interfaces.InterfaceXr
    children["node-type-sets"] = &interfaces.NodeTypeSets
    children["interface-briefs"] = &interfaces.InterfaceBriefs
    children["inventory-summary"] = &interfaces.InventorySummary
    children["interfaces"] = &interfaces.Interfaces
    children["interface-summary"] = &interfaces.InterfaceSummary
    return children
}

func (interfaces *Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Interfaces) GetParentYangName() string { return "Cisco-IOS-XR-pfi-im-cmd-oper" }

// Interfaces_InterfaceXr
// Detailed operational data for interfaces and
// configured features
type Interfaces_InterfaceXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed operational data for a particular interface. The type is slice of
    // Interfaces_InterfaceXr_Interface.
    Interface []Interfaces_InterfaceXr_Interface
}

func (interfaceXr *Interfaces_InterfaceXr) GetFilter() yfilter.YFilter { return interfaceXr.YFilter }

func (interfaceXr *Interfaces_InterfaceXr) SetFilter(yf yfilter.YFilter) { interfaceXr.YFilter = yf }

func (interfaceXr *Interfaces_InterfaceXr) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaceXr *Interfaces_InterfaceXr) GetSegmentPath() string {
    return "interface-xr"
}

func (interfaceXr *Interfaces_InterfaceXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaceXr.Interface {
            if interfaceXr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface{}
        interfaceXr.Interface = append(interfaceXr.Interface, child)
        return &interfaceXr.Interface[len(interfaceXr.Interface)-1]
    }
    return nil
}

func (interfaceXr *Interfaces_InterfaceXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceXr.Interface {
        children[interfaceXr.Interface[i].GetSegmentPath()] = &interfaceXr.Interface[i]
    }
    return children
}

func (interfaceXr *Interfaces_InterfaceXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceXr *Interfaces_InterfaceXr) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceXr *Interfaces_InterfaceXr) GetYangName() string { return "interface-xr" }

func (interfaceXr *Interfaces_InterfaceXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceXr *Interfaces_InterfaceXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceXr *Interfaces_InterfaceXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceXr *Interfaces_InterfaceXr) SetParent(parent types.Entity) { interfaceXr.parent = parent }

func (interfaceXr *Interfaces_InterfaceXr) GetParent() types.Entity { return interfaceXr.parent }

func (interfaceXr *Interfaces_InterfaceXr) GetParentYangName() string { return "interfaces" }

// Interfaces_InterfaceXr_Interface
// Detailed operational data for a particular
// interface
type Interfaces_InterfaceXr_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

    // Parent interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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

    // nV Optical Controller Information.
    NvOptical Interfaces_InterfaceXr_Interface_NvOptical
}

func (self *Interfaces_InterfaceXr_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Interfaces_InterfaceXr_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Interfaces_InterfaceXr_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "interface-type" { return "InterfaceType" }
    if yname == "hardware-type-string" { return "HardwareTypeString" }
    if yname == "state" { return "State" }
    if yname == "line-state" { return "LineState" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "encapsulation-type-string" { return "EncapsulationTypeString" }
    if yname == "mtu" { return "Mtu" }
    if yname == "is-l2-transport-enabled" { return "IsL2TransportEnabled" }
    if yname == "state-transition-count" { return "StateTransitionCount" }
    if yname == "last-state-transition-time" { return "LastStateTransitionTime" }
    if yname == "is-dampening-enabled" { return "IsDampeningEnabled" }
    if yname == "speed" { return "Speed" }
    if yname == "crc-length" { return "CrcLength" }
    if yname == "is-scramble-enabled" { return "IsScrambleEnabled" }
    if yname == "duplexity" { return "Duplexity" }
    if yname == "media-type" { return "MediaType" }
    if yname == "link-type" { return "LinkType" }
    if yname == "in-flow-control" { return "InFlowControl" }
    if yname == "out-flow-control" { return "OutFlowControl" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "max-bandwidth" { return "MaxBandwidth" }
    if yname == "keepalive" { return "Keepalive" }
    if yname == "is-l2-looped" { return "IsL2Looped" }
    if yname == "parent-interface-name" { return "ParentInterfaceName" }
    if yname == "loopback-configuration" { return "LoopbackConfiguration" }
    if yname == "description" { return "Description" }
    if yname == "is-maintenance-enabled" { return "IsMaintenanceEnabled" }
    if yname == "is-data-inverted" { return "IsDataInverted" }
    if yname == "transport-mode" { return "TransportMode" }
    if yname == "fast-shutdown" { return "FastShutdown" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "dampening-information" { return "DampeningInformation" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "burned-in-address" { return "BurnedInAddress" }
    if yname == "carrier-delay" { return "CarrierDelay" }
    if yname == "arp-information" { return "ArpInformation" }
    if yname == "ip-information" { return "IpInformation" }
    if yname == "encapsulation-information" { return "EncapsulationInformation" }
    if yname == "interface-type-information" { return "InterfaceTypeInformation" }
    if yname == "data-rates" { return "DataRates" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    if yname == "l2-interface-statistics" { return "L2InterfaceStatistics" }
    if yname == "nv-optical" { return "NvOptical" }
    return ""
}

func (self *Interfaces_InterfaceXr_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Interfaces_InterfaceXr_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dampening-information" {
        return &self.DampeningInformation
    }
    if childYangName == "mac-address" {
        return &self.MacAddress
    }
    if childYangName == "burned-in-address" {
        return &self.BurnedInAddress
    }
    if childYangName == "carrier-delay" {
        return &self.CarrierDelay
    }
    if childYangName == "arp-information" {
        return &self.ArpInformation
    }
    if childYangName == "ip-information" {
        return &self.IpInformation
    }
    if childYangName == "encapsulation-information" {
        return &self.EncapsulationInformation
    }
    if childYangName == "interface-type-information" {
        return &self.InterfaceTypeInformation
    }
    if childYangName == "data-rates" {
        return &self.DataRates
    }
    if childYangName == "interface-statistics" {
        return &self.InterfaceStatistics
    }
    if childYangName == "l2-interface-statistics" {
        return &self.L2InterfaceStatistics
    }
    if childYangName == "nv-optical" {
        return &self.NvOptical
    }
    return nil
}

func (self *Interfaces_InterfaceXr_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dampening-information"] = &self.DampeningInformation
    children["mac-address"] = &self.MacAddress
    children["burned-in-address"] = &self.BurnedInAddress
    children["carrier-delay"] = &self.CarrierDelay
    children["arp-information"] = &self.ArpInformation
    children["ip-information"] = &self.IpInformation
    children["encapsulation-information"] = &self.EncapsulationInformation
    children["interface-type-information"] = &self.InterfaceTypeInformation
    children["data-rates"] = &self.DataRates
    children["interface-statistics"] = &self.InterfaceStatistics
    children["l2-interface-statistics"] = &self.L2InterfaceStatistics
    children["nv-optical"] = &self.NvOptical
    return children
}

func (self *Interfaces_InterfaceXr_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-handle"] = self.InterfaceHandle
    leafs["interface-type"] = self.InterfaceType
    leafs["hardware-type-string"] = self.HardwareTypeString
    leafs["state"] = self.State
    leafs["line-state"] = self.LineState
    leafs["encapsulation"] = self.Encapsulation
    leafs["encapsulation-type-string"] = self.EncapsulationTypeString
    leafs["mtu"] = self.Mtu
    leafs["is-l2-transport-enabled"] = self.IsL2TransportEnabled
    leafs["state-transition-count"] = self.StateTransitionCount
    leafs["last-state-transition-time"] = self.LastStateTransitionTime
    leafs["is-dampening-enabled"] = self.IsDampeningEnabled
    leafs["speed"] = self.Speed
    leafs["crc-length"] = self.CrcLength
    leafs["is-scramble-enabled"] = self.IsScrambleEnabled
    leafs["duplexity"] = self.Duplexity
    leafs["media-type"] = self.MediaType
    leafs["link-type"] = self.LinkType
    leafs["in-flow-control"] = self.InFlowControl
    leafs["out-flow-control"] = self.OutFlowControl
    leafs["bandwidth"] = self.Bandwidth
    leafs["max-bandwidth"] = self.MaxBandwidth
    leafs["keepalive"] = self.Keepalive
    leafs["is-l2-looped"] = self.IsL2Looped
    leafs["parent-interface-name"] = self.ParentInterfaceName
    leafs["loopback-configuration"] = self.LoopbackConfiguration
    leafs["description"] = self.Description
    leafs["is-maintenance-enabled"] = self.IsMaintenanceEnabled
    leafs["is-data-inverted"] = self.IsDataInverted
    leafs["transport-mode"] = self.TransportMode
    leafs["fast-shutdown"] = self.FastShutdown
    leafs["if-index"] = self.IfIndex
    return leafs
}

func (self *Interfaces_InterfaceXr_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Interfaces_InterfaceXr_Interface) GetYangName() string { return "interface" }

func (self *Interfaces_InterfaceXr_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Interfaces_InterfaceXr_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Interfaces_InterfaceXr_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Interfaces_InterfaceXr_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Interfaces_InterfaceXr_Interface) GetParent() types.Entity { return self.parent }

func (self *Interfaces_InterfaceXr_Interface) GetParentYangName() string { return "interface-xr" }

// Interfaces_InterfaceXr_Interface_DampeningInformation
// State dampening information
type Interfaces_InterfaceXr_Interface_DampeningInformation struct {
    parent types.Entity
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

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetFilter() yfilter.YFilter { return dampeningInformation.YFilter }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) SetFilter(yf yfilter.YFilter) { dampeningInformation.YFilter = yf }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetGoName(yname string) string {
    if yname == "penalty" { return "Penalty" }
    if yname == "is-suppressed-enabled" { return "IsSuppressedEnabled" }
    if yname == "seconds-remaining" { return "SecondsRemaining" }
    if yname == "half-life" { return "HalfLife" }
    if yname == "reuse-threshold" { return "ReuseThreshold" }
    if yname == "suppress-threshold" { return "SuppressThreshold" }
    if yname == "maximum-suppress-time" { return "MaximumSuppressTime" }
    if yname == "restart-penalty" { return "RestartPenalty" }
    return ""
}

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetSegmentPath() string {
    return "dampening-information"
}

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["penalty"] = dampeningInformation.Penalty
    leafs["is-suppressed-enabled"] = dampeningInformation.IsSuppressedEnabled
    leafs["seconds-remaining"] = dampeningInformation.SecondsRemaining
    leafs["half-life"] = dampeningInformation.HalfLife
    leafs["reuse-threshold"] = dampeningInformation.ReuseThreshold
    leafs["suppress-threshold"] = dampeningInformation.SuppressThreshold
    leafs["maximum-suppress-time"] = dampeningInformation.MaximumSuppressTime
    leafs["restart-penalty"] = dampeningInformation.RestartPenalty
    return leafs
}

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetBundleName() string { return "cisco_ios_xr" }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetYangName() string { return "dampening-information" }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) SetParent(parent types.Entity) { dampeningInformation.parent = parent }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetParent() types.Entity { return dampeningInformation.parent }

func (dampeningInformation *Interfaces_InterfaceXr_Interface_DampeningInformation) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_MacAddress
// Interface MAC address
type Interfaces_InterfaceXr_Interface_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = macAddress.Address
    return leafs
}

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Interfaces_InterfaceXr_Interface_MacAddress) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_BurnedInAddress
// Interface burned in address
type Interfaces_InterfaceXr_Interface_BurnedInAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetFilter() yfilter.YFilter { return burnedInAddress.YFilter }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) SetFilter(yf yfilter.YFilter) { burnedInAddress.YFilter = yf }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetSegmentPath() string {
    return "burned-in-address"
}

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = burnedInAddress.Address
    return leafs
}

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetBundleName() string { return "cisco_ios_xr" }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetYangName() string { return "burned-in-address" }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) SetParent(parent types.Entity) { burnedInAddress.parent = parent }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetParent() types.Entity { return burnedInAddress.parent }

func (burnedInAddress *Interfaces_InterfaceXr_Interface_BurnedInAddress) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_CarrierDelay
// Carrier Delay
type Interfaces_InterfaceXr_Interface_CarrierDelay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Carrier delay on state up (ms). The type is interface{} with range:
    // 0..4294967295.
    CarrierDelayUp interface{}

    // Carrier delay on state down (ms). The type is interface{} with range:
    // 0..4294967295.
    CarrierDelayDown interface{}
}

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetFilter() yfilter.YFilter { return carrierDelay.YFilter }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) SetFilter(yf yfilter.YFilter) { carrierDelay.YFilter = yf }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetGoName(yname string) string {
    if yname == "carrier-delay-up" { return "CarrierDelayUp" }
    if yname == "carrier-delay-down" { return "CarrierDelayDown" }
    return ""
}

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetSegmentPath() string {
    return "carrier-delay"
}

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["carrier-delay-up"] = carrierDelay.CarrierDelayUp
    leafs["carrier-delay-down"] = carrierDelay.CarrierDelayDown
    return leafs
}

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetBundleName() string { return "cisco_ios_xr" }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetYangName() string { return "carrier-delay" }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) SetParent(parent types.Entity) { carrierDelay.parent = parent }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetParent() types.Entity { return carrierDelay.parent }

func (carrierDelay *Interfaces_InterfaceXr_Interface_CarrierDelay) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_ArpInformation
// Interface ARP type and timeout
type Interfaces_InterfaceXr_Interface_ArpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ARP timeout in seconds. Only valid if 'ARPIsLearningDisabled' is 'false'.
    // The type is interface{} with range: 0..4294967295. Units are second.
    ArpTimeout interface{}

    // ARP type name. The type is string.
    ArpTypeName interface{}

    // Whether the interface has dynamic learning disabled. The type is bool.
    ArpIsLearningDisabled interface{}
}

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetFilter() yfilter.YFilter { return arpInformation.YFilter }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) SetFilter(yf yfilter.YFilter) { arpInformation.YFilter = yf }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetGoName(yname string) string {
    if yname == "arp-timeout" { return "ArpTimeout" }
    if yname == "arp-type-name" { return "ArpTypeName" }
    if yname == "arp-is-learning-disabled" { return "ArpIsLearningDisabled" }
    return ""
}

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetSegmentPath() string {
    return "arp-information"
}

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["arp-timeout"] = arpInformation.ArpTimeout
    leafs["arp-type-name"] = arpInformation.ArpTypeName
    leafs["arp-is-learning-disabled"] = arpInformation.ArpIsLearningDisabled
    return leafs
}

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetYangName() string { return "arp-information" }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) SetParent(parent types.Entity) { arpInformation.parent = parent }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetParent() types.Entity { return arpInformation.parent }

func (arpInformation *Interfaces_InterfaceXr_Interface_ArpInformation) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_IpInformation
// Interface IP address info
type Interfaces_InterfaceXr_Interface_IpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Interface subnet mask length. The type is interface{} with range:
    // 0..4294967295.
    SubnetMaskLength interface{}
}

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetFilter() yfilter.YFilter { return ipInformation.YFilter }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) SetFilter(yf yfilter.YFilter) { ipInformation.YFilter = yf }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "subnet-mask-length" { return "SubnetMaskLength" }
    return ""
}

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetSegmentPath() string {
    return "ip-information"
}

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = ipInformation.IpAddress
    leafs["subnet-mask-length"] = ipInformation.SubnetMaskLength
    return leafs
}

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetYangName() string { return "ip-information" }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) SetParent(parent types.Entity) { ipInformation.parent = parent }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetParent() types.Entity { return ipInformation.parent }

func (ipInformation *Interfaces_InterfaceXr_Interface_IpInformation) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation
// Information specific to the encapsulation
type Interfaces_InterfaceXr_Interface_EncapsulationInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EncapsulationType. The type is ImCmdEncapsEnum.
    EncapsulationType interface{}

    // Frame Relay information.
    FrameRelayInformation Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation

    // VLAN 802.1q information.
    Dot1QInformation Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation

    // PPP information.
    PppInformation Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation
}

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetFilter() yfilter.YFilter { return encapsulationInformation.YFilter }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) SetFilter(yf yfilter.YFilter) { encapsulationInformation.YFilter = yf }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetGoName(yname string) string {
    if yname == "encapsulation-type" { return "EncapsulationType" }
    if yname == "frame-relay-information" { return "FrameRelayInformation" }
    if yname == "dot1q-information" { return "Dot1QInformation" }
    if yname == "ppp-information" { return "PppInformation" }
    return ""
}

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetSegmentPath() string {
    return "encapsulation-information"
}

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frame-relay-information" {
        return &encapsulationInformation.FrameRelayInformation
    }
    if childYangName == "dot1q-information" {
        return &encapsulationInformation.Dot1QInformation
    }
    if childYangName == "ppp-information" {
        return &encapsulationInformation.PppInformation
    }
    return nil
}

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frame-relay-information"] = &encapsulationInformation.FrameRelayInformation
    children["dot1q-information"] = &encapsulationInformation.Dot1QInformation
    children["ppp-information"] = &encapsulationInformation.PppInformation
    return children
}

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation-type"] = encapsulationInformation.EncapsulationType
    return leafs
}

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetBundleName() string { return "cisco_ios_xr" }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetYangName() string { return "encapsulation-information" }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) SetParent(parent types.Entity) { encapsulationInformation.parent = parent }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetParent() types.Entity { return encapsulationInformation.parent }

func (encapsulationInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation
// Frame Relay information
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation struct {
    parent types.Entity
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

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetFilter() yfilter.YFilter { return frameRelayInformation.YFilter }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) SetFilter(yf yfilter.YFilter) { frameRelayInformation.YFilter = yf }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetGoName(yname string) string {
    if yname == "fr-encapsulation-type" { return "FrEncapsulationType" }
    if yname == "lmi-type" { return "LmiType" }
    if yname == "lmidlci" { return "Lmidlci" }
    if yname == "is-nni" { return "IsNni" }
    if yname == "is-dte" { return "IsDte" }
    if yname == "is-lmi-up" { return "IsLmiUp" }
    if yname == "is-lmi-nni-dce-up" { return "IsLmiNniDceUp" }
    if yname == "is-lmi-enabled" { return "IsLmiEnabled" }
    if yname == "enquiries-received" { return "EnquiriesReceived" }
    if yname == "enquiries-sent" { return "EnquiriesSent" }
    if yname == "status-received" { return "StatusReceived" }
    if yname == "status-sent" { return "StatusSent" }
    if yname == "update-status-received" { return "UpdateStatusReceived" }
    if yname == "update-status-sent" { return "UpdateStatusSent" }
    return ""
}

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetSegmentPath() string {
    return "frame-relay-information"
}

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fr-encapsulation-type"] = frameRelayInformation.FrEncapsulationType
    leafs["lmi-type"] = frameRelayInformation.LmiType
    leafs["lmidlci"] = frameRelayInformation.Lmidlci
    leafs["is-nni"] = frameRelayInformation.IsNni
    leafs["is-dte"] = frameRelayInformation.IsDte
    leafs["is-lmi-up"] = frameRelayInformation.IsLmiUp
    leafs["is-lmi-nni-dce-up"] = frameRelayInformation.IsLmiNniDceUp
    leafs["is-lmi-enabled"] = frameRelayInformation.IsLmiEnabled
    leafs["enquiries-received"] = frameRelayInformation.EnquiriesReceived
    leafs["enquiries-sent"] = frameRelayInformation.EnquiriesSent
    leafs["status-received"] = frameRelayInformation.StatusReceived
    leafs["status-sent"] = frameRelayInformation.StatusSent
    leafs["update-status-received"] = frameRelayInformation.UpdateStatusReceived
    leafs["update-status-sent"] = frameRelayInformation.UpdateStatusSent
    return leafs
}

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetBundleName() string { return "cisco_ios_xr" }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetYangName() string { return "frame-relay-information" }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) SetParent(parent types.Entity) { frameRelayInformation.parent = parent }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetParent() types.Entity { return frameRelayInformation.parent }

func (frameRelayInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_FrameRelayInformation) GetParentYangName() string { return "encapsulation-information" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation
// VLAN 802.1q information
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation type and tag stack.
    EncapsulationDetails Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails
}

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetFilter() yfilter.YFilter { return dot1QInformation.YFilter }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) SetFilter(yf yfilter.YFilter) { dot1QInformation.YFilter = yf }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetGoName(yname string) string {
    if yname == "encapsulation-details" { return "EncapsulationDetails" }
    return ""
}

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetSegmentPath() string {
    return "dot1q-information"
}

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encapsulation-details" {
        return &dot1QInformation.EncapsulationDetails
    }
    return nil
}

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encapsulation-details"] = &dot1QInformation.EncapsulationDetails
    return children
}

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetBundleName() string { return "cisco_ios_xr" }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetYangName() string { return "dot1q-information" }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) SetParent(parent types.Entity) { dot1QInformation.parent = parent }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetParent() types.Entity { return dot1QInformation.parent }

func (dot1QInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation) GetParentYangName() string { return "encapsulation-information" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails
// Encapsulation type and tag stack
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails struct {
    parent types.Entity
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
    Dot1AdTag interface{}

    // 802.1ad native tag value. The type is interface{} with range: 0..65535.
    Dot1AdNativeTag interface{}

    // 802.1ad Outer tag value. The type is interface{} with range: 0..65535.
    Dot1AdOuterTag interface{}

    // Stack value.
    Stack Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack

    // Service Instance encapsulation.
    ServiceInstanceDetails Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails

    // 802.1ad 802.1Q stack value.
    Dot1AdDot1QStack Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack
}

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetFilter() yfilter.YFilter { return encapsulationDetails.YFilter }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) SetFilter(yf yfilter.YFilter) { encapsulationDetails.YFilter = yf }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetGoName(yname string) string {
    if yname == "vlan-encapsulation" { return "VlanEncapsulation" }
    if yname == "tag" { return "Tag" }
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "native-tag" { return "NativeTag" }
    if yname == "dot1ad-tag" { return "Dot1AdTag" }
    if yname == "dot1ad-native-tag" { return "Dot1AdNativeTag" }
    if yname == "dot1ad-outer-tag" { return "Dot1AdOuterTag" }
    if yname == "stack" { return "Stack" }
    if yname == "service-instance-details" { return "ServiceInstanceDetails" }
    if yname == "dot1ad-dot1q-stack" { return "Dot1AdDot1QStack" }
    return ""
}

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetSegmentPath() string {
    return "encapsulation-details"
}

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack" {
        return &encapsulationDetails.Stack
    }
    if childYangName == "service-instance-details" {
        return &encapsulationDetails.ServiceInstanceDetails
    }
    if childYangName == "dot1ad-dot1q-stack" {
        return &encapsulationDetails.Dot1AdDot1QStack
    }
    return nil
}

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stack"] = &encapsulationDetails.Stack
    children["service-instance-details"] = &encapsulationDetails.ServiceInstanceDetails
    children["dot1ad-dot1q-stack"] = &encapsulationDetails.Dot1AdDot1QStack
    return children
}

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-encapsulation"] = encapsulationDetails.VlanEncapsulation
    leafs["tag"] = encapsulationDetails.Tag
    leafs["outer-tag"] = encapsulationDetails.OuterTag
    leafs["native-tag"] = encapsulationDetails.NativeTag
    leafs["dot1ad-tag"] = encapsulationDetails.Dot1AdTag
    leafs["dot1ad-native-tag"] = encapsulationDetails.Dot1AdNativeTag
    leafs["dot1ad-outer-tag"] = encapsulationDetails.Dot1AdOuterTag
    return leafs
}

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetBundleName() string { return "cisco_ios_xr" }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetYangName() string { return "encapsulation-details" }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) SetParent(parent types.Entity) { encapsulationDetails.parent = parent }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetParent() types.Entity { return encapsulationDetails.parent }

func (encapsulationDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails) GetParentYangName() string { return "dot1q-information" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack
// Stack value
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetFilter() yfilter.YFilter { return stack.YFilter }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) SetFilter(yf yfilter.YFilter) { stack.YFilter = yf }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetGoName(yname string) string {
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "second-tag" { return "SecondTag" }
    return ""
}

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetSegmentPath() string {
    return "stack"
}

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outer-tag"] = stack.OuterTag
    leafs["second-tag"] = stack.SecondTag
    return leafs
}

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetBundleName() string { return "cisco_ios_xr" }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetYangName() string { return "stack" }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) SetParent(parent types.Entity) { stack.parent = parent }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetParent() types.Entity { return stack.parent }

func (stack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Stack) GetParentYangName() string { return "encapsulation-details" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails
// Service Instance encapsulation
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails struct {
    parent types.Entity
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
    LocalTrafficStack Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack

    // Tags to match on ingress packets. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch.
    TagsToMatch []Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch

    // VLAN tags pushed on egress. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe.
    Pushe []Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe
}

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetFilter() yfilter.YFilter { return serviceInstanceDetails.YFilter }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) SetFilter(yf yfilter.YFilter) { serviceInstanceDetails.YFilter = yf }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetGoName(yname string) string {
    if yname == "payload-ethertype" { return "PayloadEthertype" }
    if yname == "tags-popped" { return "TagsPopped" }
    if yname == "is-exact-match" { return "IsExactMatch" }
    if yname == "is-native-vlan" { return "IsNativeVlan" }
    if yname == "is-native-preserving" { return "IsNativePreserving" }
    if yname == "source-mac-match" { return "SourceMacMatch" }
    if yname == "destination-mac-match" { return "DestinationMacMatch" }
    if yname == "local-traffic-stack" { return "LocalTrafficStack" }
    if yname == "tags-to-match" { return "TagsToMatch" }
    if yname == "pushe" { return "Pushe" }
    return ""
}

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetSegmentPath() string {
    return "service-instance-details"
}

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-traffic-stack" {
        return &serviceInstanceDetails.LocalTrafficStack
    }
    if childYangName == "tags-to-match" {
        for _, c := range serviceInstanceDetails.TagsToMatch {
            if serviceInstanceDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch{}
        serviceInstanceDetails.TagsToMatch = append(serviceInstanceDetails.TagsToMatch, child)
        return &serviceInstanceDetails.TagsToMatch[len(serviceInstanceDetails.TagsToMatch)-1]
    }
    if childYangName == "pushe" {
        for _, c := range serviceInstanceDetails.Pushe {
            if serviceInstanceDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe{}
        serviceInstanceDetails.Pushe = append(serviceInstanceDetails.Pushe, child)
        return &serviceInstanceDetails.Pushe[len(serviceInstanceDetails.Pushe)-1]
    }
    return nil
}

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-traffic-stack"] = &serviceInstanceDetails.LocalTrafficStack
    for i := range serviceInstanceDetails.TagsToMatch {
        children[serviceInstanceDetails.TagsToMatch[i].GetSegmentPath()] = &serviceInstanceDetails.TagsToMatch[i]
    }
    for i := range serviceInstanceDetails.Pushe {
        children[serviceInstanceDetails.Pushe[i].GetSegmentPath()] = &serviceInstanceDetails.Pushe[i]
    }
    return children
}

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["payload-ethertype"] = serviceInstanceDetails.PayloadEthertype
    leafs["tags-popped"] = serviceInstanceDetails.TagsPopped
    leafs["is-exact-match"] = serviceInstanceDetails.IsExactMatch
    leafs["is-native-vlan"] = serviceInstanceDetails.IsNativeVlan
    leafs["is-native-preserving"] = serviceInstanceDetails.IsNativePreserving
    leafs["source-mac-match"] = serviceInstanceDetails.SourceMacMatch
    leafs["destination-mac-match"] = serviceInstanceDetails.DestinationMacMatch
    return leafs
}

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetBundleName() string { return "cisco_ios_xr" }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetYangName() string { return "service-instance-details" }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) SetParent(parent types.Entity) { serviceInstanceDetails.parent = parent }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetParent() types.Entity { return serviceInstanceDetails.parent }

func (serviceInstanceDetails *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails) GetParentYangName() string { return "encapsulation-details" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack
// VLAN tags for locally-sourced traffic
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN tags for locally-sourced traffic. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag.
    LocalTrafficTag []Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
}

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetFilter() yfilter.YFilter { return localTrafficStack.YFilter }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) SetFilter(yf yfilter.YFilter) { localTrafficStack.YFilter = yf }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetGoName(yname string) string {
    if yname == "local-traffic-tag" { return "LocalTrafficTag" }
    return ""
}

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetSegmentPath() string {
    return "local-traffic-stack"
}

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-traffic-tag" {
        for _, c := range localTrafficStack.LocalTrafficTag {
            if localTrafficStack.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag{}
        localTrafficStack.LocalTrafficTag = append(localTrafficStack.LocalTrafficTag, child)
        return &localTrafficStack.LocalTrafficTag[len(localTrafficStack.LocalTrafficTag)-1]
    }
    return nil
}

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localTrafficStack.LocalTrafficTag {
        children[localTrafficStack.LocalTrafficTag[i].GetSegmentPath()] = &localTrafficStack.LocalTrafficTag[i]
    }
    return children
}

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetBundleName() string { return "cisco_ios_xr" }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetYangName() string { return "local-traffic-stack" }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) SetParent(parent types.Entity) { localTrafficStack.parent = parent }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetParent() types.Entity { return localTrafficStack.parent }

func (localTrafficStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack) GetParentYangName() string { return "service-instance-details" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag
// VLAN tags for locally-sourced traffic
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetFilter() yfilter.YFilter { return localTrafficTag.YFilter }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) SetFilter(yf yfilter.YFilter) { localTrafficTag.YFilter = yf }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetSegmentPath() string {
    return "local-traffic-tag"
}

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = localTrafficTag.Ethertype
    leafs["vlan-id"] = localTrafficTag.VlanId
    return leafs
}

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetBundleName() string { return "cisco_ios_xr" }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetYangName() string { return "local-traffic-tag" }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) SetParent(parent types.Entity) { localTrafficTag.parent = parent }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetParent() types.Entity { return localTrafficTag.parent }

func (localTrafficTag *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_LocalTrafficStack_LocalTrafficTag) GetParentYangName() string { return "local-traffic-stack" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch
// Tags to match on ingress packets
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag to match. The type is EfpTagEtype.
    Ethertype interface{}

    // Priority to match. The type is EfpTagPriority.
    Priority interface{}

    // VLAN Ids to match. The type is slice of
    // Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange.
    VlanRange []Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
}

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetFilter() yfilter.YFilter { return tagsToMatch.YFilter }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) SetFilter(yf yfilter.YFilter) { tagsToMatch.YFilter = yf }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "priority" { return "Priority" }
    if yname == "vlan-range" { return "VlanRange" }
    return ""
}

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetSegmentPath() string {
    return "tags-to-match"
}

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vlan-range" {
        for _, c := range tagsToMatch.VlanRange {
            if tagsToMatch.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange{}
        tagsToMatch.VlanRange = append(tagsToMatch.VlanRange, child)
        return &tagsToMatch.VlanRange[len(tagsToMatch.VlanRange)-1]
    }
    return nil
}

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tagsToMatch.VlanRange {
        children[tagsToMatch.VlanRange[i].GetSegmentPath()] = &tagsToMatch.VlanRange[i]
    }
    return children
}

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = tagsToMatch.Ethertype
    leafs["priority"] = tagsToMatch.Priority
    return leafs
}

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetBundleName() string { return "cisco_ios_xr" }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetYangName() string { return "tags-to-match" }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) SetParent(parent types.Entity) { tagsToMatch.parent = parent }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetParent() types.Entity { return tagsToMatch.parent }

func (tagsToMatch *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch) GetParentYangName() string { return "service-instance-details" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange
// VLAN Ids to match
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN ID Low. The type is interface{} with range: 0..65535.
    VlanIdLow interface{}

    // VLAN ID High. The type is interface{} with range: 0..65535.
    VlanIdHigh interface{}
}

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetFilter() yfilter.YFilter { return vlanRange.YFilter }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) SetFilter(yf yfilter.YFilter) { vlanRange.YFilter = yf }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetGoName(yname string) string {
    if yname == "vlan-id-low" { return "VlanIdLow" }
    if yname == "vlan-id-high" { return "VlanIdHigh" }
    return ""
}

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetSegmentPath() string {
    return "vlan-range"
}

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-id-low"] = vlanRange.VlanIdLow
    leafs["vlan-id-high"] = vlanRange.VlanIdHigh
    return leafs
}

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetBundleName() string { return "cisco_ios_xr" }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetYangName() string { return "vlan-range" }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) SetParent(parent types.Entity) { vlanRange.parent = parent }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetParent() types.Entity { return vlanRange.parent }

func (vlanRange *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_TagsToMatch_VlanRange) GetParentYangName() string { return "tags-to-match" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe
// VLAN tags pushed on egress
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethertype of tag. The type is EfpTagEtype.
    Ethertype interface{}

    // VLAN Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetFilter() yfilter.YFilter { return pushe.YFilter }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) SetFilter(yf yfilter.YFilter) { pushe.YFilter = yf }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetGoName(yname string) string {
    if yname == "ethertype" { return "Ethertype" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetSegmentPath() string {
    return "pushe"
}

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethertype"] = pushe.Ethertype
    leafs["vlan-id"] = pushe.VlanId
    return leafs
}

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetBundleName() string { return "cisco_ios_xr" }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetYangName() string { return "pushe" }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) SetParent(parent types.Entity) { pushe.parent = parent }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetParent() types.Entity { return pushe.parent }

func (pushe *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_ServiceInstanceDetails_Pushe) GetParentYangName() string { return "service-instance-details" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack
// 802.1ad 802.1Q stack value
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Outer tag value. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Second tag value. The type is interface{} with range: 0..65535.
    SecondTag interface{}
}

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetFilter() yfilter.YFilter { return dot1AdDot1QStack.YFilter }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) SetFilter(yf yfilter.YFilter) { dot1AdDot1QStack.YFilter = yf }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetGoName(yname string) string {
    if yname == "outer-tag" { return "OuterTag" }
    if yname == "second-tag" { return "SecondTag" }
    return ""
}

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetSegmentPath() string {
    return "dot1ad-dot1q-stack"
}

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["outer-tag"] = dot1AdDot1QStack.OuterTag
    leafs["second-tag"] = dot1AdDot1QStack.SecondTag
    return leafs
}

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetBundleName() string { return "cisco_ios_xr" }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetYangName() string { return "dot1ad-dot1q-stack" }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) SetParent(parent types.Entity) { dot1AdDot1QStack.parent = parent }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetParent() types.Entity { return dot1AdDot1QStack.parent }

func (dot1AdDot1QStack *Interfaces_InterfaceXr_Interface_EncapsulationInformation_Dot1QInformation_EncapsulationDetails_Dot1AdDot1QStack) GetParentYangName() string { return "encapsulation-details" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation
// PPP information
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation struct {
    parent types.Entity
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
    NcpInfoArray []Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray
}

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetFilter() yfilter.YFilter { return pppInformation.YFilter }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) SetFilter(yf yfilter.YFilter) { pppInformation.YFilter = yf }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetGoName(yname string) string {
    if yname == "lcp-state" { return "LcpState" }
    if yname == "is-loopback-detected" { return "IsLoopbackDetected" }
    if yname == "keepalive-period" { return "KeepalivePeriod" }
    if yname == "is-mp-bundle-member" { return "IsMpBundleMember" }
    if yname == "is-multilink-open" { return "IsMultilinkOpen" }
    if yname == "ncp-info-array" { return "NcpInfoArray" }
    return ""
}

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetSegmentPath() string {
    return "ppp-information"
}

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ncp-info-array" {
        for _, c := range pppInformation.NcpInfoArray {
            if pppInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray{}
        pppInformation.NcpInfoArray = append(pppInformation.NcpInfoArray, child)
        return &pppInformation.NcpInfoArray[len(pppInformation.NcpInfoArray)-1]
    }
    return nil
}

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pppInformation.NcpInfoArray {
        children[pppInformation.NcpInfoArray[i].GetSegmentPath()] = &pppInformation.NcpInfoArray[i]
    }
    return children
}

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lcp-state"] = pppInformation.LcpState
    leafs["is-loopback-detected"] = pppInformation.IsLoopbackDetected
    leafs["keepalive-period"] = pppInformation.KeepalivePeriod
    leafs["is-mp-bundle-member"] = pppInformation.IsMpBundleMember
    leafs["is-multilink-open"] = pppInformation.IsMultilinkOpen
    return leafs
}

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetBundleName() string { return "cisco_ios_xr" }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetYangName() string { return "ppp-information" }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) SetParent(parent types.Entity) { pppInformation.parent = parent }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetParent() types.Entity { return pppInformation.parent }

func (pppInformation *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation) GetParentYangName() string { return "encapsulation-information" }

// Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray
// Array of per-NCP data
type Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NCP state value. The type is PppFsmState.
    NcpState interface{}

    // NCP state identifier. The type is NcpIdent.
    NcpIdentifier interface{}
}

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetFilter() yfilter.YFilter { return ncpInfoArray.YFilter }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) SetFilter(yf yfilter.YFilter) { ncpInfoArray.YFilter = yf }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetGoName(yname string) string {
    if yname == "ncp-state" { return "NcpState" }
    if yname == "ncp-identifier" { return "NcpIdentifier" }
    return ""
}

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetSegmentPath() string {
    return "ncp-info-array"
}

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ncp-state"] = ncpInfoArray.NcpState
    leafs["ncp-identifier"] = ncpInfoArray.NcpIdentifier
    return leafs
}

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetBundleName() string { return "cisco_ios_xr" }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetYangName() string { return "ncp-info-array" }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) SetParent(parent types.Entity) { ncpInfoArray.parent = parent }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetParent() types.Entity { return ncpInfoArray.parent }

func (ncpInfoArray *Interfaces_InterfaceXr_Interface_EncapsulationInformation_PppInformation_NcpInfoArray) GetParentYangName() string { return "ppp-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation
// Information specific to the interface type
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation struct {
    parent types.Entity
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

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetFilter() yfilter.YFilter { return interfaceTypeInformation.YFilter }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) SetFilter(yf yfilter.YFilter) { interfaceTypeInformation.YFilter = yf }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetGoName(yname string) string {
    if yname == "interface-type-info" { return "InterfaceTypeInfo" }
    if yname == "srp-information" { return "SrpInformation" }
    if yname == "tunnel-information" { return "TunnelInformation" }
    if yname == "bundle-information" { return "BundleInformation" }
    if yname == "serial-information" { return "SerialInformation" }
    if yname == "sonet-pos-information" { return "SonetPosInformation" }
    if yname == "tunnel-gre-information" { return "TunnelGreInformation" }
    if yname == "pseudowire-head-end-information" { return "PseudowireHeadEndInformation" }
    if yname == "cem-information" { return "CemInformation" }
    if yname == "gcc-information" { return "GccInformation" }
    return ""
}

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetSegmentPath() string {
    return "interface-type-information"
}

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srp-information" {
        return &interfaceTypeInformation.SrpInformation
    }
    if childYangName == "tunnel-information" {
        return &interfaceTypeInformation.TunnelInformation
    }
    if childYangName == "bundle-information" {
        return &interfaceTypeInformation.BundleInformation
    }
    if childYangName == "serial-information" {
        return &interfaceTypeInformation.SerialInformation
    }
    if childYangName == "sonet-pos-information" {
        return &interfaceTypeInformation.SonetPosInformation
    }
    if childYangName == "tunnel-gre-information" {
        return &interfaceTypeInformation.TunnelGreInformation
    }
    if childYangName == "pseudowire-head-end-information" {
        return &interfaceTypeInformation.PseudowireHeadEndInformation
    }
    if childYangName == "cem-information" {
        return &interfaceTypeInformation.CemInformation
    }
    if childYangName == "gcc-information" {
        return &interfaceTypeInformation.GccInformation
    }
    return nil
}

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["srp-information"] = &interfaceTypeInformation.SrpInformation
    children["tunnel-information"] = &interfaceTypeInformation.TunnelInformation
    children["bundle-information"] = &interfaceTypeInformation.BundleInformation
    children["serial-information"] = &interfaceTypeInformation.SerialInformation
    children["sonet-pos-information"] = &interfaceTypeInformation.SonetPosInformation
    children["tunnel-gre-information"] = &interfaceTypeInformation.TunnelGreInformation
    children["pseudowire-head-end-information"] = &interfaceTypeInformation.PseudowireHeadEndInformation
    children["cem-information"] = &interfaceTypeInformation.CemInformation
    children["gcc-information"] = &interfaceTypeInformation.GccInformation
    return children
}

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-type-info"] = interfaceTypeInformation.InterfaceTypeInfo
    return leafs
}

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetYangName() string { return "interface-type-information" }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) SetParent(parent types.Entity) { interfaceTypeInformation.parent = parent }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetParent() types.Entity { return interfaceTypeInformation.parent }

func (interfaceTypeInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation
// SRP interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRP-specific data.
    SrpInformation Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation

    // SRP-specific packet and byte counters.
    SrpStatistics Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetFilter() yfilter.YFilter { return srpInformation.YFilter }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) SetFilter(yf yfilter.YFilter) { srpInformation.YFilter = yf }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetGoName(yname string) string {
    if yname == "srp-information" { return "SrpInformation" }
    if yname == "srp-statistics" { return "SrpStatistics" }
    return ""
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetSegmentPath() string {
    return "srp-information"
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srp-information" {
        return &srpInformation.SrpInformation
    }
    if childYangName == "srp-statistics" {
        return &srpInformation.SrpStatistics
    }
    return nil
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["srp-information"] = &srpInformation.SrpInformation
    children["srp-statistics"] = &srpInformation.SrpStatistics
    return children
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetYangName() string { return "srp-information" }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) SetParent(parent types.Entity) { srpInformation.parent = parent }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetParent() types.Entity { return srpInformation.parent }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation
// SRP-specific data
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation struct {
    parent types.Entity
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

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetFilter() yfilter.YFilter { return srpInformation.YFilter }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) SetFilter(yf yfilter.YFilter) { srpInformation.YFilter = yf }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetGoName(yname string) string {
    if yname == "ips-info" { return "IpsInfo" }
    if yname == "topology-info" { return "TopologyInfo" }
    if yname == "srr-info" { return "SrrInfo" }
    if yname == "rate-limit-info" { return "RateLimitInfo" }
    return ""
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetSegmentPath() string {
    return "srp-information"
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ips-info" {
        return &srpInformation.IpsInfo
    }
    if childYangName == "topology-info" {
        return &srpInformation.TopologyInfo
    }
    if childYangName == "srr-info" {
        return &srpInformation.SrrInfo
    }
    if childYangName == "rate-limit-info" {
        return &srpInformation.RateLimitInfo
    }
    return nil
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ips-info"] = &srpInformation.IpsInfo
    children["topology-info"] = &srpInformation.TopologyInfo
    children["srr-info"] = &srpInformation.SrrInfo
    children["rate-limit-info"] = &srpInformation.RateLimitInfo
    return children
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetYangName() string { return "srp-information" }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) SetParent(parent types.Entity) { srpInformation.parent = parent }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetParent() types.Entity { return srpInformation.parent }

func (srpInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation) GetParentYangName() string { return "srp-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo
// SRP IPS information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is the interfaceadministratively down. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAdminDown interface{}

    // IPS information. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation.
    LocalInformation []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation
}

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetFilter() yfilter.YFilter { return ipsInfo.YFilter }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) SetFilter(yf yfilter.YFilter) { ipsInfo.YFilter = yf }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetGoName(yname string) string {
    if yname == "is-admin-down" { return "IsAdminDown" }
    if yname == "local-information" { return "LocalInformation" }
    return ""
}

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetSegmentPath() string {
    return "ips-info"
}

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-information" {
        for _, c := range ipsInfo.LocalInformation {
            if ipsInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation{}
        ipsInfo.LocalInformation = append(ipsInfo.LocalInformation, child)
        return &ipsInfo.LocalInformation[len(ipsInfo.LocalInformation)-1]
    }
    return nil
}

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipsInfo.LocalInformation {
        children[ipsInfo.LocalInformation[i].GetSegmentPath()] = &ipsInfo.LocalInformation[i]
    }
    return children
}

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-admin-down"] = ipsInfo.IsAdminDown
    return leafs
}

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetYangName() string { return "ips-info" }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) SetParent(parent types.Entity) { ipsInfo.parent = parent }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetParent() types.Entity { return ipsInfo.parent }

func (ipsInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo) GetParentYangName() string { return "srp-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation
// IPS information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation struct {
    parent types.Entity
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

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetFilter() yfilter.YFilter { return localInformation.YFilter }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) SetFilter(yf yfilter.YFilter) { localInformation.YFilter = yf }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "is-inter-card-bus-enabled" { return "IsInterCardBusEnabled" }
    if yname == "wtr-timer-period" { return "WtrTimerPeriod" }
    if yname == "side-a" { return "SideA" }
    if yname == "side-b" { return "SideB" }
    return ""
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetSegmentPath() string {
    return "local-information"
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "side-a" {
        return &localInformation.SideA
    }
    if childYangName == "side-b" {
        return &localInformation.SideB
    }
    return nil
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["side-a"] = &localInformation.SideA
    children["side-b"] = &localInformation.SideB
    return children
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = localInformation.MacAddress
    leafs["is-inter-card-bus-enabled"] = localInformation.IsInterCardBusEnabled
    leafs["wtr-timer-period"] = localInformation.WtrTimerPeriod
    return leafs
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetBundleName() string { return "cisco_ios_xr" }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetYangName() string { return "local-information" }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) SetParent(parent types.Entity) { localInformation.parent = parent }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetParent() types.Entity { return localInformation.parent }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation) GetParentYangName() string { return "ips-info" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA
// Side A IPS details
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA struct {
    parent types.Entity
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
    AssertedFailure []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure
}

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetFilter() yfilter.YFilter { return sideA.YFilter }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) SetFilter(yf yfilter.YFilter) { sideA.YFilter = yf }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "wrap-state" { return "WrapState" }
    if yname == "packet-sent-timer" { return "PacketSentTimer" }
    if yname == "send-timer-time-remaining" { return "SendTimerTimeRemaining" }
    if yname == "wtr-timer-remaining" { return "WtrTimerRemaining" }
    if yname == "self-detected-request" { return "SelfDetectedRequest" }
    if yname == "remote-request" { return "RemoteRequest" }
    if yname == "rx-neighbor-mac-address" { return "RxNeighborMacAddress" }
    if yname == "rx-message-type" { return "RxMessageType" }
    if yname == "rx-path-type" { return "RxPathType" }
    if yname == "rx-ttl" { return "RxTtl" }
    if yname == "rx-packet-test" { return "RxPacketTest" }
    if yname == "tx-neighbor-mac-address" { return "TxNeighborMacAddress" }
    if yname == "tx-message-type" { return "TxMessageType" }
    if yname == "tx-path-type" { return "TxPathType" }
    if yname == "tx-ttl" { return "TxTtl" }
    if yname == "tx-packet-test" { return "TxPacketTest" }
    if yname == "delay-keep-alive-trigger" { return "DelayKeepAliveTrigger" }
    if yname == "asserted-failure" { return "AssertedFailure" }
    return ""
}

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetSegmentPath() string {
    return "side-a"
}

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asserted-failure" {
        for _, c := range sideA.AssertedFailure {
            if sideA.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure{}
        sideA.AssertedFailure = append(sideA.AssertedFailure, child)
        return &sideA.AssertedFailure[len(sideA.AssertedFailure)-1]
    }
    return nil
}

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sideA.AssertedFailure {
        children[sideA.AssertedFailure[i].GetSegmentPath()] = &sideA.AssertedFailure[i]
    }
    return children
}

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = sideA.MacAddress
    leafs["wrap-state"] = sideA.WrapState
    leafs["packet-sent-timer"] = sideA.PacketSentTimer
    leafs["send-timer-time-remaining"] = sideA.SendTimerTimeRemaining
    leafs["wtr-timer-remaining"] = sideA.WtrTimerRemaining
    leafs["self-detected-request"] = sideA.SelfDetectedRequest
    leafs["remote-request"] = sideA.RemoteRequest
    leafs["rx-neighbor-mac-address"] = sideA.RxNeighborMacAddress
    leafs["rx-message-type"] = sideA.RxMessageType
    leafs["rx-path-type"] = sideA.RxPathType
    leafs["rx-ttl"] = sideA.RxTtl
    leafs["rx-packet-test"] = sideA.RxPacketTest
    leafs["tx-neighbor-mac-address"] = sideA.TxNeighborMacAddress
    leafs["tx-message-type"] = sideA.TxMessageType
    leafs["tx-path-type"] = sideA.TxPathType
    leafs["tx-ttl"] = sideA.TxTtl
    leafs["tx-packet-test"] = sideA.TxPacketTest
    leafs["delay-keep-alive-trigger"] = sideA.DelayKeepAliveTrigger
    return leafs
}

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetBundleName() string { return "cisco_ios_xr" }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetYangName() string { return "side-a" }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) SetParent(parent types.Entity) { sideA.parent = parent }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetParent() types.Entity { return sideA.parent }

func (sideA *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA) GetParentYangName() string { return "local-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure
// Failures presently asserted
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure struct {
    parent types.Entity
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

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetFilter() yfilter.YFilter { return assertedFailure.YFilter }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) SetFilter(yf yfilter.YFilter) { assertedFailure.YFilter = yf }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "reported-state" { return "ReportedState" }
    if yname == "debounced-state" { return "DebouncedState" }
    if yname == "current-state" { return "CurrentState" }
    if yname == "stable-time" { return "StableTime" }
    if yname == "debounced-delay" { return "DebouncedDelay" }
    return ""
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetSegmentPath() string {
    return "asserted-failure"
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = assertedFailure.Type
    leafs["reported-state"] = assertedFailure.ReportedState
    leafs["debounced-state"] = assertedFailure.DebouncedState
    leafs["current-state"] = assertedFailure.CurrentState
    leafs["stable-time"] = assertedFailure.StableTime
    leafs["debounced-delay"] = assertedFailure.DebouncedDelay
    return leafs
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetBundleName() string { return "cisco_ios_xr" }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetYangName() string { return "asserted-failure" }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) SetParent(parent types.Entity) { assertedFailure.parent = parent }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetParent() types.Entity { return assertedFailure.parent }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideA_AssertedFailure) GetParentYangName() string { return "side-a" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB
// Side B IPS details
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB struct {
    parent types.Entity
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
    AssertedFailure []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure
}

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetFilter() yfilter.YFilter { return sideB.YFilter }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) SetFilter(yf yfilter.YFilter) { sideB.YFilter = yf }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "wrap-state" { return "WrapState" }
    if yname == "packet-sent-timer" { return "PacketSentTimer" }
    if yname == "send-timer-time-remaining" { return "SendTimerTimeRemaining" }
    if yname == "wtr-timer-remaining" { return "WtrTimerRemaining" }
    if yname == "self-detected-request" { return "SelfDetectedRequest" }
    if yname == "remote-request" { return "RemoteRequest" }
    if yname == "rx-neighbor-mac-address" { return "RxNeighborMacAddress" }
    if yname == "rx-message-type" { return "RxMessageType" }
    if yname == "rx-path-type" { return "RxPathType" }
    if yname == "rx-ttl" { return "RxTtl" }
    if yname == "rx-packet-test" { return "RxPacketTest" }
    if yname == "tx-neighbor-mac-address" { return "TxNeighborMacAddress" }
    if yname == "tx-message-type" { return "TxMessageType" }
    if yname == "tx-path-type" { return "TxPathType" }
    if yname == "tx-ttl" { return "TxTtl" }
    if yname == "tx-packet-test" { return "TxPacketTest" }
    if yname == "delay-keep-alive-trigger" { return "DelayKeepAliveTrigger" }
    if yname == "asserted-failure" { return "AssertedFailure" }
    return ""
}

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetSegmentPath() string {
    return "side-b"
}

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asserted-failure" {
        for _, c := range sideB.AssertedFailure {
            if sideB.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure{}
        sideB.AssertedFailure = append(sideB.AssertedFailure, child)
        return &sideB.AssertedFailure[len(sideB.AssertedFailure)-1]
    }
    return nil
}

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sideB.AssertedFailure {
        children[sideB.AssertedFailure[i].GetSegmentPath()] = &sideB.AssertedFailure[i]
    }
    return children
}

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = sideB.MacAddress
    leafs["wrap-state"] = sideB.WrapState
    leafs["packet-sent-timer"] = sideB.PacketSentTimer
    leafs["send-timer-time-remaining"] = sideB.SendTimerTimeRemaining
    leafs["wtr-timer-remaining"] = sideB.WtrTimerRemaining
    leafs["self-detected-request"] = sideB.SelfDetectedRequest
    leafs["remote-request"] = sideB.RemoteRequest
    leafs["rx-neighbor-mac-address"] = sideB.RxNeighborMacAddress
    leafs["rx-message-type"] = sideB.RxMessageType
    leafs["rx-path-type"] = sideB.RxPathType
    leafs["rx-ttl"] = sideB.RxTtl
    leafs["rx-packet-test"] = sideB.RxPacketTest
    leafs["tx-neighbor-mac-address"] = sideB.TxNeighborMacAddress
    leafs["tx-message-type"] = sideB.TxMessageType
    leafs["tx-path-type"] = sideB.TxPathType
    leafs["tx-ttl"] = sideB.TxTtl
    leafs["tx-packet-test"] = sideB.TxPacketTest
    leafs["delay-keep-alive-trigger"] = sideB.DelayKeepAliveTrigger
    return leafs
}

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetBundleName() string { return "cisco_ios_xr" }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetYangName() string { return "side-b" }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) SetParent(parent types.Entity) { sideB.parent = parent }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetParent() types.Entity { return sideB.parent }

func (sideB *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB) GetParentYangName() string { return "local-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure
// Failures presently asserted
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure struct {
    parent types.Entity
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

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetFilter() yfilter.YFilter { return assertedFailure.YFilter }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) SetFilter(yf yfilter.YFilter) { assertedFailure.YFilter = yf }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "reported-state" { return "ReportedState" }
    if yname == "debounced-state" { return "DebouncedState" }
    if yname == "current-state" { return "CurrentState" }
    if yname == "stable-time" { return "StableTime" }
    if yname == "debounced-delay" { return "DebouncedDelay" }
    return ""
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetSegmentPath() string {
    return "asserted-failure"
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = assertedFailure.Type
    leafs["reported-state"] = assertedFailure.ReportedState
    leafs["debounced-state"] = assertedFailure.DebouncedState
    leafs["current-state"] = assertedFailure.CurrentState
    leafs["stable-time"] = assertedFailure.StableTime
    leafs["debounced-delay"] = assertedFailure.DebouncedDelay
    return leafs
}

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetBundleName() string { return "cisco_ios_xr" }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetYangName() string { return "asserted-failure" }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) SetParent(parent types.Entity) { assertedFailure.parent = parent }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetParent() types.Entity { return assertedFailure.parent }

func (assertedFailure *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_IpsInfo_LocalInformation_SideB_AssertedFailure) GetParentYangName() string { return "side-b" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo
// SRP topology information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is the interfaceadministratively down. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAdminDown interface{}

    // Detailed SRP topology information. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation.
    LocalInformation []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation
}

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetFilter() yfilter.YFilter { return topologyInfo.YFilter }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) SetFilter(yf yfilter.YFilter) { topologyInfo.YFilter = yf }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetGoName(yname string) string {
    if yname == "is-admin-down" { return "IsAdminDown" }
    if yname == "local-information" { return "LocalInformation" }
    return ""
}

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetSegmentPath() string {
    return "topology-info"
}

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-information" {
        for _, c := range topologyInfo.LocalInformation {
            if topologyInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation{}
        topologyInfo.LocalInformation = append(topologyInfo.LocalInformation, child)
        return &topologyInfo.LocalInformation[len(topologyInfo.LocalInformation)-1]
    }
    return nil
}

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range topologyInfo.LocalInformation {
        children[topologyInfo.LocalInformation[i].GetSegmentPath()] = &topologyInfo.LocalInformation[i]
    }
    return children
}

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-admin-down"] = topologyInfo.IsAdminDown
    return leafs
}

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetBundleName() string { return "cisco_ios_xr" }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetYangName() string { return "topology-info" }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) SetParent(parent types.Entity) { topologyInfo.parent = parent }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetParent() types.Entity { return topologyInfo.parent }

func (topologyInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo) GetParentYangName() string { return "srp-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation
// Detailed SRP topology information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation struct {
    parent types.Entity
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
    RingNode []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetFilter() yfilter.YFilter { return localInformation.YFilter }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) SetFilter(yf yfilter.YFilter) { localInformation.YFilter = yf }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetGoName(yname string) string {
    if yname == "topology-timer" { return "TopologyTimer" }
    if yname == "next-topology-packet-delay" { return "NextTopologyPacketDelay" }
    if yname == "time-since-last-topology-packet-received" { return "TimeSinceLastTopologyPacketReceived" }
    if yname == "time-since-last-topology-change" { return "TimeSinceLastTopologyChange" }
    if yname == "number-of-nodes-on-ring" { return "NumberOfNodesOnRing" }
    if yname == "ring-node" { return "RingNode" }
    return ""
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetSegmentPath() string {
    return "local-information"
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ring-node" {
        for _, c := range localInformation.RingNode {
            if localInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode{}
        localInformation.RingNode = append(localInformation.RingNode, child)
        return &localInformation.RingNode[len(localInformation.RingNode)-1]
    }
    return nil
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localInformation.RingNode {
        children[localInformation.RingNode[i].GetSegmentPath()] = &localInformation.RingNode[i]
    }
    return children
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-timer"] = localInformation.TopologyTimer
    leafs["next-topology-packet-delay"] = localInformation.NextTopologyPacketDelay
    leafs["time-since-last-topology-packet-received"] = localInformation.TimeSinceLastTopologyPacketReceived
    leafs["time-since-last-topology-change"] = localInformation.TimeSinceLastTopologyChange
    leafs["number-of-nodes-on-ring"] = localInformation.NumberOfNodesOnRing
    return leafs
}

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetBundleName() string { return "cisco_ios_xr" }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetYangName() string { return "local-information" }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) SetParent(parent types.Entity) { localInformation.parent = parent }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetParent() types.Entity { return localInformation.parent }

func (localInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation) GetParentYangName() string { return "topology-info" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode
// List of nodes on the ring info
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode struct {
    parent types.Entity
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

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetFilter() yfilter.YFilter { return ringNode.YFilter }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) SetFilter(yf yfilter.YFilter) { ringNode.YFilter = yf }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetGoName(yname string) string {
    if yname == "hop-count" { return "HopCount" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "is-wrapped" { return "IsWrapped" }
    if yname == "is-srr-supported" { return "IsSrrSupported" }
    if yname == "node-name" { return "NodeName" }
    return ""
}

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetSegmentPath() string {
    return "ring-node"
}

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hop-count"] = ringNode.HopCount
    leafs["mac-address"] = ringNode.MacAddress
    leafs["ipv4-address"] = ringNode.Ipv4Address
    leafs["is-wrapped"] = ringNode.IsWrapped
    leafs["is-srr-supported"] = ringNode.IsSrrSupported
    leafs["node-name"] = ringNode.NodeName
    return leafs
}

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetBundleName() string { return "cisco_ios_xr" }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetYangName() string { return "ring-node" }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) SetParent(parent types.Entity) { ringNode.parent = parent }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetParent() types.Entity { return ringNode.parent }

func (ringNode *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_TopologyInfo_LocalInformation_RingNode) GetParentYangName() string { return "local-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo
// SRP SRR information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is the interfaceadministratively down. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAdminDown interface{}

    // SRR enabled. The type is interface{} with range: -2147483648..2147483647.
    IsSrrEnabled interface{}

    // SRP information. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo.
    SrrDetailedInfo []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo
}

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetFilter() yfilter.YFilter { return srrInfo.YFilter }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) SetFilter(yf yfilter.YFilter) { srrInfo.YFilter = yf }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetGoName(yname string) string {
    if yname == "is-admin-down" { return "IsAdminDown" }
    if yname == "is-srr-enabled" { return "IsSrrEnabled" }
    if yname == "srr-detailed-info" { return "SrrDetailedInfo" }
    return ""
}

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetSegmentPath() string {
    return "srr-info"
}

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srr-detailed-info" {
        for _, c := range srrInfo.SrrDetailedInfo {
            if srrInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo{}
        srrInfo.SrrDetailedInfo = append(srrInfo.SrrDetailedInfo, child)
        return &srrInfo.SrrDetailedInfo[len(srrInfo.SrrDetailedInfo)-1]
    }
    return nil
}

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range srrInfo.SrrDetailedInfo {
        children[srrInfo.SrrDetailedInfo[i].GetSegmentPath()] = &srrInfo.SrrDetailedInfo[i]
    }
    return children
}

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-admin-down"] = srrInfo.IsAdminDown
    leafs["is-srr-enabled"] = srrInfo.IsSrrEnabled
    return leafs
}

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetBundleName() string { return "cisco_ios_xr" }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetYangName() string { return "srr-info" }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) SetParent(parent types.Entity) { srrInfo.parent = parent }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetParent() types.Entity { return srrInfo.parent }

func (srrInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo) GetParentYangName() string { return "srp-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo
// SRP information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo struct {
    parent types.Entity
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
    NodesOnRing []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing

    // nodes not in topology map. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing.
    NodesNotOnRing []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing
}

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetFilter() yfilter.YFilter { return srrDetailedInfo.YFilter }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) SetFilter(yf yfilter.YFilter) { srrDetailedInfo.YFilter = yf }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetGoName(yname string) string {
    if yname == "version-number" { return "VersionNumber" }
    if yname == "is-wrong-version-received" { return "IsWrongVersionReceived" }
    if yname == "last-wrong-version-receive-time" { return "LastWrongVersionReceiveTime" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "node-state" { return "NodeState" }
    if yname == "is-outer-ring-in-use" { return "IsOuterRingInUse" }
    if yname == "is-inner-ring-in-use" { return "IsInnerRingInUse" }
    if yname == "is-announce" { return "IsAnnounce" }
    if yname == "outer-fail-type" { return "OuterFailType" }
    if yname == "inner-fail-type" { return "InnerFailType" }
    if yname == "packet-send-timer" { return "PacketSendTimer" }
    if yname == "next-srr-packet-send-time" { return "NextSrrPacketSendTime" }
    if yname == "single-ring-bw" { return "SingleRingBw" }
    if yname == "wtr-time" { return "WtrTime" }
    if yname == "wtr-timer-remaining-outer-ring" { return "WtrTimerRemainingOuterRing" }
    if yname == "wtr-timer-remaining-inner-ring" { return "WtrTimerRemainingInnerRing" }
    if yname == "nodes-on-ring" { return "NodesOnRing" }
    if yname == "nodes-not-on-ring" { return "NodesNotOnRing" }
    return ""
}

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetSegmentPath() string {
    return "srr-detailed-info"
}

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes-on-ring" {
        for _, c := range srrDetailedInfo.NodesOnRing {
            if srrDetailedInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing{}
        srrDetailedInfo.NodesOnRing = append(srrDetailedInfo.NodesOnRing, child)
        return &srrDetailedInfo.NodesOnRing[len(srrDetailedInfo.NodesOnRing)-1]
    }
    if childYangName == "nodes-not-on-ring" {
        for _, c := range srrDetailedInfo.NodesNotOnRing {
            if srrDetailedInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing{}
        srrDetailedInfo.NodesNotOnRing = append(srrDetailedInfo.NodesNotOnRing, child)
        return &srrDetailedInfo.NodesNotOnRing[len(srrDetailedInfo.NodesNotOnRing)-1]
    }
    return nil
}

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range srrDetailedInfo.NodesOnRing {
        children[srrDetailedInfo.NodesOnRing[i].GetSegmentPath()] = &srrDetailedInfo.NodesOnRing[i]
    }
    for i := range srrDetailedInfo.NodesNotOnRing {
        children[srrDetailedInfo.NodesNotOnRing[i].GetSegmentPath()] = &srrDetailedInfo.NodesNotOnRing[i]
    }
    return children
}

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version-number"] = srrDetailedInfo.VersionNumber
    leafs["is-wrong-version-received"] = srrDetailedInfo.IsWrongVersionReceived
    leafs["last-wrong-version-receive-time"] = srrDetailedInfo.LastWrongVersionReceiveTime
    leafs["mac-address"] = srrDetailedInfo.MacAddress
    leafs["node-state"] = srrDetailedInfo.NodeState
    leafs["is-outer-ring-in-use"] = srrDetailedInfo.IsOuterRingInUse
    leafs["is-inner-ring-in-use"] = srrDetailedInfo.IsInnerRingInUse
    leafs["is-announce"] = srrDetailedInfo.IsAnnounce
    leafs["outer-fail-type"] = srrDetailedInfo.OuterFailType
    leafs["inner-fail-type"] = srrDetailedInfo.InnerFailType
    leafs["packet-send-timer"] = srrDetailedInfo.PacketSendTimer
    leafs["next-srr-packet-send-time"] = srrDetailedInfo.NextSrrPacketSendTime
    leafs["single-ring-bw"] = srrDetailedInfo.SingleRingBw
    leafs["wtr-time"] = srrDetailedInfo.WtrTime
    leafs["wtr-timer-remaining-outer-ring"] = srrDetailedInfo.WtrTimerRemainingOuterRing
    leafs["wtr-timer-remaining-inner-ring"] = srrDetailedInfo.WtrTimerRemainingInnerRing
    return leafs
}

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetBundleName() string { return "cisco_ios_xr" }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetYangName() string { return "srr-detailed-info" }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) SetParent(parent types.Entity) { srrDetailedInfo.parent = parent }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetParent() types.Entity { return srrDetailedInfo.parent }

func (srrDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo) GetParentYangName() string { return "srr-info" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing
// List of nodes on the ring info
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing struct {
    parent types.Entity
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

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetFilter() yfilter.YFilter { return nodesOnRing.YFilter }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) SetFilter(yf yfilter.YFilter) { nodesOnRing.YFilter = yf }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "srr-entry-exits" { return "SrrEntryExits" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "outer-failure" { return "OuterFailure" }
    if yname == "inner-failure" { return "InnerFailure" }
    if yname == "is-last-announce-received" { return "IsLastAnnounceReceived" }
    if yname == "last-announce-received-time" { return "LastAnnounceReceivedTime" }
    return ""
}

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetSegmentPath() string {
    return "nodes-on-ring"
}

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodesOnRing.NodeName
    leafs["srr-entry-exits"] = nodesOnRing.SrrEntryExits
    leafs["mac-address"] = nodesOnRing.MacAddress
    leafs["outer-failure"] = nodesOnRing.OuterFailure
    leafs["inner-failure"] = nodesOnRing.InnerFailure
    leafs["is-last-announce-received"] = nodesOnRing.IsLastAnnounceReceived
    leafs["last-announce-received-time"] = nodesOnRing.LastAnnounceReceivedTime
    return leafs
}

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetBundleName() string { return "cisco_ios_xr" }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetYangName() string { return "nodes-on-ring" }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) SetParent(parent types.Entity) { nodesOnRing.parent = parent }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetParent() types.Entity { return nodesOnRing.parent }

func (nodesOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesOnRing) GetParentYangName() string { return "srr-detailed-info" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing
// nodes not in topology map
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing struct {
    parent types.Entity
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

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetFilter() yfilter.YFilter { return nodesNotOnRing.YFilter }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) SetFilter(yf yfilter.YFilter) { nodesNotOnRing.YFilter = yf }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "srr-entry-exits" { return "SrrEntryExits" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "outer-failure" { return "OuterFailure" }
    if yname == "inner-failure" { return "InnerFailure" }
    if yname == "is-last-announce-received" { return "IsLastAnnounceReceived" }
    if yname == "last-announce-received-time" { return "LastAnnounceReceivedTime" }
    return ""
}

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetSegmentPath() string {
    return "nodes-not-on-ring"
}

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodesNotOnRing.NodeName
    leafs["srr-entry-exits"] = nodesNotOnRing.SrrEntryExits
    leafs["mac-address"] = nodesNotOnRing.MacAddress
    leafs["outer-failure"] = nodesNotOnRing.OuterFailure
    leafs["inner-failure"] = nodesNotOnRing.InnerFailure
    leafs["is-last-announce-received"] = nodesNotOnRing.IsLastAnnounceReceived
    leafs["last-announce-received-time"] = nodesNotOnRing.LastAnnounceReceivedTime
    return leafs
}

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetBundleName() string { return "cisco_ios_xr" }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetYangName() string { return "nodes-not-on-ring" }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) SetParent(parent types.Entity) { nodesNotOnRing.parent = parent }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetParent() types.Entity { return nodesNotOnRing.parent }

func (nodesNotOnRing *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_SrrInfo_SrrDetailedInfo_NodesNotOnRing) GetParentYangName() string { return "srr-detailed-info" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo
// SRP rate limit information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is the interfaceadministratively down. The type is interface{} with range:
    // -2147483648..2147483647.
    IsAdminDown interface{}

    // SRP rate limit information. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo.
    RateLimitDetailedInfo []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo
}

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetFilter() yfilter.YFilter { return rateLimitInfo.YFilter }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) SetFilter(yf yfilter.YFilter) { rateLimitInfo.YFilter = yf }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetGoName(yname string) string {
    if yname == "is-admin-down" { return "IsAdminDown" }
    if yname == "rate-limit-detailed-info" { return "RateLimitDetailedInfo" }
    return ""
}

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetSegmentPath() string {
    return "rate-limit-info"
}

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rate-limit-detailed-info" {
        for _, c := range rateLimitInfo.RateLimitDetailedInfo {
            if rateLimitInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo{}
        rateLimitInfo.RateLimitDetailedInfo = append(rateLimitInfo.RateLimitDetailedInfo, child)
        return &rateLimitInfo.RateLimitDetailedInfo[len(rateLimitInfo.RateLimitDetailedInfo)-1]
    }
    return nil
}

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rateLimitInfo.RateLimitDetailedInfo {
        children[rateLimitInfo.RateLimitDetailedInfo[i].GetSegmentPath()] = &rateLimitInfo.RateLimitDetailedInfo[i]
    }
    return children
}

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-admin-down"] = rateLimitInfo.IsAdminDown
    return leafs
}

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetBundleName() string { return "cisco_ios_xr" }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetYangName() string { return "rate-limit-info" }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) SetParent(parent types.Entity) { rateLimitInfo.parent = parent }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetParent() types.Entity { return rateLimitInfo.parent }

func (rateLimitInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo) GetParentYangName() string { return "srp-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo
// SRP rate limit information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum SRP priority for high-priority transmit queue. The type is
    // interface{} with range: 0..65535.
    MinPriorityValue interface{}
}

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetFilter() yfilter.YFilter { return rateLimitDetailedInfo.YFilter }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) SetFilter(yf yfilter.YFilter) { rateLimitDetailedInfo.YFilter = yf }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetGoName(yname string) string {
    if yname == "min-priority-value" { return "MinPriorityValue" }
    return ""
}

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetSegmentPath() string {
    return "rate-limit-detailed-info"
}

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min-priority-value"] = rateLimitDetailedInfo.MinPriorityValue
    return leafs
}

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetBundleName() string { return "cisco_ios_xr" }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetYangName() string { return "rate-limit-detailed-info" }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) SetParent(parent types.Entity) { rateLimitDetailedInfo.parent = parent }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetParent() types.Entity { return rateLimitDetailedInfo.parent }

func (rateLimitDetailedInfo *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpInformation_RateLimitInfo_RateLimitDetailedInfo) GetParentYangName() string { return "rate-limit-info" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics
// SRP-specific packet and byte counters
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics struct {
    parent types.Entity
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

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetFilter() yfilter.YFilter { return srpStatistics.YFilter }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) SetFilter(yf yfilter.YFilter) { srpStatistics.YFilter = yf }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetGoName(yname string) string {
    if yname == "data-rate-interval" { return "DataRateInterval" }
    if yname == "side-a-data-rate" { return "SideADataRate" }
    if yname == "side-b-data-rate" { return "SideBDataRate" }
    if yname == "side-a-errors" { return "SideAErrors" }
    if yname == "side-b-errors" { return "SideBErrors" }
    return ""
}

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetSegmentPath() string {
    return "srp-statistics"
}

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "side-a-data-rate" {
        return &srpStatistics.SideADataRate
    }
    if childYangName == "side-b-data-rate" {
        return &srpStatistics.SideBDataRate
    }
    if childYangName == "side-a-errors" {
        return &srpStatistics.SideAErrors
    }
    if childYangName == "side-b-errors" {
        return &srpStatistics.SideBErrors
    }
    return nil
}

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["side-a-data-rate"] = &srpStatistics.SideADataRate
    children["side-b-data-rate"] = &srpStatistics.SideBDataRate
    children["side-a-errors"] = &srpStatistics.SideAErrors
    children["side-b-errors"] = &srpStatistics.SideBErrors
    return children
}

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data-rate-interval"] = srpStatistics.DataRateInterval
    return leafs
}

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetYangName() string { return "srp-statistics" }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) SetParent(parent types.Entity) { srpStatistics.parent = parent }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetParent() types.Entity { return srpStatistics.parent }

func (srpStatistics *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics) GetParentYangName() string { return "srp-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate
// Data rates for side A interface
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate struct {
    parent types.Entity
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

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetFilter() yfilter.YFilter { return sideADataRate.YFilter }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) SetFilter(yf yfilter.YFilter) { sideADataRate.YFilter = yf }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetGoName(yname string) string {
    if yname == "bit-rate-sent" { return "BitRateSent" }
    if yname == "packet-rate-sent" { return "PacketRateSent" }
    if yname == "bit-rate-received" { return "BitRateReceived" }
    if yname == "packet-rate-received" { return "PacketRateReceived" }
    return ""
}

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetSegmentPath() string {
    return "side-a-data-rate"
}

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bit-rate-sent"] = sideADataRate.BitRateSent
    leafs["packet-rate-sent"] = sideADataRate.PacketRateSent
    leafs["bit-rate-received"] = sideADataRate.BitRateReceived
    leafs["packet-rate-received"] = sideADataRate.PacketRateReceived
    return leafs
}

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetBundleName() string { return "cisco_ios_xr" }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetYangName() string { return "side-a-data-rate" }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) SetParent(parent types.Entity) { sideADataRate.parent = parent }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetParent() types.Entity { return sideADataRate.parent }

func (sideADataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideADataRate) GetParentYangName() string { return "srp-statistics" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate
// Data rates for side B interface
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate struct {
    parent types.Entity
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

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetFilter() yfilter.YFilter { return sideBDataRate.YFilter }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) SetFilter(yf yfilter.YFilter) { sideBDataRate.YFilter = yf }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetGoName(yname string) string {
    if yname == "bit-rate-sent" { return "BitRateSent" }
    if yname == "packet-rate-sent" { return "PacketRateSent" }
    if yname == "bit-rate-received" { return "BitRateReceived" }
    if yname == "packet-rate-received" { return "PacketRateReceived" }
    return ""
}

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetSegmentPath() string {
    return "side-b-data-rate"
}

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bit-rate-sent"] = sideBDataRate.BitRateSent
    leafs["packet-rate-sent"] = sideBDataRate.PacketRateSent
    leafs["bit-rate-received"] = sideBDataRate.BitRateReceived
    leafs["packet-rate-received"] = sideBDataRate.PacketRateReceived
    return leafs
}

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetBundleName() string { return "cisco_ios_xr" }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetYangName() string { return "side-b-data-rate" }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) SetParent(parent types.Entity) { sideBDataRate.parent = parent }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetParent() types.Entity { return sideBDataRate.parent }

func (sideBDataRate *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBDataRate) GetParentYangName() string { return "srp-statistics" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors
// Errors for side A interface
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors struct {
    parent types.Entity
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

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetFilter() yfilter.YFilter { return sideAErrors.YFilter }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) SetFilter(yf yfilter.YFilter) { sideAErrors.YFilter = yf }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetGoName(yname string) string {
    if yname == "error-packets-received" { return "ErrorPacketsReceived" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-insufficient-resource-events" { return "InputInsufficientResourceEvents" }
    if yname == "mac-aborts-received" { return "MacAbortsReceived" }
    if yname == "mac-runt-packets-received" { return "MacRuntPacketsReceived" }
    if yname == "mac-giant-packets-received" { return "MacGiantPacketsReceived" }
    if yname == "framer-runt-packets-received" { return "FramerRuntPacketsReceived" }
    if yname == "framer-giant-packets-received" { return "FramerGiantPacketsReceived" }
    if yname == "framer-aborts-received" { return "FramerAbortsReceived" }
    return ""
}

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetSegmentPath() string {
    return "side-a-errors"
}

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["error-packets-received"] = sideAErrors.ErrorPacketsReceived
    leafs["crc-errors"] = sideAErrors.CrcErrors
    leafs["input-insufficient-resource-events"] = sideAErrors.InputInsufficientResourceEvents
    leafs["mac-aborts-received"] = sideAErrors.MacAbortsReceived
    leafs["mac-runt-packets-received"] = sideAErrors.MacRuntPacketsReceived
    leafs["mac-giant-packets-received"] = sideAErrors.MacGiantPacketsReceived
    leafs["framer-runt-packets-received"] = sideAErrors.FramerRuntPacketsReceived
    leafs["framer-giant-packets-received"] = sideAErrors.FramerGiantPacketsReceived
    leafs["framer-aborts-received"] = sideAErrors.FramerAbortsReceived
    return leafs
}

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetBundleName() string { return "cisco_ios_xr" }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetYangName() string { return "side-a-errors" }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) SetParent(parent types.Entity) { sideAErrors.parent = parent }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetParent() types.Entity { return sideAErrors.parent }

func (sideAErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideAErrors) GetParentYangName() string { return "srp-statistics" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors
// Errors for side B interface
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors struct {
    parent types.Entity
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

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetFilter() yfilter.YFilter { return sideBErrors.YFilter }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) SetFilter(yf yfilter.YFilter) { sideBErrors.YFilter = yf }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetGoName(yname string) string {
    if yname == "error-packets-received" { return "ErrorPacketsReceived" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-insufficient-resource-events" { return "InputInsufficientResourceEvents" }
    if yname == "mac-aborts-received" { return "MacAbortsReceived" }
    if yname == "mac-runt-packets-received" { return "MacRuntPacketsReceived" }
    if yname == "mac-giant-packets-received" { return "MacGiantPacketsReceived" }
    if yname == "framer-runt-packets-received" { return "FramerRuntPacketsReceived" }
    if yname == "framer-giant-packets-received" { return "FramerGiantPacketsReceived" }
    if yname == "framer-aborts-received" { return "FramerAbortsReceived" }
    return ""
}

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetSegmentPath() string {
    return "side-b-errors"
}

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["error-packets-received"] = sideBErrors.ErrorPacketsReceived
    leafs["crc-errors"] = sideBErrors.CrcErrors
    leafs["input-insufficient-resource-events"] = sideBErrors.InputInsufficientResourceEvents
    leafs["mac-aborts-received"] = sideBErrors.MacAbortsReceived
    leafs["mac-runt-packets-received"] = sideBErrors.MacRuntPacketsReceived
    leafs["mac-giant-packets-received"] = sideBErrors.MacGiantPacketsReceived
    leafs["framer-runt-packets-received"] = sideBErrors.FramerRuntPacketsReceived
    leafs["framer-giant-packets-received"] = sideBErrors.FramerGiantPacketsReceived
    leafs["framer-aborts-received"] = sideBErrors.FramerAbortsReceived
    return leafs
}

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetBundleName() string { return "cisco_ios_xr" }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetYangName() string { return "side-b-errors" }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) SetParent(parent types.Entity) { sideBErrors.parent = parent }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetParent() types.Entity { return sideBErrors.parent }

func (sideBErrors *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SrpInformation_SrpStatistics_SideBErrors) GetParentYangName() string { return "srp-statistics" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation
// Tunnel interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation struct {
    parent types.Entity
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

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetFilter() yfilter.YFilter { return tunnelInformation.YFilter }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) SetFilter(yf yfilter.YFilter) { tunnelInformation.YFilter = yf }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetGoName(yname string) string {
    if yname == "source-name" { return "SourceName" }
    if yname == "source-ipv4-address" { return "SourceIpv4Address" }
    if yname == "destination-ipv4-address" { return "DestinationIpv4Address" }
    if yname == "tunnel-type" { return "TunnelType" }
    if yname == "key" { return "Key" }
    if yname == "ttl" { return "Ttl" }
    return ""
}

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetSegmentPath() string {
    return "tunnel-information"
}

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-name"] = tunnelInformation.SourceName
    leafs["source-ipv4-address"] = tunnelInformation.SourceIpv4Address
    leafs["destination-ipv4-address"] = tunnelInformation.DestinationIpv4Address
    leafs["tunnel-type"] = tunnelInformation.TunnelType
    leafs["key"] = tunnelInformation.Key
    leafs["ttl"] = tunnelInformation.Ttl
    return leafs
}

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetYangName() string { return "tunnel-information" }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) SetParent(parent types.Entity) { tunnelInformation.parent = parent }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetParent() types.Entity { return tunnelInformation.parent }

func (tunnelInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation
// Bundle interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of bundle members and their properties. The type is slice of
    // Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member.
    Member []Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member
}

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetFilter() yfilter.YFilter { return bundleInformation.YFilter }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) SetFilter(yf yfilter.YFilter) { bundleInformation.YFilter = yf }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetGoName(yname string) string {
    if yname == "member" { return "Member" }
    return ""
}

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetSegmentPath() string {
    return "bundle-information"
}

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member" {
        for _, c := range bundleInformation.Member {
            if bundleInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member{}
        bundleInformation.Member = append(bundleInformation.Member, child)
        return &bundleInformation.Member[len(bundleInformation.Member)-1]
    }
    return nil
}

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bundleInformation.Member {
        children[bundleInformation.Member[i].GetSegmentPath()] = &bundleInformation.Member[i]
    }
    return children
}

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetYangName() string { return "bundle-information" }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) SetParent(parent types.Entity) { bundleInformation.parent = parent }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetParent() types.Entity { return bundleInformation.parent }

func (bundleInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member
// List of bundle members and their properties
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Member's interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetFilter() yfilter.YFilter { return member.YFilter }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) SetFilter(yf yfilter.YFilter) { member.YFilter = yf }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "port-priority" { return "PortPriority" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "underlying-link-id" { return "UnderlyingLinkId" }
    if yname == "link-order-number" { return "LinkOrderNumber" }
    if yname == "iccp-node" { return "IccpNode" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "lacp-enabled" { return "LacpEnabled" }
    if yname == "member-type" { return "MemberType" }
    if yname == "member-name" { return "MemberName" }
    if yname == "counters" { return "Counters" }
    if yname == "link-data" { return "LinkData" }
    if yname == "member-mux-data" { return "MemberMuxData" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetSegmentPath() string {
    return "member"
}

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &member.Counters
    }
    if childYangName == "link-data" {
        return &member.LinkData
    }
    if childYangName == "member-mux-data" {
        return &member.MemberMuxData
    }
    if childYangName == "mac-address" {
        return &member.MacAddress
    }
    return nil
}

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &member.Counters
    children["link-data"] = &member.LinkData
    children["member-mux-data"] = &member.MemberMuxData
    children["mac-address"] = &member.MacAddress
    return children
}

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = member.InterfaceName
    leafs["port-priority"] = member.PortPriority
    leafs["port-number"] = member.PortNumber
    leafs["underlying-link-id"] = member.UnderlyingLinkId
    leafs["link-order-number"] = member.LinkOrderNumber
    leafs["iccp-node"] = member.IccpNode
    leafs["bandwidth"] = member.Bandwidth
    leafs["lacp-enabled"] = member.LacpEnabled
    leafs["member-type"] = member.MemberType
    leafs["member-name"] = member.MemberName
    return leafs
}

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetBundleName() string { return "cisco_ios_xr" }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetYangName() string { return "member" }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) SetParent(parent types.Entity) { member.parent = parent }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetParent() types.Entity { return member.parent }

func (member *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member) GetParentYangName() string { return "bundle-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters
// Counters data about member link
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters struct {
    parent types.Entity
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

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetGoName(yname string) string {
    if yname == "lacpd-us-received" { return "LacpdUsReceived" }
    if yname == "lacpd-us-transmitted" { return "LacpdUsTransmitted" }
    if yname == "marker-packets-received" { return "MarkerPacketsReceived" }
    if yname == "marker-responses-transmitted" { return "MarkerResponsesTransmitted" }
    if yname == "illegal-packets-received" { return "IllegalPacketsReceived" }
    if yname == "excess-lacpd-us-received" { return "ExcessLacpdUsReceived" }
    if yname == "excess-marker-packets-received" { return "ExcessMarkerPacketsReceived" }
    if yname == "defaulted" { return "Defaulted" }
    if yname == "expired" { return "Expired" }
    if yname == "last-cleared-sec" { return "LastClearedSec" }
    if yname == "last-cleared-nsec" { return "LastClearedNsec" }
    return ""
}

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lacpd-us-received"] = counters.LacpdUsReceived
    leafs["lacpd-us-transmitted"] = counters.LacpdUsTransmitted
    leafs["marker-packets-received"] = counters.MarkerPacketsReceived
    leafs["marker-responses-transmitted"] = counters.MarkerResponsesTransmitted
    leafs["illegal-packets-received"] = counters.IllegalPacketsReceived
    leafs["excess-lacpd-us-received"] = counters.ExcessLacpdUsReceived
    leafs["excess-marker-packets-received"] = counters.ExcessMarkerPacketsReceived
    leafs["defaulted"] = counters.Defaulted
    leafs["expired"] = counters.Expired
    leafs["last-cleared-sec"] = counters.LastClearedSec
    leafs["last-cleared-nsec"] = counters.LastClearedNsec
    return leafs
}

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetYangName() string { return "counters" }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_Counters) GetParentYangName() string { return "member" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData
// Lacp data about member link
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Member's interface handle. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetFilter() yfilter.YFilter { return linkData.YFilter }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) SetFilter(yf yfilter.YFilter) { linkData.YFilter = yf }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetGoName(yname string) string {
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "actor-system-priority" { return "ActorSystemPriority" }
    if yname == "actor-system-mac-address" { return "ActorSystemMacAddress" }
    if yname == "actor-operational-key" { return "ActorOperationalKey" }
    if yname == "partner-system-priority" { return "PartnerSystemPriority" }
    if yname == "partner-system-mac-address" { return "PartnerSystemMacAddress" }
    if yname == "partner-operational-key" { return "PartnerOperationalKey" }
    if yname == "selected-aggregator-id" { return "SelectedAggregatorId" }
    if yname == "attached-aggregator-id" { return "AttachedAggregatorId" }
    if yname == "actor-port-id" { return "ActorPortId" }
    if yname == "actor-port-priority" { return "ActorPortPriority" }
    if yname == "partner-port-id" { return "PartnerPortId" }
    if yname == "partner-port-priority" { return "PartnerPortPriority" }
    if yname == "actor-port-state" { return "ActorPortState" }
    if yname == "partner-port-state" { return "PartnerPortState" }
    return ""
}

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetSegmentPath() string {
    return "link-data"
}

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-handle"] = linkData.InterfaceHandle
    leafs["actor-system-priority"] = linkData.ActorSystemPriority
    leafs["actor-system-mac-address"] = linkData.ActorSystemMacAddress
    leafs["actor-operational-key"] = linkData.ActorOperationalKey
    leafs["partner-system-priority"] = linkData.PartnerSystemPriority
    leafs["partner-system-mac-address"] = linkData.PartnerSystemMacAddress
    leafs["partner-operational-key"] = linkData.PartnerOperationalKey
    leafs["selected-aggregator-id"] = linkData.SelectedAggregatorId
    leafs["attached-aggregator-id"] = linkData.AttachedAggregatorId
    leafs["actor-port-id"] = linkData.ActorPortId
    leafs["actor-port-priority"] = linkData.ActorPortPriority
    leafs["partner-port-id"] = linkData.PartnerPortId
    leafs["partner-port-priority"] = linkData.PartnerPortPriority
    leafs["actor-port-state"] = linkData.ActorPortState
    leafs["partner-port-state"] = linkData.PartnerPortState
    return leafs
}

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetBundleName() string { return "cisco_ios_xr" }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetYangName() string { return "link-data" }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) SetParent(parent types.Entity) { linkData.parent = parent }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetParent() types.Entity { return linkData.parent }

func (linkData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_LinkData) GetParentYangName() string { return "member" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData
// Mux state machine data
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData struct {
    parent types.Entity
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

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetFilter() yfilter.YFilter { return memberMuxData.YFilter }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) SetFilter(yf yfilter.YFilter) { memberMuxData.YFilter = yf }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetGoName(yname string) string {
    if yname == "mux-state" { return "MuxState" }
    if yname == "error" { return "Error" }
    if yname == "member-mux-state-reason" { return "MemberMuxStateReason" }
    if yname == "member-state" { return "MemberState" }
    if yname == "mux-state-reason" { return "MuxStateReason" }
    if yname == "member-mux-state-reason-data" { return "MemberMuxStateReasonData" }
    return ""
}

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetSegmentPath() string {
    return "member-mux-data"
}

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member-mux-state-reason-data" {
        return &memberMuxData.MemberMuxStateReasonData
    }
    return nil
}

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["member-mux-state-reason-data"] = &memberMuxData.MemberMuxStateReasonData
    return children
}

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mux-state"] = memberMuxData.MuxState
    leafs["error"] = memberMuxData.Error
    leafs["member-mux-state-reason"] = memberMuxData.MemberMuxStateReason
    leafs["member-state"] = memberMuxData.MemberState
    leafs["mux-state-reason"] = memberMuxData.MuxStateReason
    return leafs
}

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetBundleName() string { return "cisco_ios_xr" }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetYangName() string { return "member-mux-data" }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) SetParent(parent types.Entity) { memberMuxData.parent = parent }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetParent() types.Entity { return memberMuxData.parent }

func (memberMuxData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData) GetParentYangName() string { return "member" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData
// Data regarding the reason for last Mux state
// change
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The item the reason applies to. The type is BmStateReasonTarget.
    ReasonType interface{}

    // The severity of the reason. The type is BmSeverity.
    Severity interface{}
}

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetFilter() yfilter.YFilter { return memberMuxStateReasonData.YFilter }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) SetFilter(yf yfilter.YFilter) { memberMuxStateReasonData.YFilter = yf }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetGoName(yname string) string {
    if yname == "reason-type" { return "ReasonType" }
    if yname == "severity" { return "Severity" }
    return ""
}

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetSegmentPath() string {
    return "member-mux-state-reason-data"
}

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reason-type"] = memberMuxStateReasonData.ReasonType
    leafs["severity"] = memberMuxStateReasonData.Severity
    return leafs
}

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetBundleName() string { return "cisco_ios_xr" }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetYangName() string { return "member-mux-state-reason-data" }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) SetParent(parent types.Entity) { memberMuxStateReasonData.parent = parent }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetParent() types.Entity { return memberMuxStateReasonData.parent }

func (memberMuxStateReasonData *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MemberMuxData_MemberMuxStateReasonData) GetParentYangName() string { return "member-mux-data" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress
// MAC address of this member (deprecated)
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = macAddress.Address
    return leafs
}

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_BundleInformation_Member_MacAddress) GetParentYangName() string { return "member" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation
// Serial interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeslots separated by : or - from 1 to 31. : indicates individual timeslot
    // and - represents a range. E.g. 1-3:5 represents timeslots 1, 2, 3, and 5.
    // The type is string.
    Timeslots interface{}
}

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetFilter() yfilter.YFilter { return serialInformation.YFilter }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) SetFilter(yf yfilter.YFilter) { serialInformation.YFilter = yf }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetGoName(yname string) string {
    if yname == "timeslots" { return "Timeslots" }
    return ""
}

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetSegmentPath() string {
    return "serial-information"
}

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeslots"] = serialInformation.Timeslots
    return leafs
}

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetBundleName() string { return "cisco_ios_xr" }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetYangName() string { return "serial-information" }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) SetParent(parent types.Entity) { serialInformation.parent = parent }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetParent() types.Entity { return serialInformation.parent }

func (serialInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SerialInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation
// SONET POS interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // APS state. The type is SonetApsEt.
    ApsState interface{}
}

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetFilter() yfilter.YFilter { return sonetPosInformation.YFilter }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) SetFilter(yf yfilter.YFilter) { sonetPosInformation.YFilter = yf }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetGoName(yname string) string {
    if yname == "aps-state" { return "ApsState" }
    return ""
}

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetSegmentPath() string {
    return "sonet-pos-information"
}

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aps-state"] = sonetPosInformation.ApsState
    return leafs
}

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetBundleName() string { return "cisco_ios_xr" }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetYangName() string { return "sonet-pos-information" }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) SetParent(parent types.Entity) { sonetPosInformation.parent = parent }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetParent() types.Entity { return sonetPosInformation.parent }

func (sonetPosInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_SonetPosInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation
// Tunnel GRE interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation struct {
    parent types.Entity
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

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetFilter() yfilter.YFilter { return tunnelGreInformation.YFilter }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) SetFilter(yf yfilter.YFilter) { tunnelGreInformation.YFilter = yf }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetGoName(yname string) string {
    if yname == "source-name" { return "SourceName" }
    if yname == "destination-ip-address-length" { return "DestinationIpAddressLength" }
    if yname == "tunnel-tos" { return "TunnelTos" }
    if yname == "tunnel-ttl" { return "TunnelTtl" }
    if yname == "key" { return "Key" }
    if yname == "keepalive-period" { return "KeepalivePeriod" }
    if yname == "keepalive-maximum-retry" { return "KeepaliveMaximumRetry" }
    if yname == "tunnel-mode" { return "TunnelMode" }
    if yname == "tunnel-mode-direction" { return "TunnelModeDirection" }
    if yname == "keepalive-state" { return "KeepaliveState" }
    if yname == "df-bit-state" { return "DfBitState" }
    if yname == "key-bit-state" { return "KeyBitState" }
    if yname == "source-ip-address" { return "SourceIpAddress" }
    if yname == "destination-ip-address" { return "DestinationIpAddress" }
    return ""
}

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetSegmentPath() string {
    return "tunnel-gre-information"
}

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-ip-address" {
        return &tunnelGreInformation.SourceIpAddress
    }
    if childYangName == "destination-ip-address" {
        return &tunnelGreInformation.DestinationIpAddress
    }
    return nil
}

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source-ip-address"] = &tunnelGreInformation.SourceIpAddress
    children["destination-ip-address"] = &tunnelGreInformation.DestinationIpAddress
    return children
}

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-name"] = tunnelGreInformation.SourceName
    leafs["destination-ip-address-length"] = tunnelGreInformation.DestinationIpAddressLength
    leafs["tunnel-tos"] = tunnelGreInformation.TunnelTos
    leafs["tunnel-ttl"] = tunnelGreInformation.TunnelTtl
    leafs["key"] = tunnelGreInformation.Key
    leafs["keepalive-period"] = tunnelGreInformation.KeepalivePeriod
    leafs["keepalive-maximum-retry"] = tunnelGreInformation.KeepaliveMaximumRetry
    leafs["tunnel-mode"] = tunnelGreInformation.TunnelMode
    leafs["tunnel-mode-direction"] = tunnelGreInformation.TunnelModeDirection
    leafs["keepalive-state"] = tunnelGreInformation.KeepaliveState
    leafs["df-bit-state"] = tunnelGreInformation.DfBitState
    leafs["key-bit-state"] = tunnelGreInformation.KeyBitState
    return leafs
}

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetYangName() string { return "tunnel-gre-information" }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) SetParent(parent types.Entity) { tunnelGreInformation.parent = parent }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetParent() types.Entity { return tunnelGreInformation.parent }

func (tunnelGreInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress
// Tunnel source IP address
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress struct {
    parent types.Entity
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

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetFilter() yfilter.YFilter { return sourceIpAddress.YFilter }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) SetFilter(yf yfilter.YFilter) { sourceIpAddress.YFilter = yf }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetSegmentPath() string {
    return "source-ip-address"
}

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = sourceIpAddress.Afi
    leafs["ipv4"] = sourceIpAddress.Ipv4
    leafs["ipv6"] = sourceIpAddress.Ipv6
    return leafs
}

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetYangName() string { return "source-ip-address" }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) SetParent(parent types.Entity) { sourceIpAddress.parent = parent }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetParent() types.Entity { return sourceIpAddress.parent }

func (sourceIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_SourceIpAddress) GetParentYangName() string { return "tunnel-gre-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress
// Tunnel destination IP address
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress struct {
    parent types.Entity
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

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetFilter() yfilter.YFilter { return destinationIpAddress.YFilter }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) SetFilter(yf yfilter.YFilter) { destinationIpAddress.YFilter = yf }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetSegmentPath() string {
    return "destination-ip-address"
}

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = destinationIpAddress.Afi
    leafs["ipv4"] = destinationIpAddress.Ipv4
    leafs["ipv6"] = destinationIpAddress.Ipv6
    return leafs
}

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetYangName() string { return "destination-ip-address" }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) SetParent(parent types.Entity) { destinationIpAddress.parent = parent }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetParent() types.Entity { return destinationIpAddress.parent }

func (destinationIpAddress *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_TunnelGreInformation_DestinationIpAddress) GetParentYangName() string { return "tunnel-gre-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation
// PseudowireHeadEnd interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface list Name. The type is string.
    InterfaceListName interface{}

    // L2 Overhead. The type is interface{} with range: 0..4294967295.
    L2Overhead interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}
}

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetFilter() yfilter.YFilter { return pseudowireHeadEndInformation.YFilter }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) SetFilter(yf yfilter.YFilter) { pseudowireHeadEndInformation.YFilter = yf }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetGoName(yname string) string {
    if yname == "interface-list-name" { return "InterfaceListName" }
    if yname == "l2-overhead" { return "L2Overhead" }
    if yname == "internal-label" { return "InternalLabel" }
    return ""
}

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetSegmentPath() string {
    return "pseudowire-head-end-information"
}

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-list-name"] = pseudowireHeadEndInformation.InterfaceListName
    leafs["l2-overhead"] = pseudowireHeadEndInformation.L2Overhead
    leafs["internal-label"] = pseudowireHeadEndInformation.InternalLabel
    return leafs
}

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetBundleName() string { return "cisco_ios_xr" }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetYangName() string { return "pseudowire-head-end-information" }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) SetParent(parent types.Entity) { pseudowireHeadEndInformation.parent = parent }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetParent() types.Entity { return pseudowireHeadEndInformation.parent }

func (pseudowireHeadEndInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_PseudowireHeadEndInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation
// Cem interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation struct {
    parent types.Entity
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

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetFilter() yfilter.YFilter { return cemInformation.YFilter }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) SetFilter(yf yfilter.YFilter) { cemInformation.YFilter = yf }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetGoName(yname string) string {
    if yname == "timeslots" { return "Timeslots" }
    if yname == "payload" { return "Payload" }
    if yname == "dejitter-buffer" { return "DejitterBuffer" }
    if yname == "framing" { return "Framing" }
    return ""
}

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetSegmentPath() string {
    return "cem-information"
}

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeslots"] = cemInformation.Timeslots
    leafs["payload"] = cemInformation.Payload
    leafs["dejitter-buffer"] = cemInformation.DejitterBuffer
    leafs["framing"] = cemInformation.Framing
    return leafs
}

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetBundleName() string { return "cisco_ios_xr" }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetYangName() string { return "cem-information" }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) SetParent(parent types.Entity) { cemInformation.parent = parent }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetParent() types.Entity { return cemInformation.parent }

func (cemInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_CemInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation
// GCC interface information
type Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Derived State. The type is GccDerState.
    DerivedMode interface{}

    // Sec State . The type is GccSecState.
    SecState interface{}
}

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetFilter() yfilter.YFilter { return gccInformation.YFilter }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) SetFilter(yf yfilter.YFilter) { gccInformation.YFilter = yf }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetGoName(yname string) string {
    if yname == "derived-mode" { return "DerivedMode" }
    if yname == "sec-state" { return "SecState" }
    return ""
}

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetSegmentPath() string {
    return "gcc-information"
}

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["derived-mode"] = gccInformation.DerivedMode
    leafs["sec-state"] = gccInformation.SecState
    return leafs
}

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetBundleName() string { return "cisco_ios_xr" }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetYangName() string { return "gcc-information" }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) SetParent(parent types.Entity) { gccInformation.parent = parent }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetParent() types.Entity { return gccInformation.parent }

func (gccInformation *Interfaces_InterfaceXr_Interface_InterfaceTypeInformation_GccInformation) GetParentYangName() string { return "interface-type-information" }

// Interfaces_InterfaceXr_Interface_DataRates
// Packet and byte rates
type Interfaces_InterfaceXr_Interface_DataRates struct {
    parent types.Entity
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

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetFilter() yfilter.YFilter { return dataRates.YFilter }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) SetFilter(yf yfilter.YFilter) { dataRates.YFilter = yf }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetGoName(yname string) string {
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    if yname == "peak-input-data-rate" { return "PeakInputDataRate" }
    if yname == "peak-input-packet-rate" { return "PeakInputPacketRate" }
    if yname == "peak-output-data-rate" { return "PeakOutputDataRate" }
    if yname == "peak-output-packet-rate" { return "PeakOutputPacketRate" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "load-interval" { return "LoadInterval" }
    if yname == "output-load" { return "OutputLoad" }
    if yname == "input-load" { return "InputLoad" }
    if yname == "reliability" { return "Reliability" }
    return ""
}

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetSegmentPath() string {
    return "data-rates"
}

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input-data-rate"] = dataRates.InputDataRate
    leafs["input-packet-rate"] = dataRates.InputPacketRate
    leafs["output-data-rate"] = dataRates.OutputDataRate
    leafs["output-packet-rate"] = dataRates.OutputPacketRate
    leafs["peak-input-data-rate"] = dataRates.PeakInputDataRate
    leafs["peak-input-packet-rate"] = dataRates.PeakInputPacketRate
    leafs["peak-output-data-rate"] = dataRates.PeakOutputDataRate
    leafs["peak-output-packet-rate"] = dataRates.PeakOutputPacketRate
    leafs["bandwidth"] = dataRates.Bandwidth
    leafs["load-interval"] = dataRates.LoadInterval
    leafs["output-load"] = dataRates.OutputLoad
    leafs["input-load"] = dataRates.InputLoad
    leafs["reliability"] = dataRates.Reliability
    return leafs
}

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetBundleName() string { return "cisco_ios_xr" }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetYangName() string { return "data-rates" }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) SetParent(parent types.Entity) { dataRates.parent = parent }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetParent() types.Entity { return dataRates.parent }

func (dataRates *Interfaces_InterfaceXr_Interface_DataRates) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_InterfaceStatistics
// Packet, byte and error counters
type Interfaces_InterfaceXr_Interface_InterfaceStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // StatsType. The type is ImCmdStatsEnum.
    StatsType interface{}

    // Packet, byte and all error counters.
    FullInterfaceStats Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats

    // Packet, byte and selected error counters.
    BasicInterfaceStats Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats
}

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetFilter() yfilter.YFilter { return interfaceStatistics.YFilter }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) SetFilter(yf yfilter.YFilter) { interfaceStatistics.YFilter = yf }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetGoName(yname string) string {
    if yname == "stats-type" { return "StatsType" }
    if yname == "full-interface-stats" { return "FullInterfaceStats" }
    if yname == "basic-interface-stats" { return "BasicInterfaceStats" }
    return ""
}

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetSegmentPath() string {
    return "interface-statistics"
}

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "full-interface-stats" {
        return &interfaceStatistics.FullInterfaceStats
    }
    if childYangName == "basic-interface-stats" {
        return &interfaceStatistics.BasicInterfaceStats
    }
    return nil
}

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["full-interface-stats"] = &interfaceStatistics.FullInterfaceStats
    children["basic-interface-stats"] = &interfaceStatistics.BasicInterfaceStats
    return children
}

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stats-type"] = interfaceStatistics.StatsType
    return leafs
}

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetYangName() string { return "interface-statistics" }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) SetParent(parent types.Entity) { interfaceStatistics.parent = parent }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetParent() types.Entity { return interfaceStatistics.parent }

func (interfaceStatistics *Interfaces_InterfaceXr_Interface_InterfaceStatistics) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats
// Packet, byte and all error counters
type Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats struct {
    parent types.Entity
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

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetFilter() yfilter.YFilter { return fullInterfaceStats.YFilter }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) SetFilter(yf yfilter.YFilter) { fullInterfaceStats.YFilter = yf }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "multicast-packets-received" { return "MulticastPacketsReceived" }
    if yname == "broadcast-packets-received" { return "BroadcastPacketsReceived" }
    if yname == "multicast-packets-sent" { return "MulticastPacketsSent" }
    if yname == "broadcast-packets-sent" { return "BroadcastPacketsSent" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "runt-packets-received" { return "RuntPacketsReceived" }
    if yname == "giant-packets-received" { return "GiantPacketsReceived" }
    if yname == "throttled-packets-received" { return "ThrottledPacketsReceived" }
    if yname == "parity-packets-received" { return "ParityPacketsReceived" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "crc-errors" { return "CrcErrors" }
    if yname == "input-overruns" { return "InputOverruns" }
    if yname == "framing-errors-received" { return "FramingErrorsReceived" }
    if yname == "input-ignored-packets" { return "InputIgnoredPackets" }
    if yname == "input-aborts" { return "InputAborts" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "output-underruns" { return "OutputUnderruns" }
    if yname == "output-buffer-failures" { return "OutputBufferFailures" }
    if yname == "output-buffers-swapped-out" { return "OutputBuffersSwappedOut" }
    if yname == "applique" { return "Applique" }
    if yname == "resets" { return "Resets" }
    if yname == "carrier-transitions" { return "CarrierTransitions" }
    if yname == "availability-flag" { return "AvailabilityFlag" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetSegmentPath() string {
    return "full-interface-stats"
}

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = fullInterfaceStats.PacketsReceived
    leafs["bytes-received"] = fullInterfaceStats.BytesReceived
    leafs["packets-sent"] = fullInterfaceStats.PacketsSent
    leafs["bytes-sent"] = fullInterfaceStats.BytesSent
    leafs["multicast-packets-received"] = fullInterfaceStats.MulticastPacketsReceived
    leafs["broadcast-packets-received"] = fullInterfaceStats.BroadcastPacketsReceived
    leafs["multicast-packets-sent"] = fullInterfaceStats.MulticastPacketsSent
    leafs["broadcast-packets-sent"] = fullInterfaceStats.BroadcastPacketsSent
    leafs["output-drops"] = fullInterfaceStats.OutputDrops
    leafs["output-queue-drops"] = fullInterfaceStats.OutputQueueDrops
    leafs["input-drops"] = fullInterfaceStats.InputDrops
    leafs["input-queue-drops"] = fullInterfaceStats.InputQueueDrops
    leafs["runt-packets-received"] = fullInterfaceStats.RuntPacketsReceived
    leafs["giant-packets-received"] = fullInterfaceStats.GiantPacketsReceived
    leafs["throttled-packets-received"] = fullInterfaceStats.ThrottledPacketsReceived
    leafs["parity-packets-received"] = fullInterfaceStats.ParityPacketsReceived
    leafs["unknown-protocol-packets-received"] = fullInterfaceStats.UnknownProtocolPacketsReceived
    leafs["input-errors"] = fullInterfaceStats.InputErrors
    leafs["crc-errors"] = fullInterfaceStats.CrcErrors
    leafs["input-overruns"] = fullInterfaceStats.InputOverruns
    leafs["framing-errors-received"] = fullInterfaceStats.FramingErrorsReceived
    leafs["input-ignored-packets"] = fullInterfaceStats.InputIgnoredPackets
    leafs["input-aborts"] = fullInterfaceStats.InputAborts
    leafs["output-errors"] = fullInterfaceStats.OutputErrors
    leafs["output-underruns"] = fullInterfaceStats.OutputUnderruns
    leafs["output-buffer-failures"] = fullInterfaceStats.OutputBufferFailures
    leafs["output-buffers-swapped-out"] = fullInterfaceStats.OutputBuffersSwappedOut
    leafs["applique"] = fullInterfaceStats.Applique
    leafs["resets"] = fullInterfaceStats.Resets
    leafs["carrier-transitions"] = fullInterfaceStats.CarrierTransitions
    leafs["availability-flag"] = fullInterfaceStats.AvailabilityFlag
    leafs["last-data-time"] = fullInterfaceStats.LastDataTime
    leafs["seconds-since-last-clear-counters"] = fullInterfaceStats.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = fullInterfaceStats.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = fullInterfaceStats.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = fullInterfaceStats.SecondsSincePacketSent
    return leafs
}

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetBundleName() string { return "cisco_ios_xr" }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetYangName() string { return "full-interface-stats" }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) SetParent(parent types.Entity) { fullInterfaceStats.parent = parent }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetParent() types.Entity { return fullInterfaceStats.parent }

func (fullInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_FullInterfaceStats) GetParentYangName() string { return "interface-statistics" }

// Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats
// Packet, byte and selected error counters
type Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats struct {
    parent types.Entity
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

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetFilter() yfilter.YFilter { return basicInterfaceStats.YFilter }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) SetFilter(yf yfilter.YFilter) { basicInterfaceStats.YFilter = yf }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "input-drops" { return "InputDrops" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "input-errors" { return "InputErrors" }
    if yname == "unknown-protocol-packets-received" { return "UnknownProtocolPacketsReceived" }
    if yname == "output-drops" { return "OutputDrops" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "output-errors" { return "OutputErrors" }
    if yname == "last-data-time" { return "LastDataTime" }
    if yname == "seconds-since-last-clear-counters" { return "SecondsSinceLastClearCounters" }
    if yname == "last-discontinuity-time" { return "LastDiscontinuityTime" }
    if yname == "seconds-since-packet-received" { return "SecondsSincePacketReceived" }
    if yname == "seconds-since-packet-sent" { return "SecondsSincePacketSent" }
    return ""
}

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetSegmentPath() string {
    return "basic-interface-stats"
}

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = basicInterfaceStats.PacketsReceived
    leafs["bytes-received"] = basicInterfaceStats.BytesReceived
    leafs["packets-sent"] = basicInterfaceStats.PacketsSent
    leafs["bytes-sent"] = basicInterfaceStats.BytesSent
    leafs["input-drops"] = basicInterfaceStats.InputDrops
    leafs["input-queue-drops"] = basicInterfaceStats.InputQueueDrops
    leafs["input-errors"] = basicInterfaceStats.InputErrors
    leafs["unknown-protocol-packets-received"] = basicInterfaceStats.UnknownProtocolPacketsReceived
    leafs["output-drops"] = basicInterfaceStats.OutputDrops
    leafs["output-queue-drops"] = basicInterfaceStats.OutputQueueDrops
    leafs["output-errors"] = basicInterfaceStats.OutputErrors
    leafs["last-data-time"] = basicInterfaceStats.LastDataTime
    leafs["seconds-since-last-clear-counters"] = basicInterfaceStats.SecondsSinceLastClearCounters
    leafs["last-discontinuity-time"] = basicInterfaceStats.LastDiscontinuityTime
    leafs["seconds-since-packet-received"] = basicInterfaceStats.SecondsSincePacketReceived
    leafs["seconds-since-packet-sent"] = basicInterfaceStats.SecondsSincePacketSent
    return leafs
}

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetBundleName() string { return "cisco_ios_xr" }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetYangName() string { return "basic-interface-stats" }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) SetParent(parent types.Entity) { basicInterfaceStats.parent = parent }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetParent() types.Entity { return basicInterfaceStats.parent }

func (basicInterfaceStats *Interfaces_InterfaceXr_Interface_InterfaceStatistics_BasicInterfaceStats) GetParentYangName() string { return "interface-statistics" }

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics
// L2 Protocol Statistics
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats type value. The type is interface{} with range: 0..4294967295.
    StatsType interface{}

    // Bag contents. The type is StatsTypeContents.
    Contents interface{}

    // Identifier.
    StatsId Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId

    // Block Array. The type is slice of
    // Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray.
    BlockArray []Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray

    // Element Array. The type is slice of
    // Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray.
    ElementArray []Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray
}

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetFilter() yfilter.YFilter { return l2InterfaceStatistics.YFilter }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) SetFilter(yf yfilter.YFilter) { l2InterfaceStatistics.YFilter = yf }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetGoName(yname string) string {
    if yname == "stats-type" { return "StatsType" }
    if yname == "contents" { return "Contents" }
    if yname == "stats-id" { return "StatsId" }
    if yname == "block-array" { return "BlockArray" }
    if yname == "element-array" { return "ElementArray" }
    return ""
}

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetSegmentPath() string {
    return "l2-interface-statistics"
}

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats-id" {
        return &l2InterfaceStatistics.StatsId
    }
    if childYangName == "block-array" {
        for _, c := range l2InterfaceStatistics.BlockArray {
            if l2InterfaceStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray{}
        l2InterfaceStatistics.BlockArray = append(l2InterfaceStatistics.BlockArray, child)
        return &l2InterfaceStatistics.BlockArray[len(l2InterfaceStatistics.BlockArray)-1]
    }
    if childYangName == "element-array" {
        for _, c := range l2InterfaceStatistics.ElementArray {
            if l2InterfaceStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray{}
        l2InterfaceStatistics.ElementArray = append(l2InterfaceStatistics.ElementArray, child)
        return &l2InterfaceStatistics.ElementArray[len(l2InterfaceStatistics.ElementArray)-1]
    }
    return nil
}

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats-id"] = &l2InterfaceStatistics.StatsId
    for i := range l2InterfaceStatistics.BlockArray {
        children[l2InterfaceStatistics.BlockArray[i].GetSegmentPath()] = &l2InterfaceStatistics.BlockArray[i]
    }
    for i := range l2InterfaceStatistics.ElementArray {
        children[l2InterfaceStatistics.ElementArray[i].GetSegmentPath()] = &l2InterfaceStatistics.ElementArray[i]
    }
    return children
}

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stats-type"] = l2InterfaceStatistics.StatsType
    leafs["contents"] = l2InterfaceStatistics.Contents
    return leafs
}

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetYangName() string { return "l2-interface-statistics" }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) SetParent(parent types.Entity) { l2InterfaceStatistics.parent = parent }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetParent() types.Entity { return l2InterfaceStatistics.parent }

func (l2InterfaceStatistics *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics) GetParentYangName() string { return "interface" }

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId
// Identifier
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // id type. The type is StatsId.
    IdType interface{}

    // Unused. The type is interface{} with range: 0..4294967295.
    Unused interface{}

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Feature ID. The type is interface{} with range: 0..4294967295.
    FeatureId interface{}

    // ID. The type is interface{} with range: 0..4294967295.
    Id interface{}
}

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetFilter() yfilter.YFilter { return statsId.YFilter }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) SetFilter(yf yfilter.YFilter) { statsId.YFilter = yf }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetGoName(yname string) string {
    if yname == "id-type" { return "IdType" }
    if yname == "unused" { return "Unused" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "node-id" { return "NodeId" }
    if yname == "feature-id" { return "FeatureId" }
    if yname == "id" { return "Id" }
    return ""
}

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetSegmentPath() string {
    return "stats-id"
}

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id-type"] = statsId.IdType
    leafs["unused"] = statsId.Unused
    leafs["interface-handle"] = statsId.InterfaceHandle
    leafs["node-id"] = statsId.NodeId
    leafs["feature-id"] = statsId.FeatureId
    leafs["id"] = statsId.Id
    return leafs
}

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetBundleName() string { return "cisco_ios_xr" }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetYangName() string { return "stats-id" }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) SetParent(parent types.Entity) { statsId.parent = parent }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetParent() types.Entity { return statsId.parent }

func (statsId *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_StatsId) GetParentYangName() string { return "l2-interface-statistics" }

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray
// Block Array
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is StatsCounter.
    Type interface{}

    // count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // data. The type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Data interface{}
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetFilter() yfilter.YFilter { return blockArray.YFilter }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) SetFilter(yf yfilter.YFilter) { blockArray.YFilter = yf }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "count" { return "Count" }
    if yname == "data" { return "Data" }
    return ""
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetSegmentPath() string {
    return "block-array"
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = blockArray.Type
    leafs["count"] = blockArray.Count
    leafs["data"] = blockArray.Data
    return leafs
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetBundleName() string { return "cisco_ios_xr" }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetYangName() string { return "block-array" }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) SetParent(parent types.Entity) { blockArray.parent = parent }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetParent() types.Entity { return blockArray.parent }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_BlockArray) GetParentYangName() string { return "l2-interface-statistics" }

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray
// Element Array
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // key. The type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Key interface{}

    // block array. The type is slice of
    // Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray.
    BlockArray []Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray
}

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetFilter() yfilter.YFilter { return elementArray.YFilter }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) SetFilter(yf yfilter.YFilter) { elementArray.YFilter = yf }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetGoName(yname string) string {
    if yname == "key" { return "Key" }
    if yname == "block-array" { return "BlockArray" }
    return ""
}

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetSegmentPath() string {
    return "element-array"
}

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "block-array" {
        for _, c := range elementArray.BlockArray {
            if elementArray.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray{}
        elementArray.BlockArray = append(elementArray.BlockArray, child)
        return &elementArray.BlockArray[len(elementArray.BlockArray)-1]
    }
    return nil
}

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range elementArray.BlockArray {
        children[elementArray.BlockArray[i].GetSegmentPath()] = &elementArray.BlockArray[i]
    }
    return children
}

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key"] = elementArray.Key
    return leafs
}

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetBundleName() string { return "cisco_ios_xr" }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetYangName() string { return "element-array" }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) SetParent(parent types.Entity) { elementArray.parent = parent }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetParent() types.Entity { return elementArray.parent }

func (elementArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray) GetParentYangName() string { return "l2-interface-statistics" }

// Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray
// block array
type Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is StatsCounter.
    Type interface{}

    // count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // data. The type is string with pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Data interface{}
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetFilter() yfilter.YFilter { return blockArray.YFilter }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) SetFilter(yf yfilter.YFilter) { blockArray.YFilter = yf }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "count" { return "Count" }
    if yname == "data" { return "Data" }
    return ""
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetSegmentPath() string {
    return "block-array"
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = blockArray.Type
    leafs["count"] = blockArray.Count
    leafs["data"] = blockArray.Data
    return leafs
}

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetBundleName() string { return "cisco_ios_xr" }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetYangName() string { return "block-array" }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) SetParent(parent types.Entity) { blockArray.parent = parent }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetParent() types.Entity { return blockArray.parent }

func (blockArray *Interfaces_InterfaceXr_Interface_L2InterfaceStatistics_ElementArray_BlockArray) GetParentYangName() string { return "element-array" }

// Interfaces_InterfaceXr_Interface_NvOptical
// nV Optical Controller Information
type Interfaces_InterfaceXr_Interface_NvOptical struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Controller that nV controller maps to. The type is string.
    Controller interface{}
}

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetFilter() yfilter.YFilter { return nvOptical.YFilter }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) SetFilter(yf yfilter.YFilter) { nvOptical.YFilter = yf }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetGoName(yname string) string {
    if yname == "controller" { return "Controller" }
    return ""
}

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetSegmentPath() string {
    return "nv-optical"
}

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["controller"] = nvOptical.Controller
    return leafs
}

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetBundleName() string { return "cisco_ios_xr" }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetYangName() string { return "nv-optical" }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) SetParent(parent types.Entity) { nvOptical.parent = parent }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetParent() types.Entity { return nvOptical.parent }

func (nvOptical *Interfaces_InterfaceXr_Interface_NvOptical) GetParentYangName() string { return "interface" }

// Interfaces_NodeTypeSets
// Node and/or interface type specific view of
// interface summary data
type Interfaces_NodeTypeSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary data for all interfaces on a particular node. The type is slice of
    // Interfaces_NodeTypeSets_NodeTypeSet.
    NodeTypeSet []Interfaces_NodeTypeSets_NodeTypeSet
}

func (nodeTypeSets *Interfaces_NodeTypeSets) GetFilter() yfilter.YFilter { return nodeTypeSets.YFilter }

func (nodeTypeSets *Interfaces_NodeTypeSets) SetFilter(yf yfilter.YFilter) { nodeTypeSets.YFilter = yf }

func (nodeTypeSets *Interfaces_NodeTypeSets) GetGoName(yname string) string {
    if yname == "node-type-set" { return "NodeTypeSet" }
    return ""
}

func (nodeTypeSets *Interfaces_NodeTypeSets) GetSegmentPath() string {
    return "node-type-sets"
}

func (nodeTypeSets *Interfaces_NodeTypeSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-type-set" {
        for _, c := range nodeTypeSets.NodeTypeSet {
            if nodeTypeSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_NodeTypeSets_NodeTypeSet{}
        nodeTypeSets.NodeTypeSet = append(nodeTypeSets.NodeTypeSet, child)
        return &nodeTypeSets.NodeTypeSet[len(nodeTypeSets.NodeTypeSet)-1]
    }
    return nil
}

func (nodeTypeSets *Interfaces_NodeTypeSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeTypeSets.NodeTypeSet {
        children[nodeTypeSets.NodeTypeSet[i].GetSegmentPath()] = &nodeTypeSets.NodeTypeSet[i]
    }
    return children
}

func (nodeTypeSets *Interfaces_NodeTypeSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeTypeSets *Interfaces_NodeTypeSets) GetBundleName() string { return "cisco_ios_xr" }

func (nodeTypeSets *Interfaces_NodeTypeSets) GetYangName() string { return "node-type-sets" }

func (nodeTypeSets *Interfaces_NodeTypeSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeTypeSets *Interfaces_NodeTypeSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeTypeSets *Interfaces_NodeTypeSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeTypeSets *Interfaces_NodeTypeSets) SetParent(parent types.Entity) { nodeTypeSets.parent = parent }

func (nodeTypeSets *Interfaces_NodeTypeSets) GetParent() types.Entity { return nodeTypeSets.parent }

func (nodeTypeSets *Interfaces_NodeTypeSets) GetParentYangName() string { return "interfaces" }

// Interfaces_NodeTypeSets_NodeTypeSet
// Summary data for all interfaces on a particular
// node
type Interfaces_NodeTypeSets_NodeTypeSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The location to filter on. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // The interface type to filter on. The type is InterfaceTypeSet.
    TypeSetName interface{}

    // Interface summary information.
    InterfaceSummary Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary
}

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetFilter() yfilter.YFilter { return nodeTypeSet.YFilter }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) SetFilter(yf yfilter.YFilter) { nodeTypeSet.YFilter = yf }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "type-set-name" { return "TypeSetName" }
    if yname == "interface-summary" { return "InterfaceSummary" }
    return ""
}

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetSegmentPath() string {
    return "node-type-set"
}

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-summary" {
        return &nodeTypeSet.InterfaceSummary
    }
    return nil
}

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-summary"] = &nodeTypeSet.InterfaceSummary
    return children
}

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodeTypeSet.NodeName
    leafs["type-set-name"] = nodeTypeSet.TypeSetName
    return leafs
}

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetBundleName() string { return "cisco_ios_xr" }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetYangName() string { return "node-type-set" }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) SetParent(parent types.Entity) { nodeTypeSet.parent = parent }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetParent() types.Entity { return nodeTypeSet.parent }

func (nodeTypeSet *Interfaces_NodeTypeSets_NodeTypeSet) GetParentYangName() string { return "node-type-sets" }

// Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary
// Interface summary information
type Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Counts for all interfaces.
    InterfaceCounts Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts

    // List of per interface type summary information. The type is slice of
    // Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType.
    InterfaceType []Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType
}

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetFilter() yfilter.YFilter { return interfaceSummary.YFilter }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) SetFilter(yf yfilter.YFilter) { interfaceSummary.YFilter = yf }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetGoName(yname string) string {
    if yname == "interface-counts" { return "InterfaceCounts" }
    if yname == "interface-type" { return "InterfaceType" }
    return ""
}

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetSegmentPath() string {
    return "interface-summary"
}

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-counts" {
        return &interfaceSummary.InterfaceCounts
    }
    if childYangName == "interface-type" {
        for _, c := range interfaceSummary.InterfaceType {
            if interfaceSummary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType{}
        interfaceSummary.InterfaceType = append(interfaceSummary.InterfaceType, child)
        return &interfaceSummary.InterfaceType[len(interfaceSummary.InterfaceType)-1]
    }
    return nil
}

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-counts"] = &interfaceSummary.InterfaceCounts
    for i := range interfaceSummary.InterfaceType {
        children[interfaceSummary.InterfaceType[i].GetSegmentPath()] = &interfaceSummary.InterfaceType[i]
    }
    return children
}

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetYangName() string { return "interface-summary" }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) SetParent(parent types.Entity) { interfaceSummary.parent = parent }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetParent() types.Entity { return interfaceSummary.parent }

func (interfaceSummary *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary) GetParentYangName() string { return "node-type-set" }

// Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts
// Counts for all interfaces
type Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts struct {
    parent types.Entity
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

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetFilter() yfilter.YFilter { return interfaceCounts.YFilter }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) SetFilter(yf yfilter.YFilter) { interfaceCounts.YFilter = yf }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetGoName(yname string) string {
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "up-interface-count" { return "UpInterfaceCount" }
    if yname == "down-interface-count" { return "DownInterfaceCount" }
    if yname == "admin-down-interface-count" { return "AdminDownInterfaceCount" }
    return ""
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetSegmentPath() string {
    return "interface-counts"
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-count"] = interfaceCounts.InterfaceCount
    leafs["up-interface-count"] = interfaceCounts.UpInterfaceCount
    leafs["down-interface-count"] = interfaceCounts.DownInterfaceCount
    leafs["admin-down-interface-count"] = interfaceCounts.AdminDownInterfaceCount
    return leafs
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetYangName() string { return "interface-counts" }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) SetParent(parent types.Entity) { interfaceCounts.parent = parent }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetParent() types.Entity { return interfaceCounts.parent }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceCounts) GetParentYangName() string { return "interface-summary" }

// Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType
// List of per interface type summary information
type Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the interface type. The type is string.
    InterfaceTypeName interface{}

    // Description of the interface type. The type is string.
    InterfaceTypeDescription interface{}

    // Counts for interfaces of this type.
    InterfaceCounts Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts
}

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetFilter() yfilter.YFilter { return interfaceType.YFilter }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) SetFilter(yf yfilter.YFilter) { interfaceType.YFilter = yf }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetGoName(yname string) string {
    if yname == "interface-type-name" { return "InterfaceTypeName" }
    if yname == "interface-type-description" { return "InterfaceTypeDescription" }
    if yname == "interface-counts" { return "InterfaceCounts" }
    return ""
}

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetSegmentPath() string {
    return "interface-type"
}

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-counts" {
        return &interfaceType.InterfaceCounts
    }
    return nil
}

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-counts"] = &interfaceType.InterfaceCounts
    return children
}

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-type-name"] = interfaceType.InterfaceTypeName
    leafs["interface-type-description"] = interfaceType.InterfaceTypeDescription
    return leafs
}

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetYangName() string { return "interface-type" }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) SetParent(parent types.Entity) { interfaceType.parent = parent }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetParent() types.Entity { return interfaceType.parent }

func (interfaceType *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType) GetParentYangName() string { return "interface-summary" }

// Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts
// Counts for interfaces of this type
type Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts struct {
    parent types.Entity
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

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetFilter() yfilter.YFilter { return interfaceCounts.YFilter }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) SetFilter(yf yfilter.YFilter) { interfaceCounts.YFilter = yf }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetGoName(yname string) string {
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "up-interface-count" { return "UpInterfaceCount" }
    if yname == "down-interface-count" { return "DownInterfaceCount" }
    if yname == "admin-down-interface-count" { return "AdminDownInterfaceCount" }
    return ""
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetSegmentPath() string {
    return "interface-counts"
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-count"] = interfaceCounts.InterfaceCount
    leafs["up-interface-count"] = interfaceCounts.UpInterfaceCount
    leafs["down-interface-count"] = interfaceCounts.DownInterfaceCount
    leafs["admin-down-interface-count"] = interfaceCounts.AdminDownInterfaceCount
    return leafs
}

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetYangName() string { return "interface-counts" }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) SetParent(parent types.Entity) { interfaceCounts.parent = parent }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetParent() types.Entity { return interfaceCounts.parent }

func (interfaceCounts *Interfaces_NodeTypeSets_NodeTypeSet_InterfaceSummary_InterfaceType_InterfaceCounts) GetParentYangName() string { return "interface-type" }

// Interfaces_InterfaceBriefs
// Brief operational data for interfaces
type Interfaces_InterfaceBriefs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief operational attributes for a particular interface. The type is slice
    // of Interfaces_InterfaceBriefs_InterfaceBrief.
    InterfaceBrief []Interfaces_InterfaceBriefs_InterfaceBrief
}

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetFilter() yfilter.YFilter { return interfaceBriefs.YFilter }

func (interfaceBriefs *Interfaces_InterfaceBriefs) SetFilter(yf yfilter.YFilter) { interfaceBriefs.YFilter = yf }

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetGoName(yname string) string {
    if yname == "interface-brief" { return "InterfaceBrief" }
    return ""
}

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetSegmentPath() string {
    return "interface-briefs"
}

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-brief" {
        for _, c := range interfaceBriefs.InterfaceBrief {
            if interfaceBriefs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceBriefs_InterfaceBrief{}
        interfaceBriefs.InterfaceBrief = append(interfaceBriefs.InterfaceBrief, child)
        return &interfaceBriefs.InterfaceBrief[len(interfaceBriefs.InterfaceBrief)-1]
    }
    return nil
}

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceBriefs.InterfaceBrief {
        children[interfaceBriefs.InterfaceBrief[i].GetSegmentPath()] = &interfaceBriefs.InterfaceBrief[i]
    }
    return children
}

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetYangName() string { return "interface-briefs" }

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceBriefs *Interfaces_InterfaceBriefs) SetParent(parent types.Entity) { interfaceBriefs.parent = parent }

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetParent() types.Entity { return interfaceBriefs.parent }

func (interfaceBriefs *Interfaces_InterfaceBriefs) GetParentYangName() string { return "interfaces" }

// Interfaces_InterfaceBriefs_InterfaceBrief
// Brief operational attributes for a particular
// interface
type Interfaces_InterfaceBriefs_InterfaceBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Parent Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetFilter() yfilter.YFilter { return interfaceBrief.YFilter }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) SetFilter(yf yfilter.YFilter) { interfaceBrief.YFilter = yf }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "parent-interface" { return "ParentInterface" }
    if yname == "type" { return "Type" }
    if yname == "state" { return "State" }
    if yname == "actual-state" { return "ActualState" }
    if yname == "line-state" { return "LineState" }
    if yname == "actual-line-state" { return "ActualLineState" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "encapsulation-type-string" { return "EncapsulationTypeString" }
    if yname == "mtu" { return "Mtu" }
    if yname == "sub-interface-mtu-overhead" { return "SubInterfaceMtuOverhead" }
    if yname == "l2-transport" { return "L2Transport" }
    if yname == "bandwidth" { return "Bandwidth" }
    return ""
}

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetSegmentPath() string {
    return "interface-brief" + "[interface-name='" + fmt.Sprintf("%v", interfaceBrief.InterfaceName) + "']"
}

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceBrief.InterfaceName
    leafs["interface"] = interfaceBrief.Interface
    leafs["parent-interface"] = interfaceBrief.ParentInterface
    leafs["type"] = interfaceBrief.Type
    leafs["state"] = interfaceBrief.State
    leafs["actual-state"] = interfaceBrief.ActualState
    leafs["line-state"] = interfaceBrief.LineState
    leafs["actual-line-state"] = interfaceBrief.ActualLineState
    leafs["encapsulation"] = interfaceBrief.Encapsulation
    leafs["encapsulation-type-string"] = interfaceBrief.EncapsulationTypeString
    leafs["mtu"] = interfaceBrief.Mtu
    leafs["sub-interface-mtu-overhead"] = interfaceBrief.SubInterfaceMtuOverhead
    leafs["l2-transport"] = interfaceBrief.L2Transport
    leafs["bandwidth"] = interfaceBrief.Bandwidth
    return leafs
}

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetYangName() string { return "interface-brief" }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) SetParent(parent types.Entity) { interfaceBrief.parent = parent }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetParent() types.Entity { return interfaceBrief.parent }

func (interfaceBrief *Interfaces_InterfaceBriefs_InterfaceBrief) GetParentYangName() string { return "interface-briefs" }

// Interfaces_InventorySummary
// Inventory summary information
type Interfaces_InventorySummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Counts for all interfaces.
    InterfaceCounts Interfaces_InventorySummary_InterfaceCounts

    // List of per interface type summary information. The type is slice of
    // Interfaces_InventorySummary_InterfaceType.
    InterfaceType []Interfaces_InventorySummary_InterfaceType
}

func (inventorySummary *Interfaces_InventorySummary) GetFilter() yfilter.YFilter { return inventorySummary.YFilter }

func (inventorySummary *Interfaces_InventorySummary) SetFilter(yf yfilter.YFilter) { inventorySummary.YFilter = yf }

func (inventorySummary *Interfaces_InventorySummary) GetGoName(yname string) string {
    if yname == "interface-counts" { return "InterfaceCounts" }
    if yname == "interface-type" { return "InterfaceType" }
    return ""
}

func (inventorySummary *Interfaces_InventorySummary) GetSegmentPath() string {
    return "inventory-summary"
}

func (inventorySummary *Interfaces_InventorySummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-counts" {
        return &inventorySummary.InterfaceCounts
    }
    if childYangName == "interface-type" {
        for _, c := range inventorySummary.InterfaceType {
            if inventorySummary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InventorySummary_InterfaceType{}
        inventorySummary.InterfaceType = append(inventorySummary.InterfaceType, child)
        return &inventorySummary.InterfaceType[len(inventorySummary.InterfaceType)-1]
    }
    return nil
}

func (inventorySummary *Interfaces_InventorySummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-counts"] = &inventorySummary.InterfaceCounts
    for i := range inventorySummary.InterfaceType {
        children[inventorySummary.InterfaceType[i].GetSegmentPath()] = &inventorySummary.InterfaceType[i]
    }
    return children
}

func (inventorySummary *Interfaces_InventorySummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inventorySummary *Interfaces_InventorySummary) GetBundleName() string { return "cisco_ios_xr" }

func (inventorySummary *Interfaces_InventorySummary) GetYangName() string { return "inventory-summary" }

func (inventorySummary *Interfaces_InventorySummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventorySummary *Interfaces_InventorySummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventorySummary *Interfaces_InventorySummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventorySummary *Interfaces_InventorySummary) SetParent(parent types.Entity) { inventorySummary.parent = parent }

func (inventorySummary *Interfaces_InventorySummary) GetParent() types.Entity { return inventorySummary.parent }

func (inventorySummary *Interfaces_InventorySummary) GetParentYangName() string { return "interfaces" }

// Interfaces_InventorySummary_InterfaceCounts
// Counts for all interfaces
type Interfaces_InventorySummary_InterfaceCounts struct {
    parent types.Entity
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

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetFilter() yfilter.YFilter { return interfaceCounts.YFilter }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) SetFilter(yf yfilter.YFilter) { interfaceCounts.YFilter = yf }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetGoName(yname string) string {
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "up-interface-count" { return "UpInterfaceCount" }
    if yname == "down-interface-count" { return "DownInterfaceCount" }
    if yname == "admin-down-interface-count" { return "AdminDownInterfaceCount" }
    return ""
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetSegmentPath() string {
    return "interface-counts"
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-count"] = interfaceCounts.InterfaceCount
    leafs["up-interface-count"] = interfaceCounts.UpInterfaceCount
    leafs["down-interface-count"] = interfaceCounts.DownInterfaceCount
    leafs["admin-down-interface-count"] = interfaceCounts.AdminDownInterfaceCount
    return leafs
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetYangName() string { return "interface-counts" }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) SetParent(parent types.Entity) { interfaceCounts.parent = parent }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetParent() types.Entity { return interfaceCounts.parent }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceCounts) GetParentYangName() string { return "inventory-summary" }

// Interfaces_InventorySummary_InterfaceType
// List of per interface type summary information
type Interfaces_InventorySummary_InterfaceType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the interface type. The type is string.
    InterfaceTypeName interface{}

    // Description of the interface type. The type is string.
    InterfaceTypeDescription interface{}

    // Counts for interfaces of this type.
    InterfaceCounts Interfaces_InventorySummary_InterfaceType_InterfaceCounts
}

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetFilter() yfilter.YFilter { return interfaceType.YFilter }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) SetFilter(yf yfilter.YFilter) { interfaceType.YFilter = yf }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetGoName(yname string) string {
    if yname == "interface-type-name" { return "InterfaceTypeName" }
    if yname == "interface-type-description" { return "InterfaceTypeDescription" }
    if yname == "interface-counts" { return "InterfaceCounts" }
    return ""
}

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetSegmentPath() string {
    return "interface-type"
}

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-counts" {
        return &interfaceType.InterfaceCounts
    }
    return nil
}

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-counts"] = &interfaceType.InterfaceCounts
    return children
}

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-type-name"] = interfaceType.InterfaceTypeName
    leafs["interface-type-description"] = interfaceType.InterfaceTypeDescription
    return leafs
}

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetYangName() string { return "interface-type" }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) SetParent(parent types.Entity) { interfaceType.parent = parent }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetParent() types.Entity { return interfaceType.parent }

func (interfaceType *Interfaces_InventorySummary_InterfaceType) GetParentYangName() string { return "inventory-summary" }

// Interfaces_InventorySummary_InterfaceType_InterfaceCounts
// Counts for interfaces of this type
type Interfaces_InventorySummary_InterfaceType_InterfaceCounts struct {
    parent types.Entity
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

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetFilter() yfilter.YFilter { return interfaceCounts.YFilter }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) SetFilter(yf yfilter.YFilter) { interfaceCounts.YFilter = yf }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetGoName(yname string) string {
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "up-interface-count" { return "UpInterfaceCount" }
    if yname == "down-interface-count" { return "DownInterfaceCount" }
    if yname == "admin-down-interface-count" { return "AdminDownInterfaceCount" }
    return ""
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetSegmentPath() string {
    return "interface-counts"
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-count"] = interfaceCounts.InterfaceCount
    leafs["up-interface-count"] = interfaceCounts.UpInterfaceCount
    leafs["down-interface-count"] = interfaceCounts.DownInterfaceCount
    leafs["admin-down-interface-count"] = interfaceCounts.AdminDownInterfaceCount
    return leafs
}

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetYangName() string { return "interface-counts" }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) SetParent(parent types.Entity) { interfaceCounts.parent = parent }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetParent() types.Entity { return interfaceCounts.parent }

func (interfaceCounts *Interfaces_InventorySummary_InterfaceType_InterfaceCounts) GetParentYangName() string { return "interface-type" }

// Interfaces_Interfaces
// Descriptions for interfaces
type Interfaces_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Description for a particular interface. The type is slice of
    // Interfaces_Interfaces_Interface.
    Interface []Interfaces_Interfaces_Interface
}

func (interfaces *Interfaces_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Interfaces_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Interfaces_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Interfaces_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Interfaces_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Interfaces_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Interfaces_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Interfaces_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Interfaces_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Interfaces_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Interfaces_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Interfaces_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Interfaces_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Interfaces_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Interfaces_Interfaces) GetParentYangName() string { return "interfaces" }

// Interfaces_Interfaces_Interface
// Description for a particular interface
type Interfaces_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (self *Interfaces_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Interfaces_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Interfaces_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "state" { return "State" }
    if yname == "line-state" { return "LineState" }
    if yname == "description" { return "Description" }
    return ""
}

func (self *Interfaces_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Interfaces_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Interfaces_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Interfaces_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["state"] = self.State
    leafs["line-state"] = self.LineState
    leafs["description"] = self.Description
    return leafs
}

func (self *Interfaces_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Interfaces_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Interfaces_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Interfaces_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Interfaces_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Interfaces_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Interfaces_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Interfaces_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Interfaces_InterfaceSummary
// Interface summary information
type Interfaces_InterfaceSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Counts for all interfaces.
    InterfaceCounts Interfaces_InterfaceSummary_InterfaceCounts

    // List of per interface type summary information. The type is slice of
    // Interfaces_InterfaceSummary_InterfaceType.
    InterfaceType []Interfaces_InterfaceSummary_InterfaceType
}

func (interfaceSummary *Interfaces_InterfaceSummary) GetFilter() yfilter.YFilter { return interfaceSummary.YFilter }

func (interfaceSummary *Interfaces_InterfaceSummary) SetFilter(yf yfilter.YFilter) { interfaceSummary.YFilter = yf }

func (interfaceSummary *Interfaces_InterfaceSummary) GetGoName(yname string) string {
    if yname == "interface-counts" { return "InterfaceCounts" }
    if yname == "interface-type" { return "InterfaceType" }
    return ""
}

func (interfaceSummary *Interfaces_InterfaceSummary) GetSegmentPath() string {
    return "interface-summary"
}

func (interfaceSummary *Interfaces_InterfaceSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-counts" {
        return &interfaceSummary.InterfaceCounts
    }
    if childYangName == "interface-type" {
        for _, c := range interfaceSummary.InterfaceType {
            if interfaceSummary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_InterfaceSummary_InterfaceType{}
        interfaceSummary.InterfaceType = append(interfaceSummary.InterfaceType, child)
        return &interfaceSummary.InterfaceType[len(interfaceSummary.InterfaceType)-1]
    }
    return nil
}

func (interfaceSummary *Interfaces_InterfaceSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-counts"] = &interfaceSummary.InterfaceCounts
    for i := range interfaceSummary.InterfaceType {
        children[interfaceSummary.InterfaceType[i].GetSegmentPath()] = &interfaceSummary.InterfaceType[i]
    }
    return children
}

func (interfaceSummary *Interfaces_InterfaceSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSummary *Interfaces_InterfaceSummary) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSummary *Interfaces_InterfaceSummary) GetYangName() string { return "interface-summary" }

func (interfaceSummary *Interfaces_InterfaceSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSummary *Interfaces_InterfaceSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSummary *Interfaces_InterfaceSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSummary *Interfaces_InterfaceSummary) SetParent(parent types.Entity) { interfaceSummary.parent = parent }

func (interfaceSummary *Interfaces_InterfaceSummary) GetParent() types.Entity { return interfaceSummary.parent }

func (interfaceSummary *Interfaces_InterfaceSummary) GetParentYangName() string { return "interfaces" }

// Interfaces_InterfaceSummary_InterfaceCounts
// Counts for all interfaces
type Interfaces_InterfaceSummary_InterfaceCounts struct {
    parent types.Entity
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

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetFilter() yfilter.YFilter { return interfaceCounts.YFilter }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) SetFilter(yf yfilter.YFilter) { interfaceCounts.YFilter = yf }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetGoName(yname string) string {
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "up-interface-count" { return "UpInterfaceCount" }
    if yname == "down-interface-count" { return "DownInterfaceCount" }
    if yname == "admin-down-interface-count" { return "AdminDownInterfaceCount" }
    return ""
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetSegmentPath() string {
    return "interface-counts"
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-count"] = interfaceCounts.InterfaceCount
    leafs["up-interface-count"] = interfaceCounts.UpInterfaceCount
    leafs["down-interface-count"] = interfaceCounts.DownInterfaceCount
    leafs["admin-down-interface-count"] = interfaceCounts.AdminDownInterfaceCount
    return leafs
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetYangName() string { return "interface-counts" }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) SetParent(parent types.Entity) { interfaceCounts.parent = parent }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetParent() types.Entity { return interfaceCounts.parent }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceCounts) GetParentYangName() string { return "interface-summary" }

// Interfaces_InterfaceSummary_InterfaceType
// List of per interface type summary information
type Interfaces_InterfaceSummary_InterfaceType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the interface type. The type is string.
    InterfaceTypeName interface{}

    // Description of the interface type. The type is string.
    InterfaceTypeDescription interface{}

    // Counts for interfaces of this type.
    InterfaceCounts Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts
}

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetFilter() yfilter.YFilter { return interfaceType.YFilter }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) SetFilter(yf yfilter.YFilter) { interfaceType.YFilter = yf }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetGoName(yname string) string {
    if yname == "interface-type-name" { return "InterfaceTypeName" }
    if yname == "interface-type-description" { return "InterfaceTypeDescription" }
    if yname == "interface-counts" { return "InterfaceCounts" }
    return ""
}

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetSegmentPath() string {
    return "interface-type"
}

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-counts" {
        return &interfaceType.InterfaceCounts
    }
    return nil
}

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-counts"] = &interfaceType.InterfaceCounts
    return children
}

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-type-name"] = interfaceType.InterfaceTypeName
    leafs["interface-type-description"] = interfaceType.InterfaceTypeDescription
    return leafs
}

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetYangName() string { return "interface-type" }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) SetParent(parent types.Entity) { interfaceType.parent = parent }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetParent() types.Entity { return interfaceType.parent }

func (interfaceType *Interfaces_InterfaceSummary_InterfaceType) GetParentYangName() string { return "interface-summary" }

// Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts
// Counts for interfaces of this type
type Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts struct {
    parent types.Entity
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

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetFilter() yfilter.YFilter { return interfaceCounts.YFilter }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) SetFilter(yf yfilter.YFilter) { interfaceCounts.YFilter = yf }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetGoName(yname string) string {
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "up-interface-count" { return "UpInterfaceCount" }
    if yname == "down-interface-count" { return "DownInterfaceCount" }
    if yname == "admin-down-interface-count" { return "AdminDownInterfaceCount" }
    return ""
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetSegmentPath() string {
    return "interface-counts"
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-count"] = interfaceCounts.InterfaceCount
    leafs["up-interface-count"] = interfaceCounts.UpInterfaceCount
    leafs["down-interface-count"] = interfaceCounts.DownInterfaceCount
    leafs["admin-down-interface-count"] = interfaceCounts.AdminDownInterfaceCount
    return leafs
}

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetYangName() string { return "interface-counts" }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) SetParent(parent types.Entity) { interfaceCounts.parent = parent }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetParent() types.Entity { return interfaceCounts.parent }

func (interfaceCounts *Interfaces_InterfaceSummary_InterfaceType_InterfaceCounts) GetParentYangName() string { return "interface-type" }

