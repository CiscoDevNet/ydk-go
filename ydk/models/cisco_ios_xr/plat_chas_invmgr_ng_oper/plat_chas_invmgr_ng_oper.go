// This module contains a collection of YANG definitions
// for Cisco IOS-XR plat-chas-invmgr-ng package operational data.
// 
// This module contains definitions
// for the following management objects:
//   platform: Platform information
//   platform-inventory: platform inventory
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package plat_chas_invmgr_ng_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package plat_chas_invmgr_ng_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-plat-chas-invmgr-ng-oper platform}", reflect.TypeOf(Platform{}))
    ydk.RegisterEntity("Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform", reflect.TypeOf(Platform{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-plat-chas-invmgr-ng-oper platform-inventory}", reflect.TypeOf(PlatformInventory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory", reflect.TypeOf(PlatformInventory{}))
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of racks.
    Racks Platform_Racks
}

func (platform *Platform) GetEntityData() *types.CommonEntityData {
    platform.EntityData.YFilter = platform.YFilter
    platform.EntityData.YangName = "platform"
    platform.EntityData.BundleName = "cisco_ios_xr"
    platform.EntityData.ParentYangName = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper"
    platform.EntityData.SegmentPath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform"
    platform.EntityData.AbsolutePath = platform.EntityData.SegmentPath
    platform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platform.EntityData.Children = types.NewOrderedMap()
    platform.EntityData.Children.Append("racks", types.YChild{"Racks", &platform.Racks})
    platform.EntityData.Leafs = types.NewOrderedMap()

    platform.EntityData.YListKeys = []string {}

    return &(platform.EntityData)
}

// Platform_Racks
// Table of racks
type Platform_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack name. The type is slice of Platform_Racks_Rack.
    Rack []*Platform_Racks_Rack
}

func (racks *Platform_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "platform"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/" + racks.EntityData.SegmentPath
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = types.NewOrderedMap()
    racks.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range racks.Rack {
        racks.EntityData.Children.Append(types.GetSegmentPath(racks.Rack[i]), types.YChild{"Rack", racks.Rack[i]})
    }
    racks.EntityData.Leafs = types.NewOrderedMap()

    racks.EntityData.YListKeys = []string {}

    return &(racks.EntityData)
}

// Platform_Racks_Rack
// Rack name
type Platform_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    RackName interface{}

    // Table of slots.
    Slots Platform_Racks_Rack_Slots
}

func (rack *Platform_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.RackName, "rack-name")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/racks/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("slots", types.YChild{"Slots", &rack.Slots})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("rack-name", types.YLeaf{"RackName", rack.RackName})

    rack.EntityData.YListKeys = []string {"RackName"}

    return &(rack.EntityData)
}

// Platform_Racks_Rack_Slots
// Table of slots
type Platform_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot name. The type is slice of Platform_Racks_Rack_Slots_Slot.
    Slot []*Platform_Racks_Rack_Slots_Slot
}

func (slots *Platform_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/racks/rack/" + slots.EntityData.SegmentPath
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = types.NewOrderedMap()
    slots.EntityData.Children.Append("slot", types.YChild{"Slot", nil})
    for i := range slots.Slot {
        slots.EntityData.Children.Append(types.GetSegmentPath(slots.Slot[i]), types.YChild{"Slot", slots.Slot[i]})
    }
    slots.EntityData.Leafs = types.NewOrderedMap()

    slots.EntityData.YListKeys = []string {}

    return &(slots.EntityData)
}

// Platform_Racks_Rack_Slots_Slot
// Slot name
type Platform_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (slot *Platform_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.SlotName, "slot-name")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/racks/rack/slots/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("instances", types.YChild{"Instances", &slot.Instances})
    slot.EntityData.Children.Append("vm", types.YChild{"Vm", &slot.Vm})
    slot.EntityData.Children.Append("state", types.YChild{"State", &slot.State})
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("slot-name", types.YLeaf{"SlotName", slot.SlotName})

    slot.EntityData.YListKeys = []string {"SlotName"}

    return &(slot.EntityData)
}

// Platform_Racks_Rack_Slots_Slot_Instances
// Table of Instances
type Platform_Racks_Rack_Slots_Slot_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance name. The type is slice of
    // Platform_Racks_Rack_Slots_Slot_Instances_Instance.
    Instance []*Platform_Racks_Rack_Slots_Slot_Instances_Instance
}

func (instances *Platform_Racks_Rack_Slots_Slot_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "slot"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/racks/rack/slots/slot/" + instances.EntityData.SegmentPath
    instances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instances.EntityData.Children = types.NewOrderedMap()
    instances.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range instances.Instance {
        instances.EntityData.Children.Append(types.GetSegmentPath(instances.Instance[i]), types.YChild{"Instance", instances.Instance[i]})
    }
    instances.EntityData.Leafs = types.NewOrderedMap()

    instances.EntityData.YListKeys = []string {}

    return &(instances.EntityData)
}

// Platform_Racks_Rack_Slots_Slot_Instances_Instance
// Instance name
type Platform_Racks_Rack_Slots_Slot_Instances_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Instance name. The type is string.
    InstanceName interface{}

    // State information.
    State Platform_Racks_Rack_Slots_Slot_Instances_Instance_State
}

func (instance *Platform_Racks_Rack_Slots_Slot_Instances_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instances"
    instance.EntityData.SegmentPath = "instance" + types.AddKeyToken(instance.InstanceName, "instance-name")
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/racks/rack/slots/slot/instances/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("state", types.YChild{"State", &instance.State})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", instance.InstanceName})

    instance.EntityData.YListKeys = []string {"InstanceName"}

    return &(instance.EntityData)
}

