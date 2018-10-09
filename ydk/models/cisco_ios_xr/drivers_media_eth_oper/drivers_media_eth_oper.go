// This module contains a collection of YANG definitions
// for Cisco IOS-XR drivers-media-eth package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ethernet-interface: Ethernet operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package drivers_media_eth_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package drivers_media_eth_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-drivers-media-eth-oper ethernet-interface}", reflect.TypeOf(EthernetInterface{}))
    ydk.RegisterEntity("Cisco-IOS-XR-drivers-media-eth-oper:ethernet-interface", reflect.TypeOf(EthernetInterface{}))
}

// EtherLinkState represents .5.1.1.4
type EtherLinkState string

const (
    // State undefined
    EtherLinkState_state_undefined EtherLinkState = "state-undefined"

    // Initializing, true state not yet known
    EtherLinkState_unknown_state EtherLinkState = "unknown-state"

    // Link or light normal, loopback normal
    EtherLinkState_available EtherLinkState = "available"

    // Link loss or low light, no loopback
    EtherLinkState_not_available EtherLinkState = "not-available"

    // Remote fault with no detail
    EtherLinkState_remote_fault EtherLinkState = "remote-fault"

    // Invalid signal, applies only to 10BASE-FB
    EtherLinkState_invalid_signal EtherLinkState = "invalid-signal"

    // Remote fault, reason known to be jabber
    EtherLinkState_remote_jabber EtherLinkState = "remote-jabber"

    // Remote fault, reason known to be far-end link
    // loss
    EtherLinkState_link_loss EtherLinkState = "link-loss"

    // Remote fault, reason known to be test
    EtherLinkState_remote_test EtherLinkState = "remote-test"

    // Offline (applies to auto-negotiation)
    EtherLinkState_offline EtherLinkState = "offline"

    // Auto-Negotiation Error
    EtherLinkState_auto_neg_error EtherLinkState = "auto-neg-error"

    // PMD/PMA receive link fault
    EtherLinkState_pmd_link_fault EtherLinkState = "pmd-link-fault"

    // WIS loss of frames
    EtherLinkState_frame_loss EtherLinkState = "frame-loss"

    // WIS loss of signal
    EtherLinkState_signal_loss EtherLinkState = "signal-loss"

    // PCS receive link fault
    EtherLinkState_link_fault EtherLinkState = "link-fault"

    // PCS Bit Error Rate monitor reporting excessive
    // error rate
    EtherLinkState_excessive_ber EtherLinkState = "excessive-ber"

    // DTE XGXS receive link fault
    EtherLinkState_dxs_link_fault EtherLinkState = "dxs-link-fault"

    // PHY XGXS transmit link fault
    EtherLinkState_pxs_link_fault EtherLinkState = "pxs-link-fault"

    // Security failure (not a valid part)
    EtherLinkState_security EtherLinkState = "security"

    // The optics for the port are not present
    EtherLinkState_phy_not_present EtherLinkState = "phy-not-present"

    // License error (No advanced optical license)
    EtherLinkState_no_optic_license EtherLinkState = "no-optic-license"

    // Module is not supported
    EtherLinkState_unsupported_module EtherLinkState = "unsupported-module"

    // DWDM Laser shutdown
    EtherLinkState_dwdm_laser_shut EtherLinkState = "dwdm-laser-shut"

    // WANPHY Laser shutdown
    EtherLinkState_wanphy_laser_shut EtherLinkState = "wanphy-laser-shut"

    // Incompatible configuration
    EtherLinkState_incompatible_config EtherLinkState = "incompatible-config"

    // System error
    EtherLinkState_system_error EtherLinkState = "system-error"

    // WAN Framing Error
    EtherLinkState_wan_framing_error EtherLinkState = "wan-framing-error"

    // OTN Framing Error
    EtherLinkState_otn_framing_error EtherLinkState = "otn-framing-error"

    // Link is shutdown
    EtherLinkState_shutdown EtherLinkState = "shutdown"
)

// EthernetBertPattern represents Ethernet test patterns (IEEE spec 36A/48A)
type EthernetBertPattern string

const (
    // no test pattern
    EthernetBertPattern_no_test_pattern EthernetBertPattern = "no-test-pattern"

    // high frequency
    EthernetBertPattern_high_frequency EthernetBertPattern = "high-frequency"

    // low frequency
    EthernetBertPattern_low_frequency EthernetBertPattern = "low-frequency"

    // mixed frequency
    EthernetBertPattern_mixed_frequency EthernetBertPattern = "mixed-frequency"

    // continuous random
    EthernetBertPattern_continuous_random EthernetBertPattern = "continuous-random"

    // continuous jitter
    EthernetBertPattern_continuous_jitter EthernetBertPattern = "continuous-jitter"

    // long continuous random
    EthernetBertPattern_long_continuous_random EthernetBertPattern = "long-continuous-random"

    // short continuous random
    EthernetBertPattern_short_continuous_random EthernetBertPattern = "short-continuous-random"

    // pseudorandom seed a
    EthernetBertPattern_pseudorandom_seed_a EthernetBertPattern = "pseudorandom-seed-a"

    // pseudorandom seed b
    EthernetBertPattern_pseudorandom_seed_b EthernetBertPattern = "pseudorandom-seed-b"

    // prbs31
    EthernetBertPattern_prbs31 EthernetBertPattern = "prbs31"

    // square wave
    EthernetBertPattern_square_wave EthernetBertPattern = "square-wave"

    // pseudorandom
    EthernetBertPattern_pseudorandom EthernetBertPattern = "pseudorandom"

    // ethernet bert pattern types
    EthernetBertPattern_ethernet_bert_pattern_types EthernetBertPattern = "ethernet-bert-pattern-types"
)

// EthernetPortEnable represents Port admin state
type EthernetPortEnable string

const (
    // Port disabled, both directions
    EthernetPortEnable_disabled EthernetPortEnable = "disabled"

    // Port enabled rx direction only
    EthernetPortEnable_rx_enabled EthernetPortEnable = "rx-enabled"

    // Port enabled tx direction only
    EthernetPortEnable_tx_enabled EthernetPortEnable = "tx-enabled"

    // Port enabled, both directions
    EthernetPortEnable_enabled EthernetPortEnable = "enabled"
)

// EthCtrlrAlarmState represents Ethernet alarm state
type EthCtrlrAlarmState string

const (
    // Not supported on this interface
    EthCtrlrAlarmState_alarm_not_supported EthCtrlrAlarmState = "alarm-not-supported"

    // Alarm set
    EthCtrlrAlarmState_alarm_set EthCtrlrAlarmState = "alarm-set"

    // Alarm not set
    EthCtrlrAlarmState_alarm_not_set EthCtrlrAlarmState = "alarm-not-set"
)

// EthernetDev represents Ethernet dev
type EthernetDev string

const (
    // no device
    EthernetDev_no_device EthernetDev = "no-device"

    // pma pmd
    EthernetDev_pma_pmd EthernetDev = "pma-pmd"

    // wis
    EthernetDev_wis EthernetDev = "wis"

    // pcs
    EthernetDev_pcs EthernetDev = "pcs"

    // phy xs
    EthernetDev_phy_xs EthernetDev = "phy-xs"

    // dte xs
    EthernetDev_dte_xs EthernetDev = "dte-xs"

    // ethernet num dev
    EthernetDev_ethernet_num_dev EthernetDev = "ethernet-num-dev"
)

// EtherPhyPresent represents Ether phy present
type EtherPhyPresent string

const (
    // No PHY present
    EtherPhyPresent_phy_not_present EtherPhyPresent = "phy-not-present"

    // PHY is present
    EtherPhyPresent_phy_present EtherPhyPresent = "phy-present"

    // State is unknown
    EtherPhyPresent_no_information EtherPhyPresent = "no-information"
)

// EthernetDevIf represents Ethernet dev if
type EthernetDevIf string

const (
    // no interface
    EthernetDevIf_no_interface EthernetDevIf = "no-interface"

    // xgmii
    EthernetDevIf_xgmii EthernetDevIf = "xgmii"

    // xaui
    EthernetDevIf_xaui EthernetDevIf = "xaui"

    // ethernet num dev if
    EthernetDevIf_ethernet_num_dev_if EthernetDevIf = "ethernet-num-dev-if"
)

// EtherFlowcontrol represents Flowcontrol type
type EtherFlowcontrol string

const (
    // No flow control (disabled)
    EtherFlowcontrol_no_flowcontrol EtherFlowcontrol = "no-flowcontrol"

    // Traffic egress (pause frames ingress)
    EtherFlowcontrol_egress EtherFlowcontrol = "egress"

    // Traffic ingress (pause frames egress)
    EtherFlowcontrol_ingress EtherFlowcontrol = "ingress"

    // On both ingress and egress
    EtherFlowcontrol_bidirectional EtherFlowcontrol = "bidirectional"
)

// EtherDomAlarm represents Ether dom alarm
type EtherDomAlarm string

const (
    // DOM Alarm information is not available
    EtherDomAlarm_no_information EtherDomAlarm = "no-information"

    // Alarm high
    EtherDomAlarm_alarm_high EtherDomAlarm = "alarm-high"

    // Warning high
    EtherDomAlarm_warning_high EtherDomAlarm = "warning-high"

    // Within normal parameters
    EtherDomAlarm_normal EtherDomAlarm = "normal"

    // Warning low
    EtherDomAlarm_warning_low EtherDomAlarm = "warning-low"

    // Alarm low
    EtherDomAlarm_alarm_low EtherDomAlarm = "alarm-low"
)

// EtherPfc represents Priority flowcontrol type
type EtherPfc string

const (
    // No priority flow control (disabled)
    EtherPfc_no_pfc EtherPfc = "no-pfc"

    // Priority flow control enabled
    EtherPfc_on EtherPfc = "on"
)

// EthernetBertErrCnt represents Ethernet bert err cnt
type EthernetBertErrCnt string

const (
    // no count type
    EthernetBertErrCnt_no_count_type EthernetBertErrCnt = "no-count-type"

    // bit error count
    EthernetBertErrCnt_bit_error_count EthernetBertErrCnt = "bit-error-count"

    // frame error count
    EthernetBertErrCnt_frame_error_count EthernetBertErrCnt = "frame-error-count"

    // block error count
    EthernetBertErrCnt_block_error_count EthernetBertErrCnt = "block-error-count"

    // ethernet bert err cnt types
    EthernetBertErrCnt_ethernet_bert_err_cnt_types EthernetBertErrCnt = "ethernet-bert-err-cnt-types"
)

// EthernetDuplex represents Duplexity
type EthernetDuplex string

const (
    // ethernet duplex invalid
    EthernetDuplex_ethernet_duplex_invalid EthernetDuplex = "ethernet-duplex-invalid"

    // half duplex
    EthernetDuplex_half_duplex EthernetDuplex = "half-duplex"

    // full duplex
    EthernetDuplex_full_duplex EthernetDuplex = "full-duplex"
)

// EthernetIpg represents Inter packet gap
type EthernetIpg string

const (
    // IEEE standard value of 12
    EthernetIpg_standard EthernetIpg = "standard"

    // Non-standard value of 16
    EthernetIpg_non_standard EthernetIpg = "non-standard"
)

// EthernetSpeed represents Speed
type EthernetSpeed string

const (
    // ethernet speed invalid
    EthernetSpeed_ethernet_speed_invalid EthernetSpeed = "ethernet-speed-invalid"

    // ten mbps
    EthernetSpeed_ten_mbps EthernetSpeed = "ten-mbps"

    // hundred mbps
    EthernetSpeed_hundred_mbps EthernetSpeed = "hundred-mbps"

    // one gbps
    EthernetSpeed_one_gbps EthernetSpeed = "one-gbps"

    // ten gbps
    EthernetSpeed_ten_gbps EthernetSpeed = "ten-gbps"

    // twenty five gbps
    EthernetSpeed_twenty_five_gbps EthernetSpeed = "twenty-five-gbps"

    // forty gbps
    EthernetSpeed_forty_gbps EthernetSpeed = "forty-gbps"

    // fifty gbps
    EthernetSpeed_fifty_gbps EthernetSpeed = "fifty-gbps"

    // hundred gbps
    EthernetSpeed_hundred_gbps EthernetSpeed = "hundred-gbps"

    // two hundred gbps
    EthernetSpeed_two_hundred_gbps EthernetSpeed = "two-hundred-gbps"

    // four hundred gbps
    EthernetSpeed_four_hundred_gbps EthernetSpeed = "four-hundred-gbps"

    // ethernet speed types count
    EthernetSpeed_ethernet_speed_types_count EthernetSpeed = "ethernet-speed-types-count"
)

// EtherLedState represents Ether led state
type EtherLedState string

const (
    // LED state is unknown
    EtherLedState_led_state_unknown EtherLedState = "led-state-unknown"

    // LED is off
    EtherLedState_led_off EtherLedState = "led-off"

    // LED is green
    EtherLedState_green_on EtherLedState = "green-on"

    // LED is flashing green
    EtherLedState_green_flashing EtherLedState = "green-flashing"

    // LED is yellow
    EtherLedState_yellow_on EtherLedState = "yellow-on"

    // LED is flashing yellow
    EtherLedState_yellow_flashing EtherLedState = "yellow-flashing"

    // LED is red
    EtherLedState_red_on EtherLedState = "red-on"

    // LED is flashing red
    EtherLedState_red_flashing EtherLedState = "red-flashing"
)

// EthernetFec represents FEC type
type EthernetFec string

const (
    // FEC not configured
    EthernetFec_not_configured EthernetFec = "not-configured"

    // Reed-Solomon encoding
    EthernetFec_standard EthernetFec = "standard"

    // FEC explicitly disabled
    EthernetFec_disabled EthernetFec = "disabled"

    // BASE-R encoding
    EthernetFec_base_r EthernetFec = "base-r"
)

// EthernetMedia represents 30.5.1.1.2
type EthernetMedia string

