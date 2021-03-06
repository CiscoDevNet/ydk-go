// This module contains a collection of YANG definitions
// for Cisco IOS-XR installmgr package
// admin-plane operational data.
// 
// This module contains definitions
// for the following management objects:
//   install: Information of software packages and install
//     operations
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package installmgr_admin_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package installmgr_admin_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-installmgr-admin-oper install}", reflect.TypeOf(Install{}))
    ydk.RegisterEntity("Cisco-IOS-XR-installmgr-admin-oper:install", reflect.TypeOf(Install{}))
}

// InstmgrIssuAbortMethod represents Abort method
type InstmgrIssuAbortMethod string

const (
    // Unknown abort method
    InstmgrIssuAbortMethod_method_undefined InstmgrIssuAbortMethod = "method-undefined"

    // No abort is required
    InstmgrIssuAbortMethod_method_no_operation InstmgrIssuAbortMethod = "method-no-operation"

    // Abort will reload standby nodes
    InstmgrIssuAbortMethod_method_standby_reload InstmgrIssuAbortMethod = "method-standby-reload"

    // Abort will reload the whole system
    InstmgrIssuAbortMethod_method_system_reload InstmgrIssuAbortMethod = "method-system-reload"

    // Abort will rollback
    InstmgrIssuAbortMethod_method_rollback InstmgrIssuAbortMethod = "method-rollback"

    // Abort is not possible
    InstmgrIssuAbortMethod_method_not_possible InstmgrIssuAbortMethod = "method-not-possible"

    // Abort is not possible by SDR user
    InstmgrIssuAbortMethod_method_admin_only InstmgrIssuAbortMethod = "method-admin-only"
)

// InstmgrBagRequestTrigger represents The trigger type of an install request
type InstmgrBagRequestTrigger string

const (
    // Request triggered by CLI
    InstmgrBagRequestTrigger_cli InstmgrBagRequestTrigger = "cli"

    // Request triggered by XML
    InstmgrBagRequestTrigger_xr_xml InstmgrBagRequestTrigger = "xr-xml"
)

// InstmgrGroup represents Group type
type InstmgrGroup string

const (
    // Undefined grouping
    InstmgrGroup_inst_pkg_group_undefined InstmgrGroup = "inst-pkg-group-undefined"

    // Packages are grouped
    InstmgrGroup_inst_pkg_group_grouped InstmgrGroup = "inst-pkg-group-grouped"

    // Packages are all individual
    InstmgrGroup_inst_pkg_group_individual InstmgrGroup = "inst-pkg-group-individual"
)

// IsmCardTypeFamily represents Ism card type family
type IsmCardTypeFamily string

const (
    // Active NDSC RPs
    IsmCardTypeFamily_ndsc_active_rp IsmCardTypeFamily = "ndsc-active-rp"

    // Standby NDSC RPs
    IsmCardTypeFamily_ndsc_standby_rp IsmCardTypeFamily = "ndsc-standby-rp"

    // Active DRP
    IsmCardTypeFamily_active_drp IsmCardTypeFamily = "active-drp"

    // Standby DRP
    IsmCardTypeFamily_standby_drp IsmCardTypeFamily = "standby-drp"

    // Active dSC
    IsmCardTypeFamily_dsc_active_rp IsmCardTypeFamily = "dsc-active-rp"

    // Standby dSC
    IsmCardTypeFamily_dsc_standby_rp IsmCardTypeFamily = "dsc-standby-rp"

    // Active non-root SC
    IsmCardTypeFamily_active_sc IsmCardTypeFamily = "active-sc"

    // Standby non-root SC
    IsmCardTypeFamily_standby_sc IsmCardTypeFamily = "standby-sc"

    // Active root SC
    IsmCardTypeFamily_root_active_sc IsmCardTypeFamily = "root-active-sc"

    // Standby root SC
    IsmCardTypeFamily_root_standby_sc IsmCardTypeFamily = "root-standby-sc"

    // Line card
    IsmCardTypeFamily_lc IsmCardTypeFamily = "lc"

    // Non-Fabric SP
    IsmCardTypeFamily_sp IsmCardTypeFamily = "sp"

    // Fabric SP
    IsmCardTypeFamily_fabric_sp IsmCardTypeFamily = "fabric-sp"

    // SPA
    IsmCardTypeFamily_spa IsmCardTypeFamily = "spa"
)

// InstmgrBagUserMsgCategory represents Instmgr bag user msg category
type InstmgrBagUserMsgCategory string

const (
    // User error
    InstmgrBagUserMsgCategory_user_error InstmgrBagUserMsgCategory = "user-error"

    // Non-specific message
    InstmgrBagUserMsgCategory_non_specific InstmgrBagUserMsgCategory = "non-specific"

    // Warning message
    InstmgrBagUserMsgCategory_warning InstmgrBagUserMsgCategory = "warning"

    // Information message
    InstmgrBagUserMsgCategory_information InstmgrBagUserMsgCategory = "information"

    // User prompt
    InstmgrBagUserMsgCategory_user_prompt InstmgrBagUserMsgCategory = "user-prompt"

    // Log message
    InstmgrBagUserMsgCategory_log InstmgrBagUserMsgCategory = "log"

    // System error
    InstmgrBagUserMsgCategory_system_error InstmgrBagUserMsgCategory = "system-error"

    // User response
    InstmgrBagUserMsgCategory_user_response InstmgrBagUserMsgCategory = "user-response"
)

// InstallmgrIsmNodeConforming represents ISSU manage node inventory type
type InstallmgrIsmNodeConforming string

const (
    // Conforming Nodes
    InstallmgrIsmNodeConforming_conforming InstallmgrIsmNodeConforming = "conforming"

    // Non-conforming nodes
    InstallmgrIsmNodeConforming_none_conforming InstallmgrIsmNodeConforming = "none-conforming"

    // Node Upgrade failed
    InstallmgrIsmNodeConforming_upgrade_fail InstallmgrIsmNodeConforming = "upgrade-fail"

    // Non-conforming SPAs
    InstallmgrIsmNodeConforming_none_conforming_spa InstallmgrIsmNodeConforming = "none-conforming-spa"

    // SPA Upgrade failed
    InstallmgrIsmNodeConforming_spa_upgrade_fail InstallmgrIsmNodeConforming = "spa-upgrade-fail"
)

// InstmgrInstallPhase represents Current operation phase
type InstmgrInstallPhase string

const (
    // Unknown operational phase
    InstmgrInstallPhase_inst_phase_unknown InstmgrInstallPhase = "inst-phase-unknown"

    // Downloading
    InstmgrInstallPhase_inst_phase_download InstmgrInstallPhase = "inst-phase-download"

    // Performing software changes
    InstmgrInstallPhase_inst_phase_sw_change InstmgrInstallPhase = "inst-phase-sw-change"

    // Cleaning up after op
    InstmgrInstallPhase_inst_phase_cleaning_up InstmgrInstallPhase = "inst-phase-cleaning-up"
)

// InstmgrIssuAbortImpact represents Abort impact
type InstmgrIssuAbortImpact string

const (
    // Unknown abort impact
    InstmgrIssuAbortImpact_undefined InstmgrIssuAbortImpact = "undefined"

    // Abort is hitless
    InstmgrIssuAbortImpact_hitless InstmgrIssuAbortImpact = "hitless"

    // Abort will not affect traffic
    InstmgrIssuAbortImpact_traffic_outage InstmgrIssuAbortImpact = "traffic-outage"

    // Abort impact: n/a
    InstmgrIssuAbortImpact_not_applicable InstmgrIssuAbortImpact = "not-applicable"
)

// InstmgrIsmNodeState represents ISSU manager node state
type InstmgrIsmNodeState string

const (
    // No ISSU in progress
    InstmgrIsmNodeState_none InstmgrIsmNodeState = "none"

    // Node GSP ready
    InstmgrIsmNodeState_issu_node_gsp_ready InstmgrIsmNodeState = "issu-node-gsp-ready"

    // Load shut done
    InstmgrIsmNodeState_load_shut_done InstmgrIsmNodeState = "load-shut-done"

    // Standby management nodes upgrade done
    InstmgrIsmNodeState_standby_management_upgrade_done InstmgrIsmNodeState = "standby-management-upgrade-done"

    // Fabric nodes upgrade done
    InstmgrIsmNodeState_fabric_upgrade_done InstmgrIsmNodeState = "fabric-upgrade-done"

    // iMDR preparation ACK received
    InstmgrIsmNodeState_imdr_preparation_ack_received InstmgrIsmNodeState = "imdr-preparation-ack-received"

    // iMDR preparation ACK failed
    InstmgrIsmNodeState_imdr_preparation_failed InstmgrIsmNodeState = "imdr-preparation-failed"

    // iMDR start AVK received
    InstmgrIsmNodeState_imdr_start_ack_received InstmgrIsmNodeState = "imdr-start-ack-received"

    // iMDR start failed
    InstmgrIsmNodeState_imdr_start_failed InstmgrIsmNodeState = "imdr-start-failed"

    // iMDR complete ACK received
    InstmgrIsmNodeState_imdr_complete_ack_received InstmgrIsmNodeState = "imdr-complete-ack-received"

    // iMDR complete failed
    InstmgrIsmNodeState_imdr_complete_failed InstmgrIsmNodeState = "imdr-complete-failed"

    // Standby management nodes ready
    InstmgrIsmNodeState_standby_management_ready InstmgrIsmNodeState = "standby-management-ready"

    // FO acked
    InstmgrIsmNodeState_fo_acknowledged InstmgrIsmNodeState = "fo-acknowledged"

    // FO complete
    InstmgrIsmNodeState_fo_complete InstmgrIsmNodeState = "fo-complete"

    // Standby nodes ready after FO
    InstmgrIsmNodeState_standby_ready_after_fo InstmgrIsmNodeState = "standby-ready-after-fo"

    // Node is ready after iMDR
    InstmgrIsmNodeState_iam_ready_afteri_mdr InstmgrIsmNodeState = "iam-ready-afteri-mdr"

    // NSF ready
    InstmgrIsmNodeState_nsf_ready InstmgrIsmNodeState = "nsf-ready"

    // NSF begin ACK received
    InstmgrIsmNodeState_nsf_begin_ack_received InstmgrIsmNodeState = "nsf-begin-ack-received"

    // iMDR done
    InstmgrIsmNodeState_imdr_done InstmgrIsmNodeState = "imdr-done"

    // Unshut done
    InstmgrIsmNodeState_unshut_done InstmgrIsmNodeState = "unshut-done"

    // Run done
    InstmgrIsmNodeState_run_done InstmgrIsmNodeState = "run-done"

    // iMDR abort sent
    InstmgrIsmNodeState_imdr_abort_sent InstmgrIsmNodeState = "imdr-abort-sent"

    // iMDR abort ACK Received
    InstmgrIsmNodeState_imdr_abort_ack_received InstmgrIsmNodeState = "imdr-abort-ack-received"

    // iMDR abort failed
    InstmgrIsmNodeState_imdr_abort_failed InstmgrIsmNodeState = "imdr-abort-failed"

    // Standby management nodes downgrade done
    InstmgrIsmNodeState_standby_management_downgrade_done InstmgrIsmNodeState = "standby-management-downgrade-done"

    // Fabric nodes downgrade done
    InstmgrIsmNodeState_fabric_downgrade_done InstmgrIsmNodeState = "fabric-downgrade-done"

    // Node reloaded during ISSU
    InstmgrIsmNodeState_reload_during_issu InstmgrIsmNodeState = "reload-during-issu"

    // Node time out
    InstmgrIsmNodeState_timneout InstmgrIsmNodeState = "timneout"

    // Fabric upgrade failed
    InstmgrIsmNodeState_fabric_upgrade_failed InstmgrIsmNodeState = "fabric-upgrade-failed"

    // Unsupported hardware
    InstmgrIsmNodeState_unsupported_hw InstmgrIsmNodeState = "unsupported-hw"

    // Node unreachable
    InstmgrIsmNodeState_not_reachable InstmgrIsmNodeState = "not-reachable"

    // Max node state
    InstmgrIsmNodeState_max InstmgrIsmNodeState = "max"
)

// InstmgrPkg represents Package type
type InstmgrPkg string

const (
    // Undefined package
    InstmgrPkg_inst_pkg_type_undefined InstmgrPkg = "inst-pkg-type-undefined"

    // Root package
    InstmgrPkg_inst_pkg_type_root InstmgrPkg = "inst-pkg-type-root"

    // Standard package
    InstmgrPkg_inst_pkg_type_standard InstmgrPkg = "inst-pkg-type-standard"

    // Internal package
    InstmgrPkg_inst_pkg_type_internal InstmgrPkg = "inst-pkg-type-internal"
)

// InstmgrCardState represents Instmgr card state
type InstmgrCardState string

const (
    // instmgr card not present
    InstmgrCardState_instmgr_card_not_present InstmgrCardState = "instmgr-card-not-present"

    // instmgr card present
    InstmgrCardState_instmgr_card_present InstmgrCardState = "instmgr-card-present"

    // instmgr card reset
    InstmgrCardState_instmgr_card_reset InstmgrCardState = "instmgr-card-reset"

    // instmgr card booting
    InstmgrCardState_instmgr_card_booting InstmgrCardState = "instmgr-card-booting"

    // instmgr card mbi booting
    InstmgrCardState_instmgr_card_mbi_booting InstmgrCardState = "instmgr-card-mbi-booting"

    // instmgr card running mbi
    InstmgrCardState_instmgr_card_running_mbi InstmgrCardState = "instmgr-card-running-mbi"

    // instmgr card running ena
    InstmgrCardState_instmgr_card_running_ena InstmgrCardState = "instmgr-card-running-ena"

    // instmgr card bring down
    InstmgrCardState_instmgr_card_bring_down InstmgrCardState = "instmgr-card-bring-down"

    // instmgr card ena failure
    InstmgrCardState_instmgr_card_ena_failure InstmgrCardState = "instmgr-card-ena-failure"

    // instmgr card f diag run
    InstmgrCardState_instmgr_card_f_diag_run InstmgrCardState = "instmgr-card-f-diag-run"

    // instmgr card f diag failure
    InstmgrCardState_instmgr_card_f_diag_failure InstmgrCardState = "instmgr-card-f-diag-failure"

    // instmgr card powered
    InstmgrCardState_instmgr_card_powered InstmgrCardState = "instmgr-card-powered"

    // instmgr card unpowered
    InstmgrCardState_instmgr_card_unpowered InstmgrCardState = "instmgr-card-unpowered"

    // instmgr card mdr
    InstmgrCardState_instmgr_card_mdr InstmgrCardState = "instmgr-card-mdr"

    // instmgr card mdr running mbi
    InstmgrCardState_instmgr_card_mdr_running_mbi InstmgrCardState = "instmgr-card-mdr-running-mbi"

    // instmgr card main t mode
    InstmgrCardState_instmgr_card_main_t_mode InstmgrCardState = "instmgr-card-main-t-mode"

    // instmgr card admin down
    InstmgrCardState_instmgr_card_admin_down InstmgrCardState = "instmgr-card-admin-down"

    // instmgr card no mon
    InstmgrCardState_instmgr_card_no_mon InstmgrCardState = "instmgr-card-no-mon"

    // instmgr card unknown
    InstmgrCardState_instmgr_card_unknown InstmgrCardState = "instmgr-card-unknown"

    // instmgr card failed
    InstmgrCardState_instmgr_card_failed InstmgrCardState = "instmgr-card-failed"

    // instmgr card ok
    InstmgrCardState_instmgr_card_ok InstmgrCardState = "instmgr-card-ok"

    // instmgr card missing
    InstmgrCardState_instmgr_card_missing InstmgrCardState = "instmgr-card-missing"

    // instmgr card field diag downloading
    InstmgrCardState_instmgr_card_field_diag_downloading InstmgrCardState = "instmgr-card-field-diag-downloading"

    // instmgr card field diag unmonitor
    InstmgrCardState_instmgr_card_field_diag_unmonitor InstmgrCardState = "instmgr-card-field-diag-unmonitor"

    // instmgr card fabric field diag unmonitor
    InstmgrCardState_instmgr_card_fabric_field_diag_unmonitor InstmgrCardState = "instmgr-card-fabric-field-diag-unmonitor"

    // instmgr card field diag rp launching
    InstmgrCardState_instmgr_card_field_diag_rp_launching InstmgrCardState = "instmgr-card-field-diag-rp-launching"

    // instmgr card field diag running
    InstmgrCardState_instmgr_card_field_diag_running InstmgrCardState = "instmgr-card-field-diag-running"

    // instmgr card field diag pass
    InstmgrCardState_instmgr_card_field_diag_pass InstmgrCardState = "instmgr-card-field-diag-pass"

    // instmgr card field diag fail
    InstmgrCardState_instmgr_card_field_diag_fail InstmgrCardState = "instmgr-card-field-diag-fail"

    // instmgr card field diag timeout
    InstmgrCardState_instmgr_card_field_diag_timeout InstmgrCardState = "instmgr-card-field-diag-timeout"

    // instmgr card disabled
    InstmgrCardState_instmgr_card_disabled InstmgrCardState = "instmgr-card-disabled"

    // instmgr card spa booting
    InstmgrCardState_instmgr_card_spa_booting InstmgrCardState = "instmgr-card-spa-booting"

    // instmgr card not allowed online
    InstmgrCardState_instmgr_card_not_allowed_online InstmgrCardState = "instmgr-card-not-allowed-online"

    // instmgr card stopped
    InstmgrCardState_instmgr_card_stopped InstmgrCardState = "instmgr-card-stopped"

    // instmgr card incompatible fw ver
    InstmgrCardState_instmgr_card_incompatible_fw_ver InstmgrCardState = "instmgr-card-incompatible-fw-ver"

    // instmgr card fpd hold
    InstmgrCardState_instmgr_card_fpd_hold InstmgrCardState = "instmgr-card-fpd-hold"

    // instmgr card updating fpd
    InstmgrCardState_instmgr_card_updating_fpd InstmgrCardState = "instmgr-card-updating-fpd"

    // instmgr card num states
    InstmgrCardState_instmgr_card_num_states InstmgrCardState = "instmgr-card-num-states"
)

// InstmgrNodeRole represents Node role
type InstmgrNodeRole string

const (
    // Redundency unknown
    InstmgrNodeRole_redundency_unknown InstmgrNodeRole = "redundency-unknown"

    // Redundency active
    InstmgrNodeRole_redundency_active InstmgrNodeRole = "redundency-active"

    // Redundency standby
    InstmgrNodeRole_redundency_standby InstmgrNodeRole = "redundency-standby"

    // Redundency unusable
    InstmgrNodeRole_redundency_unusable InstmgrNodeRole = "redundency-unusable"
)

// InstmgrRequest represents Instmgr request
type InstmgrRequest string

const (
    // install add
    InstmgrRequest_add InstmgrRequest = "add"

    // install accept
    InstmgrRequest_accept InstmgrRequest = "accept"

    // install clean
    InstmgrRequest_clean InstmgrRequest = "clean"

    // install activate
    InstmgrRequest_activate InstmgrRequest = "activate"

    // install deactivate
    InstmgrRequest_deactivate InstmgrRequest = "deactivate"

    // install commit
    InstmgrRequest_commit InstmgrRequest = "commit"

    // install verify
    InstmgrRequest_verify InstmgrRequest = "verify"

    // install rollback
    InstmgrRequest_rollback InstmgrRequest = "rollback"

    // install clear rollback oldest
    InstmgrRequest_clear_rollback InstmgrRequest = "clear-rollback"

    // install clear historylog
    InstmgrRequest_clear_log InstmgrRequest = "clear-log"

    // (Deprecated) install healthcheck
    InstmgrRequest_health_check InstmgrRequest = "health-check"

    // install run/accept/complete
    InstmgrRequest_operation InstmgrRequest = "operation"

    // install auto-abort-timer stop
    InstmgrRequest_stop_timer InstmgrRequest = "stop-timer"

    // install label
    InstmgrRequest_label InstmgrRequest = "label"

    // clear install label
    InstmgrRequest_clear_label InstmgrRequest = "clear-label"

    // install extend
    InstmgrRequest_extend InstmgrRequest = "extend"
)

// InstmgrIsmFsmState represents Install manager FSM state
type InstmgrIsmFsmState string

