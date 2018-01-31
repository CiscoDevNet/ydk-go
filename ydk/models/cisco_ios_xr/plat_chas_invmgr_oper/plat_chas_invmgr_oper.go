// This module contains a collection of YANG definitions
// for Cisco IOS-XR plat-chas-invmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   platform: Platform information
//   platform-inventory: platform inventory
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package plat_chas_invmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package plat_chas_invmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-plat-chas-invmgr-oper platform}", reflect.TypeOf(Platform{}))
    ydk.RegisterEntity("Cisco-IOS-XR-plat-chas-invmgr-oper:platform", reflect.TypeOf(Platform{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-plat-chas-invmgr-oper platform-inventory}", reflect.TypeOf(PlatformInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-plat-chas-invmgr-oper:platform-inventory", reflect.TypeOf(PlatformInventory{}))
}

// InvAdminState represents Inv admin state
type InvAdminState string

const (
    // admin state invalid
    InvAdminState_admin_state_invalid InvAdminState = "admin-state-invalid"

    // admin up
    InvAdminState_admin_up InvAdminState = "admin-up"

    // admin down
    InvAdminState_admin_down InvAdminState = "admin-down"
)

// InvResetReason represents Inv reset reason
type InvResetReason string

const (
    // module reset reason unknown
    InvResetReason_module_reset_reason_unknown InvResetReason = "module-reset-reason-unknown"

    // module reset reason powerup
    InvResetReason_module_reset_reason_powerup InvResetReason = "module-reset-reason-powerup"

    // module reset reason user shutdown
    InvResetReason_module_reset_reason_user_shutdown InvResetReason = "module-reset-reason-user-shutdown"

    // module reset reason user reload
    InvResetReason_module_reset_reason_user_reload InvResetReason = "module-reset-reason-user-reload"

    // module reset reason auto reload
    InvResetReason_module_reset_reason_auto_reload InvResetReason = "module-reset-reason-auto-reload"

    // module reset reason environment
    InvResetReason_module_reset_reason_environment InvResetReason = "module-reset-reason-environment"

    // module reset reason user unpower
    InvResetReason_module_reset_reason_user_unpower InvResetReason = "module-reset-reason-user-unpower"
)

// InvCardState represents Inv card state
type InvCardState string

const (
    // inv card not present
    InvCardState_inv_card_not_present InvCardState = "inv-card-not-present"

    // inv card present
    InvCardState_inv_card_present InvCardState = "inv-card-present"

    // inv card reset
    InvCardState_inv_card_reset InvCardState = "inv-card-reset"

    // inv card booting
    InvCardState_inv_card_booting InvCardState = "inv-card-booting"

    // inv card mbi booting
    InvCardState_inv_card_mbi_booting InvCardState = "inv-card-mbi-booting"

    // inv card running mbi
    InvCardState_inv_card_running_mbi InvCardState = "inv-card-running-mbi"

    // inv card running ena
    InvCardState_inv_card_running_ena InvCardState = "inv-card-running-ena"

    // inv card bring down
    InvCardState_inv_card_bring_down InvCardState = "inv-card-bring-down"

    // inv card ena failure
    InvCardState_inv_card_ena_failure InvCardState = "inv-card-ena-failure"

    // inv card f diag run
    InvCardState_inv_card_f_diag_run InvCardState = "inv-card-f-diag-run"

    // inv card f diag failure
    InvCardState_inv_card_f_diag_failure InvCardState = "inv-card-f-diag-failure"

    // inv card powered
    InvCardState_inv_card_powered InvCardState = "inv-card-powered"

    // inv card unpowered
    InvCardState_inv_card_unpowered InvCardState = "inv-card-unpowered"

    // inv card mdr
    InvCardState_inv_card_mdr InvCardState = "inv-card-mdr"

    // inv card mdr running mbi
    InvCardState_inv_card_mdr_running_mbi InvCardState = "inv-card-mdr-running-mbi"

    // inv card main t mode
    InvCardState_inv_card_main_t_mode InvCardState = "inv-card-main-t-mode"

    // inv card admin down
    InvCardState_inv_card_admin_down InvCardState = "inv-card-admin-down"

    // inv card no mon
    InvCardState_inv_card_no_mon InvCardState = "inv-card-no-mon"

    // inv card unknown
    InvCardState_inv_card_unknown InvCardState = "inv-card-unknown"

    // inv card failed
    InvCardState_inv_card_failed InvCardState = "inv-card-failed"

    // inv card ok
    InvCardState_inv_card_ok InvCardState = "inv-card-ok"

    // inv card missing
    InvCardState_inv_card_missing InvCardState = "inv-card-missing"

    // inv card field diag downloading
    InvCardState_inv_card_field_diag_downloading InvCardState = "inv-card-field-diag-downloading"

    // inv card field diag unmonitor
    InvCardState_inv_card_field_diag_unmonitor InvCardState = "inv-card-field-diag-unmonitor"

    // inv card fabric field diag unmonitor
    InvCardState_inv_card_fabric_field_diag_unmonitor InvCardState = "inv-card-fabric-field-diag-unmonitor"

    // inv card field diag rp launching
    InvCardState_inv_card_field_diag_rp_launching InvCardState = "inv-card-field-diag-rp-launching"

    // inv card field diag running
    InvCardState_inv_card_field_diag_running InvCardState = "inv-card-field-diag-running"

    // inv card field diag pass
    InvCardState_inv_card_field_diag_pass InvCardState = "inv-card-field-diag-pass"

    // inv card field diag fail
    InvCardState_inv_card_field_diag_fail InvCardState = "inv-card-field-diag-fail"

    // inv card field diag timeout
    InvCardState_inv_card_field_diag_timeout InvCardState = "inv-card-field-diag-timeout"

    // inv card disabled
    InvCardState_inv_card_disabled InvCardState = "inv-card-disabled"

    // inv card spa booting
    InvCardState_inv_card_spa_booting InvCardState = "inv-card-spa-booting"

    // inv card not allowed online
    InvCardState_inv_card_not_allowed_online InvCardState = "inv-card-not-allowed-online"

    // inv card stopped
    InvCardState_inv_card_stopped InvCardState = "inv-card-stopped"

    // inv card incompatible fw ver
    InvCardState_inv_card_incompatible_fw_ver InvCardState = "inv-card-incompatible-fw-ver"

    // inv card fpd hold
    InvCardState_inv_card_fpd_hold InvCardState = "inv-card-fpd-hold"

    // inv card node prep
    InvCardState_inv_card_node_prep InvCardState = "inv-card-node-prep"

    // inv card updating fpd
    InvCardState_inv_card_updating_fpd InvCardState = "inv-card-updating-fpd"

    // inv card num states
    InvCardState_inv_card_num_states InvCardState = "inv-card-num-states"
)

// InvMonitorState represents Inv monitor state
type InvMonitorState string

const (
    // unmonitored
    InvMonitorState_unmonitored InvMonitorState = "unmonitored"

    // monitored
    InvMonitorState_monitored InvMonitorState = "monitored"
)

// NodeState represents Node state detail
type NodeState string

const (
    // Not present
    NodeState_not_present NodeState = "not-present"

    // Present
    NodeState_present NodeState = "present"

    // Reset
    NodeState_reset NodeState = "reset"

    // Card booting or rommon
    NodeState_rommon NodeState = "rommon"

    // MBI booting
    NodeState_mbi_boot NodeState = "mbi-boot"

    // Running MBI
    NodeState_mbi_run NodeState = "mbi-run"

    // Running ENA
    NodeState_xr_run NodeState = "xr-run"

    // Bringdown
    NodeState_bring_down NodeState = "bring-down"

    // ENA failure
    NodeState_xr_fail NodeState = "xr-fail"

    // Running FDIAG
    NodeState_fdiag_run NodeState = "fdiag-run"

    // FDIAG failure
    NodeState_fdiag_fail NodeState = "fdiag-fail"

    // Powered
    NodeState_power NodeState = "power"

    // Unpowered
    NodeState_unpower NodeState = "unpower"

    // MDR warm reload
    NodeState_mdr_warm_reload NodeState = "mdr-warm-reload"

    // MDR running MBI
    NodeState_mdr_mbi_run NodeState = "mdr-mbi-run"

    // Maintenance mode
    NodeState_maintenance_mode NodeState = "maintenance-mode"

    // Admin down
    NodeState_admin_down NodeState = "admin-down"

    // No MON
    NodeState_not_monitor NodeState = "not-monitor"

    // Unknown
    NodeState_unknown_card NodeState = "unknown-card"

    // Failed
    NodeState_failed NodeState = "failed"

    // OK
    NodeState_ok NodeState = "ok"

    // Missing
    NodeState_missing NodeState = "missing"

    // Field diag downloading
    NodeState_diag_download NodeState = "diag-download"

    // Field diag unmonitor
    NodeState_diag_not_monitor NodeState = "diag-not-monitor"

    // Fabric field diag unmonitor
    NodeState_fabric_diag_not_monitor NodeState = "fabric-diag-not-monitor"

    // Field diag RP launching
    NodeState_diag_rp_launch NodeState = "diag-rp-launch"

    // Field diag running
    NodeState_diag_run NodeState = "diag-run"

    // Field diag pass
    NodeState_diag_pass NodeState = "diag-pass"

    // Field diag fail
    NodeState_diag_fail NodeState = "diag-fail"

    // Field diag timeout
    NodeState_diag_timeout NodeState = "diag-timeout"

    // Disable
    NodeState_disable NodeState = "disable"

    // SPA booting
    NodeState_spa_boot NodeState = "spa-boot"

    // Not allowed online
    NodeState_not_allowed_online NodeState = "not-allowed-online"

    // Stopped
    NodeState_stop NodeState = "stop"

    // Incompatible FW version
    NodeState_incomp_version NodeState = "incomp-version"

    // FPD hold
    NodeState_fpd_hold NodeState = "fpd-hold"

    // XR preparation
    NodeState_xr_preparation NodeState = "xr-preparation"

    // Sync ready state
    NodeState_sync_ready NodeState = "sync-ready"

    // Node isolate state
    NodeState_xr_isolate NodeState = "xr-isolate"

    // Ready
    NodeState_ready NodeState = "ready"

    // Invalid
    NodeState_invalid NodeState = "invalid"

    // Operational
    NodeState_operational NodeState = "operational"

    // Operational lock
    NodeState_operational_lock NodeState = "operational-lock"

    // Going down
    NodeState_going_down NodeState = "going-down"

    // Going offline
    NodeState_going_offline NodeState = "going-offline"

    // Going online
    NodeState_going_online NodeState = "going-online"

    // Offline
    NodeState_offline NodeState = "offline"

    // Up
    NodeState_up NodeState = "up"

    // Down
    NodeState_down NodeState = "down"

    // Max
    NodeState_max NodeState = "max"

    // Unknown
    NodeState_unknown NodeState = "unknown"
)

