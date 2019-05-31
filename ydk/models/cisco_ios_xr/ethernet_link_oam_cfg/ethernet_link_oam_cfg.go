// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-link-oam package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg,
//   Cisco-IOS-XR-l2-eth-infra-cfg,
//   Cisco-IOS-XR-ifmgr-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_link_oam_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_link_oam_cfg"))
}

// EtherLinkOamThresholdWindowMultiplierEnum represents Ether link oam threshold window multiplier enum
type EtherLinkOamThresholdWindowMultiplierEnum string

const (
    // Do not use a multiplier
    EtherLinkOamThresholdWindowMultiplierEnum_none EtherLinkOamThresholdWindowMultiplierEnum = "none"

    // Use multiplier of 1000
    EtherLinkOamThresholdWindowMultiplierEnum_thousand EtherLinkOamThresholdWindowMultiplierEnum = "thousand"

    // Use multiplier of 1,000,000
    EtherLinkOamThresholdWindowMultiplierEnum_million EtherLinkOamThresholdWindowMultiplierEnum = "million"

    // Use multiplier of 1,000,000,000
    EtherLinkOamThresholdWindowMultiplierEnum_billion EtherLinkOamThresholdWindowMultiplierEnum = "billion"
)

// EtherLinkOamThresholdUnitsFramesEnum represents Ether link oam threshold units frames enum
type EtherLinkOamThresholdUnitsFramesEnum string

const (
    // Define threshold in frames
    EtherLinkOamThresholdUnitsFramesEnum_frames EtherLinkOamThresholdUnitsFramesEnum = "frames"

    // Define threshold in parts per million
    EtherLinkOamThresholdUnitsFramesEnum_ppm EtherLinkOamThresholdUnitsFramesEnum = "ppm"
)

// EtherLinkOamThresholdUnitsSymbolsEnum represents Ether link oam threshold units symbols enum
type EtherLinkOamThresholdUnitsSymbolsEnum string

const (
    // Define threshold in symbols
    EtherLinkOamThresholdUnitsSymbolsEnum_symbols EtherLinkOamThresholdUnitsSymbolsEnum = "symbols"

    // Define threshold in parts per million
    EtherLinkOamThresholdUnitsSymbolsEnum_ppm EtherLinkOamThresholdUnitsSymbolsEnum = "ppm"
)

// EtherLinkOamWindowUnitsSymbolsEnum represents Ether link oam window units symbols enum
type EtherLinkOamWindowUnitsSymbolsEnum string

const (
    // Define window in milliseconds
    EtherLinkOamWindowUnitsSymbolsEnum_milliseconds EtherLinkOamWindowUnitsSymbolsEnum = "milliseconds"

    // Define window in symbols
    EtherLinkOamWindowUnitsSymbolsEnum_symbols EtherLinkOamWindowUnitsSymbolsEnum = "symbols"
)

// EtherLinkOamWindowUnitsFramesEnum represents Ether link oam window units frames enum
type EtherLinkOamWindowUnitsFramesEnum string

const (
    // Define window in milliseconds
    EtherLinkOamWindowUnitsFramesEnum_milliseconds EtherLinkOamWindowUnitsFramesEnum = "milliseconds"

    // Define window in frames
    EtherLinkOamWindowUnitsFramesEnum_frames EtherLinkOamWindowUnitsFramesEnum = "frames"
)

// EtherLinkOamRequireModeEnum represents Ether link oam require mode enum
type EtherLinkOamRequireModeEnum string

const (
    // Ethernet Link OAM Passive mode
    EtherLinkOamRequireModeEnum_passive EtherLinkOamRequireModeEnum = "passive"

    // Ethernet Link OAM Active mode
    EtherLinkOamRequireModeEnum_active EtherLinkOamRequireModeEnum = "active"

    // Ethernet Link OAM mode not required
    EtherLinkOamRequireModeEnum_dont_care EtherLinkOamRequireModeEnum = "dont-care"
)

// EtherLinkOamEventActionEnumEfd represents Ether link oam event action enum efd
type EtherLinkOamEventActionEnumEfd string

const (
    // Perform no action
    EtherLinkOamEventActionEnumEfd_disable EtherLinkOamEventActionEnumEfd = "disable"

    // Disable the interface
    EtherLinkOamEventActionEnumEfd_error_disable EtherLinkOamEventActionEnumEfd = "error-disable"

    // Log the event
    EtherLinkOamEventActionEnumEfd_log EtherLinkOamEventActionEnumEfd = "log"

    // EFD the interface
    EtherLinkOamEventActionEnumEfd_efd EtherLinkOamEventActionEnumEfd = "efd"
)

// EtherLinkOamEventActionPrimEnum represents Ether link oam event action prim enum
type EtherLinkOamEventActionPrimEnum string

const (
    // Perform no action
    EtherLinkOamEventActionPrimEnum_disable EtherLinkOamEventActionPrimEnum = "disable"

    // Log the event
    EtherLinkOamEventActionPrimEnum_log EtherLinkOamEventActionPrimEnum = "log"
)

// EtherLinkOamModeEnum represents Ether link oam mode enum
type EtherLinkOamModeEnum string

const (
    // Ethernet Link OAM Passive mode
    EtherLinkOamModeEnum_passive EtherLinkOamModeEnum = "passive"

    // Ethernet Link OAM Active mode
    EtherLinkOamModeEnum_active EtherLinkOamModeEnum = "active"
)

// EtherLinkOamEventActionEnum represents Ether link oam event action enum
type EtherLinkOamEventActionEnum string

const (
    // Perform no action
    EtherLinkOamEventActionEnum_disable EtherLinkOamEventActionEnum = "disable"

    // Disable the interface
    EtherLinkOamEventActionEnum_error_disable EtherLinkOamEventActionEnum = "error-disable"

    // Log the event
    EtherLinkOamEventActionEnum_log EtherLinkOamEventActionEnum = "log"
)

// EtherLinkOamHelloIntervalEnum represents Ether link oam hello interval enum
type EtherLinkOamHelloIntervalEnum string

const (
    // 1 s OAM hello interval
    EtherLinkOamHelloIntervalEnum_Y_1s EtherLinkOamHelloIntervalEnum = "1s"

    // 100 ms OAM hello interval
    EtherLinkOamHelloIntervalEnum_Y_100ms EtherLinkOamHelloIntervalEnum = "100ms"
)

