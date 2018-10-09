// This module contains a collection of YANG definitions
// for cellwan operational data.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package cellwan_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cellwan_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-cellwan-oper cellwan-oper-data}", reflect.TypeOf(CellwanOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-cellwan-oper:cellwan-oper-data", reflect.TypeOf(CellwanOperData{}))
}

// ModemService represents Modem service indicates the current network service type for cellular modem
type ModemService string

const (
    // Cellular Network service type is circuit-switched
    ModemService_service_type_circuit_switched ModemService = "service-type-circuit-switched"

    // Cellular Network service type is packet-switched
    ModemService_service_type_packet_switched ModemService = "service-type-packet-switched"

    // Cellular Network service type is combined
    ModemService_service_type_combined ModemService = "service-type-combined"

    // Cellular Network service type is invalid
    ModemService_service_type_invalid ModemService = "service-type-invalid"

    // Cellular Network service type is unknown
    ModemService_service_type_unknown ModemService = "service-type-unknown"
)

// ServiceStatus represents Service status indicates the current network service for cellular modem
type ServiceStatus string

const (
    // Cellular Network status is in normal state
    ServiceStatus_service_status_normal ServiceStatus = "service-status-normal"

    // Cellular Network status is in emergency state
    ServiceStatus_service_status_emergency ServiceStatus = "service-status-emergency"

    // Cellular Network status is in no service state
    ServiceStatus_service_status_no_service ServiceStatus = "service-status-no-service"

    // Cellular Network status is in unknown state
    ServiceStatus_service_status_unknown ServiceStatus = "service-status-unknown"
)

// RadioBandwidth represents Radio bandwidth indicates the current cellular radio bandwidth selected in MHz
type RadioBandwidth string

const (
    // Cellular radio bandwidth is 1.4 MHz
    RadioBandwidth_bandwidth_1_dot_4_mhz RadioBandwidth = "bandwidth-1-dot-4-mhz"

    // Cellular radio bandwidth is 3 MHz
    RadioBandwidth_bandwidth_3_mhz RadioBandwidth = "bandwidth-3-mhz"

    // Cellular radio bandwidth is 5 MHz
    RadioBandwidth_bandwidth_5_mhz RadioBandwidth = "bandwidth-5-mhz"

    // Cellular radio bandwidth is 10 MHz
    RadioBandwidth_bandwidth_10_mhz RadioBandwidth = "bandwidth-10-mhz"

    // Cellular radio bandwidth is 15 MHz
    RadioBandwidth_bandwidth_15_mhz RadioBandwidth = "bandwidth-15-mhz"

    // Cellular radio bandwidth is 20 MHz
    RadioBandwidth_bandwidth_20_mhz RadioBandwidth = "bandwidth-20-mhz"

    // Cellular radio bandwidth is invalid
    RadioBandwidth_bandwidth_invalid RadioBandwidth = "bandwidth-invalid"

    // Cellular radio bandwidth is unknown
    RadioBandwidth_bandwidth_unknown RadioBandwidth = "bandwidth-unknown"
)

// RatTechnology represents RAT technology indicates the current radio technology selected
type RatTechnology string

const (
    // Radio technology selected is none
    RatTechnology_system_mode_none RatTechnology = "system-mode-none"

    // Radio technology selected is GPRS (General Packet Radio Service)
    RatTechnology_system_mode_gprs RatTechnology = "system-mode-gprs"

    // Radio technology selected is EDGE (Enhanced Data rates for GSM Evolution)
    RatTechnology_system_mode_edge RatTechnology = "system-mode-edge"

    // Radio technology selected is UMTS (Universal Mobile Telecommunications System)
    RatTechnology_system_mode_umts RatTechnology = "system-mode-umts"

    // Radio technology selected is HSDPA (High Speed Downlink Packet Access)
    RatTechnology_system_mode_hsdpa RatTechnology = "system-mode-hsdpa"

    // Radio technology selected is HSUPA (High Speed Uplink Packet Access)
    RatTechnology_system_mode_hsupa RatTechnology = "system-mode-hsupa"

    // Radio technology selected is HSPA (High Speed Packet Access)
    RatTechnology_system_mode_hspa RatTechnology = "system-mode-hspa"

    // Radio technology selected is HSPA+ (Evolved High Speed Packet Access)
    RatTechnology_system_mode_hspa_plus RatTechnology = "system-mode-hspa-plus"

    // Radio technology selected is LTE-FDD (Long Term Evolution-Frequency Division Duplex)
    RatTechnology_system_mode_lte_fdd RatTechnology = "system-mode-lte-fdd"

    // Radio technology selected is LTE-TDD(Long Term Evolution-Time Division Duplex)
    RatTechnology_system_mode_lte_tdd RatTechnology = "system-mode-lte-tdd"

    // Radio technology selected is LTE (Long Term Evolution) /
    //   eHRPD (Evolved High Rate Packet Data) /
    //   1xRTT (Single carrier (1x) radio transmission technology)
    RatTechnology_system_mode_lte_e_hrpd_1x_rtt RatTechnology = "system-mode-lte-e-hrpd-1x-rtt"

    // Radio technology selected is LTE (Long Term Evolution) / 
    //   eHRPD (Evolved High Rate Packet Data) / 
    //   EVDO (Evolution-Data Optimized) 
    RatTechnology_system_mode_lte_e_hrpd_evdo RatTechnology = "system-mode-lte-e-hrpd-evdo"

    // Radio technology selected is EVDO (Evolution-Data Optimized)
    RatTechnology_system_mode_evdo RatTechnology = "system-mode-evdo"

    // Radio technology selected is EVDO (Evolution-Data Optimized) / 
    //   REVA (Evolution Data Optimized - Rev A)
    RatTechnology_system_mode_evdo_reva RatTechnology = "system-mode-evdo-reva"

    // Radio technology selected is HSDPA (High Speed Downlink Packet Access) 
    //   & WCDMA (Wideband Code Division Multiple Access)
    RatTechnology_system_mode_hsdpa_n_wcdma RatTechnology = "system-mode-hsdpa-n-wcdma"

    // Radio technology selected is WCDMA (Wideband Code Division Multiple Access)
    //   & HSUPA (High Speed Uplink Packet Access) 
    RatTechnology_system_mode_wcdma_n_hsupa RatTechnology = "system-mode-wcdma-n-hsupa"

    // Radio technology selected is HSDPA (High Speed Downlink Packet Access) 
    //   & HSUPA (High Speed Uplink Packet Access) 
    RatTechnology_system_mode_hsdpa_n_hsupa RatTechnology = "system-mode-hsdpa-n-hsupa"

    // Radio technology selected is HSDPA+ (High Speed Downlink Packet Access plus) 
    //   & WCDMA (Wideband Code Division Multiple Access)
    RatTechnology_system_mode_hsdpa_plus_n_wcdma RatTechnology = "system-mode-hsdpa-plus-n-wcdma"

    // Radio technology selected is HSDPA+ (High Speed Downlink Packet Access plus) 
    //   & HSUPA (High Speed Uplink Packet Access) 
    RatTechnology_system_mode_hsdpa_plus_n_hsupa RatTechnology = "system-mode-hsdpa-plus-n-hsupa"

    // Radio technology selected is 
    //   DC HSDPA+ (Dual Carrier High Speed Downlink Packet Access plus) 
    //   & WCDMA (Wideband Code Division Multiple Access)
    RatTechnology_system_mode_dc_hsdpa_plus_n_wcdma RatTechnology = "system-mode-dc-hsdpa-plus-n-wcdma"

    // Radio technology selected is 
    //   DC HSDPA+ (Dual Carrier High Speed Downlink Packet Access plus)
    //   & HSUPA (High Speed Uplink Packet Access)
    RatTechnology_system_mode_dc_hsdpa_plus_n_hsupa RatTechnology = "system-mode-dc-hsdpa-plus-n-hsupa"

    // Radio technology selected is null bearer
    RatTechnology_sysyem_mode_null_bearer RatTechnology = "sysyem-mode-null-bearer"

    // Radio technology selected is unknown
    RatTechnology_system_mode_unknown RatTechnology = "system-mode-unknown"
)

