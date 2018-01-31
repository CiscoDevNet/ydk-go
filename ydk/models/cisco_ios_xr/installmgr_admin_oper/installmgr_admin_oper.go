// This module contains a collection of YANG definitions
// for Cisco IOS-XR installmgr package
// admin-plane operational data.
// 
// This module contains definitions
// for the following management objects:
//   install: Information of software packages and install
//     operations
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Install operation log history size. The type is interface{} with range:
    // -2147483648..2147483647.
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

func (install *Install) GetFilter() yfilter.YFilter { return install.YFilter }

func (install *Install) SetFilter(yf yfilter.YFilter) { install.YFilter = yf }

func (install *Install) GetGoName(yname string) string {
    if yname == "log-size" { return "LogSize" }
    if yname == "configuration-registers" { return "ConfigurationRegisters" }
    if yname == "request-statuses" { return "RequestStatuses" }
    if yname == "boot-variables" { return "BootVariables" }
    if yname == "software" { return "Software" }
    if yname == "software-inventory" { return "SoftwareInventory" }
    if yname == "issu" { return "Issu" }
    if yname == "boot-image" { return "BootImage" }
    if yname == "logs" { return "Logs" }
    return ""
}

func (install *Install) GetSegmentPath() string {
    return "Cisco-IOS-XR-installmgr-admin-oper:install"
}

func (install *Install) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configuration-registers" {
        return &install.ConfigurationRegisters
    }
    if childYangName == "request-statuses" {
        return &install.RequestStatuses
    }
    if childYangName == "boot-variables" {
        return &install.BootVariables
    }
    if childYangName == "software" {
        return &install.Software
    }
    if childYangName == "software-inventory" {
        return &install.SoftwareInventory
    }
    if childYangName == "issu" {
        return &install.Issu
    }
    if childYangName == "boot-image" {
        return &install.BootImage
    }
    if childYangName == "logs" {
        return &install.Logs
    }
    return nil
}

func (install *Install) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configuration-registers"] = &install.ConfigurationRegisters
    children["request-statuses"] = &install.RequestStatuses
    children["boot-variables"] = &install.BootVariables
    children["software"] = &install.Software
    children["software-inventory"] = &install.SoftwareInventory
    children["issu"] = &install.Issu
    children["boot-image"] = &install.BootImage
    children["logs"] = &install.Logs
    return children
}

func (install *Install) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-size"] = install.LogSize
    return leafs
}

func (install *Install) GetBundleName() string { return "cisco_ios_xr" }

func (install *Install) GetYangName() string { return "install" }

func (install *Install) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (install *Install) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (install *Install) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (install *Install) SetParent(parent types.Entity) { install.parent = parent }

func (install *Install) GetParent() types.Entity { return install.parent }

func (install *Install) GetParentYangName() string { return "Cisco-IOS-XR-installmgr-admin-oper" }

// Install_ConfigurationRegisters
// Configuration register (confreg) information
type Install_ConfigurationRegisters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration register for specific node. The type is slice of
    // Install_ConfigurationRegisters_ConfigurationRegister.
    ConfigurationRegister []Install_ConfigurationRegisters_ConfigurationRegister
}

func (configurationRegisters *Install_ConfigurationRegisters) GetFilter() yfilter.YFilter { return configurationRegisters.YFilter }

func (configurationRegisters *Install_ConfigurationRegisters) SetFilter(yf yfilter.YFilter) { configurationRegisters.YFilter = yf }

func (configurationRegisters *Install_ConfigurationRegisters) GetGoName(yname string) string {
    if yname == "configuration-register" { return "ConfigurationRegister" }
    return ""
}

func (configurationRegisters *Install_ConfigurationRegisters) GetSegmentPath() string {
    return "configuration-registers"
}

func (configurationRegisters *Install_ConfigurationRegisters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configuration-register" {
        for _, c := range configurationRegisters.ConfigurationRegister {
            if configurationRegisters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_ConfigurationRegisters_ConfigurationRegister{}
        configurationRegisters.ConfigurationRegister = append(configurationRegisters.ConfigurationRegister, child)
        return &configurationRegisters.ConfigurationRegister[len(configurationRegisters.ConfigurationRegister)-1]
    }
    return nil
}

func (configurationRegisters *Install_ConfigurationRegisters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configurationRegisters.ConfigurationRegister {
        children[configurationRegisters.ConfigurationRegister[i].GetSegmentPath()] = &configurationRegisters.ConfigurationRegister[i]
    }
    return children
}

func (configurationRegisters *Install_ConfigurationRegisters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configurationRegisters *Install_ConfigurationRegisters) GetBundleName() string { return "cisco_ios_xr" }

func (configurationRegisters *Install_ConfigurationRegisters) GetYangName() string { return "configuration-registers" }

func (configurationRegisters *Install_ConfigurationRegisters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configurationRegisters *Install_ConfigurationRegisters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configurationRegisters *Install_ConfigurationRegisters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configurationRegisters *Install_ConfigurationRegisters) SetParent(parent types.Entity) { configurationRegisters.parent = parent }

func (configurationRegisters *Install_ConfigurationRegisters) GetParent() types.Entity { return configurationRegisters.parent }

func (configurationRegisters *Install_ConfigurationRegisters) GetParentYangName() string { return "install" }

// Install_ConfigurationRegisters_ConfigurationRegister
// Configuration register for specific node
type Install_ConfigurationRegisters_ConfigurationRegister struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Configuration register value. The type is string with pattern:
    // [0-9a-fA-F]{1,8}. This attribute is mandatory.
    ConfigRegister interface{}
}

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetFilter() yfilter.YFilter { return configurationRegister.YFilter }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) SetFilter(yf yfilter.YFilter) { configurationRegister.YFilter = yf }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "config-register" { return "ConfigRegister" }
    return ""
}

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetSegmentPath() string {
    return "configuration-register" + "[node-name='" + fmt.Sprintf("%v", configurationRegister.NodeName) + "']"
}

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = configurationRegister.NodeName
    leafs["config-register"] = configurationRegister.ConfigRegister
    return leafs
}

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetBundleName() string { return "cisco_ios_xr" }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetYangName() string { return "configuration-register" }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) SetParent(parent types.Entity) { configurationRegister.parent = parent }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetParent() types.Entity { return configurationRegister.parent }

func (configurationRegister *Install_ConfigurationRegisters_ConfigurationRegister) GetParentYangName() string { return "configuration-registers" }

// Install_RequestStatuses
// Install operation request status
type Install_RequestStatuses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Request status Information. The type is slice of
    // Install_RequestStatuses_RequestStatus.
    RequestStatus []Install_RequestStatuses_RequestStatus
}

func (requestStatuses *Install_RequestStatuses) GetFilter() yfilter.YFilter { return requestStatuses.YFilter }

func (requestStatuses *Install_RequestStatuses) SetFilter(yf yfilter.YFilter) { requestStatuses.YFilter = yf }

func (requestStatuses *Install_RequestStatuses) GetGoName(yname string) string {
    if yname == "request-status" { return "RequestStatus" }
    return ""
}

func (requestStatuses *Install_RequestStatuses) GetSegmentPath() string {
    return "request-statuses"
}

func (requestStatuses *Install_RequestStatuses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "request-status" {
        for _, c := range requestStatuses.RequestStatus {
            if requestStatuses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_RequestStatuses_RequestStatus{}
        requestStatuses.RequestStatus = append(requestStatuses.RequestStatus, child)
        return &requestStatuses.RequestStatus[len(requestStatuses.RequestStatus)-1]
    }
    return nil
}

func (requestStatuses *Install_RequestStatuses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range requestStatuses.RequestStatus {
        children[requestStatuses.RequestStatus[i].GetSegmentPath()] = &requestStatuses.RequestStatus[i]
    }
    return children
}

func (requestStatuses *Install_RequestStatuses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (requestStatuses *Install_RequestStatuses) GetBundleName() string { return "cisco_ios_xr" }

func (requestStatuses *Install_RequestStatuses) GetYangName() string { return "request-statuses" }

func (requestStatuses *Install_RequestStatuses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requestStatuses *Install_RequestStatuses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requestStatuses *Install_RequestStatuses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requestStatuses *Install_RequestStatuses) SetParent(parent types.Entity) { requestStatuses.parent = parent }

func (requestStatuses *Install_RequestStatuses) GetParent() types.Entity { return requestStatuses.parent }

func (requestStatuses *Install_RequestStatuses) GetParentYangName() string { return "install" }

// Install_RequestStatuses_RequestStatus
// Request status Information
type Install_RequestStatuses_RequestStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Install operation request ID. The type is
    // interface{} with range: -2147483648..2147483647.
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
    IssuMessage []Install_RequestStatuses_RequestStatus_IssuMessage

    // Messages output to the user. The type is slice of
    // Install_RequestStatuses_RequestStatus_Message.
    Message []Install_RequestStatuses_RequestStatus_Message
}

func (requestStatus *Install_RequestStatuses_RequestStatus) GetFilter() yfilter.YFilter { return requestStatus.YFilter }

func (requestStatus *Install_RequestStatuses_RequestStatus) SetFilter(yf yfilter.YFilter) { requestStatus.YFilter = yf }

func (requestStatus *Install_RequestStatuses_RequestStatus) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "percentage" { return "Percentage" }
    if yname == "abort-state" { return "AbortState" }
    if yname == "downloaded-bytes" { return "DownloadedBytes" }
    if yname == "unanswered-query" { return "UnansweredQuery" }
    if yname == "operation-phase" { return "OperationPhase" }
    if yname == "request-information" { return "RequestInformation" }
    if yname == "abort-status" { return "AbortStatus" }
    if yname == "incremental-install-information" { return "IncrementalInstallInformation" }
    if yname == "issu-message" { return "IssuMessage" }
    if yname == "message" { return "Message" }
    return ""
}

func (requestStatus *Install_RequestStatuses_RequestStatus) GetSegmentPath() string {
    return "request-status" + "[request-id='" + fmt.Sprintf("%v", requestStatus.RequestId) + "']"
}

func (requestStatus *Install_RequestStatuses_RequestStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "request-information" {
        return &requestStatus.RequestInformation
    }
    if childYangName == "abort-status" {
        return &requestStatus.AbortStatus
    }
    if childYangName == "incremental-install-information" {
        return &requestStatus.IncrementalInstallInformation
    }
    if childYangName == "issu-message" {
        for _, c := range requestStatus.IssuMessage {
            if requestStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_RequestStatuses_RequestStatus_IssuMessage{}
        requestStatus.IssuMessage = append(requestStatus.IssuMessage, child)
        return &requestStatus.IssuMessage[len(requestStatus.IssuMessage)-1]
    }
    if childYangName == "message" {
        for _, c := range requestStatus.Message {
            if requestStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_RequestStatuses_RequestStatus_Message{}
        requestStatus.Message = append(requestStatus.Message, child)
        return &requestStatus.Message[len(requestStatus.Message)-1]
    }
    return nil
}

func (requestStatus *Install_RequestStatuses_RequestStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["request-information"] = &requestStatus.RequestInformation
    children["abort-status"] = &requestStatus.AbortStatus
    children["incremental-install-information"] = &requestStatus.IncrementalInstallInformation
    for i := range requestStatus.IssuMessage {
        children[requestStatus.IssuMessage[i].GetSegmentPath()] = &requestStatus.IssuMessage[i]
    }
    for i := range requestStatus.Message {
        children[requestStatus.Message[i].GetSegmentPath()] = &requestStatus.Message[i]
    }
    return children
}

func (requestStatus *Install_RequestStatuses_RequestStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = requestStatus.RequestId
    leafs["percentage"] = requestStatus.Percentage
    leafs["abort-state"] = requestStatus.AbortState
    leafs["downloaded-bytes"] = requestStatus.DownloadedBytes
    leafs["unanswered-query"] = requestStatus.UnansweredQuery
    leafs["operation-phase"] = requestStatus.OperationPhase
    return leafs
}

func (requestStatus *Install_RequestStatuses_RequestStatus) GetBundleName() string { return "cisco_ios_xr" }

func (requestStatus *Install_RequestStatuses_RequestStatus) GetYangName() string { return "request-status" }

func (requestStatus *Install_RequestStatuses_RequestStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requestStatus *Install_RequestStatuses_RequestStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requestStatus *Install_RequestStatuses_RequestStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requestStatus *Install_RequestStatuses_RequestStatus) SetParent(parent types.Entity) { requestStatus.parent = parent }

func (requestStatus *Install_RequestStatuses_RequestStatus) GetParent() types.Entity { return requestStatus.parent }

func (requestStatus *Install_RequestStatuses_RequestStatus) GetParentYangName() string { return "request-statuses" }

// Install_RequestStatuses_RequestStatus_RequestInformation
// Requested install operation
type Install_RequestStatuses_RequestStatus_RequestInformation struct {
    parent types.Entity
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

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetFilter() yfilter.YFilter { return requestInformation.YFilter }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) SetFilter(yf yfilter.YFilter) { requestInformation.YFilter = yf }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "user-id" { return "UserId" }
    if yname == "trigger-type" { return "TriggerType" }
    if yname == "request-option" { return "RequestOption" }
    if yname == "operation-type" { return "OperationType" }
    if yname == "operation-detail" { return "OperationDetail" }
    return ""
}

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetSegmentPath() string {
    return "request-information"
}

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = requestInformation.RequestId
    leafs["user-id"] = requestInformation.UserId
    leafs["trigger-type"] = requestInformation.TriggerType
    leafs["request-option"] = requestInformation.RequestOption
    leafs["operation-type"] = requestInformation.OperationType
    leafs["operation-detail"] = requestInformation.OperationDetail
    return leafs
}

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetBundleName() string { return "cisco_ios_xr" }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetYangName() string { return "request-information" }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) SetParent(parent types.Entity) { requestInformation.parent = parent }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetParent() types.Entity { return requestInformation.parent }

func (requestInformation *Install_RequestStatuses_RequestStatus_RequestInformation) GetParentYangName() string { return "request-status" }

// Install_RequestStatuses_RequestStatus_AbortStatus
// How the abort will occur
type Install_RequestStatuses_RequestStatus_AbortStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Method of abort. The type is InstmgrIssuAbortMethod.
    AbortMethod interface{}

    // Impact of abort. The type is InstmgrIssuAbortImpact.
    AbortImpact interface{}
}

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetFilter() yfilter.YFilter { return abortStatus.YFilter }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) SetFilter(yf yfilter.YFilter) { abortStatus.YFilter = yf }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetGoName(yname string) string {
    if yname == "abort-method" { return "AbortMethod" }
    if yname == "abort-impact" { return "AbortImpact" }
    return ""
}

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetSegmentPath() string {
    return "abort-status"
}

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["abort-method"] = abortStatus.AbortMethod
    leafs["abort-impact"] = abortStatus.AbortImpact
    return leafs
}

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetBundleName() string { return "cisco_ios_xr" }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetYangName() string { return "abort-status" }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) SetParent(parent types.Entity) { abortStatus.parent = parent }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetParent() types.Entity { return abortStatus.parent }

func (abortStatus *Install_RequestStatuses_RequestStatus_AbortStatus) GetParentYangName() string { return "request-status" }

// Install_RequestStatuses_RequestStatus_IncrementalInstallInformation
// Incremental Install information
type Install_RequestStatuses_RequestStatus_IncrementalInstallInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install direction. The type is InstmgrBagIiDirection.
    Direction interface{}

    // First error during incremental install. The type is string.
    IiError interface{}

    // Participating nodes. The type is slice of
    // Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes.
    Nodes []Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes
}

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetFilter() yfilter.YFilter { return incrementalInstallInformation.YFilter }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) SetFilter(yf yfilter.YFilter) { incrementalInstallInformation.YFilter = yf }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "ii-error" { return "IiError" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetSegmentPath() string {
    return "incremental-install-information"
}

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        for _, c := range incrementalInstallInformation.Nodes {
            if incrementalInstallInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes{}
        incrementalInstallInformation.Nodes = append(incrementalInstallInformation.Nodes, child)
        return &incrementalInstallInformation.Nodes[len(incrementalInstallInformation.Nodes)-1]
    }
    return nil
}

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range incrementalInstallInformation.Nodes {
        children[incrementalInstallInformation.Nodes[i].GetSegmentPath()] = &incrementalInstallInformation.Nodes[i]
    }
    return children
}

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = incrementalInstallInformation.Direction
    leafs["ii-error"] = incrementalInstallInformation.IiError
    return leafs
}

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetBundleName() string { return "cisco_ios_xr" }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetYangName() string { return "incremental-install-information" }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) SetParent(parent types.Entity) { incrementalInstallInformation.parent = parent }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetParent() types.Entity { return incrementalInstallInformation.parent }

func (incrementalInstallInformation *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation) GetParentYangName() string { return "request-status" }

// Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes
// Participating nodes
type Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node identifier. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // State. The type is InstmgrBagIiState.
    State interface{}
}

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "state" { return "State" }
    return ""
}

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodes.NodeName
    leafs["state"] = nodes.State
    return leafs
}

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetYangName() string { return "nodes" }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Install_RequestStatuses_RequestStatus_IncrementalInstallInformation_Nodes) GetParentYangName() string { return "incremental-install-information" }

// Install_RequestStatuses_RequestStatus_IssuMessage
// Messages related to ISSU operations
type Install_RequestStatuses_RequestStatus_IssuMessage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_RequestStatuses_RequestStatus_IssuMessage_Scope
}

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetFilter() yfilter.YFilter { return issuMessage.YFilter }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) SetFilter(yf yfilter.YFilter) { issuMessage.YFilter = yf }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "message" { return "Message" }
    if yname == "scope" { return "Scope" }
    return ""
}

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetSegmentPath() string {
    return "issu-message"
}

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scope" {
        return &issuMessage.Scope
    }
    return nil
}

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scope"] = &issuMessage.Scope
    return children
}

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = issuMessage.Category
    leafs["message"] = issuMessage.Message
    return leafs
}

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetBundleName() string { return "cisco_ios_xr" }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetYangName() string { return "issu-message" }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) SetParent(parent types.Entity) { issuMessage.parent = parent }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetParent() types.Entity { return issuMessage.parent }

func (issuMessage *Install_RequestStatuses_RequestStatus_IssuMessage) GetParentYangName() string { return "request-status" }

// Install_RequestStatuses_RequestStatus_IssuMessage_Scope
// Scope of the message
type Install_RequestStatuses_RequestStatus_IssuMessage_Scope struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which LRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetFilter() yfilter.YFilter { return scope.YFilter }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) SetFilter(yf yfilter.YFilter) { scope.YFilter = yf }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetGoName(yname string) string {
    if yname == "admin-read" { return "AdminRead" }
    if yname == "affected-sd-rs" { return "AffectedSdRs" }
    return ""
}

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetSegmentPath() string {
    return "scope"
}

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-read"] = scope.AdminRead
    leafs["affected-sd-rs"] = scope.AffectedSdRs
    return leafs
}

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetBundleName() string { return "cisco_ios_xr" }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetYangName() string { return "scope" }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) SetParent(parent types.Entity) { scope.parent = parent }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetParent() types.Entity { return scope.parent }

func (scope *Install_RequestStatuses_RequestStatus_IssuMessage_Scope) GetParentYangName() string { return "issu-message" }

