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

// PacketSessStatus represents Packet Session Status indicates the Cellular packet session state
type PacketSessStatus string

const (
    // Cellular packet session status is inactive
    PacketSessStatus_packet_session_status_inactive PacketSessStatus = "packet-session-status-inactive"

    // Cellular packet session status is active
    PacketSessStatus_packet_session_status_active PacketSessStatus = "packet-session-status-active"
)

// ProfileScope represents Profile Scope indicates the IP address scope
type ProfileScope string

const (
    // IP address scope is Global
    ProfileScope_scope_global ProfileScope = "scope-global"

    // IP address scope is Link
    ProfileScope_scope_link ProfileScope = "scope-link"
)

// CellwanOperData
// Cellwan operational data
type CellwanOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Cellwan modem hardware data. The type is slice of
    // CellwanOperData_CellwanHardware.
    CellwanHardware []CellwanOperData_CellwanHardware

    // Cellwan modem radio data. The type is slice of
    // CellwanOperData_CellwanRadio.
    CellwanRadio []CellwanOperData_CellwanRadio

    // Cellwan modem network data. The type is slice of
    // CellwanOperData_CellwanNetwork.
    CellwanNetwork []CellwanOperData_CellwanNetwork

    // Cellwan modem connection data. The type is slice of
    // CellwanOperData_CellwanConnection.
    CellwanConnection []CellwanOperData_CellwanConnection
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

    cellwanOperData.EntityData.Children = make(map[string]types.YChild)
    cellwanOperData.EntityData.Children["cellwan-hardware"] = types.YChild{"CellwanHardware", nil}
    for i := range cellwanOperData.CellwanHardware {
        cellwanOperData.EntityData.Children[types.GetSegmentPath(&cellwanOperData.CellwanHardware[i])] = types.YChild{"CellwanHardware", &cellwanOperData.CellwanHardware[i]}
    }
    cellwanOperData.EntityData.Children["cellwan-radio"] = types.YChild{"CellwanRadio", nil}
    for i := range cellwanOperData.CellwanRadio {
        cellwanOperData.EntityData.Children[types.GetSegmentPath(&cellwanOperData.CellwanRadio[i])] = types.YChild{"CellwanRadio", &cellwanOperData.CellwanRadio[i]}
    }
    cellwanOperData.EntityData.Children["cellwan-network"] = types.YChild{"CellwanNetwork", nil}
    for i := range cellwanOperData.CellwanNetwork {
        cellwanOperData.EntityData.Children[types.GetSegmentPath(&cellwanOperData.CellwanNetwork[i])] = types.YChild{"CellwanNetwork", &cellwanOperData.CellwanNetwork[i]}
    }
    cellwanOperData.EntityData.Children["cellwan-connection"] = types.YChild{"CellwanConnection", nil}
    for i := range cellwanOperData.CellwanConnection {
        cellwanOperData.EntityData.Children[types.GetSegmentPath(&cellwanOperData.CellwanConnection[i])] = types.YChild{"CellwanConnection", &cellwanOperData.CellwanConnection[i]}
    }
    cellwanOperData.EntityData.Leafs = make(map[string]types.YLeaf)
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
    cellwanHardware.EntityData.SegmentPath = "cellwan-hardware" + "[cellular-interface='" + fmt.Sprintf("%v", cellwanHardware.CellularInterface) + "']"
    cellwanHardware.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanHardware.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanHardware.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanHardware.EntityData.Children = make(map[string]types.YChild)
    cellwanHardware.EntityData.Leafs = make(map[string]types.YLeaf)
    cellwanHardware.EntityData.Leafs["cellular-interface"] = types.YLeaf{"CellularInterface", cellwanHardware.CellularInterface}
    cellwanHardware.EntityData.Leafs["cellular-firmware-version"] = types.YLeaf{"CellularFirmwareVersion", cellwanHardware.CellularFirmwareVersion}
    cellwanHardware.EntityData.Leafs["cellular-firmware-build-time"] = types.YLeaf{"CellularFirmwareBuildTime", cellwanHardware.CellularFirmwareBuildTime}
    cellwanHardware.EntityData.Leafs["cellular-hardware-version"] = types.YLeaf{"CellularHardwareVersion", cellwanHardware.CellularHardwareVersion}
    cellwanHardware.EntityData.Leafs["cellular-device-model"] = types.YLeaf{"CellularDeviceModel", cellwanHardware.CellularDeviceModel}
    cellwanHardware.EntityData.Leafs["cellular-imsi"] = types.YLeaf{"CellularImsi", cellwanHardware.CellularImsi}
    cellwanHardware.EntityData.Leafs["cellular-imei"] = types.YLeaf{"CellularImei", cellwanHardware.CellularImei}
    cellwanHardware.EntityData.Leafs["cellular-iccid"] = types.YLeaf{"CellularIccid", cellwanHardware.CellularIccid}
    cellwanHardware.EntityData.Leafs["cellular-msisdn"] = types.YLeaf{"CellularMsisdn", cellwanHardware.CellularMsisdn}
    cellwanHardware.EntityData.Leafs["cellular-fsn"] = types.YLeaf{"CellularFsn", cellwanHardware.CellularFsn}
    cellwanHardware.EntityData.Leafs["cellular-modem-status"] = types.YLeaf{"CellularModemStatus", cellwanHardware.CellularModemStatus}
    cellwanHardware.EntityData.Leafs["cellular-modem-temperature"] = types.YLeaf{"CellularModemTemperature", cellwanHardware.CellularModemTemperature}
    cellwanHardware.EntityData.Leafs["cellular-pri-skuid"] = types.YLeaf{"CellularPriSkuid", cellwanHardware.CellularPriSkuid}
    cellwanHardware.EntityData.Leafs["cellular-pri-version"] = types.YLeaf{"CellularPriVersion", cellwanHardware.CellularPriVersion}
    cellwanHardware.EntityData.Leafs["cellular-carrier"] = types.YLeaf{"CellularCarrier", cellwanHardware.CellularCarrier}
    cellwanHardware.EntityData.Leafs["cellular-oem-pri-version"] = types.YLeaf{"CellularOemPriVersion", cellwanHardware.CellularOemPriVersion}
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
    cellwanRadio.EntityData.SegmentPath = "cellwan-radio" + "[cellular-interface='" + fmt.Sprintf("%v", cellwanRadio.CellularInterface) + "']"
    cellwanRadio.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanRadio.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanRadio.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanRadio.EntityData.Children = make(map[string]types.YChild)
    cellwanRadio.EntityData.Leafs = make(map[string]types.YLeaf)
    cellwanRadio.EntityData.Leafs["cellular-interface"] = types.YLeaf{"CellularInterface", cellwanRadio.CellularInterface}
    cellwanRadio.EntityData.Leafs["technology"] = types.YLeaf{"Technology", cellwanRadio.Technology}
    cellwanRadio.EntityData.Leafs["radio-power-mode"] = types.YLeaf{"RadioPowerMode", cellwanRadio.RadioPowerMode}
    cellwanRadio.EntityData.Leafs["radio-rx-channel"] = types.YLeaf{"RadioRxChannel", cellwanRadio.RadioRxChannel}
    cellwanRadio.EntityData.Leafs["radio-tx-channel"] = types.YLeaf{"RadioTxChannel", cellwanRadio.RadioTxChannel}
    cellwanRadio.EntityData.Leafs["radio-band"] = types.YLeaf{"RadioBand", cellwanRadio.RadioBand}
    cellwanRadio.EntityData.Leafs["bandwidth"] = types.YLeaf{"Bandwidth", cellwanRadio.Bandwidth}
    cellwanRadio.EntityData.Leafs["radio-rssi"] = types.YLeaf{"RadioRssi", cellwanRadio.RadioRssi}
    cellwanRadio.EntityData.Leafs["radio-rsrp"] = types.YLeaf{"RadioRsrp", cellwanRadio.RadioRsrp}
    cellwanRadio.EntityData.Leafs["radio-rsrq"] = types.YLeaf{"RadioRsrq", cellwanRadio.RadioRsrq}
    cellwanRadio.EntityData.Leafs["radio-snr"] = types.YLeaf{"RadioSnr", cellwanRadio.RadioSnr}
    cellwanRadio.EntityData.Leafs["radio-rat-preference"] = types.YLeaf{"RadioRatPreference", cellwanRadio.RadioRatPreference}
    cellwanRadio.EntityData.Leafs["radio-rat-selected"] = types.YLeaf{"RadioRatSelected", cellwanRadio.RadioRatSelected}
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
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
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
    cellwanNetwork.EntityData.SegmentPath = "cellwan-network" + "[cellular-interface='" + fmt.Sprintf("%v", cellwanNetwork.CellularInterface) + "']"
    cellwanNetwork.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanNetwork.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanNetwork.EntityData.Children = make(map[string]types.YChild)
    cellwanNetwork.EntityData.Leafs = make(map[string]types.YLeaf)
    cellwanNetwork.EntityData.Leafs["cellular-interface"] = types.YLeaf{"CellularInterface", cellwanNetwork.CellularInterface}
    cellwanNetwork.EntityData.Leafs["cellular-modem-time"] = types.YLeaf{"CellularModemTime", cellwanNetwork.CellularModemTime}
    cellwanNetwork.EntityData.Leafs["radio-access-technology-selected"] = types.YLeaf{"RadioAccessTechnologySelected", cellwanNetwork.RadioAccessTechnologySelected}
    cellwanNetwork.EntityData.Leafs["current-service-status"] = types.YLeaf{"CurrentServiceStatus", cellwanNetwork.CurrentServiceStatus}
    cellwanNetwork.EntityData.Leafs["current-system-identifier"] = types.YLeaf{"CurrentSystemIdentifier", cellwanNetwork.CurrentSystemIdentifier}
    cellwanNetwork.EntityData.Leafs["current-network-identifier"] = types.YLeaf{"CurrentNetworkIdentifier", cellwanNetwork.CurrentNetworkIdentifier}
    cellwanNetwork.EntityData.Leafs["current-service-type"] = types.YLeaf{"CurrentServiceType", cellwanNetwork.CurrentServiceType}
    cellwanNetwork.EntityData.Leafs["current-roaming-status"] = types.YLeaf{"CurrentRoamingStatus", cellwanNetwork.CurrentRoamingStatus}
    cellwanNetwork.EntityData.Leafs["network-selection-mode"] = types.YLeaf{"NetworkSelectionMode", cellwanNetwork.NetworkSelectionMode}
    cellwanNetwork.EntityData.Leafs["network-name"] = types.YLeaf{"NetworkName", cellwanNetwork.NetworkName}
    cellwanNetwork.EntityData.Leafs["mobile-country-code"] = types.YLeaf{"MobileCountryCode", cellwanNetwork.MobileCountryCode}
    cellwanNetwork.EntityData.Leafs["mobile-network-code"] = types.YLeaf{"MobileNetworkCode", cellwanNetwork.MobileNetworkCode}
    cellwanNetwork.EntityData.Leafs["packet-switch-domain-state"] = types.YLeaf{"PacketSwitchDomainState", cellwanNetwork.PacketSwitchDomainState}
    cellwanNetwork.EntityData.Leafs["lte-carrier-aggregation"] = types.YLeaf{"LteCarrierAggregation", cellwanNetwork.LteCarrierAggregation}
    cellwanNetwork.EntityData.Leafs["registration-state"] = types.YLeaf{"RegistrationState", cellwanNetwork.RegistrationState}
    cellwanNetwork.EntityData.Leafs["emm-state"] = types.YLeaf{"EmmState", cellwanNetwork.EmmState}
    cellwanNetwork.EntityData.Leafs["emm-substate"] = types.YLeaf{"EmmSubstate", cellwanNetwork.EmmSubstate}
    cellwanNetwork.EntityData.Leafs["location-area-code"] = types.YLeaf{"LocationAreaCode", cellwanNetwork.LocationAreaCode}
    cellwanNetwork.EntityData.Leafs["tracking-area-code"] = types.YLeaf{"TrackingAreaCode", cellwanNetwork.TrackingAreaCode}
    cellwanNetwork.EntityData.Leafs["cell-id"] = types.YLeaf{"CellId", cellwanNetwork.CellId}
    cellwanNetwork.EntityData.Leafs["negotiated-network-mtu"] = types.YLeaf{"NegotiatedNetworkMtu", cellwanNetwork.NegotiatedNetworkMtu}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv4Addr interface{}

    // Cellular interface IPv6 Address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Addr interface{}

    // Cellular interface primary IPv4 DNS Address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv4DnsPri interface{}

    // Cellular interface secondary IPv4 DNS Address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv4DnsSec interface{}

    // Cellular interface primary IPv6 DNS Address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6DnsPri interface{}

    // Cellular interface secondary IPv6 DNS Address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6DnsSec interface{}

    // Cellular interface profile IP address scope. The type is ProfileScope.
    Scope interface{}
}

