// This module contains a collection of YANG definitions
// for Cisco IOS-XR drivers-media-eth package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ethernet-interface: Ethernet operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet controller statistics table.
    Statistics EthernetInterface_Statistics

    // Ethernet controller info table.
    Interfaces EthernetInterface_Interfaces

    // Ethernet controller BERT table.
    Berts EthernetInterface_Berts
}

func (ethernetInterface *EthernetInterface) GetFilter() yfilter.YFilter { return ethernetInterface.YFilter }

func (ethernetInterface *EthernetInterface) SetFilter(yf yfilter.YFilter) { ethernetInterface.YFilter = yf }

func (ethernetInterface *EthernetInterface) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "berts" { return "Berts" }
    return ""
}

func (ethernetInterface *EthernetInterface) GetSegmentPath() string {
    return "Cisco-IOS-XR-drivers-media-eth-oper:ethernet-interface"
}

func (ethernetInterface *EthernetInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &ethernetInterface.Statistics
    }
    if childYangName == "interfaces" {
        return &ethernetInterface.Interfaces
    }
    if childYangName == "berts" {
        return &ethernetInterface.Berts
    }
    return nil
}

func (ethernetInterface *EthernetInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &ethernetInterface.Statistics
    children["interfaces"] = &ethernetInterface.Interfaces
    children["berts"] = &ethernetInterface.Berts
    return children
}

func (ethernetInterface *EthernetInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetInterface *EthernetInterface) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetInterface *EthernetInterface) GetYangName() string { return "ethernet-interface" }

func (ethernetInterface *EthernetInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetInterface *EthernetInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetInterface *EthernetInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetInterface *EthernetInterface) SetParent(parent types.Entity) { ethernetInterface.parent = parent }

func (ethernetInterface *EthernetInterface) GetParent() types.Entity { return ethernetInterface.parent }

func (ethernetInterface *EthernetInterface) GetParentYangName() string { return "Cisco-IOS-XR-drivers-media-eth-oper" }

// EthernetInterface_Statistics
// Ethernet controller statistics table
type EthernetInterface_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet statistics information. The type is slice of
    // EthernetInterface_Statistics_Statistic.
    Statistic []EthernetInterface_Statistics_Statistic
}

func (statistics *EthernetInterface_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *EthernetInterface_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *EthernetInterface_Statistics) GetGoName(yname string) string {
    if yname == "statistic" { return "Statistic" }
    return ""
}

func (statistics *EthernetInterface_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *EthernetInterface_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic" {
        for _, c := range statistics.Statistic {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetInterface_Statistics_Statistic{}
        statistics.Statistic = append(statistics.Statistic, child)
        return &statistics.Statistic[len(statistics.Statistic)-1]
    }
    return nil
}

func (statistics *EthernetInterface_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Statistic {
        children[statistics.Statistic[i].GetSegmentPath()] = &statistics.Statistic[i]
    }
    return children
}

func (statistics *EthernetInterface_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *EthernetInterface_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *EthernetInterface_Statistics) GetYangName() string { return "statistics" }

func (statistics *EthernetInterface_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *EthernetInterface_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *EthernetInterface_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *EthernetInterface_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *EthernetInterface_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *EthernetInterface_Statistics) GetParentYangName() string { return "ethernet-interface" }

// EthernetInterface_Statistics_Statistic
// Ethernet statistics information
type EthernetInterface_Statistics_Statistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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
    Received8021QFrames interface{}

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
    Rfc3635Dot3StatsAlignmentErrors interface{}

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
    Transmitted8021QFrames interface{}

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

func (statistic *EthernetInterface_Statistics_Statistic) GetFilter() yfilter.YFilter { return statistic.YFilter }

func (statistic *EthernetInterface_Statistics_Statistic) SetFilter(yf yfilter.YFilter) { statistic.YFilter = yf }

func (statistic *EthernetInterface_Statistics_Statistic) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "received-total-bytes" { return "ReceivedTotalBytes" }
    if yname == "received-good-bytes" { return "ReceivedGoodBytes" }
    if yname == "received-total-frames" { return "ReceivedTotalFrames" }
    if yname == "received8021q-frames" { return "Received8021QFrames" }
    if yname == "received-pause-frames" { return "ReceivedPauseFrames" }
    if yname == "received-unknown-opcodes" { return "ReceivedUnknownOpcodes" }
    if yname == "received-total64-octet-frames" { return "ReceivedTotal64OctetFrames" }
    if yname == "received-total-octet-frames-from65-to127" { return "ReceivedTotalOctetFramesFrom65To127" }
    if yname == "received-total-octet-frames-from128-to255" { return "ReceivedTotalOctetFramesFrom128To255" }
    if yname == "received-total-octet-frames-from256-to511" { return "ReceivedTotalOctetFramesFrom256To511" }
    if yname == "received-total-octet-frames-from512-to1023" { return "ReceivedTotalOctetFramesFrom512To1023" }
    if yname == "received-total-octet-frames-from1024-to1518" { return "ReceivedTotalOctetFramesFrom1024To1518" }
    if yname == "received-total-octet-frames-from1519-to-max" { return "ReceivedTotalOctetFramesFrom1519ToMax" }
    if yname == "received-good-frames" { return "ReceivedGoodFrames" }
    if yname == "received-unicast-frames" { return "ReceivedUnicastFrames" }
    if yname == "received-multicast-frames" { return "ReceivedMulticastFrames" }
    if yname == "received-broadcast-frames" { return "ReceivedBroadcastFrames" }
    if yname == "number-of-buffer-overrun-packets-dropped" { return "NumberOfBufferOverrunPacketsDropped" }
    if yname == "number-of-aborted-packets-dropped" { return "NumberOfAbortedPacketsDropped" }
    if yname == "numberof-invalid-vlan-id-packets-dropped" { return "NumberofInvalidVlanIdPacketsDropped" }
    if yname == "invalid-dest-mac-drop-packets" { return "InvalidDestMacDropPackets" }
    if yname == "invalid-encap-drop-packets" { return "InvalidEncapDropPackets" }
    if yname == "number-of-miscellaneous-packets-dropped" { return "NumberOfMiscellaneousPacketsDropped" }
    if yname == "dropped-giant-packets-greaterthan-mru" { return "DroppedGiantPacketsGreaterthanMru" }
    if yname == "dropped-ether-stats-undersize-pkts" { return "DroppedEtherStatsUndersizePkts" }
    if yname == "dropped-jabbers-packets-greaterthan-mru" { return "DroppedJabbersPacketsGreaterthanMru" }
    if yname == "dropped-ether-stats-fragments" { return "DroppedEtherStatsFragments" }
    if yname == "dropped-packets-with-crc-align-errors" { return "DroppedPacketsWithCrcAlignErrors" }
    if yname == "ether-stats-collisions" { return "EtherStatsCollisions" }
    if yname == "symbol-errors" { return "SymbolErrors" }
    if yname == "dropped-miscellaneous-error-packets" { return "DroppedMiscellaneousErrorPackets" }
    if yname == "rfc2819-ether-stats-oversized-pkts" { return "Rfc2819EtherStatsOversizedPkts" }
    if yname == "rfc2819-ether-stats-jabbers" { return "Rfc2819EtherStatsJabbers" }
    if yname == "rfc2819-ether-stats-crc-align-errors" { return "Rfc2819EtherStatsCrcAlignErrors" }
    if yname == "rfc3635dot3-stats-alignment-errors" { return "Rfc3635Dot3StatsAlignmentErrors" }
    if yname == "total-bytes-transmitted" { return "TotalBytesTransmitted" }
    if yname == "total-good-bytes-transmitted" { return "TotalGoodBytesTransmitted" }
    if yname == "total-frames-transmitted" { return "TotalFramesTransmitted" }
    if yname == "transmitted8021q-frames" { return "Transmitted8021QFrames" }
    if yname == "transmitted-total-pause-frames" { return "TransmittedTotalPauseFrames" }
    if yname == "transmitted-total64-octet-frames" { return "TransmittedTotal64OctetFrames" }
    if yname == "transmitted-total-octet-frames-from65-to127" { return "TransmittedTotalOctetFramesFrom65To127" }
    if yname == "transmitted-total-octet-frames-from128-to255" { return "TransmittedTotalOctetFramesFrom128To255" }
    if yname == "transmitted-total-octet-frames-from256-to511" { return "TransmittedTotalOctetFramesFrom256To511" }
    if yname == "transmitted-total-octet-frames-from512-to1023" { return "TransmittedTotalOctetFramesFrom512To1023" }
    if yname == "transmitted-total-octet-frames-from1024-to1518" { return "TransmittedTotalOctetFramesFrom1024To1518" }
    if yname == "transmitted-total-octet-frames-from1518-to-max" { return "TransmittedTotalOctetFramesFrom1518ToMax" }
    if yname == "transmitted-good-frames" { return "TransmittedGoodFrames" }
    if yname == "transmitted-unicast-frames" { return "TransmittedUnicastFrames" }
    if yname == "transmitted-multicast-frames" { return "TransmittedMulticastFrames" }
    if yname == "transmitted-broadcast-frames" { return "TransmittedBroadcastFrames" }
    if yname == "buffer-underrun-packet-drops" { return "BufferUnderrunPacketDrops" }
    if yname == "aborted-packet-drops" { return "AbortedPacketDrops" }
    if yname == "uncounted-dropped-frames" { return "UncountedDroppedFrames" }
    if yname == "miscellaneous-output-errors" { return "MiscellaneousOutputErrors" }
    return ""
}

