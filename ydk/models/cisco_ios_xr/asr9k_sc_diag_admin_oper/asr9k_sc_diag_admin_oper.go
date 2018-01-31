// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-sc-diag package
// admin-plane operational data.
// 
// This module contains definitions
// for the following management objects:
//   diag: Diag admin operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_sc_diag_admin_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_sc_diag_admin_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-sc-diag-admin-oper diag}", reflect.TypeOf(Diag{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-sc-diag-admin-oper:diag", reflect.TypeOf(Diag{}))
}

// DiagProcessor represents Processor types
type DiagProcessor string

const (
    // Processor type 8614D
    DiagProcessor_mpc8614d DiagProcessor = "mpc8614d"
)

// DiagSlot represents Slot types
type DiagSlot string

const (
    // Slot type is fan tray
    DiagSlot_fan_tray DiagSlot = "fan-tray"

    // Slot type is power module
    DiagSlot_power_module DiagSlot = "power-module"

    // Slot type is module
    DiagSlot_module DiagSlot = "module"
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

// DiagNode represents Node types
type DiagNode string

const (
    // Node type is node
    DiagNode_node DiagNode = "node"

    // Node type is SPA
    DiagNode_spa DiagNode = "spa"
)

// Diag
// Diag admin operational data
type Diag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Diag operational data for available racks.
    Racks Diag_Racks
}

func (diag *Diag) GetFilter() yfilter.YFilter { return diag.YFilter }

func (diag *Diag) SetFilter(yf yfilter.YFilter) { diag.YFilter = yf }

func (diag *Diag) GetGoName(yname string) string {
    if yname == "racks" { return "Racks" }
    return ""
}

func (diag *Diag) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-sc-diag-admin-oper:diag"
}

func (diag *Diag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &diag.Racks
    }
    return nil
}

func (diag *Diag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &diag.Racks
    return children
}

func (diag *Diag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diag *Diag) GetBundleName() string { return "cisco_ios_xr" }

func (diag *Diag) GetYangName() string { return "diag" }

func (diag *Diag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diag *Diag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diag *Diag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diag *Diag) SetParent(parent types.Entity) { diag.parent = parent }

func (diag *Diag) GetParent() types.Entity { return diag.parent }

func (diag *Diag) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-sc-diag-admin-oper" }

// Diag_Racks
// Diag operational data for available racks
type Diag_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Diag operational data for a particular rack. The type is slice of
    // Diag_Racks_Rack.
    Rack []Diag_Racks_Rack
}

func (racks *Diag_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *Diag_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *Diag_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *Diag_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *Diag_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *Diag_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *Diag_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *Diag_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *Diag_Racks) GetYangName() string { return "racks" }

func (racks *Diag_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *Diag_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *Diag_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *Diag_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *Diag_Racks) GetParent() types.Entity { return racks.parent }

func (racks *Diag_Racks) GetParentYangName() string { return "diag" }

// Diag_Racks_Rack
// Diag operational data for a particular rack
type Diag_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack name. The type is interface{} with range:
    // -2147483648..2147483647.
    RackName interface{}

    // Fan tray table operational data.
    FanTraies Diag_Racks_Rack_FanTraies

    // Power supply table operational data.
    PowerSupplies Diag_Racks_Rack_PowerSupplies

    // Diag operational data for all available slots.
    Slots Diag_Racks_Rack_Slots

    // Chassis information.
    Chassis Diag_Racks_Rack_Chassis

    // Summary information.
    Summary Diag_Racks_Rack_Summary
}

func (rack *Diag_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *Diag_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *Diag_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack-name" { return "RackName" }
    if yname == "fan-traies" { return "FanTraies" }
    if yname == "power-supplies" { return "PowerSupplies" }
    if yname == "slots" { return "Slots" }
    if yname == "chassis" { return "Chassis" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (rack *Diag_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack-name='" + fmt.Sprintf("%v", rack.RackName) + "']"
}

func (rack *Diag_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fan-traies" {
        return &rack.FanTraies
    }
    if childYangName == "power-supplies" {
        return &rack.PowerSupplies
    }
    if childYangName == "slots" {
        return &rack.Slots
    }
    if childYangName == "chassis" {
        return &rack.Chassis
    }
    if childYangName == "summary" {
        return &rack.Summary
    }
    return nil
}

func (rack *Diag_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fan-traies"] = &rack.FanTraies
    children["power-supplies"] = &rack.PowerSupplies
    children["slots"] = &rack.Slots
    children["chassis"] = &rack.Chassis
    children["summary"] = &rack.Summary
    return children
}

func (rack *Diag_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-name"] = rack.RackName
    return leafs
}

func (rack *Diag_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *Diag_Racks_Rack) GetYangName() string { return "rack" }

func (rack *Diag_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *Diag_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *Diag_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *Diag_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *Diag_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *Diag_Racks_Rack) GetParentYangName() string { return "racks" }

// Diag_Racks_Rack_FanTraies
// Fan tray table operational data
type Diag_Racks_Rack_FanTraies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fan tray operational data. The type is slice of
    // Diag_Racks_Rack_FanTraies_FanTray.
    FanTray []Diag_Racks_Rack_FanTraies_FanTray
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetFilter() yfilter.YFilter { return fanTraies.YFilter }

func (fanTraies *Diag_Racks_Rack_FanTraies) SetFilter(yf yfilter.YFilter) { fanTraies.YFilter = yf }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetGoName(yname string) string {
    if yname == "fan-tray" { return "FanTray" }
    return ""
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetSegmentPath() string {
    return "fan-traies"
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fan-tray" {
        for _, c := range fanTraies.FanTray {
            if fanTraies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_FanTraies_FanTray{}
        fanTraies.FanTray = append(fanTraies.FanTray, child)
        return &fanTraies.FanTray[len(fanTraies.FanTray)-1]
    }
    return nil
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fanTraies.FanTray {
        children[fanTraies.FanTray[i].GetSegmentPath()] = &fanTraies.FanTray[i]
    }
    return children
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fanTraies *Diag_Racks_Rack_FanTraies) GetBundleName() string { return "cisco_ios_xr" }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetYangName() string { return "fan-traies" }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fanTraies *Diag_Racks_Rack_FanTraies) SetParent(parent types.Entity) { fanTraies.parent = parent }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetParent() types.Entity { return fanTraies.parent }

func (fanTraies *Diag_Racks_Rack_FanTraies) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_FanTraies_FanTray
// Fan tray operational data
type Diag_Racks_Rack_FanTraies_FanTray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fan tray name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    FanTrayName interface{}

    // Diag detailed information.
    Detail Diag_Racks_Rack_FanTraies_FanTray_Detail
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetFilter() yfilter.YFilter { return fanTray.YFilter }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) SetFilter(yf yfilter.YFilter) { fanTray.YFilter = yf }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetGoName(yname string) string {
    if yname == "fan-tray-name" { return "FanTrayName" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetSegmentPath() string {
    return "fan-tray" + "[fan-tray-name='" + fmt.Sprintf("%v", fanTray.FanTrayName) + "']"
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &fanTray.Detail
    }
    return nil
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &fanTray.Detail
    return children
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fan-tray-name"] = fanTray.FanTrayName
    return leafs
}

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetBundleName() string { return "cisco_ios_xr" }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetYangName() string { return "fan-tray" }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) SetParent(parent types.Entity) { fanTray.parent = parent }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetParent() types.Entity { return fanTray.parent }

func (fanTray *Diag_Racks_Rack_FanTraies_FanTray) GetParentYangName() string { return "fan-traies" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail
// Diag detailed information
type Diag_Racks_Rack_FanTraies_FanTray_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node information.
    Node Diag_Racks_Rack_FanTraies_FanTray_Detail_Node

    // SPA information.
    Spa Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa
}

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "spa" { return "Spa" }
    return ""
}

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &detail.Node
    }
    if childYangName == "spa" {
        return &detail.Spa
    }
    return nil
}

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &detail.Node
    children["spa"] = &detail.Spa
    return children
}

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetYangName() string { return "detail" }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Diag_Racks_Rack_FanTraies_FanTray_Detail) GetParentYangName() string { return "fan-tray" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node
// Node information
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describes in user-readable terms. The type is string.
    Description interface{}

    // Main serial number. The type is string.
    SerialNumber interface{}

    // Top assembly number. The type is string.
    Tan interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // Chip hardware revision. The type is string.
    ChipHardwareRevision interface{}

    // New deviation number. The type is interface{} with range: 0..4294967295.
    NewDeviationNumber interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Module operational state. The type is NodeState.
    BoardState interface{}

    // Motherboard PLD version. The type is string.
    PldMotherboard interface{}

    // Power PLD version. The type is string.
    PldPower interface{}

    // MONLIB version. The type is string.
    Monlib interface{}

    // ROMMON version. The type is string.
    Rommon interface{}

    // Processor type. The type is DiagProcessor.
    Cpu0 interface{}

    // Programmable logic device information.
    Pld Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld

    // Hardware revision.
    HardwareRevision Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision

    // CBC active partition.
    CbcActivePartition Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition

    // CBC inactive partition.
    CbcInactivePartition Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition
}

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "tan" { return "Tan" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "clei" { return "Clei" }
    if yname == "board-state" { return "BoardState" }
    if yname == "pld-motherboard" { return "PldMotherboard" }
    if yname == "pld-power" { return "PldPower" }
    if yname == "monlib" { return "Monlib" }
    if yname == "rommon" { return "Rommon" }
    if yname == "cpu0" { return "Cpu0" }
    if yname == "pld" { return "Pld" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "cbc-active-partition" { return "CbcActivePartition" }
    if yname == "cbc-inactive-partition" { return "CbcInactivePartition" }
    return ""
}

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetSegmentPath() string {
    return "node"
}

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pld" {
        return &node.Pld
    }
    if childYangName == "hardware-revision" {
        return &node.HardwareRevision
    }
    if childYangName == "cbc-active-partition" {
        return &node.CbcActivePartition
    }
    if childYangName == "cbc-inactive-partition" {
        return &node.CbcInactivePartition
    }
    return nil
}

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pld"] = &node.Pld
    children["hardware-revision"] = &node.HardwareRevision
    children["cbc-active-partition"] = &node.CbcActivePartition
    children["cbc-inactive-partition"] = &node.CbcInactivePartition
    return children
}

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = node.Description
    leafs["serial-number"] = node.SerialNumber
    leafs["tan"] = node.Tan
    leafs["pid"] = node.Pid
    leafs["vid"] = node.Vid
    leafs["chip-hardware-revision"] = node.ChipHardwareRevision
    leafs["new-deviation-number"] = node.NewDeviationNumber
    leafs["clei"] = node.Clei
    leafs["board-state"] = node.BoardState
    leafs["pld-motherboard"] = node.PldMotherboard
    leafs["pld-power"] = node.PldPower
    leafs["monlib"] = node.Monlib
    leafs["rommon"] = node.Rommon
    leafs["cpu0"] = node.Cpu0
    return leafs
}

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetYangName() string { return "node" }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetParent() types.Entity { return node.parent }

func (node *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld
// Programmable logic device information
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Processor PLD version. The type is interface{} with range: 0..4294967295.
    Type interface{}

    // HigherVerion. The type is interface{} with range: 0..4294967295.
    ProcessorHigherVersion interface{}

    // LowerVersion. The type is interface{} with range: 0..4294967295.
    ProcessorLowerVersion interface{}
}

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetFilter() yfilter.YFilter { return pld.YFilter }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) SetFilter(yf yfilter.YFilter) { pld.YFilter = yf }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "processor-higher-version" { return "ProcessorHigherVersion" }
    if yname == "processor-lower-version" { return "ProcessorLowerVersion" }
    return ""
}

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetSegmentPath() string {
    return "pld"
}

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = pld.Type
    leafs["processor-higher-version"] = pld.ProcessorHigherVersion
    leafs["processor-lower-version"] = pld.ProcessorLowerVersion
    return leafs
}

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetBundleName() string { return "cisco_ios_xr" }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetYangName() string { return "pld" }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) SetParent(parent types.Entity) { pld.parent = parent }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetParent() types.Entity { return pld.parent }

