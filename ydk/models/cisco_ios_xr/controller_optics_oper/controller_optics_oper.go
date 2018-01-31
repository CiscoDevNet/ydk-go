// This module contains a collection of YANG definitions
// for Cisco IOS-XR controller-optics package operational data.
// 
// This module contains definitions
// for the following management objects:
//   optics-oper: Optics operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // All Optics Port operational data.
    OpticsPorts OpticsOper_OpticsPorts
}

func (opticsOper *OpticsOper) GetFilter() yfilter.YFilter { return opticsOper.YFilter }

func (opticsOper *OpticsOper) SetFilter(yf yfilter.YFilter) { opticsOper.YFilter = yf }

func (opticsOper *OpticsOper) GetGoName(yname string) string {
    if yname == "optics-ports" { return "OpticsPorts" }
    return ""
}

func (opticsOper *OpticsOper) GetSegmentPath() string {
    return "Cisco-IOS-XR-controller-optics-oper:optics-oper"
}

func (opticsOper *OpticsOper) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optics-ports" {
        return &opticsOper.OpticsPorts
    }
    return nil
}

func (opticsOper *OpticsOper) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optics-ports"] = &opticsOper.OpticsPorts
    return children
}

func (opticsOper *OpticsOper) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticsOper *OpticsOper) GetBundleName() string { return "cisco_ios_xr" }

func (opticsOper *OpticsOper) GetYangName() string { return "optics-oper" }

func (opticsOper *OpticsOper) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsOper *OpticsOper) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsOper *OpticsOper) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsOper *OpticsOper) SetParent(parent types.Entity) { opticsOper.parent = parent }

func (opticsOper *OpticsOper) GetParent() types.Entity { return opticsOper.parent }

func (opticsOper *OpticsOper) GetParentYangName() string { return "Cisco-IOS-XR-controller-optics-oper" }

// OpticsOper_OpticsPorts
// All Optics Port operational data
type OpticsOper_OpticsPorts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Optics operational data. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort.
    OpticsPort []OpticsOper_OpticsPorts_OpticsPort
}

func (opticsPorts *OpticsOper_OpticsPorts) GetFilter() yfilter.YFilter { return opticsPorts.YFilter }

func (opticsPorts *OpticsOper_OpticsPorts) SetFilter(yf yfilter.YFilter) { opticsPorts.YFilter = yf }

func (opticsPorts *OpticsOper_OpticsPorts) GetGoName(yname string) string {
    if yname == "optics-port" { return "OpticsPort" }
    return ""
}

func (opticsPorts *OpticsOper_OpticsPorts) GetSegmentPath() string {
    return "optics-ports"
}

func (opticsPorts *OpticsOper_OpticsPorts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optics-port" {
        for _, c := range opticsPorts.OpticsPort {
            if opticsPorts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticsOper_OpticsPorts_OpticsPort{}
        opticsPorts.OpticsPort = append(opticsPorts.OpticsPort, child)
        return &opticsPorts.OpticsPort[len(opticsPorts.OpticsPort)-1]
    }
    return nil
}

func (opticsPorts *OpticsOper_OpticsPorts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opticsPorts.OpticsPort {
        children[opticsPorts.OpticsPort[i].GetSegmentPath()] = &opticsPorts.OpticsPort[i]
    }
    return children
}

func (opticsPorts *OpticsOper_OpticsPorts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticsPorts *OpticsOper_OpticsPorts) GetBundleName() string { return "cisco_ios_xr" }

func (opticsPorts *OpticsOper_OpticsPorts) GetYangName() string { return "optics-ports" }

func (opticsPorts *OpticsOper_OpticsPorts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsPorts *OpticsOper_OpticsPorts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsPorts *OpticsOper_OpticsPorts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsPorts *OpticsOper_OpticsPorts) SetParent(parent types.Entity) { opticsPorts.parent = parent }

func (opticsPorts *OpticsOper_OpticsPorts) GetParent() types.Entity { return opticsPorts.parent }

func (opticsPorts *OpticsOper_OpticsPorts) GetParentYangName() string { return "optics-oper" }

// OpticsOper_OpticsPorts_OpticsPort
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // Optics operational data.
    OpticsDwdmCarrrierChannelMap OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap

    // Optics operational data.
    OpticsInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo

    // All Optics Port operational data.
    OpticsLanes OpticsOper_OpticsPorts_OpticsPort_OpticsLanes

    // Optics operational data.
    OpticsDbInfo OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo
}

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetFilter() yfilter.YFilter { return opticsPort.YFilter }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) SetFilter(yf yfilter.YFilter) { opticsPort.YFilter = yf }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "optics-dwdm-carrrier-channel-map" { return "OpticsDwdmCarrrierChannelMap" }
    if yname == "optics-info" { return "OpticsInfo" }
    if yname == "optics-lanes" { return "OpticsLanes" }
    if yname == "optics-db-info" { return "OpticsDbInfo" }
    return ""
}

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetSegmentPath() string {
    return "optics-port" + "[name='" + fmt.Sprintf("%v", opticsPort.Name) + "']"
}

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optics-dwdm-carrrier-channel-map" {
        return &opticsPort.OpticsDwdmCarrrierChannelMap
    }
    if childYangName == "optics-info" {
        return &opticsPort.OpticsInfo
    }
    if childYangName == "optics-lanes" {
        return &opticsPort.OpticsLanes
    }
    if childYangName == "optics-db-info" {
        return &opticsPort.OpticsDbInfo
    }
    return nil
}

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["optics-dwdm-carrrier-channel-map"] = &opticsPort.OpticsDwdmCarrrierChannelMap
    children["optics-info"] = &opticsPort.OpticsInfo
    children["optics-lanes"] = &opticsPort.OpticsLanes
    children["optics-db-info"] = &opticsPort.OpticsDbInfo
    return children
}

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = opticsPort.Name
    return leafs
}

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetBundleName() string { return "cisco_ios_xr" }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetYangName() string { return "optics-port" }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) SetParent(parent types.Entity) { opticsPort.parent = parent }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetParent() types.Entity { return opticsPort.parent }

func (opticsPort *OpticsOper_OpticsPorts_OpticsPort) GetParentYangName() string { return "optics-ports" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap struct {
    parent types.Entity
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
    // OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo.
    DwdmCarrierMapInfo []OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo
}

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetFilter() yfilter.YFilter { return opticsDwdmCarrrierChannelMap.YFilter }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) SetFilter(yf yfilter.YFilter) { opticsDwdmCarrrierChannelMap.YFilter = yf }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetGoName(yname string) string {
    if yname == "dwdm-carrier-band" { return "DwdmCarrierBand" }
    if yname == "dwdm-carrier-min" { return "DwdmCarrierMin" }
    if yname == "dwdm-carrier-max" { return "DwdmCarrierMax" }
    if yname == "dwdm-carrier-map-info" { return "DwdmCarrierMapInfo" }
    return ""
}

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetSegmentPath() string {
    return "optics-dwdm-carrrier-channel-map"
}

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dwdm-carrier-map-info" {
        for _, c := range opticsDwdmCarrrierChannelMap.DwdmCarrierMapInfo {
            if opticsDwdmCarrrierChannelMap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo{}
        opticsDwdmCarrrierChannelMap.DwdmCarrierMapInfo = append(opticsDwdmCarrrierChannelMap.DwdmCarrierMapInfo, child)
        return &opticsDwdmCarrrierChannelMap.DwdmCarrierMapInfo[len(opticsDwdmCarrrierChannelMap.DwdmCarrierMapInfo)-1]
    }
    return nil
}

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opticsDwdmCarrrierChannelMap.DwdmCarrierMapInfo {
        children[opticsDwdmCarrrierChannelMap.DwdmCarrierMapInfo[i].GetSegmentPath()] = &opticsDwdmCarrrierChannelMap.DwdmCarrierMapInfo[i]
    }
    return children
}

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dwdm-carrier-band"] = opticsDwdmCarrrierChannelMap.DwdmCarrierBand
    leafs["dwdm-carrier-min"] = opticsDwdmCarrrierChannelMap.DwdmCarrierMin
    leafs["dwdm-carrier-max"] = opticsDwdmCarrrierChannelMap.DwdmCarrierMax
    return leafs
}

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetBundleName() string { return "cisco_ios_xr" }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetYangName() string { return "optics-dwdm-carrrier-channel-map" }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) SetParent(parent types.Entity) { opticsDwdmCarrrierChannelMap.parent = parent }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetParent() types.Entity { return opticsDwdmCarrrierChannelMap.parent }

func (opticsDwdmCarrrierChannelMap *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap) GetParentYangName() string { return "optics-port" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo
// DWDM carrier mapping info
type OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetFilter() yfilter.YFilter { return dwdmCarrierMapInfo.YFilter }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) SetFilter(yf yfilter.YFilter) { dwdmCarrierMapInfo.YFilter = yf }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetGoName(yname string) string {
    if yname == "itu-chan-num" { return "ItuChanNum" }
    if yname == "g694-chan-num" { return "G694ChanNum" }
    if yname == "frequency" { return "Frequency" }
    if yname == "wavelength" { return "Wavelength" }
    return ""
}

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetSegmentPath() string {
    return "dwdm-carrier-map-info"
}

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["itu-chan-num"] = dwdmCarrierMapInfo.ItuChanNum
    leafs["g694-chan-num"] = dwdmCarrierMapInfo.G694ChanNum
    leafs["frequency"] = dwdmCarrierMapInfo.Frequency
    leafs["wavelength"] = dwdmCarrierMapInfo.Wavelength
    return leafs
}

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetBundleName() string { return "cisco_ios_xr" }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetYangName() string { return "dwdm-carrier-map-info" }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) SetParent(parent types.Entity) { dwdmCarrierMapInfo.parent = parent }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetParent() types.Entity { return dwdmCarrierMapInfo.parent }