const (
    // No ISSU in progress
    InstmgrIsmFsmState_idle InstmgrIsmFsmState = "idle"

    // LOAD init
    InstmgrIsmFsmState_init_done InstmgrIsmFsmState = "init-done"

    // LOAD preparation
    InstmgrIsmFsmState_load_shut InstmgrIsmFsmState = "load-shut"

    // LOAD wait
    InstmgrIsmFsmState_load_wait InstmgrIsmFsmState = "load-wait"

    // LOAD root SC FO
    InstmgrIsmFsmState_load_stp_root_before InstmgrIsmFsmState = "load-stp-root-before"

    // LOAD standby ROOT SC Upgrade
    InstmgrIsmFsmState_load_standby_root_sc_upgrade InstmgrIsmFsmState = "load-standby-root-sc-upgrade"

    // LOAD standby management upgrade
    InstmgrIsmFsmState_load_standby_management_upgrade InstmgrIsmFsmState = "load-standby-management-upgrade"

    // LOAD NDSC FO
    InstmgrIsmFsmState_load_stp_root_after InstmgrIsmFsmState = "load-stp-root-after"

    // LOAD fabric upgrade
    InstmgrIsmFsmState_load_fabric_upgrade InstmgrIsmFsmState = "load-fabric-upgrade"

    // LOAD ISSU ready
    InstmgrIsmFsmState_load_management_issu_ready InstmgrIsmFsmState = "load-management-issu-ready"

    // LOAD done
    InstmgrIsmFsmState_load_done InstmgrIsmFsmState = "load-done"

    // RUN preparation
    InstmgrIsmFsmState_run_prep InstmgrIsmFsmState = "run-prep"

    // RUN wait
    InstmgrIsmFsmState_run_wait InstmgrIsmFsmState = "run-wait"

    // RUN iMDR preparation
    InstmgrIsmFsmState_runi_mdr_prep InstmgrIsmFsmState = "runi-mdr-prep"

    // RUN iMDR start
    InstmgrIsmFsmState_runi_mdr_start InstmgrIsmFsmState = "runi-mdr-start"

    // RUN iMDR complete
    InstmgrIsmFsmState_runi_mdr_complete InstmgrIsmFsmState = "runi-mdr-complete"

    // RUN make standby ready
    InstmgrIsmFsmState_run_make_standby_ready InstmgrIsmFsmState = "run-make-standby-ready"

    // RUN root SC FO
    InstmgrIsmFsmState_run_root_scfo InstmgrIsmFsmState = "run-root-scfo"

    // RUN NDSC FO
    InstmgrIsmFsmState_run_ndscfo InstmgrIsmFsmState = "run-ndscfo"

    // RUN transient1
    InstmgrIsmFsmState_run_transient1 InstmgrIsmFsmState = "run-transient1"

    // RUN DSC FO
    InstmgrIsmFsmState_run_dscfo InstmgrIsmFsmState = "run-dscfo"

    // RUN FO compelte
    InstmgrIsmFsmState_run_fo_complete InstmgrIsmFsmState = "run-fo-complete"

    // Run STP Root Return
    InstmgrIsmFsmState_run_stp_root_return InstmgrIsmFsmState = "run-stp-root-return"

    // RUN iMDR continue
    InstmgrIsmFsmState_runi_mdr_continue InstmgrIsmFsmState = "runi-mdr-continue"

    // RUN I am ready after iMDR
    InstmgrIsmFsmState_run_am_i_ready_afteri_mdr InstmgrIsmFsmState = "run-am-i-ready-afteri-mdr"

    // RUN NSF ready
    InstmgrIsmFsmState_run_nsf_ready InstmgrIsmFsmState = "run-nsf-ready"

    // RUN iMDR begin
    InstmgrIsmFsmState_run_nsf_begin InstmgrIsmFsmState = "run-nsf-begin"

    // RUN iMDR done
    InstmgrIsmFsmState_runi_mdr_done InstmgrIsmFsmState = "runi-mdr-done"

    // RUN mgmt issu ready
    InstmgrIsmFsmState_run_management_issu_ready InstmgrIsmFsmState = "run-management-issu-ready"

    // RUN unshut
    InstmgrIsmFsmState_run_un_shut InstmgrIsmFsmState = "run-un-shut"

    // RUN done
    InstmgrIsmFsmState_run_is_done InstmgrIsmFsmState = "run-is-done"

    // Max ISSU state
    InstmgrIsmFsmState_state_max InstmgrIsmFsmState = "state-max"
)

// InstmgrBagIiDirection represents The Incremental Install direction
type InstmgrBagIiDirection string

const (
    // Not incremental install operation
    InstmgrBagIiDirection_not_incremental InstmgrBagIiDirection = "not-incremental"

    // Installing
    InstmgrBagIiDirection_installing InstmgrBagIiDirection = "installing"

    // Unwinding
    InstmgrBagIiDirection_unwinding InstmgrBagIiDirection = "unwinding"
)

// InstmgrPiCard represents PI card types
type InstmgrPiCard string

const (
    // Card type RP
    InstmgrPiCard_type_rp InstmgrPiCard = "type-rp"

    // Card Type DRP
    InstmgrPiCard_type_drp InstmgrPiCard = "type-drp"

    // Card type  LC
    InstmgrPiCard_type_lc InstmgrPiCard = "type-lc"

    // Card type SC
    InstmgrPiCard_type_sc InstmgrPiCard = "type-sc"

    // Card type SP
    InstmgrPiCard_type_sp InstmgrPiCard = "type-sp"

    // Card type other
    InstmgrPiCard_type_other InstmgrPiCard = "type-other"
)

// InstmgrBagIiState represents The Incremental Install state of an install
type InstmgrBagIiState string

const (
    // Node to be upraded
    InstmgrBagIiState_idle InstmgrBagIiState = "idle"

    // Node is being upraded
    InstmgrBagIiState_in_progress InstmgrBagIiState = "in-progress"

    // Node upgraded successfully
    InstmgrBagIiState_completed InstmgrBagIiState = "completed"

    // Node reverted to the old S/W
    InstmgrBagIiState_aborted InstmgrBagIiState = "aborted"

    // Node rebooted and held in MBI
    InstmgrBagIiState_rebooted InstmgrBagIiState = "rebooted"
)

// InstmgrBagLogEntryUserMsgCategory represents Category type
type InstmgrBagLogEntryUserMsgCategory string

const (
    // User error
    InstmgrBagLogEntryUserMsgCategory_user_error InstmgrBagLogEntryUserMsgCategory = "user-error"

    // Non-specific message
    InstmgrBagLogEntryUserMsgCategory_non_specific InstmgrBagLogEntryUserMsgCategory = "non-specific"

    // Warning message
    InstmgrBagLogEntryUserMsgCategory_warning InstmgrBagLogEntryUserMsgCategory = "warning"

    // Information message
    InstmgrBagLogEntryUserMsgCategory_information InstmgrBagLogEntryUserMsgCategory = "information"

    // User prompt
    InstmgrBagLogEntryUserMsgCategory_user_prompt InstmgrBagLogEntryUserMsgCategory = "user-prompt"

    // Log message
    InstmgrBagLogEntryUserMsgCategory_log InstmgrBagLogEntryUserMsgCategory = "log"

    // System error
    InstmgrBagLogEntryUserMsgCategory_system_error InstmgrBagLogEntryUserMsgCategory = "system-error"

    // User response
    InstmgrBagLogEntryUserMsgCategory_user_response InstmgrBagLogEntryUserMsgCategory = "user-response"
)

// InstmgrBagAbortState represents The abortable state of an install command
type InstmgrBagAbortState string

const (
    // Operation can be aborted
    InstmgrBagAbortState_abortable InstmgrBagAbortState = "abortable"

    // Operation can no longer be aborted
    InstmgrBagAbortState_no_longer_abortable InstmgrBagAbortState = "no-longer-abortable"

    // Operation cannot be aborted
    InstmgrBagAbortState_never_abortable InstmgrBagAbortState = "never-abortable"

    // Operation has been aborted
    InstmgrBagAbortState_already_aborted InstmgrBagAbortState = "already-aborted"
)

// Install
// Information of software packages and install
// operations
type Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install operation log history size. The type is interface{} with range:
    // 0..4294967295.
    LogSize interface{}

    // Configuration register (confreg) information.
    ConfigurationRegisters Install_ConfigurationRegisters

    // Install operation request status.
    RequestStatuses Install_RequestStatuses

    // Boot variable information.
    BootVariables Install_BootVariables

    // Software package,component and alias information.
    Software Install_Software

    // Information of install operations.
    SoftwareInventory Install_SoftwareInventory

    // Information of install ISSU operations.
    Issu Install_Issu

    // System Boot Image.
    BootImage Install_BootImage

    // Install operation log.
    Logs Install_Logs
}

func (install *Install) GetEntityData() *types.CommonEntityData {
    install.EntityData.YFilter = install.YFilter
    install.EntityData.YangName = "install"
    install.EntityData.BundleName = "cisco_ios_xr"
    install.EntityData.ParentYangName = "Cisco-IOS-XR-installmgr-admin-oper"
    install.EntityData.SegmentPath = "Cisco-IOS-XR-installmgr-admin-oper:install"
    install.EntityData.AbsolutePath = install.EntityData.SegmentPath
    install.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    install.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    install.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    install.EntityData.Children = types.NewOrderedMap()
    install.EntityData.Children.Append("configuration-registers", types.YChild{"ConfigurationRegisters", &install.ConfigurationRegisters})
    install.EntityData.Children.Append("request-statuses", types.YChild{"RequestStatuses", &install.RequestStatuses})
    install.EntityData.Children.Append("boot-variables", types.YChild{"BootVariables", &install.BootVariables})
    install.EntityData.Children.Append("software", types.YChild{"Software", &install.Software})
    install.EntityData.Children.Append("software-inventory", types.YChild{"SoftwareInventory", &install.SoftwareInventory})
    install.EntityData.Children.Append("issu", types.YChild{"Issu", &install.Issu})
    install.EntityData.Children.Append("boot-image", types.YChild{"BootImage", &install.BootImage})
    install.EntityData.Children.Append("logs", types.YChild{"Logs", &install.Logs})
    install.EntityData.Leafs = types.NewOrderedMap()
    install.EntityData.Leafs.Append("log-size", types.YLeaf{"LogSize", install.LogSize})

    install.EntityData.YListKeys = []string {}

    return &(install.EntityData)
}

// Install_ConfigurationRegisters
// Configuration register (confreg) information
type Install_ConfigurationRegisters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration register for specific node. The type is slice of
    // Install_ConfigurationRegisters_ConfigurationRegister.
    ConfigurationRegister []*Install_ConfigurationRegisters_ConfigurationRegister
}

func (configurationRegisters *Install_ConfigurationRegisters) GetEntityData() *types.CommonEntityData {
    configurationRegisters.EntityData.YFilter = configurationRegisters.YFilter
    configurationRegisters.EntityData.YangName = "configuration-registers"
    configurationRegisters.EntityData.BundleName = "cisco_ios_xr"
    configurationRegisters.EntityData.ParentYangName = "install"
    configurationRegisters.EntityData.SegmentPath = "configuration-registers"
    configurationRegisters.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/" + configurationRegisters.EntityData.SegmentPath
    configurationRegisters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationRegisters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationRegisters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationRegisters.EntityData.Children = types.NewOrderedMap()
    configurationRegisters.EntityData.Children.Append("configuration-register", types.YChild{"ConfigurationRegister", nil})
    for i := range configurationRegisters.ConfigurationRegister {
        configurationRegisters.EntityData.Children.Append(types.GetSegmentPath(configurationRegisters.ConfigurationRegister[i]), types.YChild{"ConfigurationRegister", configurationRegisters.ConfigurationRegister[i]})
    }
    configurationRegisters.EntityData.Leafs = types.NewOrderedMap()

    configurationRegisters.EntityData.YListKeys = []string {}

    return &(configurationRegisters.EntityData)
}

// Install_ConfigurationRegisters_ConfigurationRegister
// Configuration register for specific node
type Install_ConfigurationRegisters_ConfigurationRegister struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Configuration register value. The type is string with pattern:
    // [0-9a-fA-F]{1,8}. This attribute is mandatory.
    ConfigRegister interface{}
}

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetEntityData() *types.CommonEntityData {
    configurationRegister.EntityData.YFilter = configurationRegister.YFilter
    configurationRegister.EntityData.YangName = "configuration-register"
    configurationRegister.EntityData.BundleName = "cisco_ios_xr"
    configurationRegister.EntityData.ParentYangName = "configuration-registers"
    configurationRegister.EntityData.SegmentPath = "configuration-register" + types.AddKeyToken(configurationRegister.NodeName, "node-name")
    configurationRegister.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/configuration-registers/" + configurationRegister.EntityData.SegmentPath
    configurationRegister.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationRegister.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationRegister.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationRegister.EntityData.Children = types.NewOrderedMap()
    configurationRegister.EntityData.Leafs = types.NewOrderedMap()
    configurationRegister.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", configurationRegister.NodeName})
    configurationRegister.EntityData.Leafs.Append("config-register", types.YLeaf{"ConfigRegister", configurationRegister.ConfigRegister})

    configurationRegister.EntityData.YListKeys = []string {"NodeName"}

    return &(configurationRegister.EntityData)
}

// Install_RequestStatuses
// Install operation request status
type Install_RequestStatuses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request status Information. The type is slice of
    // Install_RequestStatuses_RequestStatus.
    RequestStatus []*Install_RequestStatuses_RequestStatus
}

func (requestStatuses *Install_RequestStatuses) GetEntityData() *types.CommonEntityData {
    requestStatuses.EntityData.YFilter = requestStatuses.YFilter
    requestStatuses.EntityData.YangName = "request-statuses"
    requestStatuses.EntityData.BundleName = "cisco_ios_xr"
    requestStatuses.EntityData.ParentYangName = "install"
    requestStatuses.EntityData.SegmentPath = "request-statuses"
    requestStatuses.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/" + requestStatuses.EntityData.SegmentPath
    requestStatuses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requestStatuses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requestStatuses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requestStatuses.EntityData.Children = types.NewOrderedMap()
    requestStatuses.EntityData.Children.Append("request-status", types.YChild{"RequestStatus", nil})
    for i := range requestStatuses.RequestStatus {
        requestStatuses.EntityData.Children.Append(types.GetSegmentPath(requestStatuses.RequestStatus[i]), types.YChild{"RequestStatus", requestStatuses.RequestStatus[i]})
    }
    requestStatuses.EntityData.Leafs = types.NewOrderedMap()

    requestStatuses.EntityData.YListKeys = []string {}

    return &(requestStatuses.EntityData)
}

// Install_RequestStatuses_RequestStatus
// Request status Information
type Install_RequestStatuses_RequestStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Install operation request ID. The type is
    // interface{} with range: 0..4294967295.
    RequestId interface{}

    // Percentage completed. The type is interface{} with range: 0..255. Units are
    // percentage.
    Percentage interface{}

    // Abort state. The type is InstmgrBagAbortState.
    AbortState interface{}

    // Downloaded bytes. The type is interface{} with range: 0..4294967295. Units
    // are byte.
    DownloadedBytes interface{}

    // Whether the operation is blocked waiting for a response from the user. The
    // type is bool.
    UnansweredQuery interface{}

    // Phase of the operation. The type is InstmgrInstallPhase.
    OperationPhase interface{}

    // Requested install operation.
    RequestInformation Install_RequestStatuses_RequestStatus_RequestInformation

    // How the abort will occur.
    AbortStatus Install_RequestStatuses_RequestStatus_AbortStatus

    // Incremental Install information.
    IncrementalInstallInformation Install_RequestStatuses_RequestStatus_IncrementalInstallInformation

    // Messages related to ISSU operations. The type is slice of
    // Install_RequestStatuses_RequestStatus_IssuMessage.
    IssuMessage []*Install_RequestStatuses_RequestStatus_IssuMessage

    // Messages output to the user. The type is slice of
    // Install_RequestStatuses_RequestStatus_Message.
    Message []*Install_RequestStatuses_RequestStatus_Message
}

func (requestStatus *Install_RequestStatuses_RequestStatus) GetEntityData() *types.CommonEntityData {
    requestStatus.EntityData.YFilter = requestStatus.YFilter
    requestStatus.EntityData.YangName = "request-status"
    requestStatus.EntityData.BundleName = "cisco_ios_xr"
    requestStatus.EntityData.ParentYangName = "request-statuses"
    requestStatus.EntityData.SegmentPath = "request-status" + types.AddKeyToken(requestStatus.RequestId, "request-id")
    requestStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/" + requestStatus.EntityData.SegmentPath
    requestStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requestStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requestStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requestStatus.EntityData.Children = types.NewOrderedMap()
    requestStatus.EntityData.Children.Append("request-information", types.YChild{"RequestInformation", &requestStatus.RequestInformation})
    requestStatus.EntityData.Children.Append("abort-status", types.YChild{"AbortStatus", &requestStatus.AbortStatus})
    requestStatus.EntityData.Children.Append("incremental-install-information", types.YChild{"IncrementalInstallInformation", &requestStatus.IncrementalInstallInformation})
    requestStatus.EntityData.Children.Append("issu-message", types.YChild{"IssuMessage", nil})
    for i := range requestStatus.IssuMessage {
        types.SetYListKey(requestStatus.IssuMessage[i], i)
        requestStatus.EntityData.Children.Append(types.GetSegmentPath(requestStatus.IssuMessage[i]), types.YChild{"IssuMessage", requestStatus.IssuMessage[i]})
    }
    requestStatus.EntityData.Children.Append("message", types.YChild{"Message", nil})
    for i := range requestStatus.Message {
        types.SetYListKey(requestStatus.Message[i], i)
        requestStatus.EntityData.Children.Append(types.GetSegmentPath(requestStatus.Message[i]), types.YChild{"Message", requestStatus.Message[i]})
    }
    requestStatus.EntityData.Leafs = types.NewOrderedMap()
    requestStatus.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", requestStatus.RequestId})
    requestStatus.EntityData.Leafs.Append("percentage", types.YLeaf{"Percentage", requestStatus.Percentage})
    requestStatus.EntityData.Leafs.Append("abort-state", types.YLeaf{"AbortState", requestStatus.AbortState})
    requestStatus.EntityData.Leafs.Append("downloaded-bytes", types.YLeaf{"DownloadedBytes", requestStatus.DownloadedBytes})
    requestStatus.EntityData.Leafs.Append("unanswered-query", types.YLeaf{"UnansweredQuery", requestStatus.UnansweredQuery})
    requestStatus.EntityData.Leafs.Append("operation-phase", types.YLeaf{"OperationPhase", requestStatus.OperationPhase})

    requestStatus.EntityData.YListKeys = []string {"RequestId"}

    return &(requestStatus.EntityData)
}

// Install_RequestStatuses_RequestStatus_RequestInformation
// Requested install operation
type Install_RequestStatuses_RequestStatus_RequestInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install id of operation. The type is interface{} with range: 0..4294967295.
    RequestId interface{}

    // User ID that started the reqeust. The type is string.
    UserId interface{}

    // Request trigger type. The type is InstmgrBagRequestTrigger.
    TriggerType interface{}

    // Options affecting processing of install requests. The type is interface{}
    // with range: -2147483648..2147483647.
    RequestOption interface{}

    // Requested operation type. The type is InstmgrRequest.
    OperationType interface{}

    // Detail operation information. The type is string.
    OperationDetail interface{}
}

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetEntityData() *types.CommonEntityData {
    requestInformation.EntityData.YFilter = requestInformation.YFilter
    requestInformation.EntityData.YangName = "request-information"
    requestInformation.EntityData.BundleName = "cisco_ios_xr"
    requestInformation.EntityData.ParentYangName = "request-status"
    requestInformation.EntityData.SegmentPath = "request-information"
    requestInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/request-status/" + requestInformation.EntityData.SegmentPath
    requestInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requestInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requestInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requestInformation.EntityData.Children = types.NewOrderedMap()
    requestInformation.EntityData.Leafs = types.NewOrderedMap()
    requestInformation.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", requestInformation.RequestId})
    requestInformation.EntityData.Leafs.Append("user-id", types.YLeaf{"UserId", requestInformation.UserId})
    requestInformation.EntityData.Leafs.Append("trigger-type", types.YLeaf{"TriggerType", requestInformation.TriggerType})
    requestInformation.EntityData.Leafs.Append("request-option", types.YLeaf{"RequestOption", requestInformation.RequestOption})
    requestInformation.EntityData.Leafs.Append("operation-type", types.YLeaf{"OperationType", requestInformation.OperationType})
    requestInformation.EntityData.Leafs.Append("operation-detail", types.YLeaf{"OperationDetail", requestInformation.OperationDetail})

    requestInformation.EntityData.YListKeys = []string {}

    return &(requestInformation.EntityData)
}

