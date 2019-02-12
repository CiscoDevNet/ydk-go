// This module contains a collection of YANG definitions
// for Cisco IOS-XR controller-optics package operational data.
// 
// This module contains definitions
// for the following management objects:
//   optics-oper: Optics operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package controller_optics_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package controller_optics_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-controller-optics-oper optics-oper}", reflect.TypeOf(OpticsOper{}))
    ydk.RegisterEntity("Cisco-IOS-XR-controller-optics-oper:optics-oper", reflect.TypeOf(OpticsOper{}))
}

// OpticsAmplifierGainRange represents Optics amplifier gain range
type OpticsAmplifierGainRange string

const (
    // Invalid
    OpticsAmplifierGainRange_optics_amplifier_gain_range_invalid OpticsAmplifierGainRange = "optics-amplifier-gain-range-invalid"

    // Normal
    OpticsAmplifierGainRange_optics_amplifier_gain_range_normal OpticsAmplifierGainRange = "optics-amplifier-gain-range-normal"

    // Extended
    OpticsAmplifierGainRange_optics_amplifier_gain_range_ext_end_ed OpticsAmplifierGainRange = "optics-amplifier-gain-range-ext-end-ed"
)

// OpticsModulation represents Optics modulation
type OpticsModulation string

const (
    // BPSK
    OpticsModulation_mod_bpsk OpticsModulation = "mod-bpsk"

    // QPSK
    OpticsModulation_mod_qpsk OpticsModulation = "mod-qpsk"

    // 8QAM
    OpticsModulation_mod_8qam OpticsModulation = "mod-8qam"

    // 16QAM
    OpticsModulation_mod_16qam OpticsModulation = "mod-16qam"

    // 32QAM
    OpticsModulation_mod_32qam OpticsModulation = "mod-32qam"

    // 64QAM
    OpticsModulation_mod_64qam OpticsModulation = "mod-64qam"

    // BPSK QPSK
    OpticsModulation_mod_bpsk_qpsk OpticsModulation = "mod-bpsk-qpsk"

    // QPSK 8QAM
    OpticsModulation_mod_qpsk_8qam OpticsModulation = "mod-qpsk-8qam"

    // 8QAM 16QAM
    OpticsModulation_mod_8qam_16qam OpticsModulation = "mod-8qam-16qam"

    // 16QAM 32QAM
    OpticsModulation_mode_16qam_32qam OpticsModulation = "mode-16qam-32qam"

    // 32QAM 64QAM
    OpticsModulation_mod_32qam_64qam OpticsModulation = "mod-32qam-64qam"
)

// OpticsAmplifierControlMode represents Optics amplifier control mode
type OpticsAmplifierControlMode string

const (
    // Automatic
    OpticsAmplifierControlMode_automatic OpticsAmplifierControlMode = "automatic"

    // Manual
    OpticsAmplifierControlMode_manual OpticsAmplifierControlMode = "manual"
)

// EthernetPmd represents Ethernet Pmd Type
type EthernetPmd string

const (
    // Not set
    EthernetPmd_optics_eth_not_set EthernetPmd = "optics-eth-not-set"

    // 10GBASE LRM
    EthernetPmd_optics_eth_10gbase_lrm EthernetPmd = "optics-eth-10gbase-lrm"

    // 10GBASE LR
    EthernetPmd_optics_eth_10gbase_lr EthernetPmd = "optics-eth-10gbase-lr"

    // 10GBASE ZR
    EthernetPmd_optics_eth_10gbase_zr EthernetPmd = "optics-eth-10gbase-zr"

    // 10GBASE ER
    EthernetPmd_optics_eth_10gbase_er EthernetPmd = "optics-eth-10gbase-er"

    // 10GBASE SR
    EthernetPmd_optics_eth_10gbase_sr EthernetPmd = "optics-eth-10gbase-sr"

    // 10GBASE T
    EthernetPmd_optics_eth_10gbase EthernetPmd = "optics-eth-10gbase"

    // 40GBASE CR4
    EthernetPmd_optics_eth_40gbase_cr4 EthernetPmd = "optics-eth-40gbase-cr4"

    // 40GBASE SR4
    EthernetPmd_optics_eth_40gbase_sr4 EthernetPmd = "optics-eth-40gbase-sr4"

    // 40GBASE LR4
    EthernetPmd_optics_eth_40gbase_lr4 EthernetPmd = "optics-eth-40gbase-lr4"

    // 40GBASE ER4
    EthernetPmd_optics_eth_40gbase_er4 EthernetPmd = "optics-eth-40gbase-er4"

    // 40GBASE PSM4
    EthernetPmd_optics_eth_40gbase_psm4 EthernetPmd = "optics-eth-40gbase-psm4"

    // 40GBASE CSR4
    EthernetPmd_optics_eth_40gbase_csr4 EthernetPmd = "optics-eth-40gbase-csr4"

    // 40GBASE SR BD
    EthernetPmd_optics_eth_40gbase_sr_bd EthernetPmd = "optics-eth-40gbase-sr-bd"

    // 40G AOC
    EthernetPmd_optics_eth_40g_aoc EthernetPmd = "optics-eth-40g-aoc"

    // 4X10GBASE LR
    EthernetPmd_optics_eth_4x10gbase_lr EthernetPmd = "optics-eth-4x10gbase-lr"

    // 4X10GBASE SR
    EthernetPmd_optics_eth_4x10gbase_sr EthernetPmd = "optics-eth-4x10gbase-sr"

    // 100G AOC
    EthernetPmd_optics_eth_100g_aoc EthernetPmd = "optics-eth-100g-aoc"

    // 100G ACC
    EthernetPmd_optics_eth_100g_acc EthernetPmd = "optics-eth-100g-acc"

    // 100GBASE SR10
    EthernetPmd_optics_eth_100gbase_sr10 EthernetPmd = "optics-eth-100gbase-sr10"

    // 100GBASE SR4
    EthernetPmd_optics_eth_100gbase_sr4 EthernetPmd = "optics-eth-100gbase-sr4"

    // 100GBASE LR4
    EthernetPmd_optics_eth_100gbase_lr4 EthernetPmd = "optics-eth-100gbase-lr4"

    // 100GBASE ER4
    EthernetPmd_optics_eth_100gbase_er4 EthernetPmd = "optics-eth-100gbase-er4"

    // 100GBASE CWDM4
    EthernetPmd_optics_eth_100gbase_cwdm4 EthernetPmd = "optics-eth-100gbase-cwdm4"

    // 100GBASE CLR4
    EthernetPmd_optics_eth_100gbase_clr4 EthernetPmd = "optics-eth-100gbase-clr4"

    // 100GBASE PSM4
    EthernetPmd_optics_eth_100gbase_psm4 EthernetPmd = "optics-eth-100gbase-psm4"

    // 100GBASE CR4
    EthernetPmd_optics_eth_100gbase_cr4 EthernetPmd = "optics-eth-100gbase-cr4"

    // 100GBASE AL
    EthernetPmd_optics_eth_100gbase_al EthernetPmd = "optics-eth-100gbase-al"

    // 100GBASE PL
    EthernetPmd_optics_eth_100gbase_pl EthernetPmd = "optics-eth-100gbase-pl"

    // Eth Undefined
    EthernetPmd_optics_eth_undefined EthernetPmd = "optics-eth-undefined"
)

// OpticsWaveBand represents Optics wave band
type OpticsWaveBand string

const (
    // C BAND
    OpticsWaveBand_c_band OpticsWaveBand = "c-band"

    // L BAND
    OpticsWaveBand_l_band OpticsWaveBand = "l-band"

    // C ODD
    OpticsWaveBand_c_band_odd OpticsWaveBand = "c-band-odd"

    // C EVEN
    OpticsWaveBand_c_band_even OpticsWaveBand = "c-band-even"

    // Invalid
    OpticsWaveBand_invalid_band OpticsWaveBand = "invalid-band"
)

// OpticsAinsStateEt represents Optics ains state et
type OpticsAinsStateEt string

const (
    // None
    OpticsAinsStateEt_none OpticsAinsStateEt = "none"

    // Running
    OpticsAinsStateEt_active_running OpticsAinsStateEt = "active-running"

    // Pending
    OpticsAinsStateEt_active_pending OpticsAinsStateEt = "active-pending"
)

// FiberConnector represents Fiber connector
type FiberConnector string

const (
    // Not Set
    FiberConnector_optics_connect_or_not_set FiberConnector = "optics-connect-or-not-set"

    // SC
    FiberConnector_optics_sc_connect_or FiberConnector = "optics-sc-connect-or"

    // LC
    FiberConnector_optics_lc_connect_or FiberConnector = "optics-lc-connect-or"

    // MPO
    FiberConnector_optics_mpo_connect_or FiberConnector = "optics-mpo-connect-or"

    // Undefined
    FiberConnector_optics_undefined_connect_or FiberConnector = "optics-undefined-connect-or"
)

// OpticsFormFactor represents Optics form factor
type OpticsFormFactor string

const (
    // Not set
    OpticsFormFactor_not_set OpticsFormFactor = "not-set"

    // Invalid
    OpticsFormFactor_invalid OpticsFormFactor = "invalid"

    // CPAK
    OpticsFormFactor_cpak OpticsFormFactor = "cpak"

    // CXP
    OpticsFormFactor_cxp OpticsFormFactor = "cxp"

    // SFP+
    OpticsFormFactor_sfp_plus OpticsFormFactor = "sfp-plus"

    // QSFP
    OpticsFormFactor_qsfp OpticsFormFactor = "qsfp"

    // QSFP+
    OpticsFormFactor_qsfp_plus OpticsFormFactor = "qsfp-plus"

    // QSFP28
    OpticsFormFactor_qsfp28 OpticsFormFactor = "qsfp28"

    // SFP
    OpticsFormFactor_sfp OpticsFormFactor = "sfp"

    // CFP
    OpticsFormFactor_cfp OpticsFormFactor = "cfp"

    // CFP2
    OpticsFormFactor_cfp2 OpticsFormFactor = "cfp2"

    // CFP2 ACO
    OpticsFormFactor_cfp2_aco OpticsFormFactor = "cfp2-aco"

    // CFP2 DCO
    OpticsFormFactor_cfp2_dco OpticsFormFactor = "cfp2-dco"

    // CFP4
    OpticsFormFactor_cfp4 OpticsFormFactor = "cfp4"

    // XFP
    OpticsFormFactor_xfp OpticsFormFactor = "xfp"

    // X2
    OpticsFormFactor_x2 OpticsFormFactor = "x2"

    // Non pluggable
    OpticsFormFactor_non_pluggable OpticsFormFactor = "non-pluggable"

    // Other
    OpticsFormFactor_other OpticsFormFactor = "other"
)

// SonetApplicationCode represents Sonet application code
type SonetApplicationCode string

const (
    // Not Set
    SonetApplicationCode_optics_sonet_not_set SonetApplicationCode = "optics-sonet-not-set"

    // VSR2000 3R2
    SonetApplicationCode_optics_vsr2000_3r2 SonetApplicationCode = "optics-vsr2000-3r2"

    // VSR2000 3R3
    SonetApplicationCode_optics_vsr2000_3r3 SonetApplicationCode = "optics-vsr2000-3r3"

    // VSR2000 3R5
    SonetApplicationCode_optics_vsr2000_3r5 SonetApplicationCode = "optics-vsr2000-3r5"

    // Undefined
    SonetApplicationCode_optics_sonet_undefined SonetApplicationCode = "optics-sonet-undefined"
)

// OpticsControllerState represents Optics controller state
type OpticsControllerState string

const (
    // Up
    OpticsControllerState_optics_state_up OpticsControllerState = "optics-state-up"

    // Down
    OpticsControllerState_optics_state_down OpticsControllerState = "optics-state-down"

    // Administratively Down
    OpticsControllerState_optics_state_admin_down OpticsControllerState = "optics-state-admin-down"
)

// OpticsAmplifierSafetyControlMode represents Optics amplifier safety control mode
type OpticsAmplifierSafetyControlMode string

const (
    // Invalid
    OpticsAmplifierSafetyControlMode_optics_amplifier_safety_mode_invalid OpticsAmplifierSafetyControlMode = "optics-amplifier-safety-mode-invalid"

    // auto
    OpticsAmplifierSafetyControlMode_optics_amplifier_safety_mode_auto OpticsAmplifierSafetyControlMode = "optics-amplifier-safety-mode-auto"

    // disabled
    OpticsAmplifierSafetyControlMode_optics_amplifier_safety_mode_disabled OpticsAmplifierSafetyControlMode = "optics-amplifier-safety-mode-disabled"
)

// OpticsLaserState represents Optics laser state
type OpticsLaserState string

const (
    // On
    OpticsLaserState_on OpticsLaserState = "on"

    // Off
    OpticsLaserState_off OpticsLaserState = "off"

    // Unknown
    OpticsLaserState_unknown OpticsLaserState = "unknown"

    // Apr
    OpticsLaserState_apr OpticsLaserState = "apr"

    // N/A
    OpticsLaserState_na OpticsLaserState = "na"
)

// OpticsFec represents Optics fec
type OpticsFec string

const (
    // FEC NONE
    OpticsFec_fec_none OpticsFec = "fec-none"

    // ENHANCED FEC H15
    OpticsFec_fec_hg15 OpticsFec = "fec-hg15"

    // ENHANCED FEC H25
    OpticsFec_fec_hg25 OpticsFec = "fec-hg25"

    // FEC HG15 DE
    OpticsFec_fec_hg15_de OpticsFec = "fec-hg15-de"

    // FEC HG25 DE
    OpticsFec_fec_hg25_de OpticsFec = "fec-hg25-de"

    // FEC ENABLED
    OpticsFec_fec_enabled OpticsFec = "fec-enabled"

    // FEC DISABLED
    OpticsFec_fec_not_set OpticsFec = "fec-not-set"

    // FEC CL91
    OpticsFec_fec_cl91 OpticsFec = "fec-cl91"
)

// OpticsPortStatus represents Optics port status
type OpticsPortStatus string

const (
    // Active
    OpticsPortStatus_active OpticsPortStatus = "active"

    // Standby
    OpticsPortStatus_standby OpticsPortStatus = "standby"
)

// OpticsPhy represents Optics phy type
type OpticsPhy string

const (
    // Not set
    OpticsPhy_not_set OpticsPhy = "not-set"

    // Invalid
    OpticsPhy_invalid OpticsPhy = "invalid"

    // Long reach 4 lanes
    OpticsPhy_long_reach_four_lanes OpticsPhy = "long-reach-four-lanes"

    // Short reach 10 lanes
    OpticsPhy_short_reach_ten_lanes OpticsPhy = "short-reach-ten-lanes"

    // Short reach 1 lane
    OpticsPhy_short_reach_one_lane OpticsPhy = "short-reach-one-lane"

    // Long reach 1 lane
    OpticsPhy_long_reach_one_lane OpticsPhy = "long-reach-one-lane"

    // Short reach 4 lanes
    OpticsPhy_short_reach_four_lanes OpticsPhy = "short-reach-four-lanes"

    // Copper 4 lanes
    OpticsPhy_copper_four_lanes OpticsPhy = "copper-four-lanes"

    // Active optical cable
    OpticsPhy_active_optical_cable OpticsPhy = "active-optical-cable"

    // Long reach 4 lanes
    OpticsPhy_fourty_gig_e_long_reach_four_lanes OpticsPhy = "fourty-gig-e-long-reach-four-lanes"

    // Short reach 4 lanes
    OpticsPhy_fourty_gig_e_short_reach_four_lanes OpticsPhy = "fourty-gig-e-short-reach-four-lanes"

    // CWDM 4 lanes
    OpticsPhy_cwdm_four_lanes OpticsPhy = "cwdm-four-lanes"

    // Extended reach 4 lanes
    OpticsPhy_extended_reach_four_lanes OpticsPhy = "extended-reach-four-lanes"

    // PSM 4 lanes
    OpticsPhy_psm_four_lanes OpticsPhy = "psm-four-lanes"

    // Active copper cable
    OpticsPhy_active_copper_cable OpticsPhy = "active-copper-cable"

    // Extended reach 4 lanes
    OpticsPhy_fourty_gig_e_extended_reach_four_lanes OpticsPhy = "fourty-gig-e-extended-reach-four-lanes"

    // Short reach 1 lane
    OpticsPhy_four_x_ten_gig_e_short_reach_one_lane OpticsPhy = "four-x-ten-gig-e-short-reach-one-lane"

    // PSM 4 lanes
    OpticsPhy_fourty_gig_epsm_four_lanes OpticsPhy = "fourty-gig-epsm-four-lanes"

    // Copper 4 lanes
    OpticsPhy_fourty_gig_e_copper_four_lanes OpticsPhy = "fourty-gig-e-copper-four-lanes"

    // Long reach MM 1 lane
    OpticsPhy_long_reach_mm_one_lane OpticsPhy = "long-reach-mm-one-lane"

    // Copper Short reach 4 lanes
    OpticsPhy_copper_short_reach OpticsPhy = "copper-short-reach"

    // Short reach BD 4 lanes
    OpticsPhy_short_reach_srbd OpticsPhy = "short-reach-srbd"

    // Copper One Lane
    OpticsPhy_copper_one_lane OpticsPhy = "copper-one-lane"

    // Long reach 1 lane
    OpticsPhy_four_x_ten_gig_e_long_reach_one_lane OpticsPhy = "four-x-ten-gig-e-long-reach-one-lane"

    // Active optical cable
    OpticsPhy_fourty_gig_eaoc_four_lanes OpticsPhy = "fourty-gig-eaoc-four-lanes"

    // Extended One Lane
    OpticsPhy_extended_one_lane OpticsPhy = "extended-one-lane"

    // ZR One Lane
    OpticsPhy_zr_one_lane OpticsPhy = "zr-one-lane"

    // DWDM One Lane
    OpticsPhy_dwdm_one_lane OpticsPhy = "dwdm-one-lane"

    // SX One Lane
    OpticsPhy_sx_one_lane OpticsPhy = "sx-one-lane"

    // LX One Lane
    OpticsPhy_lx_one_lane OpticsPhy = "lx-one-lane"

    // EX One Lane
    OpticsPhy_ex_one_lane OpticsPhy = "ex-one-lane"

    // ZX One Lane
    OpticsPhy_zx_one_lane OpticsPhy = "zx-one-lane"

    // BASE_T One Lane
    OpticsPhy_ba_set_one_lane OpticsPhy = "ba-set-one-lane"

    // Active Optical 1 Lane
    OpticsPhy_aoc_one_lane OpticsPhy = "aoc-one-lane"

    // Active Copper 1 Lane
    OpticsPhy_active_copper_one_lane OpticsPhy = "active-copper-one-lane"

    // Active Copper 4 Lanes
    OpticsPhy_fourty_gig_eacu_four_lanes OpticsPhy = "fourty-gig-eacu-four-lanes"

    // Active Copper 1 Lane
    OpticsPhy_four_x_ten_gig_eacu_one_lanes OpticsPhy = "four-x-ten-gig-eacu-one-lanes"

    // Copper 1 Lanes
    OpticsPhy_four_x_ten_gig_ecu_one_lanes OpticsPhy = "four-x-ten-gig-ecu-one-lanes"

    // Active Optics 1 Lane
    OpticsPhy_four_x_ten_gig_eaoc_one_lanes OpticsPhy = "four-x-ten-gig-eaoc-one-lanes"

    // Short reach 1 lane
    OpticsPhy_twenty_five_gig_short_reach_one_lane OpticsPhy = "twenty-five-gig-short-reach-one-lane"

    // Long reach 1 lane
    OpticsPhy_twenty_five_gig_long_reach_one_lane OpticsPhy = "twenty-five-gig-long-reach-one-lane"

    // Extended reach 1 lane
    OpticsPhy_twenty_five_gig_extended_reach_one_lane OpticsPhy = "twenty-five-gig-extended-reach-one-lane"

    // Copper One Lane
    OpticsPhy_twenty_five_gig_copper_one_lane OpticsPhy = "twenty-five-gig-copper-one-lane"

    // Active Optical One Lane
    OpticsPhy_twenty_five_gig_active_optical_one_lane OpticsPhy = "twenty-five-gig-active-optical-one-lane"

    // 100GE DWDM Optics
    OpticsPhy_hundred_gig_edwdm_two OpticsPhy = "hundred-gig-edwdm-two"

    // PLR4 Optics
    OpticsPhy_fourty_gig_plr4_four_lanes OpticsPhy = "fourty-gig-plr4-four-lanes"

    // ESR4 Optics
    OpticsPhy_fourty_gig_esr4_four_lanes OpticsPhy = "fourty-gig-esr4-four-lanes"

    // SMSR Optics
    OpticsPhy_smsr_four_lanes OpticsPhy = "smsr-four-lanes"

    // Cazadero R
    OpticsPhy_cazadero_rqsa OpticsPhy = "cazadero-rqsa"

    // CFP2 DWDM Optics
    OpticsPhy_trunk_port_cfp2 OpticsPhy = "trunk-port-cfp2"

    // Short reach 1 lane
    OpticsPhy_short_reach1_lane OpticsPhy = "short-reach1-lane"

    // Inmd reach 1 lane
    OpticsPhy_inmd_reach1lane OpticsPhy = "inmd-reach1lane"

    // Long reach 1 lane
    OpticsPhy_long_reach1_lane OpticsPhy = "long-reach1-lane"

    // Copper 1 Lanes
    OpticsPhy_twenty_five_gig_ecu_one_lanes OpticsPhy = "twenty-five-gig-ecu-one-lanes"

    // ExtentedReach4Lane
    OpticsPhy_hundred_gig_e OpticsPhy = "hundred-gig-e"

    // 10G BX optics
    OpticsPhy_ten_gig_bx OpticsPhy = "ten-gig-bx"

    // 1G BX optics
    OpticsPhy_one_geige OpticsPhy = "one-geige"

    // Long Reach 1 Lane
    OpticsPhy_ten_x_ten_gig_e_long_reach_one_lane OpticsPhy = "ten-x-ten-gig-e-long-reach-one-lane"

    // Extended Reach 1 Lane
    OpticsPhy_ten_x_ten_gig_e_extended_reach_one_lane OpticsPhy = "ten-x-ten-gig-e-extended-reach-one-lane"

    // Passive Copper One Lane
    OpticsPhy_passive_copper_one_lane OpticsPhy = "passive-copper-one-lane"

    // TenGigE CWDM One Lane
    OpticsPhy_ten_gig_ecwdm OpticsPhy = "ten-gig-ecwdm"

    // One GigE CWDM One Lane
    OpticsPhy_one_gig_ecwdm OpticsPhy = "one-gig-ecwdm"

    // One GigE DWDM One Lane
    OpticsPhy_one_gig_edwdm OpticsPhy = "one-gig-edwdm"

    // FX One Lane
    OpticsPhy_fx_one_lane OpticsPhy = "fx-one-lane"

    // TenGigE Multi Rate DWDM 1 Lane
    OpticsPhy_ten_gig_emrdwdm OpticsPhy = "ten-gig-emrdwdm"

    // TenGigE Edge Performance 1 Lane
    OpticsPhy_ten_gig_e_edge_performance OpticsPhy = "ten-gig-e-edge-performance"

    // OneGig CSFP optics
    OpticsPhy_one_gig_csfp OpticsPhy = "one-gig-csfp"
)

// OpticsTas represents Optics tas
type OpticsTas string

const (
    // Out Of Service
    OpticsTas_tas_ui_oos OpticsTas = "tas-ui-oos"

    // Maintenance
    OpticsTas_tas_ui_main OpticsTas = "tas-ui-main"

    // In Service
    OpticsTas_tas_ui_is OpticsTas = "tas-ui-is"

    // Automatic In Service
    OpticsTas_tas_ui_ains OpticsTas = "tas-ui-ains"
)

// Optics represents Optics
type Optics string

const (
    // Unknow Optics Type
    Optics_optics_unknown Optics = "optics-unknown"

    // Grey Optics
    Optics_optics_grey Optics = "optics-grey"

    // DWDM Optics
    Optics_optics_dwdm Optics = "optics-dwdm"

    // CWDM Optics
    Optics_optics_cwdm Optics = "optics-cwdm"
)

// OtnApplicationCode represents Otn application code
type OtnApplicationCode string

const (
    // Not Set
    OtnApplicationCode_optics_not_set OtnApplicationCode = "optics-not-set"

    // P1L1 2D1
    OtnApplicationCode_optics_p1l1_2d1 OtnApplicationCode = "optics-p1l1-2d1"

    // P1S1 2D2
    OtnApplicationCode_optics_p1s1_2d2 OtnApplicationCode = "optics-p1s1-2d2"

    // P1L1 2D2
    OtnApplicationCode_optics_p1l1_2d2 OtnApplicationCode = "optics-p1l1-2d2"

    // Undefined
    OtnApplicationCode_optics_undefined OtnApplicationCode = "optics-undefined"
)

// OpticsLedState represents Optics led state
type OpticsLedState string

const (
    // Off
    OpticsLedState_off OpticsLedState = "off"

    // Green
    OpticsLedState_green_on OpticsLedState = "green-on"

    // Green Flashing
    OpticsLedState_green_flashing OpticsLedState = "green-flashing"

    // Yellow
    OpticsLedState_yellow_on OpticsLedState = "yellow-on"

    // Yellow Flashing
    OpticsLedState_yellow_flashing OpticsLedState = "yellow-flashing"

    // Red
    OpticsLedState_red_on OpticsLedState = "red-on"

    // Red Flashing
    OpticsLedState_red_flashing OpticsLedState = "red-flashing"
)

// OpticsPort represents Optics port
type OpticsPort string

const (
    // Com
    OpticsPort_com OpticsPort = "com"

    // Line
    OpticsPort_line OpticsPort = "line"

    // Osc
    OpticsPort_osc OpticsPort = "osc"

    // Com Check
    OpticsPort_com_check OpticsPort = "com-check"

    // Working
    OpticsPort_work OpticsPort = "work"

    // Protected
    OpticsPort_prot OpticsPort = "prot"
)

// OpticsOper
// Optics operational data
type OpticsOper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All Optics Port operational data.
    OpticsPorts OpticsOper_OpticsPorts
}

func (opticsOper *OpticsOper) GetEntityData() *types.CommonEntityData {
    opticsOper.EntityData.YFilter = opticsOper.YFilter
    opticsOper.EntityData.YangName = "optics-oper"
    opticsOper.EntityData.BundleName = "cisco_ios_xr"
    opticsOper.EntityData.ParentYangName = "Cisco-IOS-XR-controller-optics-oper"
    opticsOper.EntityData.SegmentPath = "Cisco-IOS-XR-controller-optics-oper:optics-oper"
    opticsOper.EntityData.AbsolutePath = opticsOper.EntityData.SegmentPath
    opticsOper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsOper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsOper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsOper.EntityData.Children = types.NewOrderedMap()
    opticsOper.EntityData.Children.Append("optics-ports", types.YChild{"OpticsPorts", &opticsOper.OpticsPorts})
    opticsOper.EntityData.Leafs = types.NewOrderedMap()

    opticsOper.EntityData.YListKeys = []string {}

    return &(opticsOper.EntityData)
}

// OpticsOper_OpticsPorts
// All Optics Port operational data
type OpticsOper_OpticsPorts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optics operational data. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort.
    OpticsPort []*OpticsOper_OpticsPorts_OpticsPort
}