func (dwdmCarrierMapInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDwdmCarrrierChannelMap_DwdmCarrierMapInfo) GetParentYangName() string { return "optics-dwdm-carrrier-channel-map" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo struct {
    parent types.Entity
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

    // DWDM Carrier frequency read from hw in the unit 1THz. The type is string.
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

    // LBC high threshold default value in unit of 0 .001mA. The type is
    // interface{} with range: -2147483648..2147483647.
    LbcThHighDefault interface{}

    // LBC low threshold default value in units of 0 .001mA. The type is
    // interface{} with range: -2147483648..2147483647.
    LbcThLowDefault interface{}

    // Temp Low threshold value in the units 0.01 C. The type is interface{} with
    // range: -2147483648..2147483647.
    TempLowThreshold interface{}

    // Temp High threshold value in the units of 0.01 C. The type is interface{}
    // with range: -2147483648..2147483647.
    TempHighThreshold interface{}

    // Volt Low threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    VoltLowThreshold interface{}

    // Volt High threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
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

    // Showing Current Colour of led state. The type is OpticsLedState.
    LedState interface{}

    // Optics controller state: Up, Down or Administratively Down. The type is
    // OpticsControllerState.
    ControllerState interface{}

    // Optics form factor. The type is OpticsFormFactor.
    FormFactor interface{}

    // Optics physical type. The type is OpticsPhy.
    PhyType interface{}

    // Configured Tx power value in 0.1 dB. The type is interface{} with range:
    // -2147483648..2147483647.
    CfgTxPower interface{}

    // TX Power Configuration is supported or not. The type is bool.
    CfgTxPowerConfigurable interface{}

    // Temperature value. The type is interface{} with range:
    // -2147483648..2147483647.
    Temperature interface{}

    // Voltage value. The type is interface{} with range: -2147483648..2147483647.
    Voltage interface{}

    // Display Volt/Temp ?. The type is bool.
    DisplayVoltTemp interface{}

    // CD Configurable is supported or not. The type is bool.
    CdConfigurable interface{}

    // Optics FEC. The type is OpticsFec.
    OpticsFec interface{}

    // Showing port type. The type is OpticsPort.
    PortType interface{}

    // Showing port status. The type is OpticsPortStatus.
    PortStatus interface{}

    // Rx Voa Attenuation in the unit of 0.01dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    RxVoaAttenuation interface{}

    // Tx Voa Attenuation in the unit of 0.01dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    TxVoaAttenuation interface{}

    // Ampli Gain in the unit of 0.01dBm. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliGain interface{}

    // Ampli Tilt in the unit of 0.01dBm. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliTilt interface{}

    // rx power th configurable. The type is bool.
    RxPowerThConfigurable interface{}

    // tx power th configurable. The type is bool.
    TxPowerThConfigurable interface{}

    // rx voa attenuation config val. The type is interface{} with range:
    // -2147483648..2147483647.
    RxVoaAttenuationConfigVal interface{}

    // tx voa attenuation config val. The type is interface{} with range:
    // -2147483648..2147483647.
    TxVoaAttenuationConfigVal interface{}

    // ampli control mode config val. The type is OpticsAmplifierControlMode.
    AmpliControlModeConfigVal interface{}

    // ampli gain range config val. The type is OpticsAmplifierGainRange.
    AmpliGainRangeConfigVal interface{}

    // ampli gain config val. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliGainConfigVal interface{}

    // ampli tilt config val. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliTiltConfigVal interface{}

    // ampli channel power config val. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliChannelPowerConfigVal interface{}

    // channel power max delta config val. The type is interface{} with range:
    // -2147483648..2147483647.
    ChannelPowerMaxDeltaConfigVal interface{}

    // ampli gain thr deg low config val. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliGainThrDegLowConfigVal interface{}

    // ampli gain thr deg high config val. The type is interface{} with range:
    // -2147483648..2147483647.
    AmpliGainThrDegHighConfigVal interface{}

    // osri config val. The type is bool.
    OsriConfigVal interface{}

    // safety control mode config val. The type is
    // OpticsAmplifierSafetyControlMode.
    SafetyControlModeConfigVal interface{}

    // Total Receive Power for Multi-Lane Optics in dBm. The type is interface{}
    // with range: -2147483648..2147483647.
    TotalRxPower interface{}

    // Total Transmit Power for Multi-Lane Optics in dBm. The type is interface{}
    // with range: -2147483648..2147483647.
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

    // LBC high Warning threshold default value in unit of 0.001mA. The type is
    // interface{} with range: -2147483648..2147483647.
    LbcThHighWarningDefault interface{}

    // LBC low warning threshold default value in units of 0.001mA. The type is
    // interface{} with range: -2147483648..2147483647.
    LbcThLowWarningDefault interface{}

    // Temp Low warning threshold value in the units 0 .01 C. The type is
    // interface{} with range: -2147483648..2147483647.
    TempLowWarningThreshold interface{}

    // Temp High warning threshold value in the units of 0.01 C. The type is
    // interface{} with range: -2147483648..2147483647.
    TempHighWarningThreshold interface{}

    // Volt Low warning threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    VoltLowWarningThreshold interface{}

    // Volt High warning threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    VoltHighWarningThreshold interface{}

    // Ampli gain range. The type is OpticsAmplifierGainRange.
    AmpliGainRange interface{}

    // Safety control mode. The type is OpticsAmplifierSafetyControlMode.
    SafetyControlMode interface{}

    // OSRI. The type is bool.
    Osri interface{}

    // Controller description string. The type is string.
    Description interface{}

    // Is the Optics type string valid ?. The type is bool.
    IsOpticsTypeStringValid interface{}

    // optics type String. The type is string.
    OpticsTypeStr interface{}

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

    // OTS Spectrum information.
    SpectrumInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo

    // Lane information. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData.
    LaneData []OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData
}

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetFilter() yfilter.YFilter { return opticsInfo.YFilter }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) SetFilter(yf yfilter.YFilter) { opticsInfo.YFilter = yf }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetGoName(yname string) string {
    if yname == "transport-admin-state" { return "TransportAdminState" }
    if yname == "optics-present" { return "OpticsPresent" }
    if yname == "optics-type" { return "OpticsType" }
    if yname == "derived-optics-type" { return "DerivedOpticsType" }
    if yname == "optics-module" { return "OpticsModule" }
    if yname == "dwdm-carrier-band" { return "DwdmCarrierBand" }
    if yname == "dwdm-carrier-channel" { return "DwdmCarrierChannel" }
    if yname == "dwdm-carrier-frequency" { return "DwdmCarrierFrequency" }
    if yname == "dwdm-carrier-wavelength" { return "DwdmCarrierWavelength" }
    if yname == "grey-wavelength" { return "GreyWavelength" }
    if yname == "rx-low-threshold" { return "RxLowThreshold" }
    if yname == "rx-high-threshold" { return "RxHighThreshold" }
    if yname == "lbc-high-threshold" { return "LbcHighThreshold" }
    if yname == "tx-low-threshold" { return "TxLowThreshold" }
    if yname == "tx-high-threshold" { return "TxHighThreshold" }
    if yname == "lbc-th-high-default" { return "LbcThHighDefault" }
    if yname == "lbc-th-low-default" { return "LbcThLowDefault" }
    if yname == "temp-low-threshold" { return "TempLowThreshold" }
    if yname == "temp-high-threshold" { return "TempHighThreshold" }
    if yname == "volt-low-threshold" { return "VoltLowThreshold" }
    if yname == "volt-high-threshold" { return "VoltHighThreshold" }
    if yname == "cd" { return "Cd" }
    if yname == "cd-min" { return "CdMin" }
    if yname == "cd-max" { return "CdMax" }
    if yname == "cd-low-threshold" { return "CdLowThreshold" }
    if yname == "cd-high-threshold" { return "CdHighThreshold" }
    if yname == "osnr-low-threshold" { return "OsnrLowThreshold" }
    if yname == "dgd-high-threshold" { return "DgdHighThreshold" }
    if yname == "polarization-mode-dispersion" { return "PolarizationModeDispersion" }
    if yname == "second-order-polarization-mode-dispersion" { return "SecondOrderPolarizationModeDispersion" }
    if yname == "optical-signal-to-noise-ratio" { return "OpticalSignalToNoiseRatio" }
    if yname == "polarization-dependent-loss" { return "PolarizationDependentLoss" }
    if yname == "polarization-change-rate" { return "PolarizationChangeRate" }
    if yname == "differential-group-delay" { return "DifferentialGroupDelay" }
    if yname == "phase-noise" { return "PhaseNoise" }
    if yname == "pm-enable" { return "PmEnable" }
    if yname == "laser-state" { return "LaserState" }
    if yname == "led-state" { return "LedState" }
    if yname == "controller-state" { return "ControllerState" }
    if yname == "form-factor" { return "FormFactor" }
    if yname == "phy-type" { return "PhyType" }
    if yname == "cfg-tx-power" { return "CfgTxPower" }
    if yname == "cfg-tx-power-configurable" { return "CfgTxPowerConfigurable" }
    if yname == "temperature" { return "Temperature" }
    if yname == "voltage" { return "Voltage" }
    if yname == "display-volt-temp" { return "DisplayVoltTemp" }
    if yname == "cd-configurable" { return "CdConfigurable" }
    if yname == "optics-fec" { return "OpticsFec" }
    if yname == "port-type" { return "PortType" }
    if yname == "port-status" { return "PortStatus" }
    if yname == "rx-voa-attenuation" { return "RxVoaAttenuation" }
    if yname == "tx-voa-attenuation" { return "TxVoaAttenuation" }
    if yname == "ampli-gain" { return "AmpliGain" }
    if yname == "ampli-tilt" { return "AmpliTilt" }
    if yname == "rx-power-th-configurable" { return "RxPowerThConfigurable" }
    if yname == "tx-power-th-configurable" { return "TxPowerThConfigurable" }
    if yname == "rx-voa-attenuation-config-val" { return "RxVoaAttenuationConfigVal" }
    if yname == "tx-voa-attenuation-config-val" { return "TxVoaAttenuationConfigVal" }
    if yname == "ampli-control-mode-config-val" { return "AmpliControlModeConfigVal" }
    if yname == "ampli-gain-range-config-val" { return "AmpliGainRangeConfigVal" }
    if yname == "ampli-gain-config-val" { return "AmpliGainConfigVal" }
    if yname == "ampli-tilt-config-val" { return "AmpliTiltConfigVal" }
    if yname == "ampli-channel-power-config-val" { return "AmpliChannelPowerConfigVal" }
    if yname == "channel-power-max-delta-config-val" { return "ChannelPowerMaxDeltaConfigVal" }
    if yname == "ampli-gain-thr-deg-low-config-val" { return "AmpliGainThrDegLowConfigVal" }
    if yname == "ampli-gain-thr-deg-high-config-val" { return "AmpliGainThrDegHighConfigVal" }
    if yname == "osri-config-val" { return "OsriConfigVal" }
    if yname == "safety-control-mode-config-val" { return "SafetyControlModeConfigVal" }
    if yname == "total-rx-power" { return "TotalRxPower" }
    if yname == "total-tx-power" { return "TotalTxPower" }
    if yname == "is-bo-configured" { return "IsBoConfigured" }
    if yname == "is-ext-param-valid" { return "IsExtParamValid" }
    if yname == "alarm-detected" { return "AlarmDetected" }
    if yname == "rx-low-warning-threshold" { return "RxLowWarningThreshold" }
    if yname == "rx-high-warning-threshold" { return "RxHighWarningThreshold" }
    if yname == "tx-low-warning-threshold" { return "TxLowWarningThreshold" }
    if yname == "tx-high-warning-threshold" { return "TxHighWarningThreshold" }
    if yname == "lbc-th-high-warning-default" { return "LbcThHighWarningDefault" }
    if yname == "lbc-th-low-warning-default" { return "LbcThLowWarningDefault" }
    if yname == "temp-low-warning-threshold" { return "TempLowWarningThreshold" }
    if yname == "temp-high-warning-threshold" { return "TempHighWarningThreshold" }
    if yname == "volt-low-warning-threshold" { return "VoltLowWarningThreshold" }
    if yname == "volt-high-warning-threshold" { return "VoltHighWarningThreshold" }
    if yname == "ampli-gain-range" { return "AmpliGainRange" }
    if yname == "safety-control-mode" { return "SafetyControlMode" }
    if yname == "osri" { return "Osri" }
    if yname == "description" { return "Description" }
    if yname == "is-optics-type-string-valid" { return "IsOpticsTypeStringValid" }
    if yname == "optics-type-str" { return "OpticsTypeStr" }
    if yname == "network-srlg-info" { return "NetworkSrlgInfo" }
    if yname == "optics-alarm-info" { return "OpticsAlarmInfo" }
    if yname == "ots-alarm-info" { return "OtsAlarmInfo" }
    if yname == "transceiver-info" { return "TransceiverInfo" }
    if yname == "ext-param-val" { return "ExtParamVal" }
    if yname == "ext-param-threshold-val" { return "ExtParamThresholdVal" }
    if yname == "spectrum-info" { return "SpectrumInfo" }
    if yname == "lane-data" { return "LaneData" }
    return ""
}

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetSegmentPath() string {
    return "optics-info"
}

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-srlg-info" {
        return &opticsInfo.NetworkSrlgInfo
    }
    if childYangName == "optics-alarm-info" {
        return &opticsInfo.OpticsAlarmInfo
    }
    if childYangName == "ots-alarm-info" {
        return &opticsInfo.OtsAlarmInfo
    }
    if childYangName == "transceiver-info" {
        return &opticsInfo.TransceiverInfo
    }
    if childYangName == "ext-param-val" {
        return &opticsInfo.ExtParamVal
    }
    if childYangName == "ext-param-threshold-val" {
        return &opticsInfo.ExtParamThresholdVal
    }
    if childYangName == "spectrum-info" {
        return &opticsInfo.SpectrumInfo
    }
    if childYangName == "lane-data" {
        for _, c := range opticsInfo.LaneData {
            if opticsInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData{}
        opticsInfo.LaneData = append(opticsInfo.LaneData, child)
        return &opticsInfo.LaneData[len(opticsInfo.LaneData)-1]
    }
    return nil
}

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-srlg-info"] = &opticsInfo.NetworkSrlgInfo
    children["optics-alarm-info"] = &opticsInfo.OpticsAlarmInfo
    children["ots-alarm-info"] = &opticsInfo.OtsAlarmInfo
    children["transceiver-info"] = &opticsInfo.TransceiverInfo
    children["ext-param-val"] = &opticsInfo.ExtParamVal
    children["ext-param-threshold-val"] = &opticsInfo.ExtParamThresholdVal
    children["spectrum-info"] = &opticsInfo.SpectrumInfo
    for i := range opticsInfo.LaneData {
        children[opticsInfo.LaneData[i].GetSegmentPath()] = &opticsInfo.LaneData[i]
    }
    return children
}

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport-admin-state"] = opticsInfo.TransportAdminState
    leafs["optics-present"] = opticsInfo.OpticsPresent
    leafs["optics-type"] = opticsInfo.OpticsType
    leafs["derived-optics-type"] = opticsInfo.DerivedOpticsType
    leafs["optics-module"] = opticsInfo.OpticsModule
    leafs["dwdm-carrier-band"] = opticsInfo.DwdmCarrierBand
    leafs["dwdm-carrier-channel"] = opticsInfo.DwdmCarrierChannel
    leafs["dwdm-carrier-frequency"] = opticsInfo.DwdmCarrierFrequency
    leafs["dwdm-carrier-wavelength"] = opticsInfo.DwdmCarrierWavelength
    leafs["grey-wavelength"] = opticsInfo.GreyWavelength
    leafs["rx-low-threshold"] = opticsInfo.RxLowThreshold
    leafs["rx-high-threshold"] = opticsInfo.RxHighThreshold
    leafs["lbc-high-threshold"] = opticsInfo.LbcHighThreshold
    leafs["tx-low-threshold"] = opticsInfo.TxLowThreshold
    leafs["tx-high-threshold"] = opticsInfo.TxHighThreshold
    leafs["lbc-th-high-default"] = opticsInfo.LbcThHighDefault
    leafs["lbc-th-low-default"] = opticsInfo.LbcThLowDefault
    leafs["temp-low-threshold"] = opticsInfo.TempLowThreshold
    leafs["temp-high-threshold"] = opticsInfo.TempHighThreshold
    leafs["volt-low-threshold"] = opticsInfo.VoltLowThreshold
    leafs["volt-high-threshold"] = opticsInfo.VoltHighThreshold
    leafs["cd"] = opticsInfo.Cd
    leafs["cd-min"] = opticsInfo.CdMin
    leafs["cd-max"] = opticsInfo.CdMax
    leafs["cd-low-threshold"] = opticsInfo.CdLowThreshold
    leafs["cd-high-threshold"] = opticsInfo.CdHighThreshold
    leafs["osnr-low-threshold"] = opticsInfo.OsnrLowThreshold
    leafs["dgd-high-threshold"] = opticsInfo.DgdHighThreshold
    leafs["polarization-mode-dispersion"] = opticsInfo.PolarizationModeDispersion
    leafs["second-order-polarization-mode-dispersion"] = opticsInfo.SecondOrderPolarizationModeDispersion
    leafs["optical-signal-to-noise-ratio"] = opticsInfo.OpticalSignalToNoiseRatio
    leafs["polarization-dependent-loss"] = opticsInfo.PolarizationDependentLoss
    leafs["polarization-change-rate"] = opticsInfo.PolarizationChangeRate
    leafs["differential-group-delay"] = opticsInfo.DifferentialGroupDelay
    leafs["phase-noise"] = opticsInfo.PhaseNoise
    leafs["pm-enable"] = opticsInfo.PmEnable
    leafs["laser-state"] = opticsInfo.LaserState
    leafs["led-state"] = opticsInfo.LedState
    leafs["controller-state"] = opticsInfo.ControllerState
    leafs["form-factor"] = opticsInfo.FormFactor
    leafs["phy-type"] = opticsInfo.PhyType
    leafs["cfg-tx-power"] = opticsInfo.CfgTxPower
    leafs["cfg-tx-power-configurable"] = opticsInfo.CfgTxPowerConfigurable
    leafs["temperature"] = opticsInfo.Temperature
    leafs["voltage"] = opticsInfo.Voltage
    leafs["display-volt-temp"] = opticsInfo.DisplayVoltTemp
    leafs["cd-configurable"] = opticsInfo.CdConfigurable
    leafs["optics-fec"] = opticsInfo.OpticsFec
    leafs["port-type"] = opticsInfo.PortType
    leafs["port-status"] = opticsInfo.PortStatus
    leafs["rx-voa-attenuation"] = opticsInfo.RxVoaAttenuation
    leafs["tx-voa-attenuation"] = opticsInfo.TxVoaAttenuation
    leafs["ampli-gain"] = opticsInfo.AmpliGain
    leafs["ampli-tilt"] = opticsInfo.AmpliTilt
    leafs["rx-power-th-configurable"] = opticsInfo.RxPowerThConfigurable
    leafs["tx-power-th-configurable"] = opticsInfo.TxPowerThConfigurable
    leafs["rx-voa-attenuation-config-val"] = opticsInfo.RxVoaAttenuationConfigVal
    leafs["tx-voa-attenuation-config-val"] = opticsInfo.TxVoaAttenuationConfigVal
    leafs["ampli-control-mode-config-val"] = opticsInfo.AmpliControlModeConfigVal
    leafs["ampli-gain-range-config-val"] = opticsInfo.AmpliGainRangeConfigVal
    leafs["ampli-gain-config-val"] = opticsInfo.AmpliGainConfigVal
    leafs["ampli-tilt-config-val"] = opticsInfo.AmpliTiltConfigVal
    leafs["ampli-channel-power-config-val"] = opticsInfo.AmpliChannelPowerConfigVal
    leafs["channel-power-max-delta-config-val"] = opticsInfo.ChannelPowerMaxDeltaConfigVal
    leafs["ampli-gain-thr-deg-low-config-val"] = opticsInfo.AmpliGainThrDegLowConfigVal
    leafs["ampli-gain-thr-deg-high-config-val"] = opticsInfo.AmpliGainThrDegHighConfigVal
    leafs["osri-config-val"] = opticsInfo.OsriConfigVal
    leafs["safety-control-mode-config-val"] = opticsInfo.SafetyControlModeConfigVal
    leafs["total-rx-power"] = opticsInfo.TotalRxPower
    leafs["total-tx-power"] = opticsInfo.TotalTxPower
    leafs["is-bo-configured"] = opticsInfo.IsBoConfigured
    leafs["is-ext-param-valid"] = opticsInfo.IsExtParamValid
    leafs["alarm-detected"] = opticsInfo.AlarmDetected
    leafs["rx-low-warning-threshold"] = opticsInfo.RxLowWarningThreshold
    leafs["rx-high-warning-threshold"] = opticsInfo.RxHighWarningThreshold
    leafs["tx-low-warning-threshold"] = opticsInfo.TxLowWarningThreshold
    leafs["tx-high-warning-threshold"] = opticsInfo.TxHighWarningThreshold
    leafs["lbc-th-high-warning-default"] = opticsInfo.LbcThHighWarningDefault
    leafs["lbc-th-low-warning-default"] = opticsInfo.LbcThLowWarningDefault
    leafs["temp-low-warning-threshold"] = opticsInfo.TempLowWarningThreshold
    leafs["temp-high-warning-threshold"] = opticsInfo.TempHighWarningThreshold
    leafs["volt-low-warning-threshold"] = opticsInfo.VoltLowWarningThreshold
    leafs["volt-high-warning-threshold"] = opticsInfo.VoltHighWarningThreshold
    leafs["ampli-gain-range"] = opticsInfo.AmpliGainRange
    leafs["safety-control-mode"] = opticsInfo.SafetyControlMode
    leafs["osri"] = opticsInfo.Osri
    leafs["description"] = opticsInfo.Description
    leafs["is-optics-type-string-valid"] = opticsInfo.IsOpticsTypeStringValid
    leafs["optics-type-str"] = opticsInfo.OpticsTypeStr
    return leafs
}

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetYangName() string { return "optics-info" }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) SetParent(parent types.Entity) { opticsInfo.parent = parent }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetParent() types.Entity { return opticsInfo.parent }

