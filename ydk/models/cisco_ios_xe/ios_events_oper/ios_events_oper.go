// This module contains a collection of YANG definitions
// for asynchronous events from network element.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package ios_events_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ios_events_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-ios-events-oper ios-events}", reflect.TypeOf(IosEvents{}))
    ydk.RegisterEntity("Cisco-IOS-XE-ios-events-oper:ios-events", reflect.TypeOf(IosEvents{}))
}

// HardwareSensorType represents Hardware Sensor Type
type HardwareSensorType string

const (
    // Hardware sensor board
    HardwareSensorType_hw_sensor_board HardwareSensorType = "hw-sensor-board"

    // Hardware sensor CPU junction
    HardwareSensorType_hw_sensor_cpu_junction HardwareSensorType = "hw-sensor-cpu-junction"

    // Hardware sensor DRAM
    HardwareSensorType_hw_sensor_dram HardwareSensorType = "hw-sensor-dram"

    // Hardware sensor PIM
    HardwareSensorType_hw_sensor_pim HardwareSensorType = "hw-sensor-pim"
)

// InterfaceNotifState represents Interface Notification state
type InterfaceNotifState string

const (
    InterfaceNotifState_interface_notif_state_up InterfaceNotifState = "interface-notif-state-up"

    InterfaceNotifState_interface_notif_state_down InterfaceNotifState = "interface-notif-state-down"
)

// NotificationModuleState represents Notification module state
type NotificationModuleState string

const (
    // Notification module inserted state
    NotificationModuleState_notf_module_state_inserted NotificationModuleState = "notf-module-state-inserted"

    // Notification module removed state
    NotificationModuleState_notf_module_state_removed NotificationModuleState = "notf-module-state-removed"
)

// UtdIpsAlertPriorityVal represents Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority
type UtdIpsAlertPriorityVal string

const (
    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is unknown
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_unknown UtdIpsAlertPriorityVal = "utd-ips-alert-priority-unknown"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is emergency
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_emerg UtdIpsAlertPriorityVal = "utd-ips-alert-priority-emerg"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is alert
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_alert UtdIpsAlertPriorityVal = "utd-ips-alert-priority-alert"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is critical
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_crit UtdIpsAlertPriorityVal = "utd-ips-alert-priority-crit"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is error
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_error UtdIpsAlertPriorityVal = "utd-ips-alert-priority-error"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is warning
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_warn UtdIpsAlertPriorityVal = "utd-ips-alert-priority-warn"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is notice
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_notice UtdIpsAlertPriorityVal = "utd-ips-alert-priority-notice"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is info
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_info UtdIpsAlertPriorityVal = "utd-ips-alert-priority-info"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert priority is debug
    UtdIpsAlertPriorityVal_utd_ips_alert_priority_debug UtdIpsAlertPriorityVal = "utd-ips-alert-priority-debug"
)

// OspfIntfState represents OSPF interface states
type OspfIntfState string

const (
    OspfIntfState_ospf_ifs_down OspfIntfState = "ospf-ifs-down"

    OspfIntfState_ospf_ifs_loopback OspfIntfState = "ospf-ifs-loopback"

    OspfIntfState_ospf_ifs_waiting OspfIntfState = "ospf-ifs-waiting"

    OspfIntfState_ospf_ifs_point_to_m_point OspfIntfState = "ospf-ifs-point-to-m-point"

    OspfIntfState_ospf_ifs_point_to_point OspfIntfState = "ospf-ifs-point-to-point"

    OspfIntfState_ospf_ifs_dr OspfIntfState = "ospf-ifs-dr"

    OspfIntfState_ospf_ifs_backup OspfIntfState = "ospf-ifs-backup"

    OspfIntfState_ospf_ifs_dr_other OspfIntfState = "ospf-ifs-dr-other"

    OspfIntfState_ospf_ifs_depend_upon OspfIntfState = "ospf-ifs-depend-upon"
)

// UtdIpsAlertActionVal represents Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert action
type UtdIpsAlertActionVal string