func (statistic *EthernetInterface_Statistics_Statistic) GetSegmentPath() string {
    return "statistic" + "[interface-name='" + fmt.Sprintf("%v", statistic.InterfaceName) + "']"
}

func (statistic *EthernetInterface_Statistics_Statistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistic *EthernetInterface_Statistics_Statistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistic *EthernetInterface_Statistics_Statistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = statistic.InterfaceName
    leafs["received-total-bytes"] = statistic.ReceivedTotalBytes
    leafs["received-good-bytes"] = statistic.ReceivedGoodBytes
    leafs["received-total-frames"] = statistic.ReceivedTotalFrames
    leafs["received8021q-frames"] = statistic.Received8021QFrames
    leafs["received-pause-frames"] = statistic.ReceivedPauseFrames
    leafs["received-unknown-opcodes"] = statistic.ReceivedUnknownOpcodes
    leafs["received-total64-octet-frames"] = statistic.ReceivedTotal64OctetFrames
    leafs["received-total-octet-frames-from65-to127"] = statistic.ReceivedTotalOctetFramesFrom65To127
    leafs["received-total-octet-frames-from128-to255"] = statistic.ReceivedTotalOctetFramesFrom128To255
    leafs["received-total-octet-frames-from256-to511"] = statistic.ReceivedTotalOctetFramesFrom256To511
    leafs["received-total-octet-frames-from512-to1023"] = statistic.ReceivedTotalOctetFramesFrom512To1023
    leafs["received-total-octet-frames-from1024-to1518"] = statistic.ReceivedTotalOctetFramesFrom1024To1518
    leafs["received-total-octet-frames-from1519-to-max"] = statistic.ReceivedTotalOctetFramesFrom1519ToMax
    leafs["received-good-frames"] = statistic.ReceivedGoodFrames
    leafs["received-unicast-frames"] = statistic.ReceivedUnicastFrames
    leafs["received-multicast-frames"] = statistic.ReceivedMulticastFrames
    leafs["received-broadcast-frames"] = statistic.ReceivedBroadcastFrames
    leafs["number-of-buffer-overrun-packets-dropped"] = statistic.NumberOfBufferOverrunPacketsDropped
    leafs["number-of-aborted-packets-dropped"] = statistic.NumberOfAbortedPacketsDropped
    leafs["numberof-invalid-vlan-id-packets-dropped"] = statistic.NumberofInvalidVlanIdPacketsDropped
    leafs["invalid-dest-mac-drop-packets"] = statistic.InvalidDestMacDropPackets
    leafs["invalid-encap-drop-packets"] = statistic.InvalidEncapDropPackets
    leafs["number-of-miscellaneous-packets-dropped"] = statistic.NumberOfMiscellaneousPacketsDropped
    leafs["dropped-giant-packets-greaterthan-mru"] = statistic.DroppedGiantPacketsGreaterthanMru
    leafs["dropped-ether-stats-undersize-pkts"] = statistic.DroppedEtherStatsUndersizePkts
    leafs["dropped-jabbers-packets-greaterthan-mru"] = statistic.DroppedJabbersPacketsGreaterthanMru
    leafs["dropped-ether-stats-fragments"] = statistic.DroppedEtherStatsFragments
    leafs["dropped-packets-with-crc-align-errors"] = statistic.DroppedPacketsWithCrcAlignErrors
    leafs["ether-stats-collisions"] = statistic.EtherStatsCollisions
    leafs["symbol-errors"] = statistic.SymbolErrors
    leafs["dropped-miscellaneous-error-packets"] = statistic.DroppedMiscellaneousErrorPackets
    leafs["rfc2819-ether-stats-oversized-pkts"] = statistic.Rfc2819EtherStatsOversizedPkts
    leafs["rfc2819-ether-stats-jabbers"] = statistic.Rfc2819EtherStatsJabbers
    leafs["rfc2819-ether-stats-crc-align-errors"] = statistic.Rfc2819EtherStatsCrcAlignErrors
    leafs["rfc3635dot3-stats-alignment-errors"] = statistic.Rfc3635Dot3StatsAlignmentErrors
    leafs["total-bytes-transmitted"] = statistic.TotalBytesTransmitted
    leafs["total-good-bytes-transmitted"] = statistic.TotalGoodBytesTransmitted
    leafs["total-frames-transmitted"] = statistic.TotalFramesTransmitted
    leafs["transmitted8021q-frames"] = statistic.Transmitted8021QFrames
    leafs["transmitted-total-pause-frames"] = statistic.TransmittedTotalPauseFrames
    leafs["transmitted-total64-octet-frames"] = statistic.TransmittedTotal64OctetFrames
    leafs["transmitted-total-octet-frames-from65-to127"] = statistic.TransmittedTotalOctetFramesFrom65To127
    leafs["transmitted-total-octet-frames-from128-to255"] = statistic.TransmittedTotalOctetFramesFrom128To255
    leafs["transmitted-total-octet-frames-from256-to511"] = statistic.TransmittedTotalOctetFramesFrom256To511
    leafs["transmitted-total-octet-frames-from512-to1023"] = statistic.TransmittedTotalOctetFramesFrom512To1023
    leafs["transmitted-total-octet-frames-from1024-to1518"] = statistic.TransmittedTotalOctetFramesFrom1024To1518
    leafs["transmitted-total-octet-frames-from1518-to-max"] = statistic.TransmittedTotalOctetFramesFrom1518ToMax
    leafs["transmitted-good-frames"] = statistic.TransmittedGoodFrames
    leafs["transmitted-unicast-frames"] = statistic.TransmittedUnicastFrames
    leafs["transmitted-multicast-frames"] = statistic.TransmittedMulticastFrames
    leafs["transmitted-broadcast-frames"] = statistic.TransmittedBroadcastFrames
    leafs["buffer-underrun-packet-drops"] = statistic.BufferUnderrunPacketDrops
    leafs["aborted-packet-drops"] = statistic.AbortedPacketDrops
    leafs["uncounted-dropped-frames"] = statistic.UncountedDroppedFrames
    leafs["miscellaneous-output-errors"] = statistic.MiscellaneousOutputErrors
    return leafs
}

func (statistic *EthernetInterface_Statistics_Statistic) GetBundleName() string { return "cisco_ios_xr" }

func (statistic *EthernetInterface_Statistics_Statistic) GetYangName() string { return "statistic" }

func (statistic *EthernetInterface_Statistics_Statistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistic *EthernetInterface_Statistics_Statistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistic *EthernetInterface_Statistics_Statistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistic *EthernetInterface_Statistics_Statistic) SetParent(parent types.Entity) { statistic.parent = parent }

func (statistic *EthernetInterface_Statistics_Statistic) GetParent() types.Entity { return statistic.parent }

func (statistic *EthernetInterface_Statistics_Statistic) GetParentYangName() string { return "statistics" }

// EthernetInterface_Interfaces
// Ethernet controller info table
type EthernetInterface_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet controller information. The type is slice of
    // EthernetInterface_Interfaces_Interface.
    Interface []EthernetInterface_Interfaces_Interface
}

func (interfaces *EthernetInterface_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *EthernetInterface_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *EthernetInterface_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *EthernetInterface_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *EthernetInterface_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetInterface_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *EthernetInterface_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *EthernetInterface_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *EthernetInterface_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *EthernetInterface_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *EthernetInterface_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *EthernetInterface_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *EthernetInterface_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *EthernetInterface_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *EthernetInterface_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *EthernetInterface_Interfaces) GetParentYangName() string { return "ethernet-interface" }

// EthernetInterface_Interfaces_Interface
// Ethernet controller information
type EthernetInterface_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (self *EthernetInterface_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *EthernetInterface_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *EthernetInterface_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "oper-state-up" { return "OperStateUp" }
    if yname == "phy-info" { return "PhyInfo" }
    if yname == "layer1-info" { return "Layer1Info" }
    if yname == "mac-info" { return "MacInfo" }
    if yname == "transport-info" { return "TransportInfo" }
    return ""
}