func (opticsInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo) GetParentYangName() string { return "optics-port" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo
// Network SRLG information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network Srlg Array. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray.
    NetworkSrlgArray []OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetFilter() yfilter.YFilter { return networkSrlgInfo.YFilter }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) SetFilter(yf yfilter.YFilter) { networkSrlgInfo.YFilter = yf }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetGoName(yname string) string {
    if yname == "network-srlg-array" { return "NetworkSrlgArray" }
    return ""
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetSegmentPath() string {
    return "network-srlg-info"
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-srlg-array" {
        for _, c := range networkSrlgInfo.NetworkSrlgArray {
            if networkSrlgInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray{}
        networkSrlgInfo.NetworkSrlgArray = append(networkSrlgInfo.NetworkSrlgArray, child)
        return &networkSrlgInfo.NetworkSrlgArray[len(networkSrlgInfo.NetworkSrlgArray)-1]
    }
    return nil
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkSrlgInfo.NetworkSrlgArray {
        children[networkSrlgInfo.NetworkSrlgArray[i].GetSegmentPath()] = &networkSrlgInfo.NetworkSrlgArray[i]
    }
    return children
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetBundleName() string { return "cisco_ios_xr" }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetYangName() string { return "network-srlg-info" }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) SetParent(parent types.Entity) { networkSrlgInfo.parent = parent }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetParent() types.Entity { return networkSrlgInfo.parent }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo) GetParentYangName() string { return "optics-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray
// Network Srlg Array
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array to maintain set number. The type is interface{} with range:
    // 0..4294967295.
    SetNumber interface{}

    // Network Srlg. The type is slice of interface{} with range: 0..4294967295.
    NetworkSrlg []interface{}
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetFilter() yfilter.YFilter { return networkSrlgArray.YFilter }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) SetFilter(yf yfilter.YFilter) { networkSrlgArray.YFilter = yf }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetGoName(yname string) string {
    if yname == "set-number" { return "SetNumber" }
    if yname == "network-srlg" { return "NetworkSrlg" }
    return ""
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetSegmentPath() string {
    return "network-srlg-array"
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-number"] = networkSrlgArray.SetNumber
    leafs["network-srlg"] = networkSrlgArray.NetworkSrlg
    return leafs
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetBundleName() string { return "cisco_ios_xr" }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetYangName() string { return "network-srlg-array" }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) SetParent(parent types.Entity) { networkSrlgArray.parent = parent }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetParent() types.Entity { return networkSrlgArray.parent }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_NetworkSrlgInfo_NetworkSrlgArray) GetParentYangName() string { return "network-srlg-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo
// Optics Alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // High Rx Power in uints of 0.1 dBm.
    HighRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower

    // Low Rx Power in uints of 0.1 dBm.
    LowRxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower

    // High Tx Power in uints of 0.1 dBm.
    HighTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower

    // Low Tx Power in uints of 0.1 dBm.
    LowTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower

    // High laser bias current in units of percentage.
    HighLbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc

    // High Rx1 Power in uints of 0.1 dBm.
    HighRx1Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power

    // High Rx2 Power in uints of 0.1 dBm.
    HighRx2Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power

    // High Rx3 Power in uints of 0.1 dBm.
    HighRx3Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power

    // High Rx4 Power in uints of 0.1 dBm.
    HighRx4Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power

    // Low Rx1 Power in uints of 0.1 dBm.
    LowRx1Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power

    // Low Rx2 Power in uints of 0.1 dBm.
    LowRx2Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power

    // Low Rx3 Power in uints of 0.1 dBm.
    LowRx3Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power

    // Low Rx4 Power in uints of 0.1 dBm.
    LowRx4Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power

    // High Tx1 Power in uints of 0.1 dBm.
    HighTx1Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power

    // High Tx2 Power in uints of 0.1 dBm.
    HighTx2Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power

    // High Tx3 Power in uints of 0.1 dBm.
    HighTx3Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power

    // High Tx4 Power in uints of 0.1 dBm.
    HighTx4Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power

    // Low Tx1 Power in uints of 0.1 dBm.
    LowTx1Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power

    // Low Tx2 Power in uints of 0.1 dBm.
    LowTx2Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power

    // Low Tx3 Power in uints of 0.1 dBm.
    LowTx3Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power

    // Low Tx4 Power in uints of 0.1 dBm.
    LowTx4Power OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power

    // High Tx1 laser bias current in units of percentage.
    HighTx1Lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc

    // High Tx2 laser bias current in units of percentage.
    HighTx2Lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc

    // High Tx3 laser bias current in units of percentage.
    HighTx3Lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc

    // High Tx4 laser bias current in units of percentage.
    HighTx4Lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc

    // Low Tx1 laser bias current in units of percentage.
    LowTx1Lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc

    // Low Tx2 laser bias current in units of percentage.
    LowTx2Lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc

    // Low Tx3 laser bias current in units of percentage.
    LowTx3Lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc

    // Low Tx4 laser bias current in units of percentage.
    LowTx4Lbc OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc

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

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetFilter() yfilter.YFilter { return opticsAlarmInfo.YFilter }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) SetFilter(yf yfilter.YFilter) { opticsAlarmInfo.YFilter = yf }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetGoName(yname string) string {
    if yname == "high-rx-power" { return "HighRxPower" }
    if yname == "low-rx-power" { return "LowRxPower" }
    if yname == "high-tx-power" { return "HighTxPower" }
    if yname == "low-tx-power" { return "LowTxPower" }
    if yname == "high-lbc" { return "HighLbc" }
    if yname == "high-rx1-power" { return "HighRx1Power" }
    if yname == "high-rx2-power" { return "HighRx2Power" }
    if yname == "high-rx3-power" { return "HighRx3Power" }
    if yname == "high-rx4-power" { return "HighRx4Power" }
    if yname == "low-rx1-power" { return "LowRx1Power" }
    if yname == "low-rx2-power" { return "LowRx2Power" }
    if yname == "low-rx3-power" { return "LowRx3Power" }
    if yname == "low-rx4-power" { return "LowRx4Power" }
    if yname == "high-tx1-power" { return "HighTx1Power" }
    if yname == "high-tx2-power" { return "HighTx2Power" }
    if yname == "high-tx3-power" { return "HighTx3Power" }
    if yname == "high-tx4-power" { return "HighTx4Power" }
    if yname == "low-tx1-power" { return "LowTx1Power" }
    if yname == "low-tx2-power" { return "LowTx2Power" }
    if yname == "low-tx3-power" { return "LowTx3Power" }
    if yname == "low-tx4-power" { return "LowTx4Power" }
    if yname == "high-tx1lbc" { return "HighTx1Lbc" }
    if yname == "high-tx2lbc" { return "HighTx2Lbc" }
    if yname == "high-tx3lbc" { return "HighTx3Lbc" }
    if yname == "high-tx4lbc" { return "HighTx4Lbc" }
    if yname == "low-tx1lbc" { return "LowTx1Lbc" }
    if yname == "low-tx2lbc" { return "LowTx2Lbc" }
    if yname == "low-tx3lbc" { return "LowTx3Lbc" }
    if yname == "low-tx4lbc" { return "LowTx4Lbc" }
    if yname == "rx-los" { return "RxLos" }
    if yname == "tx-los" { return "TxLos" }
    if yname == "rx-lol" { return "RxLol" }
    if yname == "tx-lol" { return "TxLol" }
    if yname == "tx-fault" { return "TxFault" }
    if yname == "hidgd" { return "Hidgd" }
    if yname == "oorcd" { return "Oorcd" }
    if yname == "osnr" { return "Osnr" }
    if yname == "wvlool" { return "Wvlool" }
    if yname == "mea" { return "Mea" }
    if yname == "imp-removal" { return "ImpRemoval" }
    if yname == "rx-loc" { return "RxLoc" }
    if yname == "amp-gain-deg-low" { return "AmpGainDegLow" }
    if yname == "amp-gain-deg-high" { return "AmpGainDegHigh" }
    if yname == "txpwr-mismatch" { return "TxpwrMismatch" }
    return ""
}

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetSegmentPath() string {
    return "optics-alarm-info"
}

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "high-rx-power" {
        return &opticsAlarmInfo.HighRxPower
    }
    if childYangName == "low-rx-power" {
        return &opticsAlarmInfo.LowRxPower
    }
    if childYangName == "high-tx-power" {
        return &opticsAlarmInfo.HighTxPower
    }
    if childYangName == "low-tx-power" {
        return &opticsAlarmInfo.LowTxPower
    }
    if childYangName == "high-lbc" {
        return &opticsAlarmInfo.HighLbc
    }
    if childYangName == "high-rx1-power" {
        return &opticsAlarmInfo.HighRx1Power
    }
    if childYangName == "high-rx2-power" {
        return &opticsAlarmInfo.HighRx2Power
    }
    if childYangName == "high-rx3-power" {
        return &opticsAlarmInfo.HighRx3Power
    }
    if childYangName == "high-rx4-power" {
        return &opticsAlarmInfo.HighRx4Power
    }
    if childYangName == "low-rx1-power" {
        return &opticsAlarmInfo.LowRx1Power
    }
    if childYangName == "low-rx2-power" {
        return &opticsAlarmInfo.LowRx2Power
    }
    if childYangName == "low-rx3-power" {
        return &opticsAlarmInfo.LowRx3Power
    }
    if childYangName == "low-rx4-power" {
        return &opticsAlarmInfo.LowRx4Power
    }
    if childYangName == "high-tx1-power" {
        return &opticsAlarmInfo.HighTx1Power
    }
    if childYangName == "high-tx2-power" {
        return &opticsAlarmInfo.HighTx2Power
    }
    if childYangName == "high-tx3-power" {
        return &opticsAlarmInfo.HighTx3Power
    }
    if childYangName == "high-tx4-power" {
        return &opticsAlarmInfo.HighTx4Power
    }
    if childYangName == "low-tx1-power" {
        return &opticsAlarmInfo.LowTx1Power
    }
    if childYangName == "low-tx2-power" {
        return &opticsAlarmInfo.LowTx2Power
    }
    if childYangName == "low-tx3-power" {
        return &opticsAlarmInfo.LowTx3Power
    }
    if childYangName == "low-tx4-power" {
        return &opticsAlarmInfo.LowTx4Power
    }
    if childYangName == "high-tx1lbc" {
        return &opticsAlarmInfo.HighTx1Lbc
    }
    if childYangName == "high-tx2lbc" {
        return &opticsAlarmInfo.HighTx2Lbc
    }
    if childYangName == "high-tx3lbc" {
        return &opticsAlarmInfo.HighTx3Lbc
    }
    if childYangName == "high-tx4lbc" {
        return &opticsAlarmInfo.HighTx4Lbc
    }
    if childYangName == "low-tx1lbc" {
        return &opticsAlarmInfo.LowTx1Lbc
    }
    if childYangName == "low-tx2lbc" {
        return &opticsAlarmInfo.LowTx2Lbc
    }
    if childYangName == "low-tx3lbc" {
        return &opticsAlarmInfo.LowTx3Lbc
    }
    if childYangName == "low-tx4lbc" {
        return &opticsAlarmInfo.LowTx4Lbc
    }
    if childYangName == "rx-los" {
        return &opticsAlarmInfo.RxLos
    }
    if childYangName == "tx-los" {
        return &opticsAlarmInfo.TxLos
    }
    if childYangName == "rx-lol" {
        return &opticsAlarmInfo.RxLol
    }
    if childYangName == "tx-lol" {
        return &opticsAlarmInfo.TxLol
    }
    if childYangName == "tx-fault" {
        return &opticsAlarmInfo.TxFault
    }
    if childYangName == "hidgd" {
        return &opticsAlarmInfo.Hidgd
    }
    if childYangName == "oorcd" {
        return &opticsAlarmInfo.Oorcd
    }
    if childYangName == "osnr" {
        return &opticsAlarmInfo.Osnr
    }
    if childYangName == "wvlool" {
        return &opticsAlarmInfo.Wvlool
    }
    if childYangName == "mea" {
        return &opticsAlarmInfo.Mea
    }
    if childYangName == "imp-removal" {
        return &opticsAlarmInfo.ImpRemoval
    }
    if childYangName == "rx-loc" {
        return &opticsAlarmInfo.RxLoc
    }
    if childYangName == "amp-gain-deg-low" {
        return &opticsAlarmInfo.AmpGainDegLow
    }
    if childYangName == "amp-gain-deg-high" {
        return &opticsAlarmInfo.AmpGainDegHigh
    }
    if childYangName == "txpwr-mismatch" {
        return &opticsAlarmInfo.TxpwrMismatch
    }
    return nil
}

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["high-rx-power"] = &opticsAlarmInfo.HighRxPower
    children["low-rx-power"] = &opticsAlarmInfo.LowRxPower
    children["high-tx-power"] = &opticsAlarmInfo.HighTxPower
    children["low-tx-power"] = &opticsAlarmInfo.LowTxPower
    children["high-lbc"] = &opticsAlarmInfo.HighLbc
    children["high-rx1-power"] = &opticsAlarmInfo.HighRx1Power
    children["high-rx2-power"] = &opticsAlarmInfo.HighRx2Power
    children["high-rx3-power"] = &opticsAlarmInfo.HighRx3Power
    children["high-rx4-power"] = &opticsAlarmInfo.HighRx4Power
    children["low-rx1-power"] = &opticsAlarmInfo.LowRx1Power
    children["low-rx2-power"] = &opticsAlarmInfo.LowRx2Power
    children["low-rx3-power"] = &opticsAlarmInfo.LowRx3Power
    children["low-rx4-power"] = &opticsAlarmInfo.LowRx4Power
    children["high-tx1-power"] = &opticsAlarmInfo.HighTx1Power
    children["high-tx2-power"] = &opticsAlarmInfo.HighTx2Power
    children["high-tx3-power"] = &opticsAlarmInfo.HighTx3Power
    children["high-tx4-power"] = &opticsAlarmInfo.HighTx4Power
    children["low-tx1-power"] = &opticsAlarmInfo.LowTx1Power
    children["low-tx2-power"] = &opticsAlarmInfo.LowTx2Power
    children["low-tx3-power"] = &opticsAlarmInfo.LowTx3Power
    children["low-tx4-power"] = &opticsAlarmInfo.LowTx4Power
    children["high-tx1lbc"] = &opticsAlarmInfo.HighTx1Lbc
    children["high-tx2lbc"] = &opticsAlarmInfo.HighTx2Lbc
    children["high-tx3lbc"] = &opticsAlarmInfo.HighTx3Lbc
    children["high-tx4lbc"] = &opticsAlarmInfo.HighTx4Lbc
    children["low-tx1lbc"] = &opticsAlarmInfo.LowTx1Lbc
    children["low-tx2lbc"] = &opticsAlarmInfo.LowTx2Lbc
    children["low-tx3lbc"] = &opticsAlarmInfo.LowTx3Lbc
    children["low-tx4lbc"] = &opticsAlarmInfo.LowTx4Lbc
    children["rx-los"] = &opticsAlarmInfo.RxLos
    children["tx-los"] = &opticsAlarmInfo.TxLos
    children["rx-lol"] = &opticsAlarmInfo.RxLol
    children["tx-lol"] = &opticsAlarmInfo.TxLol
    children["tx-fault"] = &opticsAlarmInfo.TxFault
    children["hidgd"] = &opticsAlarmInfo.Hidgd
    children["oorcd"] = &opticsAlarmInfo.Oorcd
    children["osnr"] = &opticsAlarmInfo.Osnr
    children["wvlool"] = &opticsAlarmInfo.Wvlool
    children["mea"] = &opticsAlarmInfo.Mea
    children["imp-removal"] = &opticsAlarmInfo.ImpRemoval
    children["rx-loc"] = &opticsAlarmInfo.RxLoc
    children["amp-gain-deg-low"] = &opticsAlarmInfo.AmpGainDegLow
    children["amp-gain-deg-high"] = &opticsAlarmInfo.AmpGainDegHigh
    children["txpwr-mismatch"] = &opticsAlarmInfo.TxpwrMismatch
    return children
}

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetYangName() string { return "optics-alarm-info" }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) SetParent(parent types.Entity) { opticsAlarmInfo.parent = parent }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetParent() types.Entity { return opticsAlarmInfo.parent }

func (opticsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo) GetParentYangName() string { return "optics-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower
// High Rx Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetFilter() yfilter.YFilter { return highRxPower.YFilter }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) SetFilter(yf yfilter.YFilter) { highRxPower.YFilter = yf }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetSegmentPath() string {
    return "high-rx-power"
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highRxPower.IsDetected
    leafs["counter"] = highRxPower.Counter
    return leafs
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetBundleName() string { return "cisco_ios_xr" }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetYangName() string { return "high-rx-power" }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) SetParent(parent types.Entity) { highRxPower.parent = parent }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetParent() types.Entity { return highRxPower.parent }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRxPower) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower
// Low Rx Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetFilter() yfilter.YFilter { return lowRxPower.YFilter }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) SetFilter(yf yfilter.YFilter) { lowRxPower.YFilter = yf }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetSegmentPath() string {
    return "low-rx-power"
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowRxPower.IsDetected
    leafs["counter"] = lowRxPower.Counter
    return leafs
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetBundleName() string { return "cisco_ios_xr" }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetYangName() string { return "low-rx-power" }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) SetParent(parent types.Entity) { lowRxPower.parent = parent }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetParent() types.Entity { return lowRxPower.parent }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRxPower) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower
// High Tx Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetFilter() yfilter.YFilter { return highTxPower.YFilter }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) SetFilter(yf yfilter.YFilter) { highTxPower.YFilter = yf }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetSegmentPath() string {
    return "high-tx-power"
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTxPower.IsDetected
    leafs["counter"] = highTxPower.Counter
    return leafs
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetBundleName() string { return "cisco_ios_xr" }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetYangName() string { return "high-tx-power" }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) SetParent(parent types.Entity) { highTxPower.parent = parent }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetParent() types.Entity { return highTxPower.parent }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTxPower) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower
// Low Tx Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetFilter() yfilter.YFilter { return lowTxPower.YFilter }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) SetFilter(yf yfilter.YFilter) { lowTxPower.YFilter = yf }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetSegmentPath() string {
    return "low-tx-power"
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTxPower.IsDetected
    leafs["counter"] = lowTxPower.Counter
    return leafs
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetBundleName() string { return "cisco_ios_xr" }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetYangName() string { return "low-tx-power" }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) SetParent(parent types.Entity) { lowTxPower.parent = parent }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetParent() types.Entity { return lowTxPower.parent }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTxPower) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc
// High laser bias current in units of percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetFilter() yfilter.YFilter { return highLbc.YFilter }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) SetFilter(yf yfilter.YFilter) { highLbc.YFilter = yf }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetSegmentPath() string {
    return "high-lbc"
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highLbc.IsDetected
    leafs["counter"] = highLbc.Counter
    return leafs
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetBundleName() string { return "cisco_ios_xr" }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetYangName() string { return "high-lbc" }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) SetParent(parent types.Entity) { highLbc.parent = parent }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetParent() types.Entity { return highLbc.parent }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighLbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power
// High Rx1 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetFilter() yfilter.YFilter { return highRx1Power.YFilter }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) SetFilter(yf yfilter.YFilter) { highRx1Power.YFilter = yf }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetSegmentPath() string {
    return "high-rx1-power"
}

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highRx1Power.IsDetected
    leafs["counter"] = highRx1Power.Counter
    return leafs
}

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetBundleName() string { return "cisco_ios_xr" }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetYangName() string { return "high-rx1-power" }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) SetParent(parent types.Entity) { highRx1Power.parent = parent }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetParent() types.Entity { return highRx1Power.parent }