// Install_RequestStatuses_RequestStatus_AbortStatus
// How the abort will occur
type Install_RequestStatuses_RequestStatus_AbortStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Method of abort. The type is InstmgrIssuAbortMethod.
    AbortMethod interface{}

    // Impact of abort. The type is InstmgrIssuAbortImpact.
    AbortImpact interface{}
}

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetEntityData() *types.CommonEntityData {
    abortStatus.EntityData.YFilter = abortStatus.YFilter
    abortStatus.EntityData.YangName = "abort-status"
    abortStatus.EntityData.BundleName = "cisco_ios_xr"
    abortStatus.EntityData.ParentYangName = "request-status"
    abortStatus.EntityData.SegmentPath = "abort-status"
    abortStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/request-status/" + abortStatus.EntityData.SegmentPath
    abortStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    abortStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    abortStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    abortStatus.EntityData.Children = types.NewOrderedMap()
    abortStatus.EntityData.Leafs = types.NewOrderedMap()
    abortStatus.EntityData.Leafs.Append("abort-method", types.YLeaf{"AbortMethod", abortStatus.AbortMethod})
    abortStatus.EntityData.Leafs.Append("abort-impact", types.YLeaf{"AbortImpact", abortStatus.AbortImpact})

    abortStatus.EntityData.YListKeys = []string {}

    return &(abortStatus.EntityData)
}

// Install_RequestStatuses_RequestStatus_IncrementalInstallInformation
// Incremental Install information
type Install_RequestStatuses_RequestStatus_IncrementalInstallInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install direction. The type is InstmgrBagIiDirection.
    Direction interface{}

    // First error during incremental install. The type is string.
    IiError interface{}

    // Participating nodes. The type is slice of
    // Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes.
    Nodes []*Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes
}

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetEntityData() *types.CommonEntityData {
    incrementalInstallInformation.EntityData.YFilter = incrementalInstallInformation.YFilter
    incrementalInstallInformation.EntityData.YangName = "incremental-install-information"
    incrementalInstallInformation.EntityData.BundleName = "cisco_ios_xr"
    incrementalInstallInformation.EntityData.ParentYangName = "request-status"
    incrementalInstallInformation.EntityData.SegmentPath = "incremental-install-information"
    incrementalInstallInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/request-status/" + incrementalInstallInformation.EntityData.SegmentPath
    incrementalInstallInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incrementalInstallInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incrementalInstallInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incrementalInstallInformation.EntityData.Children = types.NewOrderedMap()
    incrementalInstallInformation.EntityData.Children.Append("nodes", types.YChild{"Nodes", nil})
    for i := range incrementalInstallInformation.Nodes {
        types.SetYListKey(incrementalInstallInformation.Nodes[i], i)
        incrementalInstallInformation.EntityData.Children.Append(types.GetSegmentPath(incrementalInstallInformation.Nodes[i]), types.YChild{"Nodes", incrementalInstallInformation.Nodes[i]})
    }
    incrementalInstallInformation.EntityData.Leafs = types.NewOrderedMap()
    incrementalInstallInformation.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", incrementalInstallInformation.Direction})
    incrementalInstallInformation.EntityData.Leafs.Append("ii-error", types.YLeaf{"IiError", incrementalInstallInformation.IiError})

    incrementalInstallInformation.EntityData.YListKeys = []string {}

    return &(incrementalInstallInformation.EntityData)
}

// Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes
// Participating nodes
type Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Node identifier. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // State. The type is InstmgrBagIiState.
    State interface{}
}

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "incremental-install-information"
    nodes.EntityData.SegmentPath = "nodes" + types.AddNoKeyToken(nodes)
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/request-status/incremental-install-information/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Leafs = types.NewOrderedMap()
    nodes.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", nodes.NodeName})
    nodes.EntityData.Leafs.Append("state", types.YLeaf{"State", nodes.State})

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Install_RequestStatuses_RequestStatus_IssuMessage
// Messages related to ISSU operations
type Install_RequestStatuses_RequestStatus_IssuMessage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Category of the message. The type is InstmgrBagUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_RequestStatuses_RequestStatus_IssuMessage_Scope
}

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetEntityData() *types.CommonEntityData {
    issuMessage.EntityData.YFilter = issuMessage.YFilter
    issuMessage.EntityData.YangName = "issu-message"
    issuMessage.EntityData.BundleName = "cisco_ios_xr"
    issuMessage.EntityData.ParentYangName = "request-status"
    issuMessage.EntityData.SegmentPath = "issu-message" + types.AddNoKeyToken(issuMessage)
    issuMessage.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/request-status/" + issuMessage.EntityData.SegmentPath
    issuMessage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issuMessage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issuMessage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issuMessage.EntityData.Children = types.NewOrderedMap()
    issuMessage.EntityData.Children.Append("scope", types.YChild{"Scope", &issuMessage.Scope})
    issuMessage.EntityData.Leafs = types.NewOrderedMap()
    issuMessage.EntityData.Leafs.Append("category", types.YLeaf{"Category", issuMessage.Category})
    issuMessage.EntityData.Leafs.Append("message", types.YLeaf{"Message", issuMessage.Message})

    issuMessage.EntityData.YListKeys = []string {}

    return &(issuMessage.EntityData)
}

// Install_RequestStatuses_RequestStatus_IssuMessage_Scope
// Scope of the message
type Install_RequestStatuses_RequestStatus_IssuMessage_Scope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which LRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetEntityData() *types.CommonEntityData {
    scope.EntityData.YFilter = scope.YFilter
    scope.EntityData.YangName = "scope"
    scope.EntityData.BundleName = "cisco_ios_xr"
    scope.EntityData.ParentYangName = "issu-message"
    scope.EntityData.SegmentPath = "scope"
    scope.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/request-status/issu-message/" + scope.EntityData.SegmentPath
    scope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scope.EntityData.Children = types.NewOrderedMap()
    scope.EntityData.Leafs = types.NewOrderedMap()
    scope.EntityData.Leafs.Append("admin-read", types.YLeaf{"AdminRead", scope.AdminRead})
    scope.EntityData.Leafs.Append("affected-sd-rs", types.YLeaf{"AffectedSdRs", scope.AffectedSdRs})

    scope.EntityData.YListKeys = []string {}

    return &(scope.EntityData)
}

// Install_RequestStatuses_RequestStatus_Message
// Messages output to the user
type Install_RequestStatuses_RequestStatus_Message struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Category of the message. The type is InstmgrBagUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_RequestStatuses_RequestStatus_Message_Scope
}

func (message *Install_RequestStatuses_RequestStatus_Message) GetEntityData() *types.CommonEntityData {
    message.EntityData.YFilter = message.YFilter
    message.EntityData.YangName = "message"
    message.EntityData.BundleName = "cisco_ios_xr"
    message.EntityData.ParentYangName = "request-status"
    message.EntityData.SegmentPath = "message" + types.AddNoKeyToken(message)
    message.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/request-status/" + message.EntityData.SegmentPath
    message.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    message.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    message.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    message.EntityData.Children = types.NewOrderedMap()
    message.EntityData.Children.Append("scope", types.YChild{"Scope", &message.Scope})
    message.EntityData.Leafs = types.NewOrderedMap()
    message.EntityData.Leafs.Append("category", types.YLeaf{"Category", message.Category})
    message.EntityData.Leafs.Append("message", types.YLeaf{"Message", message.Message})

    message.EntityData.YListKeys = []string {}

    return &(message.EntityData)
}

// Install_RequestStatuses_RequestStatus_Message_Scope
// Scope of the message
type Install_RequestStatuses_RequestStatus_Message_Scope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which LRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetEntityData() *types.CommonEntityData {
    scope.EntityData.YFilter = scope.YFilter
    scope.EntityData.YangName = "scope"
    scope.EntityData.BundleName = "cisco_ios_xr"
    scope.EntityData.ParentYangName = "message"
    scope.EntityData.SegmentPath = "scope"
    scope.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/request-statuses/request-status/message/" + scope.EntityData.SegmentPath
    scope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scope.EntityData.Children = types.NewOrderedMap()
    scope.EntityData.Leafs = types.NewOrderedMap()
    scope.EntityData.Leafs.Append("admin-read", types.YLeaf{"AdminRead", scope.AdminRead})
    scope.EntityData.Leafs.Append("affected-sd-rs", types.YLeaf{"AffectedSdRs", scope.AffectedSdRs})

    scope.EntityData.YListKeys = []string {}

    return &(scope.EntityData)
}

// Install_BootVariables
// Boot variable information
type Install_BootVariables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Boot variable for specific node. The type is slice of
    // Install_BootVariables_BootVariable.
    BootVariable []*Install_BootVariables_BootVariable
}

func (bootVariables *Install_BootVariables) GetEntityData() *types.CommonEntityData {
    bootVariables.EntityData.YFilter = bootVariables.YFilter
    bootVariables.EntityData.YangName = "boot-variables"
    bootVariables.EntityData.BundleName = "cisco_ios_xr"
    bootVariables.EntityData.ParentYangName = "install"
    bootVariables.EntityData.SegmentPath = "boot-variables"
    bootVariables.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/" + bootVariables.EntityData.SegmentPath
    bootVariables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootVariables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootVariables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootVariables.EntityData.Children = types.NewOrderedMap()
    bootVariables.EntityData.Children.Append("boot-variable", types.YChild{"BootVariable", nil})
    for i := range bootVariables.BootVariable {
        bootVariables.EntityData.Children.Append(types.GetSegmentPath(bootVariables.BootVariable[i]), types.YChild{"BootVariable", bootVariables.BootVariable[i]})
    }
    bootVariables.EntityData.Leafs = types.NewOrderedMap()

    bootVariables.EntityData.YListKeys = []string {}

    return &(bootVariables.EntityData)
}

// Install_BootVariables_BootVariable
// Boot variable for specific node
type Install_BootVariables_BootVariable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Boot variable value. The type is string. This attribute is mandatory.
    BootVariable interface{}
}

func (bootVariable *Install_BootVariables_BootVariable) GetEntityData() *types.CommonEntityData {
    bootVariable.EntityData.YFilter = bootVariable.YFilter
    bootVariable.EntityData.YangName = "boot-variable"
    bootVariable.EntityData.BundleName = "cisco_ios_xr"
    bootVariable.EntityData.ParentYangName = "boot-variables"
    bootVariable.EntityData.SegmentPath = "boot-variable" + types.AddKeyToken(bootVariable.NodeName, "node-name")
    bootVariable.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/boot-variables/" + bootVariable.EntityData.SegmentPath
    bootVariable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootVariable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootVariable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootVariable.EntityData.Children = types.NewOrderedMap()
    bootVariable.EntityData.Leafs = types.NewOrderedMap()
    bootVariable.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", bootVariable.NodeName})
    bootVariable.EntityData.Leafs.Append("boot-variable", types.YLeaf{"BootVariable", bootVariable.BootVariable})

    bootVariable.EntityData.YListKeys = []string {"NodeName"}

    return &(bootVariable.EntityData)
}

// Install_Software
// Software package,component and alias information
type Install_Software struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package alias information.
    AliasDevices Install_Software_AliasDevices

    // Package information.
    PackageDevices Install_Software_PackageDevices

    // Software component information.
    ComponentDevices Install_Software_ComponentDevices
}

func (software *Install_Software) GetEntityData() *types.CommonEntityData {
    software.EntityData.YFilter = software.YFilter
    software.EntityData.YangName = "software"
    software.EntityData.BundleName = "cisco_ios_xr"
    software.EntityData.ParentYangName = "install"
    software.EntityData.SegmentPath = "software"
    software.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/" + software.EntityData.SegmentPath
    software.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    software.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    software.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    software.EntityData.Children = types.NewOrderedMap()
    software.EntityData.Children.Append("alias-devices", types.YChild{"AliasDevices", &software.AliasDevices})
    software.EntityData.Children.Append("package-devices", types.YChild{"PackageDevices", &software.PackageDevices})
    software.EntityData.Children.Append("component-devices", types.YChild{"ComponentDevices", &software.ComponentDevices})
    software.EntityData.Leafs = types.NewOrderedMap()

    software.EntityData.YListKeys = []string {}

    return &(software.EntityData)
}

// Install_Software_AliasDevices
// Package alias information
type Install_Software_AliasDevices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package alias information for specific device. The type is slice of
    // Install_Software_AliasDevices_AliasDevice.
    AliasDevice []*Install_Software_AliasDevices_AliasDevice
}

func (aliasDevices *Install_Software_AliasDevices) GetEntityData() *types.CommonEntityData {
    aliasDevices.EntityData.YFilter = aliasDevices.YFilter
    aliasDevices.EntityData.YangName = "alias-devices"
    aliasDevices.EntityData.BundleName = "cisco_ios_xr"
    aliasDevices.EntityData.ParentYangName = "software"
    aliasDevices.EntityData.SegmentPath = "alias-devices"
    aliasDevices.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/" + aliasDevices.EntityData.SegmentPath
    aliasDevices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aliasDevices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aliasDevices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aliasDevices.EntityData.Children = types.NewOrderedMap()
    aliasDevices.EntityData.Children.Append("alias-device", types.YChild{"AliasDevice", nil})
    for i := range aliasDevices.AliasDevice {
        aliasDevices.EntityData.Children.Append(types.GetSegmentPath(aliasDevices.AliasDevice[i]), types.YChild{"AliasDevice", aliasDevices.AliasDevice[i]})
    }
    aliasDevices.EntityData.Leafs = types.NewOrderedMap()

    aliasDevices.EntityData.YListKeys = []string {}

    return &(aliasDevices.EntityData)
}

// Install_Software_AliasDevices_AliasDevice
// Package alias information for specific device
type Install_Software_AliasDevices_AliasDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Device Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    DeviceName interface{}

    // alias information.
    Aliases Install_Software_AliasDevices_AliasDevice_Aliases
}

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetEntityData() *types.CommonEntityData {
    aliasDevice.EntityData.YFilter = aliasDevice.YFilter
    aliasDevice.EntityData.YangName = "alias-device"
    aliasDevice.EntityData.BundleName = "cisco_ios_xr"
    aliasDevice.EntityData.ParentYangName = "alias-devices"
    aliasDevice.EntityData.SegmentPath = "alias-device" + types.AddKeyToken(aliasDevice.DeviceName, "device-name")
    aliasDevice.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/alias-devices/" + aliasDevice.EntityData.SegmentPath
    aliasDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aliasDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aliasDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aliasDevice.EntityData.Children = types.NewOrderedMap()
    aliasDevice.EntityData.Children.Append("aliases", types.YChild{"Aliases", &aliasDevice.Aliases})
    aliasDevice.EntityData.Leafs = types.NewOrderedMap()
    aliasDevice.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", aliasDevice.DeviceName})

    aliasDevice.EntityData.YListKeys = []string {"DeviceName"}

    return &(aliasDevice.EntityData)
}

// Install_Software_AliasDevices_AliasDevice_Aliases
// alias information
type Install_Software_AliasDevices_AliasDevice_Aliases struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Aliases for specific package. The type is slice of
    // Install_Software_AliasDevices_AliasDevice_Aliases_Alias.
    Alias []*Install_Software_AliasDevices_AliasDevice_Aliases_Alias
}

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetEntityData() *types.CommonEntityData {
    aliases.EntityData.YFilter = aliases.YFilter
    aliases.EntityData.YangName = "aliases"
    aliases.EntityData.BundleName = "cisco_ios_xr"
    aliases.EntityData.ParentYangName = "alias-device"
    aliases.EntityData.SegmentPath = "aliases"
    aliases.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/alias-devices/alias-device/" + aliases.EntityData.SegmentPath
    aliases.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aliases.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aliases.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aliases.EntityData.Children = types.NewOrderedMap()
    aliases.EntityData.Children.Append("alias", types.YChild{"Alias", nil})
    for i := range aliases.Alias {
        aliases.EntityData.Children.Append(types.GetSegmentPath(aliases.Alias[i]), types.YChild{"Alias", aliases.Alias[i]})
    }
    aliases.EntityData.Leafs = types.NewOrderedMap()

    aliases.EntityData.YListKeys = []string {}

    return &(aliases.EntityData)
}

// Install_Software_AliasDevices_AliasDevice_Aliases_Alias
// Aliases for specific package
type Install_Software_AliasDevices_AliasDevice_Aliases_Alias struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Package Name. The type is string.
    PackageName interface{}

    // Alias names. The type is string. This attribute is mandatory.
    AliasNames interface{}
}

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetEntityData() *types.CommonEntityData {
    alias.EntityData.YFilter = alias.YFilter
    alias.EntityData.YangName = "alias"
    alias.EntityData.BundleName = "cisco_ios_xr"
    alias.EntityData.ParentYangName = "aliases"
    alias.EntityData.SegmentPath = "alias" + types.AddKeyToken(alias.PackageName, "package-name")
    alias.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/alias-devices/alias-device/aliases/" + alias.EntityData.SegmentPath
    alias.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alias.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alias.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alias.EntityData.Children = types.NewOrderedMap()
    alias.EntityData.Leafs = types.NewOrderedMap()
    alias.EntityData.Leafs.Append("package-name", types.YLeaf{"PackageName", alias.PackageName})
    alias.EntityData.Leafs.Append("alias-names", types.YLeaf{"AliasNames", alias.AliasNames})

    alias.EntityData.YListKeys = []string {"PackageName"}

    return &(alias.EntityData)
}

// Install_Software_PackageDevices
// Package information
type Install_Software_PackageDevices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package information for specific device. The type is slice of
    // Install_Software_PackageDevices_PackageDevice.
    PackageDevice []*Install_Software_PackageDevices_PackageDevice
}

func (packageDevices *Install_Software_PackageDevices) GetEntityData() *types.CommonEntityData {
    packageDevices.EntityData.YFilter = packageDevices.YFilter
    packageDevices.EntityData.YangName = "package-devices"
    packageDevices.EntityData.BundleName = "cisco_ios_xr"
    packageDevices.EntityData.ParentYangName = "software"
    packageDevices.EntityData.SegmentPath = "package-devices"
    packageDevices.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/" + packageDevices.EntityData.SegmentPath
    packageDevices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packageDevices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packageDevices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packageDevices.EntityData.Children = types.NewOrderedMap()
    packageDevices.EntityData.Children.Append("package-device", types.YChild{"PackageDevice", nil})
    for i := range packageDevices.PackageDevice {
        packageDevices.EntityData.Children.Append(types.GetSegmentPath(packageDevices.PackageDevice[i]), types.YChild{"PackageDevice", packageDevices.PackageDevice[i]})
    }
    packageDevices.EntityData.Leafs = types.NewOrderedMap()

    packageDevices.EntityData.YListKeys = []string {}

    return &(packageDevices.EntityData)
}

// Install_Software_PackageDevices_PackageDevice
// Package information for specific device
type Install_Software_PackageDevices_PackageDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Device Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    DeviceName interface{}

    // Package information for specific package.
    Packages Install_Software_PackageDevices_PackageDevice_Packages
}

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetEntityData() *types.CommonEntityData {
    packageDevice.EntityData.YFilter = packageDevice.YFilter
    packageDevice.EntityData.YangName = "package-device"
    packageDevice.EntityData.BundleName = "cisco_ios_xr"
    packageDevice.EntityData.ParentYangName = "package-devices"
    packageDevice.EntityData.SegmentPath = "package-device" + types.AddKeyToken(packageDevice.DeviceName, "device-name")
    packageDevice.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/package-devices/" + packageDevice.EntityData.SegmentPath
    packageDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packageDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packageDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packageDevice.EntityData.Children = types.NewOrderedMap()
    packageDevice.EntityData.Children.Append("packages", types.YChild{"Packages", &packageDevice.Packages})
    packageDevice.EntityData.Leafs = types.NewOrderedMap()
    packageDevice.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", packageDevice.DeviceName})

    packageDevice.EntityData.YListKeys = []string {"DeviceName"}

    return &(packageDevice.EntityData)
}

