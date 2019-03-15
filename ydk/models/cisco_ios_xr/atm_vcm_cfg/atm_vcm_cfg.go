// This module contains a collection of YANG definitions
// for Cisco IOS-XR atm-vcm package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-l2vpn-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package atm_vcm_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package atm_vcm_cfg"))
}

// AtmVpiBitsMode represents Atm vpi bits mode
type AtmVpiBitsMode string

const (
    // 12-bits VPI cell format
    AtmVpiBitsMode_twelve AtmVpiBitsMode = "twelve"
)

// AtmPvcTestMode represents Atm pvc test mode
type AtmPvcTestMode string

const (
    // Loop mode applicable to L2/L3 PVC
    AtmPvcTestMode_loop AtmPvcTestMode = "loop"

    // Reserved mode applicable to L2 PVC
    AtmPvcTestMode_reserved AtmPvcTestMode = "reserved"
)

// AtmPvpTestMode represents Atm pvp test mode
type AtmPvpTestMode string

const (
    // Loop mode
    AtmPvpTestMode_loop AtmPvpTestMode = "loop"
)

