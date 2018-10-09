// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_span_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_span_datatypes"))
}

// SpanSessionClassOld represents Span session class old
type SpanSessionClassOld string

const (
    // Mirror Ethernet packets
    SpanSessionClassOld_true_ SpanSessionClassOld = "true"
)

// SpanSessionClass represents Span session class
type SpanSessionClass string

const (
    // Mirror Ethernet packets
    SpanSessionClass_ethernet SpanSessionClass = "ethernet"

    // Mirror IPv4 packets
    SpanSessionClass_ipv4 SpanSessionClass = "ipv4"

    // Mirror IPv6 packets
    SpanSessionClass_ipv6 SpanSessionClass = "ipv6"

    // Mirror MPLS-encapsulated IPv4 packets
    SpanSessionClass_mpls_ipv4 SpanSessionClass = "mpls-ipv4"

    // Mirror MPLS-encapsulated IPv6 packets
    SpanSessionClass_mpls_ipv6 SpanSessionClass = "mpls-ipv6"
)