// Install_RequestStatuses_RequestStatus_Message
// Messages output to the user
type Install_RequestStatuses_RequestStatus_Message struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_RequestStatuses_RequestStatus_Message_Scope
}

func (message *Install_RequestStatuses_RequestStatus_Message) GetFilter() yfilter.YFilter { return message.YFilter }

func (message *Install_RequestStatuses_RequestStatus_Message) SetFilter(yf yfilter.YFilter) { message.YFilter = yf }

func (message *Install_RequestStatuses_RequestStatus_Message) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "message" { return "Message" }
    if yname == "scope" { return "Scope" }
    return ""
}

func (message *Install_RequestStatuses_RequestStatus_Message) GetSegmentPath() string {
    return "message"
}

func (message *Install_RequestStatuses_RequestStatus_Message) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scope" {
        return &message.Scope
    }
    return nil
}

func (message *Install_RequestStatuses_RequestStatus_Message) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scope"] = &message.Scope
    return children
}

func (message *Install_RequestStatuses_RequestStatus_Message) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = message.Category
    leafs["message"] = message.Message
    return leafs
}

func (message *Install_RequestStatuses_RequestStatus_Message) GetBundleName() string { return "cisco_ios_xr" }

func (message *Install_RequestStatuses_RequestStatus_Message) GetYangName() string { return "message" }

func (message *Install_RequestStatuses_RequestStatus_Message) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (message *Install_RequestStatuses_RequestStatus_Message) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (message *Install_RequestStatuses_RequestStatus_Message) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (message *Install_RequestStatuses_RequestStatus_Message) SetParent(parent types.Entity) { message.parent = parent }

func (message *Install_RequestStatuses_RequestStatus_Message) GetParent() types.Entity { return message.parent }

func (message *Install_RequestStatuses_RequestStatus_Message) GetParentYangName() string { return "request-status" }

// Install_RequestStatuses_RequestStatus_Message_Scope
// Scope of the message
type Install_RequestStatuses_RequestStatus_Message_Scope struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which LRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetFilter() yfilter.YFilter { return scope.YFilter }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) SetFilter(yf yfilter.YFilter) { scope.YFilter = yf }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetGoName(yname string) string {
    if yname == "admin-read" { return "AdminRead" }
    if yname == "affected-sd-rs" { return "AffectedSdRs" }
    return ""
}

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetSegmentPath() string {
    return "scope"
}

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-read"] = scope.AdminRead
    leafs["affected-sd-rs"] = scope.AffectedSdRs
    return leafs
}

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetBundleName() string { return "cisco_ios_xr" }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetYangName() string { return "scope" }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) SetParent(parent types.Entity) { scope.parent = parent }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetParent() types.Entity { return scope.parent }

func (scope *Install_RequestStatuses_RequestStatus_Message_Scope) GetParentYangName() string { return "message" }

// Install_BootVariables
// Boot variable information
type Install_BootVariables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Boot variable for specific node. The type is slice of
    // Install_BootVariables_BootVariable.
    BootVariable []Install_BootVariables_BootVariable
}

func (bootVariables *Install_BootVariables) GetFilter() yfilter.YFilter { return bootVariables.YFilter }

func (bootVariables *Install_BootVariables) SetFilter(yf yfilter.YFilter) { bootVariables.YFilter = yf }

func (bootVariables *Install_BootVariables) GetGoName(yname string) string {
    if yname == "boot-variable" { return "BootVariable" }
    return ""
}

func (bootVariables *Install_BootVariables) GetSegmentPath() string {
    return "boot-variables"
}

func (bootVariables *Install_BootVariables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "boot-variable" {
        for _, c := range bootVariables.BootVariable {
            if bootVariables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_BootVariables_BootVariable{}
        bootVariables.BootVariable = append(bootVariables.BootVariable, child)
        return &bootVariables.BootVariable[len(bootVariables.BootVariable)-1]
    }
    return nil
}

func (bootVariables *Install_BootVariables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bootVariables.BootVariable {
        children[bootVariables.BootVariable[i].GetSegmentPath()] = &bootVariables.BootVariable[i]
    }
    return children
}

func (bootVariables *Install_BootVariables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bootVariables *Install_BootVariables) GetBundleName() string { return "cisco_ios_xr" }

func (bootVariables *Install_BootVariables) GetYangName() string { return "boot-variables" }

func (bootVariables *Install_BootVariables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootVariables *Install_BootVariables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootVariables *Install_BootVariables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootVariables *Install_BootVariables) SetParent(parent types.Entity) { bootVariables.parent = parent }

func (bootVariables *Install_BootVariables) GetParent() types.Entity { return bootVariables.parent }

func (bootVariables *Install_BootVariables) GetParentYangName() string { return "install" }

// Install_BootVariables_BootVariable
// Boot variable for specific node
type Install_BootVariables_BootVariable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Boot variable value. The type is string. This attribute is mandatory.
    BootVariable interface{}
}

func (bootVariable *Install_BootVariables_BootVariable) GetFilter() yfilter.YFilter { return bootVariable.YFilter }

func (bootVariable *Install_BootVariables_BootVariable) SetFilter(yf yfilter.YFilter) { bootVariable.YFilter = yf }

func (bootVariable *Install_BootVariables_BootVariable) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "boot-variable" { return "BootVariable" }
    return ""
}

func (bootVariable *Install_BootVariables_BootVariable) GetSegmentPath() string {
    return "boot-variable" + "[node-name='" + fmt.Sprintf("%v", bootVariable.NodeName) + "']"
}

func (bootVariable *Install_BootVariables_BootVariable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootVariable *Install_BootVariables_BootVariable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootVariable *Install_BootVariables_BootVariable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = bootVariable.NodeName
    leafs["boot-variable"] = bootVariable.BootVariable
    return leafs
}

func (bootVariable *Install_BootVariables_BootVariable) GetBundleName() string { return "cisco_ios_xr" }

func (bootVariable *Install_BootVariables_BootVariable) GetYangName() string { return "boot-variable" }

func (bootVariable *Install_BootVariables_BootVariable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootVariable *Install_BootVariables_BootVariable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootVariable *Install_BootVariables_BootVariable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootVariable *Install_BootVariables_BootVariable) SetParent(parent types.Entity) { bootVariable.parent = parent }

func (bootVariable *Install_BootVariables_BootVariable) GetParent() types.Entity { return bootVariable.parent }

func (bootVariable *Install_BootVariables_BootVariable) GetParentYangName() string { return "boot-variables" }

// Install_Software
// Software package,component and alias information
type Install_Software struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package alias information.
    AliasDevices Install_Software_AliasDevices

    // Package information.
    PackageDevices Install_Software_PackageDevices

    // Software component information.
    ComponentDevices Install_Software_ComponentDevices
}

func (software *Install_Software) GetFilter() yfilter.YFilter { return software.YFilter }

func (software *Install_Software) SetFilter(yf yfilter.YFilter) { software.YFilter = yf }

func (software *Install_Software) GetGoName(yname string) string {
    if yname == "alias-devices" { return "AliasDevices" }
    if yname == "package-devices" { return "PackageDevices" }
    if yname == "component-devices" { return "ComponentDevices" }
    return ""
}

func (software *Install_Software) GetSegmentPath() string {
    return "software"
}

func (software *Install_Software) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alias-devices" {
        return &software.AliasDevices
    }
    if childYangName == "package-devices" {
        return &software.PackageDevices
    }
    if childYangName == "component-devices" {
        return &software.ComponentDevices
    }
    return nil
}

func (software *Install_Software) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["alias-devices"] = &software.AliasDevices
    children["package-devices"] = &software.PackageDevices
    children["component-devices"] = &software.ComponentDevices
    return children
}

func (software *Install_Software) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (software *Install_Software) GetBundleName() string { return "cisco_ios_xr" }

func (software *Install_Software) GetYangName() string { return "software" }

func (software *Install_Software) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (software *Install_Software) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (software *Install_Software) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (software *Install_Software) SetParent(parent types.Entity) { software.parent = parent }

func (software *Install_Software) GetParent() types.Entity { return software.parent }

func (software *Install_Software) GetParentYangName() string { return "install" }

// Install_Software_AliasDevices
// Package alias information
type Install_Software_AliasDevices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package alias information for specific device. The type is slice of
    // Install_Software_AliasDevices_AliasDevice.
    AliasDevice []Install_Software_AliasDevices_AliasDevice
}

func (aliasDevices *Install_Software_AliasDevices) GetFilter() yfilter.YFilter { return aliasDevices.YFilter }

func (aliasDevices *Install_Software_AliasDevices) SetFilter(yf yfilter.YFilter) { aliasDevices.YFilter = yf }

func (aliasDevices *Install_Software_AliasDevices) GetGoName(yname string) string {
    if yname == "alias-device" { return "AliasDevice" }
    return ""
}

func (aliasDevices *Install_Software_AliasDevices) GetSegmentPath() string {
    return "alias-devices"
}

func (aliasDevices *Install_Software_AliasDevices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alias-device" {
        for _, c := range aliasDevices.AliasDevice {
            if aliasDevices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Software_AliasDevices_AliasDevice{}
        aliasDevices.AliasDevice = append(aliasDevices.AliasDevice, child)
        return &aliasDevices.AliasDevice[len(aliasDevices.AliasDevice)-1]
    }
    return nil
}

func (aliasDevices *Install_Software_AliasDevices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range aliasDevices.AliasDevice {
        children[aliasDevices.AliasDevice[i].GetSegmentPath()] = &aliasDevices.AliasDevice[i]
    }
    return children
}

func (aliasDevices *Install_Software_AliasDevices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aliasDevices *Install_Software_AliasDevices) GetBundleName() string { return "cisco_ios_xr" }

func (aliasDevices *Install_Software_AliasDevices) GetYangName() string { return "alias-devices" }

func (aliasDevices *Install_Software_AliasDevices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aliasDevices *Install_Software_AliasDevices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aliasDevices *Install_Software_AliasDevices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aliasDevices *Install_Software_AliasDevices) SetParent(parent types.Entity) { aliasDevices.parent = parent }

func (aliasDevices *Install_Software_AliasDevices) GetParent() types.Entity { return aliasDevices.parent }

func (aliasDevices *Install_Software_AliasDevices) GetParentYangName() string { return "software" }

// Install_Software_AliasDevices_AliasDevice
// Package alias information for specific device
type Install_Software_AliasDevices_AliasDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Device Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    DeviceName interface{}

    // alias information.
    Aliases Install_Software_AliasDevices_AliasDevice_Aliases
}

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetFilter() yfilter.YFilter { return aliasDevice.YFilter }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) SetFilter(yf yfilter.YFilter) { aliasDevice.YFilter = yf }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "aliases" { return "Aliases" }
    return ""
}

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetSegmentPath() string {
    return "alias-device" + "[device-name='" + fmt.Sprintf("%v", aliasDevice.DeviceName) + "']"
}

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aliases" {
        return &aliasDevice.Aliases
    }
    return nil
}

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aliases"] = &aliasDevice.Aliases
    return children
}

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = aliasDevice.DeviceName
    return leafs
}

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetBundleName() string { return "cisco_ios_xr" }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetYangName() string { return "alias-device" }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) SetParent(parent types.Entity) { aliasDevice.parent = parent }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetParent() types.Entity { return aliasDevice.parent }

func (aliasDevice *Install_Software_AliasDevices_AliasDevice) GetParentYangName() string { return "alias-devices" }

// Install_Software_AliasDevices_AliasDevice_Aliases
// alias information
type Install_Software_AliasDevices_AliasDevice_Aliases struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Aliases for specific package. The type is slice of
    // Install_Software_AliasDevices_AliasDevice_Aliases_Alias.
    Alias []Install_Software_AliasDevices_AliasDevice_Aliases_Alias
}

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetFilter() yfilter.YFilter { return aliases.YFilter }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) SetFilter(yf yfilter.YFilter) { aliases.YFilter = yf }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetGoName(yname string) string {
    if yname == "alias" { return "Alias" }
    return ""
}

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetSegmentPath() string {
    return "aliases"
}

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alias" {
        for _, c := range aliases.Alias {
            if aliases.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Software_AliasDevices_AliasDevice_Aliases_Alias{}
        aliases.Alias = append(aliases.Alias, child)
        return &aliases.Alias[len(aliases.Alias)-1]
    }
    return nil
}

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range aliases.Alias {
        children[aliases.Alias[i].GetSegmentPath()] = &aliases.Alias[i]
    }
    return children
}

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetBundleName() string { return "cisco_ios_xr" }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetYangName() string { return "aliases" }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) SetParent(parent types.Entity) { aliases.parent = parent }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetParent() types.Entity { return aliases.parent }

func (aliases *Install_Software_AliasDevices_AliasDevice_Aliases) GetParentYangName() string { return "alias-device" }

// Install_Software_AliasDevices_AliasDevice_Aliases_Alias
// Aliases for specific package
type Install_Software_AliasDevices_AliasDevice_Aliases_Alias struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Package Name. The type is string.
    PackageName interface{}

    // Alias names. The type is string. This attribute is mandatory.
    AliasNames interface{}
}

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetFilter() yfilter.YFilter { return alias.YFilter }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) SetFilter(yf yfilter.YFilter) { alias.YFilter = yf }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetGoName(yname string) string {
    if yname == "package-name" { return "PackageName" }
    if yname == "alias-names" { return "AliasNames" }
    return ""
}

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetSegmentPath() string {
    return "alias" + "[package-name='" + fmt.Sprintf("%v", alias.PackageName) + "']"
}

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["package-name"] = alias.PackageName
    leafs["alias-names"] = alias.AliasNames
    return leafs
}

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetBundleName() string { return "cisco_ios_xr" }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetYangName() string { return "alias" }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) SetParent(parent types.Entity) { alias.parent = parent }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetParent() types.Entity { return alias.parent }

func (alias *Install_Software_AliasDevices_AliasDevice_Aliases_Alias) GetParentYangName() string { return "aliases" }

// Install_Software_PackageDevices
// Package information
type Install_Software_PackageDevices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package information for specific device. The type is slice of
    // Install_Software_PackageDevices_PackageDevice.
    PackageDevice []Install_Software_PackageDevices_PackageDevice
}

func (packageDevices *Install_Software_PackageDevices) GetFilter() yfilter.YFilter { return packageDevices.YFilter }

func (packageDevices *Install_Software_PackageDevices) SetFilter(yf yfilter.YFilter) { packageDevices.YFilter = yf }

func (packageDevices *Install_Software_PackageDevices) GetGoName(yname string) string {
    if yname == "package-device" { return "PackageDevice" }
    return ""
}

func (packageDevices *Install_Software_PackageDevices) GetSegmentPath() string {
    return "package-devices"
}

func (packageDevices *Install_Software_PackageDevices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package-device" {
        for _, c := range packageDevices.PackageDevice {
            if packageDevices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Software_PackageDevices_PackageDevice{}
        packageDevices.PackageDevice = append(packageDevices.PackageDevice, child)
        return &packageDevices.PackageDevice[len(packageDevices.PackageDevice)-1]
    }
    return nil
}

func (packageDevices *Install_Software_PackageDevices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range packageDevices.PackageDevice {
        children[packageDevices.PackageDevice[i].GetSegmentPath()] = &packageDevices.PackageDevice[i]
    }
    return children
}

func (packageDevices *Install_Software_PackageDevices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packageDevices *Install_Software_PackageDevices) GetBundleName() string { return "cisco_ios_xr" }

func (packageDevices *Install_Software_PackageDevices) GetYangName() string { return "package-devices" }

func (packageDevices *Install_Software_PackageDevices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packageDevices *Install_Software_PackageDevices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packageDevices *Install_Software_PackageDevices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packageDevices *Install_Software_PackageDevices) SetParent(parent types.Entity) { packageDevices.parent = parent }

func (packageDevices *Install_Software_PackageDevices) GetParent() types.Entity { return packageDevices.parent }

func (packageDevices *Install_Software_PackageDevices) GetParentYangName() string { return "software" }

// Install_Software_PackageDevices_PackageDevice
// Package information for specific device
type Install_Software_PackageDevices_PackageDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Device Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    DeviceName interface{}

    // Package information for specific package.
    Packages Install_Software_PackageDevices_PackageDevice_Packages
}

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetFilter() yfilter.YFilter { return packageDevice.YFilter }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) SetFilter(yf yfilter.YFilter) { packageDevice.YFilter = yf }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "packages" { return "Packages" }
    return ""
}

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetSegmentPath() string {
    return "package-device" + "[device-name='" + fmt.Sprintf("%v", packageDevice.DeviceName) + "']"
}

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "packages" {
        return &packageDevice.Packages
    }
    return nil
}

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["packages"] = &packageDevice.Packages
    return children
}

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = packageDevice.DeviceName
    return leafs
}

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetBundleName() string { return "cisco_ios_xr" }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetYangName() string { return "package-device" }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) SetParent(parent types.Entity) { packageDevice.parent = parent }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetParent() types.Entity { return packageDevice.parent }

func (packageDevice *Install_Software_PackageDevices_PackageDevice) GetParentYangName() string { return "package-devices" }

// Install_Software_PackageDevices_PackageDevice_Packages
// Package information for specific package
type Install_Software_PackageDevices_PackageDevice_Packages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package information. The type is slice of
    // Install_Software_PackageDevices_PackageDevice_Packages_Package.
    Package []Install_Software_PackageDevices_PackageDevice_Packages_Package
}

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetFilter() yfilter.YFilter { return packages.YFilter }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) SetFilter(yf yfilter.YFilter) { packages.YFilter = yf }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetGoName(yname string) string {
    if yname == "package" { return "Package" }
    return ""
}

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetSegmentPath() string {
    return "packages"
}

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        for _, c := range packages.Package {
            if packages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Software_PackageDevices_PackageDevice_Packages_Package{}
        packages.Package = append(packages.Package, child)
        return &packages.Package[len(packages.Package)-1]
    }
    return nil
}

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range packages.Package {
        children[packages.Package[i].GetSegmentPath()] = &packages.Package[i]
    }
    return children
}

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetBundleName() string { return "cisco_ios_xr" }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetYangName() string { return "packages" }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) SetParent(parent types.Entity) { packages.parent = parent }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetParent() types.Entity { return packages.parent }

func (packages *Install_Software_PackageDevices_PackageDevice_Packages) GetParentYangName() string { return "package-device" }