// RatPreference represents RAT preference indicates the current radio technology preference
type RatPreference string

const (
    // Radio technology preference is no svc
    RatPreference_lte_radio_tech_no_svc RatPreference = "lte-radio-tech-no-svc"

    // Radio technology preference is CDMA (Code Division Multiple Access) 
    // / 1xRTT (Single carrier (1x) radio transmission technology)
    RatPreference_lte_radio_tech_cdma_1_xrtt RatPreference = "lte-radio-tech-cdma-1-xrtt"

    // Radio technology preference is CDMA (Code Division Multiple Access) 
    // / EVDO (Evolution-Data Optimized) 
    RatPreference_lte_radio_tech_cdma_evdo RatPreference = "lte-radio-tech-cdma-evdo"

    // Radio technology preference is AMPS (Advanced Mobile Phone Service)
    RatPreference_lte_radio_tech_amps RatPreference = "lte-radio-tech-amps"

    // Radio technology preference is GSM (Global System for Mobile Communications)
    RatPreference_lte_radio_tech_gsm RatPreference = "lte-radio-tech-gsm"

    // Radio technology preference is UMTS (Universal Mobile Telecommunications Service)
    RatPreference_lte_radio_tech_umts RatPreference = "lte-radio-tech-umts"

    // Radio technology preference is WLAN (wireless local area network)
    RatPreference_lte_radio_tech_wlan RatPreference = "lte-radio-tech-wlan"

    // Radio technology preference is GPRS (General Packet Radio Service)
    RatPreference_lte_radio_tech_gprs RatPreference = "lte-radio-tech-gprs"

    // Radio technology preference is LTE (Long-Term Evolution)
    RatPreference_lte_radio_tech_lte RatPreference = "lte-radio-tech-lte"

    // Radio technology preference is AUTO 
    RatPreference_lte_radio_tech_auto RatPreference = "lte-radio-tech-auto"

    // Radio technology preference is Hybrid CDMA (Hybrid Code Division Multiple Access)
    RatPreference_lte_radio_tech_hybrid_cdma RatPreference = "lte-radio-tech-hybrid-cdma"

    // Radio technology preference is WCDMA (Wideband Code Division Multiplexing Access)
    RatPreference_lte_radio_tech_wcdma RatPreference = "lte-radio-tech-wcdma"

    // Radio technology preference is GWL 
    RatPreference_lte_radio_tech_gwl RatPreference = "lte-radio-tech-gwl"

    // Radio technology preference is EDGE (Enhanced Data rates for GSM Evolution)
    RatPreference_lte_radio_tech_edge RatPreference = "lte-radio-tech-edge"

    // Radio technology preference is HSDPA (High Speed Downlink Packet Access)
    // & WCDMA (Wideband Code Division Multiplexing Access)
    RatPreference_lte_radio_tech_hsdpa_n_wcdma RatPreference = "lte-radio-tech-hsdpa-n-wcdma"

    // Radio technology preference is WCDMA (Wideband Code Division Multiplexing Access)
    //   & HSUPA (High Speed Uplink Packet Access)
    RatPreference_lte_radio_tech_wcdma_n_hsupa RatPreference = "lte-radio-tech-wcdma-n-hsupa"

    // Radio technology preference is HSDPA (High Speed Downlink Packet Access)
    //   & HSUPA (High Speed Uplink Packet Access)
    RatPreference_lte_radio_tech_hsdpa_n_hsupa RatPreference = "lte-radio-tech-hsdpa-n-hsupa"

    // Radio technology preference is HSDPA+ (High Speed Downlink Packet Access plus) 
    //   & WCDMA (Wideband Code Division Multiplexing Access)
    RatPreference_lte_radio_tech_hsdpa_plus_n_wcdma RatPreference = "lte-radio-tech-hsdpa-plus-n-wcdma"

    // Radio technology preference is HSDPA+ (High Speed Downlink Packet Access plus)
    //   & HSUPA (High Speed Uplink Packet Access)
    RatPreference_lte_radio_tech_hsdpa_plus_n_hsupa RatPreference = "lte-radio-tech-hsdpa-plus-n-hsupa"

    // Radio technology preference is 
    //   DC HSDPA+ (Dual Carrier High Speed Downlink Packet Access plus) 
    //   & WCDMA (Wideband Code Division Multiplexing Access)
    RatPreference_lte_radio_tech_dc_hsdpa_plus_n_wcdma RatPreference = "lte-radio-tech-dc-hsdpa-plus-n-wcdma"

    // Radio technology preference is 
    //   DC HSDPA+ (Dual Carrier High Speed Downlink Packet Access plus) 
    //   & HSUPA (High Speed Uplink Packet Access)
    RatPreference_lte_radio_tech_dc_hsdpa_plus_n_hsupa RatPreference = "lte-radio-tech-dc-hsdpa-plus-n-hsupa"

    // Radio technology preference is Null Bearer 
    RatPreference_lte_radio_tech_null_bearer RatPreference = "lte-radio-tech-null-bearer"

    // Radio technology preference is Unknown 
    RatPreference_lte_radio_tech_unknown RatPreference = "lte-radio-tech-unknown"

    // Radio technology preference is unchanged 
    RatPreference_lte_radio_tech_no_change RatPreference = "lte-radio-tech-no-change"
)

// CwanGpsModeSelected represents Cellular modem GPS mode selection status
type CwanGpsModeSelected string

const (
    // Cellular modem GPS mode is disabled
    CwanGpsModeSelected_gps_mode_disable CwanGpsModeSelected = "gps-mode-disable"

    // Cellular modem GPS mode is standalone
    CwanGpsModeSelected_gps_mode_standalone CwanGpsModeSelected = "gps-mode-standalone"

    // Cellular modem GPS mode is ms-based
    CwanGpsModeSelected_gps_mode_mbased CwanGpsModeSelected = "gps-mode-mbased"

    // Cellular modem GPS mode is ms-assist
    CwanGpsModeSelected_gps_mode_msassist CwanGpsModeSelected = "gps-mode-msassist"
)

// CwanGpsFeatureState represents Cellular modem GPS feature state status
type CwanGpsFeatureState string

const (
    // Cellular modem GPS feature state is Disabled
    CwanGpsFeatureState_gps_disabled CwanGpsFeatureState = "gps-disabled"

    // Cellular modem GPS feature state is Enabled
    CwanGpsFeatureState_gps_enabled CwanGpsFeatureState = "gps-enabled"
)

// PacketSessStatus represents Packet Session Status indicates the Cellular packet session state
type PacketSessStatus string

const (
    // Cellular packet session status is inactive
    PacketSessStatus_packet_session_status_inactive PacketSessStatus = "packet-session-status-inactive"

    // Cellular packet session status is active
    PacketSessStatus_packet_session_status_active PacketSessStatus = "packet-session-status-active"
)

// CwanGpsPortSelected represents Cellular modem GPS port selection status
type CwanGpsPortSelected string

