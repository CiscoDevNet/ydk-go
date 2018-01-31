// This module contains a collection of YANG definitions
// for Cisco IOS-XR icpe-infra package operational data.
// 
// This module contains definitions
// for the following management objects:
//   nv-satellite: Satellite operational information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package icpe_infra_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package icpe_infra_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-icpe-infra-oper nv-satellite}", reflect.TypeOf(NvSatellite{}))
    ydk.RegisterEntity("Cisco-IOS-XR-icpe-infra-oper:nv-satellite", reflect.TypeOf(NvSatellite{}))
}

// IcpeOpmSyncFsmState represents Icpe opm sync fsm state
type IcpeOpmSyncFsmState string

const (
    // Split brain
    IcpeOpmSyncFsmState_icpe_opm_sync_fsm_state_split_brain IcpeOpmSyncFsmState = "icpe-opm-sync-fsm-state-split-brain"

    // Waiting
    IcpeOpmSyncFsmState_icpe_opm_sync_fsm_state_waiting IcpeOpmSyncFsmState = "icpe-opm-sync-fsm-state-waiting"

    // Whole brain
    IcpeOpmSyncFsmState_icpe_opm_sync_fsm_state_whole_brain IcpeOpmSyncFsmState = "icpe-opm-sync-fsm-state-whole-brain"
)

// IcpeOperSdacpSessState represents Icpe oper sdacp sess state
type IcpeOperSdacpSessState string

const (
    // Not created
    IcpeOperSdacpSessState_icpe_oper_sdacp_sess_state_not_created IcpeOperSdacpSessState = "icpe-oper-sdacp-sess-state-not-created"

    // Created
    IcpeOperSdacpSessState_icpe_oper_sdacp_sess_state_created IcpeOperSdacpSessState = "icpe-oper-sdacp-sess-state-created"

    // Authentication not OK
    IcpeOperSdacpSessState_icpe_oper_sdacp_sess_state_authentication_not_ok IcpeOperSdacpSessState = "icpe-oper-sdacp-sess-state-authentication-not-ok"

    // Authentication OK
    IcpeOperSdacpSessState_icpe_oper_sdacp_sess_state_authentication_ok IcpeOperSdacpSessState = "icpe-oper-sdacp-sess-state-authentication-ok"

    // Version not OK
    IcpeOperSdacpSessState_icpe_oper_sdacp_sess_state_version_not_ok IcpeOperSdacpSessState = "icpe-oper-sdacp-sess-state-version-not-ok"

    // Up
    IcpeOperSdacpSessState_icpe_oper_sdacp_sess_state_up IcpeOperSdacpSessState = "icpe-oper-sdacp-sess-state-up"

    // ISSU
    IcpeOperSdacpSessState_icpe_oper_sdacp_sess_state_issu IcpeOperSdacpSessState = "icpe-oper-sdacp-sess-state-issu"
)

// IcpeOpmTransportState represents Icpe opm transport state
type IcpeOpmTransportState string

const (
    // Disconnected
    IcpeOpmTransportState_icpe_opm_transport_state_disconnected IcpeOpmTransportState = "icpe-opm-transport-state-disconnected"

    // ICCP unavailable
    IcpeOpmTransportState_icpe_opm_transport_state_iccp_unavailable IcpeOpmTransportState = "icpe-opm-transport-state-iccp-unavailable"

    // No member present
    IcpeOpmTransportState_icpe_opm_transport_state_no_member_present IcpeOpmTransportState = "icpe-opm-transport-state-no-member-present"

    // Member down
    IcpeOpmTransportState_icpe_opm_transport_state_member_down IcpeOpmTransportState = "icpe-opm-transport-state-member-down"

    // Member not reachable
    IcpeOpmTransportState_icpe_opm_transport_state_member_not_reachable IcpeOpmTransportState = "icpe-opm-transport-state-member-not-reachable"

    // Waiting for app connect
    IcpeOpmTransportState_icpe_opm_transport_state_waiting_for_app_connect IcpeOpmTransportState = "icpe-opm-transport-state-waiting-for-app-connect"

    // Waiting for app connect response
    IcpeOpmTransportState_icpe_opm_transport_state_waiting_for_app_connect_response IcpeOpmTransportState = "icpe-opm-transport-state-waiting-for-app-connect-response"

    // Connected
    IcpeOpmTransportState_icpe_opm_transport_state_connected IcpeOpmTransportState = "icpe-opm-transport-state-connected"
)

// IcpeOpmAuthFsmState represents Icpe opm auth fsm state
type IcpeOpmAuthFsmState string

const (
    // Unauth
    IcpeOpmAuthFsmState_icpe_opm_auth_fsm_state_unauth IcpeOpmAuthFsmState = "icpe-opm-auth-fsm-state-unauth"

    // Waiting
    IcpeOpmAuthFsmState_icpe_opm_auth_fsm_state_waiting IcpeOpmAuthFsmState = "icpe-opm-auth-fsm-state-waiting"

    // Waiting for auth
    IcpeOpmAuthFsmState_icpe_opm_auth_fsm_state_waiting_for_auth IcpeOpmAuthFsmState = "icpe-opm-auth-fsm-state-waiting-for-auth"

    // Waiting for reply
    IcpeOpmAuthFsmState_icpe_opm_auth_fsm_state_waiting_for_reply IcpeOpmAuthFsmState = "icpe-opm-auth-fsm-state-waiting-for-reply"

    // Authed
    IcpeOpmAuthFsmState_icpe_opm_auth_fsm_state_authed IcpeOpmAuthFsmState = "icpe-opm-auth-fsm-state-authed"
)

// IcpeOpmController represents Icpe opm controller
type IcpeOpmController string

const (
    // Unknown
    IcpeOpmController_icpe_opm_controller_unknown IcpeOpmController = "icpe-opm-controller-unknown"

    // Primary
    IcpeOpmController_icpe_opm_controller_primary IcpeOpmController = "icpe-opm-controller-primary"

    // Secondary
    IcpeOpmController_icpe_opm_controller_secondary IcpeOpmController = "icpe-opm-controller-secondary"
)

// IcpeOpmResyncFsmState represents Icpe opm resync fsm state
type IcpeOpmResyncFsmState string

const (
    // Not open
    IcpeOpmResyncFsmState_icpe_opm_resync_fsm_state_not_open IcpeOpmResyncFsmState = "icpe-opm-resync-fsm-state-not-open"

    // Stable
    IcpeOpmResyncFsmState_icpe_opm_resync_fsm_state_stable IcpeOpmResyncFsmState = "icpe-opm-resync-fsm-state-stable"

    // In resync
    IcpeOpmResyncFsmState_icpe_opm_resync_fsm_state_in_resync IcpeOpmResyncFsmState = "icpe-opm-resync-fsm-state-in-resync"

    // Queued
    IcpeOpmResyncFsmState_icpe_opm_resync_fsm_state_queued IcpeOpmResyncFsmState = "icpe-opm-resync-fsm-state-queued"

    // Resync req
    IcpeOpmResyncFsmState_icpe_opm_resync_fsm_state_resync_req IcpeOpmResyncFsmState = "icpe-opm-resync-fsm-state-resync-req"
)

// IcpeOpmChanFsmState represents Icpe opm chan fsm state
type IcpeOpmChanFsmState string

const (
    // Down
    IcpeOpmChanFsmState_icpe_opm_chan_fsm_state_down IcpeOpmChanFsmState = "icpe-opm-chan-fsm-state-down"

    // Closed
    IcpeOpmChanFsmState_icpe_opm_chan_fsm_state_closed IcpeOpmChanFsmState = "icpe-opm-chan-fsm-state-closed"

    // Opening
    IcpeOpmChanFsmState_icpe_opm_chan_fsm_state_opening IcpeOpmChanFsmState = "icpe-opm-chan-fsm-state-opening"

    // Opened
    IcpeOpmChanFsmState_icpe_opm_chan_fsm_state_opened IcpeOpmChanFsmState = "icpe-opm-chan-fsm-state-opened"

    // Open
    IcpeOpmChanFsmState_icpe_opm_chan_fsm_state_open IcpeOpmChanFsmState = "icpe-opm-chan-fsm-state-open"
)

// IcpeOpmSessState represents Icpe opm sess state
type IcpeOpmSessState string

const (
    // Disconnected
    IcpeOpmSessState_icpe_opm_sess_state_disconnected IcpeOpmSessState = "icpe-opm-sess-state-disconnected"

    // Connecting
    IcpeOpmSessState_icpe_opm_sess_state_connecting IcpeOpmSessState = "icpe-opm-sess-state-connecting"

    // Authenticating
    IcpeOpmSessState_icpe_opm_sess_state_authenticating IcpeOpmSessState = "icpe-opm-sess-state-authenticating"

    // Arbitrating
    IcpeOpmSessState_icpe_opm_sess_state_arbitrating IcpeOpmSessState = "icpe-opm-sess-state-arbitrating"

    // Waiting for resyncs
    IcpeOpmSessState_icpe_opm_sess_state_waiting_for_resyncs IcpeOpmSessState = "icpe-opm-sess-state-waiting-for-resyncs"

    // Connected
    IcpeOpmSessState_icpe_opm_sess_state_connected IcpeOpmSessState = "icpe-opm-sess-state-connected"
)

// IcpeOperInstallState represents Icpe oper install state
type IcpeOperInstallState string

const (
    // Stable
    IcpeOperInstallState_icpe_oper_install_state_stable IcpeOperInstallState = "icpe-oper-install-state-stable"

    // Transferring
    IcpeOperInstallState_icpe_oper_install_state_transferring IcpeOperInstallState = "icpe-oper-install-state-transferring"

    // Transferred
    IcpeOperInstallState_icpe_oper_install_state_transferred IcpeOperInstallState = "icpe-oper-install-state-transferred"

    // Installing
    IcpeOperInstallState_icpe_oper_install_state_installing IcpeOperInstallState = "icpe-oper-install-state-installing"

    // In progress
    IcpeOperInstallState_icpe_oper_install_state_in_progress IcpeOperInstallState = "icpe-oper-install-state-in-progress"
)

// IcpeOperPort represents Icpe oper port
type IcpeOperPort string

const (
    // Unknown
    IcpeOperPort_icpe_oper_port_unknown IcpeOperPort = "icpe-oper-port-unknown"

    // Gigabit ethernet
    IcpeOperPort_icpe_oper_port_gigabit_ethernet IcpeOperPort = "icpe-oper-port-gigabit-ethernet"

    // Ten gig e
    IcpeOperPort_icpe_oper_port_ten_gig_e IcpeOperPort = "icpe-oper-port-ten-gig-e"
)

// IcpeOperFabricPort represents Icpe oper fabric port
type IcpeOperFabricPort string

const (
    // Unknown
    IcpeOperFabricPort_icpe_oper_fabric_port_unknown IcpeOperFabricPort = "icpe-oper-fabric-port-unknown"

    // n v fabric- gig e
    IcpeOperFabricPort_icpe_oper_fabric_port_n_v_fabric_gig_e IcpeOperFabricPort = "icpe-oper-fabric-port-n-v-fabric-gig-e"

    // n v fabric- ten gig e
    IcpeOperFabricPort_icpe_oper_fabric_port_n_v_fabric_ten_gig_e IcpeOperFabricPort = "icpe-oper-fabric-port-n-v-fabric-ten-gig-e"

    // n v fabric- hundred gig e
    IcpeOperFabricPort_icpe_oper_fabric_port_n_v_fabric_hundred_gig_e IcpeOperFabricPort = "icpe-oper-fabric-port-n-v-fabric-hundred-gig-e"
)

// IcpeInstallPkgSupp represents Icpe install pkg supp
type IcpeInstallPkgSupp string

const (
    // Unknown
    IcpeInstallPkgSupp_icpe_install_pkg_supp_unknown IcpeInstallPkgSupp = "icpe-install-pkg-supp-unknown"

    // Not supported
    IcpeInstallPkgSupp_icpe_install_pkg_supp_not_supported IcpeInstallPkgSupp = "icpe-install-pkg-supp-not-supported"

    // Supported
    IcpeInstallPkgSupp_icpe_install_pkg_supp_supported IcpeInstallPkgSupp = "icpe-install-pkg-supp-supported"
)

// IcpeGcoOperControlReason represents Icpe gco oper control reason
type IcpeGcoOperControlReason string

const (
    // Unknown error
    IcpeGcoOperControlReason_icpe_gco_oper_control_reason_unknown_error IcpeGcoOperControlReason = "icpe-gco-oper-control-reason-unknown-error"

    // Wrong chassis type
    IcpeGcoOperControlReason_icpe_gco_oper_control_reason_wrong_chassis_type IcpeGcoOperControlReason = "icpe-gco-oper-control-reason-wrong-chassis-type"

    // Wrong chassis serial
    IcpeGcoOperControlReason_icpe_gco_oper_control_reason_wrong_chassis_serial IcpeGcoOperControlReason = "icpe-gco-oper-control-reason-wrong-chassis-serial"

    // Needs to upgrade
    IcpeGcoOperControlReason_icpe_gco_oper_control_reason_needs_to_upgrade IcpeGcoOperControlReason = "icpe-gco-oper-control-reason-needs-to-upgrade"

    // None
    IcpeGcoOperControlReason_icpe_gco_oper_control_reason_none IcpeGcoOperControlReason = "icpe-gco-oper-control-reason-none"
)

// IcpeInstallSatState represents Icpe install sat state
type IcpeInstallSatState string

const (
    // Unknown
    IcpeInstallSatState_icpe_install_sat_state_unknown IcpeInstallSatState = "icpe-install-sat-state-unknown"

    // Not initiated
    IcpeInstallSatState_icpe_install_sat_state_not_initiat_ed IcpeInstallSatState = "icpe-install-sat-state-not-initiat-ed"

    // Transferring
    IcpeInstallSatState_icpe_install_sat_state_transferring IcpeInstallSatState = "icpe-install-sat-state-transferring"

    // Activating
    IcpeInstallSatState_icpe_install_sat_state_activating IcpeInstallSatState = "icpe-install-sat-state-activating"

    // Updating
    IcpeInstallSatState_icpe_install_sat_state_updating IcpeInstallSatState = "icpe-install-sat-state-updating"

    // Deactivating
    IcpeInstallSatState_icpe_install_sat_state_deactivating IcpeInstallSatState = "icpe-install-sat-state-deactivating"

    // Removing
    IcpeInstallSatState_icpe_install_sat_state_removing IcpeInstallSatState = "icpe-install-sat-state-removing"

    // Success
    IcpeInstallSatState_icpe_install_sat_state_success IcpeInstallSatState = "icpe-install-sat-state-success"

    // Failure
    IcpeInstallSatState_icpe_install_sat_state_failure IcpeInstallSatState = "icpe-install-sat-state-failure"

    // Multiple ops
    IcpeInstallSatState_icpe_install_sat_state_multiple_ops IcpeInstallSatState = "icpe-install-sat-state-multiple-ops"

    // Aborted
    IcpeInstallSatState_icpe_install_sat_state_aborted IcpeInstallSatState = "icpe-install-sat-state-aborted"

    // Protocol version
    IcpeInstallSatState_icpe_install_sat_state_protocol_version IcpeInstallSatState = "icpe-install-sat-state-protocol-version"

    // Pkg not present
    IcpeInstallSatState_icpe_install_sat_state_pkg_not_present IcpeInstallSatState = "icpe-install-sat-state-pkg-not-present"

    // No image
    IcpeInstallSatState_icpe_install_sat_state_no_image IcpeInstallSatState = "icpe-install-sat-state-no-image"

    // No such file
    IcpeInstallSatState_icpe_install_sat_state_no_such_file IcpeInstallSatState = "icpe-install-sat-state-no-such-file"

    // Sat uncfgd
    IcpeInstallSatState_icpe_install_sat_state_sat_uncfgd IcpeInstallSatState = "icpe-install-sat-state-sat-uncfgd"

    // Processing
    IcpeInstallSatState_icpe_install_sat_state_processing IcpeInstallSatState = "icpe-install-sat-state-processing"
)

// IcpeOpmArbitrationFsmState represents Icpe opm arbitration fsm state
type IcpeOpmArbitrationFsmState string

const (
    // Unarbitrated
    IcpeOpmArbitrationFsmState_icpe_opm_arbitration_fsm_state_unarbitrated IcpeOpmArbitrationFsmState = "icpe-opm-arbitration-fsm-state-unarbitrated"

    // Waiting
    IcpeOpmArbitrationFsmState_icpe_opm_arbitration_fsm_state_waiting IcpeOpmArbitrationFsmState = "icpe-opm-arbitration-fsm-state-waiting"

    // Arbitrating
    IcpeOpmArbitrationFsmState_icpe_opm_arbitration_fsm_state_arbitrating IcpeOpmArbitrationFsmState = "icpe-opm-arbitration-fsm-state-arbitrating"

    // Arbitrated
    IcpeOpmArbitrationFsmState_icpe_opm_arbitration_fsm_state_arbitrated IcpeOpmArbitrationFsmState = "icpe-opm-arbitration-fsm-state-arbitrated"
)

// IcpeOperConflict represents Icpe oper conflict
type IcpeOperConflict string

const (
    // Not calculated
    IcpeOperConflict_icpe_oper_conflict_not_calculated IcpeOperConflict = "icpe-oper-conflict-not-calculated"

    // No conflict
    IcpeOperConflict_icpe_oper_conflict_no_conflict IcpeOperConflict = "icpe-oper-conflict-no-conflict"

    // Satellite not configured
    IcpeOperConflict_icpe_oper_conflict_satellite_not_configured IcpeOperConflict = "icpe-oper-conflict-satellite-not-configured"

    // Satellite no type
    IcpeOperConflict_icpe_oper_conflict_satellite_no_type IcpeOperConflict = "icpe-oper-conflict-satellite-no-type"

    // Satellite id invalid
    IcpeOperConflict_icpe_oper_conflict_satellite_id_invalid IcpeOperConflict = "icpe-oper-conflict-satellite-id-invalid"

    // Satellite no ipv4 addr
    IcpeOperConflict_icpe_oper_conflict_satellite_no_ipv4_addr IcpeOperConflict = "icpe-oper-conflict-satellite-no-ipv4-addr"

    // Satellite conflicting ipv4 addr
    IcpeOperConflict_icpe_oper_conflict_satellite_conflicting_ipv4_addr IcpeOperConflict = "icpe-oper-conflict-satellite-conflicting-ipv4-addr"

    // No configured links
    IcpeOperConflict_icpe_oper_conflict_no_configured_links IcpeOperConflict = "icpe-oper-conflict-no-configured-links"

    // No discovered links
    IcpeOperConflict_icpe_oper_conflict_no_discovered_links IcpeOperConflict = "icpe-oper-conflict-no-discovered-links"

    // Invalid ports
    IcpeOperConflict_icpe_oper_conflict_invalid_ports IcpeOperConflict = "icpe-oper-conflict-invalid-ports"

    // Ports overlap
    IcpeOperConflict_icpe_oper_conflict_ports_overlap IcpeOperConflict = "icpe-oper-conflict-ports-overlap"

    // Waiting for ipv4 addr
    IcpeOperConflict_icpe_oper_conflict_waiting_for_ipv4_addr IcpeOperConflict = "icpe-oper-conflict-waiting-for-ipv4-addr"

    // Waiting for VRF
    IcpeOperConflict_icpe_oper_conflict_waiting_for_vrf IcpeOperConflict = "icpe-oper-conflict-waiting-for-vrf"

    // Different ipv4 addr
    IcpeOperConflict_icpe_oper_conflict_different_ipv4_addr IcpeOperConflict = "icpe-oper-conflict-different-ipv4-addr"

    // Different VRF
    IcpeOperConflict_icpe_oper_conflict_different_vrf IcpeOperConflict = "icpe-oper-conflict-different-vrf"

    // Satellite link ipv4 overlap
    IcpeOperConflict_icpe_oper_conflict_satellite_link_ipv4_overlap IcpeOperConflict = "icpe-oper-conflict-satellite-link-ipv4-overlap"

    // Waiting for ident
    IcpeOperConflict_icpe_oper_conflict_waiting_for_ident IcpeOperConflict = "icpe-oper-conflict-waiting-for-ident"

    // Multiple ids
    IcpeOperConflict_icpe_oper_conflict_multiple_ids IcpeOperConflict = "icpe-oper-conflict-multiple-ids"

    // Multiple satellites
    IcpeOperConflict_icpe_oper_conflict_multiple_satellites IcpeOperConflict = "icpe-oper-conflict-multiple-satellites"

    // Ident rejected
    IcpeOperConflict_icpe_oper_conflict_ident_rejected IcpeOperConflict = "icpe-oper-conflict-ident-rejected"

    // Interface down
    IcpeOperConflict_icpe_oper_conflict_interface_down IcpeOperConflict = "icpe-oper-conflict-interface-down"

    // Auto IP unavailable
    IcpeOperConflict_icpe_oper_conflict_auto_ip_unavailable IcpeOperConflict = "icpe-oper-conflict-auto-ip-unavailable"

    // Satellite auto IP link manual IP
    IcpeOperConflict_icpe_oper_conflict_satellite_auto_ip_link_manual_ip IcpeOperConflict = "icpe-oper-conflict-satellite-auto-ip-link-manual-ip"

    // Link auto IP satellite manual IP
    IcpeOperConflict_icpe_oper_conflict_link_auto_ip_satellite_manual_ip IcpeOperConflict = "icpe-oper-conflict-link-auto-ip-satellite-manual-ip"

    // Serial num mismatch
    IcpeOperConflict_icpe_oper_conflict_serial_num_mismatch IcpeOperConflict = "icpe-oper-conflict-serial-num-mismatch"

    // Satellite not identified
    IcpeOperConflict_icpe_oper_conflict_satellite_not_identified IcpeOperConflict = "icpe-oper-conflict-satellite-not-identified"

    // Satellite unsupported type
    IcpeOperConflict_icpe_oper_conflict_satellite_unsupported_type IcpeOperConflict = "icpe-oper-conflict-satellite-unsupported-type"

    // Satellite partition unsupported
    IcpeOperConflict_icpe_oper_conflict_satellite_partition_unsupported IcpeOperConflict = "icpe-oper-conflict-satellite-partition-unsupported"

    // Satellite no serial number
    IcpeOperConflict_icpe_oper_conflict_satellite_no_serial_number IcpeOperConflict = "icpe-oper-conflict-satellite-no-serial-number"

    // Satellite conflicting serial number
    IcpeOperConflict_icpe_oper_conflict_satellite_conflicting_serial_number IcpeOperConflict = "icpe-oper-conflict-satellite-conflicting-serial-number"

    // Satellite link waiting for arp
    IcpeOperConflict_icpe_oper_conflict_satellite_link_waiting_for_arp IcpeOperConflict = "icpe-oper-conflict-satellite-link-waiting-for-arp"

    // Host PE isolated split brain
    IcpeOperConflict_icpe_oper_conflict_host_pe_isolated_split_brain IcpeOperConflict = "icpe-oper-conflict-host-pe-isolated-split-brain"

    // Fabric ICCP group inconsistent
    IcpeOperConflict_icpe_oper_conflict_fabric_iccp_group_inconsistent IcpeOperConflict = "icpe-oper-conflict-fabric-iccp-group-inconsistent"

    // Invalid ICCP group
    IcpeOperConflict_icpe_oper_conflict_invalid_iccp_group IcpeOperConflict = "icpe-oper-conflict-invalid-iccp-group"

    // Port rejected
    IcpeOperConflict_icpe_oper_conflict_port_rejected IcpeOperConflict = "icpe-oper-conflict-port-rejected"

    // Satellite ICL not supported
    IcpeOperConflict_icpe_oper_conflict_satellite_icl_not_supported IcpeOperConflict = "icpe-oper-conflict-satellite-icl-not-supported"

    // Multiple serial number
    IcpeOperConflict_icpe_oper_conflict_multiple_serial_number IcpeOperConflict = "icpe-oper-conflict-multiple-serial-number"

    // Multiple MAC address
    IcpeOperConflict_icpe_oper_conflict_multiple_mac_address IcpeOperConflict = "icpe-oper-conflict-multiple-mac-address"
)

