// This MIB defines state textual conventions.
// 
// Copyright (C) The Internet Society 2005.  This version
// of this MIB module is part of RFC 4268;  see the RFC
// itself for full legal notices.
package entity_state_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package entity_state_tc_mib"))
}

// EntityAdminState represents report administrative state.
type EntityAdminState string

const (
    EntityAdminState_unknown EntityAdminState = "unknown"

    EntityAdminState_locked EntityAdminState = "locked"

    EntityAdminState_shuttingDown EntityAdminState = "shuttingDown"

    EntityAdminState_unlocked EntityAdminState = "unlocked"
)

// EntityOperState represents resource is unable to report operational state.
type EntityOperState string

const (
    EntityOperState_unknown EntityOperState = "unknown"

    EntityOperState_disabled EntityOperState = "disabled"

    EntityOperState_enabled EntityOperState = "enabled"

    EntityOperState_testing EntityOperState = "testing"
)

// EntityUsageState represents that this resource is unable to report usage state.
type EntityUsageState string

const (
    EntityUsageState_unknown EntityUsageState = "unknown"

    EntityUsageState_idle EntityUsageState = "idle"

    EntityUsageState_active EntityUsageState = "active"

    EntityUsageState_busy EntityUsageState = "busy"
)

// EntityStandbyStatus represents report standby state.
type EntityStandbyStatus string

const (
    EntityStandbyStatus_unknown EntityStandbyStatus = "unknown"

    EntityStandbyStatus_hotStandby EntityStandbyStatus = "hotStandby"

    EntityStandbyStatus_coldStandby EntityStandbyStatus = "coldStandby"

    EntityStandbyStatus_providingService EntityStandbyStatus = "providingService"
)

