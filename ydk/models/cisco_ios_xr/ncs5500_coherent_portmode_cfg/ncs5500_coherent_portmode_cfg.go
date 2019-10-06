// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs5500-coherent-portmode package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// ModSel represents Mod sel
type ModSel string

const (
    // select qpsk for modulation
    ModSel_qpsk ModSel = "qpsk"

    // select 8qam for modulation
    ModSel_Y_8qam ModSel = "8qam"

    // select 16qam for modulation
    ModSel_Y_16qam ModSel = "16qam"
)

// FecSel represents Fec sel
type FecSel string

const (
    // 15%-SD Forward Error Correction
    FecSel_Y_15sdfec FecSel = "15sdfec"

    // 25%-SD Forward Error Correction
    FecSel_Y_25sdfec FecSel = "25sdfec"

    // 15%-SD Forward Error Correction with Diff
    FecSel_Y_15sdfecde FecSel = "15sdfecde"

    // 7%-STAIRCASE Forward Error Correction
    FecSel_otu7staircase FecSel = "otu7staircase"
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