// IcpeOperVerCheckState represents Icpe oper ver check state
type IcpeOperVerCheckState string

const (
    // Unknown
    IcpeOperVerCheckState_icpe_oper_ver_check_state_unknown IcpeOperVerCheckState = "icpe-oper-ver-check-state-unknown"

    // Not compatible
    IcpeOperVerCheckState_icpe_oper_ver_check_state_not_compatible IcpeOperVerCheckState = "icpe-oper-ver-check-state-not-compatible"

    // Current version
    IcpeOperVerCheckState_icpe_oper_ver_check_state_current_version IcpeOperVerCheckState = "icpe-oper-ver-check-state-current-version"

    // Compatible older
    IcpeOperVerCheckState_icpe_oper_ver_check_state_compatible_older IcpeOperVerCheckState = "icpe-oper-ver-check-state-compatible-older"

    // Compatible newer
    IcpeOperVerCheckState_icpe_oper_ver_check_state_compatible_newer IcpeOperVerCheckState = "icpe-oper-ver-check-state-compatible-newer"
)

// IcpeOperMultichassisRedundancy represents Icpe oper multichassis redundancy
type IcpeOperMultichassisRedundancy string

const (
    // Not redundant
    IcpeOperMultichassisRedundancy_icpe_oper_multi_chassis_redundancy_not_redundant IcpeOperMultichassisRedundancy = "icpe-oper-multi-chassis-redundancy-not-redundant"

    // Active
    IcpeOperMultichassisRedundancy_icpe_oper_multi_chassis_redundancy_active IcpeOperMultichassisRedundancy = "icpe-oper-multi-chassis-redundancy-active"

    // Standby
    IcpeOperMultichassisRedundancy_icpe_oper_multi_chassis_redundancy_standby IcpeOperMultichassisRedundancy = "icpe-oper-multi-chassis-redundancy-standby"
)

// IcpeOperDiscdLinkState represents Icpe oper discd link state
type IcpeOperDiscdLinkState string

const (
    // Stopped
    IcpeOperDiscdLinkState_icpe_oper_discd_link_state_stopped IcpeOperDiscdLinkState = "icpe-oper-discd-link-state-stopped"

    // Probing
    IcpeOperDiscdLinkState_icpe_oper_discd_link_state_probing IcpeOperDiscdLinkState = "icpe-oper-discd-link-state-probing"

    // Configuring
    IcpeOperDiscdLinkState_icpe_oper_discd_link_state_configuring IcpeOperDiscdLinkState = "icpe-oper-discd-link-state-configuring"

    // Ready
    IcpeOperDiscdLinkState_icpe_oper_discd_link_state_ready IcpeOperDiscdLinkState = "icpe-oper-discd-link-state-ready"
)

// IcpeOperTopoRemoteSource represents Icpe oper topo remote source
type IcpeOperTopoRemoteSource string

const (
    // Unknown
    IcpeOperTopoRemoteSource_icpe_oper_topo_remote_source_unknown IcpeOperTopoRemoteSource = "icpe-oper-topo-remote-source-unknown"

    // Remote ICL ID
    IcpeOperTopoRemoteSource_icpe_oper_topo_remote_source_remote_icl_id IcpeOperTopoRemoteSource = "icpe-oper-topo-remote-source-remote-icl-id"

    // Remote satellite MAC
    IcpeOperTopoRemoteSource_icpe_oper_topo_remote_source_remote_satellite_mac IcpeOperTopoRemoteSource = "icpe-oper-topo-remote-source-remote-satellite-mac"

    // Remote host MAC
    IcpeOperTopoRemoteSource_icpe_oper_topo_remote_source_remote_host_mac IcpeOperTopoRemoteSource = "icpe-oper-topo-remote-source-remote-host-mac"

    // Direct satellite
    IcpeOperTopoRemoteSource_icpe_oper_topo_remote_source_direct_satellite IcpeOperTopoRemoteSource = "icpe-oper-topo-remote-source-direct-satellite"

    // Direct host
    IcpeOperTopoRemoteSource_icpe_oper_topo_remote_source_direct_host IcpeOperTopoRemoteSource = "icpe-oper-topo-remote-source-direct-host"
)

// IcpeOpticalSyncState represents Icpe optical sync state
type IcpeOpticalSyncState string

const (
    // Unknown
    IcpeOpticalSyncState_icpe_optical_sync_state_unknown IcpeOpticalSyncState = "icpe-optical-sync-state-unknown"

    // Syncing
    IcpeOpticalSyncState_icpe_optical_sync_state_syncing IcpeOpticalSyncState = "icpe-optical-sync-state-syncing"

    // Synced
    IcpeOpticalSyncState_icpe_optical_sync_state_synced IcpeOpticalSyncState = "icpe-optical-sync-state-synced"

    // Not connected
    IcpeOpticalSyncState_icpe_optical_sync_state_not_connected IcpeOpticalSyncState = "icpe-optical-sync-state-not-connected"
)

// NvSatellite
// Satellite operational information
type NvSatellite struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed breakdown of reload status table.
    ReloadOpStatuses NvSatellite_ReloadOpStatuses

    // nV Satellite Redundancy Protocol Information table.
    SdacpRedundancies NvSatellite_SdacpRedundancies

    // Detailed breakdown of install status table.
    InstallShows NvSatellite_InstallShows

    // Satellite status information table.
    SatelliteStatuses NvSatellite_SatelliteStatuses

    // Satellite priority information table.
    SatellitePriorities NvSatellite_SatellitePriorities

    // Satellite remote version information table.
    SatelliteVersions NvSatellite_SatelliteVersions

    // Satellite Topology Information table.
    SatelliteTopologies NvSatellite_SatelliteTopologies

    // Satellite Install Reference Information.
    InstallReferenceInfo NvSatellite_InstallReferenceInfo

    // Current percentage of install table.
    InstallOpProgresses NvSatellite_InstallOpProgresses

    // Detailed breakdown of reload status table.
    ReloadStatuses NvSatellite_ReloadStatuses

    // Satellite Install Information.
    Install NvSatellite_Install

    // Detailed breakdown of install status table.
    InstallOpStatuses NvSatellite_InstallOpStatuses

    // ICPE GCO operational information.
    SatelliteProperties NvSatellite_SatelliteProperties

    // ICPE Configured interface state information table.
    SdacpDiscovery2S NvSatellite_SdacpDiscovery2S

    // ICPE DPM operational information table.
    IcpeDpms NvSatellite_IcpeDpms

    // SDAC Protocol Discovery information table.
    SdacpControls NvSatellite_SdacpControls
}

func (nvSatellite *NvSatellite) GetFilter() yfilter.YFilter { return nvSatellite.YFilter }

func (nvSatellite *NvSatellite) SetFilter(yf yfilter.YFilter) { nvSatellite.YFilter = yf }

func (nvSatellite *NvSatellite) GetGoName(yname string) string {
    if yname == "reload-op-statuses" { return "ReloadOpStatuses" }
    if yname == "sdacp-redundancies" { return "SdacpRedundancies" }
    if yname == "install-shows" { return "InstallShows" }
    if yname == "satellite-statuses" { return "SatelliteStatuses" }
    if yname == "satellite-priorities" { return "SatellitePriorities" }
    if yname == "satellite-versions" { return "SatelliteVersions" }
    if yname == "satellite-topologies" { return "SatelliteTopologies" }
    if yname == "install-reference-info" { return "InstallReferenceInfo" }
    if yname == "install-op-progresses" { return "InstallOpProgresses" }
    if yname == "reload-statuses" { return "ReloadStatuses" }
    if yname == "install" { return "Install" }
    if yname == "install-op-statuses" { return "InstallOpStatuses" }
    if yname == "satellite-properties" { return "SatelliteProperties" }
    if yname == "Cisco-IOS-XR-icpe-sdacp-oper:sdacp-discovery2s" { return "SdacpDiscovery2S" }
    if yname == "Cisco-IOS-XR-icpe-sdacp-oper:icpe-dpms" { return "IcpeDpms" }
    if yname == "Cisco-IOS-XR-icpe-sdacp-oper:sdacp-controls" { return "SdacpControls" }
    return ""
}

func (nvSatellite *NvSatellite) GetSegmentPath() string {
    return "Cisco-IOS-XR-icpe-infra-oper:nv-satellite"
}

func (nvSatellite *NvSatellite) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reload-op-statuses" {
        return &nvSatellite.ReloadOpStatuses
    }
    if childYangName == "sdacp-redundancies" {
        return &nvSatellite.SdacpRedundancies
    }
    if childYangName == "install-shows" {
        return &nvSatellite.InstallShows
    }
    if childYangName == "satellite-statuses" {
        return &nvSatellite.SatelliteStatuses
    }
    if childYangName == "satellite-priorities" {
        return &nvSatellite.SatellitePriorities
    }
    if childYangName == "satellite-versions" {
        return &nvSatellite.SatelliteVersions
    }
    if childYangName == "satellite-topologies" {
        return &nvSatellite.SatelliteTopologies
    }
    if childYangName == "install-reference-info" {
        return &nvSatellite.InstallReferenceInfo
    }
    if childYangName == "install-op-progresses" {
        return &nvSatellite.InstallOpProgresses
    }
    if childYangName == "reload-statuses" {
        return &nvSatellite.ReloadStatuses
    }
    if childYangName == "install" {
        return &nvSatellite.Install
    }
    if childYangName == "install-op-statuses" {
        return &nvSatellite.InstallOpStatuses
    }
    if childYangName == "satellite-properties" {
        return &nvSatellite.SatelliteProperties
    }
    if childYangName == "Cisco-IOS-XR-icpe-sdacp-oper:sdacp-discovery2s" {
        return &nvSatellite.SdacpDiscovery2S
    }
    if childYangName == "Cisco-IOS-XR-icpe-sdacp-oper:icpe-dpms" {
        return &nvSatellite.IcpeDpms
    }
    if childYangName == "Cisco-IOS-XR-icpe-sdacp-oper:sdacp-controls" {
        return &nvSatellite.SdacpControls
    }
    return nil
}

func (nvSatellite *NvSatellite) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["reload-op-statuses"] = &nvSatellite.ReloadOpStatuses
    children["sdacp-redundancies"] = &nvSatellite.SdacpRedundancies
    children["install-shows"] = &nvSatellite.InstallShows
    children["satellite-statuses"] = &nvSatellite.SatelliteStatuses
    children["satellite-priorities"] = &nvSatellite.SatellitePriorities
    children["satellite-versions"] = &nvSatellite.SatelliteVersions
    children["satellite-topologies"] = &nvSatellite.SatelliteTopologies
    children["install-reference-info"] = &nvSatellite.InstallReferenceInfo
    children["install-op-progresses"] = &nvSatellite.InstallOpProgresses
    children["reload-statuses"] = &nvSatellite.ReloadStatuses
    children["install"] = &nvSatellite.Install
    children["install-op-statuses"] = &nvSatellite.InstallOpStatuses
    children["satellite-properties"] = &nvSatellite.SatelliteProperties
    children["Cisco-IOS-XR-icpe-sdacp-oper:sdacp-discovery2s"] = &nvSatellite.SdacpDiscovery2S
    children["Cisco-IOS-XR-icpe-sdacp-oper:icpe-dpms"] = &nvSatellite.IcpeDpms
    children["Cisco-IOS-XR-icpe-sdacp-oper:sdacp-controls"] = &nvSatellite.SdacpControls
    return children
}

func (nvSatellite *NvSatellite) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nvSatellite *NvSatellite) GetBundleName() string { return "cisco_ios_xr" }

func (nvSatellite *NvSatellite) GetYangName() string { return "nv-satellite" }

func (nvSatellite *NvSatellite) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nvSatellite *NvSatellite) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nvSatellite *NvSatellite) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nvSatellite *NvSatellite) SetParent(parent types.Entity) { nvSatellite.parent = parent }

func (nvSatellite *NvSatellite) GetParent() types.Entity { return nvSatellite.parent }

func (nvSatellite *NvSatellite) GetParentYangName() string { return "Cisco-IOS-XR-icpe-infra-oper" }

// NvSatellite_ReloadOpStatuses
// Detailed breakdown of reload status table
type NvSatellite_ReloadOpStatuses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed breakdown of reload status. The type is slice of
    // NvSatellite_ReloadOpStatuses_ReloadOpStatus.
    ReloadOpStatus []NvSatellite_ReloadOpStatuses_ReloadOpStatus
}

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetFilter() yfilter.YFilter { return reloadOpStatuses.YFilter }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) SetFilter(yf yfilter.YFilter) { reloadOpStatuses.YFilter = yf }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetGoName(yname string) string {
    if yname == "reload-op-status" { return "ReloadOpStatus" }
    return ""
}

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetSegmentPath() string {
    return "reload-op-statuses"
}

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reload-op-status" {
        for _, c := range reloadOpStatuses.ReloadOpStatus {
            if reloadOpStatuses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_ReloadOpStatuses_ReloadOpStatus{}
        reloadOpStatuses.ReloadOpStatus = append(reloadOpStatuses.ReloadOpStatus, child)
        return &reloadOpStatuses.ReloadOpStatus[len(reloadOpStatuses.ReloadOpStatus)-1]
    }
    return nil
}

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range reloadOpStatuses.ReloadOpStatus {
        children[reloadOpStatuses.ReloadOpStatus[i].GetSegmentPath()] = &reloadOpStatuses.ReloadOpStatus[i]
    }
    return children
}

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetBundleName() string { return "cisco_ios_xr" }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetYangName() string { return "reload-op-statuses" }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) SetParent(parent types.Entity) { reloadOpStatuses.parent = parent }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetParent() types.Entity { return reloadOpStatuses.parent }

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_ReloadOpStatuses_ReloadOpStatus
// Detailed breakdown of reload status
type NvSatellite_ReloadOpStatuses_ReloadOpStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Operation ID. The type is interface{} with range:
    // 0..4294967295.
    OperationId interface{}

    // Operation ID. The type is interface{} with range: 0..4294967295.
    OperationIdXr interface{}

    // Satellite range. The type is string.
    SatelliteRange interface{}

    // Sats not initiated. The type is slice of interface{} with range: 0..65535.
    SatsNotInitiated []interface{}

    // Sats reloading. The type is slice of interface{} with range: 0..65535.
    SatsReloading []interface{}

    // Sats reloaded. The type is slice of interface{} with range: 0..65535.
    SatsReloaded []interface{}

    // Sats reload failed. The type is slice of interface{} with range: 0..65535.
    SatsReloadFailed []interface{}
}

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetFilter() yfilter.YFilter { return reloadOpStatus.YFilter }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) SetFilter(yf yfilter.YFilter) { reloadOpStatus.YFilter = yf }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetGoName(yname string) string {
    if yname == "operation-id" { return "OperationId" }
    if yname == "operation-id-xr" { return "OperationIdXr" }
    if yname == "satellite-range" { return "SatelliteRange" }
    if yname == "sats-not-initiated" { return "SatsNotInitiated" }
    if yname == "sats-reloading" { return "SatsReloading" }
    if yname == "sats-reloaded" { return "SatsReloaded" }
    if yname == "sats-reload-failed" { return "SatsReloadFailed" }
    return ""
}

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetSegmentPath() string {
    return "reload-op-status" + "[operation-id='" + fmt.Sprintf("%v", reloadOpStatus.OperationId) + "']"
}

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-id"] = reloadOpStatus.OperationId
    leafs["operation-id-xr"] = reloadOpStatus.OperationIdXr
    leafs["satellite-range"] = reloadOpStatus.SatelliteRange
    leafs["sats-not-initiated"] = reloadOpStatus.SatsNotInitiated
    leafs["sats-reloading"] = reloadOpStatus.SatsReloading
    leafs["sats-reloaded"] = reloadOpStatus.SatsReloaded
    leafs["sats-reload-failed"] = reloadOpStatus.SatsReloadFailed
    return leafs
}

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetBundleName() string { return "cisco_ios_xr" }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetYangName() string { return "reload-op-status" }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) SetParent(parent types.Entity) { reloadOpStatus.parent = parent }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetParent() types.Entity { return reloadOpStatus.parent }

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetParentYangName() string { return "reload-op-statuses" }

// NvSatellite_SdacpRedundancies
// nV Satellite Redundancy Protocol Information
// table
type NvSatellite_SdacpRedundancies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // nV Satellite Redundancy Protocol Information. The type is slice of
    // NvSatellite_SdacpRedundancies_SdacpRedundancy.
    SdacpRedundancy []NvSatellite_SdacpRedundancies_SdacpRedundancy
}

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetFilter() yfilter.YFilter { return sdacpRedundancies.YFilter }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) SetFilter(yf yfilter.YFilter) { sdacpRedundancies.YFilter = yf }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetGoName(yname string) string {
    if yname == "sdacp-redundancy" { return "SdacpRedundancy" }
    return ""
}

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetSegmentPath() string {
    return "sdacp-redundancies"
}

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sdacp-redundancy" {
        for _, c := range sdacpRedundancies.SdacpRedundancy {
            if sdacpRedundancies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpRedundancies_SdacpRedundancy{}
        sdacpRedundancies.SdacpRedundancy = append(sdacpRedundancies.SdacpRedundancy, child)
        return &sdacpRedundancies.SdacpRedundancy[len(sdacpRedundancies.SdacpRedundancy)-1]
    }
    return nil
}

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sdacpRedundancies.SdacpRedundancy {
        children[sdacpRedundancies.SdacpRedundancy[i].GetSegmentPath()] = &sdacpRedundancies.SdacpRedundancy[i]
    }
    return children
}

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetBundleName() string { return "cisco_ios_xr" }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetYangName() string { return "sdacp-redundancies" }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) SetParent(parent types.Entity) { sdacpRedundancies.parent = parent }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetParent() types.Entity { return sdacpRedundancies.parent }

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_SdacpRedundancies_SdacpRedundancy
// nV Satellite Redundancy Protocol Information
type NvSatellite_SdacpRedundancies_SdacpRedundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ICCP group. The type is interface{} with range:
    // 0..4294967295.
    IccpGroup interface{}

    // ICCP group. The type is interface{} with range: 0..4294967295.
    IccpGroupXr interface{}

    // Protocol state. The type is IcpeOpmSessState.
    ProtocolState interface{}

    // Transport state. The type is IcpeOpmTransportState.
    TransportState interface{}

    // Authentication state. The type is IcpeOpmAuthFsmState.
    AuthenticationState interface{}

    // Arbitration state. The type is IcpeOpmArbitrationFsmState.
    ArbitrationState interface{}

    // Synchronization state. The type is IcpeOpmSyncFsmState.
    SynchronizationState interface{}

    // Primacy. The type is IcpeOpmController.
    Primacy interface{}

    // System MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SystemMac interface{}

    // Isolated. The type is bool.
    Isolated interface{}

    // Timestamp.
    ProtocolStateTimestamp NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp

    // Channels on this session table. The type is slice of
    // NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel.
    Channel []NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel
}

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetFilter() yfilter.YFilter { return sdacpRedundancy.YFilter }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) SetFilter(yf yfilter.YFilter) { sdacpRedundancy.YFilter = yf }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetGoName(yname string) string {
    if yname == "iccp-group" { return "IccpGroup" }
    if yname == "iccp-group-xr" { return "IccpGroupXr" }
    if yname == "protocol-state" { return "ProtocolState" }
    if yname == "transport-state" { return "TransportState" }
    if yname == "authentication-state" { return "AuthenticationState" }
    if yname == "arbitration-state" { return "ArbitrationState" }
    if yname == "synchronization-state" { return "SynchronizationState" }
    if yname == "primacy" { return "Primacy" }
    if yname == "system-mac" { return "SystemMac" }
    if yname == "isolated" { return "Isolated" }
    if yname == "protocol-state-timestamp" { return "ProtocolStateTimestamp" }
    if yname == "channel" { return "Channel" }
    return ""
}

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetSegmentPath() string {
    return "sdacp-redundancy" + "[iccp-group='" + fmt.Sprintf("%v", sdacpRedundancy.IccpGroup) + "']"
}

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol-state-timestamp" {
        return &sdacpRedundancy.ProtocolStateTimestamp
    }
    if childYangName == "channel" {
        for _, c := range sdacpRedundancy.Channel {
            if sdacpRedundancy.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel{}
        sdacpRedundancy.Channel = append(sdacpRedundancy.Channel, child)
        return &sdacpRedundancy.Channel[len(sdacpRedundancy.Channel)-1]
    }
    return nil
}

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocol-state-timestamp"] = &sdacpRedundancy.ProtocolStateTimestamp
    for i := range sdacpRedundancy.Channel {
        children[sdacpRedundancy.Channel[i].GetSegmentPath()] = &sdacpRedundancy.Channel[i]
    }
    return children
}

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["iccp-group"] = sdacpRedundancy.IccpGroup
    leafs["iccp-group-xr"] = sdacpRedundancy.IccpGroupXr
    leafs["protocol-state"] = sdacpRedundancy.ProtocolState
    leafs["transport-state"] = sdacpRedundancy.TransportState
    leafs["authentication-state"] = sdacpRedundancy.AuthenticationState
    leafs["arbitration-state"] = sdacpRedundancy.ArbitrationState
    leafs["synchronization-state"] = sdacpRedundancy.SynchronizationState
    leafs["primacy"] = sdacpRedundancy.Primacy
    leafs["system-mac"] = sdacpRedundancy.SystemMac
    leafs["isolated"] = sdacpRedundancy.Isolated
    return leafs
}

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetBundleName() string { return "cisco_ios_xr" }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetYangName() string { return "sdacp-redundancy" }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) SetParent(parent types.Entity) { sdacpRedundancy.parent = parent }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetParent() types.Entity { return sdacpRedundancy.parent }

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetParentYangName() string { return "sdacp-redundancies" }

// NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp
// Timestamp
type NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetFilter() yfilter.YFilter { return protocolStateTimestamp.YFilter }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) SetFilter(yf yfilter.YFilter) { protocolStateTimestamp.YFilter = yf }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetSegmentPath() string {
    return "protocol-state-timestamp"
}

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = protocolStateTimestamp.Seconds
    leafs["nanoseconds"] = protocolStateTimestamp.Nanoseconds
    return leafs
}

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetBundleName() string { return "cisco_ios_xr" }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetYangName() string { return "protocol-state-timestamp" }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) SetParent(parent types.Entity) { protocolStateTimestamp.parent = parent }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetParent() types.Entity { return protocolStateTimestamp.parent }

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetParentYangName() string { return "sdacp-redundancy" }

// NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel
// Channels on this session table
type NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Channel ID. The type is interface{} with range: 0..4294967295.
    ChannelId interface{}

    // Chan state. The type is IcpeOpmChanFsmState.
    ChanState interface{}

    // Resync state. The type is IcpeOpmResyncFsmState.
    ResyncState interface{}

    // Control messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ControlMessagesSent interface{}

    // Normal messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    NormalMessagesSent interface{}

    // Control messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    ControlMessagesReceived interface{}

    // Normal messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    NormalMessagesReceived interface{}

    // Timestamp.
    ChannelStateTimestamp NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp

    // Timestamp.
    ResyncStateTimestamp NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp
}

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetFilter() yfilter.YFilter { return channel.YFilter }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) SetFilter(yf yfilter.YFilter) { channel.YFilter = yf }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetGoName(yname string) string {
    if yname == "channel-id" { return "ChannelId" }
    if yname == "chan-state" { return "ChanState" }
    if yname == "resync-state" { return "ResyncState" }
    if yname == "control-messages-sent" { return "ControlMessagesSent" }
    if yname == "normal-messages-sent" { return "NormalMessagesSent" }
    if yname == "control-messages-received" { return "ControlMessagesReceived" }
    if yname == "normal-messages-received" { return "NormalMessagesReceived" }
    if yname == "channel-state-timestamp" { return "ChannelStateTimestamp" }
    if yname == "resync-state-timestamp" { return "ResyncStateTimestamp" }
    return ""
}

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetSegmentPath() string {
    return "channel"
}

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "channel-state-timestamp" {
        return &channel.ChannelStateTimestamp
    }
    if childYangName == "resync-state-timestamp" {
        return &channel.ResyncStateTimestamp
    }
    return nil
}

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["channel-state-timestamp"] = &channel.ChannelStateTimestamp
    children["resync-state-timestamp"] = &channel.ResyncStateTimestamp
    return children
}

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["channel-id"] = channel.ChannelId
    leafs["chan-state"] = channel.ChanState
    leafs["resync-state"] = channel.ResyncState
    leafs["control-messages-sent"] = channel.ControlMessagesSent
    leafs["normal-messages-sent"] = channel.NormalMessagesSent
    leafs["control-messages-received"] = channel.ControlMessagesReceived
    leafs["normal-messages-received"] = channel.NormalMessagesReceived
    return leafs
}

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetBundleName() string { return "cisco_ios_xr" }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetYangName() string { return "channel" }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) SetParent(parent types.Entity) { channel.parent = parent }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetParent() types.Entity { return channel.parent }

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetParentYangName() string { return "sdacp-redundancy" }

// NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp
// Timestamp
type NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetFilter() yfilter.YFilter { return channelStateTimestamp.YFilter }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) SetFilter(yf yfilter.YFilter) { channelStateTimestamp.YFilter = yf }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetSegmentPath() string {
    return "channel-state-timestamp"
}

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = channelStateTimestamp.Seconds
    leafs["nanoseconds"] = channelStateTimestamp.Nanoseconds
    return leafs
}

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetBundleName() string { return "cisco_ios_xr" }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetYangName() string { return "channel-state-timestamp" }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) SetParent(parent types.Entity) { channelStateTimestamp.parent = parent }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetParent() types.Entity { return channelStateTimestamp.parent }

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetParentYangName() string { return "channel" }

// NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp
// Timestamp
type NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetFilter() yfilter.YFilter { return resyncStateTimestamp.YFilter }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) SetFilter(yf yfilter.YFilter) { resyncStateTimestamp.YFilter = yf }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetSegmentPath() string {
    return "resync-state-timestamp"
}

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resyncStateTimestamp.Seconds
    leafs["nanoseconds"] = resyncStateTimestamp.Nanoseconds
    return leafs
}

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetBundleName() string { return "cisco_ios_xr" }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetYangName() string { return "resync-state-timestamp" }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) SetParent(parent types.Entity) { resyncStateTimestamp.parent = parent }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetParent() types.Entity { return resyncStateTimestamp.parent }

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetParentYangName() string { return "channel" }

// NvSatellite_InstallShows
// Detailed breakdown of install status table
type NvSatellite_InstallShows struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed breakdown of install status. The type is slice of
    // NvSatellite_InstallShows_InstallShow.
    InstallShow []NvSatellite_InstallShows_InstallShow
}

func (installShows *NvSatellite_InstallShows) GetFilter() yfilter.YFilter { return installShows.YFilter }

func (installShows *NvSatellite_InstallShows) SetFilter(yf yfilter.YFilter) { installShows.YFilter = yf }

func (installShows *NvSatellite_InstallShows) GetGoName(yname string) string {
    if yname == "install-show" { return "InstallShow" }
    return ""
}

func (installShows *NvSatellite_InstallShows) GetSegmentPath() string {
    return "install-shows"
}

func (installShows *NvSatellite_InstallShows) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "install-show" {
        for _, c := range installShows.InstallShow {
            if installShows.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_InstallShows_InstallShow{}
        installShows.InstallShow = append(installShows.InstallShow, child)
        return &installShows.InstallShow[len(installShows.InstallShow)-1]
    }
    return nil
}

func (installShows *NvSatellite_InstallShows) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range installShows.InstallShow {
        children[installShows.InstallShow[i].GetSegmentPath()] = &installShows.InstallShow[i]
    }
    return children
}

func (installShows *NvSatellite_InstallShows) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (installShows *NvSatellite_InstallShows) GetBundleName() string { return "cisco_ios_xr" }

func (installShows *NvSatellite_InstallShows) GetYangName() string { return "install-shows" }

func (installShows *NvSatellite_InstallShows) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (installShows *NvSatellite_InstallShows) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (installShows *NvSatellite_InstallShows) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (installShows *NvSatellite_InstallShows) SetParent(parent types.Entity) { installShows.parent = parent }

func (installShows *NvSatellite_InstallShows) GetParent() types.Entity { return installShows.parent }

func (installShows *NvSatellite_InstallShows) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_InstallShows_InstallShow
// Detailed breakdown of install status
type NvSatellite_InstallShows_InstallShow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Operation ID. The type is interface{} with range:
    // 0..4294967295.
    OperationId interface{}

    // Operation ID. The type is interface{} with range: 0..4294967295.
    OperationIdXr interface{}

    // Satellite range. The type is string.
    SatelliteRange interface{}

    // Operation type. The type is interface{} with range: 0..65535.
    OperationType interface{}

    // Progress percentage. The type is interface{} with range: 0..65535. Units
    // are percentage.
    ProgressPercentage interface{}

    // Start time. The type is interface{} with range: 0..4294967295.
    StartTime interface{}

    // End time. The type is interface{} with range: 0..4294967295.
    EndTime interface{}

    // Ref state. The type is IcpeInstallSatState.
    RefState interface{}

    // Sats not initiated. The type is slice of interface{} with range: 0..65535.
    SatsNotInitiated []interface{}

    // Sats transferring. The type is slice of interface{} with range: 0..65535.
    SatsTransferring []interface{}

    // Sats activating. The type is slice of interface{} with range: 0..65535.
    SatsActivating []interface{}

    // Sats updating. The type is slice of interface{} with range: 0..65535.
    SatsUpdating []interface{}

    // Sats deactivating. The type is slice of interface{} with range: 0..65535.
    SatsDeactivating []interface{}

    // Sats removing. The type is slice of interface{} with range: 0..65535.
    SatsRemoving []interface{}

    // Sats transfer failed. The type is slice of interface{} with range:
    // 0..65535.
    SatsTransferFailed []interface{}

    // Sats activate failed. The type is slice of interface{} with range:
    // 0..65535.
    SatsActivateFailed []interface{}

    // Sats update failed. The type is slice of interface{} with range: 0..65535.
    SatsUpdateFailed []interface{}

    // Sats deactivate failed. The type is slice of interface{} with range:
    // 0..65535.
    SatsDeactivateFailed []interface{}

    // Sats remove failed. The type is slice of interface{} with range: 0..65535.
    SatsRemoveFailed []interface{}

    // Sats transfer aborted. The type is slice of interface{} with range:
    // 0..65535.
    SatsTransferAborted []interface{}

    // Sats activate aborted. The type is slice of interface{} with range:
    // 0..65535.
    SatsActivateAborted []interface{}

    // Sats update aborted. The type is slice of interface{} with range: 0..65535.
    SatsUpdateAborted []interface{}

    // Sats deactivate aborted. The type is slice of interface{} with range:
    // 0..65535.
    SatsDeactivateAborted []interface{}

    // Sats remove aborted. The type is slice of interface{} with range: 0..65535.
    SatsRemoveAborted []interface{}

    // Sats no operation. The type is slice of interface{} with range: 0..65535.
    SatsNoOperation []interface{}

    // Sats completed. The type is slice of interface{} with range: 0..65535.
    SatsCompleted []interface{}

    // Name strings. The type is slice of string.
    NameString []interface{}

    // Breakdown per satellite table. The type is slice of
    // NvSatellite_InstallShows_InstallShow_Satellite.
    Satellite []NvSatellite_InstallShows_InstallShow_Satellite
}

func (installShow *NvSatellite_InstallShows_InstallShow) GetFilter() yfilter.YFilter { return installShow.YFilter }

func (installShow *NvSatellite_InstallShows_InstallShow) SetFilter(yf yfilter.YFilter) { installShow.YFilter = yf }

func (installShow *NvSatellite_InstallShows_InstallShow) GetGoName(yname string) string {
    if yname == "operation-id" { return "OperationId" }
    if yname == "operation-id-xr" { return "OperationIdXr" }
    if yname == "satellite-range" { return "SatelliteRange" }
    if yname == "operation-type" { return "OperationType" }
    if yname == "progress-percentage" { return "ProgressPercentage" }
    if yname == "start-time" { return "StartTime" }
    if yname == "end-time" { return "EndTime" }
    if yname == "ref-state" { return "RefState" }
    if yname == "sats-not-initiated" { return "SatsNotInitiated" }
    if yname == "sats-transferring" { return "SatsTransferring" }
    if yname == "sats-activating" { return "SatsActivating" }
    if yname == "sats-updating" { return "SatsUpdating" }
    if yname == "sats-deactivating" { return "SatsDeactivating" }
    if yname == "sats-removing" { return "SatsRemoving" }
    if yname == "sats-transfer-failed" { return "SatsTransferFailed" }
    if yname == "sats-activate-failed" { return "SatsActivateFailed" }
    if yname == "sats-update-failed" { return "SatsUpdateFailed" }
    if yname == "sats-deactivate-failed" { return "SatsDeactivateFailed" }
    if yname == "sats-remove-failed" { return "SatsRemoveFailed" }
    if yname == "sats-transfer-aborted" { return "SatsTransferAborted" }
    if yname == "sats-activate-aborted" { return "SatsActivateAborted" }
    if yname == "sats-update-aborted" { return "SatsUpdateAborted" }
    if yname == "sats-deactivate-aborted" { return "SatsDeactivateAborted" }
    if yname == "sats-remove-aborted" { return "SatsRemoveAborted" }
    if yname == "sats-no-operation" { return "SatsNoOperation" }
    if yname == "sats-completed" { return "SatsCompleted" }
    if yname == "name-string" { return "NameString" }
    if yname == "satellite" { return "Satellite" }
    return ""
}

func (installShow *NvSatellite_InstallShows_InstallShow) GetSegmentPath() string {
    return "install-show" + "[operation-id='" + fmt.Sprintf("%v", installShow.OperationId) + "']"
}

func (installShow *NvSatellite_InstallShows_InstallShow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "satellite" {
        for _, c := range installShow.Satellite {
            if installShow.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_InstallShows_InstallShow_Satellite{}
        installShow.Satellite = append(installShow.Satellite, child)
        return &installShow.Satellite[len(installShow.Satellite)-1]
    }
    return nil
}

func (installShow *NvSatellite_InstallShows_InstallShow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range installShow.Satellite {
        children[installShow.Satellite[i].GetSegmentPath()] = &installShow.Satellite[i]
    }
    return children
}

func (installShow *NvSatellite_InstallShows_InstallShow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-id"] = installShow.OperationId
    leafs["operation-id-xr"] = installShow.OperationIdXr
    leafs["satellite-range"] = installShow.SatelliteRange
    leafs["operation-type"] = installShow.OperationType
    leafs["progress-percentage"] = installShow.ProgressPercentage
    leafs["start-time"] = installShow.StartTime
    leafs["end-time"] = installShow.EndTime
    leafs["ref-state"] = installShow.RefState
    leafs["sats-not-initiated"] = installShow.SatsNotInitiated
    leafs["sats-transferring"] = installShow.SatsTransferring
    leafs["sats-activating"] = installShow.SatsActivating
    leafs["sats-updating"] = installShow.SatsUpdating
    leafs["sats-deactivating"] = installShow.SatsDeactivating
    leafs["sats-removing"] = installShow.SatsRemoving
    leafs["sats-transfer-failed"] = installShow.SatsTransferFailed
    leafs["sats-activate-failed"] = installShow.SatsActivateFailed
    leafs["sats-update-failed"] = installShow.SatsUpdateFailed
    leafs["sats-deactivate-failed"] = installShow.SatsDeactivateFailed
    leafs["sats-remove-failed"] = installShow.SatsRemoveFailed
    leafs["sats-transfer-aborted"] = installShow.SatsTransferAborted
    leafs["sats-activate-aborted"] = installShow.SatsActivateAborted
    leafs["sats-update-aborted"] = installShow.SatsUpdateAborted
    leafs["sats-deactivate-aborted"] = installShow.SatsDeactivateAborted
    leafs["sats-remove-aborted"] = installShow.SatsRemoveAborted
    leafs["sats-no-operation"] = installShow.SatsNoOperation
    leafs["sats-completed"] = installShow.SatsCompleted
    leafs["name-string"] = installShow.NameString
    return leafs
}

func (installShow *NvSatellite_InstallShows_InstallShow) GetBundleName() string { return "cisco_ios_xr" }

func (installShow *NvSatellite_InstallShows_InstallShow) GetYangName() string { return "install-show" }

func (installShow *NvSatellite_InstallShows_InstallShow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (installShow *NvSatellite_InstallShows_InstallShow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (installShow *NvSatellite_InstallShows_InstallShow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (installShow *NvSatellite_InstallShows_InstallShow) SetParent(parent types.Entity) { installShow.parent = parent }

func (installShow *NvSatellite_InstallShows_InstallShow) GetParent() types.Entity { return installShow.parent }

func (installShow *NvSatellite_InstallShows_InstallShow) GetParentYangName() string { return "install-shows" }

// NvSatellite_InstallShows_InstallShow_Satellite
// Breakdown per satellite table
type NvSatellite_InstallShows_InstallShow_Satellite struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteId interface{}

    // State. The type is IcpeInstallSatState.
    State interface{}

    // Percentage. The type is interface{} with range: 0..65535. Units are
    // percentage.
    Percentage interface{}

    // Retries. The type is interface{} with range: 0..65535.
    Retries interface{}

    // Start time. The type is interface{} with range: 0..4294967295.
    StartTime interface{}

    // End time. The type is interface{} with range: 0..4294967295.
    EndTime interface{}

    // Info. The type is string.
    Info interface{}
}

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetFilter() yfilter.YFilter { return satellite.YFilter }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) SetFilter(yf yfilter.YFilter) { satellite.YFilter = yf }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "state" { return "State" }
    if yname == "percentage" { return "Percentage" }
    if yname == "retries" { return "Retries" }
    if yname == "start-time" { return "StartTime" }
    if yname == "end-time" { return "EndTime" }
    if yname == "info" { return "Info" }
    return ""
}

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetSegmentPath() string {
    return "satellite"
}

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = satellite.SatelliteId
    leafs["state"] = satellite.State
    leafs["percentage"] = satellite.Percentage
    leafs["retries"] = satellite.Retries
    leafs["start-time"] = satellite.StartTime
    leafs["end-time"] = satellite.EndTime
    leafs["info"] = satellite.Info
    return leafs
}

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetBundleName() string { return "cisco_ios_xr" }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetYangName() string { return "satellite" }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) SetParent(parent types.Entity) { satellite.parent = parent }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetParent() types.Entity { return satellite.parent }

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetParentYangName() string { return "install-show" }

