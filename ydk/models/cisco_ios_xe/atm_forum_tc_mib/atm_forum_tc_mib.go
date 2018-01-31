package atm_forum_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package atm_forum_tc_mib"))
}

// TruthValue represents Boolean values use this data type from RFC-1903
type TruthValue string

const (
    TruthValue_true TruthValue = "true"

    TruthValue_false TruthValue = "false"
)

// AtmServiceCategory represents ATM Service Categories use this data type
type AtmServiceCategory string

const (
    AtmServiceCategory_other AtmServiceCategory = "other"

    AtmServiceCategory_cbr AtmServiceCategory = "cbr"

    AtmServiceCategory_rtVbr AtmServiceCategory = "rtVbr"

    AtmServiceCategory_nrtVbr AtmServiceCategory = "nrtVbr"

    AtmServiceCategory_abr AtmServiceCategory = "abr"

    AtmServiceCategory_ubr AtmServiceCategory = "ubr"
)