func (pld *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_Pld) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision.
    HardwareRevision []Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetGoName(yname string) string {
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-revision" {
        for _, c := range hardwareRevision.HardwareRevision {
            if hardwareRevision.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision{}
        hardwareRevision.HardwareRevision = append(hardwareRevision.HardwareRevision, child)
        return &hardwareRevision.HardwareRevision[len(hardwareRevision.HardwareRevision)-1]
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareRevision.HardwareRevision {
        children[hardwareRevision.HardwareRevision[i].GetSegmentPath()] = &hardwareRevision.HardwareRevision[i]
    }
    return children
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node decsription. The type is string.
    NodeDescription interface{}

    // Version information. The type is string.
    Version interface{}

    // Hardware version.
    HwRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev

    // Firmware version.
    FwRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev

    // Software version.
    SwRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev

    // DIMM version information.
    DimmRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev

    // SSD version information.
    SsdRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetGoName(yname string) string {
    if yname == "node-description" { return "NodeDescription" }
    if yname == "version" { return "Version" }
    if yname == "hw-rev" { return "HwRev" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "sw-rev" { return "SwRev" }
    if yname == "dimm-rev" { return "DimmRev" }
    if yname == "ssd-rev" { return "SsdRev" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-rev" {
        return &hardwareRevision.HwRev
    }
    if childYangName == "fw-rev" {
        return &hardwareRevision.FwRev
    }
    if childYangName == "sw-rev" {
        return &hardwareRevision.SwRev
    }
    if childYangName == "dimm-rev" {
        return &hardwareRevision.DimmRev
    }
    if childYangName == "ssd-rev" {
        return &hardwareRevision.SsdRev
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-rev"] = &hardwareRevision.HwRev
    children["fw-rev"] = &hardwareRevision.FwRev
    children["sw-rev"] = &hardwareRevision.SwRev
    children["dimm-rev"] = &hardwareRevision.DimmRev
    children["ssd-rev"] = &hardwareRevision.SsdRev
    return children
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-description"] = hardwareRevision.NodeDescription
    leafs["version"] = hardwareRevision.Version
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetFilter() yfilter.YFilter { return hwRev.YFilter }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) SetFilter(yf yfilter.YFilter) { hwRev.YFilter = yf }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetSegmentPath() string {
    return "hw-rev"
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = hwRev.MajorRevision
    leafs["minor-revision"] = hwRev.MinorRevision
    return leafs
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetBundleName() string { return "cisco_ios_xr" }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetYangName() string { return "hw-rev" }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) SetParent(parent types.Entity) { hwRev.parent = parent }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetParent() types.Entity { return hwRev.parent }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetFilter() yfilter.YFilter { return fwRev.YFilter }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) SetFilter(yf yfilter.YFilter) { fwRev.YFilter = yf }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetSegmentPath() string {
    return "fw-rev"
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = fwRev.MajorRevision
    leafs["minor-revision"] = fwRev.MinorRevision
    return leafs
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetBundleName() string { return "cisco_ios_xr" }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetYangName() string { return "fw-rev" }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) SetParent(parent types.Entity) { fwRev.parent = parent }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetParent() types.Entity { return fwRev.parent }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetFilter() yfilter.YFilter { return swRev.YFilter }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) SetFilter(yf yfilter.YFilter) { swRev.YFilter = yf }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetSegmentPath() string {
    return "sw-rev"
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = swRev.MajorRevision
    leafs["minor-revision"] = swRev.MinorRevision
    return leafs
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetBundleName() string { return "cisco_ios_xr" }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetYangName() string { return "sw-rev" }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) SetParent(parent types.Entity) { swRev.parent = parent }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetParent() types.Entity { return swRev.parent }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size in MB. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Speed in MHz. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Locator information. The type is string.
    Locator interface{}

    // Column address strobe latency in clock cycles. The type is interface{} with
    // range: 0..4294967295.
    Cas interface{}
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetFilter() yfilter.YFilter { return dimmRev.YFilter }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) SetFilter(yf yfilter.YFilter) { dimmRev.YFilter = yf }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "speed" { return "Speed" }
    if yname == "locator" { return "Locator" }
    if yname == "cas" { return "Cas" }
    return ""
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetSegmentPath() string {
    return "dimm-rev"
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = dimmRev.Size
    leafs["speed"] = dimmRev.Speed
    leafs["locator"] = dimmRev.Locator
    leafs["cas"] = dimmRev.Cas
    return leafs
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetBundleName() string { return "cisco_ios_xr" }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetYangName() string { return "dimm-rev" }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) SetParent(parent types.Entity) { dimmRev.parent = parent }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetParent() types.Entity { return dimmRev.parent }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetFilter() yfilter.YFilter { return ssdRev.YFilter }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) SetFilter(yf yfilter.YFilter) { ssdRev.YFilter = yf }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetSegmentPath() string {
    return "ssd-rev"
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ssdRev.Number
    leafs["fw-rev"] = ssdRev.FwRev
    leafs["serial-number"] = ssdRev.SerialNumber
    return leafs
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetBundleName() string { return "cisco_ios_xr" }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetYangName() string { return "ssd-rev" }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) SetParent(parent types.Entity) { ssdRev.parent = parent }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetParent() types.Entity { return ssdRev.parent }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition
// CBC active partition
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetFilter() yfilter.YFilter { return cbcActivePartition.YFilter }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) SetFilter(yf yfilter.YFilter) { cbcActivePartition.YFilter = yf }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetSegmentPath() string {
    return "cbc-active-partition"
}

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = cbcActivePartition.MajorRevision
    leafs["minor-revision"] = cbcActivePartition.MinorRevision
    return leafs
}

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetBundleName() string { return "cisco_ios_xr" }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetYangName() string { return "cbc-active-partition" }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) SetParent(parent types.Entity) { cbcActivePartition.parent = parent }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetParent() types.Entity { return cbcActivePartition.parent }

func (cbcActivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcActivePartition) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition
// CBC inactive partition
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetFilter() yfilter.YFilter { return cbcInactivePartition.YFilter }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) SetFilter(yf yfilter.YFilter) { cbcInactivePartition.YFilter = yf }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetSegmentPath() string {
    return "cbc-inactive-partition"
}

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = cbcInactivePartition.MajorRevision
    leafs["minor-revision"] = cbcInactivePartition.MinorRevision
    return leafs
}

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetBundleName() string { return "cisco_ios_xr" }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetYangName() string { return "cbc-inactive-partition" }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) SetParent(parent types.Entity) { cbcInactivePartition.parent = parent }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetParent() types.Entity { return cbcInactivePartition.parent }

func (cbcInactivePartition *Diag_Racks_Rack_FanTraies_FanTray_Detail_Node_CbcInactivePartition) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa
// SPA information
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // SPA name. The type is string.
    Name interface{}

    // Format is C-BN-V where C = Class Code, BN = Base Number, and V = Version.
    // The type is string.
    PcaUnitNumber interface{}

    // PCA revision information. The type is string.
    PcaRevision interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Node state. The type is NodeState.
    NodeState interface{}

    // Main.
    Main Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main

    // Hardware revision.
    HardwareRevision Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision
}

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetFilter() yfilter.YFilter { return spa.YFilter }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) SetFilter(yf yfilter.YFilter) { spa.YFilter = yf }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "name" { return "Name" }
    if yname == "pca-unit-number" { return "PcaUnitNumber" }
    if yname == "pca-revision" { return "PcaRevision" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "clei" { return "Clei" }
    if yname == "node-state" { return "NodeState" }
    if yname == "main" { return "Main" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetSegmentPath() string {
    return "spa"
}

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "main" {
        return &spa.Main
    }
    if childYangName == "hardware-revision" {
        return &spa.HardwareRevision
    }
    return nil
}

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["main"] = &spa.Main
    children["hardware-revision"] = &spa.HardwareRevision
    return children
}

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = spa.Node
    leafs["name"] = spa.Name
    leafs["pca-unit-number"] = spa.PcaUnitNumber
    leafs["pca-revision"] = spa.PcaRevision
    leafs["pid"] = spa.Pid
    leafs["vid"] = spa.Vid
    leafs["clei"] = spa.Clei
    leafs["node-state"] = spa.NodeState
    return leafs
}

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetBundleName() string { return "cisco_ios_xr" }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetYangName() string { return "spa" }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) SetParent(parent types.Entity) { spa.parent = parent }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetParent() types.Entity { return spa.parent }

func (spa *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main
// Main
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board type. The type is interface{} with range: 0..4294967295.
    BoardType interface{}

    // Top assembly number. The type is string.
    Tan interface{}

    // Top assembly number revision. The type is string.
    TanRevision interface{}

    // Deviation number. The type is interface{} with range: 0..4294967295.
    DeviationNumber interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetFilter() yfilter.YFilter { return main.YFilter }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) SetFilter(yf yfilter.YFilter) { main.YFilter = yf }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetGoName(yname string) string {
    if yname == "board-type" { return "BoardType" }
    if yname == "tan" { return "Tan" }
    if yname == "tan-revision" { return "TanRevision" }
    if yname == "deviation-number" { return "DeviationNumber" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetSegmentPath() string {
    return "main"
}

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["board-type"] = main.BoardType
    leafs["tan"] = main.Tan
    leafs["tan-revision"] = main.TanRevision
    leafs["deviation-number"] = main.DeviationNumber
    leafs["serial-number"] = main.SerialNumber
    return leafs
}

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetBundleName() string { return "cisco_ios_xr" }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetYangName() string { return "main" }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) SetParent(parent types.Entity) { main.parent = parent }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetParent() types.Entity { return main.parent }

func (main *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_Main) GetParentYangName() string { return "spa" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision.
    HardwareRevision []Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetGoName(yname string) string {
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-revision" {
        for _, c := range hardwareRevision.HardwareRevision {
            if hardwareRevision.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision{}
        hardwareRevision.HardwareRevision = append(hardwareRevision.HardwareRevision, child)
        return &hardwareRevision.HardwareRevision[len(hardwareRevision.HardwareRevision)-1]
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareRevision.HardwareRevision {
        children[hardwareRevision.HardwareRevision[i].GetSegmentPath()] = &hardwareRevision.HardwareRevision[i]
    }
    return children
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision) GetParentYangName() string { return "spa" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node decsription. The type is string.
    NodeDescription interface{}

    // Version information. The type is string.
    Version interface{}

    // Hardware version.
    HwRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev

    // Firmware version.
    FwRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev

    // Software version.
    SwRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev

    // DIMM version information.
    DimmRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev

    // SSD version information.
    SsdRev Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetGoName(yname string) string {
    if yname == "node-description" { return "NodeDescription" }
    if yname == "version" { return "Version" }
    if yname == "hw-rev" { return "HwRev" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "sw-rev" { return "SwRev" }
    if yname == "dimm-rev" { return "DimmRev" }
    if yname == "ssd-rev" { return "SsdRev" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-rev" {
        return &hardwareRevision.HwRev
    }
    if childYangName == "fw-rev" {
        return &hardwareRevision.FwRev
    }
    if childYangName == "sw-rev" {
        return &hardwareRevision.SwRev
    }
    if childYangName == "dimm-rev" {
        return &hardwareRevision.DimmRev
    }
    if childYangName == "ssd-rev" {
        return &hardwareRevision.SsdRev
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-rev"] = &hardwareRevision.HwRev
    children["fw-rev"] = &hardwareRevision.FwRev
    children["sw-rev"] = &hardwareRevision.SwRev
    children["dimm-rev"] = &hardwareRevision.DimmRev
    children["ssd-rev"] = &hardwareRevision.SsdRev
    return children
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-description"] = hardwareRevision.NodeDescription
    leafs["version"] = hardwareRevision.Version
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetFilter() yfilter.YFilter { return hwRev.YFilter }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) SetFilter(yf yfilter.YFilter) { hwRev.YFilter = yf }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetSegmentPath() string {
    return "hw-rev"
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = hwRev.MajorRevision
    leafs["minor-revision"] = hwRev.MinorRevision
    return leafs
}

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetBundleName() string { return "cisco_ios_xr" }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetYangName() string { return "hw-rev" }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) SetParent(parent types.Entity) { hwRev.parent = parent }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetParent() types.Entity { return hwRev.parent }

func (hwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetFilter() yfilter.YFilter { return fwRev.YFilter }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) SetFilter(yf yfilter.YFilter) { fwRev.YFilter = yf }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetSegmentPath() string {
    return "fw-rev"
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = fwRev.MajorRevision
    leafs["minor-revision"] = fwRev.MinorRevision
    return leafs
}

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetBundleName() string { return "cisco_ios_xr" }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetYangName() string { return "fw-rev" }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) SetParent(parent types.Entity) { fwRev.parent = parent }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetParent() types.Entity { return fwRev.parent }

func (fwRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetFilter() yfilter.YFilter { return swRev.YFilter }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) SetFilter(yf yfilter.YFilter) { swRev.YFilter = yf }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetSegmentPath() string {
    return "sw-rev"
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = swRev.MajorRevision
    leafs["minor-revision"] = swRev.MinorRevision
    return leafs
}

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetBundleName() string { return "cisco_ios_xr" }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetYangName() string { return "sw-rev" }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) SetParent(parent types.Entity) { swRev.parent = parent }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetParent() types.Entity { return swRev.parent }

func (swRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size in MB. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Speed in MHz. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Locator information. The type is string.
    Locator interface{}

    // Column address strobe latency in clock cycles. The type is interface{} with
    // range: 0..4294967295.
    Cas interface{}
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetFilter() yfilter.YFilter { return dimmRev.YFilter }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) SetFilter(yf yfilter.YFilter) { dimmRev.YFilter = yf }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "speed" { return "Speed" }
    if yname == "locator" { return "Locator" }
    if yname == "cas" { return "Cas" }
    return ""
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetSegmentPath() string {
    return "dimm-rev"
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = dimmRev.Size
    leafs["speed"] = dimmRev.Speed
    leafs["locator"] = dimmRev.Locator
    leafs["cas"] = dimmRev.Cas
    return leafs
}

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetBundleName() string { return "cisco_ios_xr" }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetYangName() string { return "dimm-rev" }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) SetParent(parent types.Entity) { dimmRev.parent = parent }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetParent() types.Entity { return dimmRev.parent }

func (dimmRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetFilter() yfilter.YFilter { return ssdRev.YFilter }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) SetFilter(yf yfilter.YFilter) { ssdRev.YFilter = yf }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetSegmentPath() string {
    return "ssd-rev"
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ssdRev.Number
    leafs["fw-rev"] = ssdRev.FwRev
    leafs["serial-number"] = ssdRev.SerialNumber
    return leafs
}

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetBundleName() string { return "cisco_ios_xr" }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetYangName() string { return "ssd-rev" }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) SetParent(parent types.Entity) { ssdRev.parent = parent }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetParent() types.Entity { return ssdRev.parent }

