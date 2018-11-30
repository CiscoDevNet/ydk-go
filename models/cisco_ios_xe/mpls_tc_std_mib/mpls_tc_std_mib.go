// Copyright (C) The Internet Society (2004). The
// initial version of this MIB module was published
// in RFC 3811. For full legal notices see the RFC
// itself or see:
// http://www.ietf.org/copyrights/ianamib.html
// 
// This MIB module defines TEXTUAL-CONVENTIONs
// for concepts used in Multiprotocol Label
// Switching (MPLS) networks.
package mpls_tc_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_tc_std_mib"))
}

// MplsLabelDistributionMethod represents on Demand.
type MplsLabelDistributionMethod string

const (
    MplsLabelDistributionMethod_downstreamOnDemand MplsLabelDistributionMethod = "downstreamOnDemand"

    MplsLabelDistributionMethod_downstreamUnsolicited MplsLabelDistributionMethod = "downstreamUnsolicited"
)

// MplsRetentionMode represents a valid next hop or not.
type MplsRetentionMode string

const (
    MplsRetentionMode_conservative MplsRetentionMode = "conservative"

    MplsRetentionMode_liberal MplsRetentionMode = "liberal"
)

// MplsLdpLabelType represents frameRelay(3).
type MplsLdpLabelType string

const (
    MplsLdpLabelType_generic MplsLdpLabelType = "generic"

    MplsLdpLabelType_atm MplsLdpLabelType = "atm"

    MplsLdpLabelType_frameRelay MplsLdpLabelType = "frameRelay"
)

// TeHopAddressType represents type changes (e.g., from ipv6(2) to ipv4(1)).
type TeHopAddressType string

const (
    TeHopAddressType_unknown TeHopAddressType = "unknown"

    TeHopAddressType_ipv4 TeHopAddressType = "ipv4"

    TeHopAddressType_ipv6 TeHopAddressType = "ipv6"

    TeHopAddressType_asnumber TeHopAddressType = "asnumber"

    TeHopAddressType_unnum TeHopAddressType = "unnum"

    TeHopAddressType_lspid TeHopAddressType = "lspid"
)

// MplsLspType represents                          LSR.
type MplsLspType string

const (
    MplsLspType_unknown MplsLspType = "unknown"

    MplsLspType_terminatingLsp MplsLspType = "terminatingLsp"

    MplsLspType_originatingLsp MplsLspType = "originatingLsp"

    MplsLspType_crossConnectingLsp MplsLspType = "crossConnectingLsp"
)

// MplsOwner represents different choice.
type MplsOwner string

const (
    MplsOwner_unknown MplsOwner = "unknown"

    MplsOwner_other MplsOwner = "other"

    MplsOwner_snmp MplsOwner = "snmp"

    MplsOwner_ldp MplsOwner = "ldp"

    MplsOwner_crldp MplsOwner = "crldp"

    MplsOwner_rsvpTe MplsOwner = "rsvpTe"

    MplsOwner_policyAgent MplsOwner = "policyAgent"
)

