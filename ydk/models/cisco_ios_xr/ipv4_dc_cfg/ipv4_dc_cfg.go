// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-dc package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_dc_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_dc_cfg"))
}

// Opt61Sub represents Opt61 sub
type Opt61Sub string

const (
    // ascii
    Opt61Sub_ascii Opt61Sub = "ascii"

    // sn chassis
    Opt61Sub_sn_chassis Opt61Sub = "sn-chassis"
)

// Option61 represents Option61
type Option61 string

const (
    // OPTION61
    Option61_option61 Option61 = "option61"
)

// Option60 represents Option60
type Option60 string

const (
    // OPTION60
    Option60_option60 Option60 = "option60"
)

// Option77 represents Option77
type Option77 string

const (
    // OPTION77
    Option77_option77 Option77 = "option77"
)