// NvSatellite_SatelliteStatuses
// Satellite status information table
type NvSatellite_SatelliteStatuses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite status information. The type is slice of
    // NvSatellite_SatelliteStatuses_SatelliteStatus.
    SatelliteStatus []NvSatellite_SatelliteStatuses_SatelliteStatus
}

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetFilter() yfilter.YFilter { return satelliteStatuses.YFilter }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) SetFilter(yf yfilter.YFilter) { satelliteStatuses.YFilter = yf }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetGoName(yname string) string {
    if yname == "satellite-status" { return "SatelliteStatus" }
    return ""
}

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetSegmentPath() string {
    return "satellite-statuses"
}

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "satellite-status" {
        for _, c := range satelliteStatuses.SatelliteStatus {
            if satelliteStatuses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteStatuses_SatelliteStatus{}
        satelliteStatuses.SatelliteStatus = append(satelliteStatuses.SatelliteStatus, child)
        return &satelliteStatuses.SatelliteStatus[len(satelliteStatuses.SatelliteStatus)-1]
    }
    return nil
}

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range satelliteStatuses.SatelliteStatus {
        children[satelliteStatuses.SatelliteStatus[i].GetSegmentPath()] = &satelliteStatuses.SatelliteStatus[i]
    }
    return children
}

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetYangName() string { return "satellite-statuses" }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) SetParent(parent types.Entity) { satelliteStatuses.parent = parent }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetParent() types.Entity { return satelliteStatuses.parent }

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_SatelliteStatuses_SatelliteStatus
// Satellite status information
type NvSatellite_SatelliteStatuses_SatelliteStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite ID. The type is interface{} with range:
    // 100..65534.
    SatelliteId interface{}

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteIdXr interface{}

    // Version check state. The type is IcpeOperVerCheckState.
    VersionCheckState interface{}

    // Remote version present. The type is bool.
    RemoteVersionPresent interface{}

    // Type. The type is string.
    Type interface{}

    // Ethernet fabric supported. The type is bool.
    EthernetFabricSupported interface{}

    // Optical supported. The type is bool.
    OpticalSupported interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // IP address present. The type is bool.
    IpAddressPresent interface{}

    // IP address auto. The type is bool.
    IpAddressAuto interface{}

    // IPV6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // IPV6 address present. The type is bool.
    Ipv6AddressPresent interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // VRFID. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}

    // Description. The type is string.
    Description interface{}

    // Description present. The type is bool.
    DescriptionPresent interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // MAC address present. The type is bool.
    MacAddressPresent interface{}

    // Configured serial number. The type is string.
    ConfiguredSerialNumber interface{}

    // Configured serial number present. The type is bool.
    ConfiguredSerialNumberPresent interface{}

    // Received serial number. The type is string.
    ReceivedSerialNumber interface{}

    // Received serial number present. The type is bool.
    ReceivedSerialNumberPresent interface{}

    // Password. The type is string.
    Password interface{}

    // Password error. The type is string.
    PasswordError interface{}

    // Received hostname. The type is string.
    ReceivedHostName interface{}

    // Cfgd timeout. The type is interface{} with range: 0..4294967295.
    CfgdTimeout interface{}

    // Timeout warning. The type is interface{} with range: 0..4294967295.
    TimeoutWarning interface{}

    // Conflict reason. The type is IcpeOperConflict.
    ConflictReason interface{}

    // Conflict context. The type is string.
    ConflictContext interface{}

    // Redundancy ICCP group. The type is interface{} with range: 0..4294967295.
    RedundancyIccpGroup interface{}

    // Recovery delay time left. The type is interface{} with range: 0..65535.
    RecoveryDelayTimeLeft interface{}

    // Host treating as active. The type is bool.
    HostTreatingAsActive interface{}

    // Satellite treating as active. The type is bool.
    SatelliteTreatingAsActive interface{}

    // SDACP session state. The type is IcpeOperSdacpSessState.
    SdacpSessionState interface{}

    // SDACP session failure reason. The type is IcpeGcoOperControlReason.
    SdacpSessionFailureReason interface{}

    // Install state. The type is IcpeOperInstallState.
    InstallState interface{}

    // Remote version. The type is slice of string.
    RemoteVersion []interface{}

    // Candidate Fabric Ports on this Satellite.
    CandidateFabricPorts NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts

    // Optical Satellite Status.
    OpticalStatus NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus

    // Timestamp.
    RedundancyOutOfSyncTimestamp NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp

    // Configured Links on this Satellite table. The type is slice of
    // NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink.
    ConfiguredLink []NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink
}

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetFilter() yfilter.YFilter { return satelliteStatus.YFilter }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) SetFilter(yf yfilter.YFilter) { satelliteStatus.YFilter = yf }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "satellite-id-xr" { return "SatelliteIdXr" }
    if yname == "version-check-state" { return "VersionCheckState" }
    if yname == "remote-version-present" { return "RemoteVersionPresent" }
    if yname == "type" { return "Type" }
    if yname == "ethernet-fabric-supported" { return "EthernetFabricSupported" }
    if yname == "optical-supported" { return "OpticalSupported" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "ip-address-present" { return "IpAddressPresent" }
    if yname == "ip-address-auto" { return "IpAddressAuto" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "ipv6-address-present" { return "Ipv6AddressPresent" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrfid" { return "Vrfid" }
    if yname == "description" { return "Description" }
    if yname == "description-present" { return "DescriptionPresent" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "mac-address-present" { return "MacAddressPresent" }
    if yname == "configured-serial-number" { return "ConfiguredSerialNumber" }
    if yname == "configured-serial-number-present" { return "ConfiguredSerialNumberPresent" }
    if yname == "received-serial-number" { return "ReceivedSerialNumber" }
    if yname == "received-serial-number-present" { return "ReceivedSerialNumberPresent" }
    if yname == "password" { return "Password" }
    if yname == "password-error" { return "PasswordError" }
    if yname == "received-host-name" { return "ReceivedHostName" }
    if yname == "cfgd-timeout" { return "CfgdTimeout" }
    if yname == "timeout-warning" { return "TimeoutWarning" }
    if yname == "conflict-reason" { return "ConflictReason" }
    if yname == "conflict-context" { return "ConflictContext" }
    if yname == "redundancy-iccp-group" { return "RedundancyIccpGroup" }
    if yname == "recovery-delay-time-left" { return "RecoveryDelayTimeLeft" }
    if yname == "host-treating-as-active" { return "HostTreatingAsActive" }
    if yname == "satellite-treating-as-active" { return "SatelliteTreatingAsActive" }
    if yname == "sdacp-session-state" { return "SdacpSessionState" }
    if yname == "sdacp-session-failure-reason" { return "SdacpSessionFailureReason" }
    if yname == "install-state" { return "InstallState" }
    if yname == "remote-version" { return "RemoteVersion" }
    if yname == "candidate-fabric-ports" { return "CandidateFabricPorts" }
    if yname == "optical-status" { return "OpticalStatus" }
    if yname == "redundancy-out-of-sync-timestamp" { return "RedundancyOutOfSyncTimestamp" }
    if yname == "configured-link" { return "ConfiguredLink" }
    return ""
}

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetSegmentPath() string {
    return "satellite-status" + "[satellite-id='" + fmt.Sprintf("%v", satelliteStatus.SatelliteId) + "']"
}

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-fabric-ports" {
        return &satelliteStatus.CandidateFabricPorts
    }
    if childYangName == "optical-status" {
        return &satelliteStatus.OpticalStatus
    }
    if childYangName == "redundancy-out-of-sync-timestamp" {
        return &satelliteStatus.RedundancyOutOfSyncTimestamp
    }
    if childYangName == "configured-link" {
        for _, c := range satelliteStatus.ConfiguredLink {
            if satelliteStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink{}
        satelliteStatus.ConfiguredLink = append(satelliteStatus.ConfiguredLink, child)
        return &satelliteStatus.ConfiguredLink[len(satelliteStatus.ConfiguredLink)-1]
    }
    return nil
}

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["candidate-fabric-ports"] = &satelliteStatus.CandidateFabricPorts
    children["optical-status"] = &satelliteStatus.OpticalStatus
    children["redundancy-out-of-sync-timestamp"] = &satelliteStatus.RedundancyOutOfSyncTimestamp
    for i := range satelliteStatus.ConfiguredLink {
        children[satelliteStatus.ConfiguredLink[i].GetSegmentPath()] = &satelliteStatus.ConfiguredLink[i]
    }
    return children
}

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = satelliteStatus.SatelliteId
    leafs["satellite-id-xr"] = satelliteStatus.SatelliteIdXr
    leafs["version-check-state"] = satelliteStatus.VersionCheckState
    leafs["remote-version-present"] = satelliteStatus.RemoteVersionPresent
    leafs["type"] = satelliteStatus.Type
    leafs["ethernet-fabric-supported"] = satelliteStatus.EthernetFabricSupported
    leafs["optical-supported"] = satelliteStatus.OpticalSupported
    leafs["ip-address"] = satelliteStatus.IpAddress
    leafs["ip-address-present"] = satelliteStatus.IpAddressPresent
    leafs["ip-address-auto"] = satelliteStatus.IpAddressAuto
    leafs["ipv6-address"] = satelliteStatus.Ipv6Address
    leafs["ipv6-address-present"] = satelliteStatus.Ipv6AddressPresent
    leafs["vrf-name"] = satelliteStatus.VrfName
    leafs["vrfid"] = satelliteStatus.Vrfid
    leafs["description"] = satelliteStatus.Description
    leafs["description-present"] = satelliteStatus.DescriptionPresent
    leafs["mac-address"] = satelliteStatus.MacAddress
    leafs["mac-address-present"] = satelliteStatus.MacAddressPresent
    leafs["configured-serial-number"] = satelliteStatus.ConfiguredSerialNumber
    leafs["configured-serial-number-present"] = satelliteStatus.ConfiguredSerialNumberPresent
    leafs["received-serial-number"] = satelliteStatus.ReceivedSerialNumber
    leafs["received-serial-number-present"] = satelliteStatus.ReceivedSerialNumberPresent
    leafs["password"] = satelliteStatus.Password
    leafs["password-error"] = satelliteStatus.PasswordError
    leafs["received-host-name"] = satelliteStatus.ReceivedHostName
    leafs["cfgd-timeout"] = satelliteStatus.CfgdTimeout
    leafs["timeout-warning"] = satelliteStatus.TimeoutWarning
    leafs["conflict-reason"] = satelliteStatus.ConflictReason
    leafs["conflict-context"] = satelliteStatus.ConflictContext
    leafs["redundancy-iccp-group"] = satelliteStatus.RedundancyIccpGroup
    leafs["recovery-delay-time-left"] = satelliteStatus.RecoveryDelayTimeLeft
    leafs["host-treating-as-active"] = satelliteStatus.HostTreatingAsActive
    leafs["satellite-treating-as-active"] = satelliteStatus.SatelliteTreatingAsActive
    leafs["sdacp-session-state"] = satelliteStatus.SdacpSessionState
    leafs["sdacp-session-failure-reason"] = satelliteStatus.SdacpSessionFailureReason
    leafs["install-state"] = satelliteStatus.InstallState
    leafs["remote-version"] = satelliteStatus.RemoteVersion
    return leafs
}

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetYangName() string { return "satellite-status" }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) SetParent(parent types.Entity) { satelliteStatus.parent = parent }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetParent() types.Entity { return satelliteStatus.parent }

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetParentYangName() string { return "satellite-statuses" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts
// Candidate Fabric Ports on this Satellite
type NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Channel up. The type is bool.
    ChannelUp interface{}

    // Out of sync. The type is bool.
    OutOfSync interface{}

    // Error string. The type is string.
    ErrorString interface{}

    // Configured Candidate Fabric Ports table. The type is slice of
    // NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort.
    ConfiguredPort []NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort

    // Current Candidate Fabric Ports on this Satellite table. The type is slice
    // of
    // NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort.
    CurrentPort []NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort
}

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetFilter() yfilter.YFilter { return candidateFabricPorts.YFilter }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) SetFilter(yf yfilter.YFilter) { candidateFabricPorts.YFilter = yf }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetGoName(yname string) string {
    if yname == "channel-up" { return "ChannelUp" }
    if yname == "out-of-sync" { return "OutOfSync" }
    if yname == "error-string" { return "ErrorString" }
    if yname == "configured-port" { return "ConfiguredPort" }
    if yname == "current-port" { return "CurrentPort" }
    return ""
}

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetSegmentPath() string {
    return "candidate-fabric-ports"
}

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-port" {
        for _, c := range candidateFabricPorts.ConfiguredPort {
            if candidateFabricPorts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort{}
        candidateFabricPorts.ConfiguredPort = append(candidateFabricPorts.ConfiguredPort, child)
        return &candidateFabricPorts.ConfiguredPort[len(candidateFabricPorts.ConfiguredPort)-1]
    }
    if childYangName == "current-port" {
        for _, c := range candidateFabricPorts.CurrentPort {
            if candidateFabricPorts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort{}
        candidateFabricPorts.CurrentPort = append(candidateFabricPorts.CurrentPort, child)
        return &candidateFabricPorts.CurrentPort[len(candidateFabricPorts.CurrentPort)-1]
    }
    return nil
}

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range candidateFabricPorts.ConfiguredPort {
        children[candidateFabricPorts.ConfiguredPort[i].GetSegmentPath()] = &candidateFabricPorts.ConfiguredPort[i]
    }
    for i := range candidateFabricPorts.CurrentPort {
        children[candidateFabricPorts.CurrentPort[i].GetSegmentPath()] = &candidateFabricPorts.CurrentPort[i]
    }
    return children
}

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["channel-up"] = candidateFabricPorts.ChannelUp
    leafs["out-of-sync"] = candidateFabricPorts.OutOfSync
    leafs["error-string"] = candidateFabricPorts.ErrorString
    return leafs
}

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetBundleName() string { return "cisco_ios_xr" }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetYangName() string { return "candidate-fabric-ports" }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) SetParent(parent types.Entity) { candidateFabricPorts.parent = parent }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetParent() types.Entity { return candidateFabricPorts.parent }

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetParentYangName() string { return "satellite-status" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort
// Configured Candidate Fabric Ports table
type NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port type. The type is IcpeOperFabricPort.
    PortType interface{}

    // Slot. The type is interface{} with range: 0..65535.
    Slot interface{}

    // Subslot. The type is interface{} with range: 0..65535.
    Subslot interface{}

    // Port. The type is interface{} with range: 0..65535.
    Port interface{}

    // Valid. The type is bool.
    Valid interface{}
}

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetFilter() yfilter.YFilter { return configuredPort.YFilter }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) SetFilter(yf yfilter.YFilter) { configuredPort.YFilter = yf }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetGoName(yname string) string {
    if yname == "port-type" { return "PortType" }
    if yname == "slot" { return "Slot" }
    if yname == "subslot" { return "Subslot" }
    if yname == "port" { return "Port" }
    if yname == "valid" { return "Valid" }
    return ""
}

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetSegmentPath() string {
    return "configured-port"
}

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-type"] = configuredPort.PortType
    leafs["slot"] = configuredPort.Slot
    leafs["subslot"] = configuredPort.Subslot
    leafs["port"] = configuredPort.Port
    leafs["valid"] = configuredPort.Valid
    return leafs
}

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetBundleName() string { return "cisco_ios_xr" }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetYangName() string { return "configured-port" }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) SetParent(parent types.Entity) { configuredPort.parent = parent }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetParent() types.Entity { return configuredPort.parent }

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetParentYangName() string { return "candidate-fabric-ports" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort
// Current Candidate Fabric Ports on this Satellite
// table
type NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port type. The type is IcpeOperFabricPort.
    PortType interface{}

    // Slot. The type is interface{} with range: 0..65535.
    Slot interface{}

    // Subslot. The type is interface{} with range: 0..65535.
    Subslot interface{}

    // Port. The type is interface{} with range: 0..65535.
    Port interface{}

    // Permanent. The type is bool.
    Permanent interface{}

    // Requested. The type is bool.
    Requested interface{}
}

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetFilter() yfilter.YFilter { return currentPort.YFilter }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) SetFilter(yf yfilter.YFilter) { currentPort.YFilter = yf }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetGoName(yname string) string {
    if yname == "port-type" { return "PortType" }
    if yname == "slot" { return "Slot" }
    if yname == "subslot" { return "Subslot" }
    if yname == "port" { return "Port" }
    if yname == "permanent" { return "Permanent" }
    if yname == "requested" { return "Requested" }
    return ""
}

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetSegmentPath() string {
    return "current-port"
}

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-type"] = currentPort.PortType
    leafs["slot"] = currentPort.Slot
    leafs["subslot"] = currentPort.Subslot
    leafs["port"] = currentPort.Port
    leafs["permanent"] = currentPort.Permanent
    leafs["requested"] = currentPort.Requested
    return leafs
}

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetBundleName() string { return "cisco_ios_xr" }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetYangName() string { return "current-port" }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) SetParent(parent types.Entity) { currentPort.parent = parent }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetParent() types.Entity { return currentPort.parent }

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetParentYangName() string { return "candidate-fabric-ports" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus
// Optical Satellite Status
type NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Chassis sync state. The type is IcpeOpticalSyncState.
    ChassisSyncState interface{}

    // Application State table. The type is slice of
    // NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application.
    Application []NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application
}

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetFilter() yfilter.YFilter { return opticalStatus.YFilter }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) SetFilter(yf yfilter.YFilter) { opticalStatus.YFilter = yf }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetGoName(yname string) string {
    if yname == "chassis-sync-state" { return "ChassisSyncState" }
    if yname == "application" { return "Application" }
    return ""
}

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetSegmentPath() string {
    return "optical-status"
}

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "application" {
        for _, c := range opticalStatus.Application {
            if opticalStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application{}
        opticalStatus.Application = append(opticalStatus.Application, child)
        return &opticalStatus.Application[len(opticalStatus.Application)-1]
    }
    return nil
}

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opticalStatus.Application {
        children[opticalStatus.Application[i].GetSegmentPath()] = &opticalStatus.Application[i]
    }
    return children
}

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-sync-state"] = opticalStatus.ChassisSyncState
    return leafs
}

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetBundleName() string { return "cisco_ios_xr" }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetYangName() string { return "optical-status" }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) SetParent(parent types.Entity) { opticalStatus.parent = parent }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetParent() types.Entity { return opticalStatus.parent }

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetParentYangName() string { return "satellite-status" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application
// Application State table
type NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is string.
    Name interface{}

    // Sync state. The type is IcpeOpticalSyncState.
    SyncState interface{}
}

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetFilter() yfilter.YFilter { return application.YFilter }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) SetFilter(yf yfilter.YFilter) { application.YFilter = yf }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "sync-state" { return "SyncState" }
    return ""
}

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetSegmentPath() string {
    return "application"
}

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = application.Name
    leafs["sync-state"] = application.SyncState
    return leafs
}

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetBundleName() string { return "cisco_ios_xr" }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetYangName() string { return "application" }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) SetParent(parent types.Entity) { application.parent = parent }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetParent() types.Entity { return application.parent }

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetParentYangName() string { return "optical-status" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp
// Timestamp
type NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetFilter() yfilter.YFilter { return redundancyOutOfSyncTimestamp.YFilter }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) SetFilter(yf yfilter.YFilter) { redundancyOutOfSyncTimestamp.YFilter = yf }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetSegmentPath() string {
    return "redundancy-out-of-sync-timestamp"
}

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = redundancyOutOfSyncTimestamp.Seconds
    leafs["nanoseconds"] = redundancyOutOfSyncTimestamp.Nanoseconds
    return leafs
}

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetBundleName() string { return "cisco_ios_xr" }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetYangName() string { return "redundancy-out-of-sync-timestamp" }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) SetParent(parent types.Entity) { redundancyOutOfSyncTimestamp.parent = parent }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetParent() types.Entity { return redundancyOutOfSyncTimestamp.parent }

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetParentYangName() string { return "satellite-status" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink
// Configured Links on this Satellite table
type NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // IP address auto. The type is bool.
    IpAddressAuto interface{}

    // VRF ID present. The type is bool.
    VrfIdPresent interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Minimum preferred links. The type is interface{} with range: 0..4294967295.
    MinimumPreferredLinks interface{}

    // Number active links. The type is interface{} with range: 0..4294967295.
    NumberActiveLinks interface{}

    // Min links satisfied. The type is bool.
    MinLinksSatisfied interface{}

    // Minimum required links. The type is interface{} with range: 0..4294967295.
    MinimumRequiredLinks interface{}

    // Required min links satisfied. The type is bool.
    RequiredMinLinksSatisfied interface{}

    // Conflict reason. The type is IcpeOperConflict.
    ConflictReason interface{}

    // Conflict context. The type is string.
    ConflictContext interface{}

    // Port ranges table. The type is slice of
    // NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange.
    PortRange []NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange

    // Discovered Links table. The type is slice of
    // NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink.
    DiscoveredLink []NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink
}

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetFilter() yfilter.YFilter { return configuredLink.YFilter }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) SetFilter(yf yfilter.YFilter) { configuredLink.YFilter = yf }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetGoName(yname string) string {
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "ip-address-auto" { return "IpAddressAuto" }
    if yname == "vrf-id-present" { return "VrfIdPresent" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "minimum-preferred-links" { return "MinimumPreferredLinks" }
    if yname == "number-active-links" { return "NumberActiveLinks" }
    if yname == "min-links-satisfied" { return "MinLinksSatisfied" }
    if yname == "minimum-required-links" { return "MinimumRequiredLinks" }
    if yname == "required-min-links-satisfied" { return "RequiredMinLinksSatisfied" }
    if yname == "conflict-reason" { return "ConflictReason" }
    if yname == "conflict-context" { return "ConflictContext" }
    if yname == "port-range" { return "PortRange" }
    if yname == "discovered-link" { return "DiscoveredLink" }
    return ""
}

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetSegmentPath() string {
    return "configured-link"
}

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-range" {
        for _, c := range configuredLink.PortRange {
            if configuredLink.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange{}
        configuredLink.PortRange = append(configuredLink.PortRange, child)
        return &configuredLink.PortRange[len(configuredLink.PortRange)-1]
    }
    if childYangName == "discovered-link" {
        for _, c := range configuredLink.DiscoveredLink {
            if configuredLink.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink{}
        configuredLink.DiscoveredLink = append(configuredLink.DiscoveredLink, child)
        return &configuredLink.DiscoveredLink[len(configuredLink.DiscoveredLink)-1]
    }
    return nil
}

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configuredLink.PortRange {
        children[configuredLink.PortRange[i].GetSegmentPath()] = &configuredLink.PortRange[i]
    }
    for i := range configuredLink.DiscoveredLink {
        children[configuredLink.DiscoveredLink[i].GetSegmentPath()] = &configuredLink.DiscoveredLink[i]
    }
    return children
}

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-handle"] = configuredLink.InterfaceHandle
    leafs["ip-address"] = configuredLink.IpAddress
    leafs["ip-address-auto"] = configuredLink.IpAddressAuto
    leafs["vrf-id-present"] = configuredLink.VrfIdPresent
    leafs["vrf-id"] = configuredLink.VrfId
    leafs["minimum-preferred-links"] = configuredLink.MinimumPreferredLinks
    leafs["number-active-links"] = configuredLink.NumberActiveLinks
    leafs["min-links-satisfied"] = configuredLink.MinLinksSatisfied
    leafs["minimum-required-links"] = configuredLink.MinimumRequiredLinks
    leafs["required-min-links-satisfied"] = configuredLink.RequiredMinLinksSatisfied
    leafs["conflict-reason"] = configuredLink.ConflictReason
    leafs["conflict-context"] = configuredLink.ConflictContext
    return leafs
}

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetBundleName() string { return "cisco_ios_xr" }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetYangName() string { return "configured-link" }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) SetParent(parent types.Entity) { configuredLink.parent = parent }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetParent() types.Entity { return configuredLink.parent }

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetParentYangName() string { return "satellite-status" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange
// Port ranges table
type NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slot. The type is interface{} with range: 0..4294967295.
    Slot interface{}

    // Subslot. The type is interface{} with range: 0..4294967295.
    Subslot interface{}

    // Low port. The type is interface{} with range: 0..4294967295.
    LowPort interface{}

    // High port. The type is interface{} with range: 0..4294967295.
    HighPort interface{}

    // Port type. The type is IcpeOperPort.
    PortType interface{}

    // Conflict reason. The type is IcpeOperConflict.
    ConflictReason interface{}

    // Conflict context. The type is string.
    ConflictContext interface{}
}

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetFilter() yfilter.YFilter { return portRange.YFilter }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) SetFilter(yf yfilter.YFilter) { portRange.YFilter = yf }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    if yname == "subslot" { return "Subslot" }
    if yname == "low-port" { return "LowPort" }
    if yname == "high-port" { return "HighPort" }
    if yname == "port-type" { return "PortType" }
    if yname == "conflict-reason" { return "ConflictReason" }
    if yname == "conflict-context" { return "ConflictContext" }
    return ""
}

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetSegmentPath() string {
    return "port-range"
}

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot"] = portRange.Slot
    leafs["subslot"] = portRange.Subslot
    leafs["low-port"] = portRange.LowPort
    leafs["high-port"] = portRange.HighPort
    leafs["port-type"] = portRange.PortType
    leafs["conflict-reason"] = portRange.ConflictReason
    leafs["conflict-context"] = portRange.ConflictContext
    return leafs
}

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetBundleName() string { return "cisco_ios_xr" }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetYangName() string { return "port-range" }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) SetParent(parent types.Entity) { portRange.parent = parent }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetParent() types.Entity { return portRange.parent }

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetParentYangName() string { return "configured-link" }

// NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink
// Discovered Links table
type NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // State. The type is IcpeOperDiscdLinkState.
    State interface{}

    // Conflict reason. The type is IcpeOperConflict.
    ConflictReason interface{}

    // Conflict context. The type is string.
    ConflictContext interface{}
}

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetFilter() yfilter.YFilter { return discoveredLink.YFilter }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) SetFilter(yf yfilter.YFilter) { discoveredLink.YFilter = yf }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetGoName(yname string) string {
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "state" { return "State" }
    if yname == "conflict-reason" { return "ConflictReason" }
    if yname == "conflict-context" { return "ConflictContext" }
    return ""
}

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetSegmentPath() string {
    return "discovered-link"
}

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-handle"] = discoveredLink.InterfaceHandle
    leafs["state"] = discoveredLink.State
    leafs["conflict-reason"] = discoveredLink.ConflictReason
    leafs["conflict-context"] = discoveredLink.ConflictContext
    return leafs
}

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetBundleName() string { return "cisco_ios_xr" }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetYangName() string { return "discovered-link" }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) SetParent(parent types.Entity) { discoveredLink.parent = parent }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetParent() types.Entity { return discoveredLink.parent }

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetParentYangName() string { return "configured-link" }

// NvSatellite_SatellitePriorities
// Satellite priority information table
type NvSatellite_SatellitePriorities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite priority information. The type is slice of
    // NvSatellite_SatellitePriorities_SatellitePriority.
    SatellitePriority []NvSatellite_SatellitePriorities_SatellitePriority
}

func (satellitePriorities *NvSatellite_SatellitePriorities) GetFilter() yfilter.YFilter { return satellitePriorities.YFilter }

func (satellitePriorities *NvSatellite_SatellitePriorities) SetFilter(yf yfilter.YFilter) { satellitePriorities.YFilter = yf }

func (satellitePriorities *NvSatellite_SatellitePriorities) GetGoName(yname string) string {
    if yname == "satellite-priority" { return "SatellitePriority" }
    return ""
}

func (satellitePriorities *NvSatellite_SatellitePriorities) GetSegmentPath() string {
    return "satellite-priorities"
}

func (satellitePriorities *NvSatellite_SatellitePriorities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "satellite-priority" {
        for _, c := range satellitePriorities.SatellitePriority {
            if satellitePriorities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatellitePriorities_SatellitePriority{}
        satellitePriorities.SatellitePriority = append(satellitePriorities.SatellitePriority, child)
        return &satellitePriorities.SatellitePriority[len(satellitePriorities.SatellitePriority)-1]
    }
    return nil
}

func (satellitePriorities *NvSatellite_SatellitePriorities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range satellitePriorities.SatellitePriority {
        children[satellitePriorities.SatellitePriority[i].GetSegmentPath()] = &satellitePriorities.SatellitePriority[i]
    }
    return children
}

func (satellitePriorities *NvSatellite_SatellitePriorities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (satellitePriorities *NvSatellite_SatellitePriorities) GetBundleName() string { return "cisco_ios_xr" }

func (satellitePriorities *NvSatellite_SatellitePriorities) GetYangName() string { return "satellite-priorities" }

func (satellitePriorities *NvSatellite_SatellitePriorities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satellitePriorities *NvSatellite_SatellitePriorities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satellitePriorities *NvSatellite_SatellitePriorities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satellitePriorities *NvSatellite_SatellitePriorities) SetParent(parent types.Entity) { satellitePriorities.parent = parent }

func (satellitePriorities *NvSatellite_SatellitePriorities) GetParent() types.Entity { return satellitePriorities.parent }

func (satellitePriorities *NvSatellite_SatellitePriorities) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_SatellitePriorities_SatellitePriority
// Satellite priority information
type NvSatellite_SatellitePriorities_SatellitePriority struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite ID. The type is interface{} with range:
    // 100..65534.
    SatelliteId interface{}

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteIdXr interface{}

    // RG ID. The type is interface{} with range: 0..4294967295.
    Rgid interface{}

    // Best path hops. The type is interface{} with range: 0..4294967295.
    BestPathHops interface{}

    // Configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // Host priority. The type is interface{} with range: 0..18446744073709551615.
    HostPriority interface{}

    // Partner priority. The type is interface{} with range:
    // 0..18446744073709551615.
    PartnerPriority interface{}

    // Multichassis redundancy. The type is IcpeOperMultichassisRedundancy.
    MultichassisRedundancy interface{}
}

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetFilter() yfilter.YFilter { return satellitePriority.YFilter }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) SetFilter(yf yfilter.YFilter) { satellitePriority.YFilter = yf }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "satellite-id-xr" { return "SatelliteIdXr" }
    if yname == "rgid" { return "Rgid" }
    if yname == "best-path-hops" { return "BestPathHops" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "host-priority" { return "HostPriority" }
    if yname == "partner-priority" { return "PartnerPriority" }
    if yname == "multichassis-redundancy" { return "MultichassisRedundancy" }
    return ""
}

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetSegmentPath() string {
    return "satellite-priority" + "[satellite-id='" + fmt.Sprintf("%v", satellitePriority.SatelliteId) + "']"
}

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = satellitePriority.SatelliteId
    leafs["satellite-id-xr"] = satellitePriority.SatelliteIdXr
    leafs["rgid"] = satellitePriority.Rgid
    leafs["best-path-hops"] = satellitePriority.BestPathHops
    leafs["configured-priority"] = satellitePriority.ConfiguredPriority
    leafs["host-priority"] = satellitePriority.HostPriority
    leafs["partner-priority"] = satellitePriority.PartnerPriority
    leafs["multichassis-redundancy"] = satellitePriority.MultichassisRedundancy
    return leafs
}

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetBundleName() string { return "cisco_ios_xr" }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetYangName() string { return "satellite-priority" }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) SetParent(parent types.Entity) { satellitePriority.parent = parent }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetParent() types.Entity { return satellitePriority.parent }

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetParentYangName() string { return "satellite-priorities" }

// NvSatellite_SatelliteVersions
// Satellite remote version information table
type NvSatellite_SatelliteVersions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite remote version information. The type is slice of
    // NvSatellite_SatelliteVersions_SatelliteVersion.
    SatelliteVersion []NvSatellite_SatelliteVersions_SatelliteVersion
}

func (satelliteVersions *NvSatellite_SatelliteVersions) GetFilter() yfilter.YFilter { return satelliteVersions.YFilter }

func (satelliteVersions *NvSatellite_SatelliteVersions) SetFilter(yf yfilter.YFilter) { satelliteVersions.YFilter = yf }

func (satelliteVersions *NvSatellite_SatelliteVersions) GetGoName(yname string) string {
    if yname == "satellite-version" { return "SatelliteVersion" }
    return ""
}

func (satelliteVersions *NvSatellite_SatelliteVersions) GetSegmentPath() string {
    return "satellite-versions"
}