// Install_Software_PackageDevices_PackageDevice_Packages_Package
// Package information
type Install_Software_PackageDevices_PackageDevice_Packages_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    SubPkg []Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg
}

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetGoName(yname string) string {
    if yname == "package-name" { return "PackageName" }
    if yname == "name" { return "Name" }
    if yname == "version" { return "Version" }
    if yname == "description" { return "Description" }
    if yname == "release" { return "Release" }
    if yname == "vendor" { return "Vendor" }
    if yname == "date" { return "Date" }
    if yname == "source" { return "Source" }
    if yname == "base" { return "Base" }
    if yname == "bootable" { return "Bootable" }
    if yname == "composite" { return "Composite" }
    if yname == "restart-info" { return "RestartInfo" }
    if yname == "package-type" { return "PackageType" }
    if yname == "group-type" { return "GroupType" }
    if yname == "depth" { return "Depth" }
    if yname == "uncompressed-size" { return "UncompressedSize" }
    if yname == "compressed-size" { return "CompressedSize" }
    if yname == "cards" { return "Cards" }
    if yname == "sub-pkg" { return "SubPkg" }
    return ""
}

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetSegmentPath() string {
    return "package" + "[package-name='" + fmt.Sprintf("%v", self.PackageName) + "']"
}

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sub-pkg" {
        for _, c := range self.SubPkg {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg{}
        self.SubPkg = append(self.SubPkg, child)
        return &self.SubPkg[len(self.SubPkg)-1]
    }
    return nil
}

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.SubPkg {
        children[self.SubPkg[i].GetSegmentPath()] = &self.SubPkg[i]
    }
    return children
}

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["package-name"] = self.PackageName
    leafs["name"] = self.Name
    leafs["version"] = self.Version
    leafs["description"] = self.Description
    leafs["release"] = self.Release
    leafs["vendor"] = self.Vendor
    leafs["date"] = self.Date
    leafs["source"] = self.Source
    leafs["base"] = self.Base
    leafs["bootable"] = self.Bootable
    leafs["composite"] = self.Composite
    leafs["restart-info"] = self.RestartInfo
    leafs["package-type"] = self.PackageType
    leafs["group-type"] = self.GroupType
    leafs["depth"] = self.Depth
    leafs["uncompressed-size"] = self.UncompressedSize
    leafs["compressed-size"] = self.CompressedSize
    leafs["cards"] = self.Cards
    return leafs
}

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetYangName() string { return "package" }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetParent() types.Entity { return self.parent }

func (self *Install_Software_PackageDevices_PackageDevice_Packages_Package) GetParentYangName() string { return "packages" }

// Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg
// Sub-package info of the package
type Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the sub-package. The type is string.
    Name interface{}

    // Node types of the package. The type is interface{} with range:
    // 0..18446744073709551615.
    NodeTypes interface{}
}

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetFilter() yfilter.YFilter { return subPkg.YFilter }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) SetFilter(yf yfilter.YFilter) { subPkg.YFilter = yf }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "node-types" { return "NodeTypes" }
    return ""
}

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetSegmentPath() string {
    return "sub-pkg"
}

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = subPkg.Name
    leafs["node-types"] = subPkg.NodeTypes
    return leafs
}

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetBundleName() string { return "cisco_ios_xr" }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetYangName() string { return "sub-pkg" }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) SetParent(parent types.Entity) { subPkg.parent = parent }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetParent() types.Entity { return subPkg.parent }

func (subPkg *Install_Software_PackageDevices_PackageDevice_Packages_Package_SubPkg) GetParentYangName() string { return "package" }

// Install_Software_ComponentDevices
// Software component information
type Install_Software_ComponentDevices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component information for specific device. The type is slice of
    // Install_Software_ComponentDevices_ComponentDevice.
    ComponentDevice []Install_Software_ComponentDevices_ComponentDevice
}

func (componentDevices *Install_Software_ComponentDevices) GetFilter() yfilter.YFilter { return componentDevices.YFilter }

func (componentDevices *Install_Software_ComponentDevices) SetFilter(yf yfilter.YFilter) { componentDevices.YFilter = yf }

func (componentDevices *Install_Software_ComponentDevices) GetGoName(yname string) string {
    if yname == "component-device" { return "ComponentDevice" }
    return ""
}

func (componentDevices *Install_Software_ComponentDevices) GetSegmentPath() string {
    return "component-devices"
}

func (componentDevices *Install_Software_ComponentDevices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "component-device" {
        for _, c := range componentDevices.ComponentDevice {
            if componentDevices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Software_ComponentDevices_ComponentDevice{}
        componentDevices.ComponentDevice = append(componentDevices.ComponentDevice, child)
        return &componentDevices.ComponentDevice[len(componentDevices.ComponentDevice)-1]
    }
    return nil
}

func (componentDevices *Install_Software_ComponentDevices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range componentDevices.ComponentDevice {
        children[componentDevices.ComponentDevice[i].GetSegmentPath()] = &componentDevices.ComponentDevice[i]
    }
    return children
}

func (componentDevices *Install_Software_ComponentDevices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (componentDevices *Install_Software_ComponentDevices) GetBundleName() string { return "cisco_ios_xr" }

func (componentDevices *Install_Software_ComponentDevices) GetYangName() string { return "component-devices" }

func (componentDevices *Install_Software_ComponentDevices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (componentDevices *Install_Software_ComponentDevices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (componentDevices *Install_Software_ComponentDevices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (componentDevices *Install_Software_ComponentDevices) SetParent(parent types.Entity) { componentDevices.parent = parent }

func (componentDevices *Install_Software_ComponentDevices) GetParent() types.Entity { return componentDevices.parent }

func (componentDevices *Install_Software_ComponentDevices) GetParentYangName() string { return "software" }

// Install_Software_ComponentDevices_ComponentDevice
// Component information for specific device
type Install_Software_ComponentDevices_ComponentDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Device Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    DeviceName interface{}

    // Software package information.
    ComponentPackages Install_Software_ComponentDevices_ComponentDevice_ComponentPackages
}

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetFilter() yfilter.YFilter { return componentDevice.YFilter }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) SetFilter(yf yfilter.YFilter) { componentDevice.YFilter = yf }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "component-packages" { return "ComponentPackages" }
    return ""
}

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetSegmentPath() string {
    return "component-device" + "[device-name='" + fmt.Sprintf("%v", componentDevice.DeviceName) + "']"
}

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "component-packages" {
        return &componentDevice.ComponentPackages
    }
    return nil
}

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["component-packages"] = &componentDevice.ComponentPackages
    return children
}

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = componentDevice.DeviceName
    return leafs
}

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetBundleName() string { return "cisco_ios_xr" }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetYangName() string { return "component-device" }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) SetParent(parent types.Entity) { componentDevice.parent = parent }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetParent() types.Entity { return componentDevice.parent }

func (componentDevice *Install_Software_ComponentDevices_ComponentDevice) GetParentYangName() string { return "component-devices" }

// Install_Software_ComponentDevices_ComponentDevice_ComponentPackages
// Software package information
type Install_Software_ComponentDevices_ComponentDevice_ComponentPackages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component information for specific package. The type is slice of
    // Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage.
    ComponentPackage []Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage
}

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetFilter() yfilter.YFilter { return componentPackages.YFilter }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) SetFilter(yf yfilter.YFilter) { componentPackages.YFilter = yf }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetGoName(yname string) string {
    if yname == "component-package" { return "ComponentPackage" }
    return ""
}

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetSegmentPath() string {
    return "component-packages"
}

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "component-package" {
        for _, c := range componentPackages.ComponentPackage {
            if componentPackages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage{}
        componentPackages.ComponentPackage = append(componentPackages.ComponentPackage, child)
        return &componentPackages.ComponentPackage[len(componentPackages.ComponentPackage)-1]
    }
    return nil
}

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range componentPackages.ComponentPackage {
        children[componentPackages.ComponentPackage[i].GetSegmentPath()] = &componentPackages.ComponentPackage[i]
    }
    return children
}

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetBundleName() string { return "cisco_ios_xr" }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetYangName() string { return "component-packages" }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) SetParent(parent types.Entity) { componentPackages.parent = parent }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetParent() types.Entity { return componentPackages.parent }

func (componentPackages *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages) GetParentYangName() string { return "component-device" }

// Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage
// Component information for specific package
type Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Package Name. The type is string.
    PackageName interface{}

    // Component information. The type is slice of
    // Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component.
    Component []Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component
}

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetFilter() yfilter.YFilter { return componentPackage.YFilter }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) SetFilter(yf yfilter.YFilter) { componentPackage.YFilter = yf }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetGoName(yname string) string {
    if yname == "package-name" { return "PackageName" }
    if yname == "component" { return "Component" }
    return ""
}

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetSegmentPath() string {
    return "component-package" + "[package-name='" + fmt.Sprintf("%v", componentPackage.PackageName) + "']"
}

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "component" {
        for _, c := range componentPackage.Component {
            if componentPackage.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component{}
        componentPackage.Component = append(componentPackage.Component, child)
        return &componentPackage.Component[len(componentPackage.Component)-1]
    }
    return nil
}

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range componentPackage.Component {
        children[componentPackage.Component[i].GetSegmentPath()] = &componentPackage.Component[i]
    }
    return children
}

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["package-name"] = componentPackage.PackageName
    return leafs
}

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetBundleName() string { return "cisco_ios_xr" }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetYangName() string { return "component-package" }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) SetParent(parent types.Entity) { componentPackage.parent = parent }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetParent() types.Entity { return componentPackage.parent }

func (componentPackage *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage) GetParentYangName() string { return "component-packages" }

// Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component
// Component information
type Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetFilter() yfilter.YFilter { return component.YFilter }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) SetFilter(yf yfilter.YFilter) { component.YFilter = yf }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetGoName(yname string) string {
    if yname == "component-name" { return "ComponentName" }
    if yname == "name" { return "Name" }
    if yname == "version" { return "Version" }
    if yname == "release" { return "Release" }
    if yname == "description" { return "Description" }
    if yname == "files" { return "Files" }
    return ""
}

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetSegmentPath() string {
    return "component" + "[component-name='" + fmt.Sprintf("%v", component.ComponentName) + "']"
}

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["component-name"] = component.ComponentName
    leafs["name"] = component.Name
    leafs["version"] = component.Version
    leafs["release"] = component.Release
    leafs["description"] = component.Description
    leafs["files"] = component.Files
    return leafs
}

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetBundleName() string { return "cisco_ios_xr" }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetYangName() string { return "component" }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) SetParent(parent types.Entity) { component.parent = parent }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetParent() types.Entity { return component.parent }

func (component *Install_Software_ComponentDevices_ComponentDevice_ComponentPackages_ComponentPackage_Component) GetParentYangName() string { return "component-package" }