const (
    // IEEE 802.3/802.3ae clause 30.2.5
    EthernetMedia_ethernet_other EthernetMedia = "ethernet-other"

    // Initializing, true state or type not yet known
    EthernetMedia_ethernet_unknown EthernetMedia = "ethernet-unknown"

    // No internal MAU, view from AUI
    EthernetMedia_ethernet_aui EthernetMedia = "ethernet-aui"

    // Thick coax MAU
    EthernetMedia_ethernet_10base5 EthernetMedia = "ethernet-10base5"

    // FOIRL MAU as specified in 9.9
    EthernetMedia_ethernet_foirl EthernetMedia = "ethernet-foirl"

    // Thin coax MAU
    EthernetMedia_ethernet_10base2 EthernetMedia = "ethernet-10base2"

    // Broadband DTE MAU
    EthernetMedia_ethernet_10broad36 EthernetMedia = "ethernet-10broad36"

    // UTP MAU, duplexity unknown
    EthernetMedia_ethernet_10base EthernetMedia = "ethernet-10base"

    // UTP MAU, half duplex
    EthernetMedia_ethernet_10base_thd EthernetMedia = "ethernet-10base-thd"

    // UTP MAU, full duplex
    EthernetMedia_ethernet_10base_tfd EthernetMedia = "ethernet-10base-tfd"

    // Passive fiber MAU
    EthernetMedia_ethernet_10base_fp EthernetMedia = "ethernet-10base-fp"

    // Synchronous fiber MAU
    EthernetMedia_ethernet_10base_fb EthernetMedia = "ethernet-10base-fb"

    // Asynchronous fiber MAU, duplexity unknown
    EthernetMedia_ethernet_10base_fl EthernetMedia = "ethernet-10base-fl"

    // Asynchronous fiber MAU, half duplex
    EthernetMedia_ethernet_10base_flhd EthernetMedia = "ethernet-10base-flhd"

    // Asynchronous fiber MAU, full duplex
    EthernetMedia_ethernet_10base_flfd EthernetMedia = "ethernet-10base-flfd"

    // Four-pair Category 3 UTP
    EthernetMedia_ethernet_100base_t4 EthernetMedia = "ethernet-100base-t4"

    // Two-pair Category 5 UTP, duplexity unknown
    EthernetMedia_ethernet_100base_tx EthernetMedia = "ethernet-100base-tx"

    // Two-pair Category 5 UTP, half duplex
    EthernetMedia_ethernet_100base_txhd EthernetMedia = "ethernet-100base-txhd"

    // Two-pair Category 5 UTP, full duplex
    EthernetMedia_ethernet_100base_txfd EthernetMedia = "ethernet-100base-txfd"

    // X fiber over PMD, duplexity unknown
    EthernetMedia_ethernet_100base_fx EthernetMedia = "ethernet-100base-fx"

    // X fiber over PMD, half duplex
    EthernetMedia_ethernet_100base_fxhd EthernetMedia = "ethernet-100base-fxhd"

    // X fiber over PMD, full duplex
    EthernetMedia_ethernet_100base_fxfd EthernetMedia = "ethernet-100base-fxfd"

    // X fiber over PMD (40km), duplexity unknown
    EthernetMedia_ethernet_100base_ex EthernetMedia = "ethernet-100base-ex"

    // X fiber over PMD (40km), half duplex
    EthernetMedia_ethernet_100base_exhd EthernetMedia = "ethernet-100base-exhd"

    // X fiber over PMD (40km), full duplex
    EthernetMedia_ethernet_100base_exfd EthernetMedia = "ethernet-100base-exfd"

    // Two-pair Category 3 UTP, duplexity unknown
    EthernetMedia_ethernet_100base_t2 EthernetMedia = "ethernet-100base-t2"

    // Two-pair Category 3 UTP, half duplex
    EthernetMedia_ethernet_100base_t2hd EthernetMedia = "ethernet-100base-t2hd"

    // Two-pair Category 3 UTP, full duplex
    EthernetMedia_ethernet_100base_t2fd EthernetMedia = "ethernet-100base-t2fd"

    // X PCS/PMA, duplexity unknown
    EthernetMedia_ethernet_1000base_x EthernetMedia = "ethernet-1000base-x"

    // X 1000BASE-XHDX PCS/PMA, half duplex
    EthernetMedia_ethernet_1000base_xhd EthernetMedia = "ethernet-1000base-xhd"

    // X PCS/PMA, full duplex
    EthernetMedia_ethernet_1000base_xfd EthernetMedia = "ethernet-1000base-xfd"

    // X fiber over long-wl laser PMD, duplexity
    // unknown
    EthernetMedia_ethernet_1000base_lx EthernetMedia = "ethernet-1000base-lx"

    // X fiber over long-wl laser PMD, half duplex
    EthernetMedia_ethernet_1000base_lxhd EthernetMedia = "ethernet-1000base-lxhd"

    // X fiber over long-wl laser PMD, full duplex
    EthernetMedia_ethernet_1000base_lxfd EthernetMedia = "ethernet-1000base-lxfd"

    // X fiber over short-wl laser PMD, duplexity
    // unknown
    EthernetMedia_ethernet_1000base_sx EthernetMedia = "ethernet-1000base-sx"

    // X fiber over short-wl laser PMD, half duplex
    EthernetMedia_ethernet_1000base_sxhd EthernetMedia = "ethernet-1000base-sxhd"

    // X fiber over short-wl laser PMD, full duplex
    EthernetMedia_ethernet_1000base_sxfd EthernetMedia = "ethernet-1000base-sxfd"

    // X copper over 150-Ohm balanced PMD, duplexity
    // unknown
    EthernetMedia_ethernet_1000base_cx EthernetMedia = "ethernet-1000base-cx"

    // X copper over 150-Ohm balancedPMD, half duplex
    EthernetMedia_ethernet_1000base_cxhd EthernetMedia = "ethernet-1000base-cxhd"

    // X copper over 150-Ohm balancedPMD, full duplex
    EthernetMedia_ethernet_1000base_cxfd EthernetMedia = "ethernet-1000base-cxfd"

    // Four-pair Category 5 UTP PHY, duplexity unknown
    EthernetMedia_ethernet_1000base EthernetMedia = "ethernet-1000base"

    // Four-pair Category 5 UTP PHY, half duplex
    EthernetMedia_ethernet_1000base_thd EthernetMedia = "ethernet-1000base-thd"

    // Four-pair Category 5 UTP PHY, full duplex
    EthernetMedia_ethernet_1000base_tfd EthernetMedia = "ethernet-1000base-tfd"

    // X PCS/PMA 
    EthernetMedia_ethernet_10gbase_x EthernetMedia = "ethernet-10gbase-x"

    // X fiber over 4 lane 1310nm optics
    EthernetMedia_ethernet_10gbase_lx4 EthernetMedia = "ethernet-10gbase-lx4"

    // R PCS/PMA
    EthernetMedia_ethernet_10gbase_r EthernetMedia = "ethernet-10gbase-r"

    // R fiber over 1550nm optics
    EthernetMedia_ethernet_10gbase_er EthernetMedia = "ethernet-10gbase-er"

    // R fiber over 1310nm optics
    EthernetMedia_ethernet_10gbase_lr EthernetMedia = "ethernet-10gbase-lr"

    // R fiber over 850nm optics
    EthernetMedia_ethernet_10gbase_sr EthernetMedia = "ethernet-10gbase-sr"

    // W PCS/PMA
    EthernetMedia_ethernet_10gbase_w EthernetMedia = "ethernet-10gbase-w"

    // W fiber over 1550nm optics
    EthernetMedia_ethernet_10gbase_ew EthernetMedia = "ethernet-10gbase-ew"

    // W fiber over 1310nm optics
    EthernetMedia_ethernet_10gbase_lw EthernetMedia = "ethernet-10gbase-lw"

    // W fiber over 850nm optics
    EthernetMedia_ethernet_10gbase_sw EthernetMedia = "ethernet-10gbase-sw"

    // Single-mode fiber over 1550nm optics (Cisco)
    EthernetMedia_ethernet_1000base_zx EthernetMedia = "ethernet-1000base-zx"

    // CWDM with unknown wavelength optics
    EthernetMedia_ethernet_1000base_cwdm EthernetMedia = "ethernet-1000base-cwdm"

    // CWDM with 1470nm optics
    EthernetMedia_ethernet_1000base_cwdm_1470 EthernetMedia = "ethernet-1000base-cwdm-1470"

    // CWDM with 1490nm optics
    EthernetMedia_ethernet_1000base_cwdm_1490 EthernetMedia = "ethernet-1000base-cwdm-1490"

    // CWDM with 1510nm optics
    EthernetMedia_ethernet_1000base_cwdm_1510 EthernetMedia = "ethernet-1000base-cwdm-1510"

    // CWDM with 1530nm optics
    EthernetMedia_ethernet_1000base_cwdm_1530 EthernetMedia = "ethernet-1000base-cwdm-1530"

    // CWDM with 1550nm optics
    EthernetMedia_ethernet_1000base_cwdm_1550 EthernetMedia = "ethernet-1000base-cwdm-1550"

    // CWDM with 1570nm optics
    EthernetMedia_ethernet_1000base_cwdm_1570 EthernetMedia = "ethernet-1000base-cwdm-1570"

    // CWDM with 1590nm optics
    EthernetMedia_ethernet_1000base_cwdm_1590 EthernetMedia = "ethernet-1000base-cwdm-1590"

    // CWDM with 1610nm optics
    EthernetMedia_ethernet_1000base_cwdm_1610 EthernetMedia = "ethernet-1000base-cwdm-1610"

    // Cisco-defined, over 1550nm optics
    EthernetMedia_ethernet_10gbase_zr EthernetMedia = "ethernet-10gbase-zr"

    // DWDM optics
    EthernetMedia_ethernet_10gbase_dwdm EthernetMedia = "ethernet-10gbase-dwdm"

    // fiber over 4 lane optics (long reach)
    EthernetMedia_ethernet_100gbase_lr4 EthernetMedia = "ethernet-100gbase-lr4"

    // DWDM optics
    EthernetMedia_ethernet_1000base_dwdm EthernetMedia = "ethernet-1000base-dwdm"

    // DWDM with 1533nm optics
    EthernetMedia_ethernet_1000base_dwdm_1533 EthernetMedia = "ethernet-1000base-dwdm-1533"

    // DWDM with 1537nm optics
    EthernetMedia_ethernet_1000base_dwdm_1537 EthernetMedia = "ethernet-1000base-dwdm-1537"

    // DWDM with 1541nm optics
    EthernetMedia_ethernet_1000base_dwdm_1541 EthernetMedia = "ethernet-1000base-dwdm-1541"

    // DWDM with 1545nm optics
    EthernetMedia_ethernet_1000base_dwdm_1545 EthernetMedia = "ethernet-1000base-dwdm-1545"

    // DWDM with 1549nm optics
    EthernetMedia_ethernet_1000base_dwdm_1549 EthernetMedia = "ethernet-1000base-dwdm-1549"

    // DWDM with 1553nm optics
    EthernetMedia_ethernet_1000base_dwdm_1553 EthernetMedia = "ethernet-1000base-dwdm-1553"

    // DWDM with 1557nm optics
    EthernetMedia_ethernet_1000base_dwdm_1557 EthernetMedia = "ethernet-1000base-dwdm-1557"

    // DWDM with 1561nm optics
    EthernetMedia_ethernet_1000base_dwdm_1561 EthernetMedia = "ethernet-1000base-dwdm-1561"

    // fiber over 4 lane optics (long reach)
    EthernetMedia_ethernet_40gbase_lr4 EthernetMedia = "ethernet-40gbase-lr4"

    // fiber over 4 lane optics (extended reach)
    EthernetMedia_ethernet_40gbase_er4 EthernetMedia = "ethernet-40gbase-er4"

    // fiber over 4 lane optics (extended reach)
    EthernetMedia_ethernet_100gbase_er4 EthernetMedia = "ethernet-100gbase-er4"

    // X fiber over 1310nm optics
    EthernetMedia_ethernet_1000base_ex EthernetMedia = "ethernet-1000base-ex"

    // X fibre (D, 10km)
    EthernetMedia_ethernet_1000base_bx10_d EthernetMedia = "ethernet-1000base-bx10-d"

    // X fibre (U, 10km)
    EthernetMedia_ethernet_1000base_bx10_u EthernetMedia = "ethernet-1000base-bx10-u"

    // DWDM with 1561.42nm optics
    EthernetMedia_ethernet_1000base_dwdm_1561_42 EthernetMedia = "ethernet-1000base-dwdm-1561-42"

    // DWDM with 1560.61nm optics
    EthernetMedia_ethernet_1000base_dwdm_1560_61 EthernetMedia = "ethernet-1000base-dwdm-1560-61"

    // DWDM with 1559.79nm optics
    EthernetMedia_ethernet_1000base_dwdm_1559_79 EthernetMedia = "ethernet-1000base-dwdm-1559-79"

    // DWDM with 1558.98nm optics
    EthernetMedia_ethernet_1000base_dwdm_1558_98 EthernetMedia = "ethernet-1000base-dwdm-1558-98"

    // DWDM with 1558.17nm optics
    EthernetMedia_ethernet_1000base_dwdm_1558_17 EthernetMedia = "ethernet-1000base-dwdm-1558-17"

    // DWDM with 1557.36nm optics
    EthernetMedia_ethernet_1000base_dwdm_1557_36 EthernetMedia = "ethernet-1000base-dwdm-1557-36"

    // DWDM with 1556.55nm optics
    EthernetMedia_ethernet_1000base_dwdm_1556_55 EthernetMedia = "ethernet-1000base-dwdm-1556-55"

    // DWDM with 1555.75nm optics
    EthernetMedia_ethernet_1000base_dwdm_1555_75 EthernetMedia = "ethernet-1000base-dwdm-1555-75"

    // DWDM with 1554.94nm optics
    EthernetMedia_ethernet_1000base_dwdm_1554_94 EthernetMedia = "ethernet-1000base-dwdm-1554-94"

    // DWDM with 1554.13nm optics
    EthernetMedia_ethernet_1000base_dwdm_1554_13 EthernetMedia = "ethernet-1000base-dwdm-1554-13"

    // DWDM with 1553.33nm optics
    EthernetMedia_ethernet_1000base_dwdm_1553_33 EthernetMedia = "ethernet-1000base-dwdm-1553-33"

    // DWDM with 1552.52nm optics
    EthernetMedia_ethernet_1000base_dwdm_1552_52 EthernetMedia = "ethernet-1000base-dwdm-1552-52"

    // DWDM with 1551.72nm optics
    EthernetMedia_ethernet_1000base_dwdm_1551_72 EthernetMedia = "ethernet-1000base-dwdm-1551-72"

    // DWDM with 1550.92nm optics
    EthernetMedia_ethernet_1000base_dwdm_1550_92 EthernetMedia = "ethernet-1000base-dwdm-1550-92"

    // DWDM with 1550.12nm optics
    EthernetMedia_ethernet_1000base_dwdm_1550_12 EthernetMedia = "ethernet-1000base-dwdm-1550-12"

    // DWDM with 1549.32nm optics
    EthernetMedia_ethernet_1000base_dwdm_1549_32 EthernetMedia = "ethernet-1000base-dwdm-1549-32"

    // DWDM with 1548.51nm optics
    EthernetMedia_ethernet_1000base_dwdm_1548_51 EthernetMedia = "ethernet-1000base-dwdm-1548-51"

    // DWDM with 1547.72nm optics
    EthernetMedia_ethernet_1000base_dwdm_1547_72 EthernetMedia = "ethernet-1000base-dwdm-1547-72"

    // DWDM with 1546.92nm optics
    EthernetMedia_ethernet_1000base_dwdm_1546_92 EthernetMedia = "ethernet-1000base-dwdm-1546-92"

    // DWDM with 1546.12nm optics
    EthernetMedia_ethernet_1000base_dwdm_1546_12 EthernetMedia = "ethernet-1000base-dwdm-1546-12"

    // DWDM with 1545.32nm optics
    EthernetMedia_ethernet_1000base_dwdm_1545_32 EthernetMedia = "ethernet-1000base-dwdm-1545-32"

    // DWDM with 1544.53nm optics
    EthernetMedia_ethernet_1000base_dwdm_1544_53 EthernetMedia = "ethernet-1000base-dwdm-1544-53"

    // DWDM with 1543.73nm optics
    EthernetMedia_ethernet_1000base_dwdm_1543_73 EthernetMedia = "ethernet-1000base-dwdm-1543-73"

    // DWDM with 1542.94nm optics
    EthernetMedia_ethernet_1000base_dwdm_1542_94 EthernetMedia = "ethernet-1000base-dwdm-1542-94"

    // DWDM with 1542.14nm optics
    EthernetMedia_ethernet_1000base_dwdm_1542_14 EthernetMedia = "ethernet-1000base-dwdm-1542-14"

    // DWDM with 1541.35nm optics
    EthernetMedia_ethernet_1000base_dwdm_1541_35 EthernetMedia = "ethernet-1000base-dwdm-1541-35"

    // DWDM with 1540.56nm optics
    EthernetMedia_ethernet_1000base_dwdm_1540_56 EthernetMedia = "ethernet-1000base-dwdm-1540-56"

    // DWDM with 1539.77nm optics
    EthernetMedia_ethernet_1000base_dwdm_1539_77 EthernetMedia = "ethernet-1000base-dwdm-1539-77"

    // DWDM with 1538.98nm optics
    EthernetMedia_ethernet_1000base_dwdm_1538_98 EthernetMedia = "ethernet-1000base-dwdm-1538-98"

    // DWDM with 1538.19nm optics
    EthernetMedia_ethernet_1000base_dwdm_1538_19 EthernetMedia = "ethernet-1000base-dwdm-1538-19"

    // DWDM with 1537.40nm optics
    EthernetMedia_ethernet_1000base_dwdm_1537_40 EthernetMedia = "ethernet-1000base-dwdm-1537-40"

    // DWDM with 1536.61nm optics
    EthernetMedia_ethernet_1000base_dwdm_1536_61 EthernetMedia = "ethernet-1000base-dwdm-1536-61"

    // DWDM with 1535.82nm optics
    EthernetMedia_ethernet_1000base_dwdm_1535_82 EthernetMedia = "ethernet-1000base-dwdm-1535-82"

    // DWDM with 1535.04nm optics
    EthernetMedia_ethernet_1000base_dwdm_1535_04 EthernetMedia = "ethernet-1000base-dwdm-1535-04"

    // DWDM with 1534.25nm optics
    EthernetMedia_ethernet_1000base_dwdm_1534_25 EthernetMedia = "ethernet-1000base-dwdm-1534-25"

    // DWDM with 1533.47nm optics
    EthernetMedia_ethernet_1000base_dwdm_1533_47 EthernetMedia = "ethernet-1000base-dwdm-1533-47"

    // DWDM with 1532.68nm optics
    EthernetMedia_ethernet_1000base_dwdm_1532_68 EthernetMedia = "ethernet-1000base-dwdm-1532-68"

    // DWDM with 1531.90nm optics
    EthernetMedia_ethernet_1000base_dwdm_1531_90 EthernetMedia = "ethernet-1000base-dwdm-1531-90"

    // DWDM with 1531.12nm optics
    EthernetMedia_ethernet_1000base_dwdm_1531_12 EthernetMedia = "ethernet-1000base-dwdm-1531-12"

    // DWDM with 1530.33nm optics
    EthernetMedia_ethernet_1000base_dwdm_1530_33 EthernetMedia = "ethernet-1000base-dwdm-1530-33"

    // DWDM with tunable optics
    EthernetMedia_ethernet_1000base_dwdm_tunable EthernetMedia = "ethernet-1000base-dwdm-tunable"

    // DWDM with 1561.42nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1561_42 EthernetMedia = "ethernet-10gbase-dwdm-1561-42"

    // DWDM with 1560.61nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1560_61 EthernetMedia = "ethernet-10gbase-dwdm-1560-61"

    // DWDM with 1559.79nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1559_79 EthernetMedia = "ethernet-10gbase-dwdm-1559-79"

    // DWDM with 1558.98nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1558_98 EthernetMedia = "ethernet-10gbase-dwdm-1558-98"

    // DWDM with 1558.17nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1558_17 EthernetMedia = "ethernet-10gbase-dwdm-1558-17"

    // DWDM with 1557.36nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1557_36 EthernetMedia = "ethernet-10gbase-dwdm-1557-36"

    // DWDM with 1556.55nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1556_55 EthernetMedia = "ethernet-10gbase-dwdm-1556-55"

    // DWDM with 1555.75nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1555_75 EthernetMedia = "ethernet-10gbase-dwdm-1555-75"

    // DWDM with 1554.94nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1554_94 EthernetMedia = "ethernet-10gbase-dwdm-1554-94"

    // DWDM with 1554.13nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1554_13 EthernetMedia = "ethernet-10gbase-dwdm-1554-13"

    // DWDM with 1553.33nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1553_33 EthernetMedia = "ethernet-10gbase-dwdm-1553-33"

    // DWDM with 1552.52nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1552_52 EthernetMedia = "ethernet-10gbase-dwdm-1552-52"

    // DWDM with 1551.72nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1551_72 EthernetMedia = "ethernet-10gbase-dwdm-1551-72"

    // DWDM with 1550.92nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1550_92 EthernetMedia = "ethernet-10gbase-dwdm-1550-92"

    // DWDM with 1550.12nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1550_12 EthernetMedia = "ethernet-10gbase-dwdm-1550-12"

    // DWDM with 1549.32nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1549_32 EthernetMedia = "ethernet-10gbase-dwdm-1549-32"

    // DWDM with 1548.51nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1548_51 EthernetMedia = "ethernet-10gbase-dwdm-1548-51"

    // DWDM with 1547.72nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1547_72 EthernetMedia = "ethernet-10gbase-dwdm-1547-72"

    // DWDM with 1546.92nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1546_92 EthernetMedia = "ethernet-10gbase-dwdm-1546-92"

    // DWDM with 1546.12nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1546_12 EthernetMedia = "ethernet-10gbase-dwdm-1546-12"

    // DWDM with 1545.32nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1545_32 EthernetMedia = "ethernet-10gbase-dwdm-1545-32"

    // DWDM with 1544.53nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1544_53 EthernetMedia = "ethernet-10gbase-dwdm-1544-53"

    // DWDM with 1543.73nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1543_73 EthernetMedia = "ethernet-10gbase-dwdm-1543-73"

    // DWDM with 1542.94nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1542_94 EthernetMedia = "ethernet-10gbase-dwdm-1542-94"

    // DWDM with 1542.14nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1542_14 EthernetMedia = "ethernet-10gbase-dwdm-1542-14"

    // DWDM with 1541.35nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1541_35 EthernetMedia = "ethernet-10gbase-dwdm-1541-35"

    // DWDM with 1540.56nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1540_56 EthernetMedia = "ethernet-10gbase-dwdm-1540-56"

    // DWDM with 1539.77nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1539_77 EthernetMedia = "ethernet-10gbase-dwdm-1539-77"

    // DWDM with 1538.98nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1538_98 EthernetMedia = "ethernet-10gbase-dwdm-1538-98"

    // DWDM with 1538.19nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1538_19 EthernetMedia = "ethernet-10gbase-dwdm-1538-19"

    // DWDM with 1537.40nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1537_40 EthernetMedia = "ethernet-10gbase-dwdm-1537-40"

    // DWDM with 1536.61nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1536_61 EthernetMedia = "ethernet-10gbase-dwdm-1536-61"

    // DWDM with 1535.82nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1535_82 EthernetMedia = "ethernet-10gbase-dwdm-1535-82"

    // DWDM with 1535.04nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1535_04 EthernetMedia = "ethernet-10gbase-dwdm-1535-04"

    // DWDM with 1534.25nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1534_25 EthernetMedia = "ethernet-10gbase-dwdm-1534-25"

    // DWDM with 1533.47nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1533_47 EthernetMedia = "ethernet-10gbase-dwdm-1533-47"

    // DWDM with 1532.68nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1532_68 EthernetMedia = "ethernet-10gbase-dwdm-1532-68"

    // DWDM with 1531.90nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1531_90 EthernetMedia = "ethernet-10gbase-dwdm-1531-90"

    // DWDM with 1531.12nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1531_12 EthernetMedia = "ethernet-10gbase-dwdm-1531-12"

    // DWDM with 1530.33nm optics
    EthernetMedia_ethernet_10gbase_dwdm_1530_33 EthernetMedia = "ethernet-10gbase-dwdm-1530-33"

    // DWDM with tunable optics
    EthernetMedia_ethernet_10gbase_dwdm_tunable EthernetMedia = "ethernet-10gbase-dwdm-tunable"

    // DWDM with 1561.42nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1561_42 EthernetMedia = "ethernet-40gbase-dwdm-1561-42"

    // DWDM with 1560.61nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1560_61 EthernetMedia = "ethernet-40gbase-dwdm-1560-61"

    // DWDM with 1559.79nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1559_79 EthernetMedia = "ethernet-40gbase-dwdm-1559-79"

    // DWDM with 1558.98nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1558_98 EthernetMedia = "ethernet-40gbase-dwdm-1558-98"

    // DWDM with 1558.17nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1558_17 EthernetMedia = "ethernet-40gbase-dwdm-1558-17"

    // DWDM with 1557.36nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1557_36 EthernetMedia = "ethernet-40gbase-dwdm-1557-36"

    // DWDM with 1556.55nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1556_55 EthernetMedia = "ethernet-40gbase-dwdm-1556-55"

    // DWDM with 1555.75nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1555_75 EthernetMedia = "ethernet-40gbase-dwdm-1555-75"

    // DWDM with 1554.94nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1554_94 EthernetMedia = "ethernet-40gbase-dwdm-1554-94"

    // DWDM with 1554.13nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1554_13 EthernetMedia = "ethernet-40gbase-dwdm-1554-13"

    // DWDM with 1553.33nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1553_33 EthernetMedia = "ethernet-40gbase-dwdm-1553-33"

    // DWDM with 1552.52nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1552_52 EthernetMedia = "ethernet-40gbase-dwdm-1552-52"

    // DWDM with 1551.72nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1551_72 EthernetMedia = "ethernet-40gbase-dwdm-1551-72"

    // DWDM with 1550.92nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1550_92 EthernetMedia = "ethernet-40gbase-dwdm-1550-92"

    // DWDM with 1550.12nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1550_12 EthernetMedia = "ethernet-40gbase-dwdm-1550-12"

    // DWDM with 1549.32nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1549_32 EthernetMedia = "ethernet-40gbase-dwdm-1549-32"

    // DWDM with 1548.51nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1548_51 EthernetMedia = "ethernet-40gbase-dwdm-1548-51"

    // DWDM with 1547.72nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1547_72 EthernetMedia = "ethernet-40gbase-dwdm-1547-72"

    // DWDM with 1546.92nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1546_92 EthernetMedia = "ethernet-40gbase-dwdm-1546-92"

    // DWDM with 1546.12nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1546_12 EthernetMedia = "ethernet-40gbase-dwdm-1546-12"

    // DWDM with 1545.32nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1545_32 EthernetMedia = "ethernet-40gbase-dwdm-1545-32"

    // DWDM with 1544.53nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1544_53 EthernetMedia = "ethernet-40gbase-dwdm-1544-53"

    // DWDM with 1543.73nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1543_73 EthernetMedia = "ethernet-40gbase-dwdm-1543-73"

    // DWDM with 1542.94nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1542_94 EthernetMedia = "ethernet-40gbase-dwdm-1542-94"

    // DWDM with 1542.14nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1542_14 EthernetMedia = "ethernet-40gbase-dwdm-1542-14"

    // DWDM with 1541.35nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1541_35 EthernetMedia = "ethernet-40gbase-dwdm-1541-35"

    // DWDM with 1540.56nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1540_56 EthernetMedia = "ethernet-40gbase-dwdm-1540-56"

    // DWDM with 1539.77nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1539_77 EthernetMedia = "ethernet-40gbase-dwdm-1539-77"

    // DWDM with 1538.98nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1538_98 EthernetMedia = "ethernet-40gbase-dwdm-1538-98"

    // DWDM with 1538.19nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1538_19 EthernetMedia = "ethernet-40gbase-dwdm-1538-19"

    // DWDM with 1537.40nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1537_40 EthernetMedia = "ethernet-40gbase-dwdm-1537-40"

    // DWDM with 1536.61nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1536_61 EthernetMedia = "ethernet-40gbase-dwdm-1536-61"

    // DWDM with 1535.82nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1535_82 EthernetMedia = "ethernet-40gbase-dwdm-1535-82"

    // DWDM with 1535.04nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1535_04 EthernetMedia = "ethernet-40gbase-dwdm-1535-04"

    // DWDM with 1534.25nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1534_25 EthernetMedia = "ethernet-40gbase-dwdm-1534-25"

    // DWDM with 1533.47nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1533_47 EthernetMedia = "ethernet-40gbase-dwdm-1533-47"

    // DWDM with 1532.68nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1532_68 EthernetMedia = "ethernet-40gbase-dwdm-1532-68"

    // DWDM with 1531.90nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1531_90 EthernetMedia = "ethernet-40gbase-dwdm-1531-90"

    // DWDM with 1531.12nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1531_12 EthernetMedia = "ethernet-40gbase-dwdm-1531-12"

    // DWDM with 1530.33nm optics
    EthernetMedia_ethernet_40gbase_dwdm_1530_33 EthernetMedia = "ethernet-40gbase-dwdm-1530-33"

    // DWDM with tunable optics
    EthernetMedia_ethernet_40gbase_dwdm_tunable EthernetMedia = "ethernet-40gbase-dwdm-tunable"

    // DWDM with 1561.42nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1561_42 EthernetMedia = "ethernet-100gbase-dwdm-1561-42"

    // DWDM with 1560.61nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1560_61 EthernetMedia = "ethernet-100gbase-dwdm-1560-61"

    // DWDM with 1559.79nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1559_79 EthernetMedia = "ethernet-100gbase-dwdm-1559-79"

    // DWDM with 1558.98nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1558_98 EthernetMedia = "ethernet-100gbase-dwdm-1558-98"

    // DWDM with 1558.17nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1558_17 EthernetMedia = "ethernet-100gbase-dwdm-1558-17"

    // DWDM with 1557.36nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1557_36 EthernetMedia = "ethernet-100gbase-dwdm-1557-36"

    // DWDM with 1556.55nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1556_55 EthernetMedia = "ethernet-100gbase-dwdm-1556-55"

    // DWDM with 1555.75nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1555_75 EthernetMedia = "ethernet-100gbase-dwdm-1555-75"

    // DWDM with 1554.94nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1554_94 EthernetMedia = "ethernet-100gbase-dwdm-1554-94"

    // DWDM with 1554.13nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1554_13 EthernetMedia = "ethernet-100gbase-dwdm-1554-13"

    // DWDM with 1553.33nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1553_33 EthernetMedia = "ethernet-100gbase-dwdm-1553-33"

    // DWDM with 1552.52nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1552_52 EthernetMedia = "ethernet-100gbase-dwdm-1552-52"

    // DWDM with 1551.72nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1551_72 EthernetMedia = "ethernet-100gbase-dwdm-1551-72"

    // DWDM with 1550.92nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1550_92 EthernetMedia = "ethernet-100gbase-dwdm-1550-92"

    // DWDM with 1550.12nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1550_12 EthernetMedia = "ethernet-100gbase-dwdm-1550-12"

    // DWDM with 1549.32nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1549_32 EthernetMedia = "ethernet-100gbase-dwdm-1549-32"

    // DWDM with 1548.51nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1548_51 EthernetMedia = "ethernet-100gbase-dwdm-1548-51"

    // DWDM with 1547.72nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1547_72 EthernetMedia = "ethernet-100gbase-dwdm-1547-72"

    // DWDM with 1546.92nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1546_92 EthernetMedia = "ethernet-100gbase-dwdm-1546-92"

    // DWDM with 1546.12nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1546_12 EthernetMedia = "ethernet-100gbase-dwdm-1546-12"

    // DWDM with 1545.32nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1545_32 EthernetMedia = "ethernet-100gbase-dwdm-1545-32"

    // DWDM with 1544.53nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1544_53 EthernetMedia = "ethernet-100gbase-dwdm-1544-53"

    // DWDM with 1543.73nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1543_73 EthernetMedia = "ethernet-100gbase-dwdm-1543-73"

    // DWDM with 1542.94nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1542_94 EthernetMedia = "ethernet-100gbase-dwdm-1542-94"

    // DWDM with 1542.14nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1542_14 EthernetMedia = "ethernet-100gbase-dwdm-1542-14"

    // DWDM with 1541.35nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1541_35 EthernetMedia = "ethernet-100gbase-dwdm-1541-35"

    // DWDM with 1540.56nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1540_56 EthernetMedia = "ethernet-100gbase-dwdm-1540-56"

    // DWDM with 1539.77nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1539_77 EthernetMedia = "ethernet-100gbase-dwdm-1539-77"

    // DWDM with 1538.98nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1538_98 EthernetMedia = "ethernet-100gbase-dwdm-1538-98"

    // DWDM with 1538.19nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1538_19 EthernetMedia = "ethernet-100gbase-dwdm-1538-19"

    // DWDM with 1537.40nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1537_40 EthernetMedia = "ethernet-100gbase-dwdm-1537-40"

    // DWDM with 1536.61nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1536_61 EthernetMedia = "ethernet-100gbase-dwdm-1536-61"

    // DWDM with 1535.82nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1535_82 EthernetMedia = "ethernet-100gbase-dwdm-1535-82"

    // DWDM with 1535.04nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1535_04 EthernetMedia = "ethernet-100gbase-dwdm-1535-04"

    // DWDM with 1534.25nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1534_25 EthernetMedia = "ethernet-100gbase-dwdm-1534-25"

    // DWDM with 1533.47nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1533_47 EthernetMedia = "ethernet-100gbase-dwdm-1533-47"

    // DWDM with 1532.68nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1532_68 EthernetMedia = "ethernet-100gbase-dwdm-1532-68"

    // DWDM with 1531.90nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1531_90 EthernetMedia = "ethernet-100gbase-dwdm-1531-90"

    // DWDM with 1531.12nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1531_12 EthernetMedia = "ethernet-100gbase-dwdm-1531-12"

    // DWDM with 1530.33nm optics
    EthernetMedia_ethernet_100gbase_dwdm_1530_33 EthernetMedia = "ethernet-100gbase-dwdm-1530-33"

    // DWDM with tunable optics
    EthernetMedia_ethernet_100gbase_dwdm_tunable EthernetMedia = "ethernet-100gbase-dwdm-tunable"

    // 4 lane copper (backplane)
    EthernetMedia_ethernet_40gbase_kr4 EthernetMedia = "ethernet-40gbase-kr4"

    // 4 lane copper (very short reach)
    EthernetMedia_ethernet_40gbase_cr4 EthernetMedia = "ethernet-40gbase-cr4"

    // fiber over 4 lane optics (short reach)
    EthernetMedia_ethernet_40gbase_sr4 EthernetMedia = "ethernet-40gbase-sr4"

    // serial fiber (2+ km)
    EthernetMedia_ethernet_40gbase_fr EthernetMedia = "ethernet-40gbase-fr"

    // 10 lane copper (very short reach)
    EthernetMedia_ethernet_100gbase_cr10 EthernetMedia = "ethernet-100gbase-cr10"

    // MMF fiber over 10 lane optics (short reach)
    EthernetMedia_ethernet_100gbase_sr10 EthernetMedia = "ethernet-100gbase-sr10"

    // fiber over 4 lane optics (extended short reach)
    EthernetMedia_ethernet_40gbase_csr4 EthernetMedia = "ethernet-40gbase-csr4"

    // CWDM optics
    EthernetMedia_ethernet_10gbase_cwdm EthernetMedia = "ethernet-10gbase-cwdm"

    // CWDM with tunable optics
    EthernetMedia_ethernet_10gbase_cwdm_tunable EthernetMedia = "ethernet-10gbase-cwdm-tunable"

    // CWDM with 1470nm optics
    EthernetMedia_ethernet_10gbase_cwdm_1470 EthernetMedia = "ethernet-10gbase-cwdm-1470"

    // CWDM with 1490nm optics
    EthernetMedia_ethernet_10gbase_cwdm_1490 EthernetMedia = "ethernet-10gbase-cwdm-1490"

    // CWDM with 1510nm optics
    EthernetMedia_ethernet_10gbase_cwdm_1510 EthernetMedia = "ethernet-10gbase-cwdm-1510"

    // CWDM with 1530nm optics
    EthernetMedia_ethernet_10gbase_cwdm_1530 EthernetMedia = "ethernet-10gbase-cwdm-1530"

    // CWDM with 1550nm optics
    EthernetMedia_ethernet_10gbase_cwdm_1550 EthernetMedia = "ethernet-10gbase-cwdm-1550"

    // CWDM with 1570nm optics
    EthernetMedia_ethernet_10gbase_cwdm_1570 EthernetMedia = "ethernet-10gbase-cwdm-1570"

    // CWDM with 1590nm optics
    EthernetMedia_ethernet_10gbase_cwdm_1590 EthernetMedia = "ethernet-10gbase-cwdm-1590"

    // CWDM with 1610nm optics
    EthernetMedia_ethernet_10gbase_cwdm_1610 EthernetMedia = "ethernet-10gbase-cwdm-1610"

    // CWDM optics
    EthernetMedia_ethernet_40gbase_cwdm EthernetMedia = "ethernet-40gbase-cwdm"

    // CWDM with tunable optics
    EthernetMedia_ethernet_40gbase_cwdm_tunable EthernetMedia = "ethernet-40gbase-cwdm-tunable"

    // CWDM with 1470nm optics
    EthernetMedia_ethernet_40gbase_cwdm_1470 EthernetMedia = "ethernet-40gbase-cwdm-1470"

    // CWDM with 1490nm optics
    EthernetMedia_ethernet_40gbase_cwdm_1490 EthernetMedia = "ethernet-40gbase-cwdm-1490"

    // CWDM with 1510nm optics
    EthernetMedia_ethernet_40gbase_cwdm_1510 EthernetMedia = "ethernet-40gbase-cwdm-1510"

    // CWDM with 1530nm optics
    EthernetMedia_ethernet_40gbase_cwdm_1530 EthernetMedia = "ethernet-40gbase-cwdm-1530"

    // CWDM with 1550nm optics
    EthernetMedia_ethernet_40gbase_cwdm_1550 EthernetMedia = "ethernet-40gbase-cwdm-1550"

    // CWDM with 1570nm optics
    EthernetMedia_ethernet_40gbase_cwdm_1570 EthernetMedia = "ethernet-40gbase-cwdm-1570"

    // CWDM with 1590nm optics
    EthernetMedia_ethernet_40gbase_cwdm_1590 EthernetMedia = "ethernet-40gbase-cwdm-1590"

    // CWDM with 1610nm optics
    EthernetMedia_ethernet_40gbase_cwdm_1610 EthernetMedia = "ethernet-40gbase-cwdm-1610"

    // CWDM optics
    EthernetMedia_ethernet_100gbase_cwdm EthernetMedia = "ethernet-100gbase-cwdm"

    // CWDM with tunable optics
    EthernetMedia_ethernet_100gbase_cwdm_tunable EthernetMedia = "ethernet-100gbase-cwdm-tunable"

    // CWDM with 1470nm optics
    EthernetMedia_ethernet_100gbase_cwdm_1470 EthernetMedia = "ethernet-100gbase-cwdm-1470"

    // CWDM with 1490nm optics
    EthernetMedia_ethernet_100gbase_cwdm_1490 EthernetMedia = "ethernet-100gbase-cwdm-1490"

    // CWDM with 1510nm optics
    EthernetMedia_ethernet_100gbase_cwdm_1510 EthernetMedia = "ethernet-100gbase-cwdm-1510"

    // CWDM with 1530nm optics
    EthernetMedia_ethernet_100gbase_cwdm_1530 EthernetMedia = "ethernet-100gbase-cwdm-1530"

    // CWDM with 1550nm optics
    EthernetMedia_ethernet_100gbase_cwdm_1550 EthernetMedia = "ethernet-100gbase-cwdm-1550"

    // CWDM with 1570nm optics
    EthernetMedia_ethernet_100gbase_cwdm_1570 EthernetMedia = "ethernet-100gbase-cwdm-1570"

    // CWDM with 1590nm optics
    EthernetMedia_ethernet_100gbase_cwdm_1590 EthernetMedia = "ethernet-100gbase-cwdm-1590"

    // CWDM with 1610nm optics
    EthernetMedia_ethernet_100gbase_cwdm_1610 EthernetMedia = "ethernet-100gbase-cwdm-1610"

    // Electrical loopback
    EthernetMedia_ethernet_40gbase_elpb EthernetMedia = "ethernet-40gbase-elpb"

    // Electrical loopback
    EthernetMedia_ethernet_100gbase_elpb EthernetMedia = "ethernet-100gbase-elpb"

    // Fiber over 10 lane optics (long reach)
    EthernetMedia_ethernet_100gbase_lr10 EthernetMedia = "ethernet-100gbase-lr10"

    // Four-pair Category 8 STP
    EthernetMedia_ethernet_40gbase EthernetMedia = "ethernet-40gbase"

    // 4 lane copper (backplane)
    EthernetMedia_ethernet_100gbase_kp4 EthernetMedia = "ethernet-100gbase-kp4"

    // Improved 4 lane copper (backplane)
    EthernetMedia_ethernet_100gbase_kr4 EthernetMedia = "ethernet-100gbase-kr4"

    // Multimode fiber with 1310nm optics (long reach)
    EthernetMedia_ethernet_10gbase_lrm EthernetMedia = "ethernet-10gbase-lrm"

    // 4 lane X copper
    EthernetMedia_ethernet_10gbase_cx4 EthernetMedia = "ethernet-10gbase-cx4"

    // Four-pair Category 6+ UTP
    EthernetMedia_ethernet_10gbase EthernetMedia = "ethernet-10gbase"

    // 4 lane X copper (backplane)
    EthernetMedia_ethernet_10gbase_kx4 EthernetMedia = "ethernet-10gbase-kx4"

    // Copper (backplane)
    EthernetMedia_ethernet_10gbase_kr EthernetMedia = "ethernet-10gbase-kr"

    // Passive optical network
    EthernetMedia_ethernet_10gbase_pr EthernetMedia = "ethernet-10gbase-pr"

    // X fiber over 4 lane 1310nm optics
    EthernetMedia_ethernet_100base_lx EthernetMedia = "ethernet-100base-lx"

    // Single-mode fiber over 1550nm optics (Cisco)
    EthernetMedia_ethernet_100base_zx EthernetMedia = "ethernet-100base-zx"

    // X fibre (D)
    EthernetMedia_ethernet_1000base_bx_d EthernetMedia = "ethernet-1000base-bx-d"

    // X fibre (U)
    EthernetMedia_ethernet_1000base_bx_u EthernetMedia = "ethernet-1000base-bx-u"

    // X fibre (D, 20km)
    EthernetMedia_ethernet_1000base_bx20_d EthernetMedia = "ethernet-1000base-bx20-d"

    // X fibre (U, 20km)
    EthernetMedia_ethernet_1000base_bx20_u EthernetMedia = "ethernet-1000base-bx20-u"

    // X fibre (D, 40km)
    EthernetMedia_ethernet_1000base_bx40_d EthernetMedia = "ethernet-1000base-bx40-d"

    // X fibre (D, 40km)
    EthernetMedia_ethernet_1000base_bx40_da EthernetMedia = "ethernet-1000base-bx40-da"

    // X fibre (U, 40km)
    EthernetMedia_ethernet_1000base_bx40_u EthernetMedia = "ethernet-1000base-bx40-u"

    // X fibre (D, 80km)
    EthernetMedia_ethernet_1000base_bx80_d EthernetMedia = "ethernet-1000base-bx80-d"

    // X fibre (U, 80km)
    EthernetMedia_ethernet_1000base_bx80_u EthernetMedia = "ethernet-1000base-bx80-u"

    // X fibre (D, 120km)
    EthernetMedia_ethernet_1000base_bx120_d EthernetMedia = "ethernet-1000base-bx120-d"

    // X fibre (U, 120km)
    EthernetMedia_ethernet_1000base_bx120_u EthernetMedia = "ethernet-1000base-bx120-u"

    // X fibre (D)
    EthernetMedia_ethernet_10gbase_bx_d EthernetMedia = "ethernet-10gbase-bx-d"

    // X fibre (U)
    EthernetMedia_ethernet_10gbase_bx_u EthernetMedia = "ethernet-10gbase-bx-u"

    // X fibre (D, 10km)
    EthernetMedia_ethernet_10gbase_bx10_d EthernetMedia = "ethernet-10gbase-bx10-d"

    // X fibre (U, 10km)
    EthernetMedia_ethernet_10gbase_bx10_u EthernetMedia = "ethernet-10gbase-bx10-u"

    // X fibre (D, 20km)
    EthernetMedia_ethernet_10gbase_bx20_d EthernetMedia = "ethernet-10gbase-bx20-d"

    // X fibre (U, 20km)
    EthernetMedia_ethernet_10gbase_bx20_u EthernetMedia = "ethernet-10gbase-bx20-u"

    // X fibre (D, 40km)
    EthernetMedia_ethernet_10gbase_bx40_d EthernetMedia = "ethernet-10gbase-bx40-d"

    // X fibre (U, 40km)
    EthernetMedia_ethernet_10gbase_bx40_u EthernetMedia = "ethernet-10gbase-bx40-u"

    // X fibre (D, 80km)
    EthernetMedia_ethernet_10gbase_bx80_d EthernetMedia = "ethernet-10gbase-bx80-d"

    // X fibre (U, 80km)
    EthernetMedia_ethernet_10gbase_bx80_u EthernetMedia = "ethernet-10gbase-bx80-u"

    // X fibre (D, 120km)
    EthernetMedia_ethernet_10gbase_bx120_d EthernetMedia = "ethernet-10gbase-bx120-d"

    // X fibre (U, 120km)
    EthernetMedia_ethernet_10gbase_bx120_u EthernetMedia = "ethernet-10gbase-bx120-u"

    // X fiber over long-wl laser PMD, duplexity
    // unknown, dual rate
    EthernetMedia_ethernet_1000base_dr_lx EthernetMedia = "ethernet-1000base-dr-lx"

    // fiber over 4 lane optics (25km reach) (lite)
    EthernetMedia_ethernet_100gbase_er4l EthernetMedia = "ethernet-100gbase-er4l"

    // fiber over 4 lane optics (short reach)
    EthernetMedia_ethernet_100gbase_sr4 EthernetMedia = "ethernet-100gbase-sr4"

    // Bi-directional fiber over 2 lane optics (short
    // reach)
    EthernetMedia_ethernet_40gbase_sr_bd EthernetMedia = "ethernet-40gbase-sr-bd"

    // Twinaxial copper cabling
    EthernetMedia_ethernet_25gbase_cr EthernetMedia = "ethernet-25gbase-cr"

    // Twinaxial copper cabling (no RS-FEC)
    EthernetMedia_ethernet_25gbase_cr_s EthernetMedia = "ethernet-25gbase-cr-s"

    // One lane electrical backplane
    EthernetMedia_ethernet_25gbase_kr EthernetMedia = "ethernet-25gbase-kr"

    // One lane electrical backplane (no RS-FEC)
    EthernetMedia_ethernet_25gbase_kr_s EthernetMedia = "ethernet-25gbase-kr-s"

    // One lane fiber
    EthernetMedia_ethernet_25gbase_r EthernetMedia = "ethernet-25gbase-r"

    // Multimode fiber
    EthernetMedia_ethernet_25gbase_sr EthernetMedia = "ethernet-25gbase-sr"

    // DWDM optics
    EthernetMedia_ethernet_25gbase_dwdm EthernetMedia = "ethernet-25gbase-dwdm"

    // DWDM with tunable optics
    EthernetMedia_ethernet_25gbase_dwdm_tunable EthernetMedia = "ethernet-25gbase-dwdm-tunable"

    // CWDM optics
    EthernetMedia_ethernet_25gbase_cwdm EthernetMedia = "ethernet-25gbase-cwdm"

    // CWDM with tunable optics
    EthernetMedia_ethernet_25gbase_cwdm_tunable EthernetMedia = "ethernet-25gbase-cwdm-tunable"

    // 4 parallel single mode fibers
    EthernetMedia_ethernet_100gbase_psm4 EthernetMedia = "ethernet-100gbase-psm4"

    // Fiber over 10 lane optics (extended reach)
    EthernetMedia_ethernet_100gbase_er10 EthernetMedia = "ethernet-100gbase-er10"

    // Fiber over 10 lane optics (extended reach)
    // (lite)
    EthernetMedia_ethernet_100gbase_er10l EthernetMedia = "ethernet-100gbase-er10l"

    // Active copper cable
    EthernetMedia_ethernet_100gbase_acc EthernetMedia = "ethernet-100gbase-acc"

    // Active optical cable
    EthernetMedia_ethernet_100gbase_aoc EthernetMedia = "ethernet-100gbase-aoc"

    // 4 lane CWDM cable
    EthernetMedia_ethernet_100gbase_cwdm4 EthernetMedia = "ethernet-100gbase-cwdm4"

    // 4 parallel single mode fibers
    EthernetMedia_ethernet_40gbase_psm4 EthernetMedia = "ethernet-40gbase-psm4"

    // 4 lane copper (very short reach)
    EthernetMedia_ethernet_100gbase_cr4 EthernetMedia = "ethernet-100gbase-cr4"

    // Active loopback
    EthernetMedia_ethernet_100gbase_act_loop EthernetMedia = "ethernet-100gbase-act-loop"

    // Passive loopback
    EthernetMedia_ethernet_100gbase_pas_loop EthernetMedia = "ethernet-100gbase-pas-loop"

    // 2 lane copper (very short reach)
    EthernetMedia_ethernet_50gbase_cr2 EthernetMedia = "ethernet-50gbase-cr2"

    // fiber over 2 lane optics (short reach)
    EthernetMedia_ethernet_50gbase_sr2 EthernetMedia = "ethernet-50gbase-sr2"

    // 2 parallel single mode fibers
    EthernetMedia_ethernet_50gbase_psm2 EthernetMedia = "ethernet-50gbase-psm2"

    // 4 lanes Passive Copper
    EthernetMedia_ethernet_200gbase_cr4 EthernetMedia = "ethernet-200gbase-cr4"

    // Duplex SMF LC Connector 2km
    EthernetMedia_ethernet_400gbase_fr4 EthernetMedia = "ethernet-400gbase-fr4"

    // SMF with MPO-12 connector
    EthernetMedia_ethernet_400gbase_dr4 EthernetMedia = "ethernet-400gbase-dr4"

    // 4 lanes Passive Copper
    EthernetMedia_ethernet_400gbase_cr4 EthernetMedia = "ethernet-400gbase-cr4"

    // Passive Twinax cable assembly 1m
    EthernetMedia_ethernet_10gbase_cu1m EthernetMedia = "ethernet-10gbase-cu1m"

    // Passive Twinax cable assembly 1.5m
    EthernetMedia_ethernet_10gbase_cu1_5m EthernetMedia = "ethernet-10gbase-cu1-5m"

    // Passive Twinax cable assembly 2m
    EthernetMedia_ethernet_10gbase_cu2m EthernetMedia = "ethernet-10gbase-cu2m"

    // Passive Twinax cable assembly 2.5m
    EthernetMedia_ethernet_10gbase_cu2_5m EthernetMedia = "ethernet-10gbase-cu2-5m"

    // Passive Twinax cable assembly 3m
    EthernetMedia_ethernet_10gbase_cu3m EthernetMedia = "ethernet-10gbase-cu3m"

    // Passive Twinax cable assembly 5m
    EthernetMedia_ethernet_10gbase_cu5m EthernetMedia = "ethernet-10gbase-cu5m"

    // Active Twinax cable assembly 7m
    EthernetMedia_ethernet_10gbase_acu7m EthernetMedia = "ethernet-10gbase-acu7m"

    // Active Twinax cable assembly 10m
    EthernetMedia_ethernet_10gbase_acu10m EthernetMedia = "ethernet-10gbase-acu10m"

    // Active optical cable 1m
    EthernetMedia_ethernet_10gbase_aoc1m EthernetMedia = "ethernet-10gbase-aoc1m"

    // Active optical cable 2m
    EthernetMedia_ethernet_10gbase_aoc2m EthernetMedia = "ethernet-10gbase-aoc2m"

    // Active optical cable 3m
    EthernetMedia_ethernet_10gbase_aoc3m EthernetMedia = "ethernet-10gbase-aoc3m"

    // Active optical cable 5m
    EthernetMedia_ethernet_10gbase_aoc5m EthernetMedia = "ethernet-10gbase-aoc5m"

    // Active optical cable 7m
    EthernetMedia_ethernet_10gbase_aoc7m EthernetMedia = "ethernet-10gbase-aoc7m"

    // Active optical cable 10m
    EthernetMedia_ethernet_10gbase_aoc10m EthernetMedia = "ethernet-10gbase-aoc10m"

    // Active optical cable
    EthernetMedia_ethernet_40gbase_aoc EthernetMedia = "ethernet-40gbase-aoc"

    // fiber over 4 lane optics (long reach)
    EthernetMedia_ethernet_4x10g_base_lr EthernetMedia = "ethernet-4x10g-base-lr"

    // Active Twinax cable assembly 1m
    EthernetMedia_ethernet_40gbase_acu1m EthernetMedia = "ethernet-40gbase-acu1m"

    // Active Twinax cable assembly 3m
    EthernetMedia_ethernet_40gbase_acu3m EthernetMedia = "ethernet-40gbase-acu3m"

    // Active Twinax cable assembly 5m
    EthernetMedia_ethernet_40gbase_acu5m EthernetMedia = "ethernet-40gbase-acu5m"

    // Active Twinax cable assembly 7m
    EthernetMedia_ethernet_40gbase_acu7m EthernetMedia = "ethernet-40gbase-acu7m"

    // Active Twinax cable assembly 10m
    EthernetMedia_ethernet_40gbase_acu10m EthernetMedia = "ethernet-40gbase-acu10m"

    // Twinaxial copper cabling (1m)
    EthernetMedia_ethernet_25gbase_cu1m EthernetMedia = "ethernet-25gbase-cu1m"

    // Twinaxial copper cabling (2m)
    EthernetMedia_ethernet_25gbase_cu2m EthernetMedia = "ethernet-25gbase-cu2m"

    // Twinaxial copper cabling (3m)
    EthernetMedia_ethernet_25gbase_cu3m EthernetMedia = "ethernet-25gbase-cu3m"

    // Twinaxial copper cabling (5m)
    EthernetMedia_ethernet_25gbase_cu5m EthernetMedia = "ethernet-25gbase-cu5m"

    // 4 lane CWDM Lite cable
    EthernetMedia_ethernet_100gbase_sm_sr EthernetMedia = "ethernet-100gbase-sm-sr"

    // ethernet base max
    EthernetMedia_ethernet_base_max EthernetMedia = "ethernet-base-max"
)

