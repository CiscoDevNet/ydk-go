// This module contains a collection of YANG definitions
// common for all UTD operational data.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package utd_common_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package utd_common_oper"))
}

// UtdUpdateStatusVal represents Unified Threat Defense (UTD) update status
type UtdUpdateStatusVal string

const (
    // Unified Threat Defense (UTD) update status is unknown
    UtdUpdateStatusVal_utd_update_status_unknown UtdUpdateStatusVal = "utd-update-status-unknown"

    // Unified Threat Defense (UTD) update successful
    UtdUpdateStatusVal_utd_update_status_success UtdUpdateStatusVal = "utd-update-status-success"

    // Unified Threat Defense (UTD) update failed
    UtdUpdateStatusVal_utd_update_status_failure UtdUpdateStatusVal = "utd-update-status-failure"

    // Unified Threat Defense (UTD) update attemped but not required
    UtdUpdateStatusVal_utd_update_status_no_update UtdUpdateStatusVal = "utd-update-status-no-update"
)

