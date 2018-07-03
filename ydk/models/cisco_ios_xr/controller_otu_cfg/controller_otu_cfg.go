// This module contains a collection of YANG definitions
// for Cisco IOS-XR controller-otu package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package controller_otu_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package controller_otu_cfg"))
}

// OtnLoopback represents Otn loopback
type OtnLoopback string

const (
    // Line loopback
    OtnLoopback_line OtnLoopback = "line"

    // Internal loopback
    OtnLoopback_internal OtnLoopback = "internal"
)

// OtnExpTtiTypeSapi represents Otn exp tti type sapi
type OtnExpTtiTypeSapi string

const (
    // Expected TTI SAPI ASCII string
    OtnExpTtiTypeSapi_exp_tti_sapi_ascii__FWD_SLASH__sapi_ascii OtnExpTtiTypeSapi = "exp-tti-sapi-ascii/sapi-ascii"
)

// OtnPerMon represents Otn per mon
type OtnPerMon string

const (
    // Performance Monitoring Disabled
    OtnPerMon_disable OtnPerMon = "disable"

    // Performance Monitoring Enabled
    OtnPerMon_enable OtnPerMon = "enable"
)

// OtnSendTtiTypeOs represents Otn send tti type os
type OtnSendTtiTypeOs string

const (
    // Send TTI OS ASCII string
    OtnSendTtiTypeOs_send_tti_os_ascii__FWD_SLASH__os_ascii OtnSendTtiTypeOs = "send-tti-os-ascii/os-ascii"

    // Send TTI OS HEX string
    OtnSendTtiTypeOs_send_tti_os_hex__FWD_SLASH__os_hex OtnSendTtiTypeOs = "send-tti-os-hex/os-hex"
)

// OtuPattern represents Otu pattern
type OtuPattern string

const (
    // prbs pattern None
    OtuPattern_pattern_none OtuPattern = "pattern-none"

    // Prbs pattern pn31
    OtuPattern_pattern_pn31 OtuPattern = "pattern-pn31"

    // Prbs pattern pn23
    OtuPattern_pattern_pn23 OtuPattern = "pattern-pn23"

    // Prbs pattern pn11
    OtuPattern_pattern_pn11 OtuPattern = "pattern-pn11"

    // Prbs pattern inverted pn31
    OtuPattern_pattern_inverted_pn31 OtuPattern = "pattern-inverted-pn31"

    // Prbs pattern inverted pn11
    OtuPattern_pattern_inverted_pn11 OtuPattern = "pattern-inverted-pn11"

    // Prbs pattern pn15
    OtuPattern_pattern_pn15 OtuPattern = "pattern-pn15"
)

// OtnSendTtiTypeSapi represents Otn send tti type sapi
type OtnSendTtiTypeSapi string

const (
    // Send TTI SAPI ASCII string
    OtnSendTtiTypeSapi_send_tti_sapi_ascii__FWD_SLASH__sapi_ascii OtnSendTtiTypeSapi = "send-tti-sapi-ascii/sapi-ascii"
)

// OtnSendTtiTypeDapi represents Otn send tti type dapi
type OtnSendTtiTypeDapi string

const (
    // Send TTI DAPI ASCII string
    OtnSendTtiTypeDapi_send_tti_dapi_ascii__FWD_SLASH__dapi_ascii OtnSendTtiTypeDapi = "send-tti-dapi-ascii/dapi-ascii"
)

// OtuMode represents Otu mode
type OtuMode string

const (
    // prbs Mode Invalid
    OtuMode_mode_invalid OtuMode = "mode-invalid"

    // Prbs Mode Source
    OtuMode_mode_source OtuMode = "mode-source"

    // Prbs Mode Sink
    OtuMode_mode_sink OtuMode = "mode-sink"

    // Prbs Mode Source_Sink
    OtuMode_mode_source_sink OtuMode = "mode-source-sink"
)

// OtnExpTtiTypeOs represents Otn exp tti type os
type OtnExpTtiTypeOs string

const (
    // Expected TTI OS ASCII string
    OtnExpTtiTypeOs_exp_tti_os_ascii__FWD_SLASH__os_ascii OtnExpTtiTypeOs = "exp-tti-os-ascii/os-ascii"

    // Expected TTI OS HEX string
    OtnExpTtiTypeOs_exp_tti_os_hex__FWD_SLASH__os_hex OtnExpTtiTypeOs = "exp-tti-os-hex/os-hex"
)

// OtnSecAdminState represents Otn sec admin state
type OtnSecAdminState string

const (
    // In normal state
    OtnSecAdminState_normal OtnSecAdminState = "normal"

    // Under maintenance
    OtnSecAdminState_maintenance OtnSecAdminState = "maintenance"
)

// OtnExpTtiTypeFull represents Otn exp tti type full
type OtnExpTtiTypeFull string

const (
    // Expected TTI Full ASCII string
    OtnExpTtiTypeFull_exp_tti_full_ascii__FWD_SLASH__full_ascii OtnExpTtiTypeFull = "exp-tti-full-ascii/full-ascii"

    // Expected TTI hex string
    OtnExpTtiTypeFull_exp_tti_hex__FWD_SLASH__hex OtnExpTtiTypeFull = "exp-tti-hex/hex"
)

// OtuForwardErrorCorrection represents Otu forward error correction
type OtuForwardErrorCorrection string

const (
    // No Fec
    OtuForwardErrorCorrection_none OtuForwardErrorCorrection = "none"

    // Standard Fec
    OtuForwardErrorCorrection_standard OtuForwardErrorCorrection = "standard"

    // EnhancedI7 Fec
    OtuForwardErrorCorrection_enhanced_i7 OtuForwardErrorCorrection = "enhanced-i7"

    // Enhanced I4 Fec
    OtuForwardErrorCorrection_enhanced_i4 OtuForwardErrorCorrection = "enhanced-i4"

    // EnhancedSwizzle Fec
    OtuForwardErrorCorrection_enhanced_swizzle OtuForwardErrorCorrection = "enhanced-swizzle"

    // EnhancedHG20 Fec
    OtuForwardErrorCorrection_enhanced_hg20 OtuForwardErrorCorrection = "enhanced-hg20"

    // EnhancedHG7 Fec
    OtuForwardErrorCorrection_enhanced_hg7 OtuForwardErrorCorrection = "enhanced-hg7"
)

// OtnExpTtiTypeDapi represents Otn exp tti type dapi
type OtnExpTtiTypeDapi string

const (
    // Expected TTI DAPI ASCII string
    OtnExpTtiTypeDapi_exp_tti_dapi_ascii__FWD_SLASH__dapi_ascii OtnExpTtiTypeDapi = "exp-tti-dapi-ascii/dapi-ascii"
)

// OtnSendTtiTypeFull represents Otn send tti type full
type OtnSendTtiTypeFull string

const (
    // Send TTI Full ASCII string
    OtnSendTtiTypeFull_send_tti_full_ascii__FWD_SLASH__full_ascii OtnSendTtiTypeFull = "send-tti-full-ascii/full-ascii"

    // Send TTI hex string
    OtnSendTtiTypeFull_send_tti_hex__FWD_SLASH__hex OtnSendTtiTypeFull = "send-tti-hex/hex"
)

