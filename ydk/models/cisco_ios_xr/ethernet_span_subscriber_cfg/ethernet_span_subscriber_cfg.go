// This module contains a collection of YANG definitions
// for Cisco IOS-XR Ethernet-SPAN-subscriber package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_span_subscriber_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_span_subscriber_cfg"))
}

// SpanTrafficDirection represents Span traffic direction
type SpanTrafficDirection string

const (
    // Replicate only received (ingress) traffic
    SpanTrafficDirection_rx_only SpanTrafficDirection = "rx-only"

    // Replicate only transmitted (egress) traffic
    SpanTrafficDirection_tx_only SpanTrafficDirection = "tx-only"
)

// SpanMirrorInterval represents Span mirror interval
type SpanMirrorInterval string

const (
    // Mirror 1 in every 512 packets
    SpanMirrorInterval_Y_512 SpanMirrorInterval = "512"

    // Mirror 1 in every 1024 packets
    SpanMirrorInterval_Y_1k SpanMirrorInterval = "1k"

    // Mirror 1 in every 2048 packets
    SpanMirrorInterval_Y_2k SpanMirrorInterval = "2k"

    // Mirror 1 in every 4096 packets
    SpanMirrorInterval_Y_4k SpanMirrorInterval = "4k"

    // Mirror 1 in every 8192 packets
    SpanMirrorInterval_Y_8k SpanMirrorInterval = "8k"

    // Mirror 1 in every 16384 packets
    SpanMirrorInterval_Y_16k SpanMirrorInterval = "16k"
)