func (ssdRev *Diag_Racks_Rack_FanTraies_FanTray_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies
// Power supply table operational data
type Diag_Racks_Rack_PowerSupplies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Power supply operational data. The type is slice of
    // Diag_Racks_Rack_PowerSupplies_PowerSupply.
    PowerSupply []Diag_Racks_Rack_PowerSupplies_PowerSupply
}

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetFilter() yfilter.YFilter { return powerSupplies.YFilter }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) SetFilter(yf yfilter.YFilter) { powerSupplies.YFilter = yf }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetGoName(yname string) string {
    if yname == "power-supply" { return "PowerSupply" }
    return ""
}

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetSegmentPath() string {
    return "power-supplies"
}

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "power-supply" {
        for _, c := range powerSupplies.PowerSupply {
            if powerSupplies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_PowerSupplies_PowerSupply{}
        powerSupplies.PowerSupply = append(powerSupplies.PowerSupply, child)
        return &powerSupplies.PowerSupply[len(powerSupplies.PowerSupply)-1]
    }
    return nil
}

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range powerSupplies.PowerSupply {
        children[powerSupplies.PowerSupply[i].GetSegmentPath()] = &powerSupplies.PowerSupply[i]
    }
    return children
}

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetYangName() string { return "power-supplies" }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) SetParent(parent types.Entity) { powerSupplies.parent = parent }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetParent() types.Entity { return powerSupplies.parent }

func (powerSupplies *Diag_Racks_Rack_PowerSupplies) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply
// Power supply operational data
type Diag_Racks_Rack_PowerSupplies_PowerSupply struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Power supply name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PowerSupplyName interface{}

    // Diag detailed information.
    Detail Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail
}

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetFilter() yfilter.YFilter { return powerSupply.YFilter }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) SetFilter(yf yfilter.YFilter) { powerSupply.YFilter = yf }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetGoName(yname string) string {
    if yname == "power-supply-name" { return "PowerSupplyName" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetSegmentPath() string {
    return "power-supply" + "[power-supply-name='" + fmt.Sprintf("%v", powerSupply.PowerSupplyName) + "']"
}

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &powerSupply.Detail
    }
    return nil
}

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &powerSupply.Detail
    return children
}

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["power-supply-name"] = powerSupply.PowerSupplyName
    return leafs
}

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetBundleName() string { return "cisco_ios_xr" }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetYangName() string { return "power-supply" }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) SetParent(parent types.Entity) { powerSupply.parent = parent }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetParent() types.Entity { return powerSupply.parent }

func (powerSupply *Diag_Racks_Rack_PowerSupplies_PowerSupply) GetParentYangName() string { return "power-supplies" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail
// Diag detailed information
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node information.
    Node Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node

    // SPA information.
    Spa Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa
}

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "spa" { return "Spa" }
    return ""
}

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &detail.Node
    }
    if childYangName == "spa" {
        return &detail.Spa
    }
    return nil
}

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &detail.Node
    children["spa"] = &detail.Spa
    return children
}

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetYangName() string { return "detail" }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail) GetParentYangName() string { return "power-supply" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node
// Node information
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describes in user-readable terms. The type is string.
    Description interface{}

    // Main serial number. The type is string.
    SerialNumber interface{}

    // Top assembly number. The type is string.
    Tan interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // Chip hardware revision. The type is string.
    ChipHardwareRevision interface{}

    // New deviation number. The type is interface{} with range: 0..4294967295.
    NewDeviationNumber interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Module operational state. The type is NodeState.
    BoardState interface{}

    // Motherboard PLD version. The type is string.
    PldMotherboard interface{}

    // Power PLD version. The type is string.
    PldPower interface{}

    // MONLIB version. The type is string.
    Monlib interface{}

    // ROMMON version. The type is string.
    Rommon interface{}

    // Processor type. The type is DiagProcessor.
    Cpu0 interface{}

    // Programmable logic device information.
    Pld Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld

    // Hardware revision.
    HardwareRevision Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision

    // CBC active partition.
    CbcActivePartition Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition

    // CBC inactive partition.
    CbcInactivePartition Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition
}

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "tan" { return "Tan" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "clei" { return "Clei" }
    if yname == "board-state" { return "BoardState" }
    if yname == "pld-motherboard" { return "PldMotherboard" }
    if yname == "pld-power" { return "PldPower" }
    if yname == "monlib" { return "Monlib" }
    if yname == "rommon" { return "Rommon" }
    if yname == "cpu0" { return "Cpu0" }
    if yname == "pld" { return "Pld" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "cbc-active-partition" { return "CbcActivePartition" }
    if yname == "cbc-inactive-partition" { return "CbcInactivePartition" }
    return ""
}

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetSegmentPath() string {
    return "node"
}

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pld" {
        return &node.Pld
    }
    if childYangName == "hardware-revision" {
        return &node.HardwareRevision
    }
    if childYangName == "cbc-active-partition" {
        return &node.CbcActivePartition
    }
    if childYangName == "cbc-inactive-partition" {
        return &node.CbcInactivePartition
    }
    return nil
}

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pld"] = &node.Pld
    children["hardware-revision"] = &node.HardwareRevision
    children["cbc-active-partition"] = &node.CbcActivePartition
    children["cbc-inactive-partition"] = &node.CbcInactivePartition
    return children
}

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = node.Description
    leafs["serial-number"] = node.SerialNumber
    leafs["tan"] = node.Tan
    leafs["pid"] = node.Pid
    leafs["vid"] = node.Vid
    leafs["chip-hardware-revision"] = node.ChipHardwareRevision
    leafs["new-deviation-number"] = node.NewDeviationNumber
    leafs["clei"] = node.Clei
    leafs["board-state"] = node.BoardState
    leafs["pld-motherboard"] = node.PldMotherboard
    leafs["pld-power"] = node.PldPower
    leafs["monlib"] = node.Monlib
    leafs["rommon"] = node.Rommon
    leafs["cpu0"] = node.Cpu0
    return leafs
}

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetYangName() string { return "node" }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetParent() types.Entity { return node.parent }

func (node *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld
// Programmable logic device information
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Processor PLD version. The type is interface{} with range: 0..4294967295.
    Type interface{}

    // HigherVerion. The type is interface{} with range: 0..4294967295.
    ProcessorHigherVersion interface{}

    // LowerVersion. The type is interface{} with range: 0..4294967295.
    ProcessorLowerVersion interface{}
}

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetFilter() yfilter.YFilter { return pld.YFilter }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) SetFilter(yf yfilter.YFilter) { pld.YFilter = yf }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "processor-higher-version" { return "ProcessorHigherVersion" }
    if yname == "processor-lower-version" { return "ProcessorLowerVersion" }
    return ""
}

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetSegmentPath() string {
    return "pld"
}

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = pld.Type
    leafs["processor-higher-version"] = pld.ProcessorHigherVersion
    leafs["processor-lower-version"] = pld.ProcessorLowerVersion
    return leafs
}

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetBundleName() string { return "cisco_ios_xr" }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetYangName() string { return "pld" }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) SetParent(parent types.Entity) { pld.parent = parent }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetParent() types.Entity { return pld.parent }

