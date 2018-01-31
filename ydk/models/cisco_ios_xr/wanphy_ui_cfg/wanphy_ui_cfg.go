// This module contains a collection of YANG definitions
// for Cisco IOS-XR wanphy-ui package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package wanphy_ui_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package wanphy_ui_cfg"))
}

// WanphyLanMode represents Wanphy lan mode
type WanphyLanMode string

const (
    // LanMode
    WanphyLanMode_on WanphyLanMode = "on"
)

// WanphyWanMode represents Wanphy wan mode
type WanphyWanMode string

const (
    // WAN Mode
    WanphyWanMode_on WanphyWanMode = "on"
)