// Install_SoftwareInventory
// Information of install operations
type Install_SoftwareInventory struct {
    parent types.Entity
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

func (softwareInventory *Install_SoftwareInventory) GetFilter() yfilter.YFilter { return softwareInventory.YFilter }

func (softwareInventory *Install_SoftwareInventory) SetFilter(yf yfilter.YFilter) { softwareInventory.YFilter = yf }

func (softwareInventory *Install_SoftwareInventory) GetGoName(yname string) string {
    if yname == "committed" { return "Committed" }
    if yname == "inactive" { return "Inactive" }
    if yname == "requests" { return "Requests" }
    if yname == "active" { return "Active" }
    return ""
}

func (softwareInventory *Install_SoftwareInventory) GetSegmentPath() string {
    return "software-inventory"
}

func (softwareInventory *Install_SoftwareInventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "committed" {
        return &softwareInventory.Committed
    }
    if childYangName == "inactive" {
        return &softwareInventory.Inactive
    }
    if childYangName == "requests" {
        return &softwareInventory.Requests
    }
    if childYangName == "active" {
        return &softwareInventory.Active
    }
    return nil
}

func (softwareInventory *Install_SoftwareInventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["committed"] = &softwareInventory.Committed
    children["inactive"] = &softwareInventory.Inactive
    children["requests"] = &softwareInventory.Requests
    children["active"] = &softwareInventory.Active
    return children
}

func (softwareInventory *Install_SoftwareInventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (softwareInventory *Install_SoftwareInventory) GetBundleName() string { return "cisco_ios_xr" }

func (softwareInventory *Install_SoftwareInventory) GetYangName() string { return "software-inventory" }

func (softwareInventory *Install_SoftwareInventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (softwareInventory *Install_SoftwareInventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (softwareInventory *Install_SoftwareInventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (softwareInventory *Install_SoftwareInventory) SetParent(parent types.Entity) { softwareInventory.parent = parent }

func (softwareInventory *Install_SoftwareInventory) GetParent() types.Entity { return softwareInventory.parent }

func (softwareInventory *Install_SoftwareInventory) GetParentYangName() string { return "install" }

// Install_SoftwareInventory_Committed
// Committed inventory information
type Install_SoftwareInventory_Committed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summarized inventory information.
    Summary Install_SoftwareInventory_Committed_Summary

    // Software inventory.
    Inventories Install_SoftwareInventory_Committed_Inventories
}

func (committed *Install_SoftwareInventory_Committed) GetFilter() yfilter.YFilter { return committed.YFilter }

func (committed *Install_SoftwareInventory_Committed) SetFilter(yf yfilter.YFilter) { committed.YFilter = yf }

func (committed *Install_SoftwareInventory_Committed) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    if yname == "inventories" { return "Inventories" }
    return ""
}

func (committed *Install_SoftwareInventory_Committed) GetSegmentPath() string {
    return "committed"
}

func (committed *Install_SoftwareInventory_Committed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &committed.Summary
    }
    if childYangName == "inventories" {
        return &committed.Inventories
    }
    return nil
}

func (committed *Install_SoftwareInventory_Committed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &committed.Summary
    children["inventories"] = &committed.Inventories
    return children
}

func (committed *Install_SoftwareInventory_Committed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (committed *Install_SoftwareInventory_Committed) GetBundleName() string { return "cisco_ios_xr" }

func (committed *Install_SoftwareInventory_Committed) GetYangName() string { return "committed" }

func (committed *Install_SoftwareInventory_Committed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (committed *Install_SoftwareInventory_Committed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (committed *Install_SoftwareInventory_Committed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (committed *Install_SoftwareInventory_Committed) SetParent(parent types.Entity) { committed.parent = parent }

func (committed *Install_SoftwareInventory_Committed) GetParent() types.Entity { return committed.parent }

func (committed *Install_SoftwareInventory_Committed) GetParentYangName() string { return "software-inventory" }

// Install_SoftwareInventory_Committed_Summary
// Summarized inventory information
type Install_SoftwareInventory_Committed_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default load path.
    DefaultLoadPath Install_SoftwareInventory_Committed_Summary_DefaultLoadPath

    // Admin Resources load path.
    AdminLoadPath Install_SoftwareInventory_Committed_Summary_AdminLoadPath

    // SDR load paths. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_SdrLoadPath.
    SdrLoadPath []Install_SoftwareInventory_Committed_Summary_SdrLoadPath

    // Location load paths. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_LocationLoadPath.
    LocationLoadPath []Install_SoftwareInventory_Committed_Summary_LocationLoadPath
}

func (summary *Install_SoftwareInventory_Committed_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Install_SoftwareInventory_Committed_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Install_SoftwareInventory_Committed_Summary) GetGoName(yname string) string {
    if yname == "default-load-path" { return "DefaultLoadPath" }
    if yname == "admin-load-path" { return "AdminLoadPath" }
    if yname == "sdr-load-path" { return "SdrLoadPath" }
    if yname == "location-load-path" { return "LocationLoadPath" }
    return ""
}

func (summary *Install_SoftwareInventory_Committed_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Install_SoftwareInventory_Committed_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-load-path" {
        return &summary.DefaultLoadPath
    }
    if childYangName == "admin-load-path" {
        return &summary.AdminLoadPath
    }
    if childYangName == "sdr-load-path" {
        for _, c := range summary.SdrLoadPath {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_SdrLoadPath{}
        summary.SdrLoadPath = append(summary.SdrLoadPath, child)
        return &summary.SdrLoadPath[len(summary.SdrLoadPath)-1]
    }
    if childYangName == "location-load-path" {
        for _, c := range summary.LocationLoadPath {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_LocationLoadPath{}
        summary.LocationLoadPath = append(summary.LocationLoadPath, child)
        return &summary.LocationLoadPath[len(summary.LocationLoadPath)-1]
    }
    return nil
}

func (summary *Install_SoftwareInventory_Committed_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-load-path"] = &summary.DefaultLoadPath
    children["admin-load-path"] = &summary.AdminLoadPath
    for i := range summary.SdrLoadPath {
        children[summary.SdrLoadPath[i].GetSegmentPath()] = &summary.SdrLoadPath[i]
    }
    for i := range summary.LocationLoadPath {
        children[summary.LocationLoadPath[i].GetSegmentPath()] = &summary.LocationLoadPath[i]
    }
    return children
}

func (summary *Install_SoftwareInventory_Committed_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summary *Install_SoftwareInventory_Committed_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Install_SoftwareInventory_Committed_Summary) GetYangName() string { return "summary" }

func (summary *Install_SoftwareInventory_Committed_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Install_SoftwareInventory_Committed_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Install_SoftwareInventory_Committed_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Install_SoftwareInventory_Committed_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Install_SoftwareInventory_Committed_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Install_SoftwareInventory_Committed_Summary) GetParentYangName() string { return "committed" }

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath
// Default load path
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath struct {
    parent types.Entity
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
    LoadPath []Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath
}

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetFilter() yfilter.YFilter { return defaultLoadPath.YFilter }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) SetFilter(yf yfilter.YFilter) { defaultLoadPath.YFilter = yf }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "admin-match" { return "AdminMatch" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetSegmentPath() string {
    return "default-load-path"
}

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range defaultLoadPath.LoadPath {
            if defaultLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath{}
        defaultLoadPath.LoadPath = append(defaultLoadPath.LoadPath, child)
        return &defaultLoadPath.LoadPath[len(defaultLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range defaultLoadPath.StandbyLoadPath {
            if defaultLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath{}
        defaultLoadPath.StandbyLoadPath = append(defaultLoadPath.StandbyLoadPath, child)
        return &defaultLoadPath.StandbyLoadPath[len(defaultLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defaultLoadPath.LoadPath {
        children[defaultLoadPath.LoadPath[i].GetSegmentPath()] = &defaultLoadPath.LoadPath[i]
    }
    for i := range defaultLoadPath.StandbyLoadPath {
        children[defaultLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &defaultLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = defaultLoadPath.RequestId
    leafs["admin-match"] = defaultLoadPath.AdminMatch
    leafs["secure-domain-router-name"] = defaultLoadPath.SecureDomainRouterName
    return leafs
}

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetYangName() string { return "default-load-path" }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) SetParent(parent types.Entity) { defaultLoadPath.parent = parent }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetParent() types.Entity { return defaultLoadPath.parent }

func (defaultLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath
// Default load path
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath) GetParentYangName() string { return "default-load-path" }

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath) GetParentYangName() string { return "default-load-path" }

// Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath
// Admin Resources load path
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Admin Resources load path. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath.
    LoadPath []Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath
}

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetFilter() yfilter.YFilter { return adminLoadPath.YFilter }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) SetFilter(yf yfilter.YFilter) { adminLoadPath.YFilter = yf }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetSegmentPath() string {
    return "admin-load-path"
}

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range adminLoadPath.LoadPath {
            if adminLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath{}
        adminLoadPath.LoadPath = append(adminLoadPath.LoadPath, child)
        return &adminLoadPath.LoadPath[len(adminLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range adminLoadPath.StandbyLoadPath {
            if adminLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath{}
        adminLoadPath.StandbyLoadPath = append(adminLoadPath.StandbyLoadPath, child)
        return &adminLoadPath.StandbyLoadPath[len(adminLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range adminLoadPath.LoadPath {
        children[adminLoadPath.LoadPath[i].GetSegmentPath()] = &adminLoadPath.LoadPath[i]
    }
    for i := range adminLoadPath.StandbyLoadPath {
        children[adminLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &adminLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = adminLoadPath.RequestId
    return leafs
}

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetYangName() string { return "admin-load-path" }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) SetParent(parent types.Entity) { adminLoadPath.parent = parent }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetParent() types.Entity { return adminLoadPath.parent }

func (adminLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath
// Admin Resources load path
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath) GetParentYangName() string { return "admin-load-path" }

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath) GetParentYangName() string { return "admin-load-path" }

// Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Summary_AdminLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath
// SDR load paths
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath.
    LoadPath []Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath
}

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetFilter() yfilter.YFilter { return sdrLoadPath.YFilter }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) SetFilter(yf yfilter.YFilter) { sdrLoadPath.YFilter = yf }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetSegmentPath() string {
    return "sdr-load-path"
}

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range sdrLoadPath.LoadPath {
            if sdrLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath{}
        sdrLoadPath.LoadPath = append(sdrLoadPath.LoadPath, child)
        return &sdrLoadPath.LoadPath[len(sdrLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range sdrLoadPath.StandbyLoadPath {
            if sdrLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath{}
        sdrLoadPath.StandbyLoadPath = append(sdrLoadPath.StandbyLoadPath, child)
        return &sdrLoadPath.StandbyLoadPath[len(sdrLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sdrLoadPath.LoadPath {
        children[sdrLoadPath.LoadPath[i].GetSegmentPath()] = &sdrLoadPath.LoadPath[i]
    }
    for i := range sdrLoadPath.StandbyLoadPath {
        children[sdrLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &sdrLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = sdrLoadPath.RequestId
    leafs["secure-domain-router-name"] = sdrLoadPath.SecureDomainRouterName
    return leafs
}

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetYangName() string { return "sdr-load-path" }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) SetParent(parent types.Entity) { sdrLoadPath.parent = parent }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetParent() types.Entity { return sdrLoadPath.parent }

func (sdrLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath) GetParentYangName() string { return "sdr-load-path" }

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath) GetParentYangName() string { return "sdr-load-path" }

// Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Summary_SdrLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath
// Location load paths
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    LoadPath []Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath
}

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetFilter() yfilter.YFilter { return locationLoadPath.YFilter }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) SetFilter(yf yfilter.YFilter) { locationLoadPath.YFilter = yf }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "node-name" { return "NodeName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetSegmentPath() string {
    return "location-load-path"
}

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range locationLoadPath.LoadPath {
            if locationLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath{}
        locationLoadPath.LoadPath = append(locationLoadPath.LoadPath, child)
        return &locationLoadPath.LoadPath[len(locationLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range locationLoadPath.StandbyLoadPath {
            if locationLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath{}
        locationLoadPath.StandbyLoadPath = append(locationLoadPath.StandbyLoadPath, child)
        return &locationLoadPath.StandbyLoadPath[len(locationLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locationLoadPath.LoadPath {
        children[locationLoadPath.LoadPath[i].GetSegmentPath()] = &locationLoadPath.LoadPath[i]
    }
    for i := range locationLoadPath.StandbyLoadPath {
        children[locationLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &locationLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = locationLoadPath.RequestId
    leafs["secure-domain-router-name"] = locationLoadPath.SecureDomainRouterName
    leafs["node-name"] = locationLoadPath.NodeName
    return leafs
}

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetYangName() string { return "location-load-path" }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) SetParent(parent types.Entity) { locationLoadPath.parent = parent }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetParent() types.Entity { return locationLoadPath.parent }

func (locationLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath) GetParentYangName() string { return "location-load-path" }

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath) GetParentYangName() string { return "location-load-path" }

// Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Summary_LocationLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Committed_Inventories
// Software inventory
type Install_SoftwareInventory_Committed_Inventories struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inventory information for specific node. The type is slice of
    // Install_SoftwareInventory_Committed_Inventories_Inventory.
    Inventory []Install_SoftwareInventory_Committed_Inventories_Inventory
}

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetFilter() yfilter.YFilter { return inventories.YFilter }

func (inventories *Install_SoftwareInventory_Committed_Inventories) SetFilter(yf yfilter.YFilter) { inventories.YFilter = yf }

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetGoName(yname string) string {
    if yname == "inventory" { return "Inventory" }
    return ""
}

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetSegmentPath() string {
    return "inventories"
}

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inventory" {
        for _, c := range inventories.Inventory {
            if inventories.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Inventories_Inventory{}
        inventories.Inventory = append(inventories.Inventory, child)
        return &inventories.Inventory[len(inventories.Inventory)-1]
    }
    return nil
}

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventories.Inventory {
        children[inventories.Inventory[i].GetSegmentPath()] = &inventories.Inventory[i]
    }
    return children
}

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetBundleName() string { return "cisco_ios_xr" }

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetYangName() string { return "inventories" }

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventories *Install_SoftwareInventory_Committed_Inventories) SetParent(parent types.Entity) { inventories.parent = parent }

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetParent() types.Entity { return inventories.parent }

func (inventories *Install_SoftwareInventory_Committed_Inventories) GetParentYangName() string { return "committed" }

// Install_SoftwareInventory_Committed_Inventories_Inventory
// Inventory information for specific node
type Install_SoftwareInventory_Committed_Inventories_Inventory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    LoadPath []Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath
}

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetFilter() yfilter.YFilter { return inventory.YFilter }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) SetFilter(yf yfilter.YFilter) { inventory.YFilter = yf }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "major" { return "Major" }
    if yname == "minor" { return "Minor" }
    if yname == "boot-image-name" { return "BootImageName" }
    if yname == "node-type" { return "NodeType" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    return ""
}

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetSegmentPath() string {
    return "inventory" + "[node-name='" + fmt.Sprintf("%v", inventory.NodeName) + "']"
}

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range inventory.LoadPath {
            if inventory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath{}
        inventory.LoadPath = append(inventory.LoadPath, child)
        return &inventory.LoadPath[len(inventory.LoadPath)-1]
    }
    return nil
}

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventory.LoadPath {
        children[inventory.LoadPath[i].GetSegmentPath()] = &inventory.LoadPath[i]
    }
    return children
}

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = inventory.NodeName
    leafs["major"] = inventory.Major
    leafs["minor"] = inventory.Minor
    leafs["boot-image-name"] = inventory.BootImageName
    leafs["node-type"] = inventory.NodeType
    leafs["secure-domain-router-name"] = inventory.SecureDomainRouterName
    return leafs
}

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetBundleName() string { return "cisco_ios_xr" }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetYangName() string { return "inventory" }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) SetParent(parent types.Entity) { inventory.parent = parent }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetParent() types.Entity { return inventory.parent }

func (inventory *Install_SoftwareInventory_Committed_Inventories_Inventory) GetParentYangName() string { return "inventories" }

// Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath
// Load path
type Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath) GetParentYangName() string { return "inventory" }

// Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package
// Package
type Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Committed_Inventories_Inventory_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Inactive
// Inactive inventory information
type Install_SoftwareInventory_Inactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summarized inventory information.
    Summary Install_SoftwareInventory_Inactive_Summary

    // Software inventory.
    Inventories Install_SoftwareInventory_Inactive_Inventories
}

func (inactive *Install_SoftwareInventory_Inactive) GetFilter() yfilter.YFilter { return inactive.YFilter }

func (inactive *Install_SoftwareInventory_Inactive) SetFilter(yf yfilter.YFilter) { inactive.YFilter = yf }

func (inactive *Install_SoftwareInventory_Inactive) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    if yname == "inventories" { return "Inventories" }
    return ""
}

func (inactive *Install_SoftwareInventory_Inactive) GetSegmentPath() string {
    return "inactive"
}

func (inactive *Install_SoftwareInventory_Inactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &inactive.Summary
    }
    if childYangName == "inventories" {
        return &inactive.Inventories
    }
    return nil
}

func (inactive *Install_SoftwareInventory_Inactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &inactive.Summary
    children["inventories"] = &inactive.Inventories
    return children
}

func (inactive *Install_SoftwareInventory_Inactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inactive *Install_SoftwareInventory_Inactive) GetBundleName() string { return "cisco_ios_xr" }

func (inactive *Install_SoftwareInventory_Inactive) GetYangName() string { return "inactive" }

func (inactive *Install_SoftwareInventory_Inactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactive *Install_SoftwareInventory_Inactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactive *Install_SoftwareInventory_Inactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactive *Install_SoftwareInventory_Inactive) SetParent(parent types.Entity) { inactive.parent = parent }

func (inactive *Install_SoftwareInventory_Inactive) GetParent() types.Entity { return inactive.parent }

func (inactive *Install_SoftwareInventory_Inactive) GetParentYangName() string { return "software-inventory" }

// Install_SoftwareInventory_Inactive_Summary
// Summarized inventory information
type Install_SoftwareInventory_Inactive_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default load path.
    DefaultLoadPath Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath

    // Admin Resources load path.
    AdminLoadPath Install_SoftwareInventory_Inactive_Summary_AdminLoadPath

    // SDR load paths. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_SdrLoadPath.
    SdrLoadPath []Install_SoftwareInventory_Inactive_Summary_SdrLoadPath

    // Location load paths. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_LocationLoadPath.
    LocationLoadPath []Install_SoftwareInventory_Inactive_Summary_LocationLoadPath
}

func (summary *Install_SoftwareInventory_Inactive_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Install_SoftwareInventory_Inactive_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Install_SoftwareInventory_Inactive_Summary) GetGoName(yname string) string {
    if yname == "default-load-path" { return "DefaultLoadPath" }
    if yname == "admin-load-path" { return "AdminLoadPath" }
    if yname == "sdr-load-path" { return "SdrLoadPath" }
    if yname == "location-load-path" { return "LocationLoadPath" }
    return ""
}

func (summary *Install_SoftwareInventory_Inactive_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Install_SoftwareInventory_Inactive_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-load-path" {
        return &summary.DefaultLoadPath
    }
    if childYangName == "admin-load-path" {
        return &summary.AdminLoadPath
    }
    if childYangName == "sdr-load-path" {
        for _, c := range summary.SdrLoadPath {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_SdrLoadPath{}
        summary.SdrLoadPath = append(summary.SdrLoadPath, child)
        return &summary.SdrLoadPath[len(summary.SdrLoadPath)-1]
    }
    if childYangName == "location-load-path" {
        for _, c := range summary.LocationLoadPath {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_LocationLoadPath{}
        summary.LocationLoadPath = append(summary.LocationLoadPath, child)
        return &summary.LocationLoadPath[len(summary.LocationLoadPath)-1]
    }
    return nil
}

func (summary *Install_SoftwareInventory_Inactive_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-load-path"] = &summary.DefaultLoadPath
    children["admin-load-path"] = &summary.AdminLoadPath
    for i := range summary.SdrLoadPath {
        children[summary.SdrLoadPath[i].GetSegmentPath()] = &summary.SdrLoadPath[i]
    }
    for i := range summary.LocationLoadPath {
        children[summary.LocationLoadPath[i].GetSegmentPath()] = &summary.LocationLoadPath[i]
    }
    return children
}

func (summary *Install_SoftwareInventory_Inactive_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summary *Install_SoftwareInventory_Inactive_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Install_SoftwareInventory_Inactive_Summary) GetYangName() string { return "summary" }

func (summary *Install_SoftwareInventory_Inactive_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Install_SoftwareInventory_Inactive_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Install_SoftwareInventory_Inactive_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Install_SoftwareInventory_Inactive_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Install_SoftwareInventory_Inactive_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Install_SoftwareInventory_Inactive_Summary) GetParentYangName() string { return "inactive" }

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath
// Default load path
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath struct {
    parent types.Entity
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
    LoadPath []Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath
}

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetFilter() yfilter.YFilter { return defaultLoadPath.YFilter }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) SetFilter(yf yfilter.YFilter) { defaultLoadPath.YFilter = yf }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "admin-match" { return "AdminMatch" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetSegmentPath() string {
    return "default-load-path"
}

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range defaultLoadPath.LoadPath {
            if defaultLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath{}
        defaultLoadPath.LoadPath = append(defaultLoadPath.LoadPath, child)
        return &defaultLoadPath.LoadPath[len(defaultLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range defaultLoadPath.StandbyLoadPath {
            if defaultLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath{}
        defaultLoadPath.StandbyLoadPath = append(defaultLoadPath.StandbyLoadPath, child)
        return &defaultLoadPath.StandbyLoadPath[len(defaultLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defaultLoadPath.LoadPath {
        children[defaultLoadPath.LoadPath[i].GetSegmentPath()] = &defaultLoadPath.LoadPath[i]
    }
    for i := range defaultLoadPath.StandbyLoadPath {
        children[defaultLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &defaultLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = defaultLoadPath.RequestId
    leafs["admin-match"] = defaultLoadPath.AdminMatch
    leafs["secure-domain-router-name"] = defaultLoadPath.SecureDomainRouterName
    return leafs
}

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetYangName() string { return "default-load-path" }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) SetParent(parent types.Entity) { defaultLoadPath.parent = parent }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetParent() types.Entity { return defaultLoadPath.parent }

func (defaultLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath
// Default load path
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath) GetParentYangName() string { return "default-load-path" }

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath) GetParentYangName() string { return "default-load-path" }

// Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath
// Admin Resources load path
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Admin Resources load path. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath.
    LoadPath []Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath
}

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetFilter() yfilter.YFilter { return adminLoadPath.YFilter }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) SetFilter(yf yfilter.YFilter) { adminLoadPath.YFilter = yf }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetSegmentPath() string {
    return "admin-load-path"
}

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range adminLoadPath.LoadPath {
            if adminLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath{}
        adminLoadPath.LoadPath = append(adminLoadPath.LoadPath, child)
        return &adminLoadPath.LoadPath[len(adminLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range adminLoadPath.StandbyLoadPath {
            if adminLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath{}
        adminLoadPath.StandbyLoadPath = append(adminLoadPath.StandbyLoadPath, child)
        return &adminLoadPath.StandbyLoadPath[len(adminLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range adminLoadPath.LoadPath {
        children[adminLoadPath.LoadPath[i].GetSegmentPath()] = &adminLoadPath.LoadPath[i]
    }
    for i := range adminLoadPath.StandbyLoadPath {
        children[adminLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &adminLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = adminLoadPath.RequestId
    return leafs
}

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetYangName() string { return "admin-load-path" }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) SetParent(parent types.Entity) { adminLoadPath.parent = parent }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetParent() types.Entity { return adminLoadPath.parent }

func (adminLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath
// Admin Resources load path
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath) GetParentYangName() string { return "admin-load-path" }

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath) GetParentYangName() string { return "admin-load-path" }

// Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Summary_AdminLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath
// SDR load paths
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath.
    LoadPath []Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath
}

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetFilter() yfilter.YFilter { return sdrLoadPath.YFilter }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) SetFilter(yf yfilter.YFilter) { sdrLoadPath.YFilter = yf }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetSegmentPath() string {
    return "sdr-load-path"
}

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range sdrLoadPath.LoadPath {
            if sdrLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath{}
        sdrLoadPath.LoadPath = append(sdrLoadPath.LoadPath, child)
        return &sdrLoadPath.LoadPath[len(sdrLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range sdrLoadPath.StandbyLoadPath {
            if sdrLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath{}
        sdrLoadPath.StandbyLoadPath = append(sdrLoadPath.StandbyLoadPath, child)
        return &sdrLoadPath.StandbyLoadPath[len(sdrLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sdrLoadPath.LoadPath {
        children[sdrLoadPath.LoadPath[i].GetSegmentPath()] = &sdrLoadPath.LoadPath[i]
    }
    for i := range sdrLoadPath.StandbyLoadPath {
        children[sdrLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &sdrLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = sdrLoadPath.RequestId
    leafs["secure-domain-router-name"] = sdrLoadPath.SecureDomainRouterName
    return leafs
}

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetYangName() string { return "sdr-load-path" }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) SetParent(parent types.Entity) { sdrLoadPath.parent = parent }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetParent() types.Entity { return sdrLoadPath.parent }

func (sdrLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath) GetParentYangName() string { return "sdr-load-path" }

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath) GetParentYangName() string { return "sdr-load-path" }

// Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Summary_SdrLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath
// Location load paths
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    LoadPath []Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath
}

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetFilter() yfilter.YFilter { return locationLoadPath.YFilter }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) SetFilter(yf yfilter.YFilter) { locationLoadPath.YFilter = yf }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "node-name" { return "NodeName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetSegmentPath() string {
    return "location-load-path"
}

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range locationLoadPath.LoadPath {
            if locationLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath{}
        locationLoadPath.LoadPath = append(locationLoadPath.LoadPath, child)
        return &locationLoadPath.LoadPath[len(locationLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range locationLoadPath.StandbyLoadPath {
            if locationLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath{}
        locationLoadPath.StandbyLoadPath = append(locationLoadPath.StandbyLoadPath, child)
        return &locationLoadPath.StandbyLoadPath[len(locationLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locationLoadPath.LoadPath {
        children[locationLoadPath.LoadPath[i].GetSegmentPath()] = &locationLoadPath.LoadPath[i]
    }
    for i := range locationLoadPath.StandbyLoadPath {
        children[locationLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &locationLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = locationLoadPath.RequestId
    leafs["secure-domain-router-name"] = locationLoadPath.SecureDomainRouterName
    leafs["node-name"] = locationLoadPath.NodeName
    return leafs
}

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetYangName() string { return "location-load-path" }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) SetParent(parent types.Entity) { locationLoadPath.parent = parent }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetParent() types.Entity { return locationLoadPath.parent }

func (locationLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath) GetParentYangName() string { return "location-load-path" }

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath) GetParentYangName() string { return "location-load-path" }

// Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Summary_LocationLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Inactive_Inventories
// Software inventory
type Install_SoftwareInventory_Inactive_Inventories struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inventory information for specific node. The type is slice of
    // Install_SoftwareInventory_Inactive_Inventories_Inventory.
    Inventory []Install_SoftwareInventory_Inactive_Inventories_Inventory
}

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetFilter() yfilter.YFilter { return inventories.YFilter }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) SetFilter(yf yfilter.YFilter) { inventories.YFilter = yf }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetGoName(yname string) string {
    if yname == "inventory" { return "Inventory" }
    return ""
}

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetSegmentPath() string {
    return "inventories"
}

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inventory" {
        for _, c := range inventories.Inventory {
            if inventories.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Inventories_Inventory{}
        inventories.Inventory = append(inventories.Inventory, child)
        return &inventories.Inventory[len(inventories.Inventory)-1]
    }
    return nil
}

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventories.Inventory {
        children[inventories.Inventory[i].GetSegmentPath()] = &inventories.Inventory[i]
    }
    return children
}

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetBundleName() string { return "cisco_ios_xr" }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetYangName() string { return "inventories" }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) SetParent(parent types.Entity) { inventories.parent = parent }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetParent() types.Entity { return inventories.parent }

func (inventories *Install_SoftwareInventory_Inactive_Inventories) GetParentYangName() string { return "inactive" }

// Install_SoftwareInventory_Inactive_Inventories_Inventory
// Inventory information for specific node
type Install_SoftwareInventory_Inactive_Inventories_Inventory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    LoadPath []Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath
}

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetFilter() yfilter.YFilter { return inventory.YFilter }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) SetFilter(yf yfilter.YFilter) { inventory.YFilter = yf }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "major" { return "Major" }
    if yname == "minor" { return "Minor" }
    if yname == "boot-image-name" { return "BootImageName" }
    if yname == "node-type" { return "NodeType" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    return ""
}

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetSegmentPath() string {
    return "inventory" + "[node-name='" + fmt.Sprintf("%v", inventory.NodeName) + "']"
}

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range inventory.LoadPath {
            if inventory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath{}
        inventory.LoadPath = append(inventory.LoadPath, child)
        return &inventory.LoadPath[len(inventory.LoadPath)-1]
    }
    return nil
}

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventory.LoadPath {
        children[inventory.LoadPath[i].GetSegmentPath()] = &inventory.LoadPath[i]
    }
    return children
}

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = inventory.NodeName
    leafs["major"] = inventory.Major
    leafs["minor"] = inventory.Minor
    leafs["boot-image-name"] = inventory.BootImageName
    leafs["node-type"] = inventory.NodeType
    leafs["secure-domain-router-name"] = inventory.SecureDomainRouterName
    return leafs
}

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetBundleName() string { return "cisco_ios_xr" }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetYangName() string { return "inventory" }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) SetParent(parent types.Entity) { inventory.parent = parent }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetParent() types.Entity { return inventory.parent }

func (inventory *Install_SoftwareInventory_Inactive_Inventories_Inventory) GetParentYangName() string { return "inventories" }

// Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath
// Load path
type Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath) GetParentYangName() string { return "inventory" }

// Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package
// Package
type Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Inactive_Inventories_Inventory_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Requests
// Install operation requests
type Install_SoftwareInventory_Requests struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install operation request history.
    Requests Install_SoftwareInventory_Requests_Requests
}

func (requests *Install_SoftwareInventory_Requests) GetFilter() yfilter.YFilter { return requests.YFilter }

func (requests *Install_SoftwareInventory_Requests) SetFilter(yf yfilter.YFilter) { requests.YFilter = yf }

func (requests *Install_SoftwareInventory_Requests) GetGoName(yname string) string {
    if yname == "requests" { return "Requests" }
    return ""
}

func (requests *Install_SoftwareInventory_Requests) GetSegmentPath() string {
    return "requests"
}

func (requests *Install_SoftwareInventory_Requests) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "requests" {
        return &requests.Requests
    }
    return nil
}

func (requests *Install_SoftwareInventory_Requests) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["requests"] = &requests.Requests
    return children
}

func (requests *Install_SoftwareInventory_Requests) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (requests *Install_SoftwareInventory_Requests) GetBundleName() string { return "cisco_ios_xr" }

func (requests *Install_SoftwareInventory_Requests) GetYangName() string { return "requests" }

func (requests *Install_SoftwareInventory_Requests) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requests *Install_SoftwareInventory_Requests) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requests *Install_SoftwareInventory_Requests) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requests *Install_SoftwareInventory_Requests) SetParent(parent types.Entity) { requests.parent = parent }

func (requests *Install_SoftwareInventory_Requests) GetParent() types.Entity { return requests.parent }

func (requests *Install_SoftwareInventory_Requests) GetParentYangName() string { return "software-inventory" }

// Install_SoftwareInventory_Requests_Requests
// Install operation request history
type Install_SoftwareInventory_Requests_Requests struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install operation request information. The type is slice of
    // Install_SoftwareInventory_Requests_Requests_Request.
    Request []Install_SoftwareInventory_Requests_Requests_Request
}

func (requests *Install_SoftwareInventory_Requests_Requests) GetFilter() yfilter.YFilter { return requests.YFilter }

func (requests *Install_SoftwareInventory_Requests_Requests) SetFilter(yf yfilter.YFilter) { requests.YFilter = yf }

func (requests *Install_SoftwareInventory_Requests_Requests) GetGoName(yname string) string {
    if yname == "request" { return "Request" }
    return ""
}

func (requests *Install_SoftwareInventory_Requests_Requests) GetSegmentPath() string {
    return "requests"
}

func (requests *Install_SoftwareInventory_Requests_Requests) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "request" {
        for _, c := range requests.Request {
            if requests.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Requests_Requests_Request{}
        requests.Request = append(requests.Request, child)
        return &requests.Request[len(requests.Request)-1]
    }
    return nil
}

func (requests *Install_SoftwareInventory_Requests_Requests) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range requests.Request {
        children[requests.Request[i].GetSegmentPath()] = &requests.Request[i]
    }
    return children
}