func (highRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx1Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power
// High Rx2 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetFilter() yfilter.YFilter { return highRx2Power.YFilter }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) SetFilter(yf yfilter.YFilter) { highRx2Power.YFilter = yf }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetSegmentPath() string {
    return "high-rx2-power"
}

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highRx2Power.IsDetected
    leafs["counter"] = highRx2Power.Counter
    return leafs
}

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetBundleName() string { return "cisco_ios_xr" }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetYangName() string { return "high-rx2-power" }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) SetParent(parent types.Entity) { highRx2Power.parent = parent }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetParent() types.Entity { return highRx2Power.parent }

func (highRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx2Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power
// High Rx3 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetFilter() yfilter.YFilter { return highRx3Power.YFilter }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) SetFilter(yf yfilter.YFilter) { highRx3Power.YFilter = yf }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetSegmentPath() string {
    return "high-rx3-power"
}

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highRx3Power.IsDetected
    leafs["counter"] = highRx3Power.Counter
    return leafs
}

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetBundleName() string { return "cisco_ios_xr" }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetYangName() string { return "high-rx3-power" }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) SetParent(parent types.Entity) { highRx3Power.parent = parent }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetParent() types.Entity { return highRx3Power.parent }

func (highRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx3Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power
// High Rx4 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetFilter() yfilter.YFilter { return highRx4Power.YFilter }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) SetFilter(yf yfilter.YFilter) { highRx4Power.YFilter = yf }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetSegmentPath() string {
    return "high-rx4-power"
}

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highRx4Power.IsDetected
    leafs["counter"] = highRx4Power.Counter
    return leafs
}

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetBundleName() string { return "cisco_ios_xr" }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetYangName() string { return "high-rx4-power" }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) SetParent(parent types.Entity) { highRx4Power.parent = parent }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetParent() types.Entity { return highRx4Power.parent }

func (highRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighRx4Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power
// Low Rx1 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetFilter() yfilter.YFilter { return lowRx1Power.YFilter }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) SetFilter(yf yfilter.YFilter) { lowRx1Power.YFilter = yf }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetSegmentPath() string {
    return "low-rx1-power"
}

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowRx1Power.IsDetected
    leafs["counter"] = lowRx1Power.Counter
    return leafs
}

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetBundleName() string { return "cisco_ios_xr" }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetYangName() string { return "low-rx1-power" }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) SetParent(parent types.Entity) { lowRx1Power.parent = parent }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetParent() types.Entity { return lowRx1Power.parent }

func (lowRx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx1Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power
// Low Rx2 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetFilter() yfilter.YFilter { return lowRx2Power.YFilter }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) SetFilter(yf yfilter.YFilter) { lowRx2Power.YFilter = yf }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetSegmentPath() string {
    return "low-rx2-power"
}

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowRx2Power.IsDetected
    leafs["counter"] = lowRx2Power.Counter
    return leafs
}

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetBundleName() string { return "cisco_ios_xr" }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetYangName() string { return "low-rx2-power" }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) SetParent(parent types.Entity) { lowRx2Power.parent = parent }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetParent() types.Entity { return lowRx2Power.parent }

func (lowRx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx2Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power
// Low Rx3 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetFilter() yfilter.YFilter { return lowRx3Power.YFilter }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) SetFilter(yf yfilter.YFilter) { lowRx3Power.YFilter = yf }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetSegmentPath() string {
    return "low-rx3-power"
}

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowRx3Power.IsDetected
    leafs["counter"] = lowRx3Power.Counter
    return leafs
}

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetBundleName() string { return "cisco_ios_xr" }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetYangName() string { return "low-rx3-power" }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) SetParent(parent types.Entity) { lowRx3Power.parent = parent }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetParent() types.Entity { return lowRx3Power.parent }

func (lowRx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx3Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power
// Low Rx4 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetFilter() yfilter.YFilter { return lowRx4Power.YFilter }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) SetFilter(yf yfilter.YFilter) { lowRx4Power.YFilter = yf }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetSegmentPath() string {
    return "low-rx4-power"
}

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowRx4Power.IsDetected
    leafs["counter"] = lowRx4Power.Counter
    return leafs
}

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetBundleName() string { return "cisco_ios_xr" }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetYangName() string { return "low-rx4-power" }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) SetParent(parent types.Entity) { lowRx4Power.parent = parent }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetParent() types.Entity { return lowRx4Power.parent }

func (lowRx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowRx4Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power
// High Tx1 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetFilter() yfilter.YFilter { return highTx1Power.YFilter }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) SetFilter(yf yfilter.YFilter) { highTx1Power.YFilter = yf }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetSegmentPath() string {
    return "high-tx1-power"
}

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTx1Power.IsDetected
    leafs["counter"] = highTx1Power.Counter
    return leafs
}

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetBundleName() string { return "cisco_ios_xr" }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetYangName() string { return "high-tx1-power" }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) SetParent(parent types.Entity) { highTx1Power.parent = parent }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetParent() types.Entity { return highTx1Power.parent }

func (highTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power
// High Tx2 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetFilter() yfilter.YFilter { return highTx2Power.YFilter }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) SetFilter(yf yfilter.YFilter) { highTx2Power.YFilter = yf }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetSegmentPath() string {
    return "high-tx2-power"
}

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTx2Power.IsDetected
    leafs["counter"] = highTx2Power.Counter
    return leafs
}

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetBundleName() string { return "cisco_ios_xr" }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetYangName() string { return "high-tx2-power" }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) SetParent(parent types.Entity) { highTx2Power.parent = parent }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetParent() types.Entity { return highTx2Power.parent }

func (highTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power
// High Tx3 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetFilter() yfilter.YFilter { return highTx3Power.YFilter }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) SetFilter(yf yfilter.YFilter) { highTx3Power.YFilter = yf }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetSegmentPath() string {
    return "high-tx3-power"
}

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTx3Power.IsDetected
    leafs["counter"] = highTx3Power.Counter
    return leafs
}

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetBundleName() string { return "cisco_ios_xr" }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetYangName() string { return "high-tx3-power" }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) SetParent(parent types.Entity) { highTx3Power.parent = parent }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetParent() types.Entity { return highTx3Power.parent }

func (highTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power
// High Tx4 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetFilter() yfilter.YFilter { return highTx4Power.YFilter }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) SetFilter(yf yfilter.YFilter) { highTx4Power.YFilter = yf }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetSegmentPath() string {
    return "high-tx4-power"
}

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTx4Power.IsDetected
    leafs["counter"] = highTx4Power.Counter
    return leafs
}

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetBundleName() string { return "cisco_ios_xr" }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetYangName() string { return "high-tx4-power" }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) SetParent(parent types.Entity) { highTx4Power.parent = parent }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetParent() types.Entity { return highTx4Power.parent }

func (highTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power
// Low Tx1 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetFilter() yfilter.YFilter { return lowTx1Power.YFilter }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) SetFilter(yf yfilter.YFilter) { lowTx1Power.YFilter = yf }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetSegmentPath() string {
    return "low-tx1-power"
}

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTx1Power.IsDetected
    leafs["counter"] = lowTx1Power.Counter
    return leafs
}

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetBundleName() string { return "cisco_ios_xr" }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetYangName() string { return "low-tx1-power" }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) SetParent(parent types.Entity) { lowTx1Power.parent = parent }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetParent() types.Entity { return lowTx1Power.parent }

func (lowTx1Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power
// Low Tx2 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetFilter() yfilter.YFilter { return lowTx2Power.YFilter }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) SetFilter(yf yfilter.YFilter) { lowTx2Power.YFilter = yf }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetSegmentPath() string {
    return "low-tx2-power"
}

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTx2Power.IsDetected
    leafs["counter"] = lowTx2Power.Counter
    return leafs
}

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetBundleName() string { return "cisco_ios_xr" }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetYangName() string { return "low-tx2-power" }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) SetParent(parent types.Entity) { lowTx2Power.parent = parent }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetParent() types.Entity { return lowTx2Power.parent }

func (lowTx2Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power
// Low Tx3 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetFilter() yfilter.YFilter { return lowTx3Power.YFilter }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) SetFilter(yf yfilter.YFilter) { lowTx3Power.YFilter = yf }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetSegmentPath() string {
    return "low-tx3-power"
}

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTx3Power.IsDetected
    leafs["counter"] = lowTx3Power.Counter
    return leafs
}

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetBundleName() string { return "cisco_ios_xr" }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetYangName() string { return "low-tx3-power" }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) SetParent(parent types.Entity) { lowTx3Power.parent = parent }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetParent() types.Entity { return lowTx3Power.parent }

func (lowTx3Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power
// Low Tx4 Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetFilter() yfilter.YFilter { return lowTx4Power.YFilter }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) SetFilter(yf yfilter.YFilter) { lowTx4Power.YFilter = yf }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetSegmentPath() string {
    return "low-tx4-power"
}

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTx4Power.IsDetected
    leafs["counter"] = lowTx4Power.Counter
    return leafs
}

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetBundleName() string { return "cisco_ios_xr" }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetYangName() string { return "low-tx4-power" }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) SetParent(parent types.Entity) { lowTx4Power.parent = parent }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetParent() types.Entity { return lowTx4Power.parent }

func (lowTx4Power *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Power) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc
// High Tx1 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetFilter() yfilter.YFilter { return highTx1Lbc.YFilter }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) SetFilter(yf yfilter.YFilter) { highTx1Lbc.YFilter = yf }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetSegmentPath() string {
    return "high-tx1lbc"
}

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTx1Lbc.IsDetected
    leafs["counter"] = highTx1Lbc.Counter
    return leafs
}

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetBundleName() string { return "cisco_ios_xr" }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetYangName() string { return "high-tx1lbc" }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) SetParent(parent types.Entity) { highTx1Lbc.parent = parent }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetParent() types.Entity { return highTx1Lbc.parent }

func (highTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx1Lbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc
// High Tx2 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetFilter() yfilter.YFilter { return highTx2Lbc.YFilter }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) SetFilter(yf yfilter.YFilter) { highTx2Lbc.YFilter = yf }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetSegmentPath() string {
    return "high-tx2lbc"
}

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTx2Lbc.IsDetected
    leafs["counter"] = highTx2Lbc.Counter
    return leafs
}

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetBundleName() string { return "cisco_ios_xr" }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetYangName() string { return "high-tx2lbc" }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) SetParent(parent types.Entity) { highTx2Lbc.parent = parent }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetParent() types.Entity { return highTx2Lbc.parent }

func (highTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx2Lbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc
// High Tx3 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetFilter() yfilter.YFilter { return highTx3Lbc.YFilter }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) SetFilter(yf yfilter.YFilter) { highTx3Lbc.YFilter = yf }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetSegmentPath() string {
    return "high-tx3lbc"
}

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTx3Lbc.IsDetected
    leafs["counter"] = highTx3Lbc.Counter
    return leafs
}

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetBundleName() string { return "cisco_ios_xr" }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetYangName() string { return "high-tx3lbc" }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) SetParent(parent types.Entity) { highTx3Lbc.parent = parent }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetParent() types.Entity { return highTx3Lbc.parent }

func (highTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx3Lbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc
// High Tx4 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetFilter() yfilter.YFilter { return highTx4Lbc.YFilter }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) SetFilter(yf yfilter.YFilter) { highTx4Lbc.YFilter = yf }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetSegmentPath() string {
    return "high-tx4lbc"
}

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTx4Lbc.IsDetected
    leafs["counter"] = highTx4Lbc.Counter
    return leafs
}

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetBundleName() string { return "cisco_ios_xr" }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetYangName() string { return "high-tx4lbc" }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) SetParent(parent types.Entity) { highTx4Lbc.parent = parent }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetParent() types.Entity { return highTx4Lbc.parent }

func (highTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_HighTx4Lbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc
// Low Tx1 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetFilter() yfilter.YFilter { return lowTx1Lbc.YFilter }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) SetFilter(yf yfilter.YFilter) { lowTx1Lbc.YFilter = yf }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetSegmentPath() string {
    return "low-tx1lbc"
}

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTx1Lbc.IsDetected
    leafs["counter"] = lowTx1Lbc.Counter
    return leafs
}

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetBundleName() string { return "cisco_ios_xr" }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetYangName() string { return "low-tx1lbc" }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) SetParent(parent types.Entity) { lowTx1Lbc.parent = parent }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetParent() types.Entity { return lowTx1Lbc.parent }

func (lowTx1Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx1Lbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc
// Low Tx2 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetFilter() yfilter.YFilter { return lowTx2Lbc.YFilter }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) SetFilter(yf yfilter.YFilter) { lowTx2Lbc.YFilter = yf }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetSegmentPath() string {
    return "low-tx2lbc"
}

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTx2Lbc.IsDetected
    leafs["counter"] = lowTx2Lbc.Counter
    return leafs
}

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetBundleName() string { return "cisco_ios_xr" }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetYangName() string { return "low-tx2lbc" }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) SetParent(parent types.Entity) { lowTx2Lbc.parent = parent }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetParent() types.Entity { return lowTx2Lbc.parent }

func (lowTx2Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx2Lbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc
// Low Tx3 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetFilter() yfilter.YFilter { return lowTx3Lbc.YFilter }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) SetFilter(yf yfilter.YFilter) { lowTx3Lbc.YFilter = yf }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetSegmentPath() string {
    return "low-tx3lbc"
}

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTx3Lbc.IsDetected
    leafs["counter"] = lowTx3Lbc.Counter
    return leafs
}

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetBundleName() string { return "cisco_ios_xr" }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetYangName() string { return "low-tx3lbc" }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) SetParent(parent types.Entity) { lowTx3Lbc.parent = parent }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetParent() types.Entity { return lowTx3Lbc.parent }

func (lowTx3Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx3Lbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc
// Low Tx4 laser bias current in units of
// percentage
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetFilter() yfilter.YFilter { return lowTx4Lbc.YFilter }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) SetFilter(yf yfilter.YFilter) { lowTx4Lbc.YFilter = yf }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetSegmentPath() string {
    return "low-tx4lbc"
}

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTx4Lbc.IsDetected
    leafs["counter"] = lowTx4Lbc.Counter
    return leafs
}

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetBundleName() string { return "cisco_ios_xr" }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetYangName() string { return "low-tx4lbc" }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) SetParent(parent types.Entity) { lowTx4Lbc.parent = parent }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetParent() types.Entity { return lowTx4Lbc.parent }

func (lowTx4Lbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_LowTx4Lbc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos
// RX LOS
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetFilter() yfilter.YFilter { return rxLos.YFilter }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) SetFilter(yf yfilter.YFilter) { rxLos.YFilter = yf }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetSegmentPath() string {
    return "rx-los"
}

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = rxLos.IsDetected
    leafs["counter"] = rxLos.Counter
    return leafs
}

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetBundleName() string { return "cisco_ios_xr" }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetYangName() string { return "rx-los" }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) SetParent(parent types.Entity) { rxLos.parent = parent }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetParent() types.Entity { return rxLos.parent }

func (rxLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLos) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos
// TX LOS
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetFilter() yfilter.YFilter { return txLos.YFilter }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) SetFilter(yf yfilter.YFilter) { txLos.YFilter = yf }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetSegmentPath() string {
    return "tx-los"
}

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = txLos.IsDetected
    leafs["counter"] = txLos.Counter
    return leafs
}

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetBundleName() string { return "cisco_ios_xr" }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetYangName() string { return "tx-los" }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) SetParent(parent types.Entity) { txLos.parent = parent }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetParent() types.Entity { return txLos.parent }

func (txLos *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLos) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol
// RX LOL
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetFilter() yfilter.YFilter { return rxLol.YFilter }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) SetFilter(yf yfilter.YFilter) { rxLol.YFilter = yf }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetSegmentPath() string {
    return "rx-lol"
}

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = rxLol.IsDetected
    leafs["counter"] = rxLol.Counter
    return leafs
}

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetBundleName() string { return "cisco_ios_xr" }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetYangName() string { return "rx-lol" }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) SetParent(parent types.Entity) { rxLol.parent = parent }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetParent() types.Entity { return rxLol.parent }

func (rxLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLol) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol
// TX LOL
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetFilter() yfilter.YFilter { return txLol.YFilter }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) SetFilter(yf yfilter.YFilter) { txLol.YFilter = yf }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetSegmentPath() string {
    return "tx-lol"
}

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = txLol.IsDetected
    leafs["counter"] = txLol.Counter
    return leafs
}

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetBundleName() string { return "cisco_ios_xr" }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetYangName() string { return "tx-lol" }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) SetParent(parent types.Entity) { txLol.parent = parent }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetParent() types.Entity { return txLol.parent }

func (txLol *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxLol) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault
// TX Fault
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetFilter() yfilter.YFilter { return txFault.YFilter }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) SetFilter(yf yfilter.YFilter) { txFault.YFilter = yf }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetSegmentPath() string {
    return "tx-fault"
}

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = txFault.IsDetected
    leafs["counter"] = txFault.Counter
    return leafs
}

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetBundleName() string { return "cisco_ios_xr" }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetYangName() string { return "tx-fault" }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) SetParent(parent types.Entity) { txFault.parent = parent }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetParent() types.Entity { return txFault.parent }

func (txFault *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxFault) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd
// HI DGD
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetFilter() yfilter.YFilter { return hidgd.YFilter }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) SetFilter(yf yfilter.YFilter) { hidgd.YFilter = yf }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetSegmentPath() string {
    return "hidgd"
}

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = hidgd.IsDetected
    leafs["counter"] = hidgd.Counter
    return leafs
}

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetBundleName() string { return "cisco_ios_xr" }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetYangName() string { return "hidgd" }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) SetParent(parent types.Entity) { hidgd.parent = parent }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetParent() types.Entity { return hidgd.parent }