// Install_Software_PackageDevices_PackageDevice_Packages
// Package information for specific package
type Install_Software_PackageDevices_PackageDevice_Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package information. The type is slice of
    // Install_Software_PackageDevices_PackageDevice_Packages_Package.
    Package []*Install_Software_PackageDevices_PackageDevice_Packages_Package
}

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "package-device"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/package-devices/package-device/" + packages.EntityData.SegmentPath
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = types.NewOrderedMap()
    packages.EntityData.Children.Append("package", types.YChild{"Package", nil})
    for i := range packages.Package {
        packages.EntityData.Children.Append(types.GetSegmentPath(packages.Package[i]), types.YChild{"Package", packages.Package[i]})
    }
    packages.EntityData.Leafs = types.NewOrderedMap()

    packages.EntityData.YListKeys = []string {}

    return &(packages.EntityData)
}

// Install_Software_PackageDevices_PackageDevice_Packages_Package
// Package information
type Install_Software_PackageDevices_PackageDevice_Packages_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Package Name. The type is string.
    PackageName interface{}

    // Name of the package. The type is string.
    Name interface{}

    // Version of the package. The type is string.
    Version interface{}

    // Description of the package. The type is string.
    Description interface{}

    // Release that the package belongs to. The type is string.
    Release interface{}

    // Information about the vendor that supplied the package. The type is string.
    Vendor interface{}

    // Time and date that the package was built. The type is string.
    Date interface{}

    // Identifies the provider of the package. The type is string.
    Source interface{}

    // Identifies the base bundle that the package is for. The type is string.
    Base interface{}

    // TRUE if package has BOOTIMAGE tag set. The type is bool.
    Bootable interface{}

    // TRUE if package is a composite package. The type is bool.
    Composite interface{}

    // Restart info of the package. The type is string.
    RestartInfo interface{}

    // Type of the package. The type is InstmgrPkg.
    PackageType interface{}

    // Group type of the package. The type is InstmgrGroup.
    GroupType interface{}

    // Number of layers of parent packages. The type is interface{} with range:
    // 0..4294967295.
    Depth interface{}

    // Uncompressed size of package. The type is interface{} with range:
    // 0..4294967295.
    UncompressedSize interface{}

    // Compressed size of package. The type is interface{} with range:
    // 0..4294967295.
    CompressedSize interface{}

    // Card types that the package should be installed on. The type is slice of
    // string.
    Cards []interface{}

    // Sub-package info of the package. The type is slice of
    // Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg.
    SubPkg []*Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg
}

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "packages"
    self.EntityData.SegmentPath = "package" + types.AddKeyToken(self.PackageName, "package-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/package-devices/package-device/packages/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("sub-pkg", types.YChild{"SubPkg", nil})
    for i := range self.SubPkg {
        types.SetYListKey(self.SubPkg[i], i)
        self.EntityData.Children.Append(types.GetSegmentPath(self.SubPkg[i]), types.YChild{"SubPkg", self.SubPkg[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("package-name", types.YLeaf{"PackageName", self.PackageName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})
    self.EntityData.Leafs.Append("version", types.YLeaf{"Version", self.Version})
    self.EntityData.Leafs.Append("description", types.YLeaf{"Description", self.Description})
    self.EntityData.Leafs.Append("release", types.YLeaf{"Release", self.Release})
    self.EntityData.Leafs.Append("vendor", types.YLeaf{"Vendor", self.Vendor})
    self.EntityData.Leafs.Append("date", types.YLeaf{"Date", self.Date})
    self.EntityData.Leafs.Append("source", types.YLeaf{"Source", self.Source})
    self.EntityData.Leafs.Append("base", types.YLeaf{"Base", self.Base})
    self.EntityData.Leafs.Append("bootable", types.YLeaf{"Bootable", self.Bootable})
    self.EntityData.Leafs.Append("composite", types.YLeaf{"Composite", self.Composite})
    self.EntityData.Leafs.Append("restart-info", types.YLeaf{"RestartInfo", self.RestartInfo})
    self.EntityData.Leafs.Append("package-type", types.YLeaf{"PackageType", self.PackageType})
    self.EntityData.Leafs.Append("group-type", types.YLeaf{"GroupType", self.GroupType})
    self.EntityData.Leafs.Append("depth", types.YLeaf{"Depth", self.Depth})
    self.EntityData.Leafs.Append("uncompressed-size", types.YLeaf{"UncompressedSize", self.UncompressedSize})
    self.EntityData.Leafs.Append("compressed-size", types.YLeaf{"CompressedSize", self.CompressedSize})
    self.EntityData.Leafs.Append("cards", types.YLeaf{"Cards", self.Cards})

    self.EntityData.YListKeys = []string {"PackageName"}

    return &(self.EntityData)
}

// Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg
// Sub-package info of the package
type Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Name of the sub-package. The type is string.
    Name interface{}

    // Node types of the package. The type is interface{} with range:
    // 0..18446744073709551615.
    NodeTypes interface{}
}

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetEntityData() *types.CommonEntityData {
    subPkg.EntityData.YFilter = subPkg.YFilter
    subPkg.EntityData.YangName = "sub-pkg"
    subPkg.EntityData.BundleName = "cisco_ios_xr"
    subPkg.EntityData.ParentYangName = "package"
    subPkg.EntityData.SegmentPath = "sub-pkg" + types.AddNoKeyToken(subPkg)
    subPkg.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/package-devices/package-device/packages/package/" + subPkg.EntityData.SegmentPath
    subPkg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subPkg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subPkg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subPkg.EntityData.Children = types.NewOrderedMap()
    subPkg.EntityData.Leafs = types.NewOrderedMap()
    subPkg.EntityData.Leafs.Append("name", types.YLeaf{"Name", subPkg.Name})
    subPkg.EntityData.Leafs.Append("node-types", types.YLeaf{"NodeTypes", subPkg.NodeTypes})

    subPkg.EntityData.YListKeys = []string {}

    return &(subPkg.EntityData)
}

// Install_Software_ComponentDevices
// Software component information
type Install_Software_ComponentDevices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Component information for specific device. The type is slice of
    // Install_Software_ComponentDevices_ComponentDevice.
    ComponentDevice []*Install_Software_ComponentDevices_ComponentDevice
}

func (componentDevices *Install_Software_ComponentDevices) GetEntityData() *types.CommonEntityData {
    componentDevices.EntityData.YFilter = componentDevices.YFilter
    componentDevices.EntityData.YangName = "component-devices"
    componentDevices.EntityData.BundleName = "cisco_ios_xr"
    componentDevices.EntityData.ParentYangName = "software"
    componentDevices.EntityData.SegmentPath = "component-devices"
    componentDevices.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/" + componentDevices.EntityData.SegmentPath
    componentDevices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    componentDevices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    componentDevices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    componentDevices.EntityData.Children = types.NewOrderedMap()
    componentDevices.EntityData.Children.Append("component-device", types.YChild{"ComponentDevice", nil})
    for i := range componentDevices.ComponentDevice {
        componentDevices.EntityData.Children.Append(types.GetSegmentPath(componentDevices.ComponentDevice[i]), types.YChild{"ComponentDevice", componentDevices.ComponentDevice[i]})
    }
    componentDevices.EntityData.Leafs = types.NewOrderedMap()

    componentDevices.EntityData.YListKeys = []string {}

    return &(componentDevices.EntityData)
}

// Install_Software_ComponentDevices_ComponentDevice
// Component information for specific device
type Install_Software_ComponentDevices_ComponentDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Device Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    DeviceName interface{}

    // Software package information.
    ComponentPackages Install_Software_ComponentDevices_ComponentDevice_ComponentPackages
}

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetEntityData() *types.CommonEntityData {
    componentDevice.EntityData.YFilter = componentDevice.YFilter
    componentDevice.EntityData.YangName = "component-device"
    componentDevice.EntityData.BundleName = "cisco_ios_xr"
    componentDevice.EntityData.ParentYangName = "component-devices"
    componentDevice.EntityData.SegmentPath = "component-device" + types.AddKeyToken(componentDevice.DeviceName, "device-name")
    componentDevice.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/component-devices/" + componentDevice.EntityData.SegmentPath
    componentDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    componentDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    componentDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    componentDevice.EntityData.Children = types.NewOrderedMap()
    componentDevice.EntityData.Children.Append("component-packages", types.YChild{"ComponentPackages", &componentDevice.ComponentPackages})
    componentDevice.EntityData.Leafs = types.NewOrderedMap()
    componentDevice.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", componentDevice.DeviceName})

    componentDevice.EntityData.YListKeys = []string {"DeviceName"}

    return &(componentDevice.EntityData)
}

// Install_Software_ComponentDevices_ComponentDevice_ComponentPackages
// Software package information
type Install_Software_ComponentDevices_ComponentDevice_ComponentPackages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Component information for specific package. The type is slice of
    // Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage.
    ComponentPackage []*Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage
}

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetEntityData() *types.CommonEntityData {
    componentPackages.EntityData.YFilter = componentPackages.YFilter
    componentPackages.EntityData.YangName = "component-packages"
    componentPackages.EntityData.BundleName = "cisco_ios_xr"
    componentPackages.EntityData.ParentYangName = "component-device"
    componentPackages.EntityData.SegmentPath = "component-packages"
    componentPackages.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/component-devices/component-device/" + componentPackages.EntityData.SegmentPath
    componentPackages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    componentPackages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    componentPackages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    componentPackages.EntityData.Children = types.NewOrderedMap()
    componentPackages.EntityData.Children.Append("component-package", types.YChild{"ComponentPackage", nil})
    for i := range componentPackages.ComponentPackage {
        componentPackages.EntityData.Children.Append(types.GetSegmentPath(componentPackages.ComponentPackage[i]), types.YChild{"ComponentPackage", componentPackages.ComponentPackage[i]})
    }
    componentPackages.EntityData.Leafs = types.NewOrderedMap()

    componentPackages.EntityData.YListKeys = []string {}

    return &(componentPackages.EntityData)
}

// Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage
// Component information for specific package
type Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Package Name. The type is string.
    PackageName interface{}

    // Component information. The type is slice of
    // Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component.
    Component []*Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component
}

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetEntityData() *types.CommonEntityData {
    componentPackage.EntityData.YFilter = componentPackage.YFilter
    componentPackage.EntityData.YangName = "component-package"
    componentPackage.EntityData.BundleName = "cisco_ios_xr"
    componentPackage.EntityData.ParentYangName = "component-packages"
    componentPackage.EntityData.SegmentPath = "component-package" + types.AddKeyToken(componentPackage.PackageName, "package-name")
    componentPackage.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/component-devices/component-device/component-packages/" + componentPackage.EntityData.SegmentPath
    componentPackage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    componentPackage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    componentPackage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    componentPackage.EntityData.Children = types.NewOrderedMap()
    componentPackage.EntityData.Children.Append("component", types.YChild{"Component", nil})
    for i := range componentPackage.Component {
        componentPackage.EntityData.Children.Append(types.GetSegmentPath(componentPackage.Component[i]), types.YChild{"Component", componentPackage.Component[i]})
    }
    componentPackage.EntityData.Leafs = types.NewOrderedMap()
    componentPackage.EntityData.Leafs.Append("package-name", types.YLeaf{"PackageName", componentPackage.PackageName})

    componentPackage.EntityData.YListKeys = []string {"PackageName"}

    return &(componentPackage.EntityData)
}

// Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component
// Component information
type Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Component Name. The type is string.
    ComponentName interface{}

    // Name of the component. The type is string.
    Name interface{}

    // Version of the component. The type is string.
    Version interface{}

    // Release that the component belongs to. The type is string.
    Release interface{}

    // Description of the component. The type is string.
    Description interface{}

    // The set of files belonging to the component. The type is slice of string.
    Files []interface{}
}

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetEntityData() *types.CommonEntityData {
    component.EntityData.YFilter = component.YFilter
    component.EntityData.YangName = "component"
    component.EntityData.BundleName = "cisco_ios_xr"
    component.EntityData.ParentYangName = "component-package"
    component.EntityData.SegmentPath = "component" + types.AddKeyToken(component.ComponentName, "component-name")
    component.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software/component-devices/component-device/component-packages/component-package/" + component.EntityData.SegmentPath
    component.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    component.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    component.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    component.EntityData.Children = types.NewOrderedMap()
    component.EntityData.Leafs = types.NewOrderedMap()
    component.EntityData.Leafs.Append("component-name", types.YLeaf{"ComponentName", component.ComponentName})
    component.EntityData.Leafs.Append("name", types.YLeaf{"Name", component.Name})
    component.EntityData.Leafs.Append("version", types.YLeaf{"Version", component.Version})
    component.EntityData.Leafs.Append("release", types.YLeaf{"Release", component.Release})
    component.EntityData.Leafs.Append("description", types.YLeaf{"Description", component.Description})
    component.EntityData.Leafs.Append("files", types.YLeaf{"Files", component.Files})

    component.EntityData.YListKeys = []string {"ComponentName"}

    return &(component.EntityData)
}

// Install_SoftwareInventory
// Information of install operations
type Install_SoftwareInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Committed inventory information.
    Committed Install_SoftwareInventory_Committed

    // Inactive inventory information.
    Inactive Install_SoftwareInventory_Inactive

    // Install operation requests.
    Requests Install_SoftwareInventory_Requests

    // Active inventory information.
    Active Install_SoftwareInventory_Active
}

func (softwareInventory *Install_SoftwareInventory) GetEntityData() *types.CommonEntityData {
    softwareInventory.EntityData.YFilter = softwareInventory.YFilter
    softwareInventory.EntityData.YangName = "software-inventory"
    softwareInventory.EntityData.BundleName = "cisco_ios_xr"
    softwareInventory.EntityData.ParentYangName = "install"
    softwareInventory.EntityData.SegmentPath = "software-inventory"
    softwareInventory.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/" + softwareInventory.EntityData.SegmentPath
    softwareInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    softwareInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    softwareInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    softwareInventory.EntityData.Children = types.NewOrderedMap()
    softwareInventory.EntityData.Children.Append("committed", types.YChild{"Committed", &softwareInventory.Committed})
    softwareInventory.EntityData.Children.Append("inactive", types.YChild{"Inactive", &softwareInventory.Inactive})
    softwareInventory.EntityData.Children.Append("requests", types.YChild{"Requests", &softwareInventory.Requests})
    softwareInventory.EntityData.Children.Append("active", types.YChild{"Active", &softwareInventory.Active})
    softwareInventory.EntityData.Leafs = types.NewOrderedMap()

    softwareInventory.EntityData.YListKeys = []string {}

    return &(softwareInventory.EntityData)
}

// Install_SoftwareInventory_Committed
// Committed inventory information
type Install_SoftwareInventory_Committed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summarized inventory information.
    Summary Install_SoftwareInventory_Committed_Summary

    // Software inventory.
    Inventories Install_SoftwareInventory_Committed_Inventories
}

func (committed *Install_SoftwareInventory_Committed) GetEntityData() *types.CommonEntityData {
    committed.EntityData.YFilter = committed.YFilter
    committed.EntityData.YangName = "committed"
    committed.EntityData.BundleName = "cisco_ios_xr"
    committed.EntityData.ParentYangName = "software-inventory"
    committed.EntityData.SegmentPath = "committed"
    committed.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/" + committed.EntityData.SegmentPath
    committed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committed.EntityData.Children = types.NewOrderedMap()
    committed.EntityData.Children.Append("summary", types.YChild{"Summary", &committed.Summary})
    committed.EntityData.Children.Append("inventories", types.YChild{"Inventories", &committed.Inventories})
    committed.EntityData.Leafs = types.NewOrderedMap()

    committed.EntityData.YListKeys = []string {}

    return &(committed.EntityData)
}

// Install_SoftwareInventory_Committed_Summary
// Summarized inventory information
type Install_SoftwareInventory_Committed_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default load path.
    DefaultLoadPath Install_SoftwareInventory_Committed_Summary_DefaultLoadPath

    // Admin Resources load path.
    AdminLoadPath Install_SoftwareInventory_Committed_Summary_AdminLoadPath

    // SDR load paths. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_SdrLoadPath.
    SdrLoadPath []*Install_SoftwareInventory_Committed_Summary_SdrLoadPath

    // Location load paths. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_LocationLoadPath.
    LocationLoadPath []*Install_SoftwareInventory_Committed_Summary_LocationLoadPath
}

