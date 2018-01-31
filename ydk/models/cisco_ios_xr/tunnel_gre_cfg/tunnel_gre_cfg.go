// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-gre package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tunnel_gre_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_gre_cfg"))
}

// TunnelModeDirection represents Tunnel mode direction
type TunnelModeDirection string

const (
    // Decap-only tunnel
    TunnelModeDirection_decap TunnelModeDirection = "decap"

    // Encap-only tunnel
    TunnelModeDirection_encap TunnelModeDirection = "encap"
)