const (
    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert action is unknown
    UtdIpsAlertActionVal_utd_ips_alert_action_unknown UtdIpsAlertActionVal = "utd-ips-alert-action-unknown"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert action generated an alert
    UtdIpsAlertActionVal_utd_ips_alert_action_alert UtdIpsAlertActionVal = "utd-ips-alert-action-alert"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert action resulted in a drop
    UtdIpsAlertActionVal_utd_ips_alert_action_drop UtdIpsAlertActionVal = "utd-ips-alert-action-drop"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert action would have resulted in a drop if running in IPS mode
    UtdIpsAlertActionVal_utd_ips_alert_action_wdrop UtdIpsAlertActionVal = "utd-ips-alert-action-wdrop"
)

// BgpPstate represents GGP state
type BgpPstate string

const (
    BgpPstate_bgp_state_idle BgpPstate = "bgp-state-idle"

    BgpPstate_bgp_state_connect BgpPstate = "bgp-state-connect"

    BgpPstate_bgp_state_active BgpPstate = "bgp-state-active"

    BgpPstate_bgp_state_opensent BgpPstate = "bgp-state-opensent"

    BgpPstate_bgp_state_openconfirm BgpPstate = "bgp-state-openconfirm"

    BgpPstate_bgp_state_established BgpPstate = "bgp-state-established"

    BgpPstate_bgp_state_clearing BgpPstate = "bgp-state-clearing"

    BgpPstate_bgp_state_deleted BgpPstate = "bgp-state-deleted"
)

// DhcpServerStateVal represents DHCP Server state
type DhcpServerStateVal string

const (
    DhcpServerStateVal_dhcp_server_state_up DhcpServerStateVal = "dhcp-server-state-up"

    DhcpServerStateVal_dhcp_server_state_down DhcpServerStateVal = "dhcp-server-state-down"
)

// NotificationSeverity represents Notification severity
type NotificationSeverity string

const (
    NotificationSeverity_critical NotificationSeverity = "critical"

    NotificationSeverity_major NotificationSeverity = "major"

    NotificationSeverity_minor NotificationSeverity = "minor"
)

// VrrpIpAddrType represents IP Address type
type VrrpIpAddrType string

const (
    VrrpIpAddrType_vrrp_undefined VrrpIpAddrType = "vrrp-undefined"

    VrrpIpAddrType_vrrp_ipv4_address VrrpIpAddrType = "vrrp-ipv4-address"

    VrrpIpAddrType_vrrp_ipv6_address VrrpIpAddrType = "vrrp-ipv6-address"
)

// NotificationSensorState represents Notification sensor state
type NotificationSensorState string

const (
    // Sensor state green
    NotificationSensorState_sensor_state_green NotificationSensorState = "sensor-state-green"

    // Sensor state yellow
    NotificationSensorState_sensor_state_yellow NotificationSensorState = "sensor-state-yellow"

    // Sensor state red
    NotificationSensorState_sensor_state_red NotificationSensorState = "sensor-state-red"
)

// UtdUpdateTypeVal represents Unified Threat Defense (UTD) update type
type UtdUpdateTypeVal string

const (
    // Unified Threat Defense (UTD) update type is unknown
    UtdUpdateTypeVal_utd_update_type_unknown UtdUpdateTypeVal = "utd-update-type-unknown"

    // Unified Threat Defense (UTD) update is an IPS signature package update
    UtdUpdateTypeVal_utd_update_type_ips UtdUpdateTypeVal = "utd-update-type-ips"

    // Unified Threat Defense (UTD) update is a URL-Filtering DB update
    UtdUpdateTypeVal_utd_update_type_urlf UtdUpdateTypeVal = "utd-update-type-urlf"
)

// NotificationFailureState represents Notification failure state
type NotificationFailureState string

const (
    // Notification failure state ok
    NotificationFailureState_notf_failure_state_ok NotificationFailureState = "notf-failure-state-ok"

    // Notification failure state failed
    NotificationFailureState_notf_failure_state_failed NotificationFailureState = "notf-failure-state-failed"
)

// IntfAdminState represents Interface admin state
type IntfAdminState string

const (
    IntfAdminState_up IntfAdminState = "up"

    IntfAdminState_down IntfAdminState = "down"
)

// VrrpNotifProtoState represents VRRP protocol state
type VrrpNotifProtoState string

