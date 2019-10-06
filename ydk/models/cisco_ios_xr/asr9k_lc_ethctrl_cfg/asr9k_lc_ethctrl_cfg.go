// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-lc-ethctrl package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_lc_ethctrl_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_lc_ethctrl_cfg"))
}

// PermitPluggable represents Permit pluggable
type PermitPluggable string

const (
    // ALL types
    PermitPluggable_all PermitPluggable = "all"
)

// EtherCtrlTransportMode represents Ether ctrl transport mode
type EtherCtrlTransportMode string

const (
    // WAN
    EtherCtrlTransportMode_wan EtherCtrlTransportMode = "wan"

    // OTNOPUle
    EtherCtrlTransportMode_otnopu1e EtherCtrlTransportMode = "otnopu1e"

    // OTNOPU2e
    EtherCtrlTransportMode_otnopu2e EtherCtrlTransportMode = "otnopu2e"
)

// PermitPluggablePid represents Permit pluggable pid
type PermitPluggablePid string

const (
    // ALL PIDs
    PermitPluggablePid_all PermitPluggablePid = "all"
)