func (summary *Install_SoftwareInventory_Committed_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "committed"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("default-load-path", types.YChild{"DefaultLoadPath", &summary.DefaultLoadPath})
    summary.EntityData.Children.Append("admin-load-path", types.YChild{"AdminLoadPath", &summary.AdminLoadPath})
    summary.EntityData.Children.Append("sdr-load-path", types.YChild{"SdrLoadPath", nil})
    for i := range summary.SdrLoadPath {
        types.SetYListKey(summary.SdrLoadPath[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.SdrLoadPath[i]), types.YChild{"SdrLoadPath", summary.SdrLoadPath[i]})
    }
    summary.EntityData.Children.Append("location-load-path", types.YChild{"LocationLoadPath", nil})
    for i := range summary.LocationLoadPath {
        types.SetYListKey(summary.LocationLoadPath[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.LocationLoadPath[i]), types.YChild{"LocationLoadPath", summary.LocationLoadPath[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath
// Default load path
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Does this match the Admin Resources load path?. The type is bool.
    AdminMatch interface{}

    // Names of SDRs this applies to. The type is slice of string.
    SecureDomainRouterName []interface{}

    // Default load path. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath
}

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetEntityData() *types.CommonEntityData {
    defaultLoadPath.EntityData.YFilter = defaultLoadPath.YFilter
    defaultLoadPath.EntityData.YangName = "default-load-path"
    defaultLoadPath.EntityData.BundleName = "cisco_ios_xr"
    defaultLoadPath.EntityData.ParentYangName = "summary"
    defaultLoadPath.EntityData.SegmentPath = "default-load-path"
    defaultLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/" + defaultLoadPath.EntityData.SegmentPath
    defaultLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultLoadPath.EntityData.Children = types.NewOrderedMap()
    defaultLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range defaultLoadPath.LoadPath {
        types.SetYListKey(defaultLoadPath.LoadPath[i], i)
        defaultLoadPath.EntityData.Children.Append(types.GetSegmentPath(defaultLoadPath.LoadPath[i]), types.YChild{"LoadPath", defaultLoadPath.LoadPath[i]})
    }
    defaultLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range defaultLoadPath.StandbyLoadPath {
        types.SetYListKey(defaultLoadPath.StandbyLoadPath[i], i)
        defaultLoadPath.EntityData.Children.Append(types.GetSegmentPath(defaultLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", defaultLoadPath.StandbyLoadPath[i]})
    }
    defaultLoadPath.EntityData.Leafs = types.NewOrderedMap()
    defaultLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", defaultLoadPath.RequestId})
    defaultLoadPath.EntityData.Leafs.Append("admin-match", types.YLeaf{"AdminMatch", defaultLoadPath.AdminMatch})
    defaultLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", defaultLoadPath.SecureDomainRouterName})

    defaultLoadPath.EntityData.YListKeys = []string {}

    return &(defaultLoadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath
// Default load path
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "default-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/default-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/default-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "default-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/default-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/default-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath
// Admin Resources load path
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Admin Resources load path. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath
}

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetEntityData() *types.CommonEntityData {
    adminLoadPath.EntityData.YFilter = adminLoadPath.YFilter
    adminLoadPath.EntityData.YangName = "admin-load-path"
    adminLoadPath.EntityData.BundleName = "cisco_ios_xr"
    adminLoadPath.EntityData.ParentYangName = "summary"
    adminLoadPath.EntityData.SegmentPath = "admin-load-path"
    adminLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/" + adminLoadPath.EntityData.SegmentPath
    adminLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminLoadPath.EntityData.Children = types.NewOrderedMap()
    adminLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range adminLoadPath.LoadPath {
        types.SetYListKey(adminLoadPath.LoadPath[i], i)
        adminLoadPath.EntityData.Children.Append(types.GetSegmentPath(adminLoadPath.LoadPath[i]), types.YChild{"LoadPath", adminLoadPath.LoadPath[i]})
    }
    adminLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range adminLoadPath.StandbyLoadPath {
        types.SetYListKey(adminLoadPath.StandbyLoadPath[i], i)
        adminLoadPath.EntityData.Children.Append(types.GetSegmentPath(adminLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", adminLoadPath.StandbyLoadPath[i]})
    }
    adminLoadPath.EntityData.Leafs = types.NewOrderedMap()
    adminLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", adminLoadPath.RequestId})

    adminLoadPath.EntityData.YListKeys = []string {}

    return &(adminLoadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath
// Admin Resources load path
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "admin-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/admin-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/admin-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "admin-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/admin-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/admin-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath
// SDR load paths
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath
}

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetEntityData() *types.CommonEntityData {
    sdrLoadPath.EntityData.YFilter = sdrLoadPath.YFilter
    sdrLoadPath.EntityData.YangName = "sdr-load-path"
    sdrLoadPath.EntityData.BundleName = "cisco_ios_xr"
    sdrLoadPath.EntityData.ParentYangName = "summary"
    sdrLoadPath.EntityData.SegmentPath = "sdr-load-path" + types.AddNoKeyToken(sdrLoadPath)
    sdrLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/" + sdrLoadPath.EntityData.SegmentPath
    sdrLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrLoadPath.EntityData.Children = types.NewOrderedMap()
    sdrLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range sdrLoadPath.LoadPath {
        types.SetYListKey(sdrLoadPath.LoadPath[i], i)
        sdrLoadPath.EntityData.Children.Append(types.GetSegmentPath(sdrLoadPath.LoadPath[i]), types.YChild{"LoadPath", sdrLoadPath.LoadPath[i]})
    }
    sdrLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range sdrLoadPath.StandbyLoadPath {
        types.SetYListKey(sdrLoadPath.StandbyLoadPath[i], i)
        sdrLoadPath.EntityData.Children.Append(types.GetSegmentPath(sdrLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", sdrLoadPath.StandbyLoadPath[i]})
    }
    sdrLoadPath.EntityData.Leafs = types.NewOrderedMap()
    sdrLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", sdrLoadPath.RequestId})
    sdrLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", sdrLoadPath.SecureDomainRouterName})

    sdrLoadPath.EntityData.YListKeys = []string {}

    return &(sdrLoadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "sdr-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/sdr-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/sdr-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "sdr-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/sdr-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/sdr-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath
// Location load paths
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Node identifier. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath
}

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetEntityData() *types.CommonEntityData {
    locationLoadPath.EntityData.YFilter = locationLoadPath.YFilter
    locationLoadPath.EntityData.YangName = "location-load-path"
    locationLoadPath.EntityData.BundleName = "cisco_ios_xr"
    locationLoadPath.EntityData.ParentYangName = "summary"
    locationLoadPath.EntityData.SegmentPath = "location-load-path" + types.AddNoKeyToken(locationLoadPath)
    locationLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/" + locationLoadPath.EntityData.SegmentPath
    locationLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationLoadPath.EntityData.Children = types.NewOrderedMap()
    locationLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range locationLoadPath.LoadPath {
        types.SetYListKey(locationLoadPath.LoadPath[i], i)
        locationLoadPath.EntityData.Children.Append(types.GetSegmentPath(locationLoadPath.LoadPath[i]), types.YChild{"LoadPath", locationLoadPath.LoadPath[i]})
    }
    locationLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range locationLoadPath.StandbyLoadPath {
        types.SetYListKey(locationLoadPath.StandbyLoadPath[i], i)
        locationLoadPath.EntityData.Children.Append(types.GetSegmentPath(locationLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", locationLoadPath.StandbyLoadPath[i]})
    }
    locationLoadPath.EntityData.Leafs = types.NewOrderedMap()
    locationLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", locationLoadPath.RequestId})
    locationLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", locationLoadPath.SecureDomainRouterName})
    locationLoadPath.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", locationLoadPath.NodeName})

    locationLoadPath.EntityData.YListKeys = []string {}

    return &(locationLoadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "location-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/location-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/location-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "location-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/location-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/summary/location-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Committed_Inventories
// Software inventory
type Install_SoftwareInventory_Committed_Inventories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information for specific node. The type is slice of
    // Install_SoftwareInventory_Committed_Inventories_Inventory.
    Inventory []*Install_SoftwareInventory_Committed_Inventories_Inventory
}

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetEntityData() *types.CommonEntityData {
    inventories.EntityData.YFilter = inventories.YFilter
    inventories.EntityData.YangName = "inventories"
    inventories.EntityData.BundleName = "cisco_ios_xr"
    inventories.EntityData.ParentYangName = "committed"
    inventories.EntityData.SegmentPath = "inventories"
    inventories.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/" + inventories.EntityData.SegmentPath
    inventories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventories.EntityData.Children = types.NewOrderedMap()
    inventories.EntityData.Children.Append("inventory", types.YChild{"Inventory", nil})
    for i := range inventories.Inventory {
        inventories.EntityData.Children.Append(types.GetSegmentPath(inventories.Inventory[i]), types.YChild{"Inventory", inventories.Inventory[i]})
    }
    inventories.EntityData.Leafs = types.NewOrderedMap()

    inventories.EntityData.YListKeys = []string {}

    return &(inventories.EntityData)
}

// Install_SoftwareInventory_Committed_Inventories_Inventory
// Inventory information for specific node
type Install_SoftwareInventory_Committed_Inventories_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Major data version number. The type is interface{} with range:
    // 0..4294967295.
    Major interface{}

    // Minor data version number. The type is interface{} with range:
    // 0..4294967295.
    Minor interface{}

    // Name of the boot image. The type is string.
    BootImageName interface{}

    // Node's type. The type is interface{} with range: 0..18446744073709551615.
    NodeType interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath.
    LoadPath []*Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath
}

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "inventories"
    inventory.EntityData.SegmentPath = "inventory" + types.AddKeyToken(inventory.NodeName, "node-name")
    inventory.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/inventories/" + inventory.EntityData.SegmentPath
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range inventory.LoadPath {
        types.SetYListKey(inventory.LoadPath[i], i)
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.LoadPath[i]), types.YChild{"LoadPath", inventory.LoadPath[i]})
    }
    inventory.EntityData.Leafs = types.NewOrderedMap()
    inventory.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", inventory.NodeName})
    inventory.EntityData.Leafs.Append("major", types.YLeaf{"Major", inventory.Major})
    inventory.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", inventory.Minor})
    inventory.EntityData.Leafs.Append("boot-image-name", types.YLeaf{"BootImageName", inventory.BootImageName})
    inventory.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", inventory.NodeType})
    inventory.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", inventory.SecureDomainRouterName})

    inventory.EntityData.YListKeys = []string {"NodeName"}

    return &(inventory.EntityData)
}

// Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath
// Load path
type Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "inventory"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/inventories/inventory/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/committed/inventories/inventory/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive
// Inactive inventory information
type Install_SoftwareInventory_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summarized inventory information.
    Summary Install_SoftwareInventory_Inactive_Summary

    // Software inventory.
    Inventories Install_SoftwareInventory_Inactive_Inventories
}

func (inactive *Install_SoftwareInventory_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "software-inventory"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/" + inactive.EntityData.SegmentPath
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Children.Append("summary", types.YChild{"Summary", &inactive.Summary})
    inactive.EntityData.Children.Append("inventories", types.YChild{"Inventories", &inactive.Inventories})
    inactive.EntityData.Leafs = types.NewOrderedMap()

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary
// Summarized inventory information
type Install_SoftwareInventory_Inactive_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default load path.
    DefaultLoadPath Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath

    // Admin Resources load path.
    AdminLoadPath Install_SoftwareInventory_Inactive_Summary_AdminLoadPath

    // SDR load paths. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_SdrLoadPath.
    SdrLoadPath []*Install_SoftwareInventory_Inactive_Summary_SdrLoadPath

    // Location load paths. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_LocationLoadPath.
    LocationLoadPath []*Install_SoftwareInventory_Inactive_Summary_LocationLoadPath
}

func (summary *Install_SoftwareInventory_Inactive_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "inactive"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("default-load-path", types.YChild{"DefaultLoadPath", &summary.DefaultLoadPath})
    summary.EntityData.Children.Append("admin-load-path", types.YChild{"AdminLoadPath", &summary.AdminLoadPath})
    summary.EntityData.Children.Append("sdr-load-path", types.YChild{"SdrLoadPath", nil})
    for i := range summary.SdrLoadPath {
        types.SetYListKey(summary.SdrLoadPath[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.SdrLoadPath[i]), types.YChild{"SdrLoadPath", summary.SdrLoadPath[i]})
    }
    summary.EntityData.Children.Append("location-load-path", types.YChild{"LocationLoadPath", nil})
    for i := range summary.LocationLoadPath {
        types.SetYListKey(summary.LocationLoadPath[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.LocationLoadPath[i]), types.YChild{"LocationLoadPath", summary.LocationLoadPath[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath
// Default load path
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Does this match the Admin Resources load path?. The type is bool.
    AdminMatch interface{}

    // Names of SDRs this applies to. The type is slice of string.
    SecureDomainRouterName []interface{}

    // Default load path. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath
}

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetEntityData() *types.CommonEntityData {
    defaultLoadPath.EntityData.YFilter = defaultLoadPath.YFilter
    defaultLoadPath.EntityData.YangName = "default-load-path"
    defaultLoadPath.EntityData.BundleName = "cisco_ios_xr"
    defaultLoadPath.EntityData.ParentYangName = "summary"
    defaultLoadPath.EntityData.SegmentPath = "default-load-path"
    defaultLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/" + defaultLoadPath.EntityData.SegmentPath
    defaultLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultLoadPath.EntityData.Children = types.NewOrderedMap()
    defaultLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range defaultLoadPath.LoadPath {
        types.SetYListKey(defaultLoadPath.LoadPath[i], i)
        defaultLoadPath.EntityData.Children.Append(types.GetSegmentPath(defaultLoadPath.LoadPath[i]), types.YChild{"LoadPath", defaultLoadPath.LoadPath[i]})
    }
    defaultLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range defaultLoadPath.StandbyLoadPath {
        types.SetYListKey(defaultLoadPath.StandbyLoadPath[i], i)
        defaultLoadPath.EntityData.Children.Append(types.GetSegmentPath(defaultLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", defaultLoadPath.StandbyLoadPath[i]})
    }
    defaultLoadPath.EntityData.Leafs = types.NewOrderedMap()
    defaultLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", defaultLoadPath.RequestId})
    defaultLoadPath.EntityData.Leafs.Append("admin-match", types.YLeaf{"AdminMatch", defaultLoadPath.AdminMatch})
    defaultLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", defaultLoadPath.SecureDomainRouterName})

    defaultLoadPath.EntityData.YListKeys = []string {}

    return &(defaultLoadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath
// Default load path
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "default-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/default-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/default-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "default-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/default-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/default-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath
// Admin Resources load path
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Admin Resources load path. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath
}

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetEntityData() *types.CommonEntityData {
    adminLoadPath.EntityData.YFilter = adminLoadPath.YFilter
    adminLoadPath.EntityData.YangName = "admin-load-path"
    adminLoadPath.EntityData.BundleName = "cisco_ios_xr"
    adminLoadPath.EntityData.ParentYangName = "summary"
    adminLoadPath.EntityData.SegmentPath = "admin-load-path"
    adminLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/" + adminLoadPath.EntityData.SegmentPath
    adminLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminLoadPath.EntityData.Children = types.NewOrderedMap()
    adminLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range adminLoadPath.LoadPath {
        types.SetYListKey(adminLoadPath.LoadPath[i], i)
        adminLoadPath.EntityData.Children.Append(types.GetSegmentPath(adminLoadPath.LoadPath[i]), types.YChild{"LoadPath", adminLoadPath.LoadPath[i]})
    }
    adminLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range adminLoadPath.StandbyLoadPath {
        types.SetYListKey(adminLoadPath.StandbyLoadPath[i], i)
        adminLoadPath.EntityData.Children.Append(types.GetSegmentPath(adminLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", adminLoadPath.StandbyLoadPath[i]})
    }
    adminLoadPath.EntityData.Leafs = types.NewOrderedMap()
    adminLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", adminLoadPath.RequestId})

    adminLoadPath.EntityData.YListKeys = []string {}

    return &(adminLoadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath
// Admin Resources load path
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "admin-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/admin-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/admin-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "admin-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/admin-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/admin-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath
// SDR load paths
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath
}

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetEntityData() *types.CommonEntityData {
    sdrLoadPath.EntityData.YFilter = sdrLoadPath.YFilter
    sdrLoadPath.EntityData.YangName = "sdr-load-path"
    sdrLoadPath.EntityData.BundleName = "cisco_ios_xr"
    sdrLoadPath.EntityData.ParentYangName = "summary"
    sdrLoadPath.EntityData.SegmentPath = "sdr-load-path" + types.AddNoKeyToken(sdrLoadPath)
    sdrLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/" + sdrLoadPath.EntityData.SegmentPath
    sdrLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrLoadPath.EntityData.Children = types.NewOrderedMap()
    sdrLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range sdrLoadPath.LoadPath {
        types.SetYListKey(sdrLoadPath.LoadPath[i], i)
        sdrLoadPath.EntityData.Children.Append(types.GetSegmentPath(sdrLoadPath.LoadPath[i]), types.YChild{"LoadPath", sdrLoadPath.LoadPath[i]})
    }
    sdrLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range sdrLoadPath.StandbyLoadPath {
        types.SetYListKey(sdrLoadPath.StandbyLoadPath[i], i)
        sdrLoadPath.EntityData.Children.Append(types.GetSegmentPath(sdrLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", sdrLoadPath.StandbyLoadPath[i]})
    }
    sdrLoadPath.EntityData.Leafs = types.NewOrderedMap()
    sdrLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", sdrLoadPath.RequestId})
    sdrLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", sdrLoadPath.SecureDomainRouterName})

    sdrLoadPath.EntityData.YListKeys = []string {}

    return &(sdrLoadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "sdr-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/sdr-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/sdr-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "sdr-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/sdr-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/sdr-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath
// Location load paths
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Node identifier. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath
}

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetEntityData() *types.CommonEntityData {
    locationLoadPath.EntityData.YFilter = locationLoadPath.YFilter
    locationLoadPath.EntityData.YangName = "location-load-path"
    locationLoadPath.EntityData.BundleName = "cisco_ios_xr"
    locationLoadPath.EntityData.ParentYangName = "summary"
    locationLoadPath.EntityData.SegmentPath = "location-load-path" + types.AddNoKeyToken(locationLoadPath)
    locationLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/" + locationLoadPath.EntityData.SegmentPath
    locationLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationLoadPath.EntityData.Children = types.NewOrderedMap()
    locationLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range locationLoadPath.LoadPath {
        types.SetYListKey(locationLoadPath.LoadPath[i], i)
        locationLoadPath.EntityData.Children.Append(types.GetSegmentPath(locationLoadPath.LoadPath[i]), types.YChild{"LoadPath", locationLoadPath.LoadPath[i]})
    }
    locationLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range locationLoadPath.StandbyLoadPath {
        types.SetYListKey(locationLoadPath.StandbyLoadPath[i], i)
        locationLoadPath.EntityData.Children.Append(types.GetSegmentPath(locationLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", locationLoadPath.StandbyLoadPath[i]})
    }
    locationLoadPath.EntityData.Leafs = types.NewOrderedMap()
    locationLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", locationLoadPath.RequestId})
    locationLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", locationLoadPath.SecureDomainRouterName})
    locationLoadPath.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", locationLoadPath.NodeName})

    locationLoadPath.EntityData.YListKeys = []string {}

    return &(locationLoadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "location-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/location-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/location-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "location-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/location-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/summary/location-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Inactive_Inventories
// Software inventory
type Install_SoftwareInventory_Inactive_Inventories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information for specific node. The type is slice of
    // Install_SoftwareInventory_Inactive_Inventories_Inventory.
    Inventory []*Install_SoftwareInventory_Inactive_Inventories_Inventory
}

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetEntityData() *types.CommonEntityData {
    inventories.EntityData.YFilter = inventories.YFilter
    inventories.EntityData.YangName = "inventories"
    inventories.EntityData.BundleName = "cisco_ios_xr"
    inventories.EntityData.ParentYangName = "inactive"
    inventories.EntityData.SegmentPath = "inventories"
    inventories.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/" + inventories.EntityData.SegmentPath
    inventories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventories.EntityData.Children = types.NewOrderedMap()
    inventories.EntityData.Children.Append("inventory", types.YChild{"Inventory", nil})
    for i := range inventories.Inventory {
        inventories.EntityData.Children.Append(types.GetSegmentPath(inventories.Inventory[i]), types.YChild{"Inventory", inventories.Inventory[i]})
    }
    inventories.EntityData.Leafs = types.NewOrderedMap()

    inventories.EntityData.YListKeys = []string {}

    return &(inventories.EntityData)
}

// Install_SoftwareInventory_Inactive_Inventories_Inventory
// Inventory information for specific node
type Install_SoftwareInventory_Inactive_Inventories_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Major data version number. The type is interface{} with range:
    // 0..4294967295.
    Major interface{}

    // Minor data version number. The type is interface{} with range:
    // 0..4294967295.
    Minor interface{}

    // Name of the boot image. The type is string.
    BootImageName interface{}

    // Node's type. The type is interface{} with range: 0..18446744073709551615.
    NodeType interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath.
    LoadPath []*Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath
}

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "inventories"
    inventory.EntityData.SegmentPath = "inventory" + types.AddKeyToken(inventory.NodeName, "node-name")
    inventory.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/inventories/" + inventory.EntityData.SegmentPath
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range inventory.LoadPath {
        types.SetYListKey(inventory.LoadPath[i], i)
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.LoadPath[i]), types.YChild{"LoadPath", inventory.LoadPath[i]})
    }
    inventory.EntityData.Leafs = types.NewOrderedMap()
    inventory.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", inventory.NodeName})
    inventory.EntityData.Leafs.Append("major", types.YLeaf{"Major", inventory.Major})
    inventory.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", inventory.Minor})
    inventory.EntityData.Leafs.Append("boot-image-name", types.YLeaf{"BootImageName", inventory.BootImageName})
    inventory.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", inventory.NodeType})
    inventory.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", inventory.SecureDomainRouterName})

    inventory.EntityData.YListKeys = []string {"NodeName"}

    return &(inventory.EntityData)
}

// Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath
// Load path
type Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "inventory"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/inventories/inventory/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/inactive/inventories/inventory/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Requests
// Install operation requests
type Install_SoftwareInventory_Requests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install operation request history.
    Requests Install_SoftwareInventory_Requests_Requests
}

func (requests *Install_SoftwareInventory_Requests) GetEntityData() *types.CommonEntityData {
    requests.EntityData.YFilter = requests.YFilter
    requests.EntityData.YangName = "requests"
    requests.EntityData.BundleName = "cisco_ios_xr"
    requests.EntityData.ParentYangName = "software-inventory"
    requests.EntityData.SegmentPath = "requests"
    requests.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/" + requests.EntityData.SegmentPath
    requests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requests.EntityData.Children = types.NewOrderedMap()
    requests.EntityData.Children.Append("requests", types.YChild{"Requests", &requests.Requests})
    requests.EntityData.Leafs = types.NewOrderedMap()

    requests.EntityData.YListKeys = []string {}

    return &(requests.EntityData)
}

// Install_SoftwareInventory_Requests_Requests
// Install operation request history
type Install_SoftwareInventory_Requests_Requests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install operation request information. The type is slice of
    // Install_SoftwareInventory_Requests_Requests_Request.
    Request []*Install_SoftwareInventory_Requests_Requests_Request
}

func (requests *Install_SoftwareInventory_Requests_Requests) GetEntityData() *types.CommonEntityData {
    requests.EntityData.YFilter = requests.YFilter
    requests.EntityData.YangName = "requests"
    requests.EntityData.BundleName = "cisco_ios_xr"
    requests.EntityData.ParentYangName = "requests"
    requests.EntityData.SegmentPath = "requests"
    requests.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/requests/" + requests.EntityData.SegmentPath
    requests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requests.EntityData.Children = types.NewOrderedMap()
    requests.EntityData.Children.Append("request", types.YChild{"Request", nil})
    for i := range requests.Request {
        requests.EntityData.Children.Append(types.GetSegmentPath(requests.Request[i]), types.YChild{"Request", requests.Request[i]})
    }
    requests.EntityData.Leafs = types.NewOrderedMap()

    requests.EntityData.YListKeys = []string {}

    return &(requests.EntityData)
}

// Install_SoftwareInventory_Requests_Requests_Request
// Install operation request information
type Install_SoftwareInventory_Requests_Requests_Request struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Install operation request ID. The type is
    // interface{} with range: 0..4294967295.
    RequestId interface{}

    // Inventory information of install operation request.
    Inventories Install_SoftwareInventory_Requests_Requests_Request_Inventories
}

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "requests"
    request.EntityData.SegmentPath = "request" + types.AddKeyToken(request.RequestId, "request-id")
    request.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/requests/requests/" + request.EntityData.SegmentPath
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = types.NewOrderedMap()
    request.EntityData.Children.Append("inventories", types.YChild{"Inventories", &request.Inventories})
    request.EntityData.Leafs = types.NewOrderedMap()
    request.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", request.RequestId})

    request.EntityData.YListKeys = []string {"RequestId"}

    return &(request.EntityData)
}

// Install_SoftwareInventory_Requests_Requests_Request_Inventories
// Inventory information of install operation
// request
type Install_SoftwareInventory_Requests_Requests_Request_Inventories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information. The type is slice of
    // Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory.
    Inventory []*Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory
}

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetEntityData() *types.CommonEntityData {
    inventories.EntityData.YFilter = inventories.YFilter
    inventories.EntityData.YangName = "inventories"
    inventories.EntityData.BundleName = "cisco_ios_xr"
    inventories.EntityData.ParentYangName = "request"
    inventories.EntityData.SegmentPath = "inventories"
    inventories.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/requests/requests/request/" + inventories.EntityData.SegmentPath
    inventories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventories.EntityData.Children = types.NewOrderedMap()
    inventories.EntityData.Children.Append("inventory", types.YChild{"Inventory", nil})
    for i := range inventories.Inventory {
        inventories.EntityData.Children.Append(types.GetSegmentPath(inventories.Inventory[i]), types.YChild{"Inventory", inventories.Inventory[i]})
    }
    inventories.EntityData.Leafs = types.NewOrderedMap()

    inventories.EntityData.YListKeys = []string {}

    return &(inventories.EntityData)
}

// Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory
// Inventory information
type Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Major data version number. The type is interface{} with range:
    // 0..4294967295.
    Major interface{}

    // Minor data version number. The type is interface{} with range:
    // 0..4294967295.
    Minor interface{}

    // Name of the boot image. The type is string.
    BootImageName interface{}

    // Node's type. The type is interface{} with range: 0..18446744073709551615.
    NodeType interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath.
    LoadPath []*Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath
}

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "inventories"
    inventory.EntityData.SegmentPath = "inventory" + types.AddKeyToken(inventory.NodeName, "node-name")
    inventory.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/requests/requests/request/inventories/" + inventory.EntityData.SegmentPath
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range inventory.LoadPath {
        types.SetYListKey(inventory.LoadPath[i], i)
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.LoadPath[i]), types.YChild{"LoadPath", inventory.LoadPath[i]})
    }
    inventory.EntityData.Leafs = types.NewOrderedMap()
    inventory.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", inventory.NodeName})
    inventory.EntityData.Leafs.Append("major", types.YLeaf{"Major", inventory.Major})
    inventory.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", inventory.Minor})
    inventory.EntityData.Leafs.Append("boot-image-name", types.YLeaf{"BootImageName", inventory.BootImageName})
    inventory.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", inventory.NodeType})
    inventory.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", inventory.SecureDomainRouterName})

    inventory.EntityData.YListKeys = []string {"NodeName"}

    return &(inventory.EntityData)
}

// Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath
// Load path
type Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "inventory"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/requests/requests/request/inventories/inventory/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package
// Package
type Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/requests/requests/request/inventories/inventory/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active
// Active inventory information
type Install_SoftwareInventory_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summarized inventory information.
    Summary Install_SoftwareInventory_Active_Summary

    // Software inventory.
    Inventories Install_SoftwareInventory_Active_Inventories
}

func (active *Install_SoftwareInventory_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "software-inventory"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("summary", types.YChild{"Summary", &active.Summary})
    active.EntityData.Children.Append("inventories", types.YChild{"Inventories", &active.Inventories})
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// Install_SoftwareInventory_Active_Summary
// Summarized inventory information
type Install_SoftwareInventory_Active_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default load path.
    DefaultLoadPath Install_SoftwareInventory_Active_Summary_DefaultLoadPath

    // Admin Resources load path.
    AdminLoadPath Install_SoftwareInventory_Active_Summary_AdminLoadPath

    // SDR load paths. The type is slice of
    // Install_SoftwareInventory_Active_Summary_SdrLoadPath.
    SdrLoadPath []*Install_SoftwareInventory_Active_Summary_SdrLoadPath

    // Location load paths. The type is slice of
    // Install_SoftwareInventory_Active_Summary_LocationLoadPath.
    LocationLoadPath []*Install_SoftwareInventory_Active_Summary_LocationLoadPath
}

func (summary *Install_SoftwareInventory_Active_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "active"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("default-load-path", types.YChild{"DefaultLoadPath", &summary.DefaultLoadPath})
    summary.EntityData.Children.Append("admin-load-path", types.YChild{"AdminLoadPath", &summary.AdminLoadPath})
    summary.EntityData.Children.Append("sdr-load-path", types.YChild{"SdrLoadPath", nil})
    for i := range summary.SdrLoadPath {
        types.SetYListKey(summary.SdrLoadPath[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.SdrLoadPath[i]), types.YChild{"SdrLoadPath", summary.SdrLoadPath[i]})
    }
    summary.EntityData.Children.Append("location-load-path", types.YChild{"LocationLoadPath", nil})
    for i := range summary.LocationLoadPath {
        types.SetYListKey(summary.LocationLoadPath[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.LocationLoadPath[i]), types.YChild{"LocationLoadPath", summary.LocationLoadPath[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath
// Default load path
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Does this match the Admin Resources load path?. The type is bool.
    AdminMatch interface{}

    // Names of SDRs this applies to. The type is slice of string.
    SecureDomainRouterName []interface{}

    // Default load path. The type is slice of
    // Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath
}

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetEntityData() *types.CommonEntityData {
    defaultLoadPath.EntityData.YFilter = defaultLoadPath.YFilter
    defaultLoadPath.EntityData.YangName = "default-load-path"
    defaultLoadPath.EntityData.BundleName = "cisco_ios_xr"
    defaultLoadPath.EntityData.ParentYangName = "summary"
    defaultLoadPath.EntityData.SegmentPath = "default-load-path"
    defaultLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/" + defaultLoadPath.EntityData.SegmentPath
    defaultLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultLoadPath.EntityData.Children = types.NewOrderedMap()
    defaultLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range defaultLoadPath.LoadPath {
        types.SetYListKey(defaultLoadPath.LoadPath[i], i)
        defaultLoadPath.EntityData.Children.Append(types.GetSegmentPath(defaultLoadPath.LoadPath[i]), types.YChild{"LoadPath", defaultLoadPath.LoadPath[i]})
    }
    defaultLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range defaultLoadPath.StandbyLoadPath {
        types.SetYListKey(defaultLoadPath.StandbyLoadPath[i], i)
        defaultLoadPath.EntityData.Children.Append(types.GetSegmentPath(defaultLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", defaultLoadPath.StandbyLoadPath[i]})
    }
    defaultLoadPath.EntityData.Leafs = types.NewOrderedMap()
    defaultLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", defaultLoadPath.RequestId})
    defaultLoadPath.EntityData.Leafs.Append("admin-match", types.YLeaf{"AdminMatch", defaultLoadPath.AdminMatch})
    defaultLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", defaultLoadPath.SecureDomainRouterName})

    defaultLoadPath.EntityData.YListKeys = []string {}

    return &(defaultLoadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath
// Default load path
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "default-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/default-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/default-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "default-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/default-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/default-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active_Summary_AdminLoadPath
// Admin Resources load path
type Install_SoftwareInventory_Active_Summary_AdminLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Admin Resources load path. The type is slice of
    // Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath
}

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetEntityData() *types.CommonEntityData {
    adminLoadPath.EntityData.YFilter = adminLoadPath.YFilter
    adminLoadPath.EntityData.YangName = "admin-load-path"
    adminLoadPath.EntityData.BundleName = "cisco_ios_xr"
    adminLoadPath.EntityData.ParentYangName = "summary"
    adminLoadPath.EntityData.SegmentPath = "admin-load-path"
    adminLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/" + adminLoadPath.EntityData.SegmentPath
    adminLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminLoadPath.EntityData.Children = types.NewOrderedMap()
    adminLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range adminLoadPath.LoadPath {
        types.SetYListKey(adminLoadPath.LoadPath[i], i)
        adminLoadPath.EntityData.Children.Append(types.GetSegmentPath(adminLoadPath.LoadPath[i]), types.YChild{"LoadPath", adminLoadPath.LoadPath[i]})
    }
    adminLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range adminLoadPath.StandbyLoadPath {
        types.SetYListKey(adminLoadPath.StandbyLoadPath[i], i)
        adminLoadPath.EntityData.Children.Append(types.GetSegmentPath(adminLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", adminLoadPath.StandbyLoadPath[i]})
    }
    adminLoadPath.EntityData.Leafs = types.NewOrderedMap()
    adminLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", adminLoadPath.RequestId})

    adminLoadPath.EntityData.YListKeys = []string {}

    return &(adminLoadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath
// Admin Resources load path
type Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "admin-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/admin-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/admin-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "admin-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/admin-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/admin-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active_Summary_SdrLoadPath
// SDR load paths
type Install_SoftwareInventory_Active_Summary_SdrLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath
}

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetEntityData() *types.CommonEntityData {
    sdrLoadPath.EntityData.YFilter = sdrLoadPath.YFilter
    sdrLoadPath.EntityData.YangName = "sdr-load-path"
    sdrLoadPath.EntityData.BundleName = "cisco_ios_xr"
    sdrLoadPath.EntityData.ParentYangName = "summary"
    sdrLoadPath.EntityData.SegmentPath = "sdr-load-path" + types.AddNoKeyToken(sdrLoadPath)
    sdrLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/" + sdrLoadPath.EntityData.SegmentPath
    sdrLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrLoadPath.EntityData.Children = types.NewOrderedMap()
    sdrLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range sdrLoadPath.LoadPath {
        types.SetYListKey(sdrLoadPath.LoadPath[i], i)
        sdrLoadPath.EntityData.Children.Append(types.GetSegmentPath(sdrLoadPath.LoadPath[i]), types.YChild{"LoadPath", sdrLoadPath.LoadPath[i]})
    }
    sdrLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range sdrLoadPath.StandbyLoadPath {
        types.SetYListKey(sdrLoadPath.StandbyLoadPath[i], i)
        sdrLoadPath.EntityData.Children.Append(types.GetSegmentPath(sdrLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", sdrLoadPath.StandbyLoadPath[i]})
    }
    sdrLoadPath.EntityData.Leafs = types.NewOrderedMap()
    sdrLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", sdrLoadPath.RequestId})
    sdrLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", sdrLoadPath.SecureDomainRouterName})

    sdrLoadPath.EntityData.YListKeys = []string {}

    return &(sdrLoadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "sdr-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/sdr-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/sdr-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "sdr-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/sdr-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/sdr-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active_Summary_LocationLoadPath
// Location load paths
type Install_SoftwareInventory_Active_Summary_LocationLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Node identifier. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath.
    LoadPath []*Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath.
    StandbyLoadPath []*Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath
}

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetEntityData() *types.CommonEntityData {
    locationLoadPath.EntityData.YFilter = locationLoadPath.YFilter
    locationLoadPath.EntityData.YangName = "location-load-path"
    locationLoadPath.EntityData.BundleName = "cisco_ios_xr"
    locationLoadPath.EntityData.ParentYangName = "summary"
    locationLoadPath.EntityData.SegmentPath = "location-load-path" + types.AddNoKeyToken(locationLoadPath)
    locationLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/" + locationLoadPath.EntityData.SegmentPath
    locationLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationLoadPath.EntityData.Children = types.NewOrderedMap()
    locationLoadPath.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range locationLoadPath.LoadPath {
        types.SetYListKey(locationLoadPath.LoadPath[i], i)
        locationLoadPath.EntityData.Children.Append(types.GetSegmentPath(locationLoadPath.LoadPath[i]), types.YChild{"LoadPath", locationLoadPath.LoadPath[i]})
    }
    locationLoadPath.EntityData.Children.Append("standby-load-path", types.YChild{"StandbyLoadPath", nil})
    for i := range locationLoadPath.StandbyLoadPath {
        types.SetYListKey(locationLoadPath.StandbyLoadPath[i], i)
        locationLoadPath.EntityData.Children.Append(types.GetSegmentPath(locationLoadPath.StandbyLoadPath[i]), types.YChild{"StandbyLoadPath", locationLoadPath.StandbyLoadPath[i]})
    }
    locationLoadPath.EntityData.Leafs = types.NewOrderedMap()
    locationLoadPath.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", locationLoadPath.RequestId})
    locationLoadPath.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", locationLoadPath.SecureDomainRouterName})
    locationLoadPath.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", locationLoadPath.NodeName})

    locationLoadPath.EntityData.YListKeys = []string {}

    return &(locationLoadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "location-load-path"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/location-load-path/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/location-load-path/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetEntityData() *types.CommonEntityData {
    standbyLoadPath.EntityData.YFilter = standbyLoadPath.YFilter
    standbyLoadPath.EntityData.YangName = "standby-load-path"
    standbyLoadPath.EntityData.BundleName = "cisco_ios_xr"
    standbyLoadPath.EntityData.ParentYangName = "location-load-path"
    standbyLoadPath.EntityData.SegmentPath = "standby-load-path" + types.AddNoKeyToken(standbyLoadPath)
    standbyLoadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/location-load-path/" + standbyLoadPath.EntityData.SegmentPath
    standbyLoadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyLoadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyLoadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyLoadPath.EntityData.Children = types.NewOrderedMap()
    standbyLoadPath.EntityData.Children.Append("package", types.YChild{"Package", &standbyLoadPath.Package})
    standbyLoadPath.EntityData.Leafs = types.NewOrderedMap()
    standbyLoadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", standbyLoadPath.Version})
    standbyLoadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", standbyLoadPath.BuildInformation})

    standbyLoadPath.EntityData.YListKeys = []string {}

    return &(standbyLoadPath.EntityData)
}

// Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "standby-load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/summary/location-load-path/standby-load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_SoftwareInventory_Active_Inventories
// Software inventory
type Install_SoftwareInventory_Active_Inventories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory information for specific node. The type is slice of
    // Install_SoftwareInventory_Active_Inventories_Inventory.
    Inventory []*Install_SoftwareInventory_Active_Inventories_Inventory
}

func (inventories *Install_SoftwareInventory_Active_Inventories) GetEntityData() *types.CommonEntityData {
    inventories.EntityData.YFilter = inventories.YFilter
    inventories.EntityData.YangName = "inventories"
    inventories.EntityData.BundleName = "cisco_ios_xr"
    inventories.EntityData.ParentYangName = "active"
    inventories.EntityData.SegmentPath = "inventories"
    inventories.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/" + inventories.EntityData.SegmentPath
    inventories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventories.EntityData.Children = types.NewOrderedMap()
    inventories.EntityData.Children.Append("inventory", types.YChild{"Inventory", nil})
    for i := range inventories.Inventory {
        inventories.EntityData.Children.Append(types.GetSegmentPath(inventories.Inventory[i]), types.YChild{"Inventory", inventories.Inventory[i]})
    }
    inventories.EntityData.Leafs = types.NewOrderedMap()

    inventories.EntityData.YListKeys = []string {}

    return &(inventories.EntityData)
}

// Install_SoftwareInventory_Active_Inventories_Inventory
// Inventory information for specific node
type Install_SoftwareInventory_Active_Inventories_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Major data version number. The type is interface{} with range:
    // 0..4294967295.
    Major interface{}

    // Minor data version number. The type is interface{} with range:
    // 0..4294967295.
    Minor interface{}

    // Name of the boot image. The type is string.
    BootImageName interface{}

    // Node's type. The type is interface{} with range: 0..18446744073709551615.
    NodeType interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath.
    LoadPath []*Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath
}

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "inventories"
    inventory.EntityData.SegmentPath = "inventory" + types.AddKeyToken(inventory.NodeName, "node-name")
    inventory.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/inventories/" + inventory.EntityData.SegmentPath
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("load-path", types.YChild{"LoadPath", nil})
    for i := range inventory.LoadPath {
        types.SetYListKey(inventory.LoadPath[i], i)
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.LoadPath[i]), types.YChild{"LoadPath", inventory.LoadPath[i]})
    }
    inventory.EntityData.Leafs = types.NewOrderedMap()
    inventory.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", inventory.NodeName})
    inventory.EntityData.Leafs.Append("major", types.YLeaf{"Major", inventory.Major})
    inventory.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", inventory.Minor})
    inventory.EntityData.Leafs.Append("boot-image-name", types.YLeaf{"BootImageName", inventory.BootImageName})
    inventory.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", inventory.NodeType})
    inventory.EntityData.Leafs.Append("secure-domain-router-name", types.YLeaf{"SecureDomainRouterName", inventory.SecureDomainRouterName})

    inventory.EntityData.YListKeys = []string {"NodeName"}

    return &(inventory.EntityData)
}

// Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath
// Load path
type Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetEntityData() *types.CommonEntityData {
    loadPath.EntityData.YFilter = loadPath.YFilter
    loadPath.EntityData.YangName = "load-path"
    loadPath.EntityData.BundleName = "cisco_ios_xr"
    loadPath.EntityData.ParentYangName = "inventory"
    loadPath.EntityData.SegmentPath = "load-path" + types.AddNoKeyToken(loadPath)
    loadPath.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/inventories/inventory/" + loadPath.EntityData.SegmentPath
    loadPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadPath.EntityData.Children = types.NewOrderedMap()
    loadPath.EntityData.Children.Append("package", types.YChild{"Package", &loadPath.Package})
    loadPath.EntityData.Leafs = types.NewOrderedMap()
    loadPath.EntityData.Leafs.Append("version", types.YLeaf{"Version", loadPath.Version})
    loadPath.EntityData.Leafs.Append("build-information", types.YLeaf{"BuildInformation", loadPath.BuildInformation})

    loadPath.EntityData.YListKeys = []string {}

    return &(loadPath.EntityData)
}

// Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "load-path"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/software-inventory/active/inventories/inventory/load-path/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", self.DeviceName})
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_Issu
// Information of install ISSU operations
type Install_Issu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISSU manager card inventory table.
    CardInventories Install_Issu_CardInventories

    // Summarized ISSU stage information.
    Stage Install_Issu_Stage
}

func (issu *Install_Issu) GetEntityData() *types.CommonEntityData {
    issu.EntityData.YFilter = issu.YFilter
    issu.EntityData.YangName = "issu"
    issu.EntityData.BundleName = "cisco_ios_xr"
    issu.EntityData.ParentYangName = "install"
    issu.EntityData.SegmentPath = "issu"
    issu.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/" + issu.EntityData.SegmentPath
    issu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issu.EntityData.Children = types.NewOrderedMap()
    issu.EntityData.Children.Append("card-inventories", types.YChild{"CardInventories", &issu.CardInventories})
    issu.EntityData.Children.Append("stage", types.YChild{"Stage", &issu.Stage})
    issu.EntityData.Leafs = types.NewOrderedMap()

    issu.EntityData.YListKeys = []string {}

    return &(issu.EntityData)
}

// Install_Issu_CardInventories
// ISSU manager card inventory table
type Install_Issu_CardInventories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISSU manager inventory summary of the same card type. The type is slice of
    // Install_Issu_CardInventories_CardInventory.
    CardInventory []*Install_Issu_CardInventories_CardInventory
}

func (cardInventories *Install_Issu_CardInventories) GetEntityData() *types.CommonEntityData {
    cardInventories.EntityData.YFilter = cardInventories.YFilter
    cardInventories.EntityData.YangName = "card-inventories"
    cardInventories.EntityData.BundleName = "cisco_ios_xr"
    cardInventories.EntityData.ParentYangName = "issu"
    cardInventories.EntityData.SegmentPath = "card-inventories"
    cardInventories.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/issu/" + cardInventories.EntityData.SegmentPath
    cardInventories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardInventories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardInventories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardInventories.EntityData.Children = types.NewOrderedMap()
    cardInventories.EntityData.Children.Append("card-inventory", types.YChild{"CardInventory", nil})
    for i := range cardInventories.CardInventory {
        cardInventories.EntityData.Children.Append(types.GetSegmentPath(cardInventories.CardInventory[i]), types.YChild{"CardInventory", cardInventories.CardInventory[i]})
    }
    cardInventories.EntityData.Leafs = types.NewOrderedMap()

    cardInventories.EntityData.YListKeys = []string {}

    return &(cardInventories.EntityData)
}

// Install_Issu_CardInventories_CardInventory
// ISSU manager inventory summary of the same
// card type
type Install_Issu_CardInventories_CardInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. ISSU manager card type ID. The type is
    // IsmCardTypeFamily.
    CardTypeId interface{}

    // node state for all nodes. The type is slice of
    // Install_Issu_CardInventories_CardInventory_Summary.
    Summary []*Install_Issu_CardInventories_CardInventory_Summary
}

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetEntityData() *types.CommonEntityData {
    cardInventory.EntityData.YFilter = cardInventory.YFilter
    cardInventory.EntityData.YangName = "card-inventory"
    cardInventory.EntityData.BundleName = "cisco_ios_xr"
    cardInventory.EntityData.ParentYangName = "card-inventories"
    cardInventory.EntityData.SegmentPath = "card-inventory" + types.AddKeyToken(cardInventory.CardTypeId, "card-type-id")
    cardInventory.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/issu/card-inventories/" + cardInventory.EntityData.SegmentPath
    cardInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardInventory.EntityData.Children = types.NewOrderedMap()
    cardInventory.EntityData.Children.Append("summary", types.YChild{"Summary", nil})
    for i := range cardInventory.Summary {
        types.SetYListKey(cardInventory.Summary[i], i)
        cardInventory.EntityData.Children.Append(types.GetSegmentPath(cardInventory.Summary[i]), types.YChild{"Summary", cardInventory.Summary[i]})
    }
    cardInventory.EntityData.Leafs = types.NewOrderedMap()
    cardInventory.EntityData.Leafs.Append("card-type-id", types.YLeaf{"CardTypeId", cardInventory.CardTypeId})

    cardInventory.EntityData.YListKeys = []string {"CardTypeId"}

    return &(cardInventory.EntityData)
}

// Install_Issu_CardInventories_CardInventory_Summary
// node state for all nodes
type Install_Issu_CardInventories_CardInventory_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Node identifier. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Partner Node IDs. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PartnerNodeName interface{}

    // Node state. The type is InstmgrCardState.
    NodeState interface{}

    // Node roll. The type is InstmgrNodeRole.
    NodeRole interface{}

    // PI Node type. The type is InstmgrPiCard.
    NodeTypePi interface{}

    // ISSU node type. The type is string.
    NodeTypeIssu interface{}

    // Current node ISSU state. The type is InstmgrIsmNodeState.
    NodeCurrentState interface{}

    // Expected ISSU state. The type is InstmgrIsmNodeState.
    NodeExpectedState interface{}

    // Node failure reason. The type is string.
    NodeFailureReason interface{}

    // Node none-cnforming. The type is InstallmgrIsmNodeConforming.
    IsConformingNode interface{}

    // Number of attempts made. The type is interface{} with range: 0..4294967295.
    Attempts interface{}

    // Is node upgraded?. The type is bool.
    IsNodeUpgraded interface{}
}

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "card-inventory"
    summary.EntityData.SegmentPath = "summary" + types.AddNoKeyToken(summary)
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/issu/card-inventories/card-inventory/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", summary.NodeName})
    summary.EntityData.Leafs.Append("partner-node-name", types.YLeaf{"PartnerNodeName", summary.PartnerNodeName})
    summary.EntityData.Leafs.Append("node-state", types.YLeaf{"NodeState", summary.NodeState})
    summary.EntityData.Leafs.Append("node-role", types.YLeaf{"NodeRole", summary.NodeRole})
    summary.EntityData.Leafs.Append("node-type-pi", types.YLeaf{"NodeTypePi", summary.NodeTypePi})
    summary.EntityData.Leafs.Append("node-type-issu", types.YLeaf{"NodeTypeIssu", summary.NodeTypeIssu})
    summary.EntityData.Leafs.Append("node-current-state", types.YLeaf{"NodeCurrentState", summary.NodeCurrentState})
    summary.EntityData.Leafs.Append("node-expected-state", types.YLeaf{"NodeExpectedState", summary.NodeExpectedState})
    summary.EntityData.Leafs.Append("node-failure-reason", types.YLeaf{"NodeFailureReason", summary.NodeFailureReason})
    summary.EntityData.Leafs.Append("is-conforming-node", types.YLeaf{"IsConformingNode", summary.IsConformingNode})
    summary.EntityData.Leafs.Append("attempts", types.YLeaf{"Attempts", summary.Attempts})
    summary.EntityData.Leafs.Append("is-node-upgraded", types.YLeaf{"IsNodeUpgraded", summary.IsNodeUpgraded})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Install_Issu_Stage
// Summarized ISSU stage information
type Install_Issu_Stage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current ISSU state. The type is string.
    IssuState interface{}

    // ISSU operational ID. The type is interface{} with range: 0..4294967295.
    IssuOpId interface{}

    // ISSU progress percentage. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Percentage interface{}

    // ISSU aborted?. The type is bool.
    IsIssuAborted interface{}

    // ISSU aborted by ISM?. The type is bool.
    IsIssuAbortedByIsm interface{}

    // ISM FSM state. The type is InstmgrIsmFsmState.
    IssuManagerFsmState interface{}

    // Number of participating nodes. The type is interface{} with range:
    // 0..4294967295.
    ParticipatingNodeAll interface{}

    // Number of node in progress. The type is interface{} with range:
    // 0..4294967295.
    NumNodesInProgress interface{}

    // Number of nodes in LOAD phase. The type is interface{} with range:
    // 0..4294967295.
    NumOfNodesInLoad interface{}

    // Number of nodes in RUN phase. The type is interface{} with range:
    // 0..4294967295.
    NumOfNodesInRun interface{}

    // Number of none-conforming nodes. The type is interface{} with range:
    // 0..4294967295.
    NumofNcNodes interface{}

    // Nodes in progress.
    NodeInProgress Install_Issu_Stage_NodeInProgress

    // Node in LOAD phase.
    NodesInLoad Install_Issu_Stage_NodesInLoad

    // Node in RUN phase.
    NodesInRun Install_Issu_Stage_NodesInRun

    // None-conforming nodes.
    NcNodes Install_Issu_Stage_NcNodes
}

func (stage *Install_Issu_Stage) GetEntityData() *types.CommonEntityData {
    stage.EntityData.YFilter = stage.YFilter
    stage.EntityData.YangName = "stage"
    stage.EntityData.BundleName = "cisco_ios_xr"
    stage.EntityData.ParentYangName = "issu"
    stage.EntityData.SegmentPath = "stage"
    stage.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/issu/" + stage.EntityData.SegmentPath
    stage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stage.EntityData.Children = types.NewOrderedMap()
    stage.EntityData.Children.Append("node-in-progress", types.YChild{"NodeInProgress", &stage.NodeInProgress})
    stage.EntityData.Children.Append("nodes-in-load", types.YChild{"NodesInLoad", &stage.NodesInLoad})
    stage.EntityData.Children.Append("nodes-in-run", types.YChild{"NodesInRun", &stage.NodesInRun})
    stage.EntityData.Children.Append("nc-nodes", types.YChild{"NcNodes", &stage.NcNodes})
    stage.EntityData.Leafs = types.NewOrderedMap()
    stage.EntityData.Leafs.Append("issu-state", types.YLeaf{"IssuState", stage.IssuState})
    stage.EntityData.Leafs.Append("issu-op-id", types.YLeaf{"IssuOpId", stage.IssuOpId})
    stage.EntityData.Leafs.Append("percentage", types.YLeaf{"Percentage", stage.Percentage})
    stage.EntityData.Leafs.Append("is-issu-aborted", types.YLeaf{"IsIssuAborted", stage.IsIssuAborted})
    stage.EntityData.Leafs.Append("is-issu-aborted-by-ism", types.YLeaf{"IsIssuAbortedByIsm", stage.IsIssuAbortedByIsm})
    stage.EntityData.Leafs.Append("issu-manager-fsm-state", types.YLeaf{"IssuManagerFsmState", stage.IssuManagerFsmState})
    stage.EntityData.Leafs.Append("participating-node-all", types.YLeaf{"ParticipatingNodeAll", stage.ParticipatingNodeAll})
    stage.EntityData.Leafs.Append("num-nodes-in-progress", types.YLeaf{"NumNodesInProgress", stage.NumNodesInProgress})
    stage.EntityData.Leafs.Append("num-of-nodes-in-load", types.YLeaf{"NumOfNodesInLoad", stage.NumOfNodesInLoad})
    stage.EntityData.Leafs.Append("num-of-nodes-in-run", types.YLeaf{"NumOfNodesInRun", stage.NumOfNodesInRun})
    stage.EntityData.Leafs.Append("numof-nc-nodes", types.YLeaf{"NumofNcNodes", stage.NumofNcNodes})

    stage.EntityData.YListKeys = []string {}

    return &(stage.EntityData)
}

// Install_Issu_Stage_NodeInProgress
// Nodes in progress
type Install_Issu_Stage_NodeInProgress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // node. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node []interface{}
}

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetEntityData() *types.CommonEntityData {
    nodeInProgress.EntityData.YFilter = nodeInProgress.YFilter
    nodeInProgress.EntityData.YangName = "node-in-progress"
    nodeInProgress.EntityData.BundleName = "cisco_ios_xr"
    nodeInProgress.EntityData.ParentYangName = "stage"
    nodeInProgress.EntityData.SegmentPath = "node-in-progress"
    nodeInProgress.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/issu/stage/" + nodeInProgress.EntityData.SegmentPath
    nodeInProgress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInProgress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInProgress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInProgress.EntityData.Children = types.NewOrderedMap()
    nodeInProgress.EntityData.Leafs = types.NewOrderedMap()
    nodeInProgress.EntityData.Leafs.Append("node", types.YLeaf{"Node", nodeInProgress.Node})

    nodeInProgress.EntityData.YListKeys = []string {}

    return &(nodeInProgress.EntityData)
}

// Install_Issu_Stage_NodesInLoad
// Node in LOAD phase
type Install_Issu_Stage_NodesInLoad struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // node. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node []interface{}
}

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetEntityData() *types.CommonEntityData {
    nodesInLoad.EntityData.YFilter = nodesInLoad.YFilter
    nodesInLoad.EntityData.YangName = "nodes-in-load"
    nodesInLoad.EntityData.BundleName = "cisco_ios_xr"
    nodesInLoad.EntityData.ParentYangName = "stage"
    nodesInLoad.EntityData.SegmentPath = "nodes-in-load"
    nodesInLoad.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/issu/stage/" + nodesInLoad.EntityData.SegmentPath
    nodesInLoad.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodesInLoad.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodesInLoad.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodesInLoad.EntityData.Children = types.NewOrderedMap()
    nodesInLoad.EntityData.Leafs = types.NewOrderedMap()
    nodesInLoad.EntityData.Leafs.Append("node", types.YLeaf{"Node", nodesInLoad.Node})

    nodesInLoad.EntityData.YListKeys = []string {}

    return &(nodesInLoad.EntityData)
}

// Install_Issu_Stage_NodesInRun
// Node in RUN phase
type Install_Issu_Stage_NodesInRun struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // node. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node []interface{}
}

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetEntityData() *types.CommonEntityData {
    nodesInRun.EntityData.YFilter = nodesInRun.YFilter
    nodesInRun.EntityData.YangName = "nodes-in-run"
    nodesInRun.EntityData.BundleName = "cisco_ios_xr"
    nodesInRun.EntityData.ParentYangName = "stage"
    nodesInRun.EntityData.SegmentPath = "nodes-in-run"
    nodesInRun.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/issu/stage/" + nodesInRun.EntityData.SegmentPath
    nodesInRun.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodesInRun.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodesInRun.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodesInRun.EntityData.Children = types.NewOrderedMap()
    nodesInRun.EntityData.Leafs = types.NewOrderedMap()
    nodesInRun.EntityData.Leafs.Append("node", types.YLeaf{"Node", nodesInRun.Node})

    nodesInRun.EntityData.YListKeys = []string {}

    return &(nodesInRun.EntityData)
}

// Install_Issu_Stage_NcNodes
// None-conforming nodes
type Install_Issu_Stage_NcNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // node. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node []interface{}
}

func (ncNodes *Install_Issu_Stage_NcNodes) GetEntityData() *types.CommonEntityData {
    ncNodes.EntityData.YFilter = ncNodes.YFilter
    ncNodes.EntityData.YangName = "nc-nodes"
    ncNodes.EntityData.BundleName = "cisco_ios_xr"
    ncNodes.EntityData.ParentYangName = "stage"
    ncNodes.EntityData.SegmentPath = "nc-nodes"
    ncNodes.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/issu/stage/" + ncNodes.EntityData.SegmentPath
    ncNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncNodes.EntityData.Children = types.NewOrderedMap()
    ncNodes.EntityData.Leafs = types.NewOrderedMap()
    ncNodes.EntityData.Leafs.Append("node", types.YLeaf{"Node", ncNodes.Node})

    ncNodes.EntityData.YListKeys = []string {}

    return &(ncNodes.EntityData)
}

// Install_BootImage
// System Boot Image
type Install_BootImage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The boot image. The type is string.
    SystemImageFile interface{}
}

func (bootImage *Install_BootImage) GetEntityData() *types.CommonEntityData {
    bootImage.EntityData.YFilter = bootImage.YFilter
    bootImage.EntityData.YangName = "boot-image"
    bootImage.EntityData.BundleName = "cisco_ios_xr"
    bootImage.EntityData.ParentYangName = "install"
    bootImage.EntityData.SegmentPath = "boot-image"
    bootImage.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/" + bootImage.EntityData.SegmentPath
    bootImage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootImage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootImage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootImage.EntityData.Children = types.NewOrderedMap()
    bootImage.EntityData.Leafs = types.NewOrderedMap()
    bootImage.EntityData.Leafs.Append("system-image-file", types.YLeaf{"SystemImageFile", bootImage.SystemImageFile})

    bootImage.EntityData.YListKeys = []string {}

    return &(bootImage.EntityData)
}

// Install_Logs
// Install operation log
type Install_Logs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log information. The type is slice of Install_Logs_Log.
    Log []*Install_Logs_Log
}

func (logs *Install_Logs) GetEntityData() *types.CommonEntityData {
    logs.EntityData.YFilter = logs.YFilter
    logs.EntityData.YangName = "logs"
    logs.EntityData.BundleName = "cisco_ios_xr"
    logs.EntityData.ParentYangName = "install"
    logs.EntityData.SegmentPath = "logs"
    logs.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/" + logs.EntityData.SegmentPath
    logs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logs.EntityData.Children = types.NewOrderedMap()
    logs.EntityData.Children.Append("log", types.YChild{"Log", nil})
    for i := range logs.Log {
        logs.EntityData.Children.Append(types.GetSegmentPath(logs.Log[i]), types.YChild{"Log", logs.Log[i]})
    }
    logs.EntityData.Leafs = types.NewOrderedMap()

    logs.EntityData.YListKeys = []string {}

    return &(logs.EntityData)
}

// Install_Logs_Log
// Log information
type Install_Logs_Log struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Install operation request ID. The type is
    // interface{} with range: 0..4294967295.
    RequestId interface{}

    // Header information. The type is slice of Install_Logs_Log_Header.
    Header []*Install_Logs_Log_Header

    // Summary information. The type is slice of Install_Logs_Log_Summary.
    Summary []*Install_Logs_Log_Summary

    // Status Information Logs. The type is slice of Install_Logs_Log_Message.
    Message []*Install_Logs_Log_Message

    // Install changes. The type is slice of Install_Logs_Log_Change.
    Change []*Install_Logs_Log_Change

    // Install details. The type is slice of Install_Logs_Log_Detail.
    Detail []*Install_Logs_Log_Detail

    // Install communications. The type is slice of
    // Install_Logs_Log_Communication.
    Communication []*Install_Logs_Log_Communication
}

func (log *Install_Logs_Log) GetEntityData() *types.CommonEntityData {
    log.EntityData.YFilter = log.YFilter
    log.EntityData.YangName = "log"
    log.EntityData.BundleName = "cisco_ios_xr"
    log.EntityData.ParentYangName = "logs"
    log.EntityData.SegmentPath = "log" + types.AddKeyToken(log.RequestId, "request-id")
    log.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/" + log.EntityData.SegmentPath
    log.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    log.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    log.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    log.EntityData.Children = types.NewOrderedMap()
    log.EntityData.Children.Append("header", types.YChild{"Header", nil})
    for i := range log.Header {
        types.SetYListKey(log.Header[i], i)
        log.EntityData.Children.Append(types.GetSegmentPath(log.Header[i]), types.YChild{"Header", log.Header[i]})
    }
    log.EntityData.Children.Append("summary", types.YChild{"Summary", nil})
    for i := range log.Summary {
        types.SetYListKey(log.Summary[i], i)
        log.EntityData.Children.Append(types.GetSegmentPath(log.Summary[i]), types.YChild{"Summary", log.Summary[i]})
    }
    log.EntityData.Children.Append("message", types.YChild{"Message", nil})
    for i := range log.Message {
        types.SetYListKey(log.Message[i], i)
        log.EntityData.Children.Append(types.GetSegmentPath(log.Message[i]), types.YChild{"Message", log.Message[i]})
    }
    log.EntityData.Children.Append("change", types.YChild{"Change", nil})
    for i := range log.Change {
        types.SetYListKey(log.Change[i], i)
        log.EntityData.Children.Append(types.GetSegmentPath(log.Change[i]), types.YChild{"Change", log.Change[i]})
    }
    log.EntityData.Children.Append("detail", types.YChild{"Detail", nil})
    for i := range log.Detail {
        types.SetYListKey(log.Detail[i], i)
        log.EntityData.Children.Append(types.GetSegmentPath(log.Detail[i]), types.YChild{"Detail", log.Detail[i]})
    }
    log.EntityData.Children.Append("communication", types.YChild{"Communication", nil})
    for i := range log.Communication {
        types.SetYListKey(log.Communication[i], i)
        log.EntityData.Children.Append(types.GetSegmentPath(log.Communication[i]), types.YChild{"Communication", log.Communication[i]})
    }
    log.EntityData.Leafs = types.NewOrderedMap()
    log.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", log.RequestId})

    log.EntityData.YListKeys = []string {"RequestId"}

    return &(log.EntityData)
}

// Install_Logs_Log_Header
// Header information
type Install_Logs_Log_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Log contents.
    LogContents Install_Logs_Log_Header_LogContents
}

func (header *Install_Logs_Log_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "log"
    header.EntityData.SegmentPath = "header" + types.AddNoKeyToken(header)
    header.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/" + header.EntityData.SegmentPath
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = types.NewOrderedMap()
    header.EntityData.Children.Append("log-contents", types.YChild{"LogContents", &header.LogContents})
    header.EntityData.Leafs = types.NewOrderedMap()

    header.EntityData.YListKeys = []string {}

    return &(header.EntityData)
}

// Install_Logs_Log_Header_LogContents
// Log contents
type Install_Logs_Log_Header_LogContents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Header_LogContents_V3
}

func (logContents *Install_Logs_Log_Header_LogContents) GetEntityData() *types.CommonEntityData {
    logContents.EntityData.YFilter = logContents.YFilter
    logContents.EntityData.YangName = "log-contents"
    logContents.EntityData.BundleName = "cisco_ios_xr"
    logContents.EntityData.ParentYangName = "header"
    logContents.EntityData.SegmentPath = "log-contents"
    logContents.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/header/" + logContents.EntityData.SegmentPath
    logContents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logContents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logContents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logContents.EntityData.Children = types.NewOrderedMap()
    logContents.EntityData.Children.Append("v3", types.YChild{"V3", &logContents.V3})
    logContents.EntityData.Leafs = types.NewOrderedMap()
    logContents.EntityData.Leafs.Append("version", types.YLeaf{"Version", logContents.Version})

    logContents.EntityData.YListKeys = []string {}

    return &(logContents.EntityData)
}

// Install_Logs_Log_Header_LogContents_V3
// v3
type Install_Logs_Log_Header_LogContents_V3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Header_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetEntityData() *types.CommonEntityData {
    v3.EntityData.YFilter = v3.YFilter
    v3.EntityData.YangName = "v3"
    v3.EntityData.BundleName = "cisco_ios_xr"
    v3.EntityData.ParentYangName = "log-contents"
    v3.EntityData.SegmentPath = "v3"
    v3.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/header/log-contents/" + v3.EntityData.SegmentPath
    v3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v3.EntityData.Children = types.NewOrderedMap()
    v3.EntityData.Children.Append("scope", types.YChild{"Scope", &v3.Scope})
    v3.EntityData.Leafs = types.NewOrderedMap()
    v3.EntityData.Leafs.Append("category", types.YLeaf{"Category", v3.Category})
    v3.EntityData.Leafs.Append("message", types.YLeaf{"Message", v3.Message})

    v3.EntityData.YListKeys = []string {}

    return &(v3.EntityData)
}

// Install_Logs_Log_Header_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Header_LogContents_V3_Scope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetEntityData() *types.CommonEntityData {
    scope.EntityData.YFilter = scope.YFilter
    scope.EntityData.YangName = "scope"
    scope.EntityData.BundleName = "cisco_ios_xr"
    scope.EntityData.ParentYangName = "v3"
    scope.EntityData.SegmentPath = "scope"
    scope.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/header/log-contents/v3/" + scope.EntityData.SegmentPath
    scope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scope.EntityData.Children = types.NewOrderedMap()
    scope.EntityData.Leafs = types.NewOrderedMap()
    scope.EntityData.Leafs.Append("admin-read", types.YLeaf{"AdminRead", scope.AdminRead})
    scope.EntityData.Leafs.Append("affected-sd-rs", types.YLeaf{"AffectedSdRs", scope.AffectedSdRs})

    scope.EntityData.YListKeys = []string {}

    return &(scope.EntityData)
}

// Install_Logs_Log_Summary
// Summary information
type Install_Logs_Log_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Log contents.
    LogContents Install_Logs_Log_Summary_LogContents
}

func (summary *Install_Logs_Log_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "log"
    summary.EntityData.SegmentPath = "summary" + types.AddNoKeyToken(summary)
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("log-contents", types.YChild{"LogContents", &summary.LogContents})
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Install_Logs_Log_Summary_LogContents
// Log contents
type Install_Logs_Log_Summary_LogContents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Summary_LogContents_V3
}

func (logContents *Install_Logs_Log_Summary_LogContents) GetEntityData() *types.CommonEntityData {
    logContents.EntityData.YFilter = logContents.YFilter
    logContents.EntityData.YangName = "log-contents"
    logContents.EntityData.BundleName = "cisco_ios_xr"
    logContents.EntityData.ParentYangName = "summary"
    logContents.EntityData.SegmentPath = "log-contents"
    logContents.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/summary/" + logContents.EntityData.SegmentPath
    logContents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logContents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logContents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logContents.EntityData.Children = types.NewOrderedMap()
    logContents.EntityData.Children.Append("v3", types.YChild{"V3", &logContents.V3})
    logContents.EntityData.Leafs = types.NewOrderedMap()
    logContents.EntityData.Leafs.Append("version", types.YLeaf{"Version", logContents.Version})

    logContents.EntityData.YListKeys = []string {}

    return &(logContents.EntityData)
}

// Install_Logs_Log_Summary_LogContents_V3
// v3
type Install_Logs_Log_Summary_LogContents_V3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Summary_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetEntityData() *types.CommonEntityData {
    v3.EntityData.YFilter = v3.YFilter
    v3.EntityData.YangName = "v3"
    v3.EntityData.BundleName = "cisco_ios_xr"
    v3.EntityData.ParentYangName = "log-contents"
    v3.EntityData.SegmentPath = "v3"
    v3.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/summary/log-contents/" + v3.EntityData.SegmentPath
    v3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v3.EntityData.Children = types.NewOrderedMap()
    v3.EntityData.Children.Append("scope", types.YChild{"Scope", &v3.Scope})
    v3.EntityData.Leafs = types.NewOrderedMap()
    v3.EntityData.Leafs.Append("category", types.YLeaf{"Category", v3.Category})
    v3.EntityData.Leafs.Append("message", types.YLeaf{"Message", v3.Message})

    v3.EntityData.YListKeys = []string {}

    return &(v3.EntityData)
}

// Install_Logs_Log_Summary_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Summary_LogContents_V3_Scope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetEntityData() *types.CommonEntityData {
    scope.EntityData.YFilter = scope.YFilter
    scope.EntityData.YangName = "scope"
    scope.EntityData.BundleName = "cisco_ios_xr"
    scope.EntityData.ParentYangName = "v3"
    scope.EntityData.SegmentPath = "scope"
    scope.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/summary/log-contents/v3/" + scope.EntityData.SegmentPath
    scope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scope.EntityData.Children = types.NewOrderedMap()
    scope.EntityData.Leafs = types.NewOrderedMap()
    scope.EntityData.Leafs.Append("admin-read", types.YLeaf{"AdminRead", scope.AdminRead})
    scope.EntityData.Leafs.Append("affected-sd-rs", types.YLeaf{"AffectedSdRs", scope.AffectedSdRs})

    scope.EntityData.YListKeys = []string {}

    return &(scope.EntityData)
}

// Install_Logs_Log_Message
// Status Information Logs
type Install_Logs_Log_Message struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Log contents.
    LogContents Install_Logs_Log_Message_LogContents
}

func (message *Install_Logs_Log_Message) GetEntityData() *types.CommonEntityData {
    message.EntityData.YFilter = message.YFilter
    message.EntityData.YangName = "message"
    message.EntityData.BundleName = "cisco_ios_xr"
    message.EntityData.ParentYangName = "log"
    message.EntityData.SegmentPath = "message" + types.AddNoKeyToken(message)
    message.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/" + message.EntityData.SegmentPath
    message.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    message.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    message.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    message.EntityData.Children = types.NewOrderedMap()
    message.EntityData.Children.Append("log-contents", types.YChild{"LogContents", &message.LogContents})
    message.EntityData.Leafs = types.NewOrderedMap()

    message.EntityData.YListKeys = []string {}

    return &(message.EntityData)
}

// Install_Logs_Log_Message_LogContents
// Log contents
type Install_Logs_Log_Message_LogContents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Message_LogContents_V3
}

func (logContents *Install_Logs_Log_Message_LogContents) GetEntityData() *types.CommonEntityData {
    logContents.EntityData.YFilter = logContents.YFilter
    logContents.EntityData.YangName = "log-contents"
    logContents.EntityData.BundleName = "cisco_ios_xr"
    logContents.EntityData.ParentYangName = "message"
    logContents.EntityData.SegmentPath = "log-contents"
    logContents.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/message/" + logContents.EntityData.SegmentPath
    logContents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logContents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logContents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logContents.EntityData.Children = types.NewOrderedMap()
    logContents.EntityData.Children.Append("v3", types.YChild{"V3", &logContents.V3})
    logContents.EntityData.Leafs = types.NewOrderedMap()
    logContents.EntityData.Leafs.Append("version", types.YLeaf{"Version", logContents.Version})

    logContents.EntityData.YListKeys = []string {}

    return &(logContents.EntityData)
}

// Install_Logs_Log_Message_LogContents_V3
// v3
type Install_Logs_Log_Message_LogContents_V3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Message_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetEntityData() *types.CommonEntityData {
    v3.EntityData.YFilter = v3.YFilter
    v3.EntityData.YangName = "v3"
    v3.EntityData.BundleName = "cisco_ios_xr"
    v3.EntityData.ParentYangName = "log-contents"
    v3.EntityData.SegmentPath = "v3"
    v3.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/message/log-contents/" + v3.EntityData.SegmentPath
    v3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v3.EntityData.Children = types.NewOrderedMap()
    v3.EntityData.Children.Append("scope", types.YChild{"Scope", &v3.Scope})
    v3.EntityData.Leafs = types.NewOrderedMap()
    v3.EntityData.Leafs.Append("category", types.YLeaf{"Category", v3.Category})
    v3.EntityData.Leafs.Append("message", types.YLeaf{"Message", v3.Message})

    v3.EntityData.YListKeys = []string {}

    return &(v3.EntityData)
}

// Install_Logs_Log_Message_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Message_LogContents_V3_Scope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetEntityData() *types.CommonEntityData {
    scope.EntityData.YFilter = scope.YFilter
    scope.EntityData.YangName = "scope"
    scope.EntityData.BundleName = "cisco_ios_xr"
    scope.EntityData.ParentYangName = "v3"
    scope.EntityData.SegmentPath = "scope"
    scope.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/message/log-contents/v3/" + scope.EntityData.SegmentPath
    scope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scope.EntityData.Children = types.NewOrderedMap()
    scope.EntityData.Leafs = types.NewOrderedMap()
    scope.EntityData.Leafs.Append("admin-read", types.YLeaf{"AdminRead", scope.AdminRead})
    scope.EntityData.Leafs.Append("affected-sd-rs", types.YLeaf{"AffectedSdRs", scope.AffectedSdRs})

    scope.EntityData.YListKeys = []string {}

    return &(scope.EntityData)
}

// Install_Logs_Log_Change
// Install changes
type Install_Logs_Log_Change struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Log contents.
    LogContents Install_Logs_Log_Change_LogContents
}

func (change *Install_Logs_Log_Change) GetEntityData() *types.CommonEntityData {
    change.EntityData.YFilter = change.YFilter
    change.EntityData.YangName = "change"
    change.EntityData.BundleName = "cisco_ios_xr"
    change.EntityData.ParentYangName = "log"
    change.EntityData.SegmentPath = "change" + types.AddNoKeyToken(change)
    change.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/" + change.EntityData.SegmentPath
    change.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    change.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    change.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    change.EntityData.Children = types.NewOrderedMap()
    change.EntityData.Children.Append("log-contents", types.YChild{"LogContents", &change.LogContents})
    change.EntityData.Leafs = types.NewOrderedMap()

    change.EntityData.YListKeys = []string {}

    return &(change.EntityData)
}

// Install_Logs_Log_Change_LogContents
// Log contents
type Install_Logs_Log_Change_LogContents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Change_LogContents_V3
}

func (logContents *Install_Logs_Log_Change_LogContents) GetEntityData() *types.CommonEntityData {
    logContents.EntityData.YFilter = logContents.YFilter
    logContents.EntityData.YangName = "log-contents"
    logContents.EntityData.BundleName = "cisco_ios_xr"
    logContents.EntityData.ParentYangName = "change"
    logContents.EntityData.SegmentPath = "log-contents"
    logContents.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/change/" + logContents.EntityData.SegmentPath
    logContents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logContents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logContents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logContents.EntityData.Children = types.NewOrderedMap()
    logContents.EntityData.Children.Append("v3", types.YChild{"V3", &logContents.V3})
    logContents.EntityData.Leafs = types.NewOrderedMap()
    logContents.EntityData.Leafs.Append("version", types.YLeaf{"Version", logContents.Version})

    logContents.EntityData.YListKeys = []string {}

    return &(logContents.EntityData)
}

// Install_Logs_Log_Change_LogContents_V3
// v3
type Install_Logs_Log_Change_LogContents_V3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Change_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetEntityData() *types.CommonEntityData {
    v3.EntityData.YFilter = v3.YFilter
    v3.EntityData.YangName = "v3"
    v3.EntityData.BundleName = "cisco_ios_xr"
    v3.EntityData.ParentYangName = "log-contents"
    v3.EntityData.SegmentPath = "v3"
    v3.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/change/log-contents/" + v3.EntityData.SegmentPath
    v3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v3.EntityData.Children = types.NewOrderedMap()
    v3.EntityData.Children.Append("scope", types.YChild{"Scope", &v3.Scope})
    v3.EntityData.Leafs = types.NewOrderedMap()
    v3.EntityData.Leafs.Append("category", types.YLeaf{"Category", v3.Category})
    v3.EntityData.Leafs.Append("message", types.YLeaf{"Message", v3.Message})

    v3.EntityData.YListKeys = []string {}

    return &(v3.EntityData)
}

// Install_Logs_Log_Change_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Change_LogContents_V3_Scope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetEntityData() *types.CommonEntityData {
    scope.EntityData.YFilter = scope.YFilter
    scope.EntityData.YangName = "scope"
    scope.EntityData.BundleName = "cisco_ios_xr"
    scope.EntityData.ParentYangName = "v3"
    scope.EntityData.SegmentPath = "scope"
    scope.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/change/log-contents/v3/" + scope.EntityData.SegmentPath
    scope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scope.EntityData.Children = types.NewOrderedMap()
    scope.EntityData.Leafs = types.NewOrderedMap()
    scope.EntityData.Leafs.Append("admin-read", types.YLeaf{"AdminRead", scope.AdminRead})
    scope.EntityData.Leafs.Append("affected-sd-rs", types.YLeaf{"AffectedSdRs", scope.AffectedSdRs})

    scope.EntityData.YListKeys = []string {}

    return &(scope.EntityData)
}