func (hidgd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Hidgd) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd
// OOR CD
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetFilter() yfilter.YFilter { return oorcd.YFilter }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) SetFilter(yf yfilter.YFilter) { oorcd.YFilter = yf }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetSegmentPath() string {
    return "oorcd"
}

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = oorcd.IsDetected
    leafs["counter"] = oorcd.Counter
    return leafs
}

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetBundleName() string { return "cisco_ios_xr" }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetYangName() string { return "oorcd" }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) SetParent(parent types.Entity) { oorcd.parent = parent }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetParent() types.Entity { return oorcd.parent }

func (oorcd *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Oorcd) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr
// OSNR
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetFilter() yfilter.YFilter { return osnr.YFilter }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) SetFilter(yf yfilter.YFilter) { osnr.YFilter = yf }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetSegmentPath() string {
    return "osnr"
}

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = osnr.IsDetected
    leafs["counter"] = osnr.Counter
    return leafs
}

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetBundleName() string { return "cisco_ios_xr" }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetYangName() string { return "osnr" }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) SetParent(parent types.Entity) { osnr.parent = parent }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetParent() types.Entity { return osnr.parent }

func (osnr *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Osnr) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool
// WVL OOL
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetFilter() yfilter.YFilter { return wvlool.YFilter }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) SetFilter(yf yfilter.YFilter) { wvlool.YFilter = yf }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetSegmentPath() string {
    return "wvlool"
}

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = wvlool.IsDetected
    leafs["counter"] = wvlool.Counter
    return leafs
}

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetBundleName() string { return "cisco_ios_xr" }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetYangName() string { return "wvlool" }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) SetParent(parent types.Entity) { wvlool.parent = parent }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetParent() types.Entity { return wvlool.parent }

func (wvlool *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Wvlool) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea
// MEA
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetFilter() yfilter.YFilter { return mea.YFilter }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) SetFilter(yf yfilter.YFilter) { mea.YFilter = yf }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetSegmentPath() string {
    return "mea"
}

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = mea.IsDetected
    leafs["counter"] = mea.Counter
    return leafs
}

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetBundleName() string { return "cisco_ios_xr" }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetYangName() string { return "mea" }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) SetParent(parent types.Entity) { mea.parent = parent }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetParent() types.Entity { return mea.parent }

func (mea *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_Mea) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval
// IMPROPER REM
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetFilter() yfilter.YFilter { return impRemoval.YFilter }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) SetFilter(yf yfilter.YFilter) { impRemoval.YFilter = yf }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetSegmentPath() string {
    return "imp-removal"
}

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = impRemoval.IsDetected
    leafs["counter"] = impRemoval.Counter
    return leafs
}

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetBundleName() string { return "cisco_ios_xr" }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetYangName() string { return "imp-removal" }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) SetParent(parent types.Entity) { impRemoval.parent = parent }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetParent() types.Entity { return impRemoval.parent }

func (impRemoval *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_ImpRemoval) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc
// Rx LOC
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetFilter() yfilter.YFilter { return rxLoc.YFilter }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) SetFilter(yf yfilter.YFilter) { rxLoc.YFilter = yf }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetSegmentPath() string {
    return "rx-loc"
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = rxLoc.IsDetected
    leafs["counter"] = rxLoc.Counter
    return leafs
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetBundleName() string { return "cisco_ios_xr" }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetYangName() string { return "rx-loc" }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) SetParent(parent types.Entity) { rxLoc.parent = parent }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetParent() types.Entity { return rxLoc.parent }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_RxLoc) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow
// Ampli Gain Deg Low
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetFilter() yfilter.YFilter { return ampGainDegLow.YFilter }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) SetFilter(yf yfilter.YFilter) { ampGainDegLow.YFilter = yf }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetSegmentPath() string {
    return "amp-gain-deg-low"
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = ampGainDegLow.IsDetected
    leafs["counter"] = ampGainDegLow.Counter
    return leafs
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetBundleName() string { return "cisco_ios_xr" }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetYangName() string { return "amp-gain-deg-low" }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) SetParent(parent types.Entity) { ampGainDegLow.parent = parent }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetParent() types.Entity { return ampGainDegLow.parent }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegLow) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh
// Ampli Gain Deg High
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetFilter() yfilter.YFilter { return ampGainDegHigh.YFilter }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) SetFilter(yf yfilter.YFilter) { ampGainDegHigh.YFilter = yf }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetSegmentPath() string {
    return "amp-gain-deg-high"
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = ampGainDegHigh.IsDetected
    leafs["counter"] = ampGainDegHigh.Counter
    return leafs
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetBundleName() string { return "cisco_ios_xr" }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetYangName() string { return "amp-gain-deg-high" }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) SetParent(parent types.Entity) { ampGainDegHigh.parent = parent }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetParent() types.Entity { return ampGainDegHigh.parent }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_AmpGainDegHigh) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch
// TX POWER PROV MISMATCH
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetFilter() yfilter.YFilter { return txpwrMismatch.YFilter }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) SetFilter(yf yfilter.YFilter) { txpwrMismatch.YFilter = yf }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetSegmentPath() string {
    return "txpwr-mismatch"
}

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = txpwrMismatch.IsDetected
    leafs["counter"] = txpwrMismatch.Counter
    return leafs
}

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetBundleName() string { return "cisco_ios_xr" }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetYangName() string { return "txpwr-mismatch" }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) SetParent(parent types.Entity) { txpwrMismatch.parent = parent }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetParent() types.Entity { return txpwrMismatch.parent }

func (txpwrMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OpticsAlarmInfo_TxpwrMismatch) GetParentYangName() string { return "optics-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo
// Ots Alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Low Tx Power in uints of 0.1 dBm.
    LowTxPower OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower

    // Low Rx Power in uints of 0.1 dBm.
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

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetFilter() yfilter.YFilter { return otsAlarmInfo.YFilter }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) SetFilter(yf yfilter.YFilter) { otsAlarmInfo.YFilter = yf }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetGoName(yname string) string {
    if yname == "low-tx-power" { return "LowTxPower" }
    if yname == "low-rx-power" { return "LowRxPower" }
    if yname == "rx-los-p" { return "RxLosP" }
    if yname == "rx-loc" { return "RxLoc" }
    if yname == "amp-gain-deg-low" { return "AmpGainDegLow" }
    if yname == "amp-gain-deg-high" { return "AmpGainDegHigh" }
    if yname == "auto-laser-shut" { return "AutoLaserShut" }
    if yname == "auto-power-red" { return "AutoPowerRed" }
    if yname == "auto-ampli-ctrl-disabled" { return "AutoAmpliCtrlDisabled" }
    if yname == "auto-ampli-ctrl-config-mismatch" { return "AutoAmpliCtrlConfigMismatch" }
    if yname == "switch-to-protect" { return "SwitchToProtect" }
    if yname == "auto-ampli-ctrl-running" { return "AutoAmpliCtrlRunning" }
    return ""
}

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetSegmentPath() string {
    return "ots-alarm-info"
}

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "low-tx-power" {
        return &otsAlarmInfo.LowTxPower
    }
    if childYangName == "low-rx-power" {
        return &otsAlarmInfo.LowRxPower
    }
    if childYangName == "rx-los-p" {
        return &otsAlarmInfo.RxLosP
    }
    if childYangName == "rx-loc" {
        return &otsAlarmInfo.RxLoc
    }
    if childYangName == "amp-gain-deg-low" {
        return &otsAlarmInfo.AmpGainDegLow
    }
    if childYangName == "amp-gain-deg-high" {
        return &otsAlarmInfo.AmpGainDegHigh
    }
    if childYangName == "auto-laser-shut" {
        return &otsAlarmInfo.AutoLaserShut
    }
    if childYangName == "auto-power-red" {
        return &otsAlarmInfo.AutoPowerRed
    }
    if childYangName == "auto-ampli-ctrl-disabled" {
        return &otsAlarmInfo.AutoAmpliCtrlDisabled
    }
    if childYangName == "auto-ampli-ctrl-config-mismatch" {
        return &otsAlarmInfo.AutoAmpliCtrlConfigMismatch
    }
    if childYangName == "switch-to-protect" {
        return &otsAlarmInfo.SwitchToProtect
    }
    if childYangName == "auto-ampli-ctrl-running" {
        return &otsAlarmInfo.AutoAmpliCtrlRunning
    }
    return nil
}

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["low-tx-power"] = &otsAlarmInfo.LowTxPower
    children["low-rx-power"] = &otsAlarmInfo.LowRxPower
    children["rx-los-p"] = &otsAlarmInfo.RxLosP
    children["rx-loc"] = &otsAlarmInfo.RxLoc
    children["amp-gain-deg-low"] = &otsAlarmInfo.AmpGainDegLow
    children["amp-gain-deg-high"] = &otsAlarmInfo.AmpGainDegHigh
    children["auto-laser-shut"] = &otsAlarmInfo.AutoLaserShut
    children["auto-power-red"] = &otsAlarmInfo.AutoPowerRed
    children["auto-ampli-ctrl-disabled"] = &otsAlarmInfo.AutoAmpliCtrlDisabled
    children["auto-ampli-ctrl-config-mismatch"] = &otsAlarmInfo.AutoAmpliCtrlConfigMismatch
    children["switch-to-protect"] = &otsAlarmInfo.SwitchToProtect
    children["auto-ampli-ctrl-running"] = &otsAlarmInfo.AutoAmpliCtrlRunning
    return children
}

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetYangName() string { return "ots-alarm-info" }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) SetParent(parent types.Entity) { otsAlarmInfo.parent = parent }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetParent() types.Entity { return otsAlarmInfo.parent }

func (otsAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo) GetParentYangName() string { return "optics-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower
// Low Tx Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetFilter() yfilter.YFilter { return lowTxPower.YFilter }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) SetFilter(yf yfilter.YFilter) { lowTxPower.YFilter = yf }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetSegmentPath() string {
    return "low-tx-power"
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTxPower.IsDetected
    leafs["counter"] = lowTxPower.Counter
    return leafs
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetBundleName() string { return "cisco_ios_xr" }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetYangName() string { return "low-tx-power" }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) SetParent(parent types.Entity) { lowTxPower.parent = parent }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetParent() types.Entity { return lowTxPower.parent }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowTxPower) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower
// Low Rx Power in uints of 0.1 dBm
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetFilter() yfilter.YFilter { return lowRxPower.YFilter }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) SetFilter(yf yfilter.YFilter) { lowRxPower.YFilter = yf }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetSegmentPath() string {
    return "low-rx-power"
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowRxPower.IsDetected
    leafs["counter"] = lowRxPower.Counter
    return leafs
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetBundleName() string { return "cisco_ios_xr" }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetYangName() string { return "low-rx-power" }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) SetParent(parent types.Entity) { lowRxPower.parent = parent }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetParent() types.Entity { return lowRxPower.parent }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_LowRxPower) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP
// Rx LOS_P
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetFilter() yfilter.YFilter { return rxLosP.YFilter }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) SetFilter(yf yfilter.YFilter) { rxLosP.YFilter = yf }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetSegmentPath() string {
    return "rx-los-p"
}

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = rxLosP.IsDetected
    leafs["counter"] = rxLosP.Counter
    return leafs
}

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetBundleName() string { return "cisco_ios_xr" }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetYangName() string { return "rx-los-p" }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) SetParent(parent types.Entity) { rxLosP.parent = parent }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetParent() types.Entity { return rxLosP.parent }

func (rxLosP *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLosP) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc
// Rx LOC
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetFilter() yfilter.YFilter { return rxLoc.YFilter }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) SetFilter(yf yfilter.YFilter) { rxLoc.YFilter = yf }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetSegmentPath() string {
    return "rx-loc"
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = rxLoc.IsDetected
    leafs["counter"] = rxLoc.Counter
    return leafs
}

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetBundleName() string { return "cisco_ios_xr" }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetYangName() string { return "rx-loc" }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) SetParent(parent types.Entity) { rxLoc.parent = parent }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetParent() types.Entity { return rxLoc.parent }

func (rxLoc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_RxLoc) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow
// Ampli Gain Deg Low
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetFilter() yfilter.YFilter { return ampGainDegLow.YFilter }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) SetFilter(yf yfilter.YFilter) { ampGainDegLow.YFilter = yf }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetSegmentPath() string {
    return "amp-gain-deg-low"
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = ampGainDegLow.IsDetected
    leafs["counter"] = ampGainDegLow.Counter
    return leafs
}

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetBundleName() string { return "cisco_ios_xr" }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetYangName() string { return "amp-gain-deg-low" }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) SetParent(parent types.Entity) { ampGainDegLow.parent = parent }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetParent() types.Entity { return ampGainDegLow.parent }

func (ampGainDegLow *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegLow) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh
// Ampli Gain Deg High
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetFilter() yfilter.YFilter { return ampGainDegHigh.YFilter }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) SetFilter(yf yfilter.YFilter) { ampGainDegHigh.YFilter = yf }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetSegmentPath() string {
    return "amp-gain-deg-high"
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = ampGainDegHigh.IsDetected
    leafs["counter"] = ampGainDegHigh.Counter
    return leafs
}

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetBundleName() string { return "cisco_ios_xr" }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetYangName() string { return "amp-gain-deg-high" }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) SetParent(parent types.Entity) { ampGainDegHigh.parent = parent }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetParent() types.Entity { return ampGainDegHigh.parent }

func (ampGainDegHigh *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AmpGainDegHigh) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut
// Auto Laser Shut
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetFilter() yfilter.YFilter { return autoLaserShut.YFilter }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) SetFilter(yf yfilter.YFilter) { autoLaserShut.YFilter = yf }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetSegmentPath() string {
    return "auto-laser-shut"
}

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = autoLaserShut.IsDetected
    leafs["counter"] = autoLaserShut.Counter
    return leafs
}

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetBundleName() string { return "cisco_ios_xr" }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetYangName() string { return "auto-laser-shut" }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) SetParent(parent types.Entity) { autoLaserShut.parent = parent }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetParent() types.Entity { return autoLaserShut.parent }

func (autoLaserShut *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoLaserShut) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed
// Auto Power Red
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetFilter() yfilter.YFilter { return autoPowerRed.YFilter }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) SetFilter(yf yfilter.YFilter) { autoPowerRed.YFilter = yf }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetSegmentPath() string {
    return "auto-power-red"
}

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = autoPowerRed.IsDetected
    leafs["counter"] = autoPowerRed.Counter
    return leafs
}

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetBundleName() string { return "cisco_ios_xr" }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetYangName() string { return "auto-power-red" }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) SetParent(parent types.Entity) { autoPowerRed.parent = parent }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetParent() types.Entity { return autoPowerRed.parent }

func (autoPowerRed *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoPowerRed) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled
// Auto Ampli Ctrl Disabled
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetFilter() yfilter.YFilter { return autoAmpliCtrlDisabled.YFilter }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) SetFilter(yf yfilter.YFilter) { autoAmpliCtrlDisabled.YFilter = yf }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetSegmentPath() string {
    return "auto-ampli-ctrl-disabled"
}

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = autoAmpliCtrlDisabled.IsDetected
    leafs["counter"] = autoAmpliCtrlDisabled.Counter
    return leafs
}

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetBundleName() string { return "cisco_ios_xr" }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetYangName() string { return "auto-ampli-ctrl-disabled" }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) SetParent(parent types.Entity) { autoAmpliCtrlDisabled.parent = parent }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetParent() types.Entity { return autoAmpliCtrlDisabled.parent }

func (autoAmpliCtrlDisabled *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlDisabled) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch
// Auto Ampli Ctrl Config Mismatch
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetFilter() yfilter.YFilter { return autoAmpliCtrlConfigMismatch.YFilter }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) SetFilter(yf yfilter.YFilter) { autoAmpliCtrlConfigMismatch.YFilter = yf }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetSegmentPath() string {
    return "auto-ampli-ctrl-config-mismatch"
}

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = autoAmpliCtrlConfigMismatch.IsDetected
    leafs["counter"] = autoAmpliCtrlConfigMismatch.Counter
    return leafs
}

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetBundleName() string { return "cisco_ios_xr" }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetYangName() string { return "auto-ampli-ctrl-config-mismatch" }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) SetParent(parent types.Entity) { autoAmpliCtrlConfigMismatch.parent = parent }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetParent() types.Entity { return autoAmpliCtrlConfigMismatch.parent }

func (autoAmpliCtrlConfigMismatch *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlConfigMismatch) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect
// Switch To Protect
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetFilter() yfilter.YFilter { return switchToProtect.YFilter }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) SetFilter(yf yfilter.YFilter) { switchToProtect.YFilter = yf }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetSegmentPath() string {
    return "switch-to-protect"
}

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = switchToProtect.IsDetected
    leafs["counter"] = switchToProtect.Counter
    return leafs
}

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetBundleName() string { return "cisco_ios_xr" }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetYangName() string { return "switch-to-protect" }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) SetParent(parent types.Entity) { switchToProtect.parent = parent }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetParent() types.Entity { return switchToProtect.parent }