func (pld *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_Pld) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision.
    HardwareRevision []Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetGoName(yname string) string {
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-revision" {
        for _, c := range hardwareRevision.HardwareRevision {
            if hardwareRevision.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision{}
        hardwareRevision.HardwareRevision = append(hardwareRevision.HardwareRevision, child)
        return &hardwareRevision.HardwareRevision[len(hardwareRevision.HardwareRevision)-1]
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareRevision.HardwareRevision {
        children[hardwareRevision.HardwareRevision[i].GetSegmentPath()] = &hardwareRevision.HardwareRevision[i]
    }
    return children
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node decsription. The type is string.
    NodeDescription interface{}

    // Version information. The type is string.
    Version interface{}

    // Hardware version.
    HwRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev

    // Firmware version.
    FwRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev

    // Software version.
    SwRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev

    // DIMM version information.
    DimmRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev

    // SSD version information.
    SsdRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetGoName(yname string) string {
    if yname == "node-description" { return "NodeDescription" }
    if yname == "version" { return "Version" }
    if yname == "hw-rev" { return "HwRev" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "sw-rev" { return "SwRev" }
    if yname == "dimm-rev" { return "DimmRev" }
    if yname == "ssd-rev" { return "SsdRev" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-rev" {
        return &hardwareRevision.HwRev
    }
    if childYangName == "fw-rev" {
        return &hardwareRevision.FwRev
    }
    if childYangName == "sw-rev" {
        return &hardwareRevision.SwRev
    }
    if childYangName == "dimm-rev" {
        return &hardwareRevision.DimmRev
    }
    if childYangName == "ssd-rev" {
        return &hardwareRevision.SsdRev
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-rev"] = &hardwareRevision.HwRev
    children["fw-rev"] = &hardwareRevision.FwRev
    children["sw-rev"] = &hardwareRevision.SwRev
    children["dimm-rev"] = &hardwareRevision.DimmRev
    children["ssd-rev"] = &hardwareRevision.SsdRev
    return children
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-description"] = hardwareRevision.NodeDescription
    leafs["version"] = hardwareRevision.Version
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetFilter() yfilter.YFilter { return hwRev.YFilter }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) SetFilter(yf yfilter.YFilter) { hwRev.YFilter = yf }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetSegmentPath() string {
    return "hw-rev"
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = hwRev.MajorRevision
    leafs["minor-revision"] = hwRev.MinorRevision
    return leafs
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetBundleName() string { return "cisco_ios_xr" }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetYangName() string { return "hw-rev" }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) SetParent(parent types.Entity) { hwRev.parent = parent }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetParent() types.Entity { return hwRev.parent }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetFilter() yfilter.YFilter { return fwRev.YFilter }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) SetFilter(yf yfilter.YFilter) { fwRev.YFilter = yf }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetSegmentPath() string {
    return "fw-rev"
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = fwRev.MajorRevision
    leafs["minor-revision"] = fwRev.MinorRevision
    return leafs
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetBundleName() string { return "cisco_ios_xr" }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetYangName() string { return "fw-rev" }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) SetParent(parent types.Entity) { fwRev.parent = parent }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetParent() types.Entity { return fwRev.parent }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetFilter() yfilter.YFilter { return swRev.YFilter }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) SetFilter(yf yfilter.YFilter) { swRev.YFilter = yf }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetSegmentPath() string {
    return "sw-rev"
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = swRev.MajorRevision
    leafs["minor-revision"] = swRev.MinorRevision
    return leafs
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetBundleName() string { return "cisco_ios_xr" }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetYangName() string { return "sw-rev" }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) SetParent(parent types.Entity) { swRev.parent = parent }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetParent() types.Entity { return swRev.parent }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size in MB. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Speed in MHz. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Locator information. The type is string.
    Locator interface{}

    // Column address strobe latency in clock cycles. The type is interface{} with
    // range: 0..4294967295.
    Cas interface{}
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetFilter() yfilter.YFilter { return dimmRev.YFilter }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) SetFilter(yf yfilter.YFilter) { dimmRev.YFilter = yf }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "speed" { return "Speed" }
    if yname == "locator" { return "Locator" }
    if yname == "cas" { return "Cas" }
    return ""
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetSegmentPath() string {
    return "dimm-rev"
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = dimmRev.Size
    leafs["speed"] = dimmRev.Speed
    leafs["locator"] = dimmRev.Locator
    leafs["cas"] = dimmRev.Cas
    return leafs
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetBundleName() string { return "cisco_ios_xr" }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetYangName() string { return "dimm-rev" }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) SetParent(parent types.Entity) { dimmRev.parent = parent }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetParent() types.Entity { return dimmRev.parent }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetFilter() yfilter.YFilter { return ssdRev.YFilter }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) SetFilter(yf yfilter.YFilter) { ssdRev.YFilter = yf }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetSegmentPath() string {
    return "ssd-rev"
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ssdRev.Number
    leafs["fw-rev"] = ssdRev.FwRev
    leafs["serial-number"] = ssdRev.SerialNumber
    return leafs
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetBundleName() string { return "cisco_ios_xr" }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetYangName() string { return "ssd-rev" }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) SetParent(parent types.Entity) { ssdRev.parent = parent }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetParent() types.Entity { return ssdRev.parent }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition
// CBC active partition
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetFilter() yfilter.YFilter { return cbcActivePartition.YFilter }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) SetFilter(yf yfilter.YFilter) { cbcActivePartition.YFilter = yf }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetSegmentPath() string {
    return "cbc-active-partition"
}

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = cbcActivePartition.MajorRevision
    leafs["minor-revision"] = cbcActivePartition.MinorRevision
    return leafs
}

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetBundleName() string { return "cisco_ios_xr" }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetYangName() string { return "cbc-active-partition" }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) SetParent(parent types.Entity) { cbcActivePartition.parent = parent }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetParent() types.Entity { return cbcActivePartition.parent }

func (cbcActivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcActivePartition) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition
// CBC inactive partition
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetFilter() yfilter.YFilter { return cbcInactivePartition.YFilter }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) SetFilter(yf yfilter.YFilter) { cbcInactivePartition.YFilter = yf }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetSegmentPath() string {
    return "cbc-inactive-partition"
}

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = cbcInactivePartition.MajorRevision
    leafs["minor-revision"] = cbcInactivePartition.MinorRevision
    return leafs
}

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetBundleName() string { return "cisco_ios_xr" }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetYangName() string { return "cbc-inactive-partition" }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) SetParent(parent types.Entity) { cbcInactivePartition.parent = parent }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetParent() types.Entity { return cbcInactivePartition.parent }

func (cbcInactivePartition *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Node_CbcInactivePartition) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa
// SPA information
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // SPA name. The type is string.
    Name interface{}

    // Format is C-BN-V where C = Class Code, BN = Base Number, and V = Version.
    // The type is string.
    PcaUnitNumber interface{}

    // PCA revision information. The type is string.
    PcaRevision interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Node state. The type is NodeState.
    NodeState interface{}

    // Main.
    Main Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main

    // Hardware revision.
    HardwareRevision Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision
}

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetFilter() yfilter.YFilter { return spa.YFilter }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) SetFilter(yf yfilter.YFilter) { spa.YFilter = yf }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "name" { return "Name" }
    if yname == "pca-unit-number" { return "PcaUnitNumber" }
    if yname == "pca-revision" { return "PcaRevision" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "clei" { return "Clei" }
    if yname == "node-state" { return "NodeState" }
    if yname == "main" { return "Main" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetSegmentPath() string {
    return "spa"
}

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "main" {
        return &spa.Main
    }
    if childYangName == "hardware-revision" {
        return &spa.HardwareRevision
    }
    return nil
}

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["main"] = &spa.Main
    children["hardware-revision"] = &spa.HardwareRevision
    return children
}

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = spa.Node
    leafs["name"] = spa.Name
    leafs["pca-unit-number"] = spa.PcaUnitNumber
    leafs["pca-revision"] = spa.PcaRevision
    leafs["pid"] = spa.Pid
    leafs["vid"] = spa.Vid
    leafs["clei"] = spa.Clei
    leafs["node-state"] = spa.NodeState
    return leafs
}

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetBundleName() string { return "cisco_ios_xr" }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetYangName() string { return "spa" }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) SetParent(parent types.Entity) { spa.parent = parent }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetParent() types.Entity { return spa.parent }

func (spa *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main
// Main
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board type. The type is interface{} with range: 0..4294967295.
    BoardType interface{}

    // Top assembly number. The type is string.
    Tan interface{}

    // Top assembly number revision. The type is string.
    TanRevision interface{}

    // Deviation number. The type is interface{} with range: 0..4294967295.
    DeviationNumber interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetFilter() yfilter.YFilter { return main.YFilter }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) SetFilter(yf yfilter.YFilter) { main.YFilter = yf }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetGoName(yname string) string {
    if yname == "board-type" { return "BoardType" }
    if yname == "tan" { return "Tan" }
    if yname == "tan-revision" { return "TanRevision" }
    if yname == "deviation-number" { return "DeviationNumber" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetSegmentPath() string {
    return "main"
}

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["board-type"] = main.BoardType
    leafs["tan"] = main.Tan
    leafs["tan-revision"] = main.TanRevision
    leafs["deviation-number"] = main.DeviationNumber
    leafs["serial-number"] = main.SerialNumber
    return leafs
}

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetBundleName() string { return "cisco_ios_xr" }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetYangName() string { return "main" }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) SetParent(parent types.Entity) { main.parent = parent }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetParent() types.Entity { return main.parent }

func (main *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_Main) GetParentYangName() string { return "spa" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision.
    HardwareRevision []Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetGoName(yname string) string {
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-revision" {
        for _, c := range hardwareRevision.HardwareRevision {
            if hardwareRevision.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision{}
        hardwareRevision.HardwareRevision = append(hardwareRevision.HardwareRevision, child)
        return &hardwareRevision.HardwareRevision[len(hardwareRevision.HardwareRevision)-1]
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareRevision.HardwareRevision {
        children[hardwareRevision.HardwareRevision[i].GetSegmentPath()] = &hardwareRevision.HardwareRevision[i]
    }
    return children
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision) GetParentYangName() string { return "spa" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node decsription. The type is string.
    NodeDescription interface{}

    // Version information. The type is string.
    Version interface{}

    // Hardware version.
    HwRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev

    // Firmware version.
    FwRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev

    // Software version.
    SwRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev

    // DIMM version information.
    DimmRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev

    // SSD version information.
    SsdRev Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetGoName(yname string) string {
    if yname == "node-description" { return "NodeDescription" }
    if yname == "version" { return "Version" }
    if yname == "hw-rev" { return "HwRev" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "sw-rev" { return "SwRev" }
    if yname == "dimm-rev" { return "DimmRev" }
    if yname == "ssd-rev" { return "SsdRev" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-rev" {
        return &hardwareRevision.HwRev
    }
    if childYangName == "fw-rev" {
        return &hardwareRevision.FwRev
    }
    if childYangName == "sw-rev" {
        return &hardwareRevision.SwRev
    }
    if childYangName == "dimm-rev" {
        return &hardwareRevision.DimmRev
    }
    if childYangName == "ssd-rev" {
        return &hardwareRevision.SsdRev
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-rev"] = &hardwareRevision.HwRev
    children["fw-rev"] = &hardwareRevision.FwRev
    children["sw-rev"] = &hardwareRevision.SwRev
    children["dimm-rev"] = &hardwareRevision.DimmRev
    children["ssd-rev"] = &hardwareRevision.SsdRev
    return children
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-description"] = hardwareRevision.NodeDescription
    leafs["version"] = hardwareRevision.Version
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetFilter() yfilter.YFilter { return hwRev.YFilter }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) SetFilter(yf yfilter.YFilter) { hwRev.YFilter = yf }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetSegmentPath() string {
    return "hw-rev"
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = hwRev.MajorRevision
    leafs["minor-revision"] = hwRev.MinorRevision
    return leafs
}

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetBundleName() string { return "cisco_ios_xr" }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetYangName() string { return "hw-rev" }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) SetParent(parent types.Entity) { hwRev.parent = parent }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetParent() types.Entity { return hwRev.parent }

func (hwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetFilter() yfilter.YFilter { return fwRev.YFilter }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) SetFilter(yf yfilter.YFilter) { fwRev.YFilter = yf }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetSegmentPath() string {
    return "fw-rev"
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = fwRev.MajorRevision
    leafs["minor-revision"] = fwRev.MinorRevision
    return leafs
}

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetBundleName() string { return "cisco_ios_xr" }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetYangName() string { return "fw-rev" }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) SetParent(parent types.Entity) { fwRev.parent = parent }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetParent() types.Entity { return fwRev.parent }

func (fwRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetFilter() yfilter.YFilter { return swRev.YFilter }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) SetFilter(yf yfilter.YFilter) { swRev.YFilter = yf }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetSegmentPath() string {
    return "sw-rev"
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = swRev.MajorRevision
    leafs["minor-revision"] = swRev.MinorRevision
    return leafs
}

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetBundleName() string { return "cisco_ios_xr" }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetYangName() string { return "sw-rev" }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) SetParent(parent types.Entity) { swRev.parent = parent }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetParent() types.Entity { return swRev.parent }

func (swRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size in MB. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Speed in MHz. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Locator information. The type is string.
    Locator interface{}

    // Column address strobe latency in clock cycles. The type is interface{} with
    // range: 0..4294967295.
    Cas interface{}
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetFilter() yfilter.YFilter { return dimmRev.YFilter }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) SetFilter(yf yfilter.YFilter) { dimmRev.YFilter = yf }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "speed" { return "Speed" }
    if yname == "locator" { return "Locator" }
    if yname == "cas" { return "Cas" }
    return ""
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetSegmentPath() string {
    return "dimm-rev"
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = dimmRev.Size
    leafs["speed"] = dimmRev.Speed
    leafs["locator"] = dimmRev.Locator
    leafs["cas"] = dimmRev.Cas
    return leafs
}

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetBundleName() string { return "cisco_ios_xr" }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetYangName() string { return "dimm-rev" }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) SetParent(parent types.Entity) { dimmRev.parent = parent }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetParent() types.Entity { return dimmRev.parent }

func (dimmRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetFilter() yfilter.YFilter { return ssdRev.YFilter }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) SetFilter(yf yfilter.YFilter) { ssdRev.YFilter = yf }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetSegmentPath() string {
    return "ssd-rev"
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ssdRev.Number
    leafs["fw-rev"] = ssdRev.FwRev
    leafs["serial-number"] = ssdRev.SerialNumber
    return leafs
}

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetBundleName() string { return "cisco_ios_xr" }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetYangName() string { return "ssd-rev" }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) SetParent(parent types.Entity) { ssdRev.parent = parent }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetParent() types.Entity { return ssdRev.parent }

