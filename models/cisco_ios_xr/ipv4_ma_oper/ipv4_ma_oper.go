// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-ma package operational data.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ipv4-io-oper
// module with state data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_ma_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_ma_oper"))
}

// Ipv4MaOperConfig represents ipv4 client type
type Ipv4MaOperConfig string

const (
    // ipv4 ma oper client none
    Ipv4MaOperConfig_ipv4_ma_oper_client_none Ipv4MaOperConfig = "ipv4-ma-oper-client-none"

    // ipv4 ma oper non oc client
    Ipv4MaOperConfig_ipv4_ma_oper_non_oc_client Ipv4MaOperConfig = "ipv4-ma-oper-non-oc-client"

    // ipv4 ma oper oc client
    Ipv4MaOperConfig_ipv4_ma_oper_oc_client Ipv4MaOperConfig = "ipv4-ma-oper-oc-client"
)

// RpfMode represents Interface line states
type RpfMode string

const (
    // Strict RPF
    RpfMode_strict RpfMode = "strict"

    // Loose RPF
    RpfMode_loose RpfMode = "loose"
)

// Ipv4MaOperLineState represents Interface line states
type Ipv4MaOperLineState string

const (
    // Interface state is unknown
    Ipv4MaOperLineState_unknown Ipv4MaOperLineState = "unknown"

    // Interface has been shutdown
    Ipv4MaOperLineState_shutdown Ipv4MaOperLineState = "shutdown"

    // Interface state is down
    Ipv4MaOperLineState_down Ipv4MaOperLineState = "down"

    // Interface state is up
    Ipv4MaOperLineState_up Ipv4MaOperLineState = "up"
)

