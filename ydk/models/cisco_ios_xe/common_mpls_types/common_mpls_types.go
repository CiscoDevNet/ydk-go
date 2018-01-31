// This module contains a collection of generally useful YANG types
// that are specific to MPLS that can be usefully shared
// between multiple models.
// 
// Terms and Acronyms
// 
// MPLS: Multi Protocol Label Switching
package common_mpls_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package common_mpls_types"))
}

// IetfMplsLabel represents Temporary type until IETF definition
type IetfMplsLabel string

const (
    // IETF MPLS IPv4 explicit null Label (0)
    IetfMplsLabel_v4_explicit_null IetfMplsLabel = "v4-explicit-null"

    // IETF MPLS IPv6 explicit null label (2)
    IetfMplsLabel_v6_explicit_null IetfMplsLabel = "v6-explicit-null"

    // IETF MPLS implicit null label (3)
    IetfMplsLabel_implicit_null IetfMplsLabel = "implicit-null"
)

