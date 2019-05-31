// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package controller_odu_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package controller_odu_datatypes"))
}

// OtnChildControllerName represents Otn child controller name
type OtnChildControllerName string

const (
    // Create lower order odu1 controller
    OtnChildControllerName_odu1 OtnChildControllerName = "odu1"

    // Create lower order odu2 controller
    OtnChildControllerName_odu2 OtnChildControllerName = "odu2"

    // Create lower order odu3 controller
    OtnChildControllerName_odu3 OtnChildControllerName = "odu3"

    // Create lower order odu0 controller
    OtnChildControllerName_odu0 OtnChildControllerName = "odu0"

    // Create lower order odu2e controller
    OtnChildControllerName_odu2e OtnChildControllerName = "odu2e"

    // Create lower order odu1e controller
    OtnChildControllerName_odu1e OtnChildControllerName = "odu1e"

    // Create lower order odu2f controller
    OtnChildControllerName_odu2f OtnChildControllerName = "odu2f"

    // Create lower order odu3e1 controller
    OtnChildControllerName_odu3e1 OtnChildControllerName = "odu3e1"

    // Create lower order odu3e2 controller
    OtnChildControllerName_odu3e2 OtnChildControllerName = "odu3e2"
)

// OtnChildFlexControllerName represents Otn child flex controller name
type OtnChildFlexControllerName string

const (
    // Create lower order odu-flex controller
    OtnChildFlexControllerName_odu_flex OtnChildFlexControllerName = "odu-flex"
)

// OtnFlexMapping represents Otn flex mapping
type OtnFlexMapping string

const (
    // GFP-FIX Mapping
    OtnFlexMapping_gfp_f_fixed OtnFlexMapping = "gfp-f-fixed"

    // GFP-Resizable Mapping
    OtnFlexMapping_gfp_resizable OtnFlexMapping = "gfp-resizable"

    // CBR Mapping
    OtnFlexMapping_cbr OtnFlexMapping = "cbr"
)

// OtntcmMode represents Otntcm mode
type OtntcmMode string

const (
    // Set TCM Mode as transparent
    OtntcmMode_transparent OtntcmMode = "transparent"

    // Set TCM Mode as NIM
    OtntcmMode_nim OtntcmMode = "nim"

    // Set TCM Mode as operational
    OtntcmMode_operational OtntcmMode = "operational"
)

// OtnPerMon represents Otn per mon
type OtnPerMon string

const (
    // Performance Monitoring Disabled
    OtnPerMon_disable OtnPerMon = "disable"

    // Performance Monitoring Enabled
    OtnPerMon_enable OtnPerMon = "enable"
)

// OduPrbsMode represents Odu prbs mode
type OduPrbsMode string

const (
    // prbs Mode Invalid
    OduPrbsMode_mode_invalid OduPrbsMode = "mode-invalid"

    // Prbs Mode Source
    OduPrbsMode_mode_source OduPrbsMode = "mode-source"

    // Prbs Mode Sink
    OduPrbsMode_mode_sink OduPrbsMode = "mode-sink"

    // Prbs Mode Source_Sink
    OduPrbsMode_mode_source_sink OduPrbsMode = "mode-source-sink"
)

// OduTimeSlotGranularity represents Odu time slot granularity
type OduTimeSlotGranularity string

const (
    // 1.25G time slot granularity
    OduTimeSlotGranularity_Y_1__DOT__25g OduTimeSlotGranularity = "1.25g"

    // 2.5G time slot granularity
    OduTimeSlotGranularity_Y_2__DOT__5g OduTimeSlotGranularity = "2.5g"
)

// OtnSendTtiTypeOs represents Otn send tti type os
type OtnSendTtiTypeOs string

const (
    // Send TTI OS ASCII string
    OtnSendTtiTypeOs_send_tti_os_ascii__FWD_SLASH__os_ascii OtnSendTtiTypeOs = "send-tti-os-ascii/os-ascii"

    // Send TTI OS HEX string
    OtnSendTtiTypeOs_send_tti_os_hex__FWD_SLASH__os_hex OtnSendTtiTypeOs = "send-tti-os-hex/os-hex"
)

// OtnExpTtiTypeSapi represents Otn exp tti type sapi
type OtnExpTtiTypeSapi string

const (
    // Expected TTI SAPI ASCII string
    OtnExpTtiTypeSapi_exp_tti_sapi_ascii__FWD_SLASH__sapi_ascii OtnExpTtiTypeSapi = "exp-tti-sapi-ascii/sapi-ascii"
)

// OtnTermination represents Otn termination
type OtnTermination string

const (
    // Termination to ether
    OtnTermination_ether OtnTermination = "ether"
)

// OtnSendTtiTypeSapi represents Otn send tti type sapi
type OtnSendTtiTypeSapi string

const (
    // Send TTI SAPI ASCII string
    OtnSendTtiTypeSapi_send_tti_sapi_ascii__FWD_SLASH__sapi_ascii OtnSendTtiTypeSapi = "send-tti-sapi-ascii/sapi-ascii"
)