func (satelliteVersions *NvSatellite_SatelliteVersions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "satellite-version" {
        for _, c := range satelliteVersions.SatelliteVersion {
            if satelliteVersions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteVersions_SatelliteVersion{}
        satelliteVersions.SatelliteVersion = append(satelliteVersions.SatelliteVersion, child)
        return &satelliteVersions.SatelliteVersion[len(satelliteVersions.SatelliteVersion)-1]
    }
    return nil
}

func (satelliteVersions *NvSatellite_SatelliteVersions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range satelliteVersions.SatelliteVersion {
        children[satelliteVersions.SatelliteVersion[i].GetSegmentPath()] = &satelliteVersions.SatelliteVersion[i]
    }
    return children
}

func (satelliteVersions *NvSatellite_SatelliteVersions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (satelliteVersions *NvSatellite_SatelliteVersions) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteVersions *NvSatellite_SatelliteVersions) GetYangName() string { return "satellite-versions" }

func (satelliteVersions *NvSatellite_SatelliteVersions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteVersions *NvSatellite_SatelliteVersions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteVersions *NvSatellite_SatelliteVersions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteVersions *NvSatellite_SatelliteVersions) SetParent(parent types.Entity) { satelliteVersions.parent = parent }

func (satelliteVersions *NvSatellite_SatelliteVersions) GetParent() types.Entity { return satelliteVersions.parent }

func (satelliteVersions *NvSatellite_SatelliteVersions) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_SatelliteVersions_SatelliteVersion
// Satellite remote version information
type NvSatellite_SatelliteVersions_SatelliteVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite ID. The type is interface{} with range:
    // 100..65534.
    SatelliteId interface{}

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteIdXr interface{}

    // Version check state. The type is IcpeOperVerCheckState.
    VersionCheckState interface{}

    // Remote version present. The type is bool.
    RemoteVersionPresent interface{}

    // Remote version. The type is slice of string.
    RemoteVersion []interface{}

    // Satellite active version information.
    ActiveVersion NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion

    // Satellite transferred version information.
    TransferredVersion NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion

    // Satellite committed version information.
    CommittedVersion NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion
}

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetFilter() yfilter.YFilter { return satelliteVersion.YFilter }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) SetFilter(yf yfilter.YFilter) { satelliteVersion.YFilter = yf }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "satellite-id-xr" { return "SatelliteIdXr" }
    if yname == "version-check-state" { return "VersionCheckState" }
    if yname == "remote-version-present" { return "RemoteVersionPresent" }
    if yname == "remote-version" { return "RemoteVersion" }
    if yname == "active-version" { return "ActiveVersion" }
    if yname == "transferred-version" { return "TransferredVersion" }
    if yname == "committed-version" { return "CommittedVersion" }
    return ""
}

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetSegmentPath() string {
    return "satellite-version" + "[satellite-id='" + fmt.Sprintf("%v", satelliteVersion.SatelliteId) + "']"
}

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active-version" {
        return &satelliteVersion.ActiveVersion
    }
    if childYangName == "transferred-version" {
        return &satelliteVersion.TransferredVersion
    }
    if childYangName == "committed-version" {
        return &satelliteVersion.CommittedVersion
    }
    return nil
}

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["active-version"] = &satelliteVersion.ActiveVersion
    children["transferred-version"] = &satelliteVersion.TransferredVersion
    children["committed-version"] = &satelliteVersion.CommittedVersion
    return children
}

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = satelliteVersion.SatelliteId
    leafs["satellite-id-xr"] = satelliteVersion.SatelliteIdXr
    leafs["version-check-state"] = satelliteVersion.VersionCheckState
    leafs["remote-version-present"] = satelliteVersion.RemoteVersionPresent
    leafs["remote-version"] = satelliteVersion.RemoteVersion
    return leafs
}

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetYangName() string { return "satellite-version" }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) SetParent(parent types.Entity) { satelliteVersion.parent = parent }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetParent() types.Entity { return satelliteVersion.parent }

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetParentYangName() string { return "satellite-versions" }

// NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion
// Satellite active version information
type NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version check state. The type is IcpeOperVerCheckState.
    VersionCheckState interface{}

    // Remote version present. The type is bool.
    RemoteVersionPresent interface{}

    // Remote version. The type is slice of string.
    RemoteVersion []interface{}
}

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetFilter() yfilter.YFilter { return activeVersion.YFilter }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) SetFilter(yf yfilter.YFilter) { activeVersion.YFilter = yf }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetGoName(yname string) string {
    if yname == "version-check-state" { return "VersionCheckState" }
    if yname == "remote-version-present" { return "RemoteVersionPresent" }
    if yname == "remote-version" { return "RemoteVersion" }
    return ""
}

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetSegmentPath() string {
    return "active-version"
}

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version-check-state"] = activeVersion.VersionCheckState
    leafs["remote-version-present"] = activeVersion.RemoteVersionPresent
    leafs["remote-version"] = activeVersion.RemoteVersion
    return leafs
}

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetBundleName() string { return "cisco_ios_xr" }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetYangName() string { return "active-version" }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) SetParent(parent types.Entity) { activeVersion.parent = parent }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetParent() types.Entity { return activeVersion.parent }

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetParentYangName() string { return "satellite-version" }

// NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion
// Satellite transferred version information
type NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version check state. The type is IcpeOperVerCheckState.
    VersionCheckState interface{}

    // Remote version present. The type is bool.
    RemoteVersionPresent interface{}

    // Remote version. The type is slice of string.
    RemoteVersion []interface{}
}

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetFilter() yfilter.YFilter { return transferredVersion.YFilter }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) SetFilter(yf yfilter.YFilter) { transferredVersion.YFilter = yf }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetGoName(yname string) string {
    if yname == "version-check-state" { return "VersionCheckState" }
    if yname == "remote-version-present" { return "RemoteVersionPresent" }
    if yname == "remote-version" { return "RemoteVersion" }
    return ""
}

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetSegmentPath() string {
    return "transferred-version"
}

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version-check-state"] = transferredVersion.VersionCheckState
    leafs["remote-version-present"] = transferredVersion.RemoteVersionPresent
    leafs["remote-version"] = transferredVersion.RemoteVersion
    return leafs
}

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetBundleName() string { return "cisco_ios_xr" }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetYangName() string { return "transferred-version" }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) SetParent(parent types.Entity) { transferredVersion.parent = parent }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetParent() types.Entity { return transferredVersion.parent }

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetParentYangName() string { return "satellite-version" }

// NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion
// Satellite committed version information
type NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version check state. The type is IcpeOperVerCheckState.
    VersionCheckState interface{}

    // Remote version present. The type is bool.
    RemoteVersionPresent interface{}

    // Remote version. The type is slice of string.
    RemoteVersion []interface{}
}

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetFilter() yfilter.YFilter { return committedVersion.YFilter }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) SetFilter(yf yfilter.YFilter) { committedVersion.YFilter = yf }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetGoName(yname string) string {
    if yname == "version-check-state" { return "VersionCheckState" }
    if yname == "remote-version-present" { return "RemoteVersionPresent" }
    if yname == "remote-version" { return "RemoteVersion" }
    return ""
}

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetSegmentPath() string {
    return "committed-version"
}

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version-check-state"] = committedVersion.VersionCheckState
    leafs["remote-version-present"] = committedVersion.RemoteVersionPresent
    leafs["remote-version"] = committedVersion.RemoteVersion
    return leafs
}

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetBundleName() string { return "cisco_ios_xr" }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetYangName() string { return "committed-version" }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) SetParent(parent types.Entity) { committedVersion.parent = parent }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetParent() types.Entity { return committedVersion.parent }

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetParentYangName() string { return "satellite-version" }

// NvSatellite_SatelliteTopologies
// Satellite Topology Information table
type NvSatellite_SatelliteTopologies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite Topology Information. The type is slice of
    // NvSatellite_SatelliteTopologies_SatelliteTopology.
    SatelliteTopology []NvSatellite_SatelliteTopologies_SatelliteTopology
}

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetFilter() yfilter.YFilter { return satelliteTopologies.YFilter }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) SetFilter(yf yfilter.YFilter) { satelliteTopologies.YFilter = yf }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetGoName(yname string) string {
    if yname == "satellite-topology" { return "SatelliteTopology" }
    return ""
}

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetSegmentPath() string {
    return "satellite-topologies"
}

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "satellite-topology" {
        for _, c := range satelliteTopologies.SatelliteTopology {
            if satelliteTopologies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteTopologies_SatelliteTopology{}
        satelliteTopologies.SatelliteTopology = append(satelliteTopologies.SatelliteTopology, child)
        return &satelliteTopologies.SatelliteTopology[len(satelliteTopologies.SatelliteTopology)-1]
    }
    return nil
}

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range satelliteTopologies.SatelliteTopology {
        children[satelliteTopologies.SatelliteTopology[i].GetSegmentPath()] = &satelliteTopologies.SatelliteTopology[i]
    }
    return children
}

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetYangName() string { return "satellite-topologies" }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) SetParent(parent types.Entity) { satelliteTopologies.parent = parent }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetParent() types.Entity { return satelliteTopologies.parent }

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_SatelliteTopologies_SatelliteTopology
// Satellite Topology Information
type NvSatellite_SatelliteTopologies_SatelliteTopology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // Redundancy ICCP group. The type is interface{} with range: 0..4294967295.
    RedundancyIccpGroup interface{}

    // Is physical. The type is bool.
    IsPhysical interface{}

    // Ring whole. The type is bool.
    RingWhole interface{}

    // Discovered Links table. The type is slice of
    // NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink.
    DiscoveredLink []NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink

    // Satellite table. The type is slice of
    // NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite.
    Satellite []NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite
}

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetFilter() yfilter.YFilter { return satelliteTopology.YFilter }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) SetFilter(yf yfilter.YFilter) { satelliteTopology.YFilter = yf }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "redundancy-iccp-group" { return "RedundancyIccpGroup" }
    if yname == "is-physical" { return "IsPhysical" }
    if yname == "ring-whole" { return "RingWhole" }
    if yname == "discovered-link" { return "DiscoveredLink" }
    if yname == "satellite" { return "Satellite" }
    return ""
}

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetSegmentPath() string {
    return "satellite-topology" + "[interface-name='" + fmt.Sprintf("%v", satelliteTopology.InterfaceName) + "']"
}

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discovered-link" {
        for _, c := range satelliteTopology.DiscoveredLink {
            if satelliteTopology.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink{}
        satelliteTopology.DiscoveredLink = append(satelliteTopology.DiscoveredLink, child)
        return &satelliteTopology.DiscoveredLink[len(satelliteTopology.DiscoveredLink)-1]
    }
    if childYangName == "satellite" {
        for _, c := range satelliteTopology.Satellite {
            if satelliteTopology.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite{}
        satelliteTopology.Satellite = append(satelliteTopology.Satellite, child)
        return &satelliteTopology.Satellite[len(satelliteTopology.Satellite)-1]
    }
    return nil
}

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range satelliteTopology.DiscoveredLink {
        children[satelliteTopology.DiscoveredLink[i].GetSegmentPath()] = &satelliteTopology.DiscoveredLink[i]
    }
    for i := range satelliteTopology.Satellite {
        children[satelliteTopology.Satellite[i].GetSegmentPath()] = &satelliteTopology.Satellite[i]
    }
    return children
}

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = satelliteTopology.InterfaceName
    leafs["interface-name-xr"] = satelliteTopology.InterfaceNameXr
    leafs["interface-handle"] = satelliteTopology.InterfaceHandle
    leafs["redundancy-iccp-group"] = satelliteTopology.RedundancyIccpGroup
    leafs["is-physical"] = satelliteTopology.IsPhysical
    leafs["ring-whole"] = satelliteTopology.RingWhole
    return leafs
}

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetYangName() string { return "satellite-topology" }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) SetParent(parent types.Entity) { satelliteTopology.parent = parent }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetParent() types.Entity { return satelliteTopology.parent }

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetParentYangName() string { return "satellite-topologies" }

// NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink
// Discovered Links table
type NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // Discovery running. The type is bool.
    DiscoveryRunning interface{}
}

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetFilter() yfilter.YFilter { return discoveredLink.YFilter }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) SetFilter(yf yfilter.YFilter) { discoveredLink.YFilter = yf }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "discovery-running" { return "DiscoveryRunning" }
    return ""
}

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetSegmentPath() string {
    return "discovered-link"
}

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = discoveredLink.InterfaceName
    leafs["interface-handle"] = discoveredLink.InterfaceHandle
    leafs["discovery-running"] = discoveredLink.DiscoveryRunning
    return leafs
}

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetBundleName() string { return "cisco_ios_xr" }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetYangName() string { return "discovered-link" }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) SetParent(parent types.Entity) { discoveredLink.parent = parent }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetParent() types.Entity { return discoveredLink.parent }

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetParentYangName() string { return "satellite-topology" }

// NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite
// Satellite table
type NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Configured. The type is bool.
    Configured interface{}

    // Num hops. The type is interface{} with range: 0..65535.
    NumHops interface{}

    // Type. The type is string.
    Type interface{}

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteId interface{}

    // Received serial number. The type is string.
    ReceivedSerialNumber interface{}

    // Received serial number present. The type is bool.
    ReceivedSerialNumberPresent interface{}

    // VLAN ID. The type is interface{} with range: 0..65535.
    VlanId interface{}

    // Display name. The type is string.
    DisplayName interface{}

    // Conflict reason. The type is IcpeOperConflict.
    ConflictReason interface{}

    // Conflict context. The type is string.
    ConflictContext interface{}

    // Local Satellite Fabric Link table. The type is slice of
    // NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink.
    FabricLink []NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink
}

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetFilter() yfilter.YFilter { return satellite.YFilter }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) SetFilter(yf yfilter.YFilter) { satellite.YFilter = yf }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "configured" { return "Configured" }
    if yname == "num-hops" { return "NumHops" }
    if yname == "type" { return "Type" }
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "received-serial-number" { return "ReceivedSerialNumber" }
    if yname == "received-serial-number-present" { return "ReceivedSerialNumberPresent" }
    if yname == "vlan-id" { return "VlanId" }
    if yname == "display-name" { return "DisplayName" }
    if yname == "conflict-reason" { return "ConflictReason" }
    if yname == "conflict-context" { return "ConflictContext" }
    if yname == "fabric-link" { return "FabricLink" }
    return ""
}

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetSegmentPath() string {
    return "satellite"
}

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fabric-link" {
        for _, c := range satellite.FabricLink {
            if satellite.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink{}
        satellite.FabricLink = append(satellite.FabricLink, child)
        return &satellite.FabricLink[len(satellite.FabricLink)-1]
    }
    return nil
}

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range satellite.FabricLink {
        children[satellite.FabricLink[i].GetSegmentPath()] = &satellite.FabricLink[i]
    }
    return children
}

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = satellite.MacAddress
    leafs["configured"] = satellite.Configured
    leafs["num-hops"] = satellite.NumHops
    leafs["type"] = satellite.Type
    leafs["satellite-id"] = satellite.SatelliteId
    leafs["received-serial-number"] = satellite.ReceivedSerialNumber
    leafs["received-serial-number-present"] = satellite.ReceivedSerialNumberPresent
    leafs["vlan-id"] = satellite.VlanId
    leafs["display-name"] = satellite.DisplayName
    leafs["conflict-reason"] = satellite.ConflictReason
    leafs["conflict-context"] = satellite.ConflictContext
    return leafs
}

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetBundleName() string { return "cisco_ios_xr" }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetYangName() string { return "satellite" }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) SetParent(parent types.Entity) { satellite.parent = parent }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetParent() types.Entity { return satellite.parent }

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetParentYangName() string { return "satellite-topology" }

// NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink
// Local Satellite Fabric Link table
type NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICL ID. The type is interface{} with range: 0..4294967295.
    IclId interface{}

    // Interface name. The type is string.
    InterfaceName interface{}

    // Display name. The type is string.
    DisplayName interface{}

    // Redundant. The type is bool.
    Redundant interface{}

    // Active. The type is bool.
    Active interface{}

    // Obsolete. The type is bool.
    Obsolete interface{}

    // Remote Device table. The type is slice of
    // NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice.
    RemoteDevice []NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice
}

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetFilter() yfilter.YFilter { return fabricLink.YFilter }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) SetFilter(yf yfilter.YFilter) { fabricLink.YFilter = yf }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetGoName(yname string) string {
    if yname == "icl-id" { return "IclId" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "display-name" { return "DisplayName" }
    if yname == "redundant" { return "Redundant" }
    if yname == "active" { return "Active" }
    if yname == "obsolete" { return "Obsolete" }
    if yname == "remote-device" { return "RemoteDevice" }
    return ""
}

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetSegmentPath() string {
    return "fabric-link"
}

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-device" {
        for _, c := range fabricLink.RemoteDevice {
            if fabricLink.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice{}
        fabricLink.RemoteDevice = append(fabricLink.RemoteDevice, child)
        return &fabricLink.RemoteDevice[len(fabricLink.RemoteDevice)-1]
    }
    return nil
}

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fabricLink.RemoteDevice {
        children[fabricLink.RemoteDevice[i].GetSegmentPath()] = &fabricLink.RemoteDevice[i]
    }
    return children
}

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icl-id"] = fabricLink.IclId
    leafs["interface-name"] = fabricLink.InterfaceName
    leafs["display-name"] = fabricLink.DisplayName
    leafs["redundant"] = fabricLink.Redundant
    leafs["active"] = fabricLink.Active
    leafs["obsolete"] = fabricLink.Obsolete
    return leafs
}

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetBundleName() string { return "cisco_ios_xr" }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetYangName() string { return "fabric-link" }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) SetParent(parent types.Entity) { fabricLink.parent = parent }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetParent() types.Entity { return fabricLink.parent }

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetParentYangName() string { return "satellite" }

// NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice
// Remote Device table
type NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Source. The type is IcpeOperTopoRemoteSource.
    Source interface{}

    // Remote is satellite. The type is bool.
    RemoteIsSatellite interface{}

    // Remote is local host. The type is bool.
    RemoteIsLocalHost interface{}

    // ICL ID. The type is interface{} with range: 0..4294967295.
    IclId interface{}

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // Interface name. The type is string.
    InterfaceName interface{}
}

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetFilter() yfilter.YFilter { return remoteDevice.YFilter }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) SetFilter(yf yfilter.YFilter) { remoteDevice.YFilter = yf }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "source" { return "Source" }
    if yname == "remote-is-satellite" { return "RemoteIsSatellite" }
    if yname == "remote-is-local-host" { return "RemoteIsLocalHost" }
    if yname == "icl-id" { return "IclId" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetSegmentPath() string {
    return "remote-device"
}

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = remoteDevice.MacAddress
    leafs["source"] = remoteDevice.Source
    leafs["remote-is-satellite"] = remoteDevice.RemoteIsSatellite
    leafs["remote-is-local-host"] = remoteDevice.RemoteIsLocalHost
    leafs["icl-id"] = remoteDevice.IclId
    leafs["interface-handle"] = remoteDevice.InterfaceHandle
    leafs["interface-name"] = remoteDevice.InterfaceName
    return leafs
}

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetBundleName() string { return "cisco_ios_xr" }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetYangName() string { return "remote-device" }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) SetParent(parent types.Entity) { remoteDevice.parent = parent }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetParent() types.Entity { return remoteDevice.parent }

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetParentYangName() string { return "fabric-link" }

// NvSatellite_InstallReferenceInfo
// Satellite Install Reference Information
type NvSatellite_InstallReferenceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install Reference Information table.
    References NvSatellite_InstallReferenceInfo_References
}

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetFilter() yfilter.YFilter { return installReferenceInfo.YFilter }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) SetFilter(yf yfilter.YFilter) { installReferenceInfo.YFilter = yf }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetGoName(yname string) string {
    if yname == "references" { return "References" }
    return ""
}

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetSegmentPath() string {
    return "install-reference-info"
}

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "references" {
        return &installReferenceInfo.References
    }
    return nil
}

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["references"] = &installReferenceInfo.References
    return children
}

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetYangName() string { return "install-reference-info" }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) SetParent(parent types.Entity) { installReferenceInfo.parent = parent }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetParent() types.Entity { return installReferenceInfo.parent }

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_InstallReferenceInfo_References
// Install Reference Information table
type NvSatellite_InstallReferenceInfo_References struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install Reference Information. The type is slice of
    // NvSatellite_InstallReferenceInfo_References_Reference.
    Reference []NvSatellite_InstallReferenceInfo_References_Reference
}

func (references *NvSatellite_InstallReferenceInfo_References) GetFilter() yfilter.YFilter { return references.YFilter }

func (references *NvSatellite_InstallReferenceInfo_References) SetFilter(yf yfilter.YFilter) { references.YFilter = yf }

func (references *NvSatellite_InstallReferenceInfo_References) GetGoName(yname string) string {
    if yname == "reference" { return "Reference" }
    return ""
}

func (references *NvSatellite_InstallReferenceInfo_References) GetSegmentPath() string {
    return "references"
}

func (references *NvSatellite_InstallReferenceInfo_References) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reference" {
        for _, c := range references.Reference {
            if references.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_InstallReferenceInfo_References_Reference{}
        references.Reference = append(references.Reference, child)
        return &references.Reference[len(references.Reference)-1]
    }
    return nil
}

func (references *NvSatellite_InstallReferenceInfo_References) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range references.Reference {
        children[references.Reference[i].GetSegmentPath()] = &references.Reference[i]
    }
    return children
}

func (references *NvSatellite_InstallReferenceInfo_References) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (references *NvSatellite_InstallReferenceInfo_References) GetBundleName() string { return "cisco_ios_xr" }

func (references *NvSatellite_InstallReferenceInfo_References) GetYangName() string { return "references" }

func (references *NvSatellite_InstallReferenceInfo_References) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (references *NvSatellite_InstallReferenceInfo_References) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (references *NvSatellite_InstallReferenceInfo_References) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (references *NvSatellite_InstallReferenceInfo_References) SetParent(parent types.Entity) { references.parent = parent }

func (references *NvSatellite_InstallReferenceInfo_References) GetParent() types.Entity { return references.parent }

func (references *NvSatellite_InstallReferenceInfo_References) GetParentYangName() string { return "install-reference-info" }

// NvSatellite_InstallReferenceInfo_References_Reference
// Install Reference Information
type NvSatellite_InstallReferenceInfo_References_Reference struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ReferenceName interface{}

    // Reference name. The type is string.
    ReferenceNameXr interface{}

    // Reference files. The type is slice of string.
    ReferenceFile []interface{}
}

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetFilter() yfilter.YFilter { return reference.YFilter }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) SetFilter(yf yfilter.YFilter) { reference.YFilter = yf }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetGoName(yname string) string {
    if yname == "reference-name" { return "ReferenceName" }
    if yname == "reference-name-xr" { return "ReferenceNameXr" }
    if yname == "reference-file" { return "ReferenceFile" }
    return ""
}

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetSegmentPath() string {
    return "reference" + "[reference-name='" + fmt.Sprintf("%v", reference.ReferenceName) + "']"
}

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reference-name"] = reference.ReferenceName
    leafs["reference-name-xr"] = reference.ReferenceNameXr
    leafs["reference-file"] = reference.ReferenceFile
    return leafs
}

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetBundleName() string { return "cisco_ios_xr" }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetYangName() string { return "reference" }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) SetParent(parent types.Entity) { reference.parent = parent }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetParent() types.Entity { return reference.parent }

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetParentYangName() string { return "references" }

// NvSatellite_InstallOpProgresses
// Current percentage of install table
type NvSatellite_InstallOpProgresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current percentage of install. The type is slice of
    // NvSatellite_InstallOpProgresses_InstallOpProgress.
    InstallOpProgress []NvSatellite_InstallOpProgresses_InstallOpProgress
}

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetFilter() yfilter.YFilter { return installOpProgresses.YFilter }