// EtherAinsStatus represents Ether ains status
type EtherAinsStatus string

const (
    // AINS Soak timer not running
    EtherAinsStatus_ains_soak_status_none EtherAinsStatus = "ains-soak-status-none"

    // AINS Soak timer pending
    EtherAinsStatus_ains_soak_status_pending EtherAinsStatus = "ains-soak-status-pending"

    // AINS Soak timer running
    EtherAinsStatus_ains_soak_status_running EtherAinsStatus = "ains-soak-status-running"
)

// EthernetLoopback represents Loopback type
type EthernetLoopback string

const (
    // Disabled
    EthernetLoopback_no_loopback EthernetLoopback = "no-loopback"

    // Loopback in the framer
    EthernetLoopback_internal EthernetLoopback = "internal"

    // Loops peer's packets back to them
    EthernetLoopback_line EthernetLoopback = "line"

    // tx externally connected to rx
    EthernetLoopback_external EthernetLoopback = "external"
)

// EthernetInterface
// Ethernet operational data
type EthernetInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet controller statistics table.
    Statistics EthernetInterface_Statistics

    // Ethernet controller info table.
    Interfaces EthernetInterface_Interfaces

    // Ethernet controller BERT table.
    Berts EthernetInterface_Berts
}

func (ethernetInterface *EthernetInterface) GetEntityData() *types.CommonEntityData {
    ethernetInterface.EntityData.YFilter = ethernetInterface.YFilter
    ethernetInterface.EntityData.YangName = "ethernet-interface"
    ethernetInterface.EntityData.BundleName = "cisco_ios_xr"
    ethernetInterface.EntityData.ParentYangName = "Cisco-IOS-XR-drivers-media-eth-oper"
    ethernetInterface.EntityData.SegmentPath = "Cisco-IOS-XR-drivers-media-eth-oper:ethernet-interface"
    ethernetInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetInterface.EntityData.Children = types.NewOrderedMap()
    ethernetInterface.EntityData.Children.Append("statistics", types.YChild{"Statistics", &ethernetInterface.Statistics})
    ethernetInterface.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ethernetInterface.Interfaces})
    ethernetInterface.EntityData.Children.Append("berts", types.YChild{"Berts", &ethernetInterface.Berts})
    ethernetInterface.EntityData.Leafs = types.NewOrderedMap()

    ethernetInterface.EntityData.YListKeys = []string {}

    return &(ethernetInterface.EntityData)
}