const (
    VrrpNotifProtoState_vrrp_notif_init VrrpNotifProtoState = "vrrp-notif-init"

    VrrpNotifProtoState_vrrp_notif_backup VrrpNotifProtoState = "vrrp-notif-backup"

    VrrpNotifProtoState_vrrp_notif_master VrrpNotifProtoState = "vrrp-notif-master"

    VrrpNotifProtoState_vrrp_notif_recover VrrpNotifProtoState = "vrrp-notif-recover"
)

// OspfNbrState represents OSPF neighbor states
type OspfNbrState string

const (
    OspfNbrState_ospf_nbr_down OspfNbrState = "ospf-nbr-down"

    OspfNbrState_ospf_nbr_attempt OspfNbrState = "ospf-nbr-attempt"

    OspfNbrState_ospf_nbr_init OspfNbrState = "ospf-nbr-init"

    OspfNbrState_ospf_nbr_two_way OspfNbrState = "ospf-nbr-two-way"

    OspfNbrState_ospf_nbr_exstart OspfNbrState = "ospf-nbr-exstart"

    OspfNbrState_ospf_nbr_exchange OspfNbrState = "ospf-nbr-exchange"

    OspfNbrState_ospf_nbr_loading OspfNbrState = "ospf-nbr-loading"

    OspfNbrState_ospf_nbr_full OspfNbrState = "ospf-nbr-full"

    OspfNbrState_ospf_nbr_deleted OspfNbrState = "ospf-nbr-deleted"

    OspfNbrState_ospf_nbr_depend_upon OspfNbrState = "ospf-nbr-depend-upon"
)

// UtdIpsAlertClassificationVal represents Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification
type UtdIpsAlertClassificationVal string