func (installOpProgresses *NvSatellite_InstallOpProgresses) SetFilter(yf yfilter.YFilter) { installOpProgresses.YFilter = yf }

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetGoName(yname string) string {
    if yname == "install-op-progress" { return "InstallOpProgress" }
    return ""
}

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetSegmentPath() string {
    return "install-op-progresses"
}

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "install-op-progress" {
        for _, c := range installOpProgresses.InstallOpProgress {
            if installOpProgresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_InstallOpProgresses_InstallOpProgress{}
        installOpProgresses.InstallOpProgress = append(installOpProgresses.InstallOpProgress, child)
        return &installOpProgresses.InstallOpProgress[len(installOpProgresses.InstallOpProgress)-1]
    }
    return nil
}

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range installOpProgresses.InstallOpProgress {
        children[installOpProgresses.InstallOpProgress[i].GetSegmentPath()] = &installOpProgresses.InstallOpProgress[i]
    }
    return children
}

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetBundleName() string { return "cisco_ios_xr" }

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetYangName() string { return "install-op-progresses" }

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (installOpProgresses *NvSatellite_InstallOpProgresses) SetParent(parent types.Entity) { installOpProgresses.parent = parent }

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetParent() types.Entity { return installOpProgresses.parent }

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_InstallOpProgresses_InstallOpProgress
// Current percentage of install
type NvSatellite_InstallOpProgresses_InstallOpProgress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Operation ID. The type is interface{} with range:
    // 0..4294967295.
    OperationId interface{}

    // Operation ID. The type is interface{} with range: 0..4294967295.
    OperationIdXr interface{}

    // Progress percentage. The type is interface{} with range: 0..65535. Units
    // are percentage.
    ProgressPercentage interface{}

    // Satellite count. The type is interface{} with range: 0..4294967295.
    SatelliteCount interface{}
}

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetFilter() yfilter.YFilter { return installOpProgress.YFilter }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) SetFilter(yf yfilter.YFilter) { installOpProgress.YFilter = yf }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetGoName(yname string) string {
    if yname == "operation-id" { return "OperationId" }
    if yname == "operation-id-xr" { return "OperationIdXr" }
    if yname == "progress-percentage" { return "ProgressPercentage" }
    if yname == "satellite-count" { return "SatelliteCount" }
    return ""
}

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetSegmentPath() string {
    return "install-op-progress" + "[operation-id='" + fmt.Sprintf("%v", installOpProgress.OperationId) + "']"
}

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-id"] = installOpProgress.OperationId
    leafs["operation-id-xr"] = installOpProgress.OperationIdXr
    leafs["progress-percentage"] = installOpProgress.ProgressPercentage
    leafs["satellite-count"] = installOpProgress.SatelliteCount
    return leafs
}

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetBundleName() string { return "cisco_ios_xr" }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetYangName() string { return "install-op-progress" }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) SetParent(parent types.Entity) { installOpProgress.parent = parent }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetParent() types.Entity { return installOpProgress.parent }

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetParentYangName() string { return "install-op-progresses" }

// NvSatellite_ReloadStatuses
// Detailed breakdown of reload status table
type NvSatellite_ReloadStatuses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed breakdown of reload status. The type is slice of
    // NvSatellite_ReloadStatuses_ReloadStatus.
    ReloadStatus []NvSatellite_ReloadStatuses_ReloadStatus
}

func (reloadStatuses *NvSatellite_ReloadStatuses) GetFilter() yfilter.YFilter { return reloadStatuses.YFilter }

func (reloadStatuses *NvSatellite_ReloadStatuses) SetFilter(yf yfilter.YFilter) { reloadStatuses.YFilter = yf }

func (reloadStatuses *NvSatellite_ReloadStatuses) GetGoName(yname string) string {
    if yname == "reload-status" { return "ReloadStatus" }
    return ""
}

func (reloadStatuses *NvSatellite_ReloadStatuses) GetSegmentPath() string {
    return "reload-statuses"
}

func (reloadStatuses *NvSatellite_ReloadStatuses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reload-status" {
        for _, c := range reloadStatuses.ReloadStatus {
            if reloadStatuses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_ReloadStatuses_ReloadStatus{}
        reloadStatuses.ReloadStatus = append(reloadStatuses.ReloadStatus, child)
        return &reloadStatuses.ReloadStatus[len(reloadStatuses.ReloadStatus)-1]
    }
    return nil
}

func (reloadStatuses *NvSatellite_ReloadStatuses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range reloadStatuses.ReloadStatus {
        children[reloadStatuses.ReloadStatus[i].GetSegmentPath()] = &reloadStatuses.ReloadStatus[i]
    }
    return children
}

func (reloadStatuses *NvSatellite_ReloadStatuses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (reloadStatuses *NvSatellite_ReloadStatuses) GetBundleName() string { return "cisco_ios_xr" }

func (reloadStatuses *NvSatellite_ReloadStatuses) GetYangName() string { return "reload-statuses" }

func (reloadStatuses *NvSatellite_ReloadStatuses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reloadStatuses *NvSatellite_ReloadStatuses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reloadStatuses *NvSatellite_ReloadStatuses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reloadStatuses *NvSatellite_ReloadStatuses) SetParent(parent types.Entity) { reloadStatuses.parent = parent }

func (reloadStatuses *NvSatellite_ReloadStatuses) GetParent() types.Entity { return reloadStatuses.parent }

func (reloadStatuses *NvSatellite_ReloadStatuses) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_ReloadStatuses_ReloadStatus
// Detailed breakdown of reload status
type NvSatellite_ReloadStatuses_ReloadStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite range. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SatelliteRange interface{}

    // Satellite range. The type is string.
    SatelliteRangeXr interface{}

    // Sats not initiated. The type is slice of interface{} with range: 0..65535.
    SatsNotInitiated []interface{}

    // Sats reloading. The type is slice of interface{} with range: 0..65535.
    SatsReloading []interface{}

    // Sats reloaded. The type is slice of interface{} with range: 0..65535.
    SatsReloaded []interface{}

    // Sats reload failed. The type is slice of interface{} with range: 0..65535.
    SatsReloadFailed []interface{}
}

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetFilter() yfilter.YFilter { return reloadStatus.YFilter }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) SetFilter(yf yfilter.YFilter) { reloadStatus.YFilter = yf }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetGoName(yname string) string {
    if yname == "satellite-range" { return "SatelliteRange" }
    if yname == "satellite-range-xr" { return "SatelliteRangeXr" }
    if yname == "sats-not-initiated" { return "SatsNotInitiated" }
    if yname == "sats-reloading" { return "SatsReloading" }
    if yname == "sats-reloaded" { return "SatsReloaded" }
    if yname == "sats-reload-failed" { return "SatsReloadFailed" }
    return ""
}

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetSegmentPath() string {
    return "reload-status" + "[satellite-range='" + fmt.Sprintf("%v", reloadStatus.SatelliteRange) + "']"
}

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-range"] = reloadStatus.SatelliteRange
    leafs["satellite-range-xr"] = reloadStatus.SatelliteRangeXr
    leafs["sats-not-initiated"] = reloadStatus.SatsNotInitiated
    leafs["sats-reloading"] = reloadStatus.SatsReloading
    leafs["sats-reloaded"] = reloadStatus.SatsReloaded
    leafs["sats-reload-failed"] = reloadStatus.SatsReloadFailed
    return leafs
}

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetBundleName() string { return "cisco_ios_xr" }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetYangName() string { return "reload-status" }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) SetParent(parent types.Entity) { reloadStatus.parent = parent }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetParent() types.Entity { return reloadStatus.parent }

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetParentYangName() string { return "reload-statuses" }

// NvSatellite_Install
// Satellite Install Information
type NvSatellite_Install struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite Software Package Information table.
    SatelliteSoftwareVersions NvSatellite_Install_SatelliteSoftwareVersions
}

func (install *NvSatellite_Install) GetFilter() yfilter.YFilter { return install.YFilter }

func (install *NvSatellite_Install) SetFilter(yf yfilter.YFilter) { install.YFilter = yf }

func (install *NvSatellite_Install) GetGoName(yname string) string {
    if yname == "satellite-software-versions" { return "SatelliteSoftwareVersions" }
    return ""
}

func (install *NvSatellite_Install) GetSegmentPath() string {
    return "install"
}

func (install *NvSatellite_Install) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "satellite-software-versions" {
        return &install.SatelliteSoftwareVersions
    }
    return nil
}

func (install *NvSatellite_Install) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["satellite-software-versions"] = &install.SatelliteSoftwareVersions
    return children
}

func (install *NvSatellite_Install) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (install *NvSatellite_Install) GetBundleName() string { return "cisco_ios_xr" }

func (install *NvSatellite_Install) GetYangName() string { return "install" }

func (install *NvSatellite_Install) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (install *NvSatellite_Install) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (install *NvSatellite_Install) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (install *NvSatellite_Install) SetParent(parent types.Entity) { install.parent = parent }

func (install *NvSatellite_Install) GetParent() types.Entity { return install.parent }

func (install *NvSatellite_Install) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_Install_SatelliteSoftwareVersions
// Satellite Software Package Information table
type NvSatellite_Install_SatelliteSoftwareVersions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite Software Package Information. The type is slice of
    // NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion.
    SatelliteSoftwareVersion []NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion
}

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetFilter() yfilter.YFilter { return satelliteSoftwareVersions.YFilter }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) SetFilter(yf yfilter.YFilter) { satelliteSoftwareVersions.YFilter = yf }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetGoName(yname string) string {
    if yname == "satellite-software-version" { return "SatelliteSoftwareVersion" }
    return ""
}

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetSegmentPath() string {
    return "satellite-software-versions"
}

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "satellite-software-version" {
        for _, c := range satelliteSoftwareVersions.SatelliteSoftwareVersion {
            if satelliteSoftwareVersions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion{}
        satelliteSoftwareVersions.SatelliteSoftwareVersion = append(satelliteSoftwareVersions.SatelliteSoftwareVersion, child)
        return &satelliteSoftwareVersions.SatelliteSoftwareVersion[len(satelliteSoftwareVersions.SatelliteSoftwareVersion)-1]
    }
    return nil
}

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range satelliteSoftwareVersions.SatelliteSoftwareVersion {
        children[satelliteSoftwareVersions.SatelliteSoftwareVersion[i].GetSegmentPath()] = &satelliteSoftwareVersions.SatelliteSoftwareVersion[i]
    }
    return children
}

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetYangName() string { return "satellite-software-versions" }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) SetParent(parent types.Entity) { satelliteSoftwareVersions.parent = parent }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetParent() types.Entity { return satelliteSoftwareVersions.parent }

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetParentYangName() string { return "install" }

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion
// Satellite Software Package Information
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite ID. The type is interface{} with range:
    // 100..65534.
    SatelliteId interface{}

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteIdXr interface{}

    // Package support. The type is IcpeInstallPkgSupp.
    PackageSupport interface{}

    // Package Data on this Satellite.
    InstallPackageInfo NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo
}

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetFilter() yfilter.YFilter { return satelliteSoftwareVersion.YFilter }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) SetFilter(yf yfilter.YFilter) { satelliteSoftwareVersion.YFilter = yf }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "satellite-id-xr" { return "SatelliteIdXr" }
    if yname == "package-support" { return "PackageSupport" }
    if yname == "install-package-info" { return "InstallPackageInfo" }
    return ""
}

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetSegmentPath() string {
    return "satellite-software-version" + "[satellite-id='" + fmt.Sprintf("%v", satelliteSoftwareVersion.SatelliteId) + "']"
}

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "install-package-info" {
        return &satelliteSoftwareVersion.InstallPackageInfo
    }
    return nil
}

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["install-package-info"] = &satelliteSoftwareVersion.InstallPackageInfo
    return children
}

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = satelliteSoftwareVersion.SatelliteId
    leafs["satellite-id-xr"] = satelliteSoftwareVersion.SatelliteIdXr
    leafs["package-support"] = satelliteSoftwareVersion.PackageSupport
    return leafs
}

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetYangName() string { return "satellite-software-version" }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) SetParent(parent types.Entity) { satelliteSoftwareVersion.parent = parent }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetParent() types.Entity { return satelliteSoftwareVersion.parent }

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetParentYangName() string { return "satellite-software-versions" }

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo
// Package Data on this Satellite
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Active Packages running on this Satellite.
    ActivePackages NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages

    // Inactive Packages on this Satellite.
    InactivePackages NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages

    // Committed Packages running on this Satellite.
    CommittedPackages NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages
}

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetFilter() yfilter.YFilter { return installPackageInfo.YFilter }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) SetFilter(yf yfilter.YFilter) { installPackageInfo.YFilter = yf }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetGoName(yname string) string {
    if yname == "active-packages" { return "ActivePackages" }
    if yname == "inactive-packages" { return "InactivePackages" }
    if yname == "committed-packages" { return "CommittedPackages" }
    return ""
}

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetSegmentPath() string {
    return "install-package-info"
}

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active-packages" {
        return &installPackageInfo.ActivePackages
    }
    if childYangName == "inactive-packages" {
        return &installPackageInfo.InactivePackages
    }
    if childYangName == "committed-packages" {
        return &installPackageInfo.CommittedPackages
    }
    return nil
}

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["active-packages"] = &installPackageInfo.ActivePackages
    children["inactive-packages"] = &installPackageInfo.InactivePackages
    children["committed-packages"] = &installPackageInfo.CommittedPackages
    return children
}

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetBundleName() string { return "cisco_ios_xr" }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetYangName() string { return "install-package-info" }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) SetParent(parent types.Entity) { installPackageInfo.parent = parent }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetParent() types.Entity { return installPackageInfo.parent }

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetParentYangName() string { return "satellite-software-version" }

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages
// Active Packages running on this Satellite
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A package on this Satellite table. The type is slice of
    // NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package.
    Package []NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package
}

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetFilter() yfilter.YFilter { return activePackages.YFilter }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) SetFilter(yf yfilter.YFilter) { activePackages.YFilter = yf }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetGoName(yname string) string {
    if yname == "package" { return "Package" }
    return ""
}

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetSegmentPath() string {
    return "active-packages"
}

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        for _, c := range activePackages.Package {
            if activePackages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package{}
        activePackages.Package = append(activePackages.Package, child)
        return &activePackages.Package[len(activePackages.Package)-1]
    }
    return nil
}

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range activePackages.Package {
        children[activePackages.Package[i].GetSegmentPath()] = &activePackages.Package[i]
    }
    return children
}

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetBundleName() string { return "cisco_ios_xr" }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetYangName() string { return "active-packages" }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) SetParent(parent types.Entity) { activePackages.parent = parent }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetParent() types.Entity { return activePackages.parent }

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetParentYangName() string { return "install-package-info" }

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package
// A package on this Satellite table
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is string.
    Name interface{}

    // Version. The type is string.
    Version interface{}

    // Is base image. The type is bool.
    IsBaseImage interface{}
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "version" { return "Version" }
    if yname == "is-base-image" { return "IsBaseImage" }
    return ""
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetSegmentPath() string {
    return "package"
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    leafs["version"] = self.Version
    leafs["is-base-image"] = self.IsBaseImage
    return leafs
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetYangName() string { return "package" }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetParent() types.Entity { return self.parent }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetParentYangName() string { return "active-packages" }

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages
// Inactive Packages on this Satellite
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A package on this Satellite table. The type is slice of
    // NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package.
    Package []NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package
}

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetFilter() yfilter.YFilter { return inactivePackages.YFilter }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) SetFilter(yf yfilter.YFilter) { inactivePackages.YFilter = yf }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetGoName(yname string) string {
    if yname == "package" { return "Package" }
    return ""
}

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetSegmentPath() string {
    return "inactive-packages"
}

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        for _, c := range inactivePackages.Package {
            if inactivePackages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package{}
        inactivePackages.Package = append(inactivePackages.Package, child)
        return &inactivePackages.Package[len(inactivePackages.Package)-1]
    }
    return nil
}

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inactivePackages.Package {
        children[inactivePackages.Package[i].GetSegmentPath()] = &inactivePackages.Package[i]
    }
    return children
}

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetBundleName() string { return "cisco_ios_xr" }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetYangName() string { return "inactive-packages" }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) SetParent(parent types.Entity) { inactivePackages.parent = parent }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetParent() types.Entity { return inactivePackages.parent }

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetParentYangName() string { return "install-package-info" }

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package
// A package on this Satellite table
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is string.
    Name interface{}

    // Version. The type is string.
    Version interface{}

    // Is base image. The type is bool.
    IsBaseImage interface{}
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "version" { return "Version" }
    if yname == "is-base-image" { return "IsBaseImage" }
    return ""
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetSegmentPath() string {
    return "package"
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    leafs["version"] = self.Version
    leafs["is-base-image"] = self.IsBaseImage
    return leafs
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetYangName() string { return "package" }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetParent() types.Entity { return self.parent }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetParentYangName() string { return "inactive-packages" }

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages
// Committed Packages running on this Satellite
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A package on this Satellite table. The type is slice of
    // NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package.
    Package []NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package
}

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetFilter() yfilter.YFilter { return committedPackages.YFilter }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) SetFilter(yf yfilter.YFilter) { committedPackages.YFilter = yf }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetGoName(yname string) string {
    if yname == "package" { return "Package" }
    return ""
}

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetSegmentPath() string {
    return "committed-packages"
}

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package" {
        for _, c := range committedPackages.Package {
            if committedPackages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package{}
        committedPackages.Package = append(committedPackages.Package, child)
        return &committedPackages.Package[len(committedPackages.Package)-1]
    }
    return nil
}

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range committedPackages.Package {
        children[committedPackages.Package[i].GetSegmentPath()] = &committedPackages.Package[i]
    }
    return children
}

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetBundleName() string { return "cisco_ios_xr" }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetYangName() string { return "committed-packages" }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) SetParent(parent types.Entity) { committedPackages.parent = parent }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetParent() types.Entity { return committedPackages.parent }

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetParentYangName() string { return "install-package-info" }

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package
// A package on this Satellite table
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is string.
    Name interface{}

    // Version. The type is string.
    Version interface{}

    // Is base image. The type is bool.
    IsBaseImage interface{}
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "version" { return "Version" }
    if yname == "is-base-image" { return "IsBaseImage" }
    return ""
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetSegmentPath() string {
    return "package"
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    leafs["version"] = self.Version
    leafs["is-base-image"] = self.IsBaseImage
    return leafs
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetYangName() string { return "package" }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetParent() types.Entity { return self.parent }

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetParentYangName() string { return "committed-packages" }

// NvSatellite_InstallOpStatuses
// Detailed breakdown of install status table
type NvSatellite_InstallOpStatuses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed breakdown of install status. The type is slice of
    // NvSatellite_InstallOpStatuses_InstallOpStatus.
    InstallOpStatus []NvSatellite_InstallOpStatuses_InstallOpStatus
}

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetFilter() yfilter.YFilter { return installOpStatuses.YFilter }

func (installOpStatuses *NvSatellite_InstallOpStatuses) SetFilter(yf yfilter.YFilter) { installOpStatuses.YFilter = yf }

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetGoName(yname string) string {
    if yname == "install-op-status" { return "InstallOpStatus" }
    return ""
}

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetSegmentPath() string {
    return "install-op-statuses"
}

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "install-op-status" {
        for _, c := range installOpStatuses.InstallOpStatus {
            if installOpStatuses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_InstallOpStatuses_InstallOpStatus{}
        installOpStatuses.InstallOpStatus = append(installOpStatuses.InstallOpStatus, child)
        return &installOpStatuses.InstallOpStatus[len(installOpStatuses.InstallOpStatus)-1]
    }
    return nil
}

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range installOpStatuses.InstallOpStatus {
        children[installOpStatuses.InstallOpStatus[i].GetSegmentPath()] = &installOpStatuses.InstallOpStatus[i]
    }
    return children
}

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetBundleName() string { return "cisco_ios_xr" }

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetYangName() string { return "install-op-statuses" }

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (installOpStatuses *NvSatellite_InstallOpStatuses) SetParent(parent types.Entity) { installOpStatuses.parent = parent }

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetParent() types.Entity { return installOpStatuses.parent }

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_InstallOpStatuses_InstallOpStatus
// Detailed breakdown of install status
type NvSatellite_InstallOpStatuses_InstallOpStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Operation ID. The type is interface{} with range:
    // 0..4294967295.
    OperationId interface{}

    // Operation ID. The type is interface{} with range: 0..4294967295.
    OperationIdXr interface{}

    // Satellite range. The type is string.
    SatelliteRange interface{}

    // Sats not initiated. The type is slice of interface{} with range: 0..65535.
    SatsNotInitiated []interface{}

    // Sats transferring. The type is slice of interface{} with range: 0..65535.
    SatsTransferring []interface{}

    // Sats activating. The type is slice of interface{} with range: 0..65535.
    SatsActivating []interface{}

    // Sats updating. The type is slice of interface{} with range: 0..65535.
    SatsUpdating []interface{}

    // Sats deactivating. The type is slice of interface{} with range: 0..65535.
    SatsDeactivating []interface{}

    // Sats removing. The type is slice of interface{} with range: 0..65535.
    SatsRemoving []interface{}

    // Sats transfer failed. The type is slice of interface{} with range:
    // 0..65535.
    SatsTransferFailed []interface{}

    // Sats activate failed. The type is slice of interface{} with range:
    // 0..65535.
    SatsActivateFailed []interface{}

    // Sats update failed. The type is slice of interface{} with range: 0..65535.
    SatsUpdateFailed []interface{}

    // Sats deactivate failed. The type is slice of interface{} with range:
    // 0..65535.
    SatsDeactivateFailed []interface{}

    // Sats remove failed. The type is slice of interface{} with range: 0..65535.
    SatsRemoveFailed []interface{}

    // Sats transfer aborted. The type is slice of interface{} with range:
    // 0..65535.
    SatsTransferAborted []interface{}

    // Sats activate aborted. The type is slice of interface{} with range:
    // 0..65535.
    SatsActivateAborted []interface{}

    // Sats update aborted. The type is slice of interface{} with range: 0..65535.
    SatsUpdateAborted []interface{}

    // Sats deactivate aborted. The type is slice of interface{} with range:
    // 0..65535.
    SatsDeactivateAborted []interface{}

    // Sats remove aborted. The type is slice of interface{} with range: 0..65535.
    SatsRemoveAborted []interface{}

    // Sats no operation. The type is slice of interface{} with range: 0..65535.
    SatsNoOperation []interface{}

    // Sats completed. The type is slice of interface{} with range: 0..65535.
    SatsCompleted []interface{}
}

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetFilter() yfilter.YFilter { return installOpStatus.YFilter }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) SetFilter(yf yfilter.YFilter) { installOpStatus.YFilter = yf }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetGoName(yname string) string {
    if yname == "operation-id" { return "OperationId" }
    if yname == "operation-id-xr" { return "OperationIdXr" }
    if yname == "satellite-range" { return "SatelliteRange" }
    if yname == "sats-not-initiated" { return "SatsNotInitiated" }
    if yname == "sats-transferring" { return "SatsTransferring" }
    if yname == "sats-activating" { return "SatsActivating" }
    if yname == "sats-updating" { return "SatsUpdating" }
    if yname == "sats-deactivating" { return "SatsDeactivating" }
    if yname == "sats-removing" { return "SatsRemoving" }
    if yname == "sats-transfer-failed" { return "SatsTransferFailed" }
    if yname == "sats-activate-failed" { return "SatsActivateFailed" }
    if yname == "sats-update-failed" { return "SatsUpdateFailed" }
    if yname == "sats-deactivate-failed" { return "SatsDeactivateFailed" }
    if yname == "sats-remove-failed" { return "SatsRemoveFailed" }
    if yname == "sats-transfer-aborted" { return "SatsTransferAborted" }
    if yname == "sats-activate-aborted" { return "SatsActivateAborted" }
    if yname == "sats-update-aborted" { return "SatsUpdateAborted" }
    if yname == "sats-deactivate-aborted" { return "SatsDeactivateAborted" }
    if yname == "sats-remove-aborted" { return "SatsRemoveAborted" }
    if yname == "sats-no-operation" { return "SatsNoOperation" }
    if yname == "sats-completed" { return "SatsCompleted" }
    return ""
}

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetSegmentPath() string {
    return "install-op-status" + "[operation-id='" + fmt.Sprintf("%v", installOpStatus.OperationId) + "']"
}

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-id"] = installOpStatus.OperationId
    leafs["operation-id-xr"] = installOpStatus.OperationIdXr
    leafs["satellite-range"] = installOpStatus.SatelliteRange
    leafs["sats-not-initiated"] = installOpStatus.SatsNotInitiated
    leafs["sats-transferring"] = installOpStatus.SatsTransferring
    leafs["sats-activating"] = installOpStatus.SatsActivating
    leafs["sats-updating"] = installOpStatus.SatsUpdating
    leafs["sats-deactivating"] = installOpStatus.SatsDeactivating
    leafs["sats-removing"] = installOpStatus.SatsRemoving
    leafs["sats-transfer-failed"] = installOpStatus.SatsTransferFailed
    leafs["sats-activate-failed"] = installOpStatus.SatsActivateFailed
    leafs["sats-update-failed"] = installOpStatus.SatsUpdateFailed
    leafs["sats-deactivate-failed"] = installOpStatus.SatsDeactivateFailed
    leafs["sats-remove-failed"] = installOpStatus.SatsRemoveFailed
    leafs["sats-transfer-aborted"] = installOpStatus.SatsTransferAborted
    leafs["sats-activate-aborted"] = installOpStatus.SatsActivateAborted
    leafs["sats-update-aborted"] = installOpStatus.SatsUpdateAborted
    leafs["sats-deactivate-aborted"] = installOpStatus.SatsDeactivateAborted
    leafs["sats-remove-aborted"] = installOpStatus.SatsRemoveAborted
    leafs["sats-no-operation"] = installOpStatus.SatsNoOperation
    leafs["sats-completed"] = installOpStatus.SatsCompleted
    return leafs
}

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetBundleName() string { return "cisco_ios_xr" }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetYangName() string { return "install-op-status" }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) SetParent(parent types.Entity) { installOpStatus.parent = parent }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetParent() types.Entity { return installOpStatus.parent }

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetParentYangName() string { return "install-op-statuses" }