func (ssdRev *Diag_Racks_Rack_PowerSupplies_PowerSupply_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots
// Diag operational data for all available slots
type Diag_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Diag operational data for a particular slot. The type is slice of
    // Diag_Racks_Rack_Slots_Slot.
    Slot []Diag_Racks_Rack_Slots_Slot
}

func (slots *Diag_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *Diag_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *Diag_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *Diag_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *Diag_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *Diag_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *Diag_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *Diag_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *Diag_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *Diag_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *Diag_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *Diag_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *Diag_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *Diag_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *Diag_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_Slots_Slot
// Diag operational data for a particular slot
type Diag_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SlotName interface{}

    // Slot detailed information.
    Detail Diag_Racks_Rack_Slots_Slot_Detail

    // Diag operational data for all available instances.
    Instances Diag_Racks_Rack_Slots_Slot_Instances
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *Diag_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *Diag_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot-name" { return "SlotName" }
    if yname == "detail" { return "Detail" }
    if yname == "instances" { return "Instances" }
    return ""
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot-name='" + fmt.Sprintf("%v", slot.SlotName) + "']"
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &slot.Detail
    }
    if childYangName == "instances" {
        return &slot.Instances
    }
    return nil
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &slot.Detail
    children["instances"] = &slot.Instances
    return children
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot-name"] = slot.SlotName
    return leafs
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *Diag_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *Diag_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *Diag_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *Diag_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *Diag_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *Diag_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *Diag_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// Diag_Racks_Rack_Slots_Slot_Detail
// Slot detailed information
type Diag_Racks_Rack_Slots_Slot_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detail information for slot. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail.
    NodeDetail []Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail

    // Detail information for spa. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail.
    SpaDetail []Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail
}

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetGoName(yname string) string {
    if yname == "node-detail" { return "NodeDetail" }
    if yname == "spa-detail" { return "SpaDetail" }
    return ""
}

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-detail" {
        for _, c := range detail.NodeDetail {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail{}
        detail.NodeDetail = append(detail.NodeDetail, child)
        return &detail.NodeDetail[len(detail.NodeDetail)-1]
    }
    if childYangName == "spa-detail" {
        for _, c := range detail.SpaDetail {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail{}
        detail.SpaDetail = append(detail.SpaDetail, child)
        return &detail.SpaDetail[len(detail.SpaDetail)-1]
    }
    return nil
}

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detail.NodeDetail {
        children[detail.NodeDetail[i].GetSegmentPath()] = &detail.NodeDetail[i]
    }
    for i := range detail.SpaDetail {
        children[detail.SpaDetail[i].GetSegmentPath()] = &detail.SpaDetail[i]
    }
    return children
}

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetYangName() string { return "detail" }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetParentYangName() string { return "slot" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail
// Detail information for slot
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describes in user-readable terms. The type is string.
    Description interface{}

    // Main serial number. The type is string.
    SerialNumber interface{}

    // Top assembly number. The type is string.
    Tan interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // Chip hardware revision. The type is string.
    ChipHardwareRevision interface{}

    // New deviation number. The type is interface{} with range: 0..4294967295.
    NewDeviationNumber interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Module operational state. The type is NodeState.
    BoardState interface{}

    // Motherboard PLD version. The type is string.
    PldMotherboard interface{}

    // Power PLD version. The type is string.
    PldPower interface{}

    // MONLIB version. The type is string.
    Monlib interface{}

    // ROMMON version. The type is string.
    Rommon interface{}

    // Processor type. The type is DiagProcessor.
    Cpu0 interface{}

    // Programmable logic device information.
    Pld Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld

    // Hardware revision.
    HardwareRevision Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision

    // CBC active partition.
    CbcActivePartition Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition

    // CBC inactive partition.
    CbcInactivePartition Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition
}

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetFilter() yfilter.YFilter { return nodeDetail.YFilter }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) SetFilter(yf yfilter.YFilter) { nodeDetail.YFilter = yf }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "tan" { return "Tan" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "clei" { return "Clei" }
    if yname == "board-state" { return "BoardState" }
    if yname == "pld-motherboard" { return "PldMotherboard" }
    if yname == "pld-power" { return "PldPower" }
    if yname == "monlib" { return "Monlib" }
    if yname == "rommon" { return "Rommon" }
    if yname == "cpu0" { return "Cpu0" }
    if yname == "pld" { return "Pld" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "cbc-active-partition" { return "CbcActivePartition" }
    if yname == "cbc-inactive-partition" { return "CbcInactivePartition" }
    return ""
}

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetSegmentPath() string {
    return "node-detail"
}

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pld" {
        return &nodeDetail.Pld
    }
    if childYangName == "hardware-revision" {
        return &nodeDetail.HardwareRevision
    }
    if childYangName == "cbc-active-partition" {
        return &nodeDetail.CbcActivePartition
    }
    if childYangName == "cbc-inactive-partition" {
        return &nodeDetail.CbcInactivePartition
    }
    return nil
}

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pld"] = &nodeDetail.Pld
    children["hardware-revision"] = &nodeDetail.HardwareRevision
    children["cbc-active-partition"] = &nodeDetail.CbcActivePartition
    children["cbc-inactive-partition"] = &nodeDetail.CbcInactivePartition
    return children
}

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = nodeDetail.Description
    leafs["serial-number"] = nodeDetail.SerialNumber
    leafs["tan"] = nodeDetail.Tan
    leafs["pid"] = nodeDetail.Pid
    leafs["vid"] = nodeDetail.Vid
    leafs["chip-hardware-revision"] = nodeDetail.ChipHardwareRevision
    leafs["new-deviation-number"] = nodeDetail.NewDeviationNumber
    leafs["clei"] = nodeDetail.Clei
    leafs["board-state"] = nodeDetail.BoardState
    leafs["pld-motherboard"] = nodeDetail.PldMotherboard
    leafs["pld-power"] = nodeDetail.PldPower
    leafs["monlib"] = nodeDetail.Monlib
    leafs["rommon"] = nodeDetail.Rommon
    leafs["cpu0"] = nodeDetail.Cpu0
    return leafs
}

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetBundleName() string { return "cisco_ios_xr" }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetYangName() string { return "node-detail" }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) SetParent(parent types.Entity) { nodeDetail.parent = parent }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetParent() types.Entity { return nodeDetail.parent }

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld
// Programmable logic device information
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Processor PLD version. The type is interface{} with range: 0..4294967295.
    Type interface{}

    // HigherVerion. The type is interface{} with range: 0..4294967295.
    ProcessorHigherVersion interface{}

    // LowerVersion. The type is interface{} with range: 0..4294967295.
    ProcessorLowerVersion interface{}
}

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetFilter() yfilter.YFilter { return pld.YFilter }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) SetFilter(yf yfilter.YFilter) { pld.YFilter = yf }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "processor-higher-version" { return "ProcessorHigherVersion" }
    if yname == "processor-lower-version" { return "ProcessorLowerVersion" }
    return ""
}

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetSegmentPath() string {
    return "pld"
}

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = pld.Type
    leafs["processor-higher-version"] = pld.ProcessorHigherVersion
    leafs["processor-lower-version"] = pld.ProcessorLowerVersion
    return leafs
}

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetBundleName() string { return "cisco_ios_xr" }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetYangName() string { return "pld" }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) SetParent(parent types.Entity) { pld.parent = parent }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetParent() types.Entity { return pld.parent }

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetParentYangName() string { return "node-detail" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision.
    HardwareRevision []Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetGoName(yname string) string {
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-revision" {
        for _, c := range hardwareRevision.HardwareRevision {
            if hardwareRevision.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision{}
        hardwareRevision.HardwareRevision = append(hardwareRevision.HardwareRevision, child)
        return &hardwareRevision.HardwareRevision[len(hardwareRevision.HardwareRevision)-1]
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareRevision.HardwareRevision {
        children[hardwareRevision.HardwareRevision[i].GetSegmentPath()] = &hardwareRevision.HardwareRevision[i]
    }
    return children
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetParentYangName() string { return "node-detail" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node decsription. The type is string.
    NodeDescription interface{}

    // Version information. The type is string.
    Version interface{}

    // Hardware version.
    HwRev Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev

    // Firmware version.
    FwRev Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev

    // Software version.
    SwRev Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev

    // DIMM version information.
    DimmRev Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev

    // SSD version information.
    SsdRev Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetGoName(yname string) string {
    if yname == "node-description" { return "NodeDescription" }
    if yname == "version" { return "Version" }
    if yname == "hw-rev" { return "HwRev" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "sw-rev" { return "SwRev" }
    if yname == "dimm-rev" { return "DimmRev" }
    if yname == "ssd-rev" { return "SsdRev" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-rev" {
        return &hardwareRevision.HwRev
    }
    if childYangName == "fw-rev" {
        return &hardwareRevision.FwRev
    }
    if childYangName == "sw-rev" {
        return &hardwareRevision.SwRev
    }
    if childYangName == "dimm-rev" {
        return &hardwareRevision.DimmRev
    }
    if childYangName == "ssd-rev" {
        return &hardwareRevision.SsdRev
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-rev"] = &hardwareRevision.HwRev
    children["fw-rev"] = &hardwareRevision.FwRev
    children["sw-rev"] = &hardwareRevision.SwRev
    children["dimm-rev"] = &hardwareRevision.DimmRev
    children["ssd-rev"] = &hardwareRevision.SsdRev
    return children
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-description"] = hardwareRevision.NodeDescription
    leafs["version"] = hardwareRevision.Version
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetFilter() yfilter.YFilter { return hwRev.YFilter }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) SetFilter(yf yfilter.YFilter) { hwRev.YFilter = yf }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetSegmentPath() string {
    return "hw-rev"
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = hwRev.MajorRevision
    leafs["minor-revision"] = hwRev.MinorRevision
    return leafs
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetBundleName() string { return "cisco_ios_xr" }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetYangName() string { return "hw-rev" }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) SetParent(parent types.Entity) { hwRev.parent = parent }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetParent() types.Entity { return hwRev.parent }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetFilter() yfilter.YFilter { return fwRev.YFilter }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) SetFilter(yf yfilter.YFilter) { fwRev.YFilter = yf }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetSegmentPath() string {
    return "fw-rev"
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = fwRev.MajorRevision
    leafs["minor-revision"] = fwRev.MinorRevision
    return leafs
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetBundleName() string { return "cisco_ios_xr" }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetYangName() string { return "fw-rev" }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) SetParent(parent types.Entity) { fwRev.parent = parent }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetParent() types.Entity { return fwRev.parent }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetFilter() yfilter.YFilter { return swRev.YFilter }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) SetFilter(yf yfilter.YFilter) { swRev.YFilter = yf }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetSegmentPath() string {
    return "sw-rev"
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = swRev.MajorRevision
    leafs["minor-revision"] = swRev.MinorRevision
    return leafs
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetBundleName() string { return "cisco_ios_xr" }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetYangName() string { return "sw-rev" }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) SetParent(parent types.Entity) { swRev.parent = parent }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetParent() types.Entity { return swRev.parent }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size in MB. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Speed in MHz. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Locator information. The type is string.
    Locator interface{}

    // Column address strobe latency in clock cycles. The type is interface{} with
    // range: 0..4294967295.
    Cas interface{}
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetFilter() yfilter.YFilter { return dimmRev.YFilter }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) SetFilter(yf yfilter.YFilter) { dimmRev.YFilter = yf }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "speed" { return "Speed" }
    if yname == "locator" { return "Locator" }
    if yname == "cas" { return "Cas" }
    return ""
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetSegmentPath() string {
    return "dimm-rev"
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = dimmRev.Size
    leafs["speed"] = dimmRev.Speed
    leafs["locator"] = dimmRev.Locator
    leafs["cas"] = dimmRev.Cas
    return leafs
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetBundleName() string { return "cisco_ios_xr" }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetYangName() string { return "dimm-rev" }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) SetParent(parent types.Entity) { dimmRev.parent = parent }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetParent() types.Entity { return dimmRev.parent }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetFilter() yfilter.YFilter { return ssdRev.YFilter }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) SetFilter(yf yfilter.YFilter) { ssdRev.YFilter = yf }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetSegmentPath() string {
    return "ssd-rev"
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ssdRev.Number
    leafs["fw-rev"] = ssdRev.FwRev
    leafs["serial-number"] = ssdRev.SerialNumber
    return leafs
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetBundleName() string { return "cisco_ios_xr" }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetYangName() string { return "ssd-rev" }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) SetParent(parent types.Entity) { ssdRev.parent = parent }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetParent() types.Entity { return ssdRev.parent }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition
// CBC active partition
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetFilter() yfilter.YFilter { return cbcActivePartition.YFilter }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) SetFilter(yf yfilter.YFilter) { cbcActivePartition.YFilter = yf }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetSegmentPath() string {
    return "cbc-active-partition"
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = cbcActivePartition.MajorRevision
    leafs["minor-revision"] = cbcActivePartition.MinorRevision
    return leafs
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetBundleName() string { return "cisco_ios_xr" }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetYangName() string { return "cbc-active-partition" }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) SetParent(parent types.Entity) { cbcActivePartition.parent = parent }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetParent() types.Entity { return cbcActivePartition.parent }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetParentYangName() string { return "node-detail" }

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition
// CBC inactive partition
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetFilter() yfilter.YFilter { return cbcInactivePartition.YFilter }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) SetFilter(yf yfilter.YFilter) { cbcInactivePartition.YFilter = yf }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetSegmentPath() string {
    return "cbc-inactive-partition"
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = cbcInactivePartition.MajorRevision
    leafs["minor-revision"] = cbcInactivePartition.MinorRevision
    return leafs
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetBundleName() string { return "cisco_ios_xr" }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetYangName() string { return "cbc-inactive-partition" }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) SetParent(parent types.Entity) { cbcInactivePartition.parent = parent }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetParent() types.Entity { return cbcInactivePartition.parent }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetParentYangName() string { return "node-detail" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail
// Detail information for spa
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // SPA name. The type is string.
    Name interface{}

    // Format is C-BN-V where C = Class Code, BN = Base Number, and V = Version.
    // The type is string.
    PcaUnitNumber interface{}

    // PCA revision information. The type is string.
    PcaRevision interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Node state. The type is NodeState.
    NodeState interface{}

    // Main.
    Main Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main

    // Hardware revision.
    HardwareRevision Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision
}

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetFilter() yfilter.YFilter { return spaDetail.YFilter }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) SetFilter(yf yfilter.YFilter) { spaDetail.YFilter = yf }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "name" { return "Name" }
    if yname == "pca-unit-number" { return "PcaUnitNumber" }
    if yname == "pca-revision" { return "PcaRevision" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "clei" { return "Clei" }
    if yname == "node-state" { return "NodeState" }
    if yname == "main" { return "Main" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetSegmentPath() string {
    return "spa-detail"
}

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "main" {
        return &spaDetail.Main
    }
    if childYangName == "hardware-revision" {
        return &spaDetail.HardwareRevision
    }
    return nil
}

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["main"] = &spaDetail.Main
    children["hardware-revision"] = &spaDetail.HardwareRevision
    return children
}

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = spaDetail.Node
    leafs["name"] = spaDetail.Name
    leafs["pca-unit-number"] = spaDetail.PcaUnitNumber
    leafs["pca-revision"] = spaDetail.PcaRevision
    leafs["pid"] = spaDetail.Pid
    leafs["vid"] = spaDetail.Vid
    leafs["clei"] = spaDetail.Clei
    leafs["node-state"] = spaDetail.NodeState
    return leafs
}

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetBundleName() string { return "cisco_ios_xr" }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetYangName() string { return "spa-detail" }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) SetParent(parent types.Entity) { spaDetail.parent = parent }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetParent() types.Entity { return spaDetail.parent }

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main
// Main
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board type. The type is interface{} with range: 0..4294967295.
    BoardType interface{}

    // Top assembly number. The type is string.
    Tan interface{}

    // Top assembly number revision. The type is string.
    TanRevision interface{}

    // Deviation number. The type is interface{} with range: 0..4294967295.
    DeviationNumber interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetFilter() yfilter.YFilter { return main.YFilter }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) SetFilter(yf yfilter.YFilter) { main.YFilter = yf }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetGoName(yname string) string {
    if yname == "board-type" { return "BoardType" }
    if yname == "tan" { return "Tan" }
    if yname == "tan-revision" { return "TanRevision" }
    if yname == "deviation-number" { return "DeviationNumber" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetSegmentPath() string {
    return "main"
}

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["board-type"] = main.BoardType
    leafs["tan"] = main.Tan
    leafs["tan-revision"] = main.TanRevision
    leafs["deviation-number"] = main.DeviationNumber
    leafs["serial-number"] = main.SerialNumber
    return leafs
}

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetBundleName() string { return "cisco_ios_xr" }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetYangName() string { return "main" }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) SetParent(parent types.Entity) { main.parent = parent }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetParent() types.Entity { return main.parent }

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetParentYangName() string { return "spa-detail" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision.
    HardwareRevision []Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetGoName(yname string) string {
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-revision" {
        for _, c := range hardwareRevision.HardwareRevision {
            if hardwareRevision.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision{}
        hardwareRevision.HardwareRevision = append(hardwareRevision.HardwareRevision, child)
        return &hardwareRevision.HardwareRevision[len(hardwareRevision.HardwareRevision)-1]
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareRevision.HardwareRevision {
        children[hardwareRevision.HardwareRevision[i].GetSegmentPath()] = &hardwareRevision.HardwareRevision[i]
    }
    return children
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetParentYangName() string { return "spa-detail" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node decsription. The type is string.
    NodeDescription interface{}

    // Version information. The type is string.
    Version interface{}

    // Hardware version.
    HwRev Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev

    // Firmware version.
    FwRev Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev

    // Software version.
    SwRev Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev

    // DIMM version information.
    DimmRev Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev

    // SSD version information.
    SsdRev Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetGoName(yname string) string {
    if yname == "node-description" { return "NodeDescription" }
    if yname == "version" { return "Version" }
    if yname == "hw-rev" { return "HwRev" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "sw-rev" { return "SwRev" }
    if yname == "dimm-rev" { return "DimmRev" }
    if yname == "ssd-rev" { return "SsdRev" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-rev" {
        return &hardwareRevision.HwRev
    }
    if childYangName == "fw-rev" {
        return &hardwareRevision.FwRev
    }
    if childYangName == "sw-rev" {
        return &hardwareRevision.SwRev
    }
    if childYangName == "dimm-rev" {
        return &hardwareRevision.DimmRev
    }
    if childYangName == "ssd-rev" {
        return &hardwareRevision.SsdRev
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-rev"] = &hardwareRevision.HwRev
    children["fw-rev"] = &hardwareRevision.FwRev
    children["sw-rev"] = &hardwareRevision.SwRev
    children["dimm-rev"] = &hardwareRevision.DimmRev
    children["ssd-rev"] = &hardwareRevision.SsdRev
    return children
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-description"] = hardwareRevision.NodeDescription
    leafs["version"] = hardwareRevision.Version
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetFilter() yfilter.YFilter { return hwRev.YFilter }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) SetFilter(yf yfilter.YFilter) { hwRev.YFilter = yf }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetSegmentPath() string {
    return "hw-rev"
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = hwRev.MajorRevision
    leafs["minor-revision"] = hwRev.MinorRevision
    return leafs
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetBundleName() string { return "cisco_ios_xr" }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetYangName() string { return "hw-rev" }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) SetParent(parent types.Entity) { hwRev.parent = parent }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetParent() types.Entity { return hwRev.parent }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetFilter() yfilter.YFilter { return fwRev.YFilter }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) SetFilter(yf yfilter.YFilter) { fwRev.YFilter = yf }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetSegmentPath() string {
    return "fw-rev"
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = fwRev.MajorRevision
    leafs["minor-revision"] = fwRev.MinorRevision
    return leafs
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetBundleName() string { return "cisco_ios_xr" }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetYangName() string { return "fw-rev" }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) SetParent(parent types.Entity) { fwRev.parent = parent }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetParent() types.Entity { return fwRev.parent }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetFilter() yfilter.YFilter { return swRev.YFilter }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) SetFilter(yf yfilter.YFilter) { swRev.YFilter = yf }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetSegmentPath() string {
    return "sw-rev"
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = swRev.MajorRevision
    leafs["minor-revision"] = swRev.MinorRevision
    return leafs
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetBundleName() string { return "cisco_ios_xr" }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetYangName() string { return "sw-rev" }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) SetParent(parent types.Entity) { swRev.parent = parent }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetParent() types.Entity { return swRev.parent }

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size in MB. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Speed in MHz. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Locator information. The type is string.
    Locator interface{}

    // Column address strobe latency in clock cycles. The type is interface{} with
    // range: 0..4294967295.
    Cas interface{}
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetFilter() yfilter.YFilter { return dimmRev.YFilter }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) SetFilter(yf yfilter.YFilter) { dimmRev.YFilter = yf }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "speed" { return "Speed" }
    if yname == "locator" { return "Locator" }
    if yname == "cas" { return "Cas" }
    return ""
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetSegmentPath() string {
    return "dimm-rev"
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = dimmRev.Size
    leafs["speed"] = dimmRev.Speed
    leafs["locator"] = dimmRev.Locator
    leafs["cas"] = dimmRev.Cas
    return leafs
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetBundleName() string { return "cisco_ios_xr" }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetYangName() string { return "dimm-rev" }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) SetParent(parent types.Entity) { dimmRev.parent = parent }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetParent() types.Entity { return dimmRev.parent }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetFilter() yfilter.YFilter { return ssdRev.YFilter }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) SetFilter(yf yfilter.YFilter) { ssdRev.YFilter = yf }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetSegmentPath() string {
    return "ssd-rev"
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ssdRev.Number
    leafs["fw-rev"] = ssdRev.FwRev
    leafs["serial-number"] = ssdRev.SerialNumber
    return leafs
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetBundleName() string { return "cisco_ios_xr" }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetYangName() string { return "ssd-rev" }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) SetParent(parent types.Entity) { ssdRev.parent = parent }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetParent() types.Entity { return ssdRev.parent }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances
// Diag operational data for all available
// instances
type Diag_Racks_Rack_Slots_Slot_Instances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Diag operational data for a particular instance. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Instances_Instance.
    Instance []Diag_Racks_Rack_Slots_Slot_Instances_Instance
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetFilter() yfilter.YFilter { return instances.YFilter }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) SetFilter(yf yfilter.YFilter) { instances.YFilter = yf }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    return ""
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetSegmentPath() string {
    return "instances"
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        for _, c := range instances.Instance {
            if instances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot_Instances_Instance{}
        instances.Instance = append(instances.Instance, child)
        return &instances.Instance[len(instances.Instance)-1]
    }
    return nil
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range instances.Instance {
        children[instances.Instance[i].GetSegmentPath()] = &instances.Instance[i]
    }
    return children
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetBundleName() string { return "cisco_ios_xr" }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetYangName() string { return "instances" }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) SetParent(parent types.Entity) { instances.parent = parent }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetParent() types.Entity { return instances.parent }

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetParentYangName() string { return "slot" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance
// Diag operational data for a particular
// instance
type Diag_Racks_Rack_Slots_Slot_Instances_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Instance name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Diag detailed information.
    Detail Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetSegmentPath() string {
    return "instance" + "[instance-name='" + fmt.Sprintf("%v", instance.InstanceName) + "']"
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &instance.Detail
    }
    return nil
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &instance.Detail
    return children
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = instance.InstanceName
    return leafs
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetYangName() string { return "instance" }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetParent() types.Entity { return instance.parent }

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetParentYangName() string { return "instances" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail
// Diag detailed information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node information.
    Node Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node

    // SPA information.
    Spa Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "spa" { return "Spa" }
    return ""
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        return &detail.Node
    }
    if childYangName == "spa" {
        return &detail.Spa
    }
    return nil
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node"] = &detail.Node
    children["spa"] = &detail.Spa
    return children
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetYangName() string { return "detail" }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetParentYangName() string { return "instance" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node
// Node information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describes in user-readable terms. The type is string.
    Description interface{}

    // Main serial number. The type is string.
    SerialNumber interface{}

    // Top assembly number. The type is string.
    Tan interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // Chip hardware revision. The type is string.
    ChipHardwareRevision interface{}

    // New deviation number. The type is interface{} with range: 0..4294967295.
    NewDeviationNumber interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Module operational state. The type is NodeState.
    BoardState interface{}

    // Motherboard PLD version. The type is string.
    PldMotherboard interface{}

    // Power PLD version. The type is string.
    PldPower interface{}

    // MONLIB version. The type is string.
    Monlib interface{}

    // ROMMON version. The type is string.
    Rommon interface{}

    // Processor type. The type is DiagProcessor.
    Cpu0 interface{}

    // Programmable logic device information.
    Pld Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld

    // Hardware revision.
    HardwareRevision Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision

    // CBC active partition.
    CbcActivePartition Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition

    // CBC inactive partition.
    CbcInactivePartition Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition
}

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "tan" { return "Tan" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "chip-hardware-revision" { return "ChipHardwareRevision" }
    if yname == "new-deviation-number" { return "NewDeviationNumber" }
    if yname == "clei" { return "Clei" }
    if yname == "board-state" { return "BoardState" }
    if yname == "pld-motherboard" { return "PldMotherboard" }
    if yname == "pld-power" { return "PldPower" }
    if yname == "monlib" { return "Monlib" }
    if yname == "rommon" { return "Rommon" }
    if yname == "cpu0" { return "Cpu0" }
    if yname == "pld" { return "Pld" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    if yname == "cbc-active-partition" { return "CbcActivePartition" }
    if yname == "cbc-inactive-partition" { return "CbcInactivePartition" }
    return ""
}

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetSegmentPath() string {
    return "node"
}

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pld" {
        return &node.Pld
    }
    if childYangName == "hardware-revision" {
        return &node.HardwareRevision
    }
    if childYangName == "cbc-active-partition" {
        return &node.CbcActivePartition
    }
    if childYangName == "cbc-inactive-partition" {
        return &node.CbcInactivePartition
    }
    return nil
}

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pld"] = &node.Pld
    children["hardware-revision"] = &node.HardwareRevision
    children["cbc-active-partition"] = &node.CbcActivePartition
    children["cbc-inactive-partition"] = &node.CbcInactivePartition
    return children
}

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = node.Description
    leafs["serial-number"] = node.SerialNumber
    leafs["tan"] = node.Tan
    leafs["pid"] = node.Pid
    leafs["vid"] = node.Vid
    leafs["chip-hardware-revision"] = node.ChipHardwareRevision
    leafs["new-deviation-number"] = node.NewDeviationNumber
    leafs["clei"] = node.Clei
    leafs["board-state"] = node.BoardState
    leafs["pld-motherboard"] = node.PldMotherboard
    leafs["pld-power"] = node.PldPower
    leafs["monlib"] = node.Monlib
    leafs["rommon"] = node.Rommon
    leafs["cpu0"] = node.Cpu0
    return leafs
}

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetYangName() string { return "node" }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetParent() types.Entity { return node.parent }

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld
// Programmable logic device information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Processor PLD version. The type is interface{} with range: 0..4294967295.
    Type interface{}

    // HigherVerion. The type is interface{} with range: 0..4294967295.
    ProcessorHigherVersion interface{}

    // LowerVersion. The type is interface{} with range: 0..4294967295.
    ProcessorLowerVersion interface{}
}

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetFilter() yfilter.YFilter { return pld.YFilter }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) SetFilter(yf yfilter.YFilter) { pld.YFilter = yf }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "processor-higher-version" { return "ProcessorHigherVersion" }
    if yname == "processor-lower-version" { return "ProcessorLowerVersion" }
    return ""
}

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetSegmentPath() string {
    return "pld"
}

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = pld.Type
    leafs["processor-higher-version"] = pld.ProcessorHigherVersion
    leafs["processor-lower-version"] = pld.ProcessorLowerVersion
    return leafs
}

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetBundleName() string { return "cisco_ios_xr" }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetYangName() string { return "pld" }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) SetParent(parent types.Entity) { pld.parent = parent }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetParent() types.Entity { return pld.parent }

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision.
    HardwareRevision []Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetGoName(yname string) string {
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-revision" {
        for _, c := range hardwareRevision.HardwareRevision {
            if hardwareRevision.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision{}
        hardwareRevision.HardwareRevision = append(hardwareRevision.HardwareRevision, child)
        return &hardwareRevision.HardwareRevision[len(hardwareRevision.HardwareRevision)-1]
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareRevision.HardwareRevision {
        children[hardwareRevision.HardwareRevision[i].GetSegmentPath()] = &hardwareRevision.HardwareRevision[i]
    }
    return children
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node decsription. The type is string.
    NodeDescription interface{}

    // Version information. The type is string.
    Version interface{}

    // Hardware version.
    HwRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev

    // Firmware version.
    FwRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev

    // Software version.
    SwRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev

    // DIMM version information.
    DimmRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev

    // SSD version information.
    SsdRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetGoName(yname string) string {
    if yname == "node-description" { return "NodeDescription" }
    if yname == "version" { return "Version" }
    if yname == "hw-rev" { return "HwRev" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "sw-rev" { return "SwRev" }
    if yname == "dimm-rev" { return "DimmRev" }
    if yname == "ssd-rev" { return "SsdRev" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-rev" {
        return &hardwareRevision.HwRev
    }
    if childYangName == "fw-rev" {
        return &hardwareRevision.FwRev
    }
    if childYangName == "sw-rev" {
        return &hardwareRevision.SwRev
    }
    if childYangName == "dimm-rev" {
        return &hardwareRevision.DimmRev
    }
    if childYangName == "ssd-rev" {
        return &hardwareRevision.SsdRev
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-rev"] = &hardwareRevision.HwRev
    children["fw-rev"] = &hardwareRevision.FwRev
    children["sw-rev"] = &hardwareRevision.SwRev
    children["dimm-rev"] = &hardwareRevision.DimmRev
    children["ssd-rev"] = &hardwareRevision.SsdRev
    return children
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-description"] = hardwareRevision.NodeDescription
    leafs["version"] = hardwareRevision.Version
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetFilter() yfilter.YFilter { return hwRev.YFilter }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) SetFilter(yf yfilter.YFilter) { hwRev.YFilter = yf }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetSegmentPath() string {
    return "hw-rev"
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = hwRev.MajorRevision
    leafs["minor-revision"] = hwRev.MinorRevision
    return leafs
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetBundleName() string { return "cisco_ios_xr" }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetYangName() string { return "hw-rev" }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) SetParent(parent types.Entity) { hwRev.parent = parent }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetParent() types.Entity { return hwRev.parent }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetFilter() yfilter.YFilter { return fwRev.YFilter }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) SetFilter(yf yfilter.YFilter) { fwRev.YFilter = yf }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetSegmentPath() string {
    return "fw-rev"
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = fwRev.MajorRevision
    leafs["minor-revision"] = fwRev.MinorRevision
    return leafs
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetBundleName() string { return "cisco_ios_xr" }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetYangName() string { return "fw-rev" }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) SetParent(parent types.Entity) { fwRev.parent = parent }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetParent() types.Entity { return fwRev.parent }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetFilter() yfilter.YFilter { return swRev.YFilter }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) SetFilter(yf yfilter.YFilter) { swRev.YFilter = yf }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetSegmentPath() string {
    return "sw-rev"
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = swRev.MajorRevision
    leafs["minor-revision"] = swRev.MinorRevision
    return leafs
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetBundleName() string { return "cisco_ios_xr" }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetYangName() string { return "sw-rev" }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) SetParent(parent types.Entity) { swRev.parent = parent }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetParent() types.Entity { return swRev.parent }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size in MB. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Speed in MHz. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Locator information. The type is string.
    Locator interface{}

    // Column address strobe latency in clock cycles. The type is interface{} with
    // range: 0..4294967295.
    Cas interface{}
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetFilter() yfilter.YFilter { return dimmRev.YFilter }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) SetFilter(yf yfilter.YFilter) { dimmRev.YFilter = yf }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "speed" { return "Speed" }
    if yname == "locator" { return "Locator" }
    if yname == "cas" { return "Cas" }
    return ""
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetSegmentPath() string {
    return "dimm-rev"
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = dimmRev.Size
    leafs["speed"] = dimmRev.Speed
    leafs["locator"] = dimmRev.Locator
    leafs["cas"] = dimmRev.Cas
    return leafs
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetBundleName() string { return "cisco_ios_xr" }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetYangName() string { return "dimm-rev" }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) SetParent(parent types.Entity) { dimmRev.parent = parent }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetParent() types.Entity { return dimmRev.parent }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetFilter() yfilter.YFilter { return ssdRev.YFilter }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) SetFilter(yf yfilter.YFilter) { ssdRev.YFilter = yf }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetSegmentPath() string {
    return "ssd-rev"
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ssdRev.Number
    leafs["fw-rev"] = ssdRev.FwRev
    leafs["serial-number"] = ssdRev.SerialNumber
    return leafs
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetBundleName() string { return "cisco_ios_xr" }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetYangName() string { return "ssd-rev" }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) SetParent(parent types.Entity) { ssdRev.parent = parent }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetParent() types.Entity { return ssdRev.parent }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition
// CBC active partition
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetFilter() yfilter.YFilter { return cbcActivePartition.YFilter }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) SetFilter(yf yfilter.YFilter) { cbcActivePartition.YFilter = yf }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetSegmentPath() string {
    return "cbc-active-partition"
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = cbcActivePartition.MajorRevision
    leafs["minor-revision"] = cbcActivePartition.MinorRevision
    return leafs
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetBundleName() string { return "cisco_ios_xr" }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetYangName() string { return "cbc-active-partition" }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) SetParent(parent types.Entity) { cbcActivePartition.parent = parent }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetParent() types.Entity { return cbcActivePartition.parent }

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition
// CBC inactive partition
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetFilter() yfilter.YFilter { return cbcInactivePartition.YFilter }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) SetFilter(yf yfilter.YFilter) { cbcInactivePartition.YFilter = yf }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetSegmentPath() string {
    return "cbc-inactive-partition"
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = cbcInactivePartition.MajorRevision
    leafs["minor-revision"] = cbcInactivePartition.MinorRevision
    return leafs
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetBundleName() string { return "cisco_ios_xr" }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetYangName() string { return "cbc-inactive-partition" }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) SetParent(parent types.Entity) { cbcInactivePartition.parent = parent }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetParent() types.Entity { return cbcInactivePartition.parent }

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetParentYangName() string { return "node" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa
// SPA information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // SPA name. The type is string.
    Name interface{}

    // Format is C-BN-V where C = Class Code, BN = Base Number, and V = Version.
    // The type is string.
    PcaUnitNumber interface{}

    // PCA revision information. The type is string.
    PcaRevision interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Node state. The type is NodeState.
    NodeState interface{}

    // Main.
    Main Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main

    // Hardware revision.
    HardwareRevision Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision
}

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetFilter() yfilter.YFilter { return spa.YFilter }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) SetFilter(yf yfilter.YFilter) { spa.YFilter = yf }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "name" { return "Name" }
    if yname == "pca-unit-number" { return "PcaUnitNumber" }
    if yname == "pca-revision" { return "PcaRevision" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "clei" { return "Clei" }
    if yname == "node-state" { return "NodeState" }
    if yname == "main" { return "Main" }
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetSegmentPath() string {
    return "spa"
}

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "main" {
        return &spa.Main
    }
    if childYangName == "hardware-revision" {
        return &spa.HardwareRevision
    }
    return nil
}

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["main"] = &spa.Main
    children["hardware-revision"] = &spa.HardwareRevision
    return children
}

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = spa.Node
    leafs["name"] = spa.Name
    leafs["pca-unit-number"] = spa.PcaUnitNumber
    leafs["pca-revision"] = spa.PcaRevision
    leafs["pid"] = spa.Pid
    leafs["vid"] = spa.Vid
    leafs["clei"] = spa.Clei
    leafs["node-state"] = spa.NodeState
    return leafs
}

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetBundleName() string { return "cisco_ios_xr" }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetYangName() string { return "spa" }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) SetParent(parent types.Entity) { spa.parent = parent }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetParent() types.Entity { return spa.parent }

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetParentYangName() string { return "detail" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main
// Main
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board type. The type is interface{} with range: 0..4294967295.
    BoardType interface{}

    // Top assembly number. The type is string.
    Tan interface{}

    // Top assembly number revision. The type is string.
    TanRevision interface{}

    // Deviation number. The type is interface{} with range: 0..4294967295.
    DeviationNumber interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetFilter() yfilter.YFilter { return main.YFilter }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) SetFilter(yf yfilter.YFilter) { main.YFilter = yf }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetGoName(yname string) string {
    if yname == "board-type" { return "BoardType" }
    if yname == "tan" { return "Tan" }
    if yname == "tan-revision" { return "TanRevision" }
    if yname == "deviation-number" { return "DeviationNumber" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetSegmentPath() string {
    return "main"
}

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["board-type"] = main.BoardType
    leafs["tan"] = main.Tan
    leafs["tan-revision"] = main.TanRevision
    leafs["deviation-number"] = main.DeviationNumber
    leafs["serial-number"] = main.SerialNumber
    return leafs
}

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetBundleName() string { return "cisco_ios_xr" }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetYangName() string { return "main" }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) SetParent(parent types.Entity) { main.parent = parent }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetParent() types.Entity { return main.parent }

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetParentYangName() string { return "spa" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision.
    HardwareRevision []Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetGoName(yname string) string {
    if yname == "hardware-revision" { return "HardwareRevision" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-revision" {
        for _, c := range hardwareRevision.HardwareRevision {
            if hardwareRevision.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision{}
        hardwareRevision.HardwareRevision = append(hardwareRevision.HardwareRevision, child)
        return &hardwareRevision.HardwareRevision[len(hardwareRevision.HardwareRevision)-1]
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareRevision.HardwareRevision {
        children[hardwareRevision.HardwareRevision[i].GetSegmentPath()] = &hardwareRevision.HardwareRevision[i]
    }
    return children
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetParentYangName() string { return "spa" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node decsription. The type is string.
    NodeDescription interface{}

    // Version information. The type is string.
    Version interface{}

    // Hardware version.
    HwRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev

    // Firmware version.
    FwRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev

    // Software version.
    SwRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev

    // DIMM version information.
    DimmRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev

    // SSD version information.
    SsdRev Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetFilter() yfilter.YFilter { return hardwareRevision.YFilter }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) SetFilter(yf yfilter.YFilter) { hardwareRevision.YFilter = yf }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetGoName(yname string) string {
    if yname == "node-description" { return "NodeDescription" }
    if yname == "version" { return "Version" }
    if yname == "hw-rev" { return "HwRev" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "sw-rev" { return "SwRev" }
    if yname == "dimm-rev" { return "DimmRev" }
    if yname == "ssd-rev" { return "SsdRev" }
    return ""
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetSegmentPath() string {
    return "hardware-revision"
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-rev" {
        return &hardwareRevision.HwRev
    }
    if childYangName == "fw-rev" {
        return &hardwareRevision.FwRev
    }
    if childYangName == "sw-rev" {
        return &hardwareRevision.SwRev
    }
    if childYangName == "dimm-rev" {
        return &hardwareRevision.DimmRev
    }
    if childYangName == "ssd-rev" {
        return &hardwareRevision.SsdRev
    }
    return nil
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-rev"] = &hardwareRevision.HwRev
    children["fw-rev"] = &hardwareRevision.FwRev
    children["sw-rev"] = &hardwareRevision.SwRev
    children["dimm-rev"] = &hardwareRevision.DimmRev
    children["ssd-rev"] = &hardwareRevision.SsdRev
    return children
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-description"] = hardwareRevision.NodeDescription
    leafs["version"] = hardwareRevision.Version
    return leafs
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetYangName() string { return "hardware-revision" }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) SetParent(parent types.Entity) { hardwareRevision.parent = parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetParent() types.Entity { return hardwareRevision.parent }

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetFilter() yfilter.YFilter { return hwRev.YFilter }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) SetFilter(yf yfilter.YFilter) { hwRev.YFilter = yf }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetSegmentPath() string {
    return "hw-rev"
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = hwRev.MajorRevision
    leafs["minor-revision"] = hwRev.MinorRevision
    return leafs
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetBundleName() string { return "cisco_ios_xr" }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetYangName() string { return "hw-rev" }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) SetParent(parent types.Entity) { hwRev.parent = parent }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetParent() types.Entity { return hwRev.parent }

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetFilter() yfilter.YFilter { return fwRev.YFilter }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) SetFilter(yf yfilter.YFilter) { fwRev.YFilter = yf }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetSegmentPath() string {
    return "fw-rev"
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = fwRev.MajorRevision
    leafs["minor-revision"] = fwRev.MinorRevision
    return leafs
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetBundleName() string { return "cisco_ios_xr" }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetYangName() string { return "fw-rev" }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) SetParent(parent types.Entity) { fwRev.parent = parent }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetParent() types.Entity { return fwRev.parent }

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetFilter() yfilter.YFilter { return swRev.YFilter }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) SetFilter(yf yfilter.YFilter) { swRev.YFilter = yf }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetGoName(yname string) string {
    if yname == "major-revision" { return "MajorRevision" }
    if yname == "minor-revision" { return "MinorRevision" }
    return ""
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetSegmentPath() string {
    return "sw-rev"
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["major-revision"] = swRev.MajorRevision
    leafs["minor-revision"] = swRev.MinorRevision
    return leafs
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetBundleName() string { return "cisco_ios_xr" }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetYangName() string { return "sw-rev" }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) SetParent(parent types.Entity) { swRev.parent = parent }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetParent() types.Entity { return swRev.parent }

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size in MB. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Speed in MHz. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Locator information. The type is string.
    Locator interface{}

    // Column address strobe latency in clock cycles. The type is interface{} with
    // range: 0..4294967295.
    Cas interface{}
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetFilter() yfilter.YFilter { return dimmRev.YFilter }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) SetFilter(yf yfilter.YFilter) { dimmRev.YFilter = yf }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "speed" { return "Speed" }
    if yname == "locator" { return "Locator" }
    if yname == "cas" { return "Cas" }
    return ""
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetSegmentPath() string {
    return "dimm-rev"
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = dimmRev.Size
    leafs["speed"] = dimmRev.Speed
    leafs["locator"] = dimmRev.Locator
    leafs["cas"] = dimmRev.Cas
    return leafs
}

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetBundleName() string { return "cisco_ios_xr" }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetYangName() string { return "dimm-rev" }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) SetParent(parent types.Entity) { dimmRev.parent = parent }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetParent() types.Entity { return dimmRev.parent }

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetFilter() yfilter.YFilter { return ssdRev.YFilter }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) SetFilter(yf yfilter.YFilter) { ssdRev.YFilter = yf }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "fw-rev" { return "FwRev" }
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetSegmentPath() string {
    return "ssd-rev"
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ssdRev.Number
    leafs["fw-rev"] = ssdRev.FwRev
    leafs["serial-number"] = ssdRev.SerialNumber
    return leafs
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetBundleName() string { return "cisco_ios_xr" }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetYangName() string { return "ssd-rev" }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) SetParent(parent types.Entity) { ssdRev.parent = parent }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetParent() types.Entity { return ssdRev.parent }

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetParentYangName() string { return "hardware-revision" }

// Diag_Racks_Rack_Chassis
// Chassis information
type Diag_Racks_Rack_Chassis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Serial number. The type is string.
    SerialNumber interface{}

    // PID. The type is string.
    Pid interface{}

    // VID. The type is string.
    Vid interface{}

    // Describes in user-readable terms. The type is string.
    Description interface{}

    // CLEI. The type is string.
    Clei interface{}

    // Top assembly number. The type is string.
    Tan interface{}
}

