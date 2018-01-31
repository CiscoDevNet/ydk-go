// This module contains a collection of YANG definitions
// for Cisco IOS-XR drivers-icpe-ethernet package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package drivers_icpe_ethernet_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package drivers_icpe_ethernet_cfg"))
}

// ExtendedEthernetLoopback represents Extended ethernet loopback
type ExtendedEthernetLoopback string

const (
    // Internal loopback
    ExtendedEthernetLoopback_internal ExtendedEthernetLoopback = "internal"

    // Line loopback
    ExtendedEthernetLoopback_line ExtendedEthernetLoopback = "line"
)