const (
    // Cellular modem dedicated GPS port selected
    CwanGpsPortSelected_dedicated_gps_port CwanGpsPortSelected = "dedicated-gps-port"

    // Cellular modem DIV port selected for GPS
    CwanGpsPortSelected_div_gps_port CwanGpsPortSelected = "div-gps-port"

    // Cellular modem voltage no-bias port selected for GPS
    CwanGpsPortSelected_voltage_no_bias_gps_port CwanGpsPortSelected = "voltage-no-bias-gps-port"

    // Cellular modem none port selected for GPS
    CwanGpsPortSelected_gps_port_none CwanGpsPortSelected = "gps-port-none"
)

// CwRadioPowerStatus represents Radio power status indicates the current radio power mode of cellular modem
type CwRadioPowerStatus string

const (
    // Cellular modem is in online radio mode
    CwRadioPowerStatus_radio_power_mode_online CwRadioPowerStatus = "radio-power-mode-online"

    // Cellular modem is in low power radio mode
    CwRadioPowerStatus_radio_power_mode_low_power CwRadioPowerStatus = "radio-power-mode-low-power"

    // Cellular modem is in factory test radio mode
    CwRadioPowerStatus_radio_power_mode_factory_test CwRadioPowerStatus = "radio-power-mode-factory-test"

    // Cellular modem is in offline radio mode
    CwRadioPowerStatus_radio_power_mode_offline CwRadioPowerStatus = "radio-power-mode-offline"

    // Cellular modem is in reset radio mode
    CwRadioPowerStatus_radio_power_mode_reset CwRadioPowerStatus = "radio-power-mode-reset"

    // Cellular modem is in off radio mode
    CwRadioPowerStatus_radio_power_mode_off CwRadioPowerStatus = "radio-power-mode-off"

    // Cellular modem is in persistent low power radio mode
    CwRadioPowerStatus_radio_power_mode_persistent_low_power CwRadioPowerStatus = "radio-power-mode-persistent-low-power"
)

// CellwanSimUserOp represents Cellular modem SIM user operation status
type CellwanSimUserOp string

const (
    // Cellular modem SIM user operation is in None state
    CellwanSimUserOp_sim_user_op_none CellwanSimUserOp = "sim-user-op-none"

    // Cellular modem SIM user is in CHV1 (Card Holder Verification) state
    CellwanSimUserOp_sim_user_op_chv1 CellwanSimUserOp = "sim-user-op-chv1"

    // Cellular modem SIM user is in CHV2 (Card Holder Verification) state
    CellwanSimUserOp_sim_user_op_chv2 CellwanSimUserOp = "sim-user-op-chv2"

    // Cellular modem SIM user is in Unblocked CHV1 (Card Holder Verification) state
    CellwanSimUserOp_sim_user_op_unblock_chv1 CellwanSimUserOp = "sim-user-op-unblock-chv1"

    // Cellular modem SIM user is in Unblocked CHV2 (Card Holder Verification) state
    CellwanSimUserOp_sim_user_op_unblock_chv2 CellwanSimUserOp = "sim-user-op-unblock-chv2"

    // Cellular modem SIM user is in MEP (Mobile Equipment Personalization) state
    CellwanSimUserOp_sim_user_op_mep CellwanSimUserOp = "sim-user-op-mep"

    // Cellular modem SIM user operation is in Unknown state
    CellwanSimUserOp_sim_user_op_unknown CellwanSimUserOp = "sim-user-op-unknown"
)

// ModemTechnology represents Modem technology indicates the current cellular technology selected
type ModemTechnology string

const (
    // Modem technology selected is CDMA (Code Division Multiple Access) / 
    //   EVDO (Evolution-Data Optimized) / 
    //   1xRTT (Single carrier (1x) radio transmission technology)
    ModemTechnology_cdma_evdo_1x_rtt ModemTechnology = "cdma-evdo-1x-rtt"

    // Modem technology selected is GSM (Global System for Mobile Communications) 
    //   / UMTS (Universal Mobile Telecommunications Service) 
    //   / GPRS (General Packet Radio Service)
    ModemTechnology_gsm_umts_gprs ModemTechnology = "gsm-umts-gprs"

    // Modem technology selected is unknown
    ModemTechnology_tech_unknown ModemTechnology = "tech-unknown"
)

// ProfileScope represents Profile Scope indicates the IP address scope
type ProfileScope string

const (
    // IP address scope is Global
    ProfileScope_scope_global ProfileScope = "scope-global"

    // IP address scope is Link
    ProfileScope_scope_link ProfileScope = "scope-link"
)

// LteCa represents LTE CA indicates the LTE carrier aggregation status for cellular modem
type LteCa string

const (
    // LTE carrier aggregation is deconfigured
    LteCa_lte_ca_deconfigured LteCa = "lte-ca-deconfigured"

    // LTE carrier aggregation is deactivated
    LteCa_lte_ca_deactivated LteCa = "lte-ca-deactivated"

    // LTE carrier aggregation is activated
    LteCa_lte_ca_activated LteCa = "lte-ca-activated"

    // LTE carrier aggregation is unsupported
    LteCa_lte_ca_unsupported LteCa = "lte-ca-unsupported"
)

// ModemStatus represents Modem status indicates the current state of cellular modem
type ModemStatus string

const (
    // Modem is in Offline state
    ModemStatus_modem_status_offline ModemStatus = "modem-status-offline"

    // Modem is in Online state
    ModemStatus_modem_status_online ModemStatus = "modem-status-online"

    // Modem is in Low Power Mode state
    ModemStatus_modem_status_low_power ModemStatus = "modem-status-low-power"

    // Modem is in power off state
    ModemStatus_modem_status_power_off ModemStatus = "modem-status-power-off"

    // Modem is in boot ready state
    ModemStatus_modem_status_boot_ready ModemStatus = "modem-status-boot-ready"

    // Modem is in unknown state
    ModemStatus_modem_status_unknown ModemStatus = "modem-status-unknown"
)

// CwanGpsState represents Cellular modem GPS state
type CwanGpsState string

const (
    // Cellular modem is in GPS disabled state
    CwanGpsState_gps_state_disabled CwanGpsState = "gps-state-disabled"

    // Cellular modem is in GPS acquiring state
    CwanGpsState_gps_state_acquiring CwanGpsState = "gps-state-acquiring"

    // Cellular modem is in GPS enabled state
    CwanGpsState_gps_state_enabled CwanGpsState = "gps-state-enabled"

    // Cellular modem is in GPS location error state
    CwanGpsState_gps_loc_error CwanGpsState = "gps-loc-error"
)

// RegState represents Registration state indicates the current cellular network registration state
type RegState string

const (
    // Cellular Network is in not registered state
    RegState_reg_status_not_registered RegState = "reg-status-not-registered"

    // Cellular Network is in registered state
    RegState_reg_status_registered RegState = "reg-status-registered"

    // Cellular Network is in searching state
    RegState_reg_status_searching RegState = "reg-status-searching"

    // Cellular Network is in registration denied state
    RegState_reg_status_registration_denied RegState = "reg-status-registration-denied"

    // Cellular Network is in unsupported state
    RegState_reg_status_unsupported RegState = "reg-status-unsupported"
)

// CellwanSimStatus represents Cellular modem SIM status
type CellwanSimStatus string