func (self *EthernetInterface_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *EthernetInterface_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "phy-info" {
        return &self.PhyInfo
    }
    if childYangName == "layer1-info" {
        return &self.Layer1Info
    }
    if childYangName == "mac-info" {
        return &self.MacInfo
    }
    if childYangName == "transport-info" {
        return &self.TransportInfo
    }
    return nil
}

func (self *EthernetInterface_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["phy-info"] = &self.PhyInfo
    children["layer1-info"] = &self.Layer1Info
    children["mac-info"] = &self.MacInfo
    children["transport-info"] = &self.TransportInfo
    return children
}

func (self *EthernetInterface_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["admin-state"] = self.AdminState
    leafs["oper-state-up"] = self.OperStateUp
    return leafs
}

func (self *EthernetInterface_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *EthernetInterface_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *EthernetInterface_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *EthernetInterface_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *EthernetInterface_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *EthernetInterface_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *EthernetInterface_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *EthernetInterface_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// EthernetInterface_Interfaces_Interface_PhyInfo
// PHY information
type EthernetInterface_Interfaces_Interface_PhyInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port media type. The type is EthernetMedia.
    MediaType interface{}

    // Presence of PHY. The type is EtherPhyPresent.
    PhyPresent interface{}

    // Port operational loopback. The type is EthernetLoopback.
    Loopback interface{}

    // Details about the PHY.
    PhyDetails EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails

    // Forward Error Correction information.
    FecDetails EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails

    // Port operational extended loopback. The type is slice of
    // EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback.
    ExtendedLoopback []EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback
}

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetFilter() yfilter.YFilter { return phyInfo.YFilter }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) SetFilter(yf yfilter.YFilter) { phyInfo.YFilter = yf }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetGoName(yname string) string {
    if yname == "media-type" { return "MediaType" }
    if yname == "phy-present" { return "PhyPresent" }
    if yname == "loopback" { return "Loopback" }
    if yname == "phy-details" { return "PhyDetails" }
    if yname == "fec-details" { return "FecDetails" }
    if yname == "extended-loopback" { return "ExtendedLoopback" }
    return ""
}

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetSegmentPath() string {
    return "phy-info"
}

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "phy-details" {
        return &phyInfo.PhyDetails
    }
    if childYangName == "fec-details" {
        return &phyInfo.FecDetails
    }
    if childYangName == "extended-loopback" {
        for _, c := range phyInfo.ExtendedLoopback {
            if phyInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback{}
        phyInfo.ExtendedLoopback = append(phyInfo.ExtendedLoopback, child)
        return &phyInfo.ExtendedLoopback[len(phyInfo.ExtendedLoopback)-1]
    }
    return nil
}

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["phy-details"] = &phyInfo.PhyDetails
    children["fec-details"] = &phyInfo.FecDetails
    for i := range phyInfo.ExtendedLoopback {
        children[phyInfo.ExtendedLoopback[i].GetSegmentPath()] = &phyInfo.ExtendedLoopback[i]
    }
    return children
}

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["media-type"] = phyInfo.MediaType
    leafs["phy-present"] = phyInfo.PhyPresent
    leafs["loopback"] = phyInfo.Loopback
    return leafs
}

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetBundleName() string { return "cisco_ios_xr" }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetYangName() string { return "phy-info" }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) SetParent(parent types.Entity) { phyInfo.parent = parent }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetParent() types.Entity { return phyInfo.parent }

func (phyInfo *EthernetInterface_Interfaces_Interface_PhyInfo) GetParentYangName() string { return "interface" }

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails
// Details about the PHY
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails struct {
    parent types.Entity
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
    Lane []EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane
}

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetFilter() yfilter.YFilter { return phyDetails.YFilter }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) SetFilter(yf yfilter.YFilter) { phyDetails.YFilter = yf }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetGoName(yname string) string {
    if yname == "vendor" { return "Vendor" }
    if yname == "vendor-part-number" { return "VendorPartNumber" }
    if yname == "vendor-serial-number" { return "VendorSerialNumber" }
    if yname == "transceiver-temperature" { return "TransceiverTemperature" }
    if yname == "transceiver-voltage" { return "TransceiverVoltage" }
    if yname == "transceiver-tx-power" { return "TransceiverTxPower" }
    if yname == "transceiver-rx-power" { return "TransceiverRxPower" }
    if yname == "transceiver-tx-bias" { return "TransceiverTxBias" }
    if yname == "optics-wavelength" { return "OpticsWavelength" }
    if yname == "optics-type" { return "OpticsType" }
    if yname == "revision-number" { return "RevisionNumber" }
    if yname == "lane-field-validity" { return "LaneFieldValidity" }
    if yname == "dig-opt-mon-alarm-thresholds" { return "DigOptMonAlarmThresholds" }
    if yname == "dig-opt-mon-alarms" { return "DigOptMonAlarms" }
    if yname == "lane" { return "Lane" }
    return ""
}

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetSegmentPath() string {
    return "phy-details"
}

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lane-field-validity" {
        return &phyDetails.LaneFieldValidity
    }
    if childYangName == "dig-opt-mon-alarm-thresholds" {
        return &phyDetails.DigOptMonAlarmThresholds
    }
    if childYangName == "dig-opt-mon-alarms" {
        return &phyDetails.DigOptMonAlarms
    }
    if childYangName == "lane" {
        for _, c := range phyDetails.Lane {
            if phyDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane{}
        phyDetails.Lane = append(phyDetails.Lane, child)
        return &phyDetails.Lane[len(phyDetails.Lane)-1]
    }
    return nil
}

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lane-field-validity"] = &phyDetails.LaneFieldValidity
    children["dig-opt-mon-alarm-thresholds"] = &phyDetails.DigOptMonAlarmThresholds
    children["dig-opt-mon-alarms"] = &phyDetails.DigOptMonAlarms
    for i := range phyDetails.Lane {
        children[phyDetails.Lane[i].GetSegmentPath()] = &phyDetails.Lane[i]
    }
    return children
}

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vendor"] = phyDetails.Vendor
    leafs["vendor-part-number"] = phyDetails.VendorPartNumber
    leafs["vendor-serial-number"] = phyDetails.VendorSerialNumber
    leafs["transceiver-temperature"] = phyDetails.TransceiverTemperature
    leafs["transceiver-voltage"] = phyDetails.TransceiverVoltage
    leafs["transceiver-tx-power"] = phyDetails.TransceiverTxPower
    leafs["transceiver-rx-power"] = phyDetails.TransceiverRxPower
    leafs["transceiver-tx-bias"] = phyDetails.TransceiverTxBias
    leafs["optics-wavelength"] = phyDetails.OpticsWavelength
    leafs["optics-type"] = phyDetails.OpticsType
    leafs["revision-number"] = phyDetails.RevisionNumber
    return leafs
}

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetBundleName() string { return "cisco_ios_xr" }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetYangName() string { return "phy-details" }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) SetParent(parent types.Entity) { phyDetails.parent = parent }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetParent() types.Entity { return phyDetails.parent }

func (phyDetails *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails) GetParentYangName() string { return "phy-info" }

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity
// Digital Optical Monitoring (per lane
// information) validity
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity struct {
    parent types.Entity
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

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetFilter() yfilter.YFilter { return laneFieldValidity.YFilter }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) SetFilter(yf yfilter.YFilter) { laneFieldValidity.YFilter = yf }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetGoName(yname string) string {
    if yname == "wavelength-valid" { return "WavelengthValid" }
    if yname == "transmit-power-valid" { return "TransmitPowerValid" }
    if yname == "receive-power-valid" { return "ReceivePowerValid" }
    if yname == "laser-bias-valid" { return "LaserBiasValid" }
    return ""
}

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetSegmentPath() string {
    return "lane-field-validity"
}

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["wavelength-valid"] = laneFieldValidity.WavelengthValid
    leafs["transmit-power-valid"] = laneFieldValidity.TransmitPowerValid
    leafs["receive-power-valid"] = laneFieldValidity.ReceivePowerValid
    leafs["laser-bias-valid"] = laneFieldValidity.LaserBiasValid
    return leafs
}

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetBundleName() string { return "cisco_ios_xr" }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetYangName() string { return "lane-field-validity" }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) SetParent(parent types.Entity) { laneFieldValidity.parent = parent }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetParent() types.Entity { return laneFieldValidity.parent }