func (chassis *Diag_Racks_Rack_Chassis) GetFilter() yfilter.YFilter { return chassis.YFilter }

func (chassis *Diag_Racks_Rack_Chassis) SetFilter(yf yfilter.YFilter) { chassis.YFilter = yf }

func (chassis *Diag_Racks_Rack_Chassis) GetGoName(yname string) string {
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "description" { return "Description" }
    if yname == "clei" { return "Clei" }
    if yname == "tan" { return "Tan" }
    return ""
}

func (chassis *Diag_Racks_Rack_Chassis) GetSegmentPath() string {
    return "chassis"
}

func (chassis *Diag_Racks_Rack_Chassis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chassis *Diag_Racks_Rack_Chassis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chassis *Diag_Racks_Rack_Chassis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["serial-number"] = chassis.SerialNumber
    leafs["pid"] = chassis.Pid
    leafs["vid"] = chassis.Vid
    leafs["description"] = chassis.Description
    leafs["clei"] = chassis.Clei
    leafs["tan"] = chassis.Tan
    return leafs
}

func (chassis *Diag_Racks_Rack_Chassis) GetBundleName() string { return "cisco_ios_xr" }

func (chassis *Diag_Racks_Rack_Chassis) GetYangName() string { return "chassis" }

func (chassis *Diag_Racks_Rack_Chassis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chassis *Diag_Racks_Rack_Chassis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chassis *Diag_Racks_Rack_Chassis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chassis *Diag_Racks_Rack_Chassis) SetParent(parent types.Entity) { chassis.parent = parent }