// NvSatellite_SatelliteProperties
// ICPE GCO operational information
type NvSatellite_SatelliteProperties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite ID range table.
    IdRanges NvSatellite_SatelliteProperties_IdRanges
}

func (satelliteProperties *NvSatellite_SatelliteProperties) GetFilter() yfilter.YFilter { return satelliteProperties.YFilter }

func (satelliteProperties *NvSatellite_SatelliteProperties) SetFilter(yf yfilter.YFilter) { satelliteProperties.YFilter = yf }

func (satelliteProperties *NvSatellite_SatelliteProperties) GetGoName(yname string) string {
    if yname == "id-ranges" { return "IdRanges" }
    return ""
}

func (satelliteProperties *NvSatellite_SatelliteProperties) GetSegmentPath() string {
    return "satellite-properties"
}

func (satelliteProperties *NvSatellite_SatelliteProperties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "id-ranges" {
        return &satelliteProperties.IdRanges
    }
    return nil
}

func (satelliteProperties *NvSatellite_SatelliteProperties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["id-ranges"] = &satelliteProperties.IdRanges
    return children
}

func (satelliteProperties *NvSatellite_SatelliteProperties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (satelliteProperties *NvSatellite_SatelliteProperties) GetBundleName() string { return "cisco_ios_xr" }

func (satelliteProperties *NvSatellite_SatelliteProperties) GetYangName() string { return "satellite-properties" }

func (satelliteProperties *NvSatellite_SatelliteProperties) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satelliteProperties *NvSatellite_SatelliteProperties) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satelliteProperties *NvSatellite_SatelliteProperties) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satelliteProperties *NvSatellite_SatelliteProperties) SetParent(parent types.Entity) { satelliteProperties.parent = parent }

func (satelliteProperties *NvSatellite_SatelliteProperties) GetParent() types.Entity { return satelliteProperties.parent }

func (satelliteProperties *NvSatellite_SatelliteProperties) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_SatelliteProperties_IdRanges
// Satellite ID range table
type NvSatellite_SatelliteProperties_IdRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite ID range. The type is slice of
    // NvSatellite_SatelliteProperties_IdRanges_IdRange.
    IdRange []NvSatellite_SatelliteProperties_IdRanges_IdRange
}

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetFilter() yfilter.YFilter { return idRanges.YFilter }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) SetFilter(yf yfilter.YFilter) { idRanges.YFilter = yf }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetGoName(yname string) string {
    if yname == "id-range" { return "IdRange" }
    return ""
}

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetSegmentPath() string {
    return "id-ranges"
}

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "id-range" {
        for _, c := range idRanges.IdRange {
            if idRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SatelliteProperties_IdRanges_IdRange{}
        idRanges.IdRange = append(idRanges.IdRange, child)
        return &idRanges.IdRange[len(idRanges.IdRange)-1]
    }
    return nil
}

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range idRanges.IdRange {
        children[idRanges.IdRange[i].GetSegmentPath()] = &idRanges.IdRange[i]
    }
    return children
}

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetBundleName() string { return "cisco_ios_xr" }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetYangName() string { return "id-ranges" }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) SetParent(parent types.Entity) { idRanges.parent = parent }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetParent() types.Entity { return idRanges.parent }

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetParentYangName() string { return "satellite-properties" }

// NvSatellite_SatelliteProperties_IdRanges_IdRange
// Satellite ID range
type NvSatellite_SatelliteProperties_IdRanges_IdRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sat ID range. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SatIdRange interface{}

    // Min. The type is interface{} with range: 0..2147483647.
    Min interface{}

    // Max. The type is interface{} with range: 0..2147483647.
    Max interface{}
}

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetFilter() yfilter.YFilter { return idRange.YFilter }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) SetFilter(yf yfilter.YFilter) { idRange.YFilter = yf }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetGoName(yname string) string {
    if yname == "sat-id-range" { return "SatIdRange" }
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetSegmentPath() string {
    return "id-range" + "[sat-id-range='" + fmt.Sprintf("%v", idRange.SatIdRange) + "']"
}

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sat-id-range"] = idRange.SatIdRange
    leafs["min"] = idRange.Min
    leafs["max"] = idRange.Max
    return leafs
}

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetBundleName() string { return "cisco_ios_xr" }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetYangName() string { return "id-range" }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) SetParent(parent types.Entity) { idRange.parent = parent }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetParent() types.Entity { return idRange.parent }

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetParentYangName() string { return "id-ranges" }

// NvSatellite_SdacpDiscovery2S
// ICPE Configured interface state information
// table
type NvSatellite_SdacpDiscovery2S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICPE Configured interface state information. The type is slice of
    // NvSatellite_SdacpDiscovery2S_SdacpDiscovery2.
    SdacpDiscovery2 []NvSatellite_SdacpDiscovery2S_SdacpDiscovery2
}

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetFilter() yfilter.YFilter { return sdacpDiscovery2S.YFilter }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) SetFilter(yf yfilter.YFilter) { sdacpDiscovery2S.YFilter = yf }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetGoName(yname string) string {
    if yname == "sdacp-discovery2" { return "SdacpDiscovery2" }
    return ""
}

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetSegmentPath() string {
    return "Cisco-IOS-XR-icpe-sdacp-oper:sdacp-discovery2s"
}

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sdacp-discovery2" {
        for _, c := range sdacpDiscovery2S.SdacpDiscovery2 {
            if sdacpDiscovery2S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpDiscovery2S_SdacpDiscovery2{}
        sdacpDiscovery2S.SdacpDiscovery2 = append(sdacpDiscovery2S.SdacpDiscovery2, child)
        return &sdacpDiscovery2S.SdacpDiscovery2[len(sdacpDiscovery2S.SdacpDiscovery2)-1]
    }
    return nil
}

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sdacpDiscovery2S.SdacpDiscovery2 {
        children[sdacpDiscovery2S.SdacpDiscovery2[i].GetSegmentPath()] = &sdacpDiscovery2S.SdacpDiscovery2[i]
    }
    return children
}

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetBundleName() string { return "cisco_ios_xr" }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetYangName() string { return "sdacp-discovery2s" }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) SetParent(parent types.Entity) { sdacpDiscovery2S.parent = parent }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetParent() types.Entity { return sdacpDiscovery2S.parent }

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_SdacpDiscovery2S_SdacpDiscovery2
// ICPE Configured interface state information
type NvSatellite_SdacpDiscovery2S_SdacpDiscovery2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // ICPE Discovery interface state information table. The type is slice of
    // NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface.
    Interface []NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface

    // ICPE Satellite state information table. The type is slice of
    // NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite.
    Satellite []NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite
}

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetFilter() yfilter.YFilter { return sdacpDiscovery2.YFilter }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) SetFilter(yf yfilter.YFilter) { sdacpDiscovery2.YFilter = yf }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface" { return "Interface" }
    if yname == "satellite" { return "Satellite" }
    return ""
}

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetSegmentPath() string {
    return "sdacp-discovery2" + "[interface-name='" + fmt.Sprintf("%v", sdacpDiscovery2.InterfaceName) + "']"
}

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range sdacpDiscovery2.Interface {
            if sdacpDiscovery2.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface{}
        sdacpDiscovery2.Interface = append(sdacpDiscovery2.Interface, child)
        return &sdacpDiscovery2.Interface[len(sdacpDiscovery2.Interface)-1]
    }
    if childYangName == "satellite" {
        for _, c := range sdacpDiscovery2.Satellite {
            if sdacpDiscovery2.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite{}
        sdacpDiscovery2.Satellite = append(sdacpDiscovery2.Satellite, child)
        return &sdacpDiscovery2.Satellite[len(sdacpDiscovery2.Satellite)-1]
    }
    return nil
}

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sdacpDiscovery2.Interface {
        children[sdacpDiscovery2.Interface[i].GetSegmentPath()] = &sdacpDiscovery2.Interface[i]
    }
    for i := range sdacpDiscovery2.Satellite {
        children[sdacpDiscovery2.Satellite[i].GetSegmentPath()] = &sdacpDiscovery2.Satellite[i]
    }
    return children
}

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = sdacpDiscovery2.InterfaceName
    leafs["interface-name-xr"] = sdacpDiscovery2.InterfaceNameXr
    return leafs
}

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetBundleName() string { return "cisco_ios_xr" }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetYangName() string { return "sdacp-discovery2" }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) SetParent(parent types.Entity) { sdacpDiscovery2.parent = parent }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetParent() types.Entity { return sdacpDiscovery2.parent }

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetParentYangName() string { return "sdacp-discovery2s" }

// NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface
// ICPE Discovery interface state information table
type NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface status. The type is DpmProtoState.
    InterfaceStatus interface{}
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-status" { return "InterfaceStatus" }
    return ""
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-status"] = self.InterfaceStatus
    return leafs
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetYangName() string { return "interface" }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetParent() types.Entity { return self.parent }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetParentYangName() string { return "sdacp-discovery2" }

// NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite
// ICPE Satellite state information table
type NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteId interface{}

    // Satellite status. The type is DpmProtoState.
    SatelliteStatus interface{}

    // Conflict reason. The type is interface{} with range: 0..4294967295.
    ConflictReason interface{}

    // Satellite IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SatelliteIpAddress interface{}

    // Host IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    HostIpAddress interface{}

    // ICPE Discovered satellite state information table. The type is slice of
    // NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface.
    Interface []NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface
}

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetFilter() yfilter.YFilter { return satellite.YFilter }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) SetFilter(yf yfilter.YFilter) { satellite.YFilter = yf }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "satellite-status" { return "SatelliteStatus" }
    if yname == "conflict-reason" { return "ConflictReason" }
    if yname == "satellite-ip-address" { return "SatelliteIpAddress" }
    if yname == "host-ip-address" { return "HostIpAddress" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetSegmentPath() string {
    return "satellite"
}

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range satellite.Interface {
            if satellite.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface{}
        satellite.Interface = append(satellite.Interface, child)
        return &satellite.Interface[len(satellite.Interface)-1]
    }
    return nil
}

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range satellite.Interface {
        children[satellite.Interface[i].GetSegmentPath()] = &satellite.Interface[i]
    }
    return children
}

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = satellite.SatelliteId
    leafs["satellite-status"] = satellite.SatelliteStatus
    leafs["conflict-reason"] = satellite.ConflictReason
    leafs["satellite-ip-address"] = satellite.SatelliteIpAddress
    leafs["host-ip-address"] = satellite.HostIpAddress
    return leafs
}

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetBundleName() string { return "cisco_ios_xr" }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetYangName() string { return "satellite" }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) SetParent(parent types.Entity) { satellite.parent = parent }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetParent() types.Entity { return satellite.parent }

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetParentYangName() string { return "sdacp-discovery2" }

// NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface
// ICPE Discovered satellite state information
// table
type NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // Satellite status. The type is DpmProtoState.
    SatelliteStatus interface{}

    // Conflict reason. The type is interface{} with range: 0..4294967295.
    ConflictReason interface{}

    // Satellite chassis vendor. The type is string.
    SatelliteChassisVendor interface{}

    // Satellite interface ID. The type is interface{} with range: 0..4294967295.
    SatelliteInterfaceId interface{}

    // Satellite interface MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SatelliteInterfaceMac interface{}

    // Satellite chassis MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SatelliteChassisMac interface{}

    // Satellite serial id. The type is string.
    SatelliteSerialId interface{}

    // Satellite module vendor. The type is string.
    SatelliteModuleVendor interface{}
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetGoName(yname string) string {
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "satellite-status" { return "SatelliteStatus" }
    if yname == "conflict-reason" { return "ConflictReason" }
    if yname == "satellite-chassis-vendor" { return "SatelliteChassisVendor" }
    if yname == "satellite-interface-id" { return "SatelliteInterfaceId" }
    if yname == "satellite-interface-mac" { return "SatelliteInterfaceMac" }
    if yname == "satellite-chassis-mac" { return "SatelliteChassisMac" }
    if yname == "satellite-serial-id" { return "SatelliteSerialId" }
    if yname == "satellite-module-vendor" { return "SatelliteModuleVendor" }
    return ""
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-handle"] = self.InterfaceHandle
    leafs["satellite-status"] = self.SatelliteStatus
    leafs["conflict-reason"] = self.ConflictReason
    leafs["satellite-chassis-vendor"] = self.SatelliteChassisVendor
    leafs["satellite-interface-id"] = self.SatelliteInterfaceId
    leafs["satellite-interface-mac"] = self.SatelliteInterfaceMac
    leafs["satellite-chassis-mac"] = self.SatelliteChassisMac
    leafs["satellite-serial-id"] = self.SatelliteSerialId
    leafs["satellite-module-vendor"] = self.SatelliteModuleVendor
    return leafs
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetYangName() string { return "interface" }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetParent() types.Entity { return self.parent }

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetParentYangName() string { return "satellite" }

// NvSatellite_IcpeDpms
// ICPE DPM operational information table
type NvSatellite_IcpeDpms struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICPE DPM operational information. The type is slice of
    // NvSatellite_IcpeDpms_IcpeDpm.
    IcpeDpm []NvSatellite_IcpeDpms_IcpeDpm
}

func (icpeDpms *NvSatellite_IcpeDpms) GetFilter() yfilter.YFilter { return icpeDpms.YFilter }

func (icpeDpms *NvSatellite_IcpeDpms) SetFilter(yf yfilter.YFilter) { icpeDpms.YFilter = yf }

func (icpeDpms *NvSatellite_IcpeDpms) GetGoName(yname string) string {
    if yname == "icpe-dpm" { return "IcpeDpm" }
    return ""
}

func (icpeDpms *NvSatellite_IcpeDpms) GetSegmentPath() string {
    return "Cisco-IOS-XR-icpe-sdacp-oper:icpe-dpms"
}

func (icpeDpms *NvSatellite_IcpeDpms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "icpe-dpm" {
        for _, c := range icpeDpms.IcpeDpm {
            if icpeDpms.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_IcpeDpms_IcpeDpm{}
        icpeDpms.IcpeDpm = append(icpeDpms.IcpeDpm, child)
        return &icpeDpms.IcpeDpm[len(icpeDpms.IcpeDpm)-1]
    }
    return nil
}

func (icpeDpms *NvSatellite_IcpeDpms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range icpeDpms.IcpeDpm {
        children[icpeDpms.IcpeDpm[i].GetSegmentPath()] = &icpeDpms.IcpeDpm[i]
    }
    return children
}

func (icpeDpms *NvSatellite_IcpeDpms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (icpeDpms *NvSatellite_IcpeDpms) GetBundleName() string { return "cisco_ios_xr" }

func (icpeDpms *NvSatellite_IcpeDpms) GetYangName() string { return "icpe-dpms" }

func (icpeDpms *NvSatellite_IcpeDpms) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (icpeDpms *NvSatellite_IcpeDpms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (icpeDpms *NvSatellite_IcpeDpms) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (icpeDpms *NvSatellite_IcpeDpms) SetParent(parent types.Entity) { icpeDpms.parent = parent }

func (icpeDpms *NvSatellite_IcpeDpms) GetParent() types.Entity { return icpeDpms.parent }

func (icpeDpms *NvSatellite_IcpeDpms) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_IcpeDpms_IcpeDpm
// ICPE DPM operational information
type NvSatellite_IcpeDpms_IcpeDpm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Discovery interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    DiscoveryInterface interface{}

    // Discovery interface. The type is string.
    DiscoveryInterfaceXr interface{}

    // Discovery interface handle. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    DiscoveryInterfaceHandle interface{}

    // Discovery interface status. The type is DpmProtoState.
    DiscoveryInterfaceStatus interface{}

    // Ident packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    IdentPacketsReceived interface{}

    // Ready packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    ReadyPacketsReceived interface{}

    // Los packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    LosPacketsReceived interface{}

    // Invalid packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InvalidPacketsReceived interface{}

    // Configuration packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfigurationPacketsSent interface{}

    // Ack packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AckPacketsSent interface{}

    // Reject packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectPacketsSent interface{}

    // Probe packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ProbePacketsSent interface{}

    // Host ack packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    HostAckPacketsReceived interface{}

    // Host ack packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    HostAckPacketsSent interface{}

    // Secs since pkts cleaned. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecsSincePktsCleaned interface{}

    // ICPE DPM satellite operational information table. The type is slice of
    // NvSatellite_IcpeDpms_IcpeDpm_Satellite.
    Satellite []NvSatellite_IcpeDpms_IcpeDpm_Satellite

    // ICPE DPM remote host operational information table. The type is slice of
    // NvSatellite_IcpeDpms_IcpeDpm_RemoteHost.
    RemoteHost []NvSatellite_IcpeDpms_IcpeDpm_RemoteHost
}

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetFilter() yfilter.YFilter { return icpeDpm.YFilter }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) SetFilter(yf yfilter.YFilter) { icpeDpm.YFilter = yf }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetGoName(yname string) string {
    if yname == "discovery-interface" { return "DiscoveryInterface" }
    if yname == "discovery-interface-xr" { return "DiscoveryInterfaceXr" }
    if yname == "discovery-interface-handle" { return "DiscoveryInterfaceHandle" }
    if yname == "discovery-interface-status" { return "DiscoveryInterfaceStatus" }
    if yname == "ident-packets-received" { return "IdentPacketsReceived" }
    if yname == "ready-packets-received" { return "ReadyPacketsReceived" }
    if yname == "los-packets-received" { return "LosPacketsReceived" }
    if yname == "invalid-packets-received" { return "InvalidPacketsReceived" }
    if yname == "configuration-packets-sent" { return "ConfigurationPacketsSent" }
    if yname == "ack-packets-sent" { return "AckPacketsSent" }
    if yname == "reject-packets-sent" { return "RejectPacketsSent" }
    if yname == "probe-packets-sent" { return "ProbePacketsSent" }
    if yname == "host-ack-packets-received" { return "HostAckPacketsReceived" }
    if yname == "host-ack-packets-sent" { return "HostAckPacketsSent" }
    if yname == "secs-since-pkts-cleaned" { return "SecsSincePktsCleaned" }
    if yname == "satellite" { return "Satellite" }
    if yname == "remote-host" { return "RemoteHost" }
    return ""
}

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetSegmentPath() string {
    return "icpe-dpm" + "[discovery-interface='" + fmt.Sprintf("%v", icpeDpm.DiscoveryInterface) + "']"
}

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "satellite" {
        for _, c := range icpeDpm.Satellite {
            if icpeDpm.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_IcpeDpms_IcpeDpm_Satellite{}
        icpeDpm.Satellite = append(icpeDpm.Satellite, child)
        return &icpeDpm.Satellite[len(icpeDpm.Satellite)-1]
    }
    if childYangName == "remote-host" {
        for _, c := range icpeDpm.RemoteHost {
            if icpeDpm.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_IcpeDpms_IcpeDpm_RemoteHost{}
        icpeDpm.RemoteHost = append(icpeDpm.RemoteHost, child)
        return &icpeDpm.RemoteHost[len(icpeDpm.RemoteHost)-1]
    }
    return nil
}

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range icpeDpm.Satellite {
        children[icpeDpm.Satellite[i].GetSegmentPath()] = &icpeDpm.Satellite[i]
    }
    for i := range icpeDpm.RemoteHost {
        children[icpeDpm.RemoteHost[i].GetSegmentPath()] = &icpeDpm.RemoteHost[i]
    }
    return children
}

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discovery-interface"] = icpeDpm.DiscoveryInterface
    leafs["discovery-interface-xr"] = icpeDpm.DiscoveryInterfaceXr
    leafs["discovery-interface-handle"] = icpeDpm.DiscoveryInterfaceHandle
    leafs["discovery-interface-status"] = icpeDpm.DiscoveryInterfaceStatus
    leafs["ident-packets-received"] = icpeDpm.IdentPacketsReceived
    leafs["ready-packets-received"] = icpeDpm.ReadyPacketsReceived
    leafs["los-packets-received"] = icpeDpm.LosPacketsReceived
    leafs["invalid-packets-received"] = icpeDpm.InvalidPacketsReceived
    leafs["configuration-packets-sent"] = icpeDpm.ConfigurationPacketsSent
    leafs["ack-packets-sent"] = icpeDpm.AckPacketsSent
    leafs["reject-packets-sent"] = icpeDpm.RejectPacketsSent
    leafs["probe-packets-sent"] = icpeDpm.ProbePacketsSent
    leafs["host-ack-packets-received"] = icpeDpm.HostAckPacketsReceived
    leafs["host-ack-packets-sent"] = icpeDpm.HostAckPacketsSent
    leafs["secs-since-pkts-cleaned"] = icpeDpm.SecsSincePktsCleaned
    return leafs
}

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetBundleName() string { return "cisco_ios_xr" }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetYangName() string { return "icpe-dpm" }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) SetParent(parent types.Entity) { icpeDpm.parent = parent }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetParent() types.Entity { return icpeDpm.parent }

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetParentYangName() string { return "icpe-dpms" }

