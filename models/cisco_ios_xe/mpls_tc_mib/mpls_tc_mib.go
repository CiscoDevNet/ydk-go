// This MIB module defines Textual Conventions and
// OBJECT-IDENTITIES for use in documents defining
// management information bases (MIBs) for managing
// MPLS networks.
package mpls_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_tc_mib"))
}

// MplsLdpLabelTypes represents frameRelay(3).
type MplsLdpLabelTypes string

const (
    MplsLdpLabelTypes_generic MplsLdpLabelTypes = "generic"

    MplsLdpLabelTypes_atm MplsLdpLabelTypes = "atm"

    MplsLdpLabelTypes_frameRelay MplsLdpLabelTypes = "frameRelay"
)

// MplsInitialCreationSource represents component created the object.
type MplsInitialCreationSource string

const (
    MplsInitialCreationSource_other MplsInitialCreationSource = "other"

    MplsInitialCreationSource_snmp MplsInitialCreationSource = "snmp"

    MplsInitialCreationSource_ldp MplsInitialCreationSource = "ldp"

    MplsInitialCreationSource_rsvp MplsInitialCreationSource = "rsvp"

    MplsInitialCreationSource_crldp MplsInitialCreationSource = "crldp"

    MplsInitialCreationSource_policyAgent MplsInitialCreationSource = "policyAgent"

    MplsInitialCreationSource_unknown MplsInitialCreationSource = "unknown"
)

