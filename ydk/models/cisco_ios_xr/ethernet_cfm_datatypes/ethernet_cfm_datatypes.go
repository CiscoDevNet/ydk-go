// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_cfm_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_cfm_datatypes"))
}

// CfmCcmInterval represents Cfm ccm interval
type CfmCcmInterval string

const (
    // 3.3ms
    CfmCcmInterval_Y_3__DOT__3ms CfmCcmInterval = "3.3ms"

    // 10ms
    CfmCcmInterval_Y_10ms CfmCcmInterval = "10ms"

    // 100ms
    CfmCcmInterval_Y_100ms CfmCcmInterval = "100ms"

    // 1s
    CfmCcmInterval_Y_1s CfmCcmInterval = "1s"

    // 10s
    CfmCcmInterval_Y_10s CfmCcmInterval = "10s"

    // 1m
    CfmCcmInterval_Y_1m CfmCcmInterval = "1m"

    // 10m
    CfmCcmInterval_Y_10m CfmCcmInterval = "10m"
)

// BandwidthNotificationState represents Bandwidth notification state
type BandwidthNotificationState string

const (
    // Link is not degraded
    BandwidthNotificationState_ok BandwidthNotificationState = "ok"

    // Link is in degraded state
    BandwidthNotificationState_degraded BandwidthNotificationState = "degraded"
)

// CfmAisInterval represents Cfm ais interval
type CfmAisInterval string

const (
    // 1s
    CfmAisInterval_Y_1s CfmAisInterval = "1s"

    // 1m
    CfmAisInterval_Y_1m CfmAisInterval = "1m"
)

// CfmMepDir represents Cfm mep dir
type CfmMepDir string

const (
    // Up MEP
    CfmMepDir_up CfmMepDir = "up"

    // Down MEP
    CfmMepDir_down CfmMepDir = "down"
)