// Install_Logs_Log_Detail
// Install details
type Install_Logs_Log_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Log contents.
    LogContents Install_Logs_Log_Detail_LogContents
}

func (detail *Install_Logs_Log_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "log"
    detail.EntityData.SegmentPath = "detail" + types.AddNoKeyToken(detail)
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("log-contents", types.YChild{"LogContents", &detail.LogContents})
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Install_Logs_Log_Detail_LogContents
// Log contents
type Install_Logs_Log_Detail_LogContents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Detail_LogContents_V3
}

func (logContents *Install_Logs_Log_Detail_LogContents) GetEntityData() *types.CommonEntityData {
    logContents.EntityData.YFilter = logContents.YFilter
    logContents.EntityData.YangName = "log-contents"
    logContents.EntityData.BundleName = "cisco_ios_xr"
    logContents.EntityData.ParentYangName = "detail"
    logContents.EntityData.SegmentPath = "log-contents"
    logContents.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/detail/" + logContents.EntityData.SegmentPath
    logContents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logContents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logContents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logContents.EntityData.Children = types.NewOrderedMap()
    logContents.EntityData.Children.Append("v3", types.YChild{"V3", &logContents.V3})
    logContents.EntityData.Leafs = types.NewOrderedMap()
    logContents.EntityData.Leafs.Append("version", types.YLeaf{"Version", logContents.Version})

    logContents.EntityData.YListKeys = []string {}

    return &(logContents.EntityData)
}

