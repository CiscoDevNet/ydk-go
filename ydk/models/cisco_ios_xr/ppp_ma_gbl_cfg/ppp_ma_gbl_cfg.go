// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ma-gbl package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ppp_ma_gbl_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ppp_ma_gbl_cfg"))
}

// PppAuthenticationMethodGbl represents Ppp authentication method gbl
type PppAuthenticationMethodGbl string

const (
    // PAP
    PppAuthenticationMethodGbl_pap PppAuthenticationMethodGbl = "pap"

    // CHAP
    PppAuthenticationMethodGbl_chap PppAuthenticationMethodGbl = "chap"

    // MS CHAP
    PppAuthenticationMethodGbl_ms_chap PppAuthenticationMethodGbl = "ms-chap"
)