const (
    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is not set
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_none UtdIpsAlertClassificationVal = "utd-ips-alert-classification-none"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is not suspicious traffic
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_not_suspicious UtdIpsAlertClassificationVal = "utd-ips-alert-classification-not-suspicious"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is unknown traffic
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_unknown UtdIpsAlertClassificationVal = "utd-ips-alert-classification-unknown"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is potentially bad traffic
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_bad_unknown UtdIpsAlertClassificationVal = "utd-ips-alert-classification-bad-unknown"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is attempted information leak
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_attempted_recon UtdIpsAlertClassificationVal = "utd-ips-alert-classification-attempted-recon"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is information leak
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_successful_recon_limited UtdIpsAlertClassificationVal = "utd-ips-alert-classification-successful-recon-limited"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is large scale information leak
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_successful_recon_largescale UtdIpsAlertClassificationVal = "utd-ips-alert-classification-successful-recon-largescale"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is attempted denial of service
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_attempted_dos UtdIpsAlertClassificationVal = "utd-ips-alert-classification-attempted-dos"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is denial of service
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_successful_dos UtdIpsAlertClassificationVal = "utd-ips-alert-classification-successful-dos"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is attempted user privilege gain
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_attempted_user UtdIpsAlertClassificationVal = "utd-ips-alert-classification-attempted-user"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is unsuccessful user privilege gain
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_unsuccessful_user UtdIpsAlertClassificationVal = "utd-ips-alert-classification-unsuccessful-user"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is successful user privilege gain
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_successful_user UtdIpsAlertClassificationVal = "utd-ips-alert-classification-successful-user"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is attempted administrator privilege gain
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_attempted_admin UtdIpsAlertClassificationVal = "utd-ips-alert-classification-attempted-admin"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is successful administrator privilege gain
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_successful_admin UtdIpsAlertClassificationVal = "utd-ips-alert-classification-successful-admin"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is decode of an rpc query
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_rpc_portmap_decode UtdIpsAlertClassificationVal = "utd-ips-alert-classification-rpc-portmap-decode"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is executable code was detected
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_shellcode_detect UtdIpsAlertClassificationVal = "utd-ips-alert-classification-shellcode-detect"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is a suspicious string was detected
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_string_detect UtdIpsAlertClassificationVal = "utd-ips-alert-classification-string-detect"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is a suspicious filename was detected
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_suspicious_filename_detect UtdIpsAlertClassificationVal = "utd-ips-alert-classification-suspicious-filename-detect"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is an attempted login using a suspicious username was detected
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_suspicious_login UtdIpsAlertClassificationVal = "utd-ips-alert-classification-suspicious-login"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is a system call was detected
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_system_call_detect UtdIpsAlertClassificationVal = "utd-ips-alert-classification-system-call-detect"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is a tcp connection was detected
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_tcp_connection UtdIpsAlertClassificationVal = "utd-ips-alert-classification-tcp-connection"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is a network trojan was detected
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_trojan_activity UtdIpsAlertClassificationVal = "utd-ips-alert-classification-trojan-activity"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is a client was using an unusual port
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_unusual_client_port_connection UtdIpsAlertClassificationVal = "utd-ips-alert-classification-unusual-client-port-connection"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is detection of a network scan
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_network_scan UtdIpsAlertClassificationVal = "utd-ips-alert-classification-network-scan"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is detection of a denial of service attack
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_denial_of_service UtdIpsAlertClassificationVal = "utd-ips-alert-classification-denial-of-service"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is detection of a non-standard protocol or event
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_non_standard_protocol UtdIpsAlertClassificationVal = "utd-ips-alert-classification-non-standard-protocol"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is generic protocol command decode
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_protocol_command_decode UtdIpsAlertClassificationVal = "utd-ips-alert-classification-protocol-command-decode"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is access to a potentially vulnerable web application
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_web_application_activity UtdIpsAlertClassificationVal = "utd-ips-alert-classification-web-application-activity"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is web application attack
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_web_application_attack UtdIpsAlertClassificationVal = "utd-ips-alert-classification-web-application-attack"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is misc activity
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_misc_activity UtdIpsAlertClassificationVal = "utd-ips-alert-classification-misc-activity"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is misc attack
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_misc_attack UtdIpsAlertClassificationVal = "utd-ips-alert-classification-misc-attack"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is generic icmp event
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_icmp_event UtdIpsAlertClassificationVal = "utd-ips-alert-classification-icmp-event"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is inappropriate content was detected
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_inappropriate_content UtdIpsAlertClassificationVal = "utd-ips-alert-classification-inappropriate-content"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is potential corporate privacy violation
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_policy_violation UtdIpsAlertClassificationVal = "utd-ips-alert-classification-policy-violation"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is attempt to login by a default username and password
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_default_login_attempt UtdIpsAlertClassificationVal = "utd-ips-alert-classification-default-login-attempt"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is senstive data
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_sdf UtdIpsAlertClassificationVal = "utd-ips-alert-classification-sdf"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is known malicious file or file based exploit
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_file_format UtdIpsAlertClassificationVal = "utd-ips-alert-classification-file-format"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is known malware command and control traffic
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_malware_cnc UtdIpsAlertClassificationVal = "utd-ips-alert-classification-malware-cnc"

    // Unified Threat Defense (UTD) Intrusion Prevention System (IPS) alert classification is known client side exploit attempt
    UtdIpsAlertClassificationVal_utd_ips_alert_classification_client_side_exploit UtdIpsAlertClassificationVal = "utd-ips-alert-classification-client-side-exploit"
)

// FibUpdatesAfType represents FIB updates AF type
type FibUpdatesAfType string

const (
    FibUpdatesAfType_fib_updates_af_unknown FibUpdatesAfType = "fib-updates-af-unknown"

    FibUpdatesAfType_fib_updates_af_ipv4 FibUpdatesAfType = "fib-updates-af-ipv4"

    FibUpdatesAfType_fib_updates_af_ipv6 FibUpdatesAfType = "fib-updates-af-ipv6"
)

// IosEvents
// Events generated from IOS features
type IosEvents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (iosEvents *IosEvents) GetEntityData() *types.CommonEntityData {
    iosEvents.EntityData.YFilter = iosEvents.YFilter
    iosEvents.EntityData.YangName = "ios-events"
    iosEvents.EntityData.BundleName = "cisco_ios_xe"
    iosEvents.EntityData.ParentYangName = "Cisco-IOS-XE-ios-events-oper"
    iosEvents.EntityData.SegmentPath = "Cisco-IOS-XE-ios-events-oper:ios-events"
    iosEvents.EntityData.AbsolutePath = iosEvents.EntityData.SegmentPath
    iosEvents.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iosEvents.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iosEvents.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iosEvents.EntityData.Children = types.NewOrderedMap()
    iosEvents.EntityData.Leafs = types.NewOrderedMap()

    iosEvents.EntityData.YListKeys = []string {}

    return &(iosEvents.EntityData)
}

