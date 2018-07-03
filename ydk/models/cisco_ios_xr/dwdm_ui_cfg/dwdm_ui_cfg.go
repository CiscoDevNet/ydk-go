// This module contains a collection of YANG definitions
// for Cisco IOS-XR dwdm-ui package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dwdm_ui_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dwdm_ui_cfg"))
}

// WaveChannelNum represents Wave channel num
type WaveChannelNum string

const (
    // Default Wave Channel Number
    WaveChannelNum_default_ WaveChannelNum = "default"

    // Wavelength Wave Channel Number
    WaveChannelNum_channel_wavelength WaveChannelNum = "channel-wavelength"

    // Frequency Wave Channel Number
    WaveChannelNum_channel_frequency WaveChannelNum = "channel-frequency"

    // Frequency in Steps of 100MHz
    WaveChannelNum_Y_100mhz_frequency WaveChannelNum = "100mhz-frequency"
)

// DwdmLoopback represents Dwdm loopback
type DwdmLoopback string

const (
    // No Loopback
    DwdmLoopback_none DwdmLoopback = "none"

    // Line Loopback
    DwdmLoopback_line DwdmLoopback = "line"

    // Internal Loopback
    DwdmLoopback_internal DwdmLoopback = "internal"
)

// Fec represents Fec
type Fec string

const (
    // No FEC
    Fec_none Fec = "none"

    // Standard FEC
    Fec_standard Fec = "standard"

    // Enhanced FEC
    Fec_enhanced Fec = "enhanced"

    // High-Gain Hard Decision
    Fec_high_gain_hd Fec = "high-gain-hd"

    // Long-Haul Hard Decision
    Fec_long_haul_hd Fec = "long-haul-hd"

    // High-Gain Soft Decision
    Fec_high_gain_sd Fec = "high-gain-sd"

    // Long-Haul Soft Decision
    Fec_long_haul_sd Fec = "long-haul-sd"

    // Ci BCH
    Fec_ci_bch Fec = "ci-bch"

    // High-Gain Multivendor Interoperable Hard
    // Decision
    Fec_high_gain_multivendor_hd Fec = "high-gain-multivendor-hd"
)

// OduAlarm represents Odu alarm
type OduAlarm string

const (
    // ODU OCI
    OduAlarm_oci OduAlarm = "oci"

    // ODU AIS
    OduAlarm_odu_ais OduAlarm = "odu-ais"

    // ODU LCK
    OduAlarm_lck OduAlarm = "lck"

    // ODU BDI
    OduAlarm_odu_bdi OduAlarm = "odu-bdi"

    // ODU SF BER
    OduAlarm_odu_sf OduAlarm = "odu-sf"

    // ODU SD BER
    OduAlarm_odu_sd OduAlarm = "odu-sd"

    // ODU PTIM
    OduAlarm_plm OduAlarm = "plm"

    // ODU TIM
    OduAlarm_odu_tim OduAlarm = "odu-tim"
)

// PrbsMode represents Prbs mode
type PrbsMode string

const (
    // Source Mode
    PrbsMode_source PrbsMode = "source"

    // Sink Mode
    PrbsMode_sink PrbsMode = "sink"

    // Source-Sink Mode
    PrbsMode_source_sink PrbsMode = "source-sink"

    // Invalid Mode
    PrbsMode_invalid PrbsMode = "invalid"
)

// ExpectedTti represents Expected tti
type ExpectedTti string

const (
    // Expected TTI ascii string
    ExpectedTti_expected_tti_ascii ExpectedTti = "expected-tti-ascii"

    // Expected TTI hex string
    ExpectedTti_expected_tti_hex ExpectedTti = "expected-tti-hex"
)

// DwdmAdminState represents Dwdm admin state
type DwdmAdminState string

const (
    // Out of service
    DwdmAdminState_out_of_service DwdmAdminState = "out-of-service"

    // In service
    DwdmAdminState_in_service DwdmAdminState = "in-service"

    // Out of service maintenance
    DwdmAdminState_maintenance DwdmAdminState = "maintenance"

    // In service Config allowed
    DwdmAdminState_in_service_config_allowed DwdmAdminState = "in-service-config-allowed"
)

// PrbsPattern represents Prbs pattern
type PrbsPattern string

const (
    // None Pattern
    PrbsPattern_none PrbsPattern = "none"

    // Null Pattern
    PrbsPattern_null PrbsPattern = "null"

    // PN11 Pattern
    PrbsPattern_pn11 PrbsPattern = "pn11"

    // PN23 Pattern
    PrbsPattern_pn23 PrbsPattern = "pn23"

    // PN31 Pattern
    PrbsPattern_pn31 PrbsPattern = "pn31"
)

// OtuThreshold represents Otu threshold
type OtuThreshold string

const (
    // PREFEC SD BER THRESHOLD
    OtuThreshold_prefec_sd OtuThreshold = "prefec-sd"

    // PREFEC SF BER THRESHOLD
    OtuThreshold_prefec_sf OtuThreshold = "prefec-sf"

    // OTU SD BER threshold
    OtuThreshold_otu_sd OtuThreshold = "otu-sd"

    // OTU SF BER threshold
    OtuThreshold_otu_sf OtuThreshold = "otu-sf"
)

// OtuAlarm represents Otu alarm
type OtuAlarm string

const (
    // LOS
    OtuAlarm_los OtuAlarm = "los"

    // LOF
    OtuAlarm_lof OtuAlarm = "lof"

    // LOM
    OtuAlarm_lom OtuAlarm = "lom"

    // OTU IAE
    OtuAlarm_iae OtuAlarm = "iae"

    // OTU BDI
    OtuAlarm_otu_bdi OtuAlarm = "otu-bdi"

    // OTU TIM
    OtuAlarm_otu_tim OtuAlarm = "otu-tim"

    // OTU SF BER
    OtuAlarm_otu_sf OtuAlarm = "otu-sf"

    // OTU SD BER
    OtuAlarm_otu_sd OtuAlarm = "otu-sd"

    // FEC mismatch
    OtuAlarm_fec_mismatch OtuAlarm = "fec-mismatch"

    // PREFEC SD BER
    OtuAlarm_prefec_sd_ber OtuAlarm = "prefec-sd-ber"

    // PREFEC SF BER
    OtuAlarm_prefec_sf_ber OtuAlarm = "prefec-sf-ber"
)

// OduThreshold represents Odu threshold
type OduThreshold string

const (
    // ODU SD BER threshold
    OduThreshold_odu_sd OduThreshold = "odu-sd"

    // ODU SF BER threshold
    OduThreshold_odu_sf OduThreshold = "odu-sf"
)

// TxTti represents Tx tti
type TxTti string

const (
    // TX TTI ascii string
    TxTti_tx_tti_ascii TxTti = "tx-tti-ascii"

    // TX TTI hex string
    TxTti_tx_tti_hex TxTti = "tx-tti-hex"
)

// Efec represents Efec
type Efec string

const (
    // default submode to handle backward
    // compatibility
    Efec_none Efec = "none"

    // efec i.4
    Efec_i__DOT__4 Efec = "i.4"

    // efec i.7
    Efec_i__DOT__7 Efec = "i.7"
)

// Proactive represents Proactive
type Proactive string

const (
    // Proactive Protection Default Mode
    Proactive_default_ Proactive = "default"

    // Proactive Protection Graceful Mode
    Proactive_graceful Proactive = "graceful"
)

// Framing represents Framing
type Framing string

const (
    // opu1e Framing
    Framing_opu1e Framing = "opu1e"

    // opu2e Framing
    Framing_opu2e Framing = "opu2e"
)

