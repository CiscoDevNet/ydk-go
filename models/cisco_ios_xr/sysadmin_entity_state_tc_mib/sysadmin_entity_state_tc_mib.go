// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_entity_state_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_entity_state_tc_mib"))
}

// EntityStandbyStatus
type EntityStandbyStatus string

const (
    EntityStandbyStatus_unknown EntityStandbyStatus = "unknown"

    EntityStandbyStatus_hotStandby EntityStandbyStatus = "hotStandby"

    EntityStandbyStatus_coldStandby EntityStandbyStatus = "coldStandby"

    EntityStandbyStatus_providingService EntityStandbyStatus = "providingService"
)

// EntityOperState
type EntityOperState string

const (
    EntityOperState_unknown EntityOperState = "unknown"

    EntityOperState_disabled EntityOperState = "disabled"

    EntityOperState_enabled EntityOperState = "enabled"

    EntityOperState_testing EntityOperState = "testing"
)

// EntityAdminState
type EntityAdminState string

const (
    EntityAdminState_unknown EntityAdminState = "unknown"

    EntityAdminState_locked EntityAdminState = "locked"

    EntityAdminState_shuttingDown EntityAdminState = "shuttingDown"

    EntityAdminState_unlocked EntityAdminState = "unlocked"
)

// EntityUsageState
type EntityUsageState string

const (
    EntityUsageState_unknown EntityUsageState = "unknown"

    EntityUsageState_idle EntityUsageState = "idle"

    EntityUsageState_active EntityUsageState = "active"

    EntityUsageState_busy EntityUsageState = "busy"
)