// Install_Logs_Log_Detail_LogContents_V3
// v3
type Install_Logs_Log_Detail_LogContents_V3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Detail_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetEntityData() *types.CommonEntityData {
    v3.EntityData.YFilter = v3.YFilter
    v3.EntityData.YangName = "v3"
    v3.EntityData.BundleName = "cisco_ios_xr"
    v3.EntityData.ParentYangName = "log-contents"
    v3.EntityData.SegmentPath = "v3"
    v3.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/detail/log-contents/" + v3.EntityData.SegmentPath
    v3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v3.EntityData.Children = types.NewOrderedMap()
    v3.EntityData.Children.Append("scope", types.YChild{"Scope", &v3.Scope})
    v3.EntityData.Leafs = types.NewOrderedMap()
    v3.EntityData.Leafs.Append("category", types.YLeaf{"Category", v3.Category})
    v3.EntityData.Leafs.Append("message", types.YLeaf{"Message", v3.Message})

    v3.EntityData.YListKeys = []string {}

    return &(v3.EntityData)
}

// Install_Logs_Log_Detail_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Detail_LogContents_V3_Scope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetEntityData() *types.CommonEntityData {
    scope.EntityData.YFilter = scope.YFilter
    scope.EntityData.YangName = "scope"
    scope.EntityData.BundleName = "cisco_ios_xr"
    scope.EntityData.ParentYangName = "v3"
    scope.EntityData.SegmentPath = "scope"
    scope.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/detail/log-contents/v3/" + scope.EntityData.SegmentPath
    scope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scope.EntityData.Children = types.NewOrderedMap()
    scope.EntityData.Leafs = types.NewOrderedMap()
    scope.EntityData.Leafs.Append("admin-read", types.YLeaf{"AdminRead", scope.AdminRead})
    scope.EntityData.Leafs.Append("affected-sd-rs", types.YLeaf{"AffectedSdRs", scope.AffectedSdRs})

    scope.EntityData.YListKeys = []string {}

    return &(scope.EntityData)
}