func (chassis *Diag_Racks_Rack_Chassis) GetParent() types.Entity { return chassis.parent }

func (chassis *Diag_Racks_Rack_Chassis) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_Summary
// Summary information
type Diag_Racks_Rack_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary data. The type is slice of Diag_Racks_Rack_Summary_Summary.
    Summary []Diag_Racks_Rack_Summary_Summary
}

func (summary *Diag_Racks_Rack_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Diag_Racks_Rack_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Diag_Racks_Rack_Summary) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    return ""
}

func (summary *Diag_Racks_Rack_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Diag_Racks_Rack_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        for _, c := range summary.Summary {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Diag_Racks_Rack_Summary_Summary{}
        summary.Summary = append(summary.Summary, child)
        return &summary.Summary[len(summary.Summary)-1]
    }
    return nil
}

func (summary *Diag_Racks_Rack_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.Summary {
        children[summary.Summary[i].GetSegmentPath()] = &summary.Summary[i]
    }
    return children
}

func (summary *Diag_Racks_Rack_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summary *Diag_Racks_Rack_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Diag_Racks_Rack_Summary) GetYangName() string { return "summary" }

func (summary *Diag_Racks_Rack_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Diag_Racks_Rack_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Diag_Racks_Rack_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Diag_Racks_Rack_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Diag_Racks_Rack_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Diag_Racks_Rack_Summary) GetParentYangName() string { return "rack" }