func (laneFieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_LaneFieldValidity) GetParentYangName() string { return "phy-details" }

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds
// Digital Optical Monitoring alarm thresholds
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds struct {
    parent types.Entity
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

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetFilter() yfilter.YFilter { return digOptMonAlarmThresholds.YFilter }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) SetFilter(yf yfilter.YFilter) { digOptMonAlarmThresholds.YFilter = yf }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetGoName(yname string) string {
    if yname == "transceiver-temperature-alarm-high" { return "TransceiverTemperatureAlarmHigh" }
    if yname == "transceiver-temperature-warning-high" { return "TransceiverTemperatureWarningHigh" }
    if yname == "transceiver-temperature-warning-low" { return "TransceiverTemperatureWarningLow" }
    if yname == "transceiver-temperature-alarm-low" { return "TransceiverTemperatureAlarmLow" }
    if yname == "transceiver-voltage-alarm-high" { return "TransceiverVoltageAlarmHigh" }
    if yname == "transceiver-voltage-warning-high" { return "TransceiverVoltageWarningHigh" }
    if yname == "transceiver-voltage-warning-low" { return "TransceiverVoltageWarningLow" }
    if yname == "transceiver-voltage-alarm-low" { return "TransceiverVoltageAlarmLow" }
    if yname == "laser-bias-alarm-high" { return "LaserBiasAlarmHigh" }
    if yname == "laser-bias-warning-high" { return "LaserBiasWarningHigh" }
    if yname == "laser-bias-warning-low" { return "LaserBiasWarningLow" }
    if yname == "laser-bias-alarm-low" { return "LaserBiasAlarmLow" }
    if yname == "optical-transmit-power-alarm-high" { return "OpticalTransmitPowerAlarmHigh" }
    if yname == "optical-transmit-power-warning-high" { return "OpticalTransmitPowerWarningHigh" }
    if yname == "optical-transmit-power-warning-low" { return "OpticalTransmitPowerWarningLow" }
    if yname == "optical-transmit-power-alarm-low" { return "OpticalTransmitPowerAlarmLow" }
    if yname == "optical-receive-power-alarm-high" { return "OpticalReceivePowerAlarmHigh" }
    if yname == "optical-receive-power-warning-high" { return "OpticalReceivePowerWarningHigh" }
    if yname == "optical-receive-power-warning-low" { return "OpticalReceivePowerWarningLow" }
    if yname == "optical-receive-power-alarm-low" { return "OpticalReceivePowerAlarmLow" }
    if yname == "field-validity" { return "FieldValidity" }
    return ""
}

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetSegmentPath() string {
    return "dig-opt-mon-alarm-thresholds"
}

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "field-validity" {
        return &digOptMonAlarmThresholds.FieldValidity
    }
    return nil
}

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["field-validity"] = &digOptMonAlarmThresholds.FieldValidity
    return children
}

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transceiver-temperature-alarm-high"] = digOptMonAlarmThresholds.TransceiverTemperatureAlarmHigh
    leafs["transceiver-temperature-warning-high"] = digOptMonAlarmThresholds.TransceiverTemperatureWarningHigh
    leafs["transceiver-temperature-warning-low"] = digOptMonAlarmThresholds.TransceiverTemperatureWarningLow
    leafs["transceiver-temperature-alarm-low"] = digOptMonAlarmThresholds.TransceiverTemperatureAlarmLow
    leafs["transceiver-voltage-alarm-high"] = digOptMonAlarmThresholds.TransceiverVoltageAlarmHigh
    leafs["transceiver-voltage-warning-high"] = digOptMonAlarmThresholds.TransceiverVoltageWarningHigh
    leafs["transceiver-voltage-warning-low"] = digOptMonAlarmThresholds.TransceiverVoltageWarningLow
    leafs["transceiver-voltage-alarm-low"] = digOptMonAlarmThresholds.TransceiverVoltageAlarmLow
    leafs["laser-bias-alarm-high"] = digOptMonAlarmThresholds.LaserBiasAlarmHigh
    leafs["laser-bias-warning-high"] = digOptMonAlarmThresholds.LaserBiasWarningHigh
    leafs["laser-bias-warning-low"] = digOptMonAlarmThresholds.LaserBiasWarningLow
    leafs["laser-bias-alarm-low"] = digOptMonAlarmThresholds.LaserBiasAlarmLow
    leafs["optical-transmit-power-alarm-high"] = digOptMonAlarmThresholds.OpticalTransmitPowerAlarmHigh
    leafs["optical-transmit-power-warning-high"] = digOptMonAlarmThresholds.OpticalTransmitPowerWarningHigh
    leafs["optical-transmit-power-warning-low"] = digOptMonAlarmThresholds.OpticalTransmitPowerWarningLow
    leafs["optical-transmit-power-alarm-low"] = digOptMonAlarmThresholds.OpticalTransmitPowerAlarmLow
    leafs["optical-receive-power-alarm-high"] = digOptMonAlarmThresholds.OpticalReceivePowerAlarmHigh
    leafs["optical-receive-power-warning-high"] = digOptMonAlarmThresholds.OpticalReceivePowerWarningHigh
    leafs["optical-receive-power-warning-low"] = digOptMonAlarmThresholds.OpticalReceivePowerWarningLow
    leafs["optical-receive-power-alarm-low"] = digOptMonAlarmThresholds.OpticalReceivePowerAlarmLow
    return leafs
}

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetBundleName() string { return "cisco_ios_xr" }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetYangName() string { return "dig-opt-mon-alarm-thresholds" }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) SetParent(parent types.Entity) { digOptMonAlarmThresholds.parent = parent }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetParent() types.Entity { return digOptMonAlarmThresholds.parent }

func (digOptMonAlarmThresholds *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds) GetParentYangName() string { return "phy-details" }

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity
// Field validity
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity struct {
    parent types.Entity
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

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetFilter() yfilter.YFilter { return fieldValidity.YFilter }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) SetFilter(yf yfilter.YFilter) { fieldValidity.YFilter = yf }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetGoName(yname string) string {
    if yname == "temperature-valid" { return "TemperatureValid" }
    if yname == "voltage-valid" { return "VoltageValid" }
    if yname == "laser-bias-valid" { return "LaserBiasValid" }
    if yname == "transmit-power-valid" { return "TransmitPowerValid" }
    if yname == "receive-power-valid" { return "ReceivePowerValid" }
    return ""
}

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetSegmentPath() string {
    return "field-validity"
}

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["temperature-valid"] = fieldValidity.TemperatureValid
    leafs["voltage-valid"] = fieldValidity.VoltageValid
    leafs["laser-bias-valid"] = fieldValidity.LaserBiasValid
    leafs["transmit-power-valid"] = fieldValidity.TransmitPowerValid
    leafs["receive-power-valid"] = fieldValidity.ReceivePowerValid
    return leafs
}

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetBundleName() string { return "cisco_ios_xr" }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetYangName() string { return "field-validity" }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) SetParent(parent types.Entity) { fieldValidity.parent = parent }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetParent() types.Entity { return fieldValidity.parent }

func (fieldValidity *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarmThresholds_FieldValidity) GetParentYangName() string { return "dig-opt-mon-alarm-thresholds" }

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms
// Digital Optical Monitoring alarms
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms struct {
    parent types.Entity
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

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetFilter() yfilter.YFilter { return digOptMonAlarms.YFilter }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) SetFilter(yf yfilter.YFilter) { digOptMonAlarms.YFilter = yf }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetGoName(yname string) string {
    if yname == "transceiver-temperature" { return "TransceiverTemperature" }
    if yname == "transceiver-voltage" { return "TransceiverVoltage" }
    if yname == "transmit-laser-power" { return "TransmitLaserPower" }
    if yname == "received-laser-power" { return "ReceivedLaserPower" }
    if yname == "laser-bias-current" { return "LaserBiasCurrent" }
    return ""
}

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetSegmentPath() string {
    return "dig-opt-mon-alarms"
}

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transceiver-temperature"] = digOptMonAlarms.TransceiverTemperature
    leafs["transceiver-voltage"] = digOptMonAlarms.TransceiverVoltage
    leafs["transmit-laser-power"] = digOptMonAlarms.TransmitLaserPower
    leafs["received-laser-power"] = digOptMonAlarms.ReceivedLaserPower
    leafs["laser-bias-current"] = digOptMonAlarms.LaserBiasCurrent
    return leafs
}

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetBundleName() string { return "cisco_ios_xr" }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetYangName() string { return "dig-opt-mon-alarms" }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) SetParent(parent types.Entity) { digOptMonAlarms.parent = parent }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetParent() types.Entity { return digOptMonAlarms.parent }