const (
    // Cellular modem SIM status is inserted
    CellwanSimStatus_sim_status_ok CellwanSimStatus = "sim-status-ok"

    // Cellular modem SIM is not inserted
    CellwanSimStatus_sim_status_not_inserted CellwanSimStatus = "sim-status-not-inserted"

    // Cellular modem SIM is in removed state
    CellwanSimStatus_sim_status_removed CellwanSimStatus = "sim-status-removed"

    // Cellular modem SIM is in initialization failure state
    CellwanSimStatus_sim_status_init_failure CellwanSimStatus = "sim-status-init-failure"

    // Cellular modem SIM is in general failure state
    CellwanSimStatus_sim_status_general_failure CellwanSimStatus = "sim-status-general-failure"

    // Cellular modem SIM is in locked state
    CellwanSimStatus_sim_status_locked CellwanSimStatus = "sim-status-locked"

    // Cellular modem SIM is in chv1 (Card Holder Verification) blocked state
    CellwanSimStatus_sim_status_chv1_blocked CellwanSimStatus = "sim-status-chv1-blocked"

    // Cellular modem SIM is in chv2 (Card Holder Verification) blocked state
    CellwanSimStatus_sim_status_chv2_blocked CellwanSimStatus = "sim-status-chv2-blocked"

    // Cellular modem SIM is in chv1 (Card Holder Verification) rejected state
    CellwanSimStatus_sim_status_chv1_rejected CellwanSimStatus = "sim-status-chv1-rejected"

    // Cellular modem SIM is in chv2 (Card Holder Verification) rejected state
    CellwanSimStatus_sim_status_chv2_rejected CellwanSimStatus = "sim-status-chv2-rejected"

    // Cellular modem SIM is in MEP (Mobile Equipment Personalization) locked state
    CellwanSimStatus_sim_status_mep_locked CellwanSimStatus = "sim-status-mep-locked"

    // Cellular modem SIM is in network reject state
    CellwanSimStatus_sim_status_network_reject CellwanSimStatus = "sim-status-network-reject"

    // Cellular modem SIM is in unknown state
    CellwanSimStatus_sim_status_unknown CellwanSimStatus = "sim-status-unknown"
)

// CellwanChv1SimStatus represents Cellular modem SIM card holder verification (CHV1) status
type CellwanChv1SimStatus string

const (
    // SIM card holder verification (CHV1) disabled
    CellwanChv1SimStatus_chv1_verify_disabled CellwanChv1SimStatus = "chv1-verify-disabled"

    // SIM card holder verification (CHV1) enabled
    CellwanChv1SimStatus_chv1_verify_enabled CellwanChv1SimStatus = "chv1-verify-enabled"

    // SIM card holder verification (CHV1) pending
    CellwanChv1SimStatus_chv1_verify_pending CellwanChv1SimStatus = "chv1-verify-pending"
)

// CellwanOperData
// Cellwan operational data
type CellwanOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Cellwan modem hardware data. The type is slice of
    // CellwanOperData_CellwanHardware.
    CellwanHardware []*CellwanOperData_CellwanHardware

    // Cellwan modem radio data. The type is slice of
    // CellwanOperData_CellwanRadio.
    CellwanRadio []*CellwanOperData_CellwanRadio

    // Cellwan modem network data. The type is slice of
    // CellwanOperData_CellwanNetwork.
    CellwanNetwork []*CellwanOperData_CellwanNetwork

    // Cellwan modem connection data. The type is slice of
    // CellwanOperData_CellwanConnection.
    CellwanConnection []*CellwanOperData_CellwanConnection

    // Cellwan modem sim security data. The type is slice of
    // CellwanOperData_CellwanSecurity.
    CellwanSecurity []*CellwanOperData_CellwanSecurity

    // Cellwan modem sms data. The type is slice of CellwanOperData_CellwanSms.
    CellwanSms []*CellwanOperData_CellwanSms

    // Cellwan modem gps data. The type is slice of CellwanOperData_CellwanGps.
    CellwanGps []*CellwanOperData_CellwanGps
}

func (cellwanOperData *CellwanOperData) GetEntityData() *types.CommonEntityData {
    cellwanOperData.EntityData.YFilter = cellwanOperData.YFilter
    cellwanOperData.EntityData.YangName = "cellwan-oper-data"
    cellwanOperData.EntityData.BundleName = "cisco_ios_xe"
    cellwanOperData.EntityData.ParentYangName = "Cisco-IOS-XE-cellwan-oper"
    cellwanOperData.EntityData.SegmentPath = "Cisco-IOS-XE-cellwan-oper:cellwan-oper-data"
    cellwanOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanOperData.EntityData.Children = types.NewOrderedMap()
    cellwanOperData.EntityData.Children.Append("cellwan-hardware", types.YChild{"CellwanHardware", nil})
    for i := range cellwanOperData.CellwanHardware {
        cellwanOperData.EntityData.Children.Append(types.GetSegmentPath(cellwanOperData.CellwanHardware[i]), types.YChild{"CellwanHardware", cellwanOperData.CellwanHardware[i]})
    }
    cellwanOperData.EntityData.Children.Append("cellwan-radio", types.YChild{"CellwanRadio", nil})
    for i := range cellwanOperData.CellwanRadio {
        cellwanOperData.EntityData.Children.Append(types.GetSegmentPath(cellwanOperData.CellwanRadio[i]), types.YChild{"CellwanRadio", cellwanOperData.CellwanRadio[i]})
    }
    cellwanOperData.EntityData.Children.Append("cellwan-network", types.YChild{"CellwanNetwork", nil})
    for i := range cellwanOperData.CellwanNetwork {
        cellwanOperData.EntityData.Children.Append(types.GetSegmentPath(cellwanOperData.CellwanNetwork[i]), types.YChild{"CellwanNetwork", cellwanOperData.CellwanNetwork[i]})
    }
    cellwanOperData.EntityData.Children.Append("cellwan-connection", types.YChild{"CellwanConnection", nil})
    for i := range cellwanOperData.CellwanConnection {
        cellwanOperData.EntityData.Children.Append(types.GetSegmentPath(cellwanOperData.CellwanConnection[i]), types.YChild{"CellwanConnection", cellwanOperData.CellwanConnection[i]})
    }
    cellwanOperData.EntityData.Children.Append("cellwan-security", types.YChild{"CellwanSecurity", nil})
    for i := range cellwanOperData.CellwanSecurity {
        cellwanOperData.EntityData.Children.Append(types.GetSegmentPath(cellwanOperData.CellwanSecurity[i]), types.YChild{"CellwanSecurity", cellwanOperData.CellwanSecurity[i]})
    }
    cellwanOperData.EntityData.Children.Append("cellwan-sms", types.YChild{"CellwanSms", nil})
    for i := range cellwanOperData.CellwanSms {
        cellwanOperData.EntityData.Children.Append(types.GetSegmentPath(cellwanOperData.CellwanSms[i]), types.YChild{"CellwanSms", cellwanOperData.CellwanSms[i]})
    }
    cellwanOperData.EntityData.Children.Append("cellwan-gps", types.YChild{"CellwanGps", nil})
    for i := range cellwanOperData.CellwanGps {
        cellwanOperData.EntityData.Children.Append(types.GetSegmentPath(cellwanOperData.CellwanGps[i]), types.YChild{"CellwanGps", cellwanOperData.CellwanGps[i]})
    }
    cellwanOperData.EntityData.Leafs = types.NewOrderedMap()

    cellwanOperData.EntityData.YListKeys = []string {}

    return &(cellwanOperData.EntityData)
}

// CellwanOperData_CellwanHardware
// Cellwan modem hardware data
type CellwanOperData_CellwanHardware struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Cellular interface. The type is string.
    CellularInterface interface{}

    // Cellular Modem firmware version. The type is string.
    CellularFirmwareVersion interface{}

    // Cellular Modem firmware build time. The type is string.
    CellularFirmwareBuildTime interface{}

    // Cellular Modem hardware version. The type is string.
    CellularHardwareVersion interface{}

    // Cellular Modem device model type. The type is string.
    CellularDeviceModel interface{}

    // Cellular Modem IMSI. The type is string.
    CellularImsi interface{}

    // Cellular Modem IMEI. The type is string.
    CellularImei interface{}

    // Cellular Modem ICCID. The type is string.
    CellularIccid interface{}

    // Cellular Modem MSISDN. The type is string.
    CellularMsisdn interface{}

    // Cellular Modem FSN. The type is string.
    CellularFsn interface{}

    // Cellular Modem Status. The type is ModemStatus.
    CellularModemStatus interface{}

    // Cellular Modem temperature. The type is interface{} with range:
    // -32768..32767.
    CellularModemTemperature interface{}

    // Cellular Modem PRI sku id. The type is string.
    CellularPriSkuid interface{}

    // Cellular Modem PRI version. The type is string.
    CellularPriVersion interface{}

    // Cellular Modem carrier. The type is string.
    CellularCarrier interface{}

    // Cellular Modem OEM PRI version. The type is string.
    CellularOemPriVersion interface{}
}