// Install_Logs_Log_Communication
// Install communications
type Install_Logs_Log_Communication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Log contents.
    LogContents Install_Logs_Log_Communication_LogContents
}

func (communication *Install_Logs_Log_Communication) GetEntityData() *types.CommonEntityData {
    communication.EntityData.YFilter = communication.YFilter
    communication.EntityData.YangName = "communication"
    communication.EntityData.BundleName = "cisco_ios_xr"
    communication.EntityData.ParentYangName = "log"
    communication.EntityData.SegmentPath = "communication" + types.AddNoKeyToken(communication)
    communication.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/" + communication.EntityData.SegmentPath
    communication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    communication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    communication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    communication.EntityData.Children = types.NewOrderedMap()
    communication.EntityData.Children.Append("log-contents", types.YChild{"LogContents", &communication.LogContents})
    communication.EntityData.Leafs = types.NewOrderedMap()

    communication.EntityData.YListKeys = []string {}

    return &(communication.EntityData)
}

// Install_Logs_Log_Communication_LogContents
// Log contents
type Install_Logs_Log_Communication_LogContents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Communication_LogContents_V3
}

func (logContents *Install_Logs_Log_Communication_LogContents) GetEntityData() *types.CommonEntityData {
    logContents.EntityData.YFilter = logContents.YFilter
    logContents.EntityData.YangName = "log-contents"
    logContents.EntityData.BundleName = "cisco_ios_xr"
    logContents.EntityData.ParentYangName = "communication"
    logContents.EntityData.SegmentPath = "log-contents"
    logContents.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/communication/" + logContents.EntityData.SegmentPath
    logContents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logContents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logContents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logContents.EntityData.Children = types.NewOrderedMap()
    logContents.EntityData.Children.Append("v3", types.YChild{"V3", &logContents.V3})
    logContents.EntityData.Leafs = types.NewOrderedMap()
    logContents.EntityData.Leafs.Append("version", types.YLeaf{"Version", logContents.Version})

    logContents.EntityData.YListKeys = []string {}

    return &(logContents.EntityData)
}

// Install_Logs_Log_Communication_LogContents_V3
// v3
type Install_Logs_Log_Communication_LogContents_V3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Communication_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetEntityData() *types.CommonEntityData {
    v3.EntityData.YFilter = v3.YFilter
    v3.EntityData.YangName = "v3"
    v3.EntityData.BundleName = "cisco_ios_xr"
    v3.EntityData.ParentYangName = "log-contents"
    v3.EntityData.SegmentPath = "v3"
    v3.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/communication/log-contents/" + v3.EntityData.SegmentPath
    v3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v3.EntityData.Children = types.NewOrderedMap()
    v3.EntityData.Children.Append("scope", types.YChild{"Scope", &v3.Scope})
    v3.EntityData.Leafs = types.NewOrderedMap()
    v3.EntityData.Leafs.Append("category", types.YLeaf{"Category", v3.Category})
    v3.EntityData.Leafs.Append("message", types.YLeaf{"Message", v3.Message})

    v3.EntityData.YListKeys = []string {}

    return &(v3.EntityData)
}

// Install_Logs_Log_Communication_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Communication_LogContents_V3_Scope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetEntityData() *types.CommonEntityData {
    scope.EntityData.YFilter = scope.YFilter
    scope.EntityData.YangName = "scope"
    scope.EntityData.BundleName = "cisco_ios_xr"
    scope.EntityData.ParentYangName = "v3"
    scope.EntityData.SegmentPath = "scope"
    scope.EntityData.AbsolutePath = "Cisco-IOS-XR-installmgr-admin-oper:install/logs/log/communication/log-contents/v3/" + scope.EntityData.SegmentPath
    scope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scope.EntityData.Children = types.NewOrderedMap()
    scope.EntityData.Leafs = types.NewOrderedMap()
    scope.EntityData.Leafs.Append("admin-read", types.YLeaf{"AdminRead", scope.AdminRead})
    scope.EntityData.Leafs.Append("affected-sd-rs", types.YLeaf{"AffectedSdRs", scope.AffectedSdRs})

    scope.EntityData.YListKeys = []string {}

    return &(scope.EntityData)
}