func (digOptMonAlarms *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_DigOptMonAlarms) GetParentYangName() string { return "phy-details" }

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane
// Digital Optical Monitoring (per lane
// information)
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane struct {
    parent types.Entity
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

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetFilter() yfilter.YFilter { return lane.YFilter }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) SetFilter(yf yfilter.YFilter) { lane.YFilter = yf }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetGoName(yname string) string {
    if yname == "center-wavelength" { return "CenterWavelength" }
    if yname == "transmit-laser-power" { return "TransmitLaserPower" }
    if yname == "received-laser-power" { return "ReceivedLaserPower" }
    if yname == "laser-bias-current" { return "LaserBiasCurrent" }
    if yname == "lane-id" { return "LaneId" }
    if yname == "dig-opt-mon-alarm" { return "DigOptMonAlarm" }
    return ""
}

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetSegmentPath() string {
    return "lane"
}

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dig-opt-mon-alarm" {
        return &lane.DigOptMonAlarm
    }
    return nil
}

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dig-opt-mon-alarm"] = &lane.DigOptMonAlarm
    return children
}

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["center-wavelength"] = lane.CenterWavelength
    leafs["transmit-laser-power"] = lane.TransmitLaserPower
    leafs["received-laser-power"] = lane.ReceivedLaserPower
    leafs["laser-bias-current"] = lane.LaserBiasCurrent
    leafs["lane-id"] = lane.LaneId
    return leafs
}

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetBundleName() string { return "cisco_ios_xr" }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetYangName() string { return "lane" }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) SetParent(parent types.Entity) { lane.parent = parent }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetParent() types.Entity { return lane.parent }

func (lane *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane) GetParentYangName() string { return "phy-details" }

// EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm
// Digital Optical Monitoring alarms
type EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmit Laser Power Alarm. The type is EtherDomAlarm.
    TransmitLaserPower interface{}

    // Received Optical Power Alarm. The type is EtherDomAlarm.
    ReceivedLaserPower interface{}

    // Laser Bias Current Alarm. The type is EtherDomAlarm.
    LaserBiasCurrent interface{}
}

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetFilter() yfilter.YFilter { return digOptMonAlarm.YFilter }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) SetFilter(yf yfilter.YFilter) { digOptMonAlarm.YFilter = yf }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetGoName(yname string) string {
    if yname == "transmit-laser-power" { return "TransmitLaserPower" }
    if yname == "received-laser-power" { return "ReceivedLaserPower" }
    if yname == "laser-bias-current" { return "LaserBiasCurrent" }
    return ""
}

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetSegmentPath() string {
    return "dig-opt-mon-alarm"
}

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmit-laser-power"] = digOptMonAlarm.TransmitLaserPower
    leafs["received-laser-power"] = digOptMonAlarm.ReceivedLaserPower
    leafs["laser-bias-current"] = digOptMonAlarm.LaserBiasCurrent
    return leafs
}

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetBundleName() string { return "cisco_ios_xr" }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetYangName() string { return "dig-opt-mon-alarm" }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) SetParent(parent types.Entity) { digOptMonAlarm.parent = parent }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetParent() types.Entity { return digOptMonAlarm.parent }

func (digOptMonAlarm *EthernetInterface_Interfaces_Interface_PhyInfo_PhyDetails_Lane_DigOptMonAlarm) GetParentYangName() string { return "lane" }

// EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails
// Forward Error Correction information
type EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails struct {
    parent types.Entity
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

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetFilter() yfilter.YFilter { return fecDetails.YFilter }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) SetFilter(yf yfilter.YFilter) { fecDetails.YFilter = yf }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetGoName(yname string) string {
    if yname == "fec" { return "Fec" }
    if yname == "corrected-codeword-count" { return "CorrectedCodewordCount" }
    if yname == "uncorrected-codeword-count" { return "UncorrectedCodewordCount" }
    return ""
}

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetSegmentPath() string {
    return "fec-details"
}

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fec"] = fecDetails.Fec
    leafs["corrected-codeword-count"] = fecDetails.CorrectedCodewordCount
    leafs["uncorrected-codeword-count"] = fecDetails.UncorrectedCodewordCount
    return leafs
}

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetBundleName() string { return "cisco_ios_xr" }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetYangName() string { return "fec-details" }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) SetParent(parent types.Entity) { fecDetails.parent = parent }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetParent() types.Entity { return fecDetails.parent }

func (fecDetails *EthernetInterface_Interfaces_Interface_PhyInfo_FecDetails) GetParentYangName() string { return "phy-info" }

// EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback
// Port operational extended loopback
type EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Level. The type is interface{} with range: 0..4294967295.
    Level interface{}

    // Port operational loopback. The type is EthernetLoopback.
    Loopback interface{}
}

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetFilter() yfilter.YFilter { return extendedLoopback.YFilter }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) SetFilter(yf yfilter.YFilter) { extendedLoopback.YFilter = yf }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "loopback" { return "Loopback" }
    return ""
}

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetSegmentPath() string {
    return "extended-loopback"
}

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = extendedLoopback.Level
    leafs["loopback"] = extendedLoopback.Loopback
    return leafs
}

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetBundleName() string { return "cisco_ios_xr" }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetYangName() string { return "extended-loopback" }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) SetParent(parent types.Entity) { extendedLoopback.parent = parent }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetParent() types.Entity { return extendedLoopback.parent }

func (extendedLoopback *EthernetInterface_Interfaces_Interface_PhyInfo_ExtendedLoopback) GetParentYangName() string { return "phy-info" }

// EthernetInterface_Interfaces_Interface_Layer1Info
// Layer 1 information
type EthernetInterface_Interfaces_Interface_Layer1Info struct {
    parent types.Entity
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

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetFilter() yfilter.YFilter { return layer1Info.YFilter }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) SetFilter(yf yfilter.YFilter) { layer1Info.YFilter = yf }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetGoName(yname string) string {
    if yname == "link-state" { return "LinkState" }
    if yname == "led-state" { return "LedState" }
    if yname == "speed" { return "Speed" }
    if yname == "duplex" { return "Duplex" }
    if yname == "flowcontrol" { return "Flowcontrol" }
    if yname == "ipg" { return "Ipg" }
    if yname == "laser-squelch-enabled" { return "LaserSquelchEnabled" }
    if yname == "bandwidth-utilization" { return "BandwidthUtilization" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "autoneg" { return "Autoneg" }
    if yname == "current-alarms" { return "CurrentAlarms" }
    if yname == "previous-alarms" { return "PreviousAlarms" }
    if yname == "error-counts" { return "ErrorCounts" }
    if yname == "ber-monitoring" { return "BerMonitoring" }
    if yname == "opd-monitoring" { return "OpdMonitoring" }
    if yname == "pfc-info" { return "PfcInfo" }
    return ""
}

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetSegmentPath() string {
    return "layer1-info"
}

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "autoneg" {
        return &layer1Info.Autoneg
    }
    if childYangName == "current-alarms" {
        return &layer1Info.CurrentAlarms
    }
    if childYangName == "previous-alarms" {
        return &layer1Info.PreviousAlarms
    }
    if childYangName == "error-counts" {
        return &layer1Info.ErrorCounts
    }
    if childYangName == "ber-monitoring" {
        return &layer1Info.BerMonitoring
    }
    if childYangName == "opd-monitoring" {
        return &layer1Info.OpdMonitoring
    }
    if childYangName == "pfc-info" {
        return &layer1Info.PfcInfo
    }
    return nil
}

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["autoneg"] = &layer1Info.Autoneg
    children["current-alarms"] = &layer1Info.CurrentAlarms
    children["previous-alarms"] = &layer1Info.PreviousAlarms
    children["error-counts"] = &layer1Info.ErrorCounts
    children["ber-monitoring"] = &layer1Info.BerMonitoring
    children["opd-monitoring"] = &layer1Info.OpdMonitoring
    children["pfc-info"] = &layer1Info.PfcInfo
    return children
}

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-state"] = layer1Info.LinkState
    leafs["led-state"] = layer1Info.LedState
    leafs["speed"] = layer1Info.Speed
    leafs["duplex"] = layer1Info.Duplex
    leafs["flowcontrol"] = layer1Info.Flowcontrol
    leafs["ipg"] = layer1Info.Ipg
    leafs["laser-squelch-enabled"] = layer1Info.LaserSquelchEnabled
    leafs["bandwidth-utilization"] = layer1Info.BandwidthUtilization
    leafs["bandwidth"] = layer1Info.Bandwidth
    return leafs
}

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetBundleName() string { return "cisco_ios_xr" }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetYangName() string { return "layer1-info" }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) SetParent(parent types.Entity) { layer1Info.parent = parent }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetParent() types.Entity { return layer1Info.parent }

func (layer1Info *EthernetInterface_Interfaces_Interface_Layer1Info) GetParentYangName() string { return "interface" }

// EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg
// Port autonegotiation configuration settings
type EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg struct {
    parent types.Entity
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

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetFilter() yfilter.YFilter { return autoneg.YFilter }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) SetFilter(yf yfilter.YFilter) { autoneg.YFilter = yf }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetGoName(yname string) string {
    if yname == "autoneg-enabled" { return "AutonegEnabled" }
    if yname == "mask" { return "Mask" }
    if yname == "speed" { return "Speed" }
    if yname == "duplex" { return "Duplex" }
    if yname == "flowcontrol" { return "Flowcontrol" }
    if yname == "config-override" { return "ConfigOverride" }
    if yname == "fec" { return "Fec" }
    return ""
}

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetSegmentPath() string {
    return "autoneg"
}

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["autoneg-enabled"] = autoneg.AutonegEnabled
    leafs["mask"] = autoneg.Mask
    leafs["speed"] = autoneg.Speed
    leafs["duplex"] = autoneg.Duplex
    leafs["flowcontrol"] = autoneg.Flowcontrol
    leafs["config-override"] = autoneg.ConfigOverride
    leafs["fec"] = autoneg.Fec
    return leafs
}

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetBundleName() string { return "cisco_ios_xr" }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetYangName() string { return "autoneg" }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) SetParent(parent types.Entity) { autoneg.parent = parent }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetParent() types.Entity { return autoneg.parent }

