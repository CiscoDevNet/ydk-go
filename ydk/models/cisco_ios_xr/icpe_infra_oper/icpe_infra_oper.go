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

// NvSatellite
// Satellite operational information
type NvSatellite struct {
    EntityData types.CommonEntityData
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

func (nvSatellite *NvSatellite) GetEntityData() *types.CommonEntityData {
    nvSatellite.EntityData.YFilter = nvSatellite.YFilter
    nvSatellite.EntityData.YangName = "nv-satellite"
    nvSatellite.EntityData.BundleName = "cisco_ios_xr"
    nvSatellite.EntityData.ParentYangName = "Cisco-IOS-XR-icpe-infra-oper"
    nvSatellite.EntityData.SegmentPath = "Cisco-IOS-XR-icpe-infra-oper:nv-satellite"
    nvSatellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nvSatellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nvSatellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nvSatellite.EntityData.Children = make(map[string]types.YChild)
    nvSatellite.EntityData.Children["reload-op-statuses"] = types.YChild{"ReloadOpStatuses", &nvSatellite.ReloadOpStatuses}
    nvSatellite.EntityData.Children["sdacp-redundancies"] = types.YChild{"SdacpRedundancies", &nvSatellite.SdacpRedundancies}
    nvSatellite.EntityData.Children["install-shows"] = types.YChild{"InstallShows", &nvSatellite.InstallShows}
    nvSatellite.EntityData.Children["satellite-statuses"] = types.YChild{"SatelliteStatuses", &nvSatellite.SatelliteStatuses}
    nvSatellite.EntityData.Children["satellite-priorities"] = types.YChild{"SatellitePriorities", &nvSatellite.SatellitePriorities}
    nvSatellite.EntityData.Children["satellite-versions"] = types.YChild{"SatelliteVersions", &nvSatellite.SatelliteVersions}
    nvSatellite.EntityData.Children["satellite-topologies"] = types.YChild{"SatelliteTopologies", &nvSatellite.SatelliteTopologies}
    nvSatellite.EntityData.Children["install-reference-info"] = types.YChild{"InstallReferenceInfo", &nvSatellite.InstallReferenceInfo}
    nvSatellite.EntityData.Children["install-op-progresses"] = types.YChild{"InstallOpProgresses", &nvSatellite.InstallOpProgresses}
    nvSatellite.EntityData.Children["reload-statuses"] = types.YChild{"ReloadStatuses", &nvSatellite.ReloadStatuses}
    nvSatellite.EntityData.Children["install"] = types.YChild{"Install", &nvSatellite.Install}
    nvSatellite.EntityData.Children["install-op-statuses"] = types.YChild{"InstallOpStatuses", &nvSatellite.InstallOpStatuses}
    nvSatellite.EntityData.Children["satellite-properties"] = types.YChild{"SatelliteProperties", &nvSatellite.SatelliteProperties}
    nvSatellite.EntityData.Children["Cisco-IOS-XR-icpe-sdacp-oper:sdacp-discovery2s"] = types.YChild{"SdacpDiscovery2S", &nvSatellite.SdacpDiscovery2S}
    nvSatellite.EntityData.Children["Cisco-IOS-XR-icpe-sdacp-oper:icpe-dpms"] = types.YChild{"IcpeDpms", &nvSatellite.IcpeDpms}
    nvSatellite.EntityData.Children["Cisco-IOS-XR-icpe-sdacp-oper:sdacp-controls"] = types.YChild{"SdacpControls", &nvSatellite.SdacpControls}
    nvSatellite.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nvSatellite.EntityData)
}

// NvSatellite_ReloadOpStatuses
// Detailed breakdown of reload status table
type NvSatellite_ReloadOpStatuses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed breakdown of reload status. The type is slice of
    // NvSatellite_ReloadOpStatuses_ReloadOpStatus.
    ReloadOpStatus []NvSatellite_ReloadOpStatuses_ReloadOpStatus
}

func (reloadOpStatuses *NvSatellite_ReloadOpStatuses) GetEntityData() *types.CommonEntityData {
    reloadOpStatuses.EntityData.YFilter = reloadOpStatuses.YFilter
    reloadOpStatuses.EntityData.YangName = "reload-op-statuses"
    reloadOpStatuses.EntityData.BundleName = "cisco_ios_xr"
    reloadOpStatuses.EntityData.ParentYangName = "nv-satellite"
    reloadOpStatuses.EntityData.SegmentPath = "reload-op-statuses"
    reloadOpStatuses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reloadOpStatuses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reloadOpStatuses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reloadOpStatuses.EntityData.Children = make(map[string]types.YChild)
    reloadOpStatuses.EntityData.Children["reload-op-status"] = types.YChild{"ReloadOpStatus", nil}
    for i := range reloadOpStatuses.ReloadOpStatus {
        reloadOpStatuses.EntityData.Children[types.GetSegmentPath(&reloadOpStatuses.ReloadOpStatus[i])] = types.YChild{"ReloadOpStatus", &reloadOpStatuses.ReloadOpStatus[i]}
    }
    reloadOpStatuses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reloadOpStatuses.EntityData)
}

// NvSatellite_ReloadOpStatuses_ReloadOpStatus
// Detailed breakdown of reload status
type NvSatellite_ReloadOpStatuses_ReloadOpStatus struct {
    EntityData types.CommonEntityData
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

func (reloadOpStatus *NvSatellite_ReloadOpStatuses_ReloadOpStatus) GetEntityData() *types.CommonEntityData {
    reloadOpStatus.EntityData.YFilter = reloadOpStatus.YFilter
    reloadOpStatus.EntityData.YangName = "reload-op-status"
    reloadOpStatus.EntityData.BundleName = "cisco_ios_xr"
    reloadOpStatus.EntityData.ParentYangName = "reload-op-statuses"
    reloadOpStatus.EntityData.SegmentPath = "reload-op-status" + "[operation-id='" + fmt.Sprintf("%v", reloadOpStatus.OperationId) + "']"
    reloadOpStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reloadOpStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reloadOpStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reloadOpStatus.EntityData.Children = make(map[string]types.YChild)
    reloadOpStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    reloadOpStatus.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", reloadOpStatus.OperationId}
    reloadOpStatus.EntityData.Leafs["operation-id-xr"] = types.YLeaf{"OperationIdXr", reloadOpStatus.OperationIdXr}
    reloadOpStatus.EntityData.Leafs["satellite-range"] = types.YLeaf{"SatelliteRange", reloadOpStatus.SatelliteRange}
    reloadOpStatus.EntityData.Leafs["sats-not-initiated"] = types.YLeaf{"SatsNotInitiated", reloadOpStatus.SatsNotInitiated}
    reloadOpStatus.EntityData.Leafs["sats-reloading"] = types.YLeaf{"SatsReloading", reloadOpStatus.SatsReloading}
    reloadOpStatus.EntityData.Leafs["sats-reloaded"] = types.YLeaf{"SatsReloaded", reloadOpStatus.SatsReloaded}
    reloadOpStatus.EntityData.Leafs["sats-reload-failed"] = types.YLeaf{"SatsReloadFailed", reloadOpStatus.SatsReloadFailed}
    return &(reloadOpStatus.EntityData)
}

// NvSatellite_SdacpRedundancies
// nV Satellite Redundancy Protocol Information
// table
type NvSatellite_SdacpRedundancies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nV Satellite Redundancy Protocol Information. The type is slice of
    // NvSatellite_SdacpRedundancies_SdacpRedundancy.
    SdacpRedundancy []NvSatellite_SdacpRedundancies_SdacpRedundancy
}

func (sdacpRedundancies *NvSatellite_SdacpRedundancies) GetEntityData() *types.CommonEntityData {
    sdacpRedundancies.EntityData.YFilter = sdacpRedundancies.YFilter
    sdacpRedundancies.EntityData.YangName = "sdacp-redundancies"
    sdacpRedundancies.EntityData.BundleName = "cisco_ios_xr"
    sdacpRedundancies.EntityData.ParentYangName = "nv-satellite"
    sdacpRedundancies.EntityData.SegmentPath = "sdacp-redundancies"
    sdacpRedundancies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdacpRedundancies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdacpRedundancies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdacpRedundancies.EntityData.Children = make(map[string]types.YChild)
    sdacpRedundancies.EntityData.Children["sdacp-redundancy"] = types.YChild{"SdacpRedundancy", nil}
    for i := range sdacpRedundancies.SdacpRedundancy {
        sdacpRedundancies.EntityData.Children[types.GetSegmentPath(&sdacpRedundancies.SdacpRedundancy[i])] = types.YChild{"SdacpRedundancy", &sdacpRedundancies.SdacpRedundancy[i]}
    }
    sdacpRedundancies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sdacpRedundancies.EntityData)
}

// NvSatellite_SdacpRedundancies_SdacpRedundancy
// nV Satellite Redundancy Protocol Information
type NvSatellite_SdacpRedundancies_SdacpRedundancy struct {
    EntityData types.CommonEntityData
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SystemMac interface{}

    // Isolated. The type is bool.
    Isolated interface{}

    // Timestamp.
    ProtocolStateTimestamp NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp

    // Channels on this session table. The type is slice of
    // NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel.
    Channel []NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel
}

func (sdacpRedundancy *NvSatellite_SdacpRedundancies_SdacpRedundancy) GetEntityData() *types.CommonEntityData {
    sdacpRedundancy.EntityData.YFilter = sdacpRedundancy.YFilter
    sdacpRedundancy.EntityData.YangName = "sdacp-redundancy"
    sdacpRedundancy.EntityData.BundleName = "cisco_ios_xr"
    sdacpRedundancy.EntityData.ParentYangName = "sdacp-redundancies"
    sdacpRedundancy.EntityData.SegmentPath = "sdacp-redundancy" + "[iccp-group='" + fmt.Sprintf("%v", sdacpRedundancy.IccpGroup) + "']"
    sdacpRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdacpRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdacpRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdacpRedundancy.EntityData.Children = make(map[string]types.YChild)
    sdacpRedundancy.EntityData.Children["protocol-state-timestamp"] = types.YChild{"ProtocolStateTimestamp", &sdacpRedundancy.ProtocolStateTimestamp}
    sdacpRedundancy.EntityData.Children["channel"] = types.YChild{"Channel", nil}
    for i := range sdacpRedundancy.Channel {
        sdacpRedundancy.EntityData.Children[types.GetSegmentPath(&sdacpRedundancy.Channel[i])] = types.YChild{"Channel", &sdacpRedundancy.Channel[i]}
    }
    sdacpRedundancy.EntityData.Leafs = make(map[string]types.YLeaf)
    sdacpRedundancy.EntityData.Leafs["iccp-group"] = types.YLeaf{"IccpGroup", sdacpRedundancy.IccpGroup}
    sdacpRedundancy.EntityData.Leafs["iccp-group-xr"] = types.YLeaf{"IccpGroupXr", sdacpRedundancy.IccpGroupXr}
    sdacpRedundancy.EntityData.Leafs["protocol-state"] = types.YLeaf{"ProtocolState", sdacpRedundancy.ProtocolState}
    sdacpRedundancy.EntityData.Leafs["transport-state"] = types.YLeaf{"TransportState", sdacpRedundancy.TransportState}
    sdacpRedundancy.EntityData.Leafs["authentication-state"] = types.YLeaf{"AuthenticationState", sdacpRedundancy.AuthenticationState}
    sdacpRedundancy.EntityData.Leafs["arbitration-state"] = types.YLeaf{"ArbitrationState", sdacpRedundancy.ArbitrationState}
    sdacpRedundancy.EntityData.Leafs["synchronization-state"] = types.YLeaf{"SynchronizationState", sdacpRedundancy.SynchronizationState}
    sdacpRedundancy.EntityData.Leafs["primacy"] = types.YLeaf{"Primacy", sdacpRedundancy.Primacy}
    sdacpRedundancy.EntityData.Leafs["system-mac"] = types.YLeaf{"SystemMac", sdacpRedundancy.SystemMac}
    sdacpRedundancy.EntityData.Leafs["isolated"] = types.YLeaf{"Isolated", sdacpRedundancy.Isolated}
    return &(sdacpRedundancy.EntityData)
}

// NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp
// Timestamp
type NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (protocolStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_ProtocolStateTimestamp) GetEntityData() *types.CommonEntityData {
    protocolStateTimestamp.EntityData.YFilter = protocolStateTimestamp.YFilter
    protocolStateTimestamp.EntityData.YangName = "protocol-state-timestamp"
    protocolStateTimestamp.EntityData.BundleName = "cisco_ios_xr"
    protocolStateTimestamp.EntityData.ParentYangName = "sdacp-redundancy"
    protocolStateTimestamp.EntityData.SegmentPath = "protocol-state-timestamp"
    protocolStateTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolStateTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolStateTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolStateTimestamp.EntityData.Children = make(map[string]types.YChild)
    protocolStateTimestamp.EntityData.Leafs = make(map[string]types.YLeaf)
    protocolStateTimestamp.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", protocolStateTimestamp.Seconds}
    protocolStateTimestamp.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", protocolStateTimestamp.Nanoseconds}
    return &(protocolStateTimestamp.EntityData)
}

// NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel
// Channels on this session table
type NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel struct {
    EntityData types.CommonEntityData
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

func (channel *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel) GetEntityData() *types.CommonEntityData {
    channel.EntityData.YFilter = channel.YFilter
    channel.EntityData.YangName = "channel"
    channel.EntityData.BundleName = "cisco_ios_xr"
    channel.EntityData.ParentYangName = "sdacp-redundancy"
    channel.EntityData.SegmentPath = "channel"
    channel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channel.EntityData.Children = make(map[string]types.YChild)
    channel.EntityData.Children["channel-state-timestamp"] = types.YChild{"ChannelStateTimestamp", &channel.ChannelStateTimestamp}
    channel.EntityData.Children["resync-state-timestamp"] = types.YChild{"ResyncStateTimestamp", &channel.ResyncStateTimestamp}
    channel.EntityData.Leafs = make(map[string]types.YLeaf)
    channel.EntityData.Leafs["channel-id"] = types.YLeaf{"ChannelId", channel.ChannelId}
    channel.EntityData.Leafs["chan-state"] = types.YLeaf{"ChanState", channel.ChanState}
    channel.EntityData.Leafs["resync-state"] = types.YLeaf{"ResyncState", channel.ResyncState}
    channel.EntityData.Leafs["control-messages-sent"] = types.YLeaf{"ControlMessagesSent", channel.ControlMessagesSent}
    channel.EntityData.Leafs["normal-messages-sent"] = types.YLeaf{"NormalMessagesSent", channel.NormalMessagesSent}
    channel.EntityData.Leafs["control-messages-received"] = types.YLeaf{"ControlMessagesReceived", channel.ControlMessagesReceived}
    channel.EntityData.Leafs["normal-messages-received"] = types.YLeaf{"NormalMessagesReceived", channel.NormalMessagesReceived}
    return &(channel.EntityData)
}

// NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp
// Timestamp
type NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (channelStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ChannelStateTimestamp) GetEntityData() *types.CommonEntityData {
    channelStateTimestamp.EntityData.YFilter = channelStateTimestamp.YFilter
    channelStateTimestamp.EntityData.YangName = "channel-state-timestamp"
    channelStateTimestamp.EntityData.BundleName = "cisco_ios_xr"
    channelStateTimestamp.EntityData.ParentYangName = "channel"
    channelStateTimestamp.EntityData.SegmentPath = "channel-state-timestamp"
    channelStateTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelStateTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelStateTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelStateTimestamp.EntityData.Children = make(map[string]types.YChild)
    channelStateTimestamp.EntityData.Leafs = make(map[string]types.YLeaf)
    channelStateTimestamp.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", channelStateTimestamp.Seconds}
    channelStateTimestamp.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", channelStateTimestamp.Nanoseconds}
    return &(channelStateTimestamp.EntityData)
}

// NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp
// Timestamp
type NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resyncStateTimestamp *NvSatellite_SdacpRedundancies_SdacpRedundancy_Channel_ResyncStateTimestamp) GetEntityData() *types.CommonEntityData {
    resyncStateTimestamp.EntityData.YFilter = resyncStateTimestamp.YFilter
    resyncStateTimestamp.EntityData.YangName = "resync-state-timestamp"
    resyncStateTimestamp.EntityData.BundleName = "cisco_ios_xr"
    resyncStateTimestamp.EntityData.ParentYangName = "channel"
    resyncStateTimestamp.EntityData.SegmentPath = "resync-state-timestamp"
    resyncStateTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resyncStateTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resyncStateTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resyncStateTimestamp.EntityData.Children = make(map[string]types.YChild)
    resyncStateTimestamp.EntityData.Leafs = make(map[string]types.YLeaf)
    resyncStateTimestamp.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", resyncStateTimestamp.Seconds}
    resyncStateTimestamp.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", resyncStateTimestamp.Nanoseconds}
    return &(resyncStateTimestamp.EntityData)
}

