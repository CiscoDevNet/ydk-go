// Configuration and operational state parameters relating to the
// segment routing. This module defines a number of elements which are
// instantiated in multiple places throughout the OpenConfig collection
// of models.
// 
// Particularly:
//  - SRGB+LB dataplane instances - directly instantied by SR.
//  - SRGB+LB dataplane reservations - instantiated within MPLS and future SR
//                                  dataplanes.
//  - SR SID advertisements - instantiated within the relevant IGP.
//  - SR-specific counters - instantied within the relevant dataplane.
package segment_routing

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package segment_routing"))
}

// SrDataplaneType represents Routing block of SIDs.
type SrDataplaneType string

const (
    // The entity uses MPLS labels as Segment Identifiers.
    SrDataplaneType_MPLS SrDataplaneType = "MPLS"

    // The entity uses IPv6 prefixes as Segment Identifiers.
    SrDataplaneType_IPV6 SrDataplaneType = "IPV6"
)

// MplsLabel represents type for MPLS label value encoding
type MplsLabel string

const (
    // valid at the bottom of the label stack,
    // indicates that stack must be popped and packet forwarded
    // based on IPv4 header
    MplsLabel_IPV4_EXPLICIT_NULL MplsLabel = "IPV4_EXPLICIT_NULL"

    // allowed anywhere in the label stack except
    // the bottom, local router delivers packet to the local CPU
    // when this label is at the top of the stack
    MplsLabel_ROUTER_ALERT MplsLabel = "ROUTER_ALERT"

    // valid at the bottom of the label stack,
    // indicates that stack must be popped and packet forwarded
    // based on IPv6 header
    MplsLabel_IPV6_EXPLICIT_NULL MplsLabel = "IPV6_EXPLICIT_NULL"

    // assigned by local LSR but not carried in
    // packets
    MplsLabel_IMPLICIT_NULL MplsLabel = "IMPLICIT_NULL"

    // Entropy label indicator, to allow an LSR
    // to distinguish between entropy label and applicaiton
    // labels RFC 6790
    MplsLabel_ENTROPY_LABEL_INDICATOR MplsLabel = "ENTROPY_LABEL_INDICATOR"
)

