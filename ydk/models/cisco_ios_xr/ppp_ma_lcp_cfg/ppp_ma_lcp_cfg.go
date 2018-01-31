// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ma-lcp package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ppp_ma_lcp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ppp_ma_lcp_cfg"))
}

// PppAuthenticationMethod represents Ppp authentication method
type PppAuthenticationMethod string

const (
    // PAP
    PppAuthenticationMethod_pap PppAuthenticationMethod = "pap"

    // CHAP
    PppAuthenticationMethod_chap PppAuthenticationMethod = "chap"

    // MS CHAP
    PppAuthenticationMethod_ms_chap PppAuthenticationMethod = "ms-chap"
)