func (requests *Install_SoftwareInventory_Requests_Requests) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (requests *Install_SoftwareInventory_Requests_Requests) GetBundleName() string { return "cisco_ios_xr" }

func (requests *Install_SoftwareInventory_Requests_Requests) GetYangName() string { return "requests" }

func (requests *Install_SoftwareInventory_Requests_Requests) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requests *Install_SoftwareInventory_Requests_Requests) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requests *Install_SoftwareInventory_Requests_Requests) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requests *Install_SoftwareInventory_Requests_Requests) SetParent(parent types.Entity) { requests.parent = parent }

func (requests *Install_SoftwareInventory_Requests_Requests) GetParent() types.Entity { return requests.parent }

func (requests *Install_SoftwareInventory_Requests_Requests) GetParentYangName() string { return "requests" }

// Install_SoftwareInventory_Requests_Requests_Request
// Install operation request information
type Install_SoftwareInventory_Requests_Requests_Request struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Install operation request ID. The type is
    // interface{} with range: -2147483648..2147483647.
    RequestId interface{}

    // Inventory information of install operation request.
    Inventories Install_SoftwareInventory_Requests_Requests_Request_Inventories
}

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Install_SoftwareInventory_Requests_Requests_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "inventories" { return "Inventories" }
    return ""
}

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetSegmentPath() string {
    return "request" + "[request-id='" + fmt.Sprintf("%v", request.RequestId) + "']"
}

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inventories" {
        return &request.Inventories
    }
    return nil
}

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["inventories"] = &request.Inventories
    return children
}

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = request.RequestId
    return leafs
}

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetYangName() string { return "request" }

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Install_SoftwareInventory_Requests_Requests_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetParent() types.Entity { return request.parent }

func (request *Install_SoftwareInventory_Requests_Requests_Request) GetParentYangName() string { return "requests" }

// Install_SoftwareInventory_Requests_Requests_Request_Inventories
// Inventory information of install operation
// request
type Install_SoftwareInventory_Requests_Requests_Request_Inventories struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inventory information. The type is slice of
    // Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory.
    Inventory []Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory
}

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetFilter() yfilter.YFilter { return inventories.YFilter }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) SetFilter(yf yfilter.YFilter) { inventories.YFilter = yf }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetGoName(yname string) string {
    if yname == "inventory" { return "Inventory" }
    return ""
}

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetSegmentPath() string {
    return "inventories"
}

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inventory" {
        for _, c := range inventories.Inventory {
            if inventories.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory{}
        inventories.Inventory = append(inventories.Inventory, child)
        return &inventories.Inventory[len(inventories.Inventory)-1]
    }
    return nil
}

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventories.Inventory {
        children[inventories.Inventory[i].GetSegmentPath()] = &inventories.Inventory[i]
    }
    return children
}

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetBundleName() string { return "cisco_ios_xr" }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetYangName() string { return "inventories" }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) SetParent(parent types.Entity) { inventories.parent = parent }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetParent() types.Entity { return inventories.parent }

func (inventories *Install_SoftwareInventory_Requests_Requests_Request_Inventories) GetParentYangName() string { return "request" }

// Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory
// Inventory information
type Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    LoadPath []Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath
}

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetFilter() yfilter.YFilter { return inventory.YFilter }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) SetFilter(yf yfilter.YFilter) { inventory.YFilter = yf }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "major" { return "Major" }
    if yname == "minor" { return "Minor" }
    if yname == "boot-image-name" { return "BootImageName" }
    if yname == "node-type" { return "NodeType" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    return ""
}

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetSegmentPath() string {
    return "inventory" + "[node-name='" + fmt.Sprintf("%v", inventory.NodeName) + "']"
}

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range inventory.LoadPath {
            if inventory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath{}
        inventory.LoadPath = append(inventory.LoadPath, child)
        return &inventory.LoadPath[len(inventory.LoadPath)-1]
    }
    return nil
}

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventory.LoadPath {
        children[inventory.LoadPath[i].GetSegmentPath()] = &inventory.LoadPath[i]
    }
    return children
}

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = inventory.NodeName
    leafs["major"] = inventory.Major
    leafs["minor"] = inventory.Minor
    leafs["boot-image-name"] = inventory.BootImageName
    leafs["node-type"] = inventory.NodeType
    leafs["secure-domain-router-name"] = inventory.SecureDomainRouterName
    return leafs
}

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetBundleName() string { return "cisco_ios_xr" }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetYangName() string { return "inventory" }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) SetParent(parent types.Entity) { inventory.parent = parent }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetParent() types.Entity { return inventory.parent }

func (inventory *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory) GetParentYangName() string { return "inventories" }

// Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath
// Load path
type Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath) GetParentYangName() string { return "inventory" }

// Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package
// Package
type Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Requests_Requests_Request_Inventories_Inventory_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Active
// Active inventory information
type Install_SoftwareInventory_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summarized inventory information.
    Summary Install_SoftwareInventory_Active_Summary

    // Software inventory.
    Inventories Install_SoftwareInventory_Active_Inventories
}

func (active *Install_SoftwareInventory_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *Install_SoftwareInventory_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *Install_SoftwareInventory_Active) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    if yname == "inventories" { return "Inventories" }
    return ""
}

func (active *Install_SoftwareInventory_Active) GetSegmentPath() string {
    return "active"
}

func (active *Install_SoftwareInventory_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &active.Summary
    }
    if childYangName == "inventories" {
        return &active.Inventories
    }
    return nil
}

func (active *Install_SoftwareInventory_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &active.Summary
    children["inventories"] = &active.Inventories
    return children
}

func (active *Install_SoftwareInventory_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (active *Install_SoftwareInventory_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *Install_SoftwareInventory_Active) GetYangName() string { return "active" }

func (active *Install_SoftwareInventory_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *Install_SoftwareInventory_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *Install_SoftwareInventory_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *Install_SoftwareInventory_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *Install_SoftwareInventory_Active) GetParent() types.Entity { return active.parent }

func (active *Install_SoftwareInventory_Active) GetParentYangName() string { return "software-inventory" }

// Install_SoftwareInventory_Active_Summary
// Summarized inventory information
type Install_SoftwareInventory_Active_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default load path.
    DefaultLoadPath Install_SoftwareInventory_Active_Summary_DefaultLoadPath

    // Admin Resources load path.
    AdminLoadPath Install_SoftwareInventory_Active_Summary_AdminLoadPath

    // SDR load paths. The type is slice of
    // Install_SoftwareInventory_Active_Summary_SdrLoadPath.
    SdrLoadPath []Install_SoftwareInventory_Active_Summary_SdrLoadPath

    // Location load paths. The type is slice of
    // Install_SoftwareInventory_Active_Summary_LocationLoadPath.
    LocationLoadPath []Install_SoftwareInventory_Active_Summary_LocationLoadPath
}

func (summary *Install_SoftwareInventory_Active_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Install_SoftwareInventory_Active_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Install_SoftwareInventory_Active_Summary) GetGoName(yname string) string {
    if yname == "default-load-path" { return "DefaultLoadPath" }
    if yname == "admin-load-path" { return "AdminLoadPath" }
    if yname == "sdr-load-path" { return "SdrLoadPath" }
    if yname == "location-load-path" { return "LocationLoadPath" }
    return ""
}

func (summary *Install_SoftwareInventory_Active_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Install_SoftwareInventory_Active_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-load-path" {
        return &summary.DefaultLoadPath
    }
    if childYangName == "admin-load-path" {
        return &summary.AdminLoadPath
    }
    if childYangName == "sdr-load-path" {
        for _, c := range summary.SdrLoadPath {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_SdrLoadPath{}
        summary.SdrLoadPath = append(summary.SdrLoadPath, child)
        return &summary.SdrLoadPath[len(summary.SdrLoadPath)-1]
    }
    if childYangName == "location-load-path" {
        for _, c := range summary.LocationLoadPath {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_LocationLoadPath{}
        summary.LocationLoadPath = append(summary.LocationLoadPath, child)
        return &summary.LocationLoadPath[len(summary.LocationLoadPath)-1]
    }
    return nil
}

func (summary *Install_SoftwareInventory_Active_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-load-path"] = &summary.DefaultLoadPath
    children["admin-load-path"] = &summary.AdminLoadPath
    for i := range summary.SdrLoadPath {
        children[summary.SdrLoadPath[i].GetSegmentPath()] = &summary.SdrLoadPath[i]
    }
    for i := range summary.LocationLoadPath {
        children[summary.LocationLoadPath[i].GetSegmentPath()] = &summary.LocationLoadPath[i]
    }
    return children
}

func (summary *Install_SoftwareInventory_Active_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summary *Install_SoftwareInventory_Active_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Install_SoftwareInventory_Active_Summary) GetYangName() string { return "summary" }

func (summary *Install_SoftwareInventory_Active_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Install_SoftwareInventory_Active_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Install_SoftwareInventory_Active_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Install_SoftwareInventory_Active_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Install_SoftwareInventory_Active_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Install_SoftwareInventory_Active_Summary) GetParentYangName() string { return "active" }

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath
// Default load path
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath struct {
    parent types.Entity
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
    LoadPath []Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath
}

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetFilter() yfilter.YFilter { return defaultLoadPath.YFilter }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) SetFilter(yf yfilter.YFilter) { defaultLoadPath.YFilter = yf }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "admin-match" { return "AdminMatch" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetSegmentPath() string {
    return "default-load-path"
}

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range defaultLoadPath.LoadPath {
            if defaultLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath{}
        defaultLoadPath.LoadPath = append(defaultLoadPath.LoadPath, child)
        return &defaultLoadPath.LoadPath[len(defaultLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range defaultLoadPath.StandbyLoadPath {
            if defaultLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath{}
        defaultLoadPath.StandbyLoadPath = append(defaultLoadPath.StandbyLoadPath, child)
        return &defaultLoadPath.StandbyLoadPath[len(defaultLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defaultLoadPath.LoadPath {
        children[defaultLoadPath.LoadPath[i].GetSegmentPath()] = &defaultLoadPath.LoadPath[i]
    }
    for i := range defaultLoadPath.StandbyLoadPath {
        children[defaultLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &defaultLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = defaultLoadPath.RequestId
    leafs["admin-match"] = defaultLoadPath.AdminMatch
    leafs["secure-domain-router-name"] = defaultLoadPath.SecureDomainRouterName
    return leafs
}

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetYangName() string { return "default-load-path" }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) SetParent(parent types.Entity) { defaultLoadPath.parent = parent }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetParent() types.Entity { return defaultLoadPath.parent }

func (defaultLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath
// Default load path
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath) GetParentYangName() string { return "default-load-path" }

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath) GetParentYangName() string { return "default-load-path" }

// Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Summary_DefaultLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Active_Summary_AdminLoadPath
// Admin Resources load path
type Install_SoftwareInventory_Active_Summary_AdminLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // Admin Resources load path. The type is slice of
    // Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath.
    LoadPath []Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath
}

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetFilter() yfilter.YFilter { return adminLoadPath.YFilter }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) SetFilter(yf yfilter.YFilter) { adminLoadPath.YFilter = yf }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetSegmentPath() string {
    return "admin-load-path"
}

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range adminLoadPath.LoadPath {
            if adminLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath{}
        adminLoadPath.LoadPath = append(adminLoadPath.LoadPath, child)
        return &adminLoadPath.LoadPath[len(adminLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range adminLoadPath.StandbyLoadPath {
            if adminLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath{}
        adminLoadPath.StandbyLoadPath = append(adminLoadPath.StandbyLoadPath, child)
        return &adminLoadPath.StandbyLoadPath[len(adminLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range adminLoadPath.LoadPath {
        children[adminLoadPath.LoadPath[i].GetSegmentPath()] = &adminLoadPath.LoadPath[i]
    }
    for i := range adminLoadPath.StandbyLoadPath {
        children[adminLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &adminLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = adminLoadPath.RequestId
    return leafs
}

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetYangName() string { return "admin-load-path" }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) SetParent(parent types.Entity) { adminLoadPath.parent = parent }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetParent() types.Entity { return adminLoadPath.parent }

func (adminLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath
// Admin Resources load path
type Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath) GetParentYangName() string { return "admin-load-path" }

// Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath) GetParentYangName() string { return "admin-load-path" }

// Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Summary_AdminLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Active_Summary_SdrLoadPath
// SDR load paths
type Install_SoftwareInventory_Active_Summary_SdrLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install op affecting scope. The type is interface{} with range:
    // 0..4294967295.
    RequestId interface{}

    // SDR name. The type is string.
    SecureDomainRouterName interface{}

    // Load path. The type is slice of
    // Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath.
    LoadPath []Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath
}

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetFilter() yfilter.YFilter { return sdrLoadPath.YFilter }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) SetFilter(yf yfilter.YFilter) { sdrLoadPath.YFilter = yf }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetSegmentPath() string {
    return "sdr-load-path"
}

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range sdrLoadPath.LoadPath {
            if sdrLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath{}
        sdrLoadPath.LoadPath = append(sdrLoadPath.LoadPath, child)
        return &sdrLoadPath.LoadPath[len(sdrLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range sdrLoadPath.StandbyLoadPath {
            if sdrLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath{}
        sdrLoadPath.StandbyLoadPath = append(sdrLoadPath.StandbyLoadPath, child)
        return &sdrLoadPath.StandbyLoadPath[len(sdrLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sdrLoadPath.LoadPath {
        children[sdrLoadPath.LoadPath[i].GetSegmentPath()] = &sdrLoadPath.LoadPath[i]
    }
    for i := range sdrLoadPath.StandbyLoadPath {
        children[sdrLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &sdrLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = sdrLoadPath.RequestId
    leafs["secure-domain-router-name"] = sdrLoadPath.SecureDomainRouterName
    return leafs
}

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetYangName() string { return "sdr-load-path" }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) SetParent(parent types.Entity) { sdrLoadPath.parent = parent }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetParent() types.Entity { return sdrLoadPath.parent }

func (sdrLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath) GetParentYangName() string { return "sdr-load-path" }

// Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath) GetParentYangName() string { return "sdr-load-path" }

// Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Summary_SdrLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Active_Summary_LocationLoadPath
// Location load paths
type Install_SoftwareInventory_Active_Summary_LocationLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    LoadPath []Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath

    // Load paths for standby nodes. The type is slice of
    // Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath.
    StandbyLoadPath []Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath
}

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetFilter() yfilter.YFilter { return locationLoadPath.YFilter }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) SetFilter(yf yfilter.YFilter) { locationLoadPath.YFilter = yf }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "node-name" { return "NodeName" }
    if yname == "load-path" { return "LoadPath" }
    if yname == "standby-load-path" { return "StandbyLoadPath" }
    return ""
}

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetSegmentPath() string {
    return "location-load-path"
}

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range locationLoadPath.LoadPath {
            if locationLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath{}
        locationLoadPath.LoadPath = append(locationLoadPath.LoadPath, child)
        return &locationLoadPath.LoadPath[len(locationLoadPath.LoadPath)-1]
    }
    if childYangName == "standby-load-path" {
        for _, c := range locationLoadPath.StandbyLoadPath {
            if locationLoadPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath{}
        locationLoadPath.StandbyLoadPath = append(locationLoadPath.StandbyLoadPath, child)
        return &locationLoadPath.StandbyLoadPath[len(locationLoadPath.StandbyLoadPath)-1]
    }
    return nil
}

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locationLoadPath.LoadPath {
        children[locationLoadPath.LoadPath[i].GetSegmentPath()] = &locationLoadPath.LoadPath[i]
    }
    for i := range locationLoadPath.StandbyLoadPath {
        children[locationLoadPath.StandbyLoadPath[i].GetSegmentPath()] = &locationLoadPath.StandbyLoadPath[i]
    }
    return children
}

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = locationLoadPath.RequestId
    leafs["secure-domain-router-name"] = locationLoadPath.SecureDomainRouterName
    leafs["node-name"] = locationLoadPath.NodeName
    return leafs
}

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetYangName() string { return "location-load-path" }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) SetParent(parent types.Entity) { locationLoadPath.parent = parent }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetParent() types.Entity { return locationLoadPath.parent }

func (locationLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath) GetParentYangName() string { return "summary" }

// Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath
// Load path
type Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath) GetParentYangName() string { return "location-load-path" }

// Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath
// Load paths for standby nodes
type Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetFilter() yfilter.YFilter { return standbyLoadPath.YFilter }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) SetFilter(yf yfilter.YFilter) { standbyLoadPath.YFilter = yf }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetSegmentPath() string {
    return "standby-load-path"
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &standbyLoadPath.Package
    }
    return nil
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &standbyLoadPath.Package
    return children
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = standbyLoadPath.Version
    leafs["build-information"] = standbyLoadPath.BuildInformation
    return leafs
}

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetYangName() string { return "standby-load-path" }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) SetParent(parent types.Entity) { standbyLoadPath.parent = parent }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetParent() types.Entity { return standbyLoadPath.parent }

func (standbyLoadPath *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath) GetParentYangName() string { return "location-load-path" }

// Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package
// Package
type Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Summary_LocationLoadPath_StandbyLoadPath_Package) GetParentYangName() string { return "standby-load-path" }

// Install_SoftwareInventory_Active_Inventories
// Software inventory
type Install_SoftwareInventory_Active_Inventories struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inventory information for specific node. The type is slice of
    // Install_SoftwareInventory_Active_Inventories_Inventory.
    Inventory []Install_SoftwareInventory_Active_Inventories_Inventory
}

func (inventories *Install_SoftwareInventory_Active_Inventories) GetFilter() yfilter.YFilter { return inventories.YFilter }

func (inventories *Install_SoftwareInventory_Active_Inventories) SetFilter(yf yfilter.YFilter) { inventories.YFilter = yf }

func (inventories *Install_SoftwareInventory_Active_Inventories) GetGoName(yname string) string {
    if yname == "inventory" { return "Inventory" }
    return ""
}

func (inventories *Install_SoftwareInventory_Active_Inventories) GetSegmentPath() string {
    return "inventories"
}

func (inventories *Install_SoftwareInventory_Active_Inventories) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inventory" {
        for _, c := range inventories.Inventory {
            if inventories.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Inventories_Inventory{}
        inventories.Inventory = append(inventories.Inventory, child)
        return &inventories.Inventory[len(inventories.Inventory)-1]
    }
    return nil
}

func (inventories *Install_SoftwareInventory_Active_Inventories) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventories.Inventory {
        children[inventories.Inventory[i].GetSegmentPath()] = &inventories.Inventory[i]
    }
    return children
}

