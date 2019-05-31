// This module contains a collection of YANG definitions
// for Cisco IOS-XR controller-optics package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package controller_optics_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package controller_optics_cfg"))
}

// OpticsFec represents Optics fec
type OpticsFec string

const (
    // No Fec
    OpticsFec_fec_none OpticsFec = "fec-none"

    // Enhanced H15
    OpticsFec_fec_h15 OpticsFec = "fec-h15"

    // Enhanced H25
    OpticsFec_fec_h25 OpticsFec = "fec-h25"

    // Enhanced H15 DE
    OpticsFec_fec_h15_de OpticsFec = "fec-h15-de"

    // Enhanced H25 DE
    OpticsFec_fec_h25_de OpticsFec = "fec-h25-de"
)

// OpticsOtsAmpliControlMode represents Optics ots ampli control mode
type OpticsOtsAmpliControlMode string

const (
    // Automatic Amplifier Mode
    OpticsOtsAmpliControlMode_automatic OpticsOtsAmpliControlMode = "automatic"

    // Manual Amplifier Mode
    OpticsOtsAmpliControlMode_manual OpticsOtsAmpliControlMode = "manual"
)

// Threshold represents Threshold
type Threshold string

const (
    // Low Threshold
    Threshold_low Threshold = "low"

    // High Threshold
    Threshold_high Threshold = "high"
)

// OpticsDwdmCarrierParam represents Optics dwdm carrier param
type OpticsDwdmCarrierParam string

const (
    // ITU Wave Channel Number
    OpticsDwdmCarrierParam_itu_ch OpticsDwdmCarrierParam = "itu-ch"

    // Wavelength in nm
    OpticsDwdmCarrierParam_wavelength OpticsDwdmCarrierParam = "wavelength"

    // Frequency in Hertz
    OpticsDwdmCarrierParam_frequency OpticsDwdmCarrierParam = "frequency"
)

// OpticsOtsSafetyControlMode represents Optics ots safety control mode
type OpticsOtsSafetyControlMode string

const (
    // Automatic Safety Control Mode
    OpticsOtsSafetyControlMode_auto OpticsOtsSafetyControlMode = "auto"

    // Disable Safety Control Mode
    OpticsOtsSafetyControlMode_disabled OpticsOtsSafetyControlMode = "disabled"
)

// OpticsOtsAmpliGainRange represents Optics ots ampli gain range
type OpticsOtsAmpliGainRange string

const (
    // Normal Amplifier Gain Range
    OpticsOtsAmpliGainRange_normal OpticsOtsAmpliGainRange = "normal"

    // Extended Amplifier Gain Range
    OpticsOtsAmpliGainRange_extended OpticsOtsAmpliGainRange = "extended"
)

// OpticsDwdmCarrierGrid represents Optics dwdm carrier grid
type OpticsDwdmCarrierGrid string

const (
    // 50GHz Grid
    OpticsDwdmCarrierGrid_Y_50g_hz_grid OpticsDwdmCarrierGrid = "50g-hz-grid"

    // 100MHz Grid
    OpticsDwdmCarrierGrid_Y_100mhz_grid OpticsDwdmCarrierGrid = "100mhz-grid"
)

// OpticsLoopback represents Optics loopback
type OpticsLoopback string

const (
    // No Loopback
    OpticsLoopback_none OpticsLoopback = "none"

    // Internal Loopback
    OpticsLoopback_internal OpticsLoopback = "internal"

    // Line Loopback
    OpticsLoopback_line OpticsLoopback = "line"
)

