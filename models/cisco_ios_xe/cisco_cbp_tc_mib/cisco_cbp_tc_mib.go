// This MIB module defines textual conventions used by the
// CISCO-CBP-BASE-CFG-MIB, CISCO-CBP-BASE-MON-MIB, and any MIB
// modules extending these MIB modules.
package cisco_cbp_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_cbp_tc_mib"))
}

// CbpExecutionStrategy represents         regardless of whether they succeed or fail.
type CbpExecutionStrategy string

const (
    CbpExecutionStrategy_other CbpExecutionStrategy = "other"

    CbpExecutionStrategy_doUntilSuccess CbpExecutionStrategy = "doUntilSuccess"

    CbpExecutionStrategy_doUntilFailure CbpExecutionStrategy = "doUntilFailure"

    CbpExecutionStrategy_doAll CbpExecutionStrategy = "doAll"
)