func (inventories *Install_SoftwareInventory_Active_Inventories) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inventories *Install_SoftwareInventory_Active_Inventories) GetBundleName() string { return "cisco_ios_xr" }

func (inventories *Install_SoftwareInventory_Active_Inventories) GetYangName() string { return "inventories" }

func (inventories *Install_SoftwareInventory_Active_Inventories) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventories *Install_SoftwareInventory_Active_Inventories) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventories *Install_SoftwareInventory_Active_Inventories) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventories *Install_SoftwareInventory_Active_Inventories) SetParent(parent types.Entity) { inventories.parent = parent }

func (inventories *Install_SoftwareInventory_Active_Inventories) GetParent() types.Entity { return inventories.parent }

func (inventories *Install_SoftwareInventory_Active_Inventories) GetParentYangName() string { return "active" }

// Install_SoftwareInventory_Active_Inventories_Inventory
// Inventory information for specific node
type Install_SoftwareInventory_Active_Inventories_Inventory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    LoadPath []Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath
}

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetFilter() yfilter.YFilter { return inventory.YFilter }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) SetFilter(yf yfilter.YFilter) { inventory.YFilter = yf }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "major" { return "Major" }
    if yname == "minor" { return "Minor" }
    if yname == "boot-image-name" { return "BootImageName" }
    if yname == "node-type" { return "NodeType" }
    if yname == "secure-domain-router-name" { return "SecureDomainRouterName" }
    if yname == "load-path" { return "LoadPath" }
    return ""
}

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetSegmentPath() string {
    return "inventory" + "[node-name='" + fmt.Sprintf("%v", inventory.NodeName) + "']"
}

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-path" {
        for _, c := range inventory.LoadPath {
            if inventory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath{}
        inventory.LoadPath = append(inventory.LoadPath, child)
        return &inventory.LoadPath[len(inventory.LoadPath)-1]
    }
    return nil
}

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inventory.LoadPath {
        children[inventory.LoadPath[i].GetSegmentPath()] = &inventory.LoadPath[i]
    }
    return children
}

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = inventory.NodeName
    leafs["major"] = inventory.Major
    leafs["minor"] = inventory.Minor
    leafs["boot-image-name"] = inventory.BootImageName
    leafs["node-type"] = inventory.NodeType
    leafs["secure-domain-router-name"] = inventory.SecureDomainRouterName
    return leafs
}

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetBundleName() string { return "cisco_ios_xr" }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetYangName() string { return "inventory" }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) SetParent(parent types.Entity) { inventory.parent = parent }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetParent() types.Entity { return inventory.parent }

func (inventory *Install_SoftwareInventory_Active_Inventories_Inventory) GetParentYangName() string { return "inventories" }

// Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath
// Load path
type Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is string.
    Version interface{}

    // Build information. The type is string.
    BuildInformation interface{}

    // Package.
    Package Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package
}

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetFilter() yfilter.YFilter { return loadPath.YFilter }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) SetFilter(yf yfilter.YFilter) { loadPath.YFilter = yf }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "build-information" { return "BuildInformation" }
    if yname == "package" { return "Package" }
    return ""
}

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetSegmentPath() string {
    return "load-path"
}

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        return &loadPath.Package
    }
    return nil
}

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package"] = &loadPath.Package
    return children
}

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = loadPath.Version
    leafs["build-information"] = loadPath.BuildInformation
    return leafs
}

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetBundleName() string { return "cisco_ios_xr" }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetYangName() string { return "load-path" }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) SetParent(parent types.Entity) { loadPath.parent = parent }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetParent() types.Entity { return loadPath.parent }

func (loadPath *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath) GetParentYangName() string { return "inventory" }

// Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package
// Package
type Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Device name. The type is string.
    DeviceName interface{}

    // Package group name. The type is string.
    Name interface{}
}

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetGoName(yname string) string {
    if yname == "device-name" { return "DeviceName" }
    if yname == "name" { return "Name" }
    return ""
}

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetSegmentPath() string {
    return "package"
}

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-name"] = self.DeviceName
    leafs["name"] = self.Name
    return leafs
}

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetYangName() string { return "package" }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetParent() types.Entity { return self.parent }

func (self *Install_SoftwareInventory_Active_Inventories_Inventory_LoadPath_Package) GetParentYangName() string { return "load-path" }

// Install_Issu
// Information of install ISSU operations
type Install_Issu struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISSU manager card inventory table.
    CardInventories Install_Issu_CardInventories

    // Summarized ISSU stage information.
    Stage Install_Issu_Stage
}

func (issu *Install_Issu) GetFilter() yfilter.YFilter { return issu.YFilter }

func (issu *Install_Issu) SetFilter(yf yfilter.YFilter) { issu.YFilter = yf }

func (issu *Install_Issu) GetGoName(yname string) string {
    if yname == "card-inventories" { return "CardInventories" }
    if yname == "stage" { return "Stage" }
    return ""
}

func (issu *Install_Issu) GetSegmentPath() string {
    return "issu"
}

func (issu *Install_Issu) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "card-inventories" {
        return &issu.CardInventories
    }
    if childYangName == "stage" {
        return &issu.Stage
    }
    return nil
}

func (issu *Install_Issu) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["card-inventories"] = &issu.CardInventories
    children["stage"] = &issu.Stage
    return children
}

func (issu *Install_Issu) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (issu *Install_Issu) GetBundleName() string { return "cisco_ios_xr" }

func (issu *Install_Issu) GetYangName() string { return "issu" }

func (issu *Install_Issu) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (issu *Install_Issu) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (issu *Install_Issu) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (issu *Install_Issu) SetParent(parent types.Entity) { issu.parent = parent }

func (issu *Install_Issu) GetParent() types.Entity { return issu.parent }

func (issu *Install_Issu) GetParentYangName() string { return "install" }

// Install_Issu_CardInventories
// ISSU manager card inventory table
type Install_Issu_CardInventories struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISSU manager inventory summary of the same card type. The type is slice of
    // Install_Issu_CardInventories_CardInventory.
    CardInventory []Install_Issu_CardInventories_CardInventory
}

func (cardInventories *Install_Issu_CardInventories) GetFilter() yfilter.YFilter { return cardInventories.YFilter }

func (cardInventories *Install_Issu_CardInventories) SetFilter(yf yfilter.YFilter) { cardInventories.YFilter = yf }

func (cardInventories *Install_Issu_CardInventories) GetGoName(yname string) string {
    if yname == "card-inventory" { return "CardInventory" }
    return ""
}

func (cardInventories *Install_Issu_CardInventories) GetSegmentPath() string {
    return "card-inventories"
}

func (cardInventories *Install_Issu_CardInventories) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "card-inventory" {
        for _, c := range cardInventories.CardInventory {
            if cardInventories.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Issu_CardInventories_CardInventory{}
        cardInventories.CardInventory = append(cardInventories.CardInventory, child)
        return &cardInventories.CardInventory[len(cardInventories.CardInventory)-1]
    }
    return nil
}

func (cardInventories *Install_Issu_CardInventories) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cardInventories.CardInventory {
        children[cardInventories.CardInventory[i].GetSegmentPath()] = &cardInventories.CardInventory[i]
    }
    return children
}

func (cardInventories *Install_Issu_CardInventories) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cardInventories *Install_Issu_CardInventories) GetBundleName() string { return "cisco_ios_xr" }

func (cardInventories *Install_Issu_CardInventories) GetYangName() string { return "card-inventories" }

func (cardInventories *Install_Issu_CardInventories) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardInventories *Install_Issu_CardInventories) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardInventories *Install_Issu_CardInventories) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardInventories *Install_Issu_CardInventories) SetParent(parent types.Entity) { cardInventories.parent = parent }

func (cardInventories *Install_Issu_CardInventories) GetParent() types.Entity { return cardInventories.parent }

func (cardInventories *Install_Issu_CardInventories) GetParentYangName() string { return "issu" }

// Install_Issu_CardInventories_CardInventory
// ISSU manager inventory summary of the same
// card type
type Install_Issu_CardInventories_CardInventory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ISSU manager card type ID. The type is
    // IsmCardTypeFamily.
    CardTypeId interface{}

    // node state for all nodes. The type is slice of
    // Install_Issu_CardInventories_CardInventory_Summary.
    Summary []Install_Issu_CardInventories_CardInventory_Summary
}

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetFilter() yfilter.YFilter { return cardInventory.YFilter }

func (cardInventory *Install_Issu_CardInventories_CardInventory) SetFilter(yf yfilter.YFilter) { cardInventory.YFilter = yf }

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetGoName(yname string) string {
    if yname == "card-type-id" { return "CardTypeId" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetSegmentPath() string {
    return "card-inventory" + "[card-type-id='" + fmt.Sprintf("%v", cardInventory.CardTypeId) + "']"
}

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        for _, c := range cardInventory.Summary {
            if cardInventory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Issu_CardInventories_CardInventory_Summary{}
        cardInventory.Summary = append(cardInventory.Summary, child)
        return &cardInventory.Summary[len(cardInventory.Summary)-1]
    }
    return nil
}

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cardInventory.Summary {
        children[cardInventory.Summary[i].GetSegmentPath()] = &cardInventory.Summary[i]
    }
    return children
}

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-type-id"] = cardInventory.CardTypeId
    return leafs
}

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetBundleName() string { return "cisco_ios_xr" }

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetYangName() string { return "card-inventory" }

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardInventory *Install_Issu_CardInventories_CardInventory) SetParent(parent types.Entity) { cardInventory.parent = parent }

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetParent() types.Entity { return cardInventory.parent }

func (cardInventory *Install_Issu_CardInventories_CardInventory) GetParentYangName() string { return "card-inventories" }

// Install_Issu_CardInventories_CardInventory_Summary
// node state for all nodes
type Install_Issu_CardInventories_CardInventory_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "partner-node-name" { return "PartnerNodeName" }
    if yname == "node-state" { return "NodeState" }
    if yname == "node-role" { return "NodeRole" }
    if yname == "node-type-pi" { return "NodeTypePi" }
    if yname == "node-type-issu" { return "NodeTypeIssu" }
    if yname == "node-current-state" { return "NodeCurrentState" }
    if yname == "node-expected-state" { return "NodeExpectedState" }
    if yname == "node-failure-reason" { return "NodeFailureReason" }
    if yname == "is-conforming-node" { return "IsConformingNode" }
    if yname == "attempts" { return "Attempts" }
    if yname == "is-node-upgraded" { return "IsNodeUpgraded" }
    return ""
}

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = summary.NodeName
    leafs["partner-node-name"] = summary.PartnerNodeName
    leafs["node-state"] = summary.NodeState
    leafs["node-role"] = summary.NodeRole
    leafs["node-type-pi"] = summary.NodeTypePi
    leafs["node-type-issu"] = summary.NodeTypeIssu
    leafs["node-current-state"] = summary.NodeCurrentState
    leafs["node-expected-state"] = summary.NodeExpectedState
    leafs["node-failure-reason"] = summary.NodeFailureReason
    leafs["is-conforming-node"] = summary.IsConformingNode
    leafs["attempts"] = summary.Attempts
    leafs["is-node-upgraded"] = summary.IsNodeUpgraded
    return leafs
}

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetYangName() string { return "summary" }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Install_Issu_CardInventories_CardInventory_Summary) GetParentYangName() string { return "card-inventory" }

// Install_Issu_Stage
// Summarized ISSU stage information
type Install_Issu_Stage struct {
    parent types.Entity
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

func (stage *Install_Issu_Stage) GetFilter() yfilter.YFilter { return stage.YFilter }

func (stage *Install_Issu_Stage) SetFilter(yf yfilter.YFilter) { stage.YFilter = yf }

func (stage *Install_Issu_Stage) GetGoName(yname string) string {
    if yname == "issu-state" { return "IssuState" }
    if yname == "issu-op-id" { return "IssuOpId" }
    if yname == "percentage" { return "Percentage" }
    if yname == "is-issu-aborted" { return "IsIssuAborted" }
    if yname == "is-issu-aborted-by-ism" { return "IsIssuAbortedByIsm" }
    if yname == "issu-manager-fsm-state" { return "IssuManagerFsmState" }
    if yname == "participating-node-all" { return "ParticipatingNodeAll" }
    if yname == "num-nodes-in-progress" { return "NumNodesInProgress" }
    if yname == "num-of-nodes-in-load" { return "NumOfNodesInLoad" }
    if yname == "num-of-nodes-in-run" { return "NumOfNodesInRun" }
    if yname == "numof-nc-nodes" { return "NumofNcNodes" }
    if yname == "node-in-progress" { return "NodeInProgress" }
    if yname == "nodes-in-load" { return "NodesInLoad" }
    if yname == "nodes-in-run" { return "NodesInRun" }
    if yname == "nc-nodes" { return "NcNodes" }
    return ""
}

func (stage *Install_Issu_Stage) GetSegmentPath() string {
    return "stage"
}

func (stage *Install_Issu_Stage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-in-progress" {
        return &stage.NodeInProgress
    }
    if childYangName == "nodes-in-load" {
        return &stage.NodesInLoad
    }
    if childYangName == "nodes-in-run" {
        return &stage.NodesInRun
    }
    if childYangName == "nc-nodes" {
        return &stage.NcNodes
    }
    return nil
}

func (stage *Install_Issu_Stage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-in-progress"] = &stage.NodeInProgress
    children["nodes-in-load"] = &stage.NodesInLoad
    children["nodes-in-run"] = &stage.NodesInRun
    children["nc-nodes"] = &stage.NcNodes
    return children
}

func (stage *Install_Issu_Stage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["issu-state"] = stage.IssuState
    leafs["issu-op-id"] = stage.IssuOpId
    leafs["percentage"] = stage.Percentage
    leafs["is-issu-aborted"] = stage.IsIssuAborted
    leafs["is-issu-aborted-by-ism"] = stage.IsIssuAbortedByIsm
    leafs["issu-manager-fsm-state"] = stage.IssuManagerFsmState
    leafs["participating-node-all"] = stage.ParticipatingNodeAll
    leafs["num-nodes-in-progress"] = stage.NumNodesInProgress
    leafs["num-of-nodes-in-load"] = stage.NumOfNodesInLoad
    leafs["num-of-nodes-in-run"] = stage.NumOfNodesInRun
    leafs["numof-nc-nodes"] = stage.NumofNcNodes
    return leafs
}

func (stage *Install_Issu_Stage) GetBundleName() string { return "cisco_ios_xr" }

func (stage *Install_Issu_Stage) GetYangName() string { return "stage" }

func (stage *Install_Issu_Stage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stage *Install_Issu_Stage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stage *Install_Issu_Stage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stage *Install_Issu_Stage) SetParent(parent types.Entity) { stage.parent = parent }

func (stage *Install_Issu_Stage) GetParent() types.Entity { return stage.parent }

func (stage *Install_Issu_Stage) GetParentYangName() string { return "issu" }

// Install_Issu_Stage_NodeInProgress
// Nodes in progress
type Install_Issu_Stage_NodeInProgress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // node. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node []interface{}
}

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetFilter() yfilter.YFilter { return nodeInProgress.YFilter }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) SetFilter(yf yfilter.YFilter) { nodeInProgress.YFilter = yf }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetSegmentPath() string {
    return "node-in-progress"
}

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = nodeInProgress.Node
    return leafs
}

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetYangName() string { return "node-in-progress" }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) SetParent(parent types.Entity) { nodeInProgress.parent = parent }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetParent() types.Entity { return nodeInProgress.parent }

func (nodeInProgress *Install_Issu_Stage_NodeInProgress) GetParentYangName() string { return "stage" }

// Install_Issu_Stage_NodesInLoad
// Node in LOAD phase
type Install_Issu_Stage_NodesInLoad struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // node. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node []interface{}
}

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetFilter() yfilter.YFilter { return nodesInLoad.YFilter }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) SetFilter(yf yfilter.YFilter) { nodesInLoad.YFilter = yf }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetSegmentPath() string {
    return "nodes-in-load"
}

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = nodesInLoad.Node
    return leafs
}

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetBundleName() string { return "cisco_ios_xr" }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetYangName() string { return "nodes-in-load" }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) SetParent(parent types.Entity) { nodesInLoad.parent = parent }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetParent() types.Entity { return nodesInLoad.parent }

