// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-xtc-agent package configuration.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_xtc_agent_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_xtc_agent_cfg"))
}

// XtcMetricValue represents Xtc metric value
type XtcMetricValue string

const (
    // Relative metric value
    XtcMetricValue_relative XtcMetricValue = "relative"

    // Absolute metric value
    XtcMetricValue_absolute XtcMetricValue = "absolute"
)

// XtcAffinityRule represents Xtc affinity rule
type XtcAffinityRule string

const (
    // Include-all affinity rule
    XtcAffinityRule_affinity_include_all XtcAffinityRule = "affinity-include-all"

    // Exclude-any affinity rule
    XtcAffinityRule_affinity_exclude_any XtcAffinityRule = "affinity-exclude-any"

    // Include-any affinity rule
    XtcAffinityRule_affinity_include_any XtcAffinityRule = "affinity-include-any"
)

