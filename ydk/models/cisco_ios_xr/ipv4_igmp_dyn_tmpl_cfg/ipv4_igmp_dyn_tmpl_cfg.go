// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-igmp-dyn-tmpl package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_igmp_dyn_tmpl_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_igmp_dyn_tmpl_cfg"))
}

// DynTmplMulticastMode represents Dyn tmpl multicast mode
type DynTmplMulticastMode string

const (
    // QOS Correlation
    DynTmplMulticastMode_qos_correlation DynTmplMulticastMode = "qos-correlation"

    // Passive
    DynTmplMulticastMode_passive DynTmplMulticastMode = "passive"
)