func (switchToProtect *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_SwitchToProtect) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning
// Auto Ampli Ctrl Running
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetFilter() yfilter.YFilter { return autoAmpliCtrlRunning.YFilter }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) SetFilter(yf yfilter.YFilter) { autoAmpliCtrlRunning.YFilter = yf }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetSegmentPath() string {
    return "auto-ampli-ctrl-running"
}

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = autoAmpliCtrlRunning.IsDetected
    leafs["counter"] = autoAmpliCtrlRunning.Counter
    return leafs
}

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetBundleName() string { return "cisco_ios_xr" }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetYangName() string { return "auto-ampli-ctrl-running" }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) SetParent(parent types.Entity) { autoAmpliCtrlRunning.parent = parent }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetParent() types.Entity { return autoAmpliCtrlRunning.parent }

func (autoAmpliCtrlRunning *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_OtsAlarmInfo_AutoAmpliCtrlRunning) GetParentYangName() string { return "ots-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo
// Transceiver Vendor Details
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo struct {
    parent types.Entity
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

    // Internal Temperature in C. The type is interface{} with range:
    // -2147483648..2147483647.
    InternalTemperature interface{}
}

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetFilter() yfilter.YFilter { return transceiverInfo.YFilter }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) SetFilter(yf yfilter.YFilter) { transceiverInfo.YFilter = yf }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetGoName(yname string) string {
    if yname == "vendor-info" { return "VendorInfo" }
    if yname == "adapter-vendor-info" { return "AdapterVendorInfo" }
    if yname == "date" { return "Date" }
    if yname == "optics-vendor-rev" { return "OpticsVendorRev" }
    if yname == "optics-serial-no" { return "OpticsSerialNo" }
    if yname == "optics-vendor-part" { return "OpticsVendorPart" }
    if yname == "optics-type" { return "OpticsType" }
    if yname == "vendor-name" { return "VendorName" }
    if yname == "oui-no" { return "OuiNo" }
    if yname == "optics-pid" { return "OpticsPid" }
    if yname == "optics-vid" { return "OpticsVid" }
    if yname == "connector-type" { return "ConnectorType" }
    if yname == "otn-application-code" { return "OtnApplicationCode" }
    if yname == "sonet-application-code" { return "SonetApplicationCode" }
    if yname == "ethernet-compliance-code" { return "EthernetComplianceCode" }
    if yname == "internal-temperature" { return "InternalTemperature" }
    return ""
}

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetSegmentPath() string {
    return "transceiver-info"
}

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vendor-info"] = transceiverInfo.VendorInfo
    leafs["adapter-vendor-info"] = transceiverInfo.AdapterVendorInfo
    leafs["date"] = transceiverInfo.Date
    leafs["optics-vendor-rev"] = transceiverInfo.OpticsVendorRev
    leafs["optics-serial-no"] = transceiverInfo.OpticsSerialNo
    leafs["optics-vendor-part"] = transceiverInfo.OpticsVendorPart
    leafs["optics-type"] = transceiverInfo.OpticsType
    leafs["vendor-name"] = transceiverInfo.VendorName
    leafs["oui-no"] = transceiverInfo.OuiNo
    leafs["optics-pid"] = transceiverInfo.OpticsPid
    leafs["optics-vid"] = transceiverInfo.OpticsVid
    leafs["connector-type"] = transceiverInfo.ConnectorType
    leafs["otn-application-code"] = transceiverInfo.OtnApplicationCode
    leafs["sonet-application-code"] = transceiverInfo.SonetApplicationCode
    leafs["ethernet-compliance-code"] = transceiverInfo.EthernetComplianceCode
    leafs["internal-temperature"] = transceiverInfo.InternalTemperature
    return leafs
}

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetBundleName() string { return "cisco_ios_xr" }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetYangName() string { return "transceiver-info" }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) SetParent(parent types.Entity) { transceiverInfo.parent = parent }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetParent() types.Entity { return transceiverInfo.parent }

func (transceiverInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_TransceiverInfo) GetParentYangName() string { return "optics-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal
// Extended optics parameters
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal struct {
    parent types.Entity
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

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetFilter() yfilter.YFilter { return extParamVal.YFilter }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) SetFilter(yf yfilter.YFilter) { extParamVal.YFilter = yf }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetGoName(yname string) string {
    if yname == "snr-lane1" { return "SnrLane1" }
    if yname == "snr-lane2" { return "SnrLane2" }
    if yname == "isi-correction-lane1" { return "IsiCorrectionLane1" }
    if yname == "isi-correction-lane2" { return "IsiCorrectionLane2" }
    if yname == "pam-rate-lane1" { return "PamRateLane1" }
    if yname == "pam-rate-lane2" { return "PamRateLane2" }
    if yname == "pre-fec-ber" { return "PreFecBer" }
    if yname == "uncorrected-ber" { return "UncorrectedBer" }
    if yname == "tec-current-lane1" { return "TecCurrentLane1" }
    if yname == "tec-current-lane2" { return "TecCurrentLane2" }
    if yname == "laser-diff-frequency-lane1" { return "LaserDiffFrequencyLane1" }
    if yname == "laser-diff-frequency-lane2" { return "LaserDiffFrequencyLane2" }
    if yname == "laser-diff-temperature-lane1" { return "LaserDiffTemperatureLane1" }
    if yname == "laser-diff-temperature-lane2" { return "LaserDiffTemperatureLane2" }
    if yname == "pre-fec-ber-latched-min" { return "PreFecBerLatchedMin" }
    if yname == "pre-fec-ber-latched-max" { return "PreFecBerLatchedMax" }
    if yname == "pre-fec-ber-accumulated" { return "PreFecBerAccumulated" }
    if yname == "pre-fec-ber-instantaneous" { return "PreFecBerInstantaneous" }
    if yname == "uncorrected-ber-latched-min" { return "UncorrectedBerLatchedMin" }
    if yname == "uncorrected-ber-latched-max" { return "UncorrectedBerLatchedMax" }
    if yname == "uncorrected-ber-accumulated" { return "UncorrectedBerAccumulated" }
    if yname == "uncorrected-ber-instantaneous" { return "UncorrectedBerInstantaneous" }
    return ""
}

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetSegmentPath() string {
    return "ext-param-val"
}

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snr-lane1"] = extParamVal.SnrLane1
    leafs["snr-lane2"] = extParamVal.SnrLane2
    leafs["isi-correction-lane1"] = extParamVal.IsiCorrectionLane1
    leafs["isi-correction-lane2"] = extParamVal.IsiCorrectionLane2
    leafs["pam-rate-lane1"] = extParamVal.PamRateLane1
    leafs["pam-rate-lane2"] = extParamVal.PamRateLane2
    leafs["pre-fec-ber"] = extParamVal.PreFecBer
    leafs["uncorrected-ber"] = extParamVal.UncorrectedBer
    leafs["tec-current-lane1"] = extParamVal.TecCurrentLane1
    leafs["tec-current-lane2"] = extParamVal.TecCurrentLane2
    leafs["laser-diff-frequency-lane1"] = extParamVal.LaserDiffFrequencyLane1
    leafs["laser-diff-frequency-lane2"] = extParamVal.LaserDiffFrequencyLane2
    leafs["laser-diff-temperature-lane1"] = extParamVal.LaserDiffTemperatureLane1
    leafs["laser-diff-temperature-lane2"] = extParamVal.LaserDiffTemperatureLane2
    leafs["pre-fec-ber-latched-min"] = extParamVal.PreFecBerLatchedMin
    leafs["pre-fec-ber-latched-max"] = extParamVal.PreFecBerLatchedMax
    leafs["pre-fec-ber-accumulated"] = extParamVal.PreFecBerAccumulated
    leafs["pre-fec-ber-instantaneous"] = extParamVal.PreFecBerInstantaneous
    leafs["uncorrected-ber-latched-min"] = extParamVal.UncorrectedBerLatchedMin
    leafs["uncorrected-ber-latched-max"] = extParamVal.UncorrectedBerLatchedMax
    leafs["uncorrected-ber-accumulated"] = extParamVal.UncorrectedBerAccumulated
    leafs["uncorrected-ber-instantaneous"] = extParamVal.UncorrectedBerInstantaneous
    return leafs
}

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetBundleName() string { return "cisco_ios_xr" }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetYangName() string { return "ext-param-val" }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) SetParent(parent types.Entity) { extParamVal.parent = parent }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetParent() types.Entity { return extParamVal.parent }

func (extParamVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamVal) GetParentYangName() string { return "optics-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal
// Extended optics parameters threshold values
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal struct {
    parent types.Entity
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

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetFilter() yfilter.YFilter { return extParamThresholdVal.YFilter }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) SetFilter(yf yfilter.YFilter) { extParamThresholdVal.YFilter = yf }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetGoName(yname string) string {
    if yname == "snr-alarm-high-threshold" { return "SnrAlarmHighThreshold" }
    if yname == "snr-alarm-low-threshold" { return "SnrAlarmLowThreshold" }
    if yname == "snr-warn-high-threshold" { return "SnrWarnHighThreshold" }
    if yname == "snr-warn-low-threshold" { return "SnrWarnLowThreshold" }
    if yname == "isi-correction-alarm-high-threshold" { return "IsiCorrectionAlarmHighThreshold" }
    if yname == "isi-correction-alarm-low-threshold" { return "IsiCorrectionAlarmLowThreshold" }
    if yname == "isi-correction-warn-high-threshold" { return "IsiCorrectionWarnHighThreshold" }
    if yname == "isi-correction-warn-low-threshold" { return "IsiCorrectionWarnLowThreshold" }
    if yname == "pam-rate-alarm-high-threshold" { return "PamRateAlarmHighThreshold" }
    if yname == "pam-rate-alarm-low-threshold" { return "PamRateAlarmLowThreshold" }
    if yname == "pam-rate-warn-high-threshold" { return "PamRateWarnHighThreshold" }
    if yname == "pam-rate-warn-low-threshold" { return "PamRateWarnLowThreshold" }
    if yname == "pre-fec-ber-alarm-high-threshold" { return "PreFecBerAlarmHighThreshold" }
    if yname == "pre-fec-ber-alarm-low-threshold" { return "PreFecBerAlarmLowThreshold" }
    if yname == "pre-fec-ber-warn-high-threshold" { return "PreFecBerWarnHighThreshold" }
    if yname == "pre-fec-ber-warn-low-threshold" { return "PreFecBerWarnLowThreshold" }
    if yname == "uncorrected-ber-alarm-high-threshold" { return "UncorrectedBerAlarmHighThreshold" }
    if yname == "uncorrected-ber-alarm-low-threshold" { return "UncorrectedBerAlarmLowThreshold" }
    if yname == "uncorrected-ber-warn-high-threshold" { return "UncorrectedBerWarnHighThreshold" }
    if yname == "uncorrected-ber-warn-low-threshold" { return "UncorrectedBerWarnLowThreshold" }
    if yname == "tec-current-alarm-high-threshold" { return "TecCurrentAlarmHighThreshold" }
    if yname == "tec-current-alarm-low-threshold" { return "TecCurrentAlarmLowThreshold" }
    if yname == "tec-current-warn-high-threshold" { return "TecCurrentWarnHighThreshold" }
    if yname == "tec-current-warn-low-threshold" { return "TecCurrentWarnLowThreshold" }
    if yname == "laser-diff-frequency-alarm-high-threshold" { return "LaserDiffFrequencyAlarmHighThreshold" }
    if yname == "laser-diff-frequency-alarm-low-threshold" { return "LaserDiffFrequencyAlarmLowThreshold" }
    if yname == "laser-diff-frequency-warn-high-threshold" { return "LaserDiffFrequencyWarnHighThreshold" }
    if yname == "laser-diff-frequency-warn-low-threshold" { return "LaserDiffFrequencyWarnLowThreshold" }
    if yname == "laser-diff-temperature-alarm-high-threshold" { return "LaserDiffTemperatureAlarmHighThreshold" }
    if yname == "laser-diff-temperature-alarm-low-threshold" { return "LaserDiffTemperatureAlarmLowThreshold" }
    if yname == "laser-diff-temperature-warn-high-threshold" { return "LaserDiffTemperatureWarnHighThreshold" }
    if yname == "laser-diff-temperature-warn-low-threshold" { return "LaserDiffTemperatureWarnLowThreshold" }
    if yname == "pre-fec-ber-latched-min-alarm-high-threshold" { return "PreFecBerLatchedMinAlarmHighThreshold" }
    if yname == "pre-fec-ber-latched-min-alarm-low-threshold" { return "PreFecBerLatchedMinAlarmLowThreshold" }
    if yname == "pre-fec-ber-latched-min-warn-high-threshold" { return "PreFecBerLatchedMinWarnHighThreshold" }
    if yname == "pre-fec-ber-latched-min-warn-low-threshold" { return "PreFecBerLatchedMinWarnLowThreshold" }
    if yname == "pre-fec-ber-latched-max-alarm-high-threshold" { return "PreFecBerLatchedMaxAlarmHighThreshold" }
    if yname == "pre-fec-ber-latched-max-alarm-low-threshold" { return "PreFecBerLatchedMaxAlarmLowThreshold" }
    if yname == "pre-fec-ber-latched-max-warn-high-threshold" { return "PreFecBerLatchedMaxWarnHighThreshold" }
    if yname == "pre-fec-ber-latched-max-warn-low-threshold" { return "PreFecBerLatchedMaxWarnLowThreshold" }
    if yname == "pre-fec-ber-accumulated-alarm-high-threshold" { return "PreFecBerAccumulatedAlarmHighThreshold" }
    if yname == "pre-fec-ber-accumulated-alarm-low-threshold" { return "PreFecBerAccumulatedAlarmLowThreshold" }
    if yname == "pre-fec-ber-accumulated-warn-high-threshold" { return "PreFecBerAccumulatedWarnHighThreshold" }
    if yname == "pre-fec-ber-accumulated-warn-low-threshold" { return "PreFecBerAccumulatedWarnLowThreshold" }
    if yname == "pre-fec-ber-instantaneous-alarm-high-threshold" { return "PreFecBerInstantaneousAlarmHighThreshold" }
    if yname == "pre-fec-ber-instantaneous-alarm-low-threshold" { return "PreFecBerInstantaneousAlarmLowThreshold" }
    if yname == "pre-fec-ber-instantaneous-warn-high-threshold" { return "PreFecBerInstantaneousWarnHighThreshold" }
    if yname == "pre-fec-ber-instantaneous-warn-low-threshold" { return "PreFecBerInstantaneousWarnLowThreshold" }
    if yname == "uncorrected-ber-latched-min-alarm-high-threshold" { return "UncorrectedBerLatchedMinAlarmHighThreshold" }
    if yname == "uncorrected-ber-latched-min-alarm-low-threshold" { return "UncorrectedBerLatchedMinAlarmLowThreshold" }
    if yname == "uncorrected-ber-latched-min-warn-high-threshold" { return "UncorrectedBerLatchedMinWarnHighThreshold" }
    if yname == "uncorrected-ber-latched-min-warn-low-threshold" { return "UncorrectedBerLatchedMinWarnLowThreshold" }
    if yname == "uncorrected-ber-latched-max-alarm-high-threshold" { return "UncorrectedBerLatchedMaxAlarmHighThreshold" }
    if yname == "uncorrected-ber-latched-max-alarm-low-threshold" { return "UncorrectedBerLatchedMaxAlarmLowThreshold" }
    if yname == "uncorrected-ber-latched-max-warn-high-threshold" { return "UncorrectedBerLatchedMaxWarnHighThreshold" }
    if yname == "uncorrected-ber-latched-max-warn-low-threshold" { return "UncorrectedBerLatchedMaxWarnLowThreshold" }
    if yname == "uncorrected-ber-accumulated-alarm-high-threshold" { return "UncorrectedBerAccumulatedAlarmHighThreshold" }
    if yname == "uncorrected-ber-accumulated-alarm-low-threshold" { return "UncorrectedBerAccumulatedAlarmLowThreshold" }
    if yname == "uncorrected-ber-accumulated-warn-high-threshold" { return "UncorrectedBerAccumulatedWarnHighThreshold" }
    if yname == "uncorrected-ber-accumulated-warn-low-threshold" { return "UncorrectedBerAccumulatedWarnLowThreshold" }
    if yname == "uncorrected-ber-instantaneous-alarm-high-threshold" { return "UncorrectedBerInstantaneousAlarmHighThreshold" }
    if yname == "uncorrected-ber-instantaneous-alarm-low-threshold" { return "UncorrectedBerInstantaneousAlarmLowThreshold" }
    if yname == "uncorrected-ber-instantaneous-warn-high-threshold" { return "UncorrectedBerInstantaneousWarnHighThreshold" }
    if yname == "uncorrected-ber-instantaneous-warn-low-threshold" { return "UncorrectedBerInstantaneousWarnLowThreshold" }
    return ""
}

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetSegmentPath() string {
    return "ext-param-threshold-val"
}

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snr-alarm-high-threshold"] = extParamThresholdVal.SnrAlarmHighThreshold
    leafs["snr-alarm-low-threshold"] = extParamThresholdVal.SnrAlarmLowThreshold
    leafs["snr-warn-high-threshold"] = extParamThresholdVal.SnrWarnHighThreshold
    leafs["snr-warn-low-threshold"] = extParamThresholdVal.SnrWarnLowThreshold
    leafs["isi-correction-alarm-high-threshold"] = extParamThresholdVal.IsiCorrectionAlarmHighThreshold
    leafs["isi-correction-alarm-low-threshold"] = extParamThresholdVal.IsiCorrectionAlarmLowThreshold
    leafs["isi-correction-warn-high-threshold"] = extParamThresholdVal.IsiCorrectionWarnHighThreshold
    leafs["isi-correction-warn-low-threshold"] = extParamThresholdVal.IsiCorrectionWarnLowThreshold
    leafs["pam-rate-alarm-high-threshold"] = extParamThresholdVal.PamRateAlarmHighThreshold
    leafs["pam-rate-alarm-low-threshold"] = extParamThresholdVal.PamRateAlarmLowThreshold
    leafs["pam-rate-warn-high-threshold"] = extParamThresholdVal.PamRateWarnHighThreshold
    leafs["pam-rate-warn-low-threshold"] = extParamThresholdVal.PamRateWarnLowThreshold
    leafs["pre-fec-ber-alarm-high-threshold"] = extParamThresholdVal.PreFecBerAlarmHighThreshold
    leafs["pre-fec-ber-alarm-low-threshold"] = extParamThresholdVal.PreFecBerAlarmLowThreshold
    leafs["pre-fec-ber-warn-high-threshold"] = extParamThresholdVal.PreFecBerWarnHighThreshold
    leafs["pre-fec-ber-warn-low-threshold"] = extParamThresholdVal.PreFecBerWarnLowThreshold
    leafs["uncorrected-ber-alarm-high-threshold"] = extParamThresholdVal.UncorrectedBerAlarmHighThreshold
    leafs["uncorrected-ber-alarm-low-threshold"] = extParamThresholdVal.UncorrectedBerAlarmLowThreshold
    leafs["uncorrected-ber-warn-high-threshold"] = extParamThresholdVal.UncorrectedBerWarnHighThreshold
    leafs["uncorrected-ber-warn-low-threshold"] = extParamThresholdVal.UncorrectedBerWarnLowThreshold
    leafs["tec-current-alarm-high-threshold"] = extParamThresholdVal.TecCurrentAlarmHighThreshold
    leafs["tec-current-alarm-low-threshold"] = extParamThresholdVal.TecCurrentAlarmLowThreshold
    leafs["tec-current-warn-high-threshold"] = extParamThresholdVal.TecCurrentWarnHighThreshold
    leafs["tec-current-warn-low-threshold"] = extParamThresholdVal.TecCurrentWarnLowThreshold
    leafs["laser-diff-frequency-alarm-high-threshold"] = extParamThresholdVal.LaserDiffFrequencyAlarmHighThreshold
    leafs["laser-diff-frequency-alarm-low-threshold"] = extParamThresholdVal.LaserDiffFrequencyAlarmLowThreshold
    leafs["laser-diff-frequency-warn-high-threshold"] = extParamThresholdVal.LaserDiffFrequencyWarnHighThreshold
    leafs["laser-diff-frequency-warn-low-threshold"] = extParamThresholdVal.LaserDiffFrequencyWarnLowThreshold
    leafs["laser-diff-temperature-alarm-high-threshold"] = extParamThresholdVal.LaserDiffTemperatureAlarmHighThreshold
    leafs["laser-diff-temperature-alarm-low-threshold"] = extParamThresholdVal.LaserDiffTemperatureAlarmLowThreshold
    leafs["laser-diff-temperature-warn-high-threshold"] = extParamThresholdVal.LaserDiffTemperatureWarnHighThreshold
    leafs["laser-diff-temperature-warn-low-threshold"] = extParamThresholdVal.LaserDiffTemperatureWarnLowThreshold
    leafs["pre-fec-ber-latched-min-alarm-high-threshold"] = extParamThresholdVal.PreFecBerLatchedMinAlarmHighThreshold
    leafs["pre-fec-ber-latched-min-alarm-low-threshold"] = extParamThresholdVal.PreFecBerLatchedMinAlarmLowThreshold
    leafs["pre-fec-ber-latched-min-warn-high-threshold"] = extParamThresholdVal.PreFecBerLatchedMinWarnHighThreshold
    leafs["pre-fec-ber-latched-min-warn-low-threshold"] = extParamThresholdVal.PreFecBerLatchedMinWarnLowThreshold
    leafs["pre-fec-ber-latched-max-alarm-high-threshold"] = extParamThresholdVal.PreFecBerLatchedMaxAlarmHighThreshold
    leafs["pre-fec-ber-latched-max-alarm-low-threshold"] = extParamThresholdVal.PreFecBerLatchedMaxAlarmLowThreshold
    leafs["pre-fec-ber-latched-max-warn-high-threshold"] = extParamThresholdVal.PreFecBerLatchedMaxWarnHighThreshold
    leafs["pre-fec-ber-latched-max-warn-low-threshold"] = extParamThresholdVal.PreFecBerLatchedMaxWarnLowThreshold
    leafs["pre-fec-ber-accumulated-alarm-high-threshold"] = extParamThresholdVal.PreFecBerAccumulatedAlarmHighThreshold
    leafs["pre-fec-ber-accumulated-alarm-low-threshold"] = extParamThresholdVal.PreFecBerAccumulatedAlarmLowThreshold
    leafs["pre-fec-ber-accumulated-warn-high-threshold"] = extParamThresholdVal.PreFecBerAccumulatedWarnHighThreshold
    leafs["pre-fec-ber-accumulated-warn-low-threshold"] = extParamThresholdVal.PreFecBerAccumulatedWarnLowThreshold
    leafs["pre-fec-ber-instantaneous-alarm-high-threshold"] = extParamThresholdVal.PreFecBerInstantaneousAlarmHighThreshold
    leafs["pre-fec-ber-instantaneous-alarm-low-threshold"] = extParamThresholdVal.PreFecBerInstantaneousAlarmLowThreshold
    leafs["pre-fec-ber-instantaneous-warn-high-threshold"] = extParamThresholdVal.PreFecBerInstantaneousWarnHighThreshold
    leafs["pre-fec-ber-instantaneous-warn-low-threshold"] = extParamThresholdVal.PreFecBerInstantaneousWarnLowThreshold
    leafs["uncorrected-ber-latched-min-alarm-high-threshold"] = extParamThresholdVal.UncorrectedBerLatchedMinAlarmHighThreshold
    leafs["uncorrected-ber-latched-min-alarm-low-threshold"] = extParamThresholdVal.UncorrectedBerLatchedMinAlarmLowThreshold
    leafs["uncorrected-ber-latched-min-warn-high-threshold"] = extParamThresholdVal.UncorrectedBerLatchedMinWarnHighThreshold
    leafs["uncorrected-ber-latched-min-warn-low-threshold"] = extParamThresholdVal.UncorrectedBerLatchedMinWarnLowThreshold
    leafs["uncorrected-ber-latched-max-alarm-high-threshold"] = extParamThresholdVal.UncorrectedBerLatchedMaxAlarmHighThreshold
    leafs["uncorrected-ber-latched-max-alarm-low-threshold"] = extParamThresholdVal.UncorrectedBerLatchedMaxAlarmLowThreshold
    leafs["uncorrected-ber-latched-max-warn-high-threshold"] = extParamThresholdVal.UncorrectedBerLatchedMaxWarnHighThreshold
    leafs["uncorrected-ber-latched-max-warn-low-threshold"] = extParamThresholdVal.UncorrectedBerLatchedMaxWarnLowThreshold
    leafs["uncorrected-ber-accumulated-alarm-high-threshold"] = extParamThresholdVal.UncorrectedBerAccumulatedAlarmHighThreshold
    leafs["uncorrected-ber-accumulated-alarm-low-threshold"] = extParamThresholdVal.UncorrectedBerAccumulatedAlarmLowThreshold
    leafs["uncorrected-ber-accumulated-warn-high-threshold"] = extParamThresholdVal.UncorrectedBerAccumulatedWarnHighThreshold
    leafs["uncorrected-ber-accumulated-warn-low-threshold"] = extParamThresholdVal.UncorrectedBerAccumulatedWarnLowThreshold
    leafs["uncorrected-ber-instantaneous-alarm-high-threshold"] = extParamThresholdVal.UncorrectedBerInstantaneousAlarmHighThreshold
    leafs["uncorrected-ber-instantaneous-alarm-low-threshold"] = extParamThresholdVal.UncorrectedBerInstantaneousAlarmLowThreshold
    leafs["uncorrected-ber-instantaneous-warn-high-threshold"] = extParamThresholdVal.UncorrectedBerInstantaneousWarnHighThreshold
    leafs["uncorrected-ber-instantaneous-warn-low-threshold"] = extParamThresholdVal.UncorrectedBerInstantaneousWarnLowThreshold
    return leafs
}

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetBundleName() string { return "cisco_ios_xr" }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetYangName() string { return "ext-param-threshold-val" }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) SetParent(parent types.Entity) { extParamThresholdVal.parent = parent }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetParent() types.Entity { return extParamThresholdVal.parent }

