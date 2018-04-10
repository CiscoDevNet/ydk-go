// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-ltrace package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-config-mda-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_ltrace_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_ltrace_cfg"))
}

// InfraLtraceMode represents Infra ltrace mode
type InfraLtraceMode string

const (
    // Set ltrace memory allocation to static mode
    InfraLtraceMode_static InfraLtraceMode = "static"

    // Set ltrace memory allocation to dynamic mode
    InfraLtraceMode_dynamic InfraLtraceMode = "dynamic"
)

// InfraLtraceScale represents Infra ltrace scale
type InfraLtraceScale string

const (
    // Use platform-defined scale-factor
    InfraLtraceScale_Y_0 InfraLtraceScale = "0"

    // Do not scale down ltrace memory request
    InfraLtraceScale_Y_1 InfraLtraceScale = "1"

    // Scale down ltrace memory request by 2
    InfraLtraceScale_Y_2 InfraLtraceScale = "2"

    // Scale down ltrace memory request by 4
    InfraLtraceScale_Y_4 InfraLtraceScale = "4"

    // Scale down ltrace memory request by 8
    InfraLtraceScale_Y_8 InfraLtraceScale = "8"

    // Scale down ltrace memory request by 16
    InfraLtraceScale_Y_16 InfraLtraceScale = "16"
)