func (nodesInLoad *Install_Issu_Stage_NodesInLoad) GetParentYangName() string { return "stage" }

// Install_Issu_Stage_NodesInRun
// Node in RUN phase
type Install_Issu_Stage_NodesInRun struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // node. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node []interface{}
}

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetFilter() yfilter.YFilter { return nodesInRun.YFilter }

func (nodesInRun *Install_Issu_Stage_NodesInRun) SetFilter(yf yfilter.YFilter) { nodesInRun.YFilter = yf }

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetSegmentPath() string {
    return "nodes-in-run"
}

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = nodesInRun.Node
    return leafs
}

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetBundleName() string { return "cisco_ios_xr" }

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetYangName() string { return "nodes-in-run" }

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodesInRun *Install_Issu_Stage_NodesInRun) SetParent(parent types.Entity) { nodesInRun.parent = parent }

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetParent() types.Entity { return nodesInRun.parent }

func (nodesInRun *Install_Issu_Stage_NodesInRun) GetParentYangName() string { return "stage" }

// Install_Issu_Stage_NcNodes
// None-conforming nodes
type Install_Issu_Stage_NcNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // node. The type is slice of string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node []interface{}
}

func (ncNodes *Install_Issu_Stage_NcNodes) GetFilter() yfilter.YFilter { return ncNodes.YFilter }

func (ncNodes *Install_Issu_Stage_NcNodes) SetFilter(yf yfilter.YFilter) { ncNodes.YFilter = yf }

func (ncNodes *Install_Issu_Stage_NcNodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (ncNodes *Install_Issu_Stage_NcNodes) GetSegmentPath() string {
    return "nc-nodes"
}

func (ncNodes *Install_Issu_Stage_NcNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ncNodes *Install_Issu_Stage_NcNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ncNodes *Install_Issu_Stage_NcNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = ncNodes.Node
    return leafs
}

func (ncNodes *Install_Issu_Stage_NcNodes) GetBundleName() string { return "cisco_ios_xr" }

func (ncNodes *Install_Issu_Stage_NcNodes) GetYangName() string { return "nc-nodes" }

func (ncNodes *Install_Issu_Stage_NcNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncNodes *Install_Issu_Stage_NcNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncNodes *Install_Issu_Stage_NcNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncNodes *Install_Issu_Stage_NcNodes) SetParent(parent types.Entity) { ncNodes.parent = parent }

func (ncNodes *Install_Issu_Stage_NcNodes) GetParent() types.Entity { return ncNodes.parent }

func (ncNodes *Install_Issu_Stage_NcNodes) GetParentYangName() string { return "stage" }

// Install_BootImage
// System Boot Image
type Install_BootImage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The boot image. The type is string.
    SystemImageFile interface{}
}

func (bootImage *Install_BootImage) GetFilter() yfilter.YFilter { return bootImage.YFilter }

func (bootImage *Install_BootImage) SetFilter(yf yfilter.YFilter) { bootImage.YFilter = yf }

func (bootImage *Install_BootImage) GetGoName(yname string) string {
    if yname == "system-image-file" { return "SystemImageFile" }
    return ""
}

func (bootImage *Install_BootImage) GetSegmentPath() string {
    return "boot-image"
}

func (bootImage *Install_BootImage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootImage *Install_BootImage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootImage *Install_BootImage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-image-file"] = bootImage.SystemImageFile
    return leafs
}

func (bootImage *Install_BootImage) GetBundleName() string { return "cisco_ios_xr" }

func (bootImage *Install_BootImage) GetYangName() string { return "boot-image" }

func (bootImage *Install_BootImage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootImage *Install_BootImage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootImage *Install_BootImage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootImage *Install_BootImage) SetParent(parent types.Entity) { bootImage.parent = parent }

func (bootImage *Install_BootImage) GetParent() types.Entity { return bootImage.parent }

func (bootImage *Install_BootImage) GetParentYangName() string { return "install" }

// Install_Logs
// Install operation log
type Install_Logs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log information. The type is slice of Install_Logs_Log.
    Log []Install_Logs_Log
}

func (logs *Install_Logs) GetFilter() yfilter.YFilter { return logs.YFilter }

func (logs *Install_Logs) SetFilter(yf yfilter.YFilter) { logs.YFilter = yf }

func (logs *Install_Logs) GetGoName(yname string) string {
    if yname == "log" { return "Log" }
    return ""
}

func (logs *Install_Logs) GetSegmentPath() string {
    return "logs"
}

func (logs *Install_Logs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log" {
        for _, c := range logs.Log {
            if logs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Logs_Log{}
        logs.Log = append(logs.Log, child)
        return &logs.Log[len(logs.Log)-1]
    }
    return nil
}

func (logs *Install_Logs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range logs.Log {
        children[logs.Log[i].GetSegmentPath()] = &logs.Log[i]
    }
    return children
}

func (logs *Install_Logs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logs *Install_Logs) GetBundleName() string { return "cisco_ios_xr" }

func (logs *Install_Logs) GetYangName() string { return "logs" }

func (logs *Install_Logs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logs *Install_Logs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logs *Install_Logs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logs *Install_Logs) SetParent(parent types.Entity) { logs.parent = parent }

func (logs *Install_Logs) GetParent() types.Entity { return logs.parent }

func (logs *Install_Logs) GetParentYangName() string { return "install" }

// Install_Logs_Log
// Log information
type Install_Logs_Log struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Install operation request ID. The type is
    // interface{} with range: -2147483648..2147483647.
    RequestId interface{}

    // Header information. The type is slice of Install_Logs_Log_Header.
    Header []Install_Logs_Log_Header

    // Summary information. The type is slice of Install_Logs_Log_Summary.
    Summary []Install_Logs_Log_Summary

    // Status Information Logs. The type is slice of Install_Logs_Log_Message.
    Message []Install_Logs_Log_Message

    // Install changes. The type is slice of Install_Logs_Log_Change.
    Change []Install_Logs_Log_Change

    // Install details. The type is slice of Install_Logs_Log_Detail.
    Detail []Install_Logs_Log_Detail

    // Install communications. The type is slice of
    // Install_Logs_Log_Communication.
    Communication []Install_Logs_Log_Communication
}

func (log *Install_Logs_Log) GetFilter() yfilter.YFilter { return log.YFilter }

func (log *Install_Logs_Log) SetFilter(yf yfilter.YFilter) { log.YFilter = yf }

func (log *Install_Logs_Log) GetGoName(yname string) string {
    if yname == "request-id" { return "RequestId" }
    if yname == "header" { return "Header" }
    if yname == "summary" { return "Summary" }
    if yname == "message" { return "Message" }
    if yname == "change" { return "Change" }
    if yname == "detail" { return "Detail" }
    if yname == "communication" { return "Communication" }
    return ""
}

func (log *Install_Logs_Log) GetSegmentPath() string {
    return "log" + "[request-id='" + fmt.Sprintf("%v", log.RequestId) + "']"
}

func (log *Install_Logs_Log) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        for _, c := range log.Header {
            if log.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Logs_Log_Header{}
        log.Header = append(log.Header, child)
        return &log.Header[len(log.Header)-1]
    }
    if childYangName == "summary" {
        for _, c := range log.Summary {
            if log.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Logs_Log_Summary{}
        log.Summary = append(log.Summary, child)
        return &log.Summary[len(log.Summary)-1]
    }
    if childYangName == "message" {
        for _, c := range log.Message {
            if log.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Logs_Log_Message{}
        log.Message = append(log.Message, child)
        return &log.Message[len(log.Message)-1]
    }
    if childYangName == "change" {
        for _, c := range log.Change {
            if log.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Logs_Log_Change{}
        log.Change = append(log.Change, child)
        return &log.Change[len(log.Change)-1]
    }
    if childYangName == "detail" {
        for _, c := range log.Detail {
            if log.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Logs_Log_Detail{}
        log.Detail = append(log.Detail, child)
        return &log.Detail[len(log.Detail)-1]
    }
    if childYangName == "communication" {
        for _, c := range log.Communication {
            if log.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Install_Logs_Log_Communication{}
        log.Communication = append(log.Communication, child)
        return &log.Communication[len(log.Communication)-1]
    }
    return nil
}

func (log *Install_Logs_Log) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range log.Header {
        children[log.Header[i].GetSegmentPath()] = &log.Header[i]
    }
    for i := range log.Summary {
        children[log.Summary[i].GetSegmentPath()] = &log.Summary[i]
    }
    for i := range log.Message {
        children[log.Message[i].GetSegmentPath()] = &log.Message[i]
    }
    for i := range log.Change {
        children[log.Change[i].GetSegmentPath()] = &log.Change[i]
    }
    for i := range log.Detail {
        children[log.Detail[i].GetSegmentPath()] = &log.Detail[i]
    }
    for i := range log.Communication {
        children[log.Communication[i].GetSegmentPath()] = &log.Communication[i]
    }
    return children
}

func (log *Install_Logs_Log) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["request-id"] = log.RequestId
    return leafs
}

func (log *Install_Logs_Log) GetBundleName() string { return "cisco_ios_xr" }

func (log *Install_Logs_Log) GetYangName() string { return "log" }

func (log *Install_Logs_Log) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (log *Install_Logs_Log) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (log *Install_Logs_Log) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (log *Install_Logs_Log) SetParent(parent types.Entity) { log.parent = parent }

func (log *Install_Logs_Log) GetParent() types.Entity { return log.parent }

func (log *Install_Logs_Log) GetParentYangName() string { return "logs" }

// Install_Logs_Log_Header
// Header information
type Install_Logs_Log_Header struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log contents.
    LogContents Install_Logs_Log_Header_LogContents
}

func (header *Install_Logs_Log_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *Install_Logs_Log_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *Install_Logs_Log_Header) GetGoName(yname string) string {
    if yname == "log-contents" { return "LogContents" }
    return ""
}

func (header *Install_Logs_Log_Header) GetSegmentPath() string {
    return "header"
}

func (header *Install_Logs_Log_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-contents" {
        return &header.LogContents
    }
    return nil
}

func (header *Install_Logs_Log_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log-contents"] = &header.LogContents
    return children
}

func (header *Install_Logs_Log_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (header *Install_Logs_Log_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *Install_Logs_Log_Header) GetYangName() string { return "header" }

func (header *Install_Logs_Log_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *Install_Logs_Log_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *Install_Logs_Log_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *Install_Logs_Log_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *Install_Logs_Log_Header) GetParent() types.Entity { return header.parent }

func (header *Install_Logs_Log_Header) GetParentYangName() string { return "log" }

// Install_Logs_Log_Header_LogContents
// Log contents
type Install_Logs_Log_Header_LogContents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Header_LogContents_V3
}

func (logContents *Install_Logs_Log_Header_LogContents) GetFilter() yfilter.YFilter { return logContents.YFilter }

func (logContents *Install_Logs_Log_Header_LogContents) SetFilter(yf yfilter.YFilter) { logContents.YFilter = yf }

func (logContents *Install_Logs_Log_Header_LogContents) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "v3" { return "V3" }
    return ""
}

func (logContents *Install_Logs_Log_Header_LogContents) GetSegmentPath() string {
    return "log-contents"
}

func (logContents *Install_Logs_Log_Header_LogContents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "v3" {
        return &logContents.V3
    }
    return nil
}

func (logContents *Install_Logs_Log_Header_LogContents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["v3"] = &logContents.V3
    return children
}

func (logContents *Install_Logs_Log_Header_LogContents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = logContents.Version
    return leafs
}

func (logContents *Install_Logs_Log_Header_LogContents) GetBundleName() string { return "cisco_ios_xr" }

func (logContents *Install_Logs_Log_Header_LogContents) GetYangName() string { return "log-contents" }

func (logContents *Install_Logs_Log_Header_LogContents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logContents *Install_Logs_Log_Header_LogContents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logContents *Install_Logs_Log_Header_LogContents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logContents *Install_Logs_Log_Header_LogContents) SetParent(parent types.Entity) { logContents.parent = parent }

func (logContents *Install_Logs_Log_Header_LogContents) GetParent() types.Entity { return logContents.parent }

func (logContents *Install_Logs_Log_Header_LogContents) GetParentYangName() string { return "header" }

// Install_Logs_Log_Header_LogContents_V3
// v3
type Install_Logs_Log_Header_LogContents_V3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Header_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetFilter() yfilter.YFilter { return v3.YFilter }

func (v3 *Install_Logs_Log_Header_LogContents_V3) SetFilter(yf yfilter.YFilter) { v3.YFilter = yf }

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "message" { return "Message" }
    if yname == "scope" { return "Scope" }
    return ""
}

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetSegmentPath() string {
    return "v3"
}

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scope" {
        return &v3.Scope
    }
    return nil
}

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scope"] = &v3.Scope
    return children
}

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = v3.Category
    leafs["message"] = v3.Message
    return leafs
}

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetBundleName() string { return "cisco_ios_xr" }

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetYangName() string { return "v3" }

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v3 *Install_Logs_Log_Header_LogContents_V3) SetParent(parent types.Entity) { v3.parent = parent }

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetParent() types.Entity { return v3.parent }

func (v3 *Install_Logs_Log_Header_LogContents_V3) GetParentYangName() string { return "log-contents" }

// Install_Logs_Log_Header_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Header_LogContents_V3_Scope struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetFilter() yfilter.YFilter { return scope.YFilter }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) SetFilter(yf yfilter.YFilter) { scope.YFilter = yf }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetGoName(yname string) string {
    if yname == "admin-read" { return "AdminRead" }
    if yname == "affected-sd-rs" { return "AffectedSdRs" }
    return ""
}

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetSegmentPath() string {
    return "scope"
}

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-read"] = scope.AdminRead
    leafs["affected-sd-rs"] = scope.AffectedSdRs
    return leafs
}

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetBundleName() string { return "cisco_ios_xr" }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetYangName() string { return "scope" }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) SetParent(parent types.Entity) { scope.parent = parent }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetParent() types.Entity { return scope.parent }

func (scope *Install_Logs_Log_Header_LogContents_V3_Scope) GetParentYangName() string { return "v3" }

// Install_Logs_Log_Summary
// Summary information
type Install_Logs_Log_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log contents.
    LogContents Install_Logs_Log_Summary_LogContents
}

func (summary *Install_Logs_Log_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Install_Logs_Log_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Install_Logs_Log_Summary) GetGoName(yname string) string {
    if yname == "log-contents" { return "LogContents" }
    return ""
}

func (summary *Install_Logs_Log_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Install_Logs_Log_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-contents" {
        return &summary.LogContents
    }
    return nil
}

func (summary *Install_Logs_Log_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log-contents"] = &summary.LogContents
    return children
}

func (summary *Install_Logs_Log_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summary *Install_Logs_Log_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Install_Logs_Log_Summary) GetYangName() string { return "summary" }

func (summary *Install_Logs_Log_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Install_Logs_Log_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Install_Logs_Log_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Install_Logs_Log_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Install_Logs_Log_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Install_Logs_Log_Summary) GetParentYangName() string { return "log" }

// Install_Logs_Log_Summary_LogContents
// Log contents
type Install_Logs_Log_Summary_LogContents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Summary_LogContents_V3
}

func (logContents *Install_Logs_Log_Summary_LogContents) GetFilter() yfilter.YFilter { return logContents.YFilter }

func (logContents *Install_Logs_Log_Summary_LogContents) SetFilter(yf yfilter.YFilter) { logContents.YFilter = yf }

func (logContents *Install_Logs_Log_Summary_LogContents) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "v3" { return "V3" }
    return ""
}

func (logContents *Install_Logs_Log_Summary_LogContents) GetSegmentPath() string {
    return "log-contents"
}

func (logContents *Install_Logs_Log_Summary_LogContents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "v3" {
        return &logContents.V3
    }
    return nil
}

func (logContents *Install_Logs_Log_Summary_LogContents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["v3"] = &logContents.V3
    return children
}

func (logContents *Install_Logs_Log_Summary_LogContents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = logContents.Version
    return leafs
}

func (logContents *Install_Logs_Log_Summary_LogContents) GetBundleName() string { return "cisco_ios_xr" }

func (logContents *Install_Logs_Log_Summary_LogContents) GetYangName() string { return "log-contents" }

func (logContents *Install_Logs_Log_Summary_LogContents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logContents *Install_Logs_Log_Summary_LogContents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logContents *Install_Logs_Log_Summary_LogContents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logContents *Install_Logs_Log_Summary_LogContents) SetParent(parent types.Entity) { logContents.parent = parent }

func (logContents *Install_Logs_Log_Summary_LogContents) GetParent() types.Entity { return logContents.parent }

func (logContents *Install_Logs_Log_Summary_LogContents) GetParentYangName() string { return "summary" }

// Install_Logs_Log_Summary_LogContents_V3
// v3
type Install_Logs_Log_Summary_LogContents_V3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Summary_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetFilter() yfilter.YFilter { return v3.YFilter }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) SetFilter(yf yfilter.YFilter) { v3.YFilter = yf }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "message" { return "Message" }
    if yname == "scope" { return "Scope" }
    return ""
}

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetSegmentPath() string {
    return "v3"
}

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scope" {
        return &v3.Scope
    }
    return nil
}

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scope"] = &v3.Scope
    return children
}

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = v3.Category
    leafs["message"] = v3.Message
    return leafs
}

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetBundleName() string { return "cisco_ios_xr" }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetYangName() string { return "v3" }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) SetParent(parent types.Entity) { v3.parent = parent }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetParent() types.Entity { return v3.parent }

func (v3 *Install_Logs_Log_Summary_LogContents_V3) GetParentYangName() string { return "log-contents" }

// Install_Logs_Log_Summary_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Summary_LogContents_V3_Scope struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetFilter() yfilter.YFilter { return scope.YFilter }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) SetFilter(yf yfilter.YFilter) { scope.YFilter = yf }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetGoName(yname string) string {
    if yname == "admin-read" { return "AdminRead" }
    if yname == "affected-sd-rs" { return "AffectedSdRs" }
    return ""
}

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetSegmentPath() string {
    return "scope"
}

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-read"] = scope.AdminRead
    leafs["affected-sd-rs"] = scope.AffectedSdRs
    return leafs
}

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetBundleName() string { return "cisco_ios_xr" }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetYangName() string { return "scope" }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) SetParent(parent types.Entity) { scope.parent = parent }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetParent() types.Entity { return scope.parent }

func (scope *Install_Logs_Log_Summary_LogContents_V3_Scope) GetParentYangName() string { return "v3" }

// Install_Logs_Log_Message
// Status Information Logs
type Install_Logs_Log_Message struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log contents.
    LogContents Install_Logs_Log_Message_LogContents
}

func (message *Install_Logs_Log_Message) GetFilter() yfilter.YFilter { return message.YFilter }

func (message *Install_Logs_Log_Message) SetFilter(yf yfilter.YFilter) { message.YFilter = yf }

func (message *Install_Logs_Log_Message) GetGoName(yname string) string {
    if yname == "log-contents" { return "LogContents" }
    return ""
}

func (message *Install_Logs_Log_Message) GetSegmentPath() string {
    return "message"
}