// OtnSecAdminState represents Otn sec admin state
type OtnSecAdminState string

const (
    // In normal state
    OtnSecAdminState_normal OtnSecAdminState = "normal"

    // Under maintenance
    OtnSecAdminState_maintenance OtnSecAdminState = "maintenance"
)

// OtnMapping represents Otn mapping
type OtnMapping string

const (
    // gfp_f for mapping
    OtnMapping_gfp_f OtnMapping = "gfp-f"

    // bmp for mapping
    OtnMapping_bmp OtnMapping = "bmp"

    // gmp for mapping
    OtnMapping_gmp OtnMapping = "gmp"

    // gfp_f_ext for mapping
    OtnMapping_gfp_f_ext OtnMapping = "gfp-f-ext"
)

// OtnExpTtiTypeOs represents Otn exp tti type os
type OtnExpTtiTypeOs string

const (
    // Expected TTI OS ASCII string
    OtnExpTtiTypeOs_exp_tti_os_ascii__FWD_SLASH__os_ascii OtnExpTtiTypeOs = "exp-tti-os-ascii/os-ascii"

    // Expected TTI OS HEX string
    OtnExpTtiTypeOs_exp_tti_os_hex__FWD_SLASH__os_hex OtnExpTtiTypeOs = "exp-tti-os-hex/os-hex"
)

// OduDelay represents Odu delay
type OduDelay string

const (
    // Delay Disable
    OduDelay_disable OduDelay = "disable"

    // Delay Enable
    OduDelay_enable OduDelay = "enable"
)

// OtnLoopback represents Otn loopback
type OtnLoopback string

const (
    // Line loopback
    OtnLoopback_line OtnLoopback = "line"

    // Internal loopback
    OtnLoopback_internal OtnLoopback = "internal"
)

// OtnExpTtiTypeFull represents Otn exp tti type full
type OtnExpTtiTypeFull string

const (
    // Expected TTI Full ASCII string
    OtnExpTtiTypeFull_exp_tti_full_ascii__FWD_SLASH__full_ascii OtnExpTtiTypeFull = "exp-tti-full-ascii/full-ascii"

    // Expected TTI hex string
    OtnExpTtiTypeFull_exp_tti_hex__FWD_SLASH__hex OtnExpTtiTypeFull = "exp-tti-hex/hex"
)

// OtnSendTtiTypeFull represents Otn send tti type full
type OtnSendTtiTypeFull string

const (
    // Send TTI Full ASCII string
    OtnSendTtiTypeFull_send_tti_full_ascii__FWD_SLASH__full_ascii OtnSendTtiTypeFull = "send-tti-full-ascii/full-ascii"

    // Send TTI hex string
    OtnSendTtiTypeFull_send_tti_hex__FWD_SLASH__hex OtnSendTtiTypeFull = "send-tti-hex/hex"
)

// OtnExpTtiTypeDapi represents Otn exp tti type dapi
type OtnExpTtiTypeDapi string

const (
    // Expected TTI DAPI ASCII string
    OtnExpTtiTypeDapi_exp_tti_dapi_ascii__FWD_SLASH__dapi_ascii OtnExpTtiTypeDapi = "exp-tti-dapi-ascii/dapi-ascii"
)

// OtnSendTtiTypeDapi represents Otn send tti type dapi
type OtnSendTtiTypeDapi string

const (
    // Send TTI DAPI ASCII string
    OtnSendTtiTypeDapi_send_tti_dapi_ascii__FWD_SLASH__dapi_ascii OtnSendTtiTypeDapi = "send-tti-dapi-ascii/dapi-ascii"
)

// Otntcmca represents Otntcmca
type Otntcmca string

const (
    // Consequent(ial) action for Disabled
    Otntcmca_disable Otntcmca = "disable"

    // consequent(ial) action for Enabled
    Otntcmca_enable Otntcmca = "enable"
)

// Pattern represents Pattern
type Pattern string

const (
    // prbs pattern None
    Pattern_pattern_none Pattern = "pattern-none"

    // Prbs pattern pn31
    Pattern_pattern_pn31 Pattern = "pattern-pn31"

    // Prbs pattern pn23
    Pattern_pattern_pn23 Pattern = "pattern-pn23"

    // Prbs pattern pn11
    Pattern_pattern_pn11 Pattern = "pattern-pn11"

    // Prbs pattern inverted pn31
    Pattern_pattern_inverted_pn31 Pattern = "pattern-inverted-pn31"

    // Prbs pattern inverted pn11
    Pattern_pattern_inverted_pn11 Pattern = "pattern-inverted-pn11"
)

// Otnpmtimca represents Otnpmtimca
type Otnpmtimca string

const (
    // Path layer PM TIM Consequent action Disabled
    Otnpmtimca_disable Otnpmtimca = "disable"

    // Path layer PM TIM Consequent action  Enabled
    Otnpmtimca_enable Otnpmtimca = "enable"
)