// Platform_Racks_Rack_Slots_Slot_Instances_Instance_State
// State information
type Platform_Racks_Rack_Slots_Slot_Instances_Instance_State struct {
    EntityData types.CommonEntityData
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

func (state *Platform_Racks_Rack_Slots_Slot_Instances_Instance_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "instance"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/racks/rack/slots/slot/instances/instance/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", state.CardType})
    state.EntityData.Leafs.Append("card-redundancy-state", types.YLeaf{"CardRedundancyState", state.CardRedundancyState})
    state.EntityData.Leafs.Append("plim", types.YLeaf{"Plim", state.Plim})
    state.EntityData.Leafs.Append("state", types.YLeaf{"State", state.State})
    state.EntityData.Leafs.Append("is-monitored", types.YLeaf{"IsMonitored", state.IsMonitored})
    state.EntityData.Leafs.Append("is-powered", types.YLeaf{"IsPowered", state.IsPowered})
    state.EntityData.Leafs.Append("is-shutdown", types.YLeaf{"IsShutdown", state.IsShutdown})
    state.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", state.AdminState})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Platform_Racks_Rack_Slots_Slot_Vm
// VM information
type Platform_Racks_Rack_Slots_Slot_Vm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Type. The type is string.
    NodeDescription interface{}

    // Node Redundency Role. The type is string.
    RedRole interface{}

    // Partner Name. The type is string.
    PartnerName interface{}

    // SW status. The type is string.
    SoftwareStatus interface{}

    // Node IP Address. The type is string.
    NodeIp interface{}
}

func (vm *Platform_Racks_Rack_Slots_Slot_Vm) GetEntityData() *types.CommonEntityData {
    vm.EntityData.YFilter = vm.YFilter
    vm.EntityData.YangName = "vm"
    vm.EntityData.BundleName = "cisco_ios_xr"
    vm.EntityData.ParentYangName = "slot"
    vm.EntityData.SegmentPath = "vm"
    vm.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/racks/rack/slots/slot/" + vm.EntityData.SegmentPath
    vm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vm.EntityData.Children = types.NewOrderedMap()
    vm.EntityData.Leafs = types.NewOrderedMap()
    vm.EntityData.Leafs.Append("node-description", types.YLeaf{"NodeDescription", vm.NodeDescription})
    vm.EntityData.Leafs.Append("red-role", types.YLeaf{"RedRole", vm.RedRole})
    vm.EntityData.Leafs.Append("partner-name", types.YLeaf{"PartnerName", vm.PartnerName})
    vm.EntityData.Leafs.Append("software-status", types.YLeaf{"SoftwareStatus", vm.SoftwareStatus})
    vm.EntityData.Leafs.Append("node-ip", types.YLeaf{"NodeIp", vm.NodeIp})

    vm.EntityData.YListKeys = []string {}

    return &(vm.EntityData)
}

// Platform_Racks_Rack_Slots_Slot_State
// State information
type Platform_Racks_Rack_Slots_Slot_State struct {
    EntityData types.CommonEntityData
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

func (state *Platform_Racks_Rack_Slots_Slot_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "slot"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform/racks/rack/slots/slot/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", state.CardType})
    state.EntityData.Leafs.Append("card-redundancy-state", types.YLeaf{"CardRedundancyState", state.CardRedundancyState})
    state.EntityData.Leafs.Append("plim", types.YLeaf{"Plim", state.Plim})
    state.EntityData.Leafs.Append("state", types.YLeaf{"State", state.State})
    state.EntityData.Leafs.Append("is-monitored", types.YLeaf{"IsMonitored", state.IsMonitored})
    state.EntityData.Leafs.Append("is-powered", types.YLeaf{"IsPowered", state.IsPowered})
    state.EntityData.Leafs.Append("is-shutdown", types.YLeaf{"IsShutdown", state.IsShutdown})
    state.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", state.AdminState})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// PlatformInventory
// platform inventory
type PlatformInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of racks.
    Racks PlatformInventory_Racks
}

func (platformInventory *PlatformInventory) GetEntityData() *types.CommonEntityData {
    platformInventory.EntityData.YFilter = platformInventory.YFilter
    platformInventory.EntityData.YangName = "platform-inventory"
    platformInventory.EntityData.BundleName = "cisco_ios_xr"
    platformInventory.EntityData.ParentYangName = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper"
    platformInventory.EntityData.SegmentPath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory"
    platformInventory.EntityData.AbsolutePath = platformInventory.EntityData.SegmentPath
    platformInventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformInventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformInventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformInventory.EntityData.Children = types.NewOrderedMap()
    platformInventory.EntityData.Children.Append("racks", types.YChild{"Racks", &platformInventory.Racks})
    platformInventory.EntityData.Leafs = types.NewOrderedMap()

    platformInventory.EntityData.YListKeys = []string {}

    return &(platformInventory.EntityData)
}

// PlatformInventory_Racks
// Table of racks
type PlatformInventory_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack name. The type is slice of PlatformInventory_Racks_Rack.
    Rack []*PlatformInventory_Racks_Rack
}

func (racks *PlatformInventory_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "platform-inventory"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/" + racks.EntityData.SegmentPath
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = types.NewOrderedMap()
    racks.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range racks.Rack {
        racks.EntityData.Children.Append(types.GetSegmentPath(racks.Rack[i]), types.YChild{"Rack", racks.Rack[i]})
    }
    racks.EntityData.Leafs = types.NewOrderedMap()

    racks.EntityData.YListKeys = []string {}

    return &(racks.EntityData)
}

// PlatformInventory_Racks_Rack
// Rack name
type PlatformInventory_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of slots.
    Slots PlatformInventory_Racks_Rack_Slots

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Attributes
}

func (rack *PlatformInventory_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.Name, "name")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("slots", types.YChild{"Slots", &rack.Slots})
    rack.EntityData.Children.Append("attributes", types.YChild{"Attributes", &rack.Attributes})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("name", types.YLeaf{"Name", rack.Name})

    rack.EntityData.YListKeys = []string {"Name"}

    return &(rack.EntityData)
}

// PlatformInventory_Racks_Rack_Slots
// Table of slots
type PlatformInventory_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot name. The type is slice of PlatformInventory_Racks_Rack_Slots_Slot.
    Slot []*PlatformInventory_Racks_Rack_Slots_Slot
}