func (cellwanConnection *CellwanOperData_CellwanConnection) GetEntityData() *types.CommonEntityData {
    cellwanConnection.EntityData.YFilter = cellwanConnection.YFilter
    cellwanConnection.EntityData.YangName = "cellwan-connection"
    cellwanConnection.EntityData.BundleName = "cisco_ios_xe"
    cellwanConnection.EntityData.ParentYangName = "cellwan-oper-data"
    cellwanConnection.EntityData.SegmentPath = "cellwan-connection" + "[cellular-interface='" + fmt.Sprintf("%v", cellwanConnection.CellularInterface) + "']"
    cellwanConnection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellwanConnection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellwanConnection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellwanConnection.EntityData.Children = make(map[string]types.YChild)
    cellwanConnection.EntityData.Leafs = make(map[string]types.YLeaf)
    cellwanConnection.EntityData.Leafs["cellular-interface"] = types.YLeaf{"CellularInterface", cellwanConnection.CellularInterface}
    cellwanConnection.EntityData.Leafs["active-profile"] = types.YLeaf{"ActiveProfile", cellwanConnection.ActiveProfile}
    cellwanConnection.EntityData.Leafs["cellular-packet-status"] = types.YLeaf{"CellularPacketStatus", cellwanConnection.CellularPacketStatus}
    cellwanConnection.EntityData.Leafs["tx-packets"] = types.YLeaf{"TxPackets", cellwanConnection.TxPackets}
    cellwanConnection.EntityData.Leafs["rx-packets"] = types.YLeaf{"RxPackets", cellwanConnection.RxPackets}
    cellwanConnection.EntityData.Leafs["tx-bytes"] = types.YLeaf{"TxBytes", cellwanConnection.TxBytes}
    cellwanConnection.EntityData.Leafs["rx-bytes"] = types.YLeaf{"RxBytes", cellwanConnection.RxBytes}
    cellwanConnection.EntityData.Leafs["ipv4-addr"] = types.YLeaf{"Ipv4Addr", cellwanConnection.Ipv4Addr}
    cellwanConnection.EntityData.Leafs["ipv6-addr"] = types.YLeaf{"Ipv6Addr", cellwanConnection.Ipv6Addr}
    cellwanConnection.EntityData.Leafs["ipv4-dns-pri"] = types.YLeaf{"Ipv4DnsPri", cellwanConnection.Ipv4DnsPri}
    cellwanConnection.EntityData.Leafs["ipv4-dns-sec"] = types.YLeaf{"Ipv4DnsSec", cellwanConnection.Ipv4DnsSec}
    cellwanConnection.EntityData.Leafs["ipv6-dns-pri"] = types.YLeaf{"Ipv6DnsPri", cellwanConnection.Ipv6DnsPri}
    cellwanConnection.EntityData.Leafs["ipv6-dns-sec"] = types.YLeaf{"Ipv6DnsSec", cellwanConnection.Ipv6DnsSec}
    cellwanConnection.EntityData.Leafs["scope"] = types.YLeaf{"Scope", cellwanConnection.Scope}
    return &(cellwanConnection.EntityData)
}