// NvSatellite_IcpeDpms_IcpeDpm_Satellite
// ICPE DPM satellite operational information table
type NvSatellite_IcpeDpms_IcpeDpm_Satellite struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteId interface{}

    // Satellite interface ID. The type is interface{} with range: 0..4294967295.
    SatelliteInterfaceId interface{}

    // Satellite interface MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SatelliteInterfaceMac interface{}

    // Satellite IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SatelliteIpAddress interface{}

    // Host IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    HostIpAddress interface{}

    // Satellite chassis type. The type is string.
    SatelliteChassisType interface{}

    // Satellite chassis vendor. The type is string.
    SatelliteChassisVendor interface{}

    // Satellite chassis MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SatelliteChassisMac interface{}

    // Satellite serial id. The type is string.
    SatelliteSerialId interface{}

    // Satellite module type. The type is string.
    SatelliteModuleType interface{}

    // Satellite module vendor. The type is string.
    SatelliteModuleVendor interface{}

    // Satellite module MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SatelliteModuleMac interface{}

    // Conflict reason. The type is interface{} with range: 0..4294967295.
    ConflictReason interface{}

    // Received sys MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    ReceivedSysMac interface{}

    // Host chassis type. The type is string.
    HostChassisType interface{}

    // Host chassis vendor. The type is string.
    HostChassisVendor interface{}

    // Host chassis MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    HostChassisMac interface{}

    // Discovery protocol state. The type is DpmProtoState.
    DiscoveryProtocolState interface{}

    // Last IMDR state. The type is interface{} with range: 0..4294967295.
    LastImdrState interface{}

    // Current timeout. The type is interface{} with range: 0..4294967295.
    CurrentTimeout interface{}

    // Is queued for EFD. The type is bool.
    IsQueuedForEfd interface{}

    // Is queued for OC. The type is bool.
    IsQueuedForOc interface{}

    // Ifmgr state. The type is bool.
    IfmgrState interface{}

    // Legacy. The type is bool.
    Legacy interface{}

    // Deleting. The type is bool.
    Deleting interface{}

    // Ident packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    IdentPacketsReceived interface{}

    // Ready packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    ReadyPacketsReceived interface{}

    // Los packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    LosPacketsReceived interface{}

    // Invalid packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    InvalidPacketsReceived interface{}

    // Configuration packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfigurationPacketsSent interface{}

    // Ack packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AckPacketsSent interface{}

    // Reject packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectPacketsSent interface{}

    // Secs since pkts cleaned. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecsSincePktsCleaned interface{}
}

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetFilter() yfilter.YFilter { return satellite.YFilter }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) SetFilter(yf yfilter.YFilter) { satellite.YFilter = yf }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "satellite-interface-id" { return "SatelliteInterfaceId" }
    if yname == "satellite-interface-mac" { return "SatelliteInterfaceMac" }
    if yname == "satellite-ip-address" { return "SatelliteIpAddress" }
    if yname == "host-ip-address" { return "HostIpAddress" }
    if yname == "satellite-chassis-type" { return "SatelliteChassisType" }
    if yname == "satellite-chassis-vendor" { return "SatelliteChassisVendor" }
    if yname == "satellite-chassis-mac" { return "SatelliteChassisMac" }
    if yname == "satellite-serial-id" { return "SatelliteSerialId" }
    if yname == "satellite-module-type" { return "SatelliteModuleType" }
    if yname == "satellite-module-vendor" { return "SatelliteModuleVendor" }
    if yname == "satellite-module-mac" { return "SatelliteModuleMac" }
    if yname == "conflict-reason" { return "ConflictReason" }
    if yname == "received-sys-mac" { return "ReceivedSysMac" }
    if yname == "host-chassis-type" { return "HostChassisType" }
    if yname == "host-chassis-vendor" { return "HostChassisVendor" }
    if yname == "host-chassis-mac" { return "HostChassisMac" }
    if yname == "discovery-protocol-state" { return "DiscoveryProtocolState" }
    if yname == "last-imdr-state" { return "LastImdrState" }
    if yname == "current-timeout" { return "CurrentTimeout" }
    if yname == "is-queued-for-efd" { return "IsQueuedForEfd" }
    if yname == "is-queued-for-oc" { return "IsQueuedForOc" }
    if yname == "ifmgr-state" { return "IfmgrState" }
    if yname == "legacy" { return "Legacy" }
    if yname == "deleting" { return "Deleting" }
    if yname == "ident-packets-received" { return "IdentPacketsReceived" }
    if yname == "ready-packets-received" { return "ReadyPacketsReceived" }
    if yname == "los-packets-received" { return "LosPacketsReceived" }
    if yname == "invalid-packets-received" { return "InvalidPacketsReceived" }
    if yname == "configuration-packets-sent" { return "ConfigurationPacketsSent" }
    if yname == "ack-packets-sent" { return "AckPacketsSent" }
    if yname == "reject-packets-sent" { return "RejectPacketsSent" }
    if yname == "secs-since-pkts-cleaned" { return "SecsSincePktsCleaned" }
    return ""
}

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetSegmentPath() string {
    return "satellite"
}

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = satellite.SatelliteId
    leafs["satellite-interface-id"] = satellite.SatelliteInterfaceId
    leafs["satellite-interface-mac"] = satellite.SatelliteInterfaceMac
    leafs["satellite-ip-address"] = satellite.SatelliteIpAddress
    leafs["host-ip-address"] = satellite.HostIpAddress
    leafs["satellite-chassis-type"] = satellite.SatelliteChassisType
    leafs["satellite-chassis-vendor"] = satellite.SatelliteChassisVendor
    leafs["satellite-chassis-mac"] = satellite.SatelliteChassisMac
    leafs["satellite-serial-id"] = satellite.SatelliteSerialId
    leafs["satellite-module-type"] = satellite.SatelliteModuleType
    leafs["satellite-module-vendor"] = satellite.SatelliteModuleVendor
    leafs["satellite-module-mac"] = satellite.SatelliteModuleMac
    leafs["conflict-reason"] = satellite.ConflictReason
    leafs["received-sys-mac"] = satellite.ReceivedSysMac
    leafs["host-chassis-type"] = satellite.HostChassisType
    leafs["host-chassis-vendor"] = satellite.HostChassisVendor
    leafs["host-chassis-mac"] = satellite.HostChassisMac
    leafs["discovery-protocol-state"] = satellite.DiscoveryProtocolState
    leafs["last-imdr-state"] = satellite.LastImdrState
    leafs["current-timeout"] = satellite.CurrentTimeout
    leafs["is-queued-for-efd"] = satellite.IsQueuedForEfd
    leafs["is-queued-for-oc"] = satellite.IsQueuedForOc
    leafs["ifmgr-state"] = satellite.IfmgrState
    leafs["legacy"] = satellite.Legacy
    leafs["deleting"] = satellite.Deleting
    leafs["ident-packets-received"] = satellite.IdentPacketsReceived
    leafs["ready-packets-received"] = satellite.ReadyPacketsReceived
    leafs["los-packets-received"] = satellite.LosPacketsReceived
    leafs["invalid-packets-received"] = satellite.InvalidPacketsReceived
    leafs["configuration-packets-sent"] = satellite.ConfigurationPacketsSent
    leafs["ack-packets-sent"] = satellite.AckPacketsSent
    leafs["reject-packets-sent"] = satellite.RejectPacketsSent
    leafs["secs-since-pkts-cleaned"] = satellite.SecsSincePktsCleaned
    return leafs
}

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetBundleName() string { return "cisco_ios_xr" }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetYangName() string { return "satellite" }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) SetParent(parent types.Entity) { satellite.parent = parent }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetParent() types.Entity { return satellite.parent }

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetParentYangName() string { return "icpe-dpm" }

// NvSatellite_IcpeDpms_IcpeDpm_RemoteHost
// ICPE DPM remote host operational information
// table
type NvSatellite_IcpeDpms_IcpeDpm_RemoteHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Host chassis MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    HostChassisMac interface{}

    // Host interface MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    HostInterfaceMac interface{}

    // Discovery protocol state. The type is DpmProtoHostState.
    DiscoveryProtocolState interface{}

    // Route length. The type is interface{} with range: 0..4294967295.
    RouteLength interface{}

    // Current timeout. The type is interface{} with range: 0..4294967295.
    CurrentTimeout interface{}

    // Host ack packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    HostAckPacketsReceived interface{}

    // Host ack packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    HostAckPacketsSent interface{}

    // Secs since pkts cleaned. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecsSincePktsCleaned interface{}
}

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetFilter() yfilter.YFilter { return remoteHost.YFilter }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) SetFilter(yf yfilter.YFilter) { remoteHost.YFilter = yf }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetGoName(yname string) string {
    if yname == "host-chassis-mac" { return "HostChassisMac" }
    if yname == "host-interface-mac" { return "HostInterfaceMac" }
    if yname == "discovery-protocol-state" { return "DiscoveryProtocolState" }
    if yname == "route-length" { return "RouteLength" }
    if yname == "current-timeout" { return "CurrentTimeout" }
    if yname == "host-ack-packets-received" { return "HostAckPacketsReceived" }
    if yname == "host-ack-packets-sent" { return "HostAckPacketsSent" }
    if yname == "secs-since-pkts-cleaned" { return "SecsSincePktsCleaned" }
    return ""
}

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetSegmentPath() string {
    return "remote-host"
}

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-chassis-mac"] = remoteHost.HostChassisMac
    leafs["host-interface-mac"] = remoteHost.HostInterfaceMac
    leafs["discovery-protocol-state"] = remoteHost.DiscoveryProtocolState
    leafs["route-length"] = remoteHost.RouteLength
    leafs["current-timeout"] = remoteHost.CurrentTimeout
    leafs["host-ack-packets-received"] = remoteHost.HostAckPacketsReceived
    leafs["host-ack-packets-sent"] = remoteHost.HostAckPacketsSent
    leafs["secs-since-pkts-cleaned"] = remoteHost.SecsSincePktsCleaned
    return leafs
}

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetBundleName() string { return "cisco_ios_xr" }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetYangName() string { return "remote-host" }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) SetParent(parent types.Entity) { remoteHost.parent = parent }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetParent() types.Entity { return remoteHost.parent }

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetParentYangName() string { return "icpe-dpm" }

// NvSatellite_SdacpControls
// SDAC Protocol Discovery information table
type NvSatellite_SdacpControls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SDAC Protocol Discovery information. The type is slice of
    // NvSatellite_SdacpControls_SdacpControl.
    SdacpControl []NvSatellite_SdacpControls_SdacpControl
}

func (sdacpControls *NvSatellite_SdacpControls) GetFilter() yfilter.YFilter { return sdacpControls.YFilter }

func (sdacpControls *NvSatellite_SdacpControls) SetFilter(yf yfilter.YFilter) { sdacpControls.YFilter = yf }

func (sdacpControls *NvSatellite_SdacpControls) GetGoName(yname string) string {
    if yname == "sdacp-control" { return "SdacpControl" }
    return ""
}

func (sdacpControls *NvSatellite_SdacpControls) GetSegmentPath() string {
    return "Cisco-IOS-XR-icpe-sdacp-oper:sdacp-controls"
}

func (sdacpControls *NvSatellite_SdacpControls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sdacp-control" {
        for _, c := range sdacpControls.SdacpControl {
            if sdacpControls.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpControls_SdacpControl{}
        sdacpControls.SdacpControl = append(sdacpControls.SdacpControl, child)
        return &sdacpControls.SdacpControl[len(sdacpControls.SdacpControl)-1]
    }
    return nil
}

func (sdacpControls *NvSatellite_SdacpControls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sdacpControls.SdacpControl {
        children[sdacpControls.SdacpControl[i].GetSegmentPath()] = &sdacpControls.SdacpControl[i]
    }
    return children
}

func (sdacpControls *NvSatellite_SdacpControls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sdacpControls *NvSatellite_SdacpControls) GetBundleName() string { return "cisco_ios_xr" }

func (sdacpControls *NvSatellite_SdacpControls) GetYangName() string { return "sdacp-controls" }

func (sdacpControls *NvSatellite_SdacpControls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdacpControls *NvSatellite_SdacpControls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdacpControls *NvSatellite_SdacpControls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdacpControls *NvSatellite_SdacpControls) SetParent(parent types.Entity) { sdacpControls.parent = parent }

func (sdacpControls *NvSatellite_SdacpControls) GetParent() types.Entity { return sdacpControls.parent }

func (sdacpControls *NvSatellite_SdacpControls) GetParentYangName() string { return "nv-satellite" }

// NvSatellite_SdacpControls_SdacpControl
// SDAC Protocol Discovery information
type NvSatellite_SdacpControls_SdacpControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite ID. The type is interface{} with range:
    // 100..65534.
    SatelliteId interface{}

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteIdXr interface{}

    // Satellite IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SatelliteIpAddress interface{}

    // IP address auto. The type is bool.
    IpAddressAuto interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Control protocol state. The type is IcpeCpmControlFsmState.
    ControlProtocolState interface{}

    // Transport error. The type is interface{} with range: 0..4294967295.
    TransportError interface{}

    // Timestamp.
    ProtocolStateTimestamp NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp

    // Timestamp.
    TransportErrorTimestamp NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp

    // Channels on satellite table. The type is slice of
    // NvSatellite_SdacpControls_SdacpControl_Channel.
    Channel []NvSatellite_SdacpControls_SdacpControl_Channel
}

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetFilter() yfilter.YFilter { return sdacpControl.YFilter }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) SetFilter(yf yfilter.YFilter) { sdacpControl.YFilter = yf }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetGoName(yname string) string {
    if yname == "satellite-id" { return "SatelliteId" }
    if yname == "satellite-id-xr" { return "SatelliteIdXr" }
    if yname == "satellite-ip-address" { return "SatelliteIpAddress" }
    if yname == "ip-address-auto" { return "IpAddressAuto" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "control-protocol-state" { return "ControlProtocolState" }
    if yname == "transport-error" { return "TransportError" }
    if yname == "protocol-state-timestamp" { return "ProtocolStateTimestamp" }
    if yname == "transport-error-timestamp" { return "TransportErrorTimestamp" }
    if yname == "channel" { return "Channel" }
    return ""
}

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetSegmentPath() string {
    return "sdacp-control" + "[satellite-id='" + fmt.Sprintf("%v", sdacpControl.SatelliteId) + "']"
}

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol-state-timestamp" {
        return &sdacpControl.ProtocolStateTimestamp
    }
    if childYangName == "transport-error-timestamp" {
        return &sdacpControl.TransportErrorTimestamp
    }
    if childYangName == "channel" {
        for _, c := range sdacpControl.Channel {
            if sdacpControl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpControls_SdacpControl_Channel{}
        sdacpControl.Channel = append(sdacpControl.Channel, child)
        return &sdacpControl.Channel[len(sdacpControl.Channel)-1]
    }
    return nil
}

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocol-state-timestamp"] = &sdacpControl.ProtocolStateTimestamp
    children["transport-error-timestamp"] = &sdacpControl.TransportErrorTimestamp
    for i := range sdacpControl.Channel {
        children[sdacpControl.Channel[i].GetSegmentPath()] = &sdacpControl.Channel[i]
    }
    return children
}

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["satellite-id"] = sdacpControl.SatelliteId
    leafs["satellite-id-xr"] = sdacpControl.SatelliteIdXr
    leafs["satellite-ip-address"] = sdacpControl.SatelliteIpAddress
    leafs["ip-address-auto"] = sdacpControl.IpAddressAuto
    leafs["vrf-name"] = sdacpControl.VrfName
    leafs["control-protocol-state"] = sdacpControl.ControlProtocolState
    leafs["transport-error"] = sdacpControl.TransportError
    return leafs
}

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetBundleName() string { return "cisco_ios_xr" }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetYangName() string { return "sdacp-control" }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) SetParent(parent types.Entity) { sdacpControl.parent = parent }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetParent() types.Entity { return sdacpControl.parent }

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetParentYangName() string { return "sdacp-controls" }

// NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp
// Timestamp
type NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetFilter() yfilter.YFilter { return protocolStateTimestamp.YFilter }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) SetFilter(yf yfilter.YFilter) { protocolStateTimestamp.YFilter = yf }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetSegmentPath() string {
    return "protocol-state-timestamp"
}

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = protocolStateTimestamp.Seconds
    leafs["nanoseconds"] = protocolStateTimestamp.Nanoseconds
    return leafs
}

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetBundleName() string { return "cisco_ios_xr" }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetYangName() string { return "protocol-state-timestamp" }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) SetParent(parent types.Entity) { protocolStateTimestamp.parent = parent }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetParent() types.Entity { return protocolStateTimestamp.parent }

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetParentYangName() string { return "sdacp-control" }

// NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp
// Timestamp
type NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetFilter() yfilter.YFilter { return transportErrorTimestamp.YFilter }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) SetFilter(yf yfilter.YFilter) { transportErrorTimestamp.YFilter = yf }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetSegmentPath() string {
    return "transport-error-timestamp"
}

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = transportErrorTimestamp.Seconds
    leafs["nanoseconds"] = transportErrorTimestamp.Nanoseconds
    return leafs
}

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetBundleName() string { return "cisco_ios_xr" }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetYangName() string { return "transport-error-timestamp" }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) SetParent(parent types.Entity) { transportErrorTimestamp.parent = parent }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetParent() types.Entity { return transportErrorTimestamp.parent }

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetParentYangName() string { return "sdacp-control" }

// NvSatellite_SdacpControls_SdacpControl_Channel
// Channels on satellite table
type NvSatellite_SdacpControls_SdacpControl_Channel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Channel ID. The type is interface{} with range: 0..65535.
    ChannelId interface{}

    // Resync state. The type is IcpeCpmChannelResyncState.
    ResyncState interface{}

    // Channel state. The type is IcpeCpmChanFsmState.
    ChannelState interface{}

    // Control messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ControlMessagesSent interface{}

    // Normal messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    NormalMessagesSent interface{}

    // Control messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    ControlMessagesReceived interface{}

    // Normal messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    NormalMessagesReceived interface{}

    // Control messages dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    ControlMessagesDropped interface{}

    // Normal messages dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    NormalMessagesDropped interface{}

    // Secs since last cleared. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecsSinceLastCleared interface{}

    // Version. The type is interface{} with range: 0..65535.
    Version interface{}

    // Capabilities.
    Capabilities NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities

    // Timestamp.
    ResyncStateTimestamp NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp

    // Timestamp.
    ChannelStateTimestamp NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp
}

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetFilter() yfilter.YFilter { return channel.YFilter }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) SetFilter(yf yfilter.YFilter) { channel.YFilter = yf }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetGoName(yname string) string {
    if yname == "channel-id" { return "ChannelId" }
    if yname == "resync-state" { return "ResyncState" }
    if yname == "channel-state" { return "ChannelState" }
    if yname == "control-messages-sent" { return "ControlMessagesSent" }
    if yname == "normal-messages-sent" { return "NormalMessagesSent" }
    if yname == "control-messages-received" { return "ControlMessagesReceived" }
    if yname == "normal-messages-received" { return "NormalMessagesReceived" }
    if yname == "control-messages-dropped" { return "ControlMessagesDropped" }
    if yname == "normal-messages-dropped" { return "NormalMessagesDropped" }
    if yname == "secs-since-last-cleared" { return "SecsSinceLastCleared" }
    if yname == "version" { return "Version" }
    if yname == "capabilities" { return "Capabilities" }
    if yname == "resync-state-timestamp" { return "ResyncStateTimestamp" }
    if yname == "channel-state-timestamp" { return "ChannelStateTimestamp" }
    return ""
}

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetSegmentPath() string {
    return "channel"
}

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "capabilities" {
        return &channel.Capabilities
    }
    if childYangName == "resync-state-timestamp" {
        return &channel.ResyncStateTimestamp
    }
    if childYangName == "channel-state-timestamp" {
        return &channel.ChannelStateTimestamp
    }
    return nil
}

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["capabilities"] = &channel.Capabilities
    children["resync-state-timestamp"] = &channel.ResyncStateTimestamp
    children["channel-state-timestamp"] = &channel.ChannelStateTimestamp
    return children
}

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["channel-id"] = channel.ChannelId
    leafs["resync-state"] = channel.ResyncState
    leafs["channel-state"] = channel.ChannelState
    leafs["control-messages-sent"] = channel.ControlMessagesSent
    leafs["normal-messages-sent"] = channel.NormalMessagesSent
    leafs["control-messages-received"] = channel.ControlMessagesReceived
    leafs["normal-messages-received"] = channel.NormalMessagesReceived
    leafs["control-messages-dropped"] = channel.ControlMessagesDropped
    leafs["normal-messages-dropped"] = channel.NormalMessagesDropped
    leafs["secs-since-last-cleared"] = channel.SecsSinceLastCleared
    leafs["version"] = channel.Version
    return leafs
}

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetBundleName() string { return "cisco_ios_xr" }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetYangName() string { return "channel" }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) SetParent(parent types.Entity) { channel.parent = parent }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetParent() types.Entity { return channel.parent }

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetParentYangName() string { return "sdacp-control" }

// NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities
// Capabilities
type NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Capability TLVs table. The type is slice of
    // NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs.
    TlVs []NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs
}

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetFilter() yfilter.YFilter { return capabilities.YFilter }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) SetFilter(yf yfilter.YFilter) { capabilities.YFilter = yf }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetGoName(yname string) string {
    if yname == "tl-vs" { return "TlVs" }
    return ""
}

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetSegmentPath() string {
    return "capabilities"
}

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tl-vs" {
        for _, c := range capabilities.TlVs {
            if capabilities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs{}
        capabilities.TlVs = append(capabilities.TlVs, child)
        return &capabilities.TlVs[len(capabilities.TlVs)-1]
    }
    return nil
}

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range capabilities.TlVs {
        children[capabilities.TlVs[i].GetSegmentPath()] = &capabilities.TlVs[i]
    }
    return children
}

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetBundleName() string { return "cisco_ios_xr" }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetYangName() string { return "capabilities" }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) SetParent(parent types.Entity) { capabilities.parent = parent }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetParent() types.Entity { return capabilities.parent }

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetParentYangName() string { return "channel" }

// NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs
// Capability TLVs table
type NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is interface{} with range: 0..4294967295.
    Type interface{}

    // Debug. The type is string.
    Debug interface{}

    // Value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetFilter() yfilter.YFilter { return tlVs.YFilter }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) SetFilter(yf yfilter.YFilter) { tlVs.YFilter = yf }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "debug" { return "Debug" }
    if yname == "value" { return "Value" }
    return ""
}

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetSegmentPath() string {
    return "tl-vs"
}

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = tlVs.Type
    leafs["debug"] = tlVs.Debug
    leafs["value"] = tlVs.Value
    return leafs
}

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetBundleName() string { return "cisco_ios_xr" }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetYangName() string { return "tl-vs" }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) SetParent(parent types.Entity) { tlVs.parent = parent }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetParent() types.Entity { return tlVs.parent }

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetParentYangName() string { return "capabilities" }

// NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp
// Timestamp
type NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetFilter() yfilter.YFilter { return resyncStateTimestamp.YFilter }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) SetFilter(yf yfilter.YFilter) { resyncStateTimestamp.YFilter = yf }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetSegmentPath() string {
    return "resync-state-timestamp"
}

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = resyncStateTimestamp.Seconds
    leafs["nanoseconds"] = resyncStateTimestamp.Nanoseconds
    return leafs
}

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetBundleName() string { return "cisco_ios_xr" }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetYangName() string { return "resync-state-timestamp" }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) SetParent(parent types.Entity) { resyncStateTimestamp.parent = parent }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetParent() types.Entity { return resyncStateTimestamp.parent }

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetParentYangName() string { return "channel" }

// NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp
// Timestamp
type NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetFilter() yfilter.YFilter { return channelStateTimestamp.YFilter }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) SetFilter(yf yfilter.YFilter) { channelStateTimestamp.YFilter = yf }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetSegmentPath() string {
    return "channel-state-timestamp"
}

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = channelStateTimestamp.Seconds
    leafs["nanoseconds"] = channelStateTimestamp.Nanoseconds
    return leafs
}

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetBundleName() string { return "cisco_ios_xr" }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetYangName() string { return "channel-state-timestamp" }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) SetParent(parent types.Entity) { channelStateTimestamp.parent = parent }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetParent() types.Entity { return channelStateTimestamp.parent }

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetParentYangName() string { return "channel" }