func (extParamThresholdVal *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_ExtParamThresholdVal) GetParentYangName() string { return "optics-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo
// OTS Spectrum information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo struct {
    parent types.Entity
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
    // OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo.
    SpectrumSlicePowerInfo []OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo
}

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetFilter() yfilter.YFilter { return spectrumInfo.YFilter }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) SetFilter(yf yfilter.YFilter) { spectrumInfo.YFilter = yf }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetGoName(yname string) string {
    if yname == "total-spectrum-slice-count" { return "TotalSpectrumSliceCount" }
    if yname == "spectrum-slice-spacing" { return "SpectrumSliceSpacing" }
    if yname == "first-slice-wavelength" { return "FirstSliceWavelength" }
    if yname == "spectrum-slice-power-info" { return "SpectrumSlicePowerInfo" }
    return ""
}

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetSegmentPath() string {
    return "spectrum-info"
}

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spectrum-slice-power-info" {
        for _, c := range spectrumInfo.SpectrumSlicePowerInfo {
            if spectrumInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo{}
        spectrumInfo.SpectrumSlicePowerInfo = append(spectrumInfo.SpectrumSlicePowerInfo, child)
        return &spectrumInfo.SpectrumSlicePowerInfo[len(spectrumInfo.SpectrumSlicePowerInfo)-1]
    }
    return nil
}

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spectrumInfo.SpectrumSlicePowerInfo {
        children[spectrumInfo.SpectrumSlicePowerInfo[i].GetSegmentPath()] = &spectrumInfo.SpectrumSlicePowerInfo[i]
    }
    return children
}

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-spectrum-slice-count"] = spectrumInfo.TotalSpectrumSliceCount
    leafs["spectrum-slice-spacing"] = spectrumInfo.SpectrumSliceSpacing
    leafs["first-slice-wavelength"] = spectrumInfo.FirstSliceWavelength
    return leafs
}

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetBundleName() string { return "cisco_ios_xr" }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetYangName() string { return "spectrum-info" }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) SetParent(parent types.Entity) { spectrumInfo.parent = parent }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetParent() types.Entity { return spectrumInfo.parent }

func (spectrumInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo) GetParentYangName() string { return "optics-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo
// Power information of spectrum slice info
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // spectrum slice number. The type is interface{} with range: 0..4294967295.
    SliceNum interface{}

    // Lower frequency of the specified PSD. The type is interface{} with range:
    // 0..18446744073709551615.
    LowerFrequency interface{}

    // Upper frequency of the specified PSD. The type is interface{} with range:
    // 0..18446744073709551615.
    UpperFrequency interface{}

    // Received Power in dBm. The type is string with length: 0..32.
    RxPower interface{}

    // Transmit Power in dBm. The type is string with length: 0..32.
    TxPower interface{}

    // Received Power spectral density in microwatts per megahertz, uW/MHz. The
    // type is string with length: 0..32.
    RxPsd interface{}

    // Transmit Power spectral density in microwatts per megahertz, uW/MHz. The
    // type is string with length: 0..32.
    TxPsd interface{}
}

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetFilter() yfilter.YFilter { return spectrumSlicePowerInfo.YFilter }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) SetFilter(yf yfilter.YFilter) { spectrumSlicePowerInfo.YFilter = yf }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetGoName(yname string) string {
    if yname == "slice-num" { return "SliceNum" }
    if yname == "lower-frequency" { return "LowerFrequency" }
    if yname == "upper-frequency" { return "UpperFrequency" }
    if yname == "rx-power" { return "RxPower" }
    if yname == "tx-power" { return "TxPower" }
    if yname == "rx-psd" { return "RxPsd" }
    if yname == "tx-psd" { return "TxPsd" }
    return ""
}

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetSegmentPath() string {
    return "spectrum-slice-power-info"
}

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slice-num"] = spectrumSlicePowerInfo.SliceNum
    leafs["lower-frequency"] = spectrumSlicePowerInfo.LowerFrequency
    leafs["upper-frequency"] = spectrumSlicePowerInfo.UpperFrequency
    leafs["rx-power"] = spectrumSlicePowerInfo.RxPower
    leafs["tx-power"] = spectrumSlicePowerInfo.TxPower
    leafs["rx-psd"] = spectrumSlicePowerInfo.RxPsd
    leafs["tx-psd"] = spectrumSlicePowerInfo.TxPsd
    return leafs
}

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetBundleName() string { return "cisco_ios_xr" }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetYangName() string { return "spectrum-slice-power-info" }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) SetParent(parent types.Entity) { spectrumSlicePowerInfo.parent = parent }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetParent() types.Entity { return spectrumSlicePowerInfo.parent }

func (spectrumSlicePowerInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_SpectrumInfo_SpectrumSlicePowerInfo) GetParentYangName() string { return "spectrum-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData
// Lane information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The index number of the lane. The type is interface{} with range:
    // 0..4294967295.
    LaneIndex interface{}

    // Laser Bias Current in units of 0.01 percentage. The type is interface{}
    // with range: 0..4294967295. Units are percentage.
    LaserBiasCurrentPercent interface{}

    // Laser Bias Current in units of 0.01mA. The type is interface{} with range:
    // 0..4294967295.
    LaserBiasCurrentMilliAmps interface{}

    // Transmit power in the unit of 0.01dBm. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPower interface{}

    // Transponder receive power in the unit of 0.01dBm. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePower interface{}

    // Transponder receive signal power in the unit of 0.01dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    ReceiveSignalPower interface{}

    // Transmit Signal power in the unit of 0.01dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    TransmitSignalPower interface{}

    // Output frequency read from hw in the unit 100MHz. The type is interface{}
    // with range: -2147483648..2147483647.
    OutputFrequency interface{}

    // Frequency Offset read from hw in unit of MHz. The type is interface{} with
    // range: -2147483648..2147483647.
    FrequencyOffset interface{}

    // Lane Alarm Information.
    LaneAlarmInfo OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo
}

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetFilter() yfilter.YFilter { return laneData.YFilter }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) SetFilter(yf yfilter.YFilter) { laneData.YFilter = yf }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetGoName(yname string) string {
    if yname == "lane-index" { return "LaneIndex" }
    if yname == "laser-bias-current-percent" { return "LaserBiasCurrentPercent" }
    if yname == "laser-bias-current-milli-amps" { return "LaserBiasCurrentMilliAmps" }
    if yname == "transmit-power" { return "TransmitPower" }
    if yname == "receive-power" { return "ReceivePower" }
    if yname == "receive-signal-power" { return "ReceiveSignalPower" }
    if yname == "transmit-signal-power" { return "TransmitSignalPower" }
    if yname == "output-frequency" { return "OutputFrequency" }
    if yname == "frequency-offset" { return "FrequencyOffset" }
    if yname == "lane-alarm-info" { return "LaneAlarmInfo" }
    return ""
}

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetSegmentPath() string {
    return "lane-data"
}

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lane-alarm-info" {
        return &laneData.LaneAlarmInfo
    }
    return nil
}

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lane-alarm-info"] = &laneData.LaneAlarmInfo
    return children
}

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lane-index"] = laneData.LaneIndex
    leafs["laser-bias-current-percent"] = laneData.LaserBiasCurrentPercent
    leafs["laser-bias-current-milli-amps"] = laneData.LaserBiasCurrentMilliAmps
    leafs["transmit-power"] = laneData.TransmitPower
    leafs["receive-power"] = laneData.ReceivePower
    leafs["receive-signal-power"] = laneData.ReceiveSignalPower
    leafs["transmit-signal-power"] = laneData.TransmitSignalPower
    leafs["output-frequency"] = laneData.OutputFrequency
    leafs["frequency-offset"] = laneData.FrequencyOffset
    return leafs
}

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetBundleName() string { return "cisco_ios_xr" }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetYangName() string { return "lane-data" }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) SetParent(parent types.Entity) { laneData.parent = parent }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetParent() types.Entity { return laneData.parent }