func (cellwanHardware *CellwanOperData_CellwanHardware) GetEntityData() *types.CommonEntityData {
    cellwanHardware.EntityData.YFilter = cellwanHardware.YFilter
    cellwanHardware.EntityData.YangName = "cellwan-hardware"
    cellwanHardware.EntityData.BundleName = "cisco_ios_xe"
    cellwanHardware.EntityData.ParentYangName = "cellwan-oper-data"
    cellwanHardware.EntityData.SegmentPath = "cellwan-hardware" + types.AddKeyToken(cellwanHardware.CellularInterface, "cellular-interface")
    cellwanHardware.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanHardware.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanHardware.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanHardware.EntityData.Children = types.NewOrderedMap()
    cellwanHardware.EntityData.Leafs = types.NewOrderedMap()
    cellwanHardware.EntityData.Leafs.Append("cellular-interface", types.YLeaf{"CellularInterface", cellwanHardware.CellularInterface})
    cellwanHardware.EntityData.Leafs.Append("cellular-firmware-version", types.YLeaf{"CellularFirmwareVersion", cellwanHardware.CellularFirmwareVersion})
    cellwanHardware.EntityData.Leafs.Append("cellular-firmware-build-time", types.YLeaf{"CellularFirmwareBuildTime", cellwanHardware.CellularFirmwareBuildTime})
    cellwanHardware.EntityData.Leafs.Append("cellular-hardware-version", types.YLeaf{"CellularHardwareVersion", cellwanHardware.CellularHardwareVersion})
    cellwanHardware.EntityData.Leafs.Append("cellular-device-model", types.YLeaf{"CellularDeviceModel", cellwanHardware.CellularDeviceModel})
    cellwanHardware.EntityData.Leafs.Append("cellular-imsi", types.YLeaf{"CellularImsi", cellwanHardware.CellularImsi})
    cellwanHardware.EntityData.Leafs.Append("cellular-imei", types.YLeaf{"CellularImei", cellwanHardware.CellularImei})
    cellwanHardware.EntityData.Leafs.Append("cellular-iccid", types.YLeaf{"CellularIccid", cellwanHardware.CellularIccid})
    cellwanHardware.EntityData.Leafs.Append("cellular-msisdn", types.YLeaf{"CellularMsisdn", cellwanHardware.CellularMsisdn})
    cellwanHardware.EntityData.Leafs.Append("cellular-fsn", types.YLeaf{"CellularFsn", cellwanHardware.CellularFsn})
    cellwanHardware.EntityData.Leafs.Append("cellular-modem-status", types.YLeaf{"CellularModemStatus", cellwanHardware.CellularModemStatus})
    cellwanHardware.EntityData.Leafs.Append("cellular-modem-temperature", types.YLeaf{"CellularModemTemperature", cellwanHardware.CellularModemTemperature})
    cellwanHardware.EntityData.Leafs.Append("cellular-pri-skuid", types.YLeaf{"CellularPriSkuid", cellwanHardware.CellularPriSkuid})
    cellwanHardware.EntityData.Leafs.Append("cellular-pri-version", types.YLeaf{"CellularPriVersion", cellwanHardware.CellularPriVersion})
    cellwanHardware.EntityData.Leafs.Append("cellular-carrier", types.YLeaf{"CellularCarrier", cellwanHardware.CellularCarrier})
    cellwanHardware.EntityData.Leafs.Append("cellular-oem-pri-version", types.YLeaf{"CellularOemPriVersion", cellwanHardware.CellularOemPriVersion})

    cellwanHardware.EntityData.YListKeys = []string {"CellularInterface"}

    return &(cellwanHardware.EntityData)
}

// CellwanOperData_CellwanRadio
// Cellwan modem radio data
type CellwanOperData_CellwanRadio struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Cellular interface. The type is string.
    CellularInterface interface{}

    // Cellular modem technology. The type is ModemTechnology.
    Technology interface{}

    // Cellular modem radio power mode. The type is CwRadioPowerStatus.
    RadioPowerMode interface{}

    // Cellular radio receive (Rx) channel number. The type is interface{} with
    // range: 0..4294967295.
    RadioRxChannel interface{}

    // Cellular radio transmit (Tx) channel number. The type is interface{} with
    // range: 0..4294967295.
    RadioTxChannel interface{}

    // Cellular radio band number. The type is interface{} with range:
    // 0..4294967295.
    RadioBand interface{}

    // Cellular radio bandwidth in MHz. The type is RadioBandwidth.
    Bandwidth interface{}

    // Cellular RSSI - Received Signal Strength Indication in dBm. The type is
    // interface{} with range: -32768..32767.
    RadioRssi interface{}

    // Cellular RSRP - Reference Signal Received Power in dBm. The type is
    // interface{} with range: -32768..32767.
    RadioRsrp interface{}

    // Cellular RSRQ - Reference Signal Received Quality in dB. The type is
    // interface{} with range: -128..127.
    RadioRsrq interface{}

    // Cellular SNR - Signal to Noise Ratio in dB. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    RadioSnr interface{}

    // Cellular Radio Access Technology preference. The type is RatPreference.
    RadioRatPreference interface{}

    // Cellular Radio Access Technology selected. The type is RatTechnology.
    RadioRatSelected interface{}
}

func (cellwanRadio *CellwanOperData_CellwanRadio) GetEntityData() *types.CommonEntityData {
    cellwanRadio.EntityData.YFilter = cellwanRadio.YFilter
    cellwanRadio.EntityData.YangName = "cellwan-radio"
    cellwanRadio.EntityData.BundleName = "cisco_ios_xe"
    cellwanRadio.EntityData.ParentYangName = "cellwan-oper-data"
    cellwanRadio.EntityData.SegmentPath = "cellwan-radio" + types.AddKeyToken(cellwanRadio.CellularInterface, "cellular-interface")
    cellwanRadio.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanRadio.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanRadio.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanRadio.EntityData.Children = types.NewOrderedMap()
    cellwanRadio.EntityData.Leafs = types.NewOrderedMap()
    cellwanRadio.EntityData.Leafs.Append("cellular-interface", types.YLeaf{"CellularInterface", cellwanRadio.CellularInterface})
    cellwanRadio.EntityData.Leafs.Append("technology", types.YLeaf{"Technology", cellwanRadio.Technology})
    cellwanRadio.EntityData.Leafs.Append("radio-power-mode", types.YLeaf{"RadioPowerMode", cellwanRadio.RadioPowerMode})
    cellwanRadio.EntityData.Leafs.Append("radio-rx-channel", types.YLeaf{"RadioRxChannel", cellwanRadio.RadioRxChannel})
    cellwanRadio.EntityData.Leafs.Append("radio-tx-channel", types.YLeaf{"RadioTxChannel", cellwanRadio.RadioTxChannel})
    cellwanRadio.EntityData.Leafs.Append("radio-band", types.YLeaf{"RadioBand", cellwanRadio.RadioBand})
    cellwanRadio.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", cellwanRadio.Bandwidth})
    cellwanRadio.EntityData.Leafs.Append("radio-rssi", types.YLeaf{"RadioRssi", cellwanRadio.RadioRssi})
    cellwanRadio.EntityData.Leafs.Append("radio-rsrp", types.YLeaf{"RadioRsrp", cellwanRadio.RadioRsrp})
    cellwanRadio.EntityData.Leafs.Append("radio-rsrq", types.YLeaf{"RadioRsrq", cellwanRadio.RadioRsrq})
    cellwanRadio.EntityData.Leafs.Append("radio-snr", types.YLeaf{"RadioSnr", cellwanRadio.RadioSnr})
    cellwanRadio.EntityData.Leafs.Append("radio-rat-preference", types.YLeaf{"RadioRatPreference", cellwanRadio.RadioRatPreference})
    cellwanRadio.EntityData.Leafs.Append("radio-rat-selected", types.YLeaf{"RadioRatSelected", cellwanRadio.RadioRatSelected})

    cellwanRadio.EntityData.YListKeys = []string {"CellularInterface"}

    return &(cellwanRadio.EntityData)
}

