// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ptp_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ptp_datatypes"))
}

// PtpTelecomClock represents Ptp telecom clock
type PtpTelecomClock string

const (
    // Telecom grandmaster clock
    PtpTelecomClock_telecom_grandmaster PtpTelecomClock = "telecom-grandmaster"

    // Telecom boundary clock
    PtpTelecomClock_telecom_boundary PtpTelecomClock = "telecom-boundary"

    // Telecom slave clock
    PtpTelecomClock_telecom_slave PtpTelecomClock = "telecom-slave"
)

// PtpInvalidUnicastGrantRequestResponse represents Ptp invalid unicast grant request response
type PtpInvalidUnicastGrantRequestResponse string

const (
    // Reduce grant parameters
    PtpInvalidUnicastGrantRequestResponse_reduce PtpInvalidUnicastGrantRequestResponse = "reduce"

    // Deny grant
    PtpInvalidUnicastGrantRequestResponse_deny PtpInvalidUnicastGrantRequestResponse = "deny"
)

// PtpClockId represents Ptp clock id
type PtpClockId string

const (
    // Use the router's MAC
    PtpClockId_router_mac PtpClockId = "router-mac"

    // Use a user-specified MAC
    PtpClockId_user_mac PtpClockId = "user-mac"

    // Use a user-specified EUI-64 number
    PtpClockId_eui PtpClockId = "eui"
)

// PtpClockSelectionMode represents Ptp clock selection mode
type PtpClockSelectionMode string

const (
    // Use 1588v2 clock selection
    PtpClockSelectionMode_Y_1588v2 PtpClockSelectionMode = "1588v2"

    // Use Telecom Profile clock selection
    PtpClockSelectionMode_telecom_profile PtpClockSelectionMode = "telecom-profile"
)

// PtpClockOperation represents Ptp clock operation
type PtpClockOperation string

const (
    // Two-step clock operation
    PtpClockOperation_two_step PtpClockOperation = "two-step"

    // One-step clock operation
    PtpClockOperation_one_step PtpClockOperation = "one-step"
)

// PtpTimePeriod represents Ptp time period
type PtpTimePeriod string

const (
    // One
    PtpTimePeriod_Y_1 PtpTimePeriod = "1"

    // Two
    PtpTimePeriod_Y_2 PtpTimePeriod = "2"

    // Four
    PtpTimePeriod_Y_4 PtpTimePeriod = "4"

    // Eight
    PtpTimePeriod_Y_8 PtpTimePeriod = "8"

    // Sixteen
    PtpTimePeriod_Y_16 PtpTimePeriod = "16"

    // Thirty Two
    PtpTimePeriod_Y_32 PtpTimePeriod = "32"

    // Sixty Four
    PtpTimePeriod_Y_64 PtpTimePeriod = "64"

    // One Hundred and Twenty-Eight
    PtpTimePeriod_Y_128 PtpTimePeriod = "128"
)

// PtpTransport represents Ptp transport
type PtpTransport string

const (
    // Unicast communication
    PtpTransport_unicast PtpTransport = "unicast"

    // Mixed-mode communication
    PtpTransport_mixed_mode PtpTransport = "mixed-mode"

    // Multicast communication
    PtpTransport_multicast PtpTransport = "multicast"
)

// PtpDelayAsymmetryUnits represents Ptp delay asymmetry units
type PtpDelayAsymmetryUnits string

const (
    // Nanoseconds
    PtpDelayAsymmetryUnits_nanoseconds PtpDelayAsymmetryUnits = "nanoseconds"

    // Microseconds
    PtpDelayAsymmetryUnits_microseconds PtpDelayAsymmetryUnits = "microseconds"

    // Milliseconds
    PtpDelayAsymmetryUnits_milliseconds PtpDelayAsymmetryUnits = "milliseconds"
)

// PtpTimescale represents Ptp timescale
type PtpTimescale string

const (
    // PTP timescale
    PtpTimescale_ptp PtpTimescale = "ptp"

    // ARB timescale
    PtpTimescale_arb PtpTimescale = "arb"
)

// PtpClockProfile represents Ptp clock profile
type PtpClockProfile string

const (
    // Default clock profile
    PtpClockProfile_default_ PtpClockProfile = "default"

    // G.8265.1 profile
    PtpClockProfile_g82651 PtpClockProfile = "g82651"

    // G.8275.1 profile
    PtpClockProfile_g82751 PtpClockProfile = "g82751"

    // G.8275.2 profile
    PtpClockProfile_g82752 PtpClockProfile = "g82752"
)

// PtpPortState represents Ptp port state
type PtpPortState string

const (
    // Any port state allowed
    PtpPortState_any PtpPortState = "any"

    // Restrict to slave
    PtpPortState_slave_only PtpPortState = "slave-only"

    // Restrict to master
    PtpPortState_master_only PtpPortState = "master-only"
)

// PtpTimeSource represents Ptp time source
type PtpTimeSource string

const (
    // Unknown
    PtpTimeSource_unknown PtpTimeSource = "unknown"

    // Atomic Clock
    PtpTimeSource_atomic_clock PtpTimeSource = "atomic-clock"

    // GPS
    PtpTimeSource_gps PtpTimeSource = "gps"

    // Terrestrial Radio
    PtpTimeSource_terrestrial_radio PtpTimeSource = "terrestrial-radio"

    // PTP
    PtpTimeSource_ptp PtpTimeSource = "ptp"

    // NTP
    PtpTimeSource_ntp PtpTimeSource = "ntp"

    // Hand set
    PtpTimeSource_hand_set PtpTimeSource = "hand-set"

    // Other
    PtpTimeSource_other PtpTimeSource = "other"

    // Internal Oscillator
    PtpTimeSource_internal_oscillator PtpTimeSource = "internal-oscillator"
)

// PtpEncap represents Ptp encap
type PtpEncap string

const (
    // Ethernet Encapsulation
    PtpEncap_ethernet PtpEncap = "ethernet"

    // IPv4 Encapsulation
    PtpEncap_ipv4 PtpEncap = "ipv4"

    // IPv6 Encapsulation
    PtpEncap_ipv6 PtpEncap = "ipv6"
)

// PtpTime represents Ptp time
type PtpTime string

const (
    // Time interval in seconds
    PtpTime_interval PtpTime = "interval"

    // Frequency per second
    PtpTime_frequency PtpTime = "frequency"
)

// PtpClockAdvertisementMode represents Ptp clock advertisement mode
type PtpClockAdvertisementMode string

const (
    // Use 1588v2 clock selection
    PtpClockAdvertisementMode_Y_1588v2 PtpClockAdvertisementMode = "1588v2"

    // Use Telecom Profile clock selection
    PtpClockAdvertisementMode_telecom_profile PtpClockAdvertisementMode = "telecom-profile"
)