// CardRedundancyState represents Redundancy state detail
type CardRedundancyState string

const (
    // Active
    CardRedundancyState_active CardRedundancyState = "active"

    // Standby
    CardRedundancyState_standby CardRedundancyState = "standby"
)

// InvPowerAdminState represents Inv power admin state
type InvPowerAdminState string

const (
    // admin power invalid
    InvPowerAdminState_admin_power_invalid InvPowerAdminState = "admin-power-invalid"

    // admin on
    InvPowerAdminState_admin_on InvPowerAdminState = "admin-on"

    // admin off
    InvPowerAdminState_admin_off InvPowerAdminState = "admin-off"
)

// Platform
// Platform information
type Platform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of racks.
    Racks Platform_Racks
}

func (platform *Platform) GetFilter() yfilter.YFilter { return platform.YFilter }

func (platform *Platform) SetFilter(yf yfilter.YFilter) { platform.YFilter = yf }

func (platform *Platform) GetGoName(yname string) string {
    if yname == "racks" { return "Racks" }
    return ""
}

func (platform *Platform) GetSegmentPath() string {
    return "Cisco-IOS-XR-plat-chas-invmgr-oper:platform"
}

func (platform *Platform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &platform.Racks
    }
    return nil
}

func (platform *Platform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &platform.Racks
    return children
}

func (platform *Platform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platform *Platform) GetBundleName() string { return "cisco_ios_xr" }

func (platform *Platform) GetYangName() string { return "platform" }

func (platform *Platform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platform *Platform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platform *Platform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platform *Platform) SetParent(parent types.Entity) { platform.parent = parent }

func (platform *Platform) GetParent() types.Entity { return platform.parent }

func (platform *Platform) GetParentYangName() string { return "Cisco-IOS-XR-plat-chas-invmgr-oper" }

// Platform_Racks
// Table of racks
type Platform_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack name. The type is slice of Platform_Racks_Rack.
    Rack []Platform_Racks_Rack
}

func (racks *Platform_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *Platform_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *Platform_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *Platform_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *Platform_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Platform_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *Platform_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *Platform_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *Platform_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *Platform_Racks) GetYangName() string { return "racks" }

func (racks *Platform_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *Platform_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *Platform_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *Platform_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *Platform_Racks) GetParent() types.Entity { return racks.parent }

func (racks *Platform_Racks) GetParentYangName() string { return "platform" }

// Platform_Racks_Rack
// Rack name
type Platform_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    RackName interface{}

    // Table of slots.
    Slots Platform_Racks_Rack_Slots
}

func (rack *Platform_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *Platform_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *Platform_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack-name" { return "RackName" }
    if yname == "slots" { return "Slots" }
    return ""
}

func (rack *Platform_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack-name='" + fmt.Sprintf("%v", rack.RackName) + "']"
}

func (rack *Platform_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slots" {
        return &rack.Slots
    }
    return nil
}

func (rack *Platform_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slots"] = &rack.Slots
    return children
}

func (rack *Platform_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-name"] = rack.RackName
    return leafs
}

func (rack *Platform_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *Platform_Racks_Rack) GetYangName() string { return "rack" }

func (rack *Platform_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *Platform_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *Platform_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *Platform_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *Platform_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *Platform_Racks_Rack) GetParentYangName() string { return "racks" }

// Platform_Racks_Rack_Slots
// Table of slots
type Platform_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slot name. The type is slice of Platform_Racks_Rack_Slots_Slot.
    Slot []Platform_Racks_Rack_Slots_Slot
}

func (slots *Platform_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *Platform_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *Platform_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *Platform_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *Platform_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Platform_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *Platform_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *Platform_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *Platform_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *Platform_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *Platform_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *Platform_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *Platform_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *Platform_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *Platform_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *Platform_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// Platform_Racks_Rack_Slots_Slot
// Slot name
type Platform_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SlotName interface{}

    // Table of Instances.
    Instances Platform_Racks_Rack_Slots_Slot_Instances

    // VM information.
    Vm Platform_Racks_Rack_Slots_Slot_Vm

    // State information.
    State Platform_Racks_Rack_Slots_Slot_State
}

func (slot *Platform_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *Platform_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *Platform_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot-name" { return "SlotName" }
    if yname == "instances" { return "Instances" }
    if yname == "vm" { return "Vm" }
    if yname == "state" { return "State" }
    return ""
}

func (slot *Platform_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot-name='" + fmt.Sprintf("%v", slot.SlotName) + "']"
}

func (slot *Platform_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instances" {
        return &slot.Instances
    }
    if childYangName == "vm" {
        return &slot.Vm
    }
    if childYangName == "state" {
        return &slot.State
    }
    return nil
}

func (slot *Platform_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instances"] = &slot.Instances
    children["vm"] = &slot.Vm
    children["state"] = &slot.State
    return children
}

func (slot *Platform_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot-name"] = slot.SlotName
    return leafs
}

func (slot *Platform_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *Platform_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *Platform_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *Platform_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *Platform_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *Platform_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *Platform_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *Platform_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// Platform_Racks_Rack_Slots_Slot_Instances
// Table of Instances
type Platform_Racks_Rack_Slots_Slot_Instances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance name. The type is slice of
    // Platform_Racks_Rack_Slots_Slot_Instances_Instance.
    Instance []Platform_Racks_Rack_Slots_Slot_Instances_Instance
}

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetFilter() yfilter.YFilter { return instances.YFilter }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) SetFilter(yf yfilter.YFilter) { instances.YFilter = yf }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    return ""
}

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetSegmentPath() string {
    return "instances"
}

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        for _, c := range instances.Instance {
            if instances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Platform_Racks_Rack_Slots_Slot_Instances_Instance{}
        instances.Instance = append(instances.Instance, child)
        return &instances.Instance[len(instances.Instance)-1]
    }
    return nil
}

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range instances.Instance {
        children[instances.Instance[i].GetSegmentPath()] = &instances.Instance[i]
    }
    return children
}

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetBundleName() string { return "cisco_ios_xr" }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetYangName() string { return "instances" }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) SetParent(parent types.Entity) { instances.parent = parent }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetParent() types.Entity { return instances.parent }

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetParentYangName() string { return "slot" }

// Platform_Racks_Rack_Slots_Slot_Instances_Instance
// Instance name
type Platform_Racks_Rack_Slots_Slot_Instances_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Instance name. The type is string.
    InstanceName interface{}

    // State information.
    State Platform_Racks_Rack_Slots_Slot_Instances_Instance_State
}

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "state" { return "State" }
    return ""
}

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetSegmentPath() string {
    return "instance" + "[instance-name='" + fmt.Sprintf("%v", instance.InstanceName) + "']"
}

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &instance.State
    }
    return nil
}

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &instance.State
    return children
}

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = instance.InstanceName
    return leafs
}

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetYangName() string { return "instance" }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetParent() types.Entity { return instance.parent }

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetParentYangName() string { return "instances" }