// CellwanOperData_CellwanNetwork
// Cellwan modem network data
type CellwanOperData_CellwanNetwork struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Cellular interface. The type is string.
    CellularInterface interface{}

    // Current Cellular modem time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    CellularModemTime interface{}

    // Cellular Radio Access Technology selected. The type is RatTechnology.
    RadioAccessTechnologySelected interface{}

    // Cellular network current service status. The type is ServiceStatus.
    CurrentServiceStatus interface{}

    // Cellular modem current system identifier. The type is interface{} with
    // range: 0..65535.
    CurrentSystemIdentifier interface{}

    // Cellular modem current network identifier. The type is interface{} with
    // range: 0..65535.
    CurrentNetworkIdentifier interface{}

    // Cellular network current service type. The type is ModemService.
    CurrentServiceType interface{}

    // Cellular network current roaming status. The type is string.
    CurrentRoamingStatus interface{}

    // Cellular network selection mode. The type is string.
    NetworkSelectionMode interface{}

    // Cellular current network name. The type is string.
    NetworkName interface{}

    // Cellular network mobile country code. The type is interface{} with range:
    // 0..65535.
    MobileCountryCode interface{}

    // Cellular modem mobile network code. The type is interface{} with range:
    // 0..65535.
    MobileNetworkCode interface{}

    // Cellular packet switched (PS) domain state. The type is string.
    PacketSwitchDomainState interface{}

    // Cellular LTE carrier aggregation (CA) state. The type is LteCa.
    LteCarrierAggregation interface{}

    // Cellular network registration state. The type is RegState.
    RegistrationState interface{}

    // Cellular network EMM registration state. The type is string.
    EmmState interface{}

    // Cellular network EMM sub-state. The type is string.
    EmmSubstate interface{}

    // Cellular network location area code. The type is interface{} with range:
    // 0..65535.
    LocationAreaCode interface{}

    // Cellular network tracking area code. The type is interface{} with range:
    // 0..65535.
    TrackingAreaCode interface{}

    // Cellular network cell ID. The type is interface{} with range:
    // 0..18446744073709551615.
    CellId interface{}

    // Cellular modem negotiated network MTU size. The type is interface{} with
    // range: 0..65535.
    NegotiatedNetworkMtu interface{}
}

func (cellwanNetwork *CellwanOperData_CellwanNetwork) GetEntityData() *types.CommonEntityData {
    cellwanNetwork.EntityData.YFilter = cellwanNetwork.YFilter
    cellwanNetwork.EntityData.YangName = "cellwan-network"
    cellwanNetwork.EntityData.BundleName = "cisco_ios_xe"
    cellwanNetwork.EntityData.ParentYangName = "cellwan-oper-data"
    cellwanNetwork.EntityData.SegmentPath = "cellwan-network" + types.AddKeyToken(cellwanNetwork.CellularInterface, "cellular-interface")
    cellwanNetwork.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanNetwork.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanNetwork.EntityData.Children = types.NewOrderedMap()
    cellwanNetwork.EntityData.Leafs = types.NewOrderedMap()
    cellwanNetwork.EntityData.Leafs.Append("cellular-interface", types.YLeaf{"CellularInterface", cellwanNetwork.CellularInterface})
    cellwanNetwork.EntityData.Leafs.Append("cellular-modem-time", types.YLeaf{"CellularModemTime", cellwanNetwork.CellularModemTime})
    cellwanNetwork.EntityData.Leafs.Append("radio-access-technology-selected", types.YLeaf{"RadioAccessTechnologySelected", cellwanNetwork.RadioAccessTechnologySelected})
    cellwanNetwork.EntityData.Leafs.Append("current-service-status", types.YLeaf{"CurrentServiceStatus", cellwanNetwork.CurrentServiceStatus})
    cellwanNetwork.EntityData.Leafs.Append("current-system-identifier", types.YLeaf{"CurrentSystemIdentifier", cellwanNetwork.CurrentSystemIdentifier})
    cellwanNetwork.EntityData.Leafs.Append("current-network-identifier", types.YLeaf{"CurrentNetworkIdentifier", cellwanNetwork.CurrentNetworkIdentifier})
    cellwanNetwork.EntityData.Leafs.Append("current-service-type", types.YLeaf{"CurrentServiceType", cellwanNetwork.CurrentServiceType})
    cellwanNetwork.EntityData.Leafs.Append("current-roaming-status", types.YLeaf{"CurrentRoamingStatus", cellwanNetwork.CurrentRoamingStatus})
    cellwanNetwork.EntityData.Leafs.Append("network-selection-mode", types.YLeaf{"NetworkSelectionMode", cellwanNetwork.NetworkSelectionMode})
    cellwanNetwork.EntityData.Leafs.Append("network-name", types.YLeaf{"NetworkName", cellwanNetwork.NetworkName})
    cellwanNetwork.EntityData.Leafs.Append("mobile-country-code", types.YLeaf{"MobileCountryCode", cellwanNetwork.MobileCountryCode})
    cellwanNetwork.EntityData.Leafs.Append("mobile-network-code", types.YLeaf{"MobileNetworkCode", cellwanNetwork.MobileNetworkCode})
    cellwanNetwork.EntityData.Leafs.Append("packet-switch-domain-state", types.YLeaf{"PacketSwitchDomainState", cellwanNetwork.PacketSwitchDomainState})
    cellwanNetwork.EntityData.Leafs.Append("lte-carrier-aggregation", types.YLeaf{"LteCarrierAggregation", cellwanNetwork.LteCarrierAggregation})
    cellwanNetwork.EntityData.Leafs.Append("registration-state", types.YLeaf{"RegistrationState", cellwanNetwork.RegistrationState})
    cellwanNetwork.EntityData.Leafs.Append("emm-state", types.YLeaf{"EmmState", cellwanNetwork.EmmState})
    cellwanNetwork.EntityData.Leafs.Append("emm-substate", types.YLeaf{"EmmSubstate", cellwanNetwork.EmmSubstate})
    cellwanNetwork.EntityData.Leafs.Append("location-area-code", types.YLeaf{"LocationAreaCode", cellwanNetwork.LocationAreaCode})
    cellwanNetwork.EntityData.Leafs.Append("tracking-area-code", types.YLeaf{"TrackingAreaCode", cellwanNetwork.TrackingAreaCode})
    cellwanNetwork.EntityData.Leafs.Append("cell-id", types.YLeaf{"CellId", cellwanNetwork.CellId})
    cellwanNetwork.EntityData.Leafs.Append("negotiated-network-mtu", types.YLeaf{"NegotiatedNetworkMtu", cellwanNetwork.NegotiatedNetworkMtu})

    cellwanNetwork.EntityData.YListKeys = []string {"CellularInterface"}

    return &(cellwanNetwork.EntityData)
}