func (slots *PlatformInventory_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/" + slots.EntityData.SegmentPath
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = types.NewOrderedMap()
    slots.EntityData.Children.Append("slot", types.YChild{"Slot", nil})
    for i := range slots.Slot {
        slots.EntityData.Children.Append(types.GetSegmentPath(slots.Slot[i]), types.YChild{"Slot", slots.Slot[i]})
    }
    slots.EntityData.Leafs = types.NewOrderedMap()

    slots.EntityData.YListKeys = []string {}

    return &(slots.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot
// Slot name
type PlatformInventory_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of cards.
    Cards PlatformInventory_Racks_Rack_Slots_Slot_Cards

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Attributes
}

func (slot *PlatformInventory_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.Name, "name")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("cards", types.YChild{"Cards", &slot.Cards})
    slot.EntityData.Children.Append("attributes", types.YChild{"Attributes", &slot.Attributes})
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("name", types.YLeaf{"Name", slot.Name})

    slot.EntityData.YListKeys = []string {"Name"}

    return &(slot.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards
// Table of cards
type PlatformInventory_Racks_Rack_Slots_Slot_Cards struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Card number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card.
    Card []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card
}

func (cards *PlatformInventory_Racks_Rack_Slots_Slot_Cards) GetEntityData() *types.CommonEntityData {
    cards.EntityData.YFilter = cards.YFilter
    cards.EntityData.YangName = "cards"
    cards.EntityData.BundleName = "cisco_ios_xr"
    cards.EntityData.ParentYangName = "slot"
    cards.EntityData.SegmentPath = "cards"
    cards.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/" + cards.EntityData.SegmentPath
    cards.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cards.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cards.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cards.EntityData.Children = types.NewOrderedMap()
    cards.EntityData.Children.Append("card", types.YChild{"Card", nil})
    for i := range cards.Card {
        cards.EntityData.Children.Append(types.GetSegmentPath(cards.Card[i]), types.YChild{"Card", cards.Card[i]})
    }
    cards.EntityData.Leafs = types.NewOrderedMap()

    cards.EntityData.YListKeys = []string {}

    return &(cards.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card
// Card number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Card name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // HardwareInformationDir.
    HardwareInformation PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation

    // Table of subslots.
    SubSlots PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots

    // Table of port slots.
    Portses PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses

    // Table of port slots.
    PortSlots PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots

    // Table of  HW components .
    HwComponents PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes
}

func (card *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "cards"
    card.EntityData.SegmentPath = "card" + types.AddKeyToken(card.Name, "name")
    card.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/" + card.EntityData.SegmentPath
    card.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    card.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    card.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    card.EntityData.Children = types.NewOrderedMap()
    card.EntityData.Children.Append("hardware-information", types.YChild{"HardwareInformation", &card.HardwareInformation})
    card.EntityData.Children.Append("sub-slots", types.YChild{"SubSlots", &card.SubSlots})
    card.EntityData.Children.Append("portses", types.YChild{"Portses", &card.Portses})
    card.EntityData.Children.Append("port-slots", types.YChild{"PortSlots", &card.PortSlots})
    card.EntityData.Children.Append("hw-components", types.YChild{"HwComponents", &card.HwComponents})
    card.EntityData.Children.Append("sensors", types.YChild{"Sensors", &card.Sensors})
    card.EntityData.Children.Append("attributes", types.YChild{"Attributes", &card.Attributes})
    card.EntityData.Leafs = types.NewOrderedMap()
    card.EntityData.Leafs.Append("name", types.YLeaf{"Name", card.Name})

    card.EntityData.YListKeys = []string {"Name"}

    return &(card.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation
// HardwareInformationDir
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation struct {
    EntityData types.CommonEntityData
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

func (hardwareInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation) GetEntityData() *types.CommonEntityData {
    hardwareInformation.EntityData.YFilter = hardwareInformation.YFilter
    hardwareInformation.EntityData.YangName = "hardware-information"
    hardwareInformation.EntityData.BundleName = "cisco_ios_xr"
    hardwareInformation.EntityData.ParentYangName = "card"
    hardwareInformation.EntityData.SegmentPath = "hardware-information"
    hardwareInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/" + hardwareInformation.EntityData.SegmentPath
    hardwareInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareInformation.EntityData.Children = types.NewOrderedMap()
    hardwareInformation.EntityData.Children.Append("processor-information", types.YChild{"ProcessorInformation", &hardwareInformation.ProcessorInformation})
    hardwareInformation.EntityData.Children.Append("motherboard-information", types.YChild{"MotherboardInformation", &hardwareInformation.MotherboardInformation})
    hardwareInformation.EntityData.Children.Append("bootflash-information", types.YChild{"BootflashInformation", &hardwareInformation.BootflashInformation})
    hardwareInformation.EntityData.Children.Append("disk-information", types.YChild{"DiskInformation", &hardwareInformation.DiskInformation})
    hardwareInformation.EntityData.Leafs = types.NewOrderedMap()

    hardwareInformation.EntityData.YListKeys = []string {}

    return &(hardwareInformation.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation
// ProcesorInformation
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type e.g. 7457. The type is string with length: 0..255.
    ProcessorType interface{}

    // Speed e.g. 1197Mhz. The type is string with length: 0..255.
    Speed interface{}

    // Revision. e.g 1.1. The type is string with length: 0..255.
    Revision interface{}
}

func (processorInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_ProcessorInformation) GetEntityData() *types.CommonEntityData {
    processorInformation.EntityData.YFilter = processorInformation.YFilter
    processorInformation.EntityData.YangName = "processor-information"
    processorInformation.EntityData.BundleName = "cisco_ios_xr"
    processorInformation.EntityData.ParentYangName = "hardware-information"
    processorInformation.EntityData.SegmentPath = "processor-information"
    processorInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hardware-information/" + processorInformation.EntityData.SegmentPath
    processorInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processorInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processorInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processorInformation.EntityData.Children = types.NewOrderedMap()
    processorInformation.EntityData.Leafs = types.NewOrderedMap()
    processorInformation.EntityData.Leafs.Append("processor-type", types.YLeaf{"ProcessorType", processorInformation.ProcessorType})
    processorInformation.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", processorInformation.Speed})
    processorInformation.EntityData.Leafs.Append("revision", types.YLeaf{"Revision", processorInformation.Revision})

    processorInformation.EntityData.YListKeys = []string {}

    return &(processorInformation.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation
// MotherboardInformation
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation struct {
    EntityData types.CommonEntityData
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

func (motherboardInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation) GetEntityData() *types.CommonEntityData {
    motherboardInformation.EntityData.YFilter = motherboardInformation.YFilter
    motherboardInformation.EntityData.YangName = "motherboard-information"
    motherboardInformation.EntityData.BundleName = "cisco_ios_xr"
    motherboardInformation.EntityData.ParentYangName = "hardware-information"
    motherboardInformation.EntityData.SegmentPath = "motherboard-information"
    motherboardInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hardware-information/" + motherboardInformation.EntityData.SegmentPath
    motherboardInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    motherboardInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    motherboardInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    motherboardInformation.EntityData.Children = types.NewOrderedMap()
    motherboardInformation.EntityData.Children.Append("rom", types.YChild{"Rom", &motherboardInformation.Rom})
    motherboardInformation.EntityData.Children.Append("bootflash", types.YChild{"Bootflash", &motherboardInformation.Bootflash})
    motherboardInformation.EntityData.Children.Append("processor", types.YChild{"Processor", &motherboardInformation.Processor})
    motherboardInformation.EntityData.Leafs = types.NewOrderedMap()
    motherboardInformation.EntityData.Leafs.Append("main-memory-size", types.YLeaf{"MainMemorySize", motherboardInformation.MainMemorySize})
    motherboardInformation.EntityData.Leafs.Append("nvram-size", types.YLeaf{"NvramSize", motherboardInformation.NvramSize})

    motherboardInformation.EntityData.YListKeys = []string {}

    return &(motherboardInformation.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom
// ROM information
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom struct {
    EntityData types.CommonEntityData
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

func (rom *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Rom) GetEntityData() *types.CommonEntityData {
    rom.EntityData.YFilter = rom.YFilter
    rom.EntityData.YangName = "rom"
    rom.EntityData.BundleName = "cisco_ios_xr"
    rom.EntityData.ParentYangName = "motherboard-information"
    rom.EntityData.SegmentPath = "rom"
    rom.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hardware-information/motherboard-information/" + rom.EntityData.SegmentPath
    rom.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rom.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rom.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rom.EntityData.Children = types.NewOrderedMap()
    rom.EntityData.Leafs = types.NewOrderedMap()
    rom.EntityData.Leafs.Append("image-name", types.YLeaf{"ImageName", rom.ImageName})
    rom.EntityData.Leafs.Append("major-version", types.YLeaf{"MajorVersion", rom.MajorVersion})
    rom.EntityData.Leafs.Append("minor-version", types.YLeaf{"MinorVersion", rom.MinorVersion})
    rom.EntityData.Leafs.Append("micro-image-version", types.YLeaf{"MicroImageVersion", rom.MicroImageVersion})
    rom.EntityData.Leafs.Append("platform-specific", types.YLeaf{"PlatformSpecific", rom.PlatformSpecific})
    rom.EntityData.Leafs.Append("release-type", types.YLeaf{"ReleaseType", rom.ReleaseType})

    rom.EntityData.YListKeys = []string {}

    return &(rom.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash
// Bootflash information
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash struct {
    EntityData types.CommonEntityData
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

func (bootflash *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Bootflash) GetEntityData() *types.CommonEntityData {
    bootflash.EntityData.YFilter = bootflash.YFilter
    bootflash.EntityData.YangName = "bootflash"
    bootflash.EntityData.BundleName = "cisco_ios_xr"
    bootflash.EntityData.ParentYangName = "motherboard-information"
    bootflash.EntityData.SegmentPath = "bootflash"
    bootflash.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hardware-information/motherboard-information/" + bootflash.EntityData.SegmentPath
    bootflash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootflash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootflash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootflash.EntityData.Children = types.NewOrderedMap()
    bootflash.EntityData.Leafs = types.NewOrderedMap()
    bootflash.EntityData.Leafs.Append("image-name", types.YLeaf{"ImageName", bootflash.ImageName})
    bootflash.EntityData.Leafs.Append("platform-type", types.YLeaf{"PlatformType", bootflash.PlatformType})
    bootflash.EntityData.Leafs.Append("major-version", types.YLeaf{"MajorVersion", bootflash.MajorVersion})
    bootflash.EntityData.Leafs.Append("minor-version", types.YLeaf{"MinorVersion", bootflash.MinorVersion})
    bootflash.EntityData.Leafs.Append("micro-image-version", types.YLeaf{"MicroImageVersion", bootflash.MicroImageVersion})
    bootflash.EntityData.Leafs.Append("platform-specific", types.YLeaf{"PlatformSpecific", bootflash.PlatformSpecific})
    bootflash.EntityData.Leafs.Append("release-type", types.YLeaf{"ReleaseType", bootflash.ReleaseType})
    bootflash.EntityData.Leafs.Append("bootflash-type", types.YLeaf{"BootflashType", bootflash.BootflashType})
    bootflash.EntityData.Leafs.Append("bootflash-size", types.YLeaf{"BootflashSize", bootflash.BootflashSize})
    bootflash.EntityData.Leafs.Append("sector-size", types.YLeaf{"SectorSize", bootflash.SectorSize})

    bootflash.EntityData.YListKeys = []string {}

    return &(bootflash.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor
// Processor information
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type e.g. 7457. The type is string with length: 0..255.
    ProcessorType interface{}

    // Speed e.g. 1197Mhz. The type is string with length: 0..255.
    Speed interface{}

    // Revision. e.g 1.1. The type is string with length: 0..255.
    Revision interface{}
}

func (processor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_MotherboardInformation_Processor) GetEntityData() *types.CommonEntityData {
    processor.EntityData.YFilter = processor.YFilter
    processor.EntityData.YangName = "processor"
    processor.EntityData.BundleName = "cisco_ios_xr"
    processor.EntityData.ParentYangName = "motherboard-information"
    processor.EntityData.SegmentPath = "processor"
    processor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hardware-information/motherboard-information/" + processor.EntityData.SegmentPath
    processor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processor.EntityData.Children = types.NewOrderedMap()
    processor.EntityData.Leafs = types.NewOrderedMap()
    processor.EntityData.Leafs.Append("processor-type", types.YLeaf{"ProcessorType", processor.ProcessorType})
    processor.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", processor.Speed})
    processor.EntityData.Leafs.Append("revision", types.YLeaf{"Revision", processor.Revision})

    processor.EntityData.YListKeys = []string {}

    return &(processor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation
// BootflashInformation
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation struct {
    EntityData types.CommonEntityData
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

func (bootflashInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_BootflashInformation) GetEntityData() *types.CommonEntityData {
    bootflashInformation.EntityData.YFilter = bootflashInformation.YFilter
    bootflashInformation.EntityData.YangName = "bootflash-information"
    bootflashInformation.EntityData.BundleName = "cisco_ios_xr"
    bootflashInformation.EntityData.ParentYangName = "hardware-information"
    bootflashInformation.EntityData.SegmentPath = "bootflash-information"
    bootflashInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hardware-information/" + bootflashInformation.EntityData.SegmentPath
    bootflashInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootflashInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootflashInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootflashInformation.EntityData.Children = types.NewOrderedMap()
    bootflashInformation.EntityData.Leafs = types.NewOrderedMap()
    bootflashInformation.EntityData.Leafs.Append("image-name", types.YLeaf{"ImageName", bootflashInformation.ImageName})
    bootflashInformation.EntityData.Leafs.Append("platform-type", types.YLeaf{"PlatformType", bootflashInformation.PlatformType})
    bootflashInformation.EntityData.Leafs.Append("major-version", types.YLeaf{"MajorVersion", bootflashInformation.MajorVersion})
    bootflashInformation.EntityData.Leafs.Append("minor-version", types.YLeaf{"MinorVersion", bootflashInformation.MinorVersion})
    bootflashInformation.EntityData.Leafs.Append("micro-image-version", types.YLeaf{"MicroImageVersion", bootflashInformation.MicroImageVersion})
    bootflashInformation.EntityData.Leafs.Append("platform-specific", types.YLeaf{"PlatformSpecific", bootflashInformation.PlatformSpecific})
    bootflashInformation.EntityData.Leafs.Append("release-type", types.YLeaf{"ReleaseType", bootflashInformation.ReleaseType})
    bootflashInformation.EntityData.Leafs.Append("bootflash-type", types.YLeaf{"BootflashType", bootflashInformation.BootflashType})
    bootflashInformation.EntityData.Leafs.Append("bootflash-size", types.YLeaf{"BootflashSize", bootflashInformation.BootflashSize})
    bootflashInformation.EntityData.Leafs.Append("sector-size", types.YLeaf{"SectorSize", bootflashInformation.SectorSize})

    bootflashInformation.EntityData.YListKeys = []string {}

    return &(bootflashInformation.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation
// DiskInformation
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation struct {
    EntityData types.CommonEntityData
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
    Disks []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks
}

func (diskInformation *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation) GetEntityData() *types.CommonEntityData {
    diskInformation.EntityData.YFilter = diskInformation.YFilter
    diskInformation.EntityData.YangName = "disk-information"
    diskInformation.EntityData.BundleName = "cisco_ios_xr"
    diskInformation.EntityData.ParentYangName = "hardware-information"
    diskInformation.EntityData.SegmentPath = "disk-information"
    diskInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hardware-information/" + diskInformation.EntityData.SegmentPath
    diskInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diskInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diskInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diskInformation.EntityData.Children = types.NewOrderedMap()
    diskInformation.EntityData.Children.Append("disks", types.YChild{"Disks", nil})
    for i := range diskInformation.Disks {
        types.SetYListKey(diskInformation.Disks[i], i)
        diskInformation.EntityData.Children.Append(types.GetSegmentPath(diskInformation.Disks[i]), types.YChild{"Disks", diskInformation.Disks[i]})
    }
    diskInformation.EntityData.Leafs = types.NewOrderedMap()
    diskInformation.EntityData.Leafs.Append("disk-name", types.YLeaf{"DiskName", diskInformation.DiskName})
    diskInformation.EntityData.Leafs.Append("disk-size", types.YLeaf{"DiskSize", diskInformation.DiskSize})
    diskInformation.EntityData.Leafs.Append("sector-size", types.YLeaf{"SectorSize", diskInformation.SectorSize})

    diskInformation.EntityData.YListKeys = []string {}

    return &(diskInformation.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks
// Disk attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Disk name. The type is string with length: 0..255.
    DiskName interface{}

    // Disk size in mega-bytes. The type is interface{} with range: 0..4294967295.
    // Units are megabyte.
    DiskSize interface{}

    // Disk sector size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    SectorSize interface{}
}

func (disks *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HardwareInformation_DiskInformation_Disks) GetEntityData() *types.CommonEntityData {
    disks.EntityData.YFilter = disks.YFilter
    disks.EntityData.YangName = "disks"
    disks.EntityData.BundleName = "cisco_ios_xr"
    disks.EntityData.ParentYangName = "disk-information"
    disks.EntityData.SegmentPath = "disks" + types.AddNoKeyToken(disks)
    disks.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hardware-information/disk-information/" + disks.EntityData.SegmentPath
    disks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disks.EntityData.Children = types.NewOrderedMap()
    disks.EntityData.Leafs = types.NewOrderedMap()
    disks.EntityData.Leafs.Append("disk-name", types.YLeaf{"DiskName", disks.DiskName})
    disks.EntityData.Leafs.Append("disk-size", types.YLeaf{"DiskSize", disks.DiskSize})
    disks.EntityData.Leafs.Append("sector-size", types.YLeaf{"SectorSize", disks.SectorSize})

    disks.EntityData.YListKeys = []string {}

    return &(disks.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots
// Table of subslots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subslot number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot.
    SubSlot []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
}

func (subSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots) GetEntityData() *types.CommonEntityData {
    subSlots.EntityData.YFilter = subSlots.YFilter
    subSlots.EntityData.YangName = "sub-slots"
    subSlots.EntityData.BundleName = "cisco_ios_xr"
    subSlots.EntityData.ParentYangName = "card"
    subSlots.EntityData.SegmentPath = "sub-slots"
    subSlots.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/" + subSlots.EntityData.SegmentPath
    subSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subSlots.EntityData.Children = types.NewOrderedMap()
    subSlots.EntityData.Children.Append("sub-slot", types.YChild{"SubSlot", nil})
    for i := range subSlots.SubSlot {
        subSlots.EntityData.Children.Append(types.GetSegmentPath(subSlots.SubSlot[i]), types.YChild{"SubSlot", subSlots.SubSlot[i]})
    }
    subSlots.EntityData.Leafs = types.NewOrderedMap()

    subSlots.EntityData.YListKeys = []string {}

    return &(subSlots.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot
// Subslot number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subslot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Module of a subslot.
    Module PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes
}

func (subSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot) GetEntityData() *types.CommonEntityData {
    subSlot.EntityData.YFilter = subSlot.YFilter
    subSlot.EntityData.YangName = "sub-slot"
    subSlot.EntityData.BundleName = "cisco_ios_xr"
    subSlot.EntityData.ParentYangName = "sub-slots"
    subSlot.EntityData.SegmentPath = "sub-slot" + types.AddKeyToken(subSlot.Name, "name")
    subSlot.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/" + subSlot.EntityData.SegmentPath
    subSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subSlot.EntityData.Children = types.NewOrderedMap()
    subSlot.EntityData.Children.Append("module", types.YChild{"Module", &subSlot.Module})
    subSlot.EntityData.Children.Append("attributes", types.YChild{"Attributes", &subSlot.Attributes})
    subSlot.EntityData.Leafs = types.NewOrderedMap()
    subSlot.EntityData.Leafs.Append("name", types.YLeaf{"Name", subSlot.Name})

    subSlot.EntityData.YListKeys = []string {"Name"}

    return &(subSlot.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module
// Module of a subslot
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of port slots.
    PortSlots PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes
}

func (module *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module) GetEntityData() *types.CommonEntityData {
    module.EntityData.YFilter = module.YFilter
    module.EntityData.YangName = "module"
    module.EntityData.BundleName = "cisco_ios_xr"
    module.EntityData.ParentYangName = "sub-slot"
    module.EntityData.SegmentPath = "module"
    module.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/" + module.EntityData.SegmentPath
    module.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    module.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    module.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    module.EntityData.Children = types.NewOrderedMap()
    module.EntityData.Children.Append("port-slots", types.YChild{"PortSlots", &module.PortSlots})
    module.EntityData.Children.Append("sensors", types.YChild{"Sensors", &module.Sensors})
    module.EntityData.Children.Append("attributes", types.YChild{"Attributes", &module.Attributes})
    module.EntityData.Leafs = types.NewOrderedMap()

    module.EntityData.YListKeys = []string {}

    return &(module.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots
// Table of port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port slot number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot.
    PortSlot []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots) GetEntityData() *types.CommonEntityData {
    portSlots.EntityData.YFilter = portSlots.YFilter
    portSlots.EntityData.YangName = "port-slots"
    portSlots.EntityData.BundleName = "cisco_ios_xr"
    portSlots.EntityData.ParentYangName = "module"
    portSlots.EntityData.SegmentPath = "port-slots"
    portSlots.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/" + portSlots.EntityData.SegmentPath
    portSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlots.EntityData.Children = types.NewOrderedMap()
    portSlots.EntityData.Children.Append("port-slot", types.YChild{"PortSlot", nil})
    for i := range portSlots.PortSlot {
        portSlots.EntityData.Children.Append(types.GetSegmentPath(portSlots.PortSlot[i]), types.YChild{"PortSlot", portSlots.PortSlot[i]})
    }
    portSlots.EntityData.Leafs = types.NewOrderedMap()

    portSlots.EntityData.YListKeys = []string {}

    return &(portSlots.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot
// Port slot number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of port slots.
    Portses PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot) GetEntityData() *types.CommonEntityData {
    portSlot.EntityData.YFilter = portSlot.YFilter
    portSlot.EntityData.YangName = "port-slot"
    portSlot.EntityData.BundleName = "cisco_ios_xr"
    portSlot.EntityData.ParentYangName = "port-slots"
    portSlot.EntityData.SegmentPath = "port-slot" + types.AddKeyToken(portSlot.Name, "name")
    portSlot.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/" + portSlot.EntityData.SegmentPath
    portSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlot.EntityData.Children = types.NewOrderedMap()
    portSlot.EntityData.Children.Append("portses", types.YChild{"Portses", &portSlot.Portses})
    portSlot.EntityData.Children.Append("sensors", types.YChild{"Sensors", &portSlot.Sensors})
    portSlot.EntityData.Children.Append("attributes", types.YChild{"Attributes", &portSlot.Attributes})
    portSlot.EntityData.Leafs = types.NewOrderedMap()
    portSlot.EntityData.Leafs.Append("name", types.YLeaf{"Name", portSlot.Name})

    portSlot.EntityData.YListKeys = []string {"Name"}

    return &(portSlot.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses
// Table of port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports.
    Ports []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses) GetEntityData() *types.CommonEntityData {
    portses.EntityData.YFilter = portses.YFilter
    portses.EntityData.YangName = "portses"
    portses.EntityData.BundleName = "cisco_ios_xr"
    portses.EntityData.ParentYangName = "port-slot"
    portses.EntityData.SegmentPath = "portses"
    portses.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/" + portses.EntityData.SegmentPath
    portses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portses.EntityData.Children = types.NewOrderedMap()
    portses.EntityData.Children.Append("ports", types.YChild{"Ports", nil})
    for i := range portses.Ports {
        portses.EntityData.Children.Append(types.GetSegmentPath(portses.Ports[i]), types.YChild{"Ports", portses.Ports[i]})
    }
    portses.EntityData.Leafs = types.NewOrderedMap()

    portses.EntityData.YListKeys = []string {}

    return &(portses.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports
// Port number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xr"
    ports.EntityData.ParentYangName = "portses"
    ports.EntityData.SegmentPath = "ports" + types.AddKeyToken(ports.Name, "name")
    ports.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/" + ports.EntityData.SegmentPath
    ports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ports.EntityData.Children = types.NewOrderedMap()
    ports.EntityData.Children.Append("hw-components", types.YChild{"HwComponents", &ports.HwComponents})
    ports.EntityData.Children.Append("sensors", types.YChild{"Sensors", &ports.Sensors})
    ports.EntityData.Children.Append("attributes", types.YChild{"Attributes", &ports.Attributes})
    ports.EntityData.Leafs = types.NewOrderedMap()
    ports.EntityData.Leafs.Append("name", types.YLeaf{"Name", ports.Name})

    ports.EntityData.YListKeys = []string {"Name"}

    return &(ports.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents
// Table of  HW components 
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HW component number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent.
    HwComponent []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents) GetEntityData() *types.CommonEntityData {
    hwComponents.EntityData.YFilter = hwComponents.YFilter
    hwComponents.EntityData.YangName = "hw-components"
    hwComponents.EntityData.BundleName = "cisco_ios_xr"
    hwComponents.EntityData.ParentYangName = "ports"
    hwComponents.EntityData.SegmentPath = "hw-components"
    hwComponents.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/" + hwComponents.EntityData.SegmentPath
    hwComponents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponents.EntityData.Children = types.NewOrderedMap()
    hwComponents.EntityData.Children.Append("hw-component", types.YChild{"HwComponent", nil})
    for i := range hwComponents.HwComponent {
        hwComponents.EntityData.Children.Append(types.GetSegmentPath(hwComponents.HwComponent[i]), types.YChild{"HwComponent", hwComponents.HwComponent[i]})
    }
    hwComponents.EntityData.Leafs = types.NewOrderedMap()

    hwComponents.EntityData.YListKeys = []string {}

    return &(hwComponents.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent
// HW component number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. HW component name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetEntityData() *types.CommonEntityData {
    hwComponent.EntityData.YFilter = hwComponent.YFilter
    hwComponent.EntityData.YangName = "hw-component"
    hwComponent.EntityData.BundleName = "cisco_ios_xr"
    hwComponent.EntityData.ParentYangName = "hw-components"
    hwComponent.EntityData.SegmentPath = "hw-component" + types.AddKeyToken(hwComponent.Name, "name")
    hwComponent.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/" + hwComponent.EntityData.SegmentPath
    hwComponent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponent.EntityData.Children = types.NewOrderedMap()
    hwComponent.EntityData.Children.Append("sensors", types.YChild{"Sensors", &hwComponent.Sensors})
    hwComponent.EntityData.Children.Append("attributes", types.YChild{"Attributes", &hwComponent.Attributes})
    hwComponent.EntityData.Leafs = types.NewOrderedMap()
    hwComponent.EntityData.Leafs.Append("name", types.YLeaf{"Name", hwComponent.Name})

    hwComponent.EntityData.YListKeys = []string {"Name"}

    return &(hwComponent.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "hw-component"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "hw-component"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/hw-components/hw-component/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "ports"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "ports"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/portses/ports/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "port-slot"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "port-slot"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/port-slots/port-slot/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "module"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "module"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Module_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/module/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sub-slot"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_SubSlots_SubSlot_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sub-slots/sub-slot/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses
// Table of port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports.
    Ports []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses) GetEntityData() *types.CommonEntityData {
    portses.EntityData.YFilter = portses.YFilter
    portses.EntityData.YangName = "portses"
    portses.EntityData.BundleName = "cisco_ios_xr"
    portses.EntityData.ParentYangName = "card"
    portses.EntityData.SegmentPath = "portses"
    portses.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/" + portses.EntityData.SegmentPath
    portses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portses.EntityData.Children = types.NewOrderedMap()
    portses.EntityData.Children.Append("ports", types.YChild{"Ports", nil})
    for i := range portses.Ports {
        portses.EntityData.Children.Append(types.GetSegmentPath(portses.Ports[i]), types.YChild{"Ports", portses.Ports[i]})
    }
    portses.EntityData.Leafs = types.NewOrderedMap()

    portses.EntityData.YListKeys = []string {}

    return &(portses.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports
// Port number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of  HW components .
    HwComponents PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes
}

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xr"
    ports.EntityData.ParentYangName = "portses"
    ports.EntityData.SegmentPath = "ports" + types.AddKeyToken(ports.Name, "name")
    ports.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/" + ports.EntityData.SegmentPath
    ports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ports.EntityData.Children = types.NewOrderedMap()
    ports.EntityData.Children.Append("hw-components", types.YChild{"HwComponents", &ports.HwComponents})
    ports.EntityData.Children.Append("sensors", types.YChild{"Sensors", &ports.Sensors})
    ports.EntityData.Children.Append("attributes", types.YChild{"Attributes", &ports.Attributes})
    ports.EntityData.Leafs = types.NewOrderedMap()
    ports.EntityData.Leafs.Append("name", types.YLeaf{"Name", ports.Name})

    ports.EntityData.YListKeys = []string {"Name"}

    return &(ports.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents
// Table of  HW components 
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HW component number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent.
    HwComponent []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents) GetEntityData() *types.CommonEntityData {
    hwComponents.EntityData.YFilter = hwComponents.YFilter
    hwComponents.EntityData.YangName = "hw-components"
    hwComponents.EntityData.BundleName = "cisco_ios_xr"
    hwComponents.EntityData.ParentYangName = "ports"
    hwComponents.EntityData.SegmentPath = "hw-components"
    hwComponents.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/" + hwComponents.EntityData.SegmentPath
    hwComponents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponents.EntityData.Children = types.NewOrderedMap()
    hwComponents.EntityData.Children.Append("hw-component", types.YChild{"HwComponent", nil})
    for i := range hwComponents.HwComponent {
        hwComponents.EntityData.Children.Append(types.GetSegmentPath(hwComponents.HwComponent[i]), types.YChild{"HwComponent", hwComponents.HwComponent[i]})
    }
    hwComponents.EntityData.Leafs = types.NewOrderedMap()

    hwComponents.EntityData.YListKeys = []string {}

    return &(hwComponents.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent
// HW component number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. HW component name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent) GetEntityData() *types.CommonEntityData {
    hwComponent.EntityData.YFilter = hwComponent.YFilter
    hwComponent.EntityData.YangName = "hw-component"
    hwComponent.EntityData.BundleName = "cisco_ios_xr"
    hwComponent.EntityData.ParentYangName = "hw-components"
    hwComponent.EntityData.SegmentPath = "hw-component" + types.AddKeyToken(hwComponent.Name, "name")
    hwComponent.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/" + hwComponent.EntityData.SegmentPath
    hwComponent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponent.EntityData.Children = types.NewOrderedMap()
    hwComponent.EntityData.Children.Append("sensors", types.YChild{"Sensors", &hwComponent.Sensors})
    hwComponent.EntityData.Children.Append("attributes", types.YChild{"Attributes", &hwComponent.Attributes})
    hwComponent.EntityData.Leafs = types.NewOrderedMap()
    hwComponent.EntityData.Leafs.Append("name", types.YLeaf{"Name", hwComponent.Name})

    hwComponent.EntityData.YListKeys = []string {"Name"}

    return &(hwComponent.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "hw-component"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "hw-component"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/hw-components/hw-component/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "ports"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "ports"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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
    LastOperationalStateChange PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange

    // Module up time.
    ModuleUpTime PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo_ModuleUpTime
}

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/portses/ports/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots
// Table of port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port slot number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot.
    PortSlot []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
}

func (portSlots *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots) GetEntityData() *types.CommonEntityData {
    portSlots.EntityData.YFilter = portSlots.YFilter
    portSlots.EntityData.YangName = "port-slots"
    portSlots.EntityData.BundleName = "cisco_ios_xr"
    portSlots.EntityData.ParentYangName = "card"
    portSlots.EntityData.SegmentPath = "port-slots"
    portSlots.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/" + portSlots.EntityData.SegmentPath
    portSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlots.EntityData.Children = types.NewOrderedMap()
    portSlots.EntityData.Children.Append("port-slot", types.YChild{"PortSlot", nil})
    for i := range portSlots.PortSlot {
        portSlots.EntityData.Children.Append(types.GetSegmentPath(portSlots.PortSlot[i]), types.YChild{"PortSlot", portSlots.PortSlot[i]})
    }
    portSlots.EntityData.Leafs = types.NewOrderedMap()

    portSlots.EntityData.YListKeys = []string {}

    return &(portSlots.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot
// Port slot number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of port slots.
    Portses PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes
}

func (portSlot *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot) GetEntityData() *types.CommonEntityData {
    portSlot.EntityData.YFilter = portSlot.YFilter
    portSlot.EntityData.YangName = "port-slot"
    portSlot.EntityData.BundleName = "cisco_ios_xr"
    portSlot.EntityData.ParentYangName = "port-slots"
    portSlot.EntityData.SegmentPath = "port-slot" + types.AddKeyToken(portSlot.Name, "name")
    portSlot.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/" + portSlot.EntityData.SegmentPath
    portSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portSlot.EntityData.Children = types.NewOrderedMap()
    portSlot.EntityData.Children.Append("portses", types.YChild{"Portses", &portSlot.Portses})
    portSlot.EntityData.Children.Append("sensors", types.YChild{"Sensors", &portSlot.Sensors})
    portSlot.EntityData.Children.Append("attributes", types.YChild{"Attributes", &portSlot.Attributes})
    portSlot.EntityData.Leafs = types.NewOrderedMap()
    portSlot.EntityData.Leafs.Append("name", types.YLeaf{"Name", portSlot.Name})

    portSlot.EntityData.YListKeys = []string {"Name"}

    return &(portSlot.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses
// Table of port slots
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports.
    Ports []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports
}

func (portses *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses) GetEntityData() *types.CommonEntityData {
    portses.EntityData.YFilter = portses.YFilter
    portses.EntityData.YangName = "portses"
    portses.EntityData.BundleName = "cisco_ios_xr"
    portses.EntityData.ParentYangName = "port-slot"
    portses.EntityData.SegmentPath = "portses"
    portses.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/" + portses.EntityData.SegmentPath
    portses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portses.EntityData.Children = types.NewOrderedMap()
    portses.EntityData.Children.Append("ports", types.YChild{"Ports", nil})
    for i := range portses.Ports {
        portses.EntityData.Children.Append(types.GetSegmentPath(portses.Ports[i]), types.YChild{"Ports", portses.Ports[i]})
    }
    portses.EntityData.Leafs = types.NewOrderedMap()

    portses.EntityData.YListKeys = []string {}

    return &(portses.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports
// Port number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (ports *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xr"
    ports.EntityData.ParentYangName = "portses"
    ports.EntityData.SegmentPath = "ports" + types.AddKeyToken(ports.Name, "name")
    ports.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/" + ports.EntityData.SegmentPath
    ports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ports.EntityData.Children = types.NewOrderedMap()
    ports.EntityData.Children.Append("hw-components", types.YChild{"HwComponents", &ports.HwComponents})
    ports.EntityData.Children.Append("sensors", types.YChild{"Sensors", &ports.Sensors})
    ports.EntityData.Children.Append("attributes", types.YChild{"Attributes", &ports.Attributes})
    ports.EntityData.Leafs = types.NewOrderedMap()
    ports.EntityData.Leafs.Append("name", types.YLeaf{"Name", ports.Name})

    ports.EntityData.YListKeys = []string {"Name"}

    return &(ports.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents
// Table of  HW components 
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HW component number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent.
    HwComponent []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents) GetEntityData() *types.CommonEntityData {
    hwComponents.EntityData.YFilter = hwComponents.YFilter
    hwComponents.EntityData.YangName = "hw-components"
    hwComponents.EntityData.BundleName = "cisco_ios_xr"
    hwComponents.EntityData.ParentYangName = "ports"
    hwComponents.EntityData.SegmentPath = "hw-components"
    hwComponents.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/" + hwComponents.EntityData.SegmentPath
    hwComponents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponents.EntityData.Children = types.NewOrderedMap()
    hwComponents.EntityData.Children.Append("hw-component", types.YChild{"HwComponent", nil})
    for i := range hwComponents.HwComponent {
        hwComponents.EntityData.Children.Append(types.GetSegmentPath(hwComponents.HwComponent[i]), types.YChild{"HwComponent", hwComponents.HwComponent[i]})
    }
    hwComponents.EntityData.Leafs = types.NewOrderedMap()

    hwComponents.EntityData.YListKeys = []string {}

    return &(hwComponents.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent
// HW component number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. HW component name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent) GetEntityData() *types.CommonEntityData {
    hwComponent.EntityData.YFilter = hwComponent.YFilter
    hwComponent.EntityData.YangName = "hw-component"
    hwComponent.EntityData.BundleName = "cisco_ios_xr"
    hwComponent.EntityData.ParentYangName = "hw-components"
    hwComponent.EntityData.SegmentPath = "hw-component" + types.AddKeyToken(hwComponent.Name, "name")
    hwComponent.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/" + hwComponent.EntityData.SegmentPath
    hwComponent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponent.EntityData.Children = types.NewOrderedMap()
    hwComponent.EntityData.Children.Append("sensors", types.YChild{"Sensors", &hwComponent.Sensors})
    hwComponent.EntityData.Children.Append("attributes", types.YChild{"Attributes", &hwComponent.Attributes})
    hwComponent.EntityData.Leafs = types.NewOrderedMap()
    hwComponent.EntityData.Leafs.Append("name", types.YLeaf{"Name", hwComponent.Name})

    hwComponent.EntityData.YListKeys = []string {"Name"}

    return &(hwComponent.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "hw-component"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "hw-component"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/hw-components/hw-component/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "ports"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "ports"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Portses_Ports_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/portses/ports/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "port-slot"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "port-slot"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_PortSlots_PortSlot_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/port-slots/port-slot/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents
// Table of  HW components 
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HW component number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent.
    HwComponent []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
}

func (hwComponents *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents) GetEntityData() *types.CommonEntityData {
    hwComponents.EntityData.YFilter = hwComponents.YFilter
    hwComponents.EntityData.YangName = "hw-components"
    hwComponents.EntityData.BundleName = "cisco_ios_xr"
    hwComponents.EntityData.ParentYangName = "card"
    hwComponents.EntityData.SegmentPath = "hw-components"
    hwComponents.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/" + hwComponents.EntityData.SegmentPath
    hwComponents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponents.EntityData.Children = types.NewOrderedMap()
    hwComponents.EntityData.Children.Append("hw-component", types.YChild{"HwComponent", nil})
    for i := range hwComponents.HwComponent {
        hwComponents.EntityData.Children.Append(types.GetSegmentPath(hwComponents.HwComponent[i]), types.YChild{"HwComponent", hwComponents.HwComponent[i]})
    }
    hwComponents.EntityData.Leafs = types.NewOrderedMap()

    hwComponents.EntityData.YListKeys = []string {}

    return &(hwComponents.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent
// HW component number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. HW component name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Table of sensors.
    Sensors PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes
}

func (hwComponent *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent) GetEntityData() *types.CommonEntityData {
    hwComponent.EntityData.YFilter = hwComponent.YFilter
    hwComponent.EntityData.YangName = "hw-component"
    hwComponent.EntityData.BundleName = "cisco_ios_xr"
    hwComponent.EntityData.ParentYangName = "hw-components"
    hwComponent.EntityData.SegmentPath = "hw-component" + types.AddKeyToken(hwComponent.Name, "name")
    hwComponent.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/" + hwComponent.EntityData.SegmentPath
    hwComponent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwComponent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwComponent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwComponent.EntityData.Children = types.NewOrderedMap()
    hwComponent.EntityData.Children.Append("sensors", types.YChild{"Sensors", &hwComponent.Sensors})
    hwComponent.EntityData.Children.Append("attributes", types.YChild{"Attributes", &hwComponent.Attributes})
    hwComponent.EntityData.Leafs = types.NewOrderedMap()
    hwComponent.EntityData.Leafs.Append("name", types.YLeaf{"Name", hwComponent.Name})

    hwComponent.EntityData.YListKeys = []string {"Name"}

    return &(hwComponent.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "hw-component"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "hw-component"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_HwComponents_HwComponent_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/hw-components/hw-component/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors
// Table of sensors
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor number. The type is slice of
    // PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor.
    Sensor []*PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
}

func (sensors *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors) GetEntityData() *types.CommonEntityData {
    sensors.EntityData.YFilter = sensors.YFilter
    sensors.EntityData.YangName = "sensors"
    sensors.EntityData.BundleName = "cisco_ios_xr"
    sensors.EntityData.ParentYangName = "card"
    sensors.EntityData.SegmentPath = "sensors"
    sensors.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/" + sensors.EntityData.SegmentPath
    sensors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensors.EntityData.Children = types.NewOrderedMap()
    sensors.EntityData.Children.Append("sensor", types.YChild{"Sensor", nil})
    for i := range sensors.Sensor {
        sensors.EntityData.Children.Append(types.GetSegmentPath(sensors.Sensor[i]), types.YChild{"Sensor", sensors.Sensor[i]})
    }
    sensors.EntityData.Leafs = types.NewOrderedMap()

    sensors.EntityData.YListKeys = []string {}

    return &(sensors.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor
// Sensor number
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sensor name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Attributes.
    Attributes PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes
}

func (sensor *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "sensors"
    sensor.EntityData.SegmentPath = "sensor" + types.AddKeyToken(sensor.Name, "name")
    sensor.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sensors/" + sensor.EntityData.SegmentPath
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Children.Append("attributes", types.YChild{"Attributes", &sensor.Attributes})
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("name", types.YLeaf{"Name", sensor.Name})

    sensor.EntityData.YListKeys = []string {"Name"}

    return &(sensor.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "sensor"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sensors/sensor/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sensors/sensor/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sensors/sensor/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sensors/sensor/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Sensors_Sensor_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/sensors/sensor/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "card"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Cards_Card_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/cards/card/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Slots_Slot_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "slot"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Slots_Slot_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/slots/slot/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

// PlatformInventory_Racks_Rack_Attributes
// Attributes
type PlatformInventory_Racks_Rack_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entity attributes.
    BasicInfo PlatformInventory_Racks_Rack_Attributes_BasicInfo

    // Field Replaceable Unit (FRU) attributes.
    FruInfo PlatformInventory_Racks_Rack_Attributes_FruInfo
}

func (attributes *PlatformInventory_Racks_Rack_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "rack"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &attributes.BasicInfo})
    attributes.EntityData.Children.Append("fru-info", types.YChild{"FruInfo", &attributes.FruInfo})
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// PlatformInventory_Racks_Rack_Attributes_BasicInfo
// Entity attributes
type PlatformInventory_Racks_Rack_Attributes_BasicInfo struct {
    EntityData types.CommonEntityData
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

func (basicInfo *PlatformInventory_Racks_Rack_Attributes_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "attributes"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/attributes/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", basicInfo.Name})
    basicInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", basicInfo.Description})
    basicInfo.EntityData.Leafs.Append("model-name", types.YLeaf{"ModelName", basicInfo.ModelName})
    basicInfo.EntityData.Leafs.Append("hardware-revision", types.YLeaf{"HardwareRevision", basicInfo.HardwareRevision})
    basicInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", basicInfo.SerialNumber})
    basicInfo.EntityData.Leafs.Append("firmware-revision", types.YLeaf{"FirmwareRevision", basicInfo.FirmwareRevision})
    basicInfo.EntityData.Leafs.Append("software-revision", types.YLeaf{"SoftwareRevision", basicInfo.SoftwareRevision})
    basicInfo.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", basicInfo.VendorType})
    basicInfo.EntityData.Leafs.Append("is-field-replaceable-unit", types.YLeaf{"IsFieldReplaceableUnit", basicInfo.IsFieldReplaceableUnit})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Attributes_FruInfo
// Field Replaceable Unit (FRU) attributes
type PlatformInventory_Racks_Rack_Attributes_FruInfo struct {
    EntityData types.CommonEntityData
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

func (fruInfo *PlatformInventory_Racks_Rack_Attributes_FruInfo) GetEntityData() *types.CommonEntityData {
    fruInfo.EntityData.YFilter = fruInfo.YFilter
    fruInfo.EntityData.YangName = "fru-info"
    fruInfo.EntityData.BundleName = "cisco_ios_xr"
    fruInfo.EntityData.ParentYangName = "attributes"
    fruInfo.EntityData.SegmentPath = "fru-info"
    fruInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/attributes/" + fruInfo.EntityData.SegmentPath
    fruInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruInfo.EntityData.Children = types.NewOrderedMap()
    fruInfo.EntityData.Children.Append("last-operational-state-change", types.YChild{"LastOperationalStateChange", &fruInfo.LastOperationalStateChange})
    fruInfo.EntityData.Children.Append("module-up-time", types.YChild{"ModuleUpTime", &fruInfo.ModuleUpTime})
    fruInfo.EntityData.Leafs = types.NewOrderedMap()
    fruInfo.EntityData.Leafs.Append("module-administrative-state", types.YLeaf{"ModuleAdministrativeState", fruInfo.ModuleAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-power-administrative-state", types.YLeaf{"ModulePowerAdministrativeState", fruInfo.ModulePowerAdministrativeState})
    fruInfo.EntityData.Leafs.Append("module-operational-state", types.YLeaf{"ModuleOperationalState", fruInfo.ModuleOperationalState})
    fruInfo.EntityData.Leafs.Append("module-monitor-state", types.YLeaf{"ModuleMonitorState", fruInfo.ModuleMonitorState})
    fruInfo.EntityData.Leafs.Append("module-reset-reason", types.YLeaf{"ModuleResetReason", fruInfo.ModuleResetReason})

    fruInfo.EntityData.YListKeys = []string {}

    return &(fruInfo.EntityData)
}

// PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange
// Time operational state is   last changed
type PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (lastOperationalStateChange *PlatformInventory_Racks_Rack_Attributes_FruInfo_LastOperationalStateChange) GetEntityData() *types.CommonEntityData {
    lastOperationalStateChange.EntityData.YFilter = lastOperationalStateChange.YFilter
    lastOperationalStateChange.EntityData.YangName = "last-operational-state-change"
    lastOperationalStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastOperationalStateChange.EntityData.ParentYangName = "fru-info"
    lastOperationalStateChange.EntityData.SegmentPath = "last-operational-state-change"
    lastOperationalStateChange.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/attributes/fru-info/" + lastOperationalStateChange.EntityData.SegmentPath
    lastOperationalStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastOperationalStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastOperationalStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastOperationalStateChange.EntityData.Children = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", lastOperationalStateChange.TimeInSeconds})
    lastOperationalStateChange.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", lastOperationalStateChange.TimeInNanoSeconds})

    lastOperationalStateChange.EntityData.YListKeys = []string {}

    return &(lastOperationalStateChange.EntityData)
}

// PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime
// Module up time
type PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time Value in Seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    TimeInSeconds interface{}

    // Time Value in Nano-seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are nanosecond.
    TimeInNanoSeconds interface{}
}

func (moduleUpTime *PlatformInventory_Racks_Rack_Attributes_FruInfo_ModuleUpTime) GetEntityData() *types.CommonEntityData {
    moduleUpTime.EntityData.YFilter = moduleUpTime.YFilter
    moduleUpTime.EntityData.YangName = "module-up-time"
    moduleUpTime.EntityData.BundleName = "cisco_ios_xr"
    moduleUpTime.EntityData.ParentYangName = "fru-info"
    moduleUpTime.EntityData.SegmentPath = "module-up-time"
    moduleUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-plat-chas-invmgr-ng-oper:platform-inventory/racks/rack/attributes/fru-info/" + moduleUpTime.EntityData.SegmentPath
    moduleUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    moduleUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    moduleUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    moduleUpTime.EntityData.Children = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs = types.NewOrderedMap()
    moduleUpTime.EntityData.Leafs.Append("time-in-seconds", types.YLeaf{"TimeInSeconds", moduleUpTime.TimeInSeconds})
    moduleUpTime.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", moduleUpTime.TimeInNanoSeconds})

    moduleUpTime.EntityData.YListKeys = []string {}

    return &(moduleUpTime.EntityData)
}