// Platform_Racks_Rack_Slots_Slot_Instances_Instance_State
// State information
type Platform_Racks_Rack_Slots_Slot_Instances_Instance_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Card type. The type is string.
    CardType interface{}

    // Redundancy state. The type is CardRedundancyState.
    CardRedundancyState interface{}

    // PLIM. The type is string.
    Plim interface{}

    // State. The type is NodeState.
    State interface{}

    // True if power state is active. The type is bool.
    IsMonitored interface{}

    // True if monitor state is active. The type is bool.
    IsPowered interface{}

    // True if shutdown state is active. The type is bool.
    IsShutdown interface{}

    // Admin state. The type is string.
    AdminState interface{}
}

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetGoName(yname string) string {
    if yname == "card-type" { return "CardType" }
    if yname == "card-redundancy-state" { return "CardRedundancyState" }
    if yname == "plim" { return "Plim" }
    if yname == "state" { return "State" }
    if yname == "is-monitored" { return "IsMonitored" }
    if yname == "is-powered" { return "IsPowered" }
    if yname == "is-shutdown" { return "IsShutdown" }
    if yname == "admin-state" { return "AdminState" }
    return ""
}

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetSegmentPath() string {
    return "state"
}

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-type"] = state.CardType
    leafs["card-redundancy-state"] = state.CardRedundancyState
    leafs["plim"] = state.Plim
    leafs["state"] = state.State
    leafs["is-monitored"] = state.IsMonitored
    leafs["is-powered"] = state.IsPowered
    leafs["is-shutdown"] = state.IsShutdown
    leafs["admin-state"] = state.AdminState
    return leafs
}

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetYangName() string { return "state" }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetParent() types.Entity { return state.parent }

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetParentYangName() string { return "instance" }

// Platform_Racks_Rack_Slots_Slot_Vm
// VM information
type Platform_Racks_Rack_Slots_Slot_Vm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Type. The type is string.
    NodeDescriptiton interface{}

    // Node Redundency Role. The type is string.
    RedRole interface{}

    // Partner Name. The type is string.
    PartnerName interface{}

    // SW status. The type is string.
    SoftwareStatus interface{}

    // Node IP Address. The type is string.
    NodeIp interface{}
}

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetFilter() yfilter.YFilter { return vm.YFilter }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) SetFilter(yf yfilter.YFilter) { vm.YFilter = yf }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetGoName(yname string) string {
    if yname == "node-descriptiton" { return "NodeDescriptiton" }
    if yname == "red-role" { return "RedRole" }
    if yname == "partner-name" { return "PartnerName" }
    if yname == "software-status" { return "SoftwareStatus" }
    if yname == "node-ip" { return "NodeIp" }
    return ""
}

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetSegmentPath() string {
    return "vm"
}

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-descriptiton"] = vm.NodeDescriptiton
    leafs["red-role"] = vm.RedRole
    leafs["partner-name"] = vm.PartnerName
    leafs["software-status"] = vm.SoftwareStatus
    leafs["node-ip"] = vm.NodeIp
    return leafs
}

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetBundleName() string { return "cisco_ios_xr" }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetYangName() string { return "vm" }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) SetParent(parent types.Entity) { vm.parent = parent }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetParent() types.Entity { return vm.parent }

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetParentYangName() string { return "slot" }

// Platform_Racks_Rack_Slots_Slot_State
// State information
type Platform_Racks_Rack_Slots_Slot_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Card type. The type is string.
    CardType interface{}

    // Redundancy state. The type is CardRedundancyState.
    CardRedundancyState interface{}

    // PLIM. The type is string.
    Plim interface{}

    // State. The type is NodeState.
    State interface{}

    // True if power state is active. The type is bool.
    IsMonitored interface{}

    // True if monitor state is active. The type is bool.
    IsPowered interface{}

    // True if shutdown state is active. The type is bool.
    IsShutdown interface{}

    // Admin state. The type is string.
    AdminState interface{}
}

func (state *Platform_Racks_Rack_Slots_Slot_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Platform_Racks_Rack_Slots_Slot_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Platform_Racks_Rack_Slots_Slot_State) GetGoName(yname string) string {
    if yname == "card-type" { return "CardType" }
    if yname == "card-redundancy-state" { return "CardRedundancyState" }
    if yname == "plim" { return "Plim" }
    if yname == "state" { return "State" }
    if yname == "is-monitored" { return "IsMonitored" }
    if yname == "is-powered" { return "IsPowered" }
    if yname == "is-shutdown" { return "IsShutdown" }
    if yname == "admin-state" { return "AdminState" }
    return ""
}

func (state *Platform_Racks_Rack_Slots_Slot_State) GetSegmentPath() string {
    return "state"
}

func (state *Platform_Racks_Rack_Slots_Slot_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Platform_Racks_Rack_Slots_Slot_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Platform_Racks_Rack_Slots_Slot_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-type"] = state.CardType
    leafs["card-redundancy-state"] = state.CardRedundancyState
    leafs["plim"] = state.Plim
    leafs["state"] = state.State
    leafs["is-monitored"] = state.IsMonitored
    leafs["is-powered"] = state.IsPowered
    leafs["is-shutdown"] = state.IsShutdown
    leafs["admin-state"] = state.AdminState
    return leafs
}

func (state *Platform_Racks_Rack_Slots_Slot_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *Platform_Racks_Rack_Slots_Slot_State) GetYangName() string { return "state" }

func (state *Platform_Racks_Rack_Slots_Slot_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *Platform_Racks_Rack_Slots_Slot_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *Platform_Racks_Rack_Slots_Slot_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *Platform_Racks_Rack_Slots_Slot_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Platform_Racks_Rack_Slots_Slot_State) GetParent() types.Entity { return state.parent }

func (state *Platform_Racks_Rack_Slots_Slot_State) GetParentYangName() string { return "slot" }

// PlatformInventory
// platform inventory
type PlatformInventory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of racks.
    Racks PlatformInventory_Racks
}

func (platformInventory *PlatformInventory) GetFilter() yfilter.YFilter { return platformInventory.YFilter }

func (platformInventory *PlatformInventory) SetFilter(yf yfilter.YFilter) { platformInventory.YFilter = yf }

func (platformInventory *PlatformInventory) GetGoName(yname string) string {
    if yname == "racks" { return "Racks" }
    return ""
}

func (platformInventory *PlatformInventory) GetSegmentPath() string {
    return "Cisco-IOS-XR-plat-chas-invmgr-oper:platform-inventory"
}

func (platformInventory *PlatformInventory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &platformInventory.Racks
    }
    return nil
}

func (platformInventory *PlatformInventory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &platformInventory.Racks
    return children
}

func (platformInventory *PlatformInventory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformInventory *PlatformInventory) GetBundleName() string { return "cisco_ios_xr" }

func (platformInventory *PlatformInventory) GetYangName() string { return "platform-inventory" }

func (platformInventory *PlatformInventory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformInventory *PlatformInventory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformInventory *PlatformInventory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformInventory *PlatformInventory) SetParent(parent types.Entity) { platformInventory.parent = parent }

func (platformInventory *PlatformInventory) GetParent() types.Entity { return platformInventory.parent }

func (platformInventory *PlatformInventory) GetParentYangName() string { return "Cisco-IOS-XR-plat-chas-invmgr-oper" }

// PlatformInventory_Racks
// Table of racks
type PlatformInventory_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack name. The type is slice of PlatformInventory_Racks_Rack.
    Rack []PlatformInventory_Racks_Rack
}

func (racks *PlatformInventory_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *PlatformInventory_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *PlatformInventory_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *PlatformInventory_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *PlatformInventory_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *PlatformInventory_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *PlatformInventory_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *PlatformInventory_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *PlatformInventory_Racks) GetYangName() string { return "racks" }

func (racks *PlatformInventory_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *PlatformInventory_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *PlatformInventory_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *PlatformInventory_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *PlatformInventory_Racks) GetParent() types.Entity { return racks.parent }

func (racks *PlatformInventory_Racks) GetParentYangName() string { return "platform-inventory" }

// PlatformInventory_Racks_Rack
// Rack name
type PlatformInventory_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of slots.
    Slots PlatformInventory_Racks_Rack_Slots

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Attributes
}

func (rack *PlatformInventory_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *PlatformInventory_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *PlatformInventory_Racks_Rack) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "slots" { return "Slots" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (rack *PlatformInventory_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[name='" + fmt.Sprintf("%v", rack.Name) + "']"
}

func (rack *PlatformInventory_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slots" {
        return &rack.Slots
    }
    if childYangName == "attributes" {
        return &rack.Attributes
    }
    return nil
}

func (rack *PlatformInventory_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slots"] = &rack.Slots
    children["attributes"] = &rack.Attributes
    return children
}

func (rack *PlatformInventory_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = rack.Name
    return leafs
}

func (rack *PlatformInventory_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *PlatformInventory_Racks_Rack) GetYangName() string { return "rack" }

func (rack *PlatformInventory_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *PlatformInventory_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *PlatformInventory_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *PlatformInventory_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *PlatformInventory_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *PlatformInventory_Racks_Rack) GetParentYangName() string { return "racks" }