func (laneData *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData) GetParentYangName() string { return "optics-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo
// Lane Alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo struct {
    parent types.Entity
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

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetFilter() yfilter.YFilter { return laneAlarmInfo.YFilter }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) SetFilter(yf yfilter.YFilter) { laneAlarmInfo.YFilter = yf }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetGoName(yname string) string {
    if yname == "high-rx-power" { return "HighRxPower" }
    if yname == "low-rx-power" { return "LowRxPower" }
    if yname == "high-tx-power" { return "HighTxPower" }
    if yname == "low-tx-power" { return "LowTxPower" }
    if yname == "high-lbc" { return "HighLbc" }
    return ""
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetSegmentPath() string {
    return "lane-alarm-info"
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "high-rx-power" {
        return &laneAlarmInfo.HighRxPower
    }
    if childYangName == "low-rx-power" {
        return &laneAlarmInfo.LowRxPower
    }
    if childYangName == "high-tx-power" {
        return &laneAlarmInfo.HighTxPower
    }
    if childYangName == "low-tx-power" {
        return &laneAlarmInfo.LowTxPower
    }
    if childYangName == "high-lbc" {
        return &laneAlarmInfo.HighLbc
    }
    return nil
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["high-rx-power"] = &laneAlarmInfo.HighRxPower
    children["low-rx-power"] = &laneAlarmInfo.LowRxPower
    children["high-tx-power"] = &laneAlarmInfo.HighTxPower
    children["low-tx-power"] = &laneAlarmInfo.LowTxPower
    children["high-lbc"] = &laneAlarmInfo.HighLbc
    return children
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetYangName() string { return "lane-alarm-info" }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) SetParent(parent types.Entity) { laneAlarmInfo.parent = parent }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetParent() types.Entity { return laneAlarmInfo.parent }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo) GetParentYangName() string { return "lane-data" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower
// High Rx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetFilter() yfilter.YFilter { return highRxPower.YFilter }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) SetFilter(yf yfilter.YFilter) { highRxPower.YFilter = yf }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetSegmentPath() string {
    return "high-rx-power"
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highRxPower.IsDetected
    leafs["counter"] = highRxPower.Counter
    return leafs
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetBundleName() string { return "cisco_ios_xr" }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetYangName() string { return "high-rx-power" }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) SetParent(parent types.Entity) { highRxPower.parent = parent }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetParent() types.Entity { return highRxPower.parent }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighRxPower) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower
// Low Rx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetFilter() yfilter.YFilter { return lowRxPower.YFilter }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) SetFilter(yf yfilter.YFilter) { lowRxPower.YFilter = yf }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetSegmentPath() string {
    return "low-rx-power"
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowRxPower.IsDetected
    leafs["counter"] = lowRxPower.Counter
    return leafs
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetBundleName() string { return "cisco_ios_xr" }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetYangName() string { return "low-rx-power" }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) SetParent(parent types.Entity) { lowRxPower.parent = parent }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetParent() types.Entity { return lowRxPower.parent }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowRxPower) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower
// High Tx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetFilter() yfilter.YFilter { return highTxPower.YFilter }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) SetFilter(yf yfilter.YFilter) { highTxPower.YFilter = yf }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetSegmentPath() string {
    return "high-tx-power"
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTxPower.IsDetected
    leafs["counter"] = highTxPower.Counter
    return leafs
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetBundleName() string { return "cisco_ios_xr" }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetYangName() string { return "high-tx-power" }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) SetParent(parent types.Entity) { highTxPower.parent = parent }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetParent() types.Entity { return highTxPower.parent }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighTxPower) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower
// Low Tx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetFilter() yfilter.YFilter { return lowTxPower.YFilter }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) SetFilter(yf yfilter.YFilter) { lowTxPower.YFilter = yf }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetSegmentPath() string {
    return "low-tx-power"
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTxPower.IsDetected
    leafs["counter"] = lowTxPower.Counter
    return leafs
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetBundleName() string { return "cisco_ios_xr" }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetYangName() string { return "low-tx-power" }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) SetParent(parent types.Entity) { lowTxPower.parent = parent }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetParent() types.Entity { return lowTxPower.parent }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_LowTxPower) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc
// High laser bias current
type OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetFilter() yfilter.YFilter { return highLbc.YFilter }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) SetFilter(yf yfilter.YFilter) { highLbc.YFilter = yf }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetSegmentPath() string {
    return "high-lbc"
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highLbc.IsDetected
    leafs["counter"] = highLbc.Counter
    return leafs
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetBundleName() string { return "cisco_ios_xr" }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetYangName() string { return "high-lbc" }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) SetParent(parent types.Entity) { highLbc.parent = parent }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetParent() types.Entity { return highLbc.parent }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsInfo_LaneData_LaneAlarmInfo_HighLbc) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes
// All Optics Port operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Lane Information. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane.
    OpticsLane []OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane
}

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetFilter() yfilter.YFilter { return opticsLanes.YFilter }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) SetFilter(yf yfilter.YFilter) { opticsLanes.YFilter = yf }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetGoName(yname string) string {
    if yname == "optics-lane" { return "OpticsLane" }
    return ""
}

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetSegmentPath() string {
    return "optics-lanes"
}

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optics-lane" {
        for _, c := range opticsLanes.OpticsLane {
            if opticsLanes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane{}
        opticsLanes.OpticsLane = append(opticsLanes.OpticsLane, child)
        return &opticsLanes.OpticsLane[len(opticsLanes.OpticsLane)-1]
    }
    return nil
}

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range opticsLanes.OpticsLane {
        children[opticsLanes.OpticsLane[i].GetSegmentPath()] = &opticsLanes.OpticsLane[i]
    }
    return children
}

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetBundleName() string { return "cisco_ios_xr" }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetYangName() string { return "optics-lanes" }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) SetParent(parent types.Entity) { opticsLanes.parent = parent }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetParent() types.Entity { return opticsLanes.parent }

func (opticsLanes *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes) GetParentYangName() string { return "optics-port" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane
// Lane Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Lane Index. The type is interface{} with range:
    // -2147483648..2147483647.
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

    // Transmit power in the unit of 0.01dBm. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPower interface{}

    // Transponder receive power in the unit of 0.01dBm. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePower interface{}

    // Transponder receive signal power in the unit of 0.01dBm. The type is
    // interface{} with range: -2147483648..2147483647.
    ReceiveSignalPower interface{}

    // Transmit Signal power in the unit of 0.01dBm. The type is interface{} with
    // range: -2147483648..2147483647.
    TransmitSignalPower interface{}

    // Output frequency read from hw in the unit 100MHz. The type is interface{}
    // with range: -2147483648..2147483647.
    OutputFrequency interface{}

    // Frequency Offset read from hw in unit of MHz. The type is interface{} with
    // range: -2147483648..2147483647.
    FrequencyOffset interface{}

    // Lane Alarm Information.
    LaneAlarmInfo OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo
}

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetFilter() yfilter.YFilter { return opticsLane.YFilter }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) SetFilter(yf yfilter.YFilter) { opticsLane.YFilter = yf }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "lane-index" { return "LaneIndex" }
    if yname == "laser-bias-current-percent" { return "LaserBiasCurrentPercent" }
    if yname == "laser-bias-current-milli-amps" { return "LaserBiasCurrentMilliAmps" }
    if yname == "transmit-power" { return "TransmitPower" }
    if yname == "receive-power" { return "ReceivePower" }
    if yname == "receive-signal-power" { return "ReceiveSignalPower" }
    if yname == "transmit-signal-power" { return "TransmitSignalPower" }
    if yname == "output-frequency" { return "OutputFrequency" }
    if yname == "frequency-offset" { return "FrequencyOffset" }
    if yname == "lane-alarm-info" { return "LaneAlarmInfo" }
    return ""
}

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetSegmentPath() string {
    return "optics-lane" + "[number='" + fmt.Sprintf("%v", opticsLane.Number) + "']"
}

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lane-alarm-info" {
        return &opticsLane.LaneAlarmInfo
    }
    return nil
}

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lane-alarm-info"] = &opticsLane.LaneAlarmInfo
    return children
}

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = opticsLane.Number
    leafs["lane-index"] = opticsLane.LaneIndex
    leafs["laser-bias-current-percent"] = opticsLane.LaserBiasCurrentPercent
    leafs["laser-bias-current-milli-amps"] = opticsLane.LaserBiasCurrentMilliAmps
    leafs["transmit-power"] = opticsLane.TransmitPower
    leafs["receive-power"] = opticsLane.ReceivePower
    leafs["receive-signal-power"] = opticsLane.ReceiveSignalPower
    leafs["transmit-signal-power"] = opticsLane.TransmitSignalPower
    leafs["output-frequency"] = opticsLane.OutputFrequency
    leafs["frequency-offset"] = opticsLane.FrequencyOffset
    return leafs
}

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetBundleName() string { return "cisco_ios_xr" }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetYangName() string { return "optics-lane" }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) SetParent(parent types.Entity) { opticsLane.parent = parent }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetParent() types.Entity { return opticsLane.parent }

func (opticsLane *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane) GetParentYangName() string { return "optics-lanes" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo
// Lane Alarm Information
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo struct {
    parent types.Entity
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

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetFilter() yfilter.YFilter { return laneAlarmInfo.YFilter }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) SetFilter(yf yfilter.YFilter) { laneAlarmInfo.YFilter = yf }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetGoName(yname string) string {
    if yname == "high-rx-power" { return "HighRxPower" }
    if yname == "low-rx-power" { return "LowRxPower" }
    if yname == "high-tx-power" { return "HighTxPower" }
    if yname == "low-tx-power" { return "LowTxPower" }
    if yname == "high-lbc" { return "HighLbc" }
    return ""
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetSegmentPath() string {
    return "lane-alarm-info"
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "high-rx-power" {
        return &laneAlarmInfo.HighRxPower
    }
    if childYangName == "low-rx-power" {
        return &laneAlarmInfo.LowRxPower
    }
    if childYangName == "high-tx-power" {
        return &laneAlarmInfo.HighTxPower
    }
    if childYangName == "low-tx-power" {
        return &laneAlarmInfo.LowTxPower
    }
    if childYangName == "high-lbc" {
        return &laneAlarmInfo.HighLbc
    }
    return nil
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["high-rx-power"] = &laneAlarmInfo.HighRxPower
    children["low-rx-power"] = &laneAlarmInfo.LowRxPower
    children["high-tx-power"] = &laneAlarmInfo.HighTxPower
    children["low-tx-power"] = &laneAlarmInfo.LowTxPower
    children["high-lbc"] = &laneAlarmInfo.HighLbc
    return children
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetYangName() string { return "lane-alarm-info" }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) SetParent(parent types.Entity) { laneAlarmInfo.parent = parent }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetParent() types.Entity { return laneAlarmInfo.parent }

func (laneAlarmInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo) GetParentYangName() string { return "optics-lane" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower
// High Rx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetFilter() yfilter.YFilter { return highRxPower.YFilter }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) SetFilter(yf yfilter.YFilter) { highRxPower.YFilter = yf }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetSegmentPath() string {
    return "high-rx-power"
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highRxPower.IsDetected
    leafs["counter"] = highRxPower.Counter
    return leafs
}

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetBundleName() string { return "cisco_ios_xr" }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetYangName() string { return "high-rx-power" }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) SetParent(parent types.Entity) { highRxPower.parent = parent }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetParent() types.Entity { return highRxPower.parent }

func (highRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighRxPower) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower
// Low Rx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetFilter() yfilter.YFilter { return lowRxPower.YFilter }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) SetFilter(yf yfilter.YFilter) { lowRxPower.YFilter = yf }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetSegmentPath() string {
    return "low-rx-power"
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowRxPower.IsDetected
    leafs["counter"] = lowRxPower.Counter
    return leafs
}

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetBundleName() string { return "cisco_ios_xr" }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetYangName() string { return "low-rx-power" }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) SetParent(parent types.Entity) { lowRxPower.parent = parent }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetParent() types.Entity { return lowRxPower.parent }

func (lowRxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowRxPower) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower
// High Tx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetFilter() yfilter.YFilter { return highTxPower.YFilter }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) SetFilter(yf yfilter.YFilter) { highTxPower.YFilter = yf }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetSegmentPath() string {
    return "high-tx-power"
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highTxPower.IsDetected
    leafs["counter"] = highTxPower.Counter
    return leafs
}

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetBundleName() string { return "cisco_ios_xr" }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetYangName() string { return "high-tx-power" }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) SetParent(parent types.Entity) { highTxPower.parent = parent }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetParent() types.Entity { return highTxPower.parent }

func (highTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighTxPower) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower
// Low Tx Power
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetFilter() yfilter.YFilter { return lowTxPower.YFilter }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) SetFilter(yf yfilter.YFilter) { lowTxPower.YFilter = yf }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetSegmentPath() string {
    return "low-tx-power"
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = lowTxPower.IsDetected
    leafs["counter"] = lowTxPower.Counter
    return leafs
}

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetBundleName() string { return "cisco_ios_xr" }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetYangName() string { return "low-tx-power" }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) SetParent(parent types.Entity) { lowTxPower.parent = parent }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetParent() types.Entity { return lowTxPower.parent }

func (lowTxPower *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_LowTxPower) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc
// High laser bias current
type OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Alarm counter. The type is interface{} with range: 0..4294967295.
    Counter interface{}
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetFilter() yfilter.YFilter { return highLbc.YFilter }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) SetFilter(yf yfilter.YFilter) { highLbc.YFilter = yf }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetGoName(yname string) string {
    if yname == "is-detected" { return "IsDetected" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetSegmentPath() string {
    return "high-lbc"
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-detected"] = highLbc.IsDetected
    leafs["counter"] = highLbc.Counter
    return leafs
}

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetBundleName() string { return "cisco_ios_xr" }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetYangName() string { return "high-lbc" }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) SetParent(parent types.Entity) { highLbc.parent = parent }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetParent() types.Entity { return highLbc.parent }

func (highLbc *OpticsOper_OpticsPorts_OpticsPort_OpticsLanes_OpticsLane_LaneAlarmInfo_HighLbc) GetParentYangName() string { return "lane-alarm-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo
// Optics operational data
type OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transport Admin State. The type is OpticsTas.
    TransportAdminState interface{}

    // Optics controller state: Up, Down or Administratively Down. The type is
    // OpticsControllerState.
    ControllerState interface{}

    // Network SRLG information.
    NetworkSrlgInfo OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo
}

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetFilter() yfilter.YFilter { return opticsDbInfo.YFilter }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) SetFilter(yf yfilter.YFilter) { opticsDbInfo.YFilter = yf }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetGoName(yname string) string {
    if yname == "transport-admin-state" { return "TransportAdminState" }
    if yname == "controller-state" { return "ControllerState" }
    if yname == "network-srlg-info" { return "NetworkSrlgInfo" }
    return ""
}

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetSegmentPath() string {
    return "optics-db-info"
}

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-srlg-info" {
        return &opticsDbInfo.NetworkSrlgInfo
    }
    return nil
}

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-srlg-info"] = &opticsDbInfo.NetworkSrlgInfo
    return children
}

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transport-admin-state"] = opticsDbInfo.TransportAdminState
    leafs["controller-state"] = opticsDbInfo.ControllerState
    return leafs
}

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetBundleName() string { return "cisco_ios_xr" }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetYangName() string { return "optics-db-info" }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) SetParent(parent types.Entity) { opticsDbInfo.parent = parent }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetParent() types.Entity { return opticsDbInfo.parent }

func (opticsDbInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo) GetParentYangName() string { return "optics-port" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo
// Network SRLG information
type OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network Srlg Array. The type is slice of
    // OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray.
    NetworkSrlgArray []OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetFilter() yfilter.YFilter { return networkSrlgInfo.YFilter }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) SetFilter(yf yfilter.YFilter) { networkSrlgInfo.YFilter = yf }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetGoName(yname string) string {
    if yname == "network-srlg-array" { return "NetworkSrlgArray" }
    return ""
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetSegmentPath() string {
    return "network-srlg-info"
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-srlg-array" {
        for _, c := range networkSrlgInfo.NetworkSrlgArray {
            if networkSrlgInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray{}
        networkSrlgInfo.NetworkSrlgArray = append(networkSrlgInfo.NetworkSrlgArray, child)
        return &networkSrlgInfo.NetworkSrlgArray[len(networkSrlgInfo.NetworkSrlgArray)-1]
    }
    return nil
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkSrlgInfo.NetworkSrlgArray {
        children[networkSrlgInfo.NetworkSrlgArray[i].GetSegmentPath()] = &networkSrlgInfo.NetworkSrlgArray[i]
    }
    return children
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetBundleName() string { return "cisco_ios_xr" }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetYangName() string { return "network-srlg-info" }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) SetParent(parent types.Entity) { networkSrlgInfo.parent = parent }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetParent() types.Entity { return networkSrlgInfo.parent }

func (networkSrlgInfo *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo) GetParentYangName() string { return "optics-db-info" }

// OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray
// Network Srlg Array
type OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array to maintain set number. The type is interface{} with range:
    // 0..4294967295.
    SetNumber interface{}

    // Network Srlg. The type is slice of interface{} with range: 0..4294967295.
    NetworkSrlg []interface{}
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetFilter() yfilter.YFilter { return networkSrlgArray.YFilter }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) SetFilter(yf yfilter.YFilter) { networkSrlgArray.YFilter = yf }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetGoName(yname string) string {
    if yname == "set-number" { return "SetNumber" }
    if yname == "network-srlg" { return "NetworkSrlg" }
    return ""
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetSegmentPath() string {
    return "network-srlg-array"
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-number"] = networkSrlgArray.SetNumber
    leafs["network-srlg"] = networkSrlgArray.NetworkSrlg
    return leafs
}

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetBundleName() string { return "cisco_ios_xr" }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetYangName() string { return "network-srlg-array" }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) SetParent(parent types.Entity) { networkSrlgArray.parent = parent }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetParent() types.Entity { return networkSrlgArray.parent }

func (networkSrlgArray *OpticsOper_OpticsPorts_OpticsPort_OpticsDbInfo_NetworkSrlgInfo_NetworkSrlgArray) GetParentYangName() string { return "network-srlg-info" }

