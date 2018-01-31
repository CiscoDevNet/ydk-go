// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs5500-coherent-portmode package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs5500_coherent_portmode_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs5500_coherent_portmode_cfg"))
}

// DiffSel represents Diff sel
type DiffSel string

const (
    // disable differential
    DiffSel_disable DiffSel = "disable"

    // enable differential
    DiffSel_enable DiffSel = "enable"
)

// FecSel represents Fec sel
type FecSel string

const (
    // FEC 15
    FecSel_Y_15percent FecSel = "15percent"

    // FEC 25
    FecSel_Y_25percent FecSel = "25percent"
)

// SpeedSel represents Speed sel
type SpeedSel string

const (
    // Speed 100
    SpeedSel_Y_100g SpeedSel = "100g"

    // Speed 150
    SpeedSel_Y_150g SpeedSel = "150g"

    // Speed 200
    SpeedSel_Y_200g SpeedSel = "200g"
)