// PlatformInventory_Racks_Rack_Slots
// Table of slots
type PlatformInventory_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slot name. The type is slice of PlatformInventory_Racks_Rack_Slots_Slot.
    Slot []PlatformInventory_Racks_Rack_Slots_Slot
}

func (slots *PlatformInventory_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *PlatformInventory_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *PlatformInventory_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *PlatformInventory_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *PlatformInventory_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *PlatformInventory_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *PlatformInventory_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *PlatformInventory_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *PlatformInventory_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *PlatformInventory_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *PlatformInventory_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *PlatformInventory_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *PlatformInventory_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *PlatformInventory_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *PlatformInventory_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// PlatformInventory_Racks_Rack_Slots_Slot
// Slot name
type PlatformInventory_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of cards.
    Cards PlatformInventory_Racks_Rack_Slots_Slot_Cards

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Attributes
}

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "cards" { return "Cards" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[name='" + fmt.Sprintf("%v", slot.Name) + "']"
}

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cards" {
        return &slot.Cards
    }
    if childYangName == "attributes" {
        return &slot.Attributes
    }
    return nil
}

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cards"] = &slot.Cards
    children["attributes"] = &slot.Attributes
    return children
}

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = slot.Name
    return leafs
}

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards
// Table of cards
type PlatformInventory_Racks_Rack_Slots_Slot_Cards struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Card number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card.
    Card []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card
}

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetFilter() yfilter.YFilter { return cards.YFilter }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) SetFilter(yf yfilter.YFilter) { cards.YFilter = yf }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetGoName(yname string) string {
    if yname == "card" { return "Card" }
    return ""
}

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetSegmentPath() string {
    return "cards"
}

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "card" {
        for _, c := range cards.Card {
            if cards.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card{}
        cards.Card = append(cards.Card, child)
        return &cards.Card[len(cards.Card)-1]
    }
    return nil
}

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cards.Card {
        children[cards.Card[i].GetSegmentPath()] = &cards.Card[i]
    }
    return children
}

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetBundleName() string { return "cisco_ios_xr" }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetYangName() string { return "cards" }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) SetParent(parent types.Entity) { cards.parent = parent }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetParent() types.Entity { return cards.parent }

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetParentYangName() string { return "slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card
// Card number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Card name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // HardwareInformationDir.
    HardwareInformation PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation

    // Table of subslots.
    SubSlots PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots

    // Table of port slots.
    PortSlots PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots

    // Table of  HW components .
    HwComponents PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes
}

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetFilter() yfilter.YFilter { return card.YFilter }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) SetFilter(yf yfilter.YFilter) { card.YFilter = yf }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "hardware-information" { return "HardwareInformation" }
    if yname == "sub-slots" { return "SubSlots" }
    if yname == "port-slots" { return "PortSlots" }
    if yname == "hw-components" { return "HwComponents" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetSegmentPath() string {
    return "card" + "[name='" + fmt.Sprintf("%v", card.Name) + "']"
}

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-information" {
        return &card.HardwareInformation
    }
    if childYangName == "sub-slots" {
        return &card.SubSlots
    }
    if childYangName == "port-slots" {
        return &card.PortSlots
    }
    if childYangName == "hw-components" {
        return &card.HwComponents
    }
    if childYangName == "sensors" {
        return &card.Sensors
    }
    if childYangName == "attributes" {
        return &card.Attributes
    }
    return nil
}

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hardware-information"] = &card.HardwareInformation
    children["sub-slots"] = &card.SubSlots
    children["port-slots"] = &card.PortSlots
    children["hw-components"] = &card.HwComponents
    children["sensors"] = &card.Sensors
    children["attributes"] = &card.Attributes
    return children
}

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = card.Name
    return leafs
}

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetBundleName() string { return "cisco_ios_xr" }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetYangName() string { return "card" }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) SetParent(parent types.Entity) { card.parent = parent }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetParent() types.Entity { return card.parent }

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetParentYangName() string { return "cards" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation
// HardwareInformationDir
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ProcesorInformation.
    ProcessorInformation PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation

    // MotherboardInformation.
    MotherboardInformation PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation

    // BootflashInformation.
    BootflashInformation PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation

    // DiskInformation.
    DiskInformation PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation
}

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetFilter() yfilter.YFilter { return hardwareInformation.YFilter }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) SetFilter(yf yfilter.YFilter) { hardwareInformation.YFilter = yf }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetGoName(yname string) string {
    if yname == "processor-information" { return "ProcessorInformation" }
    if yname == "motherboard-information" { return "MotherboardInformation" }
    if yname == "bootflash-information" { return "BootflashInformation" }
    if yname == "disk-information" { return "DiskInformation" }
    return ""
}

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetSegmentPath() string {
    return "hardware-information"
}

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "processor-information" {
        return &hardwareInformation.ProcessorInformation
    }
    if childYangName == "motherboard-information" {
        return &hardwareInformation.MotherboardInformation
    }
    if childYangName == "bootflash-information" {
        return &hardwareInformation.BootflashInformation
    }
    if childYangName == "disk-information" {
        return &hardwareInformation.DiskInformation
    }
    return nil
}

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["processor-information"] = &hardwareInformation.ProcessorInformation
    children["motherboard-information"] = &hardwareInformation.MotherboardInformation
    children["bootflash-information"] = &hardwareInformation.BootflashInformation
    children["disk-information"] = &hardwareInformation.DiskInformation
    return children
}

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetYangName() string { return "hardware-information" }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) SetParent(parent types.Entity) { hardwareInformation.parent = parent }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetParent() types.Entity { return hardwareInformation.parent }

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetParentYangName() string { return "card" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation
// ProcesorInformation
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type e.g. 7457. The type is string with length: 0..255.
    ProcessorType interface{}

    // Speed e.g. 1197Mhz. The type is string with length: 0..255.
    Speed interface{}

    // Revision. e.g 1.1. The type is string with length: 0..255.
    Revision interface{}
}

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetFilter() yfilter.YFilter { return processorInformation.YFilter }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) SetFilter(yf yfilter.YFilter) { processorInformation.YFilter = yf }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetGoName(yname string) string {
    if yname == "processor-type" { return "ProcessorType" }
    if yname == "speed" { return "Speed" }
    if yname == "revision" { return "Revision" }
    return ""
}

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetSegmentPath() string {
    return "processor-information"
}

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processor-type"] = processorInformation.ProcessorType
    leafs["speed"] = processorInformation.Speed
    leafs["revision"] = processorInformation.Revision
    return leafs
}

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetBundleName() string { return "cisco_ios_xr" }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetYangName() string { return "processor-information" }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) SetParent(parent types.Entity) { processorInformation.parent = parent }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetParent() types.Entity { return processorInformation.parent }

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetParentYangName() string { return "hardware-information" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation
// MotherboardInformation
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MainMemorySize interface{}

    // NVRAM size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    NvramSize interface{}

    // ROM information.
    Rom PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom

    // Bootflash information.
    Bootflash PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash

    // Processor information.
    Processor PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor
}

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetFilter() yfilter.YFilter { return motherboardInformation.YFilter }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) SetFilter(yf yfilter.YFilter) { motherboardInformation.YFilter = yf }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetGoName(yname string) string {
    if yname == "main-memory-size" { return "MainMemorySize" }
    if yname == "nvram-size" { return "NvramSize" }
    if yname == "rom" { return "Rom" }
    if yname == "bootflash" { return "Bootflash" }
    if yname == "processor" { return "Processor" }
    return ""
}

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetSegmentPath() string {
    return "motherboard-information"
}

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rom" {
        return &motherboardInformation.Rom
    }
    if childYangName == "bootflash" {
        return &motherboardInformation.Bootflash
    }
    if childYangName == "processor" {
        return &motherboardInformation.Processor
    }
    return nil
}

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rom"] = &motherboardInformation.Rom
    children["bootflash"] = &motherboardInformation.Bootflash
    children["processor"] = &motherboardInformation.Processor
    return children
}

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["main-memory-size"] = motherboardInformation.MainMemorySize
    leafs["nvram-size"] = motherboardInformation.NvramSize
    return leafs
}

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetBundleName() string { return "cisco_ios_xr" }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetYangName() string { return "motherboard-information" }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) SetParent(parent types.Entity) { motherboardInformation.parent = parent }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetParent() types.Entity { return motherboardInformation.parent }

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetParentYangName() string { return "hardware-information" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom
// ROM information
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Image name. The type is string with length: 0..255.
    ImageName interface{}

    // Major version. The type is interface{} with range: 0..4294967295.
    MajorVersion interface{}

    // Minor version. The type is interface{} with range: 0..4294967295.
    MinorVersion interface{}

    // Micro image version. The type is string with length: 0..255.
    MicroImageVersion interface{}

    // Platform specific text. The type is string with length: 0..255.
    PlatformSpecific interface{}

    // Release type. The type is string with length: 0..255.
    ReleaseType interface{}
}

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetFilter() yfilter.YFilter { return rom.YFilter }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) SetFilter(yf yfilter.YFilter) { rom.YFilter = yf }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetGoName(yname string) string {
    if yname == "image-name" { return "ImageName" }
    if yname == "major-version" { return "MajorVersion" }
    if yname == "minor-version" { return "MinorVersion" }
    if yname == "micro-image-version" { return "MicroImageVersion" }
    if yname == "platform-specific" { return "PlatformSpecific" }
    if yname == "release-type" { return "ReleaseType" }
    return ""
}

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetSegmentPath() string {
    return "rom"
}

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["image-name"] = rom.ImageName
    leafs["major-version"] = rom.MajorVersion
    leafs["minor-version"] = rom.MinorVersion
    leafs["micro-image-version"] = rom.MicroImageVersion
    leafs["platform-specific"] = rom.PlatformSpecific
    leafs["release-type"] = rom.ReleaseType
    return leafs
}

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetBundleName() string { return "cisco_ios_xr" }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetYangName() string { return "rom" }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) SetParent(parent types.Entity) { rom.parent = parent }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetParent() types.Entity { return rom.parent }

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetParentYangName() string { return "motherboard-information" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash
// Bootflash information
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Image name. The type is string with length: 0..255.
    ImageName interface{}

    // Platform Type. The type is string with length: 0..255.
    PlatformType interface{}

    // Major version. The type is interface{} with range: 0..4294967295.
    MajorVersion interface{}

    // Minor version. The type is interface{} with range: 0..4294967295.
    MinorVersion interface{}

    // Micro image version. The type is string with length: 0..255.
    MicroImageVersion interface{}

    // Platform specific text. The type is string with length: 0..255.
    PlatformSpecific interface{}

    // Release type. The type is string with length: 0..255.
    ReleaseType interface{}

    // Bootflash type e.g. SIMM. The type is string with length: 0..255.
    BootflashType interface{}

    // Bootflash size in kilo-bytes. The type is interface{} with range:
    // 0..4294967295. Units are kilobyte.
    BootflashSize interface{}

    // Sector size in bytes. The type is interface{} with range: 0..4294967295.
    // Units are byte.
    SectorSize interface{}
}

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetFilter() yfilter.YFilter { return bootflash.YFilter }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) SetFilter(yf yfilter.YFilter) { bootflash.YFilter = yf }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetGoName(yname string) string {
    if yname == "image-name" { return "ImageName" }
    if yname == "platform-type" { return "PlatformType" }
    if yname == "major-version" { return "MajorVersion" }
    if yname == "minor-version" { return "MinorVersion" }
    if yname == "micro-image-version" { return "MicroImageVersion" }
    if yname == "platform-specific" { return "PlatformSpecific" }
    if yname == "release-type" { return "ReleaseType" }
    if yname == "bootflash-type" { return "BootflashType" }
    if yname == "bootflash-size" { return "BootflashSize" }
    if yname == "sector-size" { return "SectorSize" }
    return ""
}

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetSegmentPath() string {
    return "bootflash"
}

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["image-name"] = bootflash.ImageName
    leafs["platform-type"] = bootflash.PlatformType
    leafs["major-version"] = bootflash.MajorVersion
    leafs["minor-version"] = bootflash.MinorVersion
    leafs["micro-image-version"] = bootflash.MicroImageVersion
    leafs["platform-specific"] = bootflash.PlatformSpecific
    leafs["release-type"] = bootflash.ReleaseType
    leafs["bootflash-type"] = bootflash.BootflashType
    leafs["bootflash-size"] = bootflash.BootflashSize
    leafs["sector-size"] = bootflash.SectorSize
    return leafs
}

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetBundleName() string { return "cisco_ios_xr" }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetYangName() string { return "bootflash" }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) SetParent(parent types.Entity) { bootflash.parent = parent }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetParent() types.Entity { return bootflash.parent }

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetParentYangName() string { return "motherboard-information" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor
// Processor information
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type e.g. 7457. The type is string with length: 0..255.
    ProcessorType interface{}

    // Speed e.g. 1197Mhz. The type is string with length: 0..255.
    Speed interface{}

    // Revision. e.g 1.1. The type is string with length: 0..255.
    Revision interface{}
}

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetFilter() yfilter.YFilter { return processor.YFilter }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) SetFilter(yf yfilter.YFilter) { processor.YFilter = yf }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetGoName(yname string) string {
    if yname == "processor-type" { return "ProcessorType" }
    if yname == "speed" { return "Speed" }
    if yname == "revision" { return "Revision" }
    return ""
}

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetSegmentPath() string {
    return "processor"
}

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processor-type"] = processor.ProcessorType
    leafs["speed"] = processor.Speed
    leafs["revision"] = processor.Revision
    return leafs
}

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetBundleName() string { return "cisco_ios_xr" }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetYangName() string { return "processor" }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) SetParent(parent types.Entity) { processor.parent = parent }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetParent() types.Entity { return processor.parent }

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetParentYangName() string { return "motherboard-information" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation
// BootflashInformation
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Image name. The type is string with length: 0..255.
    ImageName interface{}

    // Platform Type. The type is string with length: 0..255.
    PlatformType interface{}

    // Major version. The type is interface{} with range: 0..4294967295.
    MajorVersion interface{}

    // Minor version. The type is interface{} with range: 0..4294967295.
    MinorVersion interface{}

    // Micro image version. The type is string with length: 0..255.
    MicroImageVersion interface{}

    // Platform specific text. The type is string with length: 0..255.
    PlatformSpecific interface{}

    // Release type. The type is string with length: 0..255.
    ReleaseType interface{}

    // Bootflash type e.g. SIMM. The type is string with length: 0..255.
    BootflashType interface{}

    // Bootflash size in kilo-bytes. The type is interface{} with range:
    // 0..4294967295. Units are kilobyte.
    BootflashSize interface{}

    // Sector size in bytes. The type is interface{} with range: 0..4294967295.
    // Units are byte.
    SectorSize interface{}
}

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetFilter() yfilter.YFilter { return bootflashInformation.YFilter }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) SetFilter(yf yfilter.YFilter) { bootflashInformation.YFilter = yf }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetGoName(yname string) string {
    if yname == "image-name" { return "ImageName" }
    if yname == "platform-type" { return "PlatformType" }
    if yname == "major-version" { return "MajorVersion" }
    if yname == "minor-version" { return "MinorVersion" }
    if yname == "micro-image-version" { return "MicroImageVersion" }
    if yname == "platform-specific" { return "PlatformSpecific" }
    if yname == "release-type" { return "ReleaseType" }
    if yname == "bootflash-type" { return "BootflashType" }
    if yname == "bootflash-size" { return "BootflashSize" }
    if yname == "sector-size" { return "SectorSize" }
    return ""
}

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetSegmentPath() string {
    return "bootflash-information"
}

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["image-name"] = bootflashInformation.ImageName
    leafs["platform-type"] = bootflashInformation.PlatformType
    leafs["major-version"] = bootflashInformation.MajorVersion
    leafs["minor-version"] = bootflashInformation.MinorVersion
    leafs["micro-image-version"] = bootflashInformation.MicroImageVersion
    leafs["platform-specific"] = bootflashInformation.PlatformSpecific
    leafs["release-type"] = bootflashInformation.ReleaseType
    leafs["bootflash-type"] = bootflashInformation.BootflashType
    leafs["bootflash-size"] = bootflashInformation.BootflashSize
    leafs["sector-size"] = bootflashInformation.SectorSize
    return leafs
}

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetBundleName() string { return "cisco_ios_xr" }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetYangName() string { return "bootflash-information" }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) SetParent(parent types.Entity) { bootflashInformation.parent = parent }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetParent() types.Entity { return bootflashInformation.parent }

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetParentYangName() string { return "hardware-information" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation
// DiskInformation
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // (Deprecated) Disk name. The type is string with length: 0..255.
    DiskName interface{}

    // (Deprecated) Disk size in mega-bytes. The type is interface{} with range:
    // 0..4294967295. Units are megabyte.
    DiskSize interface{}

    // (Deprecated) Disk sector size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    SectorSize interface{}

    // Disk attributes. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks.
    Disks []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks
}

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetFilter() yfilter.YFilter { return diskInformation.YFilter }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) SetFilter(yf yfilter.YFilter) { diskInformation.YFilter = yf }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetGoName(yname string) string {
    if yname == "disk-name" { return "DiskName" }
    if yname == "disk-size" { return "DiskSize" }
    if yname == "sector-size" { return "SectorSize" }
    if yname == "disks" { return "Disks" }
    return ""
}

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetSegmentPath() string {
    return "disk-information"
}

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "disks" {
        for _, c := range diskInformation.Disks {
            if diskInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks{}
        diskInformation.Disks = append(diskInformation.Disks, child)
        return &diskInformation.Disks[len(diskInformation.Disks)-1]
    }
    return nil
}

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diskInformation.Disks {
        children[diskInformation.Disks[i].GetSegmentPath()] = &diskInformation.Disks[i]
    }
    return children
}

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disk-name"] = diskInformation.DiskName
    leafs["disk-size"] = diskInformation.DiskSize
    leafs["sector-size"] = diskInformation.SectorSize
    return leafs
}

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetBundleName() string { return "cisco_ios_xr" }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetYangName() string { return "disk-information" }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) SetParent(parent types.Entity) { diskInformation.parent = parent }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetParent() types.Entity { return diskInformation.parent }

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetParentYangName() string { return "hardware-information" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks
// Disk attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disk name. The type is string with length: 0..255.
    DiskName interface{}

    // Disk size in mega-bytes. The type is interface{} with range: 0..4294967295.
    // Units are megabyte.
    DiskSize interface{}

    // Disk sector size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    SectorSize interface{}
}

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetFilter() yfilter.YFilter { return disks.YFilter }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) SetFilter(yf yfilter.YFilter) { disks.YFilter = yf }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetGoName(yname string) string {
    if yname == "disk-name" { return "DiskName" }
    if yname == "disk-size" { return "DiskSize" }
    if yname == "sector-size" { return "SectorSize" }
    return ""
}

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetSegmentPath() string {
    return "disks"
}

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disk-name"] = disks.DiskName
    leafs["disk-size"] = disks.DiskSize
    leafs["sector-size"] = disks.SectorSize
    return leafs
}

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetBundleName() string { return "cisco_ios_xr" }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetYangName() string { return "disks" }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) SetParent(parent types.Entity) { disks.parent = parent }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetParent() types.Entity { return disks.parent }

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetParentYangName() string { return "disk-information" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots
// Table of subslots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subslot number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot.
    SubSlot []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
}

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetFilter() yfilter.YFilter { return subSlots.YFilter }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) SetFilter(yf yfilter.YFilter) { subSlots.YFilter = yf }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetGoName(yname string) string {
    if yname == "sub-slot" { return "SubSlot" }
    return ""
}

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetSegmentPath() string {
    return "sub-slots"
}

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sub-slot" {
        for _, c := range subSlots.SubSlot {
            if subSlots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot{}
        subSlots.SubSlot = append(subSlots.SubSlot, child)
        return &subSlots.SubSlot[len(subSlots.SubSlot)-1]
    }
    return nil
}

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subSlots.SubSlot {
        children[subSlots.SubSlot[i].GetSegmentPath()] = &subSlots.SubSlot[i]
    }
    return children
}

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetBundleName() string { return "cisco_ios_xr" }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetYangName() string { return "sub-slots" }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) SetParent(parent types.Entity) { subSlots.parent = parent }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetParent() types.Entity { return subSlots.parent }

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetParentYangName() string { return "card" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
// Subslot number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Subslot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Module of a subslot.
    Module PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes
}

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetFilter() yfilter.YFilter { return subSlot.YFilter }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) SetFilter(yf yfilter.YFilter) { subSlot.YFilter = yf }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "module" { return "Module" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetSegmentPath() string {
    return "sub-slot" + "[name='" + fmt.Sprintf("%v", subSlot.Name) + "']"
}

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "module" {
        return &subSlot.Module
    }
    if childYangName == "attributes" {
        return &subSlot.Attributes
    }
    return nil
}

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["module"] = &subSlot.Module
    children["attributes"] = &subSlot.Attributes
    return children
}

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = subSlot.Name
    return leafs
}

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetBundleName() string { return "cisco_ios_xr" }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetYangName() string { return "sub-slot" }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) SetParent(parent types.Entity) { subSlot.parent = parent }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetParent() types.Entity { return subSlot.parent }

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetParentYangName() string { return "sub-slots" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module
// Module of a subslot
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of port slots.
    PortSlots PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes
}

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetFilter() yfilter.YFilter { return module.YFilter }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) SetFilter(yf yfilter.YFilter) { module.YFilter = yf }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetGoName(yname string) string {
    if yname == "port-slots" { return "PortSlots" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetSegmentPath() string {
    return "module"
}

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-slots" {
        return &module.PortSlots
    }
    if childYangName == "sensors" {
        return &module.Sensors
    }
    if childYangName == "attributes" {
        return &module.Attributes
    }
    return nil
}

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-slots"] = &module.PortSlots
    children["sensors"] = &module.Sensors
    children["attributes"] = &module.Attributes
    return children
}

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetBundleName() string { return "cisco_ios_xr" }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetYangName() string { return "module" }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) SetParent(parent types.Entity) { module.parent = parent }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetParent() types.Entity { return module.parent }

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetParentYangName() string { return "sub-slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots
// Table of port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port slot number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot.
    PortSlot []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetFilter() yfilter.YFilter { return portSlots.YFilter }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) SetFilter(yf yfilter.YFilter) { portSlots.YFilter = yf }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetGoName(yname string) string {
    if yname == "port-slot" { return "PortSlot" }
    return ""
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetSegmentPath() string {
    return "port-slots"
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-slot" {
        for _, c := range portSlots.PortSlot {
            if portSlots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot{}
        portSlots.PortSlot = append(portSlots.PortSlot, child)
        return &portSlots.PortSlot[len(portSlots.PortSlot)-1]
    }
    return nil
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portSlots.PortSlot {
        children[portSlots.PortSlot[i].GetSegmentPath()] = &portSlots.PortSlot[i]
    }
    return children
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetBundleName() string { return "cisco_ios_xr" }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetYangName() string { return "port-slots" }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) SetParent(parent types.Entity) { portSlots.parent = parent }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetParent() types.Entity { return portSlots.parent }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetParentYangName() string { return "module" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
// Port slot number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of spirit port slots.
    Portses PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetFilter() yfilter.YFilter { return portSlot.YFilter }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) SetFilter(yf yfilter.YFilter) { portSlot.YFilter = yf }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "portses" { return "Portses" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetSegmentPath() string {
    return "port-slot" + "[name='" + fmt.Sprintf("%v", portSlot.Name) + "']"
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "portses" {
        return &portSlot.Portses
    }
    if childYangName == "sensors" {
        return &portSlot.Sensors
    }
    if childYangName == "attributes" {
        return &portSlot.Attributes
    }
    return nil
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["portses"] = &portSlot.Portses
    children["sensors"] = &portSlot.Sensors
    children["attributes"] = &portSlot.Attributes
    return children
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = portSlot.Name
    return leafs
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetBundleName() string { return "cisco_ios_xr" }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetYangName() string { return "port-slot" }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) SetParent(parent types.Entity) { portSlot.parent = parent }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetParent() types.Entity { return portSlot.parent }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetParentYangName() string { return "port-slots" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses
// Table of spirit port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports.
    Ports []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetFilter() yfilter.YFilter { return portses.YFilter }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) SetFilter(yf yfilter.YFilter) { portses.YFilter = yf }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetGoName(yname string) string {
    if yname == "ports" { return "Ports" }
    return ""
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetSegmentPath() string {
    return "portses"
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ports" {
        for _, c := range portses.Ports {
            if portses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports{}
        portses.Ports = append(portses.Ports, child)
        return &portses.Ports[len(portses.Ports)-1]
    }
    return nil
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portses.Ports {
        children[portses.Ports[i].GetSegmentPath()] = &portses.Ports[i]
    }
    return children
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetBundleName() string { return "cisco_ios_xr" }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetYangName() string { return "portses" }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) SetParent(parent types.Entity) { portses.parent = parent }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetParent() types.Entity { return portses.parent }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetParentYangName() string { return "port-slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports
// Port number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of  HW components .
    HwComponents PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetFilter() yfilter.YFilter { return ports.YFilter }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) SetFilter(yf yfilter.YFilter) { ports.YFilter = yf }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "hw-components" { return "HwComponents" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetSegmentPath() string {
    return "ports" + "[name='" + fmt.Sprintf("%v", ports.Name) + "']"
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-components" {
        return &ports.HwComponents
    }
    if childYangName == "sensors" {
        return &ports.Sensors
    }
    if childYangName == "attributes" {
        return &ports.Attributes
    }
    return nil
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-components"] = &ports.HwComponents
    children["sensors"] = &ports.Sensors
    children["attributes"] = &ports.Attributes
    return children
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = ports.Name
    return leafs
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetBundleName() string { return "cisco_ios_xr" }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetYangName() string { return "ports" }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) SetParent(parent types.Entity) { ports.parent = parent }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetParent() types.Entity { return ports.parent }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetParentYangName() string { return "portses" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents
// Table of  HW components 
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // HW component number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent.
    HwComponent []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetFilter() yfilter.YFilter { return hwComponents.YFilter }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) SetFilter(yf yfilter.YFilter) { hwComponents.YFilter = yf }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetGoName(yname string) string {
    if yname == "hw-component" { return "HwComponent" }
    return ""
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetSegmentPath() string {
    return "hw-components"
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-component" {
        for _, c := range hwComponents.HwComponent {
            if hwComponents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent{}
        hwComponents.HwComponent = append(hwComponents.HwComponent, child)
        return &hwComponents.HwComponent[len(hwComponents.HwComponent)-1]
    }
    return nil
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwComponents.HwComponent {
        children[hwComponents.HwComponent[i].GetSegmentPath()] = &hwComponents.HwComponent[i]
    }
    return children
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetBundleName() string { return "cisco_ios_xr" }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetYangName() string { return "hw-components" }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) SetParent(parent types.Entity) { hwComponents.parent = parent }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetParent() types.Entity { return hwComponents.parent }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetParentYangName() string { return "ports" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent
// HW component number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HW component name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetFilter() yfilter.YFilter { return hwComponent.YFilter }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) SetFilter(yf yfilter.YFilter) { hwComponent.YFilter = yf }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetSegmentPath() string {
    return "hw-component" + "[name='" + fmt.Sprintf("%v", hwComponent.Name) + "']"
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensors" {
        return &hwComponent.Sensors
    }
    if childYangName == "attributes" {
        return &hwComponent.Attributes
    }
    return nil
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensors"] = &hwComponent.Sensors
    children["attributes"] = &hwComponent.Attributes
    return children
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = hwComponent.Name
    return leafs
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetBundleName() string { return "cisco_ios_xr" }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetYangName() string { return "hw-component" }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) SetParent(parent types.Entity) { hwComponent.parent = parent }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetParent() types.Entity { return hwComponent.parent }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetParentYangName() string { return "hw-components" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetParentYangName() string { return "hw-component" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetParentYangName() string { return "hw-component" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetParentYangName() string { return "ports" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetParentYangName() string { return "ports" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetParentYangName() string { return "port-slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetParentYangName() string { return "port-slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetParentYangName() string { return "module" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetParentYangName() string { return "module" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetParentYangName() string { return "sub-slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots
// Table of port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port slot number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot.
    PortSlot []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetFilter() yfilter.YFilter { return portSlots.YFilter }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) SetFilter(yf yfilter.YFilter) { portSlots.YFilter = yf }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetGoName(yname string) string {
    if yname == "port-slot" { return "PortSlot" }
    return ""
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetSegmentPath() string {
    return "port-slots"
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-slot" {
        for _, c := range portSlots.PortSlot {
            if portSlots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot{}
        portSlots.PortSlot = append(portSlots.PortSlot, child)
        return &portSlots.PortSlot[len(portSlots.PortSlot)-1]
    }
    return nil
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portSlots.PortSlot {
        children[portSlots.PortSlot[i].GetSegmentPath()] = &portSlots.PortSlot[i]
    }
    return children
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetBundleName() string { return "cisco_ios_xr" }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetYangName() string { return "port-slots" }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) SetParent(parent types.Entity) { portSlots.parent = parent }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetParent() types.Entity { return portSlots.parent }

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetParentYangName() string { return "card" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
// Port slot number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of spirit port slots.
    Portses PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetFilter() yfilter.YFilter { return portSlot.YFilter }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) SetFilter(yf yfilter.YFilter) { portSlot.YFilter = yf }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "portses" { return "Portses" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetSegmentPath() string {
    return "port-slot" + "[name='" + fmt.Sprintf("%v", portSlot.Name) + "']"
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "portses" {
        return &portSlot.Portses
    }
    if childYangName == "sensors" {
        return &portSlot.Sensors
    }
    if childYangName == "attributes" {
        return &portSlot.Attributes
    }
    return nil
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["portses"] = &portSlot.Portses
    children["sensors"] = &portSlot.Sensors
    children["attributes"] = &portSlot.Attributes
    return children
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = portSlot.Name
    return leafs
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetBundleName() string { return "cisco_ios_xr" }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetYangName() string { return "port-slot" }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) SetParent(parent types.Entity) { portSlot.parent = parent }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetParent() types.Entity { return portSlot.parent }

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetParentYangName() string { return "port-slots" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses
// Table of spirit port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports.
    Ports []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetFilter() yfilter.YFilter { return portses.YFilter }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) SetFilter(yf yfilter.YFilter) { portses.YFilter = yf }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetGoName(yname string) string {
    if yname == "ports" { return "Ports" }
    return ""
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetSegmentPath() string {
    return "portses"
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ports" {
        for _, c := range portses.Ports {
            if portses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports{}
        portses.Ports = append(portses.Ports, child)
        return &portses.Ports[len(portses.Ports)-1]
    }
    return nil
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portses.Ports {
        children[portses.Ports[i].GetSegmentPath()] = &portses.Ports[i]
    }
    return children
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetBundleName() string { return "cisco_ios_xr" }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetYangName() string { return "portses" }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) SetParent(parent types.Entity) { portses.parent = parent }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetParent() types.Entity { return portses.parent }

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetParentYangName() string { return "port-slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports
// Port number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of  HW components .
    HwComponents PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetFilter() yfilter.YFilter { return ports.YFilter }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) SetFilter(yf yfilter.YFilter) { ports.YFilter = yf }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "hw-components" { return "HwComponents" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetSegmentPath() string {
    return "ports" + "[name='" + fmt.Sprintf("%v", ports.Name) + "']"
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-components" {
        return &ports.HwComponents
    }
    if childYangName == "sensors" {
        return &ports.Sensors
    }
    if childYangName == "attributes" {
        return &ports.Attributes
    }
    return nil
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-components"] = &ports.HwComponents
    children["sensors"] = &ports.Sensors
    children["attributes"] = &ports.Attributes
    return children
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = ports.Name
    return leafs
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetBundleName() string { return "cisco_ios_xr" }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetYangName() string { return "ports" }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) SetParent(parent types.Entity) { ports.parent = parent }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetParent() types.Entity { return ports.parent }

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetParentYangName() string { return "portses" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents
// Table of  HW components 
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // HW component number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent.
    HwComponent []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetFilter() yfilter.YFilter { return hwComponents.YFilter }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) SetFilter(yf yfilter.YFilter) { hwComponents.YFilter = yf }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetGoName(yname string) string {
    if yname == "hw-component" { return "HwComponent" }
    return ""
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetSegmentPath() string {
    return "hw-components"
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-component" {
        for _, c := range hwComponents.HwComponent {
            if hwComponents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent{}
        hwComponents.HwComponent = append(hwComponents.HwComponent, child)
        return &hwComponents.HwComponent[len(hwComponents.HwComponent)-1]
    }
    return nil
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwComponents.HwComponent {
        children[hwComponents.HwComponent[i].GetSegmentPath()] = &hwComponents.HwComponent[i]
    }
    return children
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetBundleName() string { return "cisco_ios_xr" }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetYangName() string { return "hw-components" }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) SetParent(parent types.Entity) { hwComponents.parent = parent }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetParent() types.Entity { return hwComponents.parent }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetParentYangName() string { return "ports" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent
// HW component number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HW component name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetFilter() yfilter.YFilter { return hwComponent.YFilter }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) SetFilter(yf yfilter.YFilter) { hwComponent.YFilter = yf }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetSegmentPath() string {
    return "hw-component" + "[name='" + fmt.Sprintf("%v", hwComponent.Name) + "']"
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensors" {
        return &hwComponent.Sensors
    }
    if childYangName == "attributes" {
        return &hwComponent.Attributes
    }
    return nil
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensors"] = &hwComponent.Sensors
    children["attributes"] = &hwComponent.Attributes
    return children
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = hwComponent.Name
    return leafs
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetBundleName() string { return "cisco_ios_xr" }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetYangName() string { return "hw-component" }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) SetParent(parent types.Entity) { hwComponent.parent = parent }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetParent() types.Entity { return hwComponent.parent }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetParentYangName() string { return "hw-components" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetParentYangName() string { return "hw-component" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetParentYangName() string { return "hw-component" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetParentYangName() string { return "ports" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetParentYangName() string { return "ports" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetParentYangName() string { return "port-slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetParentYangName() string { return "port-slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents
// Table of  HW components 
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // HW component number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent.
    HwComponent []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetFilter() yfilter.YFilter { return hwComponents.YFilter }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) SetFilter(yf yfilter.YFilter) { hwComponents.YFilter = yf }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetGoName(yname string) string {
    if yname == "hw-component" { return "HwComponent" }
    return ""
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetSegmentPath() string {
    return "hw-components"
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-component" {
        for _, c := range hwComponents.HwComponent {
            if hwComponents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent{}
        hwComponents.HwComponent = append(hwComponents.HwComponent, child)
        return &hwComponents.HwComponent[len(hwComponents.HwComponent)-1]
    }
    return nil
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwComponents.HwComponent {
        children[hwComponents.HwComponent[i].GetSegmentPath()] = &hwComponents.HwComponent[i]
    }
    return children
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetBundleName() string { return "cisco_ios_xr" }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetYangName() string { return "hw-components" }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) SetParent(parent types.Entity) { hwComponents.parent = parent }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetParent() types.Entity { return hwComponents.parent }

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetParentYangName() string { return "card" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
// HW component number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HW component name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetFilter() yfilter.YFilter { return hwComponent.YFilter }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) SetFilter(yf yfilter.YFilter) { hwComponent.YFilter = yf }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "sensors" { return "Sensors" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetSegmentPath() string {
    return "hw-component" + "[name='" + fmt.Sprintf("%v", hwComponent.Name) + "']"
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensors" {
        return &hwComponent.Sensors
    }
    if childYangName == "attributes" {
        return &hwComponent.Attributes
    }
    return nil
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sensors"] = &hwComponent.Sensors
    children["attributes"] = &hwComponent.Attributes
    return children
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = hwComponent.Name
    return leafs
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetBundleName() string { return "cisco_ios_xr" }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetYangName() string { return "hw-component" }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) SetParent(parent types.Entity) { hwComponent.parent = parent }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetParent() types.Entity { return hwComponent.parent }

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetParentYangName() string { return "hw-components" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetParentYangName() string { return "hw-component" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetParentYangName() string { return "hw-component" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor.
    Sensor []PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetFilter() yfilter.YFilter { return sensors.YFilter }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) SetFilter(yf yfilter.YFilter) { sensors.YFilter = yf }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetGoName(yname string) string {
    if yname == "sensor" { return "Sensor" }
    return ""
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetSegmentPath() string {
    return "sensors"
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sensor" {
        for _, c := range sensors.Sensor {
            if sensors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor{}
        sensors.Sensor = append(sensors.Sensor, child)
        return &sensors.Sensor[len(sensors.Sensor)-1]
    }
    return nil
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sensors.Sensor {
        children[sensors.Sensor[i].GetSegmentPath()] = &sensors.Sensor[i]
    }
    return children
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetBundleName() string { return "cisco_ios_xr" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetYangName() string { return "sensors" }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) SetParent(parent types.Entity) { sensors.parent = parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetParent() types.Entity { return sensors.parent }

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetParentYangName() string { return "card" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "attributes" { return "Attributes" }
    return ""
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetSegmentPath() string {
    return "sensor" + "[name='" + fmt.Sprintf("%v", sensor.Name) + "']"
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &sensor.Attributes
    }
    return nil
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &sensor.Attributes
    return children
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sensor.Name
    return leafs
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetYangName() string { return "sensor" }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetParentYangName() string { return "sensors" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetParentYangName() string { return "sensor" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetParentYangName() string { return "card" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetParentYangName() string { return "slot" }

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *PlatformInventory_Racks_Rack_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetGoName(yname string) string {
    if yname == "basic-info" { return "BasicInfo" }
    if yname == "fru-info" { return "FruInfo" }
    return ""
}

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-info" {
        return &attributes.BasicInfo
    }
    if childYangName == "fru-info" {
        return &attributes.FruInfo
    }
    return nil
}

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-info"] = &attributes.BasicInfo
    children["fru-info"] = &attributes.FruInfo
    return children
}

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetBundleName() string { return "cisco_ios_xr" }

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetYangName() string { return "attributes" }

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attributes *PlatformInventory_Racks_Rack_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetParentYangName() string { return "rack" }

// PlatformInventory_Racks_Rack_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Attributes_BasicInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name string for the entity. The type is string with length: 0..255.
    Name interface{}

    // describes in user-readable terms                 what the entity in
    // question does. The type is string with length: 0..255.
    Description interface{}

    // model name. The type is string with length: 0..255.
    ModelName interface{}

    // hw revision string. The type is string with length: 0..255.
    HardwareRevision interface{}

    // serial number. The type is string with length: 0..255.
    SerialNumber interface{}

    // firmware revision string. The type is string with length: 0..255.
    FirmwareRevision interface{}

    // software revision string. The type is string with length: 0..255.
    SoftwareRevision interface{}

    // maps to the vendor OID string. The type is string with length: 0..255.
    VendorType interface{}

    // 1 if Field Replaceable Unit 0, if not. The type is bool.
    IsFieldReplaceableUnit interface{}
}

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetFilter() yfilter.YFilter { return basicInfo.YFilter }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) SetFilter(yf yfilter.YFilter) { basicInfo.YFilter = yf }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "model-name" { return "ModelName" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "software-revision" { return "SoftwareRevision" }
    if yname == "vendor-type" { return "VendorType" }
    if yname == "is-field-replaceable-unit" { return "IsFieldReplaceableUnit" }
    return ""
}

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetSegmentPath() string {
    return "basic-info"
}

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = basicInfo.Name
    leafs["description"] = basicInfo.Description
    leafs["model-name"] = basicInfo.ModelName
    leafs["hardware-revision"] = basicInfo.HardwareRevision
    leafs["serial-number"] = basicInfo.SerialNumber
    leafs["firmware-revision"] = basicInfo.FirmwareRevision
    leafs["software-revision"] = basicInfo.SoftwareRevision
    leafs["vendor-type"] = basicInfo.VendorType
    leafs["is-field-replaceable-unit"] = basicInfo.IsFieldReplaceableUnit
    return leafs
}

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetBundleName() string { return "cisco_ios_xr" }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetYangName() string { return "basic-info" }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) SetParent(parent types.Entity) { basicInfo.parent = parent }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetParent() types.Entity { return basicInfo.parent }

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Attributes_FruInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative    state. The type is InvAdminState.
    ModuleAdministrativeState interface{}

    // Power administrative state. The type is InvPowerAdminState.
    ModulePowerAdministrativeState interface{}

    // Operation state. The type is InvCardState.
    ModuleOperationalState interface{}

    // Monitor state. The type is InvMonitorState.
    ModuleMonitorState interface{}

    // Reset reason. The type is InvResetReason.
    ModuleResetReason interface{}

    // Time operational state is   last changed.
    LastOperationalStateChange PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetFilter() yfilter.YFilter { return fruInfo.YFilter }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) SetFilter(yf yfilter.YFilter) { fruInfo.YFilter = yf }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetGoName(yname string) string {
    if yname == "module-administrative-state" { return "ModuleAdministrativeState" }
    if yname == "module-power-administrative-state" { return "ModulePowerAdministrativeState" }
    if yname == "module-operational-state" { return "ModuleOperationalState" }
    if yname == "module-monitor-state" { return "ModuleMonitorState" }
    if yname == "module-reset-reason" { return "ModuleResetReason" }
    if yname == "last-operational-state-change" { return "LastOperationalStateChange" }
    if yname == "module-up-time" { return "ModuleUpTime" }
    return ""
}

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetSegmentPath() string {
    return "fru-info"
}

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-operational-state-change" {
        return &fruInfo.LastOperationalStateChange
    }
    if childYangName == "module-up-time" {
        return &fruInfo.ModuleUpTime
    }
    return nil
}

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-operational-state-change"] = &fruInfo.LastOperationalStateChange
    children["module-up-time"] = &fruInfo.ModuleUpTime
    return children
}

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-administrative-state"] = fruInfo.ModuleAdministrativeState
    leafs["module-power-administrative-state"] = fruInfo.ModulePowerAdministrativeState
    leafs["module-operational-state"] = fruInfo.ModuleOperationalState
    leafs["module-monitor-state"] = fruInfo.ModuleMonitorState
    leafs["module-reset-reason"] = fruInfo.ModuleResetReason
    return leafs
}

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetYangName() string { return "fru-info" }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) SetParent(parent types.Entity) { fruInfo.parent = parent }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetParent() types.Entity { return fruInfo.parent }

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetParentYangName() string { return "attributes" }

// PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetFilter() yfilter.YFilter { return lastOperationalStateChange.YFilter }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) SetFilter(yf yfilter.YFilter) { lastOperationalStateChange.YFilter = yf }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetSegmentPath() string {
    return "last-operational-state-change"
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = lastOperationalStateChange.TimeInSeconds
    leafs["time-in-nano-seconds"] = lastOperationalStateChange.TimeInNanoSeconds
    return leafs
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetBundleName() string { return "cisco_ios_xr" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetYangName() string { return "last-operational-state-change" }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) SetParent(parent types.Entity) { lastOperationalStateChange.parent = parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetParent() types.Entity { return lastOperationalStateChange.parent }

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetParentYangName() string { return "fru-info" }

// PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetFilter() yfilter.YFilter { return moduleUpTime.YFilter }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) SetFilter(yf yfilter.YFilter) { moduleUpTime.YFilter = yf }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetGoName(yname string) string {
    if yname == "time-in-seconds" { return "TimeInSeconds" }
    if yname == "time-in-nano-seconds" { return "TimeInNanoSeconds" }
    return ""
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetSegmentPath() string {
    return "module-up-time"
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-in-seconds"] = moduleUpTime.TimeInSeconds
    leafs["time-in-nano-seconds"] = moduleUpTime.TimeInNanoSeconds
    return leafs
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetYangName() string { return "module-up-time" }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) SetParent(parent types.Entity) { moduleUpTime.parent = parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetParent() types.Entity { return moduleUpTime.parent }

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetParentYangName() string { return "fru-info" }