func (autoneg *EthernetInterface_Interfaces_Interface_Layer1Info_Autoneg) GetParentYangName() string { return "layer1-info" }

// EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms
// Current alarms
type EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms struct {
    parent types.Entity
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

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetFilter() yfilter.YFilter { return currentAlarms.YFilter }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) SetFilter(yf yfilter.YFilter) { currentAlarms.YFilter = yf }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetGoName(yname string) string {
    if yname == "received-loss-of-signal-alarm" { return "ReceivedLossOfSignalAlarm" }
    if yname == "pcs-loss-of-block-lock-alarm" { return "PcsLossOfBlockLockAlarm" }
    if yname == "local-fault-alarm" { return "LocalFaultAlarm" }
    if yname == "remote-fault-alarm" { return "RemoteFaultAlarm" }
    if yname == "sd-ber-alarm" { return "SdBerAlarm" }
    if yname == "sf-ber-alarm" { return "SfBerAlarm" }
    if yname == "loss-of-synchronization-data-alarm" { return "LossOfSynchronizationDataAlarm" }
    if yname == "hi-ber-alarm" { return "HiBerAlarm" }
    if yname == "squelch-alarm" { return "SquelchAlarm" }
    if yname == "rx-opd-alarm" { return "RxOpdAlarm" }
    return ""
}

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetSegmentPath() string {
    return "current-alarms"
}

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-loss-of-signal-alarm"] = currentAlarms.ReceivedLossOfSignalAlarm
    leafs["pcs-loss-of-block-lock-alarm"] = currentAlarms.PcsLossOfBlockLockAlarm
    leafs["local-fault-alarm"] = currentAlarms.LocalFaultAlarm
    leafs["remote-fault-alarm"] = currentAlarms.RemoteFaultAlarm
    leafs["sd-ber-alarm"] = currentAlarms.SdBerAlarm
    leafs["sf-ber-alarm"] = currentAlarms.SfBerAlarm
    leafs["loss-of-synchronization-data-alarm"] = currentAlarms.LossOfSynchronizationDataAlarm
    leafs["hi-ber-alarm"] = currentAlarms.HiBerAlarm
    leafs["squelch-alarm"] = currentAlarms.SquelchAlarm
    leafs["rx-opd-alarm"] = currentAlarms.RxOpdAlarm
    return leafs
}

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetBundleName() string { return "cisco_ios_xr" }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetYangName() string { return "current-alarms" }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) SetParent(parent types.Entity) { currentAlarms.parent = parent }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetParent() types.Entity { return currentAlarms.parent }

func (currentAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_CurrentAlarms) GetParentYangName() string { return "layer1-info" }

// EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms
// Previous alarms
type EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms struct {
    parent types.Entity
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

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetFilter() yfilter.YFilter { return previousAlarms.YFilter }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) SetFilter(yf yfilter.YFilter) { previousAlarms.YFilter = yf }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetGoName(yname string) string {
    if yname == "received-loss-of-signal-alarm" { return "ReceivedLossOfSignalAlarm" }
    if yname == "pcs-loss-of-block-lock-alarm" { return "PcsLossOfBlockLockAlarm" }
    if yname == "local-fault-alarm" { return "LocalFaultAlarm" }
    if yname == "remote-fault-alarm" { return "RemoteFaultAlarm" }
    if yname == "sd-ber-alarm" { return "SdBerAlarm" }
    if yname == "sf-ber-alarm" { return "SfBerAlarm" }
    if yname == "loss-of-synchronization-data-alarm" { return "LossOfSynchronizationDataAlarm" }
    if yname == "hi-ber-alarm" { return "HiBerAlarm" }
    if yname == "squelch-alarm" { return "SquelchAlarm" }
    if yname == "rx-opd-alarm" { return "RxOpdAlarm" }
    return ""
}

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetSegmentPath() string {
    return "previous-alarms"
}

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-loss-of-signal-alarm"] = previousAlarms.ReceivedLossOfSignalAlarm
    leafs["pcs-loss-of-block-lock-alarm"] = previousAlarms.PcsLossOfBlockLockAlarm
    leafs["local-fault-alarm"] = previousAlarms.LocalFaultAlarm
    leafs["remote-fault-alarm"] = previousAlarms.RemoteFaultAlarm
    leafs["sd-ber-alarm"] = previousAlarms.SdBerAlarm
    leafs["sf-ber-alarm"] = previousAlarms.SfBerAlarm
    leafs["loss-of-synchronization-data-alarm"] = previousAlarms.LossOfSynchronizationDataAlarm
    leafs["hi-ber-alarm"] = previousAlarms.HiBerAlarm
    leafs["squelch-alarm"] = previousAlarms.SquelchAlarm
    leafs["rx-opd-alarm"] = previousAlarms.RxOpdAlarm
    return leafs
}

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetBundleName() string { return "cisco_ios_xr" }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetYangName() string { return "previous-alarms" }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) SetParent(parent types.Entity) { previousAlarms.parent = parent }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetParent() types.Entity { return previousAlarms.parent }

func (previousAlarms *EthernetInterface_Interfaces_Interface_Layer1Info_PreviousAlarms) GetParentYangName() string { return "layer1-info" }

// EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts
// Statistics for detected errors
type EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sync-header error count. The type is interface{} with range:
    // 0..18446744073709551615.
    SyncHeaderErrors interface{}

    // PCS BIP error count. The type is interface{} with range:
    // 0..18446744073709551615.
    PcsbipErrors interface{}
}

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetFilter() yfilter.YFilter { return errorCounts.YFilter }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) SetFilter(yf yfilter.YFilter) { errorCounts.YFilter = yf }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetGoName(yname string) string {
    if yname == "sync-header-errors" { return "SyncHeaderErrors" }
    if yname == "pcsbip-errors" { return "PcsbipErrors" }
    return ""
}

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetSegmentPath() string {
    return "error-counts"
}

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sync-header-errors"] = errorCounts.SyncHeaderErrors
    leafs["pcsbip-errors"] = errorCounts.PcsbipErrors
    return leafs
}

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetBundleName() string { return "cisco_ios_xr" }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetYangName() string { return "error-counts" }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) SetParent(parent types.Entity) { errorCounts.parent = parent }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetParent() types.Entity { return errorCounts.parent }

func (errorCounts *EthernetInterface_Interfaces_Interface_Layer1Info_ErrorCounts) GetParentYangName() string { return "layer1-info" }

// EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring
// BER monitoring details
type EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether or not BER monitoring is supported. The type is interface{} with
    // range: -2147483648..2147483647.
    Supported interface{}

    // The BER monitoring settings to be applied.
    Settings EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings
}

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetFilter() yfilter.YFilter { return berMonitoring.YFilter }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) SetFilter(yf yfilter.YFilter) { berMonitoring.YFilter = yf }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetGoName(yname string) string {
    if yname == "supported" { return "Supported" }
    if yname == "settings" { return "Settings" }
    return ""
}

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetSegmentPath() string {
    return "ber-monitoring"
}

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "settings" {
        return &berMonitoring.Settings
    }
    return nil
}

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["settings"] = &berMonitoring.Settings
    return children
}

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["supported"] = berMonitoring.Supported
    return leafs
}

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetBundleName() string { return "cisco_ios_xr" }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetYangName() string { return "ber-monitoring" }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) SetParent(parent types.Entity) { berMonitoring.parent = parent }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetParent() types.Entity { return berMonitoring.parent }

func (berMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring) GetParentYangName() string { return "layer1-info" }

// EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings
// The BER monitoring settings to be applied
type EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings struct {
    parent types.Entity
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

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetFilter() yfilter.YFilter { return settings.YFilter }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) SetFilter(yf yfilter.YFilter) { settings.YFilter = yf }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetGoName(yname string) string {
    if yname == "signal-degrade-threshold" { return "SignalDegradeThreshold" }
    if yname == "signal-degrade-alarm" { return "SignalDegradeAlarm" }
    if yname == "signal-fail-threshold" { return "SignalFailThreshold" }
    if yname == "signal-fail-alarm" { return "SignalFailAlarm" }
    if yname == "signal-remote-fault" { return "SignalRemoteFault" }
    return ""
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetSegmentPath() string {
    return "settings"
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["signal-degrade-threshold"] = settings.SignalDegradeThreshold
    leafs["signal-degrade-alarm"] = settings.SignalDegradeAlarm
    leafs["signal-fail-threshold"] = settings.SignalFailThreshold
    leafs["signal-fail-alarm"] = settings.SignalFailAlarm
    leafs["signal-remote-fault"] = settings.SignalRemoteFault
    return leafs
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetBundleName() string { return "cisco_ios_xr" }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetYangName() string { return "settings" }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) SetParent(parent types.Entity) { settings.parent = parent }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetParent() types.Entity { return settings.parent }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_BerMonitoring_Settings) GetParentYangName() string { return "ber-monitoring" }

// EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring
// OPD monitoring details
type EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether or not OPD monitoring is supported. The type is interface{} with
    // range: -2147483648..2147483647.
    Supported interface{}

    // The OPD monitoring settings to be applied.
    Settings EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings
}

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetFilter() yfilter.YFilter { return opdMonitoring.YFilter }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) SetFilter(yf yfilter.YFilter) { opdMonitoring.YFilter = yf }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetGoName(yname string) string {
    if yname == "supported" { return "Supported" }
    if yname == "settings" { return "Settings" }
    return ""
}

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetSegmentPath() string {
    return "opd-monitoring"
}

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "settings" {
        return &opdMonitoring.Settings
    }
    return nil
}

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["settings"] = &opdMonitoring.Settings
    return children
}

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["supported"] = opdMonitoring.Supported
    return leafs
}

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetBundleName() string { return "cisco_ios_xr" }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetYangName() string { return "opd-monitoring" }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) SetParent(parent types.Entity) { opdMonitoring.parent = parent }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetParent() types.Entity { return opdMonitoring.parent }

func (opdMonitoring *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring) GetParentYangName() string { return "layer1-info" }

// EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings
// The OPD monitoring settings to be applied
type EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rx-OPD alarm threshold set?. The type is interface{} with range:
    // -2147483648..2147483647.
    ReceivedOpticalPowerDegradeThresholdSet interface{}

    // Rx-OPD alarm threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    ReceivedOpticalPowerDegradeThreshold interface{}
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetFilter() yfilter.YFilter { return settings.YFilter }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) SetFilter(yf yfilter.YFilter) { settings.YFilter = yf }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetGoName(yname string) string {
    if yname == "received-optical-power-degrade-threshold-set" { return "ReceivedOpticalPowerDegradeThresholdSet" }
    if yname == "received-optical-power-degrade-threshold" { return "ReceivedOpticalPowerDegradeThreshold" }
    return ""
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetSegmentPath() string {
    return "settings"
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-optical-power-degrade-threshold-set"] = settings.ReceivedOpticalPowerDegradeThresholdSet
    leafs["received-optical-power-degrade-threshold"] = settings.ReceivedOpticalPowerDegradeThreshold
    return leafs
}

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetBundleName() string { return "cisco_ios_xr" }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetYangName() string { return "settings" }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) SetParent(parent types.Entity) { settings.parent = parent }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetParent() types.Entity { return settings.parent }

func (settings *EthernetInterface_Interfaces_Interface_Layer1Info_OpdMonitoring_Settings) GetParentYangName() string { return "opd-monitoring" }

// EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo
// Priority flow control information
type EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo struct {
    parent types.Entity
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

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetFilter() yfilter.YFilter { return pfcInfo.YFilter }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) SetFilter(yf yfilter.YFilter) { pfcInfo.YFilter = yf }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetGoName(yname string) string {
    if yname == "priority-flowcontrol" { return "PriorityFlowcontrol" }
    if yname == "priority-enabled-bitmap" { return "PriorityEnabledBitmap" }
    if yname == "rx-frame" { return "RxFrame" }
    if yname == "tx-frame" { return "TxFrame" }
    return ""
}

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetSegmentPath() string {
    return "pfc-info"
}

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority-flowcontrol"] = pfcInfo.PriorityFlowcontrol
    leafs["priority-enabled-bitmap"] = pfcInfo.PriorityEnabledBitmap
    leafs["rx-frame"] = pfcInfo.RxFrame
    leafs["tx-frame"] = pfcInfo.TxFrame
    return leafs
}

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetBundleName() string { return "cisco_ios_xr" }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetYangName() string { return "pfc-info" }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) SetParent(parent types.Entity) { pfcInfo.parent = parent }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetParent() types.Entity { return pfcInfo.parent }

func (pfcInfo *EthernetInterface_Interfaces_Interface_Layer1Info_PfcInfo) GetParentYangName() string { return "layer1-info" }

// EthernetInterface_Interfaces_Interface_MacInfo
// MAC Layer information
type EthernetInterface_Interfaces_Interface_MacInfo struct {
    parent types.Entity
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

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetFilter() yfilter.YFilter { return macInfo.YFilter }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) SetFilter(yf yfilter.YFilter) { macInfo.YFilter = yf }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetGoName(yname string) string {
    if yname == "mtu" { return "Mtu" }
    if yname == "mru" { return "Mru" }
    if yname == "burned-in-mac-address" { return "BurnedInMacAddress" }
    if yname == "operational-mac-address" { return "OperationalMacAddress" }
    if yname == "unicast-mac-filters" { return "UnicastMacFilters" }
    if yname == "multicast-mac-filters" { return "MulticastMacFilters" }
    return ""
}

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetSegmentPath() string {
    return "mac-info"
}

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unicast-mac-filters" {
        return &macInfo.UnicastMacFilters
    }
    if childYangName == "multicast-mac-filters" {
        return &macInfo.MulticastMacFilters
    }
    return nil
}

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unicast-mac-filters"] = &macInfo.UnicastMacFilters
    children["multicast-mac-filters"] = &macInfo.MulticastMacFilters
    return children
}

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mtu"] = macInfo.Mtu
    leafs["mru"] = macInfo.Mru
    leafs["burned-in-mac-address"] = macInfo.BurnedInMacAddress
    leafs["operational-mac-address"] = macInfo.OperationalMacAddress
    return leafs
}

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetBundleName() string { return "cisco_ios_xr" }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetYangName() string { return "mac-info" }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) SetParent(parent types.Entity) { macInfo.parent = parent }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetParent() types.Entity { return macInfo.parent }

func (macInfo *EthernetInterface_Interfaces_Interface_MacInfo) GetParentYangName() string { return "interface" }

// EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters
// Port unicast MAC filter information
type EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC addresses in the unicast ingress destination MAC filter. The type is
    // slice of string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    UnicastMacAddress []interface{}
}

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetFilter() yfilter.YFilter { return unicastMacFilters.YFilter }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) SetFilter(yf yfilter.YFilter) { unicastMacFilters.YFilter = yf }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetGoName(yname string) string {
    if yname == "unicast-mac-address" { return "UnicastMacAddress" }
    return ""
}

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetSegmentPath() string {
    return "unicast-mac-filters"
}

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unicast-mac-address"] = unicastMacFilters.UnicastMacAddress
    return leafs
}

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetBundleName() string { return "cisco_ios_xr" }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetYangName() string { return "unicast-mac-filters" }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) SetParent(parent types.Entity) { unicastMacFilters.parent = parent }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetParent() types.Entity { return unicastMacFilters.parent }

func (unicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_UnicastMacFilters) GetParentYangName() string { return "mac-info" }

// EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters
// Port multicast MAC filter information
type EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the port is in multicast promiscuous mode. The type is bool.
    MulticastPromiscuous interface{}

    // MAC addresses in the multicast ingress destination MAC filter. The type is
    // slice of
    // EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress.
    MulticastMacAddress []EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress
}

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetFilter() yfilter.YFilter { return multicastMacFilters.YFilter }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) SetFilter(yf yfilter.YFilter) { multicastMacFilters.YFilter = yf }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetGoName(yname string) string {
    if yname == "multicast-promiscuous" { return "MulticastPromiscuous" }
    if yname == "multicast-mac-address" { return "MulticastMacAddress" }
    return ""
}

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetSegmentPath() string {
    return "multicast-mac-filters"
}

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multicast-mac-address" {
        for _, c := range multicastMacFilters.MulticastMacAddress {
            if multicastMacFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress{}
        multicastMacFilters.MulticastMacAddress = append(multicastMacFilters.MulticastMacAddress, child)
        return &multicastMacFilters.MulticastMacAddress[len(multicastMacFilters.MulticastMacAddress)-1]
    }
    return nil
}

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range multicastMacFilters.MulticastMacAddress {
        children[multicastMacFilters.MulticastMacAddress[i].GetSegmentPath()] = &multicastMacFilters.MulticastMacAddress[i]
    }
    return children
}

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["multicast-promiscuous"] = multicastMacFilters.MulticastPromiscuous
    return leafs
}

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetBundleName() string { return "cisco_ios_xr" }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetYangName() string { return "multicast-mac-filters" }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) SetParent(parent types.Entity) { multicastMacFilters.parent = parent }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetParent() types.Entity { return multicastMacFilters.parent }