func (opticsPorts *OpticsOper_OpticsPorts) GetEntityData() *types.CommonEntityData {
    opticsPorts.EntityData.YFilter = opticsPorts.YFilter
    opticsPorts.EntityData.YangName = "optics-ports"
    opticsPorts.EntityData.BundleName = "cisco_ios_xr"
    opticsPorts.EntityData.ParentYangName = "optics-oper"
    opticsPorts.EntityData.SegmentPath = "optics-ports"
    opticsPorts.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/" + opticsPorts.EntityData.SegmentPath
    opticsPorts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsPorts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsPorts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsPorts.EntityData.Children = types.NewOrderedMap()
    opticsPorts.EntityData.Children.Append("optics-port", types.YChild{"OpticsPort", nil})
    for i := range opticsPorts.OpticsPort {
        opticsPorts.EntityData.Children.Append(types.GetSegmentPath(opticsPorts.OpticsPort[i]), types.YChild{"OpticsPort", opticsPorts.OpticsPort[i]})
    }
    opticsPorts.EntityData.Leafs = types.NewOrderedMap()

    opticsPorts.EntityData.YListKeys = []string {}

    return &(opticsPorts.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Name interface{}

    // Optics operational data.
    OpticsDwdmCarrierChannelMap OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap

    // Ots Spectrum Operational data.
    OtsSpectrumInfo OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo

    // Optics operational data.
    OpticsDwdmCarrierChannelMapFlexi OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi

    // Optics operational data.
    OpticsInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo

    // All Optics Port operational data.
    OpticsLanes OpticsOper_OpticsPorts_OpticsPort_OpticsLanes

    // Optics operational data.
    OpticsDbInfo OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo
}

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetEntityData() *types.CommonEntityData {
    opticsPort.EntityData.YFilter = opticsPort.YFilter
    opticsPort.EntityData.YangName = "optics-port"
    opticsPort.EntityData.BundleName = "cisco_ios_xr"
    opticsPort.EntityData.ParentYangName = "optics-ports"
    opticsPort.EntityData.SegmentPath = "optics-port" + types.AddKeyToken(opticsPort.Name, "name")
    opticsPort.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/" + opticsPort.EntityData.SegmentPath
    opticsPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsPort.EntityData.Children = types.NewOrderedMap()
    opticsPort.EntityData.Children.Append("optics-dwdm-carrier-channel-map", types.YChild{"OpticsDwdmCarrierChannelMap", &opticsPort.OpticsDwdmCarrierChannelMap})
    opticsPort.EntityData.Children.Append("ots-spectrum-info", types.YChild{"OtsSpectrumInfo", &opticsPort.OtsSpectrumInfo})
    opticsPort.EntityData.Children.Append("optics-dwdm-carrier-channel-map-flexi", types.YChild{"OpticsDwdmCarrierChannelMapFlexi", &opticsPort.OpticsDwdmCarrierChannelMapFlexi})
    opticsPort.EntityData.Children.Append("optics-info", types.YChild{"OpticsInfo", &opticsPort.OpticsInfo})
    opticsPort.EntityData.Children.Append("optics-lanes", types.YChild{"OpticsLanes", &opticsPort.OpticsLanes})
    opticsPort.EntityData.Children.Append("optics-db-info", types.YChild{"OpticsDbInfo", &opticsPort.OpticsDbInfo})
    opticsPort.EntityData.Leafs = types.NewOrderedMap()
    opticsPort.EntityData.Leafs.Append("name", types.YLeaf{"Name", opticsPort.Name})

    opticsPort.EntityData.YListKeys = []string {"Name"}

    return &(opticsPort.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DWDM carrier band. The type is OpticsWaveBand.
    DwdmCarrierBand interface{}

    // Lowest DWDM carrier supported. The type is interface{} with range:
    // 0..4294967295.
    DwdmCarrierMin interface{}

    // Highest DWDM carrier supported. The type is interface{} with range:
    // 0..4294967295.
    DwdmCarrierMax interface{}

    // DWDM carrier mapping info. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap_DwdmCarrierMapInfo.
    DwdmCarrierMapInfo []*OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap_DwdmCarrierMapInfo
}

func (opticsDwdmCarrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap) GetEntityData() *types.CommonEntityData {
    opticsDwdmCarrierChannelMap.EntityData.YFilter = opticsDwdmCarrierChannelMap.YFilter
    opticsDwdmCarrierChannelMap.EntityData.YangName = "optics-dwdm-carrier-channel-map"
    opticsDwdmCarrierChannelMap.EntityData.BundleName = "cisco_ios_xr"
    opticsDwdmCarrierChannelMap.EntityData.ParentYangName = "optics-port"
    opticsDwdmCarrierChannelMap.EntityData.SegmentPath = "optics-dwdm-carrier-channel-map"
    opticsDwdmCarrierChannelMap.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/" + opticsDwdmCarrierChannelMap.EntityData.SegmentPath
    opticsDwdmCarrierChannelMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsDwdmCarrierChannelMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsDwdmCarrierChannelMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsDwdmCarrierChannelMap.EntityData.Children = types.NewOrderedMap()
    opticsDwdmCarrierChannelMap.EntityData.Children.Append("dwdm-carrier-map-info", types.YChild{"DwdmCarrierMapInfo", nil})
    for i := range opticsDwdmCarrierChannelMap.DwdmCarrierMapInfo {
        types.SetYListKey(opticsDwdmCarrierChannelMap.DwdmCarrierMapInfo[i], i)
        opticsDwdmCarrierChannelMap.EntityData.Children.Append(types.GetSegmentPath(opticsDwdmCarrierChannelMap.DwdmCarrierMapInfo[i]), types.YChild{"DwdmCarrierMapInfo", opticsDwdmCarrierChannelMap.DwdmCarrierMapInfo[i]})
    }
    opticsDwdmCarrierChannelMap.EntityData.Leafs = types.NewOrderedMap()
    opticsDwdmCarrierChannelMap.EntityData.Leafs.Append("dwdm-carrier-band", types.YLeaf{"DwdmCarrierBand", opticsDwdmCarrierChannelMap.DwdmCarrierBand})
    opticsDwdmCarrierChannelMap.EntityData.Leafs.Append("dwdm-carrier-min", types.YLeaf{"DwdmCarrierMin", opticsDwdmCarrierChannelMap.DwdmCarrierMin})
    opticsDwdmCarrierChannelMap.EntityData.Leafs.Append("dwdm-carrier-max", types.YLeaf{"DwdmCarrierMax", opticsDwdmCarrierChannelMap.DwdmCarrierMax})

    opticsDwdmCarrierChannelMap.EntityData.YListKeys = []string {}

    return &(opticsDwdmCarrierChannelMap.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap_DwdmCarrierMapInfo
// DWDM carrier mapping info
type OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap_DwdmCarrierMapInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // ITU channel number. The type is interface{} with range: 0..4294967295.
    ItuChanNum interface{}

    // G694 channel number. The type is interface{} with range:
    // -2147483648..2147483647.
    G694ChanNum interface{}

    // Frequency. The type is string with length: 0..32.
    Frequency interface{}

    // Wavelength. The type is string with length: 0..32.
    Wavelength interface{}
}

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMap_DwdmCarrierMapInfo) GetEntityData() *types.CommonEntityData {
    dwdmCarrierMapInfo.EntityData.YFilter = dwdmCarrierMapInfo.YFilter
    dwdmCarrierMapInfo.EntityData.YangName = "dwdm-carrier-map-info"
    dwdmCarrierMapInfo.EntityData.BundleName = "cisco_ios_xr"
    dwdmCarrierMapInfo.EntityData.ParentYangName = "optics-dwdm-carrier-channel-map"
    dwdmCarrierMapInfo.EntityData.SegmentPath = "dwdm-carrier-map-info" + types.AddNoKeyToken(dwdmCarrierMapInfo)
    dwdmCarrierMapInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-dwdm-carrier-channel-map/" + dwdmCarrierMapInfo.EntityData.SegmentPath
    dwdmCarrierMapInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dwdmCarrierMapInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dwdmCarrierMapInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dwdmCarrierMapInfo.EntityData.Children = types.NewOrderedMap()
    dwdmCarrierMapInfo.EntityData.Leafs = types.NewOrderedMap()
    dwdmCarrierMapInfo.EntityData.Leafs.Append("itu-chan-num", types.YLeaf{"ItuChanNum", dwdmCarrierMapInfo.ItuChanNum})
    dwdmCarrierMapInfo.EntityData.Leafs.Append("g694-chan-num", types.YLeaf{"G694ChanNum", dwdmCarrierMapInfo.G694ChanNum})
    dwdmCarrierMapInfo.EntityData.Leafs.Append("frequency", types.YLeaf{"Frequency", dwdmCarrierMapInfo.Frequency})
    dwdmCarrierMapInfo.EntityData.Leafs.Append("wavelength", types.YLeaf{"Wavelength", dwdmCarrierMapInfo.Wavelength})

    dwdmCarrierMapInfo.EntityData.YListKeys = []string {}

    return &(dwdmCarrierMapInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo
// Ots Spectrum Operational data
type OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OTS Spectrum information.
    SpectrumInfo OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo
}

func (otsSpectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo) GetEntityData() *types.CommonEntityData {
    otsSpectrumInfo.EntityData.YFilter = otsSpectrumInfo.YFilter
    otsSpectrumInfo.EntityData.YangName = "ots-spectrum-info"
    otsSpectrumInfo.EntityData.BundleName = "cisco_ios_xr"
    otsSpectrumInfo.EntityData.ParentYangName = "optics-port"
    otsSpectrumInfo.EntityData.SegmentPath = "ots-spectrum-info"
    otsSpectrumInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/" + otsSpectrumInfo.EntityData.SegmentPath
    otsSpectrumInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otsSpectrumInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otsSpectrumInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otsSpectrumInfo.EntityData.Children = types.NewOrderedMap()
    otsSpectrumInfo.EntityData.Children.Append("spectrum-info", types.YChild{"SpectrumInfo", &otsSpectrumInfo.SpectrumInfo})
    otsSpectrumInfo.EntityData.Leafs = types.NewOrderedMap()

    otsSpectrumInfo.EntityData.YListKeys = []string {}

    return &(otsSpectrumInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo
// OTS Spectrum information
type OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of slices in Spectrum. The type is interface{} with range:
    // 0..4294967295.
    TotalSpectrumSliceCount interface{}

    // Spacing between spectrum slices. The type is interface{} with range:
    // 0..4294967295.
    SpectrumSliceSpacing interface{}

    // Wavelength of first slice. The type is string with length: 0..32.
    FirstSliceWavelength interface{}

    // Power information of spectrum slice info. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo_SpectrumSlicePowerInfo.
    SpectrumSlicePowerInfo []*OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo_SpectrumSlicePowerInfo
}

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo) GetEntityData() *types.CommonEntityData {
    spectrumInfo.EntityData.YFilter = spectrumInfo.YFilter
    spectrumInfo.EntityData.YangName = "spectrum-info"
    spectrumInfo.EntityData.BundleName = "cisco_ios_xr"
    spectrumInfo.EntityData.ParentYangName = "ots-spectrum-info"
    spectrumInfo.EntityData.SegmentPath = "spectrum-info"
    spectrumInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/ots-spectrum-info/" + spectrumInfo.EntityData.SegmentPath
    spectrumInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spectrumInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spectrumInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spectrumInfo.EntityData.Children = types.NewOrderedMap()
    spectrumInfo.EntityData.Children.Append("spectrum-slice-power-info", types.YChild{"SpectrumSlicePowerInfo", nil})
    for i := range spectrumInfo.SpectrumSlicePowerInfo {
        types.SetYListKey(spectrumInfo.SpectrumSlicePowerInfo[i], i)
        spectrumInfo.EntityData.Children.Append(types.GetSegmentPath(spectrumInfo.SpectrumSlicePowerInfo[i]), types.YChild{"SpectrumSlicePowerInfo", spectrumInfo.SpectrumSlicePowerInfo[i]})
    }
    spectrumInfo.EntityData.Leafs = types.NewOrderedMap()
    spectrumInfo.EntityData.Leafs.Append("total-spectrum-slice-count", types.YLeaf{"TotalSpectrumSliceCount", spectrumInfo.TotalSpectrumSliceCount})
    spectrumInfo.EntityData.Leafs.Append("spectrum-slice-spacing", types.YLeaf{"SpectrumSliceSpacing", spectrumInfo.SpectrumSliceSpacing})
    spectrumInfo.EntityData.Leafs.Append("first-slice-wavelength", types.YLeaf{"FirstSliceWavelength", spectrumInfo.FirstSliceWavelength})

    spectrumInfo.EntityData.YListKeys = []string {}

    return &(spectrumInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo_SpectrumSlicePowerInfo
// Power information of spectrum slice info
type OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo_SpectrumSlicePowerInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // spectrum slice number. The type is interface{} with range: 0..4294967295.
    SliceNum interface{}

    // Lower frequency of the specified PSD. The type is interface{} with range:
    // 0..18446744073709551615.
    LowerFrequency interface{}

    // Upper frequency of the specified PSD. The type is interface{} with range:
    // 0..18446744073709551615.
    UpperFrequency interface{}

    // Received Power in dBm multiplied by 100. The type is interface{} with
    // range: -2147483648..2147483647.
    RxPower interface{}

    // Transmit Power in dBm multiplied by 100. The type is interface{} with
    // range: -2147483648..2147483647.
    TxPower interface{}

    // Received Power spectral density in microwatts per megahertz, uW/MHz. The
    // type is string with length: 0..16.
    RxPsd interface{}

    // Transmit Power spectral density in microwatts per megahertz, uW/MHz. The
    // type is string with length: 0..16.
    TxPsd interface{}
}

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OtsSpectrumInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetEntityData() *types.CommonEntityData {
    spectrumSlicePowerInfo.EntityData.YFilter = spectrumSlicePowerInfo.YFilter
    spectrumSlicePowerInfo.EntityData.YangName = "spectrum-slice-power-info"
    spectrumSlicePowerInfo.EntityData.BundleName = "cisco_ios_xr"
    spectrumSlicePowerInfo.EntityData.ParentYangName = "spectrum-info"
    spectrumSlicePowerInfo.EntityData.SegmentPath = "spectrum-slice-power-info" + types.AddNoKeyToken(spectrumSlicePowerInfo)
    spectrumSlicePowerInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/ots-spectrum-info/spectrum-info/" + spectrumSlicePowerInfo.EntityData.SegmentPath
    spectrumSlicePowerInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spectrumSlicePowerInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spectrumSlicePowerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spectrumSlicePowerInfo.EntityData.Children = types.NewOrderedMap()
    spectrumSlicePowerInfo.EntityData.Leafs = types.NewOrderedMap()
    spectrumSlicePowerInfo.EntityData.Leafs.Append("slice-num", types.YLeaf{"SliceNum", spectrumSlicePowerInfo.SliceNum})
    spectrumSlicePowerInfo.EntityData.Leafs.Append("lower-frequency", types.YLeaf{"LowerFrequency", spectrumSlicePowerInfo.LowerFrequency})
    spectrumSlicePowerInfo.EntityData.Leafs.Append("upper-frequency", types.YLeaf{"UpperFrequency", spectrumSlicePowerInfo.UpperFrequency})
    spectrumSlicePowerInfo.EntityData.Leafs.Append("rx-power", types.YLeaf{"RxPower", spectrumSlicePowerInfo.RxPower})
    spectrumSlicePowerInfo.EntityData.Leafs.Append("tx-power", types.YLeaf{"TxPower", spectrumSlicePowerInfo.TxPower})
    spectrumSlicePowerInfo.EntityData.Leafs.Append("rx-psd", types.YLeaf{"RxPsd", spectrumSlicePowerInfo.RxPsd})
    spectrumSlicePowerInfo.EntityData.Leafs.Append("tx-psd", types.YLeaf{"TxPsd", spectrumSlicePowerInfo.TxPsd})

    spectrumSlicePowerInfo.EntityData.YListKeys = []string {}

    return &(spectrumSlicePowerInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DWDM carrier band. The type is OpticsWaveBand.
    DwdmCarrierBand interface{}

    // Lowest DWDM carrier supported. The type is interface{} with range:
    // 0..4294967295.
    DwdmCarrierMin interface{}

    // Highest DWDM carrier supported. The type is interface{} with range:
    // 0..4294967295.
    DwdmCarrierMax interface{}

    // DWDM carrier mapping info. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi_DwdmCarrierMapInfo.
    DwdmCarrierMapInfo []*OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi_DwdmCarrierMapInfo
}

func (opticsDwdmCarrierChannelMapFlexi *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi) GetEntityData() *types.CommonEntityData {
    opticsDwdmCarrierChannelMapFlexi.EntityData.YFilter = opticsDwdmCarrierChannelMapFlexi.YFilter
    opticsDwdmCarrierChannelMapFlexi.EntityData.YangName = "optics-dwdm-carrier-channel-map-flexi"
    opticsDwdmCarrierChannelMapFlexi.EntityData.BundleName = "cisco_ios_xr"
    opticsDwdmCarrierChannelMapFlexi.EntityData.ParentYangName = "optics-port"
    opticsDwdmCarrierChannelMapFlexi.EntityData.SegmentPath = "optics-dwdm-carrier-channel-map-flexi"
    opticsDwdmCarrierChannelMapFlexi.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/" + opticsDwdmCarrierChannelMapFlexi.EntityData.SegmentPath
    opticsDwdmCarrierChannelMapFlexi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsDwdmCarrierChannelMapFlexi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsDwdmCarrierChannelMapFlexi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsDwdmCarrierChannelMapFlexi.EntityData.Children = types.NewOrderedMap()
    opticsDwdmCarrierChannelMapFlexi.EntityData.Children.Append("dwdm-carrier-map-info", types.YChild{"DwdmCarrierMapInfo", nil})
    for i := range opticsDwdmCarrierChannelMapFlexi.DwdmCarrierMapInfo {
        types.SetYListKey(opticsDwdmCarrierChannelMapFlexi.DwdmCarrierMapInfo[i], i)
        opticsDwdmCarrierChannelMapFlexi.EntityData.Children.Append(types.GetSegmentPath(opticsDwdmCarrierChannelMapFlexi.DwdmCarrierMapInfo[i]), types.YChild{"DwdmCarrierMapInfo", opticsDwdmCarrierChannelMapFlexi.DwdmCarrierMapInfo[i]})
    }
    opticsDwdmCarrierChannelMapFlexi.EntityData.Leafs = types.NewOrderedMap()
    opticsDwdmCarrierChannelMapFlexi.EntityData.Leafs.Append("dwdm-carrier-band", types.YLeaf{"DwdmCarrierBand", opticsDwdmCarrierChannelMapFlexi.DwdmCarrierBand})
    opticsDwdmCarrierChannelMapFlexi.EntityData.Leafs.Append("dwdm-carrier-min", types.YLeaf{"DwdmCarrierMin", opticsDwdmCarrierChannelMapFlexi.DwdmCarrierMin})
    opticsDwdmCarrierChannelMapFlexi.EntityData.Leafs.Append("dwdm-carrier-max", types.YLeaf{"DwdmCarrierMax", opticsDwdmCarrierChannelMapFlexi.DwdmCarrierMax})

    opticsDwdmCarrierChannelMapFlexi.EntityData.YListKeys = []string {}

    return &(opticsDwdmCarrierChannelMapFlexi.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi_DwdmCarrierMapInfo
// DWDM carrier mapping info
type OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi_DwdmCarrierMapInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // ITU channel number. The type is interface{} with range: 0..4294967295.
    ItuChanNum interface{}

    // G694 channel number. The type is interface{} with range:
    // -2147483648..2147483647.
    G694ChanNum interface{}

    // Frequency. The type is string with length: 0..32.
    Frequency interface{}

    // Wavelength. The type is string with length: 0..32.
    Wavelength interface{}
}

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrierChannelMapFlexi_DwdmCarrierMapInfo) GetEntityData() *types.CommonEntityData {
    dwdmCarrierMapInfo.EntityData.YFilter = dwdmCarrierMapInfo.YFilter
    dwdmCarrierMapInfo.EntityData.YangName = "dwdm-carrier-map-info"
    dwdmCarrierMapInfo.EntityData.BundleName = "cisco_ios_xr"
    dwdmCarrierMapInfo.EntityData.ParentYangName = "optics-dwdm-carrier-channel-map-flexi"
    dwdmCarrierMapInfo.EntityData.SegmentPath = "dwdm-carrier-map-info" + types.AddNoKeyToken(dwdmCarrierMapInfo)
    dwdmCarrierMapInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-dwdm-carrier-channel-map-flexi/" + dwdmCarrierMapInfo.EntityData.SegmentPath
    dwdmCarrierMapInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dwdmCarrierMapInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dwdmCarrierMapInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dwdmCarrierMapInfo.EntityData.Children = types.NewOrderedMap()
    dwdmCarrierMapInfo.EntityData.Leafs = types.NewOrderedMap()
    dwdmCarrierMapInfo.EntityData.Leafs.Append("itu-chan-num", types.YLeaf{"ItuChanNum", dwdmCarrierMapInfo.ItuChanNum})
    dwdmCarrierMapInfo.EntityData.Leafs.Append("g694-chan-num", types.YLeaf{"G694ChanNum", dwdmCarrierMapInfo.G694ChanNum})
    dwdmCarrierMapInfo.EntityData.Leafs.Append("frequency", types.YLeaf{"Frequency", dwdmCarrierMapInfo.Frequency})
    dwdmCarrierMapInfo.EntityData.Leafs.Append("wavelength", types.YLeaf{"Wavelength", dwdmCarrierMapInfo.Wavelength})

    dwdmCarrierMapInfo.EntityData.YListKeys = []string {}

    return &(dwdmCarrierMapInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport Admin State. The type is OpticsTas.
    TransportAdminState interface{}

    // Is Optics Present?. The type is bool.
    OpticsPresent interface{}

    // Old Optics type name, Use Derived Optics type. The type is Optics.
    OpticsType interface{}

    // Derived Optics type name. The type is string.
    DerivedOpticsType interface{}

    // Optics module name. The type is string.
    OpticsModule interface{}

    // DWDM Carrier Band information. The type is OpticsWaveBand.
    DwdmCarrierBand interface{}

    // Current ITU DWDM Carrier channel number. The type is string.
    DwdmCarrierChannel interface{}

    // DWDM Carrier frequency read from hw in the units of  1THz. The type is
    // string.
    DwdmCarrierFrequency interface{}

    // Wavelength of color optics 0.001nm. The type is string.
    DwdmCarrierWavelength interface{}

    // Wavelength of grey optics 0.01nm. The type is interface{} with range:
    // 0..4294967295.
    GreyWavelength interface{}

    // Rx Low threshold value in units of 0.1dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    RxLowThreshold interface{}

    // Rx High threshold value in units of 0.1dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    RxHighThreshold interface{}

    // LBC High threshold value in units of percentage. The type is interface{}
    // with range: -2147483648..2147483647. Units are percentage.
    LbcHighThreshold interface{}

    // Tx Low threshold value in units of 0.1dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    TxLowThreshold interface{}

    // Tx High threshold value in units of 0.1dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    TxHighThreshold interface{}

    // LBC high threshold default value in units of 0 .001mA. The type is
    // interface{} with range: -2147483648..2147483647.
    LbcThHighDefault interface{}

    // LBC low threshold default value in units of 0 .001mA. The type is
    // interface{} with range: -2147483648..2147483647.
    LbcThLowDefault interface{}

    // Temp Low threshold value in the units 0.01C. The type is interface{} with
    // range: -2147483648..2147483647.
    TempLowThreshold interface{}

    // Temp High threshold value in the units of 0.01C. The type is interface{}
    // with range: -2147483648..2147483647.
    TempHighThreshold interface{}

    // Volt Low threshold value in the units of 0.01V. The type is interface{}
    // with range: -2147483648..2147483647.
    VoltLowThreshold interface{}

    // Volt High threshold value in the units of 0.01V. The type is interface{}
    // with range: -2147483648..2147483647.
    VoltHighThreshold interface{}

    // Chromatic Dispersion ps/nm. The type is interface{} with range:
    // -2147483648..2147483647.
    Cd interface{}

    // Chromatic Dispersion Min ps/nm. The type is interface{} with range:
    // -2147483648..2147483647.
    CdMin interface{}

    // Chromatic Dispersion Max ps/nm. The type is interface{} with range:
    // -2147483648..2147483647.
    CdMax interface{}

    // Chromatic Dispersion low threshold ps/nm. The type is interface{} with
    // range: -2147483648..2147483647.
    CdLowThreshold interface{}

    // Chromatic Dispersion high threshold ps/nm. The type is interface{} with
    // range: -2147483648..2147483647.
    CdHighThreshold interface{}

    // OSNR low threshold in 0.01 dB. The type is string.
    OsnrLowThreshold interface{}

    // DGD high threshold in 0.1 ps. The type is string.
    DgdHighThreshold interface{}

    // Polarization Mode Dispersion 0.1ps. The type is string.
    PolarizationModeDispersion interface{}

    // Second Order Polarization Mode Dispersion 0 .1ps^2. The type is string.
    SecondOrderPolarizationModeDispersion interface{}

    // Optical Signal to Noise Ratio dB. The type is string.
    OpticalSignalToNoiseRatio interface{}

    // Polarization Dependent Loss dB. The type is string.
    PolarizationDependentLoss interface{}

    // Polarization Change Rate rad/s. The type is string.
    PolarizationChangeRate interface{}

    // Differential Group Delay ps. The type is string.
    DifferentialGroupDelay interface{}

    // Phase Noise dB. The type is string.
    PhaseNoise interface{}

    // PmEable or Disable. The type is interface{} with range: 0..4294967295.
    PmEnable interface{}

    // Showing laser state.Either ON or OFF or unknown. The type is
    // OpticsLaserState.
    LaserState interface{}

    // Available Modulation Types. The type is OpticsModulation.
    ModulationType interface{}

    // Showing Current Colour of led state. The type is OpticsLedState.
    LedState interface{}

    // Optics controller state: Up, Down or Administratively Down. The type is
    // OpticsControllerState.
    ControllerState interface{}

    // Optics form factor. The type is OpticsFormFactor.
    FormFactor interface{}

    // Optics physical type. The type is OpticsPhy.
    PhyType interface{}

    // Configured Tx power value in units of 0.1dB. The type is interface{} with
    // range: -2147483648..2147483647.
    CfgTxPower interface{}

    // TX Power Configuration is supported or not. The type is bool.
    CfgTxPowerConfigurable interface{}

    // Temperature value in units of 0.01 C. The type is interface{} with range:
    // -2147483648..2147483647.
    Temperature interface{}

    // Voltage value in units of 0.01V. The type is interface{} with range:
    // -2147483648..2147483647.
    Voltage interface{}

    // Display Volt/Temp ?. The type is bool.
    DisplayVoltTemp interface{}

    // CD Configurable is supported or not. The type is bool.
    CdConfigurable interface{}

    // Optics FEC. The type is OpticsFec.
    OpticsFec interface{}

    // PM enabled or not. The type is interface{} with range:
    // -2147483648..2147483647.
    SkipSnmpPmTable interface{}

    // Showing port type. The type is OpticsPort.
    PortType interface{}

    // Showing port status. The type is OpticsPortStatus.
    PortStatus interface{}

    // Rx Voa Attenuation in the units of 0.01dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    RxVoaAttenuation interface{}

    // Tx Voa Attenuation in the units of 0.01dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    TxVoaAttenuation interface{}

    // Ampli Gain in the units of 0.01dBm. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliGain interface{}

    // Ampli Tilt in the units of 0.01dBm. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliTilt interface{}

    // rx power th configurable. The type is bool.
    RxPowerThConfigurable interface{}

    // tx power th configurable. The type is bool.
    TxPowerThConfigurable interface{}

    // RX VOA attenuation configured value in units of 0.1dB. The type is
    // interface{} with range: -2147483648..2147483647.
    RxVoaAttenuationConfigVal interface{}

    // TX VOA attenuation configured value in units of 0.1dB. The type is
    // interface{} with range: -2147483648..2147483647.
    TxVoaAttenuationConfigVal interface{}

    // ampli control mode config val. The type is OpticsAmplifierControlMode.
    AmpliControlModeConfigVal interface{}

    // ampli gain range config val. The type is OpticsAmplifierGainRange.
    AmpliGainRangeConfigVal interface{}

    // Ampli gain configured value in units of 0.1dB. The type is interface{} with
    // range: -2147483648..2147483647.
    AmpliGainConfigVal interface{}

    // Ampli tilt configured value in units of 0.1dB. The type is interface{} with
    // range: -2147483648..2147483647.
    AmpliTiltConfigVal interface{}

    // Ampli channel power configured value in units of 0.1dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    AmpliChannelPowerConfigVal interface{}

    // Channel power max delta configured value in units of 0.1 dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    ChannelPowerMaxDeltaConfigVal interface{}

    // Ampli gain low degrade threshold configured value in units of 0.1dBm. The
    // type is interface{} with range: -2147483648..2147483647.
    AmpliGainThrDegLowConfigVal interface{}

    // Ampli gain high degrade threshold configured value in units of 0.1dBm. The
    // type is interface{} with range: -2147483648..2147483647.
    AmpliGainThrDegHighConfigVal interface{}

    // osri config val. The type is bool.
    OsriConfigVal interface{}

    // tx config val. The type is bool.
    TxConfigVal interface{}

    // rx config val. The type is bool.
    RxConfigVal interface{}

    // safety control mode config val. The type is
    // OpticsAmplifierSafetyControlMode.
    SafetyControlModeConfigVal interface{}

    // Total Receive Power for Multi-Lane Optics in units of 0.01 dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    TotalRxPower interface{}

    // Total Transmit Power for Multi-Lane Optics in units of 0.01 dBm. The type
    // is interface{} with range: -2147483648..2147483647.
    TotalTxPower interface{}

    // Is BO configured ?. The type is bool.
    IsBoConfigured interface{}

    // Are the Extended Parameters Valid ?. The type is bool.
    IsExtParamValid interface{}

    // Are there any alarms ?. The type is bool.
    AlarmDetected interface{}

    // Rx Low Warning threshold value in units of 0 .1dBm. The type is interface{}
    // with range: -2147483648..2147483647.
    RxLowWarningThreshold interface{}

    // Rx High Warning threshold value in units of 0 .1dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    RxHighWarningThreshold interface{}

    // Tx Low Warning threshold value in units of 0 .1dBm. The type is interface{}
    // with range: -2147483648..2147483647.
    TxLowWarningThreshold interface{}

    // Tx High Warning threshold value in units of 0 .1dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    TxHighWarningThreshold interface{}

    // LBC high Warning threshold default value in units of 0.001mA. The type is
    // interface{} with range: -2147483648..2147483647.
    LbcThHighWarningDefault interface{}

    // LBC low warning threshold default value in units of 0.001mA. The type is
    // interface{} with range: -2147483648..2147483647.
    LbcThLowWarningDefault interface{}

    // Temp Low warning threshold value in the units 0 .01C. The type is
    // interface{} with range: -2147483648..2147483647.
    TempLowWarningThreshold interface{}

    // Temp High warning threshold value in the units of 0.01C. The type is
    // interface{} with range: -2147483648..2147483647.
    TempHighWarningThreshold interface{}

    // Volt Low warning threshold value in the units of 0.01V. The type is
    // interface{} with range: -2147483648..2147483647.
    VoltLowWarningThreshold interface{}

    // Volt High warning threshold value in the units of 0.01V. The type is
    // interface{} with range: -2147483648..2147483647.
    VoltHighWarningThreshold interface{}

    // Controller description string. The type is string.
    Description interface{}

    // Ampli gain range. The type is OpticsAmplifierGainRange.
    AmpliGainRange interface{}

    // Safety control mode. The type is OpticsAmplifierSafetyControlMode.
    SafetyControlMode interface{}

    // OSRI. The type is bool.
    Osri interface{}

    // TX Enable. The type is bool.
    TxEnable interface{}

    // RX Enable. The type is bool.
    RxEnable interface{}

    // Is the Optics type string valid ?. The type is bool.
    IsOpticsTypeStringValid interface{}

    // optics type String. The type is string.
    OpticsTypeStr interface{}

    // Rx Low threshold actual value in units of 0.1dBm. The type is interface{}
    // with range: -2147483648..2147483647.
    RxLowThresholdCurrent interface{}

    // RX Span Loss in units of 0.01 dB. The type is interface{} with range:
    // -2147483648..2147483647.
    RxSpanLoss interface{}

    // TX Span Loss in units of 0.01 dB. The type is interface{} with range:
    // -2147483648..2147483647.
    TxSpanLoss interface{}

    // Baud Rate in GBd. The type is string.
    BaudRate interface{}

    // Bits per Symbol. The type is string.
    BitsPerSymbol interface{}

    // Network SRLG information.
    NetworkSrlgInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo

    // Optics Alarm Information.
    OpticsAlarmInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo

    // Ots Alarm Information.
    OtsAlarmInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo

    // Transceiver Vendor Details.
    TransceiverInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo

    // Extended optics parameters.
    ExtParamVal OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal

    // Extended optics parameters threshold values.
    ExtParamThresholdVal OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal

    // Extended DOM alarm Information.
    ExtendedAlarmAlarmInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo

    // AINS information.
    AinsInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_AinsInfo

    // Lane information. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData.
    LaneData []*OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData
}

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetEntityData() *types.CommonEntityData {
    opticsInfo.EntityData.YFilter = opticsInfo.YFilter
    opticsInfo.EntityData.YangName = "optics-info"
    opticsInfo.EntityData.BundleName = "cisco_ios_xr"
    opticsInfo.EntityData.ParentYangName = "optics-port"
    opticsInfo.EntityData.SegmentPath = "optics-info"
    opticsInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/" + opticsInfo.EntityData.SegmentPath
    opticsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsInfo.EntityData.Children = types.NewOrderedMap()
    opticsInfo.EntityData.Children.Append("network-srlg-info", types.YChild{"NetworkSrlgInfo", &opticsInfo.NetworkSrlgInfo})
    opticsInfo.EntityData.Children.Append("optics-alarm-info", types.YChild{"OpticsAlarmInfo", &opticsInfo.OpticsAlarmInfo})
    opticsInfo.EntityData.Children.Append("ots-alarm-info", types.YChild{"OtsAlarmInfo", &opticsInfo.OtsAlarmInfo})
    opticsInfo.EntityData.Children.Append("transceiver-info", types.YChild{"TransceiverInfo", &opticsInfo.TransceiverInfo})
    opticsInfo.EntityData.Children.Append("ext-param-val", types.YChild{"ExtParamVal", &opticsInfo.ExtParamVal})
    opticsInfo.EntityData.Children.Append("ext-param-threshold-val", types.YChild{"ExtParamThresholdVal", &opticsInfo.ExtParamThresholdVal})
    opticsInfo.EntityData.Children.Append("extended-alarm-alarm-info", types.YChild{"ExtendedAlarmAlarmInfo", &opticsInfo.ExtendedAlarmAlarmInfo})
    opticsInfo.EntityData.Children.Append("ains-info", types.YChild{"AinsInfo", &opticsInfo.AinsInfo})
    opticsInfo.EntityData.Children.Append("lane-data", types.YChild{"LaneData", nil})
    for i := range opticsInfo.LaneData {
        types.SetYListKey(opticsInfo.LaneData[i], i)
        opticsInfo.EntityData.Children.Append(types.GetSegmentPath(opticsInfo.LaneData[i]), types.YChild{"LaneData", opticsInfo.LaneData[i]})
    }
    opticsInfo.EntityData.Leafs = types.NewOrderedMap()
    opticsInfo.EntityData.Leafs.Append("transport-admin-state", types.YLeaf{"TransportAdminState", opticsInfo.TransportAdminState})
    opticsInfo.EntityData.Leafs.Append("optics-present", types.YLeaf{"OpticsPresent", opticsInfo.OpticsPresent})
    opticsInfo.EntityData.Leafs.Append("optics-type", types.YLeaf{"OpticsType", opticsInfo.OpticsType})
    opticsInfo.EntityData.Leafs.Append("derived-optics-type", types.YLeaf{"DerivedOpticsType", opticsInfo.DerivedOpticsType})
    opticsInfo.EntityData.Leafs.Append("optics-module", types.YLeaf{"OpticsModule", opticsInfo.OpticsModule})
    opticsInfo.EntityData.Leafs.Append("dwdm-carrier-band", types.YLeaf{"DwdmCarrierBand", opticsInfo.DwdmCarrierBand})
    opticsInfo.EntityData.Leafs.Append("dwdm-carrier-channel", types.YLeaf{"DwdmCarrierChannel", opticsInfo.DwdmCarrierChannel})
    opticsInfo.EntityData.Leafs.Append("dwdm-carrier-frequency", types.YLeaf{"DwdmCarrierFrequency", opticsInfo.DwdmCarrierFrequency})
    opticsInfo.EntityData.Leafs.Append("dwdm-carrier-wavelength", types.YLeaf{"DwdmCarrierWavelength", opticsInfo.DwdmCarrierWavelength})
    opticsInfo.EntityData.Leafs.Append("grey-wavelength", types.YLeaf{"GreyWavelength", opticsInfo.GreyWavelength})
    opticsInfo.EntityData.Leafs.Append("rx-low-threshold", types.YLeaf{"RxLowThreshold", opticsInfo.RxLowThreshold})
    opticsInfo.EntityData.Leafs.Append("rx-high-threshold", types.YLeaf{"RxHighThreshold", opticsInfo.RxHighThreshold})
    opticsInfo.EntityData.Leafs.Append("lbc-high-threshold", types.YLeaf{"LbcHighThreshold", opticsInfo.LbcHighThreshold})
    opticsInfo.EntityData.Leafs.Append("tx-low-threshold", types.YLeaf{"TxLowThreshold", opticsInfo.TxLowThreshold})
    opticsInfo.EntityData.Leafs.Append("tx-high-threshold", types.YLeaf{"TxHighThreshold", opticsInfo.TxHighThreshold})
    opticsInfo.EntityData.Leafs.Append("lbc-th-high-default", types.YLeaf{"LbcThHighDefault", opticsInfo.LbcThHighDefault})
    opticsInfo.EntityData.Leafs.Append("lbc-th-low-default", types.YLeaf{"LbcThLowDefault", opticsInfo.LbcThLowDefault})
    opticsInfo.EntityData.Leafs.Append("temp-low-threshold", types.YLeaf{"TempLowThreshold", opticsInfo.TempLowThreshold})
    opticsInfo.EntityData.Leafs.Append("temp-high-threshold", types.YLeaf{"TempHighThreshold", opticsInfo.TempHighThreshold})
    opticsInfo.EntityData.Leafs.Append("volt-low-threshold", types.YLeaf{"VoltLowThreshold", opticsInfo.VoltLowThreshold})
    opticsInfo.EntityData.Leafs.Append("volt-high-threshold", types.YLeaf{"VoltHighThreshold", opticsInfo.VoltHighThreshold})
    opticsInfo.EntityData.Leafs.Append("cd", types.YLeaf{"Cd", opticsInfo.Cd})
    opticsInfo.EntityData.Leafs.Append("cd-min", types.YLeaf{"CdMin", opticsInfo.CdMin})
    opticsInfo.EntityData.Leafs.Append("cd-max", types.YLeaf{"CdMax", opticsInfo.CdMax})
    opticsInfo.EntityData.Leafs.Append("cd-low-threshold", types.YLeaf{"CdLowThreshold", opticsInfo.CdLowThreshold})
    opticsInfo.EntityData.Leafs.Append("cd-high-threshold", types.YLeaf{"CdHighThreshold", opticsInfo.CdHighThreshold})
    opticsInfo.EntityData.Leafs.Append("osnr-low-threshold", types.YLeaf{"OsnrLowThreshold", opticsInfo.OsnrLowThreshold})
    opticsInfo.EntityData.Leafs.Append("dgd-high-threshold", types.YLeaf{"DgdHighThreshold", opticsInfo.DgdHighThreshold})
    opticsInfo.EntityData.Leafs.Append("polarization-mode-dispersion", types.YLeaf{"PolarizationModeDispersion", opticsInfo.PolarizationModeDispersion})
    opticsInfo.EntityData.Leafs.Append("second-order-polarization-mode-dispersion", types.YLeaf{"SecondOrderPolarizationModeDispersion", opticsInfo.SecondOrderPolarizationModeDispersion})
    opticsInfo.EntityData.Leafs.Append("optical-signal-to-noise-ratio", types.YLeaf{"OpticalSignalToNoiseRatio", opticsInfo.OpticalSignalToNoiseRatio})
    opticsInfo.EntityData.Leafs.Append("polarization-dependent-loss", types.YLeaf{"PolarizationDependentLoss", opticsInfo.PolarizationDependentLoss})
    opticsInfo.EntityData.Leafs.Append("polarization-change-rate", types.YLeaf{"PolarizationChangeRate", opticsInfo.PolarizationChangeRate})
    opticsInfo.EntityData.Leafs.Append("differential-group-delay", types.YLeaf{"DifferentialGroupDelay", opticsInfo.DifferentialGroupDelay})
    opticsInfo.EntityData.Leafs.Append("phase-noise", types.YLeaf{"PhaseNoise", opticsInfo.PhaseNoise})
    opticsInfo.EntityData.Leafs.Append("pm-enable", types.YLeaf{"PmEnable", opticsInfo.PmEnable})
    opticsInfo.EntityData.Leafs.Append("laser-state", types.YLeaf{"LaserState", opticsInfo.LaserState})
    opticsInfo.EntityData.Leafs.Append("modulation-type", types.YLeaf{"ModulationType", opticsInfo.ModulationType})
    opticsInfo.EntityData.Leafs.Append("led-state", types.YLeaf{"LedState", opticsInfo.LedState})
    opticsInfo.EntityData.Leafs.Append("controller-state", types.YLeaf{"ControllerState", opticsInfo.ControllerState})
    opticsInfo.EntityData.Leafs.Append("form-factor", types.YLeaf{"FormFactor", opticsInfo.FormFactor})
    opticsInfo.EntityData.Leafs.Append("phy-type", types.YLeaf{"PhyType", opticsInfo.PhyType})
    opticsInfo.EntityData.Leafs.Append("cfg-tx-power", types.YLeaf{"CfgTxPower", opticsInfo.CfgTxPower})
    opticsInfo.EntityData.Leafs.Append("cfg-tx-power-configurable", types.YLeaf{"CfgTxPowerConfigurable", opticsInfo.CfgTxPowerConfigurable})
    opticsInfo.EntityData.Leafs.Append("temperature", types.YLeaf{"Temperature", opticsInfo.Temperature})
    opticsInfo.EntityData.Leafs.Append("voltage", types.YLeaf{"Voltage", opticsInfo.Voltage})
    opticsInfo.EntityData.Leafs.Append("display-volt-temp", types.YLeaf{"DisplayVoltTemp", opticsInfo.DisplayVoltTemp})
    opticsInfo.EntityData.Leafs.Append("cd-configurable", types.YLeaf{"CdConfigurable", opticsInfo.CdConfigurable})
    opticsInfo.EntityData.Leafs.Append("optics-fec", types.YLeaf{"OpticsFec", opticsInfo.OpticsFec})
    opticsInfo.EntityData.Leafs.Append("skip-snmp-pm-table", types.YLeaf{"SkipSnmpPmTable", opticsInfo.SkipSnmpPmTable})
    opticsInfo.EntityData.Leafs.Append("port-type", types.YLeaf{"PortType", opticsInfo.PortType})
    opticsInfo.EntityData.Leafs.Append("port-status", types.YLeaf{"PortStatus", opticsInfo.PortStatus})
    opticsInfo.EntityData.Leafs.Append("rx-voa-attenuation", types.YLeaf{"RxVoaAttenuation", opticsInfo.RxVoaAttenuation})
    opticsInfo.EntityData.Leafs.Append("tx-voa-attenuation", types.YLeaf{"TxVoaAttenuation", opticsInfo.TxVoaAttenuation})
    opticsInfo.EntityData.Leafs.Append("ampli-gain", types.YLeaf{"AmpliGain", opticsInfo.AmpliGain})
    opticsInfo.EntityData.Leafs.Append("ampli-tilt", types.YLeaf{"AmpliTilt", opticsInfo.AmpliTilt})
    opticsInfo.EntityData.Leafs.Append("rx-power-th-configurable", types.YLeaf{"RxPowerThConfigurable", opticsInfo.RxPowerThConfigurable})
    opticsInfo.EntityData.Leafs.Append("tx-power-th-configurable", types.YLeaf{"TxPowerThConfigurable", opticsInfo.TxPowerThConfigurable})
    opticsInfo.EntityData.Leafs.Append("rx-voa-attenuation-config-val", types.YLeaf{"RxVoaAttenuationConfigVal", opticsInfo.RxVoaAttenuationConfigVal})
    opticsInfo.EntityData.Leafs.Append("tx-voa-attenuation-config-val", types.YLeaf{"TxVoaAttenuationConfigVal", opticsInfo.TxVoaAttenuationConfigVal})
    opticsInfo.EntityData.Leafs.Append("ampli-control-mode-config-val", types.YLeaf{"AmpliControlModeConfigVal", opticsInfo.AmpliControlModeConfigVal})
    opticsInfo.EntityData.Leafs.Append("ampli-gain-range-config-val", types.YLeaf{"AmpliGainRangeConfigVal", opticsInfo.AmpliGainRangeConfigVal})
    opticsInfo.EntityData.Leafs.Append("ampli-gain-config-val", types.YLeaf{"AmpliGainConfigVal", opticsInfo.AmpliGainConfigVal})
    opticsInfo.EntityData.Leafs.Append("ampli-tilt-config-val", types.YLeaf{"AmpliTiltConfigVal", opticsInfo.AmpliTiltConfigVal})
    opticsInfo.EntityData.Leafs.Append("ampli-channel-power-config-val", types.YLeaf{"AmpliChannelPowerConfigVal", opticsInfo.AmpliChannelPowerConfigVal})
    opticsInfo.EntityData.Leafs.Append("channel-power-max-delta-config-val", types.YLeaf{"ChannelPowerMaxDeltaConfigVal", opticsInfo.ChannelPowerMaxDeltaConfigVal})
    opticsInfo.EntityData.Leafs.Append("ampli-gain-thr-deg-low-config-val", types.YLeaf{"AmpliGainThrDegLowConfigVal", opticsInfo.AmpliGainThrDegLowConfigVal})
    opticsInfo.EntityData.Leafs.Append("ampli-gain-thr-deg-high-config-val", types.YLeaf{"AmpliGainThrDegHighConfigVal", opticsInfo.AmpliGainThrDegHighConfigVal})
    opticsInfo.EntityData.Leafs.Append("osri-config-val", types.YLeaf{"OsriConfigVal", opticsInfo.OsriConfigVal})
    opticsInfo.EntityData.Leafs.Append("tx-config-val", types.YLeaf{"TxConfigVal", opticsInfo.TxConfigVal})
    opticsInfo.EntityData.Leafs.Append("rx-config-val", types.YLeaf{"RxConfigVal", opticsInfo.RxConfigVal})
    opticsInfo.EntityData.Leafs.Append("safety-control-mode-config-val", types.YLeaf{"SafetyControlModeConfigVal", opticsInfo.SafetyControlModeConfigVal})
    opticsInfo.EntityData.Leafs.Append("total-rx-power", types.YLeaf{"TotalRxPower", opticsInfo.TotalRxPower})
    opticsInfo.EntityData.Leafs.Append("total-tx-power", types.YLeaf{"TotalTxPower", opticsInfo.TotalTxPower})
    opticsInfo.EntityData.Leafs.Append("is-bo-configured", types.YLeaf{"IsBoConfigured", opticsInfo.IsBoConfigured})
    opticsInfo.EntityData.Leafs.Append("is-ext-param-valid", types.YLeaf{"IsExtParamValid", opticsInfo.IsExtParamValid})
    opticsInfo.EntityData.Leafs.Append("alarm-detected", types.YLeaf{"AlarmDetected", opticsInfo.AlarmDetected})
    opticsInfo.EntityData.Leafs.Append("rx-low-warning-threshold", types.YLeaf{"RxLowWarningThreshold", opticsInfo.RxLowWarningThreshold})
    opticsInfo.EntityData.Leafs.Append("rx-high-warning-threshold", types.YLeaf{"RxHighWarningThreshold", opticsInfo.RxHighWarningThreshold})
    opticsInfo.EntityData.Leafs.Append("tx-low-warning-threshold", types.YLeaf{"TxLowWarningThreshold", opticsInfo.TxLowWarningThreshold})
    opticsInfo.EntityData.Leafs.Append("tx-high-warning-threshold", types.YLeaf{"TxHighWarningThreshold", opticsInfo.TxHighWarningThreshold})
    opticsInfo.EntityData.Leafs.Append("lbc-th-high-warning-default", types.YLeaf{"LbcThHighWarningDefault", opticsInfo.LbcThHighWarningDefault})
    opticsInfo.EntityData.Leafs.Append("lbc-th-low-warning-default", types.YLeaf{"LbcThLowWarningDefault", opticsInfo.LbcThLowWarningDefault})
    opticsInfo.EntityData.Leafs.Append("temp-low-warning-threshold", types.YLeaf{"TempLowWarningThreshold", opticsInfo.TempLowWarningThreshold})
    opticsInfo.EntityData.Leafs.Append("temp-high-warning-threshold", types.YLeaf{"TempHighWarningThreshold", opticsInfo.TempHighWarningThreshold})
    opticsInfo.EntityData.Leafs.Append("volt-low-warning-threshold", types.YLeaf{"VoltLowWarningThreshold", opticsInfo.VoltLowWarningThreshold})
    opticsInfo.EntityData.Leafs.Append("volt-high-warning-threshold", types.YLeaf{"VoltHighWarningThreshold", opticsInfo.VoltHighWarningThreshold})
    opticsInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", opticsInfo.Description})
    opticsInfo.EntityData.Leafs.Append("ampli-gain-range", types.YLeaf{"AmpliGainRange", opticsInfo.AmpliGainRange})
    opticsInfo.EntityData.Leafs.Append("safety-control-mode", types.YLeaf{"SafetyControlMode", opticsInfo.SafetyControlMode})
    opticsInfo.EntityData.Leafs.Append("osri", types.YLeaf{"Osri", opticsInfo.Osri})
    opticsInfo.EntityData.Leafs.Append("tx-enable", types.YLeaf{"TxEnable", opticsInfo.TxEnable})
    opticsInfo.EntityData.Leafs.Append("rx-enable", types.YLeaf{"RxEnable", opticsInfo.RxEnable})
    opticsInfo.EntityData.Leafs.Append("is-optics-type-string-valid", types.YLeaf{"IsOpticsTypeStringValid", opticsInfo.IsOpticsTypeStringValid})
    opticsInfo.EntityData.Leafs.Append("optics-type-str", types.YLeaf{"OpticsTypeStr", opticsInfo.OpticsTypeStr})
    opticsInfo.EntityData.Leafs.Append("rx-low-threshold-current", types.YLeaf{"RxLowThresholdCurrent", opticsInfo.RxLowThresholdCurrent})
    opticsInfo.EntityData.Leafs.Append("rx-span-loss", types.YLeaf{"RxSpanLoss", opticsInfo.RxSpanLoss})
    opticsInfo.EntityData.Leafs.Append("tx-span-loss", types.YLeaf{"TxSpanLoss", opticsInfo.TxSpanLoss})
    opticsInfo.EntityData.Leafs.Append("baud-rate", types.YLeaf{"BaudRate", opticsInfo.BaudRate})
    opticsInfo.EntityData.Leafs.Append("bits-per-symbol", types.YLeaf{"BitsPerSymbol", opticsInfo.BitsPerSymbol})

    opticsInfo.EntityData.YListKeys = []string {}

    return &(opticsInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo
// Network SRLG information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network Srlg Array. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray.
    NetworkSrlgArray []*OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetEntityData() *types.CommonEntityData {
    networkSrlgInfo.EntityData.YFilter = networkSrlgInfo.YFilter
    networkSrlgInfo.EntityData.YangName = "network-srlg-info"
    networkSrlgInfo.EntityData.BundleName = "cisco_ios_xr"
    networkSrlgInfo.EntityData.ParentYangName = "optics-info"
    networkSrlgInfo.EntityData.SegmentPath = "network-srlg-info"
    networkSrlgInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + networkSrlgInfo.EntityData.SegmentPath
    networkSrlgInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkSrlgInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkSrlgInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkSrlgInfo.EntityData.Children = types.NewOrderedMap()
    networkSrlgInfo.EntityData.Children.Append("network-srlg-array", types.YChild{"NetworkSrlgArray", nil})
    for i := range networkSrlgInfo.NetworkSrlgArray {
        types.SetYListKey(networkSrlgInfo.NetworkSrlgArray[i], i)
        networkSrlgInfo.EntityData.Children.Append(types.GetSegmentPath(networkSrlgInfo.NetworkSrlgArray[i]), types.YChild{"NetworkSrlgArray", networkSrlgInfo.NetworkSrlgArray[i]})
    }
    networkSrlgInfo.EntityData.Leafs = types.NewOrderedMap()

    networkSrlgInfo.EntityData.YListKeys = []string {}

    return &(networkSrlgInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray
// Network Srlg Array
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Array to maintain set number. The type is interface{} with range:
    // 0..4294967295.
    SetNumber interface{}

    // Network Srlg. The type is slice of interface{} with range: 0..4294967295.
    NetworkSrlg []interface{}
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetEntityData() *types.CommonEntityData {
    networkSrlgArray.EntityData.YFilter = networkSrlgArray.YFilter
    networkSrlgArray.EntityData.YangName = "network-srlg-array"
    networkSrlgArray.EntityData.BundleName = "cisco_ios_xr"
    networkSrlgArray.EntityData.ParentYangName = "network-srlg-info"
    networkSrlgArray.EntityData.SegmentPath = "network-srlg-array" + types.AddNoKeyToken(networkSrlgArray)
    networkSrlgArray.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/network-srlg-info/" + networkSrlgArray.EntityData.SegmentPath
    networkSrlgArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkSrlgArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkSrlgArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkSrlgArray.EntityData.Children = types.NewOrderedMap()
    networkSrlgArray.EntityData.Leafs = types.NewOrderedMap()
    networkSrlgArray.EntityData.Leafs.Append("set-number", types.YLeaf{"SetNumber", networkSrlgArray.SetNumber})
    networkSrlgArray.EntityData.Leafs.Append("network-srlg", types.YLeaf{"NetworkSrlg", networkSrlgArray.NetworkSrlg})

    networkSrlgArray.EntityData.YListKeys = []string {}

    return &(networkSrlgArray.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo
// Optics Alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // High Rx Power in units of 0.1 dBm.
    HighRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower

    // Low Rx Power in units of 0.1 dBm.
    LowRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower

    // High Tx Power in units of 0.1 dBm.
    HighTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower

    // Low Tx Power in units of 0.1 dBm.
    LowTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower

    // High laser bias current in units of percentage.
    HighLbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc

    // Low Temperature alarm.
    LowTemperature OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTemperature

    // High Temperature alarm.
    HighTemperature OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTemperature

    // Low Voltage alarm.
    LowVoltage OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowVoltage

    // High Voltage alarm.
    HighVoltage OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighVoltage

    // High Rx1 Power in units of 0.1 dBm.
    HighRx1Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power

    // High Rx2 Power in units of 0.1 dBm.
    HighRx2Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power

    // High Rx3 Power in units of 0.1 dBm.
    HighRx3Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power

    // High Rx4 Power in units of 0.1 dBm.
    HighRx4Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power

    // Low Rx1 Power in units of 0.1 dBm.
    LowRx1Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power

    // Low Rx2 Power in units of 0.1 dBm.
    LowRx2Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power

    // Low Rx3 Power in units of 0.1 dBm.
    LowRx3Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power

    // Low Rx4 Power in units of 0.1 dBm.
    LowRx4Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power

    // High Tx1 Power in units of 0.1 dBm.
    HighTx1Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power

    // High Tx2 Power in units of 0.1 dBm.
    HighTx2Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power

    // High Tx3 Power in units of 0.1 dBm.
    HighTx3Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power

    // High Tx4 Power in units of 0.1 dBm.
    HighTx4Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power

    // Low Tx1 Power in units of 0.1 dBm.
    LowTx1Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power

    // Low Tx2 Power in units of 0.1 dBm.
    LowTx2Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power

    // Low Tx3 Power in units of 0.1 dBm.
    LowTx3Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power

    // Low Tx4 Power in units of 0.1 dBm.
    LowTx4Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power

    // High Tx1 laser bias current in units of percentage.
    HighTx1lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1lbc

    // High Tx2 laser bias current in units of percentage.
    HighTx2lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2lbc

    // High Tx3 laser bias current in units of percentage.
    HighTx3lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3lbc

    // High Tx4 laser bias current in units of percentage.
    HighTx4lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4lbc

    // Low Tx1 laser bias current in units of percentage.
    LowTx1lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1lbc

    // Low Tx2 laser bias current in units of percentage.
    LowTx2lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2lbc

    // Low Tx3 laser bias current in units of percentage.
    LowTx3lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3lbc

    // Low Tx4 laser bias current in units of percentage.
    LowTx4lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4lbc

    // RX LOS.
    RxLos OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos

    // TX LOS.
    TxLos OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos

    // RX LOL.
    RxLol OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol

    // TX LOL.
    TxLol OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol

    // TX Fault.
    TxFault OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault

    // HI DGD.
    Hidgd OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd

    // OOR CD.
    Oorcd OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd

    // OSNR.
    Osnr OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr

    // WVL OOL.
    Wvlool OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool

    // MEA.
    Mea OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea

    // IMPROPER REM.
    ImpRemoval OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval

    // Rx LOC.
    RxLoc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc

    // Ampli Gain Deg Low.
    AmpGainDegLow OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow

    // Ampli Gain Deg High.
    AmpGainDegHigh OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh

    // TX POWER PROV MISMATCH.
    TxpwrMismatch OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch
}

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetEntityData() *types.CommonEntityData {
    opticsAlarmInfo.EntityData.YFilter = opticsAlarmInfo.YFilter
    opticsAlarmInfo.EntityData.YangName = "optics-alarm-info"
    opticsAlarmInfo.EntityData.BundleName = "cisco_ios_xr"
    opticsAlarmInfo.EntityData.ParentYangName = "optics-info"
    opticsAlarmInfo.EntityData.SegmentPath = "optics-alarm-info"
    opticsAlarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + opticsAlarmInfo.EntityData.SegmentPath
    opticsAlarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsAlarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsAlarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsAlarmInfo.EntityData.Children = types.NewOrderedMap()
    opticsAlarmInfo.EntityData.Children.Append("high-rx-power", types.YChild{"HighRxPower", &opticsAlarmInfo.HighRxPower})
    opticsAlarmInfo.EntityData.Children.Append("low-rx-power", types.YChild{"LowRxPower", &opticsAlarmInfo.LowRxPower})
    opticsAlarmInfo.EntityData.Children.Append("high-tx-power", types.YChild{"HighTxPower", &opticsAlarmInfo.HighTxPower})
    opticsAlarmInfo.EntityData.Children.Append("low-tx-power", types.YChild{"LowTxPower", &opticsAlarmInfo.LowTxPower})
    opticsAlarmInfo.EntityData.Children.Append("high-lbc", types.YChild{"HighLbc", &opticsAlarmInfo.HighLbc})
    opticsAlarmInfo.EntityData.Children.Append("low-temperature", types.YChild{"LowTemperature", &opticsAlarmInfo.LowTemperature})
    opticsAlarmInfo.EntityData.Children.Append("high-temperature", types.YChild{"HighTemperature", &opticsAlarmInfo.HighTemperature})
    opticsAlarmInfo.EntityData.Children.Append("low-voltage", types.YChild{"LowVoltage", &opticsAlarmInfo.LowVoltage})
    opticsAlarmInfo.EntityData.Children.Append("high-voltage", types.YChild{"HighVoltage", &opticsAlarmInfo.HighVoltage})
    opticsAlarmInfo.EntityData.Children.Append("high-rx1-power", types.YChild{"HighRx1Power", &opticsAlarmInfo.HighRx1Power})
    opticsAlarmInfo.EntityData.Children.Append("high-rx2-power", types.YChild{"HighRx2Power", &opticsAlarmInfo.HighRx2Power})
    opticsAlarmInfo.EntityData.Children.Append("high-rx3-power", types.YChild{"HighRx3Power", &opticsAlarmInfo.HighRx3Power})
    opticsAlarmInfo.EntityData.Children.Append("high-rx4-power", types.YChild{"HighRx4Power", &opticsAlarmInfo.HighRx4Power})
    opticsAlarmInfo.EntityData.Children.Append("low-rx1-power", types.YChild{"LowRx1Power", &opticsAlarmInfo.LowRx1Power})
    opticsAlarmInfo.EntityData.Children.Append("low-rx2-power", types.YChild{"LowRx2Power", &opticsAlarmInfo.LowRx2Power})
    opticsAlarmInfo.EntityData.Children.Append("low-rx3-power", types.YChild{"LowRx3Power", &opticsAlarmInfo.LowRx3Power})
    opticsAlarmInfo.EntityData.Children.Append("low-rx4-power", types.YChild{"LowRx4Power", &opticsAlarmInfo.LowRx4Power})
    opticsAlarmInfo.EntityData.Children.Append("high-tx1-power", types.YChild{"HighTx1Power", &opticsAlarmInfo.HighTx1Power})
    opticsAlarmInfo.EntityData.Children.Append("high-tx2-power", types.YChild{"HighTx2Power", &opticsAlarmInfo.HighTx2Power})
    opticsAlarmInfo.EntityData.Children.Append("high-tx3-power", types.YChild{"HighTx3Power", &opticsAlarmInfo.HighTx3Power})
    opticsAlarmInfo.EntityData.Children.Append("high-tx4-power", types.YChild{"HighTx4Power", &opticsAlarmInfo.HighTx4Power})
    opticsAlarmInfo.EntityData.Children.Append("low-tx1-power", types.YChild{"LowTx1Power", &opticsAlarmInfo.LowTx1Power})
    opticsAlarmInfo.EntityData.Children.Append("low-tx2-power", types.YChild{"LowTx2Power", &opticsAlarmInfo.LowTx2Power})
    opticsAlarmInfo.EntityData.Children.Append("low-tx3-power", types.YChild{"LowTx3Power", &opticsAlarmInfo.LowTx3Power})
    opticsAlarmInfo.EntityData.Children.Append("low-tx4-power", types.YChild{"LowTx4Power", &opticsAlarmInfo.LowTx4Power})
    opticsAlarmInfo.EntityData.Children.Append("high-tx1lbc", types.YChild{"HighTx1lbc", &opticsAlarmInfo.HighTx1lbc})
    opticsAlarmInfo.EntityData.Children.Append("high-tx2lbc", types.YChild{"HighTx2lbc", &opticsAlarmInfo.HighTx2lbc})
    opticsAlarmInfo.EntityData.Children.Append("high-tx3lbc", types.YChild{"HighTx3lbc", &opticsAlarmInfo.HighTx3lbc})
    opticsAlarmInfo.EntityData.Children.Append("high-tx4lbc", types.YChild{"HighTx4lbc", &opticsAlarmInfo.HighTx4lbc})
    opticsAlarmInfo.EntityData.Children.Append("low-tx1lbc", types.YChild{"LowTx1lbc", &opticsAlarmInfo.LowTx1lbc})
    opticsAlarmInfo.EntityData.Children.Append("low-tx2lbc", types.YChild{"LowTx2lbc", &opticsAlarmInfo.LowTx2lbc})
    opticsAlarmInfo.EntityData.Children.Append("low-tx3lbc", types.YChild{"LowTx3lbc", &opticsAlarmInfo.LowTx3lbc})
    opticsAlarmInfo.EntityData.Children.Append("low-tx4lbc", types.YChild{"LowTx4lbc", &opticsAlarmInfo.LowTx4lbc})
    opticsAlarmInfo.EntityData.Children.Append("rx-los", types.YChild{"RxLos", &opticsAlarmInfo.RxLos})
    opticsAlarmInfo.EntityData.Children.Append("tx-los", types.YChild{"TxLos", &opticsAlarmInfo.TxLos})
    opticsAlarmInfo.EntityData.Children.Append("rx-lol", types.YChild{"RxLol", &opticsAlarmInfo.RxLol})
    opticsAlarmInfo.EntityData.Children.Append("tx-lol", types.YChild{"TxLol", &opticsAlarmInfo.TxLol})
    opticsAlarmInfo.EntityData.Children.Append("tx-fault", types.YChild{"TxFault", &opticsAlarmInfo.TxFault})
    opticsAlarmInfo.EntityData.Children.Append("hidgd", types.YChild{"Hidgd", &opticsAlarmInfo.Hidgd})
    opticsAlarmInfo.EntityData.Children.Append("oorcd", types.YChild{"Oorcd", &opticsAlarmInfo.Oorcd})
    opticsAlarmInfo.EntityData.Children.Append("osnr", types.YChild{"Osnr", &opticsAlarmInfo.Osnr})
    opticsAlarmInfo.EntityData.Children.Append("wvlool", types.YChild{"Wvlool", &opticsAlarmInfo.Wvlool})
    opticsAlarmInfo.EntityData.Children.Append("mea", types.YChild{"Mea", &opticsAlarmInfo.Mea})
    opticsAlarmInfo.EntityData.Children.Append("imp-removal", types.YChild{"ImpRemoval", &opticsAlarmInfo.ImpRemoval})
    opticsAlarmInfo.EntityData.Children.Append("rx-loc", types.YChild{"RxLoc", &opticsAlarmInfo.RxLoc})
    opticsAlarmInfo.EntityData.Children.Append("amp-gain-deg-low", types.YChild{"AmpGainDegLow", &opticsAlarmInfo.AmpGainDegLow})
    opticsAlarmInfo.EntityData.Children.Append("amp-gain-deg-high", types.YChild{"AmpGainDegHigh", &opticsAlarmInfo.AmpGainDegHigh})
    opticsAlarmInfo.EntityData.Children.Append("txpwr-mismatch", types.YChild{"TxpwrMismatch", &opticsAlarmInfo.TxpwrMismatch})
    opticsAlarmInfo.EntityData.Leafs = types.NewOrderedMap()

    opticsAlarmInfo.EntityData.YListKeys = []string {}

    return &(opticsAlarmInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower
// High Rx Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetEntityData() *types.CommonEntityData {
    highRxPower.EntityData.YFilter = highRxPower.YFilter
    highRxPower.EntityData.YangName = "high-rx-power"
    highRxPower.EntityData.BundleName = "cisco_ios_xr"
    highRxPower.EntityData.ParentYangName = "optics-alarm-info"
    highRxPower.EntityData.SegmentPath = "high-rx-power"
    highRxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highRxPower.EntityData.SegmentPath
    highRxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highRxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highRxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highRxPower.EntityData.Children = types.NewOrderedMap()
    highRxPower.EntityData.Leafs = types.NewOrderedMap()
    highRxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highRxPower.IsDetected})
    highRxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highRxPower.Counter})

    highRxPower.EntityData.YListKeys = []string {}

    return &(highRxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower
// Low Rx Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetEntityData() *types.CommonEntityData {
    lowRxPower.EntityData.YFilter = lowRxPower.YFilter
    lowRxPower.EntityData.YangName = "low-rx-power"
    lowRxPower.EntityData.BundleName = "cisco_ios_xr"
    lowRxPower.EntityData.ParentYangName = "optics-alarm-info"
    lowRxPower.EntityData.SegmentPath = "low-rx-power"
    lowRxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowRxPower.EntityData.SegmentPath
    lowRxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowRxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowRxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowRxPower.EntityData.Children = types.NewOrderedMap()
    lowRxPower.EntityData.Leafs = types.NewOrderedMap()
    lowRxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowRxPower.IsDetected})
    lowRxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowRxPower.Counter})

    lowRxPower.EntityData.YListKeys = []string {}

    return &(lowRxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower
// High Tx Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetEntityData() *types.CommonEntityData {
    highTxPower.EntityData.YFilter = highTxPower.YFilter
    highTxPower.EntityData.YangName = "high-tx-power"
    highTxPower.EntityData.BundleName = "cisco_ios_xr"
    highTxPower.EntityData.ParentYangName = "optics-alarm-info"
    highTxPower.EntityData.SegmentPath = "high-tx-power"
    highTxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTxPower.EntityData.SegmentPath
    highTxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTxPower.EntityData.Children = types.NewOrderedMap()
    highTxPower.EntityData.Leafs = types.NewOrderedMap()
    highTxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTxPower.IsDetected})
    highTxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTxPower.Counter})

    highTxPower.EntityData.YListKeys = []string {}

    return &(highTxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower
// Low Tx Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetEntityData() *types.CommonEntityData {
    lowTxPower.EntityData.YFilter = lowTxPower.YFilter
    lowTxPower.EntityData.YangName = "low-tx-power"
    lowTxPower.EntityData.BundleName = "cisco_ios_xr"
    lowTxPower.EntityData.ParentYangName = "optics-alarm-info"
    lowTxPower.EntityData.SegmentPath = "low-tx-power"
    lowTxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTxPower.EntityData.SegmentPath
    lowTxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTxPower.EntityData.Children = types.NewOrderedMap()
    lowTxPower.EntityData.Leafs = types.NewOrderedMap()
    lowTxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTxPower.IsDetected})
    lowTxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTxPower.Counter})

    lowTxPower.EntityData.YListKeys = []string {}

    return &(lowTxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc
// High laser bias current in units of percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetEntityData() *types.CommonEntityData {
    highLbc.EntityData.YFilter = highLbc.YFilter
    highLbc.EntityData.YangName = "high-lbc"
    highLbc.EntityData.BundleName = "cisco_ios_xr"
    highLbc.EntityData.ParentYangName = "optics-alarm-info"
    highLbc.EntityData.SegmentPath = "high-lbc"
    highLbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highLbc.EntityData.SegmentPath
    highLbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highLbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highLbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highLbc.EntityData.Children = types.NewOrderedMap()
    highLbc.EntityData.Leafs = types.NewOrderedMap()
    highLbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highLbc.IsDetected})
    highLbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highLbc.Counter})

    highLbc.EntityData.YListKeys = []string {}

    return &(highLbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTemperature
// Low Temperature alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTemperature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTemperature *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTemperature) GetEntityData() *types.CommonEntityData {
    lowTemperature.EntityData.YFilter = lowTemperature.YFilter
    lowTemperature.EntityData.YangName = "low-temperature"
    lowTemperature.EntityData.BundleName = "cisco_ios_xr"
    lowTemperature.EntityData.ParentYangName = "optics-alarm-info"
    lowTemperature.EntityData.SegmentPath = "low-temperature"
    lowTemperature.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTemperature.EntityData.SegmentPath
    lowTemperature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTemperature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTemperature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTemperature.EntityData.Children = types.NewOrderedMap()
    lowTemperature.EntityData.Leafs = types.NewOrderedMap()
    lowTemperature.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTemperature.IsDetected})
    lowTemperature.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTemperature.Counter})

    lowTemperature.EntityData.YListKeys = []string {}

    return &(lowTemperature.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTemperature
// High Temperature alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTemperature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTemperature *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTemperature) GetEntityData() *types.CommonEntityData {
    highTemperature.EntityData.YFilter = highTemperature.YFilter
    highTemperature.EntityData.YangName = "high-temperature"
    highTemperature.EntityData.BundleName = "cisco_ios_xr"
    highTemperature.EntityData.ParentYangName = "optics-alarm-info"
    highTemperature.EntityData.SegmentPath = "high-temperature"
    highTemperature.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTemperature.EntityData.SegmentPath
    highTemperature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTemperature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTemperature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTemperature.EntityData.Children = types.NewOrderedMap()
    highTemperature.EntityData.Leafs = types.NewOrderedMap()
    highTemperature.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTemperature.IsDetected})
    highTemperature.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTemperature.Counter})

    highTemperature.EntityData.YListKeys = []string {}

    return &(highTemperature.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowVoltage
// Low Voltage alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowVoltage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowVoltage *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowVoltage) GetEntityData() *types.CommonEntityData {
    lowVoltage.EntityData.YFilter = lowVoltage.YFilter
    lowVoltage.EntityData.YangName = "low-voltage"
    lowVoltage.EntityData.BundleName = "cisco_ios_xr"
    lowVoltage.EntityData.ParentYangName = "optics-alarm-info"
    lowVoltage.EntityData.SegmentPath = "low-voltage"
    lowVoltage.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowVoltage.EntityData.SegmentPath
    lowVoltage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowVoltage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowVoltage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowVoltage.EntityData.Children = types.NewOrderedMap()
    lowVoltage.EntityData.Leafs = types.NewOrderedMap()
    lowVoltage.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowVoltage.IsDetected})
    lowVoltage.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowVoltage.Counter})

    lowVoltage.EntityData.YListKeys = []string {}

    return &(lowVoltage.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighVoltage
// High Voltage alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighVoltage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highVoltage *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighVoltage) GetEntityData() *types.CommonEntityData {
    highVoltage.EntityData.YFilter = highVoltage.YFilter
    highVoltage.EntityData.YangName = "high-voltage"
    highVoltage.EntityData.BundleName = "cisco_ios_xr"
    highVoltage.EntityData.ParentYangName = "optics-alarm-info"
    highVoltage.EntityData.SegmentPath = "high-voltage"
    highVoltage.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highVoltage.EntityData.SegmentPath
    highVoltage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highVoltage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highVoltage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highVoltage.EntityData.Children = types.NewOrderedMap()
    highVoltage.EntityData.Leafs = types.NewOrderedMap()
    highVoltage.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highVoltage.IsDetected})
    highVoltage.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highVoltage.Counter})

    highVoltage.EntityData.YListKeys = []string {}

    return &(highVoltage.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power
// High Rx1 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetEntityData() *types.CommonEntityData {
    highRx1Power.EntityData.YFilter = highRx1Power.YFilter
    highRx1Power.EntityData.YangName = "high-rx1-power"
    highRx1Power.EntityData.BundleName = "cisco_ios_xr"
    highRx1Power.EntityData.ParentYangName = "optics-alarm-info"
    highRx1Power.EntityData.SegmentPath = "high-rx1-power"
    highRx1Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highRx1Power.EntityData.SegmentPath
    highRx1Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highRx1Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highRx1Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highRx1Power.EntityData.Children = types.NewOrderedMap()
    highRx1Power.EntityData.Leafs = types.NewOrderedMap()
    highRx1Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highRx1Power.IsDetected})
    highRx1Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highRx1Power.Counter})

    highRx1Power.EntityData.YListKeys = []string {}

    return &(highRx1Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power
// High Rx2 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetEntityData() *types.CommonEntityData {
    highRx2Power.EntityData.YFilter = highRx2Power.YFilter
    highRx2Power.EntityData.YangName = "high-rx2-power"
    highRx2Power.EntityData.BundleName = "cisco_ios_xr"
    highRx2Power.EntityData.ParentYangName = "optics-alarm-info"
    highRx2Power.EntityData.SegmentPath = "high-rx2-power"
    highRx2Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highRx2Power.EntityData.SegmentPath
    highRx2Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highRx2Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highRx2Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highRx2Power.EntityData.Children = types.NewOrderedMap()
    highRx2Power.EntityData.Leafs = types.NewOrderedMap()
    highRx2Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highRx2Power.IsDetected})
    highRx2Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highRx2Power.Counter})

    highRx2Power.EntityData.YListKeys = []string {}

    return &(highRx2Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power
// High Rx3 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetEntityData() *types.CommonEntityData {
    highRx3Power.EntityData.YFilter = highRx3Power.YFilter
    highRx3Power.EntityData.YangName = "high-rx3-power"
    highRx3Power.EntityData.BundleName = "cisco_ios_xr"
    highRx3Power.EntityData.ParentYangName = "optics-alarm-info"
    highRx3Power.EntityData.SegmentPath = "high-rx3-power"
    highRx3Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highRx3Power.EntityData.SegmentPath
    highRx3Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highRx3Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highRx3Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highRx3Power.EntityData.Children = types.NewOrderedMap()
    highRx3Power.EntityData.Leafs = types.NewOrderedMap()
    highRx3Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highRx3Power.IsDetected})
    highRx3Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highRx3Power.Counter})

    highRx3Power.EntityData.YListKeys = []string {}

    return &(highRx3Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power
// High Rx4 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetEntityData() *types.CommonEntityData {
    highRx4Power.EntityData.YFilter = highRx4Power.YFilter
    highRx4Power.EntityData.YangName = "high-rx4-power"
    highRx4Power.EntityData.BundleName = "cisco_ios_xr"
    highRx4Power.EntityData.ParentYangName = "optics-alarm-info"
    highRx4Power.EntityData.SegmentPath = "high-rx4-power"
    highRx4Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highRx4Power.EntityData.SegmentPath
    highRx4Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highRx4Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highRx4Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highRx4Power.EntityData.Children = types.NewOrderedMap()
    highRx4Power.EntityData.Leafs = types.NewOrderedMap()
    highRx4Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highRx4Power.IsDetected})
    highRx4Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highRx4Power.Counter})

    highRx4Power.EntityData.YListKeys = []string {}

    return &(highRx4Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power
// Low Rx1 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetEntityData() *types.CommonEntityData {
    lowRx1Power.EntityData.YFilter = lowRx1Power.YFilter
    lowRx1Power.EntityData.YangName = "low-rx1-power"
    lowRx1Power.EntityData.BundleName = "cisco_ios_xr"
    lowRx1Power.EntityData.ParentYangName = "optics-alarm-info"
    lowRx1Power.EntityData.SegmentPath = "low-rx1-power"
    lowRx1Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowRx1Power.EntityData.SegmentPath
    lowRx1Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowRx1Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowRx1Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowRx1Power.EntityData.Children = types.NewOrderedMap()
    lowRx1Power.EntityData.Leafs = types.NewOrderedMap()
    lowRx1Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowRx1Power.IsDetected})
    lowRx1Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowRx1Power.Counter})

    lowRx1Power.EntityData.YListKeys = []string {}

    return &(lowRx1Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power
// Low Rx2 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetEntityData() *types.CommonEntityData {
    lowRx2Power.EntityData.YFilter = lowRx2Power.YFilter
    lowRx2Power.EntityData.YangName = "low-rx2-power"
    lowRx2Power.EntityData.BundleName = "cisco_ios_xr"
    lowRx2Power.EntityData.ParentYangName = "optics-alarm-info"
    lowRx2Power.EntityData.SegmentPath = "low-rx2-power"
    lowRx2Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowRx2Power.EntityData.SegmentPath
    lowRx2Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowRx2Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowRx2Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowRx2Power.EntityData.Children = types.NewOrderedMap()
    lowRx2Power.EntityData.Leafs = types.NewOrderedMap()
    lowRx2Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowRx2Power.IsDetected})
    lowRx2Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowRx2Power.Counter})

    lowRx2Power.EntityData.YListKeys = []string {}

    return &(lowRx2Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power
// Low Rx3 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetEntityData() *types.CommonEntityData {
    lowRx3Power.EntityData.YFilter = lowRx3Power.YFilter
    lowRx3Power.EntityData.YangName = "low-rx3-power"
    lowRx3Power.EntityData.BundleName = "cisco_ios_xr"
    lowRx3Power.EntityData.ParentYangName = "optics-alarm-info"
    lowRx3Power.EntityData.SegmentPath = "low-rx3-power"
    lowRx3Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowRx3Power.EntityData.SegmentPath
    lowRx3Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowRx3Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowRx3Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowRx3Power.EntityData.Children = types.NewOrderedMap()
    lowRx3Power.EntityData.Leafs = types.NewOrderedMap()
    lowRx3Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowRx3Power.IsDetected})
    lowRx3Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowRx3Power.Counter})

    lowRx3Power.EntityData.YListKeys = []string {}

    return &(lowRx3Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power
// Low Rx4 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetEntityData() *types.CommonEntityData {
    lowRx4Power.EntityData.YFilter = lowRx4Power.YFilter
    lowRx4Power.EntityData.YangName = "low-rx4-power"
    lowRx4Power.EntityData.BundleName = "cisco_ios_xr"
    lowRx4Power.EntityData.ParentYangName = "optics-alarm-info"
    lowRx4Power.EntityData.SegmentPath = "low-rx4-power"
    lowRx4Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowRx4Power.EntityData.SegmentPath
    lowRx4Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowRx4Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowRx4Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowRx4Power.EntityData.Children = types.NewOrderedMap()
    lowRx4Power.EntityData.Leafs = types.NewOrderedMap()
    lowRx4Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowRx4Power.IsDetected})
    lowRx4Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowRx4Power.Counter})

    lowRx4Power.EntityData.YListKeys = []string {}

    return &(lowRx4Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power
// High Tx1 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetEntityData() *types.CommonEntityData {
    highTx1Power.EntityData.YFilter = highTx1Power.YFilter
    highTx1Power.EntityData.YangName = "high-tx1-power"
    highTx1Power.EntityData.BundleName = "cisco_ios_xr"
    highTx1Power.EntityData.ParentYangName = "optics-alarm-info"
    highTx1Power.EntityData.SegmentPath = "high-tx1-power"
    highTx1Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTx1Power.EntityData.SegmentPath
    highTx1Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTx1Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTx1Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTx1Power.EntityData.Children = types.NewOrderedMap()
    highTx1Power.EntityData.Leafs = types.NewOrderedMap()
    highTx1Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTx1Power.IsDetected})
    highTx1Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTx1Power.Counter})

    highTx1Power.EntityData.YListKeys = []string {}

    return &(highTx1Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power
// High Tx2 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetEntityData() *types.CommonEntityData {
    highTx2Power.EntityData.YFilter = highTx2Power.YFilter
    highTx2Power.EntityData.YangName = "high-tx2-power"
    highTx2Power.EntityData.BundleName = "cisco_ios_xr"
    highTx2Power.EntityData.ParentYangName = "optics-alarm-info"
    highTx2Power.EntityData.SegmentPath = "high-tx2-power"
    highTx2Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTx2Power.EntityData.SegmentPath
    highTx2Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTx2Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTx2Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTx2Power.EntityData.Children = types.NewOrderedMap()
    highTx2Power.EntityData.Leafs = types.NewOrderedMap()
    highTx2Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTx2Power.IsDetected})
    highTx2Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTx2Power.Counter})

    highTx2Power.EntityData.YListKeys = []string {}

    return &(highTx2Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power
// High Tx3 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetEntityData() *types.CommonEntityData {
    highTx3Power.EntityData.YFilter = highTx3Power.YFilter
    highTx3Power.EntityData.YangName = "high-tx3-power"
    highTx3Power.EntityData.BundleName = "cisco_ios_xr"
    highTx3Power.EntityData.ParentYangName = "optics-alarm-info"
    highTx3Power.EntityData.SegmentPath = "high-tx3-power"
    highTx3Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTx3Power.EntityData.SegmentPath
    highTx3Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTx3Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTx3Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTx3Power.EntityData.Children = types.NewOrderedMap()
    highTx3Power.EntityData.Leafs = types.NewOrderedMap()
    highTx3Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTx3Power.IsDetected})
    highTx3Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTx3Power.Counter})

    highTx3Power.EntityData.YListKeys = []string {}

    return &(highTx3Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power
// High Tx4 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetEntityData() *types.CommonEntityData {
    highTx4Power.EntityData.YFilter = highTx4Power.YFilter
    highTx4Power.EntityData.YangName = "high-tx4-power"
    highTx4Power.EntityData.BundleName = "cisco_ios_xr"
    highTx4Power.EntityData.ParentYangName = "optics-alarm-info"
    highTx4Power.EntityData.SegmentPath = "high-tx4-power"
    highTx4Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTx4Power.EntityData.SegmentPath
    highTx4Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTx4Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTx4Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTx4Power.EntityData.Children = types.NewOrderedMap()
    highTx4Power.EntityData.Leafs = types.NewOrderedMap()
    highTx4Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTx4Power.IsDetected})
    highTx4Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTx4Power.Counter})

    highTx4Power.EntityData.YListKeys = []string {}

    return &(highTx4Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power
// Low Tx1 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetEntityData() *types.CommonEntityData {
    lowTx1Power.EntityData.YFilter = lowTx1Power.YFilter
    lowTx1Power.EntityData.YangName = "low-tx1-power"
    lowTx1Power.EntityData.BundleName = "cisco_ios_xr"
    lowTx1Power.EntityData.ParentYangName = "optics-alarm-info"
    lowTx1Power.EntityData.SegmentPath = "low-tx1-power"
    lowTx1Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTx1Power.EntityData.SegmentPath
    lowTx1Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTx1Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTx1Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTx1Power.EntityData.Children = types.NewOrderedMap()
    lowTx1Power.EntityData.Leafs = types.NewOrderedMap()
    lowTx1Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTx1Power.IsDetected})
    lowTx1Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTx1Power.Counter})

    lowTx1Power.EntityData.YListKeys = []string {}

    return &(lowTx1Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power
// Low Tx2 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetEntityData() *types.CommonEntityData {
    lowTx2Power.EntityData.YFilter = lowTx2Power.YFilter
    lowTx2Power.EntityData.YangName = "low-tx2-power"
    lowTx2Power.EntityData.BundleName = "cisco_ios_xr"
    lowTx2Power.EntityData.ParentYangName = "optics-alarm-info"
    lowTx2Power.EntityData.SegmentPath = "low-tx2-power"
    lowTx2Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTx2Power.EntityData.SegmentPath
    lowTx2Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTx2Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTx2Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTx2Power.EntityData.Children = types.NewOrderedMap()
    lowTx2Power.EntityData.Leafs = types.NewOrderedMap()
    lowTx2Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTx2Power.IsDetected})
    lowTx2Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTx2Power.Counter})

    lowTx2Power.EntityData.YListKeys = []string {}

    return &(lowTx2Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power
// Low Tx3 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetEntityData() *types.CommonEntityData {
    lowTx3Power.EntityData.YFilter = lowTx3Power.YFilter
    lowTx3Power.EntityData.YangName = "low-tx3-power"
    lowTx3Power.EntityData.BundleName = "cisco_ios_xr"
    lowTx3Power.EntityData.ParentYangName = "optics-alarm-info"
    lowTx3Power.EntityData.SegmentPath = "low-tx3-power"
    lowTx3Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTx3Power.EntityData.SegmentPath
    lowTx3Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTx3Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTx3Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTx3Power.EntityData.Children = types.NewOrderedMap()
    lowTx3Power.EntityData.Leafs = types.NewOrderedMap()
    lowTx3Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTx3Power.IsDetected})
    lowTx3Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTx3Power.Counter})

    lowTx3Power.EntityData.YListKeys = []string {}

    return &(lowTx3Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power
// Low Tx4 Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetEntityData() *types.CommonEntityData {
    lowTx4Power.EntityData.YFilter = lowTx4Power.YFilter
    lowTx4Power.EntityData.YangName = "low-tx4-power"
    lowTx4Power.EntityData.BundleName = "cisco_ios_xr"
    lowTx4Power.EntityData.ParentYangName = "optics-alarm-info"
    lowTx4Power.EntityData.SegmentPath = "low-tx4-power"
    lowTx4Power.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTx4Power.EntityData.SegmentPath
    lowTx4Power.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTx4Power.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTx4Power.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTx4Power.EntityData.Children = types.NewOrderedMap()
    lowTx4Power.EntityData.Leafs = types.NewOrderedMap()
    lowTx4Power.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTx4Power.IsDetected})
    lowTx4Power.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTx4Power.Counter})

    lowTx4Power.EntityData.YListKeys = []string {}

    return &(lowTx4Power.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1lbc
// High Tx1 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1lbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx1lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1lbc) GetEntityData() *types.CommonEntityData {
    highTx1lbc.EntityData.YFilter = highTx1lbc.YFilter
    highTx1lbc.EntityData.YangName = "high-tx1lbc"
    highTx1lbc.EntityData.BundleName = "cisco_ios_xr"
    highTx1lbc.EntityData.ParentYangName = "optics-alarm-info"
    highTx1lbc.EntityData.SegmentPath = "high-tx1lbc"
    highTx1lbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTx1lbc.EntityData.SegmentPath
    highTx1lbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTx1lbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTx1lbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTx1lbc.EntityData.Children = types.NewOrderedMap()
    highTx1lbc.EntityData.Leafs = types.NewOrderedMap()
    highTx1lbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTx1lbc.IsDetected})
    highTx1lbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTx1lbc.Counter})

    highTx1lbc.EntityData.YListKeys = []string {}

    return &(highTx1lbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2lbc
// High Tx2 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2lbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx2lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2lbc) GetEntityData() *types.CommonEntityData {
    highTx2lbc.EntityData.YFilter = highTx2lbc.YFilter
    highTx2lbc.EntityData.YangName = "high-tx2lbc"
    highTx2lbc.EntityData.BundleName = "cisco_ios_xr"
    highTx2lbc.EntityData.ParentYangName = "optics-alarm-info"
    highTx2lbc.EntityData.SegmentPath = "high-tx2lbc"
    highTx2lbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTx2lbc.EntityData.SegmentPath
    highTx2lbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTx2lbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTx2lbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTx2lbc.EntityData.Children = types.NewOrderedMap()
    highTx2lbc.EntityData.Leafs = types.NewOrderedMap()
    highTx2lbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTx2lbc.IsDetected})
    highTx2lbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTx2lbc.Counter})

    highTx2lbc.EntityData.YListKeys = []string {}

    return &(highTx2lbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3lbc
// High Tx3 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3lbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx3lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3lbc) GetEntityData() *types.CommonEntityData {
    highTx3lbc.EntityData.YFilter = highTx3lbc.YFilter
    highTx3lbc.EntityData.YangName = "high-tx3lbc"
    highTx3lbc.EntityData.BundleName = "cisco_ios_xr"
    highTx3lbc.EntityData.ParentYangName = "optics-alarm-info"
    highTx3lbc.EntityData.SegmentPath = "high-tx3lbc"
    highTx3lbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTx3lbc.EntityData.SegmentPath
    highTx3lbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTx3lbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTx3lbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTx3lbc.EntityData.Children = types.NewOrderedMap()
    highTx3lbc.EntityData.Leafs = types.NewOrderedMap()
    highTx3lbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTx3lbc.IsDetected})
    highTx3lbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTx3lbc.Counter})

    highTx3lbc.EntityData.YListKeys = []string {}

    return &(highTx3lbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4lbc
// High Tx4 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4lbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx4lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4lbc) GetEntityData() *types.CommonEntityData {
    highTx4lbc.EntityData.YFilter = highTx4lbc.YFilter
    highTx4lbc.EntityData.YangName = "high-tx4lbc"
    highTx4lbc.EntityData.BundleName = "cisco_ios_xr"
    highTx4lbc.EntityData.ParentYangName = "optics-alarm-info"
    highTx4lbc.EntityData.SegmentPath = "high-tx4lbc"
    highTx4lbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + highTx4lbc.EntityData.SegmentPath
    highTx4lbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTx4lbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTx4lbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTx4lbc.EntityData.Children = types.NewOrderedMap()
    highTx4lbc.EntityData.Leafs = types.NewOrderedMap()
    highTx4lbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTx4lbc.IsDetected})
    highTx4lbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTx4lbc.Counter})

    highTx4lbc.EntityData.YListKeys = []string {}

    return &(highTx4lbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1lbc
// Low Tx1 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1lbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx1lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1lbc) GetEntityData() *types.CommonEntityData {
    lowTx1lbc.EntityData.YFilter = lowTx1lbc.YFilter
    lowTx1lbc.EntityData.YangName = "low-tx1lbc"
    lowTx1lbc.EntityData.BundleName = "cisco_ios_xr"
    lowTx1lbc.EntityData.ParentYangName = "optics-alarm-info"
    lowTx1lbc.EntityData.SegmentPath = "low-tx1lbc"
    lowTx1lbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTx1lbc.EntityData.SegmentPath
    lowTx1lbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTx1lbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTx1lbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTx1lbc.EntityData.Children = types.NewOrderedMap()
    lowTx1lbc.EntityData.Leafs = types.NewOrderedMap()
    lowTx1lbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTx1lbc.IsDetected})
    lowTx1lbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTx1lbc.Counter})

    lowTx1lbc.EntityData.YListKeys = []string {}

    return &(lowTx1lbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2lbc
// Low Tx2 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2lbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx2lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2lbc) GetEntityData() *types.CommonEntityData {
    lowTx2lbc.EntityData.YFilter = lowTx2lbc.YFilter
    lowTx2lbc.EntityData.YangName = "low-tx2lbc"
    lowTx2lbc.EntityData.BundleName = "cisco_ios_xr"
    lowTx2lbc.EntityData.ParentYangName = "optics-alarm-info"
    lowTx2lbc.EntityData.SegmentPath = "low-tx2lbc"
    lowTx2lbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTx2lbc.EntityData.SegmentPath
    lowTx2lbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTx2lbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTx2lbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTx2lbc.EntityData.Children = types.NewOrderedMap()
    lowTx2lbc.EntityData.Leafs = types.NewOrderedMap()
    lowTx2lbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTx2lbc.IsDetected})
    lowTx2lbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTx2lbc.Counter})

    lowTx2lbc.EntityData.YListKeys = []string {}

    return &(lowTx2lbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3lbc
// Low Tx3 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3lbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx3lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3lbc) GetEntityData() *types.CommonEntityData {
    lowTx3lbc.EntityData.YFilter = lowTx3lbc.YFilter
    lowTx3lbc.EntityData.YangName = "low-tx3lbc"
    lowTx3lbc.EntityData.BundleName = "cisco_ios_xr"
    lowTx3lbc.EntityData.ParentYangName = "optics-alarm-info"
    lowTx3lbc.EntityData.SegmentPath = "low-tx3lbc"
    lowTx3lbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTx3lbc.EntityData.SegmentPath
    lowTx3lbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTx3lbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTx3lbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTx3lbc.EntityData.Children = types.NewOrderedMap()
    lowTx3lbc.EntityData.Leafs = types.NewOrderedMap()
    lowTx3lbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTx3lbc.IsDetected})
    lowTx3lbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTx3lbc.Counter})

    lowTx3lbc.EntityData.YListKeys = []string {}

    return &(lowTx3lbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4lbc
// Low Tx4 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4lbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx4lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4lbc) GetEntityData() *types.CommonEntityData {
    lowTx4lbc.EntityData.YFilter = lowTx4lbc.YFilter
    lowTx4lbc.EntityData.YangName = "low-tx4lbc"
    lowTx4lbc.EntityData.BundleName = "cisco_ios_xr"
    lowTx4lbc.EntityData.ParentYangName = "optics-alarm-info"
    lowTx4lbc.EntityData.SegmentPath = "low-tx4lbc"
    lowTx4lbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + lowTx4lbc.EntityData.SegmentPath
    lowTx4lbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTx4lbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTx4lbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTx4lbc.EntityData.Children = types.NewOrderedMap()
    lowTx4lbc.EntityData.Leafs = types.NewOrderedMap()
    lowTx4lbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTx4lbc.IsDetected})
    lowTx4lbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTx4lbc.Counter})

    lowTx4lbc.EntityData.YListKeys = []string {}

    return &(lowTx4lbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos
// RX LOS
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetEntityData() *types.CommonEntityData {
    rxLos.EntityData.YFilter = rxLos.YFilter
    rxLos.EntityData.YangName = "rx-los"
    rxLos.EntityData.BundleName = "cisco_ios_xr"
    rxLos.EntityData.ParentYangName = "optics-alarm-info"
    rxLos.EntityData.SegmentPath = "rx-los"
    rxLos.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + rxLos.EntityData.SegmentPath
    rxLos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLos.EntityData.Children = types.NewOrderedMap()
    rxLos.EntityData.Leafs = types.NewOrderedMap()
    rxLos.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", rxLos.IsDetected})
    rxLos.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", rxLos.Counter})

    rxLos.EntityData.YListKeys = []string {}

    return &(rxLos.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos
// TX LOS
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetEntityData() *types.CommonEntityData {
    txLos.EntityData.YFilter = txLos.YFilter
    txLos.EntityData.YangName = "tx-los"
    txLos.EntityData.BundleName = "cisco_ios_xr"
    txLos.EntityData.ParentYangName = "optics-alarm-info"
    txLos.EntityData.SegmentPath = "tx-los"
    txLos.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + txLos.EntityData.SegmentPath
    txLos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLos.EntityData.Children = types.NewOrderedMap()
    txLos.EntityData.Leafs = types.NewOrderedMap()
    txLos.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", txLos.IsDetected})
    txLos.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", txLos.Counter})

    txLos.EntityData.YListKeys = []string {}

    return &(txLos.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol
// RX LOL
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetEntityData() *types.CommonEntityData {
    rxLol.EntityData.YFilter = rxLol.YFilter
    rxLol.EntityData.YangName = "rx-lol"
    rxLol.EntityData.BundleName = "cisco_ios_xr"
    rxLol.EntityData.ParentYangName = "optics-alarm-info"
    rxLol.EntityData.SegmentPath = "rx-lol"
    rxLol.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + rxLol.EntityData.SegmentPath
    rxLol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLol.EntityData.Children = types.NewOrderedMap()
    rxLol.EntityData.Leafs = types.NewOrderedMap()
    rxLol.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", rxLol.IsDetected})
    rxLol.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", rxLol.Counter})

    rxLol.EntityData.YListKeys = []string {}

    return &(rxLol.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol
// TX LOL
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetEntityData() *types.CommonEntityData {
    txLol.EntityData.YFilter = txLol.YFilter
    txLol.EntityData.YangName = "tx-lol"
    txLol.EntityData.BundleName = "cisco_ios_xr"
    txLol.EntityData.ParentYangName = "optics-alarm-info"
    txLol.EntityData.SegmentPath = "tx-lol"
    txLol.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + txLol.EntityData.SegmentPath
    txLol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLol.EntityData.Children = types.NewOrderedMap()
    txLol.EntityData.Leafs = types.NewOrderedMap()
    txLol.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", txLol.IsDetected})
    txLol.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", txLol.Counter})

    txLol.EntityData.YListKeys = []string {}

    return &(txLol.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault
// TX Fault
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetEntityData() *types.CommonEntityData {
    txFault.EntityData.YFilter = txFault.YFilter
    txFault.EntityData.YangName = "tx-fault"
    txFault.EntityData.BundleName = "cisco_ios_xr"
    txFault.EntityData.ParentYangName = "optics-alarm-info"
    txFault.EntityData.SegmentPath = "tx-fault"
    txFault.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + txFault.EntityData.SegmentPath
    txFault.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txFault.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txFault.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txFault.EntityData.Children = types.NewOrderedMap()
    txFault.EntityData.Leafs = types.NewOrderedMap()
    txFault.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", txFault.IsDetected})
    txFault.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", txFault.Counter})

    txFault.EntityData.YListKeys = []string {}

    return &(txFault.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd
// HI DGD
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetEntityData() *types.CommonEntityData {
    hidgd.EntityData.YFilter = hidgd.YFilter
    hidgd.EntityData.YangName = "hidgd"
    hidgd.EntityData.BundleName = "cisco_ios_xr"
    hidgd.EntityData.ParentYangName = "optics-alarm-info"
    hidgd.EntityData.SegmentPath = "hidgd"
    hidgd.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + hidgd.EntityData.SegmentPath
    hidgd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hidgd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hidgd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hidgd.EntityData.Children = types.NewOrderedMap()
    hidgd.EntityData.Leafs = types.NewOrderedMap()
    hidgd.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hidgd.IsDetected})
    hidgd.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hidgd.Counter})

    hidgd.EntityData.YListKeys = []string {}

    return &(hidgd.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd
// OOR CD
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetEntityData() *types.CommonEntityData {
    oorcd.EntityData.YFilter = oorcd.YFilter
    oorcd.EntityData.YangName = "oorcd"
    oorcd.EntityData.BundleName = "cisco_ios_xr"
    oorcd.EntityData.ParentYangName = "optics-alarm-info"
    oorcd.EntityData.SegmentPath = "oorcd"
    oorcd.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + oorcd.EntityData.SegmentPath
    oorcd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorcd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorcd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorcd.EntityData.Children = types.NewOrderedMap()
    oorcd.EntityData.Leafs = types.NewOrderedMap()
    oorcd.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", oorcd.IsDetected})
    oorcd.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", oorcd.Counter})

    oorcd.EntityData.YListKeys = []string {}

    return &(oorcd.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr
// OSNR
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetEntityData() *types.CommonEntityData {
    osnr.EntityData.YFilter = osnr.YFilter
    osnr.EntityData.YangName = "osnr"
    osnr.EntityData.BundleName = "cisco_ios_xr"
    osnr.EntityData.ParentYangName = "optics-alarm-info"
    osnr.EntityData.SegmentPath = "osnr"
    osnr.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + osnr.EntityData.SegmentPath
    osnr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    osnr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    osnr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    osnr.EntityData.Children = types.NewOrderedMap()
    osnr.EntityData.Leafs = types.NewOrderedMap()
    osnr.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", osnr.IsDetected})
    osnr.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", osnr.Counter})

    osnr.EntityData.YListKeys = []string {}

    return &(osnr.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool
// WVL OOL
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetEntityData() *types.CommonEntityData {
    wvlool.EntityData.YFilter = wvlool.YFilter
    wvlool.EntityData.YangName = "wvlool"
    wvlool.EntityData.BundleName = "cisco_ios_xr"
    wvlool.EntityData.ParentYangName = "optics-alarm-info"
    wvlool.EntityData.SegmentPath = "wvlool"
    wvlool.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + wvlool.EntityData.SegmentPath
    wvlool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wvlool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wvlool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wvlool.EntityData.Children = types.NewOrderedMap()
    wvlool.EntityData.Leafs = types.NewOrderedMap()
    wvlool.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", wvlool.IsDetected})
    wvlool.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", wvlool.Counter})

    wvlool.EntityData.YListKeys = []string {}

    return &(wvlool.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea
// MEA
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetEntityData() *types.CommonEntityData {
    mea.EntityData.YFilter = mea.YFilter
    mea.EntityData.YangName = "mea"
    mea.EntityData.BundleName = "cisco_ios_xr"
    mea.EntityData.ParentYangName = "optics-alarm-info"
    mea.EntityData.SegmentPath = "mea"
    mea.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + mea.EntityData.SegmentPath
    mea.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mea.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mea.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mea.EntityData.Children = types.NewOrderedMap()
    mea.EntityData.Leafs = types.NewOrderedMap()
    mea.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", mea.IsDetected})
    mea.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", mea.Counter})

    mea.EntityData.YListKeys = []string {}

    return &(mea.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval
// IMPROPER REM
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetEntityData() *types.CommonEntityData {
    impRemoval.EntityData.YFilter = impRemoval.YFilter
    impRemoval.EntityData.YangName = "imp-removal"
    impRemoval.EntityData.BundleName = "cisco_ios_xr"
    impRemoval.EntityData.ParentYangName = "optics-alarm-info"
    impRemoval.EntityData.SegmentPath = "imp-removal"
    impRemoval.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + impRemoval.EntityData.SegmentPath
    impRemoval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    impRemoval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    impRemoval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    impRemoval.EntityData.Children = types.NewOrderedMap()
    impRemoval.EntityData.Leafs = types.NewOrderedMap()
    impRemoval.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", impRemoval.IsDetected})
    impRemoval.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", impRemoval.Counter})

    impRemoval.EntityData.YListKeys = []string {}

    return &(impRemoval.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc
// Rx LOC
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetEntityData() *types.CommonEntityData {
    rxLoc.EntityData.YFilter = rxLoc.YFilter
    rxLoc.EntityData.YangName = "rx-loc"
    rxLoc.EntityData.BundleName = "cisco_ios_xr"
    rxLoc.EntityData.ParentYangName = "optics-alarm-info"
    rxLoc.EntityData.SegmentPath = "rx-loc"
    rxLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + rxLoc.EntityData.SegmentPath
    rxLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLoc.EntityData.Children = types.NewOrderedMap()
    rxLoc.EntityData.Leafs = types.NewOrderedMap()
    rxLoc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", rxLoc.IsDetected})
    rxLoc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", rxLoc.Counter})

    rxLoc.EntityData.YListKeys = []string {}

    return &(rxLoc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow
// Ampli Gain Deg Low
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetEntityData() *types.CommonEntityData {
    ampGainDegLow.EntityData.YFilter = ampGainDegLow.YFilter
    ampGainDegLow.EntityData.YangName = "amp-gain-deg-low"
    ampGainDegLow.EntityData.BundleName = "cisco_ios_xr"
    ampGainDegLow.EntityData.ParentYangName = "optics-alarm-info"
    ampGainDegLow.EntityData.SegmentPath = "amp-gain-deg-low"
    ampGainDegLow.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + ampGainDegLow.EntityData.SegmentPath
    ampGainDegLow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ampGainDegLow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ampGainDegLow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ampGainDegLow.EntityData.Children = types.NewOrderedMap()
    ampGainDegLow.EntityData.Leafs = types.NewOrderedMap()
    ampGainDegLow.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", ampGainDegLow.IsDetected})
    ampGainDegLow.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", ampGainDegLow.Counter})

    ampGainDegLow.EntityData.YListKeys = []string {}

    return &(ampGainDegLow.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh
// Ampli Gain Deg High
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetEntityData() *types.CommonEntityData {
    ampGainDegHigh.EntityData.YFilter = ampGainDegHigh.YFilter
    ampGainDegHigh.EntityData.YangName = "amp-gain-deg-high"
    ampGainDegHigh.EntityData.BundleName = "cisco_ios_xr"
    ampGainDegHigh.EntityData.ParentYangName = "optics-alarm-info"
    ampGainDegHigh.EntityData.SegmentPath = "amp-gain-deg-high"
    ampGainDegHigh.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + ampGainDegHigh.EntityData.SegmentPath
    ampGainDegHigh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ampGainDegHigh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ampGainDegHigh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ampGainDegHigh.EntityData.Children = types.NewOrderedMap()
    ampGainDegHigh.EntityData.Leafs = types.NewOrderedMap()
    ampGainDegHigh.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", ampGainDegHigh.IsDetected})
    ampGainDegHigh.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", ampGainDegHigh.Counter})

    ampGainDegHigh.EntityData.YListKeys = []string {}

    return &(ampGainDegHigh.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch
// TX POWER PROV MISMATCH
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetEntityData() *types.CommonEntityData {
    txpwrMismatch.EntityData.YFilter = txpwrMismatch.YFilter
    txpwrMismatch.EntityData.YangName = "txpwr-mismatch"
    txpwrMismatch.EntityData.BundleName = "cisco_ios_xr"
    txpwrMismatch.EntityData.ParentYangName = "optics-alarm-info"
    txpwrMismatch.EntityData.SegmentPath = "txpwr-mismatch"
    txpwrMismatch.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/optics-alarm-info/" + txpwrMismatch.EntityData.SegmentPath
    txpwrMismatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txpwrMismatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txpwrMismatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txpwrMismatch.EntityData.Children = types.NewOrderedMap()
    txpwrMismatch.EntityData.Leafs = types.NewOrderedMap()
    txpwrMismatch.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", txpwrMismatch.IsDetected})
    txpwrMismatch.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", txpwrMismatch.Counter})

    txpwrMismatch.EntityData.YListKeys = []string {}

    return &(txpwrMismatch.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo
// Ots Alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Low Tx Power in units of 0.1 dBm.
    LowTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower

    // Low Rx Power in units of 0.1 dBm.
    LowRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower

    // Rx LOS_P.
    RxLosP OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP

    // Rx LOC.
    RxLoc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc

    // Ampli Gain Deg Low.
    AmpGainDegLow OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow

    // Ampli Gain Deg High.
    AmpGainDegHigh OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh

    // Auto Laser Shut.
    AutoLaserShut OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut

    // Auto Power Red.
    AutoPowerRed OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed

    // Auto Ampli Ctrl Disabled.
    AutoAmpliCtrlDisabled OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled

    // Auto Ampli Ctrl Config Mismatch.
    AutoAmpliCtrlConfigMismatch OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch

    // Switch To Protect.
    SwitchToProtect OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect

    // Auto Ampli Ctrl Running.
    AutoAmpliCtrlRunning OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning
}

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetEntityData() *types.CommonEntityData {
    otsAlarmInfo.EntityData.YFilter = otsAlarmInfo.YFilter
    otsAlarmInfo.EntityData.YangName = "ots-alarm-info"
    otsAlarmInfo.EntityData.BundleName = "cisco_ios_xr"
    otsAlarmInfo.EntityData.ParentYangName = "optics-info"
    otsAlarmInfo.EntityData.SegmentPath = "ots-alarm-info"
    otsAlarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + otsAlarmInfo.EntityData.SegmentPath
    otsAlarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otsAlarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otsAlarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otsAlarmInfo.EntityData.Children = types.NewOrderedMap()
    otsAlarmInfo.EntityData.Children.Append("low-tx-power", types.YChild{"LowTxPower", &otsAlarmInfo.LowTxPower})
    otsAlarmInfo.EntityData.Children.Append("low-rx-power", types.YChild{"LowRxPower", &otsAlarmInfo.LowRxPower})
    otsAlarmInfo.EntityData.Children.Append("rx-los-p", types.YChild{"RxLosP", &otsAlarmInfo.RxLosP})
    otsAlarmInfo.EntityData.Children.Append("rx-loc", types.YChild{"RxLoc", &otsAlarmInfo.RxLoc})
    otsAlarmInfo.EntityData.Children.Append("amp-gain-deg-low", types.YChild{"AmpGainDegLow", &otsAlarmInfo.AmpGainDegLow})
    otsAlarmInfo.EntityData.Children.Append("amp-gain-deg-high", types.YChild{"AmpGainDegHigh", &otsAlarmInfo.AmpGainDegHigh})
    otsAlarmInfo.EntityData.Children.Append("auto-laser-shut", types.YChild{"AutoLaserShut", &otsAlarmInfo.AutoLaserShut})
    otsAlarmInfo.EntityData.Children.Append("auto-power-red", types.YChild{"AutoPowerRed", &otsAlarmInfo.AutoPowerRed})
    otsAlarmInfo.EntityData.Children.Append("auto-ampli-ctrl-disabled", types.YChild{"AutoAmpliCtrlDisabled", &otsAlarmInfo.AutoAmpliCtrlDisabled})
    otsAlarmInfo.EntityData.Children.Append("auto-ampli-ctrl-config-mismatch", types.YChild{"AutoAmpliCtrlConfigMismatch", &otsAlarmInfo.AutoAmpliCtrlConfigMismatch})
    otsAlarmInfo.EntityData.Children.Append("switch-to-protect", types.YChild{"SwitchToProtect", &otsAlarmInfo.SwitchToProtect})
    otsAlarmInfo.EntityData.Children.Append("auto-ampli-ctrl-running", types.YChild{"AutoAmpliCtrlRunning", &otsAlarmInfo.AutoAmpliCtrlRunning})
    otsAlarmInfo.EntityData.Leafs = types.NewOrderedMap()

    otsAlarmInfo.EntityData.YListKeys = []string {}

    return &(otsAlarmInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower
// Low Tx Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetEntityData() *types.CommonEntityData {
    lowTxPower.EntityData.YFilter = lowTxPower.YFilter
    lowTxPower.EntityData.YangName = "low-tx-power"
    lowTxPower.EntityData.BundleName = "cisco_ios_xr"
    lowTxPower.EntityData.ParentYangName = "ots-alarm-info"
    lowTxPower.EntityData.SegmentPath = "low-tx-power"
    lowTxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + lowTxPower.EntityData.SegmentPath
    lowTxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTxPower.EntityData.Children = types.NewOrderedMap()
    lowTxPower.EntityData.Leafs = types.NewOrderedMap()
    lowTxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTxPower.IsDetected})
    lowTxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTxPower.Counter})

    lowTxPower.EntityData.YListKeys = []string {}

    return &(lowTxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower
// Low Rx Power in units of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetEntityData() *types.CommonEntityData {
    lowRxPower.EntityData.YFilter = lowRxPower.YFilter
    lowRxPower.EntityData.YangName = "low-rx-power"
    lowRxPower.EntityData.BundleName = "cisco_ios_xr"
    lowRxPower.EntityData.ParentYangName = "ots-alarm-info"
    lowRxPower.EntityData.SegmentPath = "low-rx-power"
    lowRxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + lowRxPower.EntityData.SegmentPath
    lowRxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowRxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowRxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowRxPower.EntityData.Children = types.NewOrderedMap()
    lowRxPower.EntityData.Leafs = types.NewOrderedMap()
    lowRxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowRxPower.IsDetected})
    lowRxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowRxPower.Counter})

    lowRxPower.EntityData.YListKeys = []string {}

    return &(lowRxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP
// Rx LOS_P
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetEntityData() *types.CommonEntityData {
    rxLosP.EntityData.YFilter = rxLosP.YFilter
    rxLosP.EntityData.YangName = "rx-los-p"
    rxLosP.EntityData.BundleName = "cisco_ios_xr"
    rxLosP.EntityData.ParentYangName = "ots-alarm-info"
    rxLosP.EntityData.SegmentPath = "rx-los-p"
    rxLosP.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + rxLosP.EntityData.SegmentPath
    rxLosP.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLosP.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLosP.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLosP.EntityData.Children = types.NewOrderedMap()
    rxLosP.EntityData.Leafs = types.NewOrderedMap()
    rxLosP.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", rxLosP.IsDetected})
    rxLosP.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", rxLosP.Counter})

    rxLosP.EntityData.YListKeys = []string {}

    return &(rxLosP.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc
// Rx LOC
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetEntityData() *types.CommonEntityData {
    rxLoc.EntityData.YFilter = rxLoc.YFilter
    rxLoc.EntityData.YangName = "rx-loc"
    rxLoc.EntityData.BundleName = "cisco_ios_xr"
    rxLoc.EntityData.ParentYangName = "ots-alarm-info"
    rxLoc.EntityData.SegmentPath = "rx-loc"
    rxLoc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + rxLoc.EntityData.SegmentPath
    rxLoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLoc.EntityData.Children = types.NewOrderedMap()
    rxLoc.EntityData.Leafs = types.NewOrderedMap()
    rxLoc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", rxLoc.IsDetected})
    rxLoc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", rxLoc.Counter})

    rxLoc.EntityData.YListKeys = []string {}

    return &(rxLoc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow
// Ampli Gain Deg Low
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetEntityData() *types.CommonEntityData {
    ampGainDegLow.EntityData.YFilter = ampGainDegLow.YFilter
    ampGainDegLow.EntityData.YangName = "amp-gain-deg-low"
    ampGainDegLow.EntityData.BundleName = "cisco_ios_xr"
    ampGainDegLow.EntityData.ParentYangName = "ots-alarm-info"
    ampGainDegLow.EntityData.SegmentPath = "amp-gain-deg-low"
    ampGainDegLow.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + ampGainDegLow.EntityData.SegmentPath
    ampGainDegLow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ampGainDegLow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ampGainDegLow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ampGainDegLow.EntityData.Children = types.NewOrderedMap()
    ampGainDegLow.EntityData.Leafs = types.NewOrderedMap()
    ampGainDegLow.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", ampGainDegLow.IsDetected})
    ampGainDegLow.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", ampGainDegLow.Counter})

    ampGainDegLow.EntityData.YListKeys = []string {}

    return &(ampGainDegLow.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh
// Ampli Gain Deg High
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetEntityData() *types.CommonEntityData {
    ampGainDegHigh.EntityData.YFilter = ampGainDegHigh.YFilter
    ampGainDegHigh.EntityData.YangName = "amp-gain-deg-high"
    ampGainDegHigh.EntityData.BundleName = "cisco_ios_xr"
    ampGainDegHigh.EntityData.ParentYangName = "ots-alarm-info"
    ampGainDegHigh.EntityData.SegmentPath = "amp-gain-deg-high"
    ampGainDegHigh.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + ampGainDegHigh.EntityData.SegmentPath
    ampGainDegHigh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ampGainDegHigh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ampGainDegHigh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ampGainDegHigh.EntityData.Children = types.NewOrderedMap()
    ampGainDegHigh.EntityData.Leafs = types.NewOrderedMap()
    ampGainDegHigh.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", ampGainDegHigh.IsDetected})
    ampGainDegHigh.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", ampGainDegHigh.Counter})

    ampGainDegHigh.EntityData.YListKeys = []string {}

    return &(ampGainDegHigh.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut
// Auto Laser Shut
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetEntityData() *types.CommonEntityData {
    autoLaserShut.EntityData.YFilter = autoLaserShut.YFilter
    autoLaserShut.EntityData.YangName = "auto-laser-shut"
    autoLaserShut.EntityData.BundleName = "cisco_ios_xr"
    autoLaserShut.EntityData.ParentYangName = "ots-alarm-info"
    autoLaserShut.EntityData.SegmentPath = "auto-laser-shut"
    autoLaserShut.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + autoLaserShut.EntityData.SegmentPath
    autoLaserShut.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoLaserShut.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoLaserShut.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoLaserShut.EntityData.Children = types.NewOrderedMap()
    autoLaserShut.EntityData.Leafs = types.NewOrderedMap()
    autoLaserShut.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", autoLaserShut.IsDetected})
    autoLaserShut.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", autoLaserShut.Counter})

    autoLaserShut.EntityData.YListKeys = []string {}

    return &(autoLaserShut.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed
// Auto Power Red
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetEntityData() *types.CommonEntityData {
    autoPowerRed.EntityData.YFilter = autoPowerRed.YFilter
    autoPowerRed.EntityData.YangName = "auto-power-red"
    autoPowerRed.EntityData.BundleName = "cisco_ios_xr"
    autoPowerRed.EntityData.ParentYangName = "ots-alarm-info"
    autoPowerRed.EntityData.SegmentPath = "auto-power-red"
    autoPowerRed.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + autoPowerRed.EntityData.SegmentPath
    autoPowerRed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoPowerRed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoPowerRed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoPowerRed.EntityData.Children = types.NewOrderedMap()
    autoPowerRed.EntityData.Leafs = types.NewOrderedMap()
    autoPowerRed.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", autoPowerRed.IsDetected})
    autoPowerRed.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", autoPowerRed.Counter})

    autoPowerRed.EntityData.YListKeys = []string {}

    return &(autoPowerRed.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled
// Auto Ampli Ctrl Disabled
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetEntityData() *types.CommonEntityData {
    autoAmpliCtrlDisabled.EntityData.YFilter = autoAmpliCtrlDisabled.YFilter
    autoAmpliCtrlDisabled.EntityData.YangName = "auto-ampli-ctrl-disabled"
    autoAmpliCtrlDisabled.EntityData.BundleName = "cisco_ios_xr"
    autoAmpliCtrlDisabled.EntityData.ParentYangName = "ots-alarm-info"
    autoAmpliCtrlDisabled.EntityData.SegmentPath = "auto-ampli-ctrl-disabled"
    autoAmpliCtrlDisabled.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + autoAmpliCtrlDisabled.EntityData.SegmentPath
    autoAmpliCtrlDisabled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoAmpliCtrlDisabled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoAmpliCtrlDisabled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoAmpliCtrlDisabled.EntityData.Children = types.NewOrderedMap()
    autoAmpliCtrlDisabled.EntityData.Leafs = types.NewOrderedMap()
    autoAmpliCtrlDisabled.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", autoAmpliCtrlDisabled.IsDetected})
    autoAmpliCtrlDisabled.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", autoAmpliCtrlDisabled.Counter})

    autoAmpliCtrlDisabled.EntityData.YListKeys = []string {}

    return &(autoAmpliCtrlDisabled.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch
// Auto Ampli Ctrl Config Mismatch
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetEntityData() *types.CommonEntityData {
    autoAmpliCtrlConfigMismatch.EntityData.YFilter = autoAmpliCtrlConfigMismatch.YFilter
    autoAmpliCtrlConfigMismatch.EntityData.YangName = "auto-ampli-ctrl-config-mismatch"
    autoAmpliCtrlConfigMismatch.EntityData.BundleName = "cisco_ios_xr"
    autoAmpliCtrlConfigMismatch.EntityData.ParentYangName = "ots-alarm-info"
    autoAmpliCtrlConfigMismatch.EntityData.SegmentPath = "auto-ampli-ctrl-config-mismatch"
    autoAmpliCtrlConfigMismatch.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + autoAmpliCtrlConfigMismatch.EntityData.SegmentPath
    autoAmpliCtrlConfigMismatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoAmpliCtrlConfigMismatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoAmpliCtrlConfigMismatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoAmpliCtrlConfigMismatch.EntityData.Children = types.NewOrderedMap()
    autoAmpliCtrlConfigMismatch.EntityData.Leafs = types.NewOrderedMap()
    autoAmpliCtrlConfigMismatch.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", autoAmpliCtrlConfigMismatch.IsDetected})
    autoAmpliCtrlConfigMismatch.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", autoAmpliCtrlConfigMismatch.Counter})

    autoAmpliCtrlConfigMismatch.EntityData.YListKeys = []string {}

    return &(autoAmpliCtrlConfigMismatch.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect
// Switch To Protect
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetEntityData() *types.CommonEntityData {
    switchToProtect.EntityData.YFilter = switchToProtect.YFilter
    switchToProtect.EntityData.YangName = "switch-to-protect"
    switchToProtect.EntityData.BundleName = "cisco_ios_xr"
    switchToProtect.EntityData.ParentYangName = "ots-alarm-info"
    switchToProtect.EntityData.SegmentPath = "switch-to-protect"
    switchToProtect.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + switchToProtect.EntityData.SegmentPath
    switchToProtect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchToProtect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchToProtect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchToProtect.EntityData.Children = types.NewOrderedMap()
    switchToProtect.EntityData.Leafs = types.NewOrderedMap()
    switchToProtect.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", switchToProtect.IsDetected})
    switchToProtect.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", switchToProtect.Counter})

    switchToProtect.EntityData.YListKeys = []string {}

    return &(switchToProtect.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning
// Auto Ampli Ctrl Running
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetEntityData() *types.CommonEntityData {
    autoAmpliCtrlRunning.EntityData.YFilter = autoAmpliCtrlRunning.YFilter
    autoAmpliCtrlRunning.EntityData.YangName = "auto-ampli-ctrl-running"
    autoAmpliCtrlRunning.EntityData.BundleName = "cisco_ios_xr"
    autoAmpliCtrlRunning.EntityData.ParentYangName = "ots-alarm-info"
    autoAmpliCtrlRunning.EntityData.SegmentPath = "auto-ampli-ctrl-running"
    autoAmpliCtrlRunning.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/ots-alarm-info/" + autoAmpliCtrlRunning.EntityData.SegmentPath
    autoAmpliCtrlRunning.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoAmpliCtrlRunning.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoAmpliCtrlRunning.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoAmpliCtrlRunning.EntityData.Children = types.NewOrderedMap()
    autoAmpliCtrlRunning.EntityData.Leafs = types.NewOrderedMap()
    autoAmpliCtrlRunning.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", autoAmpliCtrlRunning.IsDetected})
    autoAmpliCtrlRunning.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", autoAmpliCtrlRunning.Counter})

    autoAmpliCtrlRunning.EntityData.YListKeys = []string {}

    return &(autoAmpliCtrlRunning.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo
// Transceiver Vendor Details
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vendor Information. The type is string.
    VendorInfo interface{}

    // Adapter Vendor Information. The type is string.
    AdapterVendorInfo interface{}

    // Date in Transceiver. The type is string.
    Date interface{}

    // Transceiver vendors revision number. The type is string.
    OpticsVendorRev interface{}

    // Transceiver serial number. The type is string.
    OpticsSerialNo interface{}

    // Transceiver vendors part number. The type is string.
    OpticsVendorPart interface{}

    // Transceiver optics type. The type is string.
    OpticsType interface{}

    // Transceiver optics vendor name. The type is string.
    VendorName interface{}

    // Transceiver optics type. The type is string.
    OuiNo interface{}

    // Transceiver optics pid. The type is string.
    OpticsPid interface{}

    // Transceiver optics vid. The type is string.
    OpticsVid interface{}

    // Connector type. The type is FiberConnector.
    ConnectorType interface{}

    // Otn Application Code. The type is OtnApplicationCode.
    OtnApplicationCode interface{}

    // Sonet Application Code. The type is SonetApplicationCode.
    SonetApplicationCode interface{}

    // Ethernet Compliance Code. The type is EthernetPmd.
    EthernetComplianceCode interface{}

    // Internal Temperature in units of C. The type is interface{} with range:
    // -2147483648..2147483647.
    InternalTemperature interface{}
}

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetEntityData() *types.CommonEntityData {
    transceiverInfo.EntityData.YFilter = transceiverInfo.YFilter
    transceiverInfo.EntityData.YangName = "transceiver-info"
    transceiverInfo.EntityData.BundleName = "cisco_ios_xr"
    transceiverInfo.EntityData.ParentYangName = "optics-info"
    transceiverInfo.EntityData.SegmentPath = "transceiver-info"
    transceiverInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + transceiverInfo.EntityData.SegmentPath
    transceiverInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transceiverInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transceiverInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transceiverInfo.EntityData.Children = types.NewOrderedMap()
    transceiverInfo.EntityData.Leafs = types.NewOrderedMap()
    transceiverInfo.EntityData.Leafs.Append("vendor-info", types.YLeaf{"VendorInfo", transceiverInfo.VendorInfo})
    transceiverInfo.EntityData.Leafs.Append("adapter-vendor-info", types.YLeaf{"AdapterVendorInfo", transceiverInfo.AdapterVendorInfo})
    transceiverInfo.EntityData.Leafs.Append("date", types.YLeaf{"Date", transceiverInfo.Date})
    transceiverInfo.EntityData.Leafs.Append("optics-vendor-rev", types.YLeaf{"OpticsVendorRev", transceiverInfo.OpticsVendorRev})
    transceiverInfo.EntityData.Leafs.Append("optics-serial-no", types.YLeaf{"OpticsSerialNo", transceiverInfo.OpticsSerialNo})
    transceiverInfo.EntityData.Leafs.Append("optics-vendor-part", types.YLeaf{"OpticsVendorPart", transceiverInfo.OpticsVendorPart})
    transceiverInfo.EntityData.Leafs.Append("optics-type", types.YLeaf{"OpticsType", transceiverInfo.OpticsType})
    transceiverInfo.EntityData.Leafs.Append("vendor-name", types.YLeaf{"VendorName", transceiverInfo.VendorName})
    transceiverInfo.EntityData.Leafs.Append("oui-no", types.YLeaf{"OuiNo", transceiverInfo.OuiNo})
    transceiverInfo.EntityData.Leafs.Append("optics-pid", types.YLeaf{"OpticsPid", transceiverInfo.OpticsPid})
    transceiverInfo.EntityData.Leafs.Append("optics-vid", types.YLeaf{"OpticsVid", transceiverInfo.OpticsVid})
    transceiverInfo.EntityData.Leafs.Append("connector-type", types.YLeaf{"ConnectorType", transceiverInfo.ConnectorType})
    transceiverInfo.EntityData.Leafs.Append("otn-application-code", types.YLeaf{"OtnApplicationCode", transceiverInfo.OtnApplicationCode})
    transceiverInfo.EntityData.Leafs.Append("sonet-application-code", types.YLeaf{"SonetApplicationCode", transceiverInfo.SonetApplicationCode})
    transceiverInfo.EntityData.Leafs.Append("ethernet-compliance-code", types.YLeaf{"EthernetComplianceCode", transceiverInfo.EthernetComplianceCode})
    transceiverInfo.EntityData.Leafs.Append("internal-temperature", types.YLeaf{"InternalTemperature", transceiverInfo.InternalTemperature})

    transceiverInfo.EntityData.YListKeys = []string {}

    return &(transceiverInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal
// Extended optics parameters
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Signal to Noise Ratio on Lane 1. The type is interface{} with range:
    // -2147483648..2147483647.
    SnrLane1 interface{}

    // Signal to Noise Ratio on Lane 2. The type is interface{} with range:
    // -2147483648..2147483647.
    SnrLane2 interface{}

    // Inter symbol Interference correction on Lane 1. The type is interface{}
    // with range: -2147483648..2147483647.
    IsiCorrectionLane1 interface{}

    // Inter symbol Interference correction on Lane 2. The type is interface{}
    // with range: -2147483648..2147483647.
    IsiCorrectionLane2 interface{}

    // PAM Histogram parameter on Lane 1. The type is interface{} with range:
    // -2147483648..2147483647.
    PamRateLane1 interface{}

    // PAM Histogram parameter on Lane 2. The type is interface{} with range:
    // -2147483648..2147483647.
    PamRateLane2 interface{}

    // Pre FEC BER since last counter reset. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    PreFecBer interface{}

    // Uncorrected BER since last counter reset. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    UncorrectedBer interface{}

    // Current flowing to the TEC of a cooled laser on Lane 1. The type is
    // interface{} with range: -2147483648..2147483647.
    TecCurrentLane1 interface{}

    // Current flowing to the TEC of a cooled laser on Lane 2. The type is
    // interface{} with range: -2147483648..2147483647.
    TecCurrentLane2 interface{}

    // Difference between target and actual center frequency on Lane 1. The type
    // is interface{} with range: -2147483648..2147483647.
    LaserDiffFrequencyLane1 interface{}

    // Difference between target and actual center frequency on Lane 2. The type
    // is interface{} with range: -2147483648..2147483647.
    LaserDiffFrequencyLane2 interface{}

    // Difference between target and actual temperature on Lane 1. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffTemperatureLane1 interface{}

    // Difference between target and actual temperature on Lane 2. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffTemperatureLane2 interface{}

    // Latched minimum Pre FEC BER value since last read, line ingress. The type
    // is interface{} with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMin interface{}

    // Latched maximum Pre FEC BER value since last read, line ingress. The type
    // is interface{} with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMax interface{}

    // Pre FEC BER value prior accumulation period, line ingress. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    PreFecBerAccumulated interface{}

    // Pre FEC BER value instantaneous line ingress. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    PreFecBerInstantaneous interface{}

    // Latched minimum Uncorrected BER value since last read, line ingress. The
    // type is interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMin interface{}

    // Latched maximum Uncorrected BER value since last read, line ingress. The
    // type is interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMax interface{}

    // Uncorrected BER value prior accumulation period, line ingress. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerAccumulated interface{}

    // Uncorrected BER value instantaneous line line ingress. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerInstantaneous interface{}
}

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetEntityData() *types.CommonEntityData {
    extParamVal.EntityData.YFilter = extParamVal.YFilter
    extParamVal.EntityData.YangName = "ext-param-val"
    extParamVal.EntityData.BundleName = "cisco_ios_xr"
    extParamVal.EntityData.ParentYangName = "optics-info"
    extParamVal.EntityData.SegmentPath = "ext-param-val"
    extParamVal.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + extParamVal.EntityData.SegmentPath
    extParamVal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extParamVal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extParamVal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extParamVal.EntityData.Children = types.NewOrderedMap()
    extParamVal.EntityData.Leafs = types.NewOrderedMap()
    extParamVal.EntityData.Leafs.Append("snr-lane1", types.YLeaf{"SnrLane1", extParamVal.SnrLane1})
    extParamVal.EntityData.Leafs.Append("snr-lane2", types.YLeaf{"SnrLane2", extParamVal.SnrLane2})
    extParamVal.EntityData.Leafs.Append("isi-correction-lane1", types.YLeaf{"IsiCorrectionLane1", extParamVal.IsiCorrectionLane1})
    extParamVal.EntityData.Leafs.Append("isi-correction-lane2", types.YLeaf{"IsiCorrectionLane2", extParamVal.IsiCorrectionLane2})
    extParamVal.EntityData.Leafs.Append("pam-rate-lane1", types.YLeaf{"PamRateLane1", extParamVal.PamRateLane1})
    extParamVal.EntityData.Leafs.Append("pam-rate-lane2", types.YLeaf{"PamRateLane2", extParamVal.PamRateLane2})
    extParamVal.EntityData.Leafs.Append("pre-fec-ber", types.YLeaf{"PreFecBer", extParamVal.PreFecBer})
    extParamVal.EntityData.Leafs.Append("uncorrected-ber", types.YLeaf{"UncorrectedBer", extParamVal.UncorrectedBer})
    extParamVal.EntityData.Leafs.Append("tec-current-lane1", types.YLeaf{"TecCurrentLane1", extParamVal.TecCurrentLane1})
    extParamVal.EntityData.Leafs.Append("tec-current-lane2", types.YLeaf{"TecCurrentLane2", extParamVal.TecCurrentLane2})
    extParamVal.EntityData.Leafs.Append("laser-diff-frequency-lane1", types.YLeaf{"LaserDiffFrequencyLane1", extParamVal.LaserDiffFrequencyLane1})
    extParamVal.EntityData.Leafs.Append("laser-diff-frequency-lane2", types.YLeaf{"LaserDiffFrequencyLane2", extParamVal.LaserDiffFrequencyLane2})
    extParamVal.EntityData.Leafs.Append("laser-diff-temperature-lane1", types.YLeaf{"LaserDiffTemperatureLane1", extParamVal.LaserDiffTemperatureLane1})
    extParamVal.EntityData.Leafs.Append("laser-diff-temperature-lane2", types.YLeaf{"LaserDiffTemperatureLane2", extParamVal.LaserDiffTemperatureLane2})
    extParamVal.EntityData.Leafs.Append("pre-fec-ber-latched-min", types.YLeaf{"PreFecBerLatchedMin", extParamVal.PreFecBerLatchedMin})
    extParamVal.EntityData.Leafs.Append("pre-fec-ber-latched-max", types.YLeaf{"PreFecBerLatchedMax", extParamVal.PreFecBerLatchedMax})
    extParamVal.EntityData.Leafs.Append("pre-fec-ber-accumulated", types.YLeaf{"PreFecBerAccumulated", extParamVal.PreFecBerAccumulated})
    extParamVal.EntityData.Leafs.Append("pre-fec-ber-instantaneous", types.YLeaf{"PreFecBerInstantaneous", extParamVal.PreFecBerInstantaneous})
    extParamVal.EntityData.Leafs.Append("uncorrected-ber-latched-min", types.YLeaf{"UncorrectedBerLatchedMin", extParamVal.UncorrectedBerLatchedMin})
    extParamVal.EntityData.Leafs.Append("uncorrected-ber-latched-max", types.YLeaf{"UncorrectedBerLatchedMax", extParamVal.UncorrectedBerLatchedMax})
    extParamVal.EntityData.Leafs.Append("uncorrected-ber-accumulated", types.YLeaf{"UncorrectedBerAccumulated", extParamVal.UncorrectedBerAccumulated})
    extParamVal.EntityData.Leafs.Append("uncorrected-ber-instantaneous", types.YLeaf{"UncorrectedBerInstantaneous", extParamVal.UncorrectedBerInstantaneous})

    extParamVal.EntityData.YListKeys = []string {}

    return &(extParamVal.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal
// Extended optics parameters threshold values
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // High threshold alarm for SNR. The type is interface{} with range:
    // -2147483648..2147483647.
    SnrAlarmHighThreshold interface{}

    // Low threshold alarm for SNR. The type is interface{} with range:
    // -2147483648..2147483647.
    SnrAlarmLowThreshold interface{}

    // High threshold warning for SNR. The type is interface{} with range:
    // -2147483648..2147483647.
    SnrWarnHighThreshold interface{}

    // Low threshold warning for SNR. The type is interface{} with range:
    // -2147483648..2147483647.
    SnrWarnLowThreshold interface{}

    // High threshold alarm for ISI Correction. The type is interface{} with
    // range: -2147483648..2147483647.
    IsiCorrectionAlarmHighThreshold interface{}

    // Low threshold alarm for ISI Correction. The type is interface{} with range:
    // -2147483648..2147483647.
    IsiCorrectionAlarmLowThreshold interface{}

    // High threshold warning for ISI Correction. The type is interface{} with
    // range: -2147483648..2147483647.
    IsiCorrectionWarnHighThreshold interface{}

    // Low threshold warning for ISI Correction. The type is interface{} with
    // range: -2147483648..2147483647.
    IsiCorrectionWarnLowThreshold interface{}

    // High threshold alarm for PAM Rate. The type is interface{} with range:
    // -2147483648..2147483647.
    PamRateAlarmHighThreshold interface{}

    // Low threshold alarm for PAM Rate. The type is interface{} with range:
    // -2147483648..2147483647.
    PamRateAlarmLowThreshold interface{}

    // High threshold warning for PAM Rate. The type is interface{} with range:
    // -2147483648..2147483647.
    PamRateWarnHighThreshold interface{}

    // Low threshold warning for PAM Rate. The type is interface{} with range:
    // -2147483648..2147483647.
    PamRateWarnLowThreshold interface{}

    // High threshold alarm for Pre FEC BER. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    PreFecBerAlarmHighThreshold interface{}

    // Low threshold alarm for Pre FEC BER. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    PreFecBerAlarmLowThreshold interface{}

    // High threshold warning for Pre FEC BER. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    PreFecBerWarnHighThreshold interface{}

    // Low threshold warning for Pre FEC BER. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    PreFecBerWarnLowThreshold interface{}

    // High threshold alarm for Uncorrected BER. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    UncorrectedBerAlarmHighThreshold interface{}

    // Low threshold alarm for Uncorrected BER. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    UncorrectedBerAlarmLowThreshold interface{}

    // High threshold warning for Uncorrected BER. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    UncorrectedBerWarnHighThreshold interface{}

    // Low threshold warning for Uncorrected BER. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    UncorrectedBerWarnLowThreshold interface{}

    // High threshold alarm for TEC Current. The type is interface{} with range:
    // -2147483648..2147483647.
    TecCurrentAlarmHighThreshold interface{}

    // Low threshold alarm for TEC Current. The type is interface{} with range:
    // -2147483648..2147483647.
    TecCurrentAlarmLowThreshold interface{}

    // High threshold warning for TEC Current. The type is interface{} with range:
    // -2147483648..2147483647.
    TecCurrentWarnHighThreshold interface{}

    // Low threshold warning for TEC Current. The type is interface{} with range:
    // -2147483648..2147483647.
    TecCurrentWarnLowThreshold interface{}

    // High Threshold Alarm for Differential Laser Frequency. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffFrequencyAlarmHighThreshold interface{}

    // Low Threshold Alarm for Differential Laser Frequency. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffFrequencyAlarmLowThreshold interface{}

    // High Threshold Warning for Differential Laser Frequency. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffFrequencyWarnHighThreshold interface{}

    // Low Threshold Warning for Differential Laser Frequency. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffFrequencyWarnLowThreshold interface{}

    // High Threshold Alarm for Differential Laser Temperature. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffTemperatureAlarmHighThreshold interface{}

    // Low Threshold Alarm for Differential Laser Temperature. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffTemperatureAlarmLowThreshold interface{}

    // High Threshold Warning for Differential Laser Temperature. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffTemperatureWarnHighThreshold interface{}

    // Low Threshold Warning for Differential Laser Temperature. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserDiffTemperatureWarnLowThreshold interface{}

    // High threshold alarm for Latched Min Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMinAlarmHighThreshold interface{}

    // Low threshold alarm for Latched Min Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMinAlarmLowThreshold interface{}

    // High threshold warning for Latched Min Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMinWarnHighThreshold interface{}

    // Low threshold warning for Latched Min Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMinWarnLowThreshold interface{}

    // High threshold alarm for Latched Max Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMaxAlarmHighThreshold interface{}

    // Low threshold alarm for Latched Max Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMaxAlarmLowThreshold interface{}

    // High threshold warning for Latched Max Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMaxWarnHighThreshold interface{}

    // Low threshold warning for Latched Max Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerLatchedMaxWarnLowThreshold interface{}

    // High threshold alarm for Accumulated Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerAccumulatedAlarmHighThreshold interface{}

    // Low threshold alarm for Accumulated Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerAccumulatedAlarmLowThreshold interface{}

    // High threshold warning for Accumulated Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerAccumulatedWarnHighThreshold interface{}

    // Low threshold warning for Accumulated Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerAccumulatedWarnLowThreshold interface{}

    // High threshold alarm for Instantaneous Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerInstantaneousAlarmHighThreshold interface{}

    // Low threshold alarm for Instantaneous Pre FEC BER. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    PreFecBerInstantaneousAlarmLowThreshold interface{}

    // High threshold warning for Instantaneous Pre FEC BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    PreFecBerInstantaneousWarnHighThreshold interface{}

    // Low threshold warning for Instantaneous Pre FEC BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    PreFecBerInstantaneousWarnLowThreshold interface{}

    // High threshold alarm for  Latched Min Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMinAlarmHighThreshold interface{}

    // Low threshold alarm for  Latched Min Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMinAlarmLowThreshold interface{}

    // High threshold warning for  Latched Min Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMinWarnHighThreshold interface{}

    // Low threshold alarm for Latched Min Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMinWarnLowThreshold interface{}

    // High threshold alarm for Latched_Max Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMaxAlarmHighThreshold interface{}

    // Low threshold alarm for Latched_Max Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMaxAlarmLowThreshold interface{}

    // High threshold warning Latched_Max for Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMaxWarnHighThreshold interface{}

    // Low threshold warning Latched_Max for Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerLatchedMaxWarnLowThreshold interface{}

    // High threshold alarm for Accumulated Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerAccumulatedAlarmHighThreshold interface{}

    // Low threshold alarm for Accumulated Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerAccumulatedAlarmLowThreshold interface{}

    // High threshold warning for Accumulated Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerAccumulatedWarnHighThreshold interface{}

    // Low threshold warning for Accumulated Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerAccumulatedWarnLowThreshold interface{}

    // High threshold alarm for Instantaneous Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerInstantaneousAlarmHighThreshold interface{}

    // Low threshold alarm for Instantaneous Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerInstantaneousAlarmLowThreshold interface{}

    // High threshold warning for Instantaneous Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerInstantaneousWarnHighThreshold interface{}

    // Low threshold warning for Instantaneous Uncorrected BER. The type is
    // interface{} with range: -9223372036854775808..9223372036854775807.
    UncorrectedBerInstantaneousWarnLowThreshold interface{}
}

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetEntityData() *types.CommonEntityData {
    extParamThresholdVal.EntityData.YFilter = extParamThresholdVal.YFilter
    extParamThresholdVal.EntityData.YangName = "ext-param-threshold-val"
    extParamThresholdVal.EntityData.BundleName = "cisco_ios_xr"
    extParamThresholdVal.EntityData.ParentYangName = "optics-info"
    extParamThresholdVal.EntityData.SegmentPath = "ext-param-threshold-val"
    extParamThresholdVal.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + extParamThresholdVal.EntityData.SegmentPath
    extParamThresholdVal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extParamThresholdVal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extParamThresholdVal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extParamThresholdVal.EntityData.Children = types.NewOrderedMap()
    extParamThresholdVal.EntityData.Leafs = types.NewOrderedMap()
    extParamThresholdVal.EntityData.Leafs.Append("snr-alarm-high-threshold", types.YLeaf{"SnrAlarmHighThreshold", extParamThresholdVal.SnrAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("snr-alarm-low-threshold", types.YLeaf{"SnrAlarmLowThreshold", extParamThresholdVal.SnrAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("snr-warn-high-threshold", types.YLeaf{"SnrWarnHighThreshold", extParamThresholdVal.SnrWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("snr-warn-low-threshold", types.YLeaf{"SnrWarnLowThreshold", extParamThresholdVal.SnrWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("isi-correction-alarm-high-threshold", types.YLeaf{"IsiCorrectionAlarmHighThreshold", extParamThresholdVal.IsiCorrectionAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("isi-correction-alarm-low-threshold", types.YLeaf{"IsiCorrectionAlarmLowThreshold", extParamThresholdVal.IsiCorrectionAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("isi-correction-warn-high-threshold", types.YLeaf{"IsiCorrectionWarnHighThreshold", extParamThresholdVal.IsiCorrectionWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("isi-correction-warn-low-threshold", types.YLeaf{"IsiCorrectionWarnLowThreshold", extParamThresholdVal.IsiCorrectionWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pam-rate-alarm-high-threshold", types.YLeaf{"PamRateAlarmHighThreshold", extParamThresholdVal.PamRateAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pam-rate-alarm-low-threshold", types.YLeaf{"PamRateAlarmLowThreshold", extParamThresholdVal.PamRateAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pam-rate-warn-high-threshold", types.YLeaf{"PamRateWarnHighThreshold", extParamThresholdVal.PamRateWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pam-rate-warn-low-threshold", types.YLeaf{"PamRateWarnLowThreshold", extParamThresholdVal.PamRateWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-alarm-high-threshold", types.YLeaf{"PreFecBerAlarmHighThreshold", extParamThresholdVal.PreFecBerAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-alarm-low-threshold", types.YLeaf{"PreFecBerAlarmLowThreshold", extParamThresholdVal.PreFecBerAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-warn-high-threshold", types.YLeaf{"PreFecBerWarnHighThreshold", extParamThresholdVal.PreFecBerWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-warn-low-threshold", types.YLeaf{"PreFecBerWarnLowThreshold", extParamThresholdVal.PreFecBerWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-alarm-high-threshold", types.YLeaf{"UncorrectedBerAlarmHighThreshold", extParamThresholdVal.UncorrectedBerAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-alarm-low-threshold", types.YLeaf{"UncorrectedBerAlarmLowThreshold", extParamThresholdVal.UncorrectedBerAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-warn-high-threshold", types.YLeaf{"UncorrectedBerWarnHighThreshold", extParamThresholdVal.UncorrectedBerWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-warn-low-threshold", types.YLeaf{"UncorrectedBerWarnLowThreshold", extParamThresholdVal.UncorrectedBerWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("tec-current-alarm-high-threshold", types.YLeaf{"TecCurrentAlarmHighThreshold", extParamThresholdVal.TecCurrentAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("tec-current-alarm-low-threshold", types.YLeaf{"TecCurrentAlarmLowThreshold", extParamThresholdVal.TecCurrentAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("tec-current-warn-high-threshold", types.YLeaf{"TecCurrentWarnHighThreshold", extParamThresholdVal.TecCurrentWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("tec-current-warn-low-threshold", types.YLeaf{"TecCurrentWarnLowThreshold", extParamThresholdVal.TecCurrentWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("laser-diff-frequency-alarm-high-threshold", types.YLeaf{"LaserDiffFrequencyAlarmHighThreshold", extParamThresholdVal.LaserDiffFrequencyAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("laser-diff-frequency-alarm-low-threshold", types.YLeaf{"LaserDiffFrequencyAlarmLowThreshold", extParamThresholdVal.LaserDiffFrequencyAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("laser-diff-frequency-warn-high-threshold", types.YLeaf{"LaserDiffFrequencyWarnHighThreshold", extParamThresholdVal.LaserDiffFrequencyWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("laser-diff-frequency-warn-low-threshold", types.YLeaf{"LaserDiffFrequencyWarnLowThreshold", extParamThresholdVal.LaserDiffFrequencyWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("laser-diff-temperature-alarm-high-threshold", types.YLeaf{"LaserDiffTemperatureAlarmHighThreshold", extParamThresholdVal.LaserDiffTemperatureAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("laser-diff-temperature-alarm-low-threshold", types.YLeaf{"LaserDiffTemperatureAlarmLowThreshold", extParamThresholdVal.LaserDiffTemperatureAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("laser-diff-temperature-warn-high-threshold", types.YLeaf{"LaserDiffTemperatureWarnHighThreshold", extParamThresholdVal.LaserDiffTemperatureWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("laser-diff-temperature-warn-low-threshold", types.YLeaf{"LaserDiffTemperatureWarnLowThreshold", extParamThresholdVal.LaserDiffTemperatureWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-latched-min-alarm-high-threshold", types.YLeaf{"PreFecBerLatchedMinAlarmHighThreshold", extParamThresholdVal.PreFecBerLatchedMinAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-latched-min-alarm-low-threshold", types.YLeaf{"PreFecBerLatchedMinAlarmLowThreshold", extParamThresholdVal.PreFecBerLatchedMinAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-latched-min-warn-high-threshold", types.YLeaf{"PreFecBerLatchedMinWarnHighThreshold", extParamThresholdVal.PreFecBerLatchedMinWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-latched-min-warn-low-threshold", types.YLeaf{"PreFecBerLatchedMinWarnLowThreshold", extParamThresholdVal.PreFecBerLatchedMinWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-latched-max-alarm-high-threshold", types.YLeaf{"PreFecBerLatchedMaxAlarmHighThreshold", extParamThresholdVal.PreFecBerLatchedMaxAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-latched-max-alarm-low-threshold", types.YLeaf{"PreFecBerLatchedMaxAlarmLowThreshold", extParamThresholdVal.PreFecBerLatchedMaxAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-latched-max-warn-high-threshold", types.YLeaf{"PreFecBerLatchedMaxWarnHighThreshold", extParamThresholdVal.PreFecBerLatchedMaxWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-latched-max-warn-low-threshold", types.YLeaf{"PreFecBerLatchedMaxWarnLowThreshold", extParamThresholdVal.PreFecBerLatchedMaxWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-accumulated-alarm-high-threshold", types.YLeaf{"PreFecBerAccumulatedAlarmHighThreshold", extParamThresholdVal.PreFecBerAccumulatedAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-accumulated-alarm-low-threshold", types.YLeaf{"PreFecBerAccumulatedAlarmLowThreshold", extParamThresholdVal.PreFecBerAccumulatedAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-accumulated-warn-high-threshold", types.YLeaf{"PreFecBerAccumulatedWarnHighThreshold", extParamThresholdVal.PreFecBerAccumulatedWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-accumulated-warn-low-threshold", types.YLeaf{"PreFecBerAccumulatedWarnLowThreshold", extParamThresholdVal.PreFecBerAccumulatedWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-instantaneous-alarm-high-threshold", types.YLeaf{"PreFecBerInstantaneousAlarmHighThreshold", extParamThresholdVal.PreFecBerInstantaneousAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-instantaneous-alarm-low-threshold", types.YLeaf{"PreFecBerInstantaneousAlarmLowThreshold", extParamThresholdVal.PreFecBerInstantaneousAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-instantaneous-warn-high-threshold", types.YLeaf{"PreFecBerInstantaneousWarnHighThreshold", extParamThresholdVal.PreFecBerInstantaneousWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("pre-fec-ber-instantaneous-warn-low-threshold", types.YLeaf{"PreFecBerInstantaneousWarnLowThreshold", extParamThresholdVal.PreFecBerInstantaneousWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-latched-min-alarm-high-threshold", types.YLeaf{"UncorrectedBerLatchedMinAlarmHighThreshold", extParamThresholdVal.UncorrectedBerLatchedMinAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-latched-min-alarm-low-threshold", types.YLeaf{"UncorrectedBerLatchedMinAlarmLowThreshold", extParamThresholdVal.UncorrectedBerLatchedMinAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-latched-min-warn-high-threshold", types.YLeaf{"UncorrectedBerLatchedMinWarnHighThreshold", extParamThresholdVal.UncorrectedBerLatchedMinWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-latched-min-warn-low-threshold", types.YLeaf{"UncorrectedBerLatchedMinWarnLowThreshold", extParamThresholdVal.UncorrectedBerLatchedMinWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-latched-max-alarm-high-threshold", types.YLeaf{"UncorrectedBerLatchedMaxAlarmHighThreshold", extParamThresholdVal.UncorrectedBerLatchedMaxAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-latched-max-alarm-low-threshold", types.YLeaf{"UncorrectedBerLatchedMaxAlarmLowThreshold", extParamThresholdVal.UncorrectedBerLatchedMaxAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-latched-max-warn-high-threshold", types.YLeaf{"UncorrectedBerLatchedMaxWarnHighThreshold", extParamThresholdVal.UncorrectedBerLatchedMaxWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-latched-max-warn-low-threshold", types.YLeaf{"UncorrectedBerLatchedMaxWarnLowThreshold", extParamThresholdVal.UncorrectedBerLatchedMaxWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-accumulated-alarm-high-threshold", types.YLeaf{"UncorrectedBerAccumulatedAlarmHighThreshold", extParamThresholdVal.UncorrectedBerAccumulatedAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-accumulated-alarm-low-threshold", types.YLeaf{"UncorrectedBerAccumulatedAlarmLowThreshold", extParamThresholdVal.UncorrectedBerAccumulatedAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-accumulated-warn-high-threshold", types.YLeaf{"UncorrectedBerAccumulatedWarnHighThreshold", extParamThresholdVal.UncorrectedBerAccumulatedWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-accumulated-warn-low-threshold", types.YLeaf{"UncorrectedBerAccumulatedWarnLowThreshold", extParamThresholdVal.UncorrectedBerAccumulatedWarnLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-instantaneous-alarm-high-threshold", types.YLeaf{"UncorrectedBerInstantaneousAlarmHighThreshold", extParamThresholdVal.UncorrectedBerInstantaneousAlarmHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-instantaneous-alarm-low-threshold", types.YLeaf{"UncorrectedBerInstantaneousAlarmLowThreshold", extParamThresholdVal.UncorrectedBerInstantaneousAlarmLowThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-instantaneous-warn-high-threshold", types.YLeaf{"UncorrectedBerInstantaneousWarnHighThreshold", extParamThresholdVal.UncorrectedBerInstantaneousWarnHighThreshold})
    extParamThresholdVal.EntityData.Leafs.Append("uncorrected-ber-instantaneous-warn-low-threshold", types.YLeaf{"UncorrectedBerInstantaneousWarnLowThreshold", extParamThresholdVal.UncorrectedBerInstantaneousWarnLowThreshold})

    extParamThresholdVal.EntityData.YListKeys = []string {}

    return &(extParamThresholdVal.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo
// Extended DOM alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Low SNR Alarm for Lane1.
    LoSnr OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoSnr

    // High SNR Alarm for Lane1.
    HiSnr1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiSnr1

    // Low SNR Alarm for Lane2.
    LoSnr1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoSnr1

    // High SNR Alarm for Lane2.
    HiSnr2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiSnr2

    // Low ISI Correction Alarm for Lane1.
    LoIsi1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoIsi1

    // High ISI Correction Alarm for Lane1.
    HiIsi1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiIsi1

    // Low ISI Correction Alarm for Lane2.
    LoIsi2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoIsi2

    // High ISI Correction Alarm for Lane2.
    HiIsi2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiIsi2

    // Low PAM Rate Alarm for Lane1.
    LoPam1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoPam1

    // High PAM Rate Alarm for Lane1.
    HiPam1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPam1

    // Low PAM Rate Alarm for Lane2.
    LoPam2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoPam2

    // High PAM Rate Alarm for Lane2.
    HiPam2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPam2

    // Low TEC Current Alarm for Lane1.
    LoTec1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoTec1

    // High TEC Current Alarm for Lane1.
    HiTec1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiTec1

    // Low TEC Current Alarm for Lane2.
    LoTec2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoTec2

    // High TEC Current Alarm for Lane2.
    HiTec2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiTec2

    // Low Laser Differential Frequency Alarm for Lane1.
    LoLaserFreq1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoLaserFreq1

    // High Laser Differential Frequency Alarm for Lane1.
    HiLaserFreq1 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiLaserFreq1

    // Low Laser Differential Frequency Alarm for Lane2.
    LoLaserFreq2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoLaserFreq2

    // High Laser Differential Frequency Alarm for Lane2.
    HiLaserFreq2 OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiLaserFreq2

    // High Pre FEC BER Current Accumulated Alarm.
    HiPreFecberCurAcc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberCurAcc

    // High Pre FEC BER Min Alarm.
    HiPreFecberMin OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberMin

    // High Pre FEC BER Max Alarm.
    HiPreFecberMax OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberMax

    // High Pre FEC BER Prior Accumulated Alarm.
    HiPreFecberPriorAcc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberPriorAcc

    // High Pre FEC BER Current Alarm.
    HiPreFecberCur OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberCur

    // High Uncorrected BER Current Accumulated Alarm.
    HiUncorrectedBerCurAcc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerCurAcc

    // High Uncorrected BER Min Alarm.
    HiUncorrectedBerMin OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerMin

    // High Uncorrected BER Max Alarm.
    HiUncorrectedBerMax OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerMax

    // High Uncorrected BER Prior Accumulated Alarm.
    HiUncorrectedBerPriorAcc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerPriorAcc

    // High Uncorrected BER Current Alarm.
    HiUncorrectedBerCur OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerCur
}

func (extendedAlarmAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo) GetEntityData() *types.CommonEntityData {
    extendedAlarmAlarmInfo.EntityData.YFilter = extendedAlarmAlarmInfo.YFilter
    extendedAlarmAlarmInfo.EntityData.YangName = "extended-alarm-alarm-info"
    extendedAlarmAlarmInfo.EntityData.BundleName = "cisco_ios_xr"
    extendedAlarmAlarmInfo.EntityData.ParentYangName = "optics-info"
    extendedAlarmAlarmInfo.EntityData.SegmentPath = "extended-alarm-alarm-info"
    extendedAlarmAlarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + extendedAlarmAlarmInfo.EntityData.SegmentPath
    extendedAlarmAlarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedAlarmAlarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedAlarmAlarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedAlarmAlarmInfo.EntityData.Children = types.NewOrderedMap()
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-snr", types.YChild{"LoSnr", &extendedAlarmAlarmInfo.LoSnr})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-snr1", types.YChild{"HiSnr1", &extendedAlarmAlarmInfo.HiSnr1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-snr1", types.YChild{"LoSnr1", &extendedAlarmAlarmInfo.LoSnr1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-snr2", types.YChild{"HiSnr2", &extendedAlarmAlarmInfo.HiSnr2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-isi1", types.YChild{"LoIsi1", &extendedAlarmAlarmInfo.LoIsi1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-isi1", types.YChild{"HiIsi1", &extendedAlarmAlarmInfo.HiIsi1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-isi2", types.YChild{"LoIsi2", &extendedAlarmAlarmInfo.LoIsi2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-isi2", types.YChild{"HiIsi2", &extendedAlarmAlarmInfo.HiIsi2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-pam1", types.YChild{"LoPam1", &extendedAlarmAlarmInfo.LoPam1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-pam1", types.YChild{"HiPam1", &extendedAlarmAlarmInfo.HiPam1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-pam2", types.YChild{"LoPam2", &extendedAlarmAlarmInfo.LoPam2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-pam2", types.YChild{"HiPam2", &extendedAlarmAlarmInfo.HiPam2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-tec1", types.YChild{"LoTec1", &extendedAlarmAlarmInfo.LoTec1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-tec1", types.YChild{"HiTec1", &extendedAlarmAlarmInfo.HiTec1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-tec2", types.YChild{"LoTec2", &extendedAlarmAlarmInfo.LoTec2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-tec2", types.YChild{"HiTec2", &extendedAlarmAlarmInfo.HiTec2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-laser-freq1", types.YChild{"LoLaserFreq1", &extendedAlarmAlarmInfo.LoLaserFreq1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-laser-freq1", types.YChild{"HiLaserFreq1", &extendedAlarmAlarmInfo.HiLaserFreq1})
    extendedAlarmAlarmInfo.EntityData.Children.Append("lo-laser-freq2", types.YChild{"LoLaserFreq2", &extendedAlarmAlarmInfo.LoLaserFreq2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-laser-freq2", types.YChild{"HiLaserFreq2", &extendedAlarmAlarmInfo.HiLaserFreq2})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-pre-fecber-cur-acc", types.YChild{"HiPreFecberCurAcc", &extendedAlarmAlarmInfo.HiPreFecberCurAcc})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-pre-fecber-min", types.YChild{"HiPreFecberMin", &extendedAlarmAlarmInfo.HiPreFecberMin})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-pre-fecber-max", types.YChild{"HiPreFecberMax", &extendedAlarmAlarmInfo.HiPreFecberMax})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-pre-fecber-prior-acc", types.YChild{"HiPreFecberPriorAcc", &extendedAlarmAlarmInfo.HiPreFecberPriorAcc})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-pre-fecber-cur", types.YChild{"HiPreFecberCur", &extendedAlarmAlarmInfo.HiPreFecberCur})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-uncorrected-ber-cur-acc", types.YChild{"HiUncorrectedBerCurAcc", &extendedAlarmAlarmInfo.HiUncorrectedBerCurAcc})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-uncorrected-ber-min", types.YChild{"HiUncorrectedBerMin", &extendedAlarmAlarmInfo.HiUncorrectedBerMin})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-uncorrected-ber-max", types.YChild{"HiUncorrectedBerMax", &extendedAlarmAlarmInfo.HiUncorrectedBerMax})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-uncorrected-ber-prior-acc", types.YChild{"HiUncorrectedBerPriorAcc", &extendedAlarmAlarmInfo.HiUncorrectedBerPriorAcc})
    extendedAlarmAlarmInfo.EntityData.Children.Append("hi-uncorrected-ber-cur", types.YChild{"HiUncorrectedBerCur", &extendedAlarmAlarmInfo.HiUncorrectedBerCur})
    extendedAlarmAlarmInfo.EntityData.Leafs = types.NewOrderedMap()

    extendedAlarmAlarmInfo.EntityData.YListKeys = []string {}

    return &(extendedAlarmAlarmInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoSnr
// Low SNR Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoSnr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loSnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoSnr) GetEntityData() *types.CommonEntityData {
    loSnr.EntityData.YFilter = loSnr.YFilter
    loSnr.EntityData.YangName = "lo-snr"
    loSnr.EntityData.BundleName = "cisco_ios_xr"
    loSnr.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loSnr.EntityData.SegmentPath = "lo-snr"
    loSnr.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loSnr.EntityData.SegmentPath
    loSnr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loSnr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loSnr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loSnr.EntityData.Children = types.NewOrderedMap()
    loSnr.EntityData.Leafs = types.NewOrderedMap()
    loSnr.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loSnr.IsDetected})
    loSnr.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loSnr.Counter})

    loSnr.EntityData.YListKeys = []string {}

    return &(loSnr.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiSnr1
// High SNR Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiSnr1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiSnr1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiSnr1) GetEntityData() *types.CommonEntityData {
    hiSnr1.EntityData.YFilter = hiSnr1.YFilter
    hiSnr1.EntityData.YangName = "hi-snr1"
    hiSnr1.EntityData.BundleName = "cisco_ios_xr"
    hiSnr1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiSnr1.EntityData.SegmentPath = "hi-snr1"
    hiSnr1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiSnr1.EntityData.SegmentPath
    hiSnr1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiSnr1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiSnr1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiSnr1.EntityData.Children = types.NewOrderedMap()
    hiSnr1.EntityData.Leafs = types.NewOrderedMap()
    hiSnr1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiSnr1.IsDetected})
    hiSnr1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiSnr1.Counter})

    hiSnr1.EntityData.YListKeys = []string {}

    return &(hiSnr1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoSnr1
// Low SNR Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoSnr1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loSnr1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoSnr1) GetEntityData() *types.CommonEntityData {
    loSnr1.EntityData.YFilter = loSnr1.YFilter
    loSnr1.EntityData.YangName = "lo-snr1"
    loSnr1.EntityData.BundleName = "cisco_ios_xr"
    loSnr1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loSnr1.EntityData.SegmentPath = "lo-snr1"
    loSnr1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loSnr1.EntityData.SegmentPath
    loSnr1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loSnr1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loSnr1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loSnr1.EntityData.Children = types.NewOrderedMap()
    loSnr1.EntityData.Leafs = types.NewOrderedMap()
    loSnr1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loSnr1.IsDetected})
    loSnr1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loSnr1.Counter})

    loSnr1.EntityData.YListKeys = []string {}

    return &(loSnr1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiSnr2
// High SNR Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiSnr2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiSnr2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiSnr2) GetEntityData() *types.CommonEntityData {
    hiSnr2.EntityData.YFilter = hiSnr2.YFilter
    hiSnr2.EntityData.YangName = "hi-snr2"
    hiSnr2.EntityData.BundleName = "cisco_ios_xr"
    hiSnr2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiSnr2.EntityData.SegmentPath = "hi-snr2"
    hiSnr2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiSnr2.EntityData.SegmentPath
    hiSnr2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiSnr2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiSnr2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiSnr2.EntityData.Children = types.NewOrderedMap()
    hiSnr2.EntityData.Leafs = types.NewOrderedMap()
    hiSnr2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiSnr2.IsDetected})
    hiSnr2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiSnr2.Counter})

    hiSnr2.EntityData.YListKeys = []string {}

    return &(hiSnr2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoIsi1
// Low ISI Correction Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoIsi1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loIsi1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoIsi1) GetEntityData() *types.CommonEntityData {
    loIsi1.EntityData.YFilter = loIsi1.YFilter
    loIsi1.EntityData.YangName = "lo-isi1"
    loIsi1.EntityData.BundleName = "cisco_ios_xr"
    loIsi1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loIsi1.EntityData.SegmentPath = "lo-isi1"
    loIsi1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loIsi1.EntityData.SegmentPath
    loIsi1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loIsi1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loIsi1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loIsi1.EntityData.Children = types.NewOrderedMap()
    loIsi1.EntityData.Leafs = types.NewOrderedMap()
    loIsi1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loIsi1.IsDetected})
    loIsi1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loIsi1.Counter})

    loIsi1.EntityData.YListKeys = []string {}

    return &(loIsi1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiIsi1
// High ISI Correction Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiIsi1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiIsi1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiIsi1) GetEntityData() *types.CommonEntityData {
    hiIsi1.EntityData.YFilter = hiIsi1.YFilter
    hiIsi1.EntityData.YangName = "hi-isi1"
    hiIsi1.EntityData.BundleName = "cisco_ios_xr"
    hiIsi1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiIsi1.EntityData.SegmentPath = "hi-isi1"
    hiIsi1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiIsi1.EntityData.SegmentPath
    hiIsi1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiIsi1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiIsi1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiIsi1.EntityData.Children = types.NewOrderedMap()
    hiIsi1.EntityData.Leafs = types.NewOrderedMap()
    hiIsi1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiIsi1.IsDetected})
    hiIsi1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiIsi1.Counter})

    hiIsi1.EntityData.YListKeys = []string {}

    return &(hiIsi1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoIsi2
// Low ISI Correction Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoIsi2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loIsi2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoIsi2) GetEntityData() *types.CommonEntityData {
    loIsi2.EntityData.YFilter = loIsi2.YFilter
    loIsi2.EntityData.YangName = "lo-isi2"
    loIsi2.EntityData.BundleName = "cisco_ios_xr"
    loIsi2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loIsi2.EntityData.SegmentPath = "lo-isi2"
    loIsi2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loIsi2.EntityData.SegmentPath
    loIsi2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loIsi2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loIsi2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loIsi2.EntityData.Children = types.NewOrderedMap()
    loIsi2.EntityData.Leafs = types.NewOrderedMap()
    loIsi2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loIsi2.IsDetected})
    loIsi2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loIsi2.Counter})

    loIsi2.EntityData.YListKeys = []string {}

    return &(loIsi2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiIsi2
// High ISI Correction Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiIsi2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiIsi2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiIsi2) GetEntityData() *types.CommonEntityData {
    hiIsi2.EntityData.YFilter = hiIsi2.YFilter
    hiIsi2.EntityData.YangName = "hi-isi2"
    hiIsi2.EntityData.BundleName = "cisco_ios_xr"
    hiIsi2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiIsi2.EntityData.SegmentPath = "hi-isi2"
    hiIsi2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiIsi2.EntityData.SegmentPath
    hiIsi2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiIsi2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiIsi2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiIsi2.EntityData.Children = types.NewOrderedMap()
    hiIsi2.EntityData.Leafs = types.NewOrderedMap()
    hiIsi2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiIsi2.IsDetected})
    hiIsi2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiIsi2.Counter})

    hiIsi2.EntityData.YListKeys = []string {}

    return &(hiIsi2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoPam1
// Low PAM Rate Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoPam1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loPam1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoPam1) GetEntityData() *types.CommonEntityData {
    loPam1.EntityData.YFilter = loPam1.YFilter
    loPam1.EntityData.YangName = "lo-pam1"
    loPam1.EntityData.BundleName = "cisco_ios_xr"
    loPam1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loPam1.EntityData.SegmentPath = "lo-pam1"
    loPam1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loPam1.EntityData.SegmentPath
    loPam1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loPam1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loPam1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loPam1.EntityData.Children = types.NewOrderedMap()
    loPam1.EntityData.Leafs = types.NewOrderedMap()
    loPam1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loPam1.IsDetected})
    loPam1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loPam1.Counter})

    loPam1.EntityData.YListKeys = []string {}

    return &(loPam1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPam1
// High PAM Rate Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPam1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiPam1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPam1) GetEntityData() *types.CommonEntityData {
    hiPam1.EntityData.YFilter = hiPam1.YFilter
    hiPam1.EntityData.YangName = "hi-pam1"
    hiPam1.EntityData.BundleName = "cisco_ios_xr"
    hiPam1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiPam1.EntityData.SegmentPath = "hi-pam1"
    hiPam1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiPam1.EntityData.SegmentPath
    hiPam1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiPam1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiPam1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiPam1.EntityData.Children = types.NewOrderedMap()
    hiPam1.EntityData.Leafs = types.NewOrderedMap()
    hiPam1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiPam1.IsDetected})
    hiPam1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiPam1.Counter})

    hiPam1.EntityData.YListKeys = []string {}

    return &(hiPam1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoPam2
// Low PAM Rate Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoPam2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loPam2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoPam2) GetEntityData() *types.CommonEntityData {
    loPam2.EntityData.YFilter = loPam2.YFilter
    loPam2.EntityData.YangName = "lo-pam2"
    loPam2.EntityData.BundleName = "cisco_ios_xr"
    loPam2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loPam2.EntityData.SegmentPath = "lo-pam2"
    loPam2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loPam2.EntityData.SegmentPath
    loPam2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loPam2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loPam2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loPam2.EntityData.Children = types.NewOrderedMap()
    loPam2.EntityData.Leafs = types.NewOrderedMap()
    loPam2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loPam2.IsDetected})
    loPam2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loPam2.Counter})

    loPam2.EntityData.YListKeys = []string {}

    return &(loPam2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPam2
// High PAM Rate Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPam2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiPam2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPam2) GetEntityData() *types.CommonEntityData {
    hiPam2.EntityData.YFilter = hiPam2.YFilter
    hiPam2.EntityData.YangName = "hi-pam2"
    hiPam2.EntityData.BundleName = "cisco_ios_xr"
    hiPam2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiPam2.EntityData.SegmentPath = "hi-pam2"
    hiPam2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiPam2.EntityData.SegmentPath
    hiPam2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiPam2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiPam2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiPam2.EntityData.Children = types.NewOrderedMap()
    hiPam2.EntityData.Leafs = types.NewOrderedMap()
    hiPam2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiPam2.IsDetected})
    hiPam2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiPam2.Counter})

    hiPam2.EntityData.YListKeys = []string {}

    return &(hiPam2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoTec1
// Low TEC Current Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoTec1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loTec1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoTec1) GetEntityData() *types.CommonEntityData {
    loTec1.EntityData.YFilter = loTec1.YFilter
    loTec1.EntityData.YangName = "lo-tec1"
    loTec1.EntityData.BundleName = "cisco_ios_xr"
    loTec1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loTec1.EntityData.SegmentPath = "lo-tec1"
    loTec1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loTec1.EntityData.SegmentPath
    loTec1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loTec1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loTec1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loTec1.EntityData.Children = types.NewOrderedMap()
    loTec1.EntityData.Leafs = types.NewOrderedMap()
    loTec1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loTec1.IsDetected})
    loTec1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loTec1.Counter})

    loTec1.EntityData.YListKeys = []string {}

    return &(loTec1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiTec1
// High TEC Current Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiTec1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiTec1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiTec1) GetEntityData() *types.CommonEntityData {
    hiTec1.EntityData.YFilter = hiTec1.YFilter
    hiTec1.EntityData.YangName = "hi-tec1"
    hiTec1.EntityData.BundleName = "cisco_ios_xr"
    hiTec1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiTec1.EntityData.SegmentPath = "hi-tec1"
    hiTec1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiTec1.EntityData.SegmentPath
    hiTec1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiTec1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiTec1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiTec1.EntityData.Children = types.NewOrderedMap()
    hiTec1.EntityData.Leafs = types.NewOrderedMap()
    hiTec1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiTec1.IsDetected})
    hiTec1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiTec1.Counter})

    hiTec1.EntityData.YListKeys = []string {}

    return &(hiTec1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoTec2
// Low TEC Current Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoTec2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loTec2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoTec2) GetEntityData() *types.CommonEntityData {
    loTec2.EntityData.YFilter = loTec2.YFilter
    loTec2.EntityData.YangName = "lo-tec2"
    loTec2.EntityData.BundleName = "cisco_ios_xr"
    loTec2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loTec2.EntityData.SegmentPath = "lo-tec2"
    loTec2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loTec2.EntityData.SegmentPath
    loTec2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loTec2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loTec2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loTec2.EntityData.Children = types.NewOrderedMap()
    loTec2.EntityData.Leafs = types.NewOrderedMap()
    loTec2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loTec2.IsDetected})
    loTec2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loTec2.Counter})

    loTec2.EntityData.YListKeys = []string {}

    return &(loTec2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiTec2
// High TEC Current Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiTec2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiTec2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiTec2) GetEntityData() *types.CommonEntityData {
    hiTec2.EntityData.YFilter = hiTec2.YFilter
    hiTec2.EntityData.YangName = "hi-tec2"
    hiTec2.EntityData.BundleName = "cisco_ios_xr"
    hiTec2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiTec2.EntityData.SegmentPath = "hi-tec2"
    hiTec2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiTec2.EntityData.SegmentPath
    hiTec2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiTec2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiTec2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiTec2.EntityData.Children = types.NewOrderedMap()
    hiTec2.EntityData.Leafs = types.NewOrderedMap()
    hiTec2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiTec2.IsDetected})
    hiTec2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiTec2.Counter})

    hiTec2.EntityData.YListKeys = []string {}

    return &(hiTec2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoLaserFreq1
// Low Laser Differential Frequency Alarm for Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoLaserFreq1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loLaserFreq1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoLaserFreq1) GetEntityData() *types.CommonEntityData {
    loLaserFreq1.EntityData.YFilter = loLaserFreq1.YFilter
    loLaserFreq1.EntityData.YangName = "lo-laser-freq1"
    loLaserFreq1.EntityData.BundleName = "cisco_ios_xr"
    loLaserFreq1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loLaserFreq1.EntityData.SegmentPath = "lo-laser-freq1"
    loLaserFreq1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loLaserFreq1.EntityData.SegmentPath
    loLaserFreq1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loLaserFreq1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loLaserFreq1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loLaserFreq1.EntityData.Children = types.NewOrderedMap()
    loLaserFreq1.EntityData.Leafs = types.NewOrderedMap()
    loLaserFreq1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loLaserFreq1.IsDetected})
    loLaserFreq1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loLaserFreq1.Counter})

    loLaserFreq1.EntityData.YListKeys = []string {}

    return &(loLaserFreq1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiLaserFreq1
// High Laser Differential Frequency Alarm for
// Lane1
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiLaserFreq1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiLaserFreq1 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiLaserFreq1) GetEntityData() *types.CommonEntityData {
    hiLaserFreq1.EntityData.YFilter = hiLaserFreq1.YFilter
    hiLaserFreq1.EntityData.YangName = "hi-laser-freq1"
    hiLaserFreq1.EntityData.BundleName = "cisco_ios_xr"
    hiLaserFreq1.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiLaserFreq1.EntityData.SegmentPath = "hi-laser-freq1"
    hiLaserFreq1.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiLaserFreq1.EntityData.SegmentPath
    hiLaserFreq1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiLaserFreq1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiLaserFreq1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiLaserFreq1.EntityData.Children = types.NewOrderedMap()
    hiLaserFreq1.EntityData.Leafs = types.NewOrderedMap()
    hiLaserFreq1.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiLaserFreq1.IsDetected})
    hiLaserFreq1.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiLaserFreq1.Counter})

    hiLaserFreq1.EntityData.YListKeys = []string {}

    return &(hiLaserFreq1.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoLaserFreq2
// Low Laser Differential Frequency Alarm for Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoLaserFreq2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (loLaserFreq2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_LoLaserFreq2) GetEntityData() *types.CommonEntityData {
    loLaserFreq2.EntityData.YFilter = loLaserFreq2.YFilter
    loLaserFreq2.EntityData.YangName = "lo-laser-freq2"
    loLaserFreq2.EntityData.BundleName = "cisco_ios_xr"
    loLaserFreq2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    loLaserFreq2.EntityData.SegmentPath = "lo-laser-freq2"
    loLaserFreq2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + loLaserFreq2.EntityData.SegmentPath
    loLaserFreq2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loLaserFreq2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loLaserFreq2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loLaserFreq2.EntityData.Children = types.NewOrderedMap()
    loLaserFreq2.EntityData.Leafs = types.NewOrderedMap()
    loLaserFreq2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", loLaserFreq2.IsDetected})
    loLaserFreq2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", loLaserFreq2.Counter})

    loLaserFreq2.EntityData.YListKeys = []string {}

    return &(loLaserFreq2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiLaserFreq2
// High Laser Differential Frequency Alarm for
// Lane2
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiLaserFreq2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiLaserFreq2 *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiLaserFreq2) GetEntityData() *types.CommonEntityData {
    hiLaserFreq2.EntityData.YFilter = hiLaserFreq2.YFilter
    hiLaserFreq2.EntityData.YangName = "hi-laser-freq2"
    hiLaserFreq2.EntityData.BundleName = "cisco_ios_xr"
    hiLaserFreq2.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiLaserFreq2.EntityData.SegmentPath = "hi-laser-freq2"
    hiLaserFreq2.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiLaserFreq2.EntityData.SegmentPath
    hiLaserFreq2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiLaserFreq2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiLaserFreq2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiLaserFreq2.EntityData.Children = types.NewOrderedMap()
    hiLaserFreq2.EntityData.Leafs = types.NewOrderedMap()
    hiLaserFreq2.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiLaserFreq2.IsDetected})
    hiLaserFreq2.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiLaserFreq2.Counter})

    hiLaserFreq2.EntityData.YListKeys = []string {}

    return &(hiLaserFreq2.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberCurAcc
// High Pre FEC BER Current Accumulated Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberCurAcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiPreFecberCurAcc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberCurAcc) GetEntityData() *types.CommonEntityData {
    hiPreFecberCurAcc.EntityData.YFilter = hiPreFecberCurAcc.YFilter
    hiPreFecberCurAcc.EntityData.YangName = "hi-pre-fecber-cur-acc"
    hiPreFecberCurAcc.EntityData.BundleName = "cisco_ios_xr"
    hiPreFecberCurAcc.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiPreFecberCurAcc.EntityData.SegmentPath = "hi-pre-fecber-cur-acc"
    hiPreFecberCurAcc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiPreFecberCurAcc.EntityData.SegmentPath
    hiPreFecberCurAcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiPreFecberCurAcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiPreFecberCurAcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiPreFecberCurAcc.EntityData.Children = types.NewOrderedMap()
    hiPreFecberCurAcc.EntityData.Leafs = types.NewOrderedMap()
    hiPreFecberCurAcc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiPreFecberCurAcc.IsDetected})
    hiPreFecberCurAcc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiPreFecberCurAcc.Counter})

    hiPreFecberCurAcc.EntityData.YListKeys = []string {}

    return &(hiPreFecberCurAcc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberMin
// High Pre FEC BER Min Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberMin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiPreFecberMin *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberMin) GetEntityData() *types.CommonEntityData {
    hiPreFecberMin.EntityData.YFilter = hiPreFecberMin.YFilter
    hiPreFecberMin.EntityData.YangName = "hi-pre-fecber-min"
    hiPreFecberMin.EntityData.BundleName = "cisco_ios_xr"
    hiPreFecberMin.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiPreFecberMin.EntityData.SegmentPath = "hi-pre-fecber-min"
    hiPreFecberMin.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiPreFecberMin.EntityData.SegmentPath
    hiPreFecberMin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiPreFecberMin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiPreFecberMin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiPreFecberMin.EntityData.Children = types.NewOrderedMap()
    hiPreFecberMin.EntityData.Leafs = types.NewOrderedMap()
    hiPreFecberMin.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiPreFecberMin.IsDetected})
    hiPreFecberMin.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiPreFecberMin.Counter})

    hiPreFecberMin.EntityData.YListKeys = []string {}

    return &(hiPreFecberMin.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberMax
// High Pre FEC BER Max Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberMax struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiPreFecberMax *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberMax) GetEntityData() *types.CommonEntityData {
    hiPreFecberMax.EntityData.YFilter = hiPreFecberMax.YFilter
    hiPreFecberMax.EntityData.YangName = "hi-pre-fecber-max"
    hiPreFecberMax.EntityData.BundleName = "cisco_ios_xr"
    hiPreFecberMax.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiPreFecberMax.EntityData.SegmentPath = "hi-pre-fecber-max"
    hiPreFecberMax.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiPreFecberMax.EntityData.SegmentPath
    hiPreFecberMax.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiPreFecberMax.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiPreFecberMax.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiPreFecberMax.EntityData.Children = types.NewOrderedMap()
    hiPreFecberMax.EntityData.Leafs = types.NewOrderedMap()
    hiPreFecberMax.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiPreFecberMax.IsDetected})
    hiPreFecberMax.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiPreFecberMax.Counter})

    hiPreFecberMax.EntityData.YListKeys = []string {}

    return &(hiPreFecberMax.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberPriorAcc
// High Pre FEC BER Prior Accumulated Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberPriorAcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiPreFecberPriorAcc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberPriorAcc) GetEntityData() *types.CommonEntityData {
    hiPreFecberPriorAcc.EntityData.YFilter = hiPreFecberPriorAcc.YFilter
    hiPreFecberPriorAcc.EntityData.YangName = "hi-pre-fecber-prior-acc"
    hiPreFecberPriorAcc.EntityData.BundleName = "cisco_ios_xr"
    hiPreFecberPriorAcc.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiPreFecberPriorAcc.EntityData.SegmentPath = "hi-pre-fecber-prior-acc"
    hiPreFecberPriorAcc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiPreFecberPriorAcc.EntityData.SegmentPath
    hiPreFecberPriorAcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiPreFecberPriorAcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiPreFecberPriorAcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiPreFecberPriorAcc.EntityData.Children = types.NewOrderedMap()
    hiPreFecberPriorAcc.EntityData.Leafs = types.NewOrderedMap()
    hiPreFecberPriorAcc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiPreFecberPriorAcc.IsDetected})
    hiPreFecberPriorAcc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiPreFecberPriorAcc.Counter})

    hiPreFecberPriorAcc.EntityData.YListKeys = []string {}

    return &(hiPreFecberPriorAcc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberCur
// High Pre FEC BER Current Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberCur struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiPreFecberCur *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiPreFecberCur) GetEntityData() *types.CommonEntityData {
    hiPreFecberCur.EntityData.YFilter = hiPreFecberCur.YFilter
    hiPreFecberCur.EntityData.YangName = "hi-pre-fecber-cur"
    hiPreFecberCur.EntityData.BundleName = "cisco_ios_xr"
    hiPreFecberCur.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiPreFecberCur.EntityData.SegmentPath = "hi-pre-fecber-cur"
    hiPreFecberCur.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiPreFecberCur.EntityData.SegmentPath
    hiPreFecberCur.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiPreFecberCur.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiPreFecberCur.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiPreFecberCur.EntityData.Children = types.NewOrderedMap()
    hiPreFecberCur.EntityData.Leafs = types.NewOrderedMap()
    hiPreFecberCur.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiPreFecberCur.IsDetected})
    hiPreFecberCur.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiPreFecberCur.Counter})

    hiPreFecberCur.EntityData.YListKeys = []string {}

    return &(hiPreFecberCur.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerCurAcc
// High Uncorrected BER Current Accumulated Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerCurAcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiUncorrectedBerCurAcc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerCurAcc) GetEntityData() *types.CommonEntityData {
    hiUncorrectedBerCurAcc.EntityData.YFilter = hiUncorrectedBerCurAcc.YFilter
    hiUncorrectedBerCurAcc.EntityData.YangName = "hi-uncorrected-ber-cur-acc"
    hiUncorrectedBerCurAcc.EntityData.BundleName = "cisco_ios_xr"
    hiUncorrectedBerCurAcc.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiUncorrectedBerCurAcc.EntityData.SegmentPath = "hi-uncorrected-ber-cur-acc"
    hiUncorrectedBerCurAcc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiUncorrectedBerCurAcc.EntityData.SegmentPath
    hiUncorrectedBerCurAcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiUncorrectedBerCurAcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiUncorrectedBerCurAcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiUncorrectedBerCurAcc.EntityData.Children = types.NewOrderedMap()
    hiUncorrectedBerCurAcc.EntityData.Leafs = types.NewOrderedMap()
    hiUncorrectedBerCurAcc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiUncorrectedBerCurAcc.IsDetected})
    hiUncorrectedBerCurAcc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiUncorrectedBerCurAcc.Counter})

    hiUncorrectedBerCurAcc.EntityData.YListKeys = []string {}

    return &(hiUncorrectedBerCurAcc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerMin
// High Uncorrected BER Min Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerMin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiUncorrectedBerMin *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerMin) GetEntityData() *types.CommonEntityData {
    hiUncorrectedBerMin.EntityData.YFilter = hiUncorrectedBerMin.YFilter
    hiUncorrectedBerMin.EntityData.YangName = "hi-uncorrected-ber-min"
    hiUncorrectedBerMin.EntityData.BundleName = "cisco_ios_xr"
    hiUncorrectedBerMin.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiUncorrectedBerMin.EntityData.SegmentPath = "hi-uncorrected-ber-min"
    hiUncorrectedBerMin.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiUncorrectedBerMin.EntityData.SegmentPath
    hiUncorrectedBerMin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiUncorrectedBerMin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiUncorrectedBerMin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiUncorrectedBerMin.EntityData.Children = types.NewOrderedMap()
    hiUncorrectedBerMin.EntityData.Leafs = types.NewOrderedMap()
    hiUncorrectedBerMin.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiUncorrectedBerMin.IsDetected})
    hiUncorrectedBerMin.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiUncorrectedBerMin.Counter})

    hiUncorrectedBerMin.EntityData.YListKeys = []string {}

    return &(hiUncorrectedBerMin.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerMax
// High Uncorrected BER Max Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerMax struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiUncorrectedBerMax *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerMax) GetEntityData() *types.CommonEntityData {
    hiUncorrectedBerMax.EntityData.YFilter = hiUncorrectedBerMax.YFilter
    hiUncorrectedBerMax.EntityData.YangName = "hi-uncorrected-ber-max"
    hiUncorrectedBerMax.EntityData.BundleName = "cisco_ios_xr"
    hiUncorrectedBerMax.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiUncorrectedBerMax.EntityData.SegmentPath = "hi-uncorrected-ber-max"
    hiUncorrectedBerMax.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiUncorrectedBerMax.EntityData.SegmentPath
    hiUncorrectedBerMax.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiUncorrectedBerMax.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiUncorrectedBerMax.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiUncorrectedBerMax.EntityData.Children = types.NewOrderedMap()
    hiUncorrectedBerMax.EntityData.Leafs = types.NewOrderedMap()
    hiUncorrectedBerMax.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiUncorrectedBerMax.IsDetected})
    hiUncorrectedBerMax.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiUncorrectedBerMax.Counter})

    hiUncorrectedBerMax.EntityData.YListKeys = []string {}

    return &(hiUncorrectedBerMax.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerPriorAcc
// High Uncorrected BER Prior Accumulated Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerPriorAcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiUncorrectedBerPriorAcc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerPriorAcc) GetEntityData() *types.CommonEntityData {
    hiUncorrectedBerPriorAcc.EntityData.YFilter = hiUncorrectedBerPriorAcc.YFilter
    hiUncorrectedBerPriorAcc.EntityData.YangName = "hi-uncorrected-ber-prior-acc"
    hiUncorrectedBerPriorAcc.EntityData.BundleName = "cisco_ios_xr"
    hiUncorrectedBerPriorAcc.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiUncorrectedBerPriorAcc.EntityData.SegmentPath = "hi-uncorrected-ber-prior-acc"
    hiUncorrectedBerPriorAcc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiUncorrectedBerPriorAcc.EntityData.SegmentPath
    hiUncorrectedBerPriorAcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiUncorrectedBerPriorAcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiUncorrectedBerPriorAcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiUncorrectedBerPriorAcc.EntityData.Children = types.NewOrderedMap()
    hiUncorrectedBerPriorAcc.EntityData.Leafs = types.NewOrderedMap()
    hiUncorrectedBerPriorAcc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiUncorrectedBerPriorAcc.IsDetected})
    hiUncorrectedBerPriorAcc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiUncorrectedBerPriorAcc.Counter})

    hiUncorrectedBerPriorAcc.EntityData.YListKeys = []string {}

    return &(hiUncorrectedBerPriorAcc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerCur
// High Uncorrected BER Current Alarm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerCur struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hiUncorrectedBerCur *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtendedAlarmAlarmInfo_HiUncorrectedBerCur) GetEntityData() *types.CommonEntityData {
    hiUncorrectedBerCur.EntityData.YFilter = hiUncorrectedBerCur.YFilter
    hiUncorrectedBerCur.EntityData.YangName = "hi-uncorrected-ber-cur"
    hiUncorrectedBerCur.EntityData.BundleName = "cisco_ios_xr"
    hiUncorrectedBerCur.EntityData.ParentYangName = "extended-alarm-alarm-info"
    hiUncorrectedBerCur.EntityData.SegmentPath = "hi-uncorrected-ber-cur"
    hiUncorrectedBerCur.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/extended-alarm-alarm-info/" + hiUncorrectedBerCur.EntityData.SegmentPath
    hiUncorrectedBerCur.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hiUncorrectedBerCur.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hiUncorrectedBerCur.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hiUncorrectedBerCur.EntityData.Children = types.NewOrderedMap()
    hiUncorrectedBerCur.EntityData.Leafs = types.NewOrderedMap()
    hiUncorrectedBerCur.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", hiUncorrectedBerCur.IsDetected})
    hiUncorrectedBerCur.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", hiUncorrectedBerCur.Counter})

    hiUncorrectedBerCur.EntityData.YListKeys = []string {}

    return &(hiUncorrectedBerCur.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_AinsInfo
// AINS information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_AinsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AINS State. The type is OpticsAinsStateEt.
    AinsState interface{}

    // AINS Timer in Minutes. The type is interface{} with range: 0..4294967295.
    // Units are minute.
    AinsTimerMinutes interface{}

    // AINS Remaining Seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    AinsRemainingSecs interface{}
}

func (ainsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_AinsInfo) GetEntityData() *types.CommonEntityData {
    ainsInfo.EntityData.YFilter = ainsInfo.YFilter
    ainsInfo.EntityData.YangName = "ains-info"
    ainsInfo.EntityData.BundleName = "cisco_ios_xr"
    ainsInfo.EntityData.ParentYangName = "optics-info"
    ainsInfo.EntityData.SegmentPath = "ains-info"
    ainsInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + ainsInfo.EntityData.SegmentPath
    ainsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ainsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ainsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ainsInfo.EntityData.Children = types.NewOrderedMap()
    ainsInfo.EntityData.Leafs = types.NewOrderedMap()
    ainsInfo.EntityData.Leafs.Append("ains-state", types.YLeaf{"AinsState", ainsInfo.AinsState})
    ainsInfo.EntityData.Leafs.Append("ains-timer-minutes", types.YLeaf{"AinsTimerMinutes", ainsInfo.AinsTimerMinutes})
    ainsInfo.EntityData.Leafs.Append("ains-remaining-secs", types.YLeaf{"AinsRemainingSecs", ainsInfo.AinsRemainingSecs})

    ainsInfo.EntityData.YListKeys = []string {}

    return &(ainsInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData
// Lane information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The index number of the lane. The type is interface{} with range:
    // 0..4294967295.
    LaneIndex interface{}

    // Laser Bias Current in units of 0.01 percentage. The type is interface{}
    // with range: 0..4294967295. Units are percentage.
    LaserBiasCurrentPercent interface{}

    // Laser Bias Current in units of 0.01mA. The type is interface{} with range:
    // 0..4294967295.
    LaserBiasCurrentMilliAmps interface{}

    // Transmit power in the units of 0.01dBm. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPower interface{}

    // Transponder receive power in the units of 0 .01dBm. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePower interface{}

    // Transponder receive signal power in the units of 0.01dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    ReceiveSignalPower interface{}

    // Transmit Signal power in the units of 0.01dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    TransmitSignalPower interface{}

    // Output frequency read from hw in the units of 100MHz. The type is
    // interface{} with range: -2147483648..2147483647.
    OutputFrequency interface{}

    // Frequency Offset read from hw in units of MHz. The type is interface{} with
    // range: -2147483648..2147483647.
    FrequencyOffset interface{}

    // Lane Alarm Information.
    LaneAlarmInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo
}

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetEntityData() *types.CommonEntityData {
    laneData.EntityData.YFilter = laneData.YFilter
    laneData.EntityData.YangName = "lane-data"
    laneData.EntityData.BundleName = "cisco_ios_xr"
    laneData.EntityData.ParentYangName = "optics-info"
    laneData.EntityData.SegmentPath = "lane-data" + types.AddNoKeyToken(laneData)
    laneData.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/" + laneData.EntityData.SegmentPath
    laneData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    laneData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    laneData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    laneData.EntityData.Children = types.NewOrderedMap()
    laneData.EntityData.Children.Append("lane-alarm-info", types.YChild{"LaneAlarmInfo", &laneData.LaneAlarmInfo})
    laneData.EntityData.Leafs = types.NewOrderedMap()
    laneData.EntityData.Leafs.Append("lane-index", types.YLeaf{"LaneIndex", laneData.LaneIndex})
    laneData.EntityData.Leafs.Append("laser-bias-current-percent", types.YLeaf{"LaserBiasCurrentPercent", laneData.LaserBiasCurrentPercent})
    laneData.EntityData.Leafs.Append("laser-bias-current-milli-amps", types.YLeaf{"LaserBiasCurrentMilliAmps", laneData.LaserBiasCurrentMilliAmps})
    laneData.EntityData.Leafs.Append("transmit-power", types.YLeaf{"TransmitPower", laneData.TransmitPower})
    laneData.EntityData.Leafs.Append("receive-power", types.YLeaf{"ReceivePower", laneData.ReceivePower})
    laneData.EntityData.Leafs.Append("receive-signal-power", types.YLeaf{"ReceiveSignalPower", laneData.ReceiveSignalPower})
    laneData.EntityData.Leafs.Append("transmit-signal-power", types.YLeaf{"TransmitSignalPower", laneData.TransmitSignalPower})
    laneData.EntityData.Leafs.Append("output-frequency", types.YLeaf{"OutputFrequency", laneData.OutputFrequency})
    laneData.EntityData.Leafs.Append("frequency-offset", types.YLeaf{"FrequencyOffset", laneData.FrequencyOffset})

    laneData.EntityData.YListKeys = []string {}

    return &(laneData.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo
// Lane Alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // High Rx Power.
    HighRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower

    // Low Rx Power.
    LowRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower

    // High Tx Power.
    HighTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower

    // Low Tx Power.
    LowTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower

    // High laser bias current.
    HighLbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetEntityData() *types.CommonEntityData {
    laneAlarmInfo.EntityData.YFilter = laneAlarmInfo.YFilter
    laneAlarmInfo.EntityData.YangName = "lane-alarm-info"
    laneAlarmInfo.EntityData.BundleName = "cisco_ios_xr"
    laneAlarmInfo.EntityData.ParentYangName = "lane-data"
    laneAlarmInfo.EntityData.SegmentPath = "lane-alarm-info"
    laneAlarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/lane-data/" + laneAlarmInfo.EntityData.SegmentPath
    laneAlarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    laneAlarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    laneAlarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    laneAlarmInfo.EntityData.Children = types.NewOrderedMap()
    laneAlarmInfo.EntityData.Children.Append("high-rx-power", types.YChild{"HighRxPower", &laneAlarmInfo.HighRxPower})
    laneAlarmInfo.EntityData.Children.Append("low-rx-power", types.YChild{"LowRxPower", &laneAlarmInfo.LowRxPower})
    laneAlarmInfo.EntityData.Children.Append("high-tx-power", types.YChild{"HighTxPower", &laneAlarmInfo.HighTxPower})
    laneAlarmInfo.EntityData.Children.Append("low-tx-power", types.YChild{"LowTxPower", &laneAlarmInfo.LowTxPower})
    laneAlarmInfo.EntityData.Children.Append("high-lbc", types.YChild{"HighLbc", &laneAlarmInfo.HighLbc})
    laneAlarmInfo.EntityData.Leafs = types.NewOrderedMap()

    laneAlarmInfo.EntityData.YListKeys = []string {}

    return &(laneAlarmInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower
// High Rx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetEntityData() *types.CommonEntityData {
    highRxPower.EntityData.YFilter = highRxPower.YFilter
    highRxPower.EntityData.YangName = "high-rx-power"
    highRxPower.EntityData.BundleName = "cisco_ios_xr"
    highRxPower.EntityData.ParentYangName = "lane-alarm-info"
    highRxPower.EntityData.SegmentPath = "high-rx-power"
    highRxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/lane-data/lane-alarm-info/" + highRxPower.EntityData.SegmentPath
    highRxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highRxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highRxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highRxPower.EntityData.Children = types.NewOrderedMap()
    highRxPower.EntityData.Leafs = types.NewOrderedMap()
    highRxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highRxPower.IsDetected})
    highRxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highRxPower.Counter})

    highRxPower.EntityData.YListKeys = []string {}

    return &(highRxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower
// Low Rx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetEntityData() *types.CommonEntityData {
    lowRxPower.EntityData.YFilter = lowRxPower.YFilter
    lowRxPower.EntityData.YangName = "low-rx-power"
    lowRxPower.EntityData.BundleName = "cisco_ios_xr"
    lowRxPower.EntityData.ParentYangName = "lane-alarm-info"
    lowRxPower.EntityData.SegmentPath = "low-rx-power"
    lowRxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/lane-data/lane-alarm-info/" + lowRxPower.EntityData.SegmentPath
    lowRxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowRxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowRxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowRxPower.EntityData.Children = types.NewOrderedMap()
    lowRxPower.EntityData.Leafs = types.NewOrderedMap()
    lowRxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowRxPower.IsDetected})
    lowRxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowRxPower.Counter})

    lowRxPower.EntityData.YListKeys = []string {}

    return &(lowRxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower
// High Tx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetEntityData() *types.CommonEntityData {
    highTxPower.EntityData.YFilter = highTxPower.YFilter
    highTxPower.EntityData.YangName = "high-tx-power"
    highTxPower.EntityData.BundleName = "cisco_ios_xr"
    highTxPower.EntityData.ParentYangName = "lane-alarm-info"
    highTxPower.EntityData.SegmentPath = "high-tx-power"
    highTxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/lane-data/lane-alarm-info/" + highTxPower.EntityData.SegmentPath
    highTxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTxPower.EntityData.Children = types.NewOrderedMap()
    highTxPower.EntityData.Leafs = types.NewOrderedMap()
    highTxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTxPower.IsDetected})
    highTxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTxPower.Counter})

    highTxPower.EntityData.YListKeys = []string {}

    return &(highTxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower
// Low Tx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetEntityData() *types.CommonEntityData {
    lowTxPower.EntityData.YFilter = lowTxPower.YFilter
    lowTxPower.EntityData.YangName = "low-tx-power"
    lowTxPower.EntityData.BundleName = "cisco_ios_xr"
    lowTxPower.EntityData.ParentYangName = "lane-alarm-info"
    lowTxPower.EntityData.SegmentPath = "low-tx-power"
    lowTxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/lane-data/lane-alarm-info/" + lowTxPower.EntityData.SegmentPath
    lowTxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTxPower.EntityData.Children = types.NewOrderedMap()
    lowTxPower.EntityData.Leafs = types.NewOrderedMap()
    lowTxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTxPower.IsDetected})
    lowTxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTxPower.Counter})

    lowTxPower.EntityData.YListKeys = []string {}

    return &(lowTxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc
// High laser bias current
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetEntityData() *types.CommonEntityData {
    highLbc.EntityData.YFilter = highLbc.YFilter
    highLbc.EntityData.YangName = "high-lbc"
    highLbc.EntityData.BundleName = "cisco_ios_xr"
    highLbc.EntityData.ParentYangName = "lane-alarm-info"
    highLbc.EntityData.SegmentPath = "high-lbc"
    highLbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-info/lane-data/lane-alarm-info/" + highLbc.EntityData.SegmentPath
    highLbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highLbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highLbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highLbc.EntityData.Children = types.NewOrderedMap()
    highLbc.EntityData.Leafs = types.NewOrderedMap()
    highLbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highLbc.IsDetected})
    highLbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highLbc.Counter})

    highLbc.EntityData.YListKeys = []string {}

    return &(highLbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes
// All Optics Port operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lane Information. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane.
    OpticsLane []*OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane
}

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetEntityData() *types.CommonEntityData {
    opticsLanes.EntityData.YFilter = opticsLanes.YFilter
    opticsLanes.EntityData.YangName = "optics-lanes"
    opticsLanes.EntityData.BundleName = "cisco_ios_xr"
    opticsLanes.EntityData.ParentYangName = "optics-port"
    opticsLanes.EntityData.SegmentPath = "optics-lanes"
    opticsLanes.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/" + opticsLanes.EntityData.SegmentPath
    opticsLanes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsLanes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsLanes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsLanes.EntityData.Children = types.NewOrderedMap()
    opticsLanes.EntityData.Children.Append("optics-lane", types.YChild{"OpticsLane", nil})
    for i := range opticsLanes.OpticsLane {
        opticsLanes.EntityData.Children.Append(types.GetSegmentPath(opticsLanes.OpticsLane[i]), types.YChild{"OpticsLane", opticsLanes.OpticsLane[i]})
    }
    opticsLanes.EntityData.Leafs = types.NewOrderedMap()

    opticsLanes.EntityData.YListKeys = []string {}

    return &(opticsLanes.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane
// Lane Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Lane Index. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // The index number of the lane. The type is interface{} with range:
    // 0..4294967295.
    LaneIndex interface{}

    // Laser Bias Current in units of 0.01 percentage. The type is interface{}
    // with range: 0..4294967295. Units are percentage.
    LaserBiasCurrentPercent interface{}

    // Laser Bias Current in units of 0.01mA. The type is interface{} with range:
    // 0..4294967295.
    LaserBiasCurrentMilliAmps interface{}

    // Transmit power in the units of 0.01dBm. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPower interface{}

    // Transponder receive power in the units of 0 .01dBm. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePower interface{}

    // Transponder receive signal power in the units of 0.01dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    ReceiveSignalPower interface{}

    // Transmit Signal power in the units of 0.01dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    TransmitSignalPower interface{}

    // Output frequency read from hw in the units of 100MHz. The type is
    // interface{} with range: -2147483648..2147483647.
    OutputFrequency interface{}

    // Frequency Offset read from hw in units of MHz. The type is interface{} with
    // range: -2147483648..2147483647.
    FrequencyOffset interface{}

    // Lane Alarm Information.
    LaneAlarmInfo OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo
}

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetEntityData() *types.CommonEntityData {
    opticsLane.EntityData.YFilter = opticsLane.YFilter
    opticsLane.EntityData.YangName = "optics-lane"
    opticsLane.EntityData.BundleName = "cisco_ios_xr"
    opticsLane.EntityData.ParentYangName = "optics-lanes"
    opticsLane.EntityData.SegmentPath = "optics-lane" + types.AddKeyToken(opticsLane.Number, "number")
    opticsLane.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-lanes/" + opticsLane.EntityData.SegmentPath
    opticsLane.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsLane.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsLane.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsLane.EntityData.Children = types.NewOrderedMap()
    opticsLane.EntityData.Children.Append("lane-alarm-info", types.YChild{"LaneAlarmInfo", &opticsLane.LaneAlarmInfo})
    opticsLane.EntityData.Leafs = types.NewOrderedMap()
    opticsLane.EntityData.Leafs.Append("number", types.YLeaf{"Number", opticsLane.Number})
    opticsLane.EntityData.Leafs.Append("lane-index", types.YLeaf{"LaneIndex", opticsLane.LaneIndex})
    opticsLane.EntityData.Leafs.Append("laser-bias-current-percent", types.YLeaf{"LaserBiasCurrentPercent", opticsLane.LaserBiasCurrentPercent})
    opticsLane.EntityData.Leafs.Append("laser-bias-current-milli-amps", types.YLeaf{"LaserBiasCurrentMilliAmps", opticsLane.LaserBiasCurrentMilliAmps})
    opticsLane.EntityData.Leafs.Append("transmit-power", types.YLeaf{"TransmitPower", opticsLane.TransmitPower})
    opticsLane.EntityData.Leafs.Append("receive-power", types.YLeaf{"ReceivePower", opticsLane.ReceivePower})
    opticsLane.EntityData.Leafs.Append("receive-signal-power", types.YLeaf{"ReceiveSignalPower", opticsLane.ReceiveSignalPower})
    opticsLane.EntityData.Leafs.Append("transmit-signal-power", types.YLeaf{"TransmitSignalPower", opticsLane.TransmitSignalPower})
    opticsLane.EntityData.Leafs.Append("output-frequency", types.YLeaf{"OutputFrequency", opticsLane.OutputFrequency})
    opticsLane.EntityData.Leafs.Append("frequency-offset", types.YLeaf{"FrequencyOffset", opticsLane.FrequencyOffset})

    opticsLane.EntityData.YListKeys = []string {"Number"}

    return &(opticsLane.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo
// Lane Alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // High Rx Power.
    HighRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower

    // Low Rx Power.
    LowRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower

    // High Tx Power.
    HighTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower

    // Low Tx Power.
    LowTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower

    // High laser bias current.
    HighLbc OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetEntityData() *types.CommonEntityData {
    laneAlarmInfo.EntityData.YFilter = laneAlarmInfo.YFilter
    laneAlarmInfo.EntityData.YangName = "lane-alarm-info"
    laneAlarmInfo.EntityData.BundleName = "cisco_ios_xr"
    laneAlarmInfo.EntityData.ParentYangName = "optics-lane"
    laneAlarmInfo.EntityData.SegmentPath = "lane-alarm-info"
    laneAlarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-lanes/optics-lane/" + laneAlarmInfo.EntityData.SegmentPath
    laneAlarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    laneAlarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    laneAlarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    laneAlarmInfo.EntityData.Children = types.NewOrderedMap()
    laneAlarmInfo.EntityData.Children.Append("high-rx-power", types.YChild{"HighRxPower", &laneAlarmInfo.HighRxPower})
    laneAlarmInfo.EntityData.Children.Append("low-rx-power", types.YChild{"LowRxPower", &laneAlarmInfo.LowRxPower})
    laneAlarmInfo.EntityData.Children.Append("high-tx-power", types.YChild{"HighTxPower", &laneAlarmInfo.HighTxPower})
    laneAlarmInfo.EntityData.Children.Append("low-tx-power", types.YChild{"LowTxPower", &laneAlarmInfo.LowTxPower})
    laneAlarmInfo.EntityData.Children.Append("high-lbc", types.YChild{"HighLbc", &laneAlarmInfo.HighLbc})
    laneAlarmInfo.EntityData.Leafs = types.NewOrderedMap()

    laneAlarmInfo.EntityData.YListKeys = []string {}

    return &(laneAlarmInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower
// High Rx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetEntityData() *types.CommonEntityData {
    highRxPower.EntityData.YFilter = highRxPower.YFilter
    highRxPower.EntityData.YangName = "high-rx-power"
    highRxPower.EntityData.BundleName = "cisco_ios_xr"
    highRxPower.EntityData.ParentYangName = "lane-alarm-info"
    highRxPower.EntityData.SegmentPath = "high-rx-power"
    highRxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-lanes/optics-lane/lane-alarm-info/" + highRxPower.EntityData.SegmentPath
    highRxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highRxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highRxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highRxPower.EntityData.Children = types.NewOrderedMap()
    highRxPower.EntityData.Leafs = types.NewOrderedMap()
    highRxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highRxPower.IsDetected})
    highRxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highRxPower.Counter})

    highRxPower.EntityData.YListKeys = []string {}

    return &(highRxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower
// Low Rx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetEntityData() *types.CommonEntityData {
    lowRxPower.EntityData.YFilter = lowRxPower.YFilter
    lowRxPower.EntityData.YangName = "low-rx-power"
    lowRxPower.EntityData.BundleName = "cisco_ios_xr"
    lowRxPower.EntityData.ParentYangName = "lane-alarm-info"
    lowRxPower.EntityData.SegmentPath = "low-rx-power"
    lowRxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-lanes/optics-lane/lane-alarm-info/" + lowRxPower.EntityData.SegmentPath
    lowRxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowRxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowRxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowRxPower.EntityData.Children = types.NewOrderedMap()
    lowRxPower.EntityData.Leafs = types.NewOrderedMap()
    lowRxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowRxPower.IsDetected})
    lowRxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowRxPower.Counter})

    lowRxPower.EntityData.YListKeys = []string {}

    return &(lowRxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower
// High Tx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetEntityData() *types.CommonEntityData {
    highTxPower.EntityData.YFilter = highTxPower.YFilter
    highTxPower.EntityData.YangName = "high-tx-power"
    highTxPower.EntityData.BundleName = "cisco_ios_xr"
    highTxPower.EntityData.ParentYangName = "lane-alarm-info"
    highTxPower.EntityData.SegmentPath = "high-tx-power"
    highTxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-lanes/optics-lane/lane-alarm-info/" + highTxPower.EntityData.SegmentPath
    highTxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highTxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highTxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highTxPower.EntityData.Children = types.NewOrderedMap()
    highTxPower.EntityData.Leafs = types.NewOrderedMap()
    highTxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highTxPower.IsDetected})
    highTxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highTxPower.Counter})

    highTxPower.EntityData.YListKeys = []string {}

    return &(highTxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower
// Low Tx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetEntityData() *types.CommonEntityData {
    lowTxPower.EntityData.YFilter = lowTxPower.YFilter
    lowTxPower.EntityData.YangName = "low-tx-power"
    lowTxPower.EntityData.BundleName = "cisco_ios_xr"
    lowTxPower.EntityData.ParentYangName = "lane-alarm-info"
    lowTxPower.EntityData.SegmentPath = "low-tx-power"
    lowTxPower.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-lanes/optics-lane/lane-alarm-info/" + lowTxPower.EntityData.SegmentPath
    lowTxPower.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lowTxPower.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lowTxPower.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lowTxPower.EntityData.Children = types.NewOrderedMap()
    lowTxPower.EntityData.Leafs = types.NewOrderedMap()
    lowTxPower.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lowTxPower.IsDetected})
    lowTxPower.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lowTxPower.Counter})

    lowTxPower.EntityData.YListKeys = []string {}

    return &(lowTxPower.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc
// High laser bias current
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetEntityData() *types.CommonEntityData {
    highLbc.EntityData.YFilter = highLbc.YFilter
    highLbc.EntityData.YangName = "high-lbc"
    highLbc.EntityData.BundleName = "cisco_ios_xr"
    highLbc.EntityData.ParentYangName = "lane-alarm-info"
    highLbc.EntityData.SegmentPath = "high-lbc"
    highLbc.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-lanes/optics-lane/lane-alarm-info/" + highLbc.EntityData.SegmentPath
    highLbc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    highLbc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    highLbc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    highLbc.EntityData.Children = types.NewOrderedMap()
    highLbc.EntityData.Leafs = types.NewOrderedMap()
    highLbc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", highLbc.IsDetected})
    highLbc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", highLbc.Counter})

    highLbc.EntityData.YListKeys = []string {}

    return &(highLbc.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport Admin State. The type is OpticsTas.
    TransportAdminState interface{}

    // Optics controller state: Up, Down or Administratively Down. The type is
    // OpticsControllerState.
    ControllerState interface{}

    // Network SRLG information.
    NetworkSrlgInfo OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo
}

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetEntityData() *types.CommonEntityData {
    opticsDbInfo.EntityData.YFilter = opticsDbInfo.YFilter
    opticsDbInfo.EntityData.YangName = "optics-db-info"
    opticsDbInfo.EntityData.BundleName = "cisco_ios_xr"
    opticsDbInfo.EntityData.ParentYangName = "optics-port"
    opticsDbInfo.EntityData.SegmentPath = "optics-db-info"
    opticsDbInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/" + opticsDbInfo.EntityData.SegmentPath
    opticsDbInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsDbInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsDbInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsDbInfo.EntityData.Children = types.NewOrderedMap()
    opticsDbInfo.EntityData.Children.Append("network-srlg-info", types.YChild{"NetworkSrlgInfo", &opticsDbInfo.NetworkSrlgInfo})
    opticsDbInfo.EntityData.Leafs = types.NewOrderedMap()
    opticsDbInfo.EntityData.Leafs.Append("transport-admin-state", types.YLeaf{"TransportAdminState", opticsDbInfo.TransportAdminState})
    opticsDbInfo.EntityData.Leafs.Append("controller-state", types.YLeaf{"ControllerState", opticsDbInfo.ControllerState})

    opticsDbInfo.EntityData.YListKeys = []string {}

    return &(opticsDbInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo
// Network SRLG information
type OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network Srlg Array. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray.
    NetworkSrlgArray []*OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetEntityData() *types.CommonEntityData {
    networkSrlgInfo.EntityData.YFilter = networkSrlgInfo.YFilter
    networkSrlgInfo.EntityData.YangName = "network-srlg-info"
    networkSrlgInfo.EntityData.BundleName = "cisco_ios_xr"
    networkSrlgInfo.EntityData.ParentYangName = "optics-db-info"
    networkSrlgInfo.EntityData.SegmentPath = "network-srlg-info"
    networkSrlgInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-db-info/" + networkSrlgInfo.EntityData.SegmentPath
    networkSrlgInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkSrlgInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkSrlgInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkSrlgInfo.EntityData.Children = types.NewOrderedMap()
    networkSrlgInfo.EntityData.Children.Append("network-srlg-array", types.YChild{"NetworkSrlgArray", nil})
    for i := range networkSrlgInfo.NetworkSrlgArray {
        types.SetYListKey(networkSrlgInfo.NetworkSrlgArray[i], i)
        networkSrlgInfo.EntityData.Children.Append(types.GetSegmentPath(networkSrlgInfo.NetworkSrlgArray[i]), types.YChild{"NetworkSrlgArray", networkSrlgInfo.NetworkSrlgArray[i]})
    }
    networkSrlgInfo.EntityData.Leafs = types.NewOrderedMap()

    networkSrlgInfo.EntityData.YListKeys = []string {}

    return &(networkSrlgInfo.EntityData)
}

// OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray
// Network Srlg Array
type OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Array to maintain set number. The type is interface{} with range:
    // 0..4294967295.
    SetNumber interface{}

    // Network Srlg. The type is slice of interface{} with range: 0..4294967295.
    NetworkSrlg []interface{}
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetEntityData() *types.CommonEntityData {
    networkSrlgArray.EntityData.YFilter = networkSrlgArray.YFilter
    networkSrlgArray.EntityData.YangName = "network-srlg-array"
    networkSrlgArray.EntityData.BundleName = "cisco_ios_xr"
    networkSrlgArray.EntityData.ParentYangName = "network-srlg-info"
    networkSrlgArray.EntityData.SegmentPath = "network-srlg-array" + types.AddNoKeyToken(networkSrlgArray)
    networkSrlgArray.EntityData.AbsolutePath = "Cisco-IOS-XR-controller-optics-oper:optics-oper/optics-ports/optics-port/optics-db-info/network-srlg-info/" + networkSrlgArray.EntityData.SegmentPath
    networkSrlgArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkSrlgArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkSrlgArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkSrlgArray.EntityData.Children = types.NewOrderedMap()
    networkSrlgArray.EntityData.Leafs = types.NewOrderedMap()
    networkSrlgArray.EntityData.Leafs.Append("set-number", types.YLeaf{"SetNumber", networkSrlgArray.SetNumber})
    networkSrlgArray.EntityData.Leafs.Append("network-srlg", types.YLeaf{"NetworkSrlg", networkSrlgArray.NetworkSrlg})

    networkSrlgArray.EntityData.YListKeys = []string {}

    return &(networkSrlgArray.EntityData)
}