// CellwanOperData_CellwanConnection
// Cellwan modem connection data
type CellwanOperData_CellwanConnection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Cellular interface. The type is string.
    CellularInterface interface{}

    // Cellular modem active profile number. The type is interface{} with range:
    // 0..18446744073709551615.
    ActiveProfile interface{}

    // Cellular modem packet session status. The type is PacketSessStatus.
    CellularPacketStatus interface{}

    // Cellular data packets transmitted. The type is interface{} with range:
    // 0..18446744073709551615.
    TxPackets interface{}

    // Cellular data packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPackets interface{}

    // Cellular data bytes transmitted. The type is interface{} with range:
    // 0..18446744073709551615.
    TxBytes interface{}

    // Cellular data bytes received. The type is interface{} with range:
    // 0..18446744073709551615.
    RxBytes interface{}

    // Cellular interface IPv4 Address. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv4Addr interface{}

    // Cellular interface IPv6 Address. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Addr interface{}

    // Cellular interface primary IPv4 DNS Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv4DnsPri interface{}

    // Cellular interface secondary IPv4 DNS Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv4DnsSec interface{}

    // Cellular interface primary IPv6 DNS Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6DnsPri interface{}

    // Cellular interface secondary IPv6 DNS Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6DnsSec interface{}

    // Cellular interface profile IP address scope. The type is ProfileScope.
    Scope interface{}
}

func (cellwanConnection *CellwanOperData_CellwanConnection) GetEntityData() *types.CommonEntityData {
    cellwanConnection.EntityData.YFilter = cellwanConnection.YFilter
    cellwanConnection.EntityData.YangName = "cellwan-connection"
    cellwanConnection.EntityData.BundleName = "cisco_ios_xe"
    cellwanConnection.EntityData.ParentYangName = "cellwan-oper-data"
    cellwanConnection.EntityData.SegmentPath = "cellwan-connection" + types.AddKeyToken(cellwanConnection.CellularInterface, "cellular-interface")
    cellwanConnection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanConnection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanConnection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanConnection.EntityData.Children = types.NewOrderedMap()
    cellwanConnection.EntityData.Leafs = types.NewOrderedMap()
    cellwanConnection.EntityData.Leafs.Append("cellular-interface", types.YLeaf{"CellularInterface", cellwanConnection.CellularInterface})
    cellwanConnection.EntityData.Leafs.Append("active-profile", types.YLeaf{"ActiveProfile", cellwanConnection.ActiveProfile})
    cellwanConnection.EntityData.Leafs.Append("cellular-packet-status", types.YLeaf{"CellularPacketStatus", cellwanConnection.CellularPacketStatus})
    cellwanConnection.EntityData.Leafs.Append("tx-packets", types.YLeaf{"TxPackets", cellwanConnection.TxPackets})
    cellwanConnection.EntityData.Leafs.Append("rx-packets", types.YLeaf{"RxPackets", cellwanConnection.RxPackets})
    cellwanConnection.EntityData.Leafs.Append("tx-bytes", types.YLeaf{"TxBytes", cellwanConnection.TxBytes})
    cellwanConnection.EntityData.Leafs.Append("rx-bytes", types.YLeaf{"RxBytes", cellwanConnection.RxBytes})
    cellwanConnection.EntityData.Leafs.Append("ipv4-addr", types.YLeaf{"Ipv4Addr", cellwanConnection.Ipv4Addr})
    cellwanConnection.EntityData.Leafs.Append("ipv6-addr", types.YLeaf{"Ipv6Addr", cellwanConnection.Ipv6Addr})
    cellwanConnection.EntityData.Leafs.Append("ipv4-dns-pri", types.YLeaf{"Ipv4DnsPri", cellwanConnection.Ipv4DnsPri})
    cellwanConnection.EntityData.Leafs.Append("ipv4-dns-sec", types.YLeaf{"Ipv4DnsSec", cellwanConnection.Ipv4DnsSec})
    cellwanConnection.EntityData.Leafs.Append("ipv6-dns-pri", types.YLeaf{"Ipv6DnsPri", cellwanConnection.Ipv6DnsPri})
    cellwanConnection.EntityData.Leafs.Append("ipv6-dns-sec", types.YLeaf{"Ipv6DnsSec", cellwanConnection.Ipv6DnsSec})
    cellwanConnection.EntityData.Leafs.Append("scope", types.YLeaf{"Scope", cellwanConnection.Scope})

    cellwanConnection.EntityData.YListKeys = []string {"CellularInterface"}

    return &(cellwanConnection.EntityData)
}

// CellwanOperData_CellwanSecurity
// Cellwan modem sim security data
type CellwanOperData_CellwanSecurity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Cellular interface. The type is string.
    CellularInterface interface{}

    // Cellular modem active SIM slot status. The type is interface{} with range:
    // -128..127.
    ActiveSim interface{}

    // Cellular modem SIM switchover attempts. The type is interface{} with range:
    // 0..4294967295.
    SimNumSwitchover interface{}

    // Cellular SIM Card Holder Verification (CHV1) status. The type is
    // CellwanChv1SimStatus.
    Chv1Status interface{}

    // Cellular SIM status. The type is CellwanSimStatus.
    SimStatus interface{}

    // Cellular SIM user operation. The type is CellwanSimUserOp.
    SimOper interface{}

    // Cellular SIM number of CHV1 retries remaining. The type is interface{} with
    // range: -128..127.
    NumRetries interface{}
}

func (cellwanSecurity *CellwanOperData_CellwanSecurity) GetEntityData() *types.CommonEntityData {
    cellwanSecurity.EntityData.YFilter = cellwanSecurity.YFilter
    cellwanSecurity.EntityData.YangName = "cellwan-security"
    cellwanSecurity.EntityData.BundleName = "cisco_ios_xe"
    cellwanSecurity.EntityData.ParentYangName = "cellwan-oper-data"
    cellwanSecurity.EntityData.SegmentPath = "cellwan-security" + types.AddKeyToken(cellwanSecurity.CellularInterface, "cellular-interface")
    cellwanSecurity.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanSecurity.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanSecurity.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanSecurity.EntityData.Children = types.NewOrderedMap()
    cellwanSecurity.EntityData.Leafs = types.NewOrderedMap()
    cellwanSecurity.EntityData.Leafs.Append("cellular-interface", types.YLeaf{"CellularInterface", cellwanSecurity.CellularInterface})
    cellwanSecurity.EntityData.Leafs.Append("active-sim", types.YLeaf{"ActiveSim", cellwanSecurity.ActiveSim})
    cellwanSecurity.EntityData.Leafs.Append("sim-num-switchover", types.YLeaf{"SimNumSwitchover", cellwanSecurity.SimNumSwitchover})
    cellwanSecurity.EntityData.Leafs.Append("chv1-status", types.YLeaf{"Chv1Status", cellwanSecurity.Chv1Status})
    cellwanSecurity.EntityData.Leafs.Append("sim-status", types.YLeaf{"SimStatus", cellwanSecurity.SimStatus})
    cellwanSecurity.EntityData.Leafs.Append("sim-oper", types.YLeaf{"SimOper", cellwanSecurity.SimOper})
    cellwanSecurity.EntityData.Leafs.Append("num-retries", types.YLeaf{"NumRetries", cellwanSecurity.NumRetries})

    cellwanSecurity.EntityData.YListKeys = []string {"CellularInterface"}

    return &(cellwanSecurity.EntityData)
}