func (multicastMacFilters *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters) GetParentYangName() string { return "mac-info" }

// EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress
// MAC addresses in the multicast ingress
// destination MAC filter
type EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Mask for this MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mask interface{}
}

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetFilter() yfilter.YFilter { return multicastMacAddress.YFilter }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) SetFilter(yf yfilter.YFilter) { multicastMacAddress.YFilter = yf }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "mask" { return "Mask" }
    return ""
}

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetSegmentPath() string {
    return "multicast-mac-address"
}

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = multicastMacAddress.MacAddress
    leafs["mask"] = multicastMacAddress.Mask
    return leafs
}

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetYangName() string { return "multicast-mac-address" }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) SetParent(parent types.Entity) { multicastMacAddress.parent = parent }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetParent() types.Entity { return multicastMacAddress.parent }

func (multicastMacAddress *EthernetInterface_Interfaces_Interface_MacInfo_MulticastMacFilters_MulticastMacAddress) GetParentYangName() string { return "multicast-mac-filters" }

// EthernetInterface_Interfaces_Interface_TransportInfo
// Transport state information
type EthernetInterface_Interfaces_Interface_TransportInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maintenance Mode - TRUE if enabled. The type is bool.
    MaintenanceModeEnabled interface{}

    // AINS Soak status. The type is EtherAinsStatus.
    AinsStatus interface{}

    // Total duration (seconds) of AINS soak timer. The type is interface{} with
    // range: 0..4294967295. Units are second.
    TotalDuration interface{}

    // Remaining duration (seconds) of AINS soak timer. The type is interface{}
    // with range: 0..4294967295. Units are second.
    RemainingDuration interface{}
}

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetFilter() yfilter.YFilter { return transportInfo.YFilter }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) SetFilter(yf yfilter.YFilter) { transportInfo.YFilter = yf }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetGoName(yname string) string {
    if yname == "maintenance-mode-enabled" { return "MaintenanceModeEnabled" }
    if yname == "ains-status" { return "AinsStatus" }
    if yname == "total-duration" { return "TotalDuration" }
    if yname == "remaining-duration" { return "RemainingDuration" }
    return ""
}

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetSegmentPath() string {
    return "transport-info"
}

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maintenance-mode-enabled"] = transportInfo.MaintenanceModeEnabled
    leafs["ains-status"] = transportInfo.AinsStatus
    leafs["total-duration"] = transportInfo.TotalDuration
    leafs["remaining-duration"] = transportInfo.RemainingDuration
    return leafs
}

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetBundleName() string { return "cisco_ios_xr" }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetYangName() string { return "transport-info" }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) SetParent(parent types.Entity) { transportInfo.parent = parent }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetParent() types.Entity { return transportInfo.parent }

func (transportInfo *EthernetInterface_Interfaces_Interface_TransportInfo) GetParentYangName() string { return "interface" }

// EthernetInterface_Berts
// Ethernet controller BERT table
type EthernetInterface_Berts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet BERT information. The type is slice of
    // EthernetInterface_Berts_Bert.
    Bert []EthernetInterface_Berts_Bert
}

func (berts *EthernetInterface_Berts) GetFilter() yfilter.YFilter { return berts.YFilter }

func (berts *EthernetInterface_Berts) SetFilter(yf yfilter.YFilter) { berts.YFilter = yf }

func (berts *EthernetInterface_Berts) GetGoName(yname string) string {
    if yname == "bert" { return "Bert" }
    return ""
}

func (berts *EthernetInterface_Berts) GetSegmentPath() string {
    return "berts"
}

func (berts *EthernetInterface_Berts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bert" {
        for _, c := range berts.Bert {
            if berts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetInterface_Berts_Bert{}
        berts.Bert = append(berts.Bert, child)
        return &berts.Bert[len(berts.Bert)-1]
    }
    return nil
}

func (berts *EthernetInterface_Berts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range berts.Bert {
        children[berts.Bert[i].GetSegmentPath()] = &berts.Bert[i]
    }
    return children
}

func (berts *EthernetInterface_Berts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (berts *EthernetInterface_Berts) GetBundleName() string { return "cisco_ios_xr" }

func (berts *EthernetInterface_Berts) GetYangName() string { return "berts" }

func (berts *EthernetInterface_Berts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (berts *EthernetInterface_Berts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (berts *EthernetInterface_Berts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (berts *EthernetInterface_Berts) SetParent(parent types.Entity) { berts.parent = parent }

func (berts *EthernetInterface_Berts) GetParent() types.Entity { return berts.parent }

func (berts *EthernetInterface_Berts) GetParentYangName() string { return "ethernet-interface" }

// EthernetInterface_Berts_Bert
// Ethernet BERT information
type EthernetInterface_Berts_Bert struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Remaining time for this test in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    TimeLeft interface{}

    // Port BERT interval. The type is interface{} with range: 0..4294967295.
    PortBertInterval interface{}

    // Current test status.
    BertStatus EthernetInterface_Berts_Bert_BertStatus
}

func (bert *EthernetInterface_Berts_Bert) GetFilter() yfilter.YFilter { return bert.YFilter }

func (bert *EthernetInterface_Berts_Bert) SetFilter(yf yfilter.YFilter) { bert.YFilter = yf }

func (bert *EthernetInterface_Berts_Bert) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "time-left" { return "TimeLeft" }
    if yname == "port-bert-interval" { return "PortBertInterval" }
    if yname == "bert-status" { return "BertStatus" }
    return ""
}

func (bert *EthernetInterface_Berts_Bert) GetSegmentPath() string {
    return "bert" + "[interface-name='" + fmt.Sprintf("%v", bert.InterfaceName) + "']"
}

func (bert *EthernetInterface_Berts_Bert) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bert-status" {
        return &bert.BertStatus
    }
    return nil
}

func (bert *EthernetInterface_Berts_Bert) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bert-status"] = &bert.BertStatus
    return children
}

func (bert *EthernetInterface_Berts_Bert) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bert.InterfaceName
    leafs["time-left"] = bert.TimeLeft
    leafs["port-bert-interval"] = bert.PortBertInterval
    return leafs
}

func (bert *EthernetInterface_Berts_Bert) GetBundleName() string { return "cisco_ios_xr" }

func (bert *EthernetInterface_Berts_Bert) GetYangName() string { return "bert" }

func (bert *EthernetInterface_Berts_Bert) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bert *EthernetInterface_Berts_Bert) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bert *EthernetInterface_Berts_Bert) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bert *EthernetInterface_Berts_Bert) SetParent(parent types.Entity) { bert.parent = parent }

func (bert *EthernetInterface_Berts_Bert) GetParent() types.Entity { return bert.parent }

func (bert *EthernetInterface_Berts_Bert) GetParentYangName() string { return "berts" }

// EthernetInterface_Berts_Bert_BertStatus
// Current test status
type EthernetInterface_Berts_Bert_BertStatus struct {
    parent types.Entity
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

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetFilter() yfilter.YFilter { return bertStatus.YFilter }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) SetFilter(yf yfilter.YFilter) { bertStatus.YFilter = yf }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetGoName(yname string) string {
    if yname == "bert-state-enabled" { return "BertStateEnabled" }
    if yname == "data-availability" { return "DataAvailability" }
    if yname == "receive-count" { return "ReceiveCount" }
    if yname == "transmit-count" { return "TransmitCount" }
    if yname == "receive-errors" { return "ReceiveErrors" }
    if yname == "error-type" { return "ErrorType" }
    if yname == "test-pattern" { return "TestPattern" }
    if yname == "device-under-test" { return "DeviceUnderTest" }
    if yname == "interface-device" { return "InterfaceDevice" }
    return ""
}

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetSegmentPath() string {
    return "bert-status"
}

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bert-state-enabled"] = bertStatus.BertStateEnabled
    leafs["data-availability"] = bertStatus.DataAvailability
    leafs["receive-count"] = bertStatus.ReceiveCount
    leafs["transmit-count"] = bertStatus.TransmitCount
    leafs["receive-errors"] = bertStatus.ReceiveErrors
    leafs["error-type"] = bertStatus.ErrorType
    leafs["test-pattern"] = bertStatus.TestPattern
    leafs["device-under-test"] = bertStatus.DeviceUnderTest
    leafs["interface-device"] = bertStatus.InterfaceDevice
    return leafs
}

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetBundleName() string { return "cisco_ios_xr" }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetYangName() string { return "bert-status" }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) SetParent(parent types.Entity) { bertStatus.parent = parent }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetParent() types.Entity { return bertStatus.parent }

func (bertStatus *EthernetInterface_Berts_Bert_BertStatus) GetParentYangName() string { return "bert" }