// NvSatellite_InstallShows
// Detailed breakdown of install status table
type NvSatellite_InstallShows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed breakdown of install status. The type is slice of
    // NvSatellite_InstallShows_InstallShow.
    InstallShow []NvSatellite_InstallShows_InstallShow
}

func (installShows *NvSatellite_InstallShows) GetEntityData() *types.CommonEntityData {
    installShows.EntityData.YFilter = installShows.YFilter
    installShows.EntityData.YangName = "install-shows"
    installShows.EntityData.BundleName = "cisco_ios_xr"
    installShows.EntityData.ParentYangName = "nv-satellite"
    installShows.EntityData.SegmentPath = "install-shows"
    installShows.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installShows.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installShows.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installShows.EntityData.Children = make(map[string]types.YChild)
    installShows.EntityData.Children["install-show"] = types.YChild{"InstallShow", nil}
    for i := range installShows.InstallShow {
        installShows.EntityData.Children[types.GetSegmentPath(&installShows.InstallShow[i])] = types.YChild{"InstallShow", &installShows.InstallShow[i]}
    }
    installShows.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(installShows.EntityData)
}

// NvSatellite_InstallShows_InstallShow
// Detailed breakdown of install status
type NvSatellite_InstallShows_InstallShow struct {
    EntityData types.CommonEntityData
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

func (installShow *NvSatellite_InstallShows_InstallShow) GetEntityData() *types.CommonEntityData {
    installShow.EntityData.YFilter = installShow.YFilter
    installShow.EntityData.YangName = "install-show"
    installShow.EntityData.BundleName = "cisco_ios_xr"
    installShow.EntityData.ParentYangName = "install-shows"
    installShow.EntityData.SegmentPath = "install-show" + "[operation-id='" + fmt.Sprintf("%v", installShow.OperationId) + "']"
    installShow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installShow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installShow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installShow.EntityData.Children = make(map[string]types.YChild)
    installShow.EntityData.Children["satellite"] = types.YChild{"Satellite", nil}
    for i := range installShow.Satellite {
        installShow.EntityData.Children[types.GetSegmentPath(&installShow.Satellite[i])] = types.YChild{"Satellite", &installShow.Satellite[i]}
    }
    installShow.EntityData.Leafs = make(map[string]types.YLeaf)
    installShow.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", installShow.OperationId}
    installShow.EntityData.Leafs["operation-id-xr"] = types.YLeaf{"OperationIdXr", installShow.OperationIdXr}
    installShow.EntityData.Leafs["satellite-range"] = types.YLeaf{"SatelliteRange", installShow.SatelliteRange}
    installShow.EntityData.Leafs["operation-type"] = types.YLeaf{"OperationType", installShow.OperationType}
    installShow.EntityData.Leafs["progress-percentage"] = types.YLeaf{"ProgressPercentage", installShow.ProgressPercentage}
    installShow.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", installShow.StartTime}
    installShow.EntityData.Leafs["end-time"] = types.YLeaf{"EndTime", installShow.EndTime}
    installShow.EntityData.Leafs["ref-state"] = types.YLeaf{"RefState", installShow.RefState}
    installShow.EntityData.Leafs["sats-not-initiated"] = types.YLeaf{"SatsNotInitiated", installShow.SatsNotInitiated}
    installShow.EntityData.Leafs["sats-transferring"] = types.YLeaf{"SatsTransferring", installShow.SatsTransferring}
    installShow.EntityData.Leafs["sats-activating"] = types.YLeaf{"SatsActivating", installShow.SatsActivating}
    installShow.EntityData.Leafs["sats-updating"] = types.YLeaf{"SatsUpdating", installShow.SatsUpdating}
    installShow.EntityData.Leafs["sats-deactivating"] = types.YLeaf{"SatsDeactivating", installShow.SatsDeactivating}
    installShow.EntityData.Leafs["sats-removing"] = types.YLeaf{"SatsRemoving", installShow.SatsRemoving}
    installShow.EntityData.Leafs["sats-transfer-failed"] = types.YLeaf{"SatsTransferFailed", installShow.SatsTransferFailed}
    installShow.EntityData.Leafs["sats-activate-failed"] = types.YLeaf{"SatsActivateFailed", installShow.SatsActivateFailed}
    installShow.EntityData.Leafs["sats-update-failed"] = types.YLeaf{"SatsUpdateFailed", installShow.SatsUpdateFailed}
    installShow.EntityData.Leafs["sats-deactivate-failed"] = types.YLeaf{"SatsDeactivateFailed", installShow.SatsDeactivateFailed}
    installShow.EntityData.Leafs["sats-remove-failed"] = types.YLeaf{"SatsRemoveFailed", installShow.SatsRemoveFailed}
    installShow.EntityData.Leafs["sats-transfer-aborted"] = types.YLeaf{"SatsTransferAborted", installShow.SatsTransferAborted}
    installShow.EntityData.Leafs["sats-activate-aborted"] = types.YLeaf{"SatsActivateAborted", installShow.SatsActivateAborted}
    installShow.EntityData.Leafs["sats-update-aborted"] = types.YLeaf{"SatsUpdateAborted", installShow.SatsUpdateAborted}
    installShow.EntityData.Leafs["sats-deactivate-aborted"] = types.YLeaf{"SatsDeactivateAborted", installShow.SatsDeactivateAborted}
    installShow.EntityData.Leafs["sats-remove-aborted"] = types.YLeaf{"SatsRemoveAborted", installShow.SatsRemoveAborted}
    installShow.EntityData.Leafs["sats-no-operation"] = types.YLeaf{"SatsNoOperation", installShow.SatsNoOperation}
    installShow.EntityData.Leafs["sats-completed"] = types.YLeaf{"SatsCompleted", installShow.SatsCompleted}
    installShow.EntityData.Leafs["name-string"] = types.YLeaf{"NameString", installShow.NameString}
    return &(installShow.EntityData)
}

// NvSatellite_InstallShows_InstallShow_Satellite
// Breakdown per satellite table
type NvSatellite_InstallShows_InstallShow_Satellite struct {
    EntityData types.CommonEntityData
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

func (satellite *NvSatellite_InstallShows_InstallShow_Satellite) GetEntityData() *types.CommonEntityData {
    satellite.EntityData.YFilter = satellite.YFilter
    satellite.EntityData.YangName = "satellite"
    satellite.EntityData.BundleName = "cisco_ios_xr"
    satellite.EntityData.ParentYangName = "install-show"
    satellite.EntityData.SegmentPath = "satellite"
    satellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satellite.EntityData.Children = make(map[string]types.YChild)
    satellite.EntityData.Leafs = make(map[string]types.YLeaf)
    satellite.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", satellite.SatelliteId}
    satellite.EntityData.Leafs["state"] = types.YLeaf{"State", satellite.State}
    satellite.EntityData.Leafs["percentage"] = types.YLeaf{"Percentage", satellite.Percentage}
    satellite.EntityData.Leafs["retries"] = types.YLeaf{"Retries", satellite.Retries}
    satellite.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", satellite.StartTime}
    satellite.EntityData.Leafs["end-time"] = types.YLeaf{"EndTime", satellite.EndTime}
    satellite.EntityData.Leafs["info"] = types.YLeaf{"Info", satellite.Info}
    return &(satellite.EntityData)
}

// NvSatellite_SatelliteStatuses
// Satellite status information table
type NvSatellite_SatelliteStatuses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite status information. The type is slice of
    // NvSatellite_SatelliteStatuses_SatelliteStatus.
    SatelliteStatus []NvSatellite_SatelliteStatuses_SatelliteStatus
}

func (satelliteStatuses *NvSatellite_SatelliteStatuses) GetEntityData() *types.CommonEntityData {
    satelliteStatuses.EntityData.YFilter = satelliteStatuses.YFilter
    satelliteStatuses.EntityData.YangName = "satellite-statuses"
    satelliteStatuses.EntityData.BundleName = "cisco_ios_xr"
    satelliteStatuses.EntityData.ParentYangName = "nv-satellite"
    satelliteStatuses.EntityData.SegmentPath = "satellite-statuses"
    satelliteStatuses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteStatuses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteStatuses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteStatuses.EntityData.Children = make(map[string]types.YChild)
    satelliteStatuses.EntityData.Children["satellite-status"] = types.YChild{"SatelliteStatus", nil}
    for i := range satelliteStatuses.SatelliteStatus {
        satelliteStatuses.EntityData.Children[types.GetSegmentPath(&satelliteStatuses.SatelliteStatus[i])] = types.YChild{"SatelliteStatus", &satelliteStatuses.SatelliteStatus[i]}
    }
    satelliteStatuses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(satelliteStatuses.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus
// Satellite status information
type NvSatellite_SatelliteStatuses_SatelliteStatus struct {
    EntityData types.CommonEntityData
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
    Type_ interface{}

    // Ethernet fabric supported. The type is bool.
    EthernetFabricSupported interface{}

    // Optical supported. The type is bool.
    OpticalSupported interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // IP address present. The type is bool.
    IpAddressPresent interface{}

    // IP address auto. The type is bool.
    IpAddressAuto interface{}

    // IPV6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

func (satelliteStatus *NvSatellite_SatelliteStatuses_SatelliteStatus) GetEntityData() *types.CommonEntityData {
    satelliteStatus.EntityData.YFilter = satelliteStatus.YFilter
    satelliteStatus.EntityData.YangName = "satellite-status"
    satelliteStatus.EntityData.BundleName = "cisco_ios_xr"
    satelliteStatus.EntityData.ParentYangName = "satellite-statuses"
    satelliteStatus.EntityData.SegmentPath = "satellite-status" + "[satellite-id='" + fmt.Sprintf("%v", satelliteStatus.SatelliteId) + "']"
    satelliteStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteStatus.EntityData.Children = make(map[string]types.YChild)
    satelliteStatus.EntityData.Children["candidate-fabric-ports"] = types.YChild{"CandidateFabricPorts", &satelliteStatus.CandidateFabricPorts}
    satelliteStatus.EntityData.Children["optical-status"] = types.YChild{"OpticalStatus", &satelliteStatus.OpticalStatus}
    satelliteStatus.EntityData.Children["redundancy-out-of-sync-timestamp"] = types.YChild{"RedundancyOutOfSyncTimestamp", &satelliteStatus.RedundancyOutOfSyncTimestamp}
    satelliteStatus.EntityData.Children["configured-link"] = types.YChild{"ConfiguredLink", nil}
    for i := range satelliteStatus.ConfiguredLink {
        satelliteStatus.EntityData.Children[types.GetSegmentPath(&satelliteStatus.ConfiguredLink[i])] = types.YChild{"ConfiguredLink", &satelliteStatus.ConfiguredLink[i]}
    }
    satelliteStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    satelliteStatus.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", satelliteStatus.SatelliteId}
    satelliteStatus.EntityData.Leafs["satellite-id-xr"] = types.YLeaf{"SatelliteIdXr", satelliteStatus.SatelliteIdXr}
    satelliteStatus.EntityData.Leafs["version-check-state"] = types.YLeaf{"VersionCheckState", satelliteStatus.VersionCheckState}
    satelliteStatus.EntityData.Leafs["remote-version-present"] = types.YLeaf{"RemoteVersionPresent", satelliteStatus.RemoteVersionPresent}
    satelliteStatus.EntityData.Leafs["type"] = types.YLeaf{"Type_", satelliteStatus.Type_}
    satelliteStatus.EntityData.Leafs["ethernet-fabric-supported"] = types.YLeaf{"EthernetFabricSupported", satelliteStatus.EthernetFabricSupported}
    satelliteStatus.EntityData.Leafs["optical-supported"] = types.YLeaf{"OpticalSupported", satelliteStatus.OpticalSupported}
    satelliteStatus.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", satelliteStatus.IpAddress}
    satelliteStatus.EntityData.Leafs["ip-address-present"] = types.YLeaf{"IpAddressPresent", satelliteStatus.IpAddressPresent}
    satelliteStatus.EntityData.Leafs["ip-address-auto"] = types.YLeaf{"IpAddressAuto", satelliteStatus.IpAddressAuto}
    satelliteStatus.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", satelliteStatus.Ipv6Address}
    satelliteStatus.EntityData.Leafs["ipv6-address-present"] = types.YLeaf{"Ipv6AddressPresent", satelliteStatus.Ipv6AddressPresent}
    satelliteStatus.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", satelliteStatus.VrfName}
    satelliteStatus.EntityData.Leafs["vrfid"] = types.YLeaf{"Vrfid", satelliteStatus.Vrfid}
    satelliteStatus.EntityData.Leafs["description"] = types.YLeaf{"Description", satelliteStatus.Description}
    satelliteStatus.EntityData.Leafs["description-present"] = types.YLeaf{"DescriptionPresent", satelliteStatus.DescriptionPresent}
    satelliteStatus.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", satelliteStatus.MacAddress}
    satelliteStatus.EntityData.Leafs["mac-address-present"] = types.YLeaf{"MacAddressPresent", satelliteStatus.MacAddressPresent}
    satelliteStatus.EntityData.Leafs["configured-serial-number"] = types.YLeaf{"ConfiguredSerialNumber", satelliteStatus.ConfiguredSerialNumber}
    satelliteStatus.EntityData.Leafs["configured-serial-number-present"] = types.YLeaf{"ConfiguredSerialNumberPresent", satelliteStatus.ConfiguredSerialNumberPresent}
    satelliteStatus.EntityData.Leafs["received-serial-number"] = types.YLeaf{"ReceivedSerialNumber", satelliteStatus.ReceivedSerialNumber}
    satelliteStatus.EntityData.Leafs["received-serial-number-present"] = types.YLeaf{"ReceivedSerialNumberPresent", satelliteStatus.ReceivedSerialNumberPresent}
    satelliteStatus.EntityData.Leafs["password"] = types.YLeaf{"Password", satelliteStatus.Password}
    satelliteStatus.EntityData.Leafs["password-error"] = types.YLeaf{"PasswordError", satelliteStatus.PasswordError}
    satelliteStatus.EntityData.Leafs["received-host-name"] = types.YLeaf{"ReceivedHostName", satelliteStatus.ReceivedHostName}
    satelliteStatus.EntityData.Leafs["cfgd-timeout"] = types.YLeaf{"CfgdTimeout", satelliteStatus.CfgdTimeout}
    satelliteStatus.EntityData.Leafs["timeout-warning"] = types.YLeaf{"TimeoutWarning", satelliteStatus.TimeoutWarning}
    satelliteStatus.EntityData.Leafs["conflict-reason"] = types.YLeaf{"ConflictReason", satelliteStatus.ConflictReason}
    satelliteStatus.EntityData.Leafs["conflict-context"] = types.YLeaf{"ConflictContext", satelliteStatus.ConflictContext}
    satelliteStatus.EntityData.Leafs["redundancy-iccp-group"] = types.YLeaf{"RedundancyIccpGroup", satelliteStatus.RedundancyIccpGroup}
    satelliteStatus.EntityData.Leafs["recovery-delay-time-left"] = types.YLeaf{"RecoveryDelayTimeLeft", satelliteStatus.RecoveryDelayTimeLeft}
    satelliteStatus.EntityData.Leafs["host-treating-as-active"] = types.YLeaf{"HostTreatingAsActive", satelliteStatus.HostTreatingAsActive}
    satelliteStatus.EntityData.Leafs["satellite-treating-as-active"] = types.YLeaf{"SatelliteTreatingAsActive", satelliteStatus.SatelliteTreatingAsActive}
    satelliteStatus.EntityData.Leafs["sdacp-session-state"] = types.YLeaf{"SdacpSessionState", satelliteStatus.SdacpSessionState}
    satelliteStatus.EntityData.Leafs["sdacp-session-failure-reason"] = types.YLeaf{"SdacpSessionFailureReason", satelliteStatus.SdacpSessionFailureReason}
    satelliteStatus.EntityData.Leafs["install-state"] = types.YLeaf{"InstallState", satelliteStatus.InstallState}
    satelliteStatus.EntityData.Leafs["remote-version"] = types.YLeaf{"RemoteVersion", satelliteStatus.RemoteVersion}
    return &(satelliteStatus.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts
// Candidate Fabric Ports on this Satellite
type NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts struct {
    EntityData types.CommonEntityData
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

func (candidateFabricPorts *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts) GetEntityData() *types.CommonEntityData {
    candidateFabricPorts.EntityData.YFilter = candidateFabricPorts.YFilter
    candidateFabricPorts.EntityData.YangName = "candidate-fabric-ports"
    candidateFabricPorts.EntityData.BundleName = "cisco_ios_xr"
    candidateFabricPorts.EntityData.ParentYangName = "satellite-status"
    candidateFabricPorts.EntityData.SegmentPath = "candidate-fabric-ports"
    candidateFabricPorts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateFabricPorts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateFabricPorts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateFabricPorts.EntityData.Children = make(map[string]types.YChild)
    candidateFabricPorts.EntityData.Children["configured-port"] = types.YChild{"ConfiguredPort", nil}
    for i := range candidateFabricPorts.ConfiguredPort {
        candidateFabricPorts.EntityData.Children[types.GetSegmentPath(&candidateFabricPorts.ConfiguredPort[i])] = types.YChild{"ConfiguredPort", &candidateFabricPorts.ConfiguredPort[i]}
    }
    candidateFabricPorts.EntityData.Children["current-port"] = types.YChild{"CurrentPort", nil}
    for i := range candidateFabricPorts.CurrentPort {
        candidateFabricPorts.EntityData.Children[types.GetSegmentPath(&candidateFabricPorts.CurrentPort[i])] = types.YChild{"CurrentPort", &candidateFabricPorts.CurrentPort[i]}
    }
    candidateFabricPorts.EntityData.Leafs = make(map[string]types.YLeaf)
    candidateFabricPorts.EntityData.Leafs["channel-up"] = types.YLeaf{"ChannelUp", candidateFabricPorts.ChannelUp}
    candidateFabricPorts.EntityData.Leafs["out-of-sync"] = types.YLeaf{"OutOfSync", candidateFabricPorts.OutOfSync}
    candidateFabricPorts.EntityData.Leafs["error-string"] = types.YLeaf{"ErrorString", candidateFabricPorts.ErrorString}
    return &(candidateFabricPorts.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort
// Configured Candidate Fabric Ports table
type NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort struct {
    EntityData types.CommonEntityData
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

func (configuredPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_ConfiguredPort) GetEntityData() *types.CommonEntityData {
    configuredPort.EntityData.YFilter = configuredPort.YFilter
    configuredPort.EntityData.YangName = "configured-port"
    configuredPort.EntityData.BundleName = "cisco_ios_xr"
    configuredPort.EntityData.ParentYangName = "candidate-fabric-ports"
    configuredPort.EntityData.SegmentPath = "configured-port"
    configuredPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredPort.EntityData.Children = make(map[string]types.YChild)
    configuredPort.EntityData.Leafs = make(map[string]types.YLeaf)
    configuredPort.EntityData.Leafs["port-type"] = types.YLeaf{"PortType", configuredPort.PortType}
    configuredPort.EntityData.Leafs["slot"] = types.YLeaf{"Slot", configuredPort.Slot}
    configuredPort.EntityData.Leafs["subslot"] = types.YLeaf{"Subslot", configuredPort.Subslot}
    configuredPort.EntityData.Leafs["port"] = types.YLeaf{"Port", configuredPort.Port}
    configuredPort.EntityData.Leafs["valid"] = types.YLeaf{"Valid", configuredPort.Valid}
    return &(configuredPort.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort
// Current Candidate Fabric Ports on this Satellite
// table
type NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort struct {
    EntityData types.CommonEntityData
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

func (currentPort *NvSatellite_SatelliteStatuses_SatelliteStatus_CandidateFabricPorts_CurrentPort) GetEntityData() *types.CommonEntityData {
    currentPort.EntityData.YFilter = currentPort.YFilter
    currentPort.EntityData.YangName = "current-port"
    currentPort.EntityData.BundleName = "cisco_ios_xr"
    currentPort.EntityData.ParentYangName = "candidate-fabric-ports"
    currentPort.EntityData.SegmentPath = "current-port"
    currentPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentPort.EntityData.Children = make(map[string]types.YChild)
    currentPort.EntityData.Leafs = make(map[string]types.YLeaf)
    currentPort.EntityData.Leafs["port-type"] = types.YLeaf{"PortType", currentPort.PortType}
    currentPort.EntityData.Leafs["slot"] = types.YLeaf{"Slot", currentPort.Slot}
    currentPort.EntityData.Leafs["subslot"] = types.YLeaf{"Subslot", currentPort.Subslot}
    currentPort.EntityData.Leafs["port"] = types.YLeaf{"Port", currentPort.Port}
    currentPort.EntityData.Leafs["permanent"] = types.YLeaf{"Permanent", currentPort.Permanent}
    currentPort.EntityData.Leafs["requested"] = types.YLeaf{"Requested", currentPort.Requested}
    return &(currentPort.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus
// Optical Satellite Status
type NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Chassis sync state. The type is IcpeOpticalSyncState.
    ChassisSyncState interface{}

    // Application State table. The type is slice of
    // NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application.
    Application []NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application
}

func (opticalStatus *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus) GetEntityData() *types.CommonEntityData {
    opticalStatus.EntityData.YFilter = opticalStatus.YFilter
    opticalStatus.EntityData.YangName = "optical-status"
    opticalStatus.EntityData.BundleName = "cisco_ios_xr"
    opticalStatus.EntityData.ParentYangName = "satellite-status"
    opticalStatus.EntityData.SegmentPath = "optical-status"
    opticalStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalStatus.EntityData.Children = make(map[string]types.YChild)
    opticalStatus.EntityData.Children["application"] = types.YChild{"Application", nil}
    for i := range opticalStatus.Application {
        opticalStatus.EntityData.Children[types.GetSegmentPath(&opticalStatus.Application[i])] = types.YChild{"Application", &opticalStatus.Application[i]}
    }
    opticalStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    opticalStatus.EntityData.Leafs["chassis-sync-state"] = types.YLeaf{"ChassisSyncState", opticalStatus.ChassisSyncState}
    return &(opticalStatus.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application
// Application State table
type NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is string.
    Name interface{}

    // Sync state. The type is IcpeOpticalSyncState.
    SyncState interface{}
}

func (application *NvSatellite_SatelliteStatuses_SatelliteStatus_OpticalStatus_Application) GetEntityData() *types.CommonEntityData {
    application.EntityData.YFilter = application.YFilter
    application.EntityData.YangName = "application"
    application.EntityData.BundleName = "cisco_ios_xr"
    application.EntityData.ParentYangName = "optical-status"
    application.EntityData.SegmentPath = "application"
    application.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    application.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    application.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    application.EntityData.Children = make(map[string]types.YChild)
    application.EntityData.Leafs = make(map[string]types.YLeaf)
    application.EntityData.Leafs["name"] = types.YLeaf{"Name", application.Name}
    application.EntityData.Leafs["sync-state"] = types.YLeaf{"SyncState", application.SyncState}
    return &(application.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp
// Timestamp
type NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (redundancyOutOfSyncTimestamp *NvSatellite_SatelliteStatuses_SatelliteStatus_RedundancyOutOfSyncTimestamp) GetEntityData() *types.CommonEntityData {
    redundancyOutOfSyncTimestamp.EntityData.YFilter = redundancyOutOfSyncTimestamp.YFilter
    redundancyOutOfSyncTimestamp.EntityData.YangName = "redundancy-out-of-sync-timestamp"
    redundancyOutOfSyncTimestamp.EntityData.BundleName = "cisco_ios_xr"
    redundancyOutOfSyncTimestamp.EntityData.ParentYangName = "satellite-status"
    redundancyOutOfSyncTimestamp.EntityData.SegmentPath = "redundancy-out-of-sync-timestamp"
    redundancyOutOfSyncTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancyOutOfSyncTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancyOutOfSyncTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancyOutOfSyncTimestamp.EntityData.Children = make(map[string]types.YChild)
    redundancyOutOfSyncTimestamp.EntityData.Leafs = make(map[string]types.YLeaf)
    redundancyOutOfSyncTimestamp.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", redundancyOutOfSyncTimestamp.Seconds}
    redundancyOutOfSyncTimestamp.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", redundancyOutOfSyncTimestamp.Nanoseconds}
    return &(redundancyOutOfSyncTimestamp.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink
// Configured Links on this Satellite table
type NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceHandle interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (configuredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink) GetEntityData() *types.CommonEntityData {
    configuredLink.EntityData.YFilter = configuredLink.YFilter
    configuredLink.EntityData.YangName = "configured-link"
    configuredLink.EntityData.BundleName = "cisco_ios_xr"
    configuredLink.EntityData.ParentYangName = "satellite-status"
    configuredLink.EntityData.SegmentPath = "configured-link"
    configuredLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredLink.EntityData.Children = make(map[string]types.YChild)
    configuredLink.EntityData.Children["port-range"] = types.YChild{"PortRange", nil}
    for i := range configuredLink.PortRange {
        configuredLink.EntityData.Children[types.GetSegmentPath(&configuredLink.PortRange[i])] = types.YChild{"PortRange", &configuredLink.PortRange[i]}
    }
    configuredLink.EntityData.Children["discovered-link"] = types.YChild{"DiscoveredLink", nil}
    for i := range configuredLink.DiscoveredLink {
        configuredLink.EntityData.Children[types.GetSegmentPath(&configuredLink.DiscoveredLink[i])] = types.YChild{"DiscoveredLink", &configuredLink.DiscoveredLink[i]}
    }
    configuredLink.EntityData.Leafs = make(map[string]types.YLeaf)
    configuredLink.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", configuredLink.InterfaceHandle}
    configuredLink.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", configuredLink.IpAddress}
    configuredLink.EntityData.Leafs["ip-address-auto"] = types.YLeaf{"IpAddressAuto", configuredLink.IpAddressAuto}
    configuredLink.EntityData.Leafs["vrf-id-present"] = types.YLeaf{"VrfIdPresent", configuredLink.VrfIdPresent}
    configuredLink.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", configuredLink.VrfId}
    configuredLink.EntityData.Leafs["minimum-preferred-links"] = types.YLeaf{"MinimumPreferredLinks", configuredLink.MinimumPreferredLinks}
    configuredLink.EntityData.Leafs["number-active-links"] = types.YLeaf{"NumberActiveLinks", configuredLink.NumberActiveLinks}
    configuredLink.EntityData.Leafs["min-links-satisfied"] = types.YLeaf{"MinLinksSatisfied", configuredLink.MinLinksSatisfied}
    configuredLink.EntityData.Leafs["minimum-required-links"] = types.YLeaf{"MinimumRequiredLinks", configuredLink.MinimumRequiredLinks}
    configuredLink.EntityData.Leafs["required-min-links-satisfied"] = types.YLeaf{"RequiredMinLinksSatisfied", configuredLink.RequiredMinLinksSatisfied}
    configuredLink.EntityData.Leafs["conflict-reason"] = types.YLeaf{"ConflictReason", configuredLink.ConflictReason}
    configuredLink.EntityData.Leafs["conflict-context"] = types.YLeaf{"ConflictContext", configuredLink.ConflictContext}
    return &(configuredLink.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange
// Port ranges table
type NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange struct {
    EntityData types.CommonEntityData
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

func (portRange *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_PortRange) GetEntityData() *types.CommonEntityData {
    portRange.EntityData.YFilter = portRange.YFilter
    portRange.EntityData.YangName = "port-range"
    portRange.EntityData.BundleName = "cisco_ios_xr"
    portRange.EntityData.ParentYangName = "configured-link"
    portRange.EntityData.SegmentPath = "port-range"
    portRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portRange.EntityData.Children = make(map[string]types.YChild)
    portRange.EntityData.Leafs = make(map[string]types.YLeaf)
    portRange.EntityData.Leafs["slot"] = types.YLeaf{"Slot", portRange.Slot}
    portRange.EntityData.Leafs["subslot"] = types.YLeaf{"Subslot", portRange.Subslot}
    portRange.EntityData.Leafs["low-port"] = types.YLeaf{"LowPort", portRange.LowPort}
    portRange.EntityData.Leafs["high-port"] = types.YLeaf{"HighPort", portRange.HighPort}
    portRange.EntityData.Leafs["port-type"] = types.YLeaf{"PortType", portRange.PortType}
    portRange.EntityData.Leafs["conflict-reason"] = types.YLeaf{"ConflictReason", portRange.ConflictReason}
    portRange.EntityData.Leafs["conflict-context"] = types.YLeaf{"ConflictContext", portRange.ConflictContext}
    return &(portRange.EntityData)
}

// NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink
// Discovered Links table
type NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceHandle interface{}

    // State. The type is IcpeOperDiscdLinkState.
    State interface{}

    // Conflict reason. The type is IcpeOperConflict.
    ConflictReason interface{}

    // Conflict context. The type is string.
    ConflictContext interface{}
}

func (discoveredLink *NvSatellite_SatelliteStatuses_SatelliteStatus_ConfiguredLink_DiscoveredLink) GetEntityData() *types.CommonEntityData {
    discoveredLink.EntityData.YFilter = discoveredLink.YFilter
    discoveredLink.EntityData.YangName = "discovered-link"
    discoveredLink.EntityData.BundleName = "cisco_ios_xr"
    discoveredLink.EntityData.ParentYangName = "configured-link"
    discoveredLink.EntityData.SegmentPath = "discovered-link"
    discoveredLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discoveredLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discoveredLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discoveredLink.EntityData.Children = make(map[string]types.YChild)
    discoveredLink.EntityData.Leafs = make(map[string]types.YLeaf)
    discoveredLink.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", discoveredLink.InterfaceHandle}
    discoveredLink.EntityData.Leafs["state"] = types.YLeaf{"State", discoveredLink.State}
    discoveredLink.EntityData.Leafs["conflict-reason"] = types.YLeaf{"ConflictReason", discoveredLink.ConflictReason}
    discoveredLink.EntityData.Leafs["conflict-context"] = types.YLeaf{"ConflictContext", discoveredLink.ConflictContext}
    return &(discoveredLink.EntityData)
}

// NvSatellite_SatellitePriorities
// Satellite priority information table
type NvSatellite_SatellitePriorities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite priority information. The type is slice of
    // NvSatellite_SatellitePriorities_SatellitePriority.
    SatellitePriority []NvSatellite_SatellitePriorities_SatellitePriority
}

func (satellitePriorities *NvSatellite_SatellitePriorities) GetEntityData() *types.CommonEntityData {
    satellitePriorities.EntityData.YFilter = satellitePriorities.YFilter
    satellitePriorities.EntityData.YangName = "satellite-priorities"
    satellitePriorities.EntityData.BundleName = "cisco_ios_xr"
    satellitePriorities.EntityData.ParentYangName = "nv-satellite"
    satellitePriorities.EntityData.SegmentPath = "satellite-priorities"
    satellitePriorities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satellitePriorities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satellitePriorities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satellitePriorities.EntityData.Children = make(map[string]types.YChild)
    satellitePriorities.EntityData.Children["satellite-priority"] = types.YChild{"SatellitePriority", nil}
    for i := range satellitePriorities.SatellitePriority {
        satellitePriorities.EntityData.Children[types.GetSegmentPath(&satellitePriorities.SatellitePriority[i])] = types.YChild{"SatellitePriority", &satellitePriorities.SatellitePriority[i]}
    }
    satellitePriorities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(satellitePriorities.EntityData)
}

// NvSatellite_SatellitePriorities_SatellitePriority
// Satellite priority information
type NvSatellite_SatellitePriorities_SatellitePriority struct {
    EntityData types.CommonEntityData
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

func (satellitePriority *NvSatellite_SatellitePriorities_SatellitePriority) GetEntityData() *types.CommonEntityData {
    satellitePriority.EntityData.YFilter = satellitePriority.YFilter
    satellitePriority.EntityData.YangName = "satellite-priority"
    satellitePriority.EntityData.BundleName = "cisco_ios_xr"
    satellitePriority.EntityData.ParentYangName = "satellite-priorities"
    satellitePriority.EntityData.SegmentPath = "satellite-priority" + "[satellite-id='" + fmt.Sprintf("%v", satellitePriority.SatelliteId) + "']"
    satellitePriority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satellitePriority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satellitePriority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satellitePriority.EntityData.Children = make(map[string]types.YChild)
    satellitePriority.EntityData.Leafs = make(map[string]types.YLeaf)
    satellitePriority.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", satellitePriority.SatelliteId}
    satellitePriority.EntityData.Leafs["satellite-id-xr"] = types.YLeaf{"SatelliteIdXr", satellitePriority.SatelliteIdXr}
    satellitePriority.EntityData.Leafs["rgid"] = types.YLeaf{"Rgid", satellitePriority.Rgid}
    satellitePriority.EntityData.Leafs["best-path-hops"] = types.YLeaf{"BestPathHops", satellitePriority.BestPathHops}
    satellitePriority.EntityData.Leafs["configured-priority"] = types.YLeaf{"ConfiguredPriority", satellitePriority.ConfiguredPriority}
    satellitePriority.EntityData.Leafs["host-priority"] = types.YLeaf{"HostPriority", satellitePriority.HostPriority}
    satellitePriority.EntityData.Leafs["partner-priority"] = types.YLeaf{"PartnerPriority", satellitePriority.PartnerPriority}
    satellitePriority.EntityData.Leafs["multichassis-redundancy"] = types.YLeaf{"MultichassisRedundancy", satellitePriority.MultichassisRedundancy}
    return &(satellitePriority.EntityData)
}

// NvSatellite_SatelliteVersions
// Satellite remote version information table
type NvSatellite_SatelliteVersions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite remote version information. The type is slice of
    // NvSatellite_SatelliteVersions_SatelliteVersion.
    SatelliteVersion []NvSatellite_SatelliteVersions_SatelliteVersion
}

func (satelliteVersions *NvSatellite_SatelliteVersions) GetEntityData() *types.CommonEntityData {
    satelliteVersions.EntityData.YFilter = satelliteVersions.YFilter
    satelliteVersions.EntityData.YangName = "satellite-versions"
    satelliteVersions.EntityData.BundleName = "cisco_ios_xr"
    satelliteVersions.EntityData.ParentYangName = "nv-satellite"
    satelliteVersions.EntityData.SegmentPath = "satellite-versions"
    satelliteVersions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteVersions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteVersions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteVersions.EntityData.Children = make(map[string]types.YChild)
    satelliteVersions.EntityData.Children["satellite-version"] = types.YChild{"SatelliteVersion", nil}
    for i := range satelliteVersions.SatelliteVersion {
        satelliteVersions.EntityData.Children[types.GetSegmentPath(&satelliteVersions.SatelliteVersion[i])] = types.YChild{"SatelliteVersion", &satelliteVersions.SatelliteVersion[i]}
    }
    satelliteVersions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(satelliteVersions.EntityData)
}

// NvSatellite_SatelliteVersions_SatelliteVersion
// Satellite remote version information
type NvSatellite_SatelliteVersions_SatelliteVersion struct {
    EntityData types.CommonEntityData
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

func (satelliteVersion *NvSatellite_SatelliteVersions_SatelliteVersion) GetEntityData() *types.CommonEntityData {
    satelliteVersion.EntityData.YFilter = satelliteVersion.YFilter
    satelliteVersion.EntityData.YangName = "satellite-version"
    satelliteVersion.EntityData.BundleName = "cisco_ios_xr"
    satelliteVersion.EntityData.ParentYangName = "satellite-versions"
    satelliteVersion.EntityData.SegmentPath = "satellite-version" + "[satellite-id='" + fmt.Sprintf("%v", satelliteVersion.SatelliteId) + "']"
    satelliteVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteVersion.EntityData.Children = make(map[string]types.YChild)
    satelliteVersion.EntityData.Children["active-version"] = types.YChild{"ActiveVersion", &satelliteVersion.ActiveVersion}
    satelliteVersion.EntityData.Children["transferred-version"] = types.YChild{"TransferredVersion", &satelliteVersion.TransferredVersion}
    satelliteVersion.EntityData.Children["committed-version"] = types.YChild{"CommittedVersion", &satelliteVersion.CommittedVersion}
    satelliteVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    satelliteVersion.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", satelliteVersion.SatelliteId}
    satelliteVersion.EntityData.Leafs["satellite-id-xr"] = types.YLeaf{"SatelliteIdXr", satelliteVersion.SatelliteIdXr}
    satelliteVersion.EntityData.Leafs["version-check-state"] = types.YLeaf{"VersionCheckState", satelliteVersion.VersionCheckState}
    satelliteVersion.EntityData.Leafs["remote-version-present"] = types.YLeaf{"RemoteVersionPresent", satelliteVersion.RemoteVersionPresent}
    satelliteVersion.EntityData.Leafs["remote-version"] = types.YLeaf{"RemoteVersion", satelliteVersion.RemoteVersion}
    return &(satelliteVersion.EntityData)
}

// NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion
// Satellite active version information
type NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version check state. The type is IcpeOperVerCheckState.
    VersionCheckState interface{}

    // Remote version present. The type is bool.
    RemoteVersionPresent interface{}

    // Remote version. The type is slice of string.
    RemoteVersion []interface{}
}

func (activeVersion *NvSatellite_SatelliteVersions_SatelliteVersion_ActiveVersion) GetEntityData() *types.CommonEntityData {
    activeVersion.EntityData.YFilter = activeVersion.YFilter
    activeVersion.EntityData.YangName = "active-version"
    activeVersion.EntityData.BundleName = "cisco_ios_xr"
    activeVersion.EntityData.ParentYangName = "satellite-version"
    activeVersion.EntityData.SegmentPath = "active-version"
    activeVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeVersion.EntityData.Children = make(map[string]types.YChild)
    activeVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    activeVersion.EntityData.Leafs["version-check-state"] = types.YLeaf{"VersionCheckState", activeVersion.VersionCheckState}
    activeVersion.EntityData.Leafs["remote-version-present"] = types.YLeaf{"RemoteVersionPresent", activeVersion.RemoteVersionPresent}
    activeVersion.EntityData.Leafs["remote-version"] = types.YLeaf{"RemoteVersion", activeVersion.RemoteVersion}
    return &(activeVersion.EntityData)
}

// NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion
// Satellite transferred version information
type NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version check state. The type is IcpeOperVerCheckState.
    VersionCheckState interface{}

    // Remote version present. The type is bool.
    RemoteVersionPresent interface{}

    // Remote version. The type is slice of string.
    RemoteVersion []interface{}
}

func (transferredVersion *NvSatellite_SatelliteVersions_SatelliteVersion_TransferredVersion) GetEntityData() *types.CommonEntityData {
    transferredVersion.EntityData.YFilter = transferredVersion.YFilter
    transferredVersion.EntityData.YangName = "transferred-version"
    transferredVersion.EntityData.BundleName = "cisco_ios_xr"
    transferredVersion.EntityData.ParentYangName = "satellite-version"
    transferredVersion.EntityData.SegmentPath = "transferred-version"
    transferredVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transferredVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transferredVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transferredVersion.EntityData.Children = make(map[string]types.YChild)
    transferredVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    transferredVersion.EntityData.Leafs["version-check-state"] = types.YLeaf{"VersionCheckState", transferredVersion.VersionCheckState}
    transferredVersion.EntityData.Leafs["remote-version-present"] = types.YLeaf{"RemoteVersionPresent", transferredVersion.RemoteVersionPresent}
    transferredVersion.EntityData.Leafs["remote-version"] = types.YLeaf{"RemoteVersion", transferredVersion.RemoteVersion}
    return &(transferredVersion.EntityData)
}

// NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion
// Satellite committed version information
type NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version check state. The type is IcpeOperVerCheckState.
    VersionCheckState interface{}

    // Remote version present. The type is bool.
    RemoteVersionPresent interface{}

    // Remote version. The type is slice of string.
    RemoteVersion []interface{}
}

func (committedVersion *NvSatellite_SatelliteVersions_SatelliteVersion_CommittedVersion) GetEntityData() *types.CommonEntityData {
    committedVersion.EntityData.YFilter = committedVersion.YFilter
    committedVersion.EntityData.YangName = "committed-version"
    committedVersion.EntityData.BundleName = "cisco_ios_xr"
    committedVersion.EntityData.ParentYangName = "satellite-version"
    committedVersion.EntityData.SegmentPath = "committed-version"
    committedVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committedVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committedVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committedVersion.EntityData.Children = make(map[string]types.YChild)
    committedVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    committedVersion.EntityData.Leafs["version-check-state"] = types.YLeaf{"VersionCheckState", committedVersion.VersionCheckState}
    committedVersion.EntityData.Leafs["remote-version-present"] = types.YLeaf{"RemoteVersionPresent", committedVersion.RemoteVersionPresent}
    committedVersion.EntityData.Leafs["remote-version"] = types.YLeaf{"RemoteVersion", committedVersion.RemoteVersion}
    return &(committedVersion.EntityData)
}

// NvSatellite_SatelliteTopologies
// Satellite Topology Information table
type NvSatellite_SatelliteTopologies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite Topology Information. The type is slice of
    // NvSatellite_SatelliteTopologies_SatelliteTopology.
    SatelliteTopology []NvSatellite_SatelliteTopologies_SatelliteTopology
}

func (satelliteTopologies *NvSatellite_SatelliteTopologies) GetEntityData() *types.CommonEntityData {
    satelliteTopologies.EntityData.YFilter = satelliteTopologies.YFilter
    satelliteTopologies.EntityData.YangName = "satellite-topologies"
    satelliteTopologies.EntityData.BundleName = "cisco_ios_xr"
    satelliteTopologies.EntityData.ParentYangName = "nv-satellite"
    satelliteTopologies.EntityData.SegmentPath = "satellite-topologies"
    satelliteTopologies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteTopologies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteTopologies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteTopologies.EntityData.Children = make(map[string]types.YChild)
    satelliteTopologies.EntityData.Children["satellite-topology"] = types.YChild{"SatelliteTopology", nil}
    for i := range satelliteTopologies.SatelliteTopology {
        satelliteTopologies.EntityData.Children[types.GetSegmentPath(&satelliteTopologies.SatelliteTopology[i])] = types.YChild{"SatelliteTopology", &satelliteTopologies.SatelliteTopology[i]}
    }
    satelliteTopologies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(satelliteTopologies.EntityData)
}

// NvSatellite_SatelliteTopologies_SatelliteTopology
// Satellite Topology Information
type NvSatellite_SatelliteTopologies_SatelliteTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
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

func (satelliteTopology *NvSatellite_SatelliteTopologies_SatelliteTopology) GetEntityData() *types.CommonEntityData {
    satelliteTopology.EntityData.YFilter = satelliteTopology.YFilter
    satelliteTopology.EntityData.YangName = "satellite-topology"
    satelliteTopology.EntityData.BundleName = "cisco_ios_xr"
    satelliteTopology.EntityData.ParentYangName = "satellite-topologies"
    satelliteTopology.EntityData.SegmentPath = "satellite-topology" + "[interface-name='" + fmt.Sprintf("%v", satelliteTopology.InterfaceName) + "']"
    satelliteTopology.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteTopology.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteTopology.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteTopology.EntityData.Children = make(map[string]types.YChild)
    satelliteTopology.EntityData.Children["discovered-link"] = types.YChild{"DiscoveredLink", nil}
    for i := range satelliteTopology.DiscoveredLink {
        satelliteTopology.EntityData.Children[types.GetSegmentPath(&satelliteTopology.DiscoveredLink[i])] = types.YChild{"DiscoveredLink", &satelliteTopology.DiscoveredLink[i]}
    }
    satelliteTopology.EntityData.Children["satellite"] = types.YChild{"Satellite", nil}
    for i := range satelliteTopology.Satellite {
        satelliteTopology.EntityData.Children[types.GetSegmentPath(&satelliteTopology.Satellite[i])] = types.YChild{"Satellite", &satelliteTopology.Satellite[i]}
    }
    satelliteTopology.EntityData.Leafs = make(map[string]types.YLeaf)
    satelliteTopology.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", satelliteTopology.InterfaceName}
    satelliteTopology.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", satelliteTopology.InterfaceNameXr}
    satelliteTopology.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", satelliteTopology.InterfaceHandle}
    satelliteTopology.EntityData.Leafs["redundancy-iccp-group"] = types.YLeaf{"RedundancyIccpGroup", satelliteTopology.RedundancyIccpGroup}
    satelliteTopology.EntityData.Leafs["is-physical"] = types.YLeaf{"IsPhysical", satelliteTopology.IsPhysical}
    satelliteTopology.EntityData.Leafs["ring-whole"] = types.YLeaf{"RingWhole", satelliteTopology.RingWhole}
    return &(satelliteTopology.EntityData)
}

// NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink
// Discovered Links table
type NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceHandle interface{}

    // Discovery running. The type is bool.
    DiscoveryRunning interface{}
}

func (discoveredLink *NvSatellite_SatelliteTopologies_SatelliteTopology_DiscoveredLink) GetEntityData() *types.CommonEntityData {
    discoveredLink.EntityData.YFilter = discoveredLink.YFilter
    discoveredLink.EntityData.YangName = "discovered-link"
    discoveredLink.EntityData.BundleName = "cisco_ios_xr"
    discoveredLink.EntityData.ParentYangName = "satellite-topology"
    discoveredLink.EntityData.SegmentPath = "discovered-link"
    discoveredLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discoveredLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discoveredLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discoveredLink.EntityData.Children = make(map[string]types.YChild)
    discoveredLink.EntityData.Leafs = make(map[string]types.YLeaf)
    discoveredLink.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", discoveredLink.InterfaceName}
    discoveredLink.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", discoveredLink.InterfaceHandle}
    discoveredLink.EntityData.Leafs["discovery-running"] = types.YLeaf{"DiscoveryRunning", discoveredLink.DiscoveryRunning}
    return &(discoveredLink.EntityData)
}

// NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite
// Satellite table
type NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Configured. The type is bool.
    Configured interface{}

    // Num hops. The type is interface{} with range: 0..65535.
    NumHops interface{}

    // Type. The type is string.
    Type_ interface{}

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

func (satellite *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite) GetEntityData() *types.CommonEntityData {
    satellite.EntityData.YFilter = satellite.YFilter
    satellite.EntityData.YangName = "satellite"
    satellite.EntityData.BundleName = "cisco_ios_xr"
    satellite.EntityData.ParentYangName = "satellite-topology"
    satellite.EntityData.SegmentPath = "satellite"
    satellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satellite.EntityData.Children = make(map[string]types.YChild)
    satellite.EntityData.Children["fabric-link"] = types.YChild{"FabricLink", nil}
    for i := range satellite.FabricLink {
        satellite.EntityData.Children[types.GetSegmentPath(&satellite.FabricLink[i])] = types.YChild{"FabricLink", &satellite.FabricLink[i]}
    }
    satellite.EntityData.Leafs = make(map[string]types.YLeaf)
    satellite.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", satellite.MacAddress}
    satellite.EntityData.Leafs["configured"] = types.YLeaf{"Configured", satellite.Configured}
    satellite.EntityData.Leafs["num-hops"] = types.YLeaf{"NumHops", satellite.NumHops}
    satellite.EntityData.Leafs["type"] = types.YLeaf{"Type_", satellite.Type_}
    satellite.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", satellite.SatelliteId}
    satellite.EntityData.Leafs["received-serial-number"] = types.YLeaf{"ReceivedSerialNumber", satellite.ReceivedSerialNumber}
    satellite.EntityData.Leafs["received-serial-number-present"] = types.YLeaf{"ReceivedSerialNumberPresent", satellite.ReceivedSerialNumberPresent}
    satellite.EntityData.Leafs["vlan-id"] = types.YLeaf{"VlanId", satellite.VlanId}
    satellite.EntityData.Leafs["display-name"] = types.YLeaf{"DisplayName", satellite.DisplayName}
    satellite.EntityData.Leafs["conflict-reason"] = types.YLeaf{"ConflictReason", satellite.ConflictReason}
    satellite.EntityData.Leafs["conflict-context"] = types.YLeaf{"ConflictContext", satellite.ConflictContext}
    return &(satellite.EntityData)
}

// NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink
// Local Satellite Fabric Link table
type NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink struct {
    EntityData types.CommonEntityData
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

func (fabricLink *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink) GetEntityData() *types.CommonEntityData {
    fabricLink.EntityData.YFilter = fabricLink.YFilter
    fabricLink.EntityData.YangName = "fabric-link"
    fabricLink.EntityData.BundleName = "cisco_ios_xr"
    fabricLink.EntityData.ParentYangName = "satellite"
    fabricLink.EntityData.SegmentPath = "fabric-link"
    fabricLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabricLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabricLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabricLink.EntityData.Children = make(map[string]types.YChild)
    fabricLink.EntityData.Children["remote-device"] = types.YChild{"RemoteDevice", nil}
    for i := range fabricLink.RemoteDevice {
        fabricLink.EntityData.Children[types.GetSegmentPath(&fabricLink.RemoteDevice[i])] = types.YChild{"RemoteDevice", &fabricLink.RemoteDevice[i]}
    }
    fabricLink.EntityData.Leafs = make(map[string]types.YLeaf)
    fabricLink.EntityData.Leafs["icl-id"] = types.YLeaf{"IclId", fabricLink.IclId}
    fabricLink.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", fabricLink.InterfaceName}
    fabricLink.EntityData.Leafs["display-name"] = types.YLeaf{"DisplayName", fabricLink.DisplayName}
    fabricLink.EntityData.Leafs["redundant"] = types.YLeaf{"Redundant", fabricLink.Redundant}
    fabricLink.EntityData.Leafs["active"] = types.YLeaf{"Active", fabricLink.Active}
    fabricLink.EntityData.Leafs["obsolete"] = types.YLeaf{"Obsolete", fabricLink.Obsolete}
    return &(fabricLink.EntityData)
}

// NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice
// Remote Device table
type NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Source. The type is IcpeOperTopoRemoteSource.
    Source interface{}

    // Remote is satellite. The type is bool.
    RemoteIsSatellite interface{}

    // Remote is local host. The type is bool.
    RemoteIsLocalHost interface{}

    // ICL ID. The type is interface{} with range: 0..4294967295.
    IclId interface{}

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceHandle interface{}

    // Interface name. The type is string.
    InterfaceName interface{}
}

func (remoteDevice *NvSatellite_SatelliteTopologies_SatelliteTopology_Satellite_FabricLink_RemoteDevice) GetEntityData() *types.CommonEntityData {
    remoteDevice.EntityData.YFilter = remoteDevice.YFilter
    remoteDevice.EntityData.YangName = "remote-device"
    remoteDevice.EntityData.BundleName = "cisco_ios_xr"
    remoteDevice.EntityData.ParentYangName = "fabric-link"
    remoteDevice.EntityData.SegmentPath = "remote-device"
    remoteDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteDevice.EntityData.Children = make(map[string]types.YChild)
    remoteDevice.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteDevice.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", remoteDevice.MacAddress}
    remoteDevice.EntityData.Leafs["source"] = types.YLeaf{"Source", remoteDevice.Source}
    remoteDevice.EntityData.Leafs["remote-is-satellite"] = types.YLeaf{"RemoteIsSatellite", remoteDevice.RemoteIsSatellite}
    remoteDevice.EntityData.Leafs["remote-is-local-host"] = types.YLeaf{"RemoteIsLocalHost", remoteDevice.RemoteIsLocalHost}
    remoteDevice.EntityData.Leafs["icl-id"] = types.YLeaf{"IclId", remoteDevice.IclId}
    remoteDevice.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", remoteDevice.InterfaceHandle}
    remoteDevice.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", remoteDevice.InterfaceName}
    return &(remoteDevice.EntityData)
}

// NvSatellite_InstallReferenceInfo
// Satellite Install Reference Information
type NvSatellite_InstallReferenceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install Reference Information table.
    References NvSatellite_InstallReferenceInfo_References
}

func (installReferenceInfo *NvSatellite_InstallReferenceInfo) GetEntityData() *types.CommonEntityData {
    installReferenceInfo.EntityData.YFilter = installReferenceInfo.YFilter
    installReferenceInfo.EntityData.YangName = "install-reference-info"
    installReferenceInfo.EntityData.BundleName = "cisco_ios_xr"
    installReferenceInfo.EntityData.ParentYangName = "nv-satellite"
    installReferenceInfo.EntityData.SegmentPath = "install-reference-info"
    installReferenceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installReferenceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installReferenceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installReferenceInfo.EntityData.Children = make(map[string]types.YChild)
    installReferenceInfo.EntityData.Children["references"] = types.YChild{"References", &installReferenceInfo.References}
    installReferenceInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(installReferenceInfo.EntityData)
}

// NvSatellite_InstallReferenceInfo_References
// Install Reference Information table
type NvSatellite_InstallReferenceInfo_References struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Install Reference Information. The type is slice of
    // NvSatellite_InstallReferenceInfo_References_Reference.
    Reference []NvSatellite_InstallReferenceInfo_References_Reference
}

func (references *NvSatellite_InstallReferenceInfo_References) GetEntityData() *types.CommonEntityData {
    references.EntityData.YFilter = references.YFilter
    references.EntityData.YangName = "references"
    references.EntityData.BundleName = "cisco_ios_xr"
    references.EntityData.ParentYangName = "install-reference-info"
    references.EntityData.SegmentPath = "references"
    references.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    references.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    references.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    references.EntityData.Children = make(map[string]types.YChild)
    references.EntityData.Children["reference"] = types.YChild{"Reference", nil}
    for i := range references.Reference {
        references.EntityData.Children[types.GetSegmentPath(&references.Reference[i])] = types.YChild{"Reference", &references.Reference[i]}
    }
    references.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(references.EntityData)
}

// NvSatellite_InstallReferenceInfo_References_Reference
// Install Reference Information
type NvSatellite_InstallReferenceInfo_References_Reference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ReferenceName interface{}

    // Reference name. The type is string.
    ReferenceNameXr interface{}

    // Reference files. The type is slice of string.
    ReferenceFile []interface{}
}

func (reference *NvSatellite_InstallReferenceInfo_References_Reference) GetEntityData() *types.CommonEntityData {
    reference.EntityData.YFilter = reference.YFilter
    reference.EntityData.YangName = "reference"
    reference.EntityData.BundleName = "cisco_ios_xr"
    reference.EntityData.ParentYangName = "references"
    reference.EntityData.SegmentPath = "reference" + "[reference-name='" + fmt.Sprintf("%v", reference.ReferenceName) + "']"
    reference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reference.EntityData.Children = make(map[string]types.YChild)
    reference.EntityData.Leafs = make(map[string]types.YLeaf)
    reference.EntityData.Leafs["reference-name"] = types.YLeaf{"ReferenceName", reference.ReferenceName}
    reference.EntityData.Leafs["reference-name-xr"] = types.YLeaf{"ReferenceNameXr", reference.ReferenceNameXr}
    reference.EntityData.Leafs["reference-file"] = types.YLeaf{"ReferenceFile", reference.ReferenceFile}
    return &(reference.EntityData)
}

// NvSatellite_InstallOpProgresses
// Current percentage of install table
type NvSatellite_InstallOpProgresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current percentage of install. The type is slice of
    // NvSatellite_InstallOpProgresses_InstallOpProgress.
    InstallOpProgress []NvSatellite_InstallOpProgresses_InstallOpProgress
}

func (installOpProgresses *NvSatellite_InstallOpProgresses) GetEntityData() *types.CommonEntityData {
    installOpProgresses.EntityData.YFilter = installOpProgresses.YFilter
    installOpProgresses.EntityData.YangName = "install-op-progresses"
    installOpProgresses.EntityData.BundleName = "cisco_ios_xr"
    installOpProgresses.EntityData.ParentYangName = "nv-satellite"
    installOpProgresses.EntityData.SegmentPath = "install-op-progresses"
    installOpProgresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installOpProgresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installOpProgresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installOpProgresses.EntityData.Children = make(map[string]types.YChild)
    installOpProgresses.EntityData.Children["install-op-progress"] = types.YChild{"InstallOpProgress", nil}
    for i := range installOpProgresses.InstallOpProgress {
        installOpProgresses.EntityData.Children[types.GetSegmentPath(&installOpProgresses.InstallOpProgress[i])] = types.YChild{"InstallOpProgress", &installOpProgresses.InstallOpProgress[i]}
    }
    installOpProgresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(installOpProgresses.EntityData)
}

// NvSatellite_InstallOpProgresses_InstallOpProgress
// Current percentage of install
type NvSatellite_InstallOpProgresses_InstallOpProgress struct {
    EntityData types.CommonEntityData
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

func (installOpProgress *NvSatellite_InstallOpProgresses_InstallOpProgress) GetEntityData() *types.CommonEntityData {
    installOpProgress.EntityData.YFilter = installOpProgress.YFilter
    installOpProgress.EntityData.YangName = "install-op-progress"
    installOpProgress.EntityData.BundleName = "cisco_ios_xr"
    installOpProgress.EntityData.ParentYangName = "install-op-progresses"
    installOpProgress.EntityData.SegmentPath = "install-op-progress" + "[operation-id='" + fmt.Sprintf("%v", installOpProgress.OperationId) + "']"
    installOpProgress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installOpProgress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installOpProgress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installOpProgress.EntityData.Children = make(map[string]types.YChild)
    installOpProgress.EntityData.Leafs = make(map[string]types.YLeaf)
    installOpProgress.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", installOpProgress.OperationId}
    installOpProgress.EntityData.Leafs["operation-id-xr"] = types.YLeaf{"OperationIdXr", installOpProgress.OperationIdXr}
    installOpProgress.EntityData.Leafs["progress-percentage"] = types.YLeaf{"ProgressPercentage", installOpProgress.ProgressPercentage}
    installOpProgress.EntityData.Leafs["satellite-count"] = types.YLeaf{"SatelliteCount", installOpProgress.SatelliteCount}
    return &(installOpProgress.EntityData)
}

// NvSatellite_ReloadStatuses
// Detailed breakdown of reload status table
type NvSatellite_ReloadStatuses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed breakdown of reload status. The type is slice of
    // NvSatellite_ReloadStatuses_ReloadStatus.
    ReloadStatus []NvSatellite_ReloadStatuses_ReloadStatus
}

func (reloadStatuses *NvSatellite_ReloadStatuses) GetEntityData() *types.CommonEntityData {
    reloadStatuses.EntityData.YFilter = reloadStatuses.YFilter
    reloadStatuses.EntityData.YangName = "reload-statuses"
    reloadStatuses.EntityData.BundleName = "cisco_ios_xr"
    reloadStatuses.EntityData.ParentYangName = "nv-satellite"
    reloadStatuses.EntityData.SegmentPath = "reload-statuses"
    reloadStatuses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reloadStatuses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reloadStatuses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reloadStatuses.EntityData.Children = make(map[string]types.YChild)
    reloadStatuses.EntityData.Children["reload-status"] = types.YChild{"ReloadStatus", nil}
    for i := range reloadStatuses.ReloadStatus {
        reloadStatuses.EntityData.Children[types.GetSegmentPath(&reloadStatuses.ReloadStatus[i])] = types.YChild{"ReloadStatus", &reloadStatuses.ReloadStatus[i]}
    }
    reloadStatuses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reloadStatuses.EntityData)
}

// NvSatellite_ReloadStatuses_ReloadStatus
// Detailed breakdown of reload status
type NvSatellite_ReloadStatuses_ReloadStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite range. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (reloadStatus *NvSatellite_ReloadStatuses_ReloadStatus) GetEntityData() *types.CommonEntityData {
    reloadStatus.EntityData.YFilter = reloadStatus.YFilter
    reloadStatus.EntityData.YangName = "reload-status"
    reloadStatus.EntityData.BundleName = "cisco_ios_xr"
    reloadStatus.EntityData.ParentYangName = "reload-statuses"
    reloadStatus.EntityData.SegmentPath = "reload-status" + "[satellite-range='" + fmt.Sprintf("%v", reloadStatus.SatelliteRange) + "']"
    reloadStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reloadStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reloadStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reloadStatus.EntityData.Children = make(map[string]types.YChild)
    reloadStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    reloadStatus.EntityData.Leafs["satellite-range"] = types.YLeaf{"SatelliteRange", reloadStatus.SatelliteRange}
    reloadStatus.EntityData.Leafs["satellite-range-xr"] = types.YLeaf{"SatelliteRangeXr", reloadStatus.SatelliteRangeXr}
    reloadStatus.EntityData.Leafs["sats-not-initiated"] = types.YLeaf{"SatsNotInitiated", reloadStatus.SatsNotInitiated}
    reloadStatus.EntityData.Leafs["sats-reloading"] = types.YLeaf{"SatsReloading", reloadStatus.SatsReloading}
    reloadStatus.EntityData.Leafs["sats-reloaded"] = types.YLeaf{"SatsReloaded", reloadStatus.SatsReloaded}
    reloadStatus.EntityData.Leafs["sats-reload-failed"] = types.YLeaf{"SatsReloadFailed", reloadStatus.SatsReloadFailed}
    return &(reloadStatus.EntityData)
}

// NvSatellite_Install
// Satellite Install Information
type NvSatellite_Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite Software Package Information table.
    SatelliteSoftwareVersions NvSatellite_Install_SatelliteSoftwareVersions
}

func (install *NvSatellite_Install) GetEntityData() *types.CommonEntityData {
    install.EntityData.YFilter = install.YFilter
    install.EntityData.YangName = "install"
    install.EntityData.BundleName = "cisco_ios_xr"
    install.EntityData.ParentYangName = "nv-satellite"
    install.EntityData.SegmentPath = "install"
    install.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    install.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    install.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    install.EntityData.Children = make(map[string]types.YChild)
    install.EntityData.Children["satellite-software-versions"] = types.YChild{"SatelliteSoftwareVersions", &install.SatelliteSoftwareVersions}
    install.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(install.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions
// Satellite Software Package Information table
type NvSatellite_Install_SatelliteSoftwareVersions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite Software Package Information. The type is slice of
    // NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion.
    SatelliteSoftwareVersion []NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion
}

func (satelliteSoftwareVersions *NvSatellite_Install_SatelliteSoftwareVersions) GetEntityData() *types.CommonEntityData {
    satelliteSoftwareVersions.EntityData.YFilter = satelliteSoftwareVersions.YFilter
    satelliteSoftwareVersions.EntityData.YangName = "satellite-software-versions"
    satelliteSoftwareVersions.EntityData.BundleName = "cisco_ios_xr"
    satelliteSoftwareVersions.EntityData.ParentYangName = "install"
    satelliteSoftwareVersions.EntityData.SegmentPath = "satellite-software-versions"
    satelliteSoftwareVersions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteSoftwareVersions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteSoftwareVersions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteSoftwareVersions.EntityData.Children = make(map[string]types.YChild)
    satelliteSoftwareVersions.EntityData.Children["satellite-software-version"] = types.YChild{"SatelliteSoftwareVersion", nil}
    for i := range satelliteSoftwareVersions.SatelliteSoftwareVersion {
        satelliteSoftwareVersions.EntityData.Children[types.GetSegmentPath(&satelliteSoftwareVersions.SatelliteSoftwareVersion[i])] = types.YChild{"SatelliteSoftwareVersion", &satelliteSoftwareVersions.SatelliteSoftwareVersion[i]}
    }
    satelliteSoftwareVersions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(satelliteSoftwareVersions.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion
// Satellite Software Package Information
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion struct {
    EntityData types.CommonEntityData
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

func (satelliteSoftwareVersion *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion) GetEntityData() *types.CommonEntityData {
    satelliteSoftwareVersion.EntityData.YFilter = satelliteSoftwareVersion.YFilter
    satelliteSoftwareVersion.EntityData.YangName = "satellite-software-version"
    satelliteSoftwareVersion.EntityData.BundleName = "cisco_ios_xr"
    satelliteSoftwareVersion.EntityData.ParentYangName = "satellite-software-versions"
    satelliteSoftwareVersion.EntityData.SegmentPath = "satellite-software-version" + "[satellite-id='" + fmt.Sprintf("%v", satelliteSoftwareVersion.SatelliteId) + "']"
    satelliteSoftwareVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteSoftwareVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteSoftwareVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteSoftwareVersion.EntityData.Children = make(map[string]types.YChild)
    satelliteSoftwareVersion.EntityData.Children["install-package-info"] = types.YChild{"InstallPackageInfo", &satelliteSoftwareVersion.InstallPackageInfo}
    satelliteSoftwareVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    satelliteSoftwareVersion.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", satelliteSoftwareVersion.SatelliteId}
    satelliteSoftwareVersion.EntityData.Leafs["satellite-id-xr"] = types.YLeaf{"SatelliteIdXr", satelliteSoftwareVersion.SatelliteIdXr}
    satelliteSoftwareVersion.EntityData.Leafs["package-support"] = types.YLeaf{"PackageSupport", satelliteSoftwareVersion.PackageSupport}
    return &(satelliteSoftwareVersion.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo
// Package Data on this Satellite
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active Packages running on this Satellite.
    ActivePackages NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages

    // Inactive Packages on this Satellite.
    InactivePackages NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages

    // Committed Packages running on this Satellite.
    CommittedPackages NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages
}

func (installPackageInfo *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo) GetEntityData() *types.CommonEntityData {
    installPackageInfo.EntityData.YFilter = installPackageInfo.YFilter
    installPackageInfo.EntityData.YangName = "install-package-info"
    installPackageInfo.EntityData.BundleName = "cisco_ios_xr"
    installPackageInfo.EntityData.ParentYangName = "satellite-software-version"
    installPackageInfo.EntityData.SegmentPath = "install-package-info"
    installPackageInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installPackageInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installPackageInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installPackageInfo.EntityData.Children = make(map[string]types.YChild)
    installPackageInfo.EntityData.Children["active-packages"] = types.YChild{"ActivePackages", &installPackageInfo.ActivePackages}
    installPackageInfo.EntityData.Children["inactive-packages"] = types.YChild{"InactivePackages", &installPackageInfo.InactivePackages}
    installPackageInfo.EntityData.Children["committed-packages"] = types.YChild{"CommittedPackages", &installPackageInfo.CommittedPackages}
    installPackageInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(installPackageInfo.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages
// Active Packages running on this Satellite
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A package on this Satellite table. The type is slice of
    // NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package_.
    Package_ []NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package
}

func (activePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages) GetEntityData() *types.CommonEntityData {
    activePackages.EntityData.YFilter = activePackages.YFilter
    activePackages.EntityData.YangName = "active-packages"
    activePackages.EntityData.BundleName = "cisco_ios_xr"
    activePackages.EntityData.ParentYangName = "install-package-info"
    activePackages.EntityData.SegmentPath = "active-packages"
    activePackages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activePackages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activePackages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activePackages.EntityData.Children = make(map[string]types.YChild)
    activePackages.EntityData.Children["package"] = types.YChild{"Package_", nil}
    for i := range activePackages.Package_ {
        activePackages.EntityData.Children[types.GetSegmentPath(&activePackages.Package_[i])] = types.YChild{"Package_", &activePackages.Package_[i]}
    }
    activePackages.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(activePackages.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package
// A package on this Satellite table
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is string.
    Name interface{}

    // Version. The type is string.
    Version interface{}

    // Is base image. The type is bool.
    IsBaseImage interface{}
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_ActivePackages_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "active-packages"
    self.EntityData.SegmentPath = "package"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    self.EntityData.Leafs["version"] = types.YLeaf{"Version", self.Version}
    self.EntityData.Leafs["is-base-image"] = types.YLeaf{"IsBaseImage", self.IsBaseImage}
    return &(self.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages
// Inactive Packages on this Satellite
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A package on this Satellite table. The type is slice of
    // NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package_.
    Package_ []NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package
}

func (inactivePackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages) GetEntityData() *types.CommonEntityData {
    inactivePackages.EntityData.YFilter = inactivePackages.YFilter
    inactivePackages.EntityData.YangName = "inactive-packages"
    inactivePackages.EntityData.BundleName = "cisco_ios_xr"
    inactivePackages.EntityData.ParentYangName = "install-package-info"
    inactivePackages.EntityData.SegmentPath = "inactive-packages"
    inactivePackages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactivePackages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactivePackages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactivePackages.EntityData.Children = make(map[string]types.YChild)
    inactivePackages.EntityData.Children["package"] = types.YChild{"Package_", nil}
    for i := range inactivePackages.Package_ {
        inactivePackages.EntityData.Children[types.GetSegmentPath(&inactivePackages.Package_[i])] = types.YChild{"Package_", &inactivePackages.Package_[i]}
    }
    inactivePackages.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inactivePackages.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package
// A package on this Satellite table
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is string.
    Name interface{}

    // Version. The type is string.
    Version interface{}

    // Is base image. The type is bool.
    IsBaseImage interface{}
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_InactivePackages_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "inactive-packages"
    self.EntityData.SegmentPath = "package"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    self.EntityData.Leafs["version"] = types.YLeaf{"Version", self.Version}
    self.EntityData.Leafs["is-base-image"] = types.YLeaf{"IsBaseImage", self.IsBaseImage}
    return &(self.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages
// Committed Packages running on this Satellite
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A package on this Satellite table. The type is slice of
    // NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package_.
    Package_ []NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package
}

func (committedPackages *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages) GetEntityData() *types.CommonEntityData {
    committedPackages.EntityData.YFilter = committedPackages.YFilter
    committedPackages.EntityData.YangName = "committed-packages"
    committedPackages.EntityData.BundleName = "cisco_ios_xr"
    committedPackages.EntityData.ParentYangName = "install-package-info"
    committedPackages.EntityData.SegmentPath = "committed-packages"
    committedPackages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committedPackages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committedPackages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committedPackages.EntityData.Children = make(map[string]types.YChild)
    committedPackages.EntityData.Children["package"] = types.YChild{"Package_", nil}
    for i := range committedPackages.Package_ {
        committedPackages.EntityData.Children[types.GetSegmentPath(&committedPackages.Package_[i])] = types.YChild{"Package_", &committedPackages.Package_[i]}
    }
    committedPackages.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(committedPackages.EntityData)
}

// NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package
// A package on this Satellite table
type NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is string.
    Name interface{}

    // Version. The type is string.
    Version interface{}

    // Is base image. The type is bool.
    IsBaseImage interface{}
}

func (self *NvSatellite_Install_SatelliteSoftwareVersions_SatelliteSoftwareVersion_InstallPackageInfo_CommittedPackages_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "committed-packages"
    self.EntityData.SegmentPath = "package"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    self.EntityData.Leafs["version"] = types.YLeaf{"Version", self.Version}
    self.EntityData.Leafs["is-base-image"] = types.YLeaf{"IsBaseImage", self.IsBaseImage}
    return &(self.EntityData)
}

// NvSatellite_InstallOpStatuses
// Detailed breakdown of install status table
type NvSatellite_InstallOpStatuses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed breakdown of install status. The type is slice of
    // NvSatellite_InstallOpStatuses_InstallOpStatus.
    InstallOpStatus []NvSatellite_InstallOpStatuses_InstallOpStatus
}

func (installOpStatuses *NvSatellite_InstallOpStatuses) GetEntityData() *types.CommonEntityData {
    installOpStatuses.EntityData.YFilter = installOpStatuses.YFilter
    installOpStatuses.EntityData.YangName = "install-op-statuses"
    installOpStatuses.EntityData.BundleName = "cisco_ios_xr"
    installOpStatuses.EntityData.ParentYangName = "nv-satellite"
    installOpStatuses.EntityData.SegmentPath = "install-op-statuses"
    installOpStatuses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installOpStatuses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installOpStatuses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installOpStatuses.EntityData.Children = make(map[string]types.YChild)
    installOpStatuses.EntityData.Children["install-op-status"] = types.YChild{"InstallOpStatus", nil}
    for i := range installOpStatuses.InstallOpStatus {
        installOpStatuses.EntityData.Children[types.GetSegmentPath(&installOpStatuses.InstallOpStatus[i])] = types.YChild{"InstallOpStatus", &installOpStatuses.InstallOpStatus[i]}
    }
    installOpStatuses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(installOpStatuses.EntityData)
}

// NvSatellite_InstallOpStatuses_InstallOpStatus
// Detailed breakdown of install status
type NvSatellite_InstallOpStatuses_InstallOpStatus struct {
    EntityData types.CommonEntityData
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

func (installOpStatus *NvSatellite_InstallOpStatuses_InstallOpStatus) GetEntityData() *types.CommonEntityData {
    installOpStatus.EntityData.YFilter = installOpStatus.YFilter
    installOpStatus.EntityData.YangName = "install-op-status"
    installOpStatus.EntityData.BundleName = "cisco_ios_xr"
    installOpStatus.EntityData.ParentYangName = "install-op-statuses"
    installOpStatus.EntityData.SegmentPath = "install-op-status" + "[operation-id='" + fmt.Sprintf("%v", installOpStatus.OperationId) + "']"
    installOpStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installOpStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installOpStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installOpStatus.EntityData.Children = make(map[string]types.YChild)
    installOpStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    installOpStatus.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", installOpStatus.OperationId}
    installOpStatus.EntityData.Leafs["operation-id-xr"] = types.YLeaf{"OperationIdXr", installOpStatus.OperationIdXr}
    installOpStatus.EntityData.Leafs["satellite-range"] = types.YLeaf{"SatelliteRange", installOpStatus.SatelliteRange}
    installOpStatus.EntityData.Leafs["sats-not-initiated"] = types.YLeaf{"SatsNotInitiated", installOpStatus.SatsNotInitiated}
    installOpStatus.EntityData.Leafs["sats-transferring"] = types.YLeaf{"SatsTransferring", installOpStatus.SatsTransferring}
    installOpStatus.EntityData.Leafs["sats-activating"] = types.YLeaf{"SatsActivating", installOpStatus.SatsActivating}
    installOpStatus.EntityData.Leafs["sats-updating"] = types.YLeaf{"SatsUpdating", installOpStatus.SatsUpdating}
    installOpStatus.EntityData.Leafs["sats-deactivating"] = types.YLeaf{"SatsDeactivating", installOpStatus.SatsDeactivating}
    installOpStatus.EntityData.Leafs["sats-removing"] = types.YLeaf{"SatsRemoving", installOpStatus.SatsRemoving}
    installOpStatus.EntityData.Leafs["sats-transfer-failed"] = types.YLeaf{"SatsTransferFailed", installOpStatus.SatsTransferFailed}
    installOpStatus.EntityData.Leafs["sats-activate-failed"] = types.YLeaf{"SatsActivateFailed", installOpStatus.SatsActivateFailed}
    installOpStatus.EntityData.Leafs["sats-update-failed"] = types.YLeaf{"SatsUpdateFailed", installOpStatus.SatsUpdateFailed}
    installOpStatus.EntityData.Leafs["sats-deactivate-failed"] = types.YLeaf{"SatsDeactivateFailed", installOpStatus.SatsDeactivateFailed}
    installOpStatus.EntityData.Leafs["sats-remove-failed"] = types.YLeaf{"SatsRemoveFailed", installOpStatus.SatsRemoveFailed}
    installOpStatus.EntityData.Leafs["sats-transfer-aborted"] = types.YLeaf{"SatsTransferAborted", installOpStatus.SatsTransferAborted}
    installOpStatus.EntityData.Leafs["sats-activate-aborted"] = types.YLeaf{"SatsActivateAborted", installOpStatus.SatsActivateAborted}
    installOpStatus.EntityData.Leafs["sats-update-aborted"] = types.YLeaf{"SatsUpdateAborted", installOpStatus.SatsUpdateAborted}
    installOpStatus.EntityData.Leafs["sats-deactivate-aborted"] = types.YLeaf{"SatsDeactivateAborted", installOpStatus.SatsDeactivateAborted}
    installOpStatus.EntityData.Leafs["sats-remove-aborted"] = types.YLeaf{"SatsRemoveAborted", installOpStatus.SatsRemoveAborted}
    installOpStatus.EntityData.Leafs["sats-no-operation"] = types.YLeaf{"SatsNoOperation", installOpStatus.SatsNoOperation}
    installOpStatus.EntityData.Leafs["sats-completed"] = types.YLeaf{"SatsCompleted", installOpStatus.SatsCompleted}
    return &(installOpStatus.EntityData)
}

// NvSatellite_SatelliteProperties
// ICPE GCO operational information
type NvSatellite_SatelliteProperties struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite ID range table.
    IdRanges NvSatellite_SatelliteProperties_IdRanges
}

func (satelliteProperties *NvSatellite_SatelliteProperties) GetEntityData() *types.CommonEntityData {
    satelliteProperties.EntityData.YFilter = satelliteProperties.YFilter
    satelliteProperties.EntityData.YangName = "satellite-properties"
    satelliteProperties.EntityData.BundleName = "cisco_ios_xr"
    satelliteProperties.EntityData.ParentYangName = "nv-satellite"
    satelliteProperties.EntityData.SegmentPath = "satellite-properties"
    satelliteProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteProperties.EntityData.Children = make(map[string]types.YChild)
    satelliteProperties.EntityData.Children["id-ranges"] = types.YChild{"IdRanges", &satelliteProperties.IdRanges}
    satelliteProperties.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(satelliteProperties.EntityData)
}

// NvSatellite_SatelliteProperties_IdRanges
// Satellite ID range table
type NvSatellite_SatelliteProperties_IdRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite ID range. The type is slice of
    // NvSatellite_SatelliteProperties_IdRanges_IdRange.
    IdRange []NvSatellite_SatelliteProperties_IdRanges_IdRange
}

func (idRanges *NvSatellite_SatelliteProperties_IdRanges) GetEntityData() *types.CommonEntityData {
    idRanges.EntityData.YFilter = idRanges.YFilter
    idRanges.EntityData.YangName = "id-ranges"
    idRanges.EntityData.BundleName = "cisco_ios_xr"
    idRanges.EntityData.ParentYangName = "satellite-properties"
    idRanges.EntityData.SegmentPath = "id-ranges"
    idRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idRanges.EntityData.Children = make(map[string]types.YChild)
    idRanges.EntityData.Children["id-range"] = types.YChild{"IdRange", nil}
    for i := range idRanges.IdRange {
        idRanges.EntityData.Children[types.GetSegmentPath(&idRanges.IdRange[i])] = types.YChild{"IdRange", &idRanges.IdRange[i]}
    }
    idRanges.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(idRanges.EntityData)
}

// NvSatellite_SatelliteProperties_IdRanges_IdRange
// Satellite ID range
type NvSatellite_SatelliteProperties_IdRanges_IdRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sat ID range. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SatIdRange interface{}

    // Min. The type is interface{} with range: 0..2147483647.
    Min interface{}

    // Max. The type is interface{} with range: 0..2147483647.
    Max interface{}
}

func (idRange *NvSatellite_SatelliteProperties_IdRanges_IdRange) GetEntityData() *types.CommonEntityData {
    idRange.EntityData.YFilter = idRange.YFilter
    idRange.EntityData.YangName = "id-range"
    idRange.EntityData.BundleName = "cisco_ios_xr"
    idRange.EntityData.ParentYangName = "id-ranges"
    idRange.EntityData.SegmentPath = "id-range" + "[sat-id-range='" + fmt.Sprintf("%v", idRange.SatIdRange) + "']"
    idRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idRange.EntityData.Children = make(map[string]types.YChild)
    idRange.EntityData.Leafs = make(map[string]types.YLeaf)
    idRange.EntityData.Leafs["sat-id-range"] = types.YLeaf{"SatIdRange", idRange.SatIdRange}
    idRange.EntityData.Leafs["min"] = types.YLeaf{"Min", idRange.Min}
    idRange.EntityData.Leafs["max"] = types.YLeaf{"Max", idRange.Max}
    return &(idRange.EntityData)
}

// NvSatellite_SdacpDiscovery2S
// ICPE Configured interface state information
// table
type NvSatellite_SdacpDiscovery2S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICPE Configured interface state information. The type is slice of
    // NvSatellite_SdacpDiscovery2S_SdacpDiscovery2.
    SdacpDiscovery2 []NvSatellite_SdacpDiscovery2S_SdacpDiscovery2
}

func (sdacpDiscovery2S *NvSatellite_SdacpDiscovery2S) GetEntityData() *types.CommonEntityData {
    sdacpDiscovery2S.EntityData.YFilter = sdacpDiscovery2S.YFilter
    sdacpDiscovery2S.EntityData.YangName = "sdacp-discovery2s"
    sdacpDiscovery2S.EntityData.BundleName = "cisco_ios_xr"
    sdacpDiscovery2S.EntityData.ParentYangName = "nv-satellite"
    sdacpDiscovery2S.EntityData.SegmentPath = "Cisco-IOS-XR-icpe-sdacp-oper:sdacp-discovery2s"
    sdacpDiscovery2S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdacpDiscovery2S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdacpDiscovery2S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdacpDiscovery2S.EntityData.Children = make(map[string]types.YChild)
    sdacpDiscovery2S.EntityData.Children["sdacp-discovery2"] = types.YChild{"SdacpDiscovery2", nil}
    for i := range sdacpDiscovery2S.SdacpDiscovery2 {
        sdacpDiscovery2S.EntityData.Children[types.GetSegmentPath(&sdacpDiscovery2S.SdacpDiscovery2[i])] = types.YChild{"SdacpDiscovery2", &sdacpDiscovery2S.SdacpDiscovery2[i]}
    }
    sdacpDiscovery2S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sdacpDiscovery2S.EntityData)
}

// NvSatellite_SdacpDiscovery2S_SdacpDiscovery2
// ICPE Configured interface state information
type NvSatellite_SdacpDiscovery2S_SdacpDiscovery2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // ICPE Discovery interface state information table. The type is slice of
    // NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface_.
    Interface_ []NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface

    // ICPE Satellite state information table. The type is slice of
    // NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite.
    Satellite []NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite
}

func (sdacpDiscovery2 *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2) GetEntityData() *types.CommonEntityData {
    sdacpDiscovery2.EntityData.YFilter = sdacpDiscovery2.YFilter
    sdacpDiscovery2.EntityData.YangName = "sdacp-discovery2"
    sdacpDiscovery2.EntityData.BundleName = "cisco_ios_xr"
    sdacpDiscovery2.EntityData.ParentYangName = "sdacp-discovery2s"
    sdacpDiscovery2.EntityData.SegmentPath = "sdacp-discovery2" + "[interface-name='" + fmt.Sprintf("%v", sdacpDiscovery2.InterfaceName) + "']"
    sdacpDiscovery2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdacpDiscovery2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdacpDiscovery2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdacpDiscovery2.EntityData.Children = make(map[string]types.YChild)
    sdacpDiscovery2.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range sdacpDiscovery2.Interface_ {
        sdacpDiscovery2.EntityData.Children[types.GetSegmentPath(&sdacpDiscovery2.Interface_[i])] = types.YChild{"Interface_", &sdacpDiscovery2.Interface_[i]}
    }
    sdacpDiscovery2.EntityData.Children["satellite"] = types.YChild{"Satellite", nil}
    for i := range sdacpDiscovery2.Satellite {
        sdacpDiscovery2.EntityData.Children[types.GetSegmentPath(&sdacpDiscovery2.Satellite[i])] = types.YChild{"Satellite", &sdacpDiscovery2.Satellite[i]}
    }
    sdacpDiscovery2.EntityData.Leafs = make(map[string]types.YLeaf)
    sdacpDiscovery2.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", sdacpDiscovery2.InterfaceName}
    sdacpDiscovery2.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", sdacpDiscovery2.InterfaceNameXr}
    return &(sdacpDiscovery2.EntityData)
}

// NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface
// ICPE Discovery interface state information table
type NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface status. The type is DpmProtoState.
    InterfaceStatus interface{}
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "sdacp-discovery2"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-status"] = types.YLeaf{"InterfaceStatus", self.InterfaceStatus}
    return &(self.EntityData)
}

// NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite
// ICPE Satellite state information table
type NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteId interface{}

    // Satellite status. The type is DpmProtoState.
    SatelliteStatus interface{}

    // Conflict reason. The type is interface{} with range: 0..4294967295.
    ConflictReason interface{}

    // Satellite IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SatelliteIpAddress interface{}

    // Host IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HostIpAddress interface{}

    // ICPE Discovered satellite state information table. The type is slice of
    // NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface_.
    Interface_ []NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface
}

func (satellite *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite) GetEntityData() *types.CommonEntityData {
    satellite.EntityData.YFilter = satellite.YFilter
    satellite.EntityData.YangName = "satellite"
    satellite.EntityData.BundleName = "cisco_ios_xr"
    satellite.EntityData.ParentYangName = "sdacp-discovery2"
    satellite.EntityData.SegmentPath = "satellite"
    satellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satellite.EntityData.Children = make(map[string]types.YChild)
    satellite.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range satellite.Interface_ {
        satellite.EntityData.Children[types.GetSegmentPath(&satellite.Interface_[i])] = types.YChild{"Interface_", &satellite.Interface_[i]}
    }
    satellite.EntityData.Leafs = make(map[string]types.YLeaf)
    satellite.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", satellite.SatelliteId}
    satellite.EntityData.Leafs["satellite-status"] = types.YLeaf{"SatelliteStatus", satellite.SatelliteStatus}
    satellite.EntityData.Leafs["conflict-reason"] = types.YLeaf{"ConflictReason", satellite.ConflictReason}
    satellite.EntityData.Leafs["satellite-ip-address"] = types.YLeaf{"SatelliteIpAddress", satellite.SatelliteIpAddress}
    satellite.EntityData.Leafs["host-ip-address"] = types.YLeaf{"HostIpAddress", satellite.HostIpAddress}
    return &(satellite.EntityData)
}

// NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface
// ICPE Discovered satellite state information
// table
type NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SatelliteInterfaceMac interface{}

    // Satellite chassis MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SatelliteChassisMac interface{}

    // Satellite serial id. The type is string.
    SatelliteSerialId interface{}

    // Satellite module vendor. The type is string.
    SatelliteModuleVendor interface{}
}

func (self *NvSatellite_SdacpDiscovery2S_SdacpDiscovery2_Satellite_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "satellite"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", self.InterfaceHandle}
    self.EntityData.Leafs["satellite-status"] = types.YLeaf{"SatelliteStatus", self.SatelliteStatus}
    self.EntityData.Leafs["conflict-reason"] = types.YLeaf{"ConflictReason", self.ConflictReason}
    self.EntityData.Leafs["satellite-chassis-vendor"] = types.YLeaf{"SatelliteChassisVendor", self.SatelliteChassisVendor}
    self.EntityData.Leafs["satellite-interface-id"] = types.YLeaf{"SatelliteInterfaceId", self.SatelliteInterfaceId}
    self.EntityData.Leafs["satellite-interface-mac"] = types.YLeaf{"SatelliteInterfaceMac", self.SatelliteInterfaceMac}
    self.EntityData.Leafs["satellite-chassis-mac"] = types.YLeaf{"SatelliteChassisMac", self.SatelliteChassisMac}
    self.EntityData.Leafs["satellite-serial-id"] = types.YLeaf{"SatelliteSerialId", self.SatelliteSerialId}
    self.EntityData.Leafs["satellite-module-vendor"] = types.YLeaf{"SatelliteModuleVendor", self.SatelliteModuleVendor}
    return &(self.EntityData)
}

// NvSatellite_IcpeDpms
// ICPE DPM operational information table
type NvSatellite_IcpeDpms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICPE DPM operational information. The type is slice of
    // NvSatellite_IcpeDpms_IcpeDpm.
    IcpeDpm []NvSatellite_IcpeDpms_IcpeDpm
}

func (icpeDpms *NvSatellite_IcpeDpms) GetEntityData() *types.CommonEntityData {
    icpeDpms.EntityData.YFilter = icpeDpms.YFilter
    icpeDpms.EntityData.YangName = "icpe-dpms"
    icpeDpms.EntityData.BundleName = "cisco_ios_xr"
    icpeDpms.EntityData.ParentYangName = "nv-satellite"
    icpeDpms.EntityData.SegmentPath = "Cisco-IOS-XR-icpe-sdacp-oper:icpe-dpms"
    icpeDpms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icpeDpms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icpeDpms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icpeDpms.EntityData.Children = make(map[string]types.YChild)
    icpeDpms.EntityData.Children["icpe-dpm"] = types.YChild{"IcpeDpm", nil}
    for i := range icpeDpms.IcpeDpm {
        icpeDpms.EntityData.Children[types.GetSegmentPath(&icpeDpms.IcpeDpm[i])] = types.YChild{"IcpeDpm", &icpeDpms.IcpeDpm[i]}
    }
    icpeDpms.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(icpeDpms.EntityData)
}

// NvSatellite_IcpeDpms_IcpeDpm
// ICPE DPM operational information
type NvSatellite_IcpeDpms_IcpeDpm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Discovery interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    DiscoveryInterface interface{}

    // Discovery interface. The type is string.
    DiscoveryInterfaceXr interface{}

    // Discovery interface handle. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (icpeDpm *NvSatellite_IcpeDpms_IcpeDpm) GetEntityData() *types.CommonEntityData {
    icpeDpm.EntityData.YFilter = icpeDpm.YFilter
    icpeDpm.EntityData.YangName = "icpe-dpm"
    icpeDpm.EntityData.BundleName = "cisco_ios_xr"
    icpeDpm.EntityData.ParentYangName = "icpe-dpms"
    icpeDpm.EntityData.SegmentPath = "icpe-dpm" + "[discovery-interface='" + fmt.Sprintf("%v", icpeDpm.DiscoveryInterface) + "']"
    icpeDpm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icpeDpm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icpeDpm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icpeDpm.EntityData.Children = make(map[string]types.YChild)
    icpeDpm.EntityData.Children["satellite"] = types.YChild{"Satellite", nil}
    for i := range icpeDpm.Satellite {
        icpeDpm.EntityData.Children[types.GetSegmentPath(&icpeDpm.Satellite[i])] = types.YChild{"Satellite", &icpeDpm.Satellite[i]}
    }
    icpeDpm.EntityData.Children["remote-host"] = types.YChild{"RemoteHost", nil}
    for i := range icpeDpm.RemoteHost {
        icpeDpm.EntityData.Children[types.GetSegmentPath(&icpeDpm.RemoteHost[i])] = types.YChild{"RemoteHost", &icpeDpm.RemoteHost[i]}
    }
    icpeDpm.EntityData.Leafs = make(map[string]types.YLeaf)
    icpeDpm.EntityData.Leafs["discovery-interface"] = types.YLeaf{"DiscoveryInterface", icpeDpm.DiscoveryInterface}
    icpeDpm.EntityData.Leafs["discovery-interface-xr"] = types.YLeaf{"DiscoveryInterfaceXr", icpeDpm.DiscoveryInterfaceXr}
    icpeDpm.EntityData.Leafs["discovery-interface-handle"] = types.YLeaf{"DiscoveryInterfaceHandle", icpeDpm.DiscoveryInterfaceHandle}
    icpeDpm.EntityData.Leafs["discovery-interface-status"] = types.YLeaf{"DiscoveryInterfaceStatus", icpeDpm.DiscoveryInterfaceStatus}
    icpeDpm.EntityData.Leafs["ident-packets-received"] = types.YLeaf{"IdentPacketsReceived", icpeDpm.IdentPacketsReceived}
    icpeDpm.EntityData.Leafs["ready-packets-received"] = types.YLeaf{"ReadyPacketsReceived", icpeDpm.ReadyPacketsReceived}
    icpeDpm.EntityData.Leafs["los-packets-received"] = types.YLeaf{"LosPacketsReceived", icpeDpm.LosPacketsReceived}
    icpeDpm.EntityData.Leafs["invalid-packets-received"] = types.YLeaf{"InvalidPacketsReceived", icpeDpm.InvalidPacketsReceived}
    icpeDpm.EntityData.Leafs["configuration-packets-sent"] = types.YLeaf{"ConfigurationPacketsSent", icpeDpm.ConfigurationPacketsSent}
    icpeDpm.EntityData.Leafs["ack-packets-sent"] = types.YLeaf{"AckPacketsSent", icpeDpm.AckPacketsSent}
    icpeDpm.EntityData.Leafs["reject-packets-sent"] = types.YLeaf{"RejectPacketsSent", icpeDpm.RejectPacketsSent}
    icpeDpm.EntityData.Leafs["probe-packets-sent"] = types.YLeaf{"ProbePacketsSent", icpeDpm.ProbePacketsSent}
    icpeDpm.EntityData.Leafs["host-ack-packets-received"] = types.YLeaf{"HostAckPacketsReceived", icpeDpm.HostAckPacketsReceived}
    icpeDpm.EntityData.Leafs["host-ack-packets-sent"] = types.YLeaf{"HostAckPacketsSent", icpeDpm.HostAckPacketsSent}
    icpeDpm.EntityData.Leafs["secs-since-pkts-cleaned"] = types.YLeaf{"SecsSincePktsCleaned", icpeDpm.SecsSincePktsCleaned}
    return &(icpeDpm.EntityData)
}

// NvSatellite_IcpeDpms_IcpeDpm_Satellite
// ICPE DPM satellite operational information table
type NvSatellite_IcpeDpms_IcpeDpm_Satellite struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteId interface{}

    // Satellite interface ID. The type is interface{} with range: 0..4294967295.
    SatelliteInterfaceId interface{}

    // Satellite interface MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SatelliteInterfaceMac interface{}

    // Satellite IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SatelliteIpAddress interface{}

    // Host IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HostIpAddress interface{}

    // Satellite chassis type. The type is string.
    SatelliteChassisType interface{}

    // Satellite chassis vendor. The type is string.
    SatelliteChassisVendor interface{}

    // Satellite chassis MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SatelliteChassisMac interface{}

    // Satellite serial id. The type is string.
    SatelliteSerialId interface{}

    // Satellite module type. The type is string.
    SatelliteModuleType interface{}

    // Satellite module vendor. The type is string.
    SatelliteModuleVendor interface{}

    // Satellite module MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SatelliteModuleMac interface{}

    // Conflict reason. The type is interface{} with range: 0..4294967295.
    ConflictReason interface{}

    // Received sys MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    ReceivedSysMac interface{}

    // Host chassis type. The type is string.
    HostChassisType interface{}

    // Host chassis vendor. The type is string.
    HostChassisVendor interface{}

    // Host chassis MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

func (satellite *NvSatellite_IcpeDpms_IcpeDpm_Satellite) GetEntityData() *types.CommonEntityData {
    satellite.EntityData.YFilter = satellite.YFilter
    satellite.EntityData.YangName = "satellite"
    satellite.EntityData.BundleName = "cisco_ios_xr"
    satellite.EntityData.ParentYangName = "icpe-dpm"
    satellite.EntityData.SegmentPath = "satellite"
    satellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satellite.EntityData.Children = make(map[string]types.YChild)
    satellite.EntityData.Leafs = make(map[string]types.YLeaf)
    satellite.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", satellite.SatelliteId}
    satellite.EntityData.Leafs["satellite-interface-id"] = types.YLeaf{"SatelliteInterfaceId", satellite.SatelliteInterfaceId}
    satellite.EntityData.Leafs["satellite-interface-mac"] = types.YLeaf{"SatelliteInterfaceMac", satellite.SatelliteInterfaceMac}
    satellite.EntityData.Leafs["satellite-ip-address"] = types.YLeaf{"SatelliteIpAddress", satellite.SatelliteIpAddress}
    satellite.EntityData.Leafs["host-ip-address"] = types.YLeaf{"HostIpAddress", satellite.HostIpAddress}
    satellite.EntityData.Leafs["satellite-chassis-type"] = types.YLeaf{"SatelliteChassisType", satellite.SatelliteChassisType}
    satellite.EntityData.Leafs["satellite-chassis-vendor"] = types.YLeaf{"SatelliteChassisVendor", satellite.SatelliteChassisVendor}
    satellite.EntityData.Leafs["satellite-chassis-mac"] = types.YLeaf{"SatelliteChassisMac", satellite.SatelliteChassisMac}
    satellite.EntityData.Leafs["satellite-serial-id"] = types.YLeaf{"SatelliteSerialId", satellite.SatelliteSerialId}
    satellite.EntityData.Leafs["satellite-module-type"] = types.YLeaf{"SatelliteModuleType", satellite.SatelliteModuleType}
    satellite.EntityData.Leafs["satellite-module-vendor"] = types.YLeaf{"SatelliteModuleVendor", satellite.SatelliteModuleVendor}
    satellite.EntityData.Leafs["satellite-module-mac"] = types.YLeaf{"SatelliteModuleMac", satellite.SatelliteModuleMac}
    satellite.EntityData.Leafs["conflict-reason"] = types.YLeaf{"ConflictReason", satellite.ConflictReason}
    satellite.EntityData.Leafs["received-sys-mac"] = types.YLeaf{"ReceivedSysMac", satellite.ReceivedSysMac}
    satellite.EntityData.Leafs["host-chassis-type"] = types.YLeaf{"HostChassisType", satellite.HostChassisType}
    satellite.EntityData.Leafs["host-chassis-vendor"] = types.YLeaf{"HostChassisVendor", satellite.HostChassisVendor}
    satellite.EntityData.Leafs["host-chassis-mac"] = types.YLeaf{"HostChassisMac", satellite.HostChassisMac}
    satellite.EntityData.Leafs["discovery-protocol-state"] = types.YLeaf{"DiscoveryProtocolState", satellite.DiscoveryProtocolState}
    satellite.EntityData.Leafs["last-imdr-state"] = types.YLeaf{"LastImdrState", satellite.LastImdrState}
    satellite.EntityData.Leafs["current-timeout"] = types.YLeaf{"CurrentTimeout", satellite.CurrentTimeout}
    satellite.EntityData.Leafs["is-queued-for-efd"] = types.YLeaf{"IsQueuedForEfd", satellite.IsQueuedForEfd}
    satellite.EntityData.Leafs["is-queued-for-oc"] = types.YLeaf{"IsQueuedForOc", satellite.IsQueuedForOc}
    satellite.EntityData.Leafs["ifmgr-state"] = types.YLeaf{"IfmgrState", satellite.IfmgrState}
    satellite.EntityData.Leafs["legacy"] = types.YLeaf{"Legacy", satellite.Legacy}
    satellite.EntityData.Leafs["deleting"] = types.YLeaf{"Deleting", satellite.Deleting}
    satellite.EntityData.Leafs["ident-packets-received"] = types.YLeaf{"IdentPacketsReceived", satellite.IdentPacketsReceived}
    satellite.EntityData.Leafs["ready-packets-received"] = types.YLeaf{"ReadyPacketsReceived", satellite.ReadyPacketsReceived}
    satellite.EntityData.Leafs["los-packets-received"] = types.YLeaf{"LosPacketsReceived", satellite.LosPacketsReceived}
    satellite.EntityData.Leafs["invalid-packets-received"] = types.YLeaf{"InvalidPacketsReceived", satellite.InvalidPacketsReceived}
    satellite.EntityData.Leafs["configuration-packets-sent"] = types.YLeaf{"ConfigurationPacketsSent", satellite.ConfigurationPacketsSent}
    satellite.EntityData.Leafs["ack-packets-sent"] = types.YLeaf{"AckPacketsSent", satellite.AckPacketsSent}
    satellite.EntityData.Leafs["reject-packets-sent"] = types.YLeaf{"RejectPacketsSent", satellite.RejectPacketsSent}
    satellite.EntityData.Leafs["secs-since-pkts-cleaned"] = types.YLeaf{"SecsSincePktsCleaned", satellite.SecsSincePktsCleaned}
    return &(satellite.EntityData)
}

// NvSatellite_IcpeDpms_IcpeDpm_RemoteHost
// ICPE DPM remote host operational information
// table
type NvSatellite_IcpeDpms_IcpeDpm_RemoteHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Host chassis MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    HostChassisMac interface{}

    // Host interface MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

func (remoteHost *NvSatellite_IcpeDpms_IcpeDpm_RemoteHost) GetEntityData() *types.CommonEntityData {
    remoteHost.EntityData.YFilter = remoteHost.YFilter
    remoteHost.EntityData.YangName = "remote-host"
    remoteHost.EntityData.BundleName = "cisco_ios_xr"
    remoteHost.EntityData.ParentYangName = "icpe-dpm"
    remoteHost.EntityData.SegmentPath = "remote-host"
    remoteHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteHost.EntityData.Children = make(map[string]types.YChild)
    remoteHost.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteHost.EntityData.Leafs["host-chassis-mac"] = types.YLeaf{"HostChassisMac", remoteHost.HostChassisMac}
    remoteHost.EntityData.Leafs["host-interface-mac"] = types.YLeaf{"HostInterfaceMac", remoteHost.HostInterfaceMac}
    remoteHost.EntityData.Leafs["discovery-protocol-state"] = types.YLeaf{"DiscoveryProtocolState", remoteHost.DiscoveryProtocolState}
    remoteHost.EntityData.Leafs["route-length"] = types.YLeaf{"RouteLength", remoteHost.RouteLength}
    remoteHost.EntityData.Leafs["current-timeout"] = types.YLeaf{"CurrentTimeout", remoteHost.CurrentTimeout}
    remoteHost.EntityData.Leafs["host-ack-packets-received"] = types.YLeaf{"HostAckPacketsReceived", remoteHost.HostAckPacketsReceived}
    remoteHost.EntityData.Leafs["host-ack-packets-sent"] = types.YLeaf{"HostAckPacketsSent", remoteHost.HostAckPacketsSent}
    remoteHost.EntityData.Leafs["secs-since-pkts-cleaned"] = types.YLeaf{"SecsSincePktsCleaned", remoteHost.SecsSincePktsCleaned}
    return &(remoteHost.EntityData)
}

// NvSatellite_SdacpControls
// SDAC Protocol Discovery information table
type NvSatellite_SdacpControls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SDAC Protocol Discovery information. The type is slice of
    // NvSatellite_SdacpControls_SdacpControl.
    SdacpControl []NvSatellite_SdacpControls_SdacpControl
}

func (sdacpControls *NvSatellite_SdacpControls) GetEntityData() *types.CommonEntityData {
    sdacpControls.EntityData.YFilter = sdacpControls.YFilter
    sdacpControls.EntityData.YangName = "sdacp-controls"
    sdacpControls.EntityData.BundleName = "cisco_ios_xr"
    sdacpControls.EntityData.ParentYangName = "nv-satellite"
    sdacpControls.EntityData.SegmentPath = "Cisco-IOS-XR-icpe-sdacp-oper:sdacp-controls"
    sdacpControls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdacpControls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdacpControls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdacpControls.EntityData.Children = make(map[string]types.YChild)
    sdacpControls.EntityData.Children["sdacp-control"] = types.YChild{"SdacpControl", nil}
    for i := range sdacpControls.SdacpControl {
        sdacpControls.EntityData.Children[types.GetSegmentPath(&sdacpControls.SdacpControl[i])] = types.YChild{"SdacpControl", &sdacpControls.SdacpControl[i]}
    }
    sdacpControls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sdacpControls.EntityData)
}

// NvSatellite_SdacpControls_SdacpControl
// SDAC Protocol Discovery information
type NvSatellite_SdacpControls_SdacpControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Satellite ID. The type is interface{} with range:
    // 100..65534.
    SatelliteId interface{}

    // Satellite ID. The type is interface{} with range: 0..65535.
    SatelliteIdXr interface{}

    // Satellite IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (sdacpControl *NvSatellite_SdacpControls_SdacpControl) GetEntityData() *types.CommonEntityData {
    sdacpControl.EntityData.YFilter = sdacpControl.YFilter
    sdacpControl.EntityData.YangName = "sdacp-control"
    sdacpControl.EntityData.BundleName = "cisco_ios_xr"
    sdacpControl.EntityData.ParentYangName = "sdacp-controls"
    sdacpControl.EntityData.SegmentPath = "sdacp-control" + "[satellite-id='" + fmt.Sprintf("%v", sdacpControl.SatelliteId) + "']"
    sdacpControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdacpControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdacpControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdacpControl.EntityData.Children = make(map[string]types.YChild)
    sdacpControl.EntityData.Children["protocol-state-timestamp"] = types.YChild{"ProtocolStateTimestamp", &sdacpControl.ProtocolStateTimestamp}
    sdacpControl.EntityData.Children["transport-error-timestamp"] = types.YChild{"TransportErrorTimestamp", &sdacpControl.TransportErrorTimestamp}
    sdacpControl.EntityData.Children["channel"] = types.YChild{"Channel", nil}
    for i := range sdacpControl.Channel {
        sdacpControl.EntityData.Children[types.GetSegmentPath(&sdacpControl.Channel[i])] = types.YChild{"Channel", &sdacpControl.Channel[i]}
    }
    sdacpControl.EntityData.Leafs = make(map[string]types.YLeaf)
    sdacpControl.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", sdacpControl.SatelliteId}
    sdacpControl.EntityData.Leafs["satellite-id-xr"] = types.YLeaf{"SatelliteIdXr", sdacpControl.SatelliteIdXr}
    sdacpControl.EntityData.Leafs["satellite-ip-address"] = types.YLeaf{"SatelliteIpAddress", sdacpControl.SatelliteIpAddress}
    sdacpControl.EntityData.Leafs["ip-address-auto"] = types.YLeaf{"IpAddressAuto", sdacpControl.IpAddressAuto}
    sdacpControl.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", sdacpControl.VrfName}
    sdacpControl.EntityData.Leafs["control-protocol-state"] = types.YLeaf{"ControlProtocolState", sdacpControl.ControlProtocolState}
    sdacpControl.EntityData.Leafs["transport-error"] = types.YLeaf{"TransportError", sdacpControl.TransportError}
    return &(sdacpControl.EntityData)
}

// NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp
// Timestamp
type NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (protocolStateTimestamp *NvSatellite_SdacpControls_SdacpControl_ProtocolStateTimestamp) GetEntityData() *types.CommonEntityData {
    protocolStateTimestamp.EntityData.YFilter = protocolStateTimestamp.YFilter
    protocolStateTimestamp.EntityData.YangName = "protocol-state-timestamp"
    protocolStateTimestamp.EntityData.BundleName = "cisco_ios_xr"
    protocolStateTimestamp.EntityData.ParentYangName = "sdacp-control"
    protocolStateTimestamp.EntityData.SegmentPath = "protocol-state-timestamp"
    protocolStateTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolStateTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolStateTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolStateTimestamp.EntityData.Children = make(map[string]types.YChild)
    protocolStateTimestamp.EntityData.Leafs = make(map[string]types.YLeaf)
    protocolStateTimestamp.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", protocolStateTimestamp.Seconds}
    protocolStateTimestamp.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", protocolStateTimestamp.Nanoseconds}
    return &(protocolStateTimestamp.EntityData)
}

// NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp
// Timestamp
type NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (transportErrorTimestamp *NvSatellite_SdacpControls_SdacpControl_TransportErrorTimestamp) GetEntityData() *types.CommonEntityData {
    transportErrorTimestamp.EntityData.YFilter = transportErrorTimestamp.YFilter
    transportErrorTimestamp.EntityData.YangName = "transport-error-timestamp"
    transportErrorTimestamp.EntityData.BundleName = "cisco_ios_xr"
    transportErrorTimestamp.EntityData.ParentYangName = "sdacp-control"
    transportErrorTimestamp.EntityData.SegmentPath = "transport-error-timestamp"
    transportErrorTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportErrorTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportErrorTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportErrorTimestamp.EntityData.Children = make(map[string]types.YChild)
    transportErrorTimestamp.EntityData.Leafs = make(map[string]types.YLeaf)
    transportErrorTimestamp.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", transportErrorTimestamp.Seconds}
    transportErrorTimestamp.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", transportErrorTimestamp.Nanoseconds}
    return &(transportErrorTimestamp.EntityData)
}

// NvSatellite_SdacpControls_SdacpControl_Channel
// Channels on satellite table
type NvSatellite_SdacpControls_SdacpControl_Channel struct {
    EntityData types.CommonEntityData
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

func (channel *NvSatellite_SdacpControls_SdacpControl_Channel) GetEntityData() *types.CommonEntityData {
    channel.EntityData.YFilter = channel.YFilter
    channel.EntityData.YangName = "channel"
    channel.EntityData.BundleName = "cisco_ios_xr"
    channel.EntityData.ParentYangName = "sdacp-control"
    channel.EntityData.SegmentPath = "channel"
    channel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channel.EntityData.Children = make(map[string]types.YChild)
    channel.EntityData.Children["capabilities"] = types.YChild{"Capabilities", &channel.Capabilities}
    channel.EntityData.Children["resync-state-timestamp"] = types.YChild{"ResyncStateTimestamp", &channel.ResyncStateTimestamp}
    channel.EntityData.Children["channel-state-timestamp"] = types.YChild{"ChannelStateTimestamp", &channel.ChannelStateTimestamp}
    channel.EntityData.Leafs = make(map[string]types.YLeaf)
    channel.EntityData.Leafs["channel-id"] = types.YLeaf{"ChannelId", channel.ChannelId}
    channel.EntityData.Leafs["resync-state"] = types.YLeaf{"ResyncState", channel.ResyncState}
    channel.EntityData.Leafs["channel-state"] = types.YLeaf{"ChannelState", channel.ChannelState}
    channel.EntityData.Leafs["control-messages-sent"] = types.YLeaf{"ControlMessagesSent", channel.ControlMessagesSent}
    channel.EntityData.Leafs["normal-messages-sent"] = types.YLeaf{"NormalMessagesSent", channel.NormalMessagesSent}
    channel.EntityData.Leafs["control-messages-received"] = types.YLeaf{"ControlMessagesReceived", channel.ControlMessagesReceived}
    channel.EntityData.Leafs["normal-messages-received"] = types.YLeaf{"NormalMessagesReceived", channel.NormalMessagesReceived}
    channel.EntityData.Leafs["control-messages-dropped"] = types.YLeaf{"ControlMessagesDropped", channel.ControlMessagesDropped}
    channel.EntityData.Leafs["normal-messages-dropped"] = types.YLeaf{"NormalMessagesDropped", channel.NormalMessagesDropped}
    channel.EntityData.Leafs["secs-since-last-cleared"] = types.YLeaf{"SecsSinceLastCleared", channel.SecsSinceLastCleared}
    channel.EntityData.Leafs["version"] = types.YLeaf{"Version", channel.Version}
    return &(channel.EntityData)
}

// NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities
// Capabilities
type NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Capability TLVs table. The type is slice of
    // NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs.
    TlVs []NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs
}

func (capabilities *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities) GetEntityData() *types.CommonEntityData {
    capabilities.EntityData.YFilter = capabilities.YFilter
    capabilities.EntityData.YangName = "capabilities"
    capabilities.EntityData.BundleName = "cisco_ios_xr"
    capabilities.EntityData.ParentYangName = "channel"
    capabilities.EntityData.SegmentPath = "capabilities"
    capabilities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capabilities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capabilities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capabilities.EntityData.Children = make(map[string]types.YChild)
    capabilities.EntityData.Children["tl-vs"] = types.YChild{"TlVs", nil}
    for i := range capabilities.TlVs {
        capabilities.EntityData.Children[types.GetSegmentPath(&capabilities.TlVs[i])] = types.YChild{"TlVs", &capabilities.TlVs[i]}
    }
    capabilities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(capabilities.EntityData)
}

// NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs
// Capability TLVs table
type NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is interface{} with range: 0..4294967295.
    Type_ interface{}

    // Debug. The type is string.
    Debug interface{}

    // Value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (tlVs *NvSatellite_SdacpControls_SdacpControl_Channel_Capabilities_TlVs) GetEntityData() *types.CommonEntityData {
    tlVs.EntityData.YFilter = tlVs.YFilter
    tlVs.EntityData.YangName = "tl-vs"
    tlVs.EntityData.BundleName = "cisco_ios_xr"
    tlVs.EntityData.ParentYangName = "capabilities"
    tlVs.EntityData.SegmentPath = "tl-vs"
    tlVs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tlVs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tlVs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tlVs.EntityData.Children = make(map[string]types.YChild)
    tlVs.EntityData.Leafs = make(map[string]types.YLeaf)
    tlVs.EntityData.Leafs["type"] = types.YLeaf{"Type_", tlVs.Type_}
    tlVs.EntityData.Leafs["debug"] = types.YLeaf{"Debug", tlVs.Debug}
    tlVs.EntityData.Leafs["value"] = types.YLeaf{"Value", tlVs.Value}
    return &(tlVs.EntityData)
}

// NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp
// Timestamp
type NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (resyncStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ResyncStateTimestamp) GetEntityData() *types.CommonEntityData {
    resyncStateTimestamp.EntityData.YFilter = resyncStateTimestamp.YFilter
    resyncStateTimestamp.EntityData.YangName = "resync-state-timestamp"
    resyncStateTimestamp.EntityData.BundleName = "cisco_ios_xr"
    resyncStateTimestamp.EntityData.ParentYangName = "channel"
    resyncStateTimestamp.EntityData.SegmentPath = "resync-state-timestamp"
    resyncStateTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resyncStateTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resyncStateTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resyncStateTimestamp.EntityData.Children = make(map[string]types.YChild)
    resyncStateTimestamp.EntityData.Leafs = make(map[string]types.YLeaf)
    resyncStateTimestamp.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", resyncStateTimestamp.Seconds}
    resyncStateTimestamp.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", resyncStateTimestamp.Nanoseconds}
    return &(resyncStateTimestamp.EntityData)
}

// NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp
// Timestamp
type NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (channelStateTimestamp *NvSatellite_SdacpControls_SdacpControl_Channel_ChannelStateTimestamp) GetEntityData() *types.CommonEntityData {
    channelStateTimestamp.EntityData.YFilter = channelStateTimestamp.YFilter
    channelStateTimestamp.EntityData.YangName = "channel-state-timestamp"
    channelStateTimestamp.EntityData.BundleName = "cisco_ios_xr"
    channelStateTimestamp.EntityData.ParentYangName = "channel"
    channelStateTimestamp.EntityData.SegmentPath = "channel-state-timestamp"
    channelStateTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelStateTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelStateTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelStateTimestamp.EntityData.Children = make(map[string]types.YChild)
    channelStateTimestamp.EntityData.Leafs = make(map[string]types.YLeaf)
    channelStateTimestamp.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", channelStateTimestamp.Seconds}
    channelStateTimestamp.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", channelStateTimestamp.Nanoseconds}
    return &(channelStateTimestamp.EntityData)
}

