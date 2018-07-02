// Types related to the OpenConfig Abstract Forwarding
// Table (AFT) model
package aft_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aft_types"))
}

// EncapsulationHeaderType represents head- or tail-end.
type EncapsulationHeaderType string

const (
    // The encapsulation header is a Generic Routing Encapsulation
    // header.
    EncapsulationHeaderType_GRE EncapsulationHeaderType = "GRE"

    // The encapsulation header is an IPv4 packet header
    EncapsulationHeaderType_IPV4 EncapsulationHeaderType = "IPV4"

    // The encapsulation header is an IPv6 packet header
    EncapsulationHeaderType_IPV6 EncapsulationHeaderType = "IPV6"

    // The encapsulation header is one or more MPLS labels indicated
    // by the pushed and popped label stack lists.
    EncapsulationHeaderType_MPLS EncapsulationHeaderType = "MPLS"
)