// CellwanOperData_CellwanSms
// Cellwan modem sms data
type CellwanOperData_CellwanSms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Cellular interface. The type is string.
    CellularInterface interface{}

    // Number of incoming SMS stored in Cellular modem. The type is interface{}
    // with range: 0..65535.
    InSmsCount interface{}

    // Number of incoming SMS archived in Cellular modem. The type is interface{}
    // with range: 0..65535.
    InSmsArchived interface{}

    // Number of incoming SMS deleted in Cellular modem. The type is interface{}
    // with range: 0..65535.
    InSmsDeleted interface{}

    // Number of SIM/modem SMS record allocated. The type is interface{} with
    // range: 0..65535.
    InSmsMax interface{}

    // Number of SIM/modem SMS record used. The type is interface{} with range:
    // 0..65535.
    InSmsUsed interface{}

    // Number of SMS triggerred data callback. The type is interface{} with range:
    // 0..65535.
    SmsCallbackCount interface{}

    // Number of successful archive on in SMS. The type is interface{} with range:
    // 0..65535.
    InSmsArchCount interface{}

    // Number of failed archive on in SMS. The type is interface{} with range:
    // 0..65535.
    InSmsArchErrorCount interface{}

    // Number of outgoing SMS successfully sent. The type is interface{} with
    // range: 0..65535.
    OutSmsCount interface{}

    // Number of outgoing SMS send failure. The type is interface{} with range:
    // 0..65535.
    OutSmsErrorCount interface{}

    // Number of outgoing SMS pending. The type is interface{} with range:
    // 0..65535.
    OutSmsPending interface{}

    // Number of successful archive out in SMS. The type is interface{} with
    // range: 0..65535.
    OutSmsArchCount interface{}

    // Number of failed archive on out SMS. The type is interface{} with range:
    // 0..65535.
    OutSmsArchErrorCount interface{}
}

func (cellwanSms *CellwanOperData_CellwanSms) GetEntityData() *types.CommonEntityData {
    cellwanSms.EntityData.YFilter = cellwanSms.YFilter
    cellwanSms.EntityData.YangName = "cellwan-sms"
    cellwanSms.EntityData.BundleName = "cisco_ios_xe"
    cellwanSms.EntityData.ParentYangName = "cellwan-oper-data"
    cellwanSms.EntityData.SegmentPath = "cellwan-sms" + types.AddKeyToken(cellwanSms.CellularInterface, "cellular-interface")
    cellwanSms.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanSms.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanSms.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanSms.EntityData.Children = types.NewOrderedMap()
    cellwanSms.EntityData.Leafs = types.NewOrderedMap()
    cellwanSms.EntityData.Leafs.Append("cellular-interface", types.YLeaf{"CellularInterface", cellwanSms.CellularInterface})
    cellwanSms.EntityData.Leafs.Append("in-sms-count", types.YLeaf{"InSmsCount", cellwanSms.InSmsCount})
    cellwanSms.EntityData.Leafs.Append("in-sms-archived", types.YLeaf{"InSmsArchived", cellwanSms.InSmsArchived})
    cellwanSms.EntityData.Leafs.Append("in-sms-deleted", types.YLeaf{"InSmsDeleted", cellwanSms.InSmsDeleted})
    cellwanSms.EntityData.Leafs.Append("in-sms-max", types.YLeaf{"InSmsMax", cellwanSms.InSmsMax})
    cellwanSms.EntityData.Leafs.Append("in-sms-used", types.YLeaf{"InSmsUsed", cellwanSms.InSmsUsed})
    cellwanSms.EntityData.Leafs.Append("sms-callback-count", types.YLeaf{"SmsCallbackCount", cellwanSms.SmsCallbackCount})
    cellwanSms.EntityData.Leafs.Append("in-sms-arch-count", types.YLeaf{"InSmsArchCount", cellwanSms.InSmsArchCount})
    cellwanSms.EntityData.Leafs.Append("in-sms-arch-error-count", types.YLeaf{"InSmsArchErrorCount", cellwanSms.InSmsArchErrorCount})
    cellwanSms.EntityData.Leafs.Append("out-sms-count", types.YLeaf{"OutSmsCount", cellwanSms.OutSmsCount})
    cellwanSms.EntityData.Leafs.Append("out-sms-error-count", types.YLeaf{"OutSmsErrorCount", cellwanSms.OutSmsErrorCount})
    cellwanSms.EntityData.Leafs.Append("out-sms-pending", types.YLeaf{"OutSmsPending", cellwanSms.OutSmsPending})
    cellwanSms.EntityData.Leafs.Append("out-sms-arch-count", types.YLeaf{"OutSmsArchCount", cellwanSms.OutSmsArchCount})
    cellwanSms.EntityData.Leafs.Append("out-sms-arch-error-count", types.YLeaf{"OutSmsArchErrorCount", cellwanSms.OutSmsArchErrorCount})

    cellwanSms.EntityData.YListKeys = []string {"CellularInterface"}

    return &(cellwanSms.EntityData)
}

// CellwanOperData_CellwanGps
// Cellwan modem gps data
type CellwanOperData_CellwanGps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Cellular interface. The type is string.
    CellularInterface interface{}

    // Cellular modem GPS feature state status. The type is CwanGpsFeatureState.
    GpsFeatureState interface{}

    // Cellular modem GPS port selection status. The type is CwanGpsPortSelected.
    PortSelected interface{}

    // Cellular modem GPS state status. The type is CwanGpsState.
    State interface{}

    // Cellular modem GPS mode selection status. The type is CwanGpsModeSelected.
    ModeSelected interface{}

    // Cellular modem GPS latitude details. The type is string.
    Latitude interface{}

    // Cellular modem GPS longitude details. The type is string.
    Longitude interface{}

    // Cellular modem GPS timestamp details. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Timestamp interface{}
}

func (cellwanGps *CellwanOperData_CellwanGps) GetEntityData() *types.CommonEntityData {
    cellwanGps.EntityData.YFilter = cellwanGps.YFilter
    cellwanGps.EntityData.YangName = "cellwan-gps"
    cellwanGps.EntityData.BundleName = "cisco_ios_xe"
    cellwanGps.EntityData.ParentYangName = "cellwan-oper-data"
    cellwanGps.EntityData.SegmentPath = "cellwan-gps" + types.AddKeyToken(cellwanGps.CellularInterface, "cellular-interface")
    cellwanGps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanGps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanGps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanGps.EntityData.Children = types.NewOrderedMap()
    cellwanGps.EntityData.Leafs = types.NewOrderedMap()
    cellwanGps.EntityData.Leafs.Append("cellular-interface", types.YLeaf{"CellularInterface", cellwanGps.CellularInterface})
    cellwanGps.EntityData.Leafs.Append("gps-feature-state", types.YLeaf{"GpsFeatureState", cellwanGps.GpsFeatureState})
    cellwanGps.EntityData.Leafs.Append("port-selected", types.YLeaf{"PortSelected", cellwanGps.PortSelected})
    cellwanGps.EntityData.Leafs.Append("state", types.YLeaf{"State", cellwanGps.State})
    cellwanGps.EntityData.Leafs.Append("mode-selected", types.YLeaf{"ModeSelected", cellwanGps.ModeSelected})
    cellwanGps.EntityData.Leafs.Append("latitude", types.YLeaf{"Latitude", cellwanGps.Latitude})
    cellwanGps.EntityData.Leafs.Append("longitude", types.YLeaf{"Longitude", cellwanGps.Longitude})
    cellwanGps.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", cellwanGps.Timestamp})

    cellwanGps.EntityData.YListKeys = []string {"CellularInterface"}

    return &(cellwanGps.EntityData)
}