// Diag_Racks_Rack_Summary_Summary
// Summary data
type Diag_Racks_Rack_Summary_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node type. The type is DiagNode.
    Type interface{}

    // Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Slot type. The type is DiagSlot.
    SlotType interface{}

    // Description. The type is string.
    Description interface{}
}

func (summary *Diag_Racks_Rack_Summary_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Diag_Racks_Rack_Summary_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Diag_Racks_Rack_Summary_Summary) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "node" { return "Node" }
    if yname == "slot-type" { return "SlotType" }
    if yname == "description" { return "Description" }
    return ""
}

func (summary *Diag_Racks_Rack_Summary_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Diag_Racks_Rack_Summary_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Diag_Racks_Rack_Summary_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Diag_Racks_Rack_Summary_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = summary.Type
    leafs["node"] = summary.Node
    leafs["slot-type"] = summary.SlotType
    leafs["description"] = summary.Description
    return leafs
}

func (summary *Diag_Racks_Rack_Summary_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Diag_Racks_Rack_Summary_Summary) GetYangName() string { return "summary" }

func (summary *Diag_Racks_Rack_Summary_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Diag_Racks_Rack_Summary_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Diag_Racks_Rack_Summary_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Diag_Racks_Rack_Summary_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Diag_Racks_Rack_Summary_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Diag_Racks_Rack_Summary_Summary) GetParentYangName() string { return "summary" }

