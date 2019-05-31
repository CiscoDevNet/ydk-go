// This module contains a collection of YANG definitions
// for Cisco IOS-XR drivers-media-eth package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package drivers_media_eth_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package drivers_media_eth_cfg"))
}

// EthernetAutoNegotiation represents Ethernet auto negotiation
type EthernetAutoNegotiation string

const (
    // IEEE Standard auto-negotiation
    EthernetAutoNegotiation_true_ EthernetAutoNegotiation = "true"

    // Auto-negotiation with configuration override
    EthernetAutoNegotiation_override EthernetAutoNegotiation = "override"
)

// EthernetFec represents Ethernet fec
type EthernetFec string

const (
    // Disable any FEC enabled on the interface
    EthernetFec_none EthernetFec = "none"

    // Enable standard (Reed-Solomon) FEC
    EthernetFec_standard EthernetFec = "standard"
)

// EthernetFlowCtrl represents Ethernet flow ctrl
type EthernetFlowCtrl string

const (
    // Ingress flow control (sending pause frames)
    EthernetFlowCtrl_ingress EthernetFlowCtrl = "ingress"

    // Egress flow control (received pause frames)
    EthernetFlowCtrl_egress EthernetFlowCtrl = "egress"

    // Bi-direction flow control
    EthernetFlowCtrl_bidirectional EthernetFlowCtrl = "bidirectional"
)

// EthernetDuplex represents Ethernet duplex
type EthernetDuplex string

const (
    // Full duplex
    EthernetDuplex_full EthernetDuplex = "full"

    // Half duplex
    EthernetDuplex_half EthernetDuplex = "half"
)

// EthernetLoopback represents Ethernet loopback
type EthernetLoopback string

const (
    // External loopback (using loopback connector)
    EthernetLoopback_external EthernetLoopback = "external"

    // Internal loopback
    EthernetLoopback_internal EthernetLoopback = "internal"

    // Line loopback
    EthernetLoopback_line EthernetLoopback = "line"
)

// EthernetSpeed represents Ethernet speed
type EthernetSpeed string

const (
    // 10Mbits/s
    EthernetSpeed_Y_10 EthernetSpeed = "10"

    // 100Mbits/s
    EthernetSpeed_Y_100 EthernetSpeed = "100"

    // 1Gbits/s
    EthernetSpeed_Y_1000 EthernetSpeed = "1000"
)

// EthernetIpg represents Ethernet ipg
type EthernetIpg string

const (
    // Non standard IPG
    EthernetIpg_non_standard EthernetIpg = "non-standard"
)

// EthernetPfc represents Ethernet pfc
type EthernetPfc string

const (
    // Enable priority flow control
    EthernetPfc_on EthernetPfc = "on"
)