// EthernetInterface_Statistics
// Ethernet controller statistics table
type EthernetInterface_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet statistics information. The type is slice of
    // EthernetInterface_Statistics_Statistic.
    Statistic []*EthernetInterface_Statistics_Statistic
}

func (statistics *EthernetInterface_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "ethernet-interface"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("statistic", types.YChild{"Statistic", nil})
    for i := range statistics.Statistic {
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.Statistic[i]), types.YChild{"Statistic", statistics.Statistic[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// EthernetInterface_Statistics_Statistic
// Ethernet statistics information
type EthernetInterface_Statistics_Statistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Total octets of all frames. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotalBytes interface{}

    // Total octets of all good frames. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedGoodBytes interface{}

    // All frames, good or bad. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotalFrames interface{}

    // All 802.1Q frames. The type is interface{} with range:
    // 0..18446744073709551615.
    Received8021qFrames interface{}

    // All pause frames. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPauseFrames interface{}

    // Unsupported MAC Control frames. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedUnknownOpcodes interface{}

    // All 64 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotal64OctetFrames interface{}

    // All 65-127 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotalOctetFramesFrom65To127 interface{}

    // All 128-255 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotalOctetFramesFrom128To255 interface{}

    // All 256-511 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotalOctetFramesFrom256To511 interface{}

    // All 512-1023 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotalOctetFramesFrom512To1023 interface{}

    // All 1024-1518 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotalOctetFramesFrom1024To1518 interface{}

    // All > 1518 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedTotalOctetFramesFrom1519ToMax interface{}

    // Received Good Frames. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedGoodFrames interface{}

    // Received unicast Frames. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedUnicastFrames interface{}

    // Received multicast Frames. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedMulticastFrames interface{}

    // Received broadcast Frames. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedBroadcastFrames interface{}

    // Drops due to buffer overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfBufferOverrunPacketsDropped interface{}

    // Drops due to packet abort. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfAbortedPacketsDropped interface{}

    // Drops due to invalid VLAN id. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberofInvalidVlanIdPacketsDropped interface{}

    // Drops due to the destination MAC not matching. The type is interface{} with
    // range: 0..18446744073709551615.
    InvalidDestMacDropPackets interface{}

    // Drops due to the encapsulation or ether type not matching. The type is
    // interface{} with range: 0..18446744073709551615.
    InvalidEncapDropPackets interface{}

    // Any other drops not counted. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfMiscellaneousPacketsDropped interface{}

    // Good frames > MRU, dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedGiantPacketsGreaterthanMru interface{}

    // Good frames < 64 Octet, dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedEtherStatsUndersizePkts interface{}

    // Bad Frames > MRU, dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedJabbersPacketsGreaterthanMru interface{}

    // Bad Frames < 64 Octet, dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedEtherStatsFragments interface{}

    // Frames 64 - MRU with CRC error. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPacketsWithCrcAlignErrors interface{}

    // All collision events. The type is interface{} with range:
    // 0..18446744073709551615.
    EtherStatsCollisions interface{}

    // Symbol errors. The type is interface{} with range: 0..18446744073709551615.
    SymbolErrors interface{}

    // Any other errors not counted. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedMiscellaneousErrorPackets interface{}

    // RFC2819 etherStatsOversizedPkts. The type is interface{} with range:
    // 0..18446744073709551615.
    Rfc2819EtherStatsOversizedPkts interface{}

    // RFC2819 etherStatsJabbers. The type is interface{} with range:
    // 0..18446744073709551615.
    Rfc2819EtherStatsJabbers interface{}

    // RFC2819 etherStatsCRCAlignErrors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rfc2819EtherStatsCrcAlignErrors interface{}

    // RFC3635 dot3StatsAlignmentErrors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rfc3635dot3StatsAlignmentErrors interface{}

    // Total octets of all frames. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalBytesTransmitted interface{}

    // Total octets of all good frames. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalGoodBytesTransmitted interface{}

    // All frames, good or bad. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalFramesTransmitted interface{}

    // All 802.1Q frames. The type is interface{} with range:
    // 0..18446744073709551615.
    Transmitted8021qFrames interface{}

    // All pause frames. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedTotalPauseFrames interface{}

    // All 64 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedTotal64OctetFrames interface{}

    // All 65-127 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedTotalOctetFramesFrom65To127 interface{}

    // All 128-255 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedTotalOctetFramesFrom128To255 interface{}

    // All 256-511 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedTotalOctetFramesFrom256To511 interface{}

    // All 512-1023 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedTotalOctetFramesFrom512To1023 interface{}

    // All 1024-1518 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedTotalOctetFramesFrom1024To1518 interface{}

    // All > 1518 Octet Frame Count. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedTotalOctetFramesFrom1518ToMax interface{}

    // Transmitted Good Frames. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedGoodFrames interface{}

    // Transmitted unicast Frames. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedUnicastFrames interface{}

    // Transmitted multicast Frames. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedMulticastFrames interface{}

    // Transmitted broadcast Frames. The type is interface{} with range:
    // 0..18446744073709551615.
    TransmittedBroadcastFrames interface{}

    // Drops due to buffer underrun. The type is interface{} with range:
    // 0..18446744073709551615.
    BufferUnderrunPacketDrops interface{}

    // Drops due to packet abort. The type is interface{} with range:
    // 0..18446744073709551615.
    AbortedPacketDrops interface{}

    // Any other drops not counted. The type is interface{} with range:
    // 0..18446744073709551615.
    UncountedDroppedFrames interface{}

    // Any other errors not counted. The type is interface{} with range:
    // 0..18446744073709551615.
    MiscellaneousOutputErrors interface{}
}

func (statistic *EthernetInterface_Statistics_Statistic) GetEntityData() *types.CommonEntityData {
    statistic.EntityData.YFilter = statistic.YFilter
    statistic.EntityData.YangName = "statistic"
    statistic.EntityData.BundleName = "cisco_ios_xr"
    statistic.EntityData.ParentYangName = "statistics"
    statistic.EntityData.SegmentPath = "statistic" + types.AddKeyToken(statistic.InterfaceName, "interface-name")
    statistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic.EntityData.Children = types.NewOrderedMap()
    statistic.EntityData.Leafs = types.NewOrderedMap()
    statistic.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", statistic.InterfaceName})
    statistic.EntityData.Leafs.Append("received-total-bytes", types.YLeaf{"ReceivedTotalBytes", statistic.ReceivedTotalBytes})
    statistic.EntityData.Leafs.Append("received-good-bytes", types.YLeaf{"ReceivedGoodBytes", statistic.ReceivedGoodBytes})
    statistic.EntityData.Leafs.Append("received-total-frames", types.YLeaf{"ReceivedTotalFrames", statistic.ReceivedTotalFrames})
    statistic.EntityData.Leafs.Append("received8021q-frames", types.YLeaf{"Received8021qFrames", statistic.Received8021qFrames})
    statistic.EntityData.Leafs.Append("received-pause-frames", types.YLeaf{"ReceivedPauseFrames", statistic.ReceivedPauseFrames})
    statistic.EntityData.Leafs.Append("received-unknown-opcodes", types.YLeaf{"ReceivedUnknownOpcodes", statistic.ReceivedUnknownOpcodes})
    statistic.EntityData.Leafs.Append("received-total64-octet-frames", types.YLeaf{"ReceivedTotal64OctetFrames", statistic.ReceivedTotal64OctetFrames})
    statistic.EntityData.Leafs.Append("received-total-octet-frames-from65-to127", types.YLeaf{"ReceivedTotalOctetFramesFrom65To127", statistic.ReceivedTotalOctetFramesFrom65To127})
    statistic.EntityData.Leafs.Append("received-total-octet-frames-from128-to255", types.YLeaf{"ReceivedTotalOctetFramesFrom128To255", statistic.ReceivedTotalOctetFramesFrom128To255})
    statistic.EntityData.Leafs.Append("received-total-octet-frames-from256-to511", types.YLeaf{"ReceivedTotalOctetFramesFrom256To511", statistic.ReceivedTotalOctetFramesFrom256To511})
    statistic.EntityData.Leafs.Append("received-total-octet-frames-from512-to1023", types.YLeaf{"ReceivedTotalOctetFramesFrom512To1023", statistic.ReceivedTotalOctetFramesFrom512To1023})
    statistic.EntityData.Leafs.Append("received-total-octet-frames-from1024-to1518", types.YLeaf{"ReceivedTotalOctetFramesFrom1024To1518", statistic.ReceivedTotalOctetFramesFrom1024To1518})
    statistic.EntityData.Leafs.Append("received-total-octet-frames-from1519-to-max", types.YLeaf{"ReceivedTotalOctetFramesFrom1519ToMax", statistic.ReceivedTotalOctetFramesFrom1519ToMax})
    statistic.EntityData.Leafs.Append("received-good-frames", types.YLeaf{"ReceivedGoodFrames", statistic.ReceivedGoodFrames})
    statistic.EntityData.Leafs.Append("received-unicast-frames", types.YLeaf{"ReceivedUnicastFrames", statistic.ReceivedUnicastFrames})
    statistic.EntityData.Leafs.Append("received-multicast-frames", types.YLeaf{"ReceivedMulticastFrames", statistic.ReceivedMulticastFrames})
    statistic.EntityData.Leafs.Append("received-broadcast-frames", types.YLeaf{"ReceivedBroadcastFrames", statistic.ReceivedBroadcastFrames})
    statistic.EntityData.Leafs.Append("number-of-buffer-overrun-packets-dropped", types.YLeaf{"NumberOfBufferOverrunPacketsDropped", statistic.NumberOfBufferOverrunPacketsDropped})
    statistic.EntityData.Leafs.Append("number-of-aborted-packets-dropped", types.YLeaf{"NumberOfAbortedPacketsDropped", statistic.NumberOfAbortedPacketsDropped})
    statistic.EntityData.Leafs.Append("numberof-invalid-vlan-id-packets-dropped", types.YLeaf{"NumberofInvalidVlanIdPacketsDropped", statistic.NumberofInvalidVlanIdPacketsDropped})
    statistic.EntityData.Leafs.Append("invalid-dest-mac-drop-packets", types.YLeaf{"InvalidDestMacDropPackets", statistic.InvalidDestMacDropPackets})
    statistic.EntityData.Leafs.Append("invalid-encap-drop-packets", types.YLeaf{"InvalidEncapDropPackets", statistic.InvalidEncapDropPackets})
    statistic.EntityData.Leafs.Append("number-of-miscellaneous-packets-dropped", types.YLeaf{"NumberOfMiscellaneousPacketsDropped", statistic.NumberOfMiscellaneousPacketsDropped})
    statistic.EntityData.Leafs.Append("dropped-giant-packets-greaterthan-mru", types.YLeaf{"DroppedGiantPacketsGreaterthanMru", statistic.DroppedGiantPacketsGreaterthanMru})
    statistic.EntityData.Leafs.Append("dropped-ether-stats-undersize-pkts", types.YLeaf{"DroppedEtherStatsUndersizePkts", statistic.DroppedEtherStatsUndersizePkts})
    statistic.EntityData.Leafs.Append("dropped-jabbers-packets-greaterthan-mru", types.YLeaf{"DroppedJabbersPacketsGreaterthanMru", statistic.DroppedJabbersPacketsGreaterthanMru})
    statistic.EntityData.Leafs.Append("dropped-ether-stats-fragments", types.YLeaf{"DroppedEtherStatsFragments", statistic.DroppedEtherStatsFragments})
    statistic.EntityData.Leafs.Append("dropped-packets-with-crc-align-errors", types.YLeaf{"DroppedPacketsWithCrcAlignErrors", statistic.DroppedPacketsWithCrcAlignErrors})
    statistic.EntityData.Leafs.Append("ether-stats-collisions", types.YLeaf{"EtherStatsCollisions", statistic.EtherStatsCollisions})
    statistic.EntityData.Leafs.Append("symbol-errors", types.YLeaf{"SymbolErrors", statistic.SymbolErrors})
    statistic.EntityData.Leafs.Append("dropped-miscellaneous-error-packets", types.YLeaf{"DroppedMiscellaneousErrorPackets", statistic.DroppedMiscellaneousErrorPackets})
    statistic.EntityData.Leafs.Append("rfc2819-ether-stats-oversized-pkts", types.YLeaf{"Rfc2819EtherStatsOversizedPkts", statistic.Rfc2819EtherStatsOversizedPkts})
    statistic.EntityData.Leafs.Append("rfc2819-ether-stats-jabbers", types.YLeaf{"Rfc2819EtherStatsJabbers", statistic.Rfc2819EtherStatsJabbers})
    statistic.EntityData.Leafs.Append("rfc2819-ether-stats-crc-align-errors", types.YLeaf{"Rfc2819EtherStatsCrcAlignErrors", statistic.Rfc2819EtherStatsCrcAlignErrors})
    statistic.EntityData.Leafs.Append("rfc3635dot3-stats-alignment-errors", types.YLeaf{"Rfc3635dot3StatsAlignmentErrors", statistic.Rfc3635dot3StatsAlignmentErrors})
    statistic.EntityData.Leafs.Append("total-bytes-transmitted", types.YLeaf{"TotalBytesTransmitted", statistic.TotalBytesTransmitted})
    statistic.EntityData.Leafs.Append("total-good-bytes-transmitted", types.YLeaf{"TotalGoodBytesTransmitted", statistic.TotalGoodBytesTransmitted})
    statistic.EntityData.Leafs.Append("total-frames-transmitted", types.YLeaf{"TotalFramesTransmitted", statistic.TotalFramesTransmitted})
    statistic.EntityData.Leafs.Append("transmitted8021q-frames", types.YLeaf{"Transmitted8021qFrames", statistic.Transmitted8021qFrames})
    statistic.EntityData.Leafs.Append("transmitted-total-pause-frames", types.YLeaf{"TransmittedTotalPauseFrames", statistic.TransmittedTotalPauseFrames})
    statistic.EntityData.Leafs.Append("transmitted-total64-octet-frames", types.YLeaf{"TransmittedTotal64OctetFrames", statistic.TransmittedTotal64OctetFrames})
    statistic.EntityData.Leafs.Append("transmitted-total-octet-frames-from65-to127", types.YLeaf{"TransmittedTotalOctetFramesFrom65To127", statistic.TransmittedTotalOctetFramesFrom65To127})
    statistic.EntityData.Leafs.Append("transmitted-total-octet-frames-from128-to255", types.YLeaf{"TransmittedTotalOctetFramesFrom128To255", statistic.TransmittedTotalOctetFramesFrom128To255})
    statistic.EntityData.Leafs.Append("transmitted-total-octet-frames-from256-to511", types.YLeaf{"TransmittedTotalOctetFramesFrom256To511", statistic.TransmittedTotalOctetFramesFrom256To511})
    statistic.EntityData.Leafs.Append("transmitted-total-octet-frames-from512-to1023", types.YLeaf{"TransmittedTotalOctetFramesFrom512To1023", statistic.TransmittedTotalOctetFramesFrom512To1023})
    statistic.EntityData.Leafs.Append("transmitted-total-octet-frames-from1024-to1518", types.YLeaf{"TransmittedTotalOctetFramesFrom1024To1518", statistic.TransmittedTotalOctetFramesFrom1024To1518})
    statistic.EntityData.Leafs.Append("transmitted-total-octet-frames-from1518-to-max", types.YLeaf{"TransmittedTotalOctetFramesFrom1518ToMax", statistic.TransmittedTotalOctetFramesFrom1518ToMax})
    statistic.EntityData.Leafs.Append("transmitted-good-frames", types.YLeaf{"TransmittedGoodFrames", statistic.TransmittedGoodFrames})
    statistic.EntityData.Leafs.Append("transmitted-unicast-frames", types.YLeaf{"TransmittedUnicastFrames", statistic.TransmittedUnicastFrames})
    statistic.EntityData.Leafs.Append("transmitted-multicast-frames", types.YLeaf{"TransmittedMulticastFrames", statistic.TransmittedMulticastFrames})
    statistic.EntityData.Leafs.Append("transmitted-broadcast-frames", types.YLeaf{"TransmittedBroadcastFrames", statistic.TransmittedBroadcastFrames})
    statistic.EntityData.Leafs.Append("buffer-underrun-packet-drops", types.YLeaf{"BufferUnderrunPacketDrops", statistic.BufferUnderrunPacketDrops})
    statistic.EntityData.Leafs.Append("aborted-packet-drops", types.YLeaf{"AbortedPacketDrops", statistic.AbortedPacketDrops})
    statistic.EntityData.Leafs.Append("uncounted-dropped-frames", types.YLeaf{"UncountedDroppedFrames", statistic.UncountedDroppedFrames})
    statistic.EntityData.Leafs.Append("miscellaneous-output-errors", types.YLeaf{"MiscellaneousOutputErrors", statistic.MiscellaneousOutputErrors})

    statistic.EntityData.YListKeys = []string {"InterfaceName"}

    return &(statistic.EntityData)
}

// EthernetInterface_Interfaces
// Ethernet controller info table
type EthernetInterface_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet controller information. The type is slice of
    // EthernetInterface_Interfaces_Interface.
    Interface []*EthernetInterface_Interfaces_Interface
}

func (interfaces *EthernetInterface_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ethernet-interface"
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

// EthernetInterface_Interfaces_Interface
// Ethernet controller information
type EthernetInterface_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Port Administrative State. The type is EthernetPortEnable.
    AdminState interface{}

    // Port Operational state - TRUE if up. The type is bool.
    OperStateUp interface{}

    // PHY information.
    PhyInfo EthernetInterface_Interfaces_Interface_PhyInfo

    // Layer 1 information.
    Layer1Info EthernetInterface_Interfaces_Interface_Layer1Info

    // MAC Layer information.
    MacInfo EthernetInterface_Interfaces_Interface_MacInfo

    // Transport state information.
    TransportInfo EthernetInterface_Interfaces_Interface_TransportInfo
}

func (self *EthernetInterface_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("phy-info", types.YChild{"PhyInfo", &self.PhyInfo})
    self.EntityData.Children.Append("layer1-info", types.YChild{"Layer1Info", &self.Layer1Info})
    self.EntityData.Children.Append("mac-info", types.YChild{"MacInfo", &self.MacInfo})
    self.EntityData.Children.Append("transport-info", types.YChild{"TransportInfo", &self.TransportInfo})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", self.AdminState})
    self.EntityData.Leafs.Append("oper-state-up", types.YLeaf{"OperStateUp", self.OperStateUp})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo
// PHY information
type EthernetInterface_Interfaces_Interface_PhyInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port media type. The type is EthernetMedia.
    MediaType interface{}

    // Presence of PHY. The type is EtherPhyPresent.
    PhyPresent interface{}

    // Port operational loopback. The type is EthernetLoopback.
    Loopback interface{}

    // Holdoff Time. The type is interface{} with range: 0..4294967295.
    HoldoffTime interface{}

    // Details about the PHY.
    PhyDetails EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails

    // Forward Error Correction information.
    FecDetails EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails

    // Port operational extended loopback. The type is slice of
    // EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback.
    ExtendedLoopback []*EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback
}

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetEntityData() *types.CommonEntityData {
    phyInfo.EntityData.YFilter = phyInfo.YFilter
    phyInfo.EntityData.YangName = "phy-info"
    phyInfo.EntityData.BundleName = "cisco_ios_xr"
    phyInfo.EntityData.ParentYangName = "interface"
    phyInfo.EntityData.SegmentPath = "phy-info"
    phyInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    phyInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    phyInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    phyInfo.EntityData.Children = types.NewOrderedMap()
    phyInfo.EntityData.Children.Append("phy-details", types.YChild{"PhyDetails", &phyInfo.PhyDetails})
    phyInfo.EntityData.Children.Append("fec-details", types.YChild{"FecDetails", &phyInfo.FecDetails})
    phyInfo.EntityData.Children.Append("extended-loopback", types.YChild{"ExtendedLoopback", nil})
    for i := range phyInfo.ExtendedLoopback {
        phyInfo.EntityData.Children.Append(types.GetSegmentPath(phyInfo.ExtendedLoopback[i]), types.YChild{"ExtendedLoopback", phyInfo.ExtendedLoopback[i]})
    }
    phyInfo.EntityData.Leafs = types.NewOrderedMap()
    phyInfo.EntityData.Leafs.Append("media-type", types.YLeaf{"MediaType", phyInfo.MediaType})
    phyInfo.EntityData.Leafs.Append("phy-present", types.YLeaf{"PhyPresent", phyInfo.PhyPresent})
    phyInfo.EntityData.Leafs.Append("loopback", types.YLeaf{"Loopback", phyInfo.Loopback})
    phyInfo.EntityData.Leafs.Append("holdoff-time", types.YLeaf{"HoldoffTime", phyInfo.HoldoffTime})

    phyInfo.EntityData.YListKeys = []string {}

    return &(phyInfo.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails
// Details about the PHY
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the port optics manufacturer. The type is string.
    Vendor interface{}

    // Part number for the port optics. The type is string.
    VendorPartNumber interface{}

    // Serial number for the port optics. The type is string.
    VendorSerialNumber interface{}

    // The temperature of the transceiver (mDegrees C). The type is interface{}
    // with range: -2147483648..2147483647.
    TransceiverTemperature interface{}

    // The input voltage to the transceiver (mV). The type is interface{} with
    // range: -2147483648..2147483647.
    TransceiverVoltage interface{}

    // The transceiver transmit laser power (uW). The type is interface{} with
    // range: -2147483648..2147483647.
    TransceiverTxPower interface{}

    // The transceiver receive optical power (uW). The type is interface{} with
    // range: -2147483648..2147483647.
    TransceiverRxPower interface{}

    // The laser bias of the transceiver (uA). The type is interface{} with range:
    // -2147483648..2147483647.
    TransceiverTxBias interface{}

    // Wavelength of the optics being used in nm * 1000. The type is interface{}
    // with range: 0..4294967295.
    OpticsWavelength interface{}

    // Optics module type. The type is string.
    OpticsType interface{}

    // Module revision number. The type is string.
    RevisionNumber interface{}

    // Digital Optical Monitoring (per lane information) validity.
    LaneFieldValidity EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity

    // Digital Optical Monitoring alarm thresholds.
    DigOptMonAlarmThresholds EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds

    // Digital Optical Monitoring alarms.
    DigOptMonAlarms EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms

    // Digital Optical Monitoring (per lane information). The type is slice of
    // EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane.
    Lane []*EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane
}

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetEntityData() *types.CommonEntityData {
    phyDetails.EntityData.YFilter = phyDetails.YFilter
    phyDetails.EntityData.YangName = "phy-details"
    phyDetails.EntityData.BundleName = "cisco_ios_xr"
    phyDetails.EntityData.ParentYangName = "phy-info"
    phyDetails.EntityData.SegmentPath = "phy-details"
    phyDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    phyDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    phyDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    phyDetails.EntityData.Children = types.NewOrderedMap()
    phyDetails.EntityData.Children.Append("lane-field-validity", types.YChild{"LaneFieldValidity", &phyDetails.LaneFieldValidity})
    phyDetails.EntityData.Children.Append("dig-opt-mon-alarm-thresholds", types.YChild{"DigOptMonAlarmThresholds", &phyDetails.DigOptMonAlarmThresholds})
    phyDetails.EntityData.Children.Append("dig-opt-mon-alarms", types.YChild{"DigOptMonAlarms", &phyDetails.DigOptMonAlarms})
    phyDetails.EntityData.Children.Append("lane", types.YChild{"Lane", nil})
    for i := range phyDetails.Lane {
        phyDetails.EntityData.Children.Append(types.GetSegmentPath(phyDetails.Lane[i]), types.YChild{"Lane", phyDetails.Lane[i]})
    }
    phyDetails.EntityData.Leafs = types.NewOrderedMap()
    phyDetails.EntityData.Leafs.Append("vendor", types.YLeaf{"Vendor", phyDetails.Vendor})
    phyDetails.EntityData.Leafs.Append("vendor-part-number", types.YLeaf{"VendorPartNumber", phyDetails.VendorPartNumber})
    phyDetails.EntityData.Leafs.Append("vendor-serial-number", types.YLeaf{"VendorSerialNumber", phyDetails.VendorSerialNumber})
    phyDetails.EntityData.Leafs.Append("transceiver-temperature", types.YLeaf{"TransceiverTemperature", phyDetails.TransceiverTemperature})
    phyDetails.EntityData.Leafs.Append("transceiver-voltage", types.YLeaf{"TransceiverVoltage", phyDetails.TransceiverVoltage})
    phyDetails.EntityData.Leafs.Append("transceiver-tx-power", types.YLeaf{"TransceiverTxPower", phyDetails.TransceiverTxPower})
    phyDetails.EntityData.Leafs.Append("transceiver-rx-power", types.YLeaf{"TransceiverRxPower", phyDetails.TransceiverRxPower})
    phyDetails.EntityData.Leafs.Append("transceiver-tx-bias", types.YLeaf{"TransceiverTxBias", phyDetails.TransceiverTxBias})
    phyDetails.EntityData.Leafs.Append("optics-wavelength", types.YLeaf{"OpticsWavelength", phyDetails.OpticsWavelength})
    phyDetails.EntityData.Leafs.Append("optics-type", types.YLeaf{"OpticsType", phyDetails.OpticsType})
    phyDetails.EntityData.Leafs.Append("revision-number", types.YLeaf{"RevisionNumber", phyDetails.RevisionNumber})

    phyDetails.EntityData.YListKeys = []string {}

    return &(phyDetails.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity
// Digital Optical Monitoring (per lane
// information) validity
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The wavelength 'per lane' field is valid. The type is interface{} with
    // range: -2147483648..2147483647.
    WavelengthValid interface{}

    // The transmit power 'per lane' field is valid. The type is interface{} with
    // range: -2147483648..2147483647.
    TransmitPowerValid interface{}

    // The receive power 'per lane' field is valid. The type is interface{} with
    // range: -2147483648..2147483647.
    ReceivePowerValid interface{}

    // The laser bias 'per lane' field is valid. The type is interface{} with
    // range: -2147483648..2147483647.
    LaserBiasValid interface{}
}

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetEntityData() *types.CommonEntityData {
    laneFieldValidity.EntityData.YFilter = laneFieldValidity.YFilter
    laneFieldValidity.EntityData.YangName = "lane-field-validity"
    laneFieldValidity.EntityData.BundleName = "cisco_ios_xr"
    laneFieldValidity.EntityData.ParentYangName = "phy-details"
    laneFieldValidity.EntityData.SegmentPath = "lane-field-validity"
    laneFieldValidity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    laneFieldValidity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    laneFieldValidity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    laneFieldValidity.EntityData.Children = types.NewOrderedMap()
    laneFieldValidity.EntityData.Leafs = types.NewOrderedMap()
    laneFieldValidity.EntityData.Leafs.Append("wavelength-valid", types.YLeaf{"WavelengthValid", laneFieldValidity.WavelengthValid})
    laneFieldValidity.EntityData.Leafs.Append("transmit-power-valid", types.YLeaf{"TransmitPowerValid", laneFieldValidity.TransmitPowerValid})
    laneFieldValidity.EntityData.Leafs.Append("receive-power-valid", types.YLeaf{"ReceivePowerValid", laneFieldValidity.ReceivePowerValid})
    laneFieldValidity.EntityData.Leafs.Append("laser-bias-valid", types.YLeaf{"LaserBiasValid", laneFieldValidity.LaserBiasValid})

    laneFieldValidity.EntityData.YListKeys = []string {}

    return &(laneFieldValidity.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds
// Digital Optical Monitoring alarm thresholds
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transceiver high temperature alarm threshold (mDegrees C). The type is
    // interface{} with range: -2147483648..2147483647.
    TransceiverTemperatureAlarmHigh interface{}

    // Transceiver high temperature warning threshold (mDegrees C). The type is
    // interface{} with range: -2147483648..2147483647.
    TransceiverTemperatureWarningHigh interface{}

    // Transceiver low temperature warning threshold (mDegrees C). The type is
    // interface{} with range: -2147483648..2147483647.
    TransceiverTemperatureWarningLow interface{}

    // Transceiver low temperature alarm threshold (mDegrees C). The type is
    // interface{} with range: -2147483648..2147483647.
    TransceiverTemperatureAlarmLow interface{}

    // Transceiver high voltage alarm threshold (mV). The type is interface{} with
    // range: 0..4294967295.
    TransceiverVoltageAlarmHigh interface{}

    // Transceiver high voltage warning threshold (mV). The type is interface{}
    // with range: 0..4294967295.
    TransceiverVoltageWarningHigh interface{}

    // Transceiver low voltage warning threshold (mV). The type is interface{}
    // with range: 0..4294967295.
    TransceiverVoltageWarningLow interface{}

    // Transceiver low voltage alarm threshold (mV). The type is interface{} with
    // range: 0..4294967295.
    TransceiverVoltageAlarmLow interface{}

    // Laser bias high alarm threshold (uA). The type is interface{} with range:
    // 0..4294967295.
    LaserBiasAlarmHigh interface{}

    // Laser bias high warning threshold (uA). The type is interface{} with range:
    // 0..4294967295.
    LaserBiasWarningHigh interface{}

    // Laser bias low warning threshold (uA). The type is interface{} with range:
    // 0..4294967295.
    LaserBiasWarningLow interface{}

    // Laser bias low alarm threshold (uA). The type is interface{} with range:
    // 0..4294967295.
    LaserBiasAlarmLow interface{}

    // High optical transmit power alarm threshold (uW). The type is interface{}
    // with range: 0..4294967295.
    OpticalTransmitPowerAlarmHigh interface{}

    // High optical transmit power warning threshold (uW). The type is interface{}
    // with range: 0..4294967295.
    OpticalTransmitPowerWarningHigh interface{}

    // Low optical transmit power warning threshold (uW). The type is interface{}
    // with range: 0..4294967295.
    OpticalTransmitPowerWarningLow interface{}

    // Low optical transmit power alarm threshold (uW). The type is interface{}
    // with range: 0..4294967295.
    OpticalTransmitPowerAlarmLow interface{}

    // High optical receive power alarm threshold (uW). The type is interface{}
    // with range: 0..4294967295.
    OpticalReceivePowerAlarmHigh interface{}

    // High optical receive power warning threshold (uW). The type is interface{}
    // with range: 0..4294967295.
    OpticalReceivePowerWarningHigh interface{}

    // Low optical receive power warning threshold (uW). The type is interface{}
    // with range: 0..4294967295.
    OpticalReceivePowerWarningLow interface{}

    // Low optical receive power alarm threshold (uW). The type is interface{}
    // with range: 0..4294967295.
    OpticalReceivePowerAlarmLow interface{}

    // Field validity.
    FieldValidity EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity
}

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetEntityData() *types.CommonEntityData {
    digOptMonAlarmThresholds.EntityData.YFilter = digOptMonAlarmThresholds.YFilter
    digOptMonAlarmThresholds.EntityData.YangName = "dig-opt-mon-alarm-thresholds"
    digOptMonAlarmThresholds.EntityData.BundleName = "cisco_ios_xr"
    digOptMonAlarmThresholds.EntityData.ParentYangName = "phy-details"
    digOptMonAlarmThresholds.EntityData.SegmentPath = "dig-opt-mon-alarm-thresholds"
    digOptMonAlarmThresholds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    digOptMonAlarmThresholds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    digOptMonAlarmThresholds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    digOptMonAlarmThresholds.EntityData.Children = types.NewOrderedMap()
    digOptMonAlarmThresholds.EntityData.Children.Append("field-validity", types.YChild{"FieldValidity", &digOptMonAlarmThresholds.FieldValidity})
    digOptMonAlarmThresholds.EntityData.Leafs = types.NewOrderedMap()
    digOptMonAlarmThresholds.EntityData.Leafs.Append("transceiver-temperature-alarm-high", types.YLeaf{"TransceiverTemperatureAlarmHigh", digOptMonAlarmThresholds.TransceiverTemperatureAlarmHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("transceiver-temperature-warning-high", types.YLeaf{"TransceiverTemperatureWarningHigh", digOptMonAlarmThresholds.TransceiverTemperatureWarningHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("transceiver-temperature-warning-low", types.YLeaf{"TransceiverTemperatureWarningLow", digOptMonAlarmThresholds.TransceiverTemperatureWarningLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("transceiver-temperature-alarm-low", types.YLeaf{"TransceiverTemperatureAlarmLow", digOptMonAlarmThresholds.TransceiverTemperatureAlarmLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("transceiver-voltage-alarm-high", types.YLeaf{"TransceiverVoltageAlarmHigh", digOptMonAlarmThresholds.TransceiverVoltageAlarmHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("transceiver-voltage-warning-high", types.YLeaf{"TransceiverVoltageWarningHigh", digOptMonAlarmThresholds.TransceiverVoltageWarningHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("transceiver-voltage-warning-low", types.YLeaf{"TransceiverVoltageWarningLow", digOptMonAlarmThresholds.TransceiverVoltageWarningLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("transceiver-voltage-alarm-low", types.YLeaf{"TransceiverVoltageAlarmLow", digOptMonAlarmThresholds.TransceiverVoltageAlarmLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("laser-bias-alarm-high", types.YLeaf{"LaserBiasAlarmHigh", digOptMonAlarmThresholds.LaserBiasAlarmHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("laser-bias-warning-high", types.YLeaf{"LaserBiasWarningHigh", digOptMonAlarmThresholds.LaserBiasWarningHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("laser-bias-warning-low", types.YLeaf{"LaserBiasWarningLow", digOptMonAlarmThresholds.LaserBiasWarningLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("laser-bias-alarm-low", types.YLeaf{"LaserBiasAlarmLow", digOptMonAlarmThresholds.LaserBiasAlarmLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("optical-transmit-power-alarm-high", types.YLeaf{"OpticalTransmitPowerAlarmHigh", digOptMonAlarmThresholds.OpticalTransmitPowerAlarmHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("optical-transmit-power-warning-high", types.YLeaf{"OpticalTransmitPowerWarningHigh", digOptMonAlarmThresholds.OpticalTransmitPowerWarningHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("optical-transmit-power-warning-low", types.YLeaf{"OpticalTransmitPowerWarningLow", digOptMonAlarmThresholds.OpticalTransmitPowerWarningLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("optical-transmit-power-alarm-low", types.YLeaf{"OpticalTransmitPowerAlarmLow", digOptMonAlarmThresholds.OpticalTransmitPowerAlarmLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("optical-receive-power-alarm-high", types.YLeaf{"OpticalReceivePowerAlarmHigh", digOptMonAlarmThresholds.OpticalReceivePowerAlarmHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("optical-receive-power-warning-high", types.YLeaf{"OpticalReceivePowerWarningHigh", digOptMonAlarmThresholds.OpticalReceivePowerWarningHigh})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("optical-receive-power-warning-low", types.YLeaf{"OpticalReceivePowerWarningLow", digOptMonAlarmThresholds.OpticalReceivePowerWarningLow})
    digOptMonAlarmThresholds.EntityData.Leafs.Append("optical-receive-power-alarm-low", types.YLeaf{"OpticalReceivePowerAlarmLow", digOptMonAlarmThresholds.OpticalReceivePowerAlarmLow})

    digOptMonAlarmThresholds.EntityData.YListKeys = []string {}

    return &(digOptMonAlarmThresholds.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity
// Field validity
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The temperature fields are valid. The type is interface{} with range:
    // -2147483648..2147483647.
    TemperatureValid interface{}

    // The voltage fields are valid. The type is interface{} with range:
    // -2147483648..2147483647.
    VoltageValid interface{}

    // The laser bias fields are valid. The type is interface{} with range:
    // -2147483648..2147483647.
    LaserBiasValid interface{}

    // The transmit power fields are valid. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPowerValid interface{}

    // The receive power fields are valid. The type is interface{} with range:
    // -2147483648..2147483647.
    ReceivePowerValid interface{}
}

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetEntityData() *types.CommonEntityData {
    fieldValidity.EntityData.YFilter = fieldValidity.YFilter
    fieldValidity.EntityData.YangName = "field-validity"
    fieldValidity.EntityData.BundleName = "cisco_ios_xr"
    fieldValidity.EntityData.ParentYangName = "dig-opt-mon-alarm-thresholds"
    fieldValidity.EntityData.SegmentPath = "field-validity"
    fieldValidity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fieldValidity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fieldValidity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fieldValidity.EntityData.Children = types.NewOrderedMap()
    fieldValidity.EntityData.Leafs = types.NewOrderedMap()
    fieldValidity.EntityData.Leafs.Append("temperature-valid", types.YLeaf{"TemperatureValid", fieldValidity.TemperatureValid})
    fieldValidity.EntityData.Leafs.Append("voltage-valid", types.YLeaf{"VoltageValid", fieldValidity.VoltageValid})
    fieldValidity.EntityData.Leafs.Append("laser-bias-valid", types.YLeaf{"LaserBiasValid", fieldValidity.LaserBiasValid})
    fieldValidity.EntityData.Leafs.Append("transmit-power-valid", types.YLeaf{"TransmitPowerValid", fieldValidity.TransmitPowerValid})
    fieldValidity.EntityData.Leafs.Append("receive-power-valid", types.YLeaf{"ReceivePowerValid", fieldValidity.ReceivePowerValid})

    fieldValidity.EntityData.YListKeys = []string {}

    return &(fieldValidity.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms
// Digital Optical Monitoring alarms
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transceiver Temperature Alarm. The type is EtherDomAlarm.
    TransceiverTemperature interface{}

    // Transceiver Voltage Alarm. The type is EtherDomAlarm.
    TransceiverVoltage interface{}

    // Transmit Laser Power Alarm. The type is EtherDomAlarm.
    TransmitLaserPower interface{}

    // Received Optical Power Alarm. The type is EtherDomAlarm.
    ReceivedLaserPower interface{}

    // Laser Bias Current Alarm. The type is EtherDomAlarm.
    LaserBiasCurrent interface{}
}

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetEntityData() *types.CommonEntityData {
    digOptMonAlarms.EntityData.YFilter = digOptMonAlarms.YFilter
    digOptMonAlarms.EntityData.YangName = "dig-opt-mon-alarms"
    digOptMonAlarms.EntityData.BundleName = "cisco_ios_xr"
    digOptMonAlarms.EntityData.ParentYangName = "phy-details"
    digOptMonAlarms.EntityData.SegmentPath = "dig-opt-mon-alarms"
    digOptMonAlarms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    digOptMonAlarms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    digOptMonAlarms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    digOptMonAlarms.EntityData.Children = types.NewOrderedMap()
    digOptMonAlarms.EntityData.Leafs = types.NewOrderedMap()
    digOptMonAlarms.EntityData.Leafs.Append("transceiver-temperature", types.YLeaf{"TransceiverTemperature", digOptMonAlarms.TransceiverTemperature})
    digOptMonAlarms.EntityData.Leafs.Append("transceiver-voltage", types.YLeaf{"TransceiverVoltage", digOptMonAlarms.TransceiverVoltage})
    digOptMonAlarms.EntityData.Leafs.Append("transmit-laser-power", types.YLeaf{"TransmitLaserPower", digOptMonAlarms.TransmitLaserPower})
    digOptMonAlarms.EntityData.Leafs.Append("received-laser-power", types.YLeaf{"ReceivedLaserPower", digOptMonAlarms.ReceivedLaserPower})
    digOptMonAlarms.EntityData.Leafs.Append("laser-bias-current", types.YLeaf{"LaserBiasCurrent", digOptMonAlarms.LaserBiasCurrent})

    digOptMonAlarms.EntityData.YListKeys = []string {}

    return &(digOptMonAlarms.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane
// Digital Optical Monitoring (per lane
// information)
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Center Wavelength (nm*1000). The type is interface{} with range:
    // 0..4294967295.
    CenterWavelength interface{}

    // Transmit Laser Power (dBm*1000). The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitLaserPower interface{}

    // Received Optical Power (dBm*1000). The type is interface{} with range:
    // -2147483648..2147483647.
    ReceivedLaserPower interface{}

    // Laser Bias Current (uAmps). The type is interface{} with range:
    // 0..4294967295.
    LaserBiasCurrent interface{}

    // Numerical identifier for this lane. The type is interface{} with range:
    // 0..4294967295.
    LaneId interface{}

    // Digital Optical Monitoring alarms.
    DigOptMonAlarm EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm
}

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetEntityData() *types.CommonEntityData {
    lane.EntityData.YFilter = lane.YFilter
    lane.EntityData.YangName = "lane"
    lane.EntityData.BundleName = "cisco_ios_xr"
    lane.EntityData.ParentYangName = "phy-details"
    lane.EntityData.SegmentPath = "lane"
    lane.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lane.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lane.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lane.EntityData.Children = types.NewOrderedMap()
    lane.EntityData.Children.Append("dig-opt-mon-alarm", types.YChild{"DigOptMonAlarm", &lane.DigOptMonAlarm})
    lane.EntityData.Leafs = types.NewOrderedMap()
    lane.EntityData.Leafs.Append("center-wavelength", types.YLeaf{"CenterWavelength", lane.CenterWavelength})
    lane.EntityData.Leafs.Append("transmit-laser-power", types.YLeaf{"TransmitLaserPower", lane.TransmitLaserPower})
    lane.EntityData.Leafs.Append("received-laser-power", types.YLeaf{"ReceivedLaserPower", lane.ReceivedLaserPower})
    lane.EntityData.Leafs.Append("laser-bias-current", types.YLeaf{"LaserBiasCurrent", lane.LaserBiasCurrent})
    lane.EntityData.Leafs.Append("lane-id", types.YLeaf{"LaneId", lane.LaneId})

    lane.EntityData.YListKeys = []string {}

    return &(lane.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm
// Digital Optical Monitoring alarms
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit Laser Power Alarm. The type is EtherDomAlarm.
    TransmitLaserPower interface{}

    // Received Optical Power Alarm. The type is EtherDomAlarm.
    ReceivedLaserPower interface{}

    // Laser Bias Current Alarm. The type is EtherDomAlarm.
    LaserBiasCurrent interface{}
}

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetEntityData() *types.CommonEntityData {
    digOptMonAlarm.EntityData.YFilter = digOptMonAlarm.YFilter
    digOptMonAlarm.EntityData.YangName = "dig-opt-mon-alarm"
    digOptMonAlarm.EntityData.BundleName = "cisco_ios_xr"
    digOptMonAlarm.EntityData.ParentYangName = "lane"
    digOptMonAlarm.EntityData.SegmentPath = "dig-opt-mon-alarm"
    digOptMonAlarm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    digOptMonAlarm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    digOptMonAlarm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    digOptMonAlarm.EntityData.Children = types.NewOrderedMap()
    digOptMonAlarm.EntityData.Leafs = types.NewOrderedMap()
    digOptMonAlarm.EntityData.Leafs.Append("transmit-laser-power", types.YLeaf{"TransmitLaserPower", digOptMonAlarm.TransmitLaserPower})
    digOptMonAlarm.EntityData.Leafs.Append("received-laser-power", types.YLeaf{"ReceivedLaserPower", digOptMonAlarm.ReceivedLaserPower})
    digOptMonAlarm.EntityData.Leafs.Append("laser-bias-current", types.YLeaf{"LaserBiasCurrent", digOptMonAlarm.LaserBiasCurrent})

    digOptMonAlarm.EntityData.YListKeys = []string {}

    return &(digOptMonAlarm.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails
// Forward Error Correction information
type EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port operational FEC type. The type is EthernetFec.
    Fec interface{}

    // Corrected codeword error count. The type is interface{} with range:
    // 0..18446744073709551615.
    CorrectedCodewordCount interface{}

    // Uncorrected codeword error count. The type is interface{} with range:
    // 0..18446744073709551615.
    UncorrectedCodewordCount interface{}
}

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetEntityData() *types.CommonEntityData {
    fecDetails.EntityData.YFilter = fecDetails.YFilter
    fecDetails.EntityData.YangName = "fec-details"
    fecDetails.EntityData.BundleName = "cisco_ios_xr"
    fecDetails.EntityData.ParentYangName = "phy-info"
    fecDetails.EntityData.SegmentPath = "fec-details"
    fecDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fecDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fecDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fecDetails.EntityData.Children = types.NewOrderedMap()
    fecDetails.EntityData.Leafs = types.NewOrderedMap()
    fecDetails.EntityData.Leafs.Append("fec", types.YLeaf{"Fec", fecDetails.Fec})
    fecDetails.EntityData.Leafs.Append("corrected-codeword-count", types.YLeaf{"CorrectedCodewordCount", fecDetails.CorrectedCodewordCount})
    fecDetails.EntityData.Leafs.Append("uncorrected-codeword-count", types.YLeaf{"UncorrectedCodewordCount", fecDetails.UncorrectedCodewordCount})

    fecDetails.EntityData.YListKeys = []string {}

    return &(fecDetails.EntityData)
}

// EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback
// Port operational extended loopback
type EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Level. The type is interface{} with range: 0..4294967295.
    Level interface{}

    // Port operational loopback. The type is EthernetLoopback.
    Loopback interface{}
}

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetEntityData() *types.CommonEntityData {
    extendedLoopback.EntityData.YFilter = extendedLoopback.YFilter
    extendedLoopback.EntityData.YangName = "extended-loopback"
    extendedLoopback.EntityData.BundleName = "cisco_ios_xr"
    extendedLoopback.EntityData.ParentYangName = "phy-info"
    extendedLoopback.EntityData.SegmentPath = "extended-loopback"
    extendedLoopback.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedLoopback.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedLoopback.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedLoopback.EntityData.Children = types.NewOrderedMap()
    extendedLoopback.EntityData.Leafs = types.NewOrderedMap()
    extendedLoopback.EntityData.Leafs.Append("level", types.YLeaf{"Level", extendedLoopback.Level})
    extendedLoopback.EntityData.Leafs.Append("loopback", types.YLeaf{"Loopback", extendedLoopback.Loopback})

    extendedLoopback.EntityData.YListKeys = []string {}

    return &(extendedLoopback.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info
// Layer 1 information
type EthernetInterface_Interfaces_Interface_Layer1Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link state. The type is EtherLinkState.
    LinkState interface{}

    // State of the LED. The type is EtherLedState.
    LedState interface{}

    // Port operational speed. The type is EthernetSpeed.
    Speed interface{}

    // Port operational duplexity. The type is EthernetDuplex.
    Duplex interface{}

    // Port operational flow control. The type is EtherFlowcontrol.
    Flowcontrol interface{}

    // Port operational inter-packet-gap. The type is EthernetIpg.
    Ipg interface{}

    // Laser Squelch - TRUE if enabled. The type is bool.
    LaserSquelchEnabled interface{}

    // Bandwidth utilization (hundredths of a percent). The type is interface{}
    // with range: 0..4294967295. Units are percentage.
    BandwidthUtilization interface{}

    // Port operational bandwidth. The type is interface{} with range:
    // 0..18446744073709551615.
    Bandwidth interface{}

    // Port autonegotiation configuration settings.
    Autoneg EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg

    // Current alarms.
    CurrentAlarms EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms

    // Previous alarms.
    PreviousAlarms EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms

    // Statistics for detected errors.
    ErrorCounts EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts

    // BER monitoring details.
    BerMonitoring EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring

    // OPD monitoring details.
    OpdMonitoring EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring

    // Priority flow control information.
    PfcInfo EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo
}

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetEntityData() *types.CommonEntityData {
    layer1Info.EntityData.YFilter = layer1Info.YFilter
    layer1Info.EntityData.YangName = "layer1-info"
    layer1Info.EntityData.BundleName = "cisco_ios_xr"
    layer1Info.EntityData.ParentYangName = "interface"
    layer1Info.EntityData.SegmentPath = "layer1-info"
    layer1Info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    layer1Info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    layer1Info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    layer1Info.EntityData.Children = types.NewOrderedMap()
    layer1Info.EntityData.Children.Append("autoneg", types.YChild{"Autoneg", &layer1Info.Autoneg})
    layer1Info.EntityData.Children.Append("current-alarms", types.YChild{"CurrentAlarms", &layer1Info.CurrentAlarms})
    layer1Info.EntityData.Children.Append("previous-alarms", types.YChild{"PreviousAlarms", &layer1Info.PreviousAlarms})
    layer1Info.EntityData.Children.Append("error-counts", types.YChild{"ErrorCounts", &layer1Info.ErrorCounts})
    layer1Info.EntityData.Children.Append("ber-monitoring", types.YChild{"BerMonitoring", &layer1Info.BerMonitoring})
    layer1Info.EntityData.Children.Append("opd-monitoring", types.YChild{"OpdMonitoring", &layer1Info.OpdMonitoring})
    layer1Info.EntityData.Children.Append("pfc-info", types.YChild{"PfcInfo", &layer1Info.PfcInfo})
    layer1Info.EntityData.Leafs = types.NewOrderedMap()
    layer1Info.EntityData.Leafs.Append("link-state", types.YLeaf{"LinkState", layer1Info.LinkState})
    layer1Info.EntityData.Leafs.Append("led-state", types.YLeaf{"LedState", layer1Info.LedState})
    layer1Info.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", layer1Info.Speed})
    layer1Info.EntityData.Leafs.Append("duplex", types.YLeaf{"Duplex", layer1Info.Duplex})
    layer1Info.EntityData.Leafs.Append("flowcontrol", types.YLeaf{"Flowcontrol", layer1Info.Flowcontrol})
    layer1Info.EntityData.Leafs.Append("ipg", types.YLeaf{"Ipg", layer1Info.Ipg})
    layer1Info.EntityData.Leafs.Append("laser-squelch-enabled", types.YLeaf{"LaserSquelchEnabled", layer1Info.LaserSquelchEnabled})
    layer1Info.EntityData.Leafs.Append("bandwidth-utilization", types.YLeaf{"BandwidthUtilization", layer1Info.BandwidthUtilization})
    layer1Info.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", layer1Info.Bandwidth})

    layer1Info.EntityData.YListKeys = []string {}

    return &(layer1Info.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg
// Port autonegotiation configuration settings
type EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE if autonegotiation is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    AutonegEnabled interface{}

    // Validity mask: 0x1 speed, 0x2 duplex, 0x4 flowcontrol, 0x8 fec. The type is
    // interface{} with range: 0..4294967295.
    Mask interface{}

    // Restricted speed (if relevant bit is set in mask). The type is
    // EthernetSpeed.
    Speed interface{}

    // Restricted duplex (if relevant bit is set in mask). The type is
    // EthernetDuplex.
    Duplex interface{}

    // Restricted flowcontrol (if relevant bit is set in mask). The type is
    // EtherFlowcontrol.
    Flowcontrol interface{}

    // If true, configuration overrides negotiated settings.  If false, negotiated
    // settings in effect. The type is interface{} with range:
    // -2147483648..2147483647.
    ConfigOverride interface{}

    // Restricted FEC (if revelevant bit is set in mask). The type is EthernetFec.
    Fec interface{}
}

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetEntityData() *types.CommonEntityData {
    autoneg.EntityData.YFilter = autoneg.YFilter
    autoneg.EntityData.YangName = "autoneg"
    autoneg.EntityData.BundleName = "cisco_ios_xr"
    autoneg.EntityData.ParentYangName = "layer1-info"
    autoneg.EntityData.SegmentPath = "autoneg"
    autoneg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoneg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoneg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoneg.EntityData.Children = types.NewOrderedMap()
    autoneg.EntityData.Leafs = types.NewOrderedMap()
    autoneg.EntityData.Leafs.Append("autoneg-enabled", types.YLeaf{"AutonegEnabled", autoneg.AutonegEnabled})
    autoneg.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", autoneg.Mask})
    autoneg.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", autoneg.Speed})
    autoneg.EntityData.Leafs.Append("duplex", types.YLeaf{"Duplex", autoneg.Duplex})
    autoneg.EntityData.Leafs.Append("flowcontrol", types.YLeaf{"Flowcontrol", autoneg.Flowcontrol})
    autoneg.EntityData.Leafs.Append("config-override", types.YLeaf{"ConfigOverride", autoneg.ConfigOverride})
    autoneg.EntityData.Leafs.Append("fec", types.YLeaf{"Fec", autoneg.Fec})

    autoneg.EntityData.YListKeys = []string {}

    return &(autoneg.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms
// Current alarms
type EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received Loss of Signal. The type is EthCtrlrAlarmState.
    ReceivedLossOfSignalAlarm interface{}

    // PCS Loss of Block Lock. The type is EthCtrlrAlarmState.
    PcsLossOfBlockLockAlarm interface{}

    // Local Fault. The type is EthCtrlrAlarmState.
    LocalFaultAlarm interface{}

    // Remote Fault. The type is EthCtrlrAlarmState.
    RemoteFaultAlarm interface{}

    // SD BER. The type is EthCtrlrAlarmState.
    SdBerAlarm interface{}

    // SF BER. The type is EthCtrlrAlarmState.
    SfBerAlarm interface{}

    // Loss of Synchronization Data. The type is EthCtrlrAlarmState.
    LossOfSynchronizationDataAlarm interface{}

    // Hi BER. The type is EthCtrlrAlarmState.
    HiBerAlarm interface{}

    // Squelch. The type is EthCtrlrAlarmState.
    SquelchAlarm interface{}

    // Rx OPD Alarm. The type is EthCtrlrAlarmState.
    RxOpdAlarm interface{}
}

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetEntityData() *types.CommonEntityData {
    currentAlarms.EntityData.YFilter = currentAlarms.YFilter
    currentAlarms.EntityData.YangName = "current-alarms"
    currentAlarms.EntityData.BundleName = "cisco_ios_xr"
    currentAlarms.EntityData.ParentYangName = "layer1-info"
    currentAlarms.EntityData.SegmentPath = "current-alarms"
    currentAlarms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentAlarms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentAlarms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentAlarms.EntityData.Children = types.NewOrderedMap()
    currentAlarms.EntityData.Leafs = types.NewOrderedMap()
    currentAlarms.EntityData.Leafs.Append("received-loss-of-signal-alarm", types.YLeaf{"ReceivedLossOfSignalAlarm", currentAlarms.ReceivedLossOfSignalAlarm})
    currentAlarms.EntityData.Leafs.Append("pcs-loss-of-block-lock-alarm", types.YLeaf{"PcsLossOfBlockLockAlarm", currentAlarms.PcsLossOfBlockLockAlarm})
    currentAlarms.EntityData.Leafs.Append("local-fault-alarm", types.YLeaf{"LocalFaultAlarm", currentAlarms.LocalFaultAlarm})
    currentAlarms.EntityData.Leafs.Append("remote-fault-alarm", types.YLeaf{"RemoteFaultAlarm", currentAlarms.RemoteFaultAlarm})
    currentAlarms.EntityData.Leafs.Append("sd-ber-alarm", types.YLeaf{"SdBerAlarm", currentAlarms.SdBerAlarm})
    currentAlarms.EntityData.Leafs.Append("sf-ber-alarm", types.YLeaf{"SfBerAlarm", currentAlarms.SfBerAlarm})
    currentAlarms.EntityData.Leafs.Append("loss-of-synchronization-data-alarm", types.YLeaf{"LossOfSynchronizationDataAlarm", currentAlarms.LossOfSynchronizationDataAlarm})
    currentAlarms.EntityData.Leafs.Append("hi-ber-alarm", types.YLeaf{"HiBerAlarm", currentAlarms.HiBerAlarm})
    currentAlarms.EntityData.Leafs.Append("squelch-alarm", types.YLeaf{"SquelchAlarm", currentAlarms.SquelchAlarm})
    currentAlarms.EntityData.Leafs.Append("rx-opd-alarm", types.YLeaf{"RxOpdAlarm", currentAlarms.RxOpdAlarm})

    currentAlarms.EntityData.YListKeys = []string {}

    return &(currentAlarms.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms
// Previous alarms
type EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received Loss of Signal. The type is EthCtrlrAlarmState.
    ReceivedLossOfSignalAlarm interface{}

    // PCS Loss of Block Lock. The type is EthCtrlrAlarmState.
    PcsLossOfBlockLockAlarm interface{}

    // Local Fault. The type is EthCtrlrAlarmState.
    LocalFaultAlarm interface{}

    // Remote Fault. The type is EthCtrlrAlarmState.
    RemoteFaultAlarm interface{}

    // SD BER. The type is EthCtrlrAlarmState.
    SdBerAlarm interface{}

    // SF BER. The type is EthCtrlrAlarmState.
    SfBerAlarm interface{}

    // Loss of Synchronization Data. The type is EthCtrlrAlarmState.
    LossOfSynchronizationDataAlarm interface{}

    // Hi BER. The type is EthCtrlrAlarmState.
    HiBerAlarm interface{}

    // Squelch. The type is EthCtrlrAlarmState.
    SquelchAlarm interface{}

    // Rx OPD Alarm. The type is EthCtrlrAlarmState.
    RxOpdAlarm interface{}
}

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetEntityData() *types.CommonEntityData {
    previousAlarms.EntityData.YFilter = previousAlarms.YFilter
    previousAlarms.EntityData.YangName = "previous-alarms"
    previousAlarms.EntityData.BundleName = "cisco_ios_xr"
    previousAlarms.EntityData.ParentYangName = "layer1-info"
    previousAlarms.EntityData.SegmentPath = "previous-alarms"
    previousAlarms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    previousAlarms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    previousAlarms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    previousAlarms.EntityData.Children = types.NewOrderedMap()
    previousAlarms.EntityData.Leafs = types.NewOrderedMap()
    previousAlarms.EntityData.Leafs.Append("received-loss-of-signal-alarm", types.YLeaf{"ReceivedLossOfSignalAlarm", previousAlarms.ReceivedLossOfSignalAlarm})
    previousAlarms.EntityData.Leafs.Append("pcs-loss-of-block-lock-alarm", types.YLeaf{"PcsLossOfBlockLockAlarm", previousAlarms.PcsLossOfBlockLockAlarm})
    previousAlarms.EntityData.Leafs.Append("local-fault-alarm", types.YLeaf{"LocalFaultAlarm", previousAlarms.LocalFaultAlarm})
    previousAlarms.EntityData.Leafs.Append("remote-fault-alarm", types.YLeaf{"RemoteFaultAlarm", previousAlarms.RemoteFaultAlarm})
    previousAlarms.EntityData.Leafs.Append("sd-ber-alarm", types.YLeaf{"SdBerAlarm", previousAlarms.SdBerAlarm})
    previousAlarms.EntityData.Leafs.Append("sf-ber-alarm", types.YLeaf{"SfBerAlarm", previousAlarms.SfBerAlarm})
    previousAlarms.EntityData.Leafs.Append("loss-of-synchronization-data-alarm", types.YLeaf{"LossOfSynchronizationDataAlarm", previousAlarms.LossOfSynchronizationDataAlarm})
    previousAlarms.EntityData.Leafs.Append("hi-ber-alarm", types.YLeaf{"HiBerAlarm", previousAlarms.HiBerAlarm})
    previousAlarms.EntityData.Leafs.Append("squelch-alarm", types.YLeaf{"SquelchAlarm", previousAlarms.SquelchAlarm})
    previousAlarms.EntityData.Leafs.Append("rx-opd-alarm", types.YLeaf{"RxOpdAlarm", previousAlarms.RxOpdAlarm})

    previousAlarms.EntityData.YListKeys = []string {}

    return &(previousAlarms.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts
// Statistics for detected errors
type EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sync-header error count. The type is interface{} with range:
    // 0..18446744073709551615.
    SyncHeaderErrors interface{}

    // PCS BIP error count. The type is interface{} with range:
    // 0..18446744073709551615.
    PcsbipErrors interface{}
}

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetEntityData() *types.CommonEntityData {
    errorCounts.EntityData.YFilter = errorCounts.YFilter
    errorCounts.EntityData.YangName = "error-counts"
    errorCounts.EntityData.BundleName = "cisco_ios_xr"
    errorCounts.EntityData.ParentYangName = "layer1-info"
    errorCounts.EntityData.SegmentPath = "error-counts"
    errorCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorCounts.EntityData.Children = types.NewOrderedMap()
    errorCounts.EntityData.Leafs = types.NewOrderedMap()
    errorCounts.EntityData.Leafs.Append("sync-header-errors", types.YLeaf{"SyncHeaderErrors", errorCounts.SyncHeaderErrors})
    errorCounts.EntityData.Leafs.Append("pcsbip-errors", types.YLeaf{"PcsbipErrors", errorCounts.PcsbipErrors})

    errorCounts.EntityData.YListKeys = []string {}

    return &(errorCounts.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring
// BER monitoring details
type EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether or not BER monitoring is supported. The type is interface{} with
    // range: -2147483648..2147483647.
    Supported interface{}

    // The BER monitoring settings to be applied.
    Settings EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings

    // The BER state.
    State EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_State
}

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetEntityData() *types.CommonEntityData {
    berMonitoring.EntityData.YFilter = berMonitoring.YFilter
    berMonitoring.EntityData.YangName = "ber-monitoring"
    berMonitoring.EntityData.BundleName = "cisco_ios_xr"
    berMonitoring.EntityData.ParentYangName = "layer1-info"
    berMonitoring.EntityData.SegmentPath = "ber-monitoring"
    berMonitoring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    berMonitoring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    berMonitoring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    berMonitoring.EntityData.Children = types.NewOrderedMap()
    berMonitoring.EntityData.Children.Append("settings", types.YChild{"Settings", &berMonitoring.Settings})
    berMonitoring.EntityData.Children.Append("state", types.YChild{"State", &berMonitoring.State})
    berMonitoring.EntityData.Leafs = types.NewOrderedMap()
    berMonitoring.EntityData.Leafs.Append("supported", types.YLeaf{"Supported", berMonitoring.Supported})

    berMonitoring.EntityData.YListKeys = []string {}

    return &(berMonitoring.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings
// The BER monitoring settings to be applied
type EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BER threshold for signal to degrade. The type is interface{} with range:
    // 0..4294967295.
    SignalDegradeThreshold interface{}

    // Report alarm to indicate signal degrade. The type is interface{} with
    // range: -2147483648..2147483647.
    SignalDegradeAlarm interface{}

    // BER threshold for signal to fail. The type is interface{} with range:
    // 0..4294967295.
    SignalFailThreshold interface{}

    // Report alarm to indicate signal failure. The type is interface{} with
    // range: -2147483648..2147483647.
    SignalFailAlarm interface{}

    // Whether drivers should signal remote faults. The type is interface{} with
    // range: -2147483648..2147483647.
    SignalRemoteFault interface{}
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetEntityData() *types.CommonEntityData {
    settings.EntityData.YFilter = settings.YFilter
    settings.EntityData.YangName = "settings"
    settings.EntityData.BundleName = "cisco_ios_xr"
    settings.EntityData.ParentYangName = "ber-monitoring"
    settings.EntityData.SegmentPath = "settings"
    settings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    settings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    settings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    settings.EntityData.Children = types.NewOrderedMap()
    settings.EntityData.Leafs = types.NewOrderedMap()
    settings.EntityData.Leafs.Append("signal-degrade-threshold", types.YLeaf{"SignalDegradeThreshold", settings.SignalDegradeThreshold})
    settings.EntityData.Leafs.Append("signal-degrade-alarm", types.YLeaf{"SignalDegradeAlarm", settings.SignalDegradeAlarm})
    settings.EntityData.Leafs.Append("signal-fail-threshold", types.YLeaf{"SignalFailThreshold", settings.SignalFailThreshold})
    settings.EntityData.Leafs.Append("signal-fail-alarm", types.YLeaf{"SignalFailAlarm", settings.SignalFailAlarm})
    settings.EntityData.Leafs.Append("signal-remote-fault", types.YLeaf{"SignalRemoteFault", settings.SignalRemoteFault})

    settings.EntityData.YListKeys = []string {}

    return &(settings.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_State
// The BER state
type EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current SD-BER. The type is interface{} with range: 0..4294967295.
    SdCurrentBer interface{}

    // Current SF-BER. The type is interface{} with range: 0..4294967295.
    SfCurrentBer interface{}
}

func (state *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "ber-monitoring"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("sd-current-ber", types.YLeaf{"SdCurrentBer", state.SdCurrentBer})
    state.EntityData.Leafs.Append("sf-current-ber", types.YLeaf{"SfCurrentBer", state.SfCurrentBer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring
// OPD monitoring details
type EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether or not OPD monitoring is supported. The type is interface{} with
    // range: -2147483648..2147483647.
    Supported interface{}

    // The OPD monitoring settings to be applied.
    Settings EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings
}

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetEntityData() *types.CommonEntityData {
    opdMonitoring.EntityData.YFilter = opdMonitoring.YFilter
    opdMonitoring.EntityData.YangName = "opd-monitoring"
    opdMonitoring.EntityData.BundleName = "cisco_ios_xr"
    opdMonitoring.EntityData.ParentYangName = "layer1-info"
    opdMonitoring.EntityData.SegmentPath = "opd-monitoring"
    opdMonitoring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opdMonitoring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opdMonitoring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opdMonitoring.EntityData.Children = types.NewOrderedMap()
    opdMonitoring.EntityData.Children.Append("settings", types.YChild{"Settings", &opdMonitoring.Settings})
    opdMonitoring.EntityData.Leafs = types.NewOrderedMap()
    opdMonitoring.EntityData.Leafs.Append("supported", types.YLeaf{"Supported", opdMonitoring.Supported})

    opdMonitoring.EntityData.YListKeys = []string {}

    return &(opdMonitoring.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings
// The OPD monitoring settings to be applied
type EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx-OPD alarm threshold set?. The type is interface{} with range:
    // -2147483648..2147483647.
    ReceivedOpticalPowerDegradeThresholdSet interface{}

    // Rx-OPD alarm threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    ReceivedOpticalPowerDegradeThreshold interface{}
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetEntityData() *types.CommonEntityData {
    settings.EntityData.YFilter = settings.YFilter
    settings.EntityData.YangName = "settings"
    settings.EntityData.BundleName = "cisco_ios_xr"
    settings.EntityData.ParentYangName = "opd-monitoring"
    settings.EntityData.SegmentPath = "settings"
    settings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    settings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    settings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    settings.EntityData.Children = types.NewOrderedMap()
    settings.EntityData.Leafs = types.NewOrderedMap()
    settings.EntityData.Leafs.Append("received-optical-power-degrade-threshold-set", types.YLeaf{"ReceivedOpticalPowerDegradeThresholdSet", settings.ReceivedOpticalPowerDegradeThresholdSet})
    settings.EntityData.Leafs.Append("received-optical-power-degrade-threshold", types.YLeaf{"ReceivedOpticalPowerDegradeThreshold", settings.ReceivedOpticalPowerDegradeThreshold})

    settings.EntityData.YListKeys = []string {}

    return &(settings.EntityData)
}

// EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo
// Priority flow control information
type EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port operational priority flow control. The type is EtherPfc.
    PriorityFlowcontrol interface{}

    // Priority bitmap. The type is interface{} with range: 0..255.
    PriorityEnabledBitmap interface{}

    // RX Frame counts. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    RxFrame []interface{}

    // TX Frame counts. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    TxFrame []interface{}
}

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetEntityData() *types.CommonEntityData {
    pfcInfo.EntityData.YFilter = pfcInfo.YFilter
    pfcInfo.EntityData.YangName = "pfc-info"
    pfcInfo.EntityData.BundleName = "cisco_ios_xr"
    pfcInfo.EntityData.ParentYangName = "layer1-info"
    pfcInfo.EntityData.SegmentPath = "pfc-info"
    pfcInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pfcInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pfcInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pfcInfo.EntityData.Children = types.NewOrderedMap()
    pfcInfo.EntityData.Leafs = types.NewOrderedMap()
    pfcInfo.EntityData.Leafs.Append("priority-flowcontrol", types.YLeaf{"PriorityFlowcontrol", pfcInfo.PriorityFlowcontrol})
    pfcInfo.EntityData.Leafs.Append("priority-enabled-bitmap", types.YLeaf{"PriorityEnabledBitmap", pfcInfo.PriorityEnabledBitmap})
    pfcInfo.EntityData.Leafs.Append("rx-frame", types.YLeaf{"RxFrame", pfcInfo.RxFrame})
    pfcInfo.EntityData.Leafs.Append("tx-frame", types.YLeaf{"TxFrame", pfcInfo.TxFrame})

    pfcInfo.EntityData.YListKeys = []string {}

    return &(pfcInfo.EntityData)
}

// EthernetInterface_Interfaces_Interface_MacInfo
// MAC Layer information
type EthernetInterface_Interfaces_Interface_MacInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port operational MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Port operational MRU. The type is interface{} with range: 0..4294967295.
    Mru interface{}

    // Port Burned-In MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BurnedInMacAddress interface{}

    // Port operational MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    OperationalMacAddress interface{}

    // Port unicast MAC filter information.
    UnicastMacFilters EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters

    // Port multicast MAC filter information.
    MulticastMacFilters EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters
}

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetEntityData() *types.CommonEntityData {
    macInfo.EntityData.YFilter = macInfo.YFilter
    macInfo.EntityData.YangName = "mac-info"
    macInfo.EntityData.BundleName = "cisco_ios_xr"
    macInfo.EntityData.ParentYangName = "interface"
    macInfo.EntityData.SegmentPath = "mac-info"
    macInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macInfo.EntityData.Children = types.NewOrderedMap()
    macInfo.EntityData.Children.Append("unicast-mac-filters", types.YChild{"UnicastMacFilters", &macInfo.UnicastMacFilters})
    macInfo.EntityData.Children.Append("multicast-mac-filters", types.YChild{"MulticastMacFilters", &macInfo.MulticastMacFilters})
    macInfo.EntityData.Leafs = types.NewOrderedMap()
    macInfo.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", macInfo.Mtu})
    macInfo.EntityData.Leafs.Append("mru", types.YLeaf{"Mru", macInfo.Mru})
    macInfo.EntityData.Leafs.Append("burned-in-mac-address", types.YLeaf{"BurnedInMacAddress", macInfo.BurnedInMacAddress})
    macInfo.EntityData.Leafs.Append("operational-mac-address", types.YLeaf{"OperationalMacAddress", macInfo.OperationalMacAddress})

    macInfo.EntityData.YListKeys = []string {}

    return &(macInfo.EntityData)
}

// EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters
// Port unicast MAC filter information
type EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC addresses in the unicast ingress destination MAC filter. The type is
    // slice of string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    UnicastMacAddress []interface{}
}

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetEntityData() *types.CommonEntityData {
    unicastMacFilters.EntityData.YFilter = unicastMacFilters.YFilter
    unicastMacFilters.EntityData.YangName = "unicast-mac-filters"
    unicastMacFilters.EntityData.BundleName = "cisco_ios_xr"
    unicastMacFilters.EntityData.ParentYangName = "mac-info"
    unicastMacFilters.EntityData.SegmentPath = "unicast-mac-filters"
    unicastMacFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unicastMacFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unicastMacFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unicastMacFilters.EntityData.Children = types.NewOrderedMap()
    unicastMacFilters.EntityData.Leafs = types.NewOrderedMap()
    unicastMacFilters.EntityData.Leafs.Append("unicast-mac-address", types.YLeaf{"UnicastMacAddress", unicastMacFilters.UnicastMacAddress})

    unicastMacFilters.EntityData.YListKeys = []string {}

    return &(unicastMacFilters.EntityData)
}

// EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters
// Port multicast MAC filter information
type EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the port is in multicast promiscuous mode. The type is bool.
    MulticastPromiscuous interface{}

    // MAC addresses in the multicast ingress destination MAC filter. The type is
    // slice of
    // EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress.
    MulticastMacAddress []*EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress
}

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetEntityData() *types.CommonEntityData {
    multicastMacFilters.EntityData.YFilter = multicastMacFilters.YFilter
    multicastMacFilters.EntityData.YangName = "multicast-mac-filters"
    multicastMacFilters.EntityData.BundleName = "cisco_ios_xr"
    multicastMacFilters.EntityData.ParentYangName = "mac-info"
    multicastMacFilters.EntityData.SegmentPath = "multicast-mac-filters"
    multicastMacFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastMacFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastMacFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastMacFilters.EntityData.Children = types.NewOrderedMap()
    multicastMacFilters.EntityData.Children.Append("multicast-mac-address", types.YChild{"MulticastMacAddress", nil})
    for i := range multicastMacFilters.MulticastMacAddress {
        multicastMacFilters.EntityData.Children.Append(types.GetSegmentPath(multicastMacFilters.MulticastMacAddress[i]), types.YChild{"MulticastMacAddress", multicastMacFilters.MulticastMacAddress[i]})
    }
    multicastMacFilters.EntityData.Leafs = types.NewOrderedMap()
    multicastMacFilters.EntityData.Leafs.Append("multicast-promiscuous", types.YLeaf{"MulticastPromiscuous", multicastMacFilters.MulticastPromiscuous})

    multicastMacFilters.EntityData.YListKeys = []string {}

    return &(multicastMacFilters.EntityData)
}

// EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress
// MAC addresses in the multicast ingress
// destination MAC filter
type EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Mask for this MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mask interface{}
}

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetEntityData() *types.CommonEntityData {
    multicastMacAddress.EntityData.YFilter = multicastMacAddress.YFilter
    multicastMacAddress.EntityData.YangName = "multicast-mac-address"
    multicastMacAddress.EntityData.BundleName = "cisco_ios_xr"
    multicastMacAddress.EntityData.ParentYangName = "multicast-mac-filters"
    multicastMacAddress.EntityData.SegmentPath = "multicast-mac-address"
    multicastMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastMacAddress.EntityData.Children = types.NewOrderedMap()
    multicastMacAddress.EntityData.Leafs = types.NewOrderedMap()
    multicastMacAddress.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", multicastMacAddress.MacAddress})
    multicastMacAddress.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", multicastMacAddress.Mask})

    multicastMacAddress.EntityData.YListKeys = []string {}

    return &(multicastMacAddress.EntityData)
}

// EthernetInterface_Interfaces_Interface_TransportInfo
// Transport state information
type EthernetInterface_Interfaces_Interface_TransportInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maintenance Mode - TRUE if enabled. The type is bool.
    MaintenanceModeEnabled interface{}

    // AINS Soak status. The type is EtherAinsStatus.
    AinsStatus interface{}

    // Total duration (minutes) of AINS soak timer. The type is interface{} with
    // range: 0..4294967295. Units are minute.
    TotalDuration interface{}

    // Remaining duration (seconds) of AINS soak timer. The type is interface{}
    // with range: 0..4294967295. Units are second.
    RemainingDuration interface{}
}

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetEntityData() *types.CommonEntityData {
    transportInfo.EntityData.YFilter = transportInfo.YFilter
    transportInfo.EntityData.YangName = "transport-info"
    transportInfo.EntityData.BundleName = "cisco_ios_xr"
    transportInfo.EntityData.ParentYangName = "interface"
    transportInfo.EntityData.SegmentPath = "transport-info"
    transportInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportInfo.EntityData.Children = types.NewOrderedMap()
    transportInfo.EntityData.Leafs = types.NewOrderedMap()
    transportInfo.EntityData.Leafs.Append("maintenance-mode-enabled", types.YLeaf{"MaintenanceModeEnabled", transportInfo.MaintenanceModeEnabled})
    transportInfo.EntityData.Leafs.Append("ains-status", types.YLeaf{"AinsStatus", transportInfo.AinsStatus})
    transportInfo.EntityData.Leafs.Append("total-duration", types.YLeaf{"TotalDuration", transportInfo.TotalDuration})
    transportInfo.EntityData.Leafs.Append("remaining-duration", types.YLeaf{"RemainingDuration", transportInfo.RemainingDuration})

    transportInfo.EntityData.YListKeys = []string {}

    return &(transportInfo.EntityData)
}

// EthernetInterface_Berts
// Ethernet controller BERT table
type EthernetInterface_Berts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet BERT information. The type is slice of
    // EthernetInterface_Berts_Bert.
    Bert []*EthernetInterface_Berts_Bert
}

func (berts *EthernetInterface_Berts) GetEntityData() *types.CommonEntityData {
    berts.EntityData.YFilter = berts.YFilter
    berts.EntityData.YangName = "berts"
    berts.EntityData.BundleName = "cisco_ios_xr"
    berts.EntityData.ParentYangName = "ethernet-interface"
    berts.EntityData.SegmentPath = "berts"
    berts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    berts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    berts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    berts.EntityData.Children = types.NewOrderedMap()
    berts.EntityData.Children.Append("bert", types.YChild{"Bert", nil})
    for i := range berts.Bert {
        berts.EntityData.Children.Append(types.GetSegmentPath(berts.Bert[i]), types.YChild{"Bert", berts.Bert[i]})
    }
    berts.EntityData.Leafs = types.NewOrderedMap()

    berts.EntityData.YListKeys = []string {}

    return &(berts.EntityData)
}

// EthernetInterface_Berts_Bert
// Ethernet BERT information
type EthernetInterface_Berts_Bert struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Remaining time for this test in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    TimeLeft interface{}

    // Port BERT interval. The type is interface{} with range: 0..4294967295.
    PortBertInterval interface{}

    // Current test status.
    BertStatus EthernetInterface_Berts_Bert_BertStatus
}

func (bert *EthernetInterface_Berts_Bert) GetEntityData() *types.CommonEntityData {
    bert.EntityData.YFilter = bert.YFilter
    bert.EntityData.YangName = "bert"
    bert.EntityData.BundleName = "cisco_ios_xr"
    bert.EntityData.ParentYangName = "berts"
    bert.EntityData.SegmentPath = "bert" + types.AddKeyToken(bert.InterfaceName, "interface-name")
    bert.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bert.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bert.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bert.EntityData.Children = types.NewOrderedMap()
    bert.EntityData.Children.Append("bert-status", types.YChild{"BertStatus", &bert.BertStatus})
    bert.EntityData.Leafs = types.NewOrderedMap()
    bert.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", bert.InterfaceName})
    bert.EntityData.Leafs.Append("time-left", types.YLeaf{"TimeLeft", bert.TimeLeft})
    bert.EntityData.Leafs.Append("port-bert-interval", types.YLeaf{"PortBertInterval", bert.PortBertInterval})

    bert.EntityData.YListKeys = []string {"InterfaceName"}

    return &(bert.EntityData)
}

// EthernetInterface_Berts_Bert_BertStatus
// Current test status
type EthernetInterface_Berts_Bert_BertStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State. The type is bool.
    BertStateEnabled interface{}

    // Flag indicating available data. The type is interface{} with range:
    // 0..4294967295.
    DataAvailability interface{}

    // Receive count (if 0x1 set in flag). The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveCount interface{}

    // Transmit count (if 0x2 set in flag). The type is interface{} with range:
    // 0..18446744073709551615.
    TransmitCount interface{}

    // Received errors (if 0x4 set in flag). The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveErrors interface{}

    // Bit, block or frame error. The type is EthernetBertErrCnt.
    ErrorType interface{}

    // Test pattern. The type is EthernetBertPattern.
    TestPattern interface{}

    // Device being tested. The type is EthernetDev.
    DeviceUnderTest interface{}

    // Interface being tested. The type is EthernetDevIf.
    InterfaceDevice interface{}
}

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetEntityData() *types.CommonEntityData {
    bertStatus.EntityData.YFilter = bertStatus.YFilter
    bertStatus.EntityData.YangName = "bert-status"
    bertStatus.EntityData.BundleName = "cisco_ios_xr"
    bertStatus.EntityData.ParentYangName = "bert"
    bertStatus.EntityData.SegmentPath = "bert-status"
    bertStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bertStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bertStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bertStatus.EntityData.Children = types.NewOrderedMap()
    bertStatus.EntityData.Leafs = types.NewOrderedMap()
    bertStatus.EntityData.Leafs.Append("bert-state-enabled", types.YLeaf{"BertStateEnabled", bertStatus.BertStateEnabled})
    bertStatus.EntityData.Leafs.Append("data-availability", types.YLeaf{"DataAvailability", bertStatus.DataAvailability})
    bertStatus.EntityData.Leafs.Append("receive-count", types.YLeaf{"ReceiveCount", bertStatus.ReceiveCount})
    bertStatus.EntityData.Leafs.Append("transmit-count", types.YLeaf{"TransmitCount", bertStatus.TransmitCount})
    bertStatus.EntityData.Leafs.Append("receive-errors", types.YLeaf{"ReceiveErrors", bertStatus.ReceiveErrors})
    bertStatus.EntityData.Leafs.Append("error-type", types.YLeaf{"ErrorType", bertStatus.ErrorType})
    bertStatus.EntityData.Leafs.Append("test-pattern", types.YLeaf{"TestPattern", bertStatus.TestPattern})
    bertStatus.EntityData.Leafs.Append("device-under-test", types.YLeaf{"DeviceUnderTest", bertStatus.DeviceUnderTest})
    bertStatus.EntityData.Leafs.Append("interface-device", types.YLeaf{"InterfaceDevice", bertStatus.InterfaceDevice})

    bertStatus.EntityData.YListKeys = []string {}

    return &(bertStatus.EntityData)
}

