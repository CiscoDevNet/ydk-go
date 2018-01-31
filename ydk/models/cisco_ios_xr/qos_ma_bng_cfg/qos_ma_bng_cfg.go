// This module contains a collection of YANG definitions
// for Cisco IOS-XR qos-ma-bng package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package qos_ma_bng_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package qos_ma_bng_cfg"))
}

// Qosl2DataLink represents Qosl2 data link
type Qosl2DataLink string

const (
    // ATM adaption layer AAL5
    Qosl2DataLink_aal5 Qosl2DataLink = "aal5"
)

// Qosl2Encap represents Qosl2 encap
type Qosl2Encap string

const (
    // snap-pppoa encap used between the DSLAM and CPE
    Qosl2Encap_snap_pppoa Qosl2Encap = "snap-pppoa"

    // mux-pppoa encap used between the DSLAM and CPE
    Qosl2Encap_mux_pppoa Qosl2Encap = "mux-pppoa"

    // snap-1483routed encap used between the DSLAM
    // and CPE
    Qosl2Encap_snap1483_routed Qosl2Encap = "snap1483-routed"

    // mux-1483routed encap used between the DSLAM and
    // CPE
    Qosl2Encap_mux1483_routed Qosl2Encap = "mux1483-routed"

    // snap-rbe encap used between the DSLAM and CPE
    Qosl2Encap_snap_rbe Qosl2Encap = "snap-rbe"

    // snap-dot1q-rbe encap used between the DSLAM and
    // CPE
    Qosl2Encap_snap_dot1qrbe Qosl2Encap = "snap-dot1qrbe"

    // mux-rbe encap used between the DSLAM and CPE
    Qosl2Encap_mux_rbe Qosl2Encap = "mux-rbe"

    // mux-dot1q-rbe encap used between the DSLAM and
    // CPE
    Qosl2Encap_mux_dot1qrbe Qosl2Encap = "mux-dot1qrbe"
)