func (message *Install_Logs_Log_Message) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-contents" {
        return &message.LogContents
    }
    return nil
}

func (message *Install_Logs_Log_Message) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log-contents"] = &message.LogContents
    return children
}

func (message *Install_Logs_Log_Message) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (message *Install_Logs_Log_Message) GetBundleName() string { return "cisco_ios_xr" }

func (message *Install_Logs_Log_Message) GetYangName() string { return "message" }

func (message *Install_Logs_Log_Message) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (message *Install_Logs_Log_Message) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (message *Install_Logs_Log_Message) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (message *Install_Logs_Log_Message) SetParent(parent types.Entity) { message.parent = parent }

func (message *Install_Logs_Log_Message) GetParent() types.Entity { return message.parent }

func (message *Install_Logs_Log_Message) GetParentYangName() string { return "log" }

// Install_Logs_Log_Message_LogContents
// Log contents
type Install_Logs_Log_Message_LogContents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Message_LogContents_V3
}

func (logContents *Install_Logs_Log_Message_LogContents) GetFilter() yfilter.YFilter { return logContents.YFilter }

func (logContents *Install_Logs_Log_Message_LogContents) SetFilter(yf yfilter.YFilter) { logContents.YFilter = yf }

func (logContents *Install_Logs_Log_Message_LogContents) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "v3" { return "V3" }
    return ""
}

func (logContents *Install_Logs_Log_Message_LogContents) GetSegmentPath() string {
    return "log-contents"
}

func (logContents *Install_Logs_Log_Message_LogContents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "v3" {
        return &logContents.V3
    }
    return nil
}

func (logContents *Install_Logs_Log_Message_LogContents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["v3"] = &logContents.V3
    return children
}

func (logContents *Install_Logs_Log_Message_LogContents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = logContents.Version
    return leafs
}

func (logContents *Install_Logs_Log_Message_LogContents) GetBundleName() string { return "cisco_ios_xr" }

func (logContents *Install_Logs_Log_Message_LogContents) GetYangName() string { return "log-contents" }

func (logContents *Install_Logs_Log_Message_LogContents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logContents *Install_Logs_Log_Message_LogContents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logContents *Install_Logs_Log_Message_LogContents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logContents *Install_Logs_Log_Message_LogContents) SetParent(parent types.Entity) { logContents.parent = parent }

func (logContents *Install_Logs_Log_Message_LogContents) GetParent() types.Entity { return logContents.parent }

func (logContents *Install_Logs_Log_Message_LogContents) GetParentYangName() string { return "message" }

// Install_Logs_Log_Message_LogContents_V3
// v3
type Install_Logs_Log_Message_LogContents_V3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Message_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetFilter() yfilter.YFilter { return v3.YFilter }

func (v3 *Install_Logs_Log_Message_LogContents_V3) SetFilter(yf yfilter.YFilter) { v3.YFilter = yf }

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "message" { return "Message" }
    if yname == "scope" { return "Scope" }
    return ""
}

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetSegmentPath() string {
    return "v3"
}

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scope" {
        return &v3.Scope
    }
    return nil
}

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scope"] = &v3.Scope
    return children
}

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = v3.Category
    leafs["message"] = v3.Message
    return leafs
}

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetBundleName() string { return "cisco_ios_xr" }

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetYangName() string { return "v3" }

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v3 *Install_Logs_Log_Message_LogContents_V3) SetParent(parent types.Entity) { v3.parent = parent }

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetParent() types.Entity { return v3.parent }

func (v3 *Install_Logs_Log_Message_LogContents_V3) GetParentYangName() string { return "log-contents" }

// Install_Logs_Log_Message_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Message_LogContents_V3_Scope struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetFilter() yfilter.YFilter { return scope.YFilter }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) SetFilter(yf yfilter.YFilter) { scope.YFilter = yf }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetGoName(yname string) string {
    if yname == "admin-read" { return "AdminRead" }
    if yname == "affected-sd-rs" { return "AffectedSdRs" }
    return ""
}

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetSegmentPath() string {
    return "scope"
}

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-read"] = scope.AdminRead
    leafs["affected-sd-rs"] = scope.AffectedSdRs
    return leafs
}

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetBundleName() string { return "cisco_ios_xr" }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetYangName() string { return "scope" }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) SetParent(parent types.Entity) { scope.parent = parent }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetParent() types.Entity { return scope.parent }

func (scope *Install_Logs_Log_Message_LogContents_V3_Scope) GetParentYangName() string { return "v3" }

// Install_Logs_Log_Change
// Install changes
type Install_Logs_Log_Change struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log contents.
    LogContents Install_Logs_Log_Change_LogContents
}

func (change *Install_Logs_Log_Change) GetFilter() yfilter.YFilter { return change.YFilter }

func (change *Install_Logs_Log_Change) SetFilter(yf yfilter.YFilter) { change.YFilter = yf }

func (change *Install_Logs_Log_Change) GetGoName(yname string) string {
    if yname == "log-contents" { return "LogContents" }
    return ""
}

func (change *Install_Logs_Log_Change) GetSegmentPath() string {
    return "change"
}

func (change *Install_Logs_Log_Change) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-contents" {
        return &change.LogContents
    }
    return nil
}

func (change *Install_Logs_Log_Change) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log-contents"] = &change.LogContents
    return children
}

func (change *Install_Logs_Log_Change) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (change *Install_Logs_Log_Change) GetBundleName() string { return "cisco_ios_xr" }

func (change *Install_Logs_Log_Change) GetYangName() string { return "change" }

func (change *Install_Logs_Log_Change) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (change *Install_Logs_Log_Change) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (change *Install_Logs_Log_Change) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (change *Install_Logs_Log_Change) SetParent(parent types.Entity) { change.parent = parent }

func (change *Install_Logs_Log_Change) GetParent() types.Entity { return change.parent }

func (change *Install_Logs_Log_Change) GetParentYangName() string { return "log" }

// Install_Logs_Log_Change_LogContents
// Log contents
type Install_Logs_Log_Change_LogContents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Change_LogContents_V3
}

func (logContents *Install_Logs_Log_Change_LogContents) GetFilter() yfilter.YFilter { return logContents.YFilter }

func (logContents *Install_Logs_Log_Change_LogContents) SetFilter(yf yfilter.YFilter) { logContents.YFilter = yf }

func (logContents *Install_Logs_Log_Change_LogContents) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "v3" { return "V3" }
    return ""
}

func (logContents *Install_Logs_Log_Change_LogContents) GetSegmentPath() string {
    return "log-contents"
}

func (logContents *Install_Logs_Log_Change_LogContents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "v3" {
        return &logContents.V3
    }
    return nil
}

func (logContents *Install_Logs_Log_Change_LogContents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["v3"] = &logContents.V3
    return children
}

func (logContents *Install_Logs_Log_Change_LogContents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = logContents.Version
    return leafs
}

func (logContents *Install_Logs_Log_Change_LogContents) GetBundleName() string { return "cisco_ios_xr" }

func (logContents *Install_Logs_Log_Change_LogContents) GetYangName() string { return "log-contents" }

func (logContents *Install_Logs_Log_Change_LogContents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logContents *Install_Logs_Log_Change_LogContents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logContents *Install_Logs_Log_Change_LogContents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logContents *Install_Logs_Log_Change_LogContents) SetParent(parent types.Entity) { logContents.parent = parent }

func (logContents *Install_Logs_Log_Change_LogContents) GetParent() types.Entity { return logContents.parent }

func (logContents *Install_Logs_Log_Change_LogContents) GetParentYangName() string { return "change" }

// Install_Logs_Log_Change_LogContents_V3
// v3
type Install_Logs_Log_Change_LogContents_V3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Change_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetFilter() yfilter.YFilter { return v3.YFilter }

func (v3 *Install_Logs_Log_Change_LogContents_V3) SetFilter(yf yfilter.YFilter) { v3.YFilter = yf }

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "message" { return "Message" }
    if yname == "scope" { return "Scope" }
    return ""
}

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetSegmentPath() string {
    return "v3"
}

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scope" {
        return &v3.Scope
    }
    return nil
}

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scope"] = &v3.Scope
    return children
}

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = v3.Category
    leafs["message"] = v3.Message
    return leafs
}

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetBundleName() string { return "cisco_ios_xr" }

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetYangName() string { return "v3" }

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v3 *Install_Logs_Log_Change_LogContents_V3) SetParent(parent types.Entity) { v3.parent = parent }

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetParent() types.Entity { return v3.parent }

func (v3 *Install_Logs_Log_Change_LogContents_V3) GetParentYangName() string { return "log-contents" }

// Install_Logs_Log_Change_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Change_LogContents_V3_Scope struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetFilter() yfilter.YFilter { return scope.YFilter }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) SetFilter(yf yfilter.YFilter) { scope.YFilter = yf }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetGoName(yname string) string {
    if yname == "admin-read" { return "AdminRead" }
    if yname == "affected-sd-rs" { return "AffectedSdRs" }
    return ""
}

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetSegmentPath() string {
    return "scope"
}

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-read"] = scope.AdminRead
    leafs["affected-sd-rs"] = scope.AffectedSdRs
    return leafs
}

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetBundleName() string { return "cisco_ios_xr" }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetYangName() string { return "scope" }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) SetParent(parent types.Entity) { scope.parent = parent }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetParent() types.Entity { return scope.parent }

func (scope *Install_Logs_Log_Change_LogContents_V3_Scope) GetParentYangName() string { return "v3" }

// Install_Logs_Log_Detail
// Install details
type Install_Logs_Log_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log contents.
    LogContents Install_Logs_Log_Detail_LogContents
}

func (detail *Install_Logs_Log_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Install_Logs_Log_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Install_Logs_Log_Detail) GetGoName(yname string) string {
    if yname == "log-contents" { return "LogContents" }
    return ""
}

func (detail *Install_Logs_Log_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Install_Logs_Log_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-contents" {
        return &detail.LogContents
    }
    return nil
}

func (detail *Install_Logs_Log_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log-contents"] = &detail.LogContents
    return children
}

func (detail *Install_Logs_Log_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detail *Install_Logs_Log_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Install_Logs_Log_Detail) GetYangName() string { return "detail" }

func (detail *Install_Logs_Log_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Install_Logs_Log_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Install_Logs_Log_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Install_Logs_Log_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Install_Logs_Log_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Install_Logs_Log_Detail) GetParentYangName() string { return "log" }

// Install_Logs_Log_Detail_LogContents
// Log contents
type Install_Logs_Log_Detail_LogContents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Detail_LogContents_V3
}

func (logContents *Install_Logs_Log_Detail_LogContents) GetFilter() yfilter.YFilter { return logContents.YFilter }

func (logContents *Install_Logs_Log_Detail_LogContents) SetFilter(yf yfilter.YFilter) { logContents.YFilter = yf }

func (logContents *Install_Logs_Log_Detail_LogContents) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "v3" { return "V3" }
    return ""
}

func (logContents *Install_Logs_Log_Detail_LogContents) GetSegmentPath() string {
    return "log-contents"
}

func (logContents *Install_Logs_Log_Detail_LogContents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "v3" {
        return &logContents.V3
    }
    return nil
}

func (logContents *Install_Logs_Log_Detail_LogContents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["v3"] = &logContents.V3
    return children
}

func (logContents *Install_Logs_Log_Detail_LogContents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = logContents.Version
    return leafs
}

func (logContents *Install_Logs_Log_Detail_LogContents) GetBundleName() string { return "cisco_ios_xr" }

func (logContents *Install_Logs_Log_Detail_LogContents) GetYangName() string { return "log-contents" }

func (logContents *Install_Logs_Log_Detail_LogContents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logContents *Install_Logs_Log_Detail_LogContents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logContents *Install_Logs_Log_Detail_LogContents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logContents *Install_Logs_Log_Detail_LogContents) SetParent(parent types.Entity) { logContents.parent = parent }

func (logContents *Install_Logs_Log_Detail_LogContents) GetParent() types.Entity { return logContents.parent }

func (logContents *Install_Logs_Log_Detail_LogContents) GetParentYangName() string { return "detail" }

// Install_Logs_Log_Detail_LogContents_V3
// v3
type Install_Logs_Log_Detail_LogContents_V3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Detail_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetFilter() yfilter.YFilter { return v3.YFilter }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) SetFilter(yf yfilter.YFilter) { v3.YFilter = yf }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "message" { return "Message" }
    if yname == "scope" { return "Scope" }
    return ""
}

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetSegmentPath() string {
    return "v3"
}

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scope" {
        return &v3.Scope
    }
    return nil
}

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scope"] = &v3.Scope
    return children
}

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = v3.Category
    leafs["message"] = v3.Message
    return leafs
}

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetBundleName() string { return "cisco_ios_xr" }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetYangName() string { return "v3" }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) SetParent(parent types.Entity) { v3.parent = parent }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetParent() types.Entity { return v3.parent }

func (v3 *Install_Logs_Log_Detail_LogContents_V3) GetParentYangName() string { return "log-contents" }

// Install_Logs_Log_Detail_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Detail_LogContents_V3_Scope struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetFilter() yfilter.YFilter { return scope.YFilter }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) SetFilter(yf yfilter.YFilter) { scope.YFilter = yf }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetGoName(yname string) string {
    if yname == "admin-read" { return "AdminRead" }
    if yname == "affected-sd-rs" { return "AffectedSdRs" }
    return ""
}

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetSegmentPath() string {
    return "scope"
}

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-read"] = scope.AdminRead
    leafs["affected-sd-rs"] = scope.AffectedSdRs
    return leafs
}

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetBundleName() string { return "cisco_ios_xr" }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetYangName() string { return "scope" }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) SetParent(parent types.Entity) { scope.parent = parent }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetParent() types.Entity { return scope.parent }

func (scope *Install_Logs_Log_Detail_LogContents_V3_Scope) GetParentYangName() string { return "v3" }

// Install_Logs_Log_Communication
// Install communications
type Install_Logs_Log_Communication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log contents.
    LogContents Install_Logs_Log_Communication_LogContents
}

func (communication *Install_Logs_Log_Communication) GetFilter() yfilter.YFilter { return communication.YFilter }

func (communication *Install_Logs_Log_Communication) SetFilter(yf yfilter.YFilter) { communication.YFilter = yf }

func (communication *Install_Logs_Log_Communication) GetGoName(yname string) string {
    if yname == "log-contents" { return "LogContents" }
    return ""
}

func (communication *Install_Logs_Log_Communication) GetSegmentPath() string {
    return "communication"
}

func (communication *Install_Logs_Log_Communication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "log-contents" {
        return &communication.LogContents
    }
    return nil
}

func (communication *Install_Logs_Log_Communication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["log-contents"] = &communication.LogContents
    return children
}

func (communication *Install_Logs_Log_Communication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (communication *Install_Logs_Log_Communication) GetBundleName() string { return "cisco_ios_xr" }

func (communication *Install_Logs_Log_Communication) GetYangName() string { return "communication" }

func (communication *Install_Logs_Log_Communication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (communication *Install_Logs_Log_Communication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (communication *Install_Logs_Log_Communication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (communication *Install_Logs_Log_Communication) SetParent(parent types.Entity) { communication.parent = parent }

func (communication *Install_Logs_Log_Communication) GetParent() types.Entity { return communication.parent }

func (communication *Install_Logs_Log_Communication) GetParentYangName() string { return "log" }

// Install_Logs_Log_Communication_LogContents
// Log contents
type Install_Logs_Log_Communication_LogContents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // v3.
    V3 Install_Logs_Log_Communication_LogContents_V3
}

func (logContents *Install_Logs_Log_Communication_LogContents) GetFilter() yfilter.YFilter { return logContents.YFilter }

func (logContents *Install_Logs_Log_Communication_LogContents) SetFilter(yf yfilter.YFilter) { logContents.YFilter = yf }

func (logContents *Install_Logs_Log_Communication_LogContents) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "v3" { return "V3" }
    return ""
}

func (logContents *Install_Logs_Log_Communication_LogContents) GetSegmentPath() string {
    return "log-contents"
}

func (logContents *Install_Logs_Log_Communication_LogContents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "v3" {
        return &logContents.V3
    }
    return nil
}

func (logContents *Install_Logs_Log_Communication_LogContents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["v3"] = &logContents.V3
    return children
}

func (logContents *Install_Logs_Log_Communication_LogContents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = logContents.Version
    return leafs
}

func (logContents *Install_Logs_Log_Communication_LogContents) GetBundleName() string { return "cisco_ios_xr" }

func (logContents *Install_Logs_Log_Communication_LogContents) GetYangName() string { return "log-contents" }

func (logContents *Install_Logs_Log_Communication_LogContents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logContents *Install_Logs_Log_Communication_LogContents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logContents *Install_Logs_Log_Communication_LogContents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logContents *Install_Logs_Log_Communication_LogContents) SetParent(parent types.Entity) { logContents.parent = parent }

func (logContents *Install_Logs_Log_Communication_LogContents) GetParent() types.Entity { return logContents.parent }

func (logContents *Install_Logs_Log_Communication_LogContents) GetParentYangName() string { return "communication" }

// Install_Logs_Log_Communication_LogContents_V3
// v3
type Install_Logs_Log_Communication_LogContents_V3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of the message. The type is InstmgrBagLogEntryUserMsgCategory.
    Category interface{}

    // Message. The type is string.
    Message interface{}

    // Scope of the message.
    Scope Install_Logs_Log_Communication_LogContents_V3_Scope
}

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetFilter() yfilter.YFilter { return v3.YFilter }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) SetFilter(yf yfilter.YFilter) { v3.YFilter = yf }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "message" { return "Message" }
    if yname == "scope" { return "Scope" }
    return ""
}

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetSegmentPath() string {
    return "v3"
}

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "scope" {
        return &v3.Scope
    }
    return nil
}

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["scope"] = &v3.Scope
    return children
}

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = v3.Category
    leafs["message"] = v3.Message
    return leafs
}

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetBundleName() string { return "cisco_ios_xr" }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetYangName() string { return "v3" }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) SetParent(parent types.Entity) { v3.parent = parent }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetParent() types.Entity { return v3.parent }

func (v3 *Install_Logs_Log_Communication_LogContents_V3) GetParentYangName() string { return "log-contents" }

// Install_Logs_Log_Communication_LogContents_V3_Scope
// Scope of the message
type Install_Logs_Log_Communication_LogContents_V3_Scope struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Does the admin want to read this?. The type is bool.
    AdminRead interface{}

    // Which SDRs are affected by the message. The type is interface{} with range:
    // 0..4294967295.
    AffectedSdRs interface{}
}

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetFilter() yfilter.YFilter { return scope.YFilter }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) SetFilter(yf yfilter.YFilter) { scope.YFilter = yf }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetGoName(yname string) string {
    if yname == "admin-read" { return "AdminRead" }
    if yname == "affected-sd-rs" { return "AffectedSdRs" }
    return ""
}

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetSegmentPath() string {
    return "scope"
}

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-read"] = scope.AdminRead
    leafs["affected-sd-rs"] = scope.AffectedSdRs
    return leafs
}

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetBundleName() string { return "cisco_ios_xr" }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetYangName() string { return "scope" }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) SetParent(parent types.Entity) { scope.parent = parent }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetParent() types.Entity { return scope.parent }

func (scope *Install_Logs_Log_Communication_LogContents_V3_Scope) GetParentYangName() string { return "v3" }

